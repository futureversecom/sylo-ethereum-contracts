pragma solidity ^0.6.0;
pragma experimental ABIEncoderV2;

import "truffle/Assert.sol";
import "../contracts/Listings.sol";


contract ListingsTest {
  Listings listings;

  function beforeEach() public {
    listings = new Listings();
  }

  function testSettingListing() public {

    Listings.Listing memory listing = Listings.Listing("/ip4/127.0.0.1/udp/1234");

    listings.setListing(listing);

    Listings.Listing memory retrieved  = listings.getListing(address(this));

    Assert.equal(listing.multiAddr, retrieved.multiAddr, "Expected listing to be set");
  }
}