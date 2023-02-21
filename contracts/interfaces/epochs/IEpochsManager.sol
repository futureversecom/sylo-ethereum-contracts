// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

interface IEpochsManager {
    /**
     * @dev This struct will hold all network parameters that will be static
     * for the entire epoch. This value will be stored in a mapping, where the
     * key is the current epoch id.
     */
    struct Epoch {
        // time related variables
        uint256 startBlock; // Block the epoch was initialized
        uint256 duration; // Minimum time epoch will be alive measured in number of blocks
        uint256 endBlock; // Block the epoch ended (and when the next epoch was initialized)
        // Zero here represents the epoch has not yet ended.

        // registry variables
        uint16 defaultPayoutPercentage;
        // ticketing variables
        uint16 decayRate;
        uint256 faceValue;
        uint128 baseLiveWinProb;
        uint128 expiredWinProb;
        uint256 ticketDuration;
    }

    function initializeEpoch() external returns (uint256);

    function setEpochDuration(uint256 _epochDuration) external;

    function getCurrentActiveEpoch() external view returns (uint256, Epoch memory);

    function joinNextEpoch() external;

    function getEpoch(uint256 epochId) external view returns (Epoch memory);

    function getNextEpochId() external view returns (uint256);
}
