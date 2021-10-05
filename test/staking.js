const BN = require("bn.js");
const crypto = require("crypto");
const sodium = require('libsodium-wrappers-sumo');
const utils = require('./utils.js');

const Token = artifacts.require("SyloToken");

// Chi Squared goodness of fit test
const chi2gof = require('@stdlib/stats/chi2gof');

const utils = require('./utils');

contract('Staking', accounts => {
  let token;
  let epochsManager;
  let rewardsManager;
  let ticketingParameters;
  let ticketing;
  let directory;
  let listings;
  let stakingManager;

  const epochId = 1;

  before(async () => {
    token = await Token.new({ from: accounts[1] });
  });

  beforeEach(async () => {
    const contracts = await utils.initializeContracts(accounts[1], token.address);
    epochsManager = contracts.epochsManager;
    rewardsManager = contracts.rewardsManager;
    ticketingParameters = contracts.ticketingParameters;
    ticketing = contracts.ticketing;
    directory = contracts.directory;
    listings = contracts.listings;
    stakingManager = contracts.stakingManager;

    await token.approve(stakingManager.address, 10000, { from: accounts[1] });
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

    const stakeEntry = await stakingManager.getStakeEntry(accounts[1], accounts[1]);

    assert.equal(
      stakeEntry.amount.toString(),
      '100',
      "A stake entry with 100 tokens should be managed by the contract"
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

  it('should be able to get total stake for a stakee', async () => {
    for (let i = 2; i < 10; i++) {
      await token.transfer(accounts[i], 1000, { from: accounts[1] });
      await token.approve(stakingManager.address, 1000, { from: accounts[i] });
      await stakingManager.addStake(10, accounts[1], { from: accounts[i] });

      const stakeAmount = await stakingManager.getCurrentStakerAmount(accounts[1], accounts[i]);
      assert.equal(
        stakeAmount.toString(),
        '10',
        "Expected contract to hold staker's stake"
      );
    }

    const totalStake = await stakingManager.getStakeeTotalManagedStake(accounts[1]);

    assert.equal(
      totalStake.toString(),
      '80',
      "Expected contract to track all stake entries"
    );
  });

  it('should store the epochId the stake entry was updated at', async () => {
    await directory.transferOwnership(epochsManager.address, { from: accounts[1] });
    await epochsManager.initializeEpoch({ from: accounts[1] });

    await stakingManager.addStake(100, accounts[1], { from: accounts[1] });

    const stakeEntry = await stakingManager.getStakeEntry(accounts[1], accounts[1]);

    assert.equal(
      parseInt(stakeEntry.epochId),
      1,
      "Stake entry should track the epoch id it was updated at"
    );
  })

  it('should be able to scan', async () => {
    await stakingManager.addStake(1, accounts[1], { from: accounts[1] });
    await directory.joinNextDirectory({ from: accounts[1] });
    await directory.setCurrentDirectory(epochId, { from: accounts[1] });

    await directory.scan(new BN(0));
  });

  it.only('should be able to scan empty directory', async () => {
    await directory.constructDirectory({ from: accounts[1] });

    const address = await directory.scan(new BN(0));

    assert.equal(
      address.toString(),
      '0x0000000000000000000000000000000000000000',
      "Expected empty directory to scan to zero address"
    );
  });

  it('should distribute scan results amongst stakees proportionally - all equal [ @skip-on-coverage ]', async () => {
    for (let i = 0; i < accounts.length; i++) {
      await stakingManager.addStake(1, accounts[i], { from: accounts[1] });
      await directory.joinNextDirectory({ from: accounts[i] });
    }

    await directory.setCurrentDirectory(epochId, { from: accounts[1] });

    let expectedResults = {}
    for (let i = 0; i < accounts.length; i++) {
      expectedResults[accounts[i]] = 1/10 * 1000;
    }

    await testScanResults(1000, expectedResults);
  }).timeout(0);

  it('should distribute scan results amongst stakees proportionally - varied stake amounts [ @skip-on-coverage ]', async () => {
    let totalStake = 0;
    for (let i = 0; i < accounts.length; i++) {
      await stakingManager.addStake(i + 1, accounts[i], { from: accounts[1] });
      await directory.joinNextDirectory({ from: accounts[i] });
      totalStake += i + 1;
    }

    await directory.setCurrentDirectory(epochId, { from: accounts[1] });

    let expectedResults = {}
    for (let i = 0; i < accounts.length; i++) {
      expectedResults[accounts[i]] = parseInt((i+1)/totalStake * 1000);
    }

    await testScanResults(1000, expectedResults);
  }).timeout(0);

  it('should be able to scan after unlocking all stake [ @skip-on-coverage ]', async () => {
    await stakingManager.addStake(1, accounts[0], { from: accounts[1] });
    await stakingManager.addStake(1, accounts[1], { from: accounts[1] });
    await stakingManager.addStake(1, accounts[2], { from: accounts[1] });

    await stakingManager.unlockStake(1, accounts[0], { from: accounts[1] });
    await stakingManager.unlockStake(1, accounts[1], { from: accounts[1] });
    await stakingManager.unlockStake(1, accounts[2], { from: accounts[1] });

    await directory.setCurrentDirectory(epochId, { from: accounts[1] });

    const address = await directory.scan(0);

    assert.equal(address, '0x0000000000000000000000000000000000000000', "Expected zero address");
  });

  it('can not join directory without a stake [ @skip-on-coverage ]', async () => {
    await stakingManager.addStake(1, accounts[1], { from: accounts[1] });
    await stakingManager.unlockStake(1, accounts[1], { from: accounts[1] });

    directory.joinNextDirectory({ from: accounts[1] })
      .then(() => {
        assert.fail("Join directory should fail as no stake for this epoch");
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

  async function testScanResults(iterations, expectedResults) {
    const results = await collectScanResults(iterations);

    let x = [];
    let y = [];

    for (let key of Object.keys(expectedResults)) {
      x.push(results[key]);
      y.push(expectedResults[key])
    }

    const chiResult = chi2gof(x, y).toJSON();

    assert.isNotOk(chiResult.rejected, "Expected scan result to pass goodness-of-fit test");
  }

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
