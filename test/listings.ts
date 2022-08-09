import { ethers } from 'hardhat';
import { Signer } from 'ethers';
import { Listings, MockOracle, Seekers } from '../typechain';
import { assert, expect } from 'chai';
import utils from './utils';

describe('Listing', () => {
  let accounts: Signer[];
  let owner: string;

  let listings: Listings;
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
    listings = contracts.listings;
    mockOracle = contracts.mockOracle;
    seekers = contracts.seekers;
  });

  it('requires default payout percentage to not exceed 100% when initializing', async () => {
    const Listings = await ethers.getContractFactory('Listings');
    listings = await Listings.deploy();
    await expect(
      listings.initialize(seekers.address, 10001, 100),
    ).to.be.revertedWith('The payout percentage can not exceed 100 percent');
  });

  it('can allow owner to set default payout percentage', async () => {
    await expect(listings.setDefaultPayoutPercentage(2000))
      .to.emit(listings, 'DefaultPayoutPercentageUpdated')
      .withArgs(2000);

    const p = await listings.defaultPayoutPercentage();
    assert.equal(
      p,
      2000,
      'Expected default payout percentage to be correctly updated',
    );
  });

  it('can set listing', async () => {
    await listings.setListing('http://api', 1);

    const listing = await listings.getListing(owner);

    assert.equal(
      listing.publicEndpoint,
      'http://api',
      'Expected listings to have correct address',
    );
    assert.equal(
      listing.minDelegatedStake.toNumber(),
      1,
      'Expected listing to have correct min delegated stake',
    );
  });

  it('requires default payout percentage to not exceed 100%', async () => {
    await expect(listings.setDefaultPayoutPercentage(10001)).to.be.revertedWith(
      'The payout percentage can not exceed 100 percent',
    );
  });

  it('can set seeker account with valid proof', async () => {
    const seekerAccount = accounts[1];
    const seekerAddress = await seekerAccount.getAddress();

    await utils.setSeekerListing(
      listings,
      mockOracle,
      seekers,
      accounts[0],
      accounts[1],
      1,
    );

    const listing = await listings.getListing(owner);

    expect(listing.seekerAccount).to.equal(seekerAddress);
    expect(listing.seekerId).to.equal(1);
  });

  it('fails to set seeker account with invalid blocks for proof', async () => {
    const seekerAccount = accounts[1];
    const seekerAddress = await seekerAccount.getAddress();

    await utils.setSeekerOwnership(mockOracle, seekers, 1, seekerAddress);

    const block = await ethers.provider.getBlockNumber();

    const hash = ethers.utils.solidityKeccak256(
      ['string', 'uint256', 'address', 'uint256'],
      [await listings.getPrefix(), 1, owner, block],
    );
    const proofMessage = ethers.utils.arrayify(hash);
    const proof = await seekerAccount.signMessage(proofMessage);

    await expect(
      listings.setSeekerAccount(seekerAddress, 1, 1000, proof),
    ).to.be.revertedWith('Proof can not be set for a future block');

    await utils.advanceBlock(200);

    await expect(
      listings.setSeekerAccount(seekerAddress, 1, 10, proof),
    ).to.be.revertedWith('Proof is expired');
  });

  it('fails to set seeker account with invalid proof', async () => {
    const seekerAccount = accounts[1];
    const seekerAddress = await seekerAccount.getAddress();

    await utils.setSeekerOwnership(mockOracle, seekers, 1, seekerAddress);

    const block = await ethers.provider.getBlockNumber();

    const hash = ethers.utils.solidityKeccak256(
      ['string', 'uint256', 'address', 'uint256'],
      [await listings.getPrefix(), 1, owner, block],
    );
    const proofMessage = ethers.utils.arrayify(hash);

    // sign proof with wrong account
    const proof = await accounts[0].signMessage(proofMessage);

    await expect(
      listings.setSeekerAccount(seekerAddress, 1, block, proof),
    ).to.be.revertedWith('Proof must be signed by specified seeker account');
  });

  it("fails to set seeker account if ownership hasn't been verified", async () => {
    const seekerAccount = accounts[1];
    const seekerAddress = await seekerAccount.getAddress();

    const block = await ethers.provider.getBlockNumber();

    const prefix = await listings.getPrefix();
    const accountAddress = await accounts[0].getAddress();
    const proofMessage = `${prefix}:${1}:${accountAddress.toLowerCase()}:${block.toString()}`;

    const signature = await seekerAccount.signMessage(proofMessage);

    await expect(
      listings.setSeekerAccount(seekerAddress, 1, block, signature),
    ).to.be.revertedWith('Seeker account must own the specified seeker');
  });

  it('can revoke seeker account', async () => {
    const seekerAccount = accounts[1];

    await utils.setSeekerListing(
      listings,
      mockOracle,
      seekers,
      accounts[0],
      accounts[1],
      1,
    );

    await listings.connect(seekerAccount).revokeSeekerAccount(owner);

    const listing = await listings.getListing(owner);

    expect(listing.seekerAccount).to.equal(ethers.constants.AddressZero);
  });

  it('can only revoke seeker account as seeker account', async () => {
    await utils.setSeekerListing(
      listings,
      mockOracle,
      seekers,
      accounts[0],
      accounts[1],
      1,
    );

    await expect(listings.revokeSeekerAccount(owner)).to.be.revertedWith(
      'Seeker account and msg.sender must be equal',
    );
  });

  it('requires listing to not have empty public endpoint string', async () => {
    await expect(listings.setListing('', 1)).to.be.revertedWith(
      'Public endpoint can not be empty',
    );
  });
});
