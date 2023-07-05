import { HardhatRuntimeEnvironment } from 'hardhat/types';
import {
  DeployFunction,
  DeployOptions,
  DeployResult,
  Receipt,
  TxOptions,
} from 'hardhat-deploy/types';
import path from 'path';
import * as fs from 'fs/promises';
import { ethers, network } from 'hardhat';
import * as configs from '../deployments/genesis.config';
import { ContractNames, DeployedContractNames } from '../common/contracts';

type ContractParams = {
  name: string;
  args: unknown[];
};

type ContractMap = {
  [key: string]: DeployResult;
};

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const [deployer] = await ethers.getSigners();
  const deploy = hre.deployments.deploy.bind(hre.deployments);
  const execute = hre.deployments.execute.bind(hre.deployments);

  const config = getConfig(network.name);

  const contracts: ContractMap = {};

  // DEPLOY CONTRACTS
  if (config.SyloToken == '') {
    config.SyloToken = (
      await deployContract('SyloToken', deployer.address, false, deploy)
    ).address;
  }
  if (config.Seekers == '') {
    config.Seekers = (
      await deployContract('TestSeekers', deployer.address, false, deploy)
    ).address;
  }
  for (const name of Object.values(DeployedContractNames)) {
    contracts[name] = await deployContract(
      name,
      deployer.address,
      true,
      deploy,
    );
  }

  // INITIALIZE CONTRACTS
  const initializeParams: ContractParams[] = [
    {
      name: ContractNames.authorizedAccounts,
      args: [],
    },
    {
      name: ContractNames.registries,
      args: [config.Seekers, config.Registries.defaultPayoutPercentage],
    },
    {
      name: ContractNames.ticketingParameters,
      args: [
        config.TicketingParameters.faceValue,
        config.TicketingParameters.baseLiveWinProb,
        config.TicketingParameters.expiredWinProb,
        config.TicketingParameters.decayRate,
        config.TicketingParameters.ticketDuration,
      ],
    },
    {
      name: ContractNames.epochsManager,
      args: [
        config.Seekers,
        contracts[ContractNames.directory].address,
        contracts[ContractNames.registries].address,
        contracts[ContractNames.ticketingParameters].address,
        config.EpochsManager.epochDuration,
      ],
    },
    {
      name: ContractNames.stakingManager,
      args: [
        config.SyloToken,
        contracts[ContractNames.rewardsManager].address,
        contracts[ContractNames.epochsManager].address,
        contracts[ContractNames.stakingManager].address,
        config.StakingManager.minimumStakeProportion,
      ],
    },
    {
      name: ContractNames.rewardsManager,
      args: [
        config.SyloToken,
        contracts[ContractNames.stakingManager].address,
        contracts[ContractNames.epochsManager].address,
      ],
    },
    {
      name: ContractNames.directory,
      args: [
        contracts[ContractNames.stakingManager].address,
        contracts[ContractNames.rewardsManager].address,
      ],
    },
    {
      name: ContractNames.syloTicketing,
      args: [
        config.SyloToken,
        contracts[ContractNames.registries].address,
        contracts[ContractNames.stakingManager].address,
        contracts[ContractNames.directory].address,
        contracts[ContractNames.epochsManager].address,
        contracts[ContractNames.rewardsManager].address,
        contracts[ContractNames.authorizedAccounts].address,
        config.Ticketing.unlockDuration,
      ],
    },
  ];
  for (const { name, args } of initializeParams) {
    await initializeContract(name, args, deployer.address, execute);
  }

  // ADD MANAGERS
  const addManagerParams: ContractParams[] = [
    {
      name: ContractNames.directory,
      args: [contracts[ContractNames.epochsManager].address],
    },
    {
      name: ContractNames.rewardsManager,
      args: [contracts[ContractNames.syloTicketing].address],
    },
    {
      name: ContractNames.rewardsManager,
      args: [contracts[ContractNames.stakingManager].address],
    },
    {
      name: ContractNames.rewardsManager,
      args: [contracts[ContractNames.epochsManager].address],
    },
  ];
  for (const { name, args } of addManagerParams) {
    await addManager(name, deployer.address, args[0] as string, execute);
  }

  await saveContracts(deployer.address, network.name, contracts, config);
};

export default func;

function getConfig(networkName: string): configs.ContractParameters {
  switch (networkName) {
    case 'poricni-dev':
      return configs.PorciniDevParameters;
    case 'poricni-testing':
      return configs.PorciniTestingParameters;
    case 'locahost':
      return configs.GanacheTestnetParameters;
    default:
      return configs.GenesisParameters;
  }
}

async function deployContract(
  contractName: string,
  deployer: string,
  useProxy: boolean,
  deploy: (name: string, options: DeployOptions) => Promise<DeployResult>,
): Promise<DeployResult> {
  const proxy = useProxy
    ? {
        proxyContract: 'OpenZeppelinTransparentProxy',
      }
    : false;

  const result = await deploy(contractName, {
    from: deployer,
    log: true,
    proxy: proxy,
    autoMine: true, // speed up deployment on local network (ganache, hardhat), no effect on live networks
  });

  printEmptyLine();

  return result;
}

async function initializeContract(
  contractName: string,
  args: unknown[],
  deployer: string,
  execute: (
    name: string,
    options: TxOptions,
    methodName: string,
    ...args: unknown[]
  ) => Promise<Receipt>,
): Promise<Receipt> {
  const result = await execute(
    contractName,
    { from: deployer, log: true },
    'initialize',
    ...args,
  );

  printEmptyLine();

  return result;
}

async function addManager(
  contractName: string,
  deployer: string,
  manager: string,
  execute: (
    name: string,
    options: TxOptions,
    methodName: string,
    ...args: unknown[]
  ) => Promise<Receipt>,
): Promise<Receipt> {
  const result = await execute(
    contractName,
    { from: deployer, log: true },
    'addManager',
    manager,
  );

  printEmptyLine();

  return result;
}

async function saveContracts(
  deployer: string,
  networkName: string,
  contracts: ContractMap,
  config: configs.ContractParameters,
) {
  const contractDeployInfo = {
    deployer,
    syloToken: config.SyloToken,
    authorizedAccounts: contracts[ContractNames.authorizedAccounts].address,
    registries: contracts[ContractNames.registries].address,
    ticketingParameters: contracts[ContractNames.ticketingParameters].address,
    epochsManager: contracts[ContractNames.epochsManager].address,
    stakingManager: contracts[ContractNames.stakingManager].address,
    rewardsManager: contracts[ContractNames.rewardsManager].address,
    directory: contracts[ContractNames.directory].address,
    syloTicketing: contracts[ContractNames.syloTicketing].address,
    seekers: config.Seekers,
  };

  const filePath = path.join(process.cwd(), 'deployments');

  await fs.writeFile(
    `${filePath}/${networkName}_deployment_phase_two.json`,
    Buffer.from(JSON.stringify(contractDeployInfo, null, ' '), 'utf8'),
  );
}

function printEmptyLine() {
  console.log('');
}
