import { ethers } from 'hardhat';
import { assert, expect } from 'chai';
import { Signer, Wallet } from 'ethers';
import { BigNumber } from '@ethersproject/bignumber';
import {
  AuthorizedAccounts,
  Directory,
  EpochsManager,
  Registries,
  RewardsManager,
  StakingManager,
  SyloTicketing,
  SyloToken,
  TestSeekers,
  TicketingParameters,
  ISyloTicketing__factory,
  TestFuturepassRegistrar,
} from '../../typechain-types';
import crypto from 'crypto';
import {
  compareExpectedBalance,
  toSOLOs,
  addStakes,
  testClaims,
  checkAfterRedeem,
  setSeekerRegistry,
  createWinningTicket,
  createCommit,
} from './utils';
import utils from '../utils';

describe('Ticketing', () => {
  let accounts: Signer[];
  let owner: string;

  const faceValue = toSOLOs(1000);
  const epochDuration = 1;

  let token: SyloToken;
  let epochsManager: EpochsManager;
  let rewardsManager: RewardsManager;
  let syloTicketing: SyloTicketing;
  let ticketingParameters: TicketingParameters;
  let directory: Directory;
  let registries: Registries;
  let stakingManager: StakingManager;
  let seekers: TestSeekers;
  let authorizedAccounts: AuthorizedAccounts;
  let futurepassRegistrar: TestFuturepassRegistrar;

  enum Permission {
    TicketSigning,
  }

  before(async () => {
    accounts = await ethers.getSigners();
    // first account is implicitly used as deployer of contracts in hardhat
    owner = await accounts[0].getAddress();

    // add more fund to owner account so it can fund main wallets
    // to authorize accounts and add permissions
    // account 10 to 15 are unused so we can get ETH from them
    for (let i = 10; i <= 15; i++) {
      await accounts[i].sendTransaction({
        to: owner,
        value: ethers.parseEther('5000.0'),
      });
    }

    const Token = await ethers.getContractFactory('SyloToken');
    token = await Token.deploy();
  });

  beforeEach(async () => {
    const contracts = await utils.initializeContracts(owner, token, {
      faceValue,
      epochDuration,
    });
    epochsManager = contracts.epochsManager;
    rewardsManager = contracts.rewardsManager;
    syloTicketing = contracts.syloTicketing;
    ticketingParameters = contracts.ticketingParameters;
    directory = contracts.directory;
    registries = contracts.registries;
    stakingManager = contracts.stakingManager;
    seekers = contracts.seekers;
    authorizedAccounts = contracts.authorizedAccounts;
    futurepassRegistrar = contracts.futurepassRegistrar;

    await token.approve(await stakingManager.getAddress(), toSOLOs(10000000));
    await token.approve(await syloTicketing.getAddress(), toSOLOs(10000000));
  });

  it('ticketing cannot be initialized twice', async () => {
    await expect(
      syloTicketing.initialize(
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        0,
      ),
    ).to.be.revertedWith('Initializable: contract is already initialized');
  });

  it('ticketing cannot be initialized with invalid parameters', async () => {
    const Ticketing = await ethers.getContractFactory('SyloTicketing');
    const ticketing = await Ticketing.deploy();

    await expect(
      ticketing.initialize(
        ethers.ZeroAddress,
        await registries.getAddress(),
        await stakingManager.getAddress(),
        await directory.getAddress(),
        await epochsManager.getAddress(),
        await rewardsManager.getAddress(),
        await authorizedAccounts.getAddress(),
        await futurepassRegistrar.getAddress(),
        0,
      ),
    ).to.be.revertedWithCustomError(ticketing, 'TokenCannotBeZeroAddress');

    await expect(
      ticketing.initialize(
        await token.getAddress(),
        await registries.getAddress(),
        await stakingManager.getAddress(),
        await directory.getAddress(),
        await epochsManager.getAddress(),
        await rewardsManager.getAddress(),
        await authorizedAccounts.getAddress(),
        await futurepassRegistrar.getAddress(),
        0,
      ),
    ).to.be.revertedWithCustomError(ticketing, 'UnlockDurationCannotBeZero');
  });

  it('ticketing parameters cannot be initialized twice', async () => {
    await expect(
      ticketingParameters.initialize(1, 1, 1, 1, 1, {
        from: owner,
      }),
    ).to.be.revertedWith('Initializable: contract is already initialized');
  });

  it('ticketing parameters cannot be initialized with invalid parameters', async () => {
    const TicketingParameters = await ethers.getContractFactory(
      'TicketingParameters',
    );
    const ticketingParameters = await TicketingParameters.deploy();

    await expect(
      ticketingParameters.initialize(0, 1, 1, 1, 1),
    ).to.be.revertedWithCustomError(
      ticketingParameters,
      'FaceValueCannotBeZero',
    );

    await expect(
      ticketingParameters.initialize(1, 1, 1, 1, 0),
    ).to.be.revertedWithCustomError(
      ticketingParameters,
      'TicketDurationCannotBeZero',
    );
  });

  it('rewards manager cannot be initialized twice', async () => {
    await expect(
      rewardsManager.initialize(
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        ethers.ZeroAddress,
      ),
    ).to.be.revertedWith('Initializable: contract is already initialized');
  });

  it('rewards manager cannot be initialized with invalid parameters', async () => {
    const RewardsManager = await ethers.getContractFactory('RewardsManager');
    const rewardsManager = await RewardsManager.deploy();

    await expect(
      rewardsManager.initialize(
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        ethers.ZeroAddress,
      ),
    ).to.be.revertedWithCustomError(rewardsManager, 'TokenCannotBeZeroAddress');

    await expect(
      rewardsManager.initialize(
        await token.getAddress(),
        ethers.ZeroAddress,
        ethers.ZeroAddress,
      ),
    )
      .to.be.revertedWithCustomError(
        rewardsManager,
        'TargetContractCannotBeZeroAddress',
      )
      .withArgs('StakingManager');

    await expect(
      rewardsManager.initialize(
        await token.getAddress(),
        await stakingManager.getAddress(),
        ethers.ZeroAddress,
      ),
    )
      .to.be.revertedWithCustomError(
        rewardsManager,
        'TargetContractCannotBeZeroAddress',
      )
      .withArgs('EpochsManager');
  });

  it('ticketing can check for support interface', async () => {
    let interfaceId = BigNumber.from(0);
    ISyloTicketing__factory.createInterface().forEachFunction(f => {
      interfaceId = interfaceId.xor(
        ethers.id(f.format('sighash')).substring(0, 10),
      );
    });

    const supportInterface = await syloTicketing.supportsInterface(
      interfaceId.toHexString(),
    );
    expect(supportInterface).to.be.true;
  });

  it('should be able to set parameters after initialization', async () => {
    await expect(ticketingParameters.setFaceValue(777))
      .to.emit(ticketingParameters, 'FaceValueUpdated')
      .withArgs(777);

    await expect(ticketingParameters.setBaseLiveWinProb(888))
      .to.emit(ticketingParameters, 'BaseLiveWinProbUpdated')
      .withArgs(888);

    await expect(ticketingParameters.setExpiredWinProb(999))
      .to.emit(ticketingParameters, 'ExpiredWinProbUpdated')
      .withArgs(999);

    await expect(ticketingParameters.setDecayRate(1111))
      .to.emit(ticketingParameters, 'DecayRateUpdated')
      .withArgs(1111);

    await expect(ticketingParameters.setTicketDuration(2222))
      .to.emit(ticketingParameters, 'TicketDurationUpdated')
      .withArgs(2222);

    await expect(syloTicketing.setUnlockDuration(3333))
      .to.emit(syloTicketing, 'UnlockDurationUpdated')
      .withArgs(3333);

    const currentfaceValue = await ticketingParameters.faceValue();
    assert.equal(
      currentfaceValue,
      777n,
      'Expected face value to be correctly set',
    );

    const baseLiveWinProb = await ticketingParameters.baseLiveWinProb();
    assert.equal(
      baseLiveWinProb,
      888n,
      'Expected base live win prob to be correctly set',
    );

    const expiredWinProb = await ticketingParameters.expiredWinProb();
    assert.equal(
      expiredWinProb,
      999n,
      'Expected expired win prob to be correctly set',
    );

    const decayRate = await ticketingParameters.decayRate();
    assert.equal(decayRate, 1111n, 'Expected decay rate to be correctly set');

    const ticketDuration = await ticketingParameters.ticketDuration();
    assert.equal(
      ticketDuration,
      2222n,
      'Expected ticket duration to be correctly set',
    );

    const unlockDuration = await syloTicketing.unlockDuration();
    assert.equal(
      unlockDuration,
      3333n,
      'Expected unlock duration to be correctly set',
    );
  });

  it('not owner cannot set parameters', async () => {
    const notOwner = accounts[1];

    await expect(
      ticketingParameters.connect(notOwner).setFaceValue(777),
    ).to.be.revertedWith('Ownable: caller is not the owner');

    await expect(
      ticketingParameters.connect(notOwner).setBaseLiveWinProb(888),
    ).to.be.revertedWith('Ownable: caller is not the owner');

    await expect(
      ticketingParameters.connect(notOwner).setExpiredWinProb(999),
    ).to.be.revertedWith('Ownable: caller is not the owner');

    await expect(
      ticketingParameters.connect(notOwner).setDecayRate(1111),
    ).to.be.revertedWith('Ownable: caller is not the owner');

    await expect(
      ticketingParameters.connect(notOwner).setTicketDuration(2222),
    ).to.be.revertedWith('Ownable: caller is not the owner');

    await expect(
      syloTicketing.connect(notOwner).setUnlockDuration(3333),
    ).to.be.revertedWith('Ownable: caller is not the owner');
  });

  it('ticketing cannot set parameters with invalid arguments', async () => {
    await expect(
      syloTicketing.setUnlockDuration(0),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'UnlockDurationCannotBeZero',
    );
  });

  it('ticketing parameters cannot set parameters with invalid arguments', async () => {
    await expect(
      ticketingParameters.setFaceValue(0),
    ).to.be.revertedWithCustomError(
      ticketingParameters,
      'FaceValueCannotBeZero',
    );

    await expect(
      ticketingParameters.setTicketDuration(0),
    ).to.be.revertedWithCustomError(
      ticketingParameters,
      'TicketDurationCannotBeZero',
    );
  });

  it('can remove managers from rewards manager', async () => {
    await rewardsManager.removeManager(await stakingManager.getAddress());
    const b = await rewardsManager.managers(await stakingManager.getAddress());
    assert.equal(b, 0n, 'Expected staking manager to be removed as manager');
  });

  it('only managers can call functions with the onlyManager constraint', async () => {
    await expect(
      rewardsManager.incrementRewardPool(owner, 10000),
    ).to.be.revertedWithCustomError(rewardsManager, 'OnlyManagers');
  });

  it('cannot increment reward pool with invalid arguments', async () => {
    await rewardsManager.addManager(owner);

    await expect(
      rewardsManager.incrementRewardPool(ethers.ZeroAddress, 10000),
    ).to.be.revertedWithCustomError(
      rewardsManager,
      'StakeeCannotBeZeroAddress',
    );

    await expect(
      rewardsManager.incrementRewardPool(owner, 0),
    ).to.be.revertedWithCustomError(rewardsManager, 'AmountCannotBeZero');
  });

  it('only managers can call functions with the onlyManager constraint', async () => {
    await expect(
      rewardsManager.initializeNextRewardPool(owner),
    ).to.be.revertedWithCustomError(rewardsManager, 'OnlyManagers');
  });

  it('cannot initialize next reward pool with invalid arguments', async () => {
    await rewardsManager.addManager(owner);

    await expect(
      rewardsManager.initializeNextRewardPool(ethers.ZeroAddress),
    ).to.be.revertedWithCustomError(
      rewardsManager,
      'StakeeCannotBeZeroAddress',
    );
  });

  it('should be able to deposit escrow', async () => {
    const alice = Wallet.createRandom();
    await syloTicketing.depositEscrow(50, alice.address);

    const deposit = await syloTicketing.deposits(alice.address);
    assert.equal(deposit.escrow.toString(), '50', 'Expected 50 in escrow');
  });

  it('should not be able to deposit escrow with invalid arguments', async () => {
    await expect(
      syloTicketing.depositEscrow(0, owner),
    ).to.be.revertedWithCustomError(syloTicketing, 'EscrowAmountCannotBeZero');

    await expect(
      syloTicketing.depositEscrow(50, ethers.ZeroAddress),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'AccountCannotBeZeroAddress',
    );
  });

  it('should be able to deposit penalty', async () => {
    const alice = Wallet.createRandom();
    await syloTicketing.depositPenalty(50, alice.address);

    const deposit = await syloTicketing.deposits(alice.address);
    assert.equal(deposit.penalty.toString(), '50', 'Expected 50 in escrow');
  });

  it('should not be able to deposit penalty with invalid arguments', async () => {
    await expect(
      syloTicketing.depositPenalty(0, owner),
    ).to.be.revertedWithCustomError(syloTicketing, 'PenaltyAmountCannotBeZero');

    await expect(
      syloTicketing.depositPenalty(50, ethers.ZeroAddress),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'AccountCannotBeZeroAddress',
    );
  });

  it('should be able to deposit escrow multiple times', async () => {
    const alice = Wallet.createRandom();
    await syloTicketing.depositEscrow(50, alice.address);
    await syloTicketing.depositEscrow(50, alice.address);

    const deposit = await syloTicketing.deposits(alice.address);
    assert.equal(deposit.escrow.toString(), '100', 'Expected 100 in escrow');
  });

  it('should be able to deposit to penalty multiple times', async () => {
    const alice = Wallet.createRandom();
    await syloTicketing.depositPenalty(50, alice.address);
    await syloTicketing.depositPenalty(50, alice.address);

    const deposit = await syloTicketing.deposits(alice.address);
    assert.equal(deposit.penalty.toString(), '100', 'Expected 100 in penalty');
  });

  it('should fail to withdraw without unlocking', async () => {
    await syloTicketing.depositEscrow(50, owner);
    await expect(syloTicketing.withdraw()).to.be.revertedWithCustomError(
      syloTicketing,
      'UnlockingNotInProcess',
    );
  });

  it('should fail to unlock without deposit', async () => {
    await expect(syloTicketing.unlockDeposits()).to.be.revertedWithCustomError(
      syloTicketing,
      'NoEsrowAndPenalty',
    );
  });

  it('should be able to unlock', async () => {
    await syloTicketing.depositEscrow(50, owner);
    await syloTicketing.unlockDeposits({ from: owner });

    const deposit = await syloTicketing.deposits(owner);
    expect(deposit.unlockAt).to.be.greaterThan(0n);
  });

  it('should fail to unlock if already unlocking', async () => {
    await syloTicketing.depositEscrow(50, owner);
    await syloTicketing.unlockDeposits();

    await expect(syloTicketing.unlockDeposits()).to.be.revertedWithCustomError(
      syloTicketing,
      'UnlockingInProcess',
    );
  });

  it('should fail to lock if already locked', async () => {
    await syloTicketing.depositEscrow(50, owner);
    await expect(syloTicketing.lockDeposits()).to.be.revertedWithCustomError(
      syloTicketing,
      'UnlockingNotInProcess',
    );
  });

  it('should be able to lock deposit while it is unlocked', async () => {
    await syloTicketing.depositEscrow(50, owner);
    await syloTicketing.unlockDeposits();

    await syloTicketing.lockDeposits();

    const deposit = await syloTicketing.deposits(owner);
    assert.equal(
      deposit.unlockAt.toString(),
      '0',
      'Expected deposit to move out of unlocking phase',
    );
  });

  it('should fail to deposit while unlocking', async () => {
    await syloTicketing.depositEscrow(50, owner);
    await syloTicketing.unlockDeposits();

    await expect(
      syloTicketing.depositEscrow(10, owner),
    ).to.be.revertedWithCustomError(syloTicketing, 'UnlockingInProcess');
    await expect(
      syloTicketing.depositPenalty(10, owner),
    ).to.be.revertedWithCustomError(syloTicketing, 'UnlockingInProcess');
  });

  it('should be able to withdraw after unlocking phase has completed', async () => {
    await syloTicketing.depositEscrow(50, owner);
    await syloTicketing.unlockDeposits();

    await utils.advanceBlock(11);

    const balanceBeforeWithdrawal = await token.balanceOf(owner);

    await syloTicketing.withdraw();

    const balanceAfterWithdrawal = await token.balanceOf(owner);

    expect(balanceAfterWithdrawal).to.be.equal(balanceBeforeWithdrawal + 50n);

    // can now deposit again
    await syloTicketing.depositEscrow(50, owner);
  });

  it('should fail to withdraw if deposits not unlocked', async () => {
    await syloTicketing.depositEscrow(50, owner);
    await expect(syloTicketing.withdraw()).to.be.revertedWithCustomError(
      syloTicketing,
      'UnlockingNotInProcess',
    );
  });

  it('should fail to withdraw if still unlocking', async () => {
    await syloTicketing.depositEscrow(50, owner);
    await syloTicketing.unlockDeposits();

    await expect(syloTicketing.withdraw()).to.be.revertedWithCustomError(
      syloTicketing,
      'UnlockingNotCompleted',
    );
  });

  it('should be able to initialize next reward pool', async () => {
    await stakingManager.addStake(30, owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    const currentBlock = await ethers.provider.getBlockNumber();

    await epochsManager.joinNextEpoch();

    const rewardPool = await rewardsManager.getRewardPool(
      await epochsManager.getNextEpochId(),
      owner,
    );

    expect(rewardPool.initializedAt).to.be.greaterThan(BigInt(currentBlock));

    assert.equal(
      rewardPool.totalActiveStake.toString(),
      '30',
      'Expected reward pool to correctly track the stake at the time it was created',
    );
  });

  it('can not initialize reward pool more than once', async () => {
    await stakingManager.addStake(30, owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);
    await epochsManager.joinNextEpoch();

    // change the seeker but node should still be prevented from
    // initializing the reward pool again
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 2);
    await expect(epochsManager.joinNextEpoch()).to.be.revertedWithCustomError(
      rewardsManager,
      'RewardPoolAlreadyExist',
    );
  });

  it('can not initialize reward pool more than once for the same seekers', async () => {
    await stakingManager.addStake(30, owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);
    await epochsManager.joinNextEpoch();

    await expect(epochsManager.joinNextEpoch())
      .to.be.revertedWithCustomError(epochsManager, 'SeekerAlreadyJoinedEpoch')
      .withArgs(1, 1);
  });

  it('should not be able to initialize next reward pool without stake', async () => {
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);
    await expect(epochsManager.joinNextEpoch()).to.be.revertedWithCustomError(
      rewardsManager,
      'NoStakeToCreateRewardPool',
    );
  });

  it('cannot redeem ticket with invalid signature', async () => {
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    const { ticket, redeemerRand } = await createWinningTicket(
      syloTicketing,
      epochsManager,
      alice,
      bob,
      owner,
    );

    const senderSig = '0x00';
    const receiverSig = '0x00';

    await expect(
      syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig),
    ).to.be.revertedWith('ECDSA: invalid signature length');
  });

  it('cannot redeem ticket with invalid redeemer rand', async () => {
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    const { ticket, senderSig, receiverSig } = await createWinningTicket(
      syloTicketing,
      epochsManager,
      alice,
      bob,
      owner,
    );

    const redeemerRand = 999;

    await expect(
      syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig),
    ).to.be.revertedWithCustomError(syloTicketing, 'RedeemerCommitMismatch');
  });

  it('cannot redeem ticket if associated epoch does not exist', async () => {
    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    const { ticket, redeemerRand, senderSig, receiverSig } =
      await createWinningTicket(
        syloTicketing,
        epochsManager,
        alice,
        bob,
        owner,
        1,
      );

    await expect(
      syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig),
    ).to.be.revertedWithCustomError(syloTicketing, 'TicketEpochNotFound');
  });

  it('can not calculate winning probability if associated epoch does not exist', async () => {
    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    const { ticket } = await createWinningTicket(
      syloTicketing,
      epochsManager,
      alice,
      bob,
      owner,
    );

    ticket.epochId = 1;

    await expect(
      syloTicketing.calculateWinningProbability(
        ticket.epochId,
        ticket.generationBlock,
      ),
    ).to.be.revertedWithCustomError(syloTicketing, 'TicketEpochNotFound');
  });

  it('cannot redeem ticket if generated for a future block', async () => {
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    const { ticket, redeemerRand, senderSig, receiverSig } =
      await createWinningTicket(
        syloTicketing,
        epochsManager,
        alice,
        bob,
        owner,
      );

    const updatedTicket = { ...ticket, generationBlock: 100000 };

    await expect(
      syloTicketing.redeem(updatedTicket, redeemerRand, senderSig, receiverSig),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'TicketCannotBeFromFutureBlock',
    );
  });

  it('cannot redeem ticket using sender delegated account to sign without permission', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    const delegatedWallet = Wallet.createRandom();

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, redeemerRand, senderSig, receiverSig } =
      await createWinningTicket(
        syloTicketing,
        epochsManager,
        alice,
        bob,
        owner,
        1,
        delegatedWallet,
        delegatedWallet,
      );

    await expect(
      syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'InvalidSenderTicketSigningPermission',
    );
  });

  it('cannot redeem ticket using receiver delegated account to sign without permission', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    const delegatedWallet = Wallet.createRandom();

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);
    await accounts[0].sendTransaction({
      to: alice.address,
      value: ethers.parseEther('2000.0'),
    });

    // alice adds this account as delegated account with permission to withdraw deposit
    const permission: Permission[] = [Permission.TicketSigning];
    const provider = ethers.provider;
    const aliceConnected = alice.connect(provider);
    await authorizedAccounts
      .connect(aliceConnected)
      .authorizeAccount(delegatedWallet.address, permission);

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);
    await accounts[0].sendTransaction({
      to: alice.address,
      value: ethers.parseEther('2000.0'),
    });

    const { ticket, redeemerRand, senderSig, receiverSig } =
      await createWinningTicket(
        syloTicketing,
        epochsManager,
        alice,
        bob,
        owner,
        1,
        delegatedWallet,
        delegatedWallet,
      );

    await expect(
      syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'InvalidReceiverTicketSigningPermission',
    );
  });

  it('cannot redeem ticket using sender delegated account to sign after unauthorizing account', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    const delegatedWallet = Wallet.createRandom();

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);
    await accounts[0].sendTransaction({
      to: alice.address,
      value: ethers.parseEther('2000.0'),
    });

    // alice adds this account as delegated account with permission to withdraw deposit
    const permission: Permission[] = [Permission.TicketSigning];
    const provider = ethers.provider;
    const aliceConnected = alice.connect(provider);
    await authorizedAccounts
      .connect(aliceConnected)
      .authorizeAccount(delegatedWallet.address, permission);
    await authorizedAccounts
      .connect(aliceConnected)
      .unauthorizeAccount(delegatedWallet.address);

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);
    await accounts[0].sendTransaction({
      to: alice.address,
      value: ethers.parseEther('2000.0'),
    });

    const { ticket, redeemerRand, senderSig, receiverSig } =
      await createWinningTicket(
        syloTicketing,
        epochsManager,
        alice,
        bob,
        owner,
        1,
        delegatedWallet,
      );

    await expect(
      syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'InvalidSenderTicketSigningPermission',
    );

    await authorizedAccounts
      .connect(aliceConnected)
      .authorizeAccount(delegatedWallet.address, permission);
    await authorizedAccounts
      .connect(aliceConnected)
      .removePermissions(delegatedWallet.address, [Permission.TicketSigning]);

    await expect(
      syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'InvalidSenderTicketSigningPermission',
    );
  });

  it('can redeem ticket using sender delegated account if removing permission is called after creating ticket', async () => {
    // make sure ticket wins as we call removePermissions after creating ticket
    await ticketingParameters.setDecayRate(1);

    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    const delegatedWallet = Wallet.createRandom();

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);
    await accounts[0].sendTransaction({
      to: alice.address,
      value: ethers.parseEther('2000.0'),
    });

    // alice adds this account as delegated account with permission to withdraw deposit
    const permission: Permission[] = [Permission.TicketSigning];
    const provider = ethers.provider;
    const aliceConnected = alice.connect(provider);
    await authorizedAccounts
      .connect(aliceConnected)
      .authorizeAccount(delegatedWallet.address, permission);

    const { ticket, redeemerRand, senderSig, receiverSig } =
      await createWinningTicket(
        syloTicketing,
        epochsManager,
        alice,
        bob,
        owner,
        1,
        delegatedWallet,
      );

    await authorizedAccounts
      .connect(aliceConnected)
      .removePermissions(delegatedWallet.address, permission);

    await syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig);

    await checkAfterRedeem(
      syloTicketing,
      rewardsManager,
      owner,
      alice,
      1000,
      50,
      500,
    );
  });

  it('can redeem ticket using sender delegated account if unauthorizing account is called after creating ticket', async () => {
    // make sure ticket wins as we call unauthorizeAccount after creating ticket
    await ticketingParameters.setDecayRate(1);

    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    const delegatedWallet = Wallet.createRandom();

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);
    await accounts[0].sendTransaction({
      to: alice.address,
      value: ethers.parseEther('2000.0'),
    });

    // alice adds this account as delegated account with permission to withdraw deposit
    const permission: Permission[] = [Permission.TicketSigning];
    const provider = ethers.provider;
    const aliceConnected = alice.connect(provider);
    await authorizedAccounts
      .connect(aliceConnected)
      .authorizeAccount(delegatedWallet.address, permission);

    const { ticket, redeemerRand, senderSig, receiverSig } =
      await createWinningTicket(
        syloTicketing,
        epochsManager,
        alice,
        bob,
        owner,
        1,
        delegatedWallet,
      );

    await authorizedAccounts
      .connect(aliceConnected)
      .unauthorizeAccount(delegatedWallet.address);

    await syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig);

    await checkAfterRedeem(
      syloTicketing,
      rewardsManager,
      owner,
      alice,
      1000,
      50,
      500,
    );
  });

  it('can redeem ticket using sender authorized account to sign with valid permission', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    const delegatedWallet = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);
    await accounts[0].sendTransaction({
      to: alice.address,
      value: ethers.parseEther('2000.0'),
    });

    // alice adds this account as delegated account with permission to withdraw deposit
    const permission: Permission[] = [Permission.TicketSigning];
    const provider = ethers.provider;
    const aliceConnected = alice.connect(provider);
    await authorizedAccounts
      .connect(aliceConnected)
      .authorizeAccount(delegatedWallet.address, permission);

    const { ticket, redeemerRand, senderSig, receiverSig } =
      await createWinningTicket(
        syloTicketing,
        epochsManager,
        alice,
        bob,
        owner,
        undefined,
        delegatedWallet,
      );

    await syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig);

    await checkAfterRedeem(
      syloTicketing,
      rewardsManager,
      owner,
      alice,
      1000,
      50,
      500,
    );
  });

  it('can redeem ticket using sender and receiver authorized accounts to sign with valid permission', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();

    const aliceDelegatedWallet = Wallet.createRandom();
    const bobDelegatedWallet = Wallet.createRandom();

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    await accounts[0].sendTransaction({
      to: alice.address,
      value: ethers.parseEther('2000.0'),
    });
    await accounts[0].sendTransaction({
      to: bob.address,
      value: ethers.parseEther('2000.0'),
    });

    // alice adds this account as delegated account with permission to withdraw deposit
    const permission: Permission[] = [Permission.TicketSigning];
    await authorizedAccounts
      .connect(alice.connect(ethers.provider))
      .authorizeAccount(aliceDelegatedWallet.address, permission);

    // bob adds this account as delegated account with permission to withdraw deposit
    await authorizedAccounts
      .connect(bob.connect(ethers.provider))
      .authorizeAccount(bobDelegatedWallet.address, permission);

    const { ticket, redeemerRand, senderSig, receiverSig } =
      await createWinningTicket(
        syloTicketing,
        epochsManager,
        alice,
        bob,
        owner,
        undefined,
        aliceDelegatedWallet,
        bobDelegatedWallet,
      );

    await syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig);

    await checkAfterRedeem(
      syloTicketing,
      rewardsManager,
      owner,
      alice,
      1000,
      50,
      500,
    );
  });

  it('can not calculate winning probablility if not generated during associated epoch', async () => {
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    const { ticket } = await createWinningTicket(
      syloTicketing,
      epochsManager,
      alice,
      bob,
      owner,
    );

    const updatedTicket = { ...ticket, generationBlock: 1 };

    await expect(
      syloTicketing.calculateWinningProbability(
        updatedTicket.epochId,
        updatedTicket.generationBlock,
      ),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'TicketNotCreatedInTheEpoch',
    );
  });

  it('cannot calculate winning probability if generated after epoch end block', async () => {
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    const { ticket } = await createWinningTicket(
      syloTicketing,
      epochsManager,
      alice,
      bob,
      owner,
    );

    await epochsManager.initializeEpoch();

    const updatedTicket = {
      ...ticket,
      generationBlock: 10000000,
    };

    await expect(
      syloTicketing.calculateWinningProbability(
        updatedTicket.epochId,
        updatedTicket.generationBlock,
      ),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'TicketNotCreatedInTheEpoch',
    );
  });

  it('calculate winning ticket panic for invalid ticket', async () => {
    await epochsManager.initializeEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    const { ticket } = await createWinningTicket(
      syloTicketing,
      epochsManager,
      alice,
      bob,
      owner,
    );

    const updatedTicket = { ...ticket, generationBlock: 10000 };

    await expect(
      syloTicketing.calculateWinningProbability(
        updatedTicket.epochId,
        updatedTicket.generationBlock,
      ),
    ).to.be.revertedWithPanic(0x11); // Overflow: Arithmetic operation underflowed or overflowed
  });

  it('cannot redeem ticket if node has not joined directory', async () => {
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, redeemerRand, senderSig, receiverSig } =
      await createWinningTicket(
        syloTicketing,
        epochsManager,
        alice,
        bob,
        owner,
      );

    await expect(
      syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'RedeemerMustHaveJoinedEpoch',
    );
  });

  it('cannot redeem ticket if node has not initialized reward pool', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await directory.addManager(owner);
    await directory.joinNextDirectory(owner);

    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, redeemerRand, senderSig, receiverSig } =
      await createWinningTicket(
        syloTicketing,
        epochsManager,
        alice,
        bob,
        owner,
      );

    await expect(
      syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig),
    ).to.be.revertedWithCustomError(rewardsManager, 'RewardPoolNotExist');
  });

  it('cannot redeem invalid ticket', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, redeemerRand, senderSig, receiverSig } =
      await createWinningTicket(
        syloTicketing,
        epochsManager,
        alice,
        bob,
        owner,
      );

    await expect(
      syloTicketing.redeem(
        {
          ...ticket,
          sender: { ...ticket.sender, main: ethers.ZeroAddress },
        },

        redeemerRand,
        senderSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'TicketSenderCannotBeZeroAddress',
    );

    await expect(
      syloTicketing.redeem(
        {
          ...ticket,
          receiver: { ...ticket.receiver, main: ethers.ZeroAddress },
        },

        redeemerRand,
        senderSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'TicketReceiverCannotBeZeroAddress',
    );

    await expect(
      syloTicketing.redeem(
        {
          ...ticket,
          redeemer: ethers.ZeroAddress,
        },

        redeemerRand,
        senderSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'TicketRedeemerCannotBeZeroAddress',
    );

    await expect(
      syloTicketing.redeem(
        {
          ...ticket,
          redeemerCommit:
            '0x0000000000000000000000000000000000000000000000000000000000000000',
        },

        redeemerRand,
        senderSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(syloTicketing, 'RedeemerCommitMismatch');

    const malformedSig =
      '0xdebcaaaa727df04bdc990083d88ed7c8e6e9897ff18b7d968867a8bc024cbdbe10ca52eebd67a14b7b493f5c00ed9dab7b96ef62916f25afc631d336f7b2ae1e1b';
    await expect(
      syloTicketing.redeem(ticket, redeemerRand, malformedSig, receiverSig),
    ).to.be.revertedWithCustomError(syloTicketing, 'InvalidSenderSignature');

    await expect(
      syloTicketing.redeem(ticket, redeemerRand, senderSig, malformedSig),
    ).to.be.revertedWithCustomError(syloTicketing, 'InvalidReceiverSignature');
  });

  it('rejects non winning ticket', async () => {
    // redeploy contracts with win chance of 0%
    const contracts = await utils.initializeContracts(owner, token, {
      baseLiveWinProb: 0,
    });
    await token.approve(
      await contracts.stakingManager.getAddress(),
      toSOLOs(100000),
    );
    await contracts.stakingManager.addStake(toSOLOs(1), owner);
    await utils.setSeekerRegistry(
      contracts.registries,
      contracts.seekers,
      accounts[0],
      accounts[1],
      1,
    );

    await contracts.epochsManager.joinNextEpoch();

    await contracts.directory.transferOwnership(
      await contracts.epochsManager.getAddress(),
    );
    await contracts.epochsManager.initializeEpoch();

    await token.approve(
      await contracts.syloTicketing.getAddress(),
      toSOLOs(100000),
    );
    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await contracts.syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await contracts.syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, redeemerRand, senderSig, receiverSig } =
      await createWinningTicket(
        syloTicketing,
        epochsManager,
        alice,
        bob,
        owner,
        1,
      );

    await utils.advanceBlock(5);

    await expect(
      contracts.syloTicketing.redeem(
        ticket,
        redeemerRand,
        senderSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(syloTicketing, 'TicketNotWinning');
  });

  it('can redeem winning ticket', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, redeemerRand, senderSig, receiverSig } =
      await createWinningTicket(
        syloTicketing,
        epochsManager,
        alice,
        bob,
        owner,
      );

    await syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig);

    const deposit = await syloTicketing.deposits(alice.address);
    assert.equal(
      deposit.escrow,
      toSOLOs(1000),
      'Expected ticket payout to be substracted from escrow',
    );
    assert.equal(
      deposit.penalty,
      toSOLOs(50),
      'Expected penalty to not be changed',
    );

    const pendingReward = await rewardsManager.getPendingRewards(owner);

    assert.equal(
      pendingReward,
      toSOLOs(500),
      'Expected balance of pending rewards to have added the ticket face value',
    );
  });

  it('cannot redeem ticket more than once', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, redeemerRand, senderSig, receiverSig } =
      await createWinningTicket(
        syloTicketing,
        epochsManager,
        alice,
        bob,
        owner,
      );

    await syloTicketing.redeem(
      ticket,

      redeemerRand,
      senderSig,
      receiverSig,
    );
    await expect(
      syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig),
    ).to.be.revertedWithCustomError(syloTicketing, 'TicketAlreadyRedeemed');
  });

  it('not manager cannot update pending rewards', async () => {
    await expect(
      rewardsManager.connect(accounts[1]).updatePendingRewards(owner, owner),
    ).to.be.revertedWithCustomError(rewardsManager, 'OnlyManagers');
  });

  it('burns penalty on insufficient escrow', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(5), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, redeemerRand, senderSig, receiverSig } =
      await createWinningTicket(
        syloTicketing,
        epochsManager,
        alice,
        bob,
        owner,
      );

    const initialTicketingBalance = await token.balanceOf(
      await syloTicketing.getAddress(),
    );

    await expect(
      syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig),
    )
      .to.emit(syloTicketing, 'SenderPenaltyBurnt')
      .withArgs(alice.address);

    const deposit = await syloTicketing.deposits(alice.address);
    assert.equal(
      deposit.escrow.toString(),
      '0',
      'Expected entire escrow to be used',
    );
    assert.equal(
      deposit.penalty.toString(),
      '0',
      'Expected entire penalty to be burned',
    );

    const pendingReward = await rewardsManager.getPendingRewards(owner);

    assert.equal(
      pendingReward.toString(),
      '2500000000000000000',
      'Expected unclaimed balance to have added the remaining available escrow',
    );

    const ticketingBalance = await token.balanceOf(
      await syloTicketing.getAddress(),
    );
    assert.equal(
      ticketingBalance,
      initialTicketingBalance - toSOLOs(55),
      'Expected tokens from ticket contract to be removed',
    );

    const deadBalance = await token.balanceOf(
      '0x000000000000000000000000000000000000dEaD',
    );
    assert.equal(
      deadBalance.toString(),
      '50000000000000000000',
      'Expected dead address to receive burned tokens',
    );
  });

  it('fails to to claim non existent rewards', async () => {
    await expect(
      rewardsManager.claimStakingRewards(owner),
    ).to.be.revertedWithCustomError(rewardsManager, 'NoRewardToClaim');
  });

  it('can claim ticketing rewards', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(50000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    for (let i = 0; i < 10; i++) {
      const { ticket, redeemerRand, senderSig, receiverSig } =
        await createWinningTicket(
          syloTicketing,
          epochsManager,
          alice,
          bob,
          owner,
        );

      await syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig);
    }

    await epochsManager.initializeEpoch();

    const initialBalance = await token.balanceOf(owner);

    await rewardsManager.claimStakingRewards(owner);

    const postBalance = await token.balanceOf(owner);
    // Expect the node have the entire reward balance added to their account
    const expectedPostBalance = initialBalance + toSOLOs(10000);

    compareExpectedBalance(expectedPostBalance, postBalance);

    const pendingReward = await rewardsManager.getPendingRewards(owner);

    compareExpectedBalance(pendingReward, 0);

    // check total rewards in the previous epoch after claiming
    const nextEpochId = await epochsManager.getNextEpochId();
    const rewardPoolStakersTotal =
      await rewardsManager.getRewardPoolStakersTotal(nextEpochId - 2n, owner);

    assert.equal(
      rewardPoolStakersTotal,
      toSOLOs(500 * 10), // 500 is added to the stakers reward total on each redemption (50% of 1000)
      'Expected reward pool stakers total in the previous epoch to be 5000 SOLOs',
    );
  });

  it('delegated stakers should be able to claim rewards', async () => {
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    const { proportions } = await addStakes(token, stakingManager, owner, [
      { account: accounts[0], stake: 3 },
      // have account 2 as a delegated staker
      { account: accounts[2], stake: 2 },
    ]);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(50000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    for (let i = 0; i < 10; i++) {
      const { ticket, redeemerRand, senderSig, receiverSig } =
        await createWinningTicket(
          syloTicketing,
          epochsManager,
          alice,
          bob,
          owner,
        );

      await syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig);
    }

    const totalWinnings = 5000;

    await epochsManager.initializeEpoch();

    await testClaims(token, rewardsManager, owner, [
      {
        account: accounts[0],
        claim: proportions[0] * totalWinnings + 5000, // add the node's fee,
      },
      {
        account: accounts[2],
        claim: proportions[1] * totalWinnings,
      },
    ]);
  });

  it('should have rewards be automatically removed from pending when stake is updated', async () => {
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await addStakes(token, stakingManager, owner, [
      { account: accounts[0], stake: 10000 },
      // have account 2 as a delegated staker
      { account: accounts[2], stake: 2 },
    ]);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(50000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    for (let i = 0; i < 10; i++) {
      const { ticket, redeemerRand, senderSig, receiverSig } =
        await createWinningTicket(
          syloTicketing,
          epochsManager,
          alice,
          bob,
          owner,
        );

      await syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig);
    }

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const claimBeforeAddingStake = await rewardsManager.calculateStakerClaim(
      owner,
      await accounts[2].getAddress(),
    );
    const pendingRewardBeforeAddingStake =
      await rewardsManager.getPendingRewards(owner);

    // add more stake
    await addStakes(token, stakingManager, owner, [
      { account: accounts[2], stake: 1 },
    ]);

    // pending reward after stake is added should have previous claim removed
    const pendingRewardAfterAddingStake =
      await rewardsManager.getPendingRewards(owner);

    assert.equal(
      pendingRewardBeforeAddingStake - pendingRewardAfterAddingStake,
      claimBeforeAddingStake,
      'Expected claim to be removed from pending reward after adding stake',
    );

    await rewardsManager.connect(accounts[2]).claimStakingRewards(owner);

    for (let i = 0; i < 10; i++) {
      const { ticket, redeemerRand, senderSig, receiverSig } =
        await createWinningTicket(
          syloTicketing,
          epochsManager,
          alice,
          bob,
          owner,
        );

      await syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig);
    }

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const claimBeforeRemovingStake = await rewardsManager.calculateStakerClaim(
      owner,
      await accounts[2].getAddress(),
    );

    const pendingRewardBeforeRemovingStake =
      await rewardsManager.getPendingRewards(owner);

    // remove some stake
    await stakingManager.connect(accounts[2]).unlockStake(1, owner);

    const pendingRewardAfterRemovingStake =
      await rewardsManager.getPendingRewards(owner);

    assert.equal(
      pendingRewardBeforeRemovingStake - pendingRewardAfterRemovingStake,
      claimBeforeRemovingStake,
      'Expected claim to be removed from pending reward after removing stake',
    );
  });

  it('can calculate staker claim if reward total is 0', async () => {
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await addStakes(token, stakingManager, owner, [
      { account: accounts[0], stake: 3 },
      // have account 2 as a delegated staker
      { account: accounts[2], stake: 10 },
    ]);

    async function calculateClaims() {
      expect(await rewardsManager.calculateStakerClaim(owner, owner)).to.equal(
        0,
      );
      expect(
        await rewardsManager.calculateStakerClaim(
          owner,
          await accounts[2].getAddress(),
        ),
      ).to.equal(0);
    }

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    await calculateClaims();

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    await calculateClaims();
  });

  it('cannot calculate staker claim with invalid arguments)', async () => {
    await expect(
      rewardsManager.calculateStakerClaim(ethers.ZeroAddress, owner),
    ).to.be.revertedWithCustomError(
      rewardsManager,
      'StakeeCannotBeZeroAddress',
    );

    await expect(
      rewardsManager.calculateStakerClaim(owner, ethers.ZeroAddress),
    ).to.be.revertedWithCustomError(
      rewardsManager,
      'StakerCannotBeZeroAddress',
    );
  });

  it('can not claim reward more than once for the same epoch', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(50000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, redeemerRand, senderSig, receiverSig } =
      await createWinningTicket(
        syloTicketing,
        epochsManager,
        alice,
        bob,
        owner,
      );

    await syloTicketing.redeem(
      ticket,

      redeemerRand,
      senderSig,
      receiverSig,
    );

    await epochsManager.initializeEpoch();

    await rewardsManager.claimStakingRewards(owner);

    const lastClaim = await rewardsManager.getLastClaim(owner, owner);
    expect(lastClaim.claimedAt).to.be.above(0);

    await expect(
      rewardsManager.claimStakingRewards(owner),
    ).to.be.revertedWithCustomError(rewardsManager, 'NoRewardToClaim');
  });

  it('should be able to correctly calculate staking rewards for multiple epochs when managed stake is the same', async () => {
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    const { proportions } = await addStakes(token, stakingManager, owner, [
      { account: accounts[0], stake: 1000 },
      // have accounts 1, 2 and 3 as delegated stakers
      { account: accounts[1], stake: 250 },
      { account: accounts[2], stake: 400 },
      { account: accounts[3], stake: 350 },
    ]);

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(500000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    for (let j = 0; j < 3; j++) {
      await epochsManager.joinNextEpoch();
      await epochsManager.initializeEpoch();

      // 500 is added to the stakers reward total on each redemption (50% of 1000)
      for (let i = 0; i < 6; i++) {
        const { ticket, redeemerRand, senderSig, receiverSig } =
          await createWinningTicket(
            syloTicketing,
            epochsManager,
            alice,
            bob,
            owner,
          );

        await syloTicketing.redeem(
          ticket,
          redeemerRand,
          senderSig,
          receiverSig,
        );
      }
    }

    await epochsManager.initializeEpoch();

    const pendingReward = await rewardsManager.getPendingRewards(owner);

    // the total unclaimed stake reward should 3 * 6 * 500 = 9000
    const totalWinnings = 9000;

    compareExpectedBalance(toSOLOs(totalWinnings), pendingReward);

    // verify each staker will receive the correct amount of reward if they were to claim now
    await testClaims(token, rewardsManager, owner, [
      { account: accounts[1], claim: proportions[1] * totalWinnings },
      { account: accounts[2], claim: proportions[2] * totalWinnings },
      { account: accounts[3], claim: proportions[3] * totalWinnings },
    ]);
  });

  it('should be able to stake, accumulate rewards, and claim more than once as a delegated staker', async () => {
    const { proportions } = await addStakes(token, stakingManager, owner, [
      { account: accounts[0], stake: 1000 },
      // have accounts 1, 2 and 3 as delegated stakers
      { account: accounts[1], stake: 250 },
      { account: accounts[2], stake: 400 },
      { account: accounts[3], stake: 350 },
    ]);

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(500000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);
    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    for (let runs = 0; runs < 10; runs++) {
      // 250 is added to the delegated stakers reward total on each redemption
      for (let i = 0; i < 5; i++) {
        const { ticket, redeemerRand, senderSig, receiverSig } =
          await createWinningTicket(
            syloTicketing,
            epochsManager,
            alice,
            bob,
            owner,
          );

        await syloTicketing.redeem(
          ticket,
          redeemerRand,
          senderSig,
          receiverSig,
        );
      }

      await epochsManager.joinNextEpoch();
      await epochsManager.initializeEpoch();

      const totalWinnings = 2500;

      // No accounts adds or withdraws stake, so stake proportions remain constant
      await testClaims(token, rewardsManager, owner, [
        {
          account: accounts[0],
          claim: proportions[0] * totalWinnings + 2500,
        },
        {
          account: accounts[1],
          claim: proportions[1] * totalWinnings,
        },
        {
          account: accounts[2],
          claim: proportions[2] * totalWinnings,
        },
        {
          account: accounts[3],
          claim: proportions[3] * totalWinnings,
        },
      ]);
    }
  });

  it('should be able to correctly calculate staking rewards for multiple epochs when managed stake increases', async () => {
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await addStakes(token, stakingManager, owner, [
      { account: accounts[0], stake: 1000 },
      // have accounts 1, 2 and 3 as delegated stakers
      { account: accounts[1], stake: 250 },
      { account: accounts[2], stake: 400 },
      { account: accounts[3], stake: 350 },
    ]);

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(50000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    // have account 4 add stake midway through
    for (let j = 0; j < 3; j++) {
      if (j == 1) {
        await addStakes(token, stakingManager, owner, [
          { account: accounts[4], stake: 500 },
        ]);
      }

      await epochsManager.joinNextEpoch();
      await epochsManager.initializeEpoch();

      // 500 is added to the stakers reward total on each redemption (50% of 1000)
      for (let i = 0; i < 6; i++) {
        const { ticket, redeemerRand, senderSig, receiverSig } =
          await createWinningTicket(
            syloTicketing,
            epochsManager,
            alice,
            bob,
            owner,
          );

        await syloTicketing.redeem(
          ticket,
          redeemerRand,
          senderSig,
          receiverSig,
        );
      }
    }

    await epochsManager.initializeEpoch();

    const epochTwoActiveStake = await rewardsManager.getRewardPoolActiveStake(
      2,
      owner,
    );

    // account 4's reward should be the sum of the rewards gained in both epoch 2 and 3
    // multiplied by the proportion of the stake held when their stake became active
    const stakeClaimFive = await rewardsManager.calculateStakerClaim(
      owner,
      await accounts[4].getAddress(),
    );
    const s = toSOLOs(500); // initial stake
    const r = toSOLOs(2 * 6 * 500); // accumulated reward
    const expectedStakeClaimFive = (s * BigInt(r)) / epochTwoActiveStake;
    compareExpectedBalance(expectedStakeClaimFive, stakeClaimFive);

    // for accounts 1, 2, and 3, the total managed stake that becomes active
    // changes from epoch 2, thus to calculate the expected reward, we need
    // to calculate the expected reward for epoch 1 using different stake proportions
    // than for epochs 2 and 3
    for (let i = 1; i < 4; i++) {
      const initialStake = await stakingManager.getCurrentStakerAmount(
        owner,
        await accounts[i].getAddress(),
      );
      const epochOneActiveStake = await rewardsManager.getRewardPoolActiveStake(
        1,
        owner,
      );
      const epochOneReward =
        (toSOLOs(6 * 500) * initialStake) / epochOneActiveStake;

      const remainingReward =
        (toSOLOs(2 * 6 * 500) * initialStake) / epochTwoActiveStake;

      const totalExpectedReward = epochOneReward + remainingReward;
      const stakerClaim = await rewardsManager.calculateStakerClaim(
        owner,
        await accounts[i].getAddress(),
      );

      compareExpectedBalance(totalExpectedReward, stakerClaim);
    }

    // ensure each staker is actually able to claim
    for (let i = 1; i < 5; i++) {
      await rewardsManager.connect(accounts[i]).claimStakingRewards(owner);
    }
  });

  it('should be able to correctly calculate staking rewards for multiple epochs when managed stake decreases', async () => {
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await addStakes(token, stakingManager, owner, [
      { account: accounts[0], stake: 1000 },
      // have accounts 1, 2 and 3 as delegated stakers
      { account: accounts[1], stake: 250 },
      { account: accounts[2], stake: 400 },
      { account: accounts[3], stake: 350 },
    ]);

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(50000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    // have account 1 unlock stake midway through
    for (let j = 0; j < 3; j++) {
      if (j == 1) {
        await stakingManager
          .connect(accounts[1])
          .unlockStake(toSOLOs(250), owner);
      }

      await epochsManager.joinNextEpoch();
      await epochsManager.initializeEpoch();

      for (let i = 0; i < 6; i++) {
        const { ticket, redeemerRand, senderSig, receiverSig } =
          await createWinningTicket(
            syloTicketing,
            epochsManager,
            alice,
            bob,
            owner,
          );

        await syloTicketing.redeem(
          ticket,
          redeemerRand,
          senderSig,
          receiverSig,
        );
      }
    }

    await epochsManager.initializeEpoch();

    const epochTwoActiveStake = await rewardsManager.getRewardPoolActiveStake(
      2,
      owner,
    );

    // for accounts 2, and 3, the total managed stake that becomes active
    // changes from epoch 2, thus to calculate the expected reward, we need
    // to caluclate the expected reward for epoch 1 using different stake proportions
    // than for epochs 2 and 3
    for (let i = 2; i < 4; i++) {
      const initialStake = await stakingManager.getCurrentStakerAmount(
        owner,
        await accounts[i].getAddress(),
      );
      const epochOneActiveStake = await rewardsManager.getRewardPoolActiveStake(
        1,
        owner,
      );
      const epochOneReward =
        (toSOLOs(6 * 500) * initialStake) / epochOneActiveStake;

      const remainingReward =
        (toSOLOs(2 * 6 * 500) * initialStake) / epochTwoActiveStake;

      const totalExpectedReward = epochOneReward + remainingReward;
      const stakerClaim = await rewardsManager.calculateStakerClaim(
        owner,
        await accounts[i].getAddress(),
      );

      compareExpectedBalance(totalExpectedReward, stakerClaim);
    }

    // ensure each staker is actually able to claim
    for (let i = 2; i < 4; i++) {
      await rewardsManager.connect(accounts[i]).claimStakingRewards(owner);
    }
  });

  // XXX: Set to skip as it is a very long test and sometimes breaks the local
  // truffle test network/client. However this should be manually run if any significant changes
  // to the Rewards contract calculation is made.
  // TODO: Create script to spin up new test network to run this test locally or for CI automatically.
  it('should calculate updated stake and rewards over several ticket redemptions without significant precision loss [ @skip-on-coverage ]', async () => {
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    const { proportions } = await addStakes(token, stakingManager, owner, [
      { account: accounts[0], stake: 1000 },
      // have accounts 1, 2 and 3 as delegated stakers
      { account: accounts[1], stake: 250 },
      { account: accounts[2], stake: 400 },
      { account: accounts[3], stake: 350 },
    ]);

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(1000 * 500), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const iterations = 450;

    for (let i = 0; i < iterations; i++) {
      const { ticket, redeemerRand, senderSig, receiverSig } =
        await createWinningTicket(
          syloTicketing,
          epochsManager,
          alice,
          bob,
          owner,
        );

      await syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig);
    }

    await epochsManager.initializeEpoch();

    await testClaims(token, rewardsManager, owner, [
      { account: accounts[1], claim: iterations * 500 * proportions[1] },
      { account: accounts[2], claim: iterations * 500 * proportions[2] },
      { account: accounts[3], claim: iterations * 500 * proportions[3] },
    ]);
  }).timeout(0);

  it('should decay winning probability as ticket approaches expiry', async () => {
    // deploy another ticketing contract with simpler parameters
    const contracts = await utils.initializeContracts(owner, token, {
      faceValue,
      baseLiveWinProb: 100000,
      expiredWinProb: 1000,
      decayRate: 8000,
      ticketDuration: 100,
    });
    epochsManager = contracts.epochsManager;
    rewardsManager = contracts.rewardsManager;
    syloTicketing = contracts.syloTicketing;
    directory = contracts.directory;
    registries = contracts.registries;
    stakingManager = contracts.stakingManager;

    await directory.transferOwnership(await epochsManager.getAddress());
    await rewardsManager.addManager(await syloTicketing.getAddress());
    await rewardsManager.addManager(await epochsManager.getAddress());

    await token.approve(await syloTicketing.getAddress(), toSOLOs(10000));
    await token.approve(await stakingManager.getAddress(), toSOLOs(10000));

    await stakingManager.addStake(toSOLOs(1), owner);
    await utils.setSeekerRegistry(
      contracts.registries,
      contracts.seekers,
      accounts[0],
      accounts[1],
      1,
    );

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(5000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket } = await createWinningTicket(
      syloTicketing,
      epochsManager,
      alice,
      bob,
      owner,
    );

    // advance the block halfway to ticket expiry
    await utils.advanceBlock(51);

    // check if the probability has decayed 50% of the maximum decayed value (80%)
    const expectedProbability = 100000 - 0.5 * 0.8 * 100000;

    const decayedProbability = await syloTicketing.calculateWinningProbability(
      ticket.epochId,
      ticket.generationBlock,
    );

    assert.equal(
      decayedProbability.toString(),
      expectedProbability.toString(),
      'Expected probablity of ticket winning to decay',
    );
  });

  it('should be able to correctly calculate staker rewards if node was not active for multiple epochs', async () => {
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    const { proportions } = await addStakes(token, stakingManager, owner, [
      { account: accounts[0], stake: 1000 },
      // have accounts 1, 2 and 3 as delegated stakers
      { account: accounts[1], stake: 250 },
      { account: accounts[2], stake: 400 },
      { account: accounts[3], stake: 350 },
    ]);

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(500000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    // the node doesn't participate for several epochs
    for (let i = 0; i < 3; i++) {
      await utils.advanceBlock(11);
      await epochsManager.initializeEpoch();
    }

    for (let j = 0; j < 3; j++) {
      await epochsManager.joinNextEpoch();
      await epochsManager.initializeEpoch();

      // 500 is added to the stakers reward total on each redemption (50% of 1000)
      for (let i = 0; i < 6; i++) {
        const { ticket, redeemerRand, senderSig, receiverSig } =
          await createWinningTicket(
            syloTicketing,
            epochsManager,
            alice,
            bob,
            owner,
          );

        await syloTicketing.redeem(
          ticket,
          redeemerRand,
          senderSig,
          receiverSig,
        );
      }
    }

    await epochsManager.initializeEpoch();

    const pendingReward = await rewardsManager.getPendingRewards(owner);

    // the total unclaimed stake reward should 3 * 6 * 500 = 9000
    const totalWinnings = 9000;

    compareExpectedBalance(toSOLOs(totalWinnings), pendingReward);

    // verify each staker will receive the correct amount of reward if they were to claim now
    await testClaims(token, rewardsManager, owner, [
      { account: accounts[1], claim: proportions[1] * totalWinnings },
      { account: accounts[2], claim: proportions[2] * totalWinnings },
      { account: accounts[3], claim: proportions[3] * totalWinnings },
    ]);
  });

  it('claiming staking rewards only claims up to the previous epoch', async () => {
    await addStakes(token, stakingManager, owner, [
      { account: accounts[0], stake: 1 },
      { account: accounts[1], stake: 1 },
    ]);

    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(50000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    for (let j = 0; j < 3; j++) {
      for (let i = 0; i < 10; i++) {
        const { ticket, redeemerRand, senderSig, receiverSig } =
          await createWinningTicket(
            syloTicketing,
            epochsManager,
            alice,
            bob,
            owner,
          );

        await syloTicketing.redeem(
          ticket,
          redeemerRand,
          senderSig,
          receiverSig,
        );
      }

      if (j != 2) {
        await epochsManager.joinNextEpoch();
        await epochsManager.initializeEpoch();
      }
    }

    // should only be able to claim for the first two epochs
    let claim = await rewardsManager.calculateStakerClaim(
      owner,
      await accounts[1].getAddress(),
    );
    compareExpectedBalance(claim, toSOLOs(5000));

    // initialize epoch and check all three epochs can be claimed
    await epochsManager.initializeEpoch();

    claim = await rewardsManager.calculateStakerClaim(
      owner,
      await accounts[1].getAddress(),
    );
    compareExpectedBalance(claim, toSOLOs(7500));
  });

  it('rewards generated in the current epoch are not claimable', async () => {
    await addStakes(token, stakingManager, owner, [
      { account: accounts[0], stake: 1 },
      { account: accounts[1], stake: 1 },
    ]);

    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(50000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    for (let j = 0; j < 2; j++) {
      for (let i = 0; i < 10; i++) {
        const { ticket, redeemerRand, senderSig, receiverSig } =
          await createWinningTicket(
            syloTicketing,
            epochsManager,
            alice,
            bob,
            owner,
          );

        await syloTicketing.redeem(
          ticket,
          redeemerRand,
          senderSig,
          receiverSig,
        );
      }

      await epochsManager.joinNextEpoch();
      await epochsManager.initializeEpoch();
    }

    // claim rewards from the previous epoch
    await rewardsManager.connect(accounts[1]).claimStakingRewards(owner);

    let claim = await rewardsManager.calculateStakerClaim(
      owner,
      await accounts[1].getAddress(),
    );

    compareExpectedBalance(claim, 0);

    // generate rewards for the current epoch
    for (let i = 0; i < 10; i++) {
      const { ticket, redeemerRand, senderSig, receiverSig } =
        await createWinningTicket(
          syloTicketing,
          epochsManager,
          alice,
          bob,
          owner,
        );

      await syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig);
    }

    // confirm claim is still 0
    claim = await rewardsManager.calculateStakerClaim(
      owner,
      await accounts[1].getAddress(),
    );

    compareExpectedBalance(claim, 0);

    await expect(
      rewardsManager.connect(accounts[1]).claimStakingRewards(owner),
    ).to.be.revertedWithCustomError(rewardsManager, 'NoRewardToClaim');
  });

  it('can claim staking rewards again after previous ended', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(50000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const initialBalance = await token.balanceOf(owner);

    for (let i = 0; i < 10; i++) {
      const { ticket, redeemerRand, senderSig, receiverSig } =
        await createWinningTicket(
          syloTicketing,
          epochsManager,
          alice,
          bob,
          owner,
        );

      await syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig);
    }

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    await rewardsManager.claimStakingRewards(owner);

    for (let i = 0; i < 10; i++) {
      const { ticket, redeemerRand, senderSig, receiverSig } =
        await createWinningTicket(
          syloTicketing,
          epochsManager,
          alice,
          bob,
          owner,
        );

      await syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig);
    }

    await epochsManager.initializeEpoch();

    await rewardsManager.claimStakingRewards(owner);

    const postBalance = await token.balanceOf(owner);

    // expect to receive total balance of all redeemed tickets
    compareExpectedBalance(postBalance, initialBalance + toSOLOs(20000));
  });

  it('can claim staking rewards if node already joined next epoch', async () => {
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    const { proportions } = await addStakes(token, stakingManager, owner, [
      { account: accounts[0], stake: 1000 },
      // have accounts 1, 2 and 3 as delegated stakers
      { account: accounts[1], stake: 250 },
      { account: accounts[2], stake: 400 },
      { account: accounts[3], stake: 350 },
    ]);

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(50000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    for (let j = 0; j < 2; j++) {
      await epochsManager.joinNextEpoch();
      await epochsManager.initializeEpoch();

      for (let i = 0; i < 10; i++) {
        const { ticket, redeemerRand, senderSig, receiverSig } =
          await createWinningTicket(
            syloTicketing,
            epochsManager,
            alice,
            bob,
            owner,
          );

        await syloTicketing.redeem(
          ticket,
          redeemerRand,
          senderSig,
          receiverSig,
        );
      }
    }

    // 10 tickets redeemed per epoch
    const totalWinnings = 5000;

    // have the node join the next epoch and test stakers are able
    // to claim for the first epoch
    await epochsManager.joinNextEpoch();
    await testClaims(token, rewardsManager, owner, [
      {
        account: accounts[1],
        claim: proportions[1] * totalWinnings,
      },
      {
        account: accounts[2],
        claim: proportions[2] * totalWinnings,
      },
      {
        account: accounts[3],
        claim: proportions[3] * totalWinnings,
      },
    ]);

    // confirm the second epoch can be claimed for
    await epochsManager.initializeEpoch();
    await testClaims(token, rewardsManager, owner, [
      {
        account: accounts[1],
        claim: proportions[1] * totalWinnings,
      },
      {
        account: accounts[2],
        claim: proportions[2] * totalWinnings,
      },
      {
        account: accounts[3],
        claim: proportions[3] * totalWinnings,
      },
    ]);
  });

  it('can claim staking rewards if node already joined next epoch but skipped the current epoch', async () => {
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    const { proportions } = await addStakes(token, stakingManager, owner, [
      { account: accounts[0], stake: 1000 },
      // have accounts 1, 2 and 3 as delegated stakers
      { account: accounts[1], stake: 250 },
      { account: accounts[2], stake: 400 },
      { account: accounts[3], stake: 350 },
    ]);

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(50000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    for (let j = 0; j < 2; j++) {
      await epochsManager.joinNextEpoch();
      await epochsManager.initializeEpoch();

      for (let i = 0; i < 10; i++) {
        const { ticket, redeemerRand, senderSig, receiverSig } =
          await createWinningTicket(
            syloTicketing,
            epochsManager,
            alice,
            bob,
            owner,
          );

        await syloTicketing.redeem(
          ticket,
          redeemerRand,
          senderSig,
          receiverSig,
        );
      }
    }

    // 20 tickets redeemed
    const totalWinnings = 10000;

    // initialize the next epoch but have the node skip the next one
    await epochsManager.initializeEpoch();
    await epochsManager.joinNextEpoch();

    // confirm all epochs can be claimed for
    await testClaims(token, rewardsManager, owner, [
      {
        account: accounts[1],
        claim: proportions[1] * totalWinnings,
      },
      {
        account: accounts[2],
        claim: proportions[2] * totalWinnings,
      },
      {
        account: accounts[3],
        claim: proportions[3] * totalWinnings,
      },
    ]);
  });

  it('cannot claim staking rewards with invalid arguments', async () => {
    await expect(
      rewardsManager.claimStakingRewards(ethers.ZeroAddress),
    ).to.be.revertedWithCustomError(
      rewardsManager,
      'StakeeCannotBeZeroAddress',
    );
  });

  it('claim calculation returns 0 if no rewards redeemed for an epoch', async () => {
    for (let i = 2; i < 5; i++) {
      await token.transfer(await accounts[i].getAddress(), toSOLOs(1000));
      await token
        .connect(accounts[i])
        .approve(await stakingManager.getAddress(), toSOLOs(1000));
    }

    await stakingManager.addStake(toSOLOs(1000), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    // have account 2, 3 and 4 as delegated stakers with varying levels of stake
    await stakingManager.connect(accounts[2]).addStake(toSOLOs(250), owner);
    await stakingManager.connect(accounts[3]).addStake(toSOLOs(400), owner);
    await stakingManager.connect(accounts[4]).addStake(toSOLOs(350), owner);

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(50000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    for (let i = 0; i < 10; i++) {
      const { ticket, redeemerRand, senderSig, receiverSig } =
        await createWinningTicket(
          syloTicketing,
          epochsManager,
          alice,
          bob,
          owner,
        );

      await syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig);
    }

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    // have each staker claim for the first epoch
    for (let i = 2; i < 5; i++) {
      await rewardsManager.connect(accounts[i]).claimStakingRewards(owner);
    }

    // start the next epoch without rewards being redeemed
    await epochsManager.initializeEpoch();

    // check that for each staker, the calculated claim is 0
    for (let i = 2; i < 5; i++) {
      const claim = await rewardsManager.calculateStakerClaim(
        owner,
        await accounts[i].getAddress(),
      );
      expect(claim).to.equal(0);
    }
  });

  it('can retrieve total staking rewards for an epoch', async () => {
    await stakingManager.addStake(toSOLOs(1000), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(50000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    for (let i = 0; i < 3; i++) {
      await epochsManager.joinNextEpoch();
      await epochsManager.initializeEpoch();

      for (let j = 0; j < 3 * (i + 1); j++) {
        const { ticket, redeemerRand, senderSig, receiverSig } =
          await createWinningTicket(
            syloTicketing,
            epochsManager,
            alice,
            bob,
            owner,
          );

        await syloTicketing.redeem(
          ticket,
          redeemerRand,
          senderSig,
          receiverSig,
        );
      }
    }

    for (let i = 0; i < 3; i++) {
      const totalReward = await rewardsManager.getTotalEpochRewards(i + 1);
      const totalStakingReward =
        await rewardsManager.getTotalEpochStakingRewards(i + 1);

      expect(totalReward).to.equal(toSOLOs(1000) * 3n * BigInt(i + 1));
      expect(totalStakingReward).to.equal(toSOLOs(500) * 3n * BigInt(i + 1));
    }
  });

  it('returns 0 winning probability if ticket has expired', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await registries.register('0.0.0.0/0');

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();

    await epochsManager.initializeEpoch();

    const { ticket } = await createWinningTicket(
      syloTicketing,
      epochsManager,
      alice,
      bob,
      owner,
    );

    // advance the block all the way to ticket expiry
    await utils.advanceBlock(21);

    const p = await syloTicketing.calculateWinningProbability(
      ticket.epochId,
      ticket.generationBlock,
    );

    assert.equal(p, 0n, 'Expected probability to be 0');
  });

  it('reverts when reward pool stake is signficanlty less than reward', async () => {
    // The node's stake is 2**63-1 times smaller than what the
    // reward will be.
    await ticketingParameters.setFaceValue(ethers.parseEther('10000'));
    await stakingManager.addStake(50, owner);

    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bob = Wallet.createRandom();
    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, redeemerRand, senderSig, receiverSig } =
      await createWinningTicket(
        syloTicketing,
        epochsManager,
        alice,
        bob,
        owner,
      );

    await expect(
      syloTicketing.redeem(ticket, redeemerRand, senderSig, receiverSig),
    ).to.be.revertedWithCustomError(rewardsManager, 'InvalidFixedPointResult');
  });

  it('simulates scenario between sender, node, and oracle', async () => {
    const sender = Wallet.createRandom();
    const receiver = Wallet.createRandom();
    const node = owner;

    // set up the node's stake and registry
    await stakingManager.addStake(toSOLOs(1), node);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    // set up the sender's escrow
    await syloTicketing.depositEscrow(50, sender.address);
    await syloTicketing.depositPenalty(toSOLOs(50), sender.address);

    // have the node generate random numbers
    const nodeRand = BigInt('0x' + crypto.randomBytes(32).toString('hex'));

    const generationBlock = BigInt(
      (await ethers.provider.getBlockNumber()) + 1,
    );

    // create commits from those random numbers
    const nodeCommit = createCommit(generationBlock, nodeRand);

    const epochId = await epochsManager.currentIteration();
    // create the ticket to be given to the node
    const ticket = {
      epochId,
      sender: {
        main: sender.address,
        delegated: ethers.ZeroAddress,
      },
      receiver: {
        main: receiver.address,
        delegated: ethers.ZeroAddress,
      },
      redeemer: node,
      generationBlock,
      redeemerCommit: nodeCommit,
    };

    // have sender sign the hash of the ticket
    const ticketHash = await syloTicketing.getTicketHash(ticket);
    const senderSig = await sender.signMessage(ethers.getBytes(ticketHash));
    const receiverSig = await receiver.signMessage(ethers.getBytes(ticketHash));

    // once secret has been revealed, the node can now redeem the ticket
    await syloTicketing.redeem(ticket, nodeRand, senderSig, receiverSig, {
      from: node,
    });
  });
});
