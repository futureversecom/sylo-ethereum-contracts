// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

library SyloUtils {
    uint256 constant PERCENTEGE_DENOMINATOR = 100;

    /*
     * Multiply a value by a given percentage.
     * Converts the provided uint128 value to uint256 to avoid
     * any reverts on overflow.
     */
    function percOf(uint128 value, uint256 percentage) internal pure returns (uint256) {
        return value * percentage / PERCENTEGE_DENOMINATOR;
    }
}