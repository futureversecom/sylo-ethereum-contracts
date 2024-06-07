// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "./SeekerStatsOracle.sol";

interface ISeekerStakingManager {
    struct StakedSeeker {
        uint256 seekerId;
        address node;
        address user;
    }

    function stakeSeeker(
        address node,
        SeekerStatsOracle.Seeker calldata seeker,
        bytes calldata seekerStatsProof
    ) external;

    function transferStakedSeeker(address fromNode, address toNode, uint256 seekerId) external;

    function unstakeSeeker(address node, uint256 seekerId) external;

    function getStakedSeekersByNode(address node) external view returns (uint256[] memory);

    function getStakedSeekersByUser(address node) external view returns (uint256[] memory);
}
