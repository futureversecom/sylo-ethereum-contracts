// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.13;

library SyloUtils {
    /**
     * @dev Percentages are expressed as a ratio where 10000 is the denominator.
     * A large denominator allows for more precision, e.g representing 12.5%
     * can be done as 1250 / 10000
     */
    uint16 constant public PERCENTAGE_DENOMINATOR = 10000;

    /**
     * @dev Multiply a value by a given percentage. Converts the provided
     * uint128 value to uint256 to avoid any reverts on overflow.
     * @param value The value to multiply.
     * @param percentage The percentage, as a ratio of 10000.
     */
    function percOf(uint128 value, uint16 percentage) internal pure returns (uint256) {
        return uint256(value) * percentage / PERCENTAGE_DENOMINATOR;
    }

    /**
     * @dev Return a fraction as a percentage.
     * @param numerator The numerator limited to a uint128 value to prevent
     * phantom overflow.
     * @param denominator The denominator.
     * @return The percentage, as a ratio of 10000.
     */
    function asPerc(uint128 numerator, uint256 denominator) internal pure returns(uint16) {
        return uint16(uint256(numerator) * PERCENTAGE_DENOMINATOR / denominator);
    }
}