// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/Ownable2StepUpgradeable.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";

import "./IProtocolTimeManager.sol";

contract ProtocolTimeManager is
    IProtocolTimeManager,
    Initializable,
    Ownable2StepUpgradeable,
    ERC165
{
    using EnumerableMap for EnumerableMap.UintToUintMap;

    /**
     * @notice Holds the start time for the cycle/period intervals in unix
     */
    uint256 start;

    /**
     * @notice Holds the cycle duration in unix
     */
    uint256 cycleDuration;

    /**
     * @notice Holds the period duration in unix
     */
    uint256 periodDuration;

    /**
     * @notice A iterable map used to track cycle duration updates.
     * Indexed by the timestamp of the block that made the duration update
     */
    EnumerableMap.UintToUintMap cycleDurationUpdates;

    /**
     * @notice A iterable map used to track period duration updates.
     * Indexed by the timestamp of the block that made the duration update
     */
    EnumerableMap.UintToUintMap periodDurationUpdates;

    error CannotInitializeWithZeroStart();
    error CannotInitializeWithZeroCycleDuration();
    error CannotInitializeWithZeroPeriodDuration();
    error CannotSetProtocolStartWithZeroStart();
    error CannotSetCycleDurationWithZeroDuration();
    error CannotSetStartInFuture();
    error CannotSetZeroPeriodDuration();
    error CannotSetZeroCycleDuration();

    function initialize(uint256 _cycleDuration, uint256 _periodDuration) external initializer {
        if (_cycleDuration == 0) {
            revert CannotInitializeWithZeroCycleDuration();
        }
        if (_periodDuration == 0) {
            revert CannotInitializeWithZeroPeriodDuration();
        }

        Ownable2StepUpgradeable.__Ownable2Step_init();

        start = block.timestamp;
        _setCycleDuration(_cycleDuration);
        _setPeriodDuration(_periodDuration);
    }

    /**
     * @notice Returns true if the contract implements the interface defined by
     * `interfaceId` from ERC165.
     */
    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return interfaceId == type(IProtocolTimeManager).interfaceId;
    }

    /**
     * @notice Sets the start time for the cycle/period intervals
     * @param _start The start time for the cycle/period intervals in unix
     */
    function setProtocolStart(uint256 _start) external onlyOwner {
        if (_start == 0) {
            revert CannotSetProtocolStartWithZeroStart();
        }
        if (_start > block.timestamp) {
            revert CannotSetStartInFuture();
        }
        start = _start;
    }

    /**
     * @notice Sets the cycle duration
     * @param _cycleDuration The duration for which cycles should last in unix
     */
    function setCycleDuration(uint256 _cycleDuration) external onlyOwner {
        _setCycleDuration(_cycleDuration);
    }

    /**
     * @notice Sets the period duration
     * @param _periodDuration The duration for which periods should last in unix
     */
    function setPeriodDuration(uint256 _periodDuration) external onlyOwner {
        _setPeriodDuration(_periodDuration);
    }

    /**
     * @notice Sets the cycle duration
     * @param _cycleDuration The duration for which cycles should last in unix
     */
    function _setCycleDuration(uint256 _cycleDuration) internal {
        if (_cycleDuration == 0) {
            revert CannotSetZeroCycleDuration();
        }
        cycleDuration = _cycleDuration;
        uint256 timeSinceStart = block.timestamp - start;
        cycleDurationUpdates.set(timeSinceStart, cycleDuration);
    }

    /**
     * @notice Sets the period duration
     * @param _periodDuration The duration for which periods should last in unix
     */
    function _setPeriodDuration(uint256 _periodDuration) internal {
        if (_periodDuration == 0) {
            revert CannotSetZeroPeriodDuration();
        }
        periodDuration = _periodDuration;
        uint256 timeSinceStart = block.timestamp - start;
        periodDurationUpdates.set(timeSinceStart, periodDuration);
    }

    /**
     * @notice Get the current cycle
     */
    function getCurrentCycle() external view returns (uint256) {
        return _getCurrentCycle();
    }

    /**
     * @notice Get the current period
     */
    function getCurrentPeriod() external view returns (uint256) {
        return _getCurrentPeriod();
    }

    /**
     * @notice  Calculates the current cycle number based on the elapsed time
     * since the contract's start, taking into account any updates to the cycle duration.
     * This function works by iterating over the cycle durations tracked by the
     * cycleDurationUpdates map and calculating the interval to which this duration
     * applies. By finding the interval to which the duration applies the amount
     * of cycles for that interval can be calculated by dividing the interval time
     * by the duration of the cycles.
     *
     * for exmaple
     * [<------------ 200 ---------->] where each '|' represents a duration update
     * [------|----------|--------|--] and each number a cycle/period interval (excluding 200)
     * [<-60->|<-40->|<--80-->|<-20->]
     * 0      60    100      180     -
     *
     * Duration (0 -> 60):    20    cycles = (60 - 0)    / 20 = 3
     * Duration (60 -> 100):  10    cycles = (100 - 60)  / 10 = 4
     * Duration (100 -> 180): 40    cycles = (180 - 100) / 40 = 2
     *
     * Duration (180 -> -):    5    totalCycles += (200 - 180) / 5  = 13 (where 5 is the current cycle durartion)
     *                                                                   (where 200 is the totalTimeElapsed since start)
     */
    function _getCurrentCycle() internal view returns (uint256) {
        uint256 totalTimeElapsed = block.timestamp - start;
        uint256 processedCycleIntervals = 0;
        uint256 cycles = 1;

        for (uint256 i = 0; (i + 1) < cycleDurationUpdates.length(); ++i) {
            (uint256 timestamp, uint256 intervalDuration) = cycleDurationUpdates.at(i);
            (uint256 nextTimestamp, ) = cycleDurationUpdates.at(i + 1);
            uint256 cycleInterval = nextTimestamp - timestamp;

            processedCycleIntervals += cycleInterval;

            cycles += cycleInterval / intervalDuration;
        }

        uint256 remaingCycleInterval = ((totalTimeElapsed - processedCycleIntervals) /
            cycleDuration);
        uint256 currentCycle = cycles + remaingCycleInterval;
        return currentCycle;
    }

    /**
     * @notice Calculates the current period number based on the elapsed time
     * since the contract's start, taking into account any updates to the period duration.
     * This function works by iterating over the period durations tracked by the
     * periodDurationUpdates map and calculating the interval to which this duration
     * applies. By finding the interval to which the duration applies the amount
     * of periods for that interval can be calculated by dividing the interval time
     * by the duration of the periods.
     * refer to explaination above for funtionality
     */
    function _getCurrentPeriod() internal view returns (uint256) {
        uint256 totalTimeElapsed = block.timestamp - start;
        uint256 processedPeriodIntervals = 0;
        uint256 cycles = 1;

        for (uint256 i = 0; (i + 1) < periodDurationUpdates.length(); ++i) {
            (uint256 timestamp, uint256 intervalDuration) = periodDurationUpdates.at(i);
            (uint256 nextTimestamp, ) = periodDurationUpdates.at(i + 1);
            uint256 periodInterval = nextTimestamp - timestamp;

            processedPeriodIntervals += periodInterval;

            cycles += periodInterval / intervalDuration;
        }

        uint256 remaingPeriodInterval = ((totalTimeElapsed - processedPeriodIntervals) /
            periodDuration);
        uint256 currentPeriod = cycles + remaingPeriodInterval;
        return currentPeriod;
    }

    /**
     * @notice Get the cycle duration
     */
    function getCycleDuration() external view returns (uint256) {
        return cycleDuration;
    }

    /**
     * @notice Get the period duration
     */
    function getPeriodDuration() external view returns (uint256) {
        return periodDuration;
    }

    /**
     * @notice Get the start time of the cycle/period intervals
     */
    function getStart() external view returns (uint256) {
        return start;
    }

    /**
     * @notice Get the cycle and period durations
     */
    function timeNow() external view returns (uint256, uint256) {
        uint256 cycle = _getCurrentCycle();
        uint256 period = _getCurrentPeriod();
        return (cycle, period);
    }
}
