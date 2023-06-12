// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

interface IAuthorizedAccount {
    enum Permission {
        // TicketSigning permission allows the authorized account to
        // sign tickets for ticketing redemption.
        TicketSigning
    }

    /**
     * @dev This type will hold the permission type and the block number from
     * which the permission was set to be authorized and unauthorized.
     * The permission is authorized when authorizedAt >= unauthorizedAt.
     *
     * Note: authorizedAt and unauthorizedAt won't be set to 0 when the permission
     * is updated, because they are both needed when validating the permission.
     */
    struct AuthorizedPermission {
        // Permission type
        Permission permission;
        // Block number from which the permission was set to be authorized.
        // If the transaction is called in block 1, the permission is
        // authorized from block 1 (authorizedAt = block.number).
        uint256 authorizedAt;
        // Block number from which the permission was set to be unauthorized.
        // If the transaction is called in block 1, the permission is
        // unauthorized from block 2 (unauthorizedAt = block.number + 1)
        // unauthorizedAt is set that way to avoid the case where the
        // permission is authorized and unauthorized in the same block:
        // E.g. addPermission is called => authorizedAt = 1
        //      removePermission is called => unauthorizedAt = 1
        // => We cannot tell if the permission is authorized or not.
        // E.g. addPermission is called => authorizedAt = 1
        //      removePermission is called => unauthorizedAt = 2
        //      addPermission is called => authorizedAt = 1 AND update unauthorizedAt = authorizedAt = 1
        // => The permission is authorized when authorizedAt >= unauthorizedAt
        uint256 unauthorizedAt;
    }

    struct AuthorizedAccount {
        // The authorized account
        address account;
        // Block number at which the account was authorized.
        // If the transaction is called in block 1, the account is
        // authorized at block 1 (authorizedAt = block.number).
        // If the account is unauthorized, authorizedAt will be set to 0.
        uint256 authorizedAt;
        // Permission list
        AuthorizedPermission[] permissions;
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
