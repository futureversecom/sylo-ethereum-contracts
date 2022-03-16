import { ethers } from "hardhat";
import { Signer } from "ethers";
import { Listings } from "../typechain";
import { assert, expect } from "chai";

describe('Listing', () => {
  let accounts: Signer[];
  let owner: string;

  let listings: Listings;

  before(async () => {
    accounts = await ethers.getSigners();
    // first account is implicitly used as deployer of contracts in hardhat
    owner = await accounts[0].getAddress();
  });

  beforeEach(async () => {
    const Listings = await ethers.getContractFactory("Listings");
    listings = await Listings.deploy() as Listings;
    await listings.initialize(5000);
  });

  it('can allow owner to set default payout percentage', async () => {
    await expect(listings.setDefaultPayoutPercentage(2000))
      .to.emit(listings, 'DefaultPayoutPercentageUpdated')
      .withArgs(2000);

    const p = await listings.defaultPayoutPercentage();
    assert.equal(p, 2000, "Expected default payout percentage to be correctly updated");
  });

  it('can set listing', async () => {
    await listings.setListing("0.0.0.0/0", 1);

    const listing = await listings.getListing(owner);

    assert.equal(listing.multiAddr, "0.0.0.0/0", "Expected listings to have correct address");
    assert.equal(listing.minDelegatedStake.toNumber(), 1, "Expected listing to have correct min delegated stake");
  });

  it('requires default payout percentage to not exceed 100% when initializing', async () => {
    const Listings = await ethers.getContractFactory("Listings");
    listings = await Listings.deploy() as Listings;
    await expect(listings.initialize(10001))
      .to.be.revertedWith("The payout percentage can not exceed 100 percent");
  });

  it('requires default payout percentage to not exceed 100%', async () => {
    await expect(listings.setDefaultPayoutPercentage(10001))
      .to.be.revertedWith("The payout percentage can not exceed 100 percent");
  });

  it('requires listing to not have empty multiaddr string', async () => {
    await expect(listings.setListing("", 1))
      .to.be.revertedWith("Multiaddr string is empty");
  });
});