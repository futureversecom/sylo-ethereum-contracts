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

    struct StakeEntry {
        uint256 amount;
        // Block number this entry was created at
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
        // The amount of stake a staker held at a specific block number can be
        // found by finding the most recent entry prior to the specified block.
        // (refer to `getStakerAmount` for implementation)
        mapping (address => StakeEntry[]) stakeEntries;

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

        uint256 currentStake = getCurrentStakerAmount(msg.sender, stakee);

        stake.stakeEntries[msg.sender].push(
            StakeEntry(
                currentStake + amount,
                block.number
            )
        );

        stake.totalStake += amount;
    }

    function unlockStake(uint256 amount, address stakee) public returns (uint256) {
        Stake storage stake = stakes[stakee];

        uint256 currentStake = getCurrentStakerAmount(msg.sender, stakee);

        require(currentStake > 0, "Nothing to unstake");
        require(currentStake >= amount, "Cannot unlock more than staked");

        stake.stakeEntries[msg.sender].push(
            StakeEntry(
                currentStake - amount,
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

    function getStakeEntries(address staker, address stakee) public view returns (StakeEntry[] memory) {
        return stakes[stakee].stakeEntries[staker];
    }

    function getStakeeTotalStake(address stakee) public view returns (uint256) {
        return stakes[stakee].totalStake;
    }

    /*
     * Helper function that finds the stake amount for a staker at a given block.
     * It will search through the historical entries up to the specified block,
     * and return the previous value.
     */
    function getStakerAmount(address staker, address stakee, uint blockNumber) public view returns (uint256) {
        uint256 length = stakes[stakee].stakeEntries[staker].length;
        if (length == 0) {
            return 0;
        }

        uint256 l = 0;
        uint256 r = stakes[stakee].stakeEntries[staker].length - 1;

        StakeEntry memory end = stakes[stakee].stakeEntries[staker][r];
        if (end._block <= blockNumber) {
            return end.amount;
        }

        // since the entries are sorted by block number, we can perform
        // a binary search to optimize the gas cost
        while (l <= r) {
            uint index = (l + r) / 2;

            uint lower = index == 0 ? 0 : stakes[stakee].stakeEntries[staker][index - 1]._block;
            uint upper = stakes[stakee].stakeEntries[staker][index]._block;

            if (blockNumber >= lower && blockNumber < upper) {
                return stakes[stakee].stakeEntries[staker][index - 1].amount;
            } else if (blockNumber < lower) {
                r = index - 1;
            } else if (blockNumber >= upper) {
                l = index + 1;
            }
        }

        return 0;
    }

    function getCurrentStakerAmount(address staker, address stakee) public view returns (uint256) {
        return getStakerAmount(staker, stakee, block.number);
    }
}