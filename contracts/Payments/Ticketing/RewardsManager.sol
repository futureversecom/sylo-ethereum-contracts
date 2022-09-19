// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.13;

import "../../Manageable.sol";
import "../../Staking/Manager.sol";
import "../../Epochs/Manager.sol";
import "../../Utils.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "abdk-libraries-solidity/ABDKMath64x64.sol";

/**
 * @notice Handles epoch based reward pools that are incremented from redeeming tickets.
 * Nodes use this contract to set up their reward pool for the next epoch,
 * and stakers use this contract to track and claim staking rewards.
 * @dev After deployment, the SyloTicketing contract should be
 * set up as a manager to be able to call certain restricted functions.
 */
contract RewardsManager is Initializable, OwnableUpgradeable, Manageable {
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
    mapping(address => int128) cumulativeRewardFactors;

    /**
     * @notice When a staker uses the CRF to calculate their share of the reward, the
     * contract needs to track the value of the CRF, and use this value in the next calculation.
     * The key to this mapping is a hash of the Node's address and the staker's address.
     */
    mapping(bytes32 => int128) lastUsedCRFs;

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
     */
    mapping(bytes32 => uint256) public unclaimedStakingRewards;

    /**
     * @notice Tracks each Node's most recently active reward pool
     */
    mapping(address => uint256) public latestActiveRewardPools;

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

    /**
     * @notice Tracks each reward pool initialized by a Node. The key to this map
     * is derived from the epochId and the Node's address.
     */
    mapping(bytes32 => RewardPool) public rewardPools;

    function initialize(
        IERC20 token,
        StakingManager stakingManager,
        EpochsManager epochsManager
    ) external initializer {
        OwnableUpgradeable.__Ownable_init();
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
    function getRewardPool(uint256 epochId, address stakee)
        external
        view
        returns (RewardPool memory)
    {
        return rewardPools[getRewardPoolKey(epochId, stakee)];
    }

    /**
     * @notice Retrieve the total accumulated reward that will be distributed to a Node's
     * delegated stakers for a given epoch.
     * @param epochId The ID of the epoch the reward pool was initialized in.
     * @param stakee The address of the Node.
     * @return The total accumulated staker reward in SOLO.
     */
    function getRewardPoolStakersTotal(uint256 epochId, address stakee)
        external
        view
        returns (uint256)
    {
        return rewardPools[getRewardPoolKey(epochId, stakee)].stakersRewardTotal;
    }

    /**
     * @notice Retrieve the total active stake that will be used for a Node's reward
     * pool in a given epoch.
     * @param epochId The ID of the epoch the reward pool was initialized in.
     * @param stakee The address of the Node.
     * @return The total active stake for that reward pool in SOLO.
     */
    function getRewardPoolActiveStake(uint256 epochId, address stakee)
        external
        view
        returns (uint256)
    {
        return rewardPools[getRewardPoolKey(epochId, stakee)].totalActiveStake;
    }

    /**
     * @notice Retrieve the total unclaimed staking reward allocated to a Node's
     * delegated stakers.
     * @param stakee The address of the Node.
     * @return The total unclaimed staking reward in SOLO.
     */
    function getPendingRewards(address stakee) external view returns (uint256) {
        return pendingRewards[stakee];
    }

    /**
     * @notice This is used by Nodes to initialize their reward pool for
     * the next epoch. This function will revert if the caller has no stake, or
     * if the reward pool has already been initialized. The total active stake
     * for the next reward pool is calculated by summing up the total managed
     * stake held by the RewardsManager contract, plus any unclaimed staking rewards.
     */
    function initializeNextRewardPool(address stakee) external onlyManager {
        uint256 nextEpochId = _epochsManager.getNextEpochId();

        RewardPool storage nextRewardPool = rewardPools[getRewardPoolKey(nextEpochId, stakee)];
        require(
            nextRewardPool.initializedAt == 0,
            "The next reward pool has already been initialized"
        );

        uint256 totalStake = _stakingManager.getStakeeTotalManagedStake(stakee);
        require(totalStake > 0, "Must have stake to initialize a reward pool");
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
        EpochsManager.Epoch memory currentEpoch = _epochsManager.getCurrentActiveEpoch();

        RewardPool storage rewardPool = rewardPools[
            getRewardPoolKey(currentEpoch.iteration, stakee)
        ];
        require(
            rewardPool.totalActiveStake > 0,
            "Reward pool has not been initialized for the current epoch"
        );

        // Update the latest active reward pool for the node to be this pool
        if (latestActiveRewardPools[stakee] < currentEpoch.iteration) {
            latestActiveRewardPools[stakee] = currentEpoch.iteration;
        }

        uint256 stakersReward = SyloUtils.percOf(
            uint128(amount),
            currentEpoch.defaultPayoutPercentage
        );

        // transfer the node's fee reward to it's unclaimed reward value
        unclaimedStakingRewards[getStakerKey(stakee, stakee)] += (amount - stakersReward);

        // update the value of the reward owed to the stakers
        pendingRewards[stakee] += stakersReward;

        rewardPool.stakersRewardTotal += stakersReward;

        // this is the first ticket redeemed for this reward, set the initial
        // CRF value for this pool
        if (rewardPool.initialCumulativeRewardFactor == 0) {
            rewardPool.initialCumulativeRewardFactor = cumulativeRewardFactors[stakee];
        }

        cumulativeRewardFactors[stakee] =
            ABDKMath64x64.add(
                cumulativeRewardFactors[stakee],
                ABDKMath64x64.div(
                    toFixedPointSYLO(stakersReward),
                    toFixedPointSYLO(rewardPool.totalActiveStake)
                )
            );
    }


    /**
     * @dev This function utilizes the cumulative reward factors, and the staker's
     * value in stake to calculate the staker's share of the pending reward.
     */
    function calculatePendingClaim(address stakee, address staker) internal view returns (uint256) {
        StakingManager.StakeEntry memory stakeEntry = _stakingManager.getStakeEntry(
            stakee,
            staker
        );
        if (stakeEntry.amount == 0) {
            return 0;
        }

        // Retrieve the latest active reward pool
        RewardPool memory latestActivePool = rewardPools[
            getRewardPoolKey(latestActiveRewardPools[stakee], stakee)
        ];

        // Retrieve the last used CRF value
        int128 initialCumulativeRewardFactor = lastUsedCRFs[getStakerKey(stakee, staker)];
        // Claims are only made up to the last epoch, so the final CRF for this claim
        // will be the initial CRF of the latest pool
        int128 finalCumulativeRewardFactor = latestActivePool.initialCumulativeRewardFactor;

        // We convert the staker amount to SYLO as the maximum uint256 value that
        // can be used for the fixed point representation is 2^64-1.
        int128 initialStake = toFixedPointSYLO(stakeEntry.amount);

        return fromFixedPointSYLO(
            ABDKMath64x64.mul(
                initialStake,
                ABDKMath64x64.sub(
                    finalCumulativeRewardFactor,
                    initialCumulativeRewardFactor
                )
            )
        );
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
        uint256 pendingClaim = calculatePendingClaim(stakee, staker);

        return pendingClaim + unclaimedStakingRewards[getStakerKey(stakee, staker)];
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
        uint256 pendingReward = calculatePendingClaim(stakee, msg.sender);
        uint256 totalClaim = pendingReward + unclaimedStakingRewards[getStakerKey(stakee, msg.sender)];

        require(totalClaim > 0, "Nothing to claim");

        unclaimedStakingRewards[getStakerKey(stakee, msg.sender)] = 0;
        pendingRewards[stakee] -= pendingReward;

        updateLastUsedCRF(stakee, msg.sender);

        _token.transfer(msg.sender, totalClaim);

        return totalClaim;
    }

    /**
     * @notice This is called by the staking manager to transfer pending rewards
     * to unclaimed rewards for a staker. This is required as the lase used CRF
     * needs to be updated whenever stake changes.
     */
    function updatePendingRewards(address stakee, address staker) external onlyManager {
        uint256 pendingReward = calculatePendingClaim(stakee, msg.sender);

        pendingRewards[stakee] -= pendingReward;
        unclaimedStakingRewards[getStakerKey(stakee, staker)] += pendingReward;

        updateLastUsedCRF(stakee, staker);
    }

    /**
     * @dev Called whenever a staker's share of reward is calculated from the stakee's pending
     * rewards.
     */
    function updateLastUsedCRF(address stakee, address staker) internal onlyManager {
        // Retrieve the latest active reward pool
        RewardPool memory latestActivePool = rewardPools[
            getRewardPoolKey(latestActiveRewardPools[stakee], stakee)
        ];
        lastUsedCRFs[getStakerKey(stakee, staker)] = latestActivePool.initialCumulativeRewardFactor;
    }
}
