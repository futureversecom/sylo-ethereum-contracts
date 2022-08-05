import { ethers } from 'hardhat';
import { BigNumber, BigNumberish, Signer } from 'ethers';
import { toWei } from 'web3-utils';
import {
  Directory,
  EpochsManager,
  Listings,
  MockOracle,
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
  listings: Listings;
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

  const MockOracle = await ethers.getContractFactory('MockOracle');
  const mockOracle = await MockOracle.deploy();

  const Seekers = await ethers.getContractFactory('Seekers');
  const seekers = await Seekers.deploy();
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

  const Listings = await ethers.getContractFactory('Listings');
  const listings = await Listings.deploy();
  await listings.initialize(seekers.address, payoutPercentage, 100, {
    from: deployer,
  });

  const TicketingParameters = await ethers.getContractFactory(
    'TicketingParameters',
  );
  const ticketingParameters = await TicketingParameters.deploy();
  await ticketingParameters.initialize(
    faceValue,
    baseLiveWinProb,
    expiredWinProb,
    decayRate,
    ticketDuration,
    { from: deployer },
  );

  const EpochsManager = await ethers.getContractFactory('EpochsManager');
  const epochsManager = await EpochsManager.deploy();

  const StakingManager = await ethers.getContractFactory('StakingManager');
  const stakingManager = await StakingManager.deploy();

  const RewardsManager = await ethers.getContractFactory('RewardsManager');
  const rewardsManager = await RewardsManager.deploy();

  const Directory = await ethers.getContractFactory('Directory');
  const directory = await Directory.deploy();

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
    listings.address,
    ticketingParameters.address,
    epochDuration,
    { from: deployer },
  );

  const Ticketing = await ethers.getContractFactory('SyloTicketing');
  const ticketing = await Ticketing.deploy();
  await ticketing.initialize(
    tokenAddress,
    listings.address,
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
    listings,
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

async function setSeekerListing(
  listings: Listings,
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

  const prefix = await listings.getPrefix();
  const accountAddress = await account.getAddress();
  const proofMessage = `${prefix}:${tokenId}:${accountAddress.toLowerCase()}:${block.toString()}`;

  const signature = await seekerAccount.signMessage(proofMessage);

  await listings.connect(account).setListing('0.0.0.0/0', 1);

  await listings
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
  setSeekerListing,
};
