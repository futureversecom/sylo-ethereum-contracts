import { ethers } from "hardhat";
import { BigNumber, BigNumberish, Signer, Wallet } from "ethers";
import { Directory, EpochsManager, Listings, RewardsManager, StakingManager, SyloTicketing, SyloToken, TicketingParameters } from "../../typechain";
const crypto = require("crypto");
const sodium = require('libsodium-wrappers-sumo');
const eth = require('eth-lib');
import web3 from "web3";
import { soliditySha3 } from "web3-utils";
import utils from '../utils';
import { assert, expect } from "chai";

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
  let listings: Listings;
  let stakingManager: StakingManager;

  before(async () => {
    accounts = await ethers.getSigners();
    // first account is implicitly used as deployer of contracts in hardhat
    owner = await accounts[0].getAddress();

    const Token = await ethers.getContractFactory("SyloToken");
    token = await Token.deploy() as SyloToken;
  });

  beforeEach(async () => {
    const contracts = await utils.initializeContracts(owner, token.address, { faceValue, epochDuration });
    epochsManager = contracts.epochsManager;
    rewardsManager = contracts.rewardsManager;
    ticketing = contracts.ticketing;
    ticketingParameters = contracts.ticketingParameters;
    directory = contracts.directory;
    listings = contracts.listings;
    stakingManager = contracts.stakingManager;

    await token.approve(stakingManager.address, toSOLOs(10000000));
    await token.approve(ticketing.address, toSOLOs(10000000));
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

    const faceValue = await ticketingParameters.faceValue();
    assert.equal(faceValue.toNumber(), 777, "Expected face value to be correctly set");

    const baseLiveWinProb = await ticketingParameters.baseLiveWinProb();
    assert.equal(baseLiveWinProb.toNumber(), 888, "Expected base live win prob to be correctly set");

    const expiredWinProb = await ticketingParameters.expiredWinProb();
    assert.equal(expiredWinProb.toNumber(), 999, "Expected expired win prob to be correctly set");

    const decayRate = await ticketingParameters.decayRate();
    assert.equal(decayRate, 1111, "Expected decay rate to be correctly set");

    const ticketDuration = await ticketingParameters.ticketDuration();
    assert.equal(ticketDuration.toNumber(), 2222, "Expected ticket duration to be correctly set");

    const unlockDuration = await ticketing.unlockDuration();
    assert.equal(unlockDuration.toNumber(), 3333, "Expected unlock duration to be correctly set");
  });

  it('can remove managers from rewards manager', async () => {
    await rewardsManager.removeManager(stakingManager.address);
    const b = await rewardsManager.managers(stakingManager.address);
    assert.equal(
      b.toNumber(),
      0,
      "Expected staking manager to be removed as manager"
    );
  });

  it('Only managers can call functions with the onlyManager constraint', async () => {
    await expect(rewardsManager.incrementRewardPool(owner, 10000))
      .to.be.revertedWith("Only managers of this contract can call this function");
  });

  it('Only managers can call functions with the onlyManager constraint', async () => {
    await expect(rewardsManager.initializeNextRewardPool(owner))
      .to.be.revertedWith("Only managers of this contract can call this function");
  });

  it('can not set ticket duration to 0', async () => {
    await expect(ticketingParameters.setTicketDuration(0))
      .to.be.revertedWith("Ticket duration cannot be 0");
  })

  it('should be able to deposit escrow', async () => {
    const alice = Wallet.createRandom();
    await ticketing.depositEscrow(50, alice.address);

    const deposit = await ticketing.deposits(alice.address);
    assert.equal(deposit.escrow.toString(), '50', 'Expected 50 in escrow');
  });

  it('should be able to deposit penalty', async () => {
    const alice = Wallet.createRandom()
    await ticketing.depositPenalty(50, alice.address);

    const deposit = await ticketing.deposits(alice.address);
    assert.equal(deposit.penalty.toString(), '50', 'Expected 50 in escrow');
  });

  it('should be able to deposit escrow multiple times', async () => {
    const alice = Wallet.createRandom()
    await ticketing.depositEscrow(50, alice.address);
    await ticketing.depositEscrow(50, alice.address);

    const deposit = await ticketing.deposits(alice.address);
    assert.equal(deposit.escrow.toString(), '100', 'Expected 100 in escrow');
  });

  it('should be able to deposit to penalty multiple times', async () => {
    const alice = Wallet.createRandom()
    await ticketing.depositPenalty(50, alice.address);
    await ticketing.depositPenalty(50, alice.address);

    const deposit = await ticketing.deposits(alice.address);
    assert.equal(deposit.penalty.toString(), '100', 'Expected 100 in penalty');
  });

  it('should fail to withdraw without unlocking', async () => {
    await ticketing.depositEscrow(50, owner);

    await ticketing.withdraw()
      .then(() => {
        assert.fail('Withdrawing should fail');
      })
      .catch(e => {
        assert.include(e.message, 'Deposits not unlocked', 'Withdraw should fail due to not being unlocked');
      });
  });

  it('should fail to unlock without deposit', async () => {
    await ticketing.unlockDeposits()
      .then(() => {
        assert.fail('Withdrawing should fail');
      })
      .catch(e => {
        assert.include(e.message, 'Nothing to withdraw', 'Unlock should fail due to no deposit');
      });
  });

  it('should be able to unlock', async () => {
    await ticketing.depositEscrow(50, owner);
    await ticketing.unlockDeposits({ from: owner});

    const deposit = await ticketing.deposits(owner);
    assert.isAbove(deposit.unlockAt.toNumber(), 0, 'Expected deposit to go into unlocking phase');
  });

  it('should fail to unlock if already unlocking', async () => {
    await ticketing.depositEscrow(50, owner);
    await ticketing.unlockDeposits();

    await ticketing.unlockDeposits()
      .then(() => {
        assert.fail('Withdrawing should fail');
      })
      .catch(e => {
        assert.include(e.message, 'Unlock already in progress', 'Unlock should fail due to already unlocking');
      });
  });

  it('should fail to lock if already locked', async () => {
    await ticketing.depositEscrow(50, owner);
    await ticketing.lockDeposits()
      .then(() => {
        assert.fail('Locking should fail');
      })
      .catch(e => {
        assert.include(e.message, 'Not unlocking, cannot lock', 'Expect lock to fail as it deposit is already locked');
      });
  });

  it('should be able to lock deposit while it is unlocked', async () => {
    await ticketing.depositEscrow(50, owner);
    await ticketing.unlockDeposits();

    await ticketing.lockDeposits();

    const deposit = await ticketing.deposits(owner);
    assert.equal(deposit.unlockAt.toString(), '0', 'Expected deposit to move out of unlocking phase');
  });

  it('should fail to deposit while unlocking', async () => {
    await ticketing.depositEscrow(50, owner);
    await ticketing.unlockDeposits();

    await expect(ticketing.depositEscrow(10, owner))
      .to.be.revertedWith("Cannot deposit while unlocking");
    await expect(ticketing.depositPenalty(10, owner))
      .to.be.revertedWith("Cannot deposit while unlocking");
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
    await expect(ticketing.withdraw())
      .to.be.revertedWith("Deposits not unlocked");
  });

  it('should fail to withdraw if still unlocking', async () => {
    await ticketing.depositEscrow(50, owner);
    await ticketing.unlockDeposits();

    await expect(ticketing.withdraw())
      .to.be.revertedWith("Unlock period not complete");
  })

  it('should be able to initialize next reward pool', async () => {
    await stakingManager.addStake(30, owner);

    const currentBlock = await ethers.provider.getBlockNumber();
    await epochsManager.joinNextEpoch();

    const rewardPool = await rewardsManager.getRewardPool(
      await epochsManager.getNextEpochId(),
      owner
    );

    assert.isAbove(
      rewardPool.initializedAt.toNumber(),
      currentBlock,
      "Expected reward pool to track the block number it was created"
    );

    assert.equal(
      rewardPool.totalActiveStake.toString(),
      "30",
      "Expected reward pool to correctly track the stake at the time it was created"
    );
  });

  it('can not initialize reward pool more than once', async () => {
    await stakingManager.addStake(30, owner);
    await epochsManager.joinNextEpoch();
    await expect(epochsManager.joinNextEpoch())
      .to.be.revertedWith("The next reward pool has already been initialized");
  });

  it('should not be able to initialize next reward pool without stake', async () => {
    await expect(epochsManager.joinNextEpoch())
      .to.be.revertedWith("Must have stake to initialize a reward pool");
  });

  it('can not redeem ticket with invalid signature', async () => {
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom()
    const { ticket, senderRand, redeemerRand } =
      await createWinningTicket(alice, owner);

    const signature = '0x00';

    await expect(ticketing.redeem(ticket, senderRand, redeemerRand, signature))
      .to.be.revertedWith('ECDSA: invalid signature length');
  });

  it('can not redeem ticket with invalid sender rand', async () => {
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom()
    const { ticket, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    const senderRand = 999;

    await expect(ticketing.redeem(ticket, senderRand, redeemerRand, signature))
      .to.be.revertedWith('Hash of senderRand doesn\'t match senderRandHash');
  });

  it('can not redeem ticket with invalid redeemer rand', async () => {
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom()
    const { ticket, senderRand, signature } =
      await createWinningTicket(alice, owner);

    const redeemerRand = 999;

    await expect(ticketing.redeem(ticket, senderRand, redeemerRand, signature))
      .to.be.revertedWith('Hash of redeemerRand doesn\'t match redeemerRandHash');
  });

  it('can not redeem ticket if associated epoch does not exist', async () => {
    const alice = Wallet.createRandom()
    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    ticket.epochId = 1;

    await expect(ticketing.redeem(ticket, senderRand, redeemerRand, signature))
      .to.be.revertedWith('Ticket\'s associated epoch does not exist');
  });

  it('can not calculate winning probability if associated epoch does not exist', async () => {
    const alice = Wallet.createRandom()
    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    ticket.epochId = 1;

    await expect(ticketing.calculateWinningProbability(ticket))
      .to.be.revertedWith('Ticket\'s associated epoch does not exist');
  });

  it('can not redeem ticket if not generated during associated epoch', async () => {
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom()
    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    const updatedTicket = { ...ticket, generationBlock: 1 }

    await expect(ticketing.redeem(updatedTicket, senderRand, redeemerRand, signature))
      .to.be.revertedWith("This ticket was not generated during it\'s associated epoch");
  });

  it('can not calculate winning probablility if not generated during associated epoch', async () => {
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom()
    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    const updatedTicket = { ...ticket, generationBlock: 1 }

    await expect(ticketing.calculateWinningProbability(updatedTicket))
      .to.be.revertedWith("This ticket was not generated during it\'s associated epoch");
  });

  it('can not redeem ticket if node does not have a listing', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom()
    await ticketing.depositEscrow(toSOLOs(2000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    await expect(ticketing.redeem(ticket, senderRand, redeemerRand, signature))
      .to.be.revertedWith("Ticket redeemer must have a valid listing");
  });

  it('can not redeem ticket if node has not joined directory', async () => {
    await listings.setListing("0.0.0.0/0", 1);

    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom()
    await ticketing.depositEscrow(toSOLOs(2000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    await expect(ticketing.redeem(ticket, senderRand, redeemerRand, signature))
      .to.be.revertedWith("Ticket redeemer must have joined the directory for this epoch");
  });

  it('can not redeem ticket if node has not initialized reward pool', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await listings.setListing("0.0.0.0/0", 1);

    await directory.addManager(owner);
    await directory.joinNextDirectory(owner);

    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom()
    await ticketing.depositEscrow(toSOLOs(2000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    await expect(ticketing.redeem(ticket, senderRand, redeemerRand, signature))
      .to.be.revertedWith("Reward pool has not been initialized for the current epoch");
  });

  it('can not redeem invalid ticket', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await listings.setListing("0.0.0.0/0", 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom()
    await ticketing.depositEscrow(toSOLOs(2000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    let malformedTicket = { ...ticket };
    malformedTicket.sender = '0x0000000000000000000000000000000000000000';
    await expect(ticketing.redeem(malformedTicket, senderRand, redeemerRand, signature))
      .to.be.revertedWith("Ticket sender is null");

    malformedTicket = { ...ticket };
    malformedTicket.redeemer = '0x0000000000000000000000000000000000000000';
    await expect(ticketing.redeem(malformedTicket, senderRand, redeemerRand, signature))
      .to.be.revertedWith("Ticket redeemer is null");

    malformedTicket = { ...ticket };
    malformedTicket.senderCommit = '0x0000000000000000000000000000000000000000000000000000000000000000';
    await expect(ticketing.redeem(malformedTicket, senderRand, redeemerRand, signature))
      .to.be.revertedWith("Hash of senderRand doesn't match senderRandHash");

    malformedTicket = { ...ticket };
    malformedTicket.redeemerCommit = '0x0000000000000000000000000000000000000000000000000000000000000000';
    await expect(ticketing.redeem(malformedTicket, senderRand, redeemerRand, signature))
      .to.be.revertedWith("Hash of redeemerRand doesn't match redeemerRandHash");

    const malformedSig = '0xdebcaaaa727df04bdc990083d88ed7c8e6e9897ff18b7d968867a8bc024cbdbe10ca52eebd67a14b7b493f5c00ed9dab7b96ef62916f25afc631d336f7b2ae1e1b';
    await expect(ticketing.redeem(ticket, senderRand, redeemerRand, malformedSig))
      .to.be.revertedWith("Ticket doesn't have a valid signature");
  });

  it('rejects non winning ticket', async () => {
    // redeploy contracts with win chance of 0%
    const contracts = await utils.initializeContracts(owner, token.address, { baseLiveWinProb: 0 });
    await token.approve(contracts.stakingManager.address, toSOLOs(100000));
    await contracts.stakingManager.addStake(toSOLOs(1), owner);
    await contracts.listings.setListing("0.0.0.0/0", 1);

    await contracts.epochsManager.joinNextEpoch();

    await contracts.directory.transferOwnership(contracts.epochsManager.address);
    await contracts.epochsManager.initializeEpoch();

    await token.approve(contracts.ticketing.address, toSOLOs(100000));
    const alice = Wallet.createRandom();
    await contracts.ticketing.depositEscrow(toSOLOs(2000), alice.address);
    await contracts.ticketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner, 1);

    await utils.advanceBlock(5);

    await expect(contracts.ticketing.redeem(ticket, senderRand, redeemerRand, signature))
      .to.be.revertedWith("Ticket is not a winner");
  });

  it('can redeem winning ticket', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await listings.setListing("0.0.0.0/0", 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom()
    await ticketing.depositEscrow(toSOLOs(2000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    await ticketing.redeem(ticket, senderRand, redeemerRand, signature);

    const deposit = await ticketing.deposits(alice.address);
    assert.equal(deposit.escrow.toString(), toSOLOs(1000), 'Expected ticket payout to be substracted from escrow');
    assert.equal(deposit.penalty.toString(), toSOLOs(50), 'Expected penalty to not be changed');

    const unclaimedNodeReward = await rewardsManager.getUnclaimedNodeReward(owner);
    const unclaimedStakeReward = await rewardsManager.getUnclaimedStakeReward(owner);

    assert.equal(
      unclaimedNodeReward.add(unclaimedStakeReward).toString(),
      toSOLOs(1000),
      "Expected balance of unclaimed rewards to have added the ticket face value"
    );
  });

  it('can not redeem ticket more than once', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await listings.setListing("0.0.0.0/0", 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom()
    await ticketing.depositEscrow(toSOLOs(2000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
    await expect(ticketing.redeem(ticket, senderRand, redeemerRand, signature))
      .to.be.revertedWith("Ticket already redeemed");
  });

  it('burns penalty on insufficient escrow', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await listings.setListing("0.0.0.0/0", 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom()
    await ticketing.depositEscrow(toSOLOs(5), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    const initialTicketingBalance = await token.balanceOf(ticketing.address);

    await ticketing.redeem(ticket, senderRand, redeemerRand, signature);

    const deposit = await ticketing.deposits(alice.address);
    assert.equal(deposit.escrow.toString(), '0', 'Expected entire escrow to be used');
    assert.equal(deposit.penalty.toString(), '0', 'Expected entire penalty to be burned');

    const unclaimedNodeReward = await rewardsManager.getUnclaimedNodeReward(owner);
    const unclaimedStakeReward = await rewardsManager.getUnclaimedStakeReward(owner);

    assert.equal(
      unclaimedNodeReward.add(unclaimedStakeReward).toString(),
      '5000000000000000000',
      "Expected unclaimed balance to have added the remaining available escrow"
    );

    const ticketingBalance = await token.balanceOf(ticketing.address);
    assert.equal(
      ticketingBalance.toString(),
      initialTicketingBalance.sub((toSOLOs(55))).toString(),
      'Expected tokens from ticket contract to be removed'
    );

    const deadBalance = await token.balanceOf('0x000000000000000000000000000000000000dEaD');
    assert.equal(
      deadBalance.toString(),
      '50000000000000000000',
      'Expected dead address to receive burned tokens'
    );
  });

  it('should restake unclaimed staker rewards', async () => {
    await stakingManager.addStake(5, owner);
    await listings.setListing("0.0.0.0/0", 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom()
    await ticketing.depositEscrow(50, alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    for (let i = 0; i < 10; i++) {
      const { ticket, senderRand, redeemerRand, signature } =
        await createWinningTicket(alice, owner);
      await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
    }

    epochsManager.joinNextEpoch();

    const rewardPool = await rewardsManager.getRewardPool(
      await epochsManager.getNextEpochId(),
      owner
    );

    const unclaimedStakeReward = await rewardsManager.getUnclaimedStakeReward(owner);

    // check the total active stake for the next epoch includes the unclaimed rewards plus
    // the managed stake
    assert.equal(
      rewardPool.totalActiveStake.toString(),
      unclaimedStakeReward.add(5).toString(),
      "Expected total active stake for next reward pool to include unclaimed rewards"
    );
  });

  it('fails to to claim non existent rewards', async () => {
    await expect(rewardsManager.claimStakingRewards(owner))
      .to.be.revertedWith("Nothing to claim");
    await expect(rewardsManager.claimNodeRewards())
      .to.be.revertedWith("Nothing to claim");
  });

  it('can claim ticketing rewards', async () => {
    await stakingManager.addStake(web3.utils.toWei('1'), owner);
    await listings.setListing("0.0.0.0/0", 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom()
    await ticketing.depositEscrow(toSOLOs(50000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    for (let i = 0; i < 10; i++) {
      const { ticket, senderRand, redeemerRand, signature } =
        await createWinningTicket(alice, owner);

      await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
    }

    const initialBalance = await token.balanceOf(owner);

    await rewardsManager.claimNodeRewards();
    await rewardsManager.claimStakingRewards(owner);

    const postBalance = await token.balanceOf(owner);
    // Expect the node have the entire reward balance added to their account
    const expectedPostBalance = initialBalance.add(toSOLOs(10000));

    compareExpectedBalance(expectedPostBalance, postBalance);

    const unclaimedNodeReward = await rewardsManager.getUnclaimedNodeReward(owner);
    const unclaimedStakeReward = await rewardsManager.getUnclaimedStakeReward(owner);

    compareExpectedBalance(
      unclaimedNodeReward.add(unclaimedStakeReward),
      0
    );
  });

  it('delegated stakers should be able to claim rewards', async () => {
    for (let i = 2; i < 4; i++) {
      const account = await accounts[i].getAddress();
      await token.transfer(account, toSOLOs(1000));
      await token.connect(accounts[i]).approve(stakingManager.address, toSOLOs(1000));
    }

    await stakingManager.addStake(toSOLOs(3), owner);
    await listings.setListing("0.0.0.0/0", 1);

    // have account 2 as a delegated staker
    await stakingManager.connect(accounts[2]).addStake(toSOLOs(2), owner);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom()
    await ticketing.depositEscrow(toSOLOs(50000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    for (let i = 0; i < 10; i++) {
      const { ticket, senderRand, redeemerRand, signature } =
        await createWinningTicket(alice, owner);

      await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
    }

    const totalStakersReward = await rewardsManager.getRewardPoolStakersTotal(1, owner);
    assert.equal(
      totalStakersReward.toString(),
      toSOLOs(5000),
      "Expected correct amount of reward to be allocated to stakers"
    );

    const initialDelegatorTwoBalance = await token.balanceOf(await accounts[0].getAddress());
    const initialDelegatorThreeBalance = await token.balanceOf(await accounts[2].getAddress());

    await rewardsManager.claimStakingRewards(owner);
    await rewardsManager.connect(accounts[2]).claimStakingRewards(owner);

    // The stakers reward total is 5000 SYLOs. Account 0 owns 66% of the stake, and account 2 owns
    // 33% of the stake, the splits should be 3/5 * 5000 and 2/5 * 5000

    const postDelegatorTwoBalance = await token.balanceOf(await accounts[0].getAddress());
    const expectedPostDelegatorTwoBalance = initialDelegatorTwoBalance.add(toSOLOs(3 / 5 * 5000));
    compareExpectedBalance(expectedPostDelegatorTwoBalance, postDelegatorTwoBalance);

    const postDelegatorThreeBalance = await token.balanceOf(await accounts[2].getAddress());
    const expectedPostDelegatorThreeBalance = initialDelegatorThreeBalance.add(toSOLOs(2 / 5 * 5000));
    compareExpectedBalance(expectedPostDelegatorThreeBalance, postDelegatorThreeBalance);

  });

  it('should have rewards be automatically claimed when stake is updated', async () => {
    await token.transfer(await accounts[2].getAddress(), 1000 );
    await token.connect(accounts[2]).approve(stakingManager.address, toSOLOs(1000));

    await stakingManager.addStake(toSOLOs(3), owner);
    await listings.setListing("0.0.0.0/0", 1);

    // have account 2 as a delegated staker
    await stakingManager.connect(accounts[2]).addStake(toSOLOs(1), owner);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom()
    await ticketing.depositEscrow(toSOLOs(50000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    for (let i = 0; i < 10; i++) {
      const { ticket, senderRand, redeemerRand, signature } =
        await createWinningTicket(alice, owner);

      await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
    }

    // add more stake
    await stakingManager.connect(accounts[2]).addStake(toSOLOs(1), owner);

    const claimAfterAddingStake = await rewardsManager.calculateStakerClaim(owner, await accounts[2].getAddress());

    assert.equal(
      claimAfterAddingStake.toString(),
      '0',
      "Expected reward to be automatically claimed after adding stake"
    );

    for (let i = 0; i < 10; i++) {
      const { ticket, senderRand, redeemerRand, signature } =
        await createWinningTicket(alice, owner);

      await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
    }

    // remove some stake
    await stakingManager.connect(accounts[2]).unlockStake(1, owner);

    const claimAfterRemovingStake = await rewardsManager.calculateStakerClaim(owner, await accounts[2].getAddress());

    assert.equal(
      claimAfterRemovingStake.toString(),
      '0',
      "Expected reward to be automatically claimed after adding stake"
    );
  });

  it('can not claim reward more than once', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await listings.setListing("0.0.0.0/0", 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom()
    await ticketing.depositEscrow(toSOLOs(50000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, senderRand, redeemerRand, signature } =
      await createWinningTicket(alice, owner);

    await ticketing.redeem(ticket, senderRand, redeemerRand, signature);

    await rewardsManager.claimStakingRewards(owner);

    const lastClaim = await rewardsManager.getLastClaim(owner, owner);
    expect(lastClaim).to.be.above(0);

    await expect(rewardsManager.claimStakingRewards(owner))
      .to.be.revertedWith("Nothing to claim");
  });

  it('should be able to correctly calculate staking rewards for multiple epochs when managed stake is the same', async () => {
    for (let i = 2; i < 5; i++) {
      await token.transfer(await accounts[i].getAddress(), toSOLOs(1000) );
      await token.connect(accounts[i]).approve(stakingManager.address, toSOLOs(1000));
    }

    await stakingManager.addStake(toSOLOs(1000), owner);
    await listings.setListing("0.0.0.0/0", 1);

    // have accounts 2, 3 and 4 as delegated stakers with varying levels of stake
    await stakingManager.connect(accounts[2]).addStake(toSOLOs(250), owner);
    await stakingManager.connect(accounts[3]).addStake(toSOLOs(400), owner);
    await stakingManager.connect(accounts[4]).addStake(toSOLOs(350), owner);

    const alice = Wallet.createRandom()
    await ticketing.depositEscrow(toSOLOs(500000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    for (let j = 0; j < 3; j++) {
      await epochsManager.joinNextEpoch();
      await epochsManager.initializeEpoch();

      // 500 is added to the stakers reward total on each redemption (50% of 1000)
      for (let i = 0 ; i < 6; i++) {
        const { ticket, senderRand, redeemerRand, signature } =
          await createWinningTicket(alice, owner);

        await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
      }
    }

    const unclaimedStakeReward = await rewardsManager.getUnclaimedStakeReward(owner);

    // the total unclaimed stake reward should 3 * 6 * 500 = 9000
    compareExpectedBalance(toSOLOs(9000), unclaimedStakeReward);

    // verify each staker will receive the correct amount of reward if they were to claim now
    const stakerClaimTwo = await rewardsManager.calculateStakerClaim(owner, await accounts[2].getAddress());
    const expectedStakerClaimTwo = toSOLOs(9000 * 0.125);
    compareExpectedBalance(expectedStakerClaimTwo, stakerClaimTwo);

    const stakerClaimThree = await rewardsManager.calculateStakerClaim(owner, await accounts[3].getAddress());
    const expectedStakerClaimThree = toSOLOs(9000 * 0.2);
    compareExpectedBalance(expectedStakerClaimThree, stakerClaimThree);

    const stakerClaimFour = await rewardsManager.calculateStakerClaim(owner, await accounts[4].getAddress());
    const expectedStakerClaimFour = toSOLOs(9000 * 0.175);
    compareExpectedBalance(expectedStakerClaimFour, stakerClaimFour);

    // ensure each staker is actually able to claim
    for (let i = 2; i < 5; i++) {
      await rewardsManager.connect(accounts[i]).claimStakingRewards(owner);
    }
  });

  it('should be able to correctly calculate staking rewards for multiple epochs when managed stake increases', async () => {
    for (let i = 2; i < 5; i++) {
      await token.transfer(await accounts[i].getAddress(), toSOLOs(1000));
      await token.connect(accounts[i]).approve(stakingManager.address, toSOLOs(1000));
    }

    await stakingManager.addStake(toSOLOs(1000), owner);
    await listings.setListing("0.0.0.0/0", 1);

    // have account 2, 3 and 4 as delegated stakers with varying levels of stake
    await stakingManager.connect(accounts[2]).addStake(toSOLOs(250), owner);
    await stakingManager.connect(accounts[3]).addStake(toSOLOs(400), owner);
    await stakingManager.connect(accounts[4]).addStake(toSOLOs(350), owner);

    const alice = Wallet.createRandom()
    await ticketing.depositEscrow(toSOLOs(50000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    // have account 5 add stake midway through
    for (let j = 0; j < 3; j++) {
      if (j == 1) {
        await token.transfer(await accounts[5].getAddress(), toSOLOs(1000));
        await token.connect(accounts[5]).approve(stakingManager.address, toSOLOs(1000));
        // their stake will be active in the next round
        await stakingManager.connect(accounts[5]).addStake(toSOLOs(500), owner);
      }

      await epochsManager.joinNextEpoch();
      await epochsManager.initializeEpoch();

      // 7.5 is added to the stakers reward total on each redemption (50% of 15)
      for (let i = 0 ; i < 6; i++) {
        const { ticket, senderRand, redeemerRand, signature } =
          await createWinningTicket(alice, owner);

        await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
      }
    }

    const epochTwoActiveStake = await rewardsManager.getRewardPoolActiveStake(2, owner);

    // account 5's reward should be the sum of the rewards gained in both epoch 2 and 3
    // multiplied by the proportion of the stake held when their stake became active
    const stakeClaimFive = await rewardsManager.calculateStakerClaim(owner, await accounts[5].getAddress());
    const s = BigNumber.from(toSOLOs(500)); // initial stake
    const r = toSOLOs(2 * 6 * 500); // accumulated reward
    const expectedStakeClaimFive = s.mul(r).div(epochTwoActiveStake);
    compareExpectedBalance(expectedStakeClaimFive, stakeClaimFive);

    // for accounts 2, 3, and 4, the total managed stake that becomes active
    // changes from epoch 2, thus to calculate the expected reward, we need
    // to caluculate the expected reward for epoch 1 using different stake proportions
    // than for epochs 2 and 3
    for (let i = 2; i < 5; i++) {
      const initialStake = await stakingManager.getCurrentStakerAmount(owner, await accounts[i].getAddress());
      const epochOneActiveStake = await rewardsManager.getRewardPoolActiveStake(1, owner);
      const epochOneReward = BigNumber.from(toSOLOs(6 * 500)).mul(initialStake).div(epochOneActiveStake);

      const stakeAtEpochTwo = initialStake.add(epochOneReward);
      const remainingReward = stakeAtEpochTwo.mul(toSOLOs(2 * 6 * 500)).div(epochTwoActiveStake);

      const totalExpectedReward = epochOneReward.add(remainingReward);
      const stakerClaim = await rewardsManager.calculateStakerClaim(owner, await accounts[i].getAddress());

      compareExpectedBalance(totalExpectedReward, stakerClaim);
    }

    // ensure each staker is actually able to claim
    for (let i = 2; i < 6; i++) {
      await rewardsManager.connect(accounts[i]).claimStakingRewards(owner);
    }
  });

  it('should be able to correctly calculate staking rewards for multiple epochs when managed stake decreases', async () => {
    for (let i = 2; i < 5; i++) {
      await token.transfer(await accounts[i].getAddress(), toSOLOs(1000));
      await token.connect(accounts[i]).approve(stakingManager.address, toSOLOs(1000));
    }

    await stakingManager.addStake(toSOLOs(1000), owner);
    await listings.setListing("0.0.0.0/0", 1);

    // have account 2, 3 and 4 as delegated stakers with varying levels of stake
    await stakingManager.connect(accounts[2]).addStake(toSOLOs(250), owner);
    await stakingManager.connect(accounts[3]).addStake(toSOLOs(400), owner);
    await stakingManager.connect(accounts[4]).addStake(toSOLOs(350), owner);

    const alice = Wallet.createRandom()
    await ticketing.depositEscrow(toSOLOs(50000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    // have account 2 unlock stake midway through
    for (let j = 0; j < 3; j++) {
      if (j == 1) {
        await stakingManager.connect(accounts[2]).unlockStake(toSOLOs(250), owner);
      }

      await epochsManager.joinNextEpoch();
      await epochsManager.initializeEpoch();

      // 7 is added to the stakers reward total on each redemption (50% of 15)
      for (let i = 0 ; i < 6; i++) {
        const { ticket, senderRand, redeemerRand, signature } =
          await createWinningTicket(alice, owner);

        await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
      }
    }

    const epochTwoActiveStake = await rewardsManager.getRewardPoolActiveStake(2, owner);

    // for accounts 3, and 4, the total managed stake that becomes active
    // changes from epoch 2, thus to calculate the expected reward, we need
    // to caluclate the expected reward for epoch 1 using different stake proportions
    // than for epochs 2 and 3
    for (let i = 3; i < 5; i++) {
      const initialStake = await stakingManager.getCurrentStakerAmount(owner, await accounts[i].getAddress());
      const epochOneActiveStake = await rewardsManager.getRewardPoolActiveStake(1, owner);
      const epochOneReward = BigNumber.from(toSOLOs(6 * 500)).mul(initialStake).div(epochOneActiveStake);

      const stakeAtEpochTwo = initialStake.add(epochOneReward);
      const remainingReward = stakeAtEpochTwo.mul(toSOLOs(2 * 6 * 500)).div(epochTwoActiveStake);

      const totalExpectedReward = epochOneReward.add(remainingReward);
      const stakerClaim = await rewardsManager.calculateStakerClaim(owner, await accounts[i].getAddress());

      compareExpectedBalance(totalExpectedReward, stakerClaim);
    }

    // ensure each staker is actually able to claim
    for (let i = 3; i < 5; i++) {
      await rewardsManager.connect(accounts[i]).claimStakingRewards(owner);
    }
  });

  // XXX: Set to skip as it is a very long test and sometimes breaks the local
  // truffle test network/client. However this should be manually run if any significant changes
  // to the Rewards contract calculation is made.
  // TODO: Create script to spin up new test network to run this test locally or for CI automatically.
  it('should calculate updated stake and rewards over several ticket redemptions without significant precision loss [ @skip-on-coverage ]', async () => {
    for (let i = 2; i < 5; i++) {
      await token.transfer(await accounts[i].getAddress(), toSOLOs(1000));
      await token.connect(accounts[i]).approve(stakingManager.address, toSOLOs(1000));
    }

    await stakingManager.addStake(toSOLOs(1000), owner);
    await listings.setListing("0.0.0.0/0", 1);

    // have account 2, 3 and 4 as delegated stakers with varying levels of stake
    await stakingManager.connect(accounts[2]).addStake(toSOLOs(250), owner);
    await stakingManager.connect(accounts[3]).addStake(toSOLOs(400), owner);
    await stakingManager.connect(accounts[4]).addStake(toSOLOs(350), owner);

    const alice = Wallet.createRandom()
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

    const stakerClaimTwo = await rewardsManager.calculateStakerClaim(owner, await accounts[2].getAddress());
    const expectedStakerClaimTwo = toSOLOs(iterations * 500 * 0.125);
    compareExpectedBalance(expectedStakerClaimTwo, stakerClaimTwo);

    const stakerClaimThree = await rewardsManager.calculateStakerClaim(owner, await accounts[3].getAddress());
    const expectedStakerClaimThree = toSOLOs(iterations * 500 * 0.2);
    compareExpectedBalance(expectedStakerClaimThree, stakerClaimThree);

    const stakerClaimFour = await rewardsManager.calculateStakerClaim(owner, await accounts[4].getAddress());
    const expectedStakerClaimFour = toSOLOs(iterations * 500 * 0.175);
    compareExpectedBalance(expectedStakerClaimFour, stakerClaimFour);
  }).timeout(0);

  it('should decay winning probability as ticket approaches expiry', async () => {
    // deploy another ticketing contract with simpler parameters
    const contracts = await utils.initializeContracts(
      owner,
      token.address,
      { faceValue,
        baseLiveWinProb: 100000,
        expiredWinProb: 1000,
        decayRate: 8000,
        ticketDuration: 100
      });
    epochsManager = contracts.epochsManager;
    rewardsManager = contracts.rewardsManager;
    ticketing = contracts.ticketing;
    directory = contracts.directory;
    listings = contracts.listings;
    stakingManager = contracts.stakingManager;

    await directory.transferOwnership(epochsManager.address);
    await rewardsManager.addManager(ticketing.address);
    await rewardsManager.addManager(epochsManager.address);

    await token.approve(ticketing.address, toSOLOs(10000));
    await token.approve(stakingManager.address, toSOLOs(10000));

    await stakingManager.addStake(toSOLOs(1), owner);
    await listings.setListing("0.0.0.0/0", 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom()
    await ticketing.depositEscrow(toSOLOs(5000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket } =
      await createWinningTicket(alice, owner);

    // advance the block halfway to ticket expiry
    await utils.advanceBlock(51);

    // check if the probability has decayed 50% of the maximum decayed value (80%)
    const expectedProbability = 100000 - (0.5 * 0.8 * 100000);

    const decayedProbability = await ticketing.calculateWinningProbability(
      ticket
    );

    assert.equal(
      decayedProbability.toString(),
      expectedProbability.toString(),
      "Expected probablity of ticket winning to decay"
    );
  });

  it('should be able to correctly calculate staker rewards if node was not active for multiple epochs', async () => {
    for (let i = 2; i < 5; i++) {
      await token.transfer(await accounts[i].getAddress(), toSOLOs(1000) );
      await token.connect(accounts[i]).approve(stakingManager.address, toSOLOs(1000));
    }

    await stakingManager.addStake(toSOLOs(1000), owner);
    await listings.setListing("0.0.0.0/0", 1);

    // have accounts 2, 3 and 4 as delegated stakers with varying levels of stake
    await stakingManager.connect(accounts[2]).addStake(toSOLOs(250), owner);
    await stakingManager.connect(accounts[3]).addStake(toSOLOs(400), owner);
    await stakingManager.connect(accounts[4]).addStake(toSOLOs(350), owner);

    const alice = Wallet.createRandom()
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
      for (let i = 0 ; i < 6; i++) {
        const { ticket, senderRand, redeemerRand, signature } =
          await createWinningTicket(alice, owner);

        await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
      }
    }

    const unclaimedStakeReward = await rewardsManager.getUnclaimedStakeReward(owner);

    // the total unclaimed stake reward should 3 * 6 * 500 = 9000
    compareExpectedBalance(toSOLOs(9000), unclaimedStakeReward);

    // verify each staker will receive the correct amount of reward if they were to claim now
    const stakerClaimTwo = await rewardsManager.calculateStakerClaim(owner, await accounts[2].getAddress());
    const expectedStakerClaimTwo = toSOLOs(9000 * 0.125);
    compareExpectedBalance(expectedStakerClaimTwo, stakerClaimTwo);

    const stakerClaimThree = await rewardsManager.calculateStakerClaim(owner, await accounts[3].getAddress());
    const expectedStakerClaimThree = toSOLOs(9000 * 0.2);
    compareExpectedBalance(expectedStakerClaimThree, stakerClaimThree);

    const stakerClaimFour = await rewardsManager.calculateStakerClaim(owner, await accounts[4].getAddress());
    const expectedStakerClaimFour = toSOLOs(9000 * 0.175);
    compareExpectedBalance(expectedStakerClaimFour, stakerClaimFour);

    // ensure each staker is actually able to claim
    for (let i = 2; i < 5; i++) {
      await rewardsManager.connect(accounts[i]).claimStakingRewards(owner);
    }
  });

  it('allows node to claim remaining unclaimed rewards', async () => {
    await token.transfer(await accounts[2].getAddress(), toSOLOs(1000));
    await token.connect(accounts[2]).approve(stakingManager.address, toSOLOs(1000));

    await stakingManager.addStake(toSOLOs(3), owner);
    await listings.setListing("0.0.0.0/0", 1);

    // have account 2 as a delegated staker
    await stakingManager.connect(accounts[2]).addStake(toSOLOs(1), owner);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom()
    await ticketing.depositEscrow(toSOLOs(50000), alice.address);
    await ticketing.depositPenalty(toSOLOs(50), alice.address);

    for (let i = 0; i < 10; i++) {
      const { ticket, senderRand, redeemerRand, signature } =
        await createWinningTicket(alice, owner);

      await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
    }

    // all stakers have unstaked and have automatically claimed their rewards
    await stakingManager.unlockStake(toSOLOs(3), owner);
    await stakingManager.connect(accounts[2]).unlockStake(toSOLOs(1), owner);

    // all tickets redeemed at this point should just have entire face value given to node
    for (let i = 0; i < 10; i++) {
      const { ticket, senderRand, redeemerRand, signature } =
        await createWinningTicket(alice, owner);

      await ticketing.redeem(ticket, senderRand, redeemerRand, signature);
    }

    const balanceBeforeClaim = await token.balanceOf(owner);

    await rewardsManager.claimNodeRewards();

    const balanceAfterClaim = await token.balanceOf(owner);

    assert.equal(
      balanceAfterClaim.toString(),
      balanceBeforeClaim.add(toSOLOs(15000)).toString(),
      "Expected node to be able to claim remaining rewards"
    );
  });

  it('returns 0 winning probability if ticket has expired', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await listings.setListing("0.0.0.0/0", 1);

    const alice = Wallet.createRandom()

    await epochsManager.initializeEpoch();

    const { ticket } =
      await createWinningTicket(alice, owner);

    // advance the block all the way to ticket expiry
    await utils.advanceBlock(21);

    const p = await ticketing.calculateWinningProbability(
      ticket
    );

    assert.equal(
      '0',
      p.toString(),
      'Expected probability to be 0'
    );
  });

  it('simulates scenario between sender, node, and oracle', async () => {
    const sender = Wallet.createRandom();
    const node = owner;

    // set up the node's stake and listing
    await stakingManager.addStake(toSOLOs(1), node);
    await listings.setListing("0.0.0.0/0", 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    // set up the sender's escrow
    await ticketing.depositEscrow(50, sender.address);
    await ticketing.depositPenalty(toSOLOs(50), sender.address);

    // have the node and sender generate random numbers
    const nodeRand = crypto.randomBytes(32);
    const senderRand = crypto.randomBytes(32);

    // create commits from those random numbers
    const nodeCommit = soliditySha3(nodeRand)!;
    const senderCommit = soliditySha3(senderRand)!;

    const epochId = await epochsManager.getCurrentActiveEpoch().then(e => e.iteration);

    // create the ticket to be given to the node
    const ticket = {
      epochId,
      sender: sender.address,
      redeemer: node,
      generationBlock: (await ethers.provider.getBlockNumber()) + 1,
      senderCommit,
      redeemerCommit: nodeCommit
    };

    // have sender sign the hash of the ticket
    const ticketHash = await ticketing.getTicketHash(ticket);
    const signature = eth.Account.sign(ticketHash, sender.privateKey);

    // establish the oracle
    const oracle = sodium.crypto_sign_keypair('uint8array');

    // encrypt senderRandom to create the key
    const key = sodium.crypto_box_seal(
      senderRand,
      sodium.crypto_sign_ed25519_pk_to_curve25519(oracle.publicKey),
      'uint8array'
    );

    // have oracle decrypt the key and reveal the random number to the node
    const revealedSenderRand = sodium.crypto_box_seal_open(
      key,
      sodium.crypto_sign_ed25519_pk_to_curve25519(oracle.publicKey),
      sodium.crypto_sign_ed25519_sk_to_curve25519(oracle.privateKey),
      'uint8array'
    );

    // once secret has been revealed, the node can now redeem the ticket
    await ticketing.redeem(
      ticket,
      revealedSenderRand,
      nodeRand,
      signature,
      { from: node }
    );
  });

  // This test suite relies on confirming updated stakes and rewards are correctly
  // calculated after incrementing the reward pool. However due to minor precision loss, the
  // actual balance may slightly differ.
  // This function checks the difference falls within a small fraction of a single SYLO.
  function compareExpectedBalance(a: BigNumberish, b: BigNumberish) {
    const diff = BigNumber.from(a).sub(BigNumber.from(b));
    // NOTE: This essentially says that a margin of 10**4 SOLOs is acceptable, or
    // 0.00000000000001 SYLOs
    expect(diff.toNumber()).to.be.within(0, 10**4);
  }

  function toSOLOs(a: number): string {
    return web3.utils.toWei(a.toString());
  }

  async function createWinningTicket(sender: Wallet, redeemer: string, epochId?: number) {
    const senderRand = 1;
    const senderCommit = soliditySha3(senderRand)!;

    const redeemerRand = 1;
    const redeemerCommit = soliditySha3(redeemerRand)!;

    const generationBlock = await ethers.provider.getBlockNumber();

    const ticket = {
      epochId: epochId ?? await epochsManager.currentIteration(),
      sender: sender.address,
      redeemer,
      generationBlock: BigNumber.from(generationBlock + 1),
      senderCommit,
      redeemerCommit
    };

    const ticketHash = await ticketing.getTicketHash(ticket);

    const signature = eth.Account.sign(ticketHash, sender.privateKey);

    return { ticket, senderRand, redeemerRand, signature, ticketHash }
  }

});
