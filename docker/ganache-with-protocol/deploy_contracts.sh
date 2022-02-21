#!/bin/bash

echo "starting local ganache node"
ganache-cli \
  --db ganache-data \
  --deterministic \
  --mnemonic "enroll regret dial tray life phrase saddle term friend figure meat add puppy explain soup" \
  --blockTime 10 \
  --account_keys_path ./deployment/accounts.json \
  &>/dev/null &

ganache_pid=`echo $!`

echo "waiting for ganache to start up"
sleep 5

echo "deploying contracts to local node"
npx hardhat --network localhost run ./deploy/00_deploy_phase_two.ts
sleep 10
mv ./deploy/localhost_deployment_phase_two.json deployment/addresses.json
sleep 10

echo "shutting down ganache"
kill $ganache_pid
sleep 10
echo "deployment complete"
