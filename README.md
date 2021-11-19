# Sylo Ethereum Contracts

Smart Contracts used for the Sylo Network Protocol. These contracts
define the mechanisms for:
  - The ERC21 Sylo Token
  - Staking of Sylos against Sylo Node operators
  - Stake-weighted scan function used to allocate business
  within the network
  - Probabilistic Micro-payments for providing the Event Relay
  Service
  - Epochs and various network parameters

## Documentation

An [overview](docs/overview.md) is available for the Sylo Network Protocol.
Additionally, read the [contract specification](docs/spec.md) to
help understand the implementation of the contracts.

## Development

This project employs [Hardhat](https://hardhat.org/getting-started/) for development and testing.

### Setup

Ensure Node.js (>=v12.0) is installed.

`npm install`

### Build

`npm run build`

This will compile the contracts and create typechain typescript definitions.

### Running Tests

Testing is done through a local hardhat network.

`npm test`

Running this will also compile a `gasReport.txt`, which show gas costs
for each contract call.

#### Coverage

`npm run coverage`

This project attempts to maintain 100% code coverage at all times.

### Deployment

Deployment is supported by the `hardhat.config.ts` configuration. These
contracts rely on [open-zeppelin's upgrade proxy plugin](https://docs.openzeppelin.com/upgrades-plugins/1.x/) in order to be able to update
contracts post-deployment.

`npx hardhat run deploy/00_deploy_phase_two.ts` can be used to deploy
to a local hardhat network.

Deployment configurations will be saved `.openzeppelin/${network}.json`
