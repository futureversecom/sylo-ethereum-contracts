// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract StateOracleDemo {
    /// log on state oracle request
    event HiToEthereum(uint256 requestId);
    /// log on state oracle response
    event HiFromEthereum(uint256 requestId, uint256 timestamp, uint256 balance);
    /// CENNZnet ethereum state oracle precompile address
    address constant STATE_ORACLE = address(27572);

    /// Make a request for `remoteToken` balance of `who`
    ///
    /// @param remoteToken an ERC20 contract address on the remote Ethereum network
    /// @param who address to fetch the erc20 balanceOf on the remote Ethereum network
    function helloEthereum(address remoteToken, address who) payable external {
        bytes memory balanceOfCall = abi.encodeWithSignature("balanceOf(address)", who);
        bytes4 callbackSelector = this.ethereumSaysHi.selector;
        uint256 callbackGasLimit = 400_000;
        uint256 callbackBounty = 2 ether; // == 2 cpay

        // request a remote eth_call via the state oracle
        bytes memory remoteCallRequest = abi.encodeWithSignature(
            "remoteCall(address,bytes,bytes4,uint256,uint256)",
            remoteToken,
            balanceOfCall,
            callbackSelector,
            callbackGasLimit,
            callbackBounty
        );

        (bool success, bytes memory returnData) = STATE_ORACLE.call(remoteCallRequest);
        require(success);

        uint256 requestId = abi.decode(returnData, (uint256));
        emit HiToEthereum(requestId);
    }

    // Receive state oracle response
    function ethereumSaysHi(uint256 requestId, uint256 timestamp, bytes32 returnData) external {
        require(msg.sender == STATE_ORACLE, "must be state oracle");
        uint256 balanceOf = uint256(returnData);

        emit HiFromEthereum(
            requestId,
            timestamp,
            balanceOf
        );
    }
}