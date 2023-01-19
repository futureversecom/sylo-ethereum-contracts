import { ethers } from 'hardhat';
import { Signer } from 'ethers';
import { Registries, TestSeekers } from '../typechain';
import { assert, expect } from 'chai';
import utils from './utils';

describe('Registries', () => {
  let accounts: Signer[];
  let owner: string;

  let registries: Registries;
  let seekers: TestSeekers;

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
    await registries.register('http://api');

    const registry = await registries.getRegistry(owner);

    assert.equal(
      registry.publicEndpoint,
      'http://api',
      'Expected registries to have correct address',
    );
  });

  it('can retrieve all registered nodes', async () => {
    await registries.register('http://api');
    await registries.connect(accounts[1]).register('http://api');

    const nodes = await registries.getNodes();

    assert.deepEqual(nodes, [
      await accounts[0].getAddress(),
      await accounts[1].getAddress(),
    ]);
  });

  it('can query total number of registered nodes', async () => {
    await registries.register('http://api');
    await registries.connect(accounts[1]).register('http://api');

    const n = await registries.getTotalNodes();

    assert.equal(n.toNumber(), 2);
  });

  it('can retrieve a list of registries', async () => {
    const addresses = await Promise.all(accounts.map(a => a.getAddress()));

    for (let i = 0; i < 20; i++) {
      await registries.connect(accounts[i]).register(`http://api/${i}`);
    }

    const result = await registries.getRegistries(0, 20);

    assert.deepEqual(result[0], addresses, 'Expected 20 registries returned');

    for (let i = 0; i < 20; i++) {
      assert.equal(
        result[0][i],
        addresses[i],
        'Expected correct registry to be returned',
      );
      assert.equal(
        result[1][i].publicEndpoint,
        `http://api/${i}`,
        'Expected correct registry to be returned',
      );
    }
  });

  it('can retrieve a list of registries with start and end indexes', async () => {
    const addresses = await Promise.all(accounts.map(a => a.getAddress()));

    for (let i = 0; i < 20; i++) {
      await registries.connect(accounts[i]).register(`http://api/${i}`);
    }

    const result = await registries.getRegistries(5, 10);

    assert.deepEqual(
      result[0],
      addresses.slice(5, 10),
      'Expected only accounts 5 to 9 to be returned from query',
    );

    for (let i = 5; i < 10; i++) {
      assert.equal(
        result[0][i - 5],
        addresses[i],
        'Expected correct registry to be returned',
      );
      assert.equal(
        result[1][i - 5].publicEndpoint,
        `http://api/${i}`,
        'Expected correct registry to be returned',
      );
    }

    await expect(registries.getRegistries(8, 5)).to.be.revertedWith(
      'End index must be greater than start index',
    );

    await expect(registries.getRegistries(8, 21)).to.be.revertedWith(
      'End index cannot be greater than total number of registered nodes',
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

    const tokenId = 1;
    await seekers.mint(seekerAddress, tokenId);

    const block = await ethers.provider.getBlockNumber();

    const hash = ethers.utils.solidityKeccak256(
      ['string', 'uint256', 'address', 'uint256'],
      [
        await registries.getProofMessage(tokenId, seekerAddress, block),
        tokenId,
        owner,
        block,
      ],
    );
    const proofMessage = ethers.utils.arrayify(hash);
    const proof = await seekerAccount.signMessage(proofMessage);

    await expect(
      registries.setSeekerAccount(seekerAddress, tokenId, block + 1000, proof),
    ).to.be.revertedWith('Proof can not be set for a future block');

    await utils.advanceBlock(200);

    await expect(
      registries.setSeekerAccount(seekerAddress, tokenId, block, proof),
    ).to.be.revertedWith('Proof is expired');
  });

  it('fails to set seeker account with invalid proof', async () => {
    const seekerAccount = accounts[1];
    const seekerAddress = await seekerAccount.getAddress();

    const tokenId = 1;
    await seekers.mint(seekerAddress, tokenId);

    const block = await ethers.provider.getBlockNumber();

    const hash = ethers.utils.solidityKeccak256(
      ['string', 'uint256', 'address', 'uint256'],
      [
        await registries.getProofMessage(tokenId, seekerAddress, block),
        tokenId,
        owner,
        block,
      ],
    );
    const proofMessage = ethers.utils.arrayify(hash);

    // sign proof with wrong account
    const proof = await seekerAccount.signMessage(proofMessage);

    await expect(
      registries.setSeekerAccount(seekerAddress, tokenId, block, proof),
    ).to.be.revertedWith('Proof must be signed by specified seeker account');
  });

  it("fails to set seeker account if seeker isn't owned by account", async () => {
    const seekerAccount = accounts[1];
    const seekerAddress = await seekerAccount.getAddress();

    const tokenId = 1;
    await seekers.mint(await accounts[2].getAddress(), tokenId);

    const block = await ethers.provider.getBlockNumber();

    const accountAddress = await accounts[0].getAddress();
    const proofMessage = await registries.getProofMessage(
      tokenId,
      accountAddress,
      block,
    );

    const signature = await seekerAccount.signMessage(proofMessage);

    await expect(
      registries.setSeekerAccount(seekerAddress, tokenId, block, signature),
    ).to.be.revertedWith('Seeker account must own the specified seeker');
  });

  it('can revoke seeker account', async () => {
    const seekerAccount = accounts[1];

    await utils.setSeekerRegistry(
      registries,
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
    await expect(registries.register('')).to.be.revertedWith(
      'Public endpoint can not be empty',
    );
  });

  it('registered node seeker id will be reset if a new node registers with the same seeker id', async () => {
    const seekerAccount = accounts[2];

    const accountOne = accounts[3];
    const accountAddressOne = await accountOne.getAddress();

    const accountTwo = accounts[4];
    const accountAddressTwo = await accountTwo.getAddress();

    const tokenID = 100;

    await utils.setSeekerRegistry(
      registries,
      seekers,
      accountOne,
      seekerAccount,
      tokenID,
    );

    await utils.setSeekerRegistry(
      registries,
      seekers,
      accountTwo,
      seekerAccount,
      tokenID,
    );

    const regoSeekerAccountOne = await registries.getRegistry(
      accountAddressOne,
    );
    const regoSeekerAccountTwo = await registries.getRegistry(
      accountAddressTwo,
    );

    // tests that the registry for both seeker accounts don't have the same seekerID
    expect(regoSeekerAccountOne.seekerId).to.equal(0);
    expect(regoSeekerAccountTwo.seekerId).is.equal(tokenID);
  });

  it('Has the correct prefix message', async () => {
    const account = accounts[0];

    const seekerAccount = accounts[1];

    const tokenId = 100; // Seeker ID

    const block = await ethers.provider.getBlockNumber();

    await utils.setSeekerRegistry(
      registries,
      seekers,
      account,
      seekerAccount,
      tokenId,
    );

    const prefix = await registries.getProofMessage(
      tokenId,
      await account.getAddress(),
      block,
    );
    console.log(prefix);
  });
});
