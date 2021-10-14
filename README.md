# sylo-ethereum-contracts
Smart Contracts used for Sylo Incentivisation

## Developing

This project employs [Hardhat](https://hardhat.org/getting-started/) for development and testing.

### Setup

`npm i`

### Build

`npm run build`

### Test

`npm test`

### Deploying to Ropsten

Create a `.env` file in the root folder of this project with the following values

```
# Can retrieve this from 1password
ROPSTEN_MNEMONIC=XXX
# Login to Infura account and select Sylo Test Network project and find the endpoint for Ropsten network
ROPSTEN_INFURA_ENDPOINT=XXX
```

Then run
`npx hardhat --network ropsten run deploy/00_deploy_phase_two.ts`

The deployed configuration should be saved to `.openzeppelin/ropsten.json`