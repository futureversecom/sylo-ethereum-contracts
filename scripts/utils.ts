import { ethers } from 'ethers';
import * as factories from '../typechain-types';

export type ContractsJSON = {
  stakingManager: string;
  token: string;
  seekers: string;
  registries: string;
  ticketing: string;
  ticketingParameters: string;
  epochsManager: string;
};

export type Contracts = {
  stakingManager: factories.contracts.staking.StakingManager;
  token: factories.contracts.SyloToken;
  seekers: factories.contracts.mocks.TestSeekers;
  registries: factories.contracts.Registries;
  ticketing: factories.contracts.payments.SyloTicketing;
  ticketingParameters: factories.contracts.payments.ticketing.TicketingParameters;
  epochsManager: factories.contracts.epochs.EpochsManager;
};

export function conectContracts(
  contracts: ContractsJSON,
  provider: ethers.providers.JsonRpcProvider,
) {
  const stakingManager = factories.StakingManager__factory.connect(
    contracts.stakingManager,
    provider,
  );

  const token = factories.SyloToken__factory.connect(contracts.token, provider);

  const seekers = factories.TestSeekers__factory.connect(
    contracts.seekers,
    provider,
  );

  const registries = factories.Registries__factory.connect(
    contracts.registries,
    provider,
  );

  const ticketing = factories.SyloTicketing__factory.connect(
    contracts.ticketing,
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

  return {
    token,
    stakingManager,
    seekers,
    registries,
    ticketing,
    ticketingParameters,
    epochsManager,
  } as Contracts;
}
