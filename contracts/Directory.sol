pragma solidity ^0.6.0;

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
        uint256 unlockAt; // Block number a user can withdraw their stake
    }

    uint256 constant maxU256 = 2^256-1;

    /* ERC 20 compatible token we are dealing with */
    IERC20 _token;

    /*
     * The number of blocks a user must wait after calling "unlock"
     * before they can withdraw their stake
     */
    uint256 public unlockDuration;

    mapping(address => Stake) public stakes;

    address[] stakers;

    constructor(IERC20 token, uint256 _unlockDuration) public {
        _token = token;
        unlockDuration = _unlockDuration;
    }

    function setUnlockDuration(uint256 newUnlockDuration) public onlyOwner {
        unlockDuration = newUnlockDuration;
    }

    function addStake(uint256 amount) public{
        return addStakeFor(amount, msg.sender);
    }

    function addStakeFor(uint256 amount, address staker) public {

        require(staker != address(0), "Address is null");
        require(amount != 0, "Cannot stake nothing");

        Stake storage stake = getStake(staker);

        require(stake.unlockAt == 0, "Cannot stake while unlocking");

        if (stake.amount == 0) {
            stakers.push(staker);
        }

        stake.amount += amount;

        // DO Last
        _token.transferFrom(msg.sender, address(this), amount);
    }

    /* TODO allow partial withdrawl of stake */
    function unlockStake(/*uint256 amount*/) public returns (uint256) {

        Stake storage stake = getStake(msg.sender);
        require(stake.amount > 0, "Nothing to unstake");
        require(stake.unlockAt == 0, "Already unlocking");

        stake.unlockAt = block.number + unlockDuration;

        return stake.unlockAt;
    }

    function lockStake() public {
        Stake storage stake = getStake(msg.sender);

        require(stake.unlockAt > 0, "Not unlocking cannot lock");

        stake.unlockAt = 0;
    }

    function unstake() public {
        return unstakeTo(msg.sender);
    }

    function unstakeTo(address account) public {
        Stake storage stake = getStake(msg.sender);

        // require(stake.unlockAt > 0, "Stake not unlocked");
        // require(stake.unlockAt <= block.number, "Unlock period not complete");
        require(isValidStake(stake), "Stake not withdrawable");

        uint256 amount = stake.amount;
        stake.amount = 0;
        stake.unlockAt = 0;

        // This is super inefficient, as users unstake the array doesn't get any smaller
        for (uint256 i = 0; i < stakers.length; i++){
            if (stakers[i] == msg.sender) {
                stakers[i] = address(0);
            }
        }

        _token.transfer(account, amount);
    }

    function scan(uint256 rand) public view returns (address) {

        uint256 expectedVal = getTotalStake() * rand / maxU256;
        uint256 sum;

        for (uint256 i = 0; i < stakers.length; i++) {

            Stake memory stake =stakes[stakers[i]];

            if (!isValidStake(stake)) {
                continue;
            }

            sum += stake.amount;

            if (expectedVal <= sum) {
                return stakers[i];
            }
        }

        return address(0);
    }

    function getStake(address account) private view returns (Stake storage) {
        return stakes[account];
    }

    function getTotalStake() public view returns (uint256) {
        uint256 totalStake;
        for (uint256 i = 0; i < stakers.length; i++) {
            Stake memory stake = stakes[stakers[i]];
            if (isValidStake(stake)) {
                totalStake += stake.amount;
            }
        }

        return totalStake;
    }

    function isValidStake(Stake memory stake) private view returns (bool) {
        return stake.amount > 0 && stake.unlockAt < block.number;
    }
}
