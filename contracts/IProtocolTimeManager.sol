// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

interface IProtocolTimeManager {
    function setProtocolStart(uint256 _start) external;

    function setCycleDuration(uint256 duration) external;

    function setPeriodDuration(uint256 duration) external;

    function getCycleDuration() external returns (uint256);

    function getPeriodDuration() external returns (uint256);

    function timeNow() external returns (uint256, uint256);

    function getCurrentCycle() external returns (uint256);

    function getCurrentPeriod() external returns (uint256);

    function getStart() external view returns (uint256);
}
