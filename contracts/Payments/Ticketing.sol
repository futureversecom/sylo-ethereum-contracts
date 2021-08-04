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
        uint256 generationBlock; // Block number the ticket was generated
        bytes32 receiverRandHash; // keccak256 hash of receivers random value
        uint32 senderNonce;
    }

    /* ERC 20 compatible token we are dealing with */
    IERC20 _token;

    /* Sylo Listings contract */
    Listings _listings;

    /* Sylo Directory contract */
    StakingManager _stakingManager;

    // The value of a winning ticket
    uint256 public faceValue;

    // The chance of a ticket winning
    uint256 public baseWinProb;

    // A percentage value representing the proportion of the base win probability
    // that will be decayed once a ticket has expired.
    // Example: 80% decayRate indicates that a ticket will retain 20% of its
    // base win probability once it has expired.
    uint32 public decayRate;

    // A constant value added to the probability calculated from
    // multiplying the base win probability by the decay rate
    uint256 public minProbConstant;

    // The length in blocks before a ticket is considered expired.
    // The default initialization value is 80,000. This equates
    // to roughly two weeks (15s per block).
    uint256 public ticketLength;

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

    function initialize(
        IERC20 token, 
        Listings listings, 
        StakingManager stakingManager, 
        uint256 _unlockDuration,
        uint256 _faceValue,
        uint256 _baseWinProb,
        uint8 _decayRate,
        uint256 _minProbConstant,
        uint256 _ticketLength
    ) public initializer {
        OwnableUpgradeable.__Ownable_init();
        _token = token;
        _listings = listings;
        _stakingManager = stakingManager;
        unlockDuration = _unlockDuration;
        faceValue = _faceValue;
        baseWinProb = _baseWinProb;
        decayRate = _decayRate;
        minProbConstant = _minProbConstant;
        ticketLength = _ticketLength;
    }

    function setUnlockDuration(uint256 newUnlockDuration) public onlyOwner {
        unlockDuration = newUnlockDuration;
    }

    function setFaceValue(uint256 _faceValue) public onlyOwner {
        faceValue = _faceValue;
    }

    function setWinProb(uint256 _baseWinProb) public onlyOwner {
        baseWinProb = _baseWinProb;
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

        uint256 payout = calculatePayout(ticket);

        require(
            deposit.escrow + deposit.penalty >= payout,
            "Sender doesn't have enough funds to pay"
        );

        usedTickets[ticketHash] = true;

        Listings.Listing memory listing = _listings.getListing(ticket.receiver);
        require(listing.initialized == true, "Ticket receiver must have a valid listing");

        uint256 totalStake = _stakingManager.totalStakes(ticket.receiver);
        require(totalStake != 0, "Ticket receiver must have stake");

        if (faceValue > deposit.escrow) {
            _token.transfer(ticket.receiver, deposit.escrow);
            _token.transfer(address(_token), deposit.penalty);

            deposit.escrow = 0;
            deposit.penalty = 0;
        } else {
            deposit.escrow = deposit.escrow.sub(payout);

            uint256 stakersPayout = SyloUtils.percOf(payout, listing.payoutPercentage);
            
            address[] memory stakers = _stakingManager.getStakers(ticket.receiver);

            // Track any value lost from precision due to rounding down
            uint256 stakersPayoutRemainder = stakersPayout;
            for (uint32 i = 0; i < stakers.length; i++) {
                StakingManager.Stake memory stake = _stakingManager.getStake(stakers[i], ticket.receiver);
                uint256 stakerPayout = stake.amount * stakersPayout / totalStake;
                stakersPayoutRemainder -= stakerPayout;
                _token.transfer(stakers[i], stakerPayout);
            }

            // payout any remainder to the stakee
            uint256 stakeePayout = payout - stakersPayout + stakersPayoutRemainder;
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

        require(!usedTickets[ticketHash], "Ticket already redeemed");
        require(
            keccak256(abi.encodePacked(receiverRand)) == ticket.receiverRandHash,
            "Hash of receiverRand doesn't match receiverRandHash"
        );

        require(isValidTicketSig(sig, ticket.sender, ticketHash), "Ticket doesn't have a valid signature");

        uint256 realWinProb = calculateWinningProbability(ticket);
        require(isWinningTicket(sig, receiverRand, realWinProb), "Ticket is not a winner");
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

    function calculateWinningProbability(
        Ticket memory ticket
    ) public view returns (uint256) {
        uint256 elapsedDuration = block.number - ticket.generationBlock;
        uint256 elapsedPercentage = SyloUtils.toPerc(elapsedDuration, ticketLength);

        // Ticket has completely expired, preventing any chance of winning
        if (elapsedPercentage >= 100) {
            return 0;
        }

        uint256 maxDecayValue = SyloUtils.percOf(baseWinProb, decayRate);
        uint256 decayedProbability = baseWinProb - SyloUtils.percOf(maxDecayValue, elapsedPercentage);

        // add the minimum probability constant but avoid overflow
        if (type(uint256).max - minProbConstant < decayedProbability) {
            return type(uint256).max;
        } else {
            return decayedProbability + minProbConstant;
        }
    }

    function calculatePayout(
        Ticket memory ticket
    ) public view returns (uint256) {
        uint256 expiry = ticket.generationBlock + ticketLength;
        if (expiry <= block.number) {
            // ticket has expired
            return 0;
        } else {
            // determine the number of blocks the ticket will remain alive for
            uint256 remainingLifetime = ticketLength - (block.number - ticket.generationBlock);

            // determine the payout value
            uint256 payout = remainingLifetime * faceValue / ticketLength;
            return payout;
        }
    }

    function getTicketHash(Ticket memory ticket) public view returns (bytes32) {
        return keccak256(
            abi.encodePacked(
                ticket.sender,
                ticket.receiver,
                faceValue,
                baseWinProb,
                ticket.generationBlock,
                ticket.receiverRandHash,
                ticket.senderNonce
            )
        );
    }
}
