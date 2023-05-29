import { ethers } from 'hardhat';
import { Signer } from 'ethers';
import { AuthorizedAccount, SyloToken } from '../typechain-types';
import utils, { Contracts } from './utils';
import { assert, expect } from 'chai';

describe('Authorized Account', () => {
  let accounts: Signer[];
  let deployer: string;
  let mainAccount: Signer;
  let mainAccountAddress: string;
  let delegatedAccount1: string;
  let delegatedAccount2: string;
  let delegatedAccount3: string;

  let token: SyloToken;
  let authorizedAccount: AuthorizedAccount;
  let contracts: Contracts;

  enum Permission {
    DepositWithdrawal,
  }

  before(async () => {
    accounts = await ethers.getSigners();
    deployer = await accounts[0].getAddress();
    mainAccount = accounts[1];
    mainAccountAddress = await mainAccount.getAddress();
    delegatedAccount1 = await accounts[2].getAddress();
    delegatedAccount2 = await accounts[3].getAddress();
    delegatedAccount3 = await accounts[4].getAddress();

    const Token = await ethers.getContractFactory('SyloToken');
    token = await Token.deploy();
  });

  beforeEach(async () => {
    contracts = await utils.initializeContracts(deployer, token.address);
    authorizedAccount = contracts.authorizedAccount;
  });

  it('authorized account cannot be initialized again', async () => {
    await expect(authorizedAccount.initialize()).to.be.revertedWith(
      'Initializable: contract is already initialized',
    );
  });

  it('cannot add zero authorized account', async () => {
    const permission: Permission[] = [Permission.DepositWithdrawal];
    await expect(
      authorizedAccount
        .connect(mainAccount)
        .authorizeAccount(ethers.constants.AddressZero, permission),
    ).to.be.revertedWithCustomError(
      authorizedAccount,
      'AuthorizedAccountCannotBeZeroAddress',
    );
  });

  it('can add unexisted authorized account', async () => {
    const permission: Permission[] = [Permission.DepositWithdrawal];
    await authorizedAccount
      .connect(mainAccount)
      .authorizeAccount(delegatedAccount1, permission);
    const authorizedAccounts = await authorizedAccount
      .connect(mainAccount)
      .getAuthorizedAccounts(mainAccountAddress);
    assert.equal(authorizedAccounts[0].account, delegatedAccount1);
    assert.equal(authorizedAccounts.length, 1);
  });

  it('can add multiple unexisted authorized account', async () => {
    const permission: Permission[] = [Permission.DepositWithdrawal];
    await authorizedAccount
      .connect(mainAccount)
      .authorizeAccount(delegatedAccount1, permission);
    await authorizedAccount
      .connect(mainAccount)
      .authorizeAccount(delegatedAccount2, permission);
    const authorizedAccounts = await authorizedAccount
      .connect(mainAccount)
      .getAuthorizedAccounts(mainAccountAddress);
    assert.equal(authorizedAccounts[0].account, delegatedAccount1);
    assert.equal(authorizedAccounts[1].account, delegatedAccount2);
    assert.equal(authorizedAccounts.length, 2);
  });

  it('cannot add existed authorized account', async () => {
    const permission: Permission[] = [Permission.DepositWithdrawal];
    await authorizedAccount
      .connect(mainAccount)
      .authorizeAccount(delegatedAccount1, permission);
    await expect(
      authorizedAccount
        .connect(mainAccount)
        .authorizeAccount(delegatedAccount1, permission),
    ).to.be.revertedWithCustomError(authorizedAccount, 'AccountAlreadyExists');
  });

  it('cannot unauthorize invalid authorized account', async () => {
    await expect(
      authorizedAccount
        .connect(mainAccount)
        .unauthorizeAccount(ethers.constants.AddressZero),
    ).to.be.revertedWithCustomError(
      authorizedAccount,
      'AuthorizedAccountCannotBeZeroAddress',
    );
  });

  it('cannot unauthorize account if none account is available', async () => {
    await expect(
      authorizedAccount
        .connect(mainAccount)
        .unauthorizeAccount(delegatedAccount1),
    ).to.be.revertedWithCustomError(authorizedAccount, 'AccountDoesNotExist');
  });

  it('cannot unauthorize unexisted account', async () => {
    const permission: Permission[] = [Permission.DepositWithdrawal];
    await authorizedAccount
      .connect(mainAccount)
      .authorizeAccount(delegatedAccount1, permission);
    await expect(
      authorizedAccount
        .connect(mainAccount)
        .unauthorizeAccount(delegatedAccount2),
    ).to.be.revertedWithCustomError(authorizedAccount, 'AccountDoesNotExist');
  });

  it('can unauthorize existed account', async () => {
    const permission: Permission[] = [Permission.DepositWithdrawal];
    await authorizedAccount
      .connect(mainAccount)
      .authorizeAccount(delegatedAccount1, permission);
    await authorizedAccount
      .connect(mainAccount)
      .authorizeAccount(delegatedAccount2, permission);
    await authorizedAccount
      .connect(mainAccount)
      .authorizeAccount(delegatedAccount3, permission);
    let authorizedAccounts = await authorizedAccount
      .connect(mainAccount)
      .getAuthorizedAccounts(mainAccountAddress);
    assert.equal(authorizedAccounts[0].account, delegatedAccount1);
    assert.equal(authorizedAccounts[1].account, delegatedAccount2);
    assert.equal(authorizedAccounts[2].account, delegatedAccount3);
    assert.equal(authorizedAccounts.length, 3);

    await authorizedAccount
      .connect(mainAccount)
      .unauthorizeAccount(delegatedAccount2);
    authorizedAccounts = await authorizedAccount
      .connect(mainAccount)
      .getAuthorizedAccounts(mainAccountAddress);
    assert.equal(authorizedAccounts[0].account, delegatedAccount1);
    assert.equal(authorizedAccounts[1].account, delegatedAccount3);
    assert.equal(authorizedAccounts.length, 2);
  });

  it('cannot add permission for invalid delegated account', async () => {
    const permission: Permission[] = [Permission.DepositWithdrawal];
    await expect(
      authorizedAccount
        .connect(mainAccount)
        .addPermissions(ethers.constants.AddressZero, permission),
    ).to.be.revertedWithCustomError(
      authorizedAccount,
      'AuthorizedAccountCannotBeZeroAddress',
    );
  });

  it('cannot add permission for unexisted delegated account', async () => {
    const permission: Permission[] = [Permission.DepositWithdrawal];
    await authorizedAccount
      .connect(mainAccount)
      .authorizeAccount(delegatedAccount2, permission);
    await expect(
      authorizedAccount
        .connect(mainAccount)
        .addPermissions(delegatedAccount1, permission),
    ).to.be.revertedWithCustomError(authorizedAccount, 'AccountDoesNotExist');
  });

  it('can add permission for existed delegated account', async () => {
    const permission: Permission[] = [];
    const newPermission: Permission[] = [Permission.DepositWithdrawal];
    await authorizedAccount
      .connect(mainAccount)
      .authorizeAccount(delegatedAccount1, permission);
    await authorizedAccount
      .connect(mainAccount)
      .addPermissions(delegatedAccount1, newPermission);
    await authorizedAccount
      .connect(mainAccount)
      .addPermissions(delegatedAccount1, newPermission);
    const accounts = await authorizedAccount
      .connect(mainAccount)
      .getAuthorizedAccounts(mainAccountAddress);
    assert.equal(accounts[0].permissions.length, 1);
  });

  it('can add multiple permissions (with duplicated permissions) for existed delegated account', async () => {
    const permission: Permission[] = [];
    const newPermissions: Permission[] = [
      Permission.DepositWithdrawal,
      Permission.DepositWithdrawal,
    ];
    await authorizedAccount
      .connect(mainAccount)
      .authorizeAccount(delegatedAccount1, permission);
    await authorizedAccount
      .connect(mainAccount)
      .addPermissions(delegatedAccount1, newPermissions);
    const accounts = await authorizedAccount
      .connect(mainAccount)
      .getAuthorizedAccounts(mainAccountAddress);
    assert.equal(accounts[0].permissions.length, 1);
  });

  it('can add existed permission for current delegated account but permissions will not be duplicated', async () => {
    const permission: Permission[] = [Permission.DepositWithdrawal];
    await authorizedAccount
      .connect(mainAccount)
      .authorizeAccount(delegatedAccount1, permission);
    await authorizedAccount
      .connect(mainAccount)
      .addPermissions(delegatedAccount1, permission);
    const accounts = await authorizedAccount
      .connect(mainAccount)
      .getAuthorizedAccounts(mainAccountAddress);
    assert.equal(accounts[0].permissions.length, 1);
  });

  it('can remove multiple permissions (with duplicated permissions) for existed delegated account', async () => {
    const permissionsToAdd: Permission[] = [Permission.DepositWithdrawal];
    const permissionsToRemove: Permission[] = [
      Permission.DepositWithdrawal,
      Permission.DepositWithdrawal,
    ];
    const emptyPermissions: Permission[] = [];
    const permission2: Permission[] = [Permission.DepositWithdrawal];

    await authorizedAccount
      .connect(mainAccount)
      .authorizeAccount(delegatedAccount1, permissionsToAdd);
    await authorizedAccount
      .connect(mainAccount)
      .authorizeAccount(delegatedAccount2, permission2);

    let accounts = await authorizedAccount
      .connect(mainAccount)
      .getAuthorizedAccounts(mainAccountAddress);
    assert.equal(accounts[0].permissions.length, 1);

    await authorizedAccount
      .connect(mainAccount)
      .removePermissions(delegatedAccount1, emptyPermissions);

    accounts = await authorizedAccount
      .connect(mainAccount)
      .getAuthorizedAccounts(mainAccountAddress);
    assert.equal(accounts[0].permissions.length, 1);

    await authorizedAccount
      .connect(mainAccount)
      .removePermissions(delegatedAccount1, permissionsToRemove);

    accounts = await authorizedAccount
      .connect(mainAccount)
      .getAuthorizedAccounts(mainAccountAddress);
    assert.equal(accounts[0].permissions.length, 0);
  });

  it('cannot remove permission for unexisted delegated account', async () => {
    const permission: Permission[] = [Permission.DepositWithdrawal];
    await authorizedAccount
      .connect(mainAccount)
      .authorizeAccount(delegatedAccount2, permission);
    await expect(
      authorizedAccount
        .connect(mainAccount)
        .removePermissions(delegatedAccount1, permission),
    ).to.be.revertedWithCustomError(authorizedAccount, 'AccountDoesNotExist');
  });

  it('cannot remove permission with zero authorized account', async () => {
    const permission: Permission[] = [Permission.DepositWithdrawal];
    const authorizedAddress = ethers.constants.AddressZero;
    await expect(
      authorizedAccount
        .connect(mainAccount)
        .removePermissions(authorizedAddress, permission),
    ).to.be.revertedWithCustomError(
      authorizedAccount,
      'AuthorizedAccountCannotBeZeroAddress',
    );
  });

  it('can remove permission', async () => {
    const permission: Permission[] = [Permission.DepositWithdrawal];
    await authorizedAccount
      .connect(mainAccount)
      .authorizeAccount(delegatedAccount1, permission);
    await authorizedAccount
      .connect(mainAccount)
      .authorizeAccount(delegatedAccount2, permission);
    let accounts = await authorizedAccount
      .connect(mainAccount)
      .getAuthorizedAccounts(mainAccountAddress);
    assert.equal(accounts[0].permissions.length, 1);
    await authorizedAccount
      .connect(mainAccount)
      .removePermissions(delegatedAccount1, permission);
    accounts = await authorizedAccount
      .connect(mainAccount)
      .getAuthorizedAccounts(mainAccountAddress);
    assert.equal(accounts[0].permissions.length, 0);
  });

  it('cannot get authorized accounts associated with invalid main account', async () => {
    const main = ethers.constants.AddressZero;
    await expect(
      authorizedAccount.connect(mainAccount).getAuthorizedAccounts(main),
    ).to.be.revertedWithCustomError(
      authorizedAccount,
      'MainAccountCannotBeZeroAddress',
    );
  });

  it('cannot validate permission with invalid main address ', async () => {
    const main = ethers.constants.AddressZero;
    const permission = Permission.DepositWithdrawal;
    await expect(
      authorizedAccount
        .connect(mainAccount)
        .validatePermission(main, delegatedAccount1, permission),
    ).to.be.revertedWithCustomError(
      authorizedAccount,
      'MainAccountCannotBeZeroAddress',
    );
  });

  it('cannot validate permission with invalid authorized address ', async () => {
    const authorizedAddress = ethers.constants.AddressZero;
    const permission = Permission.DepositWithdrawal;
    await expect(
      authorizedAccount
        .connect(mainAccount)
        .validatePermission(mainAccountAddress, authorizedAddress, permission),
    ).to.be.revertedWithCustomError(
      authorizedAccount,
      'AuthorizedAccountCannotBeZeroAddress',
    );
  });

  it('return false when validating permission with unavailable authorized account ', async () => {
    const permission: Permission[] = [Permission.DepositWithdrawal];
    await authorizedAccount
      .connect(mainAccount)
      .authorizeAccount(delegatedAccount1, permission);
    const validate = await authorizedAccount
      .connect(mainAccount)
      .validatePermission(
        mainAccountAddress,
        delegatedAccount2,
        Permission.DepositWithdrawal,
      );
    assert.equal(validate, false);
  });

  it('return false if authorized account does not have valid permission', async () => {
    const permission: Permission[] = [];
    await authorizedAccount
      .connect(mainAccount)
      .authorizeAccount(delegatedAccount1, permission);
    const validate = await authorizedAccount
      .connect(mainAccount)
      .validatePermission(
        mainAccountAddress,
        delegatedAccount1,
        Permission.DepositWithdrawal,
      );
    assert.equal(validate, false);
  });

  it('return true if authorized account has valid permission', async () => {
    const permission: Permission[] = [Permission.DepositWithdrawal];
    await authorizedAccount
      .connect(mainAccount)
      .authorizeAccount(delegatedAccount1, permission);
    const validate = await authorizedAccount
      .connect(mainAccount)
      .validatePermission(
        mainAccountAddress,
        delegatedAccount1,
        Permission.DepositWithdrawal,
      );
    assert.equal(validate, true);
  });
});
