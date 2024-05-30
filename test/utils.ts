import { ethers } from 'hardhat';
import { SyloContracts } from '../common/contracts';

export type DeploymentOptions = {
  syloStakingManager?: {
    unlockDuration?: number;
  };
};

export async function deployContracts(
  opts: DeploymentOptions = {},
): Promise<SyloContracts> {
  const SyloTokenFactory = await ethers.getContractFactory('SyloToken');
  const syloToken = await SyloTokenFactory.deploy();

  const syloStakingManagerOpts = {
    unlockDuration: opts.syloStakingManager?.unlockDuration ?? 10,
  };

  const SyloStakingManagerFactory = await ethers.getContractFactory(
    'SyloStakingManager',
  );
  const syloStakingManager = await SyloStakingManagerFactory.deploy();

  await syloStakingManager.initialize(
    await syloToken.getAddress(),
    syloStakingManagerOpts.unlockDuration,
  );

  return {
    syloToken,
    syloStakingManager,
  };
}
