// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

// Basic ERC721 contract with an unrestricted mint function.
// Useful for mimicking the Seekers ERC721 contract for testing
// purposes.
contract TestSeekers is ERC721 {
    constructor() ERC721("Seekers", "SEEKERS") {}

    function mint(address to, uint256 tokenId) external {
        _safeMint(to, tokenId);
    }

    function exists(uint256 tokenId) external view returns (bool) {
        return _exists(tokenId);
    }
}
