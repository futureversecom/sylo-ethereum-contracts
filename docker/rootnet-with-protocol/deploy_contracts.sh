#!/bin/bash

echo "starting local rootnet node"
/usr/bin/seed \
    --dev \
    --ws-external \
    --rpc-external \
    --base-path=/mnt/data \
    &>/dev/null &

rootnet_pid=`echo $!`

echo "waiting for rootnet to start up"
sleep 5

echo "deploying contracts to local node"
npx hardhat run ./deploy/00_deploy_phase_two.ts --network rootlocal

echo "minting seekers to deployer"
npx ts-node scripts/mint_seeker.ts \
    --evm http://0.0.0.0:9933 \
    --account_pk 0x79c3b7fc0b7697b9414cb87adcb37317d1cab32818ae18c0e97ad76395d1fdcf \
    --seeker_contract 0x50f15F0a0B798b6E3A248Fccc26b6de636b7c3ef \
    --amount 10

echo "shutting down rootnet"
kill $rootnet_pid
sleep 5
echo "deployment complete"
