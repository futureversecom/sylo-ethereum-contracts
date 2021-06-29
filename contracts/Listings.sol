// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract Listings is Initializable, OwnableUpgradeable {

    struct Listing {
        // MultiAddr to connect to the account
        string multiAddr; 

        // Payout percentage to delegated stakes for winning
        // This value is currently locked to the default payout
        // percentage until epochs are implemented
        uint8 payoutPercentage;

        // The minimum amount of stake that is required to
        // add a delegated stake against this node
        uint256 minDelegatedStake;

        // Explicit property to check if an instance of this struct actually exists
        bool initialized;
    }

    mapping(address => Listing) listings;

    uint8 public defaultPayoutPercentage;

    function initialize(uint8 _defaultPayoutPercentage) public initializer {
        OwnableUpgradeable.__Ownable_init();
        setDefaultPayoutPercentage(_defaultPayoutPercentage);
    }

    function setDefaultPayoutPercentage(uint8 _defaultPayoutPercentage) public onlyOwner {
        require(
            _defaultPayoutPercentage >= 0 && _defaultPayoutPercentage <= 100,
            "The payout percentage must be a value between 0 and 100"
        );
        defaultPayoutPercentage = _defaultPayoutPercentage;
    }

    function setListing(string memory multiAddr, uint256 minDelegatedStake) public {
        // TODO validate listing?
        require(bytes(multiAddr).length != 0, "Multiaddr string is empty");

        // TODO Remove defaultPayoutPercentage once epochs are introduced
        Listing memory listing = Listing(multiAddr, defaultPayoutPercentage, minDelegatedStake, true);
        listings[msg.sender] = listing;
    }

    function getListing(address account) public view returns (Listing memory) {
        return listings[account];
    }
}
