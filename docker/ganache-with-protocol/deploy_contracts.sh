#!/bin/bash

echo "starting local ganache node"
ganache-cli \
  --db ganache-data \
  --deterministic \
  --mnemonic "enroll regret dial tray life phrase saddle term friend figure meat add puppy explain soup" \
  --blockTime 5 \
  --account_keys_path ./deployment/accounts.json \
  &>/dev/null &

ganache_pid=`echo $!`

echo "waiting for ganache to start up"
sleep 5

echo "deploying contracts to local node"
npx hardhat --network localhost run ./deploy/00_deploy_phase_two.ts
mv ./deploy/localhost_deployment_phase_two.json deployment/addresses.json

echo "minting seekers to deployer"
npx ts-node scripts/mint_seeker.ts \
    --evm http://0.0.0.0:8545 \
    --account_pk 0x150934096e7bcd0485d154edd771b4466680038a068ccca8e8b483dce8527245 \
    --seeker_contract 0x49C537a88016186Ef41713239799Fc975F9e9aFA \
    --amount 10

echo "shutting down ganache"
kill $ganache_pid
sleep 10
echo "deployment complete"
