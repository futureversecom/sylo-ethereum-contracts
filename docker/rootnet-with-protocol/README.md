# Rootnet Docker Image

This docker image runs a local root network node (seed) with the contracts
already deployed, and 10 seekers minted for the deployer.

The contract addresses are

```json
{
  "deployer": "0x25451A4de12dcCc2D166922fA938E900fCc4ED24",
  "token": "0xfAd47F21Ba0FC182691d1e14e916eA18208A45A9",
  "registries": "0xEa1a67093759de6f434182161815AfB33D9D30d6",
  "ticketing": "0xE9C62C36Bd5E4E4B322bfFB54B2B709eF6964E60",
  "ticketingParameters": "0x483a30dF2ddEA31d9E1ce7d35833bf2d759f4346",
  "directory": "0xfeEAD513aBAD98cE78645D5Fd625aE0f3ede6Dc5",
  "rewardsManager": "0x4002F76bA7db46b3Dc79507C11A08Fa389b9d6CB",
  "epochsManager": "0x161b5FA7CCCf8477Dc892941e303A159EC7D3e57",
  "stakingManager": "0x41Cd8EB92BD2f7d5eAa1202666AB9EC9F7058151",
  "seekers": "0x50f15F0a0B798b6E3A248Fccc26b6de636b7c3ef"
}
```

The chain state is saved in the directory `/mnt/data`.

The `private key` used to when starting the node and also
used to deploy the contracts is
`0x79c3b7fc0b7697b9414cb87adcb37317d1cab32818ae18c0e97ad76395d1fdcf`

Note: The context for building this image is in the root path of
this repository.

## Build docker

```sh
npm run docker-rootnet
```

## Run docker

The options used to run the local rootnet node should be set correctly
to ensure that it re-uses the same directory for the chain state.

```sh
docker run -p 9933:9933 -p 9944:9944 dn3010/sylo-ethereum-testnet:rootnet --dev --unsafe-ws-external --unsafe-rpc-external --base-path=/mnt/data
```
