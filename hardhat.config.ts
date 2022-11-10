import * as dotenv from 'dotenv';

import { HardhatUserConfig } from 'hardhat/config';

import '@nomicfoundation/hardhat-chai-matchers';
import '@nomicfoundation/hardhat-ethers';
import 'hardhat-gas-reporter';
import '@typechain/hardhat';
import 'solidity-coverage';
import 'hardhat-deploy';

dotenv.config();

const config: HardhatUserConfig = {
  solidity: {
    version: '0.8.18',
    settings: {
      optimizer: {
        enabled: true,
        runs: 200,
      },
    },
  },
  gasReporter: {
    coinmarketcap: '3da4e7e8-31fb-477a-85a8-a905ad24fd28',
    currency: 'USD',
    outputFile: 'gasReport.txt',
    noColors: true, // Needed for outputfile
  },
  networks: {
    hardhat: {
      gas: 100000000,
      blockGasLimit: 2000000000000, // Add this to allow multiple transactions in one block when testing
    },
    localhost: {
      url: 'http://0.0.0.0:8545',
      accounts: 'remote',
    },
    ropsten: {
      url: process.env.ROPSTEN_INFURA_ENDPOINT ?? '',
      accounts: {
        mnemonic: process.env.ROPSTEN_MNEMONIC ?? '',
      },
    },
    rata: {
      url: process.env.RATA_ENDPOINT ?? '',
      accounts: {
        mnemonic: process.env.ROPSTEN_MNEMONIC ?? '',
      },
    },
    nikau: {
      url: process.env.NIKAU_ENDPOINT ?? '',
      accounts: {
        mnemonic: process.env.ROPSTEN_MNEMONIC ?? '',
      },
    },
    'porcini-dev': {
      url: 'https://porcini.rootnet.app',
      accounts: {
        mnemonic: process.env.ROPSTEN_MNEMONIC ?? '',
      },
    },
    'porcini-testing': {
      url: 'https://porcini.rootnet.app',
      accounts: {
        mnemonic: process.env.ROPSTEN_MNEMONIC ?? '',
      },
    },
    rootlocal: {
      url: 'http://0.0.0.0:9933',
      accounts: [
        '0x79c3b7fc0b7697b9414cb87adcb37317d1cab32818ae18c0e97ad76395d1fdcf',
      ],
    },
  },
};

export default config;
