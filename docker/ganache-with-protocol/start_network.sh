#!/bin/bash

run_evm_mine() {
  while true
  do
    curl -s -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"evm_mine","params":[],"id":1}' http://localhost:8545
    sleep 10
  done
}

# mine block in background
run_evm_mine &

## run ganache node
node /app/dist/node/cli.js "$@"