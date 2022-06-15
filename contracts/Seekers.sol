// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.13;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "hardhat/console.sol";

contract Seekers is Initializable, OwnableUpgradeable {

    struct Owner {
        address owner;
        uint256 expiry;
    }

    mapping (uint256 => uint256) public requests;
    mapping (uint256 => Owner) private owners;

    /**
     * @notice The address of the Seekers NFT contract on ethereum mainnet.
     */
    address public seekers;

    /**
     * @notice The address of the token used for fees.
     */

    address public token;

    /**
     * @notice The address of the Oracle account performing these attestations.
     */
    address public oracle;

    /**
     * @notice The duration in blocks a response from the State Oracle
     * will be valid for. Any queries for ownership of a seeker that is older
     * than this value will return address(0).
     */
    uint256 public validDuration;

    /**
     * @notice The maximum gas the callback function should take to execute.
     */
    uint256 public callbackGasLimit;

    /**
     * @notice The amount in CPAY relayers will be rewarded for performing
     * the state bridge request.
     */
    uint256 public callbackBounty;

    function initialize(
        address _seekers,
        address _token,
        address _oracle,
        uint256 _validDuration,
        uint256 _callbackGasLimit,
        uint256 _callbackBounty
    ) external initializer {
        OwnableUpgradeable.__Ownable_init();
        seekers = _seekers;
        oracle = _oracle;
        validDuration = _validDuration;
        callbackGasLimit = _callbackGasLimit;
        callbackBounty = _callbackBounty;
    }

    function setSeekers(address _seekers) external onlyOwner {
        seekers = _seekers;
    }

    function setToken(address _token) external onlyOwner {
        token = _token;
    }

    function setOracle(address _oracle) external onlyOwner {
        oracle = _oracle;
    }

    function setValidDuration(uint256 _validDuration) external onlyOwner {
        validDuration = _validDuration;
    }

    function setCallbackGasLimit (uint256 _callbackGasLimit) external onlyOwner {
        callbackGasLimit = _callbackGasLimit;
    }

    function setCallbackBounty (uint256 _callbackBounty) external onlyOwner {
        callbackBounty = _callbackBounty;
    }

    function requestVerification(uint256 seekerId, uint256 maxFee) external payable {
        bytes memory ownerOfCall = abi.encodeWithSignature("ownerOf(uint256)", seekerId);
        bytes4 callbackSelector = this.confirmOwnership.selector;

        // request a remote eth_call via the state oracle
        bytes memory remoteCallRequest = abi.encodeWithSignature(
            "remoteCallWithFeeSwap(address,bytes,bytes4,uint256,uint256,address,uint256)",
            seekers,
            ownerOfCall,
            callbackSelector,
            callbackGasLimit,
            callbackBounty,
            token,
            maxFee
        );

        (bool success, bytes memory returnData) = oracle.call(remoteCallRequest);
        require(success, "oracle request failed");

        uint256 requestId = abi.decode(returnData, (uint256));

        requests[requestId] = seekerId;
    }


    function confirmOwnership(uint256 requestId, uint256 timestamp, bytes32 returnData) external {
        require(msg.sender == oracle, "must be state oracle");

        address owner = address(uint160(uint256(returnData)));

        uint256 seekerId = requests[requestId];

        owners[seekerId] = Owner(owner, block.number + validDuration);

        delete requests[requestId];
    }

    function ownerOf(uint256 seekerId) external view returns (address) {
        Owner memory owner = owners[seekerId];

        if (owner.owner == address(0)) {
            return address(0);
        }

        // State Oracle's response has become stale.
        if (owner.expiry < block.number) {
            return address(0);
        }

        return owner.owner;
    }

    function withdrawAll(address recipient) public onlyOwner {
        uint256 balance = address(this).balance;
        payable(recipient).transfer(balance);
    }
}