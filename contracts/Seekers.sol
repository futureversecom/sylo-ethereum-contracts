// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.13;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract Seekers is Initializable, OwnableUpgradeable {
    struct Owner {
        address owner;
        uint256 expiry;
    }

    // This is a mapping of requests Ids to Seeker tokenIds. This is
    // used in the Oracle callback function to understand which
    // tokenId the ownership request was made for.
    mapping(uint256 => uint256) public requests;

    // A mapping from Seeker tokenId and it's respective owner on Ethereum mainnet.
    mapping(uint256 => Owner) private owners;

    /**
     * @notice The address of the Seekers NFT contract on ethereum mainnet.
     */
    address public _seekers;

    /**
     * @notice The address of the token used for paying oracle fees if a fee swap
     * is used.
     */
    address public _token;

    /**
     * @notice The address of the Oracle account performing these attestations.
     */
    address public _oracle;

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
        address seekers,
        address token,
        address oracle,
        uint256 _validDuration,
        uint256 _callbackGasLimit,
        uint256 _callbackBounty
    ) external initializer {
        OwnableUpgradeable.__Ownable_init();
        _seekers = seekers;
        _oracle = oracle;
        _token = token;
        validDuration = _validDuration;
        callbackGasLimit = _callbackGasLimit;
        callbackBounty = _callbackBounty;
    }

    function setSeekers(address seekers) external onlyOwner {
        _seekers = seekers;
    }

    function setToken(address token) external onlyOwner {
        _token = token;
    }

    function setOracle(address oracle) external onlyOwner {
        _oracle = oracle;
    }

    function setValidDuration(uint256 _validDuration) external onlyOwner {
        validDuration = _validDuration;
    }

    function setCallbackGasLimit(uint256 _callbackGasLimit) external onlyOwner {
        callbackGasLimit = _callbackGasLimit;
    }

    function setCallbackBounty(uint256 _callbackBounty) external onlyOwner {
        callbackBounty = _callbackBounty;
    }

    function requestVerification(uint256 seekerId) external payable {
        bytes memory ownerOfCall = abi.encodeWithSignature("ownerOf(uint256)", seekerId);
        bytes4 callbackSelector = this.confirmOwnership.selector;

        // request a remote eth_call via the state oracle
        bytes memory remoteCallRequest = abi.encodeWithSignature(
            "remoteCall(address,bytes,bytes4,uint256,uint256)",
            _seekers,
            ownerOfCall,
            callbackSelector,
            callbackGasLimit,
            callbackBounty
        );

        (bool success, bytes memory returnData) = _oracle.call(remoteCallRequest);
        require(success, "Oracle request failed");

        uint256 requestId = abi.decode(returnData, (uint256));

        requests[requestId] = seekerId;
    }

    function requestVerificationWithFeeSwap(uint256 seekerId, uint256 maxFee) external payable {
        bytes memory ownerOfCall = abi.encodeWithSignature("ownerOf(uint256)", seekerId);
        bytes4 callbackSelector = this.confirmOwnership.selector;

        // request a remote eth_call via the state oracle
        bytes memory remoteCallRequest = abi.encodeWithSignature(
            "remoteCallWithFeeSwap(address,bytes,bytes4,uint256,uint256,address,uint256)",
            _seekers,
            ownerOfCall,
            callbackSelector,
            callbackGasLimit,
            callbackBounty,
            _token,
            maxFee
        );

        (bool success, bytes memory returnData) = _oracle.call(remoteCallRequest);
        require(success, "Oracle request failed");

        uint256 requestId = abi.decode(returnData, (uint256));

        requests[requestId] = seekerId;
    }

    function confirmOwnership(
        uint256 requestId,
        uint256 timestamp,
        bytes32 returnData
    ) external {
        require(msg.sender == _oracle, "msg.sender must be state oracle");

        address owner = address(uint160(uint256(returnData)));

        uint256 seekerId = requests[requestId];

        owners[seekerId] = Owner(owner, block.number + validDuration);

        delete requests[requestId];
    }

    function ownerOf(uint256 seekerId) external view returns (Owner memory) {
        Owner memory owner = owners[seekerId];

        // State Oracle's response has become stale.
        if (owner.expiry < block.number) {
            return Owner(address(0), 0);
        }

        return owner;
    }

    function withdrawAll(address recipient) public onlyOwner {
        uint256 balance = address(this).balance;
        payable(recipient).transfer(balance);
    }

    function withdrawAllFee(address recipient) public onlyOwner {
        IERC20 token = IERC20(_token);
        uint256 amount = token.balanceOf(address(this));
        token.transfer(recipient, amount);
    }
}
