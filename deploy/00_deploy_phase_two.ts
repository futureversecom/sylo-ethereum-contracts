import { ethers, upgrades } from 'hardhat';
import Config from './genesis.config';
import { Directory, EpochsManager, Listings, RewardsManager, StakingManager, SyloTicketing, TicketingParameters } from '../typechain';

type PhaseTwoContracts = {
  directory: Directory,
  epochsManager: EpochsManager,
  listings: Listings,
  rewardsManager: RewardsManager,
  stakingManager: StakingManager,
  ticketingParameters: TicketingParameters,
  ticketing: SyloTicketing
};

async function deployPhaseTwoContracts(config: typeof Config): Promise<PhaseTwoContracts> {
  const Listings = await ethers.getContractFactory("Listings");
  const listings = await upgrades.deployProxy(
    Listings,
    [config.Listings.defaultPayoutPercentage]
  ) as Listings;

  const TicketingParameters = await ethers.getContractFactory("TicketingParameters");
  const ticketingParameters = await upgrades.deployProxy(
    TicketingParameters,
    [ config.TicketingParameters.faceValue,
      config.TicketingParameters.baseLiveWinProb,
      config.TicketingParameters.expiredWinProb,
      config.TicketingParameters.decayRate,
      config.TicketingParameters.ticketDuration
    ]
  ) as TicketingParameters;

  // We disable automatic calling of the initializer functions as there are
  // cyclic dependencies, so we need to call those functions manually after
  // deployment.

  const EpochsManager = await ethers.getContractFactory("EpochsManager");
  const epochsManager = await upgrades.deployProxy(
    EpochsManager,
    undefined,
    { initializer: false }
  ) as EpochsManager;

  const StakingManager = await ethers.getContractFactory("StakingManager");
  const stakingManager = await upgrades.deployProxy(
    StakingManager,
    undefined,
    { initializer: false }
  ) as StakingManager;

  const RewardsManager = await ethers.getContractFactory("RewardsManager");
  const rewardsManager = await upgrades.deployProxy(
    RewardsManager,
    undefined,
    { initializer: false }
  ) as RewardsManager;

  const Directory = await ethers.getContractFactory("Directory");
  const directory = await upgrades.deployProxy(
    Directory,
    undefined,
    { initializer: false }
  ) as Directory;

  await epochsManager.initialize(
    directory.address,
    listings.address,
    ticketingParameters.address,
    config.EpochsManager.epochDuration
  );

  await stakingManager.initialize(
    config.SyloToken,
    rewardsManager.address,
    epochsManager.address,
    config.StakingManager.unlockDuration,
  );

  await rewardsManager.initialize(
    config.SyloToken,
    stakingManager.address,
    epochsManager.address
  );

  await directory.initialize(
    stakingManager.address,
    rewardsManager.address
  );

  const Ticketing = await ethers.getContractFactory("SyloTicketing");
  const ticketing = await upgrades.deployProxy(
    Ticketing,
    [ config.SyloToken,
      listings.address,
      stakingManager.address,
      directory.address,
      epochsManager.address,
      rewardsManager.address,
      config.Ticketing.unlockDuration
    ]
  ) as SyloTicketing;

  // add managers to the rewards manager contract
  await rewardsManager.addManager(ticketing.address);
  await rewardsManager.addManager(stakingManager.address);

  // set directory ownership to the epochs manager
  await directory.transferOwnership(epochsManager.address);

  return {
    listings,
    ticketing,
    ticketingParameters,
    directory,
    rewardsManager,
    epochsManager,
    stakingManager
  }
};

async function main() {
  await deployPhaseTwoContracts(Config);
}

export {
  deployPhaseTwoContracts
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
