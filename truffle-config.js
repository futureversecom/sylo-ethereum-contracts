module.exports = {
  // Uncommenting the defaults below 
  // provides for an easier quick-start with Ganache.
  // You can also follow this format for other networks;
  // see <http://truffleframework.com/docs/advanced/configuration>
  // for more details on how to specify configuration options!
  //
  networks: {
   testNetwork: {
     host: "127.0.0.1",
     port: 8555,
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
      currency: 'USD'
    }
  }
};
