import { ethers } from 'ethers';
import * as factories from '../typechain-types';

export const FixedContractNames = {
  syloToken: 'SyloToken',
  seekers: 'Seekers',
};

export const DeployedContractNames = {
  authorizedAccounts: 'AuthorizedAccounts',
  registries: 'Registries',
  ticketingParameters: 'TicketingParameters',
  epochsManager: 'EpochsManager',
  stakingManager: 'StakingManager',
  rewardsManager: 'RewardsManager',
  directory: 'Directory',
  syloTicketing: 'SyloTicketing',
};

export const ContractNames = {
  ...FixedContractNames,
  ...DeployedContractNames,
};

export type SyloContracts = {
  syloToken: factories.contracts.SyloToken;
  authorizedAccounts: factories.contracts.AuthorizedAccounts;
  registries: factories.contracts.Registries;
  ticketingParameters: factories.contracts.payments.ticketing.TicketingParameters;
  epochsManager: factories.contracts.epochs.EpochsManager;
  stakingManager: factories.contracts.staking.StakingManager;
  rewardsManager: factories.contracts.payments.ticketing.RewardsManager;
  directory: factories.contracts.staking.Directory;
  syloTicketing: factories.contracts.payments.SyloTicketing;
  seekers: factories.contracts.mocks.TestSeekers;
  futurepassRegistrar: factories.contracts.mocks.TestFuturepassRegistrar;
};

export type ContractAddresses = {
  syloToken: string;
  authorizedAccounts: string;
  registries: string;
  ticketingParameters: string;
  epochsManager: string;
  stakingManager: string;
  rewardsManager: string;
  directory: string;
  syloTicketing: string;
  seekers: string;
  futurepassRegistrar: string;
};

export function connectContracts(
  contracts: ContractAddresses,
  provider: ethers.JsonRpcProvider,
): SyloContracts {
  const syloToken = factories.SyloToken__factory.connect(
    contracts.syloToken,
    provider,
  );

  const authorizedAccounts = factories.AuthorizedAccounts__factory.connect(
    contracts.authorizedAccounts,
    provider,
  );

  const registries = factories.Registries__factory.connect(
    contracts.registries,
    provider,
  );

  const ticketingParameters = factories.TicketingParameters__factory.connect(
    contracts.ticketingParameters,
    provider,
  );

  const epochsManager = factories.EpochsManager__factory.connect(
    contracts.epochsManager,
    provider,
  );

  const stakingManager = factories.StakingManager__factory.connect(
    contracts.stakingManager,
    provider,
  );

  const rewardsManager = factories.RewardsManager__factory.connect(
    contracts.rewardsManager,
    provider,
  );

  const directory = factories.Directory__factory.connect(
    contracts.directory,
    provider,
  );

  const syloTicketing = factories.SyloTicketing__factory.connect(
    contracts.syloTicketing,
    provider,
  );

  const seekers = factories.TestSeekers__factory.connect(
    contracts.seekers,
    provider,
  );

  const futurepassRegistrar =
    factories.TestFuturepassRegistrar__factory.connect(
      contracts.futurepassRegistrar,
      provider,
    );

  return {
    syloToken,
    authorizedAccounts,
    registries,
    ticketingParameters,
    epochsManager,
    stakingManager,
    rewardsManager,
    directory,
    syloTicketing,
    seekers,
    futurepassRegistrar,
  };
}
