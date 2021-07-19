// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "./Manager.sol";
import "../Payments/Pricing/Manager.sol";
import "../Payments/Pricing/Voting.sol";

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

/*
 * The Directory contract constructs and manages a structure holding the current stakes,
 * which is queried against using the scan function. The scan function allows submitting 
 * random points which will return a staked node's address in proportion to the stake it has.
*/
contract Directory is Initializable, OwnableUpgradeable {

    uint256 constant PERC_DIVISOR = 100;

    // Nodes are excluded if their voted price exceeds service price + 10%
    uint256 constant PRICE_THRESHOLD = 110;

    StakingManager _stakingManager;

    PriceManager _priceManager;

    PriceVoting _priceVoting;

    struct DirectoryEntry {
        address stakee;
        uint256 boundary;
    }

    DirectoryEntry[] public currentDirectory;

    function initialize(
        PriceVoting priceVoting,
        PriceManager priceManager,
        StakingManager stakingManager
    ) public initializer {
        OwnableUpgradeable.__Ownable_init();
        _priceVoting = priceVoting;
        _priceManager = priceManager;
        _stakingManager = stakingManager;
    }

    /*
     * We construct the directory by iterating through each valid stakee, and
     * creating a boundary value which is a sum of the previously iterated stakee's
     * boundary value, and the current stakee's total stake. The previous boundary and
     * the current boundary essentially create a range, where if a random point were to 
     * fall within that range, it would belong to the current stakee. The boundary value
     * grows in size as each stakee is iterated, thus the final directory array
     * is sorted. This allows us to perform a binary search on the directory.
     *
     * Example
     *
     * Stakes: [ Alice/20, Bob/10, Carl/40, Dave/25 ]
     * TotalStake: 95
     *
     * Directory:
     *
     *  |-----------|------|----------------|--------|
     *     Alice/20  Bob/30     Carl/70      Dave/95
     */
    function constructDirectory() public onlyOwner {
        delete currentDirectory;

        uint lowerBoundary = 0;

        for (uint i = 0; i < _stakingManager.getCountOfStakees(); i++) {
            address stakee = _stakingManager.stakees(i);
            uint totalStake = _stakingManager.totalStakes(stakee);

            // Only add stakee to the directory after passing
            // some validation

            if (totalStake < 1) {
                continue;
            }

            uint votedPrice = _priceVoting.votes(stakee);

            // absent vote
            if (votedPrice == 0) {
                continue;
            }

            // Voted with price aboove 90th percentile and 
            // above service price + 10%
            if (votedPrice > _priceManager.currentUpperPrice() &&
                votedPrice > (_priceManager.currentServicePrice() * PRICE_THRESHOLD / PERC_DIVISOR)) {
                continue;
            }

            currentDirectory.push(DirectoryEntry(stakee, lowerBoundary + totalStake));

            lowerBoundary += totalStake;
        }
    }

    function scan(uint128 point) public view returns (address) {
        if (currentDirectory.length == 0) {
            return address(0);
        }

        uint256 totalStake = _stakingManager.getTotalStake();
        // Staking all the Sylo would only be 94 bits, so multiplying this with
        // a uint128 cannot overflow a uint256.
        uint256 expectedVal = totalStake * uint256(point) >> 128;

        uint256 l = 0;
        uint256 r = currentDirectory.length - 1;

        // perform a binary search through the directory
        while (l <= r) {
            uint index = (l + r) / 2;

            uint lower = index == 0 ? 0 : currentDirectory[index - 1].boundary;
            uint upper = currentDirectory[index].boundary;

            if (expectedVal >= lower && expectedVal < upper) {
                return currentDirectory[index].stakee;
            } else if (expectedVal < lower) {
                r = index - 1;
            } else if (expectedVal >= upper) {
                l = index + 1;
            }
        }

        return address(0);
    }
}
