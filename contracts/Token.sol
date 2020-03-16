pragma solidity ^0.6.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract SyloToken is ERC20 {

  constructor() public {
    _mint(msg.sender, 10000000000 ether);
  }
}