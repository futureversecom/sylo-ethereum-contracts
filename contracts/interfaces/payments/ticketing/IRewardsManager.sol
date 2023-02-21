// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

interface IRewardsManager {
    /**
     * @dev This type will hold the necessary information for delegated stakers
     * to make reward claims against their Node. Every Node will initialize
     * and store a new Reward Pool for each epoch they participate in.
     */
    struct RewardPool {
        // Tracks the balance of the reward pool owed to the stakers
        uint256 stakersRewardTotal;
        // Tracks the block number this reward pool was initialized
        uint256 initializedAt;
        // The total active stake for the node for will be the sum of the
        // stakes owned by its delegators and the node's own stake.
        uint256 totalActiveStake;
        // track the cumulative reward factor as of the time the first ticket
        // for this pool was redeemed
        int128 initialCumulativeRewardFactor;
    }

    struct LastClaim {
        // The epoch the claim was made.
        uint256 claimedAt;
        // The stake at the time the claim was made. This is tracked as
        // rewards can only be claimed after an epoch has ended, but the
        // user's stake may have changed by then. This field tracks the
        // staking value before the change so the reward for that epoch
        // can be manually calculated.
        uint256 stake;
    }

    function getRewardPool(
        uint256 epochId,
        address stakee
    ) external view returns (RewardPool memory);

    function getRewardPoolKey(uint256 epochId, address stakee) external pure returns (bytes32);

    function getRewardPoolStakersTotal(
        uint256 epochId,
        address stakee
    ) external view returns (uint256);

    function getRewardPoolActiveStake(
        uint256 epochId,
        address stakee
    ) external view returns (uint256);

    function getPendingRewards(address stakee) external view returns (uint256);

    function getLastClaim(address stakee, address staker) external view returns (LastClaim memory);

    function getTotalEpochRewards(uint256 epochId) external view returns (uint256);

    function getTotalEpochStakingRewards(uint256 epochId) external view returns (uint256);

    function initializeNextRewardPool(address stakee) external;

    function incrementRewardPool(address stakee, uint256 amount) external;

    function claimStakingRewards(address stakee) external returns (uint256);

    function updatePendingRewards(address stakee, address staker) external;
}
