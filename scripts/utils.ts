import { ethers } from 'ethers';
import * as factories from '../typechain-types';
import contractAddress from '../deploy/ganache_deployment_phase_two.json';

export type Contracts = {
  stakingManager: factories.contracts.staking.StakingManager;
  token: factories.contracts.SyloToken;
  seekers: factories.contracts.mocks.TestSeekers;
  registries: factories.contracts.Registries;
  ticketing: factories.contracts.payments.SyloTicketing;
  ticketingParameters: factories.contracts.payments.ticketing.TicketingParameters;
  epochsManager: factories.contracts.epochs.EpochsManager;
};

export function conectContracts(provider: ethers.providers.JsonRpcProvider) {
  const stakingManager = factories.StakingManager__factory.connect(
    contractAddress.stakingManager,
    provider,
  );

  const token = factories.SyloToken__factory.connect(
    contractAddress.token,
    provider,
  );

  const seekers = factories.TestSeekers__factory.connect(
    contractAddress.seekers,
    provider,
  );

  const registries = factories.Registries__factory.connect(
    contractAddress.registries,
    provider,
  );

  const ticketing = factories.SyloTicketing__factory.connect(
    contractAddress.ticketing,
    provider,
  );

  const ticketingParameters = factories.TicketingParameters__factory.connect(
    contractAddress.ticketingParameters,
    provider,
  );

  const epochsManager = factories.EpochsManager__factory.connect(
    contractAddress.epochsManager,
    provider,
  );

  return {
    token,
    stakingManager,
    seekers,
    registries,
    contractAddress,
    ticketing,
    ticketingParameters,
    epochsManager,
  } as Contracts;
}
