// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.13;

// This contract is solely used for mocking the oracle state bridge
contract MockOracle {

    // mock nft ownership
    mapping (uint256 => address) owners;

    uint256 public nextRequestId;

    function remoteCall(
        address target,
        bytes memory input,
        bytes4 callbackSignature,
        uint256 callbackGasLimit,
        uint256 bounty
    ) external returns (uint256) {
        nextRequestId++;

        (bool success, bytes memory returnData) = address(this).call(input);
        require(success);

        bytes memory response = abi.encodeWithSelector(callbackSignature, nextRequestId, returnData);
        return nextRequestId;
    }

    function ownerOf(uint256 tokenId) external view returns (address) {
        return owners[tokenId];
    }

    function setOwner(uint256 tokenId, address owner) external {
        owners[tokenId] = owner;
    }
}