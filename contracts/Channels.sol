// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract Channels is Initializable, OwnableUpgradeable {
    mapping(address => uint8) public channels;

    event NewChannelUpdated(uint8 newChannel);

    function initialize() external initializer {
        OwnableUpgradeable.__Ownable_init();
    }

    function setChannel(uint8 channel) external {
        channels[msg.sender] = channel;
    }

    function getChannel() external view returns (uint8) {
        return channels[msg.sender];
    }
}
