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

    function sortVotes() view public returns (address[] memory, uint256[] memory) {
        Vote[] memory copy = new Vote[](voters.length);
        for (uint i = 0; i < voters.length; i++) {
            address voter = voters[i];
            copy[i] = Vote(voter, votes[voter]);
        }

        quickSort(copy, int(0), int(copy.length - 1));

        address[] memory a = new address[](copy.length);
        uint256[] memory b = new uint256[](copy.length);

        for (uint i = 0; i < copy.length; i++) {
            a[i] = copy[i].voter;
            b[i] = copy[i].price;
        }

        return (a, b);
    }

    function quickSort(Vote[] memory data, int low, int high) internal pure {
        if (low < high) {
            int pi = partition(data, low, high);

            quickSort(data, low, pi - 1);
            quickSort(data, pi + 1, high);
        }
    }

    function partition(Vote[] memory data, int low, int high) internal pure returns (int) {
        // pivot (Element to be placed at right position)
        Vote memory pivot = data[uint(high)];  
    
        int i = low - 1;

        for (int j = low; j <= high - 1; j++) {
            // If current element is smaller than the pivot
            if (data[uint(j)].price < pivot.price) {
                i++;    // increment index of smaller element

                // swap
                Vote memory copy = data[uint(j)];
                data[uint(j)] = data[uint(i)];
                data[uint(i)] = copy;
            }
        }

        // place pivot in correct position
        data[uint(high)] = data[uint(i + 1)];
        data[uint(i + 1)] = pivot;
        return i + 1;
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
    function validateSortedVotes(Vote[] memory sortedVotes) view public {
        uint256 curr = 0;

        address[] memory knownVoters = new address[](sortedVotes.length);

        for (uint i = 0; i < sortedVotes.length; i++) {
            require(sortedVotes[i].price >= curr, "Given vote array is not sorted");
            curr = sortedVotes[i].price;

            require(votes[sortedVotes[i].voter] > 0, "Found invalid voter in sorted voter array");

            // ensure there are no duplicates in the sorted array
            for (uint j = 0; j < knownVoters.length; j++) {
                require(knownVoters[j] != sortedVotes[i].voter, "Found duplicate in sorted voter array");

                // end of array
                if (knownVoters[j] == address(0)) {
                    break;
                }
            }

            knownVoters[i] = sortedVotes[i].voter;
        }

        // If we validate each voter in the sorted array has voted, and there are no
        // duplicates, and the length of the sorted array is equal to the voter array, 
        // then all voters are accounted for
        require(sortedVotes.length == voters.length, "Not all voters were present in sorted voter array");
    }
}
