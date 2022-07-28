// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract SyloToken is ERC20 {
    constructor() ERC20("Sylo", "SYLO") {
        _mint(msg.sender, 10_000_000_000 ether);
    }
}
