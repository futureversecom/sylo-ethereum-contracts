// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "../contracts/Token.sol";

/*
 * Directory for accounts that wish to offer services to the Sylo Network
 * It provides the ability for accounts to stake to become listed
 * It also provides the functionality for a client to get a random stake weighted selected service peer
*/
contract Directory is Initializable, OwnableUpgradeable {

    uint32 constant DELEGATED_STAKER_CAP = 10;

    struct StakePointer {
        bytes32 value_;
    }

    struct Stake {
        uint256 amount; // Amount of the stake

        uint256 leftAmount; // Value of stake on left branch
        uint256 rightAmount; // Value of stake on right branch

        address stakee; // Address of peer that offers services

        StakePointer parent; // Pointear to parent
        StakePointer left; // Pointer to left child
        StakePointer right; // Pointer to right child
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

    mapping(bytes32 => Stake) public stakes;

    // Keeps track of all addresses staked to a stakee
    mapping(address => address[]) public stakers;

    // Keeps track of stakees stake amount
    mapping(address => uint256) public stakees;

    // Funds that are in the process of being unlocked
    mapping(bytes32 => Unlock) public unlockings;

    StakePointer root;

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

    function addStake_(uint256 amount, address stakee) private {
        require(stakee != address(0), "Address is null");
        require(amount != 0, "Cannot stake nothing");

        address staker = msg.sender;
        bytes32 key = getKey(staker, stakee);

        Stake storage stake = stakes[key];

        // New stake
        if (stake.amount == 0) {
            // The number of stakers allowed is capped
            require(
                stakers[stakee].length <= DELEGATED_STAKER_CAP, 
                "This node has reached its delegated staker cap"
            );
            // Find the node to add a new child
            //
            // There is no node for this key in the stake tree, so we will need
            // to traverse the tree and find a place to add the stake.
            StakePointer storage p = root;
            StakePointer memory parent;
            parent.value_ = bytes32(0);

            while(p.value_ != bytes32(0)) {
                parent = p;
                Stake storage current = stakes[parent.value_];
                // This will fill the right side of the tree first when
                // there is a tie.
                p = current.leftAmount < current.rightAmount
                    ? current.left
                    : current.right;
            }
            stake.parent = parent;
            p.value_ = key;
            stake.stakee = stakee;

            stakers[stakee].push(staker);
        }

        if (stake.parent.value_ == bytes32(0)) {
            // This is the first node, so we need to define the root.
            root.value_ = key;
        }

        updateStakeAmount(key, stake, amount);
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
        Stake storage stake = stakes[key];
        StakePointer memory stakeNodePointer;
        stakeNodePointer.value_ = key;

        require(stake.amount > 0, "Nothing to unstake");
        require(stake.amount >= amount, "Cannot unlock more than staked");

        // Unchecked here will prevent solidity from panicking due to overflow on uint256
        updateStakeAmount(stakeNodePointer.value_, stake, type(uint256).max - amount + 1 );

        // All stake being withdrawn, update the tree
        if (stake.amount == 0) {
            if (stake.leftAmount + stake.rightAmount == 0) {
                // No children have any stake value (i.e. no children)
                if (stake.parent.value_ != bytes32(0)) {
                    // Stake is a leaf, we need to disconnect it from the parent
                    removeChild(stakes[stake.parent.value_], stakeNodePointer.value_);
                } else {
                    // The only staker is removed, reset root
                    root.value_ = bytes32(0);
                }
            } else {
                StakePointer memory subtreePointer = stake.leftAmount > stake.rightAmount
                    ? stake.left
                    : stake.right;
                Stake storage subtreeNode = stakes[subtreePointer.value_];
                StakePointer memory child = subtreePointer;

                // Find the leaf of the most valuable path of the child
                while(true) {
                    StakePointer memory next = subtreeNode.leftAmount > subtreeNode.rightAmount
                        ? subtreeNode.left
                        : subtreeNode.right;
                    if (next.value_ == bytes32(0)) {
                        break;
                    }
                    child = next;
                    subtreeNode = stakes[next.value_];
                }

                StakePointer memory currentParent = subtreeNode.parent;
                Stake storage current = subtreeNode;

                if (stakeNodePointer.value_ != root.value_) {
                    // Update the child of the stake's parent to reference the new child
                    // value
                    Stake storage parent = stakes[stake.parent.value_];
                    setChild(parent, key, child.value_);
                }

                current.parent = stake.parent;

                // Update the children of current to be that of what the removed stake was
                if (currentParent.value_ != key) {

                    // Move the children of stake to current
                    fixr(stake, child.value_, current);
                    fixl(stake, child.value_, current);

                    // We set the parent of the stake to the current parent
                    // so we can recurively update the amount values
                    stake.parent = currentParent;
                    setChild(stakes[currentParent.value_], child.value_, stakeNodePointer.value_);

                    // Update all values starting from the stake node now that it is a leaf
                    applyStakeChange(key, stake, type(uint256).max - current.amount + 1, current.parent.value_); 

                    // Remove reference to the old stake
                    removeChild(stakes[currentParent.value_], stakeNodePointer.value_);
                } else if (stake.left.value_ == child.value_) {
                    fixr(stake, child.value_, current);
                } else {
                    fixl(stake, child.value_, current);
                } 

                // Update the root if thats changed
                if (current.parent.value_ == bytes32(0)) {
                    root = child;
                }
            }

            // Now that the node is unlinked from any other nodes, we can remove it
            delete stakes[key];

            // Also delete the reference to the staker
            address[] storage _stakers = stakers[stakee];
            for (uint32 i = 0; i < _stakers.length; i++) {
                if (_stakers[i] == msg.sender) {
                    _stakers[i] = _stakers[_stakers.length - 1];
                    _stakers.pop();
                    break;
                }
            }
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

        while(true) {

            Stake storage stake = stakes[current];

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

    function name(StakePointer storage p) private view returns (bytes32) {
        return p.value_;
    }

    function getKey(address staker, address stakee) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(staker, stakee));
    }

    function getStake(address stakee) private view returns (Stake storage) {
        return stakes[getKey(msg.sender, stakee)];
    }

    function getStake(address stakee, address staker) public view returns (Stake memory) {
        return stakes[getKey(staker, stakee)];
    }

    function getStakers(address stakee) public view returns (address[] memory) {
        return stakers[stakee];
    }

    function getTotalStake() public view returns (uint256) {
        if (root.value_ == bytes32(0)) {
            return 0;
        }

        Stake storage stake = stakes[root.value_];

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
        // unchecked here to allow uint to wrap on overflow
        unchecked {
            stake.amount += amount;
            stakees[stake.stakee] += amount;
        }

        applyStakeChange(key, stake, amount, bytes32(0));
    }

    // Recursively update left/right amounts of stakes for parents of an updated stake
    //
    // TODO Are recursive smart contracts subject to attack?
    function applyStakeChange(bytes32 key, Stake storage node, uint256 amount, bytes32 root_) private {
        StakePointer storage parentKey = node.parent;

        if (parentKey.value_ == root_) {
            // We are at the root, theres nothing left to update
            return;
        }

        Stake storage parent = stakes[parentKey.value_];

        // unchecked here to allow uint to wrap on overflow
        unchecked {
            if (parent.left.value_ == key) {
                parent.leftAmount += amount;
            } else {
                parent.rightAmount += amount;
            }
        }
        
        return applyStakeChange(parentKey.value_, parent, amount, root_);
    }

    // Move the left child of stake to current
    function fixl(Stake storage stake, bytes32 currentKey, Stake storage current) private {
        if (stake.left.value_ == bytes32(0)) {
            return;
        }

        stakes[stake.left.value_].parent.value_ = currentKey;
        current.left = stake.left;
        current.leftAmount = stake.leftAmount;
    }

    // Move the right child of stake to current
    function fixr(Stake storage stake, bytes32 currentKey, Stake storage current) private {
        if (stake.right.value_ == bytes32(0)) {
            return;
        }

        stakes[stake.right.value_].parent.value_ = currentKey;
        current.right = stake.right;
        current.rightAmount = stake.rightAmount;
    }

    function setChild(Stake storage stake, bytes32 oldChild, bytes32 newChild) private {
        if (stake.left.value_ == oldChild) {
            stake.left.value_ = newChild;
        } else if (stake.right.value_ == oldChild) {
            stake.right.value_ = newChild;
        } else {
            require(stake.left.value_ == oldChild
                || stake.right.value_ == oldChild, "Old child cannot be changed - it does not exist");
        }
    }

    function removeChild(Stake storage stake, bytes32 oldChild) private {
        if (stake.left.value_ == oldChild) {
            stake.left.value_ == bytes32(0);
        } else if (stake.right.value_ == oldChild) {
            stake.right.value_ = bytes32(0);
        } else {
            require(stake.left.value_ == oldChild
                || stake.right.value_ == oldChild, "Old child cannot be removed - it does not exist");
        }
    }
}
