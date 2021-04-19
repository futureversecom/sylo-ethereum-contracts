// SPDX-License-Identifier: UNLICENSED
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

    struct NodePointer {
        bytes32 value_;
    }

    struct Node {
        uint256 amount; // Amount of the stake

        uint256 leftAmount; // Value of stake on left branch
        uint256 rightAmount; // Value of stake on right branch

        address stakee; // Address of peer that offers services

        bytes32 parent; // A hash of staker + stakee
        NodePointer left; // Pointer to left child
        NodePointer right; // Pointer to right child
    }

    struct Unlock {
        uint256 amount; // Amount of stake unlocking

        uint256 unlockAt; // Block number the stake becomes withdrawable
    }

    /* ERC 20 compatible token we are dealing with */
    IERC20 _token;

    /*
     * The number of blocks a user must wait after calling "unlock"
     * before they can withdraw their stake
     */
    uint256 public unlockDuration;

    mapping(bytes32 => Node) public nodes;

    // Keeps track of stakees stake amount
    mapping(address => uint256) public stakees;

    // Funds that are in the process of being unlocked
    mapping(bytes32 => Unlock) public unlockings;

    NodePointer root;

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

        address staker = msg.sender;
        bytes32 key = getKey(staker, stakee);

        Node storage node = nodes[key];

        if (node.amount == 0) {
            // There is no node for this key in the stake tree, so we will need
            // to traverse the tree and find a place to add the stake.
            bytes32 parent = bytes32(0);
            NodePointer storage p = root;

            while(p.value_ != bytes32(0)) {
                parent = p.value_;
                Node storage current = nodes[parent];
                // To decend the tree we consider the current node to be on the
                // right. This will fill the left side of the tree first when
                // there is a tie.
                p = current.leftAmount < current.amount + current.rightAmount
                    ? current.left
                    : current.right;
            }
            node.parent = parent;
            p.value_ = key;
            node.stakee = stakee;
        }

        if (node.parent == bytes32(0)) {
            // This is the first node, so we need to define the root.
            root.value_ = key;
        }

        updateStakeAmount(key, node, amount);

        _token.transferFrom(staker, address(this), amount);
    }

    function unlockStake(uint256 amount, address stakee) public returns (uint256) {

        bytes32 key = getKey(msg.sender, stakee);
        Node storage node = nodes[key];

        require(node.amount > 0, "Nothing to unstake");
        require(node.amount >= amount, "Cannot unlock more than staked");

        updateStakeAmount(key, node, -amount);

        // All stake being withdrawn, update the tree
        if (node.amount == 0) {
            NodePointer storage child = node.leftAmount > node.rightAmount ? node.left : node.right;

            Node storage parent = nodes[node.parent];

            if (child.value_ == bytes32(0)) {
                // Stake is a leaf, we need to disconnect it from the parent
                setChild(parent, key, bytes32(0));

                // The only staker is removed, reset root
                if (node.parent == bytes32(0)) {
                    root.value_ = bytes32(0);
                }
            } else {
                Node storage current = nodes[child.value_];

                // Find the leaf of the most valuable path of the child
                while(true) {
                    NodePointer storage next = current.leftAmount > current.rightAmount ? current.left : current.right;
                    if (next.value_ == bytes32(0)) {
                        break;
                    }
                    child = next;
                    current = nodes[next.value_];
                }

                bytes32 currentParent = current.parent;

                // Move leaf to position of removed stake
                setChild(parent, key, child.value_);
                current.parent = node.parent;

                // Update the children of current to be that of what the removed stake was
                if (currentParent != key) {

                    // Move the children of stake to current
                    fixl(node, child.value_, current);
                    fixr(node, child.value_, current);

                    // Place stake where current was and
                    node.parent = currentParent; // Set parent
                    setChild(nodes[currentParent], child.value_, key); // Set parents child

                    // Update all the values
                    applyStakeChange(key, node, -current.amount, current.parent);

                    // Remove reference to the old stake
                    setChild(nodes[currentParent], key, bytes32(0));
                } else if (node.left.value_ == child.value_) {
                    fixr(node, child.value_, current);
                } else {
                    fixl(node, child.value_, current);
                }

                // Update the root if thats changed
                if (current.parent == bytes32(0)) {
                    root = child;
                }
            }

            // Now that the node is unlinked from any other nodes, we can remove it
            delete nodes[key];
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
        Node storage stake = nodes[key];

        pullUnlocking(key, amount);

        updateStakeAmount(key, stake, amount);
    }

    function unstake(address stakee) public {
        bytes32 key = getKey(msg.sender, stakee);

        Unlock storage unlock = unlockings[key];

        require(unlock.unlockAt < block.number, "Stake not yet unlocked");
        require(unlock.amount > 0, "No amount to unlock");

        uint256 amount = unlock.amount;

        delete unlockings[key];

        _token.transfer(msg.sender, amount);
    }

    function scan(uint128 point) public view returns (address) {

        // Nothing is staked
        if (root.value_ == bytes32(0)) {
            return address(0);
        }

        // Staking all the Sylo would only be 94 bits, so multiplying this with
        // a uint128 cannot overflow a uint256.
        uint256 expectedVal = getTotalStake() * uint256(point) >> 128;

        bytes32 current = root.value_;

        for (;;) {

            Node storage stake = nodes[current];

            // Prefer decending the left side.
            if (expectedVal < stake.leftAmount) {
                if (stake.left.value_ == bytes32(0)) {
                    // There is no node on the left. This is a leaf so just
                    // return this node's address.
                    return stake.stakee;
                }
                current = stake.left.value_;
                continue;
            }

            // We are not decending the left side of the tree so remove the
            // entire value of that subtree.
            expectedVal -= stake.leftAmount;

            // If value is less the the current node (the "middle"), we are
            // done.
            if (expectedVal < stake.amount) {
                return stake.stakee;
            }

            // Else decend the right side of the tree and continue the loop.
            require(stake.right.value_ != bytes32(0), "missing node on the right");
            current = stake.right.value_;
            expectedVal -= stake.amount;
        }

        return address(0);
    }

    function getKey(address staker, address stakee) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(staker, stakee));
    }

    function getStake(address stakee) private view returns (Node storage) {
        return nodes[getKey(msg.sender, stakee)];
    }

    function getTotalStake() public view returns (uint256) {
        if (root.value_ == bytes32(0)) {
            return 0;
        }

        Node storage stake = nodes[root.value_];

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

    function updateStakeAmount(bytes32 key, Node storage stake, uint256 amount) private {
        stake.amount += amount;
        stakees[stake.stakee] += amount;

        applyStakeChange(key, stake, amount, bytes32(0));
    }

    // Recursively update left/right amounts of stakes for parents of an updated stake
    function applyStakeChange(bytes32 key, Node storage stake, uint256 amount, bytes32 root_) private {
        bytes32 parentKey = stake.parent;

        if (parentKey == root_) {
            // We are at the root, theres nothing left to update
            return;
        }

        Node storage parent = nodes[parentKey];

        if (parent.left.value_ == key) {
            parent.leftAmount += amount;
        } else {
            parent.rightAmount += amount;
        }

        return applyStakeChange(parentKey, parent, amount, root_);
    }

    // Move the left child of stake to current
    function fixl(Node storage stake, bytes32 currentKey, Node storage current) private {
        if (stake.left.value_ == bytes32(0)) {
            return;
        }

        nodes[stake.left.value_].parent = currentKey;
        current.left = stake.left;
        current.leftAmount = stake.leftAmount;
    }

    // Move the right child of stake to current
    function fixr(Node storage stake, bytes32 currentKey, Node storage current) private {
        if (stake.right.value_ == bytes32(0)) {
            return;
        }

        nodes[stake.right.value_].parent = currentKey;
        current.right = stake.right;
        current.rightAmount = stake.rightAmount;
    }

    function setChild(Node storage stake, bytes32 oldKey, bytes32 newKey) private {
        if (stake.left.value_ == oldKey) {
            stake.left.value_ = newKey;
        } else {
            stake.right.value_ = newKey;
        }
    }
}
