import { ethers } from 'hardhat';
import { SyloContracts } from '../common/contracts';
import { Address } from 'hardhat-deploy/types';

export type DeploymentOptions = {
  syloStakingManager?: {
    unlockDuration?: number;
  };
  seekerStatsOralce?: {
    oracleAccount?: Address;
  };
  protocolTimeManager?: {
    cycleDuration?: number;
    periodDuration?: number;
  };
};

export async function deployContracts(
  opts: DeploymentOptions = {},
): Promise<SyloContracts> {
  // Factories
  const syloTokenFactory = await ethers.getContractFactory('SyloToken');
  const syloStakingManagerFactory = await ethers.getContractFactory(
    'SyloStakingManager',
  );
  const seekerStatsOracleFactory = await ethers.getContractFactory(
    'SeekerStatsOracle',
  );
  const seekerStakingManagerFactor = await ethers.getContractFactory(
    'SeekerStakingManager',
  );
  const seekersFactory = await ethers.getContractFactory('TestSeekers');
  const protocolTimeManagerFactory = await ethers.getContractFactory(
    'ProtocolTimeManager',
  );

  // Deploy
  const syloToken = await syloTokenFactory.deploy();
  const syloStakingManager = await syloStakingManagerFactory.deploy();
  const seekerStatsOracle = await seekerStatsOracleFactory.deploy();
  const seekers = await seekersFactory.deploy();
  const seekerStakingManager = await seekerStakingManagerFactor.deploy();
  const protocolTimeManager = await protocolTimeManagerFactory.deploy();

  // Options
  const syloStakingManagerOpts = {
    unlockDuration: opts.syloStakingManager?.unlockDuration ?? 10,
  };
  const seekerStatsOracleOpts = {
    oracleAccount:
      opts.seekerStatsOralce?.oracleAccount ??
      '0xd9D6945dfe8c1C7aFaFcDF8bf1D1c5beDfeccABF',
  };
  const protocolTimeManagerOpts = {
    cycleDuration: opts.protocolTimeManager?.cycleDuration ?? 1000,
    periodDuration: opts.protocolTimeManager?.periodDuration ?? 100,
  };

  // Initliaze
  await syloStakingManager.initialize(
    await syloToken.getAddress(),
    syloStakingManagerOpts.unlockDuration,
  );
  await seekerStatsOracle.initialize(seekerStatsOracleOpts.oracleAccount);
  await seekerStakingManager.initialize(
    await seekers.getAddress(),
    await seekerStatsOracle.getAddress(),
  );
  await protocolTimeManager.initialize(
    protocolTimeManagerOpts.cycleDuration,
    protocolTimeManagerOpts.periodDuration,
  );

  return {
    syloToken,
    syloStakingManager,
    seekerStatsOracle,
    seekerStakingManager,
    seekers,
    protocolTimeManager,
  };
}

export function getInterfaceId(abi: string[]): string {
  const iface = new ethers.Interface(abi);

  const selectors: string[] = [];

  iface.forEachFunction(f => {
    selectors.push(f.selector);
  });

  const interfaceId = selectors.reduce((id, selector) => {
    const selectorBytes = ethers.getBytes(selector);
    const idBytes = ethers.getBytes(id);
    return ethers.hexlify(
      selectorBytes.map((byte, index) => byte ^ idBytes[index]),
    );
  }, '0x00000000');

  return interfaceId;
}
