// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract Listings is Initializable, OwnableUpgradeable {

    struct Listing {
        // MultiAddr to connect to the account
        string multiAddr;

        // Percentage of a tickets value that will be rewarded to
        // delagated stakers expressed as a fraction of 10000.
        // This value is currently locked to the default payout percentage
        // until epochs are implemented.
        uint16 payoutPercentage;

        // The minimum amount of stake that is required to
        // add a delegated stake against this node
        uint256 minDelegatedStake;

        // Explicit property to check if an instance of this struct actually exists
        bool initialized;
    }

    mapping(address => Listing) public listings;

    uint16 public defaultPayoutPercentage;

    function initialize(uint16 _defaultPayoutPercentage) public initializer {
        OwnableUpgradeable.__Ownable_init();
        setDefaultPayoutPercentage(_defaultPayoutPercentage);
    }

    function setDefaultPayoutPercentage(uint16 _defaultPayoutPercentage) public onlyOwner {
        require(
            _defaultPayoutPercentage <= 10000,
            "The payout percentage can not exceed 100 percent"
        );
        defaultPayoutPercentage = _defaultPayoutPercentage;
    }

    function setListing(string memory multiAddr, uint256 minDelegatedStake) public {
        require(bytes(multiAddr).length != 0, "Multiaddr string is empty");

        // TODO Remove defaultPayoutPercentage once epochs are introduced
        Listing memory listing = Listing(multiAddr, defaultPayoutPercentage, minDelegatedStake, true);
        listings[msg.sender] = listing;
    }

    function getListing(address account) public view returns (Listing memory) {
        return listings[account];
    }
}
