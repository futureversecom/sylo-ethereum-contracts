import { BigNumber, BigNumberish } from "ethers";

type ContractParameters = {
  // Address of the existing Sylo Token
  SyloToken: string;

  EpochsManager: {
    epochDuration: BigNumberish;
  }

  Listings: {
    defaultPayoutPercentage: number;
  }

  TicketingParameters: {
    faceValue: BigNumberish;
    baseLiveWinProb: BigNumberish;
    expiredWinProb: BigNumberish;
    ticketDuration: BigNumberish;
    decayRate: number;
  }

  Ticketing: {
    unlockDuration: BigNumberish;
  }

  StakingManager: {
    unlockDuration: BigNumberish;
  }
}

const GenesisParameters: ContractParameters = {
  SyloToken: "0xf293d23bf2cdc05411ca0eddd588eb1977e8dcd4",

  EpochsManager: {
    epochDuration: 80000
  },

  Listings: {
    defaultPayoutPercentage: 5000,
  },

  TicketingParameters: {
    faceValue: 100000,
    baseLiveWinProb: BigNumber.from(2).pow(128).sub(1).div(1000),
    expiredWinProb: BigNumber.from(2).pow(128).sub(1).div(1000),
    ticketDuration: 80000,
    decayRate: 8000
  },

  Ticketing: {
    unlockDuration: 80000
  },

  StakingManager: {
    unlockDuration: 8000
  }
}

export default GenesisParameters;