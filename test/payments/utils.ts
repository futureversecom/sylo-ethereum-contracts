import { ethers } from 'hardhat';
import { assert, expect } from 'chai';
import { BigNumberish, BytesLike, HDNodeWallet, Signer } from 'ethers';
import {
  EpochsManager,
  Registries,
  RewardsManager,
  StakingManager,
  SyloTicketing,
  SyloToken,
  TestSeekers,
  SeekerPowerOracle,
  IAuthorizedAccounts,
} from '../../typechain-types';
import * as contractTypes from '../../typechain-types';
import web3 from 'web3';
import utils from '../utils';
import { SignatureType } from '../../common/enum';

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
  ticket: contractTypes.ISyloTicketing.TicketStruct;
  receiver: Signer;
  redeemerRand: number;
  senderSig: contractTypes.ISyloTicketing.UserSignatureStruct;
  receiverSig: contractTypes.ISyloTicketing.UserSignatureStruct;
  ticketHash: string;
}> {
  const generationBlock = BigInt(
    ((await syloTicketing.runner?.provider?.getBlockNumber()) ?? 0) + 1,
  );

  const redeemerRand = 1;
  const redeemerCommit = createCommit(generationBlock, redeemerRand);

  const ticket = {
    epochId: epochId ?? (await epochsManager.currentIteration()),
    sender: await sender.getAddress(),
    receiver: await receiver.getAddress(),
    redeemer,
    generationBlock,
    redeemerCommit: ethers.hexlify(redeemerCommit),
  };

  const ticketHash = await syloTicketing.getTicketHash(ticket);

  return {
    ticket,
    receiver,
    redeemerRand,
    senderSig: await createUserSignature(
      ethers.getBytes(ethers.getBytes(ticketHash)),
      senderDelegatedWallet ? SignatureType.Authorized : SignatureType.Main,
      sender,
      senderDelegatedWallet,
    ),
    receiverSig: await createUserSignature(
      ethers.getBytes(ethers.getBytes(ticketHash)),
      receiverDelegatedWallet ? SignatureType.Authorized : SignatureType.Main,
      receiver,
      receiverDelegatedWallet,
    ),
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
  ticket: contractTypes.ISyloTicketing.MultiReceiverTicketStruct;
  redeemerRand: number;
  senderSig: contractTypes.ISyloTicketing.UserSignatureStruct;
  ticketHash: string;
}> {
  const generationBlock = BigInt((await ethers.provider.getBlockNumber()) + 1);

  const redeemerRand = 1;
  const redeemerCommit = createCommit(generationBlock, redeemerRand);

  const ticket = {
    epochId: epochId ?? (await epochsManager.currentIteration()),
    sender: sender.address,
    redeemer,
    generationBlock,
    redeemerCommit: ethers.hexlify(redeemerCommit),
  };

  const ticketHash = await syloTicketing.getMultiReceiverTicketHash(ticket);

  return {
    ticket,
    redeemerRand,
    senderSig: await createUserSignature(
      ethers.getBytes(ethers.getBytes(ticketHash)),
      senderDelegatedWallet ? SignatureType.Authorized : SignatureType.Main,
      sender,
      senderDelegatedWallet,
    ),
    ticketHash,
  };
}

export async function createUserSignature(
  message: BytesLike,
  signatureType: SignatureType,
  account: Signer,
  delegated?: HDNodeWallet,
  authorizedAccounts?: contractTypes.AuthorizedAccounts,
): Promise<contractTypes.ISyloTicketing.UserSignatureStruct> {
  let signature: string;
  if (signatureType != SignatureType.Main) {
    if (!delegated) {
      throw new Error(
        'Must supply delegated account for authorized signatures',
      );
    }
    signature = await delegated.signMessage(message);
  } else {
    signature = await account.signMessage(message);
  }

  return {
    sigType: signatureType,
    signature,
    authorizedAccount:
      signatureType === SignatureType.Authorized && delegated
        ? delegated
        : ethers.ZeroAddress,
    attachedAuthorizedAccount:
      signatureType === SignatureType.AttachedAuthorized &&
      delegated &&
      authorizedAccounts
        ? await createAttachedAuthorizedAccount(
            account,
            delegated,
            authorizedAccounts,
          )
        : createEmptyAttachedAuthorizedAccount(),
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

export async function createAttachedAuthorizedAccount(
  main: Signer,
  delegatedWallet: HDNodeWallet,
  authorizedAccounts: contractTypes.AuthorizedAccounts,
): Promise<IAuthorizedAccounts.AttachedAuthorizedAccountStruct> {
  const block = await ethers.provider.getBlock('latest');

  const expiry = (block?.timestamp ?? 0) + 10000000;

  const prefix = 'prefix';
  const suffix = 'suffix';
  const infixOne = 'infix';

  const proofMessage =
    await authorizedAccounts.createAttachedAuthorizedAccountProofMessage(
      delegatedWallet.address,
      expiry,
      prefix,
      suffix,
      infixOne,
    );

  const proof = await main.signMessage(
    Buffer.from(proofMessage.slice(2), 'hex'),
  );

  return {
    account: delegatedWallet.address,
    expiry,
    proof,
    prefix,
    suffix,
    infixOne,
  };
}

export function createEmptyAttachedAuthorizedAccount(): IAuthorizedAccounts.AttachedAuthorizedAccountStruct {
  return {
    account: ethers.ZeroAddress,
    expiry: 0,
    proof: new Uint8Array(),
    prefix: '',
    suffix: '',
    infixOne: '',
  };
}
