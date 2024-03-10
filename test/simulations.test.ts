import { ethers } from 'hardhat';
import { BaseContract, Contract, MaxUint256, Signer } from 'ethers';
import {
  Directory,
  EpochsManager,
  IRewardsManager,
  Registries,
  RewardsManager,
  RewardsManager__factory,
  SeekerPowerOracle,
  StakingManager,
  SyloTicketing,
  SyloToken,
  TestSeekers,
} from '../typechain-types';
import utils from './utils';
import { assert, expect } from 'chai';
// Chi Squared goodness of fit test
import { chi2gof } from '@stdlib/stats';
import crypto from 'crypto';
import * as fs from 'fs/promises';
import { createWinningTicket, setSeekerRegistry } from './payments/utils';
import { SyloContracts } from '../common/contracts';

type Results = { [key: string]: number };

const MAX_SYLO_STAKE = ethers.parseEther('10000000000');

type Transaction = {
  From: string;
  To: string;
  BlockNumber: number;
  RawInput: string;
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  InputsMap: { [key: string]: any };
  ExtrinsicId: string;
};

describe('Simulations', () => {
  let accounts: Signer[];
  let owner: string;

  let token: SyloToken;

  before(async () => {
    accounts = await ethers.getSigners();
    // first account is implicitly used as deployer of contracts in hardhat
    owner = await accounts[0].getAddress();

    const Token = await ethers.getContractFactory('SyloToken');
    token = await Token.deploy();
  });

  it.only('simulates porcini deployment', async () => {
    // if (!process.env.RUN_PORCINI_SIM) {
    //   return;
    // }

    const contracts = await utils.onlyDeployContracts(token);

    await token.approve(await contracts.stakingManager.getAddress(), 100000);

    await token.approve(
      await contracts.syloTicketing.getAddress(),
      ethers.parseEther('100000000000'),
    );

    const impersonatedSigners: { [account: string]: Signer } = {};

    const getAndFundSigner = async (
      account: string,
      ethFund?: string,
      syloFund?: string,
    ): Promise<Signer> => {
      if (impersonatedSigners[account] != null) {
        return impersonatedSigners[account];
      }

      const signer = await ethers.getImpersonatedSigner(account);

      // send eth
      await accounts[0]
        .sendTransaction({
          to: account,
          value: ethers.parseEther(ethFund ?? '200'),
        })
        .then(_ => utils.advanceBlock(1));

      // send sylo
      await token
        .transfer(account, ethers.parseEther(syloFund ?? '1000000000'))
        .then(_ => utils.advanceBlock(1));

      await token
        .connect(signer)
        .approve(
          await contracts.syloTicketing.getAddress(),
          ethers.parseEther('1000000000'),
        )
        .then(_ => utils.advanceBlock(1));
      await token
        .connect(signer)
        .approve(
          await contracts.stakingManager.getAddress(),
          ethers.parseEther('1000000000'),
        )
        .then(_ => utils.advanceBlock(1));

      impersonatedSigners[account] = signer;

      return signer;
    };

    const startingBlock = 10900000;
    await utils.advanceBlock(startingBlock);

    console.log(
      ethers.formatEther(
        await ethers.provider.getBalance(await accounts[0].getAddress()),
      ),
    );

    // fund node accounts
    const porciniDeployer = await getAndFundSigner(
      '0x448c8e9e1816300dd052e77d2a44c990a2807d15',
      '2000',
      '1000000000',
    );
    await getAndFundSigner(
      '0xa5a5a6e97528a6ba1ee04f27582d37e9b612f6c3',
      '500',
      '100000000',
    );
    await getAndFundSigner(
      '0x3fD3fA93c55AB830D99959fD1EaBe9eCBc1E3F96',
      '500',
      '100000000',
    );
    const node4 = await getAndFundSigner(
      '0xb521c2a19Df3b949acB766DE987EFc6787584570',
      '500',
      '100000000',
    );

    console.log(await node4.getAddress());

    console.log('funded nodes');

    const contractMappings: { [to: string]: BaseContract } = {
      '0xcccccccc00000c64000000000000000000000000': contracts.syloToken,
      '0x68ecb081f49690b3453f9a4ffa65c054d74dcfe0':
        contracts.authorizedAccounts,
      '0x76c226017d7b0cecc467dc0f353268d537252607': contracts.registries,
      '0xce1bad6fac2edab78c607bfa50a9bfc29a172214':
        contracts.ticketingParameters,
      '0x9591e2f0a64e39675c9c3c28d055974fa1b93774': contracts.epochsManager,
      '0x9dbdad1af4c0ab8065a2ad0e55172590d6655cff': contracts.stakingManager,
      '0x6802094ec43b1d769aa41a6925244aedb0d294a0': contracts.rewardsManager,
      '0x75abf214870415a94b0d38840cd541e9a0f196fa': contracts.directory,
      '0x62d39eb18f77143aa3b847fdb54f95db896fc28f': contracts.syloTicketing,
      '0xaaaaaaaa00001864000000000000000000000000': contracts.seekers,
      '0xdb51e2219fc69627ae29d177e4a1c9d9129e8230': contracts.seekerPowerOracle,
    };

    const transactionDataBuf = await fs.readFile('transaction_data.json');

    const transactionData = JSON.parse(
      transactionDataBuf.toString(),
    ) as Transaction[];

    await contracts.directory
      .connect(porciniDeployer)
      .initialize(
        await contracts.stakingManager.getAddress(),
        await contracts.rewardsManager.getAddress(),
      )
      .then(_ => utils.advanceBlock(1));

    await contracts.directory
      .connect(porciniDeployer)
      .addManager(await contracts.epochsManager.getAddress());

    await contracts.seekerPowerOracle.initialize(
      await porciniDeployer.getAddress(),
    );

    const processed: { [id: string]: boolean } = {};

    let count = 0;

    let lastMod = '';
    for (const transaction of transactionData) {
      if (processed[transaction.ExtrinsicId]) {
        continue;
      }

      // ensure any seekers are 'bridged' and owned
      if (transaction.To === '0x76c226017d7b0cecc467dc0f353268d537252607') {
        const seekerAccount = transaction.InputsMap['seekerAccount'] as string;
        const seekerId = transaction.InputsMap['seekerId'] as number;
        if (seekerAccount != undefined && seekerId != undefined) {
          await contracts.seekers
            .mint(seekerAccount, seekerId)
            .then(_ => utils.advanceBlock(1));

          await contracts.seekerPowerOracle
            .connect(porciniDeployer)
            .registerSeekerPowerRestricted(seekerId, 30)
            .then(_ => utils.advanceBlock(1));
        }
      }

      if (transaction.RawInput.startsWith('0x96bb1fef')) {
        continue;
      }

      const currentBlock = await ethers.provider.getBlock('latest');
      if (currentBlock == null) {
        throw new Error(`could not get block`);
      }

      if (currentBlock.number < transaction.BlockNumber) {
        await utils.advanceBlock(
          transaction.BlockNumber - currentBlock.number - 1,
        );
      }

      const signer = await getAndFundSigner(transaction.From);

      // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
      const to = await contractMappings[transaction.To]!.getAddress();

      // we replace any inputs that use contract addresses, and replace
      // those addresses with equivalent in the local testnet
      let data = transaction.RawInput;
      for (const [key, value] of Object.entries(contractMappings)) {
        const addr = await value.getAddress();

        data = data.replace(
          new RegExp(key.slice(2).toLowerCase()),
          addr.slice(2).toLowerCase(),
        );
      }

      // we also replace input data for any nodes
      // if (transaction.RawInput.startsWith('0x2d49aa1c')) {
      // }

      // console.log(
      //   `sending transaction: to=${to} (${transaction.To}), extrinsic-id=${transaction.ExtrinsicId}, block=${transaction.BlockNumber}`,
      // );

      await signer
        .sendTransaction({
          to,
          data,
        })
        .then(tx => tx.wait(1));

      count++;

      processed[transaction.ExtrinsicId] = true;

      if (transaction.RawInput === '0xe1519a75') {
        const currentIteration =
          await contracts.epochsManager.currentIteration();
        console.log(`initialized epoch: ${currentIteration}`);
      }

      if (
        transaction.RawInput === '0x45526649' &&
        transaction.From == '0xb521c2a19df3b949acb766de987efc6787584570'
      ) {
        const currentIteration =
          await contracts.epochsManager.currentIteration();
        const newRewardPool = await contracts.rewardsManager.getRewardPool(
          currentIteration,
          '0xb521c2a19df3b949acb766de987efc6787584570',
        );

        console.log(
          `initialized new reward pool (${currentIteration}): ${renderRewardPool(
            newRewardPool,
          )}`,
        );
      }

      if (
        transaction.From == '0xffffffff0000000000000000000000000000039c' &&
        (transaction.RawInput.startsWith('0x2d49aa1c') ||
          transaction.RawInput.startsWith('0xa859f172') ||
          transaction.RawInput.startsWith('0x96bb1fef')) &&
        transaction.InputsMap['stakee'] ===
          '0xb521c2a19df3b949acb766de987efc6787584570'
      ) {
        if (transaction.RawInput.startsWith('0x2d49aa1c')) {
          lastMod = 'addStake';
        }

        if (transaction.RawInput.startsWith('0xa859f172')) {
          lastMod = 'unlockStake';
        }

        if (transaction.RawInput.startsWith('0x96bb1fef')) {
          lastMod = 'claimStakingRewards';
        }

        for (const node of [
          // '0x448c8e9e1816300dd052e77d2a44c990a2807d15',
          // '0xa5a5a6e97528a6ba1ee04f27582d37e9b612f6c3',
          // '0x3fd3fa93c55ab830d99959fd1eabe9ecbc1e3f96',
          '0xb521c2a19df3b949acb766de987efc6787584570',
        ]) {
          console.log('\n');

          console.log(
            `sending transaction: to=${to} (${transaction.To}), from=${transaction.From}, extrinsic-id=${transaction.ExtrinsicId}, block=${transaction.BlockNumber}`,
          );

          if (transaction.RawInput.startsWith('0x2d49aa1c')) {
            console.log(
              `amount: ${JSON.stringify(transaction.InputsMap['amount'])}`,
            );
          }

          if (transaction.RawInput.startsWith('0xa859f172')) {
            console.log(
              `amount: ${JSON.stringify(transaction.InputsMap['amount'])}`,
            );
          }

          console.log('updating values from', lastMod);

          console.log(
            `stake: `,
            ethers.formatEther(
              await contracts.stakingManager.getStakeeTotalManagedStake(node),
            ),
          );

          const fpAddr = '0xFFFFfFff0000000000000000000000000000039C';
          const claim = await contracts.rewardsManager.calculateStakerClaim(
            node,
            '0xFFFFfFff0000000000000000000000000000039C',
          );

          console.log(`claim: `, ethers.formatEther(claim));

          const stakerKey = await contracts.rewardsManager.getStakerKey(
            node,
            fpAddr,
          );

          const currentEpoch = await contracts.epochsManager.currentIteration();

          const lastClaim = await contracts.rewardsManager.getLastClaim(
            node,
            '0xFFFFfFff0000000000000000000000000000039C',
          );

          const initialClaim =
            await contracts.rewardsManager.calculateInitialClaim(
              stakerKey,
              node,
            );

          const finalCrf =
            await contracts.rewardsManager.getFinalCumulativeRewardFactor(
              node,
              await contracts.epochsManager.currentIteration(),
            );

          const lastClaimRewardPool =
            await contracts.rewardsManager.getRewardPool(
              lastClaim.claimedAt,
              node,
            );

          console.log('current epoch: ', currentEpoch);
          console.log(
            `last claim: ${lastClaim.claimedAt}, ${ethers.formatEther(
              lastClaim.stake,
            )}`,
          );
          console.log(`initial claim: ${ethers.formatEther(initialClaim)}`);
          console.log('final crf: ', Number(finalCrf) / 9223372036854775807);
          console.log(
            `last claim reward pool: ${renderRewardPool(lastClaimRewardPool)}`,
          );

          lastMod = '';

          const fpSigner = await getAndFundSigner(
            '0xFFFFfFff0000000000000000000000000000039C',
          );

          if (claim > 0n) {
            await contracts.rewardsManager
              .connect(fpSigner)
              .claimStakingRewards.estimateGas(node);
          }

          console.log('\n');
        }
      }
    }

    console.log('DONE');

    for (const node of [
      // '0x448c8e9e1816300dd052e77d2a44c990a2807d15',
      // '0xa5a5a6e97528a6ba1ee04f27582d37e9b612f6c3',
      // '0x3fd3fa93c55ab830d99959fd1eabe9ecbc1e3f96',
      '0xb521c2a19df3b949acb766de987efc6787584570',
    ]) {
      console.log(
        `${node} stake: `,
        ethers.formatEther(
          await contracts.stakingManager.getStakeeTotalManagedStake(node),
        ),
      );

      const claim = await contracts.rewardsManager.calculateStakerClaim(
        node,
        '0xFFFFfFff0000000000000000000000000000039C',
      );

      console.log(`${node} pending reward: `, ethers.formatEther(claim));

      const signer = await getAndFundSigner(
        '0xFFFFfFff0000000000000000000000000000039C',
      );

      const fpAddr = '0xFFFFfFff0000000000000000000000000000039C';

      const stakerKey = await contracts.rewardsManager.getStakerKey(
        node,
        fpAddr,
      );

      const currentEpoch = await contracts.epochsManager.currentIteration();

      const lastClaim = await contracts.rewardsManager.getLastClaim(
        node,
        '0xFFFFfFff0000000000000000000000000000039C',
      );

      const initialClaim = await contracts.rewardsManager.calculateInitialClaim(
        stakerKey,
        node,
      );

      const finalCrf =
        await contracts.rewardsManager.getFinalCumulativeRewardFactor(
          node,
          await contracts.epochsManager.currentIteration(),
        );

      const lastClaimRewardPool = await contracts.rewardsManager.getRewardPool(
        lastClaim.claimedAt,
        node,
      );

      const pendingClaim = await contracts.rewardsManager.calculatePendingClaim(
        stakerKey,
        '0xb521c2a19df3b949acb766de987efc6787584570',
        '0xFFFFfFff0000000000000000000000000000039C',
      );

      const initialActiveRewardPool =
        await contracts.rewardsManager.getInitialActiveRewardPool(
          stakerKey,
          '0xb521c2a19df3b949acb766de987efc6787584570',
        );

      console.log('current epoch: ', currentEpoch);
      console.log(
        `last claim: ${lastClaim.claimedAt}, ${ethers.formatEther(
          lastClaim.stake,
        )}`,
      );
      console.log(`initial claim: ${ethers.formatEther(initialClaim)}`);
      console.log('final crf: ', Number(finalCrf) / 9223372036854775807);
      console.log(
        `last claim reward pool: ${renderRewardPool(lastClaimRewardPool)}`,
      );
      console.log(
        `initial active reward pool: ${renderRewardPool(
          initialActiveRewardPool,
        )}`,
      );

      if (claim > 0n) {
        await contracts.rewardsManager
          .connect(signer)
          .claimStakingRewards(node);
      }
    }
  }).timeout(0);

  it('simulates 10 nodes and 100 stakers over several epochs', async () => {});
});

function renderRewardPool(r: IRewardsManager.RewardPoolStructOutput) {
  return `{ reward: ${ethers.formatEther(
    r.stakersRewardTotal,
  )}, stake: ${ethers.formatEther(r.totalActiveStake)}, icrf: ${
    Number(r.initialCumulativeRewardFactor) / 9223372036854775807
  }}`;
}
