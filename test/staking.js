const BN = require("bn.js");
const crypto = require("crypto");
const sodium = require('libsodium-wrappers-sumo');

const Directory = artifacts.require("Directory");
const StakingManager = artifacts.require("StakingManager");
const Token = artifacts.require("SyloToken");

const utils = require('./utils');

contract.only('Staking', accounts => {
  let token;
  let stakingManager;
  let directory;

  const epochId = '0x1000000000000000000000000000000000000000000000000000000000000000';

  before(async () => {
    token = await Token.new({ from: accounts[1] });
  });

  beforeEach(async () => {
    stakingManager = await StakingManager.new({ from: accounts[1] });
    await stakingManager.initialize(token.address, 0, { from: accounts[1] });

    await token.approve(stakingManager.address, 10000, { from: accounts[1] });

    directory = await Directory.new({ from: accounts[1] });
    await directory.initialize(
        stakingManager.address,
      { from: accounts[1] }
    );

    await directory.setCurrentDirectory(epochId, { from: accounts[1] });
  });

  it('should be able to get unlocking duration', async () => {
    await stakingManager.setUnlockDuration(100, { from: accounts[1] });
    const unlockDuration = await stakingManager.unlockDuration();
    assert.equal(unlockDuration.toNumber(), 100, "Expected unlock duration to be updated");
  });

  it('should be able to stake', async () => {
    const initialBalance = await token.balanceOf(accounts[1]);

    await stakingManager.addStake(100, accounts[1], { from: accounts[1] });

    const postStakeBalance = await token.balanceOf(accounts[1]);

    assert.equal(
      initialBalance.sub(new BN(100)).toString(),
      postStakeBalance.toString(),
      "100 tokens should be subtracted from initial balance after staking"
    );
  });

  it('should be able to unlock stake', async () => {
    await stakingManager.addStake(100, accounts[1], { from: accounts[1] });
    await stakingManager.unlockStake(100, accounts[1], { from: accounts[1] });

    const key = await stakingManager.getKey(accounts[1], accounts[1]);
    const unlocking = await stakingManager.unlockings(key);
    assert.equal(unlocking.amount.toNumber(), 100, 'Expected unlocking to exist');
  });

  it('should be able to restake when everything is unstaked', async () => {
    await stakingManager.addStake(1, accounts[1], { from: accounts[1] });

    await stakingManager.unlockStake(1, accounts[1], { from: accounts[1] });

    // Restake
    await stakingManager.addStake(1, accounts[1], { from: accounts[1] });
  });

  it('should be able to withdraw stake', async () => {
    const initialBalance = await token.balanceOf(accounts[1]);

    await stakingManager.addStake(100, accounts[1], { from: accounts[1] });
    await stakingManager.unlockStake(100, accounts[1], { from: accounts[1] });
    await stakingManager.withdrawStake(accounts[1], { from: accounts[1] });

    const postWithdrawBalance = await token.balanceOf(accounts[1]);

    assert.equal(
      initialBalance.toString(),
      postWithdrawBalance.toString(),
      "Balance should be equal to initial balance after withdrawing"
    );
  });

  it('should be able to cancel withdraw', async () => {
    await stakingManager.addStake(100, accounts[1], { from: accounts[1] });
    await stakingManager.unlockStake(100, accounts[1], { from: accounts[1] });
    await stakingManager.cancelUnlocking(50, accounts[1], { from: accounts[1] });

    const key = await stakingManager.getKey(accounts[1], accounts[1]);
    let unlocking = await stakingManager.unlockings(key);
    assert.equal(unlocking.amount.toNumber(), 50, 'Expected 50 of the unlocking amount to be cancelled');

    await stakingManager.cancelUnlocking(50, accounts[1], { from: accounts[1] });
    unlocking = await stakingManager.unlockings(key);
    assert.equal(unlocking.amount.toNumber(), 0, 'Expected entire unlocking amount to be cancelled');
  });

  it('should be able to scan', async () => {
    await stakingManager.addStake(1, accounts[1], { from: accounts[1] });
    await directory.joinDirectory(epochId, { from: accounts[1] });

    await directory.scan(new BN(0));
  });

  it('should distribute scan results amongst stakees proportionally - all equal [ @skip-on-coverage ]', async () => {
    for (let i = 0; i < accounts.length; i++) {
      await stakingManager.addStake(1, accounts[i], { from: accounts[1] });
      await directory.joinDirectory(epochId, { from: accounts[i] });
    }

    console.log(await directory.getEntries(epochId));

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

  it('should distribute scan results amongst stakees proportionally - varied stake amounts [ @skip-on-coverage ]', async () => {
    let totalStake = 0;
    for (let i = 0; i < accounts.length; i++) {
      await stakingManager.addStake(i + 1, accounts[i], { from: accounts[1] });
      await directory.joinDirectory(epochId, { from: accounts[i] });
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

  it('should be able to scan after unlocking all stake [ @skip-on-coverage ]', async () => {
    await stakingManager.addStake(1, accounts[0], { from: accounts[1] });
    await stakingManager.addStake(1, accounts[1], { from: accounts[1] });
    await stakingManager.addStake(1, accounts[2], { from: accounts[1] });

    await stakingManager.unlockStake(1, accounts[0], { from: accounts[1] });
    await stakingManager.unlockStake(1, accounts[1], { from: accounts[1] });
    await stakingManager.unlockStake(1, accounts[2], { from: accounts[1] });

    const address = await directory.scan(0);

    assert.equal(address, '0x0000000000000000000000000000000000000000', "Expected zero address");
  });

  it('can not join directory without a stake', async () => {
    await stakingManager.addStake(1, accounts[1], { from: accounts[1] });
    await stakingManager.unlockStake(1, accounts[1], { from: accounts[1] });

    directory.joinDirectory(epochId, { from: accounts[1] })
      .then(() => {
        assert.fail("Join directory should fail as no epoch");
      })
      .catch(e => {
        assert.include(e.message, "Can not join directory for next epoch without any stake");
      })

  });

  it.skip('excludes nodes from directory without a vote', async () => {
    for (let i = 0; i < accounts.length; i++) {
      await stakingManager.addStake(1, accounts[i], { from: accounts[1] });
    }

    for (let i = 0; i < accounts.length - 1; i++) {
      await priceVoting.vote(1, { from: accounts[i] });
    }

    await directory.constructDirectory({ from: accounts[1] });

    let found = false;
    for (let i = 0; i < accounts.length; i++) {
      try {
        const a = await directory.currentDirectory(i);
        if (a == accounts[9]) {
          found = true;
          break;
        }
      } catch(e) {}
    }

    assert.equal(found, false, "The account with no vote should not exist in the directory");
  });

  it.skip('excludes nodes from directory with too high voted price', async () => {
    for (let i = 0; i < accounts.length; i++) {
      await stakingManager.addStake(1, accounts[i], { from: accounts[1] });
    }

    for (let i = 0; i < accounts.length - 1; i++) {
      await priceVoting.vote(1, { from: accounts[i] });
    }

    await priceVoting.vote(5, { from: accounts[5] });

    await directory.constructDirectory({ from: accounts[1] });

    let found = false;
    for (let i = 0; i < accounts.length; i++) {
      try {
        const a = await directory.currentDirectory(i);
        if (a == accounts[5]) {
          found = true;
          break;
        }
      } catch(e) {}
    }
    assert.equal(found, false, "The account with no vote should not exist in the directory");
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
