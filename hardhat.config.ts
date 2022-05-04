import * as dotenv from "dotenv";

import { HardhatUserConfig } from "hardhat/config";
import "@nomiclabs/hardhat-etherscan";
import "@nomiclabs/hardhat-waffle";
import "@typechain/hardhat";
import "@nomiclabs/hardhat-ethers";
import "hardhat-gas-reporter";
import "solidity-coverage";
import "@openzeppelin/hardhat-upgrades";

dotenv.config();

const config: HardhatUserConfig = {
  solidity: {
    version: "0.8.13",
    settings: {
      optimizer: {
        enabled: true,
        runs: 1500,
      },
    },
  },
  gasReporter: {
    coinmarketcap: "3da4e7e8-31fb-477a-85a8-a905ad24fd28",
    currency: "USD",
    outputFile: "gasReport.txt",
    noColors: true, // Needed for outputfile
  },
  networks: {
    localhost: {
      url: "http://0.0.0.0:8545",
      accounts: "remote",
    },
    ropsten: {
      url: process.env.ROPSTEN_INFURA_ENDPOINT ?? "",
      accounts: {
        mnemonic: process.env.ROPSTEN_MNEMONIC ?? "",
      },
    },
    kovan: {
      url: process.env.KOVAN_INFURA_ENDPOINT ?? "",
      accounts: {
        mnemonic: process.env.ROPSTEN_MNEMONIC ?? "",
      },
    },
    nikau: {
      url: "https://nikau.centrality.me/public",
      accounts: {
        mnemonic: process.env.ROPSTEN_MNEMONIC ?? "",
      },
    },
  },
};

export default config;
