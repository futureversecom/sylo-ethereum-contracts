// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "abdk-libraries-solidity/ABDKMath64x64.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/math/SafeCast.sol";
import "abdk-libraries-solidity/ABDKMath64x64.sol";

import "../../libraries/Manageable.sol";
import "../../libraries/SyloUtils.sol";
import "../../staking/StakingManager.sol";
import "../../epochs/EpochsManager.sol";

/**
 * @notice Handles epoch based reward pools that are incremented from redeeming tickets.
 * Nodes use this contract to set up their reward pool for the next epoch,
 * and stakers use this contract to track and claim staking rewards.
 * @dev After deployment, the SyloTicketing contract should be
 * set up as a manager to be able to call certain restricted functions.
 */
contract RewardsManager is Initializable, Manageable {
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

    uint256 internal constant ONE_SYLO = 1 ether;

    // 64x64 Fixed point representation of 1 SYLO (10**18 >> 64)
    int128 internal constant ONE_SYLO_FIXED = 18446744073709551616000000000000000000;

    /** ERC20 Sylo token contract. */
    IERC20 public _token;

    /** Sylo Staking Manager contract. */
    StakingManager public _stakingManager;

    /** Sylo Epochs Manager. */
    EpochsManager public _epochsManager;

    /**
     * @dev Each node will manage a cumulative reward factor (CRF) that is incremented
     * whenever a ticket is redeemed. This factor acts as a single value
     * that can be used to calculate any particular staker's reward share. This
     * prevents the need to individually track each staker's proportion, and also allows
     * a claim calculation to be performed without iterating through every epoch.
     *
     * The CRF is calculated as CRF = CRF + Reward / TotalStake.
     */
    mapping(address => int128) private cumulativeRewardFactors;

    /**
     * @notice Tracks the last epoch a delegated staker made a reward claim in.
     * The key to this mapping is a hash of the Node's address and the delegated
     * stakers address.
     */
    mapping(bytes32 => LastClaim) public lastClaims;

    /**
     * @notice Tracks each Nodes total pending rewards in SOLOs. This
     * value is accumulated as Node's redeem tickets. Rewards are pending if the
     * distribution amongst the stakers has not been accounted for yet. Pending rewards
     * are transferred to unclaimed rewards once the the staker's share has been
     * calculated.
     */
    mapping(address => uint256) public pendingRewards;

    /**
     * @notice Tracks rewards for stakers after the stakers share has been calculated,
     * but has not actually been claimed by the staker.
     * The node fee reward is also added to the node's unclaimedStakingRewards.
     */
    mapping(bytes32 => uint256) public unclaimedStakingRewards;

    /**
     * @notice Tracks each Node's most recently active reward pool
     */
    mapping(address => uint256) public latestActiveRewardPools;

    /**
     * @notice Tracks total accumulated rewards in each epoch
     */
    mapping(uint256 => uint256) public totalEpochRewards;

    /**
     * @notice Tracks total accumulated staking rewards in each epoch
     */
    mapping(uint256 => uint256) public totalEpochStakingRewards;

    /**
     * @notice Tracks each reward pool initialized by a Node. The key to this map
     * is derived from the epochId and the Node's address.
     */
    mapping(bytes32 => RewardPool) public rewardPools;

    error NoRewardToClaim();
    error RewardPoolNotExist();
    error RewardPoolAlreadyExist();
    error NoStakeToCreateRewardPool();

    function initialize(
        IERC20 token,
        StakingManager stakingManager,
        EpochsManager epochsManager
    ) external initializer {
        Ownable2StepUpgradeable.__Ownable2Step_init();
        _token = token;
        _epochsManager = epochsManager;
        _stakingManager = stakingManager;
    }

    /**
     * @notice Returns the key used to index a reward pool. The key is a hash of
     * the epochId and Node's address.
     * @param epochId The epoch ID the reward pool was created in.
     * @param stakee The address of the Node.
     * @return A byte-array representing the reward pool key.
     */
    function getRewardPoolKey(uint256 epochId, address stakee) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(epochId, stakee));
    }

    /**
     * @notice Returns the key used to index staking claims. The key is a hash of
     * the Node's address and the staker's address.
     * @param stakee The address of the Node.
     * @param staker The address of the stake.
     * @return A byte-array representing the key.
     */
    function getStakerKey(address stakee, address staker) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(stakee, staker));
    }

    /**
     * @notice Retrieve the reward pool initialized by the given node, at the specified
     * epoch.
     * @param epochId The ID of the epoch the reward pool was initialized in.
     * @param stakee The address of the Node.
     * @return The reward pool.
     */
    function getRewardPool(
        uint256 epochId,
        address stakee
    ) external view returns (RewardPool memory) {
        return rewardPools[getRewardPoolKey(epochId, stakee)];
    }

    /**
     * @notice Retrieve the total accumulated reward that will be distributed to a Node's
     * delegated stakers for a given epoch.
     * @param epochId The ID of the epoch the reward pool was initialized in.
     * @param stakee The address of the Node.
     * @return The total accumulated staker reward in SOLO.
     */
    function getRewardPoolStakersTotal(
        uint256 epochId,
        address stakee
    ) external view returns (uint256) {
        return rewardPools[getRewardPoolKey(epochId, stakee)].stakersRewardTotal;
    }

    /**
     * @notice Retrieve the total active stake that will be used for a Node's reward
     * pool in a given epoch.
     * @param epochId The ID of the epoch the reward pool was initialized in.
     * @param stakee The address of the Node.
     * @return The total active stake for that reward pool in SOLO.
     */
    function getRewardPoolActiveStake(
        uint256 epochId,
        address stakee
    ) external view returns (uint256) {
        return rewardPools[getRewardPoolKey(epochId, stakee)].totalActiveStake;
    }

    /**
     * @notice Retrieve the total pending staking reward allocated to a Node's
     * stakers.
     * @param stakee The address of the Node.
     * @return The total pending staking reward in SOLO.
     */
    function getPendingRewards(address stakee) external view returns (uint256) {
        return pendingRewards[stakee];
    }

    /**
     * @notice Retrieves the ID of the epoch in which a staker last made their
     * staking claim.
     * @param stakee The address of the Node.
     * @param staker The address of the staker.
     * @return The ID of the epoch.
     */
    function getLastClaim(
        address stakee,
        address staker
    ) external view returns (LastClaim memory) {
        return lastClaims[getStakerKey(stakee, staker)];
    }

    /**
     * @notice Retrieves the total accumulated rewards for a specific epoch.
     * @param epochId The epoch id.
     * @return The total reward in that epoch, including staking rewards and fee
     * reward.
     */
    function getTotalEpochRewards(uint256 epochId) external view returns (uint256) {
        return totalEpochRewards[epochId];
    }

    /**
     * @notice Retrieves the total accumulated rewards for stakers in a specific epoch.
     * @param epochId The epoch id.
     * @return The total staking reward in that epoch.
     */
    function getTotalEpochStakingRewards(uint256 epochId) external view returns (uint256) {
        return totalEpochStakingRewards[epochId];
    }

    /**
     * @notice This is used by Nodes to initialize their reward pool for
     * the next epoch. This function will revert if the caller has no stake, or
     * if the reward pool has already been initialized. The total active stake
     * for the next reward pool is calculated by summing up the total managed
     * stake held by the RewardsManager contract.
     */
    function initializeNextRewardPool(address stakee) external onlyManager {
        uint256 nextEpochId = _epochsManager.getNextEpochId();

        RewardPool storage nextRewardPool = rewardPools[getRewardPoolKey(nextEpochId, stakee)];
        if (nextRewardPool.initializedAt != 0) {
            revert RewardPoolAlreadyExist();
        }

        uint256 totalStake = _stakingManager.getStakeeTotalManagedStake(stakee);
        if (totalStake == 0) {
            revert NoStakeToCreateRewardPool();
        }

        nextRewardPool.totalActiveStake = totalStake;

        nextRewardPool.initializedAt = block.number;
    }

    /**
     * @dev This function should be called by the Ticketing contract when a
     * ticket is successfully redeemed. The face value of the ticket
     * should be split between incrementing the node's reward balance,
     * and the reward balance for the node's delegated stakers. The face value
     * will be added to the current reward pool's balance. This function will
     * fail if the Ticketing contract has not been set as a manager.
     * @param stakee The address of the Node.
     * @param amount The face value of the ticket in SOLO.
     */
    function incrementRewardPool(address stakee, uint256 amount) external onlyManager {
        (uint256 epochId, EpochsManager.Epoch memory currentEpoch) = _epochsManager
            .getCurrentActiveEpoch();

        RewardPool storage rewardPool = rewardPools[getRewardPoolKey(epochId, stakee)];
        if (rewardPool.initializedAt == 0) {
            revert RewardPoolNotExist();
        }

        // Update the latest active reward pool for the node to be this pool
        if (latestActiveRewardPools[stakee] < epochId) {
            latestActiveRewardPools[stakee] = epochId;
        }

        uint256 stakersReward = SyloUtils.percOf(
            SafeCast.toUint128(amount),
            currentEpoch.defaultPayoutPercentage
        );

        // transfer the node's fee reward to it's unclaimed reward value
        unclaimedStakingRewards[getStakerKey(stakee, stakee)] =
            unclaimedStakingRewards[getStakerKey(stakee, stakee)] +
            (amount - stakersReward);

        // update the value of the reward owed to the stakers
        pendingRewards[stakee] = pendingRewards[stakee] + stakersReward;

        // if this is the first ticket redeemed for this reward, set the initial
        // CRF value for this pool
        if (rewardPool.stakersRewardTotal == 0) {
            rewardPool.initialCumulativeRewardFactor = cumulativeRewardFactors[stakee];
        }

        rewardPool.stakersRewardTotal = rewardPool.stakersRewardTotal + stakersReward;

        cumulativeRewardFactors[stakee] = ABDKMath64x64.add(
            cumulativeRewardFactors[stakee],
            ABDKMath64x64.div(
                toFixedPointSYLO(stakersReward),
                toFixedPointSYLO(rewardPool.totalActiveStake)
            )
        );

        totalEpochRewards[epochId] = totalEpochRewards[epochId] + amount;
        totalEpochStakingRewards[epochId] = totalEpochStakingRewards[epochId] + stakersReward;
    }

    /**
     * @dev This function utilizes the cumulative reward factors, and the staker's
     * value in stake to calculate the staker's share of the pending reward.
     */
    function calculatePendingClaim(
        bytes32 stakerKey,
        address stakee,
        address staker
    ) internal view returns (uint256) {
        uint256 claim = calculateInitialClaim(stakerKey, stakee);

        // find the first reward pool where their stake was active and had
        // generated rewards
        uint256 activeAt;
        RewardPool memory initialActivePool;

        uint256 currentEpochId = _epochsManager.currentIteration();

        for (uint256 i = lastClaims[stakerKey].claimedAt + 1; i < currentEpochId; ++i) {
            initialActivePool = rewardPools[getRewardPoolKey(i, stakee)];
            // check if node initialized a reward pool for this epoch and
            // gained rewards
            if (initialActivePool.initializedAt > 0 && initialActivePool.stakersRewardTotal > 0) {
                activeAt = i;
                break;
            }
        }

        if (activeAt == 0) {
            return claim;
        }

        StakingManager.StakeEntry memory stakeEntry = _stakingManager.getStakeEntry(
            stakee,
            staker
        );

        // We convert the staker amount to SYLO as the maximum uint256 value that
        // can be used for the fixed point representation is 2^64-1.
        int128 initialStake = toFixedPointSYLO(stakeEntry.amount);

        int128 initialCumulativeRewardFactor = initialActivePool.initialCumulativeRewardFactor;

        int128 finalCumulativeRewardFactor = getFinalCumulativeRewardFactor(
            stakee,
            currentEpochId
        );

        return
            claim +
            fromFixedPointSYLO(
                ABDKMath64x64.mul(
                    initialStake,
                    ABDKMath64x64.sub(finalCumulativeRewardFactor, initialCumulativeRewardFactor)
                )
            );
    }

    /**
     * Manually calculates the reward claim for the first epoch the claim is being
     * made for. This manual calculation is necessary as claims are only made up
     * to the previous epoch.
     */
    function calculateInitialClaim(
        bytes32 stakerKey,
        address stakee
    ) internal view returns (uint256) {
        LastClaim memory lastClaim = lastClaims[stakerKey];

        // if we have already made a claim up to the previous epoch, then
        // there is no need to calculate the initial claim
        if (_epochsManager.currentIteration() == lastClaim.claimedAt) {
            return 0;
        }

        RewardPool memory firstRewardPool = rewardPools[
            getRewardPoolKey(lastClaim.claimedAt, stakee)
        ];

        // if there was no reward pool initialized for the first epoch,
        // then there is no need to calculate the initial claim
        if (firstRewardPool.totalActiveStake == 0) {
            return 0;
        }

        return
            (firstRewardPool.stakersRewardTotal * lastClaim.stake) /
            firstRewardPool.totalActiveStake;
    }

    /**
     * Determines the cumulative reward factor to use for claim calculations. The
     * CRF will depend on when the Node last initialized a reward pool, and also when
     * the staker last made their claim.
     */
    function getFinalCumulativeRewardFactor(
        address stakee,
        uint256 currentEpochId
    ) internal view returns (int128) {
        int128 finalCumulativeRewardFactor;

        // Get the cumulative reward factor for the Node
        // for the start of this epoch, since we only perform
        // calculations up to the end of the previous epoch.
        if (latestActiveRewardPools[stakee] < currentEpochId) {
            // If the Node has not been active, then the final
            // cumulative reward factor will just be the current one.
            finalCumulativeRewardFactor = cumulativeRewardFactors[stakee];
        } else {
            // We are calculating the claim for an active epoch, the
            // final cumulative reward factor will be taken from the start of this
            // epoch (end of previous epoch).
            RewardPool storage latestRewardPool = rewardPools[
                getRewardPoolKey(latestActiveRewardPools[stakee], stakee)
            ];
            finalCumulativeRewardFactor = latestRewardPool.initialCumulativeRewardFactor;
        }

        return finalCumulativeRewardFactor;
    }

    /**
     * @notice Call this function to calculate the total reward owed to a staker.
     * This value will include all epochs since the last claim was made up to
     * the previous epoch. This will also add any pending rewards to the
     * final value as well.
     * @dev This function will utilize the cumulative reward factor to perform the
     * calculation, keeping the gas cost scaling of this function to a constant value.
     * @param stakee The address of the Node.
     * @param staker The address of the staker.
     * @return The value of the reward owed to the staker in SOLO.
     */
    function calculateStakerClaim(address stakee, address staker) public view returns (uint256) {
        bytes32 stakerKey = getStakerKey(stakee, staker);
        uint256 pendingClaim = calculatePendingClaim(stakerKey, stakee, staker);

        return pendingClaim + unclaimedStakingRewards[stakerKey];
    }

    /**
     * Helper function to convert a uint256 value in SOLOs to a 64.64 fixed point
     * representation in SYLOs while avoiding any possibility of overflow.
     * Any remainders from converting SOLO to SYLO is explicitly handled to mitigate
     * precision loss. The error when using this function is [-1/2^64, 0].
     */
    function toFixedPointSYLO(uint256 amount) internal pure returns (int128) {
        int128 fullSylos = ABDKMath64x64.fromUInt(amount / ONE_SYLO);
        int128 fracSylos = ABDKMath64x64.fromUInt(amount % ONE_SYLO); // remainder

        return ABDKMath64x64.add(fullSylos, ABDKMath64x64.div(fracSylos, ONE_SYLO_FIXED));
    }

    /**
     * Helper function to convert a 64.64 fixed point value in SYLOs to a uint256
     * representation in SOLOs while avoiding any possibility of overflow.
     */
    function fromFixedPointSYLO(int128 amount) internal pure returns (uint256) {
        uint256 fullSylos = ABDKMath64x64.toUInt(amount);
        uint256 fullSolos = fullSylos * ONE_SYLO;

        // calculate the value lost when converting the fixed point amount to a uint
        int128 fracSylos = ABDKMath64x64.sub(amount, ABDKMath64x64.fromUInt(fullSylos));
        uint256 fracSolos = ABDKMath64x64.toUInt(ABDKMath64x64.mul(fracSylos, ONE_SYLO_FIXED));

        return fullSolos + fracSolos;
    }

    /**
     * @notice Call this function to claim rewards as a staker. The
     * SYLO tokens will be transferred to the caller's account. This function will
     * fail if there exists no reward to claim. Note: Calling this will remove
     * the current unclaimed reward from being used as stake in the next round.
     * @param stakee The address of the Node to claim against.
     */
    function claimStakingRewards(address stakee) external returns (uint256) {
        bytes32 stakerKey = getStakerKey(stakee, msg.sender);
        uint256 pendingReward = calculatePendingClaim(stakerKey, stakee, msg.sender);

        uint256 totalClaim = pendingReward + unclaimedStakingRewards[stakerKey];
        if (totalClaim == 0) {
            revert NoRewardToClaim();
        }

        delete unclaimedStakingRewards[stakerKey];
        pendingRewards[stakee] = pendingRewards[stakee] - pendingReward;

        updateLastClaim(stakee, msg.sender);

        SafeERC20.safeTransfer(_token, msg.sender, totalClaim);

        return totalClaim;
    }

    /**
     * @notice This is called by the staking manager to transfer pending rewards
     * to unclaimed rewards for a staker. This is required as the last used CRF
     * needs to be updated whenever stake changes.
     */
    function updatePendingRewards(address stakee, address staker) external onlyManager {
        bytes32 stakerKey = getStakerKey(stakee, staker);
        uint256 pendingReward = calculatePendingClaim(stakerKey, stakee, staker);

        pendingRewards[stakee] = pendingRewards[stakee] - pendingReward;

        unclaimedStakingRewards[stakerKey] = unclaimedStakingRewards[stakerKey] + pendingReward;

        updateLastClaim(stakee, staker);
    }

    function updateLastClaim(address stakee, address staker) internal {
        StakingManager.StakeEntry memory stakeEntry = _stakingManager.getStakeEntry(
            stakee,
            staker
        );
        lastClaims[getStakerKey(stakee, staker)] = LastClaim(
            _epochsManager.currentIteration(),
            stakeEntry.amount
        );
    }
}
