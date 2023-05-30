import { ethers, network } from 'hardhat';
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
  let authAccountConnectMain: AuthorizedAccount;
  let contracts: Contracts;

  enum Permission {
    TicketSigning,
  }

  const permissionList: Permission[] = [Permission.TicketSigning];

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
    authAccountConnectMain = contracts.authorizedAccount.connect(mainAccount);
  });

  it('authorized account cannot be initialized again', async () => {
    await expect(authAccountConnectMain.initialize()).to.be.revertedWith(
      'Initializable: contract is already initialized',
    );
  });

  it('cannot add zero authorized account', async () => {
    await expect(
      authAccountConnectMain.authorizeAccount(
        ethers.constants.AddressZero,
        permissionList,
      ),
    ).to.be.revertedWithCustomError(
      authAccountConnectMain,
      'AuthorizedAccountCannotBeZeroAddress',
    );
  });

  it('can add unexisted authorized account', async () => {
    await authAccountConnectMain.authorizeAccount(
      delegatedAccount1,
      permissionList,
    );
    const authorizedAccounts =
      await authAccountConnectMain.getAuthorizedAccounts(mainAccountAddress);
    assert.equal(authorizedAccounts.length, 1);
    assert.equal(authorizedAccounts[0].account, delegatedAccount1);
    assert.equal(
      authorizedAccounts[0].permissions.length,
      permissionList.length,
    );
    assert.equal(
      authorizedAccounts[0].permissions[0].permission,
      permissionList[0],
    );
    assert.equal(
      authorizedAccounts[0].permissions[0].authorizedAt.toNumber(),
      await currentBlock(),
    );
    assert.equal(
      authorizedAccounts[0].permissions[0].unauthorizedAt.toNumber(),
      0,
    );
    assert.equal(
      authorizedAccounts[0].authorizedAt.toNumber(),
      await currentBlock(),
    );
  });

  it('can add multiple unexisted authorized account', async () => {
    await authAccountConnectMain.authorizeAccount(
      delegatedAccount1,
      permissionList,
    );

    await authAccountConnectMain.authorizeAccount(
      delegatedAccount2,
      permissionList,
    );

    const authorizedAccounts =
      await authAccountConnectMain.getAuthorizedAccounts(mainAccountAddress);

    assert.equal(authorizedAccounts[0].account, delegatedAccount1);
    assert.equal(authorizedAccounts[1].account, delegatedAccount2);
    assert.equal(authorizedAccounts.length, 2);
  });

  it('cannot add existed authorized account', async () => {
    await authAccountConnectMain.authorizeAccount(
      delegatedAccount1,
      permissionList,
    );
    await expect(
      authAccountConnectMain.authorizeAccount(
        delegatedAccount1,
        permissionList,
      ),
    ).to.be.revertedWithCustomError(
      authAccountConnectMain,
      'AccountAlreadyAuthorized',
    );
  });

  it('cannot unauthorize invalid authorized account', async () => {
    await expect(
      authAccountConnectMain.unauthorizeAccount(ethers.constants.AddressZero),
    ).to.be.revertedWithCustomError(
      authAccountConnectMain,
      'AuthorizedAccountCannotBeZeroAddress',
    );
  });

  it('cannot unauthorize account if none account is available', async () => {
    await expect(
      authAccountConnectMain.unauthorizeAccount(delegatedAccount1),
    ).to.be.revertedWithCustomError(
      authAccountConnectMain,
      'AccountDoesNotExist',
    );
  });

  it('cannot unauthorize unexisted account', async () => {
    await authAccountConnectMain.authorizeAccount(
      delegatedAccount1,
      permissionList,
    );
    await expect(
      authAccountConnectMain.unauthorizeAccount(delegatedAccount2),
    ).to.be.revertedWithCustomError(
      authAccountConnectMain,
      'AccountDoesNotExist',
    );
  });

  it('can unauthorize existed account', async () => {
    await authAccountConnectMain.authorizeAccount(
      delegatedAccount1,
      permissionList,
    );
    await authAccountConnectMain.authorizeAccount(
      delegatedAccount2,
      permissionList,
    );

    let result = authAccountConnectMain.authorizeAccount(
      delegatedAccount3,
      permissionList,
    );
    await expect(result)
      .to.emit(authAccountConnectMain, 'PermissionsAdded')
      .withArgs(
        await mainAccount.getAddress(),
        delegatedAccount3,
        permissionList,
      );

    let authorizedAccounts = await authAccountConnectMain.getAuthorizedAccounts(
      mainAccountAddress,
    );
    assert.equal(authorizedAccounts[0].account, delegatedAccount1);
    assert.equal(authorizedAccounts[1].account, delegatedAccount2);
    assert.equal(authorizedAccounts[2].account, delegatedAccount3);
    assert.equal(authorizedAccounts.length, 3);

    result = authAccountConnectMain.unauthorizeAccount(delegatedAccount2);
    await expect(result)
      .to.emit(authAccountConnectMain, 'PermissionsRemoved')
      .withArgs(
        await mainAccount.getAddress(),
        delegatedAccount2,
        permissionList,
      );

    authorizedAccounts = await authAccountConnectMain.getAuthorizedAccounts(
      mainAccountAddress,
    );

    assert.equal(authorizedAccounts.length, 3);
    assert.equal(authorizedAccounts[0].account, delegatedAccount1);
    assert.equal(authorizedAccounts[1].account, delegatedAccount2);
    assert.equal(authorizedAccounts[2].account, delegatedAccount3);

    assert.equal(authorizedAccounts[1].authorizedAt.toNumber(), 0);
    assert.equal(authorizedAccounts[1].permissions.length, 1);
    assert.equal(
      authorizedAccounts[1].permissions[0].permission,
      permissionList[0],
    );
    assert.equal(
      authorizedAccounts[1].permissions[0].unauthorizedAt.toNumber(),
      (await currentBlock()) + 1,
    );
    assert.isAbove(
      authorizedAccounts[1].permissions[0].unauthorizedAt.toNumber(),
      authorizedAccounts[1].permissions[0].authorizedAt.toNumber(),
    );
  });

  it('can unauthorize existed account in the same block after authorize account', async () => {
    await network.provider.send('evm_setAutomine', [false]);
    const block = await currentBlock();

    await authAccountConnectMain.authorizeAccount(
      delegatedAccount1,
      permissionList,
    );
    await authAccountConnectMain.unauthorizeAccount(delegatedAccount1);

    assert.equal(await currentBlock(), block);

    await network.provider.send('evm_mine');
    await network.provider.send('evm_setAutomine', [true]);

    const authorizedAccounts =
      await authAccountConnectMain.getAuthorizedAccounts(mainAccountAddress);

    assert.equal(authorizedAccounts.length, 1);
    assert.equal(authorizedAccounts[0].account, delegatedAccount1);
    assert.equal(authorizedAccounts[0].authorizedAt.toNumber(), 0);
    assert.equal(authorizedAccounts[0].permissions.length, 1);
    assert.equal(
      authorizedAccounts[0].permissions[0].permission,
      permissionList[0],
    );
    assert.equal(
      authorizedAccounts[0].permissions[0].authorizedAt.toNumber(),
      await currentBlock(),
    );
    assert.equal(
      authorizedAccounts[0].permissions[0].unauthorizedAt.toNumber(),
      (await currentBlock()) + 1,
    );
  });

  it('can authorize, unauthorize, and authorize again in one block', async () => {
    await network.provider.send('evm_setAutomine', [false]);
    const block = await currentBlock();

    await authAccountConnectMain.authorizeAccount(
      delegatedAccount1,
      permissionList,
    );
    await authAccountConnectMain.unauthorizeAccount(delegatedAccount1);
    await authAccountConnectMain.authorizeAccount(
      delegatedAccount1,
      permissionList,
    );

    assert.equal(await currentBlock(), block);

    await network.provider.send('evm_mine');
    await network.provider.send('evm_setAutomine', [true]);

    const authorizedAccounts =
      await authAccountConnectMain.getAuthorizedAccounts(mainAccountAddress);

    assert.equal(authorizedAccounts.length, 1);
    assert.equal(authorizedAccounts[0].account, delegatedAccount1);
    assert.equal(
      authorizedAccounts[0].authorizedAt.toNumber(),
      await currentBlock(),
    );
    assert.equal(authorizedAccounts[0].permissions.length, 1);
    assert.equal(
      authorizedAccounts[0].permissions[0].permission,
      permissionList[0],
    );
    assert.equal(
      authorizedAccounts[0].permissions[0].authorizedAt.toNumber(),
      await currentBlock(),
    );
    assert.equal(
      authorizedAccounts[0].permissions[0].unauthorizedAt.toNumber(),
      authorizedAccounts[0].permissions[0].authorizedAt.toNumber(),
    );
  });

  it('cannot add permission for invalid delegated account', async () => {
    await expect(
      authAccountConnectMain.addPermissions(
        ethers.constants.AddressZero,
        permissionList,
      ),
    ).to.be.revertedWithCustomError(
      authAccountConnectMain,
      'AuthorizedAccountCannotBeZeroAddress',
    );
  });

  it('cannot add permission for unexisted delegated account', async () => {
    await authAccountConnectMain.authorizeAccount(
      delegatedAccount2,
      permissionList,
    );
    await expect(
      authAccountConnectMain.addPermissions(delegatedAccount1, permissionList),
    ).to.be.revertedWithCustomError(
      authAccountConnectMain,
      'AccountDoesNotExist',
    );
  });

  it('can add permission for existed delegated account', async () => {
    const permission: Permission[] = [];
    const newPermission: Permission[] = [Permission.TicketSigning];
    await authAccountConnectMain.authorizeAccount(
      delegatedAccount1,
      permission,
    );

    let result = authAccountConnectMain.addPermissions(
      delegatedAccount1,
      permission,
    );

    await expect(result)
      .to.emit(authAccountConnectMain, 'PermissionsAdded')
      .withArgs(await mainAccount.getAddress(), delegatedAccount1, permission);

    result = authAccountConnectMain.addPermissions(
      delegatedAccount1,
      newPermission,
    );

    await expect(result)
      .to.emit(authAccountConnectMain, 'PermissionsAdded')
      .withArgs(
        await mainAccount.getAddress(),
        delegatedAccount1,
        newPermission,
      );

    const accounts = await authAccountConnectMain.getAuthorizedAccounts(
      mainAccountAddress,
    );

    assert.equal(accounts[0].permissions.length, 1);
    assert.equal(
      accounts[0].permissions[0].permission,
      Permission.TicketSigning,
    );
    assert.equal(
      accounts[0].permissions[0].authorizedAt.toNumber(),
      await currentBlock(),
    );
  });

  it('can add multiple permissions (with duplicated permissions) for existed delegated account', async () => {
    const permission: Permission[] = [];
    const newPermissions: Permission[] = [
      Permission.TicketSigning,
      Permission.TicketSigning,
    ];
    await authAccountConnectMain.authorizeAccount(
      delegatedAccount1,
      permission,
    );
    await authAccountConnectMain.addPermissions(
      delegatedAccount1,
      newPermissions,
    );
    const accounts = await authAccountConnectMain.getAuthorizedAccounts(
      mainAccountAddress,
    );
    assert.equal(accounts[0].permissions.length, 1);
  });

  it('can add existed permission for current delegated account but permissions will not be duplicated', async () => {
    await authAccountConnectMain.authorizeAccount(
      delegatedAccount1,
      permissionList,
    );
    await authAccountConnectMain.addPermissions(
      delegatedAccount1,
      permissionList,
    );
    const accounts = await authAccountConnectMain.getAuthorizedAccounts(
      mainAccountAddress,
    );
    assert.equal(accounts[0].permissions.length, 1);
  });

  it('can remove multiple permissions (with duplicated permissions) for existed delegated account', async () => {
    const permissionsToAdd: Permission[] = [Permission.TicketSigning];
    const permissionsToRemove: Permission[] = [
      Permission.TicketSigning,
      Permission.TicketSigning,
    ];
    const emptyPermissions: Permission[] = [];
    const permission2: Permission[] = [Permission.TicketSigning];

    await authAccountConnectMain.authorizeAccount(
      delegatedAccount1,
      permissionsToAdd,
    );
    const authorizedAtBlock = await currentBlock();

    await authAccountConnectMain.authorizeAccount(
      delegatedAccount2,
      permission2,
    );

    let accounts = await authAccountConnectMain.getAuthorizedAccounts(
      mainAccountAddress,
    );
    assert.equal(accounts[0].permissions.length, 1);

    let result = authAccountConnectMain.removePermissions(
      delegatedAccount1,
      emptyPermissions,
    );

    await expect(result)
      .to.emit(authAccountConnectMain, 'PermissionsRemoved')
      .withArgs(
        await mainAccount.getAddress(),
        delegatedAccount1,
        emptyPermissions,
      );

    accounts = await authAccountConnectMain.getAuthorizedAccounts(
      mainAccountAddress,
    );
    assert.equal(accounts[0].permissions.length, 1);

    result = authAccountConnectMain.removePermissions(
      delegatedAccount1,
      permissionsToRemove,
    );

    await expect(result)
      .to.emit(authAccountConnectMain, 'PermissionsRemoved')
      .withArgs(
        await mainAccount.getAddress(),
        delegatedAccount1,
        permissionsToRemove,
      );

    accounts = await authAccountConnectMain.getAuthorizedAccounts(
      mainAccountAddress,
    );

    assert.equal(accounts[0].permissions.length, 1);
    assert.equal(
      accounts[0].permissions[0].permission,
      Permission.TicketSigning,
    );
    assert.equal(
      accounts[0].permissions[0].authorizedAt.toNumber(),
      authorizedAtBlock,
    );
    assert.equal(
      accounts[0].permissions[0].unauthorizedAt.toNumber(),
      (await currentBlock()) + 1,
    );
    assert.isAbove(
      accounts[0].permissions[0].unauthorizedAt.toNumber(),
      accounts[0].permissions[0].authorizedAt.toNumber(),
    );
  });

  it('cannot remove permission for unexisted delegated account', async () => {
    await authAccountConnectMain.authorizeAccount(
      delegatedAccount2,
      permissionList,
    );
    await expect(
      authAccountConnectMain.removePermissions(
        delegatedAccount1,
        permissionList,
      ),
    ).to.be.revertedWithCustomError(
      authAccountConnectMain,
      'AccountDoesNotExist',
    );
  });

  it('cannot remove permission with zero authorized account', async () => {
    const authorizedAddress = ethers.constants.AddressZero;
    await expect(
      authAccountConnectMain.removePermissions(
        authorizedAddress,
        permissionList,
      ),
    ).to.be.revertedWithCustomError(
      authAccountConnectMain,
      'AuthorizedAccountCannotBeZeroAddress',
    );
  });

  it('can remove permission', async () => {
    await authAccountConnectMain.authorizeAccount(
      delegatedAccount1,
      permissionList,
    );
    const authorizedAtBlock = await currentBlock();

    await authAccountConnectMain.authorizeAccount(
      delegatedAccount2,
      permissionList,
    );

    let accounts = await authAccountConnectMain.getAuthorizedAccounts(
      mainAccountAddress,
    );

    assert.equal(accounts[0].permissions.length, 1);

    await authAccountConnectMain.removePermissions(
      delegatedAccount1,
      permissionList,
    );

    accounts = await authAccountConnectMain.getAuthorizedAccounts(
      mainAccountAddress,
    );

    assert.equal(accounts[0].permissions.length, 1);
    assert.equal(accounts[0].permissions[0].permission, permissionList[0]);
    assert.equal(
      accounts[0].permissions[0].authorizedAt.toNumber(),
      authorizedAtBlock,
    );
    assert.equal(
      accounts[0].permissions[0].unauthorizedAt.toNumber(),
      (await currentBlock()) + 1,
    );
  });

  it('cannot get authorized accounts associated with invalid main account', async () => {
    const main = ethers.constants.AddressZero;
    await expect(
      authAccountConnectMain.getAuthorizedAccounts(main),
    ).to.be.revertedWithCustomError(
      authAccountConnectMain,
      'MainAccountCannotBeZeroAddress',
    );
  });

  it('cannot validate permission with invalid main address ', async () => {
    const main = ethers.constants.AddressZero;
    const permission = Permission.TicketSigning;
    await expect(
      authAccountConnectMain.validatePermission(
        main,
        delegatedAccount1,
        permission,
        await currentBlock(),
      ),
    ).to.be.revertedWithCustomError(
      authAccountConnectMain,
      'MainAccountCannotBeZeroAddress',
    );
  });

  it('cannot validate permission with invalid authorized address ', async () => {
    const authorizedAddress = ethers.constants.AddressZero;
    const permission = Permission.TicketSigning;
    await expect(
      authAccountConnectMain.validatePermission(
        mainAccountAddress,
        authorizedAddress,
        permission,
        await currentBlock(),
      ),
    ).to.be.revertedWithCustomError(
      authAccountConnectMain,
      'AuthorizedAccountCannotBeZeroAddress',
    );
  });

  it('cannot validate permission with invalid atBlock', async () => {
    await authAccountConnectMain.authorizeAccount(
      delegatedAccount1,
      permissionList,
    );
    await expect(
      authAccountConnectMain.validatePermission(
        mainAccountAddress,
        delegatedAccount1,
        permissionList,
        0,
      ),
    ).to.be.revertedWithCustomError(
      authAccountConnectMain,
      'AtBlockNumberCannotBeZero',
    );
  });

  it('return false if the account is never authorized', async () => {
    const validate = await authAccountConnectMain.validatePermission(
      mainAccountAddress,
      delegatedAccount1,
      Permission.TicketSigning,
      await currentBlock(),
    );
    assert.equal(validate, false);
  });

  it('return false when validating permission with invalid authorized account', async () => {
    await authAccountConnectMain.authorizeAccount(
      delegatedAccount1,
      permissionList,
    );

    const validate = await authAccountConnectMain.validatePermission(
      mainAccountAddress,
      delegatedAccount2,
      Permission.TicketSigning,
      await currentBlock(),
    );

    assert.equal(validate, false);
  });

  it('return false if authorized account does not have valid permission', async () => {
    const permission: Permission[] = [];
    await authAccountConnectMain.authorizeAccount(
      delegatedAccount1,
      permission,
    );

    const validate = await authAccountConnectMain.validatePermission(
      mainAccountAddress,
      delegatedAccount1,
      Permission.TicketSigning,
      await currentBlock(),
    );

    assert.equal(validate, false);
  });

  it('can validate multiple cases with different atBlock', async () => {
    const authContract = authAccountConnectMain;

    /**
     * Symbol:
     * A: authorizedAt
     * U: unauthorizedAt
     * B: atBlock
     */

    // A = 0 => false
    let validate = await authContract.validatePermission(
      mainAccountAddress,
      delegatedAccount1,
      Permission.TicketSigning,
      1000,
    );
    assert.equal(validate, false);

    await authContract.authorizeAccount(delegatedAccount1, permissionList);
    let authBlock = await currentBlock();

    // U = 0 && A < B => true
    validate = await authContract.validatePermission(
      mainAccountAddress,
      delegatedAccount1,
      Permission.TicketSigning,
      1000,
    );
    assert.equal(validate, true);

    // A < B < U => true
    await utils.advanceBlock(5);
    await authContract.unauthorizeAccount(delegatedAccount1);
    let unauthBlock = (await currentBlock()) + 1; // unauthorizeAt = block.number + 1

    validate = await authContract.validatePermission(
      mainAccountAddress,
      delegatedAccount1,
      Permission.TicketSigning,
      (await currentBlock()) - 3,
    );
    assert.equal(validate, true);

    // A = B < U => true
    validate = await authContract.validatePermission(
      mainAccountAddress,
      delegatedAccount1,
      Permission.TicketSigning,
      authBlock,
    );
    assert.equal(validate, true);

    // B < A < U => false
    validate = await authContract.validatePermission(
      mainAccountAddress,
      delegatedAccount1,
      Permission.TicketSigning,
      authBlock - 1,
    );
    assert.equal(validate, false);

    // A < U = B => false
    validate = await authContract.validatePermission(
      mainAccountAddress,
      delegatedAccount1,
      Permission.TicketSigning,
      unauthBlock + 1,
    );
    assert.equal(validate, false);

    // A < U < B => false
    validate = await authContract.validatePermission(
      mainAccountAddress,
      delegatedAccount1,
      Permission.TicketSigning,
      unauthBlock + 123,
    );
    assert.equal(validate, false);

    // U = A = B (A is called after U) => true
    await authContract.authorizeAccount(delegatedAccount1, permissionList);
    authBlock = await currentBlock();

    validate = await authContract.validatePermission(
      mainAccountAddress,
      delegatedAccount1,
      Permission.TicketSigning,
      authBlock,
    );
    assert.equal(validate, true);

    // U = A < B => true
    validate = await authContract.validatePermission(
      mainAccountAddress,
      delegatedAccount1,
      Permission.TicketSigning,
      authBlock + 10,
    );
    assert.equal(validate, true);

    // U = A > B => false
    validate = await authContract.validatePermission(
      mainAccountAddress,
      delegatedAccount1,
      Permission.TicketSigning,
      authBlock - 3,
    );
    assert.equal(validate, false);

    // B < U < A => false
    await authContract.unauthorizeAccount(delegatedAccount1);
    unauthBlock = (await currentBlock()) + 1; // unauthorizeAt = block.number + 1
    await utils.advanceBlock(1);
    await authContract.authorizeAccount(delegatedAccount1, permissionList);
    authBlock = await currentBlock();

    validate = await authContract.validatePermission(
      mainAccountAddress,
      delegatedAccount1,
      Permission.TicketSigning,
      authBlock - 3,
    );
    assert.equal(validate, false);

    // U < A < B => true
    validate = await authContract.validatePermission(
      mainAccountAddress,
      delegatedAccount1,
      Permission.TicketSigning,
      authBlock + 3,
    );
    assert.equal(validate, true);

    // U < B < A => false
    validate = await authContract.validatePermission(
      mainAccountAddress,
      delegatedAccount1,
      Permission.TicketSigning,
      authBlock - 1,
    );
    assert.equal(validate, false);
  });

  it('return true if authorized account has valid permission', async () => {
    await authAccountConnectMain.authorizeAccount(
      delegatedAccount1,
      permissionList,
    );

    const validate = await authAccountConnectMain.validatePermission(
      mainAccountAddress,
      delegatedAccount1,
      Permission.TicketSigning,
      (await currentBlock()) + 1,
    );

    assert.equal(validate, true);
  });

  async function currentBlock() {
    return await ethers.provider.getBlockNumber();
  }
});
