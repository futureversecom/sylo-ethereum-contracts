// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.13;

// This contract is solely used for mocking the oracle state bridge
contract MockOracle {
    // mock nft ownership
    mapping(uint256 => address) public owners;

    struct Callback {
        address sender;
        bytes callback;
    }

    mapping(uint256 => Callback) private callbacks;

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

        bytes32 encoded;
        assembly {
            encoded := mload(add(returnData, 32))
        }

        bytes memory callback = abi.encodeWithSelector(
            callbackSignature,
            nextRequestId,
            block.timestamp,
            encoded
        );
        callbacks[nextRequestId] = Callback(msg.sender, callback);

        return nextRequestId;
    }

    function remoteCallWithFeeSwap(
        address target,
        bytes memory input,
        bytes4 callbackSignature,
        uint256 callbackGasLimit,
        uint256 bounty,
        address token,
        uint256 maxFee
    ) external returns (uint256) {
        nextRequestId++;

        (bool success, bytes memory returnData) = address(this).call(input);
        require(success);

        bytes32 encoded;
        assembly {
            encoded := mload(add(returnData, 32))
        }

        bytes memory callback = abi.encodeWithSelector(
            callbackSignature,
            nextRequestId,
            block.timestamp,
            encoded
        );
        callbacks[nextRequestId] = Callback(msg.sender, callback);

        return nextRequestId;
    }

    function invokeCallback() external {
        Callback storage callback = callbacks[nextRequestId];
        (bool callbackSuccess, ) = callback.sender.call(callback.callback);
        require(callbackSuccess, "Callback to sender failed");
    }

    function ownerOf(uint256 tokenId) external view returns (address) {
        return owners[tokenId];
    }

    function setOwner(uint256 tokenId, address owner) external {
        owners[tokenId] = owner;
    }
}

contract MockOracleFail {
    function remoteCall(
        address target,
        bytes memory input,
        bytes4 callbackSignature,
        uint256 callbackGasLimit,
        uint256 bounty
    ) public pure returns (uint256) {
        require(false, "MockOracleFail always fails");
    }

    function remoteCallWithFeeSwap(
        address target,
        bytes memory input,
        bytes4 callbackSignature,
        uint256 callbackGasLimit,
        uint256 bounty,
        address token,
        uint256 maxFee
    ) public pure returns (uint256) {
        require(false, "MockOracleFail always fails with fee swap");
    }
}