import * as hre from 'hardhat';
import { BigNumber, BigNumberish } from 'ethers';

type ContractParameters = {
  // Address of the existing Sylo Token
  SyloToken: string;

  Seekers: {
    seekersERC721: string;
    oracle: string;
    validDuration: BigNumberish;
    callbackGasLimit: BigNumberish;
    callbackBounty: BigNumberish;
  };

  EpochsManager: {
    epochDuration: BigNumberish;
  };

  Listings: {
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
      : '',

  Seekers: {
    seekersERC721:
      hre.network.name === 'rata'
        ? '0x856D2df6998AcA9FDC4B2CA316b0527081cee8DD'
        : hre.network.name === 'nikau'
        ? '0x49e5706c06a263ac142d0a8c06f7ab247c1d978a'
        : '',
    oracle: '0x0000000000000000000000000000000000006bb4',
    validDuration: 200,
    callbackGasLimit: 300000,
    callbackBounty: 0,
  },

  EpochsManager: {
    epochDuration: 80000,
  },

  Listings: {
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
