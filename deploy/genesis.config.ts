import * as hre from 'hardhat';
import { BigNumber, BigNumberish } from 'ethers';

type ContractParameters = {
  // Address of the existing Sylo Token
  SyloToken: string;

  // Address of the existing bridged Seekers contract
  Seekers: string;

  EpochsManager: {
    epochDuration: BigNumberish;
  };

  Registries: {
    defaultPayoutPercentage: number;
    proofDuration: number;
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
      : hre.network.name === 'porcini'
      ? '0xCCcCCcCC00000C64000000000000000000000000'
      : '',

  Seekers:
    hre.network.name === 'nikau'
      ? '0xC65fDC6c38D0a1d3524aE54ba205BDE197AbddbA'
      : hre.network.name === 'porcini'
      ? '0xAAAAAAAA00001864000000000000000000000000'
      : '',

  EpochsManager: {
    epochDuration: 80000,
  },

  Registries: {
    defaultPayoutPercentage: 5000,
    proofDuration: 100,
  },

  TicketingParameters: {
    faceValue: 100000,
    baseLiveWinProb: BigNumber.from(2).pow(128).sub(1).div(1000),
    expiredWinProb: BigNumber.from(2).pow(128).sub(1).div(1000),
    ticketDuration: 80000,
    decayRate: 8000,
  },

  Ticketing: {
    unlockDuration: 80000,
  },

  StakingManager: {
    unlockDuration: 8000,
    minimumStakeProportion: 3000,
  },
};

export default GenesisParameters;
