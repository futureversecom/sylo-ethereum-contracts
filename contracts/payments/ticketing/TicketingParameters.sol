// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@openzeppelin/contracts-upgradeable/access/Ownable2StepUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "../../libraries/SyloUtils.sol";
import "../../interfaces/payments/ticketing/ITicketingParameters.sol";

/**
 * @dev Persists the parameters for the ticketing mechanism. This contract is
 * read by the EpochManager. Extracting the parameters into another
 * contract is necessary to avoid a cyclic dependency between the ticketing
 * and epoch contracts.
 */
contract TicketingParameters is
    ITicketingParameters,
    Initializable,
    Ownable2StepUpgradeable,
    ERC165
{
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
     * ratio where 2^128 - 1 is the denominator. Note: Redeeming expired
     * tickets is currently not supported.
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
     * The value is expressed as a fraction of 100000.
     */
    uint32 public decayRate;

    /** @notice The value of a winning multi-receiver ticket in SOLO.
     * This value was added from an upgrade, so is not present int the initialize
     * method.
     */
    uint256 public multiReceiverFaceValue;

    event FaceValueUpdated(uint256 faceValue);
    event BaseLiveWinProbUpdated(uint128 baseLiveWinprob);
    event ExpiredWinProbUpdated(uint128 expiredWinProb);
    event TicketDurationUpdated(uint256 ticketDuration);
    event DecayRateUpdated(uint32 decayRate);
    event MultiReceiverFaceValueUpdated(uint256 multiReceiverFaceValue);

    error FaceValueCannotBeZero();
    error TicketDurationCannotBeZero();

    function initialize(
        uint256 _faceValue,
        uint128 _baseLiveWinProb,
        uint128 _expiredWinProb,
        uint32 _decayRate,
        uint256 _ticketDuration
    ) external initializer {
        if (_faceValue == 0) {
            revert FaceValueCannotBeZero();
        }
        if (_ticketDuration == 0) {
            revert TicketDurationCannotBeZero();
        }

        Ownable2StepUpgradeable.__Ownable2Step_init();

        faceValue = _faceValue;
        baseLiveWinProb = _baseLiveWinProb;
        expiredWinProb = _expiredWinProb;
        decayRate = _decayRate;
        ticketDuration = _ticketDuration;
    }

    /**
     * @notice Returns true if the contract implements the interface defined by
     * `interfaceId` from ERC165.
     */
    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return interfaceId == type(ITicketingParameters).interfaceId;
    }

    /**
     * @notice Set the face value for tickets in SOLO. Only callable by
     * the contract owner.
     * @param _faceValue The face value to set in SOLO.
     */
    function setFaceValue(uint256 _faceValue) external onlyOwner {
        if (_faceValue == 0) {
            revert FaceValueCannotBeZero();
        }

        faceValue = _faceValue;
        emit FaceValueUpdated(_faceValue);
    }

    /**
     * @notice Set the base live win probability of a ticket. Only callable by
     * the contract owner.
     * @param _baseLiveWinProb The probability represented as a value
     * between 0 to 2**128 - 1.
     */
    function setBaseLiveWinProb(uint128 _baseLiveWinProb) external onlyOwner {
        baseLiveWinProb = _baseLiveWinProb;
        emit BaseLiveWinProbUpdated(_baseLiveWinProb);
    }

    /**
     * @notice Set the expired win probability of a ticket. Only callable by
     * the contract owner.
     * @param _expiredWinProb The probability represented as a value
     * between 0 to 2**128 - 1.
     */
    function setExpiredWinProb(uint128 _expiredWinProb) external onlyOwner {
        expiredWinProb = _expiredWinProb;
        emit ExpiredWinProbUpdated(_expiredWinProb);
    }

    /**
     * @notice Set the decay rate of a ticket. Only callable by the
     * the contract owner.
     * @param _decayRate The decay rate as a percentage, where the
     * denominator is 10000.
     */
    function setDecayRate(uint32 _decayRate) external onlyOwner {
        decayRate = _decayRate;
        emit DecayRateUpdated(_decayRate);
    }

    /**
     * @notice Set the ticket duration of a ticket. Only callable by the
     * contract owner.
     * @param _ticketDuration The duration of a ticket in number of blocks.
     */
    function setTicketDuration(uint256 _ticketDuration) external onlyOwner {
        if (_ticketDuration == 0) {
            revert TicketDurationCannotBeZero();
        }

        ticketDuration = _ticketDuration;
        emit TicketDurationUpdated(_ticketDuration);
    }

    /**
     * @notice Retrieve the current ticketing parameters.
     * @return faceValue The face value of a ticket in SOLO.
     * @return baseLiveWinProb The base live win probability of a ticket.
     * @return expiredWinProb The expired win probability of a ticket.
     * @return ticketDuration The duration of a ticket in number of blocks.
     * @return decayRate The decay rate of a ticket.
     * @return multiReceiverFaceValue The face value of a multi-receiver ticket in SOLO.
     */
    function getTicketingParameters()
        external
        view
        returns (uint256, uint128, uint128, uint256, uint32, uint256)
    {
        return (faceValue, baseLiveWinProb, expiredWinProb, ticketDuration, decayRate, multiReceiverFaceValue);
    }

    function setMultiReceiverFaceValue(uint256 _multiReceiverFaceValue) external onlyOwner {
        if (_multiReceiverFaceValue == 0) {
            revert FaceValueCannotBeZero();
        }

        multiReceiverFaceValue = _multiReceiverFaceValue;
        emit MultiReceiverFaceValueUpdated(multiReceiverFaceValue);
    }
}
