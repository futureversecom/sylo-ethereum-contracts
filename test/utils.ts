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
