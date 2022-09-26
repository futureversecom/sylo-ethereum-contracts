import { ethers, upgrades, network } from 'hardhat';
import Config from './genesis.config';
import {
  Directory,
  EpochsManager,
  Registries,
  Registries__factory,
  RewardsManager,
  Seekers,
  StakingManager,
  SyloTicketing,
  TicketingParameters,
} from '../typechain';

type PhaseTwoContracts = {
  token: string;
  directory: Directory;
  epochsManager: EpochsManager;
  registries: Registries;
  rewardsManager: RewardsManager;
  stakingManager: StakingManager;
  ticketingParameters: TicketingParameters;
  ticketing: SyloTicketing;
  seekers: Seekers;
};

async function deployPhaseTwoContracts(
  config: typeof Config,
): Promise<PhaseTwoContracts> {
  const [deployer] = await ethers.getSigners();
  console.log(
    `Beginning deployment of phase two contracts to ${network.name} by deployer: ${deployer.address}`,
  );

  // If there doesn't exist an address for the Sylo Token for this configuration,
  // then let's deploy the token contract and set the address
  if (config.SyloToken == '') {
    const SyloTokenFactory = await ethers.getContractFactory('SyloToken');
    const syloToken = await SyloTokenFactory.deploy();
    config.SyloToken = syloToken.address;
    console.log(
      `Sylo token deployed to ${network.name} network at ${syloToken.address}`,
    );
  }

  const SeekersFactory = await ethers.getContractFactory('Seekers');
  const seekers = (await upgrades.deployProxy(SeekersFactory, [
    config.Seekers.seekersERC721,
    config.SyloToken,
    config.Seekers.oracle,
    config.Seekers.validDuration,
    config.Seekers.callbackGasLimit,
    config.Seekers.callbackBounty,
  ])) as Seekers;

  logDeployment('Seekers', seekers.address);

  const RegistriesFactory = await ethers.getContractFactory('Registries');
  const registries = (await upgrades.deployProxy(RegistriesFactory, [
    seekers.address,
    config.Registries.defaultPayoutPercentage,
    config.Registries.proofDuration,
  ])) as Registries;

  logDeployment('Registries', registries.address);

  const TicketingParametersFactory = await ethers.getContractFactory(
    'TicketingParameters',
  );
  const ticketingParameters = (await upgrades.deployProxy(
    TicketingParametersFactory,
    [
      config.TicketingParameters.faceValue,
      config.TicketingParameters.baseLiveWinProb,
      config.TicketingParameters.expiredWinProb,
      config.TicketingParameters.decayRate,
      config.TicketingParameters.ticketDuration,
    ],
  )) as TicketingParameters;

  logDeployment('TicketingParameters', ticketingParameters.address);

  // We disable automatic calling of the initializer functions as there are
  // cyclic dependencies, so we need to call those functions manually after
  // deployment.

  const EpochsManagerFactory = await ethers.getContractFactory('EpochsManager');
  const epochsManager = (await upgrades.deployProxy(
    EpochsManagerFactory,
    undefined,
    {
      initializer: false,
    },
  )) as EpochsManager;

  logDeployment('EpochsManager', epochsManager.address);

  const StakingManagerFactory = await ethers.getContractFactory(
    'StakingManager',
  );
  const stakingManager = (await upgrades.deployProxy(
    StakingManagerFactory,
    undefined,
    { initializer: false },
  )) as StakingManager;

  logDeployment('StakingManager', stakingManager.address);

  const RewardsManagerFactory = await ethers.getContractFactory(
    'RewardsManager',
  );
  const rewardsManager = (await upgrades.deployProxy(
    RewardsManagerFactory,
    undefined,
    { initializer: false },
  )) as RewardsManager;

  logDeployment('RewardsManager', rewardsManager.address);

  const DirectoryFactory = await ethers.getContractFactory('Directory');
  const directory = (await upgrades.deployProxy(DirectoryFactory, undefined, {
    initializer: false,
  })) as Directory;

  logDeployment('Directory', directory.address);

  await epochsManager.initialize(
    seekers.address,
    directory.address,
    registries.address,
    ticketingParameters.address,
    config.EpochsManager.epochDuration,
  );

  console.log('Initialized epochs manager contract');

  await stakingManager.initialize(
    config.SyloToken,
    rewardsManager.address,
    epochsManager.address,
    config.StakingManager.unlockDuration,
    config.StakingManager.minimumStakeProportion,
  );

  console.log('Initialized staking manager contract');

  await rewardsManager.initialize(
    config.SyloToken,
    stakingManager.address,
    epochsManager.address,
  );

  console.log('Initialized rewards manager contract');

  await directory.initialize(stakingManager.address, rewardsManager.address);

  console.log('Initialized directory contract');

  const TicketingFactory = await ethers.getContractFactory('SyloTicketing');
  const ticketing = (await upgrades.deployProxy(TicketingFactory, [
    config.SyloToken,
    registries.address,
    stakingManager.address,
    directory.address,
    epochsManager.address,
    rewardsManager.address,
    config.Ticketing.unlockDuration,
  ])) as SyloTicketing;

  logDeployment('Ticketing', ticketing.address);

  // add managers to the staking manager contract
  await directory.addManager(epochsManager.address);

  console.log('Added managers to directory manager contract');

  // add managers to the rewards manager contract
  await rewardsManager.addManager(ticketing.address);
  await rewardsManager.addManager(stakingManager.address);
  await rewardsManager.addManager(epochsManager.address);

  console.log('Added managers to rewards manager contract');

  console.log(
    `Deployment of phase two contracts to network ${network.name} complete by deployer: ${deployer.address}`,
  );

  return {
    token: config.SyloToken,
    registries,
    ticketing,
    ticketingParameters,
    directory,
    rewardsManager,
    epochsManager,
    stakingManager,
    seekers,
  };
}

function logDeployment(contract: string, contractAddress: string) {
  console.log(`Deployed ${contract} at ${contractAddress}`);
}

async function main() {
  const [deployer] = await ethers.getSigners();
  console.log(deployer.address);

  const registries = Registries__factory.connect(
    '0x3db64F253Ae1d4bbe31FcD533704D7E64Ef28494',
    deployer,
  );

  await registries.register('http://54.252.224.232/public/metadata', 1);
  console.log('Registered metadata');
}

export { deployPhaseTwoContracts };

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error);
    process.exit(1);
  });
