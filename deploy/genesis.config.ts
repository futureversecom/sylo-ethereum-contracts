import * as hre from 'hardhat';
import { BigNumber, BigNumberish } from 'ethers';

type ContractParameters = {
  // Address of the existing Sylo Token
  SyloToken: string;

  Seekers: {
    seekerAddress: string;
    shouldMint: boolean;
  };

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

  Seekers: {
    seekerAddress:
      hre.network.name === 'nikau'
        ? '0xC65fDC6c38D0a1d3524aE54ba205BDE197AbddbA'
        : hre.network.name === 'porcini'
        ? '0xAAAAAAAA00001864000000000000000000000000'
        : '',
    shouldMint:
      hre.network.name === 'localhost' || hre.network.name === 'rootlocal'
        ? true
        : false,
  },

  EpochsManager: {
    epochDuration: 12, // 1 minute
  },

  Registries: {
    defaultPayoutPercentage: 5000,
    proofDuration: 100,
  },

  TicketingParameters: {
    faceValue: hre.ethers.utils.parseEther('10000'),
    baseLiveWinProb: BigNumber.from(2).pow(128).sub(1), // 100%
    expiredWinProb: BigNumber.from(2).pow(128).sub(1).div(1000),
    ticketDuration: 80000,
    decayRate: 8000,
  },

  Ticketing: {
    unlockDuration: 30, // 150 seconds
  },

  StakingManager: {
    unlockDuration: 30,
    minimumStakeProportion: 2000,
  },
};

export default GenesisParameters;
