pragma solidity ^0.6.0;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts/access/Ownable.sol";
import "../contracts/Token.sol";

/*
 * Directory for accounts that wish to offer services to the Sylo Network
 * It provides the ability for accounts to stake to become listed
 * It also provides the functionality for a client to get a random stake weighted selected service peer
*/
contract Directory is Ownable {

    struct Stake {
        uint256 amount; // Amount of the stake

        uint256 leftAmount; // Value of stake on left branch
        uint256 rightAmount; // Value of stake on right branch

        address stakee; // Address of peer that offers services

        bytes32 parent; // A hash of staker + stakee
        bytes32 left; // Pointer to left child
        bytes32 right; // Pointer to right child
    }

    struct Unlock {
        uint256 amount; // Amount of stake unlocking

        uint256 unlockAt; // Block number the stake becomes withdrawable
    }

    uint256 constant maxU256 = 2^256-1;

    /* ERC 20 compatible token we are dealing with */
    IERC20 _token;

    /*
     * The number of blocks a user must wait after calling "unlock"
     * before they can withdraw their stake
     */
    uint256 public unlockDuration;

    mapping(bytes32 => Stake) public stakes;

    // Keeps track of stakees stake amount
    mapping(address => uint256) public stakees;

    // Funds that are in the process of being unlocked
    mapping(bytes32 => Unlock) public unlockings;

    bytes32 root;

    constructor(IERC20 token, uint256 _unlockDuration) public {
        _token = token;
        unlockDuration = _unlockDuration;
    }

    function setUnlockDuration(uint256 newUnlockDuration) public onlyOwner {
        unlockDuration = newUnlockDuration;
    }

    function addStake(uint256 amount, address stakee) public {

        require(stakee != address(0), "Address is null");
        require(amount != 0, "Cannot stake nothing");

        bytes32 key = getKey(msg.sender, stakee);

        Stake storage stake = stakes[key];

        // New stake
        if (stake.amount == 0) {

            // Find the node to add a new child
            bytes32 parent = root;
            Stake storage current = stakes[parent];
            while(parent != bytes32(0)) {
                bytes32 next = current.leftAmount < current.rightAmount
                    ? current.left
                    : current.right;

                if (next == bytes32(0)) {
                    break;
                }

                parent = next;
                current = stakes[parent];
            }

            // Set the new child on the node
            setChild(current, bytes32(0), key);

            stake.parent = parent;
            stake.stakee = stakee;
        }

        // First stake lets set the root
        if (stake.parent == bytes32(0)) {
            root = key;
        }

        updateStakeAmount(key, stake, amount);

        _token.transferFrom(msg.sender, address(this), amount);
    }

    function unlockStake(uint256 amount, address stakee) public returns (uint256) {

        bytes32 key = getKey(msg.sender, stakee);
        Stake storage stake = stakes[key];

        require(stake.amount > 0, "Nothing to unstake");
        require(stake.amount >= amount, "Cannot unlock more than staked");

        updateStakeAmount(key, stake, -amount);

        // All stake being withdrawn, update the tree
        if (stake.amount == 0) {
            bytes32 child = stake.leftAmount > stake.rightAmount ? stake.left : stake.right;

            Stake storage parent = stakes[stake.parent];

            if (child == bytes32(0)) {
                // Stake is a leaf, we need to disconnect it from the parent
                setChild(stake, key, bytes32(0));
            } else {
                Stake storage current = stakes[child];

                // Find the leaf of the most valuable path of the child
                while(true) {
                    bytes32 next = current.leftAmount > current.rightAmount ? current.left : current.right;
                    if (next == bytes32(0)) {
                        break;
                    }
                    child = next;
                    current = stakes[next];
                }

                bytes32 currentParent = current.parent;

                // Move leaf to position of removed stake
                setChild(parent, key, child);
                current.parent = stake.parent;

                // Update the children of current to be that of what the removed stake was
                if (currentParent != key) {

                    // Move the children of stake to current
                    fixl(stake, child, current);
                    fixr(stake, child, current);

                    // Place stake where current was and 
                    stake.parent = currentParent; // Set parent
                    setChild(stakes[currentParent], currentParent, key); // Set parents child

                    // Update all the values
                    applyStakeChange(key, stake, -current.amount, current.parent);

                    // Remove reference to the old stake
                    setChild(stakes[currentParent], key, bytes32(0));
                } else if (stake.left == child) {
                    fixr(stake, child, current);
                } else {
                    fixl(stake, child, current);
                }

                // Update the root if thats changed
                if (current.parent == bytes32(0)) {
                    root = child;
                }
            }

            // Now that the node is unlinked from any other nodes, we can remove it
            delete stakes[key];
        }

        // Keep track of when the stake can be withdrawn
        Unlock storage unlock = unlockings[key];

        uint256 unlockAt = block.number + unlockDuration;
        if (unlock.unlockAt < unlockAt) {
            unlock.unlockAt = unlockAt;
        }

        unlock.amount += amount;

        return unlockAt;
    }

    // Reverse unlocking a certain amount of stake
    function lockStake(uint256 amount, address stakee) public {
        bytes32 key = getKey(msg.sender, stakee);
        Stake storage stake = stakes[key];

        pullUnlocking(key, amount);

        updateStakeAmount(key, stake, amount);
    }

    function unstake(address stakee) public {
        bytes32 key = getKey(msg.sender, stakee);

        Unlock storage unlock = unlockings[key];

        require(unlock.unlockAt < block.number, "Stake not yet unlocked");
        require(unlock.amount > 0, "No amount to unlock");

        _token.transfer(msg.sender, unlock.amount);
    }

    // Point should be a randomly generated uint256
    function scan(uint256 point) public view returns (address) {

        // Nothing is staked
        if (root == bytes32(0)) {
            return address(0);
        }

        uint256 expectedVal = getTotalStake() * point / maxU256;

        bytes32 current = root;

        while(true) {

            Stake storage stake = stakes[current];

            if (expectedVal < stake.leftAmount) {
                current = stake.left;
                continue;
            }

            expectedVal -= stake.leftAmount;

            if (expectedVal <= stake.amount) {
                return stake.stakee;
            }

            expectedVal -= stake.amount;

            current = stake.right;
        }

        return address(0);
    }

    function getKey(address staker, address stakee) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(staker, stakee));
    }

    function getStake(address stakee) private view returns (Stake storage) {
        return stakes[getKey(msg.sender, stakee)];
    }

    function getTotalStake() public view returns (uint256) {
        if (root == bytes32(0)) {
            return 0;
        }

        Stake storage stake = stakes[root];

        return stake.amount + stake.leftAmount + stake.rightAmount;
    }

    function pullUnlocking(bytes32 key, uint256 amount) private {
        Unlock storage unlock = unlockings[key];

        // TODO guard unlockAt

        if (amount == unlock.amount) {
            delete unlockings[key];
        } else {
            require(amount < unlock.amount, "Unlock has insufficient amount");
            unlock.amount -= amount;
        }
    }

    function updateStakeAmount(bytes32 key, Stake storage stake, uint256 amount) private {
        stake.amount += amount;
        stakees[stake.stakee] += amount;

        applyStakeChange(key, stake, amount, bytes32(0));
    }

    // Recursively update left/right amounts of stakes for parents of an updated stake
    function applyStakeChange(bytes32 key, Stake storage stake, uint256 amount, bytes32 root_) private {
        bytes32 parentKey = stake.parent;

        if (parentKey == root_) {
            // We are at the root, theres nothing left to update
            return;
        }

        Stake storage parent = stakes[parentKey];

        if (parent.left == key) {
            parent.leftAmount += amount;
        } else {
            parent.rightAmount += amount;
        }

        return applyStakeChange(parentKey, parent, amount, root_);
    }

    // Move the left child of stake to current
    function fixl(Stake storage stake, bytes32 currentKey, Stake storage current) private {
        if (stake.left == bytes32(0)) {
            return;
        }

        stakes[stake.left].parent = currentKey;
        current.left = stake.left;
        current.leftAmount = stake.leftAmount;
    }

    // Move the right child of stake to current
    function fixr(Stake storage stake, bytes32 currentKey, Stake storage current) private {
        if (stake.right == bytes32(0)) {
            return;
        }

        stakes[stake.right].parent = currentKey;
        current.right = stake.right;
        current.rightAmount = stake.rightAmount;
    }

    function setChild(Stake storage stake, bytes32 oldKey, bytes32 newKey) private {
        if (stake.left == oldKey) {
            stake.left = newKey;
        } else {
            stake.right = newKey;
        }
    }
}
