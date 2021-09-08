// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "../../Utils.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

/*
 * Handles rewards from redeeming tickets. After deployment,
 * the SyloTicketing and StakingManager contracts should be set as managers
 * to be able to call certain restricted functions.
*/

contract RewardPool is Initializable, OwnableUpgradeable {
    /* ERC 20 compatible token we are dealing with */
    IERC20 _token;

    /* Mapping of ticket hashes, used to check if ticket has been redeemed */
    mapping (bytes32 => bool) public usedTickets;

    struct Pool {
        // Tracks outstanding reward balance for the stakee. This value
        // is incremented as tickets are redeemed, and is entirely cleared
        // when the stakee claims their reward
        uint256 outstandingStakeeReward;

        // Tracks the total accumulated reward over time that is designated
        // for delegated stakers. This value is incremented as tickets are redeemed,
        // but is NOT decrement when delegators claim their reward. Instead, the most
        // totalDelegatorsReward value is recorded whenever a delegator claims their reward,
        // (trakced below in delegatorClaims map), and calculating the outstanding reward
        // designated for a specific delegators is done by: totalDelegatorsReward - delegatorClaims[msg.sender].
        // This design allows us to optimize gas costs when distributing rewards.
        uint256 totalDelegatorsReward;

        mapping (address => uint256) delegatorClaims;
    }

    mapping (address => Pool) public rewardPools;

    // Certain functions of this contract should only be called by certain other
    // contracts, namely the Ticketing contract and the Staking Manager contract.
    // We use this mapping to restrict access to those functions in a similar
    // fashion to the onlyOwner construct. The uint256 is the block the
    // managing was contract was added in.
    mapping (address => uint256) managers;

    function initialize(
        IERC20 token
    ) public initializer {
        OwnableUpgradeable.__Ownable_init();
        _token = token;
    }

    function getOutstandingStakeeReward(address stakee) public view returns (uint256) {
        return rewardPools[stakee].outstandingStakeeReward;
    }

    function getTotalDelegatorsReward(address stakee) public view returns (uint256) {
        return rewardPools[stakee].totalDelegatorsReward;
    }

    function getDelegatorClaimAmount(
      address stakee,
      address staker,
      uint256 delegatedStake,
      uint256 totalStake
    ) public view returns (uint256) {
        Pool storage rewardPool = rewardPools[stakee];

        // Calculate the amount of reward that has been accumulated since the last time
        // this sender claimed their reward
        uint256 accumulatedReward = rewardPool.totalDelegatorsReward - rewardPool.delegatorClaims[staker];

        if (accumulatedReward == 0) {
            return 0;
        }

        require(
            delegatedStake > 0,
            "Must be a delegated staker or the stakee to claim rewards"
        );

        // we calculate the payout for this staker by taking their
        // proportion of stake against the total stake, and multiplying
        // that against the total reward for the stakers
        uint256 payout = delegatedStake * accumulatedReward / totalStake;

        return payout;
    }

    function incrementRewardPool(
        uint16 defaultPayoutPercentage,
        address stakee,
        uint256 amount
    ) public onlyManager {
        _token.transferFrom(msg.sender, address(this), amount);

        uint256 delegatorReward = SyloUtils.percOf(
            uint128(amount),
            defaultPayoutPercentage
        );

        rewardPools[stakee].outstandingStakeeReward += amount - delegatorReward;
        rewardPools[stakee].totalDelegatorsReward += delegatorReward;
    }

    function transferDelegatedStakerReward(
      address stakee,
      address rewardee,
      uint256 delegatedStake,
      uint256 totalStake
    ) public onlyManager {
        Pool storage rewardPool = rewardPools[stakee];

        uint256 reward = getDelegatorClaimAmount(stakee, rewardee, delegatedStake, totalStake);

        require(reward > 0, "Accumalated reward is 0");

        rewardPool.delegatorClaims[rewardee] = rewardPool.totalDelegatorsReward;

        _token.transfer(rewardee, reward);
    }

    function transferReward(
      address stakee,
      address rewardee,
      uint256 delegatedStake,
      uint256 totalStake
    ) public onlyManager {
        Pool storage rewardPool = rewardPools[stakee];

        uint256 reward = getDelegatorClaimAmount(stakee, rewardee, delegatedStake, totalStake);

        if (reward > 0) {
            rewardPool.delegatorClaims[msg.sender] = rewardPool.totalDelegatorsReward;
        }

        if (rewardee == stakee) {
          reward += rewardPool.outstandingStakeeReward;
          rewardPool.outstandingStakeeReward = 0;
        }

        _token.transfer(rewardee, reward);
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