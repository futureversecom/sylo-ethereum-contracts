// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

interface IAuthorizedAccount {
    enum Permission {
        TicketSigning
    }

    struct AuthorizedPermission {
        Permission permission; // Permission
        uint256 authorizedAt; // Block number the permission is last authorized (block.number)
        uint256 unauthorizedAt; // Block number the permission is started to be unauthorized (block.number + 1)
    }

    struct AuthorizedAccount {
        address account; // Authorized account
        uint256 authorizedAt; // Block number the main account authorized the authorized account, 0 if not authorized
        AuthorizedPermission[] permissions; // Permission list
    }

    function authorizeAccount(address authorized, Permission[] calldata permissions) external;

    function unauthorizeAccount(address authorized) external;

    function addPermissions(address authorized, Permission[] calldata permissions) external;

    function removePermissions(
        address authorized,
        Permission[] calldata permissionsToRemove
    ) external;

    function validatePermission(
        address main,
        address authorized,
        Permission permission,
        uint256 atBlock
    ) external returns (bool);

    function getAuthorizedAccounts(
        address main
    ) external view returns (AuthorizedAccount[] memory);
}
