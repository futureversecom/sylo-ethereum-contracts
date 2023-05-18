import { ethers, upgrades, network } from 'hardhat';
import Config from './genesis.config';
import {
  AuthorizedAccount,
  Directory,
  EpochsManager,
  Registries,
  RewardsManager,
  StakingManager,
  SyloTicketing,
  TicketingParameters,
} from '../typechain-types';
import * as fs from 'fs/promises';

type PhaseTwoContracts = {
  token: string;
  directory: Directory;
  epochsManager: EpochsManager;
  authorizedAccount: AuthorizedAccount;
  registries: Registries;
  rewardsManager: RewardsManager;
  stakingManager: StakingManager;
  ticketingParameters: TicketingParameters;
  ticketing: SyloTicketing;
  seekers: string;
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
    await syloToken.deployed();

    config.SyloToken = syloToken.address;
    console.log(
      `Sylo token deployed to ${network.name} network at ${syloToken.address}`,
    );
  }

  if (config.Seekers == '') {
    const TestSeekerFactory = await ethers.getContractFactory('TestSeekers');
    const seekers = await TestSeekerFactory.deploy();
    await seekers.deployed();

    config.Seekers = seekers.address;
    logDeployment('Seekers', seekers.address);
  }

  const AuthorizedAccountFactory = await ethers.getContractFactory(
    'AuthorizedAccount',
  );
  const authorizeAccount = (await upgrades.deployProxy(
    AuthorizedAccountFactory,
    undefined,
    {
      initializer: false,
    },
  )) as AuthorizedAccount;
  await authorizeAccount.deployed();

  logDeployment('AuthorizedAccount', authorizeAccount.address);

  const RegistriesFactory = await ethers.getContractFactory('Registries');
  const registries = (await upgrades.deployProxy(RegistriesFactory, undefined, {
    initializer: false,
  })) as Registries;
  await registries.deployed();

  logDeployment('Registries', registries.address);

  const TicketingParametersFactory = await ethers.getContractFactory(
    'TicketingParameters',
  );
  const ticketingParameters = (await upgrades.deployProxy(
    TicketingParametersFactory,
    undefined,
    { initializer: false },
  )) as TicketingParameters;
  await ticketingParameters.deployed();

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
  await epochsManager.deployed();

  logDeployment('EpochsManager', epochsManager.address);

  const StakingManagerFactory = await ethers.getContractFactory(
    'StakingManager',
  );
  const stakingManager = (await upgrades.deployProxy(
    StakingManagerFactory,
    undefined,
    { initializer: false },
  )) as StakingManager;
  await stakingManager.deployed();

  logDeployment('StakingManager', stakingManager.address);

  const RewardsManagerFactory = await ethers.getContractFactory(
    'RewardsManager',
  );
  const rewardsManager = (await upgrades.deployProxy(
    RewardsManagerFactory,
    undefined,
    { initializer: false },
  )) as RewardsManager;
  await rewardsManager.deployed();

  logDeployment('RewardsManager', rewardsManager.address);

  const DirectoryFactory = await ethers.getContractFactory('Directory');
  const directory = (await upgrades.deployProxy(DirectoryFactory, undefined, {
    initializer: false,
  })) as Directory;
  await directory.deployed();

  logDeployment('Directory', directory.address);

  await registries
    .initialize(config.Seekers, config.Registries.defaultPayoutPercentage)
    .then(tx => tx.wait());

  console.log('Initialized registries');

  await authorizeAccount.initialize().then(tx => tx.wait());

  console.log('Initialized authorized account');

  await ticketingParameters
    .initialize(
      config.TicketingParameters.faceValue,
      config.TicketingParameters.baseLiveWinProb,
      config.TicketingParameters.expiredWinProb,
      config.TicketingParameters.decayRate,
      config.TicketingParameters.ticketDuration,
    )
    .then(tx => tx.wait());

  console.log('Initialized ticketing parameters');

  await epochsManager
    .initialize(
      config.Seekers,
      directory.address,
      registries.address,
      ticketingParameters.address,
      config.EpochsManager.epochDuration,
    )
    .then(tx => tx.wait());

  console.log('Initialized epochs manager contract');

  await stakingManager
    .initialize(
      config.SyloToken,
      rewardsManager.address,
      epochsManager.address,
      config.StakingManager.unlockDuration,
      config.StakingManager.minimumStakeProportion,
    )
    .then(tx => tx.wait());

  console.log('Initialized staking manager contract');

  await rewardsManager
    .initialize(config.SyloToken, stakingManager.address, epochsManager.address)
    .then(tx => tx.wait());

  console.log('Initialized rewards manager contract');

  await directory
    .initialize(stakingManager.address, rewardsManager.address)
    .then(tx => tx.wait());

  console.log('Initialized directory contract');

  const TicketingFactory = await ethers.getContractFactory('SyloTicketing');
  const ticketing = (await upgrades.deployProxy(TicketingFactory, undefined, {
    initializer: false,
  })) as SyloTicketing;
  await ticketing.deployed();

  logDeployment('Ticketing', ticketing.address);

  await ticketing
    .initialize(
      config.SyloToken,
      registries.address,
      stakingManager.address,
      directory.address,
      epochsManager.address,
      rewardsManager.address,
      authorizeAccount.address,
      config.Ticketing.unlockDuration,
    )
    .then(tx => tx.wait());

  console.log('Initialized ticketing contract');

  // add managers to the staking manager contract
  await directory.addManager(epochsManager.address).then(tx => tx.wait());

  console.log('Added managers to directory manager contract');

  // add managers to the rewards manager contract
  await rewardsManager.addManager(ticketing.address).then(tx => tx.wait());

  await rewardsManager.addManager(stakingManager.address).then(tx => tx.wait());

  await rewardsManager.addManager(epochsManager.address).then(tx => tx.wait());

  console.log('Added managers to rewards manager contract');

  console.log(
    `Deployment of phase two contracts to network ${network.name} complete by deployer: ${deployer.address}`,
  );

  return {
    token: config.SyloToken,
    authorizeAccount,
    registries,
    ticketing,
    ticketingParameters,
    directory,
    rewardsManager,
    epochsManager,
    stakingManager,
    seekers: config.Seekers,
  };
}

function logDeployment(contract: string, contractAddress: string) {
  console.log(`Deployed ${contract} at ${contractAddress}`);
}

async function main() {
  const contracts = await deployPhaseTwoContracts(Config);

  const [deployer] = await ethers.getSigners();

  // write the deployed contracts to a json file
  // this is easier to read than the openzeppelin manifest
  const deployedJson = {
    deployer: deployer.address,
    token: contracts.token,
    authorizedAccount: contracts.authorizeAccount.address,
    registries: contracts.registries.address,
    ticketing: contracts.ticketing.address,
    ticketingParameters: contracts.ticketingParameters.address,
    directory: contracts.directory.address,
    rewardsManager: contracts.rewardsManager.address,
    epochsManager: contracts.epochsManager.address,
    stakingManager: contracts.stakingManager.address,
    seekers: contracts.seekers,
  };

  await fs.writeFile(
    `${__dirname}/${network.name}_deployment_phase_two.json`,
    Buffer.from(JSON.stringify(deployedJson, null, ' '), 'utf8'),
  );
}

export { deployPhaseTwoContracts };

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error);
    process.exit(1);
  });
