import { ethers } from 'hardhat';
import { BigNumber, BigNumberish, Signer } from 'ethers';
import { toWei } from 'web3-utils';
import {
  Directory,
  EpochsManager,
  Registries,
  RewardsManager,
  StakingManager,
  SyloTicketing,
  TicketingParameters,
  TestSeekers,
} from '../typechain';

type Options = {
  faceValue?: BigNumberish;
  payoutPercentage?: number;
  baseLiveWinProb?: BigNumberish;
  expiredWinProb?: BigNumberish;
  decayRate?: number;
  ticketDuration?: number;
  epochDuration?: number;
  minimumStakeProportion?: number;
  unlockDuration?: number;
};

export type Contracts = {
  registries: Registries;
  ticketing: SyloTicketing;
  ticketingParameters: TicketingParameters;
  directory: Directory;
  rewardsManager: RewardsManager;
  epochsManager: EpochsManager;
  stakingManager: StakingManager;
  seekers: TestSeekers;
};

const initializeContracts = async function (
  deployer: string,
  tokenAddress: string,
  opts: Options = {},
): Promise<Contracts> {
  const payoutPercentage = opts.payoutPercentage ? opts.payoutPercentage : 5000;

  const faceValue = opts.faceValue ?? toWei('15');
  const baseLiveWinProb =
    opts.baseLiveWinProb ?? BigNumber.from(2).pow(128).sub(1);
  const expiredWinProb = opts.expiredWinProb ?? 1000;
  const decayRate = opts.decayRate ?? 8000;
  const ticketDuration = opts.ticketDuration ?? 20;

  const epochDuration = opts.epochDuration ?? 30;

  const unlockDuration = opts.unlockDuration ?? 10;

  const minimumStakeProportion = opts.minimumStakeProportion ?? 2000;

  const SeekersFactory = await ethers.getContractFactory('TestSeekers');
  const seekers = await SeekersFactory.deploy();

  const RegistriesFactory = await ethers.getContractFactory('Registries');
  const registries = await RegistriesFactory.deploy();
  await registries.initialize(seekers.address, payoutPercentage, 100, {
    from: deployer,
  });

  const TicketingParametersFactory = await ethers.getContractFactory(
    'TicketingParameters',
  );
  const ticketingParameters = await TicketingParametersFactory.deploy();
  await ticketingParameters.initialize(
    faceValue,
    baseLiveWinProb,
    expiredWinProb,
    decayRate,
    ticketDuration,
    { from: deployer },
  );

  const EpochsManagerFactory = await ethers.getContractFactory('EpochsManager');
  const epochsManager = await EpochsManagerFactory.deploy();

  const StakingManagerFactory = await ethers.getContractFactory(
    'StakingManager',
  );
  const stakingManager = await StakingManagerFactory.deploy();

  const RewardsManagerFactory = await ethers.getContractFactory(
    'RewardsManager',
  );
  const rewardsManager = await RewardsManagerFactory.deploy();

  const DirectoryFactory = await ethers.getContractFactory('Directory');
  const directory = await DirectoryFactory.deploy();

  await stakingManager.initialize(
    tokenAddress,
    rewardsManager.address,
    epochsManager.address,
    unlockDuration,
    minimumStakeProportion,
    { from: deployer },
  );
  await rewardsManager.initialize(
    tokenAddress,
    stakingManager.address,
    epochsManager.address,
    { from: deployer },
  );
  await directory.initialize(stakingManager.address, rewardsManager.address, {
    from: deployer,
  });
  await epochsManager.initialize(
    seekers.address,
    directory.address,
    registries.address,
    ticketingParameters.address,
    epochDuration,
    { from: deployer },
  );

  const TicketingFactory = await ethers.getContractFactory('SyloTicketing');
  const ticketing = await TicketingFactory.deploy();
  await ticketing.initialize(
    tokenAddress,
    registries.address,
    stakingManager.address,
    directory.address,
    epochsManager.address,
    rewardsManager.address,
    unlockDuration,
    { from: deployer },
  );

  await rewardsManager.addManager(ticketing.address, { from: deployer });
  await rewardsManager.addManager(stakingManager.address, { from: deployer });
  await rewardsManager.addManager(epochsManager.address, { from: deployer });

  await directory.addManager(epochsManager.address);

  return {
    registries,
    ticketing,
    ticketingParameters,
    directory,
    rewardsManager,
    epochsManager,
    stakingManager,
    seekers,
  };
};

const advanceBlock = async function (i: number): Promise<void> {
  i = i || 1;
  for (let j = 0; j < i; j++) {
    await ethers.provider.send('evm_mine', []);
  }
};

async function setSeekerRegistry(
  registries: Registries,
  seekers: TestSeekers,
  account: Signer,
  seekerAccount: Signer,
  tokenId: number,
): Promise<void> {
  if (!(await seekers.exists(tokenId))) {
    await seekers.mint(await seekerAccount.getAddress(), tokenId);
  }

  const block = await ethers.provider.getBlockNumber();

  const prefix = await registries.getPrefix();
  const accountAddress = await account.getAddress();
  const proofMessage = `${prefix}:${tokenId}:${accountAddress.toLowerCase()}:${block.toString()}`;

  const signature = await seekerAccount.signMessage(proofMessage);

  await registries.connect(account).register('0.0.0.0/0', 1);

  await registries
    .connect(account)
    .setSeekerAccount(
      await seekerAccount.getAddress(),
      tokenId,
      block,
      signature,
    );
}

export default {
  initializeContracts,
  advanceBlock,
  setSeekerRegistry,
};
