// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

interface IStakingManager {
    /**
     * For every Node, there will be a mapping of the staker to a
     * StakeEntry. The stake entry tracks the amount of stake in SOLO,
     * and also when the stake was updated.
     */
    struct StakeEntry {
        uint256 amount;
        // Block number this entry was updated at
        uint256 updatedAt;
        // Epoch this entry was updated. The stake will become active
        // in the following epoch
        uint256 epochId;
    }

    /**
     * Every Node must have stake in order to participate in the Epoch.
     * Stake can be provided by the Node itself or by other accounts in
     * the network.
     */
    struct Stake {
        // Track each stake entry associated to a node
        mapping(address => StakeEntry) stakeEntries;
        // The total stake held by this contract for a node,
        // which will be the sum of all addStake and unlockStake calls
        uint256 totalManagedStake;
    }

    /**
     * This struct will track stake that is in the process of unlocking.
     */
    struct Unlock {
        uint256 amount; // Amount of stake unlocking
        uint256 unlockAt; // Block number the stake becomes withdrawable
    }

    function setUnlockDuration(uint256 _unlockDuration) external;

    function setMinimumStakeProportion(uint32 _minimumStakeProportion) external;

    function addStake(uint256 amount, address stakee) external;

    function unlockStake(uint256 amount, address stakee) external returns (uint256);

    function withdrawStake(address stakee) external;

    function cancelUnlocking(uint256 amount, address stakee) external;

    function calculateCapacityFromSeekerPower(uint256 seekerId) external view returns (uint256);

    function calculateMaxAdditionalDelegatedStake(address stakee) external view returns (uint256);

    function getTotalManagedStake() external view returns (uint256);

    function getStakeEntry(
        address stakee,
        address staker
    ) external view returns (StakeEntry memory);

    function getStakeeTotalManagedStake(address stakee) external view returns (uint256);
}
