// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "../../Staking/Manager.sol";
import "../../Epochs/Manager.sol";
import "../../Utils.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "../../../node_modules/abdk-libraries-solidity/ABDKMathQuad.sol";

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
        bytes16 initialCumulativeRewardFactor;

        // track the cumulative reward factor as a quadruple precision floating point value
        bytes16 cumulativeRewardFactor;
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
        uint256 epochId,
        address stakee,
        uint256 amount
    ) public onlyManager {
        EpochsManager.Epoch memory epoch = _epochsManager.getEpoch(epochId);
        require(epoch.startBlock > 0, "Epoch does not exist");

        RewardPool storage rewardPool = rewardPools[getKey(epochId, stakee)];
        require(
            rewardPool.totalActiveStake > 0,
            "Reward pool has not been initialized for this epoch"
        );

        uint256 stakersReward = SyloUtils.percOf(
            uint128(amount),
            epoch.defaultPayoutPercentage
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
                ABDKMathQuad.div(
                    ABDKMathQuad.fromUInt(rewardPool.stakersRewardTotal),
                    ABDKMathQuad.fromUInt(rewardPool.totalActiveStake)
                );
        } else {
            rewardPool.cumulativeRewardFactor = calculatateUpdatedCumulativeRewardFactor(
                rewardPool.initialCumulativeRewardFactor,
                rewardPool.stakersRewardTotal,
                rewardPool.totalActiveStake
            );
        }

        // in the case that the ticket was redeemed for a historical epoch,
        // we must update all proceeding cumulative reward factors
        uint256 latestIteration = latestActiveRewardPools[stakee];
        if (epochId < latestIteration) {
            bytes16 previousCumulativeRewardFactor = rewardPool.cumulativeRewardFactor;
            for (uint i = epochId + 1; i <= latestIteration; i++) {
                RewardPool storage next = rewardPools[getKey(i, stakee)];
                if (next.initializedAt > 0) {
                    bytes16 nextCumulativeRewardFactor = calculatateUpdatedCumulativeRewardFactor(
                        previousCumulativeRewardFactor,
                        next.stakersRewardTotal,
                        next.totalActiveStake
                    );
                    next.initialCumulativeRewardFactor = previousCumulativeRewardFactor;
                    next.cumulativeRewardFactor = nextCumulativeRewardFactor;
                    previousCumulativeRewardFactor = nextCumulativeRewardFactor;
                }
            }
        }
    }

    function calculatateUpdatedCumulativeRewardFactor(
        bytes16 previousCumulativeRewardFactor,
        uint256 rewardTotal,
        uint256 stakeTotal
    ) internal pure returns (bytes16) {
        return ABDKMathQuad.add(
            previousCumulativeRewardFactor,
            ABDKMathQuad.mul(
                previousCumulativeRewardFactor,
                ABDKMathQuad.div(
                    ABDKMathQuad.fromUInt(rewardTotal),
                    ABDKMathQuad.fromUInt(stakeTotal)
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

        bytes16 initialStake = ABDKMathQuad.fromUInt(stakeEntry.amount);
        bytes16 initialCumulativeRewardFactor = initialActivePool.initialCumulativeRewardFactor;

        // if the staker started staking prior to the node generating any
        // rewards (initial crf == 0), then we have to manually calculate the proprotion of reward
        // for the first epoch, and use that value as the initial stake instead
        if (initialCumulativeRewardFactor == bytes16(0)) {
            initialStake = ABDKMathQuad.add(
                initialStake,
                ABDKMathQuad.mul(
                    ABDKMathQuad.fromUInt(initialActivePool.stakersRewardTotal),
                    ABDKMathQuad.div(
                        initialStake,
                        ABDKMathQuad.fromUInt(initialActivePool.totalActiveStake)
                    )
                )
            );
            initialCumulativeRewardFactor = initialActivePool.cumulativeRewardFactor;
        }

        RewardPool storage latestRewardPool = rewardPools[getKey(
            latestActiveRewardPools[stakee], stakee
        )];

        // utilize the cumulative reward factor to calculate their updated stake amount
        uint256 updatedStake = ABDKMathQuad.toUInt(
            ABDKMathQuad.mul(
                initialStake,
                ABDKMathQuad.div(
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
      require(managers[msg.sender] > 0, "Only controllers of this contract can call this function");
      _;
    }

}