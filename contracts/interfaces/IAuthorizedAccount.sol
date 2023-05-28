// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

interface IAuthorizedAccount {
    enum Permission {
        TicketSigning
    }

    struct AuthorizedAccount {
        address account; // Address of the authorized account
        uint256 createdAt; // Block number the authorized account was created
        Permission[] permissions; // Permissions associated with delegated account
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
        Permission permission
    ) external returns (bool);

    function getAuthorizedAccounts(
        address main
    ) external view returns (AuthorizedAccount[] memory);
}
