import { ethers } from 'hardhat';
import { assert, expect } from 'chai';
import { BytesLike, HDNodeWallet, Signer, Wallet } from 'ethers';
import {
  AuthorizedAccounts,
  EpochsManager,
  Registries,
  RewardsManager,
  SeekerPowerOracle,
  StakingManager,
  SyloTicketing,
  SyloToken,
  TestFuturepassRegistrar,
  TestSeekers,
  TicketingParameters,
} from '../../typechain-types';
import {
  toSOLOs,
  setSeekerRegistry,
  createWinningMultiReceiverTicket,
  createUserSignature,
} from './utils';
import utils from '../utils';
import { SignatureType } from '../../common/enum';

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
  let seekerPowerOracle: SeekerPowerOracle;
  let futurepassRegistrar: TestFuturepassRegistrar;
  let authorizedAccounts: AuthorizedAccounts;

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
    seekerPowerOracle = contracts.seekerPowerOracle;
    futurepassRegistrar = contracts.futurepassRegistrar;
    authorizedAccounts = contracts.authorizedAccounts;

    await ticketingParameters.setMultiReceiverFaceValue(faceValue);

    await token.approve(await stakingManager.getAddress(), toSOLOs(10000000));
    await token.approve(await syloTicketing.getAddress(), toSOLOs(10000000));
  });

  it('can redeem multi receiver ticket', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(
      seekers,
      registries,
      seekerPowerOracle,
      accounts[0],
      accounts[1],
      1,
    );

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bobs = await createFuturepassReceivers(5);

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, redeemerRand, senderSig, ticketHash } =
      await createWinningMultiReceiverTicket(
        syloTicketing,
        epochsManager,
        alice,
        owner,
      );

    const receiverSig = await createUserSignature(
      ethers.getBytes(ticketHash),
      SignatureType.Main,
      bobs[0],
    );

    await syloTicketing.redeemMultiReceiver(
      ticket,
      redeemerRand,
      bobs[0].address,
      senderSig,
      receiverSig,
    );

    const deposit = await syloTicketing.deposits(alice.address);
    assert.equal(
      deposit.escrow,
      toSOLOs(1000),
      'Expected ticket payout to be subtracted from escrow',
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
    await setSeekerRegistry(
      seekers,
      registries,
      seekerPowerOracle,
      accounts[0],
      accounts[1],
      1,
    );

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bobs = await createFuturepassReceivers(5);

    await syloTicketing.depositEscrow(toSOLOs(20000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, redeemerRand, senderSig, ticketHash } =
      await createWinningMultiReceiverTicket(
        syloTicketing,
        epochsManager,
        alice,
        owner,
      );

    let previousDeposit = await syloTicketing.deposits(alice.address);
    let previousReward = await rewardsManager.getPendingRewards(owner);

    for (const bob of bobs) {
      const receiverSig = await createUserSignature(
        ethers.getBytes(ticketHash),
        SignatureType.Main,
        bob,
      );

      await syloTicketing.redeemMultiReceiver(
        ticket,
        redeemerRand,
        bob.address,
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
    const bobs = await createFuturepassReceivers(5);

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, redeemerRand, senderSig, ticketHash } =
      await createWinningMultiReceiverTicket(
        syloTicketing,
        epochsManager,
        alice,
        owner,
      );

    const receiverSig = await createUserSignature(
      ethers.getBytes(ticketHash),
      SignatureType.Main,
      bobs[0],
    );

    await expect(
      syloTicketing.redeemMultiReceiver(
        { ...ticket, sender: ethers.ZeroAddress },
        redeemerRand,
        bobs[0].address,
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
        ethers.ZeroAddress,
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
        bobs[0].address,
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
        bobs[0].address,
        senderSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(syloTicketing, 'RedeemerCommitMismatch');

    await expect(
      syloTicketing.redeemMultiReceiver(
        ticket,
        redeemerRand,
        bobs[0].address,
        {
          ...senderSig,
          authorizedAccount: bobs[0].address,
          sigType: SignatureType.Authorized,
        },
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(syloTicketing, 'InvalidSigningPermission');

    await expect(
      syloTicketing.redeemMultiReceiver(
        ticket,
        redeemerRand,
        bobs[0].address,
        senderSig,
        {
          ...receiverSig,
          authorizedAccount: alice.address,
          sigType: SignatureType.Authorized,
        },
      ),
    ).to.be.revertedWithCustomError(syloTicketing, 'InvalidSigningPermission');

    const malformedSig = {
      ...senderSig,
      signature:
        '0xdebcaaaa727df04bdc990083d88ed7c8e6e9897ff18b7d968867a8bc024cbdbe10ca52eebd67a14b7b493f5c00ed9dab7b96ef62916f25afc631d336f7b2ae1e1b',
    };
    await expect(
      syloTicketing.redeemMultiReceiver(
        ticket,
        redeemerRand,
        bobs[0].address,
        malformedSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(syloTicketing, 'InvalidSignature');

    await expect(
      syloTicketing.redeemMultiReceiver(
        ticket,
        redeemerRand,
        bobs[0].address,
        senderSig,
        malformedSig,
      ),
    ).to.be.revertedWithCustomError(syloTicketing, 'InvalidSignature');
  });

  it('can not redeem non winning ticket', async () => {
    await ticketingParameters.setBaseLiveWinProb(0);

    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(
      seekers,
      registries,
      seekerPowerOracle,
      accounts[0],
      accounts[1],
      1,
    );

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bobs = await createFuturepassReceivers(5);

    const { ticket, redeemerRand, senderSig, ticketHash } =
      await createWinningMultiReceiverTicket(
        syloTicketing,
        epochsManager,
        alice,
        owner,
      );

    const receiverSig = await createUserSignature(
      ethers.getBytes(ticketHash),
      SignatureType.Main,
      bobs[0],
    );

    await expect(
      syloTicketing.redeemMultiReceiver(
        ticket,
        redeemerRand,
        bobs[0].address,
        senderSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(syloTicketing, 'TicketNotWinning');
  });

  it('can not redeem ticket for future block', async () => {
    const alice = Wallet.createRandom();
    const bobs = await createFuturepassReceivers(5);

    const { ticket, redeemerRand, senderSig } =
      await createWinningMultiReceiverTicket(
        syloTicketing,
        epochsManager,
        alice,
        owner,
      );

    await expect(
      syloTicketing.redeemMultiReceiver(
        {
          ...ticket,
          generationBlock: BigInt(ticket.generationBlock) + 10n,
        },
        redeemerRand,
        bobs[0].address,
        senderSig,
        senderSig,
      ),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'TicketCannotBeFromFutureBlock',
    );
  });

  it('cannot redeem multi receiver ticket if node has not joined epoch', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(
      seekers,
      registries,
      seekerPowerOracle,
      accounts[0],
      accounts[1],
      1,
    );

    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bobs = await createFuturepassReceivers(5);

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, redeemerRand, senderSig, ticketHash } =
      await createWinningMultiReceiverTicket(
        syloTicketing,
        epochsManager,
        alice,
        owner,
      );

    const receiverSig = await createUserSignature(
      ethers.getBytes(ticketHash),
      SignatureType.Main,
      bobs[0],
    );

    await expect(
      syloTicketing.redeemMultiReceiver(
        ticket,
        redeemerRand,
        bobs[0].address,
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
    await setSeekerRegistry(
      seekers,
      registries,
      seekerPowerOracle,
      accounts[0],
      accounts[1],
      1,
    );

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, redeemerRand, senderSig, ticketHash } =
      await createWinningMultiReceiverTicket(
        syloTicketing,
        epochsManager,
        alice,
        owner,
      );

    const invalidReceiver = Wallet.createRandom();
    // we use a proof from a valid receiver but supply an invalid receiver
    const receiverSig = await createUserSignature(
      ethers.getBytes(ticketHash),
      SignatureType.Main,
      invalidReceiver,
    );

    await expect(
      syloTicketing.redeemMultiReceiver(
        ticket,
        redeemerRand,
        invalidReceiver.address,
        senderSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(syloTicketing, 'MissingFuturepassAccount');
  });

  it('can not redeem for the same user more than once', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(
      seekers,
      registries,
      seekerPowerOracle,
      accounts[0],
      accounts[1],
      1,
    );

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const bobs = await createFuturepassReceivers(5);

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, redeemerRand, senderSig, ticketHash } =
      await createWinningMultiReceiverTicket(
        syloTicketing,
        epochsManager,
        alice,
        owner,
      );

    const receiverSig = await createUserSignature(
      ethers.getBytes(ticketHash),
      SignatureType.Main,
      bobs[0],
    );

    await syloTicketing.redeemMultiReceiver(
      ticket,
      redeemerRand,
      bobs[0].address,
      senderSig,
      receiverSig,
    );

    await expect(
      syloTicketing.redeemMultiReceiver(
        ticket,
        redeemerRand,
        bobs[0].address,
        senderSig,
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(syloTicketing, 'TicketAlreadyRedeemed');
  });

  it('can redeem for receiver using attached authorized account', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(
      seekers,
      registries,
      seekerPowerOracle,
      accounts[0],
      accounts[1],
      1,
    );

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const [bob] = await createFuturepassReceivers(1);
    const delegatedWallet = Wallet.createRandom();

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, redeemerRand, senderSig, ticketHash } =
      await createWinningMultiReceiverTicket(
        syloTicketing,
        epochsManager,
        alice,
        owner,
      );

    const receiverSig = await createUserSignature(
      ethers.getBytes(ticketHash),
      SignatureType.Main,
      bob,
      delegatedWallet,
      authorizedAccounts,
    );

    await syloTicketing.redeemMultiReceiver(
      ticket,
      redeemerRand,
      bob.address,
      senderSig,
      receiverSig,
    );
  });

  it('cannot redeem ticket when using sender attached authorized account', async () => {
    await stakingManager.addStake(toSOLOs(1), owner);
    await setSeekerRegistry(
      seekers,
      registries,
      seekerPowerOracle,
      accounts[0],
      accounts[1],
      1,
    );

    await epochsManager.joinNextEpoch();
    await epochsManager.initializeEpoch();

    const alice = Wallet.createRandom();
    const [bob] = await createFuturepassReceivers(1);
    const delegatedWallet = Wallet.createRandom();

    await syloTicketing.depositEscrow(toSOLOs(2000), alice.address);
    await syloTicketing.depositPenalty(toSOLOs(50), alice.address);

    const { ticket, redeemerRand, senderSig, ticketHash } =
      await createWinningMultiReceiverTicket(
        syloTicketing,
        epochsManager,
        alice,
        owner,
      );

    const receiverSig = await createUserSignature(
      ethers.getBytes(ticketHash),
      SignatureType.Main,
      bob,
      delegatedWallet,
      authorizedAccounts,
    );

    await expect(
      syloTicketing.redeemMultiReceiver(
        ticket,
        redeemerRand,
        bob.address,
        { ...senderSig, sigType: SignatureType.AttachedAuthorized },
        receiverSig,
      ),
    ).to.be.revertedWithCustomError(
      syloTicketing,
      'SenderCannotUseAttachedAuthorizedAccount',
    );
  });

  async function createFuturepassReceivers(n: number): Promise<HDNodeWallet[]> {
    return Promise.all(
      Array(n)
        .fill(0)
        .map(async _ => {
          const w = Wallet.createRandom();
          // register futurepass account
          await futurepassRegistrar.create(w.address);
          return w;
        }),
    );
  }
});
