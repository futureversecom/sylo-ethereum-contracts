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
    // This value is also used to deterministically determine the
    // next epoch identifier.
    uint256 currentIteration;

    bytes32 public currentActiveEpoch;

    mapping (bytes32 => Epoch) epochs;

    event NewEpoch(bytes32 epochId);

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

    function initializeEpoch() public returns (bytes32) {
        Epoch storage current = epochs[currentActiveEpoch];

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

        bytes32 epochId = getNextEpochId();

        epochs[epochId] = nextEpoch;
        current.endBlock = block.number;

        currentIteration = nextIteration;
        currentActiveEpoch = epochId;

        emit NewEpoch(epochId);

        return epochId;
    }

    function getCurrentActiveEpoch() public view returns (Epoch memory epoch) {
        return epochs[currentActiveEpoch];
    }

    function getEpochId(uint256 iteration) public pure returns (bytes32) {
        return keccak256(
            abi.encodePacked(iteration)
        );
    }

    function getNextEpochId() public view returns (bytes32) {
        return getEpochId(currentIteration + 1);
    }

    function getEpoch(bytes32 epochId) public view returns (Epoch memory) {
        return epochs[epochId];
    }
}