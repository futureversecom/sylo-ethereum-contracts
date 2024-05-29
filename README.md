# Sylo Ethereum Contracts

Smart Contracts used for the Sylo Network Protocol. These contracts define the
mechanisms for:

- The ERC21 Sylo Token
- Staking of Sylos against Sylo Node operators
- Stake-weighted scan function used to allocate business within the network
- Probabilistic Micro-payments for providing the Event Relay Service
- Epochs and various network parameters

## 📖 Documentation

An [overview](docs/overview.md) is available for the Sylo Network Protocol.
Additionally, read the [contract specification](docs/spec.md) to help understand
the implementation of the contracts.

## 🎸 Development

This project employs [Hardhat](https://hardhat.org/getting-started/) for
development and testing.

### 🔥 Setup

Ensure Node.js (>=v18.0) is installed.

`yarn`

### 🦖 Build

`yarn build`

This will compile the contracts and create typechain typescript definitions.

### 🧪 Running Tests

Testing is done through a local hardhat network.

`yarn test`

Running this will also compile a `gasReport.txt`, which show gas costs for each
contract call.

#### 🎁 Coverage

`yarn coverage`

This project attempts to maintain 100% code coverage at all times.

### 🐳 Docker

A docker image that contains a local ethereum node with the current contracts
deployed to that network can be built with:

`npm run docker-ganache`

See the [docker readme](docker/README.md) for more
details.

### 🌥️ Deployment

#### 🚗 Tool

The `hardhat-deploy` plugin is used to manage deployments. See the
[documentation](https://github.com/wighawag/hardhat-deploy/#-hardhat-deploy) for
more details.

#### 🦁 Branch

We should use the branch `deploy-contracts` when deploying new set of contract.
It represents for the latest deployed contracts addresses for all networks.

There is no need to keep the history for the `deploy-contracts` branch, so we
can just force push new commits from the `master` branch to this branch when
needed.

#### 🌈 How to

Before deploying, we should remove the `deployments/<network>` folder (e.g.
`deployments/porcini-dev`) so all contracts will be created again, instead of
reusing Factory contracts.

Command to deploy:

```sh
npx hardhat deploy --network <network>
```

If you wanna update the deployment code, it should be in the `deploy` folder,
the file number represents for the order of execution.
