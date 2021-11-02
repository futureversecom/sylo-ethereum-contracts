// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "../Payments/Ticketing/Parameters.sol";
import "../Listings.sol";
import "../Staking/Directory.sol";

contract EpochsManager is Initializable, OwnableUpgradeable {

    /**
     * @dev This struct will hold all network parameters that will be static
     * for the entire epoch. This value will be stored in a mapping, where the
     * key is also the epoch's iteration value.
     */
    struct Epoch {
        uint256 iteration;

        // time related variables
        uint256 startBlock; // Block the epoch was initialized
        uint256 duration; // Minimum time epoch will be alive measured in number of blocks
        uint256 endBlock; // Block the epoch ended (and when the next epoch was initialized)
                          // Zero here represents the epoch has not yet ended.

        // listing variables
        uint16 defaultPayoutPercentage;

        // ticketing variables
        uint256 faceValue;
        uint128 baseLiveWinProb;
        uint128 expiredWinProb;
        uint256 ticketDuration;
        uint16 decayRate;
    }

    Directory public _directory;

    Listings public _listings;

    TicketingParameters public _ticketingParameters;

    // Define all Epoch specific parameters here.
    // When initializing an epoch, these parameters are read,
    // along with parameters from the other contracts to create the
    // new epoch.

    /**
     * @notice The duration in blocks an epoch will last for.
     */
    uint256 public epochDuration;

    /**
     * @notice The value of the integer used as the current
     * epoch's identifier. This value is incremented as each epoch
     * is initialized.
     */
    uint256 public currentIteration;

    /**
     * @notice A mapping of all epochs that have been initialized.
     */
    mapping (uint256 => Epoch) public epochs;

    event NewEpoch(uint256 epochId);

    function initialize(
        Directory directory,
        Listings listings,
        TicketingParameters ticketingParameters,
        uint256 _epochDuration
    ) public initializer {
        OwnableUpgradeable.__Ownable_init();
        _directory = directory;
        _listings = listings;
        _ticketingParameters = ticketingParameters;
        epochDuration = _epochDuration;
        currentIteration = 0;
    }

    /**
     * @notice Call this to initialize the next epoch. This is only callable
     * by the owner of the Sylo contracts. On success, a `NewEpoch` event
     * will be emitted.
     * @dev The function will read the current set of network parameters, and store
     * the parameters in a new Epoch struct. The end block of the current epoch
     * will also be set to a non-zero value.
     */
    function initializeEpoch() public returns (uint256) {
        Epoch storage current = epochs[currentIteration];

        uint256 end = current.startBlock + current.duration;
        require(end <= block.number, "Current epoch has not yet ended");

        uint256 nextIteration = currentIteration + 1;

        Epoch memory nextEpoch = Epoch(
            nextIteration,
            block.number,
            epochDuration,
            0,
            _listings.defaultPayoutPercentage(),
            _ticketingParameters.faceValue(),
            _ticketingParameters.baseLiveWinProb(),
            _ticketingParameters.expiredWinProb(),
            _ticketingParameters.ticketDuration(),
            _ticketingParameters.decayRate()
        );

        uint256 epochId = getNextEpochId();

        _directory.setCurrentDirectory(epochId);

        epochs[epochId] = nextEpoch;
        current.endBlock = block.number;

        currentIteration = nextIteration;

        emit NewEpoch(epochId);

        return epochId;
    }

    /**
     * @notice Retrieve the parameters for the current epoch.
     * @return The current Epoch parameters.
     */
    function getCurrentActiveEpoch() public view returns (Epoch memory) {
        return epochs[currentIteration];
    }

    /**
     * @notice Retrieve the integer value that will be used for the
     * next epoch id.
     * @return The next epoch id identifier.
     */
    function getNextEpochId() public view returns (uint256) {
        return currentIteration + 1;
    }

    /**
     * @notice Retrieve the epoch parameter for the given id.
     * @param epochId The id of the epoch to retrieve.
     * @return The epoch parameters associated with the id.
     */
    function getEpoch(uint256 epochId) public view returns (Epoch memory) {
        return epochs[epochId];
    }
}
