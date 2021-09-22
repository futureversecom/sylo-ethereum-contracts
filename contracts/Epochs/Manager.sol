// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "../Payments/Pricing/Manager.sol";
import "../Payments/Ticketing/Parameters.sol";
import "../Listings.sol";
import "../Staking/Directory.sol";

contract EpochsManager is Initializable, OwnableUpgradeable {

    struct Epoch {
        uint256 iteration;

        // time related variables
        uint256 startBlock; // Block the epoch was initialized
        uint256 duration; // Minimum time epoch will be alive measued in number of blocks
        uint256 endBlock; // Block the epoch ended (and when the next epoch was initialised)
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

    PriceManager _priceManager;

    Directory _directory;

    Listings _listings;

    TicketingParameters _ticketingParameters;

    /* Define all Epoch specific parameters here.
     * When initializing an epoch, these parameters are read,
     * along with parameters from the other contracts to create the
     * new epoch.
     */

    uint256 public epochDuration;

    // Increment this value as each epoch is intialized.
    // The iteration is also used as the epoch's identifier.
    uint256 public currentIteration;

    mapping (uint256 => Epoch) epochs;

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
    }

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

    function getCurrentActiveEpoch() public view returns (Epoch memory epoch) {
        return epochs[currentIteration];
    }

    function getNextEpochId() public view returns (uint256) {
        return currentIteration + 1;
    }

    function getEpoch(uint256 epochId) public view returns (Epoch memory) {
        return epochs[epochId];
    }
}