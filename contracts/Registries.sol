// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.13;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

import "./ECDSA.sol";
import "./Seekers.sol";

/**
 * @notice This contract manages Registries for Nodes. A Registry is a
 * set of parameters configured by the Node itself. A Node is required
 * to have a valid Registry to be able to participate in the network.
 */
contract Registries is Initializable, OwnableUpgradeable {
    using ECDSA for bytes32;

    string public constant SEEKER_OWNERSHIP_PREFIX =
        "This message allows your seeker to be used to operate your node.";

    struct Registry {
        // Public http/s endpoint to retrieve additional metadata
        // about the node.
        // The current metadata schema is as follows:
        //  { name: string, multiaddrs: string[] }
        string publicEndpoint;
        // The account which owns a seeker that will be used to
        // operate the Node for this registry.
        address seekerAccount;
        // The id of the seeker used to operate the node. The owner
        // of this id should be the seeker account.
        uint256 seekerId;
        // Percentage of a tickets value that will be rewarded to
        // delegated stakers expressed as a fraction of 10000.
        // This value is currently locked to the default payout percentage
        // until epochs are implemented.
        uint16 payoutPercentage;
        // The minimum amount of stake that is required to
        // add delegated stake against this node
        uint256 minDelegatedStake;
    }

    /**
     * @notice Seeker's contract for verifying ownership of a seeker.
     */
    Seekers public _seekers;

    /**
     * @notice Tracks each Node's registry.
     */
    mapping(address => Registry) public registries;

    /**
     * @notice Tracks the node addresses that registered to the contract.
     */
    address[] public nodeAddresses;

    event DefaultPayoutPercentageUpdated(uint16 defaultPayoutPercentage);

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

    function initialize(
        Seekers seekers,
        uint16 _defaultPayoutPercentage,
        uint16 _proofDuration
    ) external initializer {
        OwnableUpgradeable.__Ownable_init();
        _seekers = seekers;
        require(
            _defaultPayoutPercentage <= 10000,
            "The payout percentage can not exceed 100 percent"
        );
        defaultPayoutPercentage = _defaultPayoutPercentage;
        proofDuration = _proofDuration;
    }

    /**
     * @notice Set the global default payout percentage value. Only callable
     * by the owner.
     * @param _defaultPayoutPercentage The payout percentage as a value where the
     * denominator is 10000.
     */
    function setDefaultPayoutPercentage(uint16 _defaultPayoutPercentage) external onlyOwner {
        require(
            _defaultPayoutPercentage <= 10000,
            "The payout percentage can not exceed 100 percent"
        );
        defaultPayoutPercentage = _defaultPayoutPercentage;
        emit DefaultPayoutPercentageUpdated(_defaultPayoutPercentage);
    }

    /**
     * @notice Call this as a Node to set or update your Registry entry.
     * @param publicEndpoint The public endpoint of your Node. Essential for
     * clients to be able to retrieve additional information, such as
     * an address to establish a p2p connection.
     * @param minDelegatedStake The minimum amount of stake in SOLO that
     * a staker must add when calling StakingManager.addStake.
     */
    function register(string memory publicEndpoint, uint256 minDelegatedStake) external {
        require(bytes(publicEndpoint).length != 0, "Public endpoint can not be empty");

        if (bytes(registries[msg.sender].publicEndpoint).length == 0) {
            nodeAddresses.push(msg.sender);
        }

        registries[msg.sender].publicEndpoint = publicEndpoint;
        registries[msg.sender].minDelegatedStake = minDelegatedStake;
    }

    function setSeekerAccount(
        address seekerAccount,
        uint256 seekerId,
        uint256 proofBlock,
        bytes memory signature
    ) external {
        require(block.number >= proofBlock, "Proof can not be set for a future block");
        require(block.number - proofBlock < proofDuration, "Proof is expired");

        bytes memory proofMessage = getProofMessage(seekerId, msg.sender, proofBlock);

        bytes32 proof = keccak256(
            abi.encodePacked(
                "\x19Ethereum Signed Message:\n",
                Strings.toString(proofMessage.length),
                proofMessage
            )
        );

        require(
            ECDSA.recover(proof, signature) == seekerAccount,
            "Proof must be signed by specified seeker account"
        );

        // Now verify the seeker account actually owns the seeker
        Seekers.Owner memory owner = _seekers.ownerOf(seekerId);

        require(seekerAccount == owner.owner, "Seeker account must own the specified seeker");

        registries[msg.sender].seekerAccount = seekerAccount;
        registries[msg.sender].seekerId = seekerId;
    }

    function revokeSeekerAccount(address node) external {
        Registry storage registry = registries[node];

        require(
            registry.seekerAccount == msg.sender,
            "Seeker account and msg.sender must be equal"
        );

        registry.seekerAccount = address(0);
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
     * @notice Retrieve all registries.
     * @return The array of Registry.
     */
    function getRegistries() external view returns (address[] memory, Registry[] memory) {
        Registry[] memory _registries = new Registry[](nodeAddresses.length);

        for (uint256 i = 0; i < nodeAddresses.length; i++) {
            _registries[i] = registries[nodeAddresses[i]];
        }

        return (nodeAddresses, _registries);
    }

    /**
     * @notice Retrieves the prefix used for creating proofs.
     */
    function getPrefix() public pure returns (string memory) {
        return SEEKER_OWNERSHIP_PREFIX;
    }

    /**
     * @notice Helper function for deriving the proof message used to
     * validate seeker ownership.
     * @param seekerId The tokenId of the seeker used for operation.
     * @param node The address of the node which that will be operated
     * by the specified seeker.
     * @param proofBlock The block the proof was generated in.
     */
    function getProofMessage(
        uint256 seekerId,
        address node,
        uint256 proofBlock
    ) public pure returns (bytes memory) {
        return
            abi.encodePacked(
                SEEKER_OWNERSHIP_PREFIX,
                ":",
                Strings.toString(seekerId),
                ":",
                Strings.toHexString(uint256(uint160(node)), 20),
                ":",
                Strings.toString(proofBlock)
            );
    }
}
