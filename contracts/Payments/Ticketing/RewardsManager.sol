// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "../../Staking/Manager.sol";
import "../../Epochs/Manager.sol";
import "../../Utils.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

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

    struct Stake {
        address staker;
        uint256 amount;
    }

    struct RewardPool {
        // Tracks the balance of the reward pool as tickets are incremented
        uint256 balance;

        // Tracks the delegated stakers and their proportions of stake
        Stake[] stakes;

        uint256 totalStake;
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

    function getKey(bytes32 epochId, address stakee) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(epochId, stakee));
    }

    function getRewardBalance(bytes32 epochId, address stakee) public view returns (uint256) {
        return rewardPools[getKey(epochId, stakee)].balance;
    }

    function getRewardPoolStake(bytes32 epochId, address stakee) public view returns (uint256) {
        return rewardPools[getKey(epochId, stakee)].balance;
    }

    function getDelegatorOwedAmount(
        bytes32 epochId,
        address stakee,
        address staker
    ) public view returns (uint256) {
        EpochsManager.Epoch memory epoch = _epochsManager.getEpoch(epochId);
        require(epoch.startBlock > 0, "Epoch does not exist");

        RewardPool memory rewardPool = rewardPools[getKey(epochId, stakee)];

        if (rewardPool.balance == 0) {
            return 0;
        }

        uint256 stake = 0;
        for (uint i = 0; i < rewardPool.stakes.length; i++) {
            if (rewardPool.stakes[i].staker == staker) {
                stake = rewardPool.stakes[i].amount;
                break;
            }
        }

        if (stake == 0) {
            return 0;
        }

        uint256 delegatorReward = SyloUtils.percOf(
            uint128(rewardPool.balance),
            epoch.defaultPayoutPercentage
        );

        return calculateDelegatorPayout(stake, rewardPool.totalStake, delegatorReward);
    }

    function calculateDelegatorPayout(
        uint256 stake,
        uint256 totalStake,
        uint256 delegatorReward
    ) public pure returns (uint256) {
        // we calculate the payout for this staker by taking their
        // proportion of stake against the total stake, and multiplying
        // that against the total reward for the stakers
        return stake * delegatorReward / totalStake;
    }

    function incrementRewardPool(
        bytes32 epochId,
        address stakee,
        uint256 amount
    ) public onlyManager {
        EpochsManager.Epoch memory epoch = _epochsManager.getEpoch(epochId);
        require(epoch.startBlock > 0, "Epoch does not exist");

        RewardPool storage rewardPool = rewardPools[getKey(epochId, stakee)];
        require(
            rewardPool.totalStake > 0,
            "Reward pool has not been constructed for this epoch"
        );

        _token.transferFrom(msg.sender, address(this), amount);

        rewardPool.balance += amount;
    }

    /*
     * This function should be called by the node in order to distribute
     * the reward accumulated over the epoch.
     */
    function distributeReward(bytes32 epochId) public {
        EpochsManager.Epoch memory epoch = _epochsManager.getEpoch(epochId);
        require(epoch.startBlock > 0, "Epoch does not exist");
        require(epoch.endBlock > 0, "Epoch has not yet ended");
        require(
            block.number > epoch.endBlock + epoch.ticketDuration,
            "Can only distribute rewards once all possible tickets have been redeemed"
            " (epoch.endBlock + epoch.ticketDuration)"
        );

        bytes32 key = getKey(epochId, msg.sender);

        RewardPool memory rewardPool = rewardPools[key];
        require(
            rewardPool.balance > 0,
            "Can not distribute reward if balance is zero"
        );

        uint256 delegatorReward = SyloUtils.percOf(
            uint128(rewardPool.balance),
            epoch.defaultPayoutPercentage
        );

        uint256 totalDelegatorPayout = 0;

        // Iterate through each delegator and pay them accordingly
        for (uint i = 0; i < rewardPool.stakes.length; i++) {
            Stake memory stake = rewardPool.stakes[i];

            uint256 payout = calculateDelegatorPayout(stake.amount, rewardPool.totalStake, delegatorReward);

            // Avoid reverting if the payout is zero, which could
            // occur if the stake is too low relative to the other stakes
            if (payout == 0) {
                continue;
            }

            _token.transfer(stake.staker, payout);
            totalDelegatorPayout += payout;
        }

        // The node is paid out the remaining after paying out the delegated stakers
        uint256 stakeePayout = rewardPool.balance - totalDelegatorPayout;

        _token.transfer(msg.sender, stakeePayout);

        // The rewards have been reconciled, we can clear the storage for this pool
        delete rewardPools[key];
    }

    /*
     * This function is called by the node to initialize their own reward pool
     * for the next epoch. Calling this function will be necessary for the node
     * to participate in the next epoch.
     */
    function initializeRewardPool(bytes32 epochId) public {
        bytes32 key = getKey(epochId, msg.sender);

        RewardPool storage rewardPool = rewardPools[key];
        require(rewardPool.totalStake > 0, "Reward pool has already been initialized");

        EpochsManager.Epoch memory epoch = _epochsManager.getEpoch(epochId);
        require(epoch.endBlock == 0, "Epoch has already ended");

        address[] memory stakers = _stakingManager.getStakers(msg.sender);
        require(stakers.length > 0, "Must have stake to intitialize a reward pool");

        // record all stakers for this stakee at the time this reward pool
        // is initialized
        for (uint i = 0; i < stakers.length; i++) {
            uint256 amount = _stakingManager.getStake(stakers[i], msg.sender).amount;
            rewardPool.stakes.push(Stake(
                stakers[i],
                amount
            ));
        }

        rewardPool.totalStake = _stakingManager.totalStakes(msg.sender);
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