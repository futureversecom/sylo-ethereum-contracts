// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/access/Ownable2StepUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";

import "./interfaces/IAuthorizedAccount.sol";

/**
 * @notice Manages authorized accounts with limited permissions on behalf of main account
 * these authorized accounts are allowed to perform some certain actions in the Sylo network
 * in order to reduce the works for main account
 */
contract AuthorizedAccount is IAuthorizedAccount, Initializable, Ownable2StepUpgradeable, ERC165 {
    /**
     * @notice Tracks authorized accounts for every main account
     */
    mapping(address => AuthorizedAccount[]) public authorizedAccounts;

    error AuthorizedAccountCannotBeZeroAddress();
    error MainAccountCannotBeZeroAddress();
    error AccountAlreadyExists();
    error AccountDoesNotExist();

    function initialize() external initializer {
        Ownable2StepUpgradeable.__Ownable2Step_init();
    }

    /**
     * @notice Returns true if the contract implements the interface defined by
     * `interfaceId` from ERC165.
     */
    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return interfaceId == type(IAuthorizedAccount).interfaceId;
    }

    /**
     * @notice Adds new authorized accounts with certain permissions. This will revert if the account
     * has already existed
     * @param authorized The authorized address of the main account
     * @param permissions The list of permissions that the authorized account
     * can perform within the Sylo network.
     */
    function authorizeAccount(address authorized, Permission[] calldata permissions) external {
        if (authorized == address(0)) {
            revert AuthorizedAccountCannotBeZeroAddress();
        }

        // check if account has already existed
        AuthorizedAccount[] storage _authorizedAccounts = authorizedAccounts[msg.sender];
        for (uint i = 0; i < _authorizedAccounts.length; i++) {
            if (_authorizedAccounts[i].account == authorized) {
                revert AccountAlreadyExists();
            }
        }

        AuthorizedAccount memory newAccount = AuthorizedAccount({
            account: authorized,
            createdAt: block.number,
            permissions: permissions
        });

        authorizedAccounts[msg.sender].push(newAccount);
    }

    /**
     * @notice Removes an authorized account associated with the msg.sender.
     * This will revert if the account does not exist
     * @param authorized The address of the authorized account
     */
    function unauthorizeAccount(address authorized) external {
        if (authorized == address(0)) {
            revert AuthorizedAccountCannotBeZeroAddress();
        }

        AuthorizedAccount[] storage _authorizedAccounts = authorizedAccounts[msg.sender];
        for (uint i = 0; i < _authorizedAccounts.length; i++) {
            if (_authorizedAccounts[i].account == authorized) {
                _authorizedAccounts[i] = _authorizedAccounts[_authorizedAccounts.length - 1];
                _authorizedAccounts.pop();
                return;
            }
        }

        revert AccountDoesNotExist();
    }

    /**
     * @notice Adds new permissions to a specific authorized account. This will revert if
     * the account does not exist. Adding existing permissions will have no effect
     * @param authorized The address of authorized account
     * @param permissions The new permissions will be added for the authorized account
     */
    function addPermissions(address authorized, Permission[] calldata permissions) external {
        if (authorized == address(0)) {
            revert AuthorizedAccountCannotBeZeroAddress();
        }

        AuthorizedAccount[] storage _authorizedAccounts = authorizedAccounts[msg.sender];
        for (uint i = 0; i < _authorizedAccounts.length; i++) {
            if (_authorizedAccounts[i].account == authorized) {
                // iterate over the permissions
                for (uint j = 0; j < permissions.length; j++) {
                    // check if the permission already exists in the permissions array.
                    bool exists = false;
                    for (uint k = 0; k < _authorizedAccounts[i].permissions.length; k++) {
                        if (_authorizedAccounts[i].permissions[k] == permissions[j]) {
                            exists = true;
                            break;
                        }
                    }
                    // if the permission does not already exist, add it
                    // otherwise, just skip duplicated permission
                    if (!exists) {
                        _authorizedAccounts[i].permissions.push(permissions[j]);
                    }
                }
                return;
            }
        }
        revert AccountDoesNotExist();
    }

    /**
     * @notice Removes permissions of specific authorized account. This will revert if the account
     * does not exist.
     * @param authorized The address of authorized account
     * @param permissions The list of permissions will be removed
     */
    function removePermissions(address authorized, Permission[] calldata permissions) external {
        if (authorized == address(0)) {
            revert AuthorizedAccountCannotBeZeroAddress();
        }

        AuthorizedAccount[] storage _authorizedAccounts = authorizedAccounts[msg.sender];
        for (uint i = 0; i < _authorizedAccounts.length; i++) {
            if (_authorizedAccounts[i].account == authorized) {
                // create temporary array to store the permissions to keep
                Permission[] memory newPermissions = new Permission[](
                    _authorizedAccounts[i].permissions.length
                );

                uint counter = 0;
                // iterate over the existing permissions
                for (uint j = 0; j < _authorizedAccounts[i].permissions.length; j++) {
                    bool shouldRemove = false;
                    for (uint k = 0; k < permissions.length; k++) {
                        if (_authorizedAccounts[i].permissions[j] == permissions[k]) {
                            shouldRemove = true;
                            break;
                        }
                    }
                    // if the permission should not be removed, add it to the new permissions array.
                    if (!shouldRemove) {
                        newPermissions[counter] = _authorizedAccounts[i].permissions[j];
                        counter++;
                    }
                }

                // replace the old permissions array with the new one.
                delete _authorizedAccounts[i].permissions;
                for (uint j = 0; j < counter; j++) {
                    _authorizedAccounts[i].permissions.push(newPermissions[j]);
                }
                return;
            }
        }

        revert AccountDoesNotExist();
    }

    /**
     * @notice Validates permission of an authorized account associated with the main account.
     * @param main The address of main account
     * @param authorized The address of authorized account
     * @param permission The permission needs to be verified with the authorized account
     * @return boolean value
     */
    function validatePermission(
        address main,
        address authorized,
        Permission permission
    ) external view returns (bool) {
        if (main == address(0)) {
            revert MainAccountCannotBeZeroAddress();
        }

        if (authorized == address(0)) {
            revert AuthorizedAccountCannotBeZeroAddress();
        }

        AuthorizedAccount[] memory _authorizedAccounts = authorizedAccounts[main];
        for (uint i = 0; i < _authorizedAccounts.length; i++) {
            if (_authorizedAccounts[i].account == authorized) {
                for (uint j = 0; j < _authorizedAccounts[i].permissions.length; j++) {
                    if (_authorizedAccounts[i].permissions[j] == permission) {
                        return true;
                    }
                }
            }
        }
        return false;
    }

    /**
     * @notice Get all authorized accounts associated with a given account
     * @param main The address of main account
     * @return An array of authorized accounts
     */
    function getAuthorizedAccounts(
        address main
    ) external view returns (AuthorizedAccount[] memory) {
        if (main == address(0)) {
            revert MainAccountCannotBeZeroAddress();
        }

        return authorizedAccounts[main];
    }
}