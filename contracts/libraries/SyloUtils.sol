// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/utils/math/SafeCast.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";

error ContractNameCannotBeEmpty();
error InterfaceIdCannotBeZeroBytes();
error TargetContractCannotBeZeroAddress(string name);
error TargetNotSupportInterface(string name, bytes4 interfaceId);

library SyloUtils {
    /**
     * @dev The maximum possible SYLO that exists in the network.
     */
    uint256 public constant MAX_SYLO = 10_000_000_000 ether;

    /**
     * @dev Percentages are expressed as a ratio where 100000 is the denominator.
     * A large denominator allows for more precision, e.g representing 12.5%
     * can be done as 12500 / 100000
     */
    uint32 public constant PERCENTAGE_DENOMINATOR = 100000;

    /**
     * @dev Multiply a value by a given percentage. Converts the provided
     * uint128 value to uint256 to avoid any reverts on overflow.
     * @param value The value to multiply.
     * @param percentage The percentage, as a ratio of 100000.
     */
    function percOf(uint128 value, uint32 percentage) internal pure returns (uint256) {
        return (uint256(value) * percentage) / PERCENTAGE_DENOMINATOR;
    }

    /**
     * @dev Return a fraction as a percentage.
     * @param numerator The numerator limited to a uint128 value to prevent
     * phantom overflow.
     * @param denominator The denominator.
     * @return The percentage, as a ratio of 100000.
     */
    function asPerc(uint128 numerator, uint256 denominator) internal pure returns (uint32) {
        return SafeCast.toUint32((uint256(numerator) * PERCENTAGE_DENOMINATOR) / denominator);
    }

    /**
     * @dev Validate that a contract implements a given interface.
     * @param name The name of the contract, used in error messages.
     * @param target The address of the contract.
     * @param interfaceId The interface ID to check.
     */
    function validateContractInterface(
        string memory name,
        address target,
        bytes4 interfaceId
    ) internal view {
        if (bytes(name).length == 0) {
            revert ContractNameCannotBeEmpty();
        }
        if (target == address(0)) {
            revert TargetContractCannotBeZeroAddress(name);
        }
        if (interfaceId == bytes4(0)) {
            revert InterfaceIdCannotBeZeroBytes();
        }
        if (!ERC165(target).supportsInterface(interfaceId)) {
            revert TargetNotSupportInterface(name, interfaceId);
        }
    }
}