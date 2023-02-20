import { ethers } from 'hardhat';
import { BigNumber, BigNumberish, Signer, Wallet } from 'ethers';
import {
  Directory,
  EpochsManager,
  Registries,
  RewardsManager,
  StakingManager,
  SyloTicketing,
  SyloToken,
  TicketingParameters,
  TestSeekers,
} from '../../typechain';
import crypto from 'crypto';
import sodium from 'libsodium-wrappers-sumo';
import keccak256 from 'keccak256';
import * as secp256k1 from 'secp256k1';
import web3 from 'web3';
import utils from '../utils';
import { assert, expect } from 'chai';

describe('Ticketing', () => {
  let accounts: Signer[];
  let owner: string;

  const faceValue = toSOLOs(1000);
  const epochDuration = 1;

  let token: SyloToken;
  let epochsManager: EpochsManager;
  let rewardsManager: RewardsManager;
  let ticketing: SyloTicketing;
  let ticketingParameters: TicketingParameters;
  let directory: Directory;
  let registries: Registries;
  let stakingManager: StakingManager;
  let seekers: TestSeekers;

  before(async () => {
    accounts = await ethers.getSigners();
    // first account is implicitly used as deployer of contracts in hardhat
    owner = await accounts[0].getAddress();

    const Token = await ethers.getContractFactory('SyloToken');
    token = await Token.deploy();
  });

  beforeEach(async () => {
    const contracts = await utils.initializeContracts(owner, token.address, {
      faceValue,
      epochDuration,
    });
    epochsManager = contracts.epochsManager;
    rewardsManager = contracts.rewardsManager;
    ticketing = contracts.ticketing;
    ticketingParameters = contracts.ticketingParameters;
    directory = contracts.directory;
    registries = contracts.registries;
    stakingManager = contracts.stakingManager;
    seekers = contracts.seekers;

    await token.approve(stakingManager.address, toSOLOs(10000000));
    await token.approve(ticketing.address, toSOLOs(10000000));
  });

  it('cannot be initialize ticketing parameter when ticket duration less than or equal 0', async () => {
    const TicketingParameters = await ethers.getContractFactory(
      'TicketingParameters',
    );
    const ticketingParameters = await TicketingParameters.deploy();

    await expect(
      ticketingParameters.initialize(1, 1, 1, 1, 0, {
        from: owner,
      }),
    ).to.be.revertedWith('Ticket duration cannot be 0');
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

    await expect(ticketing.setUnlockDuration(3333))
      .to.emit(ticketing, 'UnlockDurationUpdated')
      .withArgs(3333);

    const currentfaceValue = await ticketingParameters.faceValue();
    assert.equal(
      currentfaceValue.toNumber(),
      777,
      'Expected face value to be correctly set',
    );

    const baseLiveWinProb = await ticketingParameters.baseLiveWinProb();
    assert.equal(
      baseLiveWinProb.toNumber(),
      888,
      'Expected base live win prob to be correctly set',
    );

    const expiredWinProb = await ticketingParameters.expiredWinProb();
    assert.equal(
      expiredWinProb.toNumber(),
      999,
      'Expected expired win prob to be correctly set',
    );

    const decayRate = await ticketingParameters.decayRate();
    assert.equal(decayRate, 1111, 'Expected decay rate to be correctly set');

    const ticketDuration = await ticketingParameters.ticketDuration();
    assert.equal(
      ticketDuration.toNumber(),
      2222,
      'Expected ticket duration to be correctly set',
    );

    const unlockDuration = await ticketing.unlockDuration();
    assert.equal(
      unlockDuration.toNumber(),
      3333,
      'Expected unlock duration to be correctly set',
    );
  });

  it('can remove managers from rewards manager', async () => {
    await rewardsManager.removeManager(stakingManager.address);
    const b = await rewardsManager.managers(stakingManager.address);
    assert.equal(
      b.toNumber(),
      0,
      'Expected staking manager to be removed as manager',
    );
  });

  it('only managers can call functions with the onlyManager constraint', async () => {
    await expect(
      rewardsManager.incrementRewardPool(owner, 10000),
    ).to.be.revertedWith(
      'Only managers of this contract can call this function',
    );
  });

  it('only managers can call functions with the onlyManager constraint', async () => {
    await expect(
      rewardsManager.initializeNextRewardPool(owner),
    ).to.be.revertedWith(
      'Only managers of this contract can call this function',
    );
  });

  it('can not set ticket duration to 0', async () => {
    await expect(ticketingParameters.setTicketDuration(0)).to.be.revertedWith(
      'Ticket duration cannot be 0',
    );
  });

  it('should be able to deposit escrow', async () => {
    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(50, alice.address);

    const deposit = await ticketing.deposits(alice.address);
    assert.equal(deposit.escrow.toString(), '50', 'Expected 50 in escrow');
  });

  it('should be able to deposit penalty', async () => {
    const alice = Wallet.createRandom();
    await ticketing.depositPenalty(50, alice.address);

    const deposit = await ticketing.deposits(alice.address);
    assert.equal(deposit.penalty.toString(), '50', 'Expected 50 in escrow');
  });

  it('should be able to deposit escrow multiple times', async () => {
    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(50, alice.address);
    await ticketing.depositEscrow(50, alice.address);

    const deposit = await ticketing.deposits(alice.address);
    assert.equal(deposit.escrow.toString(), '100', 'Expected 100 in escrow');
  });

  it('should be able to deposit to penalty multiple times', async () => {
    const alice = Wallet.createRandom();
    await ticketing.depositPenalty(50, alice.address);
    await ticketing.depositPenalty(50, alice.address);

    const deposit = await ticketing.deposits(alice.address);
    assert.equal(deposit.penalty.toString(), '100', 'Expected 100 in penalty');
  });

  it('should fail to withdraw without unlocking', async () => {
    await ticketing.depositEscrow(50, owner);
    await expect(ticketing.withdraw()).to.be.revertedWith(
      'Deposits not unlocked',
    );
  });

  it('should fail to unlock without deposit', async () => {
    await expect(ticketing.unlockDeposits()).to.be.revertedWith(
      'Nothing to withdraw',
    );
  });

  it('should be able to unlock', async () => {
    await ticketing.depositEscrow(50, owner);
    await ticketing.unlockDeposits({ from: owner });

    const deposit = await ticketing.deposits(owner);
    assert.isAbove(
      deposit.unlockAt.toNumber(),
      0,
      'Expected deposit to go into unlocking phase',
    );
  });

  it('should fail to unlock if already unlocking', async () => {
    await ticketing.depositEscrow(50, owner);
    await ticketing.unlockDeposits();

    await expect(ticketing.unlockDeposits()).to.be.revertedWith(
      'Unlock already in progress',
    );
  });

  it('should fail to lock if already locked', async () => {
    await ticketing.depositEscrow(50, owner);
    await expect(ticketing.lockDeposits()).to.be.revertedWith(
      'Not unlocking, cannot lock',
    );
  });

  it('should be able to lock deposit while it is unlocked', async () => {
    await ticketing.depositEscrow(50, owner);
    await ticketing.unlockDeposits();

    await ticketing.lockDeposits();

    const deposit = await ticketing.deposits(owner);
    assert.equal(
      deposit.unlockAt.toString(),
      '0',
      'Expected deposit to move out of unlocking phase',
    );
  });

  it('should fail to deposit while unlocking', async () => {
    await ticketing.depositEscrow(50, owner);
    await ticketing.unlockDeposits();

    await expect(ticketing.depositEscrow(10, owner)).to.be.revertedWith(
      'Cannot deposit while unlocking',
    );
    await expect(ticketing.depositPenalty(10, owner)).to.be.revertedWith(
      'Cannot deposit while unlocking',
    );
  });

  it('should be able to withdraw after unlocking phase has completed', async () => {
    await ticketing.depositEscrow(50, owner);
    await ticketing.unlockDeposits();

    await utils.advanceBlock(11);

    const balanceBeforeWithdrawal = await token.balanceOf(owner);

    await ticketing.withdraw();

    const balanceAfterWithdrawal = await token.balanceOf(owner);

    expect(balanceAfterWithdrawal).to.be.equal(balanceBeforeWithdrawal.add(50));

    // can now deposit again
    await ticketing.depositEscrow(50, owner);
  });

  it('should fail to withdraw if deposits not unlocked', async () => {
    await ticketing.depositEscrow(50, owner);
    await expect(ticketing.withdraw()).to.be.revertedWith(
      'Deposits not unlocked',
    );
  });

  it('should fail to withdraw if still unlocking', async () => {
    await ticketing.depositEscrow(50, owner);
    await ticketing.unlockDeposits();

    await expect(ticketing.withdraw()).to.be.revertedWith(
      'Unlock period not complete',
    );
  });

  it('should be able to initialize next reward pool', async () => {
    await stakingManager.addStake(30, owner);
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    const currentBlock = await ethers.provider.getBlockNumber();

    await epochsManager.joinNextEpoch();

    const rewardPool = await rewardsManager.getRewardPool(
      await epochsManager.getNextEpochId(),
      owner,
    );

    assert.isAbove(
      rewardPool.initializedAt.toNumber(),
      currentBlock,
      'Expected reward pool to track the block number it was created',
    );

    assert.equal(
      rewardPool.totalActiveStake.toString(),
      '30',
      'Expected reward pool to correctly track the stake at the time it was created',
    );
  });

  it('can not initialize reward pool more than once', async () => {
    await stakingManager.addStake(30, owner);
    await setSeekerRegistry(accounts[0], accounts[1], 1);
    await epochsManager.joinNextEpoch();

    // change the seeker but node should still be prevented from
    // initializing the reward pool again
    await setSeekerRegistry(accounts[0], accounts[1], 2);
    await expect(epochsManager.joinNextEpoch()).to.be.revertedWith(
      'The next reward pool has already been initialized',
    );
  });

  it('can not initialize reward pool more than once for the same seekers', async () => {
    await stakingManager.addStake(30, owner);
    await setSeekerRegistry(accounts[0], accounts[1], 1);
    await epochsManager.joinNextEpoch();

    await expect(epochsManager.joinNextEpoch()).to.be.revertedWith(
      'Seeker has already joined the next epoch',
    );
  });

  it('should not be able to initialize next reward pool without stake', async () => {
    await setSeekerRegistry(accounts[0], accounts[1], 1);
    await expect(epochsManager.joinNextEpoch()).to.be.revertedWith(
      'Must have stake to initialize a reward pool',
    );
  });

  it('can not redeem ticket with invalid signature', async () => {
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const { ticket, senderRand, redeemerRand } = await createWinningTicket(
      alice,
      owner,
    );

    const signature = '0x00';

    await expect(
      ticketing.redeem(ticket, senderRand, redeemerRand, signature),
    ).to.be.revertedWith('ECDSA: invalid signature length');
  });

  it('can not redeem ticket with invalid sender rand', async () => {
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const { ticket, redeemerRand, signature } = await createWinningTicket(
      alice,
      owner,
    );

    const senderRand = 999;

    await expect(
      ticketing.redeem(ticket, senderRand, redeemerRand, signature),
    ).to.be.revertedWith("Hash of senderRand doesn't match senderRandHash");
  });

  it('can not redeem ticket with invalid redeemer rand', async () => {
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const { ticket, senderRand, signature } = await createWinningTicket(
      alice,
      owner,
    );

    const redeemerRand = 999;

    await expect(
      ticketing.redeem(ticket, senderRand, redeemerRand, signature),
    ).to.be.revertedWith("Hash of redeemerRand doesn't match redeemerRandHash");
  });

  it('can not redeem ticket if associated epoch does not exist', async () => {
    const alice = Wallet.createRandom();
    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner, 1);

    await expect(
      ticketing.redeem(ticket, senderRand, redeemerRand, signature),
    ).to.be.revertedWith("Ticket's associated epoch does not exist");
  });

  it('can not calculate winning probability if associated epoch does not exist', async () => {
    const alice = Wallet.createRandom();
    const { ticket } = await createWinningTicket(alice, owner);

    ticket.epochId = 1;

    await expect(
      ticketing.calculateWinningProbability(ticket),
    ).to.be.revertedWith("Ticket's associated epoch does not exist");
  });

  it('can not redeem ticket if generated for a future block', async () => {
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    const updatedTicket = { ...ticket, generationBlock: 100000 };

    await expect(
      ticketing.redeem(updatedTicket, senderRand, redeemerRand, signature),
    ).to.be.revertedWith('The ticket cannot be generated for a future block');
  });

  it('can not calculate winning probablility if not generated during associated epoch', async () => {
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const { ticket } = await createWinningTicket(alice, owner);

    const updatedTicket = { ...ticket, generationBlock: 1 };

    await expect(
      ticketing.calculateWinningProbability(updatedTicket),
    ).to.be.revertedWith(
      "This ticket was not generated during it's associated epoch",
    );
  });

  it('can not redeem ticket if node has not joined directory', async () => {
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(2000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    await expect(
      ticketing.redeem(ticket, senderRand, redeemerRand, signature),
    ).to.be.revertedWith(
      'Ticket redeemer must have joined the directory for this epoch',
    );
  });

  it('can not redeem ticket if node has not initialized reward pool', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    await directory.addManager(owner);
    await directory.joinNextDirectory(owner);

    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(2000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    await expect(
      ticketing.redeem(ticket, senderRand, redeemerRand, signature),
    ).to.be.revertedWith(
      'Reward pool has not been initialized for the current epoch',
    );
  });

  it('can not redeem invalid ticket', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(2000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    let malformedTicket = { ...ticket };
    malformedTicket.sender = ethers.constants.AddressZero;
    await expect(
      ticketing.redeem(malformedTicket, senderRand, redeemerRand, signature),
    ).to.be.revertedWith('Ticket sender is null');

    malformedTicket = { ...ticket };
    malformedTicket.redeemer = ethers.constants.AddressZero;
    await expect(
      ticketing.redeem(malformedTicket, senderRand, redeemerRand, signature),
    ).to.be.revertedWith('Ticket redeemer is null');

    malformedTicket = { ...ticket };
    malformedTicket.senderCommit =
      '0x0000000000000000000000000000000000000000000000000000000000000000';
    await expect(
      ticketing.redeem(malformedTicket, senderRand, redeemerRand, signature),
    ).to.be.revertedWith("Hash of senderRand doesn't match senderRandHash");

    malformedTicket = { ...ticket };
    malformedTicket.redeemerCommit =
      '0x0000000000000000000000000000000000000000000000000000000000000000';
    await expect(
      ticketing.redeem(malformedTicket, senderRand, redeemerRand, signature),
    ).to.be.revertedWith("Hash of redeemerRand doesn't match redeemerRandHash");

    const malformedSig =
      '0xdebcaaaa727df04bdc990083d88ed7c8e6e9897ff18b7d968867a8bc024cbdbe10ca52eebd67a14b7b493f5c00ed9dab7b96ef62916f25afc631d336f7b2ae1e1b';
    await expect(
      ticketing.redeem(ticket, senderRand, redeemerRand, malformedSig),
    ).to.be.revertedWith("Ticket doesn't have a valid signature");
  });

  it('rejects non winning ticket', async () => {
    // redeploy contracts with win chance of 0%
    const contracts = await utils.initializeContracts(owner, token.address, {
      baseLiveWinProb: 0,
    });
    await token.approve(contracts.stakingManager.address, toSOLOs(100000));
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
      contracts.epochsManager.address,
    );
    await contracts.epochsManager.initializeEpoch();

    await token.approve(contracts.ticketing.address, toSOLOs(100000));
    const alice = Wallet.createRandom();
    await contracts.ticketing.depositEscrow(toSOLOs(2000), alice.address);
    await contracts.ticketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner, 1);

    await utils.advanceBlock(5);

    await expect(
      contracts.ticketing.redeem(ticket, senderRand, redeemerRand, signature),
    ).to.be.revertedWith('Ticket is not a winner');
  });

  it('can redeem winning ticket', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(2000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    await ticketing.redeem(ticket, senderRand, redeemerRand, signature);

    const deposit = await ticketing.deposits(alice.address);
    assert.equal(
      deposit.escrow.toString(),
      toSOLOs(1000),
      'Expected ticket payout to be substracted from escrow',
    );
    assert.equal(
      deposit.penalty.toString(),
      toSOLOs(50),
      'Expected penalty to not be changed',
    );

    const pendingReward = await rewardsManager.getPendingRewards(owner);

    assert.equal(
      pendingReward.toString(),
      toSOLOs(500),
      'Expected balance of pending rewards to have added the ticket face value',
    );
  });

  it('can not redeem ticket more than once', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(2000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
    await expect(
      ticketing.redeem(ticket, senderRand, redeemerRand, signature),
    ).to.be.revertedWith('Ticket already redeemed');
  });

  it('burns penalty on insufficient escrow', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(5), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    const initialTicketingBalance = await token.balanceOf(ticketing.address);

    await ticketing.redeem(ticket, senderRand, redeemerRand, signature);

    const deposit = await ticketing.deposits(alice.address);
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

    const ticketingBalance = await token.balanceOf(ticketing.address);
    assert.equal(
      ticketingBalance.toString(),
      initialTicketingBalance.sub(toSOLOs(55)).toString(),
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
    await expect(rewardsManager.claimStakingRewards(owner)).to.be.revertedWith(
      'Nothing to claim',
    );
  });

  it('can claim ticketing rewards', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(50000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    for (let i = 0; i < 10; i++) {
      const { ticket, senderRand, redeemerRand, signature } =
        await createWinningTicket(alice, owner);

      await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
    }

    await epochsManager.initializeEpoch();

    const initialBalance = await token.balanceOf(owner);

    await rewardsManager.claimStakingRewards(owner);

    const postBalance = await token.balanceOf(owner);
    // Expect the node have the entire reward balance added to their account
    const expectedPostBalance = initialBalance.add(toSOLOs(10000));

    compareExpectedBalance(expectedPostBalance, postBalance);

    const pendingReward = await rewardsManager.getPendingRewards(owner);

    compareExpectedBalance(pendingReward, 0);

    // check total rewards in the previous epoch after claiming
    const nextEpochId = await epochsManager.getNextEpochId();
    const rewardPoolStakersTotal =
      await rewardsManager.getRewardPoolStakersTotal(nextEpochId.sub(2), owner);

    assert.equal(
      rewardPoolStakersTotal.toString(),
      toSOLOs(500 * 10), // 500 is added to the stakers reward total on each redemption (50% of 1000)
      'Expected reward pool stakers total in the previous epoch to be 5000 SOLOs',
    );
  });

  it('delegated stakers should be able to claim rewards', async () => {
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    const { proportions } = await addStakes([
      { account: accounts[0], stake: 3 },
      // have account 2 as a delegated staker
      { account: accounts[2], stake: 2 },
    ]);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(50000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    for (let i = 0; i < 10; i++) {
      const { ticket, senderRand, redeemerRand, signature } =
        await createWinningTicket(alice, owner);

      await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
    }

    const totalWinnings = 5000;

    await epochsManager.initializeEpoch();

    await testClaims([
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
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    await addStakes([
      { account: accounts[0], stake: 10000 },
      // have account 2 as a delegated staker
      { account: accounts[2], stake: 2 },
    ]);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(50000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    for (let i = 0; i < 10; i++) {
      const { ticket, senderRand, redeemerRand, signature } =
        await createWinningTicket(alice, owner);

      await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
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
    await addStakes([{ account: accounts[2], stake: 1 }]);

    // pending reward after stake is added should have previous claim removed
    const pendingRewardAfterAddingStake =
      await rewardsManager.getPendingRewards(owner);

    assert.equal(
      pendingRewardBeforeAddingStake
        .sub(pendingRewardAfterAddingStake)
        .toString(),
      claimBeforeAddingStake.toString(),
      'Expected claim to be removed from pending reward after adding stake',
    );

    await rewardsManager.connect(accounts[2]).claimStakingRewards(owner);

    for (let i = 0; i < 10; i++) {
      const { ticket, senderRand, redeemerRand, signature } =
        await createWinningTicket(alice, owner);

      await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
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
      pendingRewardBeforeRemovingStake
        .sub(pendingRewardAfterRemovingStake)
        .toString(),
      claimBeforeRemovingStake.toString(),
      'Expected claim to be removed from pending reward after removing stake',
    );
  });

  it('can calculate staker claim if reward total is 0', async () => {
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    await addStakes([
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

  it('can not claim reward more than once for the same epoch', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(50000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    await ticketing.redeem(ticket, senderRand, redeemerRand, signature);

    await epochsManager.initializeEpoch();

    await rewardsManager.claimStakingRewards(owner);

    const lastClaim = await rewardsManager.getLastClaim(owner, owner);
    expect(lastClaim.claimedAt).to.be.above(0);

    await expect(rewardsManager.claimStakingRewards(owner)).to.be.revertedWith(
      'Nothing to claim',
    );
  });

  it('should be able to correctly calculate staking rewards for multiple epochs when managed stake is the same', async () => {
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    const { proportions } = await addStakes([
      { account: accounts[0], stake: 1000 },
      // have accounts 1, 2 and 3 as delegated stakers
      { account: accounts[1], stake: 250 },
      { account: accounts[2], stake: 400 },
      { account: accounts[3], stake: 350 },
    ]);

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(500000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    for (let j = 0; j < 3; j++) {
      await epochsManager.joinNextEpoch();
      await epochsManager.initializeEpoch();

      // 500 is added to the stakers reward total on each redemption (50% of 1000)
      for (let i = 0; i < 6; i++) {
        const { ticket, senderRand, redeemerRand, signature } =
          await createWinningTicket(alice, owner);

        await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
      }
    }

    await epochsManager.initializeEpoch();

    const pendingReward = await rewardsManager.getPendingRewards(owner);

    // the total unclaimed stake reward should 3 * 6 * 500 = 9000
    const totalWinnings = 9000;

    compareExpectedBalance(toSOLOs(totalWinnings), pendingReward);

    // verify each staker will receive the correct amount of reward if they were to claim now
    await testClaims([
      { account: accounts[1], claim: proportions[1] * totalWinnings },
      { account: accounts[2], claim: proportions[2] * totalWinnings },
      { account: accounts[3], claim: proportions[3] * totalWinnings },
    ]);
  });

  it('should be able to stake, accumulate rewards, and claim more than once as a delegated staker', async () => {
    const { proportions } = await addStakes([
      { account: accounts[0], stake: 1000 },
      // have accounts 1, 2 and 3 as delegated stakers
      { account: accounts[1], stake: 250 },
      { account: accounts[2], stake: 400 },
      { account: accounts[3], stake: 350 },
    ]);

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(500000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    await setSeekerRegistry(accounts[0], accounts[1], 1);
    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    for (let runs = 0; runs < 10; runs++) {
      // 250 is added to the delegated stakers reward total on each redemption
      for (let i = 0; i < 5; i++) {
        const { ticket, senderRand, redeemerRand, signature } =
          await createWinningTicket(alice, owner);

        await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
      }

      await epochsManager.joinNextEpoch();
      await epochsManager.initializeEpoch();

      const totalWinnings = 2500;

      // No accounts adds or withdraws stake, so stake proportions remain constant
      await testClaims([
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
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    await addStakes([
      { account: accounts[0], stake: 1000 },
      // have accounts 1, 2 and 3 as delegated stakers
      { account: accounts[1], stake: 250 },
      { account: accounts[2], stake: 400 },
      { account: accounts[3], stake: 350 },
    ]);

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(50000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    // have account 4 add stake midway through
    for (let j = 0; j < 3; j++) {
      if (j == 1) {
        await addStakes([{ account: accounts[4], stake: 500 }]);
      }

      await epochsManager.joinNextEpoch();
      await epochsManager.initializeEpoch();

      // 500 is added to the stakers reward total on each redemption (50% of 1000)
      for (let i = 0; i < 6; i++) {
        const { ticket, senderRand, redeemerRand, signature } =
          await createWinningTicket(alice, owner);

        await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
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
    const s = BigNumber.from(toSOLOs(500)); // initial stake
    const r = toSOLOs(2 * 6 * 500); // accumulated reward
    const expectedStakeClaimFive = s.mul(r).div(epochTwoActiveStake);
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
      const epochOneReward = BigNumber.from(toSOLOs(6 * 500))
        .mul(initialStake)
        .div(epochOneActiveStake);

      const remainingReward = BigNumber.from(toSOLOs(2 * 6 * 500))
        .mul(initialStake)
        .div(epochTwoActiveStake);

      const totalExpectedReward = epochOneReward.add(remainingReward);
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
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    await addStakes([
      { account: accounts[0], stake: 1000 },
      // have accounts 1, 2 and 3 as delegated stakers
      { account: accounts[1], stake: 250 },
      { account: accounts[2], stake: 400 },
      { account: accounts[3], stake: 350 },
    ]);

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(50000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

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
        const { ticket, senderRand, redeemerRand, signature } =
          await createWinningTicket(alice, owner);

        await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
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
      const epochOneReward = BigNumber.from(toSOLOs(6 * 500))
        .mul(initialStake)
        .div(epochOneActiveStake);

      const remainingReward = BigNumber.from(toSOLOs(2 * 6 * 500))
        .mul(initialStake)
        .div(epochTwoActiveStake);

      const totalExpectedReward = epochOneReward.add(remainingReward);
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
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    const { proportions } = await addStakes([
      { account: accounts[0], stake: 1000 },
      // have accounts 1, 2 and 3 as delegated stakers
      { account: accounts[1], stake: 250 },
      { account: accounts[2], stake: 400 },
      { account: accounts[3], stake: 350 },
    ]);

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(1000 * 500), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const iterations = 450;

    for (let i = 0; i < iterations; i++) {
      const { ticket, senderRand, redeemerRand, signature } =
        await createWinningTicket(alice, owner);

      await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
    }

    await epochsManager.initializeEpoch();

    await testClaims([
      { account: accounts[1], claim: iterations * 500 * proportions[1] },
      { account: accounts[2], claim: iterations * 500 * proportions[2] },
      { account: accounts[3], claim: iterations * 500 * proportions[3] },
    ]);
  }).timeout(0);

  it('should decay winning probability as ticket approaches expiry', async () => {
    // deploy another ticketing contract with simpler parameters
    const contracts = await utils.initializeContracts(owner, token.address, {
      faceValue,
      baseLiveWinProb: 100000,
      expiredWinProb: 1000,
      decayRate: 8000,
      ticketDuration: 100,
    });
    epochsManager = contracts.epochsManager;
    rewardsManager = contracts.rewardsManager;
    ticketing = contracts.ticketing;
    directory = contracts.directory;
    registries = contracts.registries;
    stakingManager = contracts.stakingManager;

    await directory.transferOwnership(epochsManager.address);
    await rewardsManager.addManager(ticketing.address);
    await rewardsManager.addManager(epochsManager.address);

    await token.approve(ticketing.address, toSOLOs(10000));
    await token.approve(stakingManager.address, toSOLOs(10000));

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
    await ticketing.depositEscrow(toSOLOs(5000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket } = await createWinningTicket(alice, owner);

    // advance the block halfway to ticket expiry
    await utils.advanceBlock(51);

    // check if the probability has decayed 50% of the maximum decayed value (80%)
    const expectedProbability = 100000 - 0.5 * 0.8 * 100000;

    const decayedProbability = await ticketing.calculateWinningProbability(
      ticket,
    );

    assert.equal(
      decayedProbability.toString(),
      expectedProbability.toString(),
      'Expected probablity of ticket winning to decay',
    );
  });

  it('should be able to correctly calculate staker rewards if node was not active for multiple epochs', async () => {
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    const { proportions } = await addStakes([
      { account: accounts[0], stake: 1000 },
      // have accounts 1, 2 and 3 as delegated stakers
      { account: accounts[1], stake: 250 },
      { account: accounts[2], stake: 400 },
      { account: accounts[3], stake: 350 },
    ]);

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(500000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

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
        const { ticket, senderRand, redeemerRand, signature } =
          await createWinningTicket(alice, owner);

        await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
      }
    }

    await epochsManager.initializeEpoch();

    const pendingReward = await rewardsManager.getPendingRewards(owner);

    // the total unclaimed stake reward should 3 * 6 * 500 = 9000
    const totalWinnings = 9000;

    compareExpectedBalance(toSOLOs(totalWinnings), pendingReward);

    // verify each staker will receive the correct amount of reward if they were to claim now
    await testClaims([
      { account: accounts[1], claim: proportions[1] * totalWinnings },
      { account: accounts[2], claim: proportions[2] * totalWinnings },
      { account: accounts[3], claim: proportions[3] * totalWinnings },
    ]);
  });

  it('claiming staking rewards only claims up to the previous epoch', async () => {
    await addStakes([
      { account: accounts[0], stake: 1 },
      { account: accounts[1], stake: 1 },
    ]);

    await setSeekerRegistry(accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(50000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    for (let j = 0; j < 3; j++) {
      for (let i = 0; i < 10; i++) {
        const { ticket, senderRand, redeemerRand, signature } =
          await createWinningTicket(alice, owner);

        await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
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
    await addStakes([
      { account: accounts[0], stake: 1 },
      { account: accounts[1], stake: 1 },
    ]);

    await setSeekerRegistry(accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(50000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    for (let j = 0; j < 2; j++) {
      for (let i = 0; i < 10; i++) {
        const { ticket, senderRand, redeemerRand, signature } =
          await createWinningTicket(alice, owner);

        await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
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
      const { ticket, senderRand, redeemerRand, signature } =
        await createWinningTicket(alice, owner);

      await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
    }

    // confirm claim is still 0
    claim = await rewardsManager.calculateStakerClaim(
      owner,
      await accounts[1].getAddress(),
    );

    compareExpectedBalance(claim, 0);

    await expect(
      rewardsManager.connect(accounts[1]).claimStakingRewards(owner),
    ).to.be.revertedWith('Nothing to claim');
  });

  it('can claim staking rewards again after previous ended', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(50000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    const initialBalance = await token.balanceOf(owner);

    for (let i = 0; i < 10; i++) {
      const { ticket, senderRand, redeemerRand, signature } =
        await createWinningTicket(alice, owner);

      await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
    }

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    await rewardsManager.claimStakingRewards(owner);

    for (let i = 0; i < 10; i++) {
      const { ticket, senderRand, redeemerRand, signature } =
        await createWinningTicket(alice, owner);

      await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
    }

    await epochsManager.initializeEpoch();

    await rewardsManager.claimStakingRewards(owner);

    const postBalance = await token.balanceOf(owner);

    // expect to receive total balance of all redeemed tickets
    compareExpectedBalance(postBalance, initialBalance.add(toSOLOs(20000)));
  });

  it('can claim staking rewards if node already joined next epoch', async () => {
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    const { proportions } = await addStakes([
      { account: accounts[0], stake: 1000 },
      // have accounts 1, 2 and 3 as delegated stakers
      { account: accounts[1], stake: 250 },
      { account: accounts[2], stake: 400 },
      { account: accounts[3], stake: 350 },
    ]);

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(50000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    for (let j = 0; j < 2; j++) {
      await epochsManager.joinNextEpoch();
      await epochsManager.initializeEpoch();

      for (let i = 0; i < 10; i++) {
        const { ticket, senderRand, redeemerRand, signature } =
          await createWinningTicket(alice, owner);

        await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
      }
    }

    // 10 tickets redeemed per epoch
    const totalWinnings = 5000;

    // have the node join the next epoch and test stakers are able
    // to claim for the first epoch
    await epochsManager.joinNextEpoch();
    await testClaims([
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
    await testClaims([
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
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    const { proportions } = await addStakes([
      { account: accounts[0], stake: 1000 },
      // have accounts 1, 2 and 3 as delegated stakers
      { account: accounts[1], stake: 250 },
      { account: accounts[2], stake: 400 },
      { account: accounts[3], stake: 350 },
    ]);

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(50000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    for (let j = 0; j < 2; j++) {
      await epochsManager.joinNextEpoch();
      await epochsManager.initializeEpoch();

      for (let i = 0; i < 10; i++) {
        const { ticket, senderRand, redeemerRand, signature } =
          await createWinningTicket(alice, owner);

        await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
      }
    }

    // 20 tickets redeemed
    const totalWinnings = 10000;

    // initialize the next epoch but have the node skip the next one
    await epochsManager.initializeEpoch();
    await epochsManager.joinNextEpoch();

    // confirm all epochs can be claimed for
    await testClaims([
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

  it('claim calculation returns 0 if no rewards redeemed for an epoch', async () => {
    for (let i = 2; i < 5; i++) {
      await token.transfer(await accounts[i].getAddress(), toSOLOs(1000));
      await token
        .connect(accounts[i])
        .approve(stakingManager.address, toSOLOs(1000));
    }

    await stakingManager.addStake(toSOLOs(1000), owner);
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    // have account 2, 3 and 4 as delegated stakers with varying levels of stake
    await stakingManager.connect(accounts[2]).addStake(toSOLOs(250), owner);
    await stakingManager.connect(accounts[3]).addStake(toSOLOs(400), owner);
    await stakingManager.connect(accounts[4]).addStake(toSOLOs(350), owner);

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(50000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    for (let i = 0; i < 10; i++) {
      const { ticket, senderRand, redeemerRand, signature } =
        await createWinningTicket(alice, owner);

      await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
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
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(toSOLOs(50000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    for (let i = 0; i < 3; i++) {
      await epochsManager.joinNextEpoch();
      await epochsManager.initializeEpoch();

      for (let j = 0; j < 3 * (i + 1); j++) {
        const { ticket, senderRand, redeemerRand, signature } =
          await createWinningTicket(alice, owner);

        await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
      }
    }

    for (let i = 0; i < 3; i++) {
      const totalReward = await rewardsManager.getTotalEpochRewards(i + 1);
      const totalStakingReward =
        await rewardsManager.getTotalEpochStakingRewards(i + 1);

      expect(totalReward.toString()).to.equal(
        BigNumber.from(toSOLOs(1000))
          .mul(3)
          .mul(i + 1)
          .toString(),
      );
      expect(totalStakingReward.toString()).to.equal(
        BigNumber.from(toSOLOs(500))
          .mul(3)
          .mul(i + 1)
          .toString(),
      );
    }
  });

  it('returns 0 winning probability if ticket has expired', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await registries.register('0.0.0.0/0');

    const alice = Wallet.createRandom();

    await epochsManager.initializeEpoch();

    const { ticket } = await createWinningTicket(alice, owner);

    // advance the block all the way to ticket expiry
    await utils.advanceBlock(21);

    const p = await ticketing.calculateWinningProbability(ticket);

    assert.equal('0', p.toString(), 'Expected probability to be 0');
  });

  it('simulates scenario between sender, node, and oracle', async () => {
    const sender = Wallet.createRandom();
    const node = owner;

    // set up the node's stake and registry
    await stakingManager.addStake(toSOLOs(1), node);
    await setSeekerRegistry(accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    // set up the sender's escrow
    await ticketing.depositEscrow(50, sender.address);
    await ticketing.depositPenalty(toSOLOs(50), sender.address);

    // have the node and sender generate random numbers
    const nodeRand = crypto.randomBytes(32);
    const senderRand = crypto.randomBytes(32);

    const generationBlock = (await ethers.provider.getBlockNumber()) + 1;

    // create commits from those random numbers
    const senderCommit = createCommit(generationBlock, senderRand);
    const nodeCommit = createCommit(generationBlock, nodeRand);

    const epochId = await epochsManager
      .getCurrentActiveEpoch()
      .then(e => e.iteration);

    // create the ticket to be given to the node
    const ticket = {
      epochId,
      sender: sender.address,
      redeemer: node,
      generationBlock,
      senderCommit,
      redeemerCommit: nodeCommit,
    };

    // have sender sign the hash of the ticket
    const ticketHash = await ticketing.getTicketHash(ticket);
    const sigObj = secp256k1.ecdsaSign(
      Buffer.from(ticketHash.slice(2), 'hex'),
      Buffer.from(sender.privateKey.slice(2), 'hex'),
    );
    const signature = Buffer.concat([
      sigObj.signature,
      new Uint8Array([sigObj.recid]),
    ]);

    // establish the oracle
    const oracle = sodium.crypto_sign_keypair('uint8array');

    // encrypt senderRandom to create the key
    const key = sodium.crypto_box_seal(
      senderRand,
      sodium.crypto_sign_ed25519_pk_to_curve25519(oracle.publicKey),
      'uint8array',
    );

    // have oracle decrypt the key and reveal the random number to the node
    const revealedSenderRand = sodium.crypto_box_seal_open(
      key,
      sodium.crypto_sign_ed25519_pk_to_curve25519(oracle.publicKey),
      sodium.crypto_sign_ed25519_sk_to_curve25519(oracle.privateKey),
      'uint8array',
    );

    // once secret has been revealed, the node can now redeem the ticket
    await ticketing.redeem(ticket, revealedSenderRand, nodeRand, signature, {
      from: node,
    });
  });

  // This test suite relies on confirming that updated stakes and rewards are correctly
  // calculated after incrementing the reward pool. However due to minor precision loss, the
  // actual balance may slightly differ.
  // This function checks the difference falls within a small fraction of a single SYLO.
  function compareExpectedBalance(a: BigNumberish, b: BigNumberish) {
    const diff = BigNumber.from(a).sub(BigNumber.from(b)).abs();
    // NOTE: This essentially says that a margin of 10**4 SOLOs is acceptable, or
    // 0.00000000000001 SYLOs
    expect(diff.toNumber()).to.be.within(0, 10 ** 4);
  }

  function toSOLOs(a: number): string {
    return web3.utils.toWei(a.toString());
  }

  function fromSOLOs(a: number | BigNumber | string): string {
    return web3.utils.fromWei(a.toString());
  }

  // Helper function to initialize stakes for multiple stakees,
  // and returns an array of stake proportions.
  const addStakes = async (
    stakees: { account: Signer; stake: number }[],
  ): Promise<{ totalStake: number; proportions: number[] }> => {
    const totalStake = stakees.reduce((p, c) => {
      return { ...p, stake: p.stake + c.stake };
    }).stake;

    const proportions = await Promise.all(
      stakees.map(async s => {
        const stake = toSOLOs(s.stake);

        await token.transfer(await s.account.getAddress(), stake);
        await token.connect(s.account).approve(stakingManager.address, stake);

        await stakingManager.connect(s.account).addStake(stake, owner);

        return s.stake / totalStake;
      }),
    );

    return { totalStake, proportions };
  };

  // Helper function for testing that an account's SYLO balance
  // increased as expected after making a claim
  const testClaims = async (tests: { account: Signer; claim: number }[]) => {
    for (const t of tests) {
      const claim = toSOLOs(t.claim);
      const address = await t.account.getAddress();

      const expectedBalance = await token
        .balanceOf(address)
        .then(b => b.add(claim));

      await rewardsManager.connect(t.account).claimStakingRewards(owner);

      compareExpectedBalance(expectedBalance, await token.balanceOf(address));
    }
  };

  async function setSeekerRegistry(
    account: Signer,
    seekerAccount: Signer,
    tokenId: number,
  ) {
    await utils.setSeekerRegistry(
      registries,
      seekers,
      account,
      seekerAccount,
      tokenId,
    );
  }

  async function createWinningTicket(
    sender: Wallet,
    redeemer: string,
    epochId?: number,
  ) {
    const generationBlock = (await ethers.provider.getBlockNumber()) + 1;

    const senderRand = 1;
    const senderCommit = createCommit(generationBlock, senderRand);

    const redeemerRand = 1;
    const redeemerCommit = createCommit(generationBlock, redeemerRand);

    const ticket = {
      epochId: epochId ?? (await epochsManager.currentIteration()),
      sender: sender.address,
      redeemer,
      generationBlock,
      senderCommit: '0x' + senderCommit.toString('hex'),
      redeemerCommit: '0x' + redeemerCommit.toString('hex'),
    };

    const ticketHash = await ticketing.getTicketHash(ticket);

    // eslint-disable-next-line @typescript-eslint/no-unsafe-member-access
    const { signature, recid } = secp256k1.ecdsaSign(
      Buffer.from(ticketHash.slice(2), 'hex'),
      Buffer.from(sender.privateKey.slice(2), 'hex'),
    );

    if (!signature) {
      throw new Error('failed to derive signature for ticket');
    }

    return {
      ticket,
      senderRand,
      redeemerRand,
      signature: Buffer.concat([signature, new Uint8Array([recid])]),
      ticketHash,
    };
  }

  function createCommit(
    generationBlock: BigNumberish,
    rand: BigNumberish,
  ): Buffer {
    return keccak256(
      ethers.utils.defaultAbiCoder.encode(
        ['bytes32'],
        [
          keccak256(
            ethers.utils.defaultAbiCoder.encode(
              ['uint256', 'uint256'],
              [generationBlock, rand],
            ),
          ),
        ],
      ),
    );
  }
});
