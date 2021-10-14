import { ethers, upgrades, network } from 'hardhat';
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
  const [deployer] = await ethers.getSigners();
  console.log(`Beginning deployment of phase two contracts to ${network.name} by deployer: ${deployer.address}`);

  // If there doesn't exist an address for the Sylo Token for this configuration,
  // then let's deploy the token contract and set the address
  if (config.SyloToken == "") {
    const SyloToken = await ethers.getContractFactory("SyloToken");
    const syloToken = await SyloToken.deploy();
    config.SyloToken = syloToken.address;
    console.log(`Sylo token deployed to ${network.name} network at ${syloToken.address}`);
  }

  const Listings = await ethers.getContractFactory("Listings");
  const listings = await upgrades.deployProxy(
    Listings,
    [config.Listings.defaultPayoutPercentage]
  ) as Listings;

  logDeployment("Listings", listings.address);

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

  logDeployment("TicketingParameters", ticketingParameters.address);

  // We disable automatic calling of the initializer functions as there are
  // cyclic dependencies, so we need to call those functions manually after
  // deployment.

  const EpochsManager = await ethers.getContractFactory("EpochsManager");
  const epochsManager = await upgrades.deployProxy(
    EpochsManager,
    undefined,
    { initializer: false }
  ) as EpochsManager;

  logDeployment("EpochsManager", epochsManager.address);

  const StakingManager = await ethers.getContractFactory("StakingManager");
  const stakingManager = await upgrades.deployProxy(
    StakingManager,
    undefined,
    { initializer: false }
  ) as StakingManager;

  logDeployment("StakingManager", stakingManager.address);

  const RewardsManager = await ethers.getContractFactory("RewardsManager");
  const rewardsManager = await upgrades.deployProxy(
    RewardsManager,
    undefined,
    { initializer: false }
  ) as RewardsManager;

  logDeployment("RewardsManager", rewardsManager.address);

  const Directory = await ethers.getContractFactory("Directory");
  const directory = await upgrades.deployProxy(
    Directory,
    undefined,
    { initializer: false }
  ) as Directory;

  logDeployment("Directory", directory.address);

  await epochsManager.initialize(
    directory.address,
    listings.address,
    ticketingParameters.address,
    config.EpochsManager.epochDuration
  );

  console.log('Initialized epochs manager contract');

  await stakingManager.initialize(
    config.SyloToken,
    rewardsManager.address,
    epochsManager.address,
    config.StakingManager.unlockDuration,
  );

  console.log('Initialized staking manager contract');

  await rewardsManager.initialize(
    config.SyloToken,
    stakingManager.address,
    epochsManager.address
  );

  console.log('Initialized rewards manager contract');

  await directory.initialize(
    stakingManager.address,
    rewardsManager.address
  );

  console.log('Initialized directory contract');

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

  logDeployment("Ticketing", ticketing.address);

  // add managers to the rewards manager contract
  await rewardsManager.addManager(ticketing.address);
  await rewardsManager.addManager(stakingManager.address);

  console.log('Aadded managers to rewards manager contract');

  // set directory ownership to the epochs manager
  await directory.transferOwnership(epochsManager.address);

  console.log('Transferred directory ownership to epochs manager');

  console.log(`Deployment of phase two contracts to network ${network.name} complete by deployer: ${deployer.address}`);

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

function logDeployment(contract: string, contractAddress: string) {
  console.log(`Deployed ${contract} at ${contractAddress}`);
}

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
