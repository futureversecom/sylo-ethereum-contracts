#!/bin/bash

# create a new hardhat config with private key hardcoded in
cat << EOF > hardhat.config.ts
import { HardhatUserConfig } from "hardhat/config";
import "@nomiclabs/hardhat-etherscan";
import "@nomiclabs/hardhat-waffle";
import "@typechain/hardhat";
import "@nomiclabs/hardhat-ethers";
import "hardhat-gas-reporter";
import "solidity-coverage";
import "@openzeppelin/hardhat-upgrades";

const config: HardhatUserConfig = {
  solidity: {
    version: "0.8.4",
    settings: {
      optimizer: {
        enabled: true,
        runs: 1500,
      },
    },
  },
  networks: {
    ganache: {
      url: "http://0.0.0.0:8545",
      accounts: "remote"
    }
  }
};

export default config;
EOF

echo "starting local ganache node"
ganache-cli \
  --db ganache-data \
  --deterministic \
  --mnemonic "enroll regret dial tray life phrase saddle term friend figure meat add puppy explain soup" \
  --blockTime 10 \
  --account_keys_path ./deployment/accounts.json \
  --chainId 1 \
  &>/dev/null &

ganache_pid=`echo $!`

trap 'kill -9 $ganache_pid' EXIT

sleep 5

echo "deploying contracts to local node"
npx hardhat --network ganache run ./deploy/00_deploy_phase_two.ts

mv ./deploy/ganache_deployment_phase_two.json deployment/addresses.json
