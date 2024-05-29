# Ganache Docker Image

This docker image runs a local ganache node with the contracts
already deployed, 10 seekers minted to the deployer, and other
network params set to be ready for ticket redemption.

When the image is built, an `accounts.json`
used for the network and an `addresses.json` is stored in
`/app/deployment`.

The `addresses.json` is a JSON file listing out the deployed
contracts.

```json
{
  "deployer": "0x835dF5fE77D479695a616F79A3FC3a25310eb7c6",
  "token": "0xc4db8fD3209c98290AB32693F5155c596B97Eabe",
  "registries": "0x6137Da22d887053274ddd3923C6C364F7E34d9F2",
  "ticketing": "0xB9c4F75F96cf9ba473607EC6e0e6249e39d8Ca03",
  "ticketingParameters": "0x044Be02533C31207A3Ca34CaF6f574accB375bd1",
  "directory": "0x4BF9F7cf118870485F135DaFd8B5EbBAC334ACAc",
  "rewardsManager": "0x308de5B9b84961874E3414D9B677bC36Ef48fB18",
  "epochsManager": "0xBE2cD9BECf461624a38AF9f4Df4585a3C5DbD478",
  "stakingManager": "0x6d48a17aAbB94157F75ABB83538D6175B5354361",
  "seekers": "0x49C537a88016186Ef41713239799Fc975F9e9aFA"
}
```

Use volumes to feed the accounts or addresses to a local test or
another service.

The chain state is saved in the directory `ganache-data`.

The `mnemonic` used to when starting the ganache node and also
used to deploy the contracts is
`enroll regret dial tray life phrase saddle term friend figure meat add puppy explain soup`

Note: The context for building this image is in the root path of
this repository.

## Build docker

```sh
yarn docker-ganache
```

## Run docker

The options used to run the local ganache node should be set correctly
to ensure that it re-uses the same directory for the chain state.

```sh
docker run -p 8545:8545 dn3010/sylo-ethereum-testnet:0.1.0 `bash -c ganache-cli --host 0.0.0.0 --db ganache-data --blockTime 5`
```
