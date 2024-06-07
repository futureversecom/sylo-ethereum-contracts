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
};

export async function deployContracts(
  opts: DeploymentOptions = {},
): Promise<SyloContracts> {
  // Sylo Token
  const SyloTokenFactory = await ethers.getContractFactory('SyloToken');
  const syloToken = await SyloTokenFactory.deploy();

  // Sylo Staking Manager
  const syloStakingManagerOpts = {
    unlockDuration: opts.syloStakingManager?.unlockDuration ?? 10,
  };
  const SyloStakingManagerFactory = await ethers.getContractFactory(
    'SyloStakingManager',
  );
  const syloStakingManager = await SyloStakingManagerFactory.deploy();

  // Seeker Stats Oracle
  const seekerStatsOracleOpts = {
    oracleAccount:
      opts.seekerStatsOralce?.oracleAccount ??
      '0xd9D6945dfe8c1C7aFaFcDF8bf1D1c5beDfeccABF',
  };
  const seekerStatsOracleFactory = await ethers.getContractFactory(
    'SeekerStatsOracle',
  );
  const seekerStatsOracle = await seekerStatsOracleFactory.deploy();

  // Seekers
  const SeekersFactory = await ethers.getContractFactory('TestSeekers');
  const seekers = await SeekersFactory.deploy();

  // Seeker Staking Manager
  const seekerStakingManagerFactor = await ethers.getContractFactory(
    'SeekerStakingManager',
  );
  const seekerStakingManager = await seekerStakingManagerFactor.deploy();

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

  return {
    syloToken,
    syloStakingManager,
    seekerStatsOracle,
    seekerStakingManager,
    seekers,
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
