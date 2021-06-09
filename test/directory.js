const BN = require("bn.js");
const crypto = require("crypto");
const sodium = require('libsodium-wrappers-sumo');

const Directory = artifacts.require("Directory");
const Token = artifacts.require("SyloToken");

contract('Directory', accounts => {
  let token;
  let directory;

  before(async () => {
    token = await Token.new({ from: accounts[1] });
  });

  beforeEach(async () => {
    directory = await Directory.new(token.address, 0, { from: accounts[0] });

    await token.approve(directory.address, 10000, { from: accounts[1] });
  });

  it('should be able to unstake the root', async () => {
    await directory.addStake(1, accounts[1], { from: accounts[1] });
    await directory.addStake(1, accounts[2],{ from: accounts[1] });

    const totalStake = await directory.getTotalStake();
    assert.equal(totalStake, 2, "Expected a total stake of 2");

    await directory.unlockStake(1, accounts[1], { from: accounts[1] });

    const totalStake2 = await directory.getTotalStake();
    assert.equal(totalStake2.toNumber(), 1, "Expected a total stake of 1");
  });

  it('should be able to unstake a leaf', async () => {
    await directory.addStake(1, accounts[1], { from: accounts[1] });
    await directory.addStake(1, accounts[2],{ from: accounts[1] });

    const totalStake = await directory.getTotalStake();
    assert.equal(totalStake, 2, "Expected a total stake of 2");

    await directory.unlockStake(1, accounts[2], { from: accounts[1] });

    const totalStake2 = await directory.getTotalStake();
    assert.equal(totalStake2.toNumber(), 1, "Expected a total stake of 1");
  });

  it('should be able to unstake a node with parent and child', async () => {
    for (let i = 0; i < accounts.length; i++) {
      await directory.addStake(1, accounts[i],{ from: accounts[1] });
    }

    const totalStake = await directory.getTotalStake();
    assert.equal(totalStake, accounts.length, "Expected a total stake of 2");

    await directory.unlockStake(1, accounts[2], { from: accounts[1] });

    const totalStake2 = await directory.getTotalStake();
    assert.equal(totalStake2.toNumber(), accounts.length - 1, "Expected a total stake of 1");
  });

  it('should be able to restake when everything is unstaked', async () => {
    await directory.addStake(1, accounts[1], { from: accounts[1] });

    await directory.unlockStake(1, accounts[1], { from: accounts[1] });

    // Restake
    await directory.addStake(1, accounts[1], { from: accounts[1] });
  });

  it('should distribute scan results amongst stakees proportionally', async () => {
    for (let i = 0; i < accounts.length; i++) {
      await directory.addStake(1, accounts[i], { from: accounts[1] });
      console.log('added stake to', accounts[i]);
    }

    const results = await collectScanResults(1000);
    console.log(results);
  }).timeout(0);

  // it('should be able to scan after unlocking all stake', async () => {
  //   await directory.addStake(1, accounts[0], { from: accounts[1] });
  //   await directory.addStake(1, accounts[1], { from: accounts[1] });
  //   await directory.addStake(1, accounts[2], { from: accounts[1] });
  // });

  it('should distribute scan results proporitionally after nodes unlock stake', async () => {
    for (let i = 0; i < accounts.length; i++) {
      await directory.addStake(1, accounts[i], { from: accounts[1] });
      console.log('added stake to', accounts[i]);
    }

    /***
    *             0
    *           /   \
    *          2     1
    *         / |   | \
    *        6  4   5  3
    *            \   \   \
    *             7   9   8
    */

    console.log("unstaking for address", accounts[4]);
    await debug(directory.unlockStake(1, accounts[4], { from: accounts[1] }));

    /***
    *             0
    *           /   \
    *          2     1
    *         / |   | \
    *        6  7   5  3
    *                \  \
    *                 9   8
    */

    console.log("unstaking for address", accounts[5]);
    await directory.unlockStake(1, accounts[5], { from: accounts[1] });

    /***
    *             0
    *           /   \
    *          2     1
    *         / |   | \
    *        6  7   9  3
    *                   \
    *                    8
    */

    const results = await collectScanResults(1000);
    console.log(results);

  }).timeout(0);
});

async function collectScanResults(iterations) {
  const points = {};
  const updatePoint = (address) => {
    if (!points[address]) {
      points[address] = 1;
    } else {
      points[address]++;
    }
  }

  function outputCompletion() {
    if (i >= itr) {
      return;
    }
    process.stdout.write(" " + (i/itr * 100).toPrecision(2) + "% complete\r")                                                              
    setTimeout(outputCompletion, 1000);
  }

  let i = 0;

  outputCompletion();

  console.log("collecting scan results for", iterations, "iterations...");

  while (i < iterations) {
    const kp = sodium.crypto_sign_keypair('uint8array');
    const hash = crypto.createHash("sha256");
    hash.update(kp.publicKey);
    hash.update(Buffer.from([0]));
    const point = new BN(hash.digest().subarray(0, 16));
    const address = await directory.scan(point);
    updatePoint(address);
    i++;
  }

  return points;
}