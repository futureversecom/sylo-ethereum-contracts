// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.13;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "../Payments/Ticketing/Parameters.sol";
import "../Listings.sol";
import "../Staking/Directory.sol";

contract EpochsManager is Initializable, OwnableUpgradeable {

    /// CENNZnet ethereum state oracle precompile address
    address constant STATE_ORACLE = address(27572);

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
        bool success
    );

    Directory public _directory;

    Listings public _listings;

    TicketingParameters public _ticketingParameters;

    mapping (address => uint256) public activeRequests;
    mapping (uint256 => JoinNextEpochRequest) public joinRequests;

    /**
     * @notice The address of the Seekers NFT contract on ethereum mainnet.
     */
    address public seekers;

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
        address _seekers,
        uint256 _epochDuration
    ) external initializer {
        OwnableUpgradeable.__Ownable_init();
        _directory = directory;
        _listings = listings;
        _ticketingParameters = ticketingParameters;
        seekers = _seekers;
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
     * @notice Set the external Seekers NFT contract address.
     * @param _seekers The address of the Seekers NFT contract.
     */
    function setEpochDuration(address _seekers) external onlyOwner {
        seekers = _seekers;
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

    function requestJoinNextEpoch() external returns (uint256) {
        Listings.Listing memory listing = _listings.getListing(msg.sender);

        require(
            listing.seekerAccount != address(0),
            "Node must have a valid listing to join the next epoch"
        );

        _directory._rewardsManager().validateInitializeNextRewardPool(msg.sender);
        _directory.validateJoinNextDirectory(msg.sender);

        bytes memory balanceOfCall = abi.encodeWithSignature("ownerOf(uint256)", listing.seekerId);
        bytes4 callbackSelector = this.completeJoinNextEpochRequest.selector;
        uint256 callbackGasLimit = 400_000;
        uint256 callbackBounty = 2 ether; // == 2 cpay

        // request a remote eth_call via the state oracle
        bytes memory remoteCallRequest = abi.encodeWithSignature(
            "remoteCall(address,bytes,bytes4,uint256,uint256)",
            this.seekers,
            balanceOfCall,
            callbackSelector,
            callbackGasLimit,
            callbackBounty
        );

        (bool success, bytes memory returnData) = STATE_ORACLE.call(remoteCallRequest);
        require(success);

        uint256 requestId = abi.decode(returnData, (uint256));

        // overwrites existing request
        uint256 currentActive = activeRequests[msg.sender];
        activeRequests[msg.sender] = requestId;
        joinRequests[requestId] = JoinNextEpochRequest(
            msg.sender,
            listing.seekerAccount,
            listing.seekerId
        );
        delete joinRequests[currentActive];

        return requestId;
    }

    function completeJoinNextEpochRequest(uint256 requestId, bytes32 returnData) external {
        require(msg.sender == STATE_ORACLE, "must be state oracle");

        // the account that actually owns the seeker
        address ownerOf = address(bytes20(returnData));

        // find the corresponding join request
        JoinNextEpochRequest memory request = joinRequests[requestId];

        // existing request may have been deleted already
        if (request.node == address(0)) {
            return;
        }

        // before placing the node in the next epoch, validate the node's seeker
        // account owns the specified seeker id
        if (request.seekerAccount != ownerOf) {
            emit JoinNextEpochResult(currentIteration + 1, request.node, false);
        } else {
            _directory._rewardsManager().initializeNextRewardPool(request.node);
            _directory.joinNextDirectory(request.node);
            emit JoinNextEpochResult(currentIteration + 1, request.node, true);
        }
    }

    /**
     * @notice Nodes should call this to join the next epoch. It will
     * initialize the next reward pool and set the stake for the next directory.
     * @dev This is a proxy function for `initalizeNextRewardPool` and
     * `joinNextDirectory`.
     */
    function joinNextEpoch() external {
        _directory._rewardsManager().initializeNextRewardPool(msg.sender);
        _directory.joinNextDirectory(msg.sender);
        emit EpochJoined(currentIteration + 1, msg.sender);
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
