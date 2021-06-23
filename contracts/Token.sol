// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract SyloToken is ERC20 {
    constructor() ERC20("Sylo", "SYLO") {
        _mint(msg.sender, 10000000000 ether);
    }
}
