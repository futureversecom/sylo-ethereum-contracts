#!/bin/bash

set -e

echo "starting local blockchain node"
npx hardhat node --no-deploy &>/dev/null &

blockchain_pid=`echo $!`

echo "deploying contracts"
npx hardhat --network localhost deploy

echo "initializing network"
npx hardhat --network localhost run ./scripts/init_local_network.ts

trap "trap - SIGTERM && kill -- -$$" SIGINT SIGTERM EXIT
wait
