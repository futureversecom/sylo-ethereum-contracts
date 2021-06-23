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
    directory = await Directory.new({ from: accounts[0] });
    await directory.initialize(token.address, 0, { from: accounts[0] });
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

  it('should distribute scan results amongst stakees proportionally - all equal', async () => {
    let totalStake = 0;
    for (let i = 0; i < accounts.length; i++) {
      await directory.addStake(1, accounts[i], { from: accounts[1] });
      totalStake++;
    }

    let expectedResults = {}
    for (let i = 0; i < accounts.length; i++) {
      expectedResults[accounts[i]] = 1/10 * 1000;
    }

    const results = await collectScanResults(1000);
    for (let key of Object.keys(expectedResults)) {
      const expected = expectedResults[key];
      const actual = results[key];
      console.log('For address', key, 'expected=', expected, 'actual=', actual);
    }
  }).timeout(0);

  it('should distribute scan results amongst stakees proportionally - varied stake amounts', async () => {
    let totalStake = 0;
    for (let i = 0; i < accounts.length; i++) {
      await directory.addStake(i + 1, accounts[i], { from: accounts[1] });
      totalStake += i + 1;
    }

    let expectedResults = {}
    for (let i = 0; i < accounts.length; i++) {
      expectedResults[accounts[i]] = parseInt((i+1)/totalStake * 1000);
    }

    const results = await collectScanResults(1000);
    for (let key of Object.keys(expectedResults)) {
      const expected = expectedResults[key];
      const actual = results[key];
      console.log('For address', key, 'expected=', expected, 'actual=', actual);
    }
  }).timeout(0);

  it('should be able to scan after unlocking all stake', async () => {
    await directory.addStake(1, accounts[0], { from: accounts[1] });
    await directory.addStake(1, accounts[1], { from: accounts[1] });
    await directory.addStake(1, accounts[2], { from: accounts[1] });

    await directory.unlockStake(1, accounts[0], { from: accounts[1] });
    await directory.unlockStake(1, accounts[1], { from: accounts[1] });
    await directory.unlockStake(1, accounts[2], { from: accounts[1] });

    const address = await directory.scan(0);

    assert.equal(address, '0x0000000000000000000000000000000000000000', "Expected zero address");
  });

  it('scan after root unstakes', async () => {
    for (let i = 0; i < accounts.length; i++) {
      await directory.addStake(1, accounts[i], { from: accounts[1] });
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

    await directory.unlockStake(1, accounts[0], { from: accounts[1] });

    /***
    *             8
    *           /   \
    *          2     1
    *         / |   | \
    *        6  7   5  3
    *                \ 
    *                 9   
    */

    let expectedResults = {}
    for (let i = 1; i < accounts.length; i++) {
      expectedResults[accounts[i]] = parseInt(1/9 * 1000);
    }

    const results = await collectScanResults(1000);
    for (let key of Object.keys(expectedResults)) {
      const expected = expectedResults[key];
      const actual = results[key];
      console.log('For address', key, 'expected=', expected, 'actual=', actual);
    }
  }).timeout(0);

  it('scan after nodes where parent is root unstakes', async () => {
    for (let i = 0; i < accounts.length; i++) {
      await directory.addStake(1, accounts[i], { from: accounts[1] });
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

    await directory.unlockStake(1, accounts[1], { from: accounts[1] });

    /***
    *             0
    *           /   \
    *          2     8
    *         / |   | \
    *        6  4   5  3
    *            \   \ 
    *             7   9 
    */

    await directory.unlockStake(1, accounts[2], { from: accounts[1] });

    /***
    *             0
    *           /   \
    *          7     8
    *         / |   | \
    *        6  4   5  3
    *                \ 
    *                 9 
    */

    let expectedResults = {}
    for (let i = 0; i < accounts.length; i++) {
      if (i == 1 || i == 2) continue;
      expectedResults[accounts[i]] = parseInt(1/8 * 1000);
    }

    const results = await collectScanResults(1000);
    for (let key of Object.keys(expectedResults)) {
      const expected = expectedResults[key];
      const actual = results[key];
      console.log('For address', key, 'expected=', expected, 'actual=', actual);
    }
  }).timeout(0);

  it('scan after leaf nodes unstake', async () => {
    for (let i = 0; i < accounts.length; i++) {
      await directory.addStake(1, accounts[i], { from: accounts[1] });
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

    await directory.unlockStake(1, accounts[7], { from: accounts[1] });

    /***
    *             0
    *           /   \
    *          2     8
    *         / |   | \
    *        6  4   5  3
    *                \  \
    *                 9  8
    */

    await directory.unlockStake(1, accounts[8], { from: accounts[1] });

    /***
    *             0
    *           /   \
    *          7     8
    *         / |   | \
    *        6  4   5  3
    *                \ 
    *                 9 
    */

    let expectedResults = {}
    for (let i = 0; i < accounts.length; i++) {
      if (i == 7 || i == 8) continue;
      expectedResults[accounts[i]] = parseInt(1/8 * 1000);
    }

    const results = await collectScanResults(1000);
    for (let key of Object.keys(expectedResults)) {
      const expected = expectedResults[key];
      const actual = results[key];
      console.log('For address', key, 'expected=', expected, 'actual=', actual);
    }
  }).timeout(0);

  it('scan after nodes where the child is a leaf unstake', async () => {
    for (let i = 0; i < accounts.length; i++) {
      await directory.addStake(1, accounts[i], { from: accounts[1] });
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

    await directory.unlockStake(1, accounts[4], { from: accounts[1] });

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

     let expectedResults = {}
     for (let i = 0; i < accounts.length; i++) {
       if (i == 4 || i == 5) continue;
       expectedResults[accounts[i]] = 1/8 * 1000;
     }
 
     const results = await collectScanResults(1000);
     for (let key of Object.keys(expectedResults)) {
       const expected = expectedResults[key];
       const actual = results[key];
       console.log('For address', key, 'expected=', expected, 'actual=', actual);
     }
  }).timeout(0);

  it('scan results for a tree that became severely imbalanced', async () => {
    await directory.addStake(1, accounts[0], { from: accounts[1] });
    await directory.addStake(20, accounts[1], { from: accounts[1] });
    for (let i = 2; i < accounts.length; i++) {
      await directory.addStake(1, accounts[i], { from: accounts[1] });
    }

    /***
    *             0
    *           /   \
    *          2     1(20)
    *         /  \   
    *        4    3   
    *       / \   | \     
    *      8   6  7  5
    *                 \
    *                  9
    */

    let expectedResults = {}
    expectedResults[accounts[0]] = 1/29 * 1000;
    expectedResults[accounts[1]] = 20/29 * 1000;
    for (let i = 2; i < accounts.length; i++) {
      expectedResults[accounts[i]] = 1/29 * 1000;
    }

    console.log("checking results before unlocking node(1)")
    let results = await collectScanResults(1000);
    for (let key of Object.keys(expectedResults)) {
      const expected = expectedResults[key];
      const actual = results[key];
      console.log('For address', key, 'expected=', expected, 'actual=', actual);
    }

    await directory.unlockStake(20, accounts[1], { from: accounts[1] });

    /***
    *             0
    *           / 
    *          2 
    *         /  \   
    *        4    3   
    *       / \   | \     
    *      8   6  7  5
    *                 \
    *                  9
    */

    expectedResults = {}
    for (let i = 0; i < accounts.length; i++) {
      if (i == 1) continue;
      expectedResults[accounts[i]] = 1/9 * 1000;
    }
 
    console.log("checking results after unlocking node(1)")
    results = await collectScanResults(1000);
    for (let key of Object.keys(expectedResults)) {
      const expected = expectedResults[key];
      const actual = results[key];
      console.log('For address', key, 'expected=', expected, 'actual=', actual);
    }
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
      if (i >= iterations) {
        return;
      }
      process.stdout.write(" " + (i/iterations * 100).toPrecision(2) + "% completed\r")                                                              
      setTimeout(outputCompletion, 1000);
    }
  
    let i = 0;
  
    outputCompletion();
  
    console.log("collecting scan results for", iterations, "iterations...");
  
    while (i < iterations) {
      // generate a random ed25519 key and hash with an epoch to create a
      // 'random' point value
      const kp = sodium.crypto_sign_keypair('uint8array');
      const hash = crypto.createHash("sha256");
      hash.update(kp.publicKey);
      hash.update(Buffer.from([0])); // append epoch
      const point = new BN(hash.digest().subarray(0, 16));
      const address = await directory.scan(point);
      updatePoint(address);
      i++;
    }
  
    return points;
  }
});