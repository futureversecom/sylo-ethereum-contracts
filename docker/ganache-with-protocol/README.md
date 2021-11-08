## Ganache Docker Image

This docker image runs a local ganache node with the contracts
already deployed.

When the image is built, an `accounts,json`
used for the network and an `addresses.json` is stored in
`/app/deployment`.

The `addresses.json` is a JSON file listing out the deployed
contracts.

```
{
 "deployer": "0x835dF5fE77D479695a616F79A3FC3a25310eb7c6",
 "token": "0xc4db8fD3209c98290AB32693F5155c596B97Eabe",
 "listings": "0x2f2f947095021Ff12B316Fe61FcE57Fb57C90366",
 "ticketing": "0x943E7031A7Ed0FC236173f05a3084104b81Aa480",
 "ticketingParameters":"0x075EEeD1215982b78A2e05cD2213b5f53A718a9",
 "directory": "0xa4dE4FEA5e961e5C130013CfE207f7C08148A73C",
 "rewardsManager": "0x7bFCE7796fdE3Ba0F2052959d506bdA480518edA",
 "epochsManager": "0xFB87c433852Bb2917B37b0471DFA5B369e75083A",
 "stakingManager": "0x65A9be6e97eD2F417250A819e93BCb359b1140d0"
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

### Ganache Docker Image

The options used to run the local ganache node should be set correctly
to ensure that it re-uses the same directory for the chain state.

```
docker run \
  -p 8545:8545 \
  $image_sha \
  `bash -c ganache-cli --host 0.0.0.0 --db ganache-data --blockTime 10`
```