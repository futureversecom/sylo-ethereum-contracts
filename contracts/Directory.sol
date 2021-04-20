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

        NodePointer parent; // Pointer to parent
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
    NodePointer empty;

    constructor(IERC20 token, uint256 _unlockDuration) public {
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

    function addStake_(uint256 amount, address stakee) private {
        require(stakee != address(0), "Address is null");
        require(amount != 0, "Cannot stake nothing");

        address staker = msg.sender;
        bytes32 key = getKey(staker, stakee);

        Node storage node = nodes[key];

        if (node.amount == 0) {
            // There is no node for this key in the stake tree, so we will need
            // to traverse the tree and find a place to add the stake.
            NodePointer storage p = root;
            NodePointer memory parent;
            parent.value_ = bytes32(0);

            while(p.value_ != bytes32(0)) {
                parent = p;
                Node storage current = nodes[parent.value_];
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

        if (node.parent.value_ == bytes32(0)) {
            // This is the first node, so we need to define the root.
            root.value_ = key;
        }

        updateStakeAmount(key, node, amount);
    }

    // unlockStake will immediately remove some amount of stake from the stake
    // tree and place it into the `unlockings` data structure.
    //
    // Currently, unlocking more stake while stake is being unlocked will reset
    // the `unlockAt` block for all unlocking stake since there is only one
    // location per address to store unlocking data.
    //
    // Also, unlocking all stake will remove the node from the stake tree. This
    // means that if the unlocking is cancelled, the stake may be added to a new
    // location on the stake tree.
    function unlockStake(uint256 amount, address stakee) public returns (uint256) {

        bytes32 key = getKey(msg.sender, stakee);
        Node storage stakeNode = nodes[key];
        NodePointer memory stakeNodePointer;
        stakeNodePointer.value_ = key;

        require(stakeNode.amount > 0, "Nothing to unstake");
        require(stakeNode.amount >= amount, "Cannot unlock more than staked");

        updateStakeAmount(stakeNodePointer.value_, stakeNode, -amount);

        // All stake being withdrawn, update the tree
        if (stakeNode.amount == 0) {
            if (stakeNode.leftAmount + stakeNode.rightAmount == 0) {
                // No children have any stake value (i.e. no children)
                if (stakeNode.parent.value_ == bytes32(0)) {
                    // The only node is being removed, reset root
                    root.value_ = bytes32(0);
                } else {
                    // Stake is a leaf, we need to remove it from the parent
                    removeChild(nodes[stakeNode.parent.value_], stakeNodePointer.value_);
                }
            } else {
                // There is a subtree. Find the leaf down the most valuable path
                // of the subtree.

                if (stakeNode.parent.value_ == bytes32(0)) {
                    // parent is root
                }

                // Find the leaf node that will go into the old place on the
                // tree.
                NodePointer memory subtreePointer = stakeNode.leftAmount > stakeNode.rightAmount
                    ? stakeNode.left
                    : stakeNode.right;
                Node storage subtreeNode = nodes[subtreePointer.value_];
                NodePointer memory leafPointer = subtreePointer;
                for (;;) {
                    NodePointer memory next = subtreeNode.leftAmount > subtreeNode.rightAmount
                        ? subtreeNode.left
                        : subtreeNode.right;
                    if (next.value_ == bytes32(0)) {
                        break;
                    }
                    leafPointer = next;
                    subtreeNode = nodes[next.value_];
                }

                NodePointer memory leafParent = subtreeNode.parent;
                Node storage leafNode = subtreeNode;

                // Move leaf node to position of removed stake node. It's a leaf
                // node, so it won't have any children.
                setChild(leafNode, bytes32(0), leafPointer.value_);
                leafNode.parent = stakeNode.parent;

                // Update the children of current to be that of what the removed
                // stake was.
                if (leafParent.value_ != key) {

                    // Move the children of stakeNode to leafNode
                    fixl(stakeNode, leafPointer.value_, leafNode);
                    fixr(stakeNode, leafPointer.value_, leafNode);

                    // Place stake where current was and
                    stakeNode.parent = leafParent; // Set parent
                    setChild(nodes[leafParent.value_], leafPointer.value_, stakeNodePointer.value_);

                    // Update all the values
                    applyStakeChange(key, stakeNode, -leafNode.amount, leafNode.parent.value_);

                    // Remove reference to the old stake
                    removeChild(nodes[leafParent.value_], stakeNodePointer.value_);
                } else if (stakeNode.left.value_ == leafPointer.value_) {
                    fixr(stakeNode, leafPointer.value_, leafNode);
                } else {
                    fixl(stakeNode, leafPointer.value_, leafNode);
                }

                // Update the root if thats changed
                if (leafNode.parent.value_ == bytes32(0)) {
                    root = leafPointer;
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

    // Cancel unlocking a certain amount of stake.
    function cancelUnlocking(uint256 amount, address stakee) public {
        bytes32 key = getKey(msg.sender, stakee);
        pullUnlocking(key, amount);
        addStake_(amount, stakee);
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

    function name(NodePointer storage p) private view returns (bytes32) {
        return p.value_;
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

    // Recursively update left/right amounts of stakes for parents of an updated stake.
    //
    // TODO Are recursive smart contracts subject to attack?
    function applyStakeChange(bytes32 key, Node storage node, uint256 amount, bytes32 root_) private {
        NodePointer storage parent = node.parent;

        if (parent.value_ == root_) {
            // We are at the root, theres nothing left to update
            return;
        }

        Node storage parentNode = nodes[parent.value_];

        if (parentNode.left.value_ == key) {
            parentNode.leftAmount += amount;
        } else {
            parentNode.rightAmount += amount;
        }

        return applyStakeChange(parent.value_, parentNode, amount, root_);
    }

    // Move the left child of node to current
    function fixl(Node storage node, bytes32 currentKey, Node storage current) private {
        if (node.left.value_ == bytes32(0)) {
            return;
        }

        nodes[node.left.value_].parent.value_ = currentKey;
        current.left = node.left;
        current.leftAmount = node.leftAmount;
    }

    // Move the right child of stake to current
    function fixr(Node storage node, bytes32 currentKey, Node storage current) private {
        if (node.right.value_ == bytes32(0)) {
            return;
        }

        nodes[node.right.value_].parent.value_ = currentKey;
        current.right = node.right;
        current.rightAmount = node.rightAmount;
    }

    function setChild(Node storage node, bytes32 oldChild, bytes32 newChild) private {
        if (node.left.value_ == oldChild) {
            node.left.value_ = newChild;
        } else if (node.right.value_ == oldChild) {
            node.right.value_ = newChild;
        } else {
            require(node.left.value_ == oldChild
                || node.right.value_ == oldChild, "Old child cannot be changed - it does not exist");
        }
    }

    function removeChild(Node storage node, bytes32 oldChild) private {
        if (node.left.value_ == oldChild) {
            node.left.value_ == bytes32(0);
        } else if (node.right.value_ == oldChild) {
            node.right.value_ = bytes32(0);
        } else {
            require(node.left.value_ == oldChild
                || node.right.value_ == oldChild, "Old child cannot be removed - it does not exist");
        }
    }
}
