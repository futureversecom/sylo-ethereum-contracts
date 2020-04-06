pragma solidity ^0.6.0;
pragma experimental ABIEncoderV2;


/*
 * Stores details on the accounts listed in the directory
*/
contract Listings {

  struct Listing {
    bytes32 syloId; // SyloId to verify identity once connected
    string multiAddr; // MultiAddr to connect to the account
    // TODO store tags
  }

  mapping(address => Listing) listings;

  function setListing(Listing memory listing) public {
    // TODO validate listing?
    listings[msg.sender] = listing;
  }

  function getListing(address account) public view returns (Listing memory) {
    return listings[account];
  }
}
