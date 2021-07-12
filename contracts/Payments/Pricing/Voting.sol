// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "../../Staking/Manager.sol";

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract PriceVoting is Initializable, OwnableUpgradeable {
    StakingManager _stakingManager;

    struct Vote {
        address voter;
        uint256 price;
    }

    mapping (address => uint256) public votes;

    address[] voters;

    function initialize(StakingManager stakingManager) public initializer {
        OwnableUpgradeable.__Ownable_init();
        _stakingManager = stakingManager;
    }

    function vote(uint256 price) public {
        require(_stakingManager.totalStakes(msg.sender) > 0, "Must have stake to vote");
        require(price > 0, "Voting price must be greater than 0");

        if (votes[msg.sender] == 0) {
            voters.push(msg.sender);
        }

        votes[msg.sender] = price;
    }

    // This function validates that the given list of votes is
    // sorted from lowest to highest, and each voter is included in the list,
    // and that all voters in the given list have actually voted
    function validateSortedVotes(Vote[] memory sortedVotes) view public {
        uint256 curr = 0;

        for (uint i = 0; i < sortedVotes.length; i++) {
            require(sortedVotes[i].price >= curr, "Given vote array is not sorted");

            require(votes[sortedVotes[i].voter] > 0, "Found invalid voter in sorted voter array");
        }

        // If we validate each voter in the sorted array has voted, and the length of the sorted
        // array is equal to the voter array, then all voters are accounted for
        require(sortedVotes.length == voters.length, "Not all voters were present in sorted voter array");
    }
}