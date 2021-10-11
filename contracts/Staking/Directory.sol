// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "./Manager.sol";
import "../Payments/Ticketing/RewardsManager.sol";
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

    /* Sylo Rewards Manager contract */
    RewardsManager _rewardsManager;

    struct DirectoryEntry {
        address stakee;
        uint256 boundary;
    }

    struct Directory {
        DirectoryEntry[] entries;

        mapping (address => uint256) stakes;

        uint256 totalStake;
    }

    uint256 public currentDirectory;

    // Directories are indexed by the associated epoch's id
    mapping (uint256 => Directory) directories;

    function initialize(
        StakingManager stakingManager,
        RewardsManager rewardsManager
    ) public initializer {
        OwnableUpgradeable.__Ownable_init();
        _stakingManager = stakingManager;
        _rewardsManager = rewardsManager;
    }

    function setCurrentDirectory(uint256 epochId) public onlyOwner {
        currentDirectory = epochId;
    }

    /*
     * This function is called by a node as a prerequiste to participate in the next epoch.
     * This will construct the directory as nodes join. The directory is constructed
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
    function joinNextDirectory() public {
        address stakee = msg.sender;

        uint256 managedStake = _stakingManager.getStakeeTotalManagedStake(stakee);
        uint256 stakeReward = _rewardsManager.unclaimedStakeRewards(stakee);
        uint256 totalStake = managedStake + stakeReward;
        require(totalStake > 0, "Can not join directory for next epoch without any stake");

        uint256 epochId = currentDirectory + 1;

        require(
            directories[epochId].stakes[stakee] == 0,
            "Can only join the directory once per epoch"
        );

        uint256 nextBoundary = directories[epochId].totalStake + totalStake;

        directories[epochId].entries.push(DirectoryEntry(stakee, nextBoundary));
        directories[epochId].stakes[stakee] = totalStake;
        directories[epochId].totalStake = nextBoundary;
    }

    function scan(uint128 point) public view returns (address stakee) {
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
    }

    function getTotalStakeForStakee(uint256 epochId, address stakee) public view returns (uint256) {
        return directories[epochId].stakes[stakee];
    }

    function getTotalStake(uint256 epochId) public view returns (uint256) {
        return directories[epochId].totalStake;
    }

    function getEntries(uint256 epochId) public view returns (address[] memory, uint256[] memory) {
        address[] memory stakees = new address[](directories[epochId].entries.length);
        uint256[] memory boundaries = new uint256[](directories[epochId].entries.length);
        for (uint i = 0; i < directories[epochId].entries.length; i++) {
            DirectoryEntry memory entry = directories[epochId].entries[i];
            stakees[i] = entry.stakee;
            boundaries[i] = entry.boundary;
        }
        return (stakees, boundaries);
    }
}
