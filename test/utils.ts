import { ethers } from 'hardhat';
import { BigNumber, BigNumberish, Signer } from 'ethers';
import { toWei } from 'web3-utils';
import {
  Directory,
  EpochsManager,
  MockOracle,
  Registries,
  RewardsManager,
  Seekers,
  StakingManager,
  SyloTicketing,
  TicketingParameters,
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
  seekers: Seekers;
  mockOracle: MockOracle;
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

  const MockOracleFactory = await ethers.getContractFactory('MockOracle');
  const mockOracle = await MockOracleFactory.deploy();

  const SeekersFactory = await ethers.getContractFactory('Seekers');
  const seekers = await SeekersFactory.deploy();
  await seekers.initialize(
    ethers.constants.AddressZero,
    tokenAddress,
    mockOracle.address,
    100,
    300000,
    ethers.utils.parseEther('2'),
    {
      from: deployer,
    },
  );

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
    mockOracle,
  };
};

const advanceBlock = async function (i: number): Promise<void> {
  i = i || 1;
  for (let j = 0; j < i; j++) {
    await ethers.provider.send('evm_mine', []);
  }
};

const setSeekerOwnership = async function (
  mockOracle: MockOracle,
  seekers: Seekers,
  tokenId: number,
  owner: string,
): Promise<void> {
  await mockOracle.setOwner(tokenId, owner);
  await seekers.requestVerification(tokenId);
  await mockOracle.invokeCallback();
};

async function setSeekerRegistry(
  registries: Registries,
  mockOracle: MockOracle,
  seekers: Seekers,
  account: Signer,
  seekerAccount: Signer,
  tokenId: number,
): Promise<void> {
  await setSeekerOwnership(
    mockOracle,
    seekers,
    tokenId,
    await seekerAccount.getAddress(),
  );

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
  setSeekerOwnership,
  setSeekerRegistry,
};
