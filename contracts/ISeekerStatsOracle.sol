// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

interface ISeekerStatsOracle {
    struct Seeker {
        uint256 seekerId;
        uint256 rank;
        uint256 attr_chip;
        uint256 attr_durability;
        uint256 attr_sensors;
        uint256 attr_cores;
        uint256 attr_storage;
        uint256 attr_reactor;
    }

    function createStatsMessage(Seeker calldata seeker) external pure returns (bytes memory);

    function registerSeekerRestricted(Seeker calldata seeker) external;

    function registerSeeker(Seeker calldata seeker, bytes calldata signature) external;

    function calculateAttributeCoverage(Seeker[] calldata seekers) external view returns (int256);
}
