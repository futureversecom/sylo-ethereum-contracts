// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "./Manager.sol";
import "../Payments/Pricing/Manager.sol";
import "../Payments/Pricing/Voting.sol";
import "../Utils.sol";

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

/*
 * The Directory contract constructs and manages a structure holding the current stakes,
 * which is queried against using the scan function. The scan function allows submitting
 * random points which will return a staked node's address in proportion to the stake it has.
*/
contract Directory is Initializable, OwnableUpgradeable {
    /* Sylo Staking Manager contract */
    StakingManager _stakingManager;

    struct DirectoryEntry {
        address stakee;
        uint256 boundary;
    }

    struct Directory {
        DirectoryEntry[] entries;

        mapping (address => uint256) stakes;

        uint256 totalStake;
    }

    bytes32 public currentDirectory;

    // Directories are indexed by the associated epoch's id
    mapping (bytes32 => Directory) directories;

    // Tracks the total stake held within a specific directory
    mapping (bytes32 => uint256) totalStakes;

    function initialize(
        StakingManager stakingManager
    ) public initializer {
        OwnableUpgradeable.__Ownable_init();
        _stakingManager = stakingManager;
    }

    /*
     * This function is called by a node as a prerequiste to participate in the next epoch.
     * This will construt the directory as nodes join. The directory is constructed
     * by creating a boundary value which is a sum of the current directory's total stake, and
     * the current stakee's total stake, and pushing the new boundary into the entries array.
     * The previous boundary and the current boundary essentially create a range, where if a
     * random point were to fall within that range, it would belong to the respective stakee.
     * The boundary value grows in size as each stakee joins, thus the directory array
     * always remains sorted. This allows us to perform a binary search on the directory.
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
    function joinDirectory(bytes32 epochId) public {
        address stakee = msg.sender;

        uint totalStake = _stakingManager.totalStakes(stakee);
        require(totalStake > 0, "Can not join directory for next epoch without any stake");

        require(
            directories[epochId].stakes[stakee] == 0,
            "Can only join the directory once per epoch"
        );

        uint nextBoundary = totalStakes[epochId] + totalStake;

        directories[epochId].entries.push(DirectoryEntry(stakee, nextBoundary));
        directories[epochId].stakes[stakee] = totalStake;
        totalStakes[epochId] += totalStake;
    }

    function scan(uint128 point) public view returns (address) {
        if (directories[currentDirectory].entries.length == 0) {
            return address(0);
        }

        // Staking all the Sylo would only be 94 bits, so multiplying this with
        // a uint128 cannot overflow a uint256.
        uint256 expectedVal = directories[currentDirectory].totalStake * uint256(point) >> 128;

        uint256 l = 0;
        uint256 r = directories[currentDirectory].entries.length - 1;

        // perform a binary search through the directory
        while (l <= r) {
            uint index = (l + r) / 2;

            uint lower = index == 0 ? 0 : directories[currentDirectory].entries[index - 1].boundary;
            uint upper = directories[currentDirectory].entries[index].boundary;

            if (expectedVal >= lower && expectedVal < upper) {
                return directories[currentDirectory].entries[index].stakee;
            } else if (expectedVal < lower) {
                r = index - 1;
            } else if (expectedVal >= upper) {
                l = index + 1;
            }
        }

        return address(0);
    }

    function getTotalStakeForStakee(bytes32 epochId, address stakee) public view returns (uint256) {
        return directories[epochId].stakes[stakee];
    }

    function getTotalStake(bytes32 epochId) public view returns (uint256) {
        return directories[epochId].totalStake;
    }
}
