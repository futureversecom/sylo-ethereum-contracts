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

    // This array is used to iterate through all voters
    address[] public voters;

    function initialize(StakingManager stakingManager) public initializer {
        OwnableUpgradeable.__Ownable_init();
        _stakingManager = stakingManager;
    }

    function vote(uint256 price) public {
        require(price > 0, "Voting price must be greater than 0");

        if (votes[msg.sender] == 0) {
            voters.push(msg.sender);
        }

        votes[msg.sender] = price;
    }

    function withdraw() public {
        votes[msg.sender] = 0;
        for (uint i = 0; i < voters.length; i++) {
            if (voters[i] == msg.sender) {
                voters[i] = voters[voters.length - 1];
                voters.pop();
            }
        }
    }

    function getVotes() view public returns (address[] memory, uint256[] memory) {
        address[] memory _voters = new address[](voters.length);
        uint256[] memory _votes = new uint256[](voters.length);

        for (uint i = 0; i < voters.length; i++) {
            _voters[i] = voters[i];
            _votes[i] = votes[voters[i]];
        }

        return (_voters, _votes);
    }

    // This function validates that the given list of votes is
    // sorted from lowest to highest, and each voter is included in the list,
    // and that all voters in the given list have actually voted
    function validateSortedVotes(uint256[] memory sortedIndexes) public view returns (Vote[] memory) {
        // If we validate there are no duplicates, and the length of the sorted 
        // array is equal to the voter array, then all voters are accounted for
        require(sortedIndexes.length == voters.length, "Not all voters were present in sorted voter array");

        // used to validate there are no duplicate votes
        bool[] memory seen = new bool[](voters.length);

        // used to validate that votes are in order
        uint256 curr = 0;

        // sorted voter list constructed from querying sorted indexes
        Vote[] memory sortedVotes = new Vote[](voters.length);

        for (uint i = 0; i < sortedIndexes.length; i++) {
            address voter = voters[sortedIndexes[i]];
            uint256 price = votes[voter];

            require(price >= curr, "Given vote array is not sorted");
            curr = votes[voter];

            require(!seen[sortedIndexes[i]], "Found duplicate");
            seen[sortedIndexes[i]] = true;

            sortedVotes[i] = Vote(voter, price);
        }

        return sortedVotes;
    }
}
