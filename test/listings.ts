import { ethers } from "hardhat";
import { Signer } from "ethers";
import { Listings } from "../typechain";
import { assert } from "chai";

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
    await listings.setDefaultPayoutPercentage(2000);

    const p = await listings.defaultPayoutPercentage();
    assert.equal(p, 2000, "Expected default payout percentage to be correctly updated");
  });

  it('can set listing', async () => {
    await listings.setListing("0.0.0.0/0", 1);

    const listing = await listings.getListing(owner);

    assert.equal(listing.multiAddr, "0.0.0.0/0", "Expected listings to have correct address");
    assert.equal(listing.minDelegatedStake.toNumber(), 1, "Expected listing to have correct min delegated stake");
  });

});