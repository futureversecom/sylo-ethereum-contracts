// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "../Token.sol";

/*
 * Manages stakes and delegated stakes for accounts that wish to be listed
*/
contract StakingManager is Initializable, OwnableUpgradeable {
    /* ERC 20 compatible token we are dealing with */
    IERC20 _token;

    enum StakeOperationType{ ADDSTAKE, UNLOCKSTAKE }

    struct StakeOperation {
        StakeOperationType _type;
        uint256 amount;

        // Block number this operation was created at
        uint256 _block;
    }

    /*
     * Every Node must have stake in order to participate in the Epoch.
     * Stake can be provided by the Node itself or by other accounts in
     * the network.
     */
    struct Stake {
        // For each staker associated to a node, we track the historical
        // changes in their given stake instead of just tracking a
        // single stake value. This is because other contracts will need
        // to know the state of a Node's stake at certain points in time.
        // The stake can be calculated by folding over all of the operations
        // for a particular stakee up until a specified block number
        // (refer to `getStakerAmount` for implementation)
        mapping (address => StakeOperation[]) stakerOperations;
        uint256 totalStake;
    }

    struct Unlock {
        uint256 amount; // Amount of stake unlocking
        uint256 unlockAt; // Block number the stake becomes withdrawable
    }

    mapping (address => Stake) stakes;

    /* Tracks overall total stake */
    uint256 public totalStake;

    /*
     * The number of blocks a user must wait after calling "unlock"
     * before they can withdraw their stake
     */
    uint256 public unlockDuration;

    /* Tracks funds that are in the process of being unlocked */
    mapping(bytes32 => Unlock) public unlockings;

    function initialize(IERC20 token, uint256 _unlockDuration) public initializer {
        OwnableUpgradeable.__Ownable_init();
        _token = token;
        unlockDuration = _unlockDuration;
    }

    function setUnlockDuration(uint256 newUnlockDuration) public onlyOwner {
        unlockDuration = newUnlockDuration;
    }

    function addStake(uint256 amount, address stakee) public {
        addStake_(amount, stakee);
        _token.transferFrom(msg.sender, address(this), amount);
    }

    function addStake_(uint256 amount, address stakee) internal {
        require(stakee != address(0), "Address is null");
        require(amount != 0, "Cannot stake nothing");

        Stake storage stake = stakes[stakee];

        stake.stakerOperations[msg.sender].push(
            StakeOperation(
                StakeOperationType.ADDSTAKE,
                amount,
                block.number
            )
        );

        stake.totalStake += amount;
    }

    function unlockStake(uint256 amount, address stakee) public returns (uint256) {
        Stake storage stake = stakes[stakee];

        uint256 currentStake = getStakerAmount(msg.sender, stakee, block.number);

        require(currentStake > 0, "Nothing to unstake");
        require(currentStake >= amount, "Cannot unlock more than staked");

        stake.stakerOperations[msg.sender].push(
            StakeOperation(
                StakeOperationType.UNLOCKSTAKE,
                amount,
                block.number
            )
        );

        stake.totalStake -= amount;

        bytes32 key = getKey(msg.sender, stakee);

        // Keep track of when the stake can be withdrawn
        Unlock storage unlock = unlockings[key];

        uint256 unlockAt = block.number + unlockDuration;
        if (unlock.unlockAt < unlockAt) {
            unlock.unlockAt = unlockAt;
        }

        unlock.amount += amount;

        return unlockAt;
    }

    // Withdraw any unlocked stake.
    function withdrawStake(address stakee) public {
        bytes32 key = getKey(msg.sender, stakee);

        Unlock storage unlock = unlockings[key];

        require(unlock.unlockAt < block.number, "Stake not yet unlocked");
        require(unlock.amount > 0, "No amount to withdraw");

        uint256 amount = unlock.amount;

        delete unlockings[key];

        _token.transfer(msg.sender, amount);
    }

    // Reverse unlocking a certain amount of stake
    function cancelUnlocking(uint256 amount, address stakee) public {
        bytes32 key = getKey(msg.sender, stakee);

        Unlock storage unlock = unlockings[key];

        // TODO guard unlockAt

        if (amount == unlock.amount) {
            delete unlockings[key];
        } else {
            require(amount < unlock.amount, "Unlock has insufficient amount");
            unlock.amount -= amount;
        }

        addStake_(amount, stakee);
    }

    function getKey(address staker, address stakee) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(staker, stakee));
    }

    function getTotalStake() public view returns (uint256) {
        return totalStake;
    }

    function getStakeOperations(address staker, address stakee) public view returns (StakeOperation[] memory) {
        return stakes[stakee].stakerOperations[staker];
    }

    function getStakeeTotalStake(address stakee) public view returns (uint256) {
        return stakes[stakee].totalStake;
    }

    /*
     * Helper function that returns the total stake for a staker at a given block.
     * It will fold over all historical operations up to the specified block,
     * and return the final value.
     */
    function getStakerAmount(address staker, address stakee, uint blockNumber) public view returns (uint256) {
        StakingManager.StakeOperation[] memory operations = getStakeOperations(staker, stakee);
        uint256 stake = 0;

        for (uint i = 0; i < operations.length; i++) {
            StakingManager.StakeOperation memory op = operations[i];

            // We have folded over all operations prior to the specified block
            if (op._block > blockNumber) {
                break;
            }

            if (op._type == StakeOperationType.ADDSTAKE) {
                stake += op.amount;
            } else if (op._type == StakeOperationType.UNLOCKSTAKE) {
                stake -= op.amount;
            }
        }

        return stake;
    }
}