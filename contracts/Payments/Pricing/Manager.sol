// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "./Voting.sol";
import "../../Staking/Manager.sol";

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract PriceManager is Initializable, OwnableUpgradeable {

    uint32 constant LOWER_QUARTILE_PERC = 25;
    uint32 constant UPPER_BOUNDARY_PERC = 90;
    uint256 constant PERC_DIVISOR = 100;

    /* Sylo Directory contract */
    StakingManager _stakingManager;

    /* Sylo Price Voting contract */
    PriceVoting _voting;

    uint256 public currentServicePrice = 0;
    uint256 public currentUpperPrice = 0;

    function initialize(StakingManager stakingManager, PriceVoting voting) public initializer {
        OwnableUpgradeable.__Ownable_init();
        _stakingManager = stakingManager;
        _voting = voting;
    }

    // The submitter of this transaction sorts the memory off chain,
    // and this contract only validates that it is sorted. This helps
    // to reduce gas cost.
    function calculatePrices(
        PriceVoting.Vote[] memory sortedVotes
    ) public onlyOwner returns (uint256 servicePrice, uint256 upperPrice) {
        _voting.validateSortedVotes(sortedVotes);

        uint256 totalStake = _stakingManager.totalStake();
        uint256 lowerBoundary = LOWER_QUARTILE_PERC * totalStake / PERC_DIVISOR;
        uint256 upperBoundary = UPPER_BOUNDARY_PERC * totalStake / PERC_DIVISOR;

        uint256 needle = 0;

        // Iterate through the votes, and accumalate each voter's stake
        // until we have crossed 25% of the total stake
        for (uint i = 0; i < sortedVotes.length; i++) {
            uint256 stake = _stakingManager.totalStakes(sortedVotes[i].voter);
            needle += stake;

            if (needle >= lowerBoundary) {
                currentServicePrice = sortedVotes[i].price;
                break;
            }
        }

        needle = totalStake;

        // Iterate through the votes starting from the end, and
        // substract each stake until we cross the upper boundary
        for (uint i = sortedVotes.length - 1; i != 0; i--) {
            uint256 stake = _stakingManager.totalStakes(sortedVotes[i].voter);
            needle -= stake;

            if (needle <= upperBoundary) {
                currentUpperPrice = sortedVotes[i].price;
            }
        }

        return (currentServicePrice, currentUpperPrice);
    }
}