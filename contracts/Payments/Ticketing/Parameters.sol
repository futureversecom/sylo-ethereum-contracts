// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

/*
 * Persists the parameters for the ticketing mechanism. This contract is
 * read by the EpochManager. Extracting the parameters into another
 * contract is necessary to avoid a cyclic dependency between the ticketing and epoch contracts.
 */
contract TicketingParameters is Initializable, OwnableUpgradeable {
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

    /**
     * A percentage value representing the proportion of the base win probability
     * that will be decayed once a ticket has expired.
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

    function setFaceValue(uint256 _faceValue) public onlyOwner {
        faceValue = _faceValue;
    }

    function setBaseLiveWinProb(uint128 _baseLiveWinProb) public onlyOwner {
        baseLiveWinProb = _baseLiveWinProb;
    }

    function setExpiredWinProb(uint128 _expiredWinProb) public onlyOwner {
        expiredWinProb = _expiredWinProb;
    }

    function setDecayRate(uint16 _decayRate) public onlyOwner {
        decayRate = _decayRate;
    }

    function setTicketDuration(uint256 _ticketDuration) public onlyOwner {
        require(_ticketDuration > 0, "Ticket duration cannot be 0");
        ticketDuration = _ticketDuration;
    }
}