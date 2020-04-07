# Go Eth Client for Sylo Smart Contracts

Provides a simple client interface for interacting with Sylo Smart contracts in Go

## Generating Code

All the files in the `contracts` directory are generated using go-ethereum [abigen](https://geth.ethereum.org/docs/dapp/native-bindings)

### Build contracts

From the project root
```
npm i
npm run build
node ./scripts/parseArtifacts.js

```

#### Generate code from built contracts
```
cd go-eth
go generate client.go

```