// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract Listings is Initializable {

    struct Listing {
        string multiAddr; // MultiAddr to connect to the account
        // TODO store tags
    }

    mapping(address => Listing) listings;

    function initialize() public initializer {
    }

    function setListing(Listing memory listing) public {
        // TODO validate listing?
        listings[msg.sender] = listing;
    }

    function getListing(address account) public view returns (Listing memory) {
        return listings[account];
    }
}
