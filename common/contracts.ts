import { ethers } from 'ethers';
import * as factories from '../typechain-types';

export const FixedContractNames = {
  syloToken: 'SyloToken',
};

export const DeployedContractNames = {
  syloStakingManager: 'SyloStakingManager',
  seekerStatsOracle: 'SeekerStatsOracle',
};

export const ContractNames = {
  ...FixedContractNames,
  ...DeployedContractNames,
};

export type SyloContracts = {
  syloToken: factories.contracts.SyloToken;
  syloStakingManager: factories.contracts.staking.sylo.SyloStakingManager;
  seekerStatsOracle: factories.contracts.staking.seekers.SeekerStatsOracle;
};

export type ContractAddresses = {
  syloToken: string;
  syloStakingManager: string;
  seekerStatsOracle: string;
};

export function connectContracts(
  contracts: ContractAddresses,
  provider: ethers.ContractRunner,
): SyloContracts {
  const syloToken = factories.SyloToken__factory.connect(
    contracts.syloToken,
    provider,
  );

  const syloStakingManager = factories.SyloStakingManager__factory.connect(
    contracts.syloStakingManager,
    provider,
  );

  const seekerStatsOracle = factories.SeekerStatsOracle__factory.connect(
    contracts.seekerStatsOracle,
    provider,
  );

  return {
    syloToken,
    syloStakingManager,
    seekerStatsOracle,
  };
}
