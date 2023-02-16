// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.13;

import "./Manager.sol";
import "../Payments/Ticketing/RewardsManager.sol";
import "../Utils.sol";
import "../Manageable.sol";

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/math/SafeCast.sol";

/**
 * @notice The Directory contract constructs and manages a structure holding the current stakes,
 * which is queried against using the scan function. The scan function allows submitting
 * random points which will return a staked node's address in proportion to the stake it has.
 */
contract Directory is Initializable, Manageable {
    /** Sylo Staking Manager contract */
    StakingManager public _stakingManager;

    /** Sylo Rewards Manager contract */
    RewardsManager public _rewardsManager;

    struct DirectoryEntry {
        address stakee;
        uint256 boundary;
    }

    /**
     * @dev A Directory will be stored for every epoch. The directory will be
     * constructed piece by piece as Nodes join, each adding their own
     * directory entry based on their current stake value.
     */
    struct Directory {
        DirectoryEntry[] entries;
        mapping(address => uint256) stakes;
        uint256 totalStake;
    }

    event CurrentDirectoryUpdated(uint256 currentDirectory);

    /**
     * @notice The epoch ID of the current directory.
     */
    uint256 public currentDirectory;

    /**
     * @notice Tracks every directory, which will be indexed by an epoch ID
     */
    mapping(uint256 => Directory) public directories;

    function initialize(
        StakingManager stakingManager,
        RewardsManager rewardsManager
    ) external initializer {
        OwnableUpgradeable.__Ownable_init();
        _stakingManager = stakingManager;
        _rewardsManager = rewardsManager;
    }

    /**
     * @notice This function should be called when a new epoch is initialized.
     * This will set the current directory to the specified epoch. This is only
     * callable by the owner of this contract, which should be the EpochsManager
     * contract.
     * @dev After deployment, the EpochsManager should immediately be set as
     * the owner.
     * @param epochId The ID of the specified epoch.
     */
    function setCurrentDirectory(uint256 epochId) external onlyManager {
        currentDirectory = epochId;
        emit CurrentDirectoryUpdated(epochId);
    }

    /**
     * @notice This function is called by a node as a prerequisite to participate in the next epoch.
     * @dev This will construct the directory as nodes join. The directory is constructed
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
    function joinNextDirectory(address stakee) external onlyManager {
        uint256 totalStake = _stakingManager.getStakeeTotalManagedStake(stakee);
        require(totalStake > 0, "No stake to join epoch");

        uint256 currentStake = _stakingManager.getCurrentStakerAmount(stakee, stakee);
        uint16 ownedStakeProportion = SyloUtils.asPerc(
            SafeCast.toUint128(currentStake),
            totalStake
        );

        uint16 minimumStakeProportion = _stakingManager.minimumStakeProportion();

        uint256 joiningStake;
        if (ownedStakeProportion >= minimumStakeProportion) {
            joiningStake = totalStake;
        } else {
            // if the node is below the minimum stake proportion, then we reduce
            // the stake used to join the epoch proportionally
            joiningStake = (totalStake * ownedStakeProportion) / minimumStakeProportion;
        }
        require(joiningStake > 0, "Joining stake not greater than 0");

        uint256 epochId = currentDirectory + 1;

        require(directories[epochId].stakes[stakee] == 0, "Node already joined next epoch");

        uint256 nextBoundary = directories[epochId].totalStake + joiningStake;

        directories[epochId].entries.push(DirectoryEntry(stakee, nextBoundary));
        directories[epochId].stakes[stakee] = joiningStake;
        directories[epochId].totalStake = nextBoundary;
    }

    /**
     * @notice Call this to perform a stake-weighted scan to find the Node assigned
     * to the given point of the current directory.
     * @param point The point, which will usually be a hash of a public key.
     */
    function scan(uint128 point) external view returns (address stakee) {
        return _scan(point, currentDirectory);
    }

    /**
     * @notice Call this to perform a stake-weighted scan to find the Node assigned
     * to the given point of the requested directory.
     * @param point The point, which will usually be a hash of a public key.
     * @param epochId The epoch id associated with the directory to scan.
     */
    function scanWithEpochId(
        uint128 point,
        uint256 epochId
    ) external view returns (address stakee) {
        return _scan(point, epochId);
    }

    /**
     * @notice Call this to perform a stake-weighted scan to find the Node assigned
     * to the given point of the requested directory (internal).
     * @dev The current implementation will perform a binary search through
     * the directory. This can allow gas costs to be low if this needs to be
     * used in a transaction.
     * @param point The point, which will usually be a hash of a public key.
     * @param epochId The epoch id associated with the directory to scan.
     */
    function _scan(uint128 point, uint256 epochId) internal view returns (address stakee) {
        uint256 entryLength = directories[epochId].entries.length;

        if (entryLength == 0) {
            return address(0);
        }

        // Staking all the Sylo would only be 94 bits, so multiplying this with
        // a uint128 cannot overflow a uint256.
        uint256 expectedVal = (directories[epochId].totalStake * uint256(point)) >> 128;

        uint256 left;
        uint256 right = entryLength - 1;

        // perform a binary search through the directory
        uint256 lower;
        uint256 upper;
        uint256 index;

        while (left <= right) {
            index = (left + right) >> 1;

            lower = index == 0 ? 0 : directories[epochId].entries[index - 1].boundary;
            upper = directories[epochId].entries[index].boundary;

            if (expectedVal >= lower && expectedVal < upper) {
                return directories[epochId].entries[index].stakee;
            } else if (expectedVal < lower) {
                right = index - 1;
            } else {
                // expectedVal >= upper
                left = index + 1;
            }
        }
    }

    /**
     * @notice Retrieve the total stake a Node has for the directory in the
     * specified epoch.
     * @param epochId The ID of the epoch.
     * @param stakee The address of the Node.
     * @return The amount of stake the Node has for the given directory in SOLO.
     */
    function getTotalStakeForStakee(
        uint256 epochId,
        address stakee
    ) external view returns (uint256) {
        return directories[epochId].stakes[stakee];
    }

    /**
     * @notice Retrieve the total stake for a directory in the specified epoch, which
     * will be the sum of the stakes for all Nodes participating in that epoch.
     * @param epochId The ID of the epoch.
     * @return The total amount of stake in SOLO.
     */
    function getTotalStake(uint256 epochId) external view returns (uint256) {
        return directories[epochId].totalStake;
    }

    /**
     * @notice Retrieve all entries for a directory in a specified epoch.
     * @return An array of all the directory entries.
     */
    function getEntries(
        uint256 epochId
    ) external view returns (address[] memory, uint256[] memory) {
        uint256 entryLength = directories[epochId].entries.length;

        address[] memory stakees = new address[](entryLength);
        uint256[] memory boundaries = new uint256[](entryLength);

        DirectoryEntry memory entry;
        DirectoryEntry[] memory entries = directories[epochId].entries;

        for (uint256 i; i < entryLength; ++i) {
            entry = entries[i];
            stakees[i] = entry.stakee;
            boundaries[i] = entry.boundary;
        }
        return (stakees, boundaries);
    }
}
