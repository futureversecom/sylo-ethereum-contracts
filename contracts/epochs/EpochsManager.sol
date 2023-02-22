// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@openzeppelin/contracts-upgradeable/access/Ownable2StepUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "../Registries.sol";
import "../staking/Directory.sol";
import "../interfaces/epochs/IEpochsManager.sol";
import "../payments/ticketing/TicketingParameters.sol";

contract EpochsManager is IEpochsManager, Initializable, Ownable2StepUpgradeable, ERC165 {
    Directory public _directory;

    Registries public _registries;

    IERC721 public _rootSeekers;

    TicketingParameters public _ticketingParameters;

    /**
     * @notice Track seekers that have joined for a specific epoch.
     */
    mapping(uint256 => mapping(uint256 => address)) public activeSeekers;

    // Define all Epoch specific parameters here.
    // When initializing an epoch, these parameters are read,
    // along with parameters from the other contracts to create the
    // new epoch.

    /**
     * @notice A mapping of all epochs that have been initialized.
     */
    mapping(uint256 => Epoch) public epochs;

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

    event NewEpoch(uint256 indexed epochId);
    event EpochJoined(uint256 indexed epochId, address indexed node, uint256 indexed seekerId);
    event EpochDurationUpdated(uint256 epochDuration);

    error SeekerOwnerMismatch();
    error EpochDurationCannotBeZero();
    error DirectoryCannotBeZeroAddress();
    error RegistriesCannotBeZeroAddress();
    error RootSeekerCannotBeZeroAddress();
    error EpochHasNotEnded(uint256 epochId);
    error SeekerAcountCannotBeZeroAddress();
    error TicketingParametersCannotBeZeroAddress();
    error SeekerAlreadyJoinedEpoch(uint256 epochId, uint256 seekerId);

    function initialize(
        IERC721 rootSeekers,
        Directory directory,
        Registries registries,
        TicketingParameters ticketingParameters,
        uint256 _epochDuration
    ) external initializer {
        if (address(rootSeekers) == address(0)) {
            revert RootSeekerCannotBeZeroAddress();
        }

        SyloUtils.validateContractInterface(
            "Directory",
            address(directory),
            type(IDirectory).interfaceId
        );

        SyloUtils.validateContractInterface(
            "Registries",
            address(registries),
            type(IRegistries).interfaceId
        );

        SyloUtils.validateContractInterface(
            "TicketingParameters",
            address(ticketingParameters),
            type(ITicketingParameters).interfaceId
        );

        Ownable2StepUpgradeable.__Ownable2Step_init();

        _rootSeekers = rootSeekers;
        _directory = directory;
        _registries = registries;
        _ticketingParameters = ticketingParameters;
        epochDuration = _epochDuration;
    }

    /**
     * @notice Returns true if the contract implements the interface defined by
     * `interfaceId` from ERC165.
     */
    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return interfaceId == type(IEpochsManager).interfaceId;
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
        if (end > block.number) {
            revert EpochHasNotEnded(currentIteration);
        }

        (
            uint256 faceValue,
            uint128 baseLiveWinProb,
            uint128 expiredWinProb,
            uint256 ticketDuration,
            uint16 decayRate
        ) = _ticketingParameters.getTicketingParameters();

        uint256 nextEpochId = getNextEpochId();

        epochs[nextEpochId] = Epoch(
            block.number,
            epochDuration,
            0,
            _registries.defaultPayoutPercentage(),
            decayRate,
            faceValue,
            baseLiveWinProb,
            expiredWinProb,
            ticketDuration
        );

        current.endBlock = block.number;

        currentIteration = nextEpochId;

        _directory.setCurrentDirectory(nextEpochId);

        emit NewEpoch(nextEpochId);

        return nextEpochId;
    }

    /**
     * @notice Set the epoch duration. Will take effect in the next epoch. only
     * callable by the owner.
     * @param _epochDuration The epoch duration in number of blocks.
     */
    function setEpochDuration(uint256 _epochDuration) external onlyOwner {
        if (_epochDuration == 0) {
            revert EpochDurationCannotBeZero();
        }
        epochDuration = _epochDuration;
        emit EpochDurationUpdated(epochDuration);
    }

    /**
     * @notice Retrieve the parameters for the current epoch.
     * @return The current Epoch parameters.
     */
    function getCurrentActiveEpoch() external view returns (uint256, Epoch memory) {
        return (currentIteration, epochs[currentIteration]);
    }

    /**
     * @notice Nodes should call this to join the next epoch. It will
     * initialize the next reward pool and set the stake for the next directory.
     * @dev This is a proxy function for `initalizeNextRewardPool` and
     * `joinNextDirectory`.
     */
    function joinNextEpoch() external {
        Registries.Registry memory registry = _registries.getRegistry(msg.sender);

        // validate the node's seeker ownership
        if (registry.seekerAccount == address(0)) {
            revert SeekerAcountCannotBeZeroAddress();
        }

        uint256 seekerId = registry.seekerId;

        address owner = _rootSeekers.ownerOf(seekerId);
        if (registry.seekerAccount != owner) {
            revert SeekerOwnerMismatch();
        }

        uint256 nextEpoch = getNextEpochId();
        if (activeSeekers[nextEpoch][seekerId] != address(0)) {
            revert SeekerAlreadyJoinedEpoch(nextEpoch, seekerId);
        }

        activeSeekers[nextEpoch][seekerId] = msg.sender;

        _directory._rewardsManager().initializeNextRewardPool(msg.sender);
        _directory.joinNextDirectory(msg.sender);

        emit EpochJoined(nextEpoch, msg.sender, seekerId);
    }

    /**
     * @notice Retrieve the epoch parameter for the given id.
     * @param epochId The id of the epoch to retrieve.
     * @return The epoch parameters associated with the id.
     */
    function getEpoch(uint256 epochId) external view returns (Epoch memory) {
        return epochs[epochId];
    }

    /**
     * @notice Retrieve the integer value that will be used for the
     * next epoch id.
     * @return The next epoch id identifier.
     */
    function getNextEpochId() public view returns (uint256) {
        return currentIteration + 1;
    }
}
