// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "../Listings.sol";
import "../Staking/Manager.sol";
import "../ECDSA.sol";
import "../Utils.sol";
import "../Epochs/Manager.sol";
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
        bytes32 epochId; // The epoch this ticket is associated with
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

    /* Sylo Epochs Manager.
     * This contract holds various ticketing parameters
     */
    EpochsManager _epochsManager;

    /*
     * The number of blocks a user must wait after calling "unlock"
     * before they can withdraw their funds
     */
    uint256 public unlockDuration;

    /* Mapping of user deposits to their address */
    mapping(address => Deposit) public deposits;

    /* Mapping of ticket hashes, used to check if ticket has been redeemed */
    mapping (bytes32 => bool) public usedTickets;

    // TODO define events

    function initialize(
        IERC20 token,
        Listings listings,
        StakingManager stakingManager,
        EpochsManager epochsManager,
        uint256 _unlockDuration
    ) public initializer {
        OwnableUpgradeable.__Ownable_init();
        _token = token;
        _listings = listings;
        _stakingManager = stakingManager;
        _epochsManager = epochsManager;
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
        uint256 senderRand,
        uint256 redeemerRand,
        bytes memory sig
    ) public {
        EpochsManager.Epoch memory epoch = _epochsManager.getEpoch(ticket.epochId);
        require(epoch.startBlock > 0, "Ticket's associated epoch does not exist");
        require(
            ticket.generationBlock >= epoch.startBlock &&
                (epoch.endBlock > 0 ? ticket.generationBlock < epoch.endBlock : true),
            "This ticket was not generated during it's associated epoch"
        );

        bytes32 ticketHash = getTicketHash(ticket);

        requireValidWinningTicket(ticket, ticketHash, senderRand, redeemerRand, sig, epoch);

        Listings.Listing memory listing = _listings.getListing(ticket.redeemer);
        require(listing.initialized == true, "Ticket redeemer must have a valid listing");

        usedTickets[ticketHash] = true;

        uint256 totalStake = _stakingManager.totalStakes(ticket.redeemer);
        require(totalStake != 0, "Ticket redeemer must have stake");

        rewardRedeemer(epoch, ticket, listing, totalStake);
    }

    function rewardRedeemer(
        EpochsManager.Epoch memory epoch,
        Ticket memory ticket,
        Listings.Listing memory listing,
        uint256 totalStake
    ) internal {
        Deposit storage deposit = getDeposit(ticket.sender);

        if (epoch.faceValue > deposit.escrow) {
            _token.transfer(ticket.redeemer, deposit.escrow);
            _token.transfer(address(0x000000000000000000000000000000000000dEaD), deposit.penalty);

            deposit.escrow = 0;
            deposit.penalty = 0;
        } else {
            deposit.escrow = deposit.escrow - epoch.faceValue;

            // We can safely cast faceValue to 128 bits as all Sylo Tokens
            // would fit within 94 bits
            uint256 stakersPayout = SyloUtils.percOf(uint128(epoch.faceValue), listing.payoutPercentage);

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
            uint256 stakeePayout = epoch.faceValue - stakersPayout + stakersPayoutRemainder;
            _token.transfer(ticket.redeemer, stakeePayout);
        }
    }

    function requireValidWinningTicket(
        Ticket memory ticket,
        bytes32 ticketHash,
        uint256 senderRand,
        uint256 redeemerRand,
        bytes memory sig,
        EpochsManager.Epoch memory epoch
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

        uint128 remainingProbability = calculateWinningProbability(ticket, epoch);
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
        uint256 prob = uint256(winProb) << 128 | uint256(winProb);
        return uint256(keccak256(abi.encodePacked(sig, redeemerRand))) < prob;
    }

    function calculateWinningProbability(
        Ticket memory ticket,
        EpochsManager.Epoch memory epoch
    ) public view returns (uint128) {
        uint256 elapsedDuration = block.number - ticket.generationBlock;

        // Ticket has completely expired
        if (elapsedDuration >= epoch.ticketDuration) {
            return 0;
        }

        uint256 maxDecayValue = SyloUtils.percOf(epoch.baseLiveWinProb, epoch.decayRate);

        // determine the amount of probability that has actually decayed
        // by multiplying the maximum decay value against ratio of the tickets elapsed duration
        // vs the actual ticket duration. The max decay value is calculated from a fraction of a
        // uint128 value so we cannot phantom overflow here
        uint256 decayedProbability = maxDecayValue * elapsedDuration / epoch.ticketDuration;

        // calculate the remaining probability by substracting the decayed probability
        // from the base
        return epoch.baseLiveWinProb - uint128(decayedProbability);
    }

    function getTicketHash(Ticket memory ticket) public pure returns (bytes32) {
        return keccak256(
            abi.encodePacked(
                ticket.epochId,
                ticket.sender,
                ticket.redeemer,
                ticket.generationBlock,
                ticket.senderCommit,
                ticket.redeemerCommit
            )
        );
    }
}
