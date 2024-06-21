// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "./SeekerStatsOracle.sol";

interface ISeekerStakingManager {
    struct StakedSeeker {
        uint256 seekerId; // bridged seeker id in TRN
        address node; // sylo node futureverse address
        address user; // msg.sender - the seeker owner, futureverse address in TRN
    }

    function stakeSeeker(
        address node,
        SeekerStatsOracle.Seeker calldata seeker,
        bytes calldata seekerStatsProof
    ) external;

    function stakeSeekers(
        address node,
        SeekerStatsOracle.Seeker[] calldata seekers,
        bytes[] calldata seekerStatsProofs
    ) external;

    function unstakeSeeker(address node, uint256 seekerId) external;

    function getStakedSeekersByNode(address node) external view returns (uint256[] memory);

    function getStakedSeekersByUser(address node) external view returns (uint256[] memory);
}
