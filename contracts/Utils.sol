// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

library SyloUtils {
    /*
     * Percentages are expressed as a ratio where 10000 is the denominator.
     * A large denominator allows for more precision, e.g representing 12.5%
     * can be done as 1250 / 10000
     */
    uint16 constant public PERCENTAGE_DENOMINATOR = 10000;

    /*
     * Multiply a value by a given percentage.
     * Converts the provided uint128 value to uint256 to avoid
     * any reverts on overflow.
     */
    function percOf(uint128 value, uint16 percentage) internal pure returns (uint256) {
        return uint256(value) * percentage / PERCENTAGE_DENOMINATOR;
    }
}