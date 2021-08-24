// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "../Payments/Pricing/Manager.sol";
import "../Staking/Directory.sol";

contract EpochsManager is Initializable, OwnableUpgradeable {

    struct Epoch {
        // time related variables
        uint256 startBlock;
        uint256 duration;

        // pricing variables
        uint256 currentServicePrice;
        uint256 currentUpperPrice;

        // pointer to directory constructed for this epoch
        bytes32 directoryId;

        // ticketing variables
        // uint256 faceValue;
        // uint128 baseLiveWinProb;
        // uint128 expiredWinProb;
        // uint256 ticketDuration;
        // uint256 unlockDuration;
        // uint16 decayRate;
    }

    PriceManager priceManager;

    Directory directory;

    /* Define all Epoch specific parameters here.
     * When initializing an epoch, these parameters are read,
     * along with parameters from the other contracts to create the
     * new epoch.
     */

    uint256 public epochDuration;

    Epoch public currentActiveEpoch;

    Epoch public previousActiveEpoch;

    mapping (bytes32 => Epoch) epochs;

    function initialize(PriceManager _priceManager, Directory _directory) public initializer {
        OwnableUpgradeable.__Ownable_init();
        priceManager = _priceManager;
        directory = _directory;
    }

    function initializeRound(
        uint256[] memory sortedIndexes
    ) public onlyOwner {
        uint256 end = currentActiveEpoch.startBlock + currentActiveEpoch.duration;
        require(end <= block.number, "Current epoch has not yet ended");

        // Calculate the service prices for this epoch
        (uint256 servicePrice, uint256 upperPrice) = priceManager.calculatePrices(sortedIndexes);

        bytes32 directoryId = directory.constructDirectory2();

        Epoch memory nextEpoch = Epoch(
            block.number, 
            epochDuration, 
            servicePrice, 
            upperPrice,
            directoryId
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
                epoch.duration
            )
        );
    }
}