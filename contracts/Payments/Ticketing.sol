// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "../Listings.sol";
import "../Staking/Manager.sol";
import "../ECDSA.sol";
import "../Utils.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract SyloTicketing is Initializable, OwnableUpgradeable {

    /**
     * The maximum probability value, where probability is represented
     * as an integer between 0 to 2^128 - 1.
     */
    uint128 constant MAX_PROB = type(uint128).max;

    struct Deposit {
        uint256 escrow; // Balance of users escrow
        uint256 penalty; // Balance of users penalty

        uint256 unlockAt; // Block number a user can withdraw their balances
    }

    struct Ticket {
        address sender; // Address of the ticket sender
        address redeemer; // Address of the intended recipient
        uint256 generationBlock; // Block number the ticket was generated
        bytes32 senderCommit; // Hash of the secret random number of the sender
        bytes32 redeemerCommit; // Hash of the secret random number of the redeemder
    }

    /* ERC 20 compatible token we are dealing with */
    IERC20 _token;

    /* Sylo Listings contract */
    Listings _listings;

    /* Sylo Directory contract */
    StakingManager _stakingManager;

    /* The value of a winning ticket */
    uint256 public faceValue;

    /** 
     * The probability of a ticket winning during the start of its lifetime.
     * This is a uint128 value representing the numerator in the probability
     * ratio where 2^128 - 1 is the denominator.
     */
    uint128 public baseLiveWinProb;

    /** 
     * The probability of a ticket winning after it has expired.
     * This is a uint128 value representing the numerator in the probability
     * ratio where 2^128 - 1 is the denominator.
     */
    uint128 public expiredWinProb;

    /**
     * The length in blocks before a ticket is considered expired.
     * The default initialization value is 80,000. This equates
     * to roughly two weeks (15s per block).
     */
    uint256 public ticketDuration;

    /*
     * The number of blocks a user must wait after calling "unlock"
     * before they can withdraw their funds
     */
    uint256 public unlockDuration;

    /** 
     * A percentage value representing the proportion of the base win probability
     * that will be decayed once a ticket has expired.
     * Example: 80% decayRate indicates that a ticket will retain 20% of its
     * base win probability once it has expired.
     */
    uint8 public decayRate;

    /* Mapping of user deposits to their address */
    mapping(address => Deposit) public deposits;

    /* Mapping of ticket hashes, used to check if ticket has been redeemed */
    mapping (bytes32 => bool) public usedTickets;

    // TODO define events

    function initialize(
        IERC20 token, 
        Listings listings, 
        StakingManager stakingManager, 
        uint256 _unlockDuration,
        uint256 _faceValue,
        uint128 _baseLiveWinProb,
        uint128 _expiredWinProb,
        uint8 _decayRate,
        uint256 _ticketDuration
    ) public initializer {
        OwnableUpgradeable.__Ownable_init();
        _token = token;
        _listings = listings;
        _stakingManager = stakingManager;
        unlockDuration = _unlockDuration;
        faceValue = _faceValue;
        baseLiveWinProb = _baseLiveWinProb;
        expiredWinProb = _expiredWinProb;
        decayRate = _decayRate;
        ticketDuration = _ticketDuration;
    }

    function setUnlockDuration(uint256 newUnlockDuration) public onlyOwner {
        unlockDuration = newUnlockDuration;
    }

    function setFaceValue(uint256 _faceValue) public onlyOwner {
        faceValue = _faceValue;
    }

    function setBaseLiveWinProb(uint128 _baseLiveWinProb) public onlyOwner {
        baseLiveWinProb = _baseLiveWinProb;
    }

    function setExpiredWinProb(uint128 _expiredWinProb) public onlyOwner {
        expiredWinProb = _expiredWinProb;
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
        uint256 senderRand,
        uint256 redeemerRand,
        bytes memory sig
    ) public {

        bytes32 ticketHash = getTicketHash(ticket);

        requireValidWinningTicket(ticket, ticketHash, senderRand, redeemerRand, sig);

        Deposit storage deposit = getDeposit(ticket.sender);

        require(
            deposit.escrow + deposit.penalty >= faceValue,
            "Sender doesn't have enough funds to pay"
        );

        usedTickets[ticketHash] = true;

        Listings.Listing memory listing = _listings.getListing(ticket.redeemer);
        require(listing.initialized == true, "Ticket redeemer must have a valid listing");

        uint256 totalStake = _stakingManager.totalStakes(ticket.redeemer);
        require(totalStake != 0, "Ticket redeemer must have stake");

        if (faceValue > deposit.escrow) {
            _token.transfer(ticket.redeemer, deposit.escrow);
            _token.transfer(address(_token), deposit.penalty);

            deposit.escrow = 0;
            deposit.penalty = 0;
        } else {
            deposit.escrow = deposit.escrow - faceValue;

            // We can safely cast faceValue to 128 bits as all Sylo Tokens
            // would fit within 94 bits
            uint256 stakersPayout = SyloUtils.percOf(uint128(faceValue), listing.payoutPercentage);
            
            address[] memory stakers = _stakingManager.getStakers(ticket.redeemer);

            // Track any value lost from precision due to rounding down
            uint256 stakersPayoutRemainder = stakersPayout;
            for (uint32 i = 0; i < stakers.length; i++) {
                StakingManager.Stake memory stake = _stakingManager.getStake(stakers[i], ticket.redeemer);
                uint256 stakerPayout = stake.amount * stakersPayout / totalStake;
                stakersPayoutRemainder -= stakerPayout;
                _token.transfer(stakers[i], stakerPayout);
            }

            // payout any remainder to the stakee
            uint256 stakeePayout = faceValue - stakersPayout + stakersPayoutRemainder;
            _token.transfer(ticket.redeemer, stakeePayout);
        }
    }

    function requireValidWinningTicket(
        Ticket memory ticket,
        bytes32 ticketHash,
        uint256 senderRand,
        uint256 redeemerRand,
        bytes memory sig
    ) internal view {
        require(ticket.sender != address(0), "Ticket sender is null");
        require(ticket.redeemer != address(0), "Ticket redeemer is null");

        require(!usedTickets[ticketHash], "Ticket already redeemed");

        // validate that the sender's random number has been revealed to 
        // the redeemer
        require(
            keccak256(abi.encodePacked(senderRand)) == ticket.senderCommit,
            "Hash of senderRand doesn't match senderRandHash"
        );

        // validate the redeemer has knowledge of the redeemer rand
        require(
            keccak256(abi.encodePacked(redeemerRand)) == ticket.redeemerCommit,
            "Hash of redeemerRand doesn't match redeemerRandHash"
        );

        require(isValidTicketSig(sig, ticket.sender, ticketHash), "Ticket doesn't have a valid signature");

        uint128 remainingProbability = calculateWinningProbability(ticket);
        require(isWinningTicket(sig, redeemerRand, remainingProbability), "Ticket is not a winner");
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
        uint256 redeemerRand,
        uint128 winProb
    ) internal pure returns (bool) {
        // bitshift the winProb to a 256 bit value to allow comparison to a 32 byte hash
        uint256 prob =  uint256(winProb) << 128;
        return uint256(keccak256(abi.encodePacked(sig, redeemerRand))) < prob;
    }

    function calculateWinningProbability(
        Ticket memory ticket
    ) public view returns (uint128) {
        uint256 elapsedDuration = block.number - ticket.generationBlock;

        // Ticket has completely expired
        if (elapsedDuration >= ticketDuration) {
            return 0;
        }

        uint256 maxDecayValue = SyloUtils.percOf(baseLiveWinProb, decayRate);

        // determine the amount of probability that has actually decayed
        // by multiplying the maximum decay value against ratio of the tickets elapsed duration 
        // vs the actual ticket duration
        uint256 decayedProbability = maxDecayValue * elapsedDuration / ticketDuration;

        // calculate the remaining probability by substracting the decayed probability
        // from the base
        return baseLiveWinProb - uint128(decayedProbability);
    }

    function getTicketHash(Ticket memory ticket) public view returns (bytes32) {
        return keccak256(
            abi.encodePacked(
                ticket.sender,
                ticket.redeemer,
                faceValue,
                baseLiveWinProb,
                ticket.generationBlock,
                ticket.senderCommit,
                ticket.redeemerCommit
            )
        );
    }
}
