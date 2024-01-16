import { ethers } from 'hardhat';
import { assert, expect } from 'chai';
import { BigNumberish, HDNodeWallet, Signer } from 'ethers';
import {
  EpochsManager,
  Registries,
  RewardsManager,
  StakingManager,
  SyloTicketing,
  SyloToken,
  TestSeekers,
  SeekerPowerOracle,
} from '../../typechain-types';
import * as contractTypes from '../../typechain-types';
import web3 from 'web3';
import utils from '../utils';

// This test suite relies on confirming that updated stakes and rewards are correctly
// calculated after incrementing the reward pool. However due to minor precision loss, the
// actual balance may slightly differ.
// This function checks the difference falls within a small fraction of a single SYLO.
export function compareExpectedBalance(a: BigNumberish, b: BigNumberish): void {
  const abs = (n: bigint) => (n < 0 ? -n : n);
  const diff = abs(BigInt(a) - BigInt(b));
  // NOTE: This essentially says that a margin of 10**4 SOLOs is acceptable, or
  // 0.00000000000001 SYLOs
  expect(diff).to.be.within(0, BigInt(10 ** 4));
}

export function toSOLOs(a: number): bigint {
  return BigInt(web3.utils.toWei(a.toString()));
}

export function fromSOLOs(a: number | bigint | string): string {
  return web3.utils.fromWei(a.toString());
}

// Helper function to initialize stakes for multiple stakees,
// and returns an array of stake proportions.
export const addStakes = async (
  token: SyloToken,
  stakingManager: StakingManager,
  owner: string,
  stakees: { account: Signer; stake: number }[],
): Promise<{ totalStake: number; proportions: number[] }> => {
  const totalStake = stakees.reduce((p, c) => {
    return { ...p, stake: p.stake + c.stake };
  }).stake;

  const proportions = await Promise.all(
    stakees.map(async s => {
      const stake = toSOLOs(s.stake);

      await token.transfer(await s.account.getAddress(), stake);
      await token
        .connect(s.account)
        .approve(await stakingManager.getAddress(), stake);

      await stakingManager.connect(s.account).addStake(stake, owner);

      return s.stake / totalStake;
    }),
  );

  return { totalStake, proportions };
};

// Helper function for testing that an account's SYLO balance
// increased as expected after making a claim
export const testClaims = async (
  token: SyloToken,
  rewardsManager: RewardsManager,
  owner: string,
  tests: { account: Signer; claim: number }[],
): Promise<void> => {
  for (const t of tests) {
    const claim = toSOLOs(t.claim);
    const address = await t.account.getAddress();

    const expectedBalance = await token
      .balanceOf(address)
      .then((b: bigint) => b + claim);

    await rewardsManager.connect(t.account).claimStakingRewards(owner);

    compareExpectedBalance(expectedBalance, await token.balanceOf(address));
  }
};

export const checkAfterRedeem = async (
  syloTicketing: SyloTicketing,
  rewardsManager: RewardsManager,
  owner: string,
  sender: HDNodeWallet,
  expectedEscrow: number,
  expectedPenalty: number,
  expectedRewards: number,
): Promise<void> => {
  const deposit = await syloTicketing.deposits(sender.address);
  assert.equal(
    deposit.escrow,
    toSOLOs(expectedEscrow),
    'Expected ticket payout to be substracted from escrow',
  );
  assert.equal(
    deposit.penalty,
    toSOLOs(expectedPenalty),
    'Expected penalty to not be changed',
  );

  const pendingReward = await rewardsManager.getPendingRewards(owner);

  assert.equal(
    pendingReward,
    toSOLOs(expectedRewards),
    'Expected balance of pending rewards to have added the ticket face value',
  );
};

export async function setSeekerRegistry(
  seekers: TestSeekers,
  registries: Registries,
  seekerPowerOracle: SeekerPowerOracle,
  account: Signer,
  seekerAccount: Signer,
  tokenId: number,
): Promise<void> {
  await utils.setSeekerRegistry(
    registries,
    seekers,
    seekerPowerOracle,
    account,
    seekerAccount,
    tokenId,
  );
}

