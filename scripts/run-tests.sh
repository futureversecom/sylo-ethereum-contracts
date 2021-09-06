#!/bin/bash

# start new eth network
# will use 'test' network defined in truffle-config.js
npm run ganache > /dev/null 2>&1 &

# run tests
truffle test

exit $?
