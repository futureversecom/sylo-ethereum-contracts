// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

/**
 * @notice This contract manages Listings for Nodes. A Listing is a
 * set of parameters configured by the Node itself. A Node is required
 * to have a valid Listing to be able to participate in the network.
 */
contract Listings is Initializable, OwnableUpgradeable {

    struct Listing {
        // MultiAddr to connect to the account
        string multiAddr;

        // Percentage of a tickets value that will be rewarded to
        // delegated stakers expressed as a fraction of 10000.
        // This value is currently locked to the default payout percentage
        // until epochs are implemented.
        uint16 payoutPercentage;

        // The minimum amount of stake that is required to
        // add a delegated stake against this node
        uint256 minDelegatedStake;

        // Explicit property to check if an instance of this struct actually exists
        bool initialized;
    }

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

    function initialize(uint16 _defaultPayoutPercentage) external initializer {
        OwnableUpgradeable.__Ownable_init();
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
     * @param multiAddr The libp2p multiAddr of your Node. Essential for
     * clients to be able to establish a p2p connection.
     * @param minDelegatedStake The minimum amount of stake in SOLO that
     * a staker must add when calling StakingManager.addStake.
     */
    function setListing(string memory multiAddr, uint256 minDelegatedStake) external {
        require(bytes(multiAddr).length != 0, "Multiaddr string is empty");

        // TODO Remove defaultPayoutPercentage once epochs are introduced
        Listing memory listing = Listing(multiAddr, defaultPayoutPercentage, minDelegatedStake, true);
        listings[msg.sender] = listing;
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
