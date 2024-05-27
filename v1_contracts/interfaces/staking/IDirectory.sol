// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

interface IDirectory {
    /**
     * @dev A DirectoryEntry will be stored for every node that joins the
     * network in a specific epoch. The entry will contain the stakee's
     * address, and a boundary value which is a sum of the current directory's
     * total stake, and the current stakee's total stake.
     */
    struct DirectoryEntry {
        address stakee;
        uint256 boundary;
    }

    /**
     * @dev An EpochDirectory will be stored for every epoch. The
     * directory will be constructed piece by piece as Nodes join,
     * each adding their own directory entry based on their current
     * stake value.
     */
    struct EpochDirectory {
        DirectoryEntry[] entries;
        mapping(address => uint256) stakes;
        uint256 totalStake;
    }

    function setCurrentDirectory(uint256 epochId) external;

    function joinNextDirectory(address stakee, uint256 seekerId) external;

    function scan(uint128 point) external view returns (address stakee);

    function scanWithEpochId(
        uint128 point,
        uint256 epochId
    ) external view returns (address stakee);

    function getTotalStakeForStakee(
        uint256 epochId,
        address stakee
    ) external view returns (uint256);

    function getTotalStake(uint256 epochId) external view returns (uint256);

    function getEntries(
        uint256 epochId
    ) external view returns (address[] memory, uint256[] memory);
}