export async function createWinningTicket(
  syloTicketing: SyloTicketing,
  epochsManager: EpochsManager,
  sender: Signer,
  receiver: Signer,
  redeemer: string,
  epochId?: number,
  senderDelegatedWallet?: HDNodeWallet,
  receiverDelegatedWallet?: HDNodeWallet,
): Promise<{
  ticket: contractTypes.contracts.interfaces.payments.ISyloTicketing.TicketStruct;
  receiver: Signer;
  redeemerRand: number;
  senderSig: Uint8Array;
  receiverSig: Uint8Array;
  ticketHash: string;
}> {
  const generationBlock = BigInt(
    ((await syloTicketing.runner?.provider?.getBlockNumber()) ?? 0) + 1,
  );

  const redeemerRand = 1;
  const redeemerCommit = createCommit(generationBlock, redeemerRand);

  let senderDelegatedAccount = ethers.ZeroAddress;
  if (senderDelegatedWallet) {
    senderDelegatedAccount = senderDelegatedWallet.address;
  }

  let receiverDelegatedAccount = ethers.ZeroAddress;
  if (receiverDelegatedWallet) {
    receiverDelegatedAccount = receiverDelegatedWallet.address;
  }

  const ticket = {
    epochId: epochId ?? (await epochsManager.currentIteration()),
    sender: {
      main: await sender.getAddress(),
      delegated: senderDelegatedAccount,
    },
    receiver: {
      main: await receiver.getAddress(),
      delegated: receiverDelegatedAccount,
    },
    redeemer,
    generationBlock,
    redeemerCommit: ethers.hexlify(redeemerCommit),
  };

  const ticketHash = await syloTicketing.getTicketHash(ticket);
  let senderSig = await sender.signMessage(ethers.getBytes(ticketHash));
  if (senderDelegatedWallet) {
    senderSig = await senderDelegatedWallet.signMessage(
      ethers.getBytes(ticketHash),
    );
  }
  if (!senderSig) {
    throw new Error('failed to derive sender signature for ticket');
  }

  let receiverSig = await receiver.signMessage(ethers.getBytes(ticketHash));
  if (receiverDelegatedWallet) {
    receiverSig = await receiverDelegatedWallet.signMessage(
      ethers.getBytes(ticketHash),
    );
  }
  if (!receiverSig) {
    throw new Error('failed to derive receiver signature for ticket');
  }

  return {
    ticket,
    receiver,
    redeemerRand,
    senderSig: ethers.getBytes(senderSig),
    receiverSig: ethers.getBytes(receiverSig),
    ticketHash,
  };
}

export async function createWinningMultiReceiverTicket(
  syloTicketing: SyloTicketing,
  epochsManager: EpochsManager,
  sender: HDNodeWallet,
  redeemer: string,
  epochId?: number,
  senderDelegatedWallet?: HDNodeWallet,
): Promise<{
  ticket: contractTypes.contracts.interfaces.payments.ISyloTicketing.MultiReceiverTicketStruct;
  redeemerRand: number;
  senderSig: Uint8Array;
  ticketHash: string;
}> {
  const generationBlock = BigInt((await ethers.provider.getBlockNumber()) + 1);

  const redeemerRand = 1;
  const redeemerCommit = createCommit(generationBlock, redeemerRand);

  let senderDelegatedAccount = ethers.ZeroAddress;
  if (senderDelegatedWallet) {
    senderDelegatedAccount = senderDelegatedWallet.address;
  }

  const ticket = {
    epochId: epochId ?? (await epochsManager.currentIteration()),
    sender: {
      main: sender.address,
      delegated: senderDelegatedAccount,
    },
    redeemer,
    generationBlock,
    redeemerCommit: ethers.hexlify(redeemerCommit),
  };

  const ticketHash = await syloTicketing.getMultiReceiverTicketHash(ticket);
  let senderSig = await sender.signMessage(ethers.getBytes(ticketHash));
  if (senderDelegatedWallet) {
    senderSig = await senderDelegatedWallet.signMessage(
      ethers.getBytes(ticketHash),
    );
  }
  if (!senderSig) {
    throw new Error('failed to derive sender signature for ticket');
  }

  return {
    ticket,
    redeemerRand,
    senderSig: ethers.getBytes(senderSig),
    ticketHash,
  };
}

export function createCommit(
  generationBlock: bigint,
  rand: BigNumberish,
): string {
  return ethers.solidityPackedKeccak256(
    ['bytes32'],
    [
      ethers.solidityPackedKeccak256(
        ['uint256', 'uint256'],
        [generationBlock, rand],
      ),
    ],
  );
}
