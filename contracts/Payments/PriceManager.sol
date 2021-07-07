// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "../Directory.sol";

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract PriceManager is Initializable, OwnableUpgradeable {

    uint32 constant LOWER_QUARTILE_PERC = 25;
    uint256 constant PERC_DIVISOR = 100;

    /* Sylo Directory contract */
    Directory _directory;

    uint256 public currentPrice = 0;

    struct Vote {
        address voter;
        uint256 price;
    }

    mapping (address => uint256) votes;

    address[] voters;

    function initialize(Directory directory) public initializer {
        OwnableUpgradeable.__Ownable_init();
        _directory = directory;
    }

    function vote(uint256 price) public {
        require(_directory.stakees(msg.sender) > 0, "Must have stake to vote");
        require(price > 0, "Voting price must be greater than 0");

        if (votes[msg.sender] == 0) {
            voters.push(msg.sender);
        }

        votes[msg.sender] = price;
    }

    // The submitter of this transaction sorts the memory off chain,
    // and this contract only validate that it is sorted. This helps
    // to reduce gas cost.
    function calculatePrice(Vote[] memory sortedVotes) public onlyOwner returns (uint256 price) {
        validateSortedVotes(sortedVotes);

        uint256 totalStake = _directory.getTotalStake();
        uint256 boundary = LOWER_QUARTILE_PERC * totalStake / PERC_DIVISOR;
        uint256 iteratedStake = 0;

        // Iterate through the votes, and accumalate each voter's stake
        // until we have crossed 25% of the total stake
        for (uint i = 0; i < sortedVotes.length; i++) {
            uint256 stake = _directory.stakees(sortedVotes[i].voter);
            iteratedStake += stake;

            if (iteratedStake >= boundary) {
                currentPrice = sortedVotes[i].price;
                return currentPrice;
            }
        }
    }

    // This function validates that the given list of votes is
    // sorted from lowest to highest, and each voter is included in the list,
    // and that all voters in the given list have actually voted
    function validateSortedVotes(Vote[] memory sortedVotes) internal {
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