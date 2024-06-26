// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/Ownable2StepUpgradeable.sol";

import "../SyloToken.sol";
import "../libraries/SyloUtils.sol";
import "../SeekerPowerOracle.sol";
import "../epochs/EpochsManager.sol";
import "../payments/ticketing/RewardsManager.sol";
import "../interfaces/staking/IStakingManager.sol";

/**
 * @notice Manages stakes and delegated stakes for Nodes. Holding
 * staked Sylo is necessary for a Node to participate in the
 * Sylo Network. The stake is used in stake-weighted scan function,
 * and delegated stakers are rewarded on a pro-rata basis.
 */
contract StakingManager is IStakingManager, Initializable, Ownable2StepUpgradeable, ERC165 {
    /** ERC 20 compatible token we are dealing with */
    IERC20 public _token;

    /**
     * @notice Rewards Manager contract. Any changes to stake will automatically
     * trigger a claim to any outstanding rewards.
     */
    RewardsManager public _rewardsManager;

    EpochsManager public _epochsManager;

    SeekerPowerOracle public _seekerPowerOracle;

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
    uint32 public minimumStakeProportion;

    /**
     * @notice The multiplier used in determining a Seeker's staking
     * capacity based on its power level.
     */
    uint256 public seekerPowerMultiplier;

    event UnlockDurationUpdated(uint256 unlockDuration);
    event MinimumStakeProportionUpdated(uint256 minimumStakeProportion);
    event SeekerPowerMultiplierUpdated(uint256 seekerPowerMultipler);

    error NoStakeToUnlock();
    error StakeNotYetUnlocked();
    error CannotStakeZeroAmount();
    error CannotUnlockZeroAmount();
    error TokenCannotBeZeroAddress();
    error StakeeCannotBeZeroAddress();
    error UnlockDurationCannotBeZero();
    error CannotCancelUnlockZeroAmount();
    error CannotUnlockMoreThanStaked(uint256 stakeAmount, uint256 unlockAmount);
    error StakeCapacityReached(uint256 maxCapacity, uint256 currentCapacity);
    error SeekerPowerNotRegistered(uint256 seekerId);

    function initialize(
        IERC20 token,
        RewardsManager rewardsManager,
        EpochsManager epochsManager,
        SeekerPowerOracle seekerPowerOracle,
        uint256 _unlockDuration,
        uint32 _minimumStakeProportion,
        uint256 _seekerPowerMultiplier
    ) external initializer {
        if (address(token) == address(0)) {
            revert TokenCannotBeZeroAddress();
        }

        SyloUtils.validateContractInterface(
            "RewardsManager",
            address(rewardsManager),
            type(IRewardsManager).interfaceId
        );

        SyloUtils.validateContractInterface(
            "EpochsManager",
            address(epochsManager),
            type(IEpochsManager).interfaceId
        );

        SyloUtils.validateContractInterface(
            "SeekerPowerOracle",
            address(seekerPowerOracle),
            type(ISeekerPowerOracle).interfaceId
        );

        if (_unlockDuration == 0) {
            revert UnlockDurationCannotBeZero();
        }

        Ownable2StepUpgradeable.__Ownable2Step_init();

        _token = token;
        _rewardsManager = rewardsManager;
        _epochsManager = epochsManager;
        _seekerPowerOracle = seekerPowerOracle;
        unlockDuration = _unlockDuration;
        minimumStakeProportion = _minimumStakeProportion;
        seekerPowerMultiplier = _seekerPowerMultiplier;
    }

    /**
     * @notice Returns true if the contract implements the interface defined by
     * `interfaceId` from ERC165.
     */
    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return interfaceId == type(IStakingManager).interfaceId;
    }

    /**
     * @notice Sets the unlock duration for stakes. Only callable by
     * the owner.
     * @param _unlockDuration The unlock duration in number of blocks.
     */
    function setUnlockDuration(uint256 _unlockDuration) external onlyOwner {
        if (_unlockDuration == 0) {
            revert UnlockDurationCannotBeZero();
        }

        unlockDuration = _unlockDuration;
        emit UnlockDurationUpdated(_unlockDuration);
    }

    function setSeekerPowerMultiplier(uint256 _seekerPowerMultiplier) external onlyOwner {
        seekerPowerMultiplier = _seekerPowerMultiplier;
        emit SeekerPowerMultiplierUpdated(seekerPowerMultiplier);
    }

    /**
     * @notice Sets the minimum stake proportion for Nodes. Only callable by
     * the owner.
     * @param _minimumStakeProportion The minimum stake proportion in SOLO.
     */
    function setMinimumStakeProportion(uint32 _minimumStakeProportion) external onlyOwner {
        minimumStakeProportion = _minimumStakeProportion;
        emit MinimumStakeProportionUpdated(_minimumStakeProportion);
    }

    /**
     * @notice Called by Nodes and delegated stakers to add stake.
     * This function will fail under the following conditions:
     *   - If the Node address is invalid
     *   - If the specified stake value is zero
     *   - If the additional stake causes the Node to fail to meet the
     *     minimum stake proportion requirement.
     * @param amount The amount of stake to add in SOLO.
     * @param stakee The address of the staked Node.
     */
    function addStake(uint256 amount, address stakee) external {
        if (stakee == address(0)) {
            revert StakeeCannotBeZeroAddress();
        }
        if (amount == 0) {
            revert CannotStakeZeroAmount();
        }

        _addStake(amount, stakee);
        SafeERC20.safeTransferFrom(_token, msg.sender, address(this), amount);
    }

    function _addStake(uint256 amount, address stakee) internal {
        // automatically move any pending rewards generated by their existing stake
        _rewardsManager.updatePendingRewards(stakee, msg.sender);

        uint256 currentEpochId = _epochsManager.currentIteration();

        Stake storage stake = stakes[stakee];

        uint256 currentStake = getCurrentStakerAmount(stakee, msg.sender);

        stake.stakeEntries[msg.sender] = StakeEntry(
            currentStake + amount,
            block.number,
            currentEpochId
        );

        stake.totalManagedStake = stake.totalManagedStake + amount;
        totalManagedStake = totalManagedStake + amount;
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
        if (stakee == address(0)) {
            revert StakeeCannotBeZeroAddress();
        }
        if (amount == 0) {
            revert CannotUnlockZeroAmount();
        }

        uint256 currentStake = getCurrentStakerAmount(stakee, msg.sender);

        if (currentStake == 0) {
            revert NoStakeToUnlock();
        }
        if (currentStake < amount) {
            revert CannotUnlockMoreThanStaked(currentStake, amount);
        }

        // automatically move any pending rewards generated by their existing stake
        _rewardsManager.updatePendingRewards(stakee, msg.sender);

        uint256 currentEpochId = _epochsManager.currentIteration();

        Stake storage stake = stakes[stakee];

        stake.stakeEntries[msg.sender] = StakeEntry(
            currentStake - amount,
            block.number,
            currentEpochId
        );

        stake.totalManagedStake = stake.totalManagedStake - amount;
        totalManagedStake = totalManagedStake - amount;

        bytes32 key = getKey(stakee, msg.sender);

        // Keep track of when the stake can be withdrawn
        Unlock storage unlock = unlockings[key];

        uint256 unlockAt = block.number + unlockDuration;
        if (unlock.unlockAt < unlockAt) {
            unlock.unlockAt = unlockAt;
        }

        unlock.amount = unlock.amount + amount;

        return unlockAt;
    }

    /**
     * @notice Call this function to withdraw stake that has finished unlocking.
     * This will fail if the stake has not yet unlocked.
     * @param stakee The address of the staked Node.
     */
    function withdrawStake(address stakee) external {
        if (stakee == address(0)) {
            revert StakeeCannotBeZeroAddress();
        }

        bytes32 key = getKey(stakee, msg.sender);

        Unlock storage unlock = unlockings[key];

        if (unlock.unlockAt >= block.number) {
            revert StakeNotYetUnlocked();
        }

        uint256 amount = unlock.amount;

        delete unlockings[key];

        SafeERC20.safeTransfer(_token, msg.sender, amount);
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
        if (stakee == address(0)) {
            revert StakeeCannotBeZeroAddress();
        }
        if (amount == 0) {
            revert CannotCancelUnlockZeroAmount();
        }

        bytes32 key = getKey(stakee, msg.sender);

        Unlock storage unlock = unlockings[key];

        if (amount >= unlock.amount) {
            amount = unlock.amount;
            delete unlockings[key];
        } else {
            unlock.amount = unlock.amount - amount;
        }

        _addStake(amount, stakee);
    }

    /**
     * @notice This function determines the staking capacity of
     * a Seeker based on its power level. The method will revert if
     * the Seeker's power level has not been registered with the oracle.
     *
     * Currently the algorithm is as follows:
     *    staking_capacity = seeker_power * seeker_power_multiplier;
     */
    function calculateCapacityFromSeekerPower(uint256 seekerId) external view returns (uint256) {
        uint256 seekerPower = _seekerPowerOracle.getSeekerPower(seekerId);
        if (seekerPower == 0) {
            revert SeekerPowerNotRegistered(seekerId);
        }

        // If the Seeker Power is already
        // at the maximum sylo, then we just return the max sylo value directly.
        if (seekerPower >= SyloUtils.MAX_SYLO) {
            return SyloUtils.MAX_SYLO;
        }

        uint256 capacity = seekerPower * seekerPowerMultiplier;

        return capacity > SyloUtils.MAX_SYLO ? SyloUtils.MAX_SYLO : capacity;
    }

    /**
     * @notice This function can be used to a determine a Node's staking capacity,
     * based on the minimum stake proportion constant.
     * @param stakee The address of the staked Node.
     */
    function calculateCapacityFromMinStakingProportion(
        address stakee
    ) public view returns (uint256) {
        if (stakee == address(0)) {
            revert StakeeCannotBeZeroAddress();
        }

        Stake storage stake = stakes[stakee];

        uint256 currentlyOwnedStake = stake.stakeEntries[stakee].amount;
        return (currentlyOwnedStake * SyloUtils.PERCENTAGE_DENOMINATOR) / minimumStakeProportion;
    }

    /**
     * @notice This function should be called by clients to determine how much
     * additional delegated stake can be allocated to a Node via an addStake or
     * cancelUnlocking call. This is useful to avoid a revert due to
     * the minimum stake proportion requirement not being met from the additional stake.
     * @param stakee The address of the staked Node.
     */
    function calculateMaxAdditionalDelegatedStake(address stakee) external view returns (uint256) {
        uint256 totalMaxStake = calculateCapacityFromMinStakingProportion(stakee);

        Stake storage stake = stakes[stakee];

        if (totalMaxStake < stake.totalManagedStake) {
            revert StakeCapacityReached(totalMaxStake, stake.totalManagedStake);
        }

        return totalMaxStake - stake.totalManagedStake;
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
    function getStakeEntry(
        address stakee,
        address staker
    ) external view returns (StakeEntry memory) {
        return stakes[stakee].stakeEntries[staker];
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
        if (stakee == address(0)) {
            revert StakeeCannotBeZeroAddress();
        }

        Stake storage stake = stakes[stakee];

        uint256 currentlyOwnedStake = stake.stakeEntries[stakee].amount;
        uint32 ownedStakeProportion = SyloUtils.asPerc(
            SafeCast.toUint128(currentlyOwnedStake),
            stake.totalManagedStake
        );

        return ownedStakeProportion >= minimumStakeProportion;
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
}
