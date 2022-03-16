// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "./Manager.sol";
import "../Payments/Ticketing/RewardsManager.sol";
import "../Utils.sol";
import "../Manageable.sol";

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

/**
 * @notice The Directory contract constructs and manages a structure holding the current stakes,
 * which is queried against using the scan function. The scan function allows submitting
 * random points which will return a staked node's address in proportion to the stake it has.
 */
contract Directory is Initializable, OwnableUpgradeable, Manageable {
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

        mapping (address => uint256) stakes;

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
    mapping (uint256 => Directory) public directories;

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
        uint256 managedStake = _stakingManager.getStakeeTotalManagedStake(stakee);
        uint256 stakeReward = _rewardsManager.unclaimedStakeRewards(stakee);
        uint256 totalStake = managedStake + stakeReward;
        require(totalStake > 0, "Can not join directory for next epoch without any stake");
        require(
            _stakingManager.checkMinimumStakeProportion(stakee),
            "Can not join directory without owning minimum amount of stake"
        );

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

    /**
     * @notice This method is the entry point for the directory scan.
     * It calls the derivePoint method to hash input values and send the result
     * to the _scan method to get the associated node address.
     * @param pk The user's public key.
     * @param epochId The specified epochId to select which directory to scan.
     * This will give the ability for users to scan using historical directories.
     * @param channel The specified channel that can effect the hash function's
     * output, which leads to different point and different node address. This will
     * allow user to switch node during the same epoch.
     * @return stakee The address of the node the user will be assigned to.
     */
    function scan(bytes memory pk, uint256 epochId, uint8 channel)  external view returns (address stakee) {
        uint128 point = derivePoint(pk, epochId, channel);
        return _scan(point, epochId);
    }

    /**
     * @notice This method will hash the public key, epoch Id, and channel
     * provided by user.
     * @param pk The user's public key.
     * @param epochId The specified epochId to select which directory to scan.
     * @param channel The specified channel that can effect the hash function's
     * output.
     * @return finalPoint The shifted hash value of the public key, epoch Id, and channel.
     * The hash value is shifted to reduce its size so that it can be processed
     * in the _scan method.
     */
    function derivePoint(bytes memory pk, uint256 epochId, uint8 channel) internal pure returns (uint128 finalPoint) {
        require(channel <= 48, "Only channels 0-48 are supported");

        bytes32 b = bytes32(keccak256(abi.encodePacked(pk, epochId, channel)));
        uint256 point = uint(b) >> 128;
        finalPoint = uint128(point);
        return finalPoint;
    }

    /**
     * @notice Call this to perform a stake-weighted scan to find the Node assigned
     * to the given point.
     * @dev The current implementation will perform a binary search through
     * the directory. This can allow gas costs to be low if this needs to be
     * used in a transaction.
     * @param point The point, which will usually be a hash of a public key.
     */
    function _scan(uint128 point, uint256 epochId) public view returns (address stakee) {
        if (directories[epochId].entries.length == 0) {
            return address(0);
        }

        // Staking all the Sylo would only be 94 bits, so multiplying this with
        // a uint128 cannot overflow a uint256.
        uint256 expectedVal = directories[epochId].totalStake * uint256(point) >> 128;

        uint256 left = 0;
        uint256 right = directories[epochId].entries.length - 1;

        // perform a binary search through the directory
        while (left <= right) {
            uint index = (left + right) / 2;

            uint lower = index == 0 ? 0 : directories[epochId].entries[index - 1].boundary;
            uint upper = directories[epochId].entries[index].boundary;

            if (expectedVal >= lower && expectedVal < upper) {
                return directories[epochId].entries[index].stakee;
            } else if (expectedVal < lower) {
                right = index - 1;
            } else {  // expectedVal >= upper
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
    function getTotalStakeForStakee(uint256 epochId, address stakee) external view returns (uint256) {
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
    function getEntries(uint256 epochId) external view returns (address[] memory, uint256[] memory) {
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
