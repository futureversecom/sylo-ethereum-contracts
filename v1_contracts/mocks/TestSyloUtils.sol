// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "../libraries/SyloUtils.sol";

contract TestSyloUtils {
    function percOf(uint128 value, uint32 percentage) public pure returns (uint256) {
        return SyloUtils.percOf(value, percentage);
    }

    function asPerc(uint128 numerator, uint256 denominator) public pure returns (uint32) {
        return SyloUtils.asPerc(numerator, denominator);
    }

    function validateContractInterface(
        string memory name,
        address target,
        bytes4 interfaceId
    ) public view {
        SyloUtils.validateContractInterface(name, target, interfaceId);
    }
}
