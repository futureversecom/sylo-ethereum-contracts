// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

/**
 * @notice This contract manages Channels for users to
 * switch node during epoch.
 *
 * The channel value will be hashed in the Directory's scan function
 * together with user's public key and epochId to determine the
 * node address.
 */
contract Channels is Initializable, OwnableUpgradeable {
    /**
     * @notice Tracks each user's channel
     */
    mapping(address => uint8) public channels;

    /**
     * @notice Display value when setting new channel
     */
    event NewChannelUpdated(uint8 newChannel);

    function initialize() external initializer {
        OwnableUpgradeable.__Ownable_init();
    }

    /**
     * @notice Call this as a user to set or update your channel.
     * @param channel The channel value to set (0-48).
     */
    function setChannel(uint8 channel) external {
        require(channel <= 48, "Channels cannot be more than 48");

        channels[msg.sender] = channel;
        emit NewChannelUpdated(channel);
    }

    /**
     * @notice Retrieve the channel associated with the user.
     * @return The user's channel.
     */
    function getChannel() external view returns (uint8) {
        return channels[msg.sender];
    }
}
