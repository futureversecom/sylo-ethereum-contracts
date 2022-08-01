// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.13;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "../Payments/Ticketing/Parameters.sol";
import "../Listings.sol";
import "../Staking/Directory.sol";
import "../Seekers.sol";

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

    event EpochJoined (
        uint256 epochId,
        address node,
        uint256 seekerId
    );

    Directory public _directory;

    Listings public _listings;

    Seekers public _seekers;

    TicketingParameters public _ticketingParameters;

    /**
     * @notice Track seekers that have joined for a specific epoch.
     */
    mapping (uint256 => mapping (uint256 => address)) public activeSeekers;

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
        Seekers seekers,
        Directory directory,
        Listings listings,
        TicketingParameters ticketingParameters,
        uint256 _epochDuration
    ) external initializer {
        OwnableUpgradeable.__Ownable_init();
        _seekers = seekers;
        _directory = directory;
        _listings = listings;
        _ticketingParameters = ticketingParameters;
        epochDuration = _epochDuration;
        currentIteration = 0;
    }

    /**
     * @notice Call this to initialize the next epoch. On success, a `NewEpoch` event
     * will be emitted.
     * @dev The function will read the current set of network parameters, and store
     * the parameters in a new Epoch struct. The end block of the current epoch
     * will also be set to a non-zero value.
     */
    function initializeEpoch() external returns (uint256) {
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
     * @notice Set the epoch duration. Will take effect in the next epoch. only
     * callable by the owner.
     * @param _epochDuration The epoch duration in number of blocks.
     */
    function setEpochDuration(uint256 _epochDuration) external onlyOwner {
        epochDuration = _epochDuration;
    }

    /**
     * @notice Retrieve the parameters for the current epoch.
     * @return The current Epoch parameters.
     */
    function getCurrentActiveEpoch() external view returns (Epoch memory) {
        return epochs[currentIteration];
    }

    /**
     * @notice Nodes should call this to join the next epoch. It will
     * initialize the next reward pool and set the stake for the next directory.
     * @dev This is a proxy function for `initalizeNextRewardPool` and
     * `joinNextDirectory`.
     */
    function joinNextEpoch() external {
        Listings.Listing memory listing = _listings.getListing(msg.sender);

        // validate the node's seeker ownership
        require(listing.seekerAccount != address(0), "Node must have a valid seeker account to join an epoch");

        address owner = _seekers.ownerOf(listing.seekerId);

        require(listing.seekerAccount == owner, "Node's seeker account does not match the current seeker owner");

        uint256 nextEpoch = getNextEpochId();

        require(
            activeSeekers[nextEpoch][listing.seekerId] == address(0),
            "Seeker has already joined the next epoch"
        );

        _directory._rewardsManager().initializeNextRewardPool(msg.sender);
        _directory.joinNextDirectory(msg.sender);
        activeSeekers[nextEpoch][listing.seekerId] = msg.sender;
        emit EpochJoined(currentIteration + 1, msg.sender, listing.seekerId);
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
    function getEpoch(uint256 epochId) external view returns (Epoch memory) {
        return epochs[epochId];
    }
}
