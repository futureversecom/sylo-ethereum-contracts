// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

interface ISeekerPowerOracle {
    function setOracle(address oracle) external;

    function registerSeekerPowerRestricted(uint256 seekerId, uint256 power) external;

    function registerSeekerPower(
        uint256 seekerId,
        uint256 power,
        bytes32 nonce,
        bytes calldata proof
    ) external;

    function getSeekerPower(uint256 seekerId) external view returns (uint256);

    function getProofMessage(
        uint256 seekerId,
        uint256 power,
        bytes32 nonce
    ) external pure returns (bytes memory);
}
