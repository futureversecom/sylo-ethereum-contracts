import { ethers } from 'hardhat';
import { assert, expect } from 'chai';
import { Signer, Wallet } from 'ethers';
import { StandardMerkleTree } from '@openzeppelin/merkle-tree';
import {
  EpochsManager,
  Registries,
  RewardsManager,
  StakingManager,
  SyloTicketing,
  SyloToken,
  TestSeekers,
  TicketingParameters,
} from '../../typechain-types';
import {
  toSOLOs,
  setSeekerRegistry,
  createWinningMultiReceiverTicket,
} from './utils';
import utils from '../utils';

describe('MultiReceiverTicketing', () => {
  let accounts: Signer[];
  let owner: string;

  const faceValue = toSOLOs(1000);
  const epochDuration = 1;

  let token: SyloToken;
  let epochsManager: EpochsManager;
  let rewardsManager: RewardsManager;
  let syloTicketing: SyloTicketing;
  let ticketingParameters: TicketingParameters;
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
    const contracts = await utils.initializeContracts(owner, token, {
      faceValue,
      epochDuration,
    });
    epochsManager = contracts.epochsManager;
    rewardsManager = contracts.rewardsManager;
    syloTicketing = contracts.syloTicketing;
    ticketingParameters = contracts.ticketingParameters;
    registries = contracts.registries;
    stakingManager = contracts.stakingManager;
    seekers = contracts.seekers;

    await token.approve(await stakingManager.getAddress(), toSOLOs(10000000));
    await token.approve(await syloTicketing.getAddress(), toSOLOs(10000000));
  });

  it('can redeem multi receiver ticket', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bobs = Array(5)
      .fill(0)
      .map(_ => Wallet.createRandom());

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const tree = StandardMerkleTree.of(
      bobs.map(b => [b.address]),
      ['address'],
    );

    const { ticket, redeemerRand, senderSig, ticketHash } =
      await createWinningMultiReceiverTicket(
        syloTicketing,
        epochsManager,
        alice,
        tree.root,
        owner,
      );

    const receiverSig = await bobs[0].signMessage(ethers.getBytes(ticketHash));
    const proof = tree.getProof([bobs[0].address]);

    await syloTicketing.redeemMultiReceiver(
      ticket,
      redeemerRand,
      { main: bobs[0].address, delegated: ethers.ZeroAddress },
      proof,
      senderSig,
      receiverSig,
    );

    const deposit = await syloTicketing.deposits(alice.address);
    assert.equal(
      deposit.escrow,
      toSOLOs(1000),
      'Expected ticket payout to be substracted from escrow',
    );

    const pendingReward = await rewardsManager.getPendingRewards(owner);

    assert.equal(
      pendingReward,
      toSOLOs(500),
      'Expected balance of pending rewards to have added the ticket face value',
    );
  });

  it('can redeem same multi receiver ticket for multiple different receivers', async () => {
    // setting a large ticket duration prevents the tickets from expiring
    await ticketingParameters.setTicketDuration(10000000000);

    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bobs = Array(5)
      .fill(0)
      .map(_ => Wallet.createRandom());

    await syloTicketing.depositEscrow(toSOLOs(20000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const tree = StandardMerkleTree.of(
      bobs.map(b => [b.address]),
      ['address'],
    );

    const { ticket, redeemerRand, senderSig, ticketHash } =
      await createWinningMultiReceiverTicket(
        syloTicketing,
        epochsManager,
        alice,
        tree.root,
        owner,
      );

    let previousDeposit = await syloTicketing.deposits(alice.address);
    let previousReward = await rewardsManager.getPendingRewards(owner);

    for (const bob of bobs) {
      const receiverSig = await bob.signMessage(ethers.getBytes(ticketHash));
      const proof = tree.getProof([bob.address]);

      await syloTicketing.redeemMultiReceiver(
        ticket,
        redeemerRand,
        { main: bob.address, delegated: ethers.ZeroAddress },
        proof,
        senderSig,
        receiverSig,
      );

      const deposit = await syloTicketing.deposits(alice.address);
      assert.equal(
        deposit.escrow,
        previousDeposit.escrow - toSOLOs(1000),
        'Expected ticket payout to be substracted from escrow',
      );
      previousDeposit = deposit;

      const pendingReward = await rewardsManager.getPendingRewards(owner);
      assert.equal(
        pendingReward,
        previousReward + toSOLOs(500),
        'Expected balance of pending rewards to have added the ticket face value',
      );
      previousReward = pendingReward;
    }
  });

  it('can not redeem invalid ticket', async () => {
    const alice = Wallet.createRandom();
    const bobs = Array(5)
      .fill(0)
      .map(_ => Wallet.createRandom());

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const tree = StandardMerkleTree.of(
      bobs.map(b => [b.address]),
      ['address'],
    );

    const { ticket, redeemerRand, senderSig, ticketHash } =
      await createWinningMultiReceiverTicket(
        syloTicketing,
        epochsManager,
        alice,
        tree.root,
        owner,
      );

    const receiverSig = await bobs[0].signMessage(ethers.getBytes(ticketHash));
    const proof = tree.getProof([bobs[0].address]);

    await expect(
      syloTicketing.redeemMultiReceiver(
        { ...ticket, sender: { ...ticket.sender, main: ethers.ZeroAddress } },
        redeemerRand,
        { main: bobs[0].address, delegated: ethers.ZeroAddress },
        proof,
        senderSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'TicketSenderCannotBeZeroAddress',
    );

    await expect(
      syloTicketing.redeemMultiReceiver(
        ticket,
        redeemerRand,
        { main: ethers.ZeroAddress, delegated: ethers.ZeroAddress },
        proof,
        senderSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'TicketReceiverCannotBeZeroAddress',
    );

    await expect(
      syloTicketing.redeemMultiReceiver(
        { ...ticket, redeemer: ethers.ZeroAddress },
        redeemerRand,
        { main: bobs[0].address, delegated: ethers.ZeroAddress },
        proof,
        senderSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'TicketRedeemerCannotBeZeroAddress',
    );

    await expect(
      syloTicketing.redeemMultiReceiver(
        {
          ...ticket,
          redeemerCommit:
            '0x0000000000000000000000000000000000000000000000000000000000000000',
        },
        redeemerRand,
        { main: bobs[0].address, delegated: ethers.ZeroAddress },
        proof,
        senderSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(syloTicketing, 'RedeemerCommitMismatch');

    await expect(
      syloTicketing.redeemMultiReceiver(
        {
          ...ticket,
          sender: {
            ...ticket.sender,
            delegated: Wallet.createRandom().address,
          },
        },
        redeemerRand,
        { main: bobs[0].address, delegated: ethers.ZeroAddress },
        proof,
        senderSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'InvalidSenderTicketSigningPermission',
    );

    await expect(
      syloTicketing.redeemMultiReceiver(
        ticket,
        redeemerRand,
        { main: bobs[0].address, delegated: Wallet.createRandom().address },
        proof,
        senderSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'InvalidReceiverTicketSigningPermission',
    );

    const malformedSig =
      '0xdebcaaaa727df04bdc990083d88ed7c8e6e9897ff18b7d968867a8bc024cbdbe10ca52eebd67a14b7b493f5c00ed9dab7b96ef62916f25afc631d336f7b2ae1e1b';
    await expect(
      syloTicketing.redeemMultiReceiver(
        ticket,
        redeemerRand,
        { main: bobs[0].address, delegated: ethers.ZeroAddress },
        proof,
        malformedSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(syloTicketing, 'InvalidSenderSignature');

    await expect(
      syloTicketing.redeemMultiReceiver(
        ticket,
        redeemerRand,
        { main: bobs[0].address, delegated: ethers.ZeroAddress },
        proof,
        senderSig,
        malformedSig,
      ),
    ).to.be.revertedWithCustomError(syloTicketing, 'InvalidReceiverSignature');
  });

  it('can not redeem non winning ticket', async () => {
    await ticketingParameters.setBaseLiveWinProb(0);

    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bobs = Array(5)
      .fill(0)
      .map(_ => Wallet.createRandom());

    const tree = StandardMerkleTree.of(
      bobs.map(b => [b.address]),
      ['address'],
    );

    const { ticket, redeemerRand, senderSig, ticketHash } =
      await createWinningMultiReceiverTicket(
        syloTicketing,
        epochsManager,
        alice,
        tree.root,
        owner,
      );

    const receiverSig = await bobs[0].signMessage(ethers.getBytes(ticketHash));
    const proof = tree.getProof([bobs[0].address]);

    await expect(
      syloTicketing.redeemMultiReceiver(
        ticket,
        redeemerRand,
        { main: bobs[0].address, delegated: ethers.ZeroAddress },
        proof,
        senderSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(syloTicketing, 'TicketNotWinning');
  });

  it('can not redeem ticket for future block', async () => {
    const alice = Wallet.createRandom();
    const bobs = Array(5)
      .fill(0)
      .map(_ => Wallet.createRandom());

    const tree = StandardMerkleTree.of(
      bobs.map(b => [b.address]),
      ['address'],
    );

    const { ticket, redeemerRand, senderSig } =
      await createWinningMultiReceiverTicket(
        syloTicketing,
        epochsManager,
        alice,
        tree.root,
        owner,
      );

    await expect(
      syloTicketing.redeemMultiReceiver(
        {
          ...ticket,
          generationBlock: BigInt(ticket.generationBlock) + 10n,
        },
        redeemerRand,
        { main: bobs[0].address, delegated: ethers.ZeroAddress },
        [],
        senderSig,
        new Uint8Array(0),
      ),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'TicketCannotBeFromFutureBlock',
    );
  });

  it('cannot redeem multi receiver ticket if node has not joined epoch', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bobs = Array(5)
      .fill(0)
      .map(_ => Wallet.createRandom());

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const tree = StandardMerkleTree.of(
      bobs.map(b => [b.address]),
      ['address'],
    );

    const { ticket, redeemerRand, senderSig, ticketHash } =
      await createWinningMultiReceiverTicket(
        syloTicketing,
        epochsManager,
        alice,
        tree.root,
        owner,
      );

    const receiverSig = await bobs[0].signMessage(ethers.getBytes(ticketHash));
    const proof = tree.getProof([bobs[0].address]);

    await expect(
      syloTicketing.redeemMultiReceiver(
        ticket,
        redeemerRand,
        { main: bobs[0].address, delegated: ethers.ZeroAddress },
        proof,
        senderSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'RedeemerMustHaveJoinedEpoch',
    );
  });

  it('can not redeem for non valid receiver', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bobs = Array(5)
      .fill(0)
      .map(_ => Wallet.createRandom());

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const tree = StandardMerkleTree.of(
      bobs.map(b => [b.address]),
      ['address'],
    );

    const { ticket, redeemerRand, senderSig, ticketHash } =
      await createWinningMultiReceiverTicket(
        syloTicketing,
        epochsManager,
        alice,
        tree.root,
        owner,
      );

    const invalidReceiver = Wallet.createRandom();
    // we use a proof from a valid receiver but supply an invalid receiver
    const receiverSig = await invalidReceiver.signMessage(
      ethers.getBytes(ticketHash),
    );
    const proof = tree.getProof([bobs[0].address]);

    await expect(
      syloTicketing.redeemMultiReceiver(
        ticket,
        redeemerRand,
        { main: invalidReceiver.address, delegated: ethers.ZeroAddress },
        proof,
        senderSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(syloTicketing, 'InvalidMerkleProof');
  });

  it('can not redeem for the same user more than once', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bobs = Array(5)
      .fill(0)
      .map(_ => Wallet.createRandom());

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const tree = StandardMerkleTree.of(
      bobs.map(b => [b.address]),
      ['address'],
    );

    const { ticket, redeemerRand, senderSig, ticketHash } =
      await createWinningMultiReceiverTicket(
        syloTicketing,
        epochsManager,
        alice,
        tree.root,
        owner,
      );

    const receiverSig = await bobs[0].signMessage(ethers.getBytes(ticketHash));
    const proof = tree.getProof([bobs[0].address]);

    await syloTicketing.redeemMultiReceiver(
      ticket,
      redeemerRand,
      { main: bobs[0].address, delegated: ethers.ZeroAddress },
      proof,
      senderSig,
      receiverSig,
    );

    await expect(
      syloTicketing.redeemMultiReceiver(
        ticket,
        redeemerRand,
        { main: bobs[0].address, delegated: ethers.ZeroAddress },
        proof,
        senderSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(syloTicketing, 'TicketAlreadyRedeemed');
  });

  it('can redeem a ticket if receiver is very large [ @skip-on-coverage ]', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(seekers, registries, accounts[0], accounts[1], 1);

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    // create 100_000 receivers
    const bobs = Array(10000)
      .fill(0)
      .map((_, i) => {
        if (i > 999 && i % 1000 === 0) {
          console.log(`generated ${i} wallets...`);
        }
        return Wallet.createRandom();
      });

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const tree = StandardMerkleTree.of(
      bobs.map(b => [b.address]),
      ['address'],
    );

    const { ticket, redeemerRand, senderSig, ticketHash } =
      await createWinningMultiReceiverTicket(
        syloTicketing,
        epochsManager,
        alice,
        tree.root,
        owner,
      );

    const receiverSig = await bobs[0].signMessage(ethers.getBytes(ticketHash));
    const proof = tree.getProof([bobs[0].address]);

    console.log('proof size: ', proof.length);

    console.log(
      'gas limit estimation: ',
      await syloTicketing.redeemMultiReceiver.estimateGas(
        ticket,
        redeemerRand,
        { main: bobs[0].address, delegated: ethers.ZeroAddress },
        proof,
        senderSig,
        receiverSig,
      ),
    );
    await syloTicketing.redeemMultiReceiver(
      ticket,
      redeemerRand,
      { main: bobs[0].address, delegated: ethers.ZeroAddress },
      proof,
      senderSig,
      receiverSig,
    );
  }).timeout(0);
});
