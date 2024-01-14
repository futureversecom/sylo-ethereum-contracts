// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/access/Ownable2StepUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

import "./interfaces/ISeekerPowerOracle.sol";

/**
 * @notice Acts as a source of information for Seeker Powers. Allows setting
 * a Seeker's power level via a restricted owner call. Seeker Power can also
 * be set by any account if the correct Oracle signature proof is provided.
 */
contract SeekerPowerOracle is
  ISeekerPowerOracle,
  Initializable,
  Ownable2StepUpgradeable
{
  /**
   * @notice The oracle account. This contract accepts any attestations of
   * Seeker power that have been signed by this account.
   */
  address public oracle;

  /**
   * @notice Tracks the set Seeker Power levels.
   */
  mapping(uint256 => uint256) public seekerPowers;

  event SeekerPowerUpdated(
    uint256 indexed seekerId,
    uint256 indexed power
  );

  function initialize(
    address _oracle
  ) external initializer {
    Ownable2StepUpgradeable.__Ownable2Step_init();

    oracle = _oracle;
  }

  error UnauthorizedSetSeekerCall();

  function setOracle(address _oracle) external onlyOwner {
    oracle = _oracle;
  }

  function setSeekerPowerRestricted(uint256 seekerId, uint256 power) external {
    if (msg.sender != this.owner() || msg.sender != oracle) {
      revert UnauthorizedSetSeekerCall();
    }

    seekerPowers[seekerId] = power;
    emit SeekerPowerUpdated(seekerId, power);
  }

  function setSeekerPower(uint256 seekerId, uint256 power, bytes calldata proof) external {
    bytes memory proofMessage = getProofMessage((seekerId, power);
    bytes32 ecdsaHash = ECDSA.toEthSignedMessageHash(proofMessage);

    if (ECDSA.recover(ecdsaHash, proof) != oracle) {
      revert UnauthorizedSetSeekerCall();
    }

    seekerPowers[seekerId] = power;
    emit SeekerPowerUpdated(seekerId, power);
  }

  function getSeekerPower(uint256 seekerId) external view returns (uint256) {
    return seekerPowers[seekerId];
  }

  function getProofMessage(uint256 seekerId, uint256 power) external view returns (bytes memory) {
    return abi.encodePacked(
      Strings.toString(seekerId),
      ":",
      Strings.toString(power)
    );
  }
}