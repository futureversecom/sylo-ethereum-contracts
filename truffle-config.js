module.exports = {
  networks: {
   test: {
    host: "127.0.0.1",
    port: 8545,
    network_id: "*",
    websockets: true,
    disableConfirmationListener: true
  },
  },
  compilers: {
    solc: {
      version: '0.8.4',
      settings: {
        optimizer: {
          enabled: true,
          runs: 1500
        }
      }
    }
  },
  plugins: ["solidity-coverage"],
  mocha: {
    reporter: 'eth-gas-reporter',
    reporterOptions : {
      coinmarketcap: '3da4e7e8-31fb-477a-85a8-a905ad24fd28',
      currency: 'USD',
      outputFile: 'gasReport.txt',
      noColors: 'true' // Needed for outputfile
    }
  }
};
