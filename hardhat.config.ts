import * as dotenv from 'dotenv';

import { HardhatUserConfig } from 'hardhat/config';
import '@nomiclabs/hardhat-etherscan';
import '@typechain/hardhat';
import '@nomiclabs/hardhat-ethers';
import 'hardhat-gas-reporter';
import 'solidity-coverage';
import '@openzeppelin/hardhat-upgrades';
import '@nomicfoundation/hardhat-chai-matchers';

dotenv.config();

const config: HardhatUserConfig = {
  solidity: {
    version: '0.8.13',
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
    porcini: {
      url: 'https://porcini.au.rootnet.app',
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
