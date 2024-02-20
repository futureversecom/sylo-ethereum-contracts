#!/bin/bash

set -e

echo "starting local ganache node"
ganache \
  --database.dbPath ganache-data \
  --wallet.mnemonic "enroll regret dial tray life phrase saddle term friend figure meat add puppy explain soup" \
  --miner.blockTime 1 \
  --wallet.accountKeysPath ./deployment/accounts.json \
  &>/dev/null &

ganache_pid=`echo $!`

echo "waiting for ganache to start up"
sleep 15

echo "deploying contracts to local node"
npx hardhat --network localhost deploy
mv ./deployments/localhost_deployment_phase_two.json deployment/addresses.json

echo "initialzing network"
npx hardhat --network localhost run scripts/init_local_network.ts

echo "shutting down ganache"
kill $ganache_pid
sleep 10
echo "deployment complete"
