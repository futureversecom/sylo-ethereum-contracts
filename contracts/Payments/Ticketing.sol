// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.13;

import "../Registries.sol";
import "../Staking/Directory.sol";
import "../Staking/Manager.sol";
import "../ECDSA.sol";
import "../Utils.sol";
import "../Epochs/Manager.sol";
import "./Ticketing/RewardsManager.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/math/SafeCast.sol";

/**
 * @notice The SyloTicketing contract manages the Probabilistic
 * Micro-Payment Ticketing system that pays Nodes for providing the
 * Event Relay service.
 */
contract SyloTicketing is Initializable, OwnableUpgradeable {
    /**
     * The maximum probability value, where probability is represented
     * as an integer between 0 to 2^128 - 1.
     */
    uint128 internal constant MAX_PROB = type(uint128).max;

    struct Deposit {
        uint256 escrow; // Balance of users escrow
        uint256 penalty; // Balance of users penalty
        uint256 unlockAt; // Block number a user can withdraw their balances
    }

    struct Ticket {
        uint256 epochId; // The epoch this ticket is associated with
        address sender; // Address of the ticket sender
        address redeemer; // Address of the intended recipient
        uint256 generationBlock; // Block number the ticket was generated
        bytes32 senderCommit; // Hash of the secret random number of the sender
        bytes32 redeemerCommit; // Hash of the secret random number of the redeemer
    }

    event Redemption(
        uint256 epochId,
        address sender,
        address redeemer,
        uint256 generationBlock,
        uint256 amount
    );

    /** ERC20 Sylo token contract.*/
    IERC20 public _token;

    /** Sylo Registries contract */
    Registries public _registries;

    /** Sylo Staking Manager contract */
    StakingManager public _stakingManager;

    /** Sylo Directory contract */
    Directory public _directory;

    /** Rewards Manager contract */
    RewardsManager public _rewardsManager;

    /**
     * @notice Sylo Epochs Manager.
     * @dev The ticketing parameters used when redeeming tickets
     * will be read from this contract.
     */
    EpochsManager public _epochsManager;

    event UnlockDurationUpdated(uint256 unlockDuration);

    /**
     * @notice The number of blocks a user must wait after calling "unlock"
     * before they can withdraw their funds.
     */
    uint256 public unlockDuration;

    /** @notice Mapping of user deposits */
    mapping(address => Deposit) public deposits;

    /** @notice Mapping of ticket hashes, used to check if a ticket has been redeemed */
    mapping(bytes32 => bool) public usedTickets;

    function initialize(
        IERC20 token,
        Registries registries,
        StakingManager stakingManager,
        Directory directory,
        EpochsManager epochsManager,
        RewardsManager rewardsManager,
        uint256 _unlockDuration
    ) external initializer {
        OwnableUpgradeable.__Ownable_init();
        _token = token;
        _registries = registries;
        _stakingManager = stakingManager;
        _directory = directory;
        _epochsManager = epochsManager;
        _rewardsManager = rewardsManager;
        unlockDuration = _unlockDuration;
    }

    /**
     * @notice Set the unlock duration for deposits. Only callable
     * by the owner.
     * @param _unlockDuration The unlock duration in blocks.
     */
    function setUnlockDuration(uint256 _unlockDuration) external onlyOwner {
        unlockDuration = _unlockDuration;
        emit UnlockDurationUpdated(_unlockDuration);
    }

    /**
     * @notice Use this function to deposit funds into the
     * escrow. This will fail if the deposit is currently being
     * unlocked.
     * @param amount The amount in SOLO to add to the escrow.
     * @param account The address of the account holding the escrow.
     */
    function depositEscrow(uint256 amount, address account) external {
        Deposit storage deposit = getDeposit(account);
        require(deposit.unlockAt == 0, "Cannot deposit while unlocking");

        deposit.escrow += amount;

        _token.transferFrom(msg.sender, address(this), amount);
    }

    /**
     * @notice Use this function to deposit funds into the
     * penalty. This will fail if the deposit is currently being
     * unlocked.
     * @param amount The amount in SOLO to add to the escrow.
     * @param account The address of the account holding the penalty.
     */
    function depositPenalty(uint256 amount, address account) external {
        Deposit storage deposit = getDeposit(account);
        require(deposit.unlockAt == 0, "Cannot deposit while unlocking");

        deposit.penalty += amount;

        _token.transferFrom(msg.sender, address(this), amount);
    }

    /**
     * @notice Call this function to begin unlocking deposits. This function
     * will fail if no deposit exists, or if the unlock process has
     * already begun.
     */
    function unlockDeposits() external returns (uint256) {
        Deposit storage deposit = getDeposit(msg.sender);
        require(deposit.escrow > 0 || deposit.penalty > 0, "Nothing to withdraw");
        require(deposit.unlockAt == 0, "Unlock already in progress");

        deposit.unlockAt = block.number + unlockDuration;

        return deposit.unlockAt;
    }

    /**
     * @notice Call this function to cancel any deposit that is in the
     * unlocking process.
     */
    function lockDeposits() external {
        Deposit storage deposit = getDeposit(msg.sender);
        require(deposit.unlockAt != 0, "Not unlocking, cannot lock");

        deposit.unlockAt = 0;
    }

    /**
     * @notice Call this function once the unlock duration has
     * elapsed in order to transfer the unlocked tokens to the caller's account.
     */
    function withdraw() external {
        return withdrawTo(msg.sender);
    }

    /**
     * @notice Call this function once the unlock duration has
     * elapsed in order to transfer the unlocked tokens to the specified
     * account.
     * @param account The address of the account the tokens should be
     * transferred to.
     */
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

    /**
     * @notice Nodes should call this function on completing an event
     * delivery and having the sender rand revealed. This function will fail if
     * the ticket is invalid or if the ticket is not a winner. Clients should
     * calculate if the ticket is a winner locally, but can also us the public view
     * functions: `requireValidWinningTicket` and `isWinningTicket` to check
     * that a ticket is winning.
     * @param ticket The ticket issued by the sender.
     * @param senderRand The sender random value, revealed on completing an event
     * relay.
     * @param redeemerRand The redeemer random value, generated by the Node prior
     * to performing the event relay.
     * @param sig The signature of the sender of the ticket.
     */
    function redeem(
        Ticket memory ticket,
        uint256 senderRand,
        uint256 redeemerRand,
        bytes memory sig
    ) external {
        EpochsManager.Epoch memory epoch = _epochsManager.getEpoch(ticket.epochId);
        require(
            ticket.generationBlock <= block.number,
            "The ticket cannot be generated for a future block"
        );

        bytes32 ticketHash = getTicketHash(ticket);
        requireValidWinningTicket(ticket, ticketHash, senderRand, redeemerRand, sig);

        uint256 directoryStake = _directory.getTotalStakeForStakee(
            ticket.epochId,
            ticket.redeemer
        );
        require(
            directoryStake > 0,
            "Ticket redeemer must have joined the directory for this epoch"
        );

        usedTickets[ticketHash] = true;

        uint256 rewardAmount = rewardRedeemer(epoch, ticket);

        emit Redemption(
            ticket.epochId,
            ticket.sender,
            ticket.redeemer,
            ticket.generationBlock,
            rewardAmount
        );
    }

    function rewardRedeemer(EpochsManager.Epoch memory epoch, Ticket memory ticket)
        internal
        returns (uint256)
    {
        Deposit storage deposit = getDeposit(ticket.sender);

        uint256 amount = 0;

        if (epoch.faceValue > deposit.escrow) {
            amount = deposit.escrow;
            incrementRewardPool(ticket.redeemer, deposit, amount);
            _token.transfer(address(0x000000000000000000000000000000000000dEaD), deposit.penalty);

            deposit.penalty = 0;
        } else {
            amount = epoch.faceValue;
            incrementRewardPool(ticket.redeemer, deposit, amount);
        }

        return amount;
    }

    /**
     * @notice Call this function to check if a ticket is valid and is
     * a winning ticket. It will fail if the ticket is invalid or is not
     * a winner. A ticket is invalid if:
     *      - The sender or redeemer addresses are null
     *      - The ticket has already been redeemed.
     *      - The secret random value of the sender does not match the commit
     *        in the ticket.
     *      - The secret random value of the redeemer does not match the commit
     *        in the ticket.
     *      - The signature is invalid.
     * @param ticket The ticket issued by the sender.
     * @param ticketHash The hash of the ticket. Should match the hash generated
     * by `getTicketHash`.
     * @param senderRand The sender random value, revealed on completing an event
     * relay.
     * @param redeemerRand The redeemer random value, generated by the Node prior
     * to performing the event relay.
     * @param sig The signature of the sender of the ticket.
     */
    function requireValidWinningTicket(
        Ticket memory ticket,
        bytes32 ticketHash,
        uint256 senderRand,
        uint256 redeemerRand,
        bytes memory sig
    ) public view {
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

        require(
            isValidTicketSig(sig, ticket.sender, ticketHash),
            "Ticket doesn't have a valid signature"
        );

        require(isWinningTicket(sig, ticket, senderRand, redeemerRand), "Ticket is not a winner");
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

    /**
     * @notice Use this function to check if a ticket is winning.
     * @param sig The signature of the sender of the ticket.
     * @param ticket The ticket issued by the sender, which holds the various ticketing parameters.
     * @param senderRand The sender random value, revealed on completing an event
     * relay.
     * @param redeemerRand The redeemer random value, generated by the Node prior
     * to performing the event relay.
     * @return True if a ticket is a winner.
     */
    function isWinningTicket(
        bytes memory sig,
        Ticket memory ticket,
        uint256 senderRand,
        uint256 redeemerRand
    ) public view returns (bool) {
        uint256 winProb = calculateWinningProbability(ticket);
        // bitshift the winProb to a 256 bit value to allow comparison to a 32 byte hash
        uint256 prob = (uint256(winProb) << 128) | uint256(winProb);
        return uint256(keccak256(abi.encodePacked(sig, senderRand, redeemerRand))) < prob;
    }

    /**
     * @notice This function calculates the probability of a ticket winning at
     * the block that this function was called. A ticket's winning probability
     * will decay every block since its issuance. The amount of decay will depend
     * on the decay rate parameter of the epoch the ticket was generated in.
     * @param ticket The ticket issued by the sender.
     */
    function calculateWinningProbability(Ticket memory ticket) public view returns (uint128) {
        EpochsManager.Epoch memory epoch = _epochsManager.getEpoch(ticket.epochId);
        require(epoch.startBlock > 0, "Ticket's associated epoch does not exist");
        require(
            ticket.generationBlock >= epoch.startBlock &&
                (epoch.endBlock > 0 ? ticket.generationBlock < epoch.endBlock : true),
            "This ticket was not generated during it's associated epoch"
        );

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
        uint256 decayedProbability = (maxDecayValue * elapsedDuration) / epoch.ticketDuration;

        // calculate the remaining probability by subtracting the decayed probability
        // from the base
        return epoch.baseLiveWinProb - SafeCast.toUint128(decayedProbability);
    }

    /**
     * @notice Returns the hash of the ticket. Takes all fields in a ticket
     * as inputs to the hash.
     * @return A byte-array representing the hash.
     */
    function getTicketHash(Ticket memory ticket) public pure returns (bytes32) {
        return
            keccak256(
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

    function incrementRewardPool(
        address stakee,
        Deposit storage deposit,
        uint256 amount
    ) internal {
        deposit.escrow = deposit.escrow - amount;

        _token.transfer(address(_rewardsManager), amount);
        _rewardsManager.incrementRewardPool(stakee, amount);
    }
}
