// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

interface ISeekerPowerOracle {
  function setOracle(address oracle) external;

  function setSeekerPowerRestricted(uint256 seekerId, uint256 power) external;

  function setSeekerPower(uint256 seekerId, uint256 power, bytes calldata proof) external;

  function getProofMessage(uint256 seekerId, uint256 power) external view returns (bytes memory);
}