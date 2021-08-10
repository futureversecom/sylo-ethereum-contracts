// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "truffle/Assert.sol";
import "../contracts/Listings.sol";


contract ListingsTest {
  Listings listings;

  function beforeEach() public {
    listings = new Listings(); 
    listings.initialize(5000);
  }

  function testSettingListing() public {
    string memory multiAddr = "/ip4/127.0.0.1/udp/1234";
    listings.setListing(multiAddr, 1);

    Listings.Listing memory retrieved = listings.getListing(address(this));

    Assert.equal(multiAddr, retrieved.multiAddr, "Expected listing to be set");
  }
}
