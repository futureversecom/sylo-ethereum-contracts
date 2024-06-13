// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

interface ISyloStakingManager {
    /**
     * For every Node, there will be a mapping of the staker to a
     * StakeEntry. The stake entry tracks the amount of stake in SOLO,
     * and also when the stake was updated.
     */
    struct StakeEntry {
        uint256 amount;
        // Timestamp this entry was updated at (from block timestamp)
        uint256 updatedAt;
    }

    /**
     * Every Node must have stake in order to participate in the Epoch.
     * Stake can be provided by the Node itself or by other accounts in
     * the network.
     */
    struct Stake {
        // Tracks each stake entry associated to a node
        mapping(address => StakeEntry) entries;
        // The total stake held by this contract for a node,
        // which will be the sum of all addStake and unlockStake calls
        uint256 totalManagedStake;
    }

    /**
     * This struct will track stake that is in the process of unlocking.
     */
    struct Unlocking {
        uint256 amount; // Amount of stake unlocking
        uint256 unlockAt; // Timestamp the stake becomes withdrawable
    }

    function setUnlockDuration(uint256 _unlockDuration) external;

    function addStake(address node, uint256 amount) external;

    function unlockStake(address node, uint256 amount) external returns (uint256);

    function withdrawStake(address node) external;

    function cancelUnlocking(address node, uint256 amount) external;

    function transferStake(address from, address to, uint256 amount) external;

    function getTotalManagedStake() external view returns (uint256);

    function getManagedStake(address node, address user) external view returns (StakeEntry memory);

    function getTotalManagedStakeByNode(address node) external view returns (uint256);
}
