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
    uint32 constant DELEGATED_STAKER_CAP = 10;

    struct Stake {
        uint256 amount; // Amount of the stake;
        address stakee; // Address of peer that offers services;
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

    
    // Tracks all keys
    bytes32[] keys;

    // Tracks all stakes
    mapping(bytes32 => Stake) public stakes;

    // Tracks all addresses staked to a stakee
    mapping(address => address[]) public stakers;

    // Tracks total stake for each stakee
    mapping(address => uint256) public totalStakes;

    // Tracks overall total stake
    uint256 public totalStake;

    // Tracks all stakees
    address[] public stakees;

    // Tracks funds that are in the process of being unlocked
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

        address staker = msg.sender;
        bytes32 key = getKey(staker, stakee);

        Stake storage stake = stakes[key];

        // New stake
        if (stake.amount == 0) {
            require(
                stakers[stakee].length < DELEGATED_STAKER_CAP, 
                "This node has reached its delegated staker cap"
            );
            stakers[stakee].push(staker);

            stake.amount = amount;
            stake.stakee = stakee;

            stakees.push(stakee);
        } else {
            stake.amount += amount;
        }

        totalStakes[stakee] += amount;
        totalStake += amount;
    }

    function unlockStake(uint256 amount, address stakee) public {
        bytes32 key = getKey(msg.sender, stakee);
        Stake storage stake = stakes[key];

        require(stake.amount > 0, "Nothing to unstake");
        require(stake.amount >= amount, "Cannot unlock more than staked");

        stake.amount -= amount;

        if (stake.amount == 0) {
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

        totalStakes[stakee] -= amount;
        totalStake -= amount;
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

    function getCountOfStakees() public view returns (uint count) {
        return stakees.length;
    }

    function getTotalStake() public view returns (uint256) {
        return totalStake;
    }

    function getStake(address staker, address stakee) public view returns (Stake memory) {
        return stakes[getKey(staker, stakee)];
    }

    function getStakers(address stakee) public view returns (address[] memory) {
        return stakers[stakee];
    }
}