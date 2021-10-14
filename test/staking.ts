import { ethers } from "hardhat";
import { BigNumber, Signer } from "ethers";
import { Directory, EpochsManager, Listings, RewardsManager, StakingManager, SyloTicketing, SyloToken } from "../typechain";
const crypto = require("crypto");
const sodium = require('libsodium-wrappers-sumo');
import utils from './utils';
import { assert, expect } from "chai";

// Chi Squared goodness of fit test
const chi2gof = require('@stdlib/stats/chi2gof');

type Results = { [key: string]: number };

describe('Staking', () => {
  let accounts: Signer[];
  let owner: string;

  let token: SyloToken;
  let epochsManager: EpochsManager;
  let rewardsManager: RewardsManager;
  let ticketing: SyloTicketing;
  let directory: Directory;
  let listings: Listings;
  let stakingManager: StakingManager;

  const epochId = 1;

  before(async () => {
    accounts = await ethers.getSigners();
    // first account is implicitly used as deployer of contracts in hardhat
    owner = await accounts[0].getAddress();

    const Token = await ethers.getContractFactory("SyloToken");
    token = await Token.deploy() as SyloToken;
  });

  beforeEach(async () => {
    const contracts = await utils.initializeContracts(owner, token.address);
    epochsManager = contracts.epochsManager;
    rewardsManager = contracts.rewardsManager;
    ticketing = contracts.ticketing;
    directory = contracts.directory;
    listings = contracts.listings;
    stakingManager = contracts.stakingManager;

    await token.approve(stakingManager.address, 100000);
  });

  it('should be able to get unlocking duration', async () => {
    await stakingManager.setUnlockDuration(100);
    const unlockDuration = await stakingManager.unlockDuration();
    assert.equal(unlockDuration.toNumber(), 100, "Expected unlock duration to be updated");
  });

  it('should be able to stake', async () => {
    const initialBalance = await token.balanceOf(owner);

    await stakingManager.addStake(100, owner);

    const postStakeBalance = await token.balanceOf(owner);

    assert.equal(
      initialBalance.sub(100).toString(),
      postStakeBalance.toString(),
      "100 tokens should be subtracted from initial balance after staking"
    );

    const stakeEntry = await stakingManager.getStakeEntry(owner, owner);

    assert.equal(
      stakeEntry.amount.toString(),
      '100',
      "A stake entry with 100 tokens should be managed by the contract"
    );
  });

  it('should be able to unlock stake', async () => {
    await stakingManager.addStake(100, owner);
    await stakingManager.unlockStake(100, owner);

    const key = await stakingManager.getKey(owner, owner);
    const unlocking = await stakingManager.unlockings(key);
    assert.equal(unlocking.amount.toNumber(), 100, 'Expected unlocking to exist');
  });

  it('should be able to restake when everything is unstaked', async () => {
    await stakingManager.addStake(1, owner);

    await stakingManager.unlockStake(1, owner);

    // Restake
    await stakingManager.addStake(1, owner);
  });

  it('should be able to withdraw stake', async () => {
    const initialBalance = await token.balanceOf(owner);

    await stakingManager.addStake(100, owner);
    await stakingManager.unlockStake(100, owner);
    await stakingManager.withdrawStake(owner);

    const postWithdrawBalance = await token.balanceOf(owner);

    assert.equal(
      initialBalance.toString(),
      postWithdrawBalance.toString(),
      "Balance should be equal to initial balance after withdrawing"
    );
  });

  it('should be able to get total stake for a stakee', async () => {
    for (let i = 2; i < 10; i++) {
      await token.transfer(await accounts[i].getAddress(), 1000);
      await token.connect(accounts[i]).approve(stakingManager.address, 1000);
      await stakingManager.connect(accounts[i]).addStake(10, owner);

      const stakeAmount = await stakingManager.getCurrentStakerAmount(owner, await accounts[i].getAddress());
      assert.equal(
        stakeAmount.toString(),
        '10',
        "Expected contract to hold staker's stake"
      );
    }

    const totalStake = await stakingManager.getStakeeTotalManagedStake(owner);

    assert.equal(
      totalStake.toString(),
      '80',
      "Expected contract to track all stake entries"
    );
  });

  it('should store the epochId the stake entry was updated at', async () => {
    await directory.transferOwnership(epochsManager.address);
    await epochsManager.initializeEpoch({ from: owner });

    await stakingManager.addStake(100, owner);

    const stakeEntry = await stakingManager.getStakeEntry(owner, owner);

    assert.equal(
      stakeEntry.epochId.toNumber(),
      1,
      "Stake entry should track the epoch id it was updated at"
    );
  });

  it('should not be able to join directory without stake', async () => {
    await directory.joinNextDirectory({ from: owner })
      .then(() => {
        assert.fail('Joining directory should fail without stake');
      })
      .catch(e => {
        assert.include(e.message, 'Can not join directory for next epoch without any stake');
      });
  });

  it('should be able to scan after joining directory', async () => {
    await stakingManager.addStake(1, owner);
    await directory.joinNextDirectory({ from: owner });
    await directory.setCurrentDirectory(epochId);

    await directory.scan(0);
  });

  it('should not be able to join directory more than once per epoch', async () => {
    await stakingManager.addStake(1, owner);
    await directory.joinNextDirectory({ from: owner });
    await expect(directory.joinNextDirectory({ from: owner }))
      .to.be.revertedWith('Can only join the directory once per epoch');
  });

  it('should be able to scan empty directory', async () => {
    await directory.setCurrentDirectory(epochId);

    const address = await directory.scan(0);

    assert.equal(
      address.toString(),
      '0x0000000000000000000000000000000000000000',
      "Expected empty directory to scan to zero address"
    );
  });

  it('should be able to query properties of directory', async () => {
    let expectedTotalStake = 0;
    for (let i = 0; i < accounts.length; i++) {
      await stakingManager.addStake(1, await accounts[i].getAddress());
      await directory.connect(accounts[i]).joinNextDirectory();
      expectedTotalStake += 1;
      const stake = await directory.getTotalStakeForStakee(1, await accounts[i].getAddress());
      assert.equal(
        stake.toNumber(),
        1,
        "Expected to be able to query total stake for stakee"
      );
    }

    await directory.setCurrentDirectory(epochId);

    const totalStake = await directory.getTotalStake(1);
    assert.equal(
      totalStake.toNumber(),
      expectedTotalStake,
      "Expected to return correct amount for total stake query"
    );

    const entries = await directory.getEntries(1);
    for (let i = 0; i < accounts.length; i++) {
      const address = entries[0][i];
      const boundary = entries[1][i];
      assert.equal(
        address,
        await accounts[i].getAddress(),
        "Expected entry to hold correct address"
      );
      assert.equal(
        boundary.toNumber(),
        i + 1,
        "Expected entry to hold correct boundary value"
      );
    }
  });

  it('should correctly scan accounts based on their stake proportions', async () => {
    for (let i = 0; i < 4; i++) {
      await stakingManager.addStake(1, await accounts[i].getAddress());
      await directory.connect(accounts[i]).joinNextDirectory();
    }

    await directory.setCurrentDirectory(epochId);

    const quarterPoint = BigNumber.from(2).pow(128).sub(1).div(4);
    const points = [
      '0',
      quarterPoint.add(1).toString(),
      quarterPoint.add(1).add(quarterPoint).add(1).toString(),
      quarterPoint.add(1).add(quarterPoint).add(1).add(quarterPoint).add(1).toString()
    ];

    for (let i = 0; i < 4; i++) {
      const address = await directory.scan(points[i]);
      assert.equal(address, await accounts[i].getAddress(), "Expected scan to return correct result");
    }
  });

  it('should distribute scan results amongst stakees proportionally - all equal [ @skip-on-coverage ]', async () => {
    for (let i = 0; i < accounts.length; i++) {
      await stakingManager.addStake(1, await accounts[i].getAddress());
      await directory.connect(accounts[i]).joinNextDirectory();
    }

    await directory.setCurrentDirectory(epochId);

    let expectedResults: Results = {};
    for (let i = 0; i < accounts.length; i++) {
      expectedResults[await accounts[i].getAddress()] = 1/accounts.length * 3000;
    }

    await testScanResults(3000, expectedResults);
  }).timeout(0);

  it('should distribute scan results amongst stakees proportionally - varied stake amounts [ @skip-on-coverage ]', async () => {
    let totalStake = 0;
    for (let i = 0; i < accounts.length; i++) {
      await stakingManager.addStake(i + 1, await accounts[i].getAddress());
      await directory.connect(accounts[i]).joinNextDirectory();
      totalStake += i + 1;
    }

    await directory.setCurrentDirectory(epochId);

    let expectedResults: Results = {};
    for (let i = 0; i < accounts.length; i++) {
      expectedResults[await accounts[i].getAddress()] = (i+1)/totalStake * 3000;
    }

    await testScanResults(3000, expectedResults);
  }).timeout(0);

  it('should be able to scan after unlocking all stake [ @skip-on-coverage ]', async () => {
    await stakingManager.addStake(1, owner);
    await stakingManager.addStake(1, await accounts[1].getAddress());
    await stakingManager.addStake(1, await accounts[2].getAddress());

    await stakingManager.unlockStake(1, owner);
    await stakingManager.unlockStake(1, await accounts[1].getAddress());
    await stakingManager.unlockStake(1, await accounts[2].getAddress());

    await directory.setCurrentDirectory(epochId);

    const address = await directory.scan(0);

    assert.equal(address, '0x0000000000000000000000000000000000000000', "Expected zero address");
  });

  it('can not join directory without a stake [ @skip-on-coverage ]', async () => {
    await stakingManager.addStake(1, owner);
    await stakingManager.unlockStake(1, owner);

    directory.joinNextDirectory({ from: owner })
      .then(() => {
        assert.fail("Join directory should fail as no stake for this epoch");
      })
      .catch(e => {
        assert.include(e.message, "Can not join directory for next epoch without any stake");
      })
  });

  async function testScanResults(iterations: number, expectedResults: { [key: string]: number }) {
    const results = await collectScanResults(iterations);

    let x = [];
    let y = [];

    for (let key of Object.keys(expectedResults)) {
      x.push(results[key]);
      y.push(expectedResults[key])
    }

    const chiResult = chi2gof(x, y).toJSON();

    if (chiResult.rejected) {
      assert.fail("Expected scan result to pass goodness-of-fit test");
    }
  }

  async function collectScanResults(iterations: number) {
    const points: { [key:string]: number } = {};
    const updatePoint = (address: string) => {
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
      const point = BigNumber.from(hash.digest().subarray(0, 16));
      const address = await directory.scan(point);
      updatePoint(address);
      i++;
    }

    return points;
  }
});
