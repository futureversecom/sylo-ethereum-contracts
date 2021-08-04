// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "./Voting.sol";
import "../../Utils.sol";
import "../../Staking/Manager.sol";

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract PriceManager is Initializable, OwnableUpgradeable {

    uint32 constant LOWER_QUARTILE_PERC = 25;
    uint32 constant UPPER_BOUNDARY_PERC = 90;

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

    function calculatePrices(
        uint256[] memory sortedIndexes
    ) public onlyOwner returns (uint256 servicePrice, uint256 upperPrice) {
        // we were given a list of sorted votes, validate the votes instead
        PriceVoting.Vote[] memory sortedVotes = _voting.validateSortedVotes(sortedIndexes);
        return _calculatePrices(sortedVotes);
    }

    function _calculatePrices(
        PriceVoting.Vote[] memory sortedVotes
    ) internal onlyOwner returns (uint256 servicePrice, uint256 upperPrice) {
        uint256 totalStake = _stakingManager.totalStake();

        uint256 lowerBoundary = SyloUtils.percOf(totalStake, LOWER_QUARTILE_PERC);
        uint256 upperBoundary = SyloUtils.percOf(totalStake, UPPER_BOUNDARY_PERC);

        uint256 needle = 0;

        // Iterate through the votes, and accumalate each voter's stake
        // until we have crossed 25% of the total stake
        for (uint i = 0; i < sortedVotes.length; i++) {
            uint256 stake = _stakingManager.totalStakes(sortedVotes[i].voter);

            // Only consider nodes with a valid stake
            if (stake == 0) {
                continue;
            }

            // Only consider nodes with a valid vote
            if (sortedVotes[i].price == 0) {
                continue;
            }

            needle += stake;

            if (needle >= lowerBoundary) {
                currentServicePrice = sortedVotes[i].price;
                break;
            }
        }

        needle = totalStake;

        // Iterate through the votes starting from the end, and
        // substract each stake until we cross the upper boundary
        for (uint i = sortedVotes.length; i > 0; i--) {
            uint256 stake = _stakingManager.totalStakes(sortedVotes[i - 1].voter);

            // Only consider nodes with a valid stake
            if (stake == 0) {
                continue;
            }

            // Only consider nodes with a valid vote
            if (sortedVotes[i - 1].price == 0) {
                continue;
            }

            needle -= stake;

            if (needle <= upperBoundary) {
                currentUpperPrice = sortedVotes[i - 1].price;
            }
        }

        return (currentServicePrice, currentUpperPrice);
    }
}
