import * as hre from 'hardhat';
import { BigNumberish } from 'ethers';

type ContractParameters = {
  // Address of the existing Sylo Token
  SyloToken: string;

  // Address of the existing bridged Seekers contract
  Seekers: string;

  // Address of the pre-compile futurepass registrar
  FuturepassRegistrar: string;

  EpochsManager: {
    initialEpoch: BigNumberish;
    epochDuration: BigNumberish;
  };

  Registries: {
    defaultPayoutPercentage: number;
  };

  TicketingParameters: {
    faceValue: BigNumberish;
    baseLiveWinProb: BigNumberish;
    expiredWinProb: BigNumberish;
    ticketDuration: BigNumberish;
    decayRate: number;
  };

  Ticketing: {
    unlockDuration: BigNumberish;
  };

  StakingManager: {
    unlockDuration: BigNumberish;
    minimumStakeProportion: number;
    seekerPowerMultiplier: BigNumberish;
  };

  SeekerPowerOracle: {
    oracleAccount: string;
  };
};

const GenesisParameters: ContractParameters = {
  SyloToken:
    hre.network.name === 'mainnet'
      ? '0xf293d23bf2cdc05411ca0eddd588eb1977e8dcd4'
      : hre.network.name === 'rata'
      ? '0xcCCCcCcC00004274000000000000000000000000'
      : hre.network.name === 'nikau'
      ? '0xcccCccCC000042B4000000000000000000000000'
      : '',

  Seekers:
    hre.network.name === 'nikau'
      ? '0xC65fDC6c38D0a1d3524aE54ba205BDE197AbddbA'
      : '',

  FuturepassRegistrar: '',

  EpochsManager: {
    initialEpoch: 0,
    epochDuration: 80000,
  },

  Registries: {
    defaultPayoutPercentage: 50000,
  },

  TicketingParameters: {
    faceValue: 100000,
    baseLiveWinProb: (2n ** 128n - 1n) / 1000n,
    expiredWinProb: (2n ** 128n - 1n) / 1000n,
    ticketDuration: 80000,
    decayRate: 80000,
  },

  Ticketing: {
    unlockDuration: 80000,
  },

  StakingManager: {
    unlockDuration: 8000,
    minimumStakeProportion: 20000,
    seekerPowerMultiplier: hre.ethers.parseEther('1000000'),
  },

  SeekerPowerOracle: {
    oracleAccount: '0x835dF5fE77D479695a616F79A3FC3a25310eb7c6',
  },
};

const GanacheTestnetParameters: ContractParameters = {
  SyloToken: '',

  Seekers: '',

  FuturepassRegistrar: '',

  EpochsManager: {
    initialEpoch: 0,
    epochDuration: 30, // 30 * 4 = 120 seconds = 2 minutes
  },

  Registries: {
    defaultPayoutPercentage: 50000,
  },

  TicketingParameters: {
    faceValue: 100000,
    baseLiveWinProb: (2n ** 128n - 1n) / 1000n,
    expiredWinProb: (2n ** 128n - 1n) / 1000n,
    ticketDuration: 10_000_000, // make sure the ticket never expires in the short time on testnet
    decayRate: 80000,
  },

  Ticketing: {
    unlockDuration: 30, // 30 * 4 = 120 seconds = 2 minutes
  },

  StakingManager: {
    unlockDuration: 30, // 30 * 4 = 120 seconds = 2 minutes
    minimumStakeProportion: 20000,
    seekerPowerMultiplier: hre.ethers.parseEther('1000000'),
  },

  SeekerPowerOracle: {
    oracleAccount: '',
  },
};

const PorciniDevParameters: ContractParameters = {
  SyloToken: '0xCCcCCcCC00000C64000000000000000000000000',

  Seekers: '0xAAAAAAAA00001864000000000000000000000000',

  FuturepassRegistrar: '0x000000000000000000000000000000000000FFFF',

  EpochsManager: {
    initialEpoch: 0,
    epochDuration: 17280, // 1 day
  },

  Registries: {
    defaultPayoutPercentage: 50000,
  },

  TicketingParameters: {
    faceValue: hre.ethers.parseEther('1000'),
    baseLiveWinProb: 2n ** 128n - 1n / 10n,
    expiredWinProb: 2n ** 128n - 1n,
    ticketDuration: 17280,
    decayRate: 80000,
  },

  Ticketing: {
    unlockDuration: 10,
  },

  StakingManager: {
    unlockDuration: 10,
    minimumStakeProportion: 20000,
    seekerPowerMultiplier: hre.ethers.parseEther('1000000'),
  },

  SeekerPowerOracle: {
    oracleAccount: '0x448c8e9e1816300Dd052e77D2A44c990A2807D15',
  },
};

export { GenesisParameters, GanacheTestnetParameters, PorciniDevParameters };

export type { ContractParameters };
