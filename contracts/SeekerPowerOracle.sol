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
contract SeekerPowerOracle is ISeekerPowerOracle, Initializable, Ownable2StepUpgradeable {
    /**
     * @notice The oracle account. This contract accepts any attestations of
     * Seeker power that have been signed by this account.
     */
    address public oracle;

    /**
     * @notice Tracks nonce used when register the Seeker power to
     * prevent signature re-use.
     */
    mapping(bytes32 => address) private proofNonces;

    /**
     * @notice Tracks the set of Seeker Power levels.
     */
    mapping(uint256 => uint256) public seekerPowers;

    event SeekerPowerUpdated(uint256 indexed seekerId, uint256 indexed power);

    function initialize(address _oracle) external initializer {
        Ownable2StepUpgradeable.__Ownable2Step_init();

        oracle = _oracle;
    }

    error UnauthorizedSetSeekerCall();
    error NonceCannotBeReused();

    /**
     * @notice Sets the oracle account.
     * @param _oracle The oracle account.
     */
    function setOracle(address _oracle) external onlyOwner {
        oracle = _oracle;
    }

    /**
     * @notice Registers a Seeker's power level. Only callable by the
     * owner or the oracle account.
     * @param seekerId The id of the Seeker.
     * @param power The power level of the Seeker.
     */
    function registerSeekerPowerRestricted(uint256 seekerId, uint256 power) external {
        if (msg.sender != this.owner() && msg.sender != oracle) {
            revert UnauthorizedSetSeekerCall();
        }

        seekerPowers[seekerId] = power;
        emit SeekerPowerUpdated(seekerId, power);
    }

    /**
     * @notice Registers a Seeker's power level. Callable by any account
     * but requires a proof signed by the oracle.
     * @param seekerId The id of the Seeker.
     * @param power The power level of the Seeker.
     */
    function registerSeekerPower(
        uint256 seekerId,
        uint256 power,
        bytes32 nonce,
        bytes calldata proof
    ) external {
        if (proofNonces[nonce] != address(0)) {
            revert NonceCannotBeReused();
        }

        bytes memory proofMessage = getProofMessage(seekerId, power, nonce);
        bytes32 ecdsaHash = ECDSA.toEthSignedMessageHash(proofMessage);

        if (ECDSA.recover(ecdsaHash, proof) != oracle) {
            revert UnauthorizedSetSeekerCall();
        }

        seekerPowers[seekerId] = power;
        emit SeekerPowerUpdated(seekerId, power);

        proofNonces[nonce] = oracle;
    }

    /**
     * @notice Retrieves a Seeker's stored power level.
     * @param seekerId The id of the Seeker.
     */
    function getSeekerPower(uint256 seekerId) external view returns (uint256) {
        return seekerPowers[seekerId];
    }

    /**
     * @notice Constructs a proof message for the oracle to sign.
     * @param seekerId The id of the Seeker.
     * @param power The power level of the Seeker.
     */
    function getProofMessage(
        uint256 seekerId,
        uint256 power,
        bytes32 nonce
    ) public pure returns (bytes memory) {
        return
            abi.encodePacked(
                Strings.toString(seekerId),
                ":",
                Strings.toString(power),
                ":",
                Strings.toHexString(uint256(nonce), 32)
            );
    }
}
