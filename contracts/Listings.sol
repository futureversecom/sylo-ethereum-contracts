// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.13;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "./ECDSA.sol";
import "./Seekers.sol";

/**
 * @notice This contract manages Listings for Nodes. A Listing is a
 * set of parameters configured by the Node itself. A Node is required
 * to have a valid Listing to be able to participate in the network.
 */
contract Listings is Initializable, OwnableUpgradeable {
    using ECDSA for bytes32;

    string constant SEEKER_OWNERSHIP_PREFIX = "This message allows your seeker to be used to operate your node.";

    struct Listing {
        // Public http/s endpoint to retrieve additional metadata
        // about the node.
        // The current metadata schema is as follows:
        //  { name: string, multiaddrs: string[] }
        string publicEndpoint;

        // The account which owns a seeker that will be used to
        // operate the Node for this listing.
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

        // Explicit property to check if an instance of this struct actually exists
        bool initialized;
    }

    /**
     * @notice Seeker's contract for verifying ownership of a seeker.
     */
    Seekers public _seekers;

    /**
     * @notice Tracks each Node's listing.
     */
    mapping(address => Listing) public listings;

    event DefaultPayoutPercentageUpdated(uint16 defaultPayoutPercentage);

    /**
     * @notice Payout percentage refers to the portion of a tickets reward
     * that will be allocated to the Node's stakers. This is global, and is
     * currently set for all Nodes.
     */
    uint16 public defaultPayoutPercentage;

    function initialize(
        Seekers seekers,
        uint16 _defaultPayoutPercentage
    ) external initializer {
        OwnableUpgradeable.__Ownable_init();
        _seekers = _seekers;
        require(
            _defaultPayoutPercentage <= 10000,
            "The payout percentage can not exceed 100 percent"
        );
        defaultPayoutPercentage = _defaultPayoutPercentage;
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
     * @notice Call this as a Node to set or update your Listing entry.
     * @param publicEndpoint The public endpoint of your Node. Essential for
     * clients to be able to retrieve additional information, such as
     * an address to establish a p2p connection.
     * @param minDelegatedStake The minimum amount of stake in SOLO that
     * a staker must add when calling StakingManager.addStake.
     */
    function setListing(string memory publicEndpoint, uint256 minDelegatedStake) external {
        require(bytes(publicEndpoint).length != 0, "Public endpoint can not be empty");

        listings[msg.sender].publicEndpoint = publicEndpoint;
        listings[msg.sender].minDelegatedStake = minDelegatedStake;
    }

    function setSeekerAccount(address seekerAccount, uint256 seekerId, uint256 proofBlock, bytes memory signature) external {
        // Proofs are only valid for 100 blocks since they were signed
        require(block.number - proofBlock < 100, "Proof is expired");
        bytes32 proof = keccak256(
            abi.encode(
                SEEKER_OWNERSHIP_PREFIX, seekerId, msg.sender, proofBlock
            )
        ).toEthSignedMessageHash();

        require(ECDSA.recover(proof, signature) == seekerAccount, "Proof must be signed by specified seeker account");

        // Now verify the seeker account actually owns the seeker
        address owner = _seekers.ownerOf(seekerId);

        require(seekerAccount == owner, "Seeker account must own the specified seeker");

        listings[msg.sender].seekerAccount = seekerAccount;
        listings[msg.sender].seekerId = seekerId;
    }

    function revokeSeekerAccount(address node) external {
        Listing storage listing = listings[msg.sender];

        require(
            listing.seekerAccount == msg.sender,
            "Seeker account and msg.sender must be equal"
        );

        listing.seekerAccount = address(0);
    }

    /**
     * @notice Retrieve the listing associated with a Node.
     * @param account The address of the Node.
     * @return The Node's Listing.
     */
    function getListing(address account) external view returns (Listing memory) {
        return listings[account];
    }
}
