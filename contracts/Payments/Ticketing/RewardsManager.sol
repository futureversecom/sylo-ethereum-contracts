// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "../../Staking/Manager.sol";
import "../../Epochs/Manager.sol";
import "../../Utils.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "../../../node_modules/abdk-libraries-solidity/ABDKMath64x64.sol";

/*
 * Handles epoch based reward pools that are incremented from redeeming tickets.
 * Nodes use this contract to set up their reward pool for the next epoch,
 * and also to payout delegated stakers after the epoch ends.
 * After deployment, the SyloTicketing contract should be
 * set up as a manager to be able to call certain restricted functions.
*/

contract RewardsManager is Initializable, OwnableUpgradeable {
    /* ERC 20 compatible token we are dealing with */
    IERC20 _token;

    /* Sylo Staking Manager contract. */
    StakingManager _stakingManager;

    /* Sylo Epochs Manager. */
    EpochsManager _epochsManager;

    mapping (address => uint256) public unclaimedNodeRewards;

    mapping (address => uint256) public unclaimedStakeRewards;

    /* For every node, track their most recently initialized reward pool */
    mapping (address => uint256) public latestActiveRewardPools;

    /* For every delegated staker a node has, track the last epoch they made a claim in */
    mapping (bytes32 => uint256) public lastClaims;

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

        // track the cumulative reward factor as a quadruple precision floating point value
        int128 cumulativeRewardFactor;
    }

    // Reward Pools are indexed by a key that is derived from the epochId and the stakee's address
    mapping (bytes32 => RewardPool) rewardPools;

    // Certain functions of this contract should only be called by certain other
    // contracts, namely the Ticketing contract.
    // We use this mapping to restrict access to those functions in a similar
    // fashion to the onlyOwner construct. The stored value is the block the
    // managing was contract was added in.
    mapping (address => uint256) managers;

    function initialize(
        IERC20 token,
        StakingManager stakingManager,
        EpochsManager epochsManager
    ) public initializer {
        OwnableUpgradeable.__Ownable_init();
        _token = token;
        _epochsManager = epochsManager;
        _stakingManager = stakingManager;
    }

    function getKey(uint256 epochId, address stakee) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(epochId, stakee));
    }

    function getStakerKey(address stakee, address staker) public pure returns(bytes32) {
        return keccak256(abi.encodePacked(stakee, staker));
    }

    function getLastClaim(address stakee, address staker) public view returns(uint256) {
        return lastClaims[getStakerKey(stakee, staker)];
    }

    function getRewardPool(uint256 epochId, address stakee) public view returns (RewardPool memory) {
        return rewardPools[getKey(epochId, stakee)];
    }

    function getRewardPoolStakersTotal(uint256 epochId, address stakee) public view returns (uint256) {
        return rewardPools[getKey(epochId, stakee)].stakersRewardTotal;
    }

    function getRewardPoolActiveStake(uint256 epochId, address stakee) public view returns (uint256) {
        return rewardPools[getKey(epochId, stakee)].totalActiveStake;
    }

    function getUnclaimedNodeReward(address stakee) public view returns (uint256) {
        return unclaimedNodeRewards[stakee];
    }

    function getUnclaimedStakeReward(address stakee) public view returns (uint256) {
        return unclaimedStakeRewards[stakee];
    }

    function initializeNextRewardPool() public {
        uint256 nextEpochId = _epochsManager.getNextEpochId();

        RewardPool storage nextRewardPool = rewardPools[getKey(nextEpochId, msg.sender)];
        require(
            nextRewardPool.initializedAt == 0,
            "The next reward pool has already been initialized"
        );

        uint256 totalStake = _stakingManager.getStakeeTotalManagedStake(msg.sender);
        require(totalStake > 0, "Must have stake to intitialize a reward pool");

        nextRewardPool.initializedAt = block.number;

        // Any unclaimed staker rewards will automatically be added to the
        // active stake total
        nextRewardPool.totalActiveStake = totalStake + unclaimedStakeRewards[msg.sender];

        nextRewardPool.initialCumulativeRewardFactor = rewardPools[getKey(
            latestActiveRewardPools[msg.sender],
            msg.sender
        )].cumulativeRewardFactor;

        latestActiveRewardPools[msg.sender] = nextEpochId;
    }

    /*
     * This function should be called by the Ticketing contract when a
     * ticket is successfully redeemed. The face value of the ticket
     * should be split between incrementing the node's reward balance,
     * and the reward balance for the node's delegated stakers.
     * Additionally, the cumulative reward factor will be updated, and in the
     * case the ticket was redeemed for an epoch that has already ended, any proceeding
     * cumulative reward factors will also be updated. This has the consequence of
     * the gas cost for redeeming to increase if it is redeemed much later than the epoch
     * it was generated in.
     */
    function incrementRewardPool(
        address stakee,
        uint256 amount
    ) public onlyManager {
        EpochsManager.Epoch memory currentEpoch = _epochsManager.getCurrentActiveEpoch();

        RewardPool storage rewardPool = rewardPools[getKey(currentEpoch.iteration, stakee)];
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
                    ABDKMath64x64.fromUInt(rewardPool.stakersRewardTotal),
                    ABDKMath64x64.fromUInt(rewardPool.totalActiveStake)
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
                    ABDKMath64x64.fromUInt(rewardTotal),
                    ABDKMath64x64.fromUInt(stakeTotal)
                )
            )
        );
    }

    function calculateStakerClaim(address stakee, address staker) public view returns (uint256) {
        // The staking manager will track the initial stake that was available prior
        // to becoming active
        StakingManager.StakeEntry memory stakeEntry = _stakingManager.getStakeEntry(stakee, staker);
        if (stakeEntry.amount == 0) {
            return 0;
        }

        // find the reward pool when their stake became active,
        // which will be the first reward pool after their last claim
        uint256 activeAt = 0;
        for (uint i = lastClaims[getStakerKey(stakee, staker)] + 1; i < _epochsManager.getNextEpochId(); i++) {
            RewardPool storage rewardPool = rewardPools[getKey(i, stakee)];
            // check if node initialized a reward pool for this epoch
            if (rewardPool.initializedAt > 0) {
                activeAt = i;
                break;
            }
        }

        if (activeAt == 0) {
            return 0;
        }

        RewardPool storage initialActivePool = rewardPools[getKey(activeAt, stakee)];

        int128 initialStake = ABDKMath64x64.fromUInt(stakeEntry.amount);
        int128 initialCumulativeRewardFactor = initialActivePool.initialCumulativeRewardFactor;

        // if the staker started staking prior to the node generating any
        // rewards (initial crf == 0), then we have to manually calculate the proprotion of reward
        // for the first epoch, and use that value as the initial stake instead
        if (initialCumulativeRewardFactor == int128(0)) {
            initialStake = ABDKMath64x64.add(
                initialStake,
                ABDKMath64x64.mul(
                    ABDKMath64x64.fromUInt(initialActivePool.stakersRewardTotal),
                    ABDKMath64x64.div(
                        initialStake,
                        ABDKMath64x64.fromUInt(initialActivePool.totalActiveStake)
                    )
                )
            );
            initialCumulativeRewardFactor = initialActivePool.cumulativeRewardFactor;
        }

        RewardPool storage latestRewardPool = rewardPools[getKey(
            latestActiveRewardPools[stakee], stakee
        )];

        // utilize the cumulative reward factor to calculate their updated stake amount
        uint256 updatedStake = ABDKMath64x64.toUInt(
            ABDKMath64x64.mul(
                initialStake,
                ABDKMath64x64.div(
                    latestRewardPool.cumulativeRewardFactor,
                    initialCumulativeRewardFactor
                )
            )
        );

        // this is the actual amount of rewards generated by their stake
        // since their stake became active
        return updatedStake - stakeEntry.amount;
    }

    function claimStakingRewards(address stakee) public {
        uint256 rewardClaim = calculateStakerClaim(stakee, msg.sender);
        require(rewardClaim > 0, "Nothing to claim");
        unclaimedStakeRewards[stakee] -= rewardClaim;
        lastClaims[getStakerKey(stakee, msg.sender)] = latestActiveRewardPools[stakee];
        _token.transfer(msg.sender, rewardClaim);
    }

    /*
     * This function will generally be called by the staking manager to
     * automatically claim rewards for a staker when the staker wishes to
     * update their stake amount.
     */
    function claimStakingRewardsAsManager(address stakee, address staker) public onlyManager {
        uint256 rewardClaim = calculateStakerClaim(stakee, staker);
        lastClaims[getStakerKey(stakee, staker)] = latestActiveRewardPools[stakee];
        if (rewardClaim == 0) {
            return;
        }
        unclaimedStakeRewards[stakee] -= rewardClaim;
        _token.transfer(staker, rewardClaim);
    }

    function claimNodeRewards() public {
        uint256 claim = unclaimedNodeRewards[msg.sender];
        require(claim > 0, "Nothing to claim");

        unclaimedNodeRewards[msg.sender] = 0;
        _token.transfer(msg.sender, claim);
    }

    function addManager(address manager) public onlyOwner {
      managers[manager] = block.number;
    }

    function removeManager(address manager) public onlyOwner {
      delete managers[manager];
    }

    modifier onlyManager() {
      require(managers[msg.sender] > 0, "Only managers of this contract can call this function");
      _;
    }

}