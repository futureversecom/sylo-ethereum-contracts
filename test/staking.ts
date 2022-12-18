import { ethers } from 'hardhat';
import { BigNumber, Signer } from 'ethers';
import {
  Directory,
  EpochsManager,
  Registries,
  RewardsManager,
  StakingManager,
  SyloToken,
  TestSeekers,
} from '../typechain';
import utils from './utils';
import { assert, expect } from 'chai';
// Chi Squared goodness of fit test
import { chi2gof } from '@stdlib/stats';
import crypto from 'crypto';

type Results = { [key: string]: number };

describe('Staking', () => {
  let accounts: Signer[];
  let owner: string;

  let token: SyloToken;
  let epochsManager: EpochsManager;
  let rewardsManager: RewardsManager;
  let directory: Directory;
  let stakingManager: StakingManager;
  let registries: Registries;
  let seekers: TestSeekers;

  const epochId = 1;

  before(async () => {
    accounts = await ethers.getSigners();
    // first account is implicitly used as deployer of contracts in hardhat
    owner = await accounts[0].getAddress();

    const Token = await ethers.getContractFactory('SyloToken');
    token = await Token.deploy();
  });

  beforeEach(async () => {
    const contracts = await utils.initializeContracts(owner, token.address);
    epochsManager = contracts.epochsManager;
    rewardsManager = contracts.rewardsManager;
    directory = contracts.directory;
    stakingManager = contracts.stakingManager;
    registries = contracts.registries;
    seekers = contracts.seekers;

    await token.approve(stakingManager.address, 100000);
  });

  it('should be able to set parameters after initialization', async () => {
    await expect(stakingManager.setUnlockDuration(100))
      .to.emit(stakingManager, 'UnlockDurationUpdated')
      .withArgs(100);

    await expect(stakingManager.setMinimumStakeProportion(3000))
      .to.emit(stakingManager, 'MinimumStakeProportionUpdated')
      .withArgs(3000);

    const unlockDuration = await stakingManager.unlockDuration();
    const minimumStakeProportion =
      await stakingManager.minimumStakeProportion();

    assert.equal(
      unlockDuration.toNumber(),
      100,
      'Expected unlock duration to be correctly set',
    );
    assert.equal(
      minimumStakeProportion,
      3000,
      'Expected minimum node stake to be correctly set',
    );
  });

  it('should be able to join the next epoch and directory at once', async () => {
    await stakingManager.addStake(100, owner);
    await setSeekeRegistry(accounts[0], accounts[1], 1);

    const currentEpochId = (await directory.currentDirectory()).add(1);
    const nextRewardPool = await rewardsManager.getRewardPool(
      currentEpochId,
      owner,
    );
    assert.equal(
      nextRewardPool.initializedAt.toNumber(),
      0,
      'Expected next reward pool to be uninitalized',
    );

    await epochsManager.joinNextEpoch();

    const currentRewardPool = await rewardsManager.getRewardPool(
      currentEpochId,
      owner,
    );
    assert.notEqual(
      currentRewardPool.initializedAt.toNumber(),
      0,
      'Expected reward pool to have been initalized',
    );
  });

  it('should be able to get unlocking duration', async () => {
    await stakingManager.setUnlockDuration(100);
    const unlockDuration = await stakingManager.unlockDuration();
    assert.equal(
      unlockDuration.toNumber(),
      100,
      'Expected unlock duration to be updated',
    );
  });

  it('should be able to stake', async () => {
    const initialBalance = await token.balanceOf(owner);

    await stakingManager.addStake(100, owner);

    const postStakeBalance = await token.balanceOf(owner);

    assert.equal(
      initialBalance.sub(100).toString(),
      postStakeBalance.toString(),
      '100 tokens should be subtracted from initial balance after staking',
    );

    const stakeEntry = await stakingManager.getStakeEntry(owner, owner);

    assert.equal(
      stakeEntry.amount.toString(),
      '100',
      'A stake entry with 100 tokens should be managed by the contract',
    );
  });

  it('should be able to calculate remaining stake that can be added to a stakee', async () => {
    await stakingManager.addStake(111, owner);

    const expectedRemaining = Math.floor(111 / 0.2) - 111;

    const remaining = await stakingManager.calculateMaxAdditionalDelegatedStake(
      owner,
    );

    assert.equal(
      expectedRemaining,
      remaining.toNumber(),
      'Expected remaining additional stake to be correctly calculated',
    );

    // ensure we can actually add that amount
    await token.transfer(await accounts[1].getAddress(), 1000);
    await token.connect(accounts[1]).approve(stakingManager.address, 1000);
    await stakingManager
      .connect(accounts[1])
      .addStake(expectedRemaining, owner);
  });

  it('should fail to calculate remaining stake if owned stake too low', async () => {
    await stakingManager.addStake(100, owner);

    await token.transfer(await accounts[1].getAddress(), 1000);
    await token.connect(accounts[1]).approve(stakingManager.address, 1000);
    await stakingManager.connect(accounts[1]).addStake(100, owner);

    await stakingManager.unlockStake(80, owner);

    await expect(
      stakingManager.calculateMaxAdditionalDelegatedStake(owner),
    ).to.be.revertedWith('Can not add more delegated stake to this stakee');
  });

  it('should not able to add stake to zero address', async () => {
    await expect(
      stakingManager.addStake(100, ethers.constants.AddressZero),
    ).to.be.revertedWith('Address is null');
  });

  it('should not be able to add 0 stake', async () => {
    await expect(stakingManager.addStake(0, owner)).to.be.revertedWith(
      'Stake amount must be more than 0',
    );
  });

  it('should be able to unlock stake', async () => {
    await stakingManager.addStake(100, owner);
    await stakingManager.unlockStake(100, owner);

    const key = await stakingManager.getKey(owner, owner);
    const unlocking = await stakingManager.unlockings(key);
    assert.equal(
      unlocking.amount.toNumber(),
      100,
      'Expected unlocking to exist',
    );
  });

  it('can not unlock no stake', async () => {
    await expect(stakingManager.unlockStake(100, owner)).to.be.revertedWith(
      'Nothing to unstake',
    );
  });

  it('can not unlock zero stake', async () => {
    await stakingManager.addStake(100, owner);
    await expect(stakingManager.unlockStake(0, owner)).to.be.revertedWith(
      'Cannot unlock with zero amount',
    );
  });

  it('can not unlock more stake than exists', async () => {
    await stakingManager.addStake(100, owner);
    await expect(stakingManager.unlockStake(101, owner)).to.be.revertedWith(
      'Cannot unlock more than staked',
    );
  });

  it('should update unlocking state when unlocking more stake', async () => {
    await stakingManager.addStake(100, owner);
    await stakingManager.unlockStake(40, owner);

    const key = await stakingManager.getKey(owner, owner);
    const unlockingOne = await stakingManager.unlockings(key);

    await stakingManager.unlockStake(40, owner);
    const unlockingTwo = await stakingManager.unlockings(key);

    expect(unlockingTwo.unlockAt.toNumber()).to.be.greaterThan(
      unlockingOne.unlockAt.toNumber(),
    );
  });

  it("doesn't update unlock at if existing unlock will unlock later", async () => {
    await stakingManager.addStake(100, owner);
    await stakingManager.unlockStake(40, owner);

    const key = await stakingManager.getKey(owner, owner);
    const unlockingOne = await stakingManager.unlockings(key);

    // we sit the unlock duration to a shorter value here
    await stakingManager.setUnlockDuration(1);

    await stakingManager.unlockStake(40, owner);
    const unlockingTwo = await stakingManager.unlockings(key);

    // expect the second unlocking to not overwrite the original one
    expect(unlockingTwo.unlockAt.toNumber()).to.be.equal(
      unlockingOne.unlockAt.toNumber(),
    );
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

    await utils.advanceBlock(11);

    await stakingManager.withdrawStake(owner);

    const postWithdrawBalance = await token.balanceOf(owner);

    assert.equal(
      initialBalance.toString(),
      postWithdrawBalance.toString(),
      'Balance should be equal to initial balance after withdrawing',
    );
  });

  it("should not be able to withdraw stake that hasn't unlocked", async () => {
    await stakingManager.addStake(100, owner);
    await stakingManager.unlockStake(100, owner);
    await expect(stakingManager.withdrawStake(owner)).to.be.revertedWith(
      'Stake not yet unlocked',
    );
  });

  it('should be able to cancel unlocking', async () => {
    await stakingManager.addStake(100, owner);
    await stakingManager.unlockStake(100, owner);
    await stakingManager.cancelUnlocking(100, owner);

    const key = await stakingManager.getKey(owner, owner);
    const unlocking = await stakingManager.unlockings(key);

    assert.equal(
      unlocking.amount.toNumber(),
      0,
      'Expected unlocking to be cancelled',
    );
  });

  it('should be able to cancel a portion of the unlocking', async () => {
    await stakingManager.addStake(100, owner);
    await stakingManager.unlockStake(100, owner);
    await stakingManager.cancelUnlocking(54, owner);

    const key = await stakingManager.getKey(owner, owner);
    const unlocking = await stakingManager.unlockings(key);

    assert.equal(
      unlocking.amount.toNumber(),
      46,
      'Expected only a portion of the unlocking to be cancelled',
    );
  });

  it('unlocking more than exists clears entire stake', async () => {
    await stakingManager.addStake(100, owner);
    await stakingManager.unlockStake(100, owner);
    await stakingManager.cancelUnlocking(101, owner);

    const key = await stakingManager.getKey(owner, owner);
    const unlocking = await stakingManager.unlockings(key);

    assert.equal(
      unlocking.amount.toNumber(),
      0,
      'Expected unlocking to be cancelled',
    );
  });

  it('should allow delegated stake to exceed minimum owned stake by the stakee', async () => {
    await token.transfer(await accounts[1].getAddress(), 1000);
    await token.connect(accounts[1]).approve(stakingManager.address, 1000);
    await stakingManager.connect(accounts[1]).addStake(180, owner);
  });

  it('should not allow directory to be joined with no stake', async () => {
    await directory.addManager(owner);
    await expect(directory.joinNextDirectory(owner)).to.be.revertedWith(
      'Can not join directory for next epoch without any stake',
    );
  });

  it('can not join directory after unlocking all stake', async () => {
    await stakingManager.addStake(1, owner);
    await stakingManager.unlockStake(1, owner);

    await directory.addManager(owner);

    await expect(directory.joinNextDirectory(owner)).to.be.revertedWith(
      'Can not join directory for next epoch without any stake',
    );
  });

  it('should reduce stake when joining directory with less than minimum stake', async () => {
    await stakingManager.addStake(100, owner);

    await token.transfer(await accounts[1].getAddress(), 1000);
    await token.connect(accounts[1]).approve(stakingManager.address, 1000);
    await stakingManager.connect(accounts[1]).addStake(180, owner);

    // after unlocking, Node will own less than 20% of stake
    await stakingManager.unlockStake(80, owner);

    await directory.addManager(owner);
    await directory.joinNextDirectory(owner);

    // the node now only owns 10% of the stake, which 50% of the
    // minimum stake proportion
    // the stake that the node joined with should be 50% of the managed
    // stake
    const joinedStake = await directory.getTotalStakeForStakee(1, owner);
    const managedStake = await stakingManager.getStakeeTotalManagedStake(owner);

    const meetsMinimum = await stakingManager.checkMinimumStakeProportion(
      owner,
    );
    expect(meetsMinimum).to.equal(false);

    expect(managedStake.div(2).toString()).to.equal(joinedStake);
  });

  it('should fail to join when node`s own stake is 0', async () => {
    await token.transfer(await accounts[1].getAddress(), 1000);
    await token.connect(accounts[1]).approve(stakingManager.address, 1000);
    await stakingManager.connect(accounts[1]).addStake(180, owner);

    await directory.addManager(owner);

    await expect(directory.joinNextDirectory(owner)).to.be.revertedWith(
      'Can not join directory for next epoch without any stake',
    );
  });

  it('should be able to get total stake for a stakee', async () => {
    await stakingManager.addStake(100, owner);
    for (let i = 2; i < 10; i++) {
      await token.transfer(await accounts[i].getAddress(), 1000);
      await token.connect(accounts[i]).approve(stakingManager.address, 1000);
      await stakingManager.connect(accounts[i]).addStake(10, owner);

      const stakeAmount = await stakingManager.getCurrentStakerAmount(
        owner,
        await accounts[i].getAddress(),
      );
      assert.equal(
        stakeAmount.toString(),
        '10',
        "Expected contract to hold staker's stake",
      );
    }

    const totalStake = await stakingManager.getStakeeTotalManagedStake(owner);

    assert.equal(
      totalStake.toString(),
      '180',
      'Expected contract to track all stake entries',
    );

    const meetsMinimum = await stakingManager.checkMinimumStakeProportion(
      owner,
    );
    expect(meetsMinimum).to.equal(true);
  });

  it('should store the epochId the stake entry was updated at', async () => {
    await directory.transferOwnership(epochsManager.address);
    await epochsManager.initializeEpoch({ from: owner });

    await stakingManager.addStake(100, owner);

    const stakeEntry = await stakingManager.getStakeEntry(owner, owner);

    assert.equal(
      stakeEntry.epochId.toNumber(),
      1,
      'Stake entry should track the epoch id it was updated at',
    );
  });

  it('should not be able to join directory without stake', async () => {
    await setSeekeRegistry(accounts[0], accounts[1], 1);
    await expect(epochsManager.joinNextEpoch()).to.be.revertedWith(
      'Must have stake to initialize a reward pool',
    );
  });

  it('should not be able to join directory without setting seeker account', async () => {
    await expect(epochsManager.joinNextEpoch()).to.be.revertedWith(
      'Node must have a valid seeker account to join an epoch',
    );
  });

  it('should not be able to join directory when seeker account is not seeker owner', async () => {
    await setSeekeRegistry(accounts[0], accounts[1], 1);
    await stakingManager.addStake(1, owner);

    await seekers
      .connect(accounts[1])
      .transferFrom(
        await accounts[1].getAddress(),
        await accounts[2].getAddress(),
        1,
      );

    await expect(epochsManager.joinNextEpoch()).to.be.revertedWith(
      "Node's seeker account does not match the current seeker owner",
    );
  });

  it('should not be able to join the next epoch more than once', async () => {
    await stakingManager.addStake(1, owner);
    await directory.addManager(owner);
    await directory.joinNextDirectory(owner);
    await expect(directory.joinNextDirectory(owner)).to.be.revertedWith(
      'Can only join the directory once per epoch',
    );
  });

  it('should be able to scan after joining directory', async () => {
    await setSeekeRegistry(accounts[0], accounts[1], 1);
    await stakingManager.addStake(1, owner);
    await epochsManager.joinNextEpoch();

    await directory.addManager(owner);
    await directory.setCurrentDirectory(epochId);

    await directory.scan(0);
  });

  it('should be able to scan with epoch id after joining directory', async () => {
    await setSeekeRegistry(accounts[0], accounts[1], 1);
    await stakingManager.addStake(1, owner);
    await epochsManager.joinNextEpoch();

    await directory.addManager(owner);
    await directory.setCurrentDirectory(epochId);

    await directory.scanWithEpochId(0, epochId);
  });

  it('should be able to scan empty directory', async () => {
    await directory.addManager(owner);
    await directory.setCurrentDirectory(epochId);

    const address = await directory.scan(0);

    assert.equal(
      address.toString(),
      ethers.constants.AddressZero,
      'Expected empty directory to scan to zero address',
    );
  });

  it('should be able to query properties of directory', async () => {
    let expectedTotalStake = 0;
    for (let i = 0; i < accounts.length; i++) {
      await addStakeAndJoinEpoch(accounts[i], 1, i);
      expectedTotalStake += 1;
      const stake = await directory.getTotalStakeForStakee(
        1,
        await accounts[i].getAddress(),
      );
      assert.equal(
        stake.toNumber(),
        1,
        'Expected to be able to query total stake for stakee',
      );
    }

    await directory.addManager(owner);
    await directory.setCurrentDirectory(epochId);

    const totalStake = await directory.getTotalStake(1);
    assert.equal(
      totalStake.toNumber(),
      expectedTotalStake,
      'Expected to return correct amount for total stake query',
    );

    const entries = await directory.getEntries(1);
    for (let i = 0; i < accounts.length; i++) {
      const address = entries[0][i];
      const boundary = entries[1][i];
      assert.equal(
        address,
        await accounts[i].getAddress(),
        'Expected entry to hold correct address',
      );
      assert.equal(
        boundary.toNumber(),
        i + 1,
        'Expected entry to hold correct boundary value',
      );
    }
  });

  it('should be able to get total managed stake', async () => {
    let expectedTotalStake = 0;
    for (let i = 0; i < accounts.length; i++) {
      await token.transfer(await accounts[i].getAddress(), 100);
      await token.connect(accounts[i]).approve(stakingManager.address, 100);
      await stakingManager
        .connect(accounts[i])
        .addStake(i + 1, await accounts[i].getAddress());
      expectedTotalStake += i + 1;
    }

    const totalManagedStake = await stakingManager.getTotalManagedStake();

    assert.equal(
      totalManagedStake.toNumber(),
      expectedTotalStake,
      'Expected to be able to query for total managed stake',
    );
  });

  it('should correctly scan accounts based on their stake proportions', async () => {
    for (let i = 0; i < 5; i++) {
      await addStakeAndJoinEpoch(accounts[i], 1, i);
    }

    await directory.addManager(owner);
    await directory.setCurrentDirectory(epochId);

    const fifthPoint = BigNumber.from(2).pow(128).sub(1).div(5);
    const points = [
      '0',
      fifthPoint.add(1).toString(),
      fifthPoint.mul(2).add(2).toString(),
      fifthPoint.mul(3).add(3).toString(),
      fifthPoint.mul(4).add(4).toString(),
    ];

    for (let i = 0; i < 5; i++) {
      // check scan
      const address = await directory.scan(points[i]);
      assert.equal(
        address,
        await accounts[i].getAddress(),
        'Expected scan to return correct result',
      );

      // check scan with epoch id
      const addressWithEpochId = await directory.scanWithEpochId(
        points[i],
        epochId,
      );
      assert.equal(
        addressWithEpochId,
        await accounts[i].getAddress(),
        'Expected scan with epoch id to return correct result',
      );
    }
  });

  it('should correctly scan with different epoch ids', async () => {
    async function checkScanWithEpochId(
      nodeAddress: string,
      pointValue: string,
      requestEpochId: number,
    ) {
      const address = await directory.scanWithEpochId(
        pointValue,
        requestEpochId,
      );
      assert.equal(
        address.toString(),
        nodeAddress,
        `Expected scan with epoch id to return correct address ${nodeAddress} for epoch ${requestEpochId}`,
      );
    }

    // process epoch 1
    const amountEpochOne = [250, 350, 400];
    for (let i = 0; i < amountEpochOne.length; i++) {
      await addStakeAndJoinEpoch(accounts[i], amountEpochOne[i], i);
    }
    await directory.addManager(owner);
    await directory.setCurrentDirectory(1);
    await epochsManager.initializeEpoch();

    // process epoch 2
    const amountEpochTwo = [50, 100, 100, 300, 450];
    for (let i = 0; i < amountEpochTwo.length; i++) {
      await addStakeAndJoinEpoch(accounts[i], amountEpochTwo[i], i);
    }
    await directory.addManager(owner);
    await directory.setCurrentDirectory(2);
    await epochsManager.initializeEpoch();

    // check point of node 0, epoch 1
    let point = BigNumber.from(2).pow(128).sub(1).div(8);
    await checkScanWithEpochId(
      await accounts[0].getAddress(),
      point.toString(),
      1,
    );

    // check point of node 1, epoch 1
    point = BigNumber.from(2).pow(128).sub(1).div(2);
    await checkScanWithEpochId(
      await accounts[1].getAddress(),
      point.toString(),
      1,
    );

    // check point of node 2, epoch 1
    point = BigNumber.from(2).pow(128).sub(1);
    await checkScanWithEpochId(
      await accounts[2].getAddress(),
      point.toString(),
      1,
    );

    // In epoch 2, the directory tree will be
    //
    // 300 | 450   | 500   | 300   | 450
    // 0%  | 15%   | 37.5% | 62.5% | 77.5%

    // check point of node 1, epoch 2
    point = BigNumber.from(2).pow(128).sub(1).div(4);
    await checkScanWithEpochId(
      await accounts[1].getAddress(),
      point.toString(),
      2,
    );

    // check point of node 3, epoch 2
    point = BigNumber.from(2).pow(128).sub(1).div(4).mul(3);
    await checkScanWithEpochId(
      await accounts[3].getAddress(),
      point.toString(),
      2,
    );
    
    // check epoch 3 - empty directory
    await checkScanWithEpochId(ethers.constants.AddressZero, '10000000', 4);
  });

  it('can not call functions that onlyManager constraint', async () => {
    await expect(directory.joinNextDirectory(owner)).to.be.revertedWith(
      'Only managers of this contract can call this function',
    );
  });

  it('should distribute scan results amongst stakees proportionally - all equal [ @skip-on-coverage ]', async () => {
    const numAccounts = 10;

    let totalStake = 0;
    for (let i = 0; i < numAccounts; i++) {
      await addStakeAndJoinEpoch(accounts[i], 1, i);
      totalStake += 1;
    }

    await directory.addManager(owner);
    await directory.setCurrentDirectory(epochId);

    const iterations = process.env.ITERATIONS
      ? parseInt(process.env.ITERATIONS)
      : 1000;

    console.log(
      `running all equal stake amount distribution test with ${iterations} iterations`,
    );

    const expectedResults: Results = {};
    for (let i = 0; i < numAccounts; i++) {
      expectedResults[await accounts[i].getAddress()] =
        (1 / totalStake) * iterations;
    }

    await testScanResults(iterations, expectedResults);
  }).timeout(0);

  it('should distribute scan results amongst stakees proportionally - varied stake amounts [ @skip-on-coverage ]', async () => {
    const numAccounts = 10;

    let totalStake = 0;
    for (let i = 0; i < numAccounts; i++) {
      await addStakeAndJoinEpoch(accounts[i], i + 1, i);
      totalStake += i + 1;
    }

    await directory.addManager(owner);
    await directory.setCurrentDirectory(epochId);

    const iterations = process.env.ITERATIONS
      ? parseInt(process.env.ITERATIONS)
      : 1000;

    console.log(
      `running varied stake amount distribution test with ${iterations} iterations`,
    );

    const expectedResults: Results = {};
    for (let i = 0; i < numAccounts; i++) {
      expectedResults[await accounts[i].getAddress()] =
        ((i + 1) / totalStake) * iterations;
    }

    await testScanResults(iterations, expectedResults);
  }).timeout(0);

  it('should be able to scan after unlocking all stake [ @skip-on-coverage ]', async () => {
    await stakingManager.addStake(1, owner);

    await token.transfer(await accounts[1].getAddress(), 100);
    await token.connect(accounts[1]).approve(stakingManager.address, 100);
    await stakingManager
      .connect(accounts[1])
      .addStake(1, await accounts[1].getAddress());

    await token.transfer(await accounts[2].getAddress(), 100);
    await token.connect(accounts[2]).approve(stakingManager.address, 100);
    await stakingManager
      .connect(accounts[2])
      .addStake(1, await accounts[2].getAddress());

    await stakingManager.unlockStake(1, owner);
    await stakingManager
      .connect(accounts[1])
      .unlockStake(1, await accounts[1].getAddress());
    await stakingManager
      .connect(accounts[2])
      .unlockStake(1, await accounts[2].getAddress());

    await directory.addManager(owner);
    await directory.setCurrentDirectory(epochId);

    // check scan
    const address = await directory.scan(0);

    assert.equal(
      address,
      ethers.constants.AddressZero,
      'Expected zero address',
    );

    // check scan with epoch id
    const addressWithEpochId = await directory.scanWithEpochId(0, epochId);

    assert.equal(
      addressWithEpochId,
      address,
      'Expected address from scan with epoch id to be the same as address from scan',
    );
  });

  async function setSeekeRegistry(
    account: Signer,
    seekerAccount: Signer,
    tokenId: number,
  ) {
    await utils.setSeekerRegistry(
      registries,
      seekers,
      account,
      seekerAccount,
      tokenId,
    );
  }

  async function addStakeAndJoinEpoch(
    account: Signer,
    amount: number,
    seekerId: number,
  ) {
    await token.transfer(await account.getAddress(), amount);
    await token.connect(account).approve(stakingManager.address, amount);
    await stakingManager
      .connect(account)
      .addStake(amount, await account.getAddress());
    await setSeekeRegistry(account, accounts[9], seekerId);
    await epochsManager.connect(account).joinNextEpoch();
  }

  async function testScanResults(iterations: number, expectedResults: Results) {
    const results = await collectScanResults(iterations);

    const x = [];
    const y = [];

    for (const key of Object.keys(expectedResults)) {
      x.push(results[key]);
      y.push(expectedResults[key]);
    }

    // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
    const chiResult = chi2gof(x, y).toJSON();

    // eslint-disable-next-line @typescript-eslint/no-unsafe-member-access
    if (chiResult.rejected) {
      assert.fail(
        'Expected scan result to pass goodness-of-fit test \n' +
          `Expected: ${JSON.stringify(expectedResults)} \n` +
          `Actual: ${JSON.stringify(results)} \n`,
      );
    }
  }

  async function collectScanResults(iterations: number) {
    const points: Results = {};
    const updatePoint = (address: string) => {
      if (!points[address]) {
        points[address] = 1;
      } else {
        points[address]++;
      }
    };

    function outputCompletion() {
      if (i >= iterations) {
        return;
      }
      process.stdout.write(
        ' ' + ((i / iterations) * 100).toPrecision(2) + '% completed\r',
      );
      setTimeout(outputCompletion, 1000);
    }

    const mnemonic =
      'search topple trouble similar sorry just around connect hello range predict ahead';
    const keys = [];
    for (let i = 0; i < iterations; i++) {
      keys.push(ethers.Wallet.fromMnemonic(mnemonic, `m/44'/60'/0'/${i}`));
    }

    let i = 0;

    outputCompletion();

    console.log('collecting scan results for', iterations, 'iterations...');

    while (i < iterations) {
      const hash = crypto.createHash('sha256');
      hash.update(keys[i].publicKey);
      hash.update(Buffer.from([0])); // append epoch
      const point = BigNumber.from(hash.digest().subarray(0, 16));
      const address = await directory.scan(point);
      updatePoint(address);
      i++;
    }

    return points;
  }
});
