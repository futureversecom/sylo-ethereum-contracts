// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "../Payments/Pricing/Manager.sol";
import "../Payments/Ticketing/Parameters.sol";
import "../Staking/Directory.sol";

contract EpochsManager is Initializable, OwnableUpgradeable {

    struct Epoch {
        // time related variables
        uint256 startBlock;
        uint256 duration;

        // pointer to directory constructed for this epoch
        bytes32 directoryId;

        // ticketing variables
        uint256 faceValue;
        uint128 baseLiveWinProb;
        uint128 expiredWinProb;
        uint256 ticketDuration;
        uint16 decayRate;
    }

    PriceManager _priceManager;

    Directory _directory;

    TicketingParameters _ticketingParameters;

    /* Define all Epoch specific parameters here.
     * When initializing an epoch, these parameters are read,
     * along with parameters from the other contracts to create the
     * new epoch.
     */

    uint256 public epochDuration;

    Epoch public currentActiveEpoch;

    Epoch public previousActiveEpoch;

    mapping (bytes32 => Epoch) epochs;

    function initialize(
        Directory directory, 
        TicketingParameters ticketingParameters
    ) public initializer {
        OwnableUpgradeable.__Ownable_init();
        _directory = directory;
        _ticketingParameters = ticketingParameters;
    }

    function initializeRound() public onlyOwner {
        uint256 end = currentActiveEpoch.startBlock + currentActiveEpoch.duration;
        require(end <= block.number, "Current epoch has not yet ended");

        bytes32 directoryId = _directory.constructDirectory();

        Epoch memory nextEpoch = Epoch(
            block.number, 
            epochDuration, 
            directoryId,
            _ticketingParameters.faceValue(),
            _ticketingParameters.baseLiveWinProb(),
            _ticketingParameters.expiredWinProb(),
            _ticketingParameters.ticketDuration(),
            _ticketingParameters.decayRate()
        );
        
        bytes32 id = getEpochId(nextEpoch);

        epochs[id] = nextEpoch;
        previousActiveEpoch = currentActiveEpoch;
        currentActiveEpoch = nextEpoch;
    }

    function getEpochId(Epoch memory epoch) public pure returns (bytes32) {
        return keccak256(
            abi.encodePacked(
                epoch.startBlock,
                epoch.duration,
                epoch.faceValue,
                epoch.baseLiveWinProb,
                epoch.expiredWinProb,
                epoch.ticketDuration,
                epoch.decayRate
            )
        );
    }

    function getEpoch(bytes32 epochId) public view returns (Epoch memory) {
        return epochs[epochId];
    }
}