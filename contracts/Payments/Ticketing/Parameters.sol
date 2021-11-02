// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

/**
 * @dev Persists the parameters for the ticketing mechanism. This contract is
 * read by the EpochManager. Extracting the parameters into another
 * contract is necessary to avoid a cyclic dependency between the ticketing
 * and epoch contracts.
 */
contract TicketingParameters is Initializable, OwnableUpgradeable {
    /** @notice The value of a winning ticket in SOLO. */
    uint256 public faceValue;

    /**
     * @notice The probability of a ticket winning during the start of its lifetime.
     * This is a uint128 value representing the numerator in the probability
     * ratio where 2^128 - 1 is the denominator.
     */
    uint128 public baseLiveWinProb;

    /**
     * @notice The probability of a ticket winning after it has expired.
     * This is a uint128 value representing the numerator in the probability
     * ratio where 2^128 - 1 is the denominator.
     */
    uint128 public expiredWinProb;

    /**
     * @notice The length in blocks before a ticket is considered expired.
     * The default initialization value is 80,000. This equates
     * to roughly two weeks (15s per block).
     */
    uint256 public ticketDuration;

    /**
     * @notice A percentage value representing the proportion of the base win
     * probability that will be decayed once a ticket has expired.
     * Example: 80% decayRate indicates that a ticket will decay down to 20% of its
     * base win probability upon reaching the block before its expiry.
     * The value is expressed as a fraction of 10000.
     */
    uint16 public decayRate;

    function initialize(
        uint256 _faceValue,
        uint128 _baseLiveWinProb,
        uint128 _expiredWinProb,
        uint16 _decayRate,
        uint256 _ticketDuration
    ) public initializer {
        OwnableUpgradeable.__Ownable_init();
        faceValue = _faceValue;
        baseLiveWinProb = _baseLiveWinProb;
        expiredWinProb = _expiredWinProb;
        decayRate = _decayRate;
        setTicketDuration(_ticketDuration);
    }

    /**
     * @notice Set the face value for tickets in SOLO. Only callable by
     * the contract owner.
     * @param _faceValue The face value to set in SOLO.
     */
    function setFaceValue(uint256 _faceValue) public onlyOwner {
        faceValue = _faceValue;
    }

    /**
     * @notice Set the base live win probability of a ticket. Only callable by
     * the contract owner.
     * @param _baseLiveWinProb The probability represented as a value
     * between 0 to 2**128 - 1.
     */
    function setBaseLiveWinProb(uint128 _baseLiveWinProb) public onlyOwner {
        baseLiveWinProb = _baseLiveWinProb;
    }

    /**
     * @notice Set the expired win probability of a ticket. Only callable by
     * the contract owner.
     * @param _expiredWinProb The probability represented as a value
     * between 0 to 2**128 - 1.
     */
    function setExpiredWinProb(uint128 _expiredWinProb) public onlyOwner {
        expiredWinProb = _expiredWinProb;
    }

    /**
     * @notice Set the decay rate of a ticket. Only callable by the
     * the contract owner.
     * @param _decayRate The decay rate as a percentage, where the
     * denominator is 10000.
     */
    function setDecayRate(uint16 _decayRate) public onlyOwner {
        decayRate = _decayRate;
    }

    /**
     * @notice Set the ticket duration of a ticket. Only callable by the
     * contract owner.
     * @param _ticketDuration The duration of a ticket in number of blocks.
     */
    function setTicketDuration(uint256 _ticketDuration) public onlyOwner {
        require(_ticketDuration > 0, "Ticket duration cannot be 0");
        ticketDuration = _ticketDuration;
    }
}