pragma solidity ^0.6.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/ownership/Ownable.sol";

contract SyloTicketing is Ownable {

    /* ERC 20 compatible token we are dealing with */
    IERC20 _token;

    /*
     * The number of blocks a user must wait after calling "unlock"
     * before they can withdraw their funds
     */
    uint256 public unlockDuration;

    struct Deposit {
        uint256 escrow; // Balance of users escrow
        uint256 penalty; // Balance of users penalty

        uint256 unlockAt; // Block a user can withdraw their balances
    }

    // Mapping of user deposits to their address
    mapping(address => Deposit) deposits;

    // TODO define events

    constructor(IERC20 token, uint256 _unlockDuration) public {
        _token = token;
        unlockDuration = _unlockDuration;
    }

    function setUnlockDuration(uint256 newUnlockDuration) public onlyOwner {
        unlockDuration = newUnlockDuration;
    }

    function depositEscrow(uint256 amount) public {
        Deposit storage deposit = getDeposit(msg.sender);
        require(deposit.unlockAt == 0, "Cannot deposit while unlocking");

        deposit.escrow += amount;

        _token.transferFrom(msg.sender, address(this), amount);
    }

    function depositPenalty(uint256 amount) public {
        Deposit storage deposit = getDeposit(msg.sender);
        require(deposit.unlockAt == 0, "Cannot deposit while unlocking");

        deposit.penalty += amount;

        _token.transferFrom(msg.sender, address(this), amount);
    }

    // Unlock deposits, starting the withdrawl process
    function unlock() public returns (uint256) {

        Deposit storage deposit = getDeposit(msg.sender);
        require(deposit.escrow > 0 || deposit.penalty > 0, "Nothing to withdraw");
        require(deposit.unlockAt == 0, "Unlock already in progress");

        deposit.unlockAt = block.number + unlockDuration;

        return deposit.unlockAt;
    }

    // Cancel the withdrawl process
    function lock() public {

        Deposit storage deposit = getDeposit(msg.sender);
        require(deposit.unlockAt != 0, "Not unlocking, cannot lock");

        deposit.unlockAt = 0;
    }

    // Complete the withdrawl process and withdraw the deposits
    function withdraw() public {

        Deposit storage deposit = getDeposit(msg.sender);
        require(deposit.unlockAt > 0, "Deposits not unlocked");
        require(deposit.unlockAt >= block.number, "Unlock period not complete"); // TODO should this be '>' or '>='

        uint256 amount = deposit.escrow + deposit.penalty;

        // Set values to 0
        deposit.escrow = 0;
        deposit.penalty = 0;

        // Re-lock so if more funds are deposited they must be unlocked again
        deposit.unlockAt = 0;

        _token.transfer(msg.sender, amount);
    }

    function getDeposit(address account) private view returns (Deposit storage) {
        return deposits[account];
    }

    // Supplementry to getDeposit because ABI decoder only has experimental support for returning structs
    function getDepositDetails(address account) public view returns (uint256 escrow, uint256 penalty, uint256 unlockAt) {
        Deposit storage deposit = getDeposit(account);

        return (deposit.escrow, deposit.penalty, deposit.unlockAt);
    }
}