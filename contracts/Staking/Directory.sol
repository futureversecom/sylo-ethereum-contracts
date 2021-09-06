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

    // Nodes are excluded if their voted price exceeds service price + 10%
    uint16 constant PRICE_THRESHOLD = 11000;

    StakingManager _stakingManager;

    struct DirectoryEntry {
        address stakee;
        uint256 boundary;
    }

    struct Stake {
        address staker;
        uint256 amount;
    }

    struct Directory {
        DirectoryEntry[] entries;

        // We also persist all of the stakes associated to a particular stakee for
        // this directory iteration. This record is used to appropriately divy out rewards
        // based on stake proportion at the end of an epoch.
        mapping (address => Stake[]) stakes;

        uint256 totalStake;
    }

    bytes32 public currentDirectory;

    mapping (bytes32 => Directory) directories;

    function initialize(
        StakingManager stakingManager
    ) public initializer {
        OwnableUpgradeable.__Ownable_init();
        _stakingManager = stakingManager;
    }

    /*
     * We construct the directory entries by iterating through each valid stakee, and
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
    function constructDirectory() public onlyOwner returns (bytes32 direcrtoryId) {
        bytes32 directoryId = keccak256(abi.encodePacked(block.number));

        uint lowerBoundary = 0;

        for (uint i = 0; i < _stakingManager.getCountOfStakees(); i++) {
            address stakee = _stakingManager.stakees(i);
            uint totalStake = _stakingManager.totalStakes(stakee);

            // Only add stakee to the directory after passing
            // some validation

            if (totalStake < 1) {
                continue;
            }

            directories[directoryId].entries.push(DirectoryEntry(stakee, lowerBoundary + totalStake));

            address[] memory stakers = _stakingManager.getStakers(stakee);
            for (uint j = 0; j < stakers.length; j++) {
                StakingManager.Stake memory stake = _stakingManager.getStake(stakers[j], stakee);
                directories[directoryId].stakes[stakee].push(
                    Stake(
                        stakers[j],
                        stake.amount
                    )
                );
            }

            lowerBoundary += totalStake;
        }

        directories[directoryId].totalStake = _stakingManager.getTotalStake();

        currentDirectory = directoryId;

        return directoryId;
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

    function getStakes(bytes32 directoryId, address stakee) public view returns (Stake[] memory) {
        return directories[directoryId].stakes[stakee];
    }
}
