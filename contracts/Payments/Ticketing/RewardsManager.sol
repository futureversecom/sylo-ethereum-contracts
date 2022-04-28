// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "../../Manageable.sol";
import "../../Staking/Manager.sol";
import "../../Epochs/Manager.sol";
import "../../Utils.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "abdk-libraries-solidity/ABDKMath64x64.sol";
import "hardhat/console.sol";

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
     * @notice Tracks each Nodes total unclaimed rewards in SOLOs. This value
     * accumulated as Node's redeem tickets, and tracks the portion of the
     * reward which is allocated to the Node as payment for operating
     * a Sylo Node.
     */
    mapping (address => uint256) public unclaimedNodeRewards;

    /**
     * @notice Tracks each Nodes total unclaimed staking rewards in SOLOs. This
     * value is accumulated as Node's redeem tickets, and tracks the portion of
     * the reward which is allocated to its delegated stakers.
     */
    mapping (address => uint256) public unclaimedStakeRewards;

    /**
     * @notice Tracks each Node's most recently initialized reward pool
     */
    mapping (address => uint256) public latestActiveRewardPools;

    struct LastClaim {
        // The block number the claim was made.
        uint256 claimedAt;

        // The stake at the time the claim was made. This is tracked as
        // rewards can only be claimed after an epoch has ended, but the
        // user's stake may have changed by then. This field tracks the
        // before the change so the reward for that epoch can be manually
        // calculated.
        uint256 stake;
    }

    /**
     * @notice Tracks the last epoch a delegated staker made a reward claim in.
     * The key to this mapping is a hash of the Node's address and the delegated
     * stakers address.
     */
    mapping (bytes32 => LastClaim) public lastClaims;

    /**
     * @dev This type will hold the necessary information for delegated stakers
     * to make reward claims against their Node. Every Node will initialize
     * and store a new Reward Pool for each they participate in.
     */
    struct RewardPool {
        // Tracks the balance of the reward pool owed to the stakers
        uint256 stakersRewardTotal;

        // Tracks the block number this reward pool was initialized
        uint256 initializedAt;

        // The total active stake for the node for will be the sum of the
        // stakes owned by its delegators plus the value of the unclaimed
        // staker rewards at the time this pool was initialized
        uint256 totalActiveStake;

        // track the cumulative reward factor as of the time the pool was initialized
        int128 initialCumulativeRewardFactor;

        // track the cumulative reward factor as a 64x64 fixed-point value
        int128 cumulativeRewardFactor;
    }

    /**
     * @notice Tracks each reward pool initialized by a Node. The key to this map
     * is derived from the epochId and the Node's address.
     */
    mapping (bytes32 => RewardPool) public rewardPools;

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
    function getStakerKey(address stakee, address staker) public pure returns(bytes32) {
        return keccak256(abi.encodePacked(stakee, staker));
    }

    /**
     * @notice Retrieves the ID of the epoch in which a staker last made their
     * staking claim.
     * @param stakee The address of the Node.
     * @param staker The address of the staker.
     * @return The ID of the epoch.
     */
    function getLastClaim(address stakee, address staker) external view returns (LastClaim memory) {
        return lastClaims[getStakerKey(stakee, staker)];
    }

    /**
     * @notice Retrieve the reward pool initialized by the given node, at the specified
     * epoch.
     * @param epochId The ID of the epoch the reward pool was initialized in.
     * @param stakee The address of the Node.
     * @return The reward pool.
     */
    function getRewardPool(uint256 epochId, address stakee) external view returns (RewardPool memory) {
        return rewardPools[getRewardPoolKey(epochId, stakee)];
    }

    /**
     * @notice Retrieve the total accumulated reward that will be distributed to a Node's
     * delegated stakers for a given epoch.
     * @param epochId The ID of the epoch the reward pool was initialized in.
     * @param stakee The address of the Node.
     * @return The total accumulated staker reward in SOLO.
     */
    function getRewardPoolStakersTotal(uint256 epochId, address stakee) external view returns (uint256) {
        return rewardPools[getRewardPoolKey(epochId, stakee)].stakersRewardTotal;
    }

    /**
     * @notice Retrieve the total active stake that will be used for a Node's reward
     * pool in a given epoch.
     * @param epochId The ID of the epoch the reward pool was initialized in.
     * @param stakee The address of the Node.
     * @return The total active stake for that reward pool in SOLO.
     */
    function getRewardPoolActiveStake(uint256 epochId, address stakee) external view returns (uint256) {
        return rewardPools[getRewardPoolKey(epochId, stakee)].totalActiveStake;
    }

    /**
     * @notice Retrieve the total unclaimed reward allocated to a Node as payment
     * for providing a service.
     * @param stakee The address of the Node.
     * @return The total unclaimed Node reward in SOLO.
     */
    function getUnclaimedNodeReward(address stakee) external view returns (uint256) {
        return unclaimedNodeRewards[stakee];
    }

    /**
     * @notice Retrieve the total unclaimed staking reward allocated to a Node's
     * delegated stakers.
     * @param stakee The address of the Node.
     * @return The total unclaimed staking reward in SOLO.
     */
    function getUnclaimedStakeReward(address stakee) external view returns (uint256) {
        return unclaimedStakeRewards[stakee];
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

        nextRewardPool.initializedAt = block.number;

        // Any unclaimed staker rewards will automatically be added to the
        // active stake total
        nextRewardPool.totalActiveStake = totalStake + unclaimedStakeRewards[stakee];

        nextRewardPool.initialCumulativeRewardFactor = rewardPools[getRewardPoolKey(
            latestActiveRewardPools[stakee],
            stakee
        )].cumulativeRewardFactor;

        latestActiveRewardPools[stakee] = nextEpochId;
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
    function incrementRewardPool(
        address stakee,
        uint256 amount
    ) external onlyManager {
        EpochsManager.Epoch memory currentEpoch = _epochsManager.getCurrentActiveEpoch();

        RewardPool storage rewardPool = rewardPools[getRewardPoolKey(currentEpoch.iteration, stakee)];
        require(
            rewardPool.totalActiveStake > 0,
            "Reward pool has not been initialized for the current epoch"
        );

        uint256 stakersReward = SyloUtils.percOf(
            uint128(amount),
            currentEpoch.defaultPayoutPercentage
        );

        // update the value of the reward owed to the node
        unclaimedNodeRewards[stakee] += (amount - stakersReward);

        // update the value of the reward owed to the delegated stakers
        unclaimedStakeRewards[stakee] += stakersReward;

        rewardPool.stakersRewardTotal += stakersReward;

        // if this is the first epoch the node is ever active
        // then we can't rely on the previous crf to calculate the current crf
        if (rewardPool.initialCumulativeRewardFactor == 0) {
            rewardPool.cumulativeRewardFactor =
                ABDKMath64x64.div(
                    toFixedPointSYLO(rewardPool.stakersRewardTotal),
                    toFixedPointSYLO(rewardPool.totalActiveStake)
                );
        } else {
            rewardPool.cumulativeRewardFactor = calculatateUpdatedCumulativeRewardFactor(
                rewardPool.initialCumulativeRewardFactor,
                rewardPool.stakersRewardTotal,
                rewardPool.totalActiveStake
            );
        }
    }

    function calculatateUpdatedCumulativeRewardFactor(
        int128 previousCumulativeRewardFactor,
        uint256 rewardTotal,
        uint256 stakeTotal
    ) internal pure returns (int128) {
        return ABDKMath64x64.add(
            previousCumulativeRewardFactor,
            ABDKMath64x64.mul(
                previousCumulativeRewardFactor,
                ABDKMath64x64.div(
                    toFixedPointSYLO(rewardTotal),
                    toFixedPointSYLO(stakeTotal)
                )
            )
        );
    }

    /**
     * @notice Call this function to calculate the total portion of staking reward
     * that a delegated staker is owed. This value will include all epochs since the
     * last claim was made.
     * @dev This function will utilize the cumulative reward factor to perform the
     * calculation, keeping the gas cost scaling of this function to a constant value.
     * @param stakee The address of the Node.
     * @param staker The address of the staker.
     * @return The value of the reward owed to the staker in SOLO.
     */
    function calculateStakerClaim(address stakee, address staker) public view returns (uint256) {
        // The staking manager will track the initial stake that was available prior
        // to becoming active
        StakingManager.StakeEntry memory stakeEntry = _stakingManager.getStakeEntry(stakee, staker);
        if (stakeEntry.amount == 0) {
            return 0;
        }

        uint256 claim = calculateInitialClaim(stakee, staker);

        // find the first reward pool where their stake was active and had
        // generated rewards
        uint256 activeAt = 0;
        for (uint i = lastClaims[getStakerKey(stakee, staker)].claimedAt + 1; i < _epochsManager.currentIteration(); i++) {
            RewardPool storage rewardPool = rewardPools[getRewardPoolKey(i, stakee)];
            // check if node initialized a reward pool for this epoch and
            // gained rewards
            if (rewardPool.initializedAt > 0 && rewardPool.stakersRewardTotal > 0) {
                activeAt = i;
                break;
            }
        }

        if (activeAt == 0) {
            return claim;
        }

        RewardPool storage initialActivePool = rewardPools[getRewardPoolKey(activeAt, stakee)];

        // We convert the staker amount to SYLO as the maximum uint256 value that
        // can be used for the fixed point representation is 2^64-1.
        int128 initialStake = toFixedPointSYLO(stakeEntry.amount);
        int128 initialCumulativeRewardFactor = initialActivePool.initialCumulativeRewardFactor;

        // if the staker started staking prior to the node generating any
        // rewards (initial crf == 0), then we have to manually calculate the proportion of reward
        // for the first epoch, and use that value as the initial stake instead
        if (initialCumulativeRewardFactor == int128(0)) {
            initialStake = ABDKMath64x64.add(
                initialStake,
                ABDKMath64x64.mul(
                    toFixedPointSYLO(initialActivePool.stakersRewardTotal),
                    ABDKMath64x64.div(
                        initialStake,
                        toFixedPointSYLO(initialActivePool.totalActiveStake)
                    )
                )
            );
            initialCumulativeRewardFactor = initialActivePool.cumulativeRewardFactor;
        }

        RewardPool storage latestRewardPool = rewardPools[getRewardPoolKey(
            latestActiveRewardPools[stakee], stakee
        )];

        int128 finalCumulativeRewardFactor = getFinalCumulativeRewardFactor(stakee, staker);

        // Utilize the cumulative reward factor to calculate their updated stake amount.
        // We only allow claims from up to the previous epoch, so use the initial crf
        // for calculations.
        uint256 updatedStake = fromFixedPointSYLO(
            ABDKMath64x64.mul(
                initialStake,
                ABDKMath64x64.div(
                    finalCumulativeRewardFactor,
                    initialCumulativeRewardFactor
                )
            )
        );

        // this is the actual amount of rewards generated by their stake
        // since their stake became active
        return claim + updatedStake - stakeEntry.amount;
    }

    /**
     * Manually calculates the reward claim for the first epoch the claim is being
     * made for. This manual calculation is necessary as claims are only made up
     * to the previous epoch.
     */
    function calculateInitialClaim(address stakee, address staker) internal view returns (uint256) {
        LastClaim storage lastClaim = lastClaims[getStakerKey(stakee, staker)];
        RewardPool storage firstRewardPool = rewardPools[getRewardPoolKey(lastClaim.claimedAt, stakee)];

        if (firstRewardPool.totalActiveStake == 0) {
            return 0;
        }

        return firstRewardPool.stakersRewardTotal * lastClaim.stake / firstRewardPool.totalActiveStake;
    }

    /**
     * Determines the cumulative reward factor to use for claim calculations. The
     * CRF will depend when the Node last initialized a reward pool, and also when
     * the staker last made their claim.
     */
    function getFinalCumulativeRewardFactor (address stakee, address staker) internal view returns (int128) {
        uint256 currentEpoch = _epochsManager.currentIteration();

        RewardPool storage latestRewardPool = rewardPools[getRewardPoolKey(
            latestActiveRewardPools[stakee], stakee
        )];

        int128 finalCumulativeRewardFactor = 0;

        // Get the cumulative reward factor for the Node
        // for the start of this epoch, since we only perform
        // calculations up to the end of the previous epoch.
        if (latestActiveRewardPools[stakee] < _epochsManager.currentIteration()) {
            // If the Node has not been active, then the final
            // cumulative reward factor will just be the current one.
            finalCumulativeRewardFactor = latestRewardPool.cumulativeRewardFactor;
        } else if (latestActiveRewardPools[stakee] == currentEpoch) {
            // We are calculating the claim for an active epoch, the
            // final cumulative reward factor will be taken from the start of this
            // epoch (end of previous epoch).
            finalCumulativeRewardFactor = latestRewardPool.initialCumulativeRewardFactor;
        } else {
            // The Node has already initialized the next reward pool.

            // First check if a reward pool was initialized for the current
            // epoch.
            RewardPool storage currentRewardPool = rewardPools[getRewardPoolKey(
                currentEpoch, stakee
            )];

            if (currentRewardPool.initialCumulativeRewardFactor > 0) {
                finalCumulativeRewardFactor = currentRewardPool.initialCumulativeRewardFactor;
            } else {
                // The node did not initialize a reward pool for the current epoch,
                // so the final crf will just be the current crf
                finalCumulativeRewardFactor = latestRewardPool.initialCumulativeRewardFactor;
            }
        }

        return finalCumulativeRewardFactor;
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
     * @notice Call this function to claim rewards as a delegated staker. The
     * SYLO tokens will be transferred to the caller's account. This function will
     * fail if there exists no reward to claim. Note: Calling this will remove
     * the current unclaimed reward from being used as stake in the next round.
     * @param stakee The address of the Node to claim against.
     */
    function claimStakingRewards(address stakee) external returns (uint256) {
        uint256 rewardClaim = calculateStakerClaim(stakee, msg.sender);
        require(rewardClaim > 0, "Nothing to claim");
        unclaimedStakeRewards[stakee] -= rewardClaim;
        updateLastClaim(stakee, msg.sender, rewardClaim);
        _token.transfer(msg.sender, rewardClaim);
        return rewardClaim;
    }

    /**
     * @notice This function should be called to automatically claim rewards
     * when a staker wishes to update their stake. This is only callable
     * by the StakingManager contract.
     * @dev This function will revert if the StakingManager contract has
     * not been set as a manager.
     * @param stakee The address of the Node to claim against.
     * @param staker The address of the staker.
     * @param payee The address of the account receiving the reward.
     * This will be either the StakingManager account when adding stake,
     * or the staker's account when withdrawing stake.
     */
    function claimStakingRewardsAsManager(address stakee, address staker, address payee) external onlyManager returns (uint256) {
        uint256 rewardClaim = calculateStakerClaim(stakee, staker);
        updateLastClaim(stakee, staker, rewardClaim);
        if (rewardClaim == 0) {
            return rewardClaim;
        }
        unclaimedStakeRewards[stakee] -= rewardClaim;
        _token.transfer(payee, rewardClaim);

        return rewardClaim;
    }

    function updateLastClaim(address stakee, address staker, uint256 rewardClaim) internal {
        StakingManager.StakeEntry memory stakeEntry = _stakingManager.getStakeEntry(stakee, staker);
        lastClaims[getStakerKey(stakee, staker)] = LastClaim(
            _epochsManager.currentIteration(),
            stakeEntry.amount + rewardClaim
        );
    }

    /**
     * @notice Call this function as a Node operator to claim the accumulated
     * reward for operating a Sylo Node.
     */
    function claimNodeRewards() external {
        uint256 claim = unclaimedNodeRewards[msg.sender];
        require(claim > 0, "Nothing to claim");

        unclaimedNodeRewards[msg.sender] = 0;
        _token.transfer(msg.sender, claim);
    }
}