// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.13;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "../Token.sol";
import "../Payments/Ticketing/RewardsManager.sol";
import "../Epochs/Manager.sol";
import "../Utils.sol";

/**
 * @notice Manages stakes and delegated stakes for Nodes. Holding
 * staked Sylo is necessary for a Node to participate in the
 * Sylo Network. The stake is used in stake-weighted scan function,
 * and delegated stakers are rewarded on a pro-rata basis.
 */
contract StakingManager is Initializable, OwnableUpgradeable {
    /** ERC 20 compatible token we are dealing with */
    IERC20 public _token;

    /**
     * @notice Rewards Manager contract. Any changes to stake will automatically
     * trigger a claim to any outstanding rewards.
     */
    RewardsManager public _rewardsManager;

    EpochsManager public _epochsManager;

    /**
     * For every Node, there will be a mapping of the staker to a
     * StakeEntry. The stake entry tracks the amount of stake in SOLO,
     * and also when the stake was updated.
     */
    struct StakeEntry {
        uint256 amount;
        // Block number this entry was updated at
        uint256 updatedAt;
        // Epoch this entry was updated. The stake will become active
        // in the following epoch
        uint256 epochId;
    }

    /**
     * Every Node must have stake in order to participate in the Epoch.
     * Stake can be provided by the Node itself or by other accounts in
     * the network.
     */
    struct Stake {
        // Track each stake entry associated to a node
        mapping(address => StakeEntry) stakeEntries;
        // The total stake held by this contract for a node,
        // which will be the sum of all addStake and unlockStake calls
        uint256 totalManagedStake;
    }

    /**
     * This struct will track stake that is in the process of unlocking.
     */
    struct Unlock {
        uint256 amount; // Amount of stake unlocking
        uint256 unlockAt; // Block number the stake becomes withdrawable
    }

    /**
     * @notice Tracks the managed stake for every Node.
     */
    mapping(address => Stake) public stakes;

    /** @notice Tracks overall total stake held by this contract */
    uint256 public totalManagedStake;

    /**
     * @notice Tracks funds that are in the process of being unlocked. This
     * is indexed by a key that hashes both the address of the staked Node and
     * the address of the staker.
     */
    mapping(bytes32 => Unlock) public unlockings;

    event UnlockDurationUpdated(uint256 unlockDuration);
    event MinimumStakeProportionUpdated(uint256 minimumStakeProportion);

    /**
     * @notice The number of blocks a user must wait after calling "unlock"
     * before they can withdraw their stake
     */
    uint256 public unlockDuration;

    /**
     * @notice Minimum amount of stake that a Node needs to stake
     * against itself in order to participate in the network. This is
     * represented as a percentage of the Node's total stake, where
     * the value is a ratio of 10000.
     */
    uint16 public minimumStakeProportion;

    function initialize(
        IERC20 token,
        RewardsManager rewardsManager,
        EpochsManager epochsManager,
        uint256 _unlockDuration,
        uint16 _minimumStakeProportion
    ) external initializer {
        OwnableUpgradeable.__Ownable_init();
        _token = token;
        _rewardsManager = rewardsManager;
        _epochsManager = epochsManager;
        unlockDuration = _unlockDuration;
        minimumStakeProportion = _minimumStakeProportion;
    }

    /**
     * @notice Sets the unlock duration for stakes. Only callable by
     * the owner.
     * @param _unlockDuration The unlock duration in number of blocks.
     */
    function setUnlockDuration(uint256 _unlockDuration) external onlyOwner {
        unlockDuration = _unlockDuration;
        emit UnlockDurationUpdated(_unlockDuration);
    }

    /**
     * @notice Sets the minimum stake proportion for Nodes. Only callable by
     * the owner.
     * @param _minimumStakeProportion The minimum stake proportion in SOLO.
     */
    function setMinimumStakeProportion(uint16 _minimumStakeProportion) external onlyOwner {
        minimumStakeProportion = _minimumStakeProportion;
        emit MinimumStakeProportionUpdated(_minimumStakeProportion);
    }

    /**
     * @notice Called by Nodes and delegated stakers to add stake. Calling
     * this will trigger an automatic transfer of any outstanding staking
     * rewards to the total managed staker. This function will fail under
     * the following conditions:
     *   - If the Node address is invalid
     *   - If the specified stake value is zero
     *   - If the additional stake causes the Node to fail to meet the
     *     minimum stake proportion requirement.
     * @param amount The amount of stake to add in SOLO.
     * @param stakee The address of the staked Node.
     */
    function addStake(uint256 amount, address stakee) external {
        addStake_(amount, stakee);
        _token.transferFrom(msg.sender, address(this), amount);
    }

    function addStake_(uint256 amount, address stakee) internal {
        require(stakee != address(0), "Address is null");

        // automatically claim any outstanding rewards generated by their existing stake
        uint256 reward = _rewardsManager.claimStakingRewardsAsManager(
            stakee,
            msg.sender,
            address(this)
        );

        uint256 totalToStake = amount + reward;
        require(totalToStake != 0, "Stake amount or reward must be more than 0");

        uint256 currentEpochId = _epochsManager.currentIteration();

        Stake storage stake = stakes[stakee];

        uint256 currentStake = getCurrentStakerAmount(stakee, msg.sender);

        stake.stakeEntries[msg.sender] = StakeEntry(
            currentStake + totalToStake,
            block.number,
            currentEpochId
        );

        stake.totalManagedStake += totalToStake;
        totalManagedStake += totalToStake;

        // ensure that the node's own stake is still at the minimum amount
        if (msg.sender != stakee) {
            require(
                checkMinimumStakeProportion(stakee),
                "Can not add more stake until stakee adds more stake itself"
            );
        }
    }

    /**
     * @notice Call this function to begin the unlocking process. Calling this
     * will trigger an automatic claim of any outstanding staking rewards. Any
     * stake that was already in the unlocking phase will have the specified
     * amount added to it, and its duration refreshed. This function will fail
     * under the following conditions:
     *   - If no stake exists for the caller
     *   - If the unlock amount is zero
     *   - If the unlock amount is more than what is staked
     * Note: If calling as a Node, this function will *not* revert if it causes
     * the Node to fail to meet the minimum stake proportion. However it will still
     * prevent the Node from participating in the network until the minimum is met
     * again.
     * @param amount The amount of stake to unlock in SOLO.
     * @param stakee The address of the staked Node.
     */
    function unlockStake(uint256 amount, address stakee) external returns (uint256) {
        Stake storage stake = stakes[stakee];

        uint256 currentStake = getCurrentStakerAmount(stakee, msg.sender);

        require(currentStake > 0, "Nothing to unstake");
        require(amount > 0, "Cannot unlock with zero amount");
        require(currentStake >= amount, "Cannot unlock more than staked");

        // automatically claim any outstanding rewards generated by their existing stake
        _rewardsManager.claimStakingRewardsAsManager(stakee, msg.sender, msg.sender);

        uint256 currentEpochId = _epochsManager.currentIteration();

        stake.stakeEntries[msg.sender] = StakeEntry(
            currentStake - amount,
            block.number,
            currentEpochId
        );

        stake.totalManagedStake -= amount;
        totalManagedStake -= amount;

        bytes32 key = getKey(stakee, msg.sender);

        // Keep track of when the stake can be withdrawn
        Unlock storage unlock = unlockings[key];

        uint256 unlockAt = block.number + unlockDuration;
        if (unlock.unlockAt < unlockAt) {
            unlock.unlockAt = unlockAt;
        }

        unlock.amount += amount;

        return unlockAt;
    }

    /**
     * @notice Call this function to withdraw stake that has finished unlocking.
     * This will fail if the stake has not yet unlocked.
     * @param stakee The address of the staked Node.
     */
    function withdrawStake(address stakee) external {
        bytes32 key = getKey(stakee, msg.sender);

        Unlock storage unlock = unlockings[key];

        require(unlock.unlockAt < block.number, "Stake not yet unlocked");

        uint256 amount = unlock.amount;

        delete unlockings[key];

        _token.transfer(msg.sender, amount);
    }

    /**
     * @notice Call this function to cancel any stake that is in the process
     * of unlocking. As this essentially adds back stake to the Node, this
     * will trigger an automatic claim of any outstanding staking rewards.
     * If the specified amount to cancel is greater than the stake that is
     * currently being unlocked, it will cancel the maximum stake possible.
     * @param amount The amount of unlocking stake to cancel in SOLO.
     * @param stakee The address of the staked Node.
     */
    function cancelUnlocking(uint256 amount, address stakee) external {
        bytes32 key = getKey(stakee, msg.sender);

        Unlock storage unlock = unlockings[key];

        if (amount >= unlock.amount) {
            amount = unlock.amount;
            delete unlockings[key];
        } else {
            unlock.amount -= amount;
        }

        addStake_(amount, stakee);
    }

    /**
     * @notice Retrieve the key used to index a stake entry. The key is a hash
     * which takes both address of the Node and the staker as input.
     * @param stakee The address of the staked Node.
     * @param staker The address of the staker.
     * @return A byte-array representing the key.
     */
    function getKey(address stakee, address staker) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(stakee, staker));
    }

    /**
     * @notice Retrieve the total stake being managed by this contract.
     * @return The total amount of managed stake in SOLO.
     */
    function getTotalManagedStake() external view returns (uint256) {
        return totalManagedStake;
    }

    /**
     * @notice Retrieve a stake entry.
     * @param stakee The address of the staked Node.
     * @param staker The address of the staker.
     * @return The stake entry.
     */
    function getStakeEntry(address stakee, address staker)
        external
        view
        returns (StakeEntry memory)
    {
        return stakes[stakee].stakeEntries[staker];
    }

    /**
     * @notice Retrieve the current amount of SOLO staked against a Node by
     * a specified staker.
     * @param stakee The address of the staked Node.
     * @param staker The address of the staker.
     * @return The amount of staked SOLO.
     */
    function getCurrentStakerAmount(address stakee, address staker) public view returns (uint256) {
        return stakes[stakee].stakeEntries[staker].amount;
    }

    /**
     * @notice Retrieve the total amount of SOLO staked against a Node.
     * @param stakee The address of the staked Node.
     * @return The amount of staked SOLO.
     */
    function getStakeeTotalManagedStake(address stakee) external view returns (uint256) {
        return stakes[stakee].totalManagedStake;
    }

    /**
     * @notice Check if a Node is meeting the minimum stake proportion requirement.
     * @param stakee The address of the staked Node.
     * @return True if the Node is meeting minimum stake proportion requirement.
     */
    function checkMinimumStakeProportion(address stakee) public view returns (bool) {
        Stake storage stake = stakes[stakee];

        uint256 currentlyOwnedStake = stake.stakeEntries[stakee].amount;
        uint16 ownedStakeProportion = SyloUtils.asPerc(
            uint128(currentlyOwnedStake),
            stake.totalManagedStake
        );

        return ownedStakeProportion >= minimumStakeProportion;
    }

    /**
     * @notice This function should be called by clients to determine how much
     * additional delegated stake can be allocated to a Node via an addStake or
     * cancelUnlocking call. This is useful to avoid a revert due to
     * the minimum stake proportion requirement not being met from the additional stake.
     * @param stakee The address of the staked Node.
     */
    function calculateMaxAdditionalDelegatedStake(address stakee) external view returns (uint256) {
        Stake storage stake = stakes[stakee];

        uint256 currentlyOwnedStake = stake.stakeEntries[stakee].amount;
        uint256 totalMaxStake = (currentlyOwnedStake * SyloUtils.PERCENTAGE_DENOMINATOR) /
            minimumStakeProportion;

        require(
            totalMaxStake >= stake.totalManagedStake,
            "Can not add more delegated stake to this stakee"
        );

        return totalMaxStake - stake.totalManagedStake;
    }
}
