// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/Ownable2StepUpgradeable.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";

import "./IProtocolTimeManager.sol";

import "hardhat/console.sol";

contract ProtocolTimeManager is
    IProtocolTimeManager,
    Initializable,
    Ownable2StepUpgradeable,
    ERC165
{
    struct CycleDurationUpdate {
        uint256 duration;
        // The unix timestamp this update becomes effective (takes effect at
        // the start of the next cycle
        uint256 updatesAt;
    }

    struct PeriodDurationUpdate {
        uint256 duration;
        // The cycle this duration becomes effective
        uint256 updatesAt;
    }

    /**
     * @notice Holds the unix timestamp for when the protocol starts
     */
    uint256 start;

    CycleDurationUpdate[] cycleDurationUpdates;
    PeriodDurationUpdate[] periodDurationUpdates;

    error CannotInitializeWithZeroCycleDuration();
    error CannotInitializeWithZeroPeriodDuration();
    error CannotSetStartInThePast();
    error CannotSetStartAfterProtocolHasStarted();
    error CannotSetProtocolStartToZero();
    error CannotSetZeroCycleDuration();
    error CannotSetZeroPeriodDuration();
    error ProtocolHasNotBegun();

    function initialize(uint256 _cycleDuration, uint256 _periodDuration) external initializer {
        if (_cycleDuration == 0) {
            revert CannotInitializeWithZeroCycleDuration();
        }
        if (_periodDuration == 0) {
            revert CannotInitializeWithZeroPeriodDuration();
        }

        Ownable2StepUpgradeable.__Ownable2Step_init();

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
     * @notice Sets the start time for the protocol
     * @param _start A unix timestamp representing the start time
     */
    function setProtocolStart(uint256 _start) external onlyOwner {
        if (_start == 0) {
            revert CannotSetProtocolStartToZero();
        }

        if (_start <= block.timestamp) {
            revert CannotSetStartInThePast();
        }

        if (start >= block.timestamp) {
            revert CannotSetStartAfterProtocolHasStarted();
        }

        start = _start;
        // update when the next cycle starts
        cycleDurationUpdates[0].updatesAt = _start;
    }

    function hasProtocolStarted() internal view returns (bool) {
        return block.timestamp >= start && start != 0;
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

    function _setCycleDuration(uint256 _cycleDuration) internal {
        if (_cycleDuration == 0) {
            revert CannotSetZeroCycleDuration();
        }

        // check if this is the first instance we are setting the cycle's duration
        if (cycleDurationUpdates.length == 0) {
            cycleDurationUpdates.push(CycleDurationUpdate(
                _cycleDuration,
                0 // start value will be updated once `start()` is called
            ));
            return;
        }

        CycleDurationUpdate storage lastUpdate = cycleDurationUpdates[cycleDurationUpdates.length - 1];

        // check if we are updating the duration again before the protocol has
        // started
        if (start == 0 || start > block.timestamp) {
            lastUpdate.duration = _cycleDuration;
            return;
        }

        // Check if the next cycle's duration has already been updated. In this
        // case we overwrite the existing update
        if (lastUpdate.updatesAt > block.timestamp) {
            lastUpdate.duration = _cycleDuration;
            return;
        }

        Cycle memory cycle = _getCurrentCycle();

        uint256 nextCycleStart = cycle.start + cycle.duration;

        cycleDurationUpdates.push(CycleDurationUpdate(
            _cycleDuration,
            nextCycleStart
        ));
    }

    /**
     * @notice Sets the period duration
     * @param _periodDuration The duration for which periods should last in unix
     */
    function _setPeriodDuration(uint256 _periodDuration) internal {
        if (_periodDuration == 0) {
            revert CannotSetZeroPeriodDuration();
        }

        // check if this is the first instance of setting the period duration
        if (periodDurationUpdates.length == 0) {
            periodDurationUpdates.push(PeriodDurationUpdate(
                _periodDuration,
                1
            ));
            return;
        }

        PeriodDurationUpdate storage lastUpdate = periodDurationUpdates[periodDurationUpdates.length - 1];

        // check if we are updating the duration again before the protocol has
        // started
        if (start == 0 || start > block.timestamp) {
            lastUpdate.duration = _periodDuration;
            return;
        }

        Cycle memory cycle = _getCurrentCycle();
        uint256 nextCycle = cycle.iteration + 1;

        // Check if the next cycle's period duration has already been updated. In this
        // case we overwrite the existing update
        if (lastUpdate.updatesAt == nextCycle) {
            lastUpdate.duration = _periodDuration;
            return;
        }

        periodDurationUpdates.push(PeriodDurationUpdate(
            _periodDuration,
            nextCycle
        ));
    }

    /**
     * @notice Get the current cycle
     */
    function getCurrentCycle() external view returns (Cycle memory) {
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
    function _getCurrentCycle() internal view returns (Cycle memory) {
        if (!hasProtocolStarted()) {
            revert ProtocolHasNotBegun();
        }

        uint256 totalTimeElapsed = block.timestamp - start;
        uint256 cycles = 1;

        uint256 cursor = 0;
        uint256 currentDuration = cycleDurationUpdates[cursor].duration;
        uint256 currentDurationStart = cycleDurationUpdates[cursor].updatesAt;
        uint256 currentCycleStart = 0;

        uint256 lastUpdateToProcess = 0;

        // We need to process up to the most previous cycle duration update.
        // However, if the next cycle's duration has been updated, then we
        // should only process up to the second to last update.
        if (cycleDurationUpdates[cycleDurationUpdates.length - 1].updatesAt > block.timestamp) {
            lastUpdateToProcess = cycleDurationUpdates.length - 1;
        } else {
            lastUpdateToProcess = cycleDurationUpdates.length;
        }

        while (true) {
            // we have reached the end of the cycle updates
            if (cursor + 1 == lastUpdateToProcess) {
                if (currentDurationStart > block.timestamp) {
                    break;
                }

                // add the cycles that occurred with the current duration
                uint256 cyclesAtCurrentDuration = totalTimeElapsed / currentDuration;
                cycles += cyclesAtCurrentDuration;

                currentCycleStart = currentDurationStart + cyclesAtCurrentDuration * currentDuration;
                break;
            } else {
                uint256 nextDurationUpdate = cycleDurationUpdates[cursor + 1].updatesAt;

                uint256 timeElapsedAtCurrentDuration = nextDurationUpdate - currentDurationStart;

                // add the cycles that occurred with the current duration
                uint256 cyclesAtCurrentDuration = timeElapsedAtCurrentDuration / currentDuration;
                cycles += cyclesAtCurrentDuration;

                // as we iterate through the updates, we track how much
                // of the total time remaining to account for before
                // processing the next cycle duration update
                totalTimeElapsed -= timeElapsedAtCurrentDuration;

                // update cursors
                cursor++;
                currentDuration = cycleDurationUpdates[cursor].duration;
                currentDurationStart = cycleDurationUpdates[cursor].updatesAt;
            }
        }

        return Cycle(cycles, currentCycleStart, currentDuration);
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
        if (!hasProtocolStarted()) {
            revert ProtocolHasNotBegun();
        }

        Cycle memory cycle = _getCurrentCycle();
        uint256 totalTimeElapsedWithinPeriod = block.timestamp - cycle.start;

        uint256 periodDuration = _getPeriodDuration();

        return totalTimeElapsedWithinPeriod / periodDuration;
    }

    /**
     * @notice Get the cycle duration
     */
    function getCycleDuration() external view returns (uint256) {
        return _getCycleDuration();
    }

    function _getCycleDuration() internal view returns (uint256) {
        // if the protocol has not started, then the current duration is the
        // first update
        if (!hasProtocolStarted()) {
            return cycleDurationUpdates[0].duration;
        }

        CycleDurationUpdate storage lastUpdate = cycleDurationUpdates[cycleDurationUpdates.length - 1];

        // if the last update occurred before the current timestamp, then the
        // last update holds the current duration
        if (lastUpdate.updatesAt <= block.timestamp) {
            return lastUpdate.duration;
        // else the duration has been updated for the next cycle, so the current
        // cycle duration is defined in the previous update
        } else {
            return cycleDurationUpdates[cycleDurationUpdates.length - 2].duration;
        }
    }

    /**
     * @notice Get the period duration
     */
    function getPeriodDuration() external view returns (uint256) {
        return _getPeriodDuration();
    }

    function _getPeriodDuration() internal view returns (uint256) {
        // if the protocol has not started, then the current duration is the
        // first update
        if (!hasProtocolStarted()) {
            return periodDurationUpdates[0].duration;
        }

        Cycle memory cycle = _getCurrentCycle();

        PeriodDurationUpdate storage lastUpdate = periodDurationUpdates[periodDurationUpdates.length - 1];

        // if the last update occurred before the current cycle, then the
        // last update holds the current duration
        if (lastUpdate.updatesAt <= cycle.iteration) {
            return lastUpdate.duration;
        // else the duration has been updated for the next cycle, so the current
        // period duration is defined in the previous update
        } else {
            return periodDurationUpdates[periodDurationUpdates.length - 2].duration;
        }
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
    function getTime() external view returns (uint256, uint256, Cycle memory) {
        Cycle memory cycle = _getCurrentCycle();
        uint256 period = _getCurrentPeriod();
        return (cycle.iteration, period, cycle);
    }
}
