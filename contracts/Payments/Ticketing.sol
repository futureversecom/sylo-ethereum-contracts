// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "../Listings.sol";
import "../Directory.sol";
import "../ECDSA.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract SyloTicketing is Initializable, OwnableUpgradeable {

    uint256 constant PERC_DIVISOR = 100;

    using SafeMath for uint256;

    struct Deposit {
        uint256 escrow; // Balance of users escrow
        uint256 penalty; // Balance of users penalty

        uint256 unlockAt; // Block number a user can withdraw their balances
    }

    // Properties are ordered to decrease storage size
    // https://solidity.readthedocs.io/en/v0.6.4/miscellaneous.html#layout-of-state-variables-in-storage
    struct Ticket {
        address sender; // Address of the ticket sender
        address receiver; // Address of the intended recipient
        uint256 faceValue; // The value of a winning ticket
        uint256 winProb; // The chance of a ticket winning
        uint256 expirationBlock; // Block number the ticket is valid until
        bytes32 receiverRandHash; // keccak256 hash of receivers random value
        uint32 senderNonce; // Senders ticket counter
    }

    /* ERC 20 compatible token we are dealing with */
    IERC20 _token;

    /* Sylo Listings contract */
    Listings _listings;

    /* Sylo Directory contract */
    Directory _directory;

    /*
     * The number of blocks a user must wait after calling "unlock"
     * before they can withdraw their funds
     */
    uint256 public unlockDuration;

    // Mapping of user deposits to their address
    mapping(address => Deposit) public deposits;

    // Mapping of ticket hashes, used to check if ticket has been redeemed
    mapping (bytes32 => bool) public usedTickets;

    // TODO define events

    function initialize(IERC20 token, Listings listings, Directory directory, uint256 _unlockDuration) public initializer {
        OwnableUpgradeable.__Ownable_init();
        _token = token;
        _listings = listings;
        _directory = directory;
        unlockDuration = _unlockDuration;
    }

    function setUnlockDuration(uint256 newUnlockDuration) public onlyOwner {
        unlockDuration = newUnlockDuration;
    }

    function depositEscrow(uint256 amount, address account) public {
        Deposit storage deposit = getDeposit(account);
        require(deposit.unlockAt == 0, "Cannot deposit while unlocking");

        deposit.escrow += amount;

        _token.transferFrom(msg.sender, address(this), amount);
    }

    function depositPenalty(uint256 amount, address account) public {
        Deposit storage deposit = getDeposit(account);
        require(deposit.unlockAt == 0, "Cannot deposit while unlocking");

        deposit.penalty += amount;

        _token.transferFrom(msg.sender, address(this), amount);
    }

    // Unlock deposits, starting the withdrawl process
    function unlockDeposits() public returns (uint256) {

        Deposit storage deposit = getDeposit(msg.sender);
        require(deposit.escrow > 0 || deposit.penalty > 0, "Nothing to withdraw");
        require(deposit.unlockAt == 0, "Unlock already in progress");

        deposit.unlockAt = block.number + unlockDuration;

        return deposit.unlockAt;
    }

    // Cancel the withdrawl process
    function lockDeposits() public {

        Deposit storage deposit = getDeposit(msg.sender);
        require(deposit.unlockAt != 0, "Not unlocking, cannot lock");

        deposit.unlockAt = 0;
    }

    function withdraw() public {
        return withdrawTo(msg.sender);
    }

    // Complete the withdrawl process and withdraw the deposits
    function withdrawTo(address account) public {

        Deposit storage deposit = getDeposit(msg.sender);
        require(deposit.unlockAt > 0, "Deposits not unlocked");
        require(deposit.unlockAt < block.number, "Unlock period not complete");

        uint256 amount = deposit.escrow + deposit.penalty;

        // Set values to 0
        deposit.escrow = 0;
        deposit.penalty = 0;

        // Re-lock so if more funds are deposited they must be unlocked again
        deposit.unlockAt = 0;

        _token.transfer(account, amount);
    }

    function redeem(
        Ticket memory ticket,
        uint256 receiverRand,
        bytes memory sig
    ) public {

        bytes32 ticketHash = getTicketHash(ticket);

        requireValidWinningTicket(ticket, ticketHash, receiverRand, sig);

        Deposit storage deposit = getDeposit(ticket.sender);

        require(
            deposit.escrow + deposit.penalty >= ticket.faceValue,
            "Sender doesn't have enough funds to pay"
        );

        usedTickets[ticketHash] = true;

        Listings.Listing memory listing = _listings.getListing(ticket.receiver);
        require(listing.initialized == true, "Ticket receiver must have a valid listing");

        uint256 totalStake = _directory.stakees(ticket.receiver);
        require(totalStake != 0, "Ticket receiver must have stake");

        if (ticket.faceValue > deposit.escrow) {
            _token.transfer(ticket.receiver, deposit.escrow);
            _token.transfer(address(_token), deposit.penalty);

            deposit.escrow = 0;
            deposit.penalty = 0;
        } else {
            deposit.escrow = deposit.escrow.sub(ticket.faceValue);

            uint256 stakersPayout = listing.payoutPercentage * ticket.faceValue / PERC_DIVISOR;
            
            address[] memory stakers = _directory.getStakers(ticket.receiver);

            // Track any value lost from precision due to rounding down
            uint256 stakersPayoutRemainder = stakersPayout;
            for (uint32 i = 0; i < stakers.length; i++) {
                Directory.Stake memory stake = _directory.getStake(ticket.receiver, stakers[i]);
                uint256 stakerPayout = stake.amount * stakersPayout / totalStake;
                stakersPayoutRemainder -= stakerPayout;
                _token.transfer(stakers[i], stakerPayout);
            }

            // payout any remainder to the stakee
            uint256 stakeePayout = ticket.faceValue - stakersPayout + stakersPayoutRemainder;
            _token.transfer(ticket.receiver, stakeePayout);
        }
    }

    function requireValidWinningTicket(
        Ticket memory ticket,
        bytes32 ticketHash,
        uint256 receiverRand,
        bytes memory sig
    ) internal view {
        require(ticket.sender != address(0), "Ticket sender is null");
        require(ticket.receiver != address(0), "Ticket receiver is null");

        require(
            ticket.expirationBlock == 0 || ticket.expirationBlock >= block.number,
            "Ticket has expired"
        );
        require(!usedTickets[ticketHash], "Ticket already redeemed");
        require(
            keccak256(abi.encodePacked(receiverRand)) == ticket.receiverRandHash,
            "Hash of receiverRand doesn't match receiverRandHash"
        );

        require(isValidTicketSig(sig, ticket.sender, ticketHash), "Ticket doesn't have a valid signature");
        require(isWinningTicket(sig, receiverRand, ticket.winProb), "Ticket is not a winner");
    }

    function getDeposit(address account) private view returns (Deposit storage) {
        return deposits[account];
    }

    function isValidTicketSig(
        bytes memory sig,
        address sender,
        bytes32 ticketHash
    ) internal pure returns (bool) {
        return ECDSA.recover(ticketHash, sig) == sender;
    }

    function isWinningTicket(
        bytes memory sig,
        uint256 receiverRand,
        uint256 winProb
    ) internal pure returns (bool) {
        return uint256(keccak256(abi.encodePacked(sig, receiverRand))) < winProb;
    }

    function getTicketHash(Ticket memory ticket) public pure returns (bytes32) {
        return keccak256(
            abi.encodePacked(
                ticket.sender,
                ticket.receiver,
                ticket.faceValue,
                ticket.winProb,
                ticket.expirationBlock,
                ticket.receiverRandHash,
                ticket.senderNonce
            )
        );
    }
}
