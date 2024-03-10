import { ethers } from 'hardhat';
import { Signer } from 'ethers';
import { SyloToken } from '../typechain-types';
import utils from './utils';
import { createWinningTicket } from './payments/utils';
import { SyloContracts } from '../common/contracts';

describe('Simulations', () => {
  if (!process.env.RUN_SIMULATIONS) {
    return;
  }

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

  it.only('simulates 10 nodes and 25 stakers over several epochs', async () => {
    const contracts = await utils.initializeContracts(owner, token);

    // create 10 nodes accounts
    const nodes = await Promise.all(
      Array(10)
        .fill(0)
        .map(_ => {
          return createSigner(contracts);
        }),
    );

    console.log('created nodes');

    // create 25 staker accounts
    const stakers = await Promise.all(
      Array(25)
        .fill(0)
        .map((_, i) => {
          return createSigner(contracts, (i % 5) + 1);
        }),
    );

    console.log('created stakers');

    // set the initial node stakes and registrations
    await Promise.all(
      nodes.map(async (node, i) => {
        await contracts.stakingManager
          .connect(node.signer)
          .addStake(ethers.parseEther('1000'), node.address);

        await utils.setSeekerRegistry(
          contracts.registries,
          contracts.seekers,
          contracts.seekerPowerOracle,
          node.signer,
          node.signer,
          i,
        );
      }),
    );

    console.log('setup nodes');

    const stakeEntries: { [staker: string]: { [stakee: string]: number } } = {};

    // have each staker randomly stake to 5 nodes
    await Promise.all(
      stakers.map(async staker => {
        stakeEntries[staker.address] = {};

        for (let i = 0; i < 5; i++) {
          // select random node
          const node = nodes[randomInt(9)];

          // stake between 1 and 1000000
          const stakeAmount = randomInt(1000000) + 1;

          await contracts.stakingManager
            .connect(staker.signer)
            .addStake(ethers.parseEther(stakeAmount.toString()), node.address);

          stakeEntries[staker.address][node.address] = stakeAmount;
        }
      }),
    );

    console.log('setup stakers');

    // setup bootstrapper
    const bootstrapper = await createSigner(contracts, 9);
    await contracts.syloToken.transfer(
      bootstrapper.address,
      ethers.parseEther('1000000000'),
    );
    await contracts.syloTicketing
      .connect(bootstrapper.signer)
      .depositEscrow(ethers.parseEther('1000000000'), bootstrapper.address);
    await contracts.syloTicketing
      .connect(bootstrapper.signer)
      .depositPenalty(ethers.parseEther('1000'), bootstrapper.address);

    await contracts.ticketingParameters.setFaceValue(
      ethers.parseEther('25000'),
    );
    await contracts.epochsManager.setEpochDuration(1);

    // have nodes join the first epoch
    for (const node of nodes) {
      await contracts.epochsManager.connect(node.signer).joinNextEpoch();
    }
    await contracts.epochsManager.initializeEpoch();

    // run the simulation for 50 epochs
    for (let i = 1; i < 50; i++) {
      console.log(`running epoch ${i}...`);

      await Promise.all(
        stakers.map(async staker => {
          await runRandomStakerAction(contracts, staker, stakeEntries);
        }),
      );

      // bootstrap the nodes
      for (let j = 0; j < 10; j++) {
        // scan a node
        const scannedNode = await contracts.directory.scan(
          randomInt(1000000000),
        );

        const { ticket, redeemerRand, senderSig, receiverSig } =
          await createWinningTicket(
            contracts.syloTicketing,
            contracts.epochsManager,
            bootstrapper.signer,
            bootstrapper.signer,
            scannedNode,
          );

        for (const node of nodes) {
          if (node.address == scannedNode) {
            await contracts.syloTicketing
              .connect(node.signer)
              .redeem(ticket, redeemerRand, senderSig, receiverSig);
            break;
          }
        }
      }

      for (const node of nodes) {
        await contracts.epochsManager.connect(node.signer).joinNextEpoch();
      }
      await contracts.epochsManager.initializeEpoch();

      await Promise.all(
        stakers.map(async staker => {
          await runRandomStakerAction(contracts, staker, stakeEntries);
        }),
      );
    }

    console.log('claim final rewards');

    for (const staker of stakers) {
      console.log('performing claims for: ', staker.address);

      for (const node of Object.keys(stakeEntries[staker.address])) {
        const claim = await contracts.rewardsManager.calculateStakerClaim(
          node,
          staker.address,
        );

        if (claim > 0n) {
          await contracts.rewardsManager
            .connect(staker.signer)
            .claimStakingRewards(node);
        }
      }
    }
  }).timeout(0);

  const randomInt = (max: number) => {
    return Math.floor(Math.random() * max);
  };

  // helper function to create a new funder signer
  const createSigner = async (contracts: SyloContracts, funder?: number) => {
    const signer = ethers.Wallet.createRandom(ethers.provider);
    const address = await signer.getAddress();

    await accounts[funder ?? 0].sendTransaction({
      to: address,
      value: ethers.parseEther('100'),
    });

    await contracts.syloToken.transfer(address, ethers.parseEther('100000000'));

    await contracts.syloToken
      .connect(signer)
      .approve(
        await contracts.syloTicketing.getAddress(),
        ethers.parseEther('10000000000'),
      );
    await contracts.syloToken
      .connect(signer)
      .approve(
        await contracts.stakingManager.getAddress(),
        ethers.parseEther('10000000000'),
      );

    return { address, signer };
  };

  // for each of a node's stakers, run an action that will
  // affect their stake or claim amount
  const runRandomStakerAction = async (
    contracts: SyloContracts,
    staker: { address: string; signer: Signer },
    stakeEntries: { [staker: string]: { [stakee: string]: number } },
  ) => {
    for (const node of Object.keys(stakeEntries[staker.address])) {
      const randAction = randomInt(3);

      switch (randAction) {
        case 0: // add stake
          const stakeAmount = randomInt(100000) + 1;

          await contracts.stakingManager
            .connect(staker.signer)
            .addStake(ethers.parseEther(stakeAmount.toString()), node);

          stakeEntries[staker.address][node] += stakeAmount;

          break;
        case 1: // unlock state
          const currentStake = stakeEntries[staker.address][node];
          if (currentStake > 0) {
            const unlockAmount = randomInt(currentStake) + 1;

            await contracts.stakingManager
              .connect(staker.signer)
              .unlockStake(ethers.parseEther(unlockAmount.toString()), node);

            stakeEntries[staker.address][node] -= unlockAmount;
          }

          break;
        case 2: // cancel unlocking
          const key = await contracts.stakingManager.getKey(
            node,
            staker.address,
          );
          const unlocking = await contracts.stakingManager.unlockings(key);

          if (unlocking.amount > 0) {
            await contracts.stakingManager
              .connect(staker.signer)
              .cancelUnlocking(unlocking.amount, node);

            const stakeAmount = parseInt(ethers.formatEther(unlocking.amount));

            stakeEntries[staker.address][node] += stakeAmount;
          }

          break;
        case 3: // claim
          console.log('claim');
          const claim = await contracts.rewardsManager.calculateStakerClaim(
            node,
            staker.address,
          );
          if (claim > 0n) {
            await contracts.rewardsManager
              .connect(staker.signer)
              .claimStakingRewards(node);
          }

          break;
      }
    }
  };
});
