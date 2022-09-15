import { ethers } from 'hardhat';
import { Signer } from 'ethers';
import { Registries, MockOracle, Seekers } from '../typechain';
import { assert, expect } from 'chai';
import utils from './utils';

describe('Registries', () => {
  let accounts: Signer[];
  let owner: string;

  let registries: Registries;
  let mockOracle: MockOracle;
  let seekers: Seekers;

  before(async () => {
    accounts = await ethers.getSigners();
    // first account is implicitly used as deployer of contracts in hardhat
    owner = await accounts[0].getAddress();
  });

  beforeEach(async () => {
    const Token = await ethers.getContractFactory('SyloToken');
    const token = await Token.deploy();

    const contracts = await utils.initializeContracts(owner, token.address, {
      payoutPercentage: 5000,
    });
    registries = contracts.registries;
    mockOracle = contracts.mockOracle;
    seekers = contracts.seekers;
  });

  it('requires default payout percentage to not exceed 100% when initializing', async () => {
    const Registries = await ethers.getContractFactory('Registries');
    registries = await Registries.deploy();
    await expect(
      registries.initialize(seekers.address, 10001, 100),
    ).to.be.revertedWith('The payout percentage can not exceed 100 percent');
  });

  it('can allow owner to set default payout percentage', async () => {
    await expect(registries.setDefaultPayoutPercentage(2000))
      .to.emit(registries, 'DefaultPayoutPercentageUpdated')
      .withArgs(2000);

    const p = await registries.defaultPayoutPercentage();
    assert.equal(
      p,
      2000,
      'Expected default payout percentage to be correctly updated',
    );
  });

  it('can set registry', async () => {
    await registries.register('http://api', 1);

    const registry = await registries.getRegistry(owner);

    assert.equal(
      registry.publicEndpoint,
      'http://api',
      'Expected registries to have correct address',
    );
    assert.equal(
      registry.minDelegatedStake.toNumber(),
      1,
      'Expected registry to have correct min delegated stake',
    );
  });

  it('requires default payout percentage to not exceed 100%', async () => {
    await expect(
      registries.setDefaultPayoutPercentage(10001),
    ).to.be.revertedWith('The payout percentage can not exceed 100 percent');
  });

  it('can set seeker account with valid proof', async () => {
    const seekerAccount = accounts[1];
    const seekerAddress = await seekerAccount.getAddress();

    await utils.setSeekerRegistry(
      registries,
      mockOracle,
      seekers,
      accounts[0],
      accounts[1],
      1,
    );

    const registry = await registries.getRegistry(owner);

    expect(registry.seekerAccount).to.equal(seekerAddress);
    expect(registry.seekerId).to.equal(1);
  });

  it('fails to set seeker account with invalid blocks for proof', async () => {
    const seekerAccount = accounts[1];
    const seekerAddress = await seekerAccount.getAddress();

    await utils.setSeekerOwnership(mockOracle, seekers, 1, seekerAddress);

    const block = await ethers.provider.getBlockNumber();

    const hash = ethers.utils.solidityKeccak256(
      ['string', 'uint256', 'address', 'uint256'],
      [await registries.getPrefix(), 1, owner, block],
    );
    const proofMessage = ethers.utils.arrayify(hash);
    const proof = await seekerAccount.signMessage(proofMessage);

    await expect(
      registries.setSeekerAccount(seekerAddress, 1, block + 1000, proof),
    ).to.be.revertedWith('Proof can not be set for a future block');

    await utils.advanceBlock(200);

    await expect(
      registries.setSeekerAccount(seekerAddress, 1, block, proof),
    ).to.be.revertedWith('Proof is expired');
  });

  it('fails to set seeker account with invalid proof', async () => {
    const seekerAccount = accounts[1];
    const seekerAddress = await seekerAccount.getAddress();

    await utils.setSeekerOwnership(mockOracle, seekers, 1, seekerAddress);

    const block = await ethers.provider.getBlockNumber();

    const hash = ethers.utils.solidityKeccak256(
      ['string', 'uint256', 'address', 'uint256'],
      [await registries.getPrefix(), 1, owner, block],
    );
    const proofMessage = ethers.utils.arrayify(hash);

    // sign proof with wrong account
    const proof = await accounts[0].signMessage(proofMessage);

    await expect(
      registries.setSeekerAccount(seekerAddress, 1, block, proof),
    ).to.be.revertedWith('Proof must be signed by specified seeker account');
  });

  it("fails to set seeker account if ownership hasn't been verified", async () => {
    const seekerAccount = accounts[1];
    const seekerAddress = await seekerAccount.getAddress();

    const block = await ethers.provider.getBlockNumber();

    const prefix = await registries.getPrefix();
    const accountAddress = await accounts[0].getAddress();
    const proofMessage = `${prefix}:${1}:${accountAddress.toLowerCase()}:${block.toString()}`;

    const signature = await seekerAccount.signMessage(proofMessage);

    await expect(
      registries.setSeekerAccount(seekerAddress, 1, block, signature),
    ).to.be.revertedWith('Seeker account must own the specified seeker');
  });

  it('can revoke seeker account', async () => {
    const seekerAccount = accounts[1];

    await utils.setSeekerRegistry(
      registries,
      mockOracle,
      seekers,
      accounts[0],
      accounts[1],
      1,
    );

    await registries.connect(seekerAccount).revokeSeekerAccount(owner);

    const registry = await registries.getRegistry(owner);

    expect(registry.seekerAccount).to.equal(ethers.constants.AddressZero);
  });

  it('can only revoke seeker account as seeker account', async () => {
    await utils.setSeekerRegistry(
      registries,
      mockOracle,
      seekers,
      accounts[0],
      accounts[1],
      1,
    );

    await expect(registries.revokeSeekerAccount(owner)).to.be.revertedWith(
      'Seeker account and msg.sender must be equal',
    );
  });

  it('requires registry to not have empty public endpoint string', async () => {
    await expect(registries.register('', 1)).to.be.revertedWith(
      'Public endpoint can not be empty',
    );
  });
});
