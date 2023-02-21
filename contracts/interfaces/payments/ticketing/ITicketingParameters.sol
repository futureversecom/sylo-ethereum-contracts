// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

interface ITicketingParameters {
    function setFaceValue(uint256 _faceValue) external;

    function setBaseLiveWinProb(uint128 _baseLiveWinProb) external;

    function setExpiredWinProb(uint128 _expiredWinProb) external;

    function setDecayRate(uint16 _decayRate) external;

    function setTicketDuration(uint256 _ticketDuration) external;

    function getTicketingParameters()
        external
        view
        returns (uint256, uint128, uint128, uint256, uint16);
}
