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
        // Tracks the balance of the reward pool owed to the stakee
        uint256 stakeeRewardBalance;

        // Tracks the balance of the reward pool owed to the stakers
        uint256 stakersRewardBalance;

        // Tracks the block number this reward pool was initialized
        uint256 initializedAt;

        // Tracks any users that have claimed from this pool
        mapping (address => bool) claimed;

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

    function getRewardPoolStakeeBalance(bytes32 epochId, address stakee) public view returns (uint256) {
        return rewardPools[getKey(epochId, stakee)].stakeeRewardBalance;
    }

    function getRewardPoolStakersBalance(bytes32 epochId, address stakee) public view returns (uint256) {
        return rewardPools[getKey(epochId, stakee)].stakersRewardBalance;
    }

    function getRewardPoolBalance(bytes32 epochId, address stakee) public view returns (uint256) {
        return
            getRewardPoolStakeeBalance(epochId, stakee) +
            getRewardPoolStakersBalance(epochId, stakee);
    }

    function getRewardPoolStake(bytes32 epochId, address stakee) public view returns (uint256) {
        return rewardPools[getKey(epochId, stakee)].totalStake;
    }

    function getDelegatorOwedAmount(
        bytes32 epochId,
        address stakee,
        address staker
    ) public view returns (uint256) {
        EpochsManager.Epoch memory epoch = _epochsManager.getEpoch(epochId);
        require(epoch.startBlock > 0, "Epoch does not exist");

        RewardPool storage rewardPool = rewardPools[getKey(epochId, stakee)];

        if (rewardPool.stakersRewardBalance == 0) {
            return 0;
        }

        uint256 stake = _stakingManager.getStakerAmount(staker, stakee, rewardPool.initializedAt);
        if (stake == 0) {
            return 0;
        }

        return calculateDelegatorPayout(stake, rewardPool.totalStake, rewardPool.stakersRewardBalance);
    }

    function calculateDelegatorPayout(
        uint256 stake,
        uint256 totalStake,
        uint256 delegatorReward
    ) internal pure returns (uint256) {
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

        uint256 stakersReward = SyloUtils.percOf(
            uint128(amount),
            epoch.defaultPayoutPercentage
        );

        rewardPool.stakeeRewardBalance += (amount - stakersReward);
        rewardPool.stakersRewardBalance += stakersReward;
    }

    /*
     * This function is called by the node to initialize their own reward pool
     * for the next epoch. Calling this function will be necessary for the node
     * to participate in the next epoch.
     */
    function initializeRewardPool(bytes32 epochId) public {
        bytes32 key = getKey(epochId, msg.sender);

        RewardPool storage rewardPool = rewardPools[key];
        require(rewardPool.totalStake == 0, "Reward pool has already been initialized");

        EpochsManager.Epoch memory epoch = _epochsManager.getEpoch(epochId);
        require(epoch.endBlock == 0, "Epoch has already ended");

        uint256 totalStake = _stakingManager.getStakeeTotalStake(msg.sender);
        require(totalStake > 0, "Must have stake to intitialize a reward pool");

        rewardPool.initializedAt = block.number;
        rewardPool.totalStake = totalStake;
    }

    /*
     * This function should be called by the node or staker in order to claim
     * their rewards accumalted over the specified epochs.
     */
    function claimRewards(bytes32[] memory epochIds, address stakee) public {
        uint256 totalPayout = 0;
        for (uint i = 0; i < epochIds.length; i++) {
            uint256 payout = calculateClaim(epochIds[i], stakee);
            if (payout > 0) {
                bytes32 key = getKey(epochIds[i], stakee);
                rewardPools[key].claimed[msg.sender] = true;
            }
            totalPayout += payout;
        }

        _token.transfer(msg.sender, totalPayout);
    }

    function calculateClaim(bytes32 epochId, address stakee) public view returns (uint256) {
        EpochsManager.Epoch memory epoch = _epochsManager.getEpoch(epochId);
        require(epoch.startBlock > 0, "Epoch does not exist");
        require(epoch.endBlock > 0, "Epoch has not yet ended");
        require(
            block.number > epoch.endBlock + epoch.ticketDuration,
            "Can only claim rewards once all possible tickets have been redeemed"
            " (epoch.endBlock + epoch.ticketDuration)"
        );

        bytes32 key = getKey(epochId, stakee);

        RewardPool storage rewardPool = rewardPools[key];
        require(
            rewardPool.stakeeRewardBalance + rewardPool.stakersRewardBalance > 0,
            "Can not claim reward if balance is zero"
        );
        require(
            rewardPool.claimed[msg.sender] == false,
            "Can not claim balance more than once"
        );

        uint256 stake = _stakingManager.getStakerAmount(msg.sender, stakee, rewardPool.initializedAt);
        require(
            stake > 0 || msg.sender == stakee,
            "Must have had stake for this epoch or be the stakee in order to claim reward"
        );

        uint256 payout = calculateDelegatorPayout(stake, rewardPool.totalStake, rewardPool.stakersRewardBalance);

        if (msg.sender == stakee) {
            payout += rewardPool.stakeeRewardBalance;
        }

        return payout;
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