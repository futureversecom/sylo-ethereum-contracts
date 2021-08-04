// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "abdk-libraries-solidity/ABDKMathQuad.sol";

library SyloUtils {
    /* 
     * Perform operation `a = x * y / z` that avoids phantom overflow during
     * `x * y` calculation. Will only revert if `a` does not fit into 256-bit.
     */
    function mulDiv(uint x, uint y, uint z) internal pure returns (uint) {
      return
        ABDKMathQuad.toUInt(
            ABDKMathQuad.div(
                ABDKMathQuad.mul(
                    ABDKMathQuad.fromUInt(x),
                    ABDKMathQuad.fromUInt(y)
                ),
                ABDKMathQuad.fromUInt(z)
            )
        );
    }

    /*
     * Convert a fraction to a percentage value
     */
    function toPerc(uint numerator, uint denominator) internal pure returns (uint) {
        return mulDiv(numerator, 100, denominator);
    }

    /*
     * Multiply a value by a given percentage
     */
    function percOf(uint value, uint percentage) internal pure returns (uint) {
        return mulDiv(value, percentage, 100);
    }
}