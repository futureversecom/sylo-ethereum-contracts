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

echo "shutting down rootnet"
kill $rootnet_pid
sleep 5
echo "deployment complete"
