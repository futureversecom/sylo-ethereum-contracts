// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/Ownable2StepUpgradeable.sol";

import "./libraries/SyloUtils.sol";
import "./interfaces/IRegistries.sol";

/**
 * @notice This contract manages Registries for Nodes. A Registry is a
 * set of parameters configured by the Node itself. A Node is required
 * to have a valid Registry to be able to participate in the network.
 */
contract Registries is IRegistries, Initializable, Ownable2StepUpgradeable, IERC165 {
    using ECDSA for bytes32;

    uint256 private constant MAX_SEEKER_ID = 47894;

    /**
     * @notice ERC721 contract for bridged Seekers. Used for verifying ownership
     * of a seeker.
     */
    IERC721 public _rootSeekers;

    /**
     * @notice Tracks each Node's registry.
     */
    mapping(address => IRegistries.Registry) public registries;

    /**
     * @notice Tracks the node address that each seeker id is registered with
     */
    mapping(uint256 => address) public seekerRegistration;

    /**
     * @notice Tracks the address of every registered node.
     */
    address[] public nodes;

    /**
     * @notice Tracks nonces used when registering the seeker account
     * to prevent signature re-use.
     */
    mapping(bytes32 => address) private signatureNonces;

    /**
     * @notice Payout percentage refers to the portion of a tickets reward
     * that will be allocated to the Node's stakers. This is global, and is
     * currently set for all Nodes.
     */
    uint16 public defaultPayoutPercentage;

    /**
     * @notice Proof duration states the duration in blocks that the
     * proof used to validate seeker ownership will be valid for. The
     * `setSeekerAccount` transaction will revert if the proof message
     * was signed too many blocks ago.
     */
    uint16 public proofDuration;

    event DefaultPayoutPercentageUpdated(uint16 defaultPayoutPercentage);

    error SeekerIdOutOfRange();
    error NonceCannotBeReused();
    error EndMustBeGreaterThanStart();
    error PercentageCannotExceed10000();
    error PublicEndpointCannotBeEmpty();
    error SeekerAccountMustOwnSeekerId();
    error SeekerAccountMustBeMsgSender();
    error ProofNotSignedBySeekerAccount();
    error RootSeekersCannotBeZeroAddress();
    error SeekerAccountCannotBeZeroAddress();
    error EndCannotExceedNumberOfNodes(uint256 nodeLength);

    error Test(bytes4 id);

    function initialize(
        IERC721 rootSeekers,
        uint16 _defaultPayoutPercentage
    ) external initializer {
        if (address(rootSeekers) == address(0)) {
            revert RootSeekersCannotBeZeroAddress();
        }
        if (_defaultPayoutPercentage > 10000) {
            revert PercentageCannotExceed10000();
        }

        Ownable2StepUpgradeable.__Ownable2Step_init();

        _rootSeekers = rootSeekers;
        defaultPayoutPercentage = _defaultPayoutPercentage;
    }

    /**
     * @notice Returns true if the contract implements the interface defined by
     * `interfaceId` from ERC165.
     */
    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return interfaceId == type(IRegistries).interfaceId;
    }

    /**
     * @notice Set the global default payout percentage value. Only callable
     * by the owner.
     * @param _defaultPayoutPercentage The payout percentage as a value where the
     * denominator is 10000.
     */
    function setDefaultPayoutPercentage(uint16 _defaultPayoutPercentage) external onlyOwner {
        if (_defaultPayoutPercentage > 10000) {
            revert PercentageCannotExceed10000();
        }

        defaultPayoutPercentage = _defaultPayoutPercentage;
        emit DefaultPayoutPercentageUpdated(_defaultPayoutPercentage);
    }

    /**
     * @notice Call this as a Node to set or update your Registry entry.
     * @param publicEndpoint The public endpoint of your Node. Essential for
     * clients to be able to retrieve additional information, such as
     * an address to establish a p2p connection.
     */
    function register(string calldata publicEndpoint) external {
        if (bytes(publicEndpoint).length == 0) {
            revert PublicEndpointCannotBeEmpty();
        }

        // This is the nodes first registration
        if (bytes(registries[msg.sender].publicEndpoint).length == 0) {
            nodes.push(msg.sender);
        }

        registries[msg.sender].publicEndpoint = publicEndpoint;
    }

    function setSeekerAccount(
        address seekerAccount,
        uint256 seekerId,
        bytes32 nonce,
        bytes calldata signature
    ) external {
        if (seekerAccount == address(0)) {
            revert SeekerAccountCannotBeZeroAddress();
        }
        if (seekerId > MAX_SEEKER_ID) {
            revert SeekerIdOutOfRange();
        }
        if (signatureNonces[nonce] != address(0)) {
            revert NonceCannotBeReused();
        }

        bytes memory proofMessage = getProofMessage(seekerId, msg.sender, nonce);
        bytes32 ethProof = ECDSA.toEthSignedMessageHash(proofMessage);

        if (ECDSA.recover(ethProof, signature) != seekerAccount) {
            revert ProofNotSignedBySeekerAccount();
        }

        // Now verify the seeker account actually owns the seeker
        address owner = _rootSeekers.ownerOf(seekerId);

        if (seekerAccount != owner) {
            revert SeekerAccountMustOwnSeekerId();
        }

        delete registries[seekerRegistration[seekerId]].seekerId;

        registries[msg.sender].seekerAccount = seekerAccount;
        registries[msg.sender].seekerId = seekerId;

        seekerRegistration[seekerId] = msg.sender;

        signatureNonces[nonce] = seekerAccount;
    }

    function revokeSeekerAccount(address node) external {
        Registry storage registry = registries[node];

        if (registry.seekerAccount != msg.sender) {
            revert SeekerAccountMustBeMsgSender();
        }

        delete registry.seekerAccount;
        delete seekerRegistration[registry.seekerId];
        delete registry.seekerId;
    }

    /**
     * @notice Retrieve the registry associated with a Node.
     * @param account The address of the Node.
     * @return The Node's Registry.
     */
    function getRegistry(address account) external view returns (Registry memory) {
        return registries[account];
    }

    /**
     * @notice Retrieve all registered nodes.
     * @return An array of node addresses.
     */
    function getNodes() external view returns (address[] memory) {
        return nodes;
    }

    /**
     * @notice Retrieves a list of registries. Takes in a
     * a start and end indices to allow pagination.
     * @param start The start index which is inclusive.
     * @param end The end index which is exclusive.
     * @return An array of Registries.
     */
    function getRegistries(
        uint256 start,
        uint256 end
    ) external view returns (address[] memory, Registry[] memory) {
        uint256 nodesLength = nodes.length;

        if (end <= start) {
            revert EndMustBeGreaterThanStart();
        }
        if (end > nodesLength) {
            revert EndCannotExceedNumberOfNodes(nodesLength);
        }

        address[] memory _nodes = new address[](end - start);
        Registry[] memory _registries = new Registry[](_nodes.length);

        for (uint256 i = start; i < end; ++i) {
            _nodes[i - start] = nodes[i];
            _registries[i - start] = registries[nodes[i]];
        }

        return (_nodes, _registries);
    }

    /**
     * @notice Returns the total number of registered nodes.
     * @return The number of registered nodes.
     */
    function getTotalNodes() external view returns (uint256) {
        return nodes.length;
    }

    /**
     * @notice Helper function for deriving the proof message used to
     * validate seeker ownership.
     * @param seekerId The tokenId of the seeker used for operation.
     * @param node The address of the node which that will be operated
     * by the specified seeker.
     * @param nonce The nonce used for this message.
     */
    function getProofMessage(
        uint256 seekerId,
        address node,
        bytes32 nonce
    ) public pure returns (bytes memory) {
        return
            abi.encodePacked(
                unicode"🤖 Hi frend! 🤖\n\n📜 Signing this message proves that you're the owner of this Seeker NFT and allows your Seeker to be used to operate your Seeker's Node. It's a simple but important step to ensure smooth operation.\n\nThis request will not trigger a blockchain transaction or cost any gas fees.\n\n🔥 Your node's address: ",
                Strings.toHexString(uint256(uint160(node)), 20),
                unicode"\n\n🆔 Your seeker id: ",
                Strings.toString(seekerId),
                unicode"\n\n📦 A unique random value which secures this message: ",
                Strings.toHexString(uint256(nonce), 32)
            );
    }
}
