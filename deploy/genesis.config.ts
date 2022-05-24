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
    hre.network.name == 'mainnet'
      ? '0xf293d23bf2cdc05411ca0eddd588eb1977e8dcd4'
      : '',

  Seekers: {
    seekersERC721:
      hre.network.name == "rata"
        ? "0x856D2df6998AcA9FDC4B2CA316b0527081cee8DD"
        : "",
    oracle: "0x0000000000000000000000000000000000006bb4",
    validDuration: 200,
    callbackGasLimit: 300000,
    callbackBounty: hre.ethers.utils.parseEther("2"),
  },

  EpochsManager: {
    epochDuration: 80000,
  },

  Listings: {
    defaultPayoutPercentage: 5000,
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
