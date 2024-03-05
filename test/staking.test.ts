import { ethers } from 'hardhat';
import { MaxUint256, Signer } from 'ethers';
import {
  Directory,
  EpochsManager,
  Registries,
  RewardsManager,
  SeekerPowerOracle,
  StakingManager,
  SyloToken,
  TestSeekers,
} from '../typechain-types';
import utils from './utils';
import { assert, expect } from 'chai';
// Chi Squared goodness of fit test
import { chi2gof } from '@stdlib/stats';
import crypto from 'crypto';

type Results = { [key: string]: number };

const MAX_SYLO_STAKE = ethers.parseEther('10000000000');

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
  let seekerPowerOracle: SeekerPowerOracle;

  const epochId = 1;

  const defaultSeekerId = 1;

  before(async () => {
    accounts = await ethers.getSigners();
    // first account is implicitly used as deployer of contracts in hardhat
    owner = await accounts[0].getAddress();

    const Token = await ethers.getContractFactory('SyloToken');
    token = await Token.deploy();
  });

  beforeEach(async () => {
    const contracts = await utils.initializeContracts(owner, token);
    epochsManager = contracts.epochsManager;
    rewardsManager = contracts.rewardsManager;
    directory = contracts.directory;
    stakingManager = contracts.stakingManager;
    registries = contracts.registries;
    seekers = contracts.seekers;
    seekerPowerOracle = contracts.seekerPowerOracle;

    await token.approve(await stakingManager.getAddress(), 100000);

    // set the seeker power to max for all tests by default
    await seekerPowerOracle.registerSeekerPowerRestricted(
      defaultSeekerId,
      MaxUint256,
    );
  });

  it('staking manager cannot be intialized twice', async () => {
    await expect(
      stakingManager.initialize(
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        0,
        0,
      ),
    ).to.be.revertedWith('Initializable: contract is already initialized');
  });

  it('staking manager cannot be intialized with arguments', async () => {
    const StakingManager = await ethers.getContractFactory('StakingManager');
    const stakingManager = await StakingManager.deploy();

    await expect(
      stakingManager.initialize(
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        0,
        0,
      ),
    ).to.be.revertedWithCustomError(stakingManager, 'TokenCannotBeZeroAddress');

    await expect(
      stakingManager.initialize(
        await token.getAddress(),
        await epochsManager.getAddress(), // correct contract is rewards manager
        await seekerPowerOracle.getAddress(),
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        0,
        0,
      ),
    )
      .to.be.revertedWithCustomError(
        stakingManager,
        'TargetNotSupportInterface',
      )
      .withArgs('RewardsManager', '0x3db12b5a');

    await expect(
      stakingManager.initialize(
        await token.getAddress(),
        await rewardsManager.getAddress(),
        await epochsManager.getAddress(),
        await seekerPowerOracle.getAddress(),
        0,
        0,
        0,
      ),
    ).to.be.revertedWithCustomError(
      stakingManager,
      'UnlockDurationCannotBeZero',
    );
  });

  it('directory cannot be intialized twice', async () => {
    await expect(
      directory.initialize(ethers.ZeroAddress, ethers.ZeroAddress),
    ).to.be.revertedWith('Initializable: contract is already initialized');
  });

  it('not owner cannot add manager', async () => {
    await expect(
      rewardsManager
        .connect(accounts[1])
        .addManager(await accounts[1].getAddress()),
    ).to.be.revertedWith('Ownable: caller is not the owner');
  });

  it('cannot add manager with zero address', async () => {
    await expect(
      rewardsManager.addManager(ethers.ZeroAddress),
    ).to.be.revertedWithCustomError(
      rewardsManager,
      'ManagerCannotBeZeroAddress',
    );
  });

  it('can remove owner as manager', async () => {
    await rewardsManager.removeManager(owner);
  });

  it('not owner cannot remove manager', async () => {
    await expect(
      rewardsManager
        .connect(accounts[1])
        .removeManager(await accounts[1].getAddress()),
    ).to.be.revertedWith('Ownable: caller is not the owner');
  });

  it('staking manager should be able to set parameters after initialization', async () => {
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
      unlockDuration,
      100n,
      'Expected unlock duration to be correctly set',
    );
    assert.equal(
      minimumStakeProportion,
      3000n,
      'Expected minimum node stake to be correctly set',
    );
  });

  it('staking manager should not be able to set unlock duration with zero value', async () => {
    await expect(
      stakingManager.setUnlockDuration(0),
    ).to.be.revertedWithCustomError(
      stakingManager,
      'UnlockDurationCannotBeZero',
    );
  });

  it('staking manager should not be able to set parameters before initialization', async () => {
    const StakingManager = await ethers.getContractFactory('StakingManager');
    const stakingManager = await StakingManager.deploy();

    await expect(stakingManager.setUnlockDuration(100)).to.be.revertedWith(
      'Ownable: caller is not the owner',
    );

    await expect(
      stakingManager.setMinimumStakeProportion(3000),
    ).to.be.revertedWith('Ownable: caller is not the owner');
  });

  it('not manager cannot set current directory', async () => {
    await expect(
      directory.connect(accounts[1]).setCurrentDirectory(100),
    ).to.be.revertedWithCustomError(directory, 'OnlyManagers');
  });

  it('should be able to join the next epoch and directory at once', async () => {
    await stakingManager.addStake(100, owner);
    await setSeekeRegistry(accounts[0], accounts[1], 1);

    const currentEpochId = (await directory.currentDirectory()) + 1n;
    const nextRewardPool = await rewardsManager.getRewardPool(
      currentEpochId,
      owner,
    );
    assert.equal(
      nextRewardPool.initializedAt,
      0n,
      'Expected next reward pool to be uninitalized',
    );

    await epochsManager.joinNextEpoch();

    const currentRewardPool = await rewardsManager.getRewardPool(
      currentEpochId,
      owner,
    );
    assert.notEqual(
      currentRewardPool.initializedAt,
      0n,
      'Expected reward pool to have been initalized',
    );
  });

  it('should be able to get unlocking duration', async () => {
    await stakingManager.setUnlockDuration(100);
    const unlockDuration = await stakingManager.unlockDuration();
    assert.equal(
      unlockDuration,
      100n,
      'Expected unlock duration to be updated',
    );
  });

  it('should be able to stake', async () => {
    const initialBalance = await token.balanceOf(owner);

    await stakingManager.addStake(100, owner);

    const postStakeBalance = await token.balanceOf(owner);

    assert.equal(
      initialBalance - 100n,
      postStakeBalance,
      '100 tokens should be subtracted from initial balance after staking',
    );

    const stakeEntry = await stakingManager.getStakeEntry(owner, owner);

    assert.equal(
      stakeEntry.amount,
      100n,
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
      BigInt(expectedRemaining),
      remaining,
      'Expected remaining additional stake to be correctly calculated',
    );

    // ensure we can actually add that amount
    await token.transfer(await accounts[1].getAddress(), 1000);
    await token
      .connect(accounts[1])
      .approve(await stakingManager.getAddress(), 1000);
    await stakingManager
      .connect(accounts[1])
      .addStake(expectedRemaining, owner);
  });

  it('should fail to calculate remaining stake if owned stake too low', async () => {
    await stakingManager.addStake(100, owner);

    await token.transfer(await accounts[1].getAddress(), 1000);
    await token
      .connect(accounts[1])
      .approve(await stakingManager.getAddress(), 1000);
    await stakingManager.connect(accounts[1]).addStake(100, owner);

    await stakingManager.unlockStake(80, owner);

    await expect(stakingManager.calculateMaxAdditionalDelegatedStake(owner))
      .to.be.revertedWithCustomError(stakingManager, 'StakeCapacityReached')
      .withArgs(100n, 120n);
  });

  it('cannot calculate remaining stake with invalid arguments', async () => {
    await expect(
      stakingManager.calculateMaxAdditionalDelegatedStake(ethers.ZeroAddress),
    ).to.be.revertedWithCustomError(
      stakingManager,
      'StakeeCannotBeZeroAddress',
    );
  });

  it('should not able to add stake invalid arguments', async () => {
    await expect(
      stakingManager.addStake(100, ethers.ZeroAddress),
    ).to.be.revertedWithCustomError(
      stakingManager,
      'StakeeCannotBeZeroAddress',
    );

    await expect(
      stakingManager.addStake(0, owner),
    ).to.be.revertedWithCustomError(stakingManager, 'CannotStakeZeroAmount');
  });

  it('should be able to unlock stake', async () => {
    await stakingManager.addStake(100, owner);
    await stakingManager.unlockStake(100, owner);

    const key = await stakingManager.getKey(owner, owner);
    const unlocking = await stakingManager.unlockings(key);
    assert.equal(unlocking.amount, 100n, 'Expected unlocking to exist');
  });

  it('can not unlock stake with invalid arguments', async () => {
    await expect(
      stakingManager.unlockStake(100, ethers.ZeroAddress),
    ).to.be.revertedWithCustomError(
      stakingManager,
      'StakeeCannotBeZeroAddress',
    );

    await expect(
      stakingManager.unlockStake(0, owner),
    ).to.be.revertedWithCustomError(stakingManager, 'CannotUnlockZeroAmount');
  });

  it('can not unlock if user has zero stake', async () => {
    await expect(
      stakingManager.unlockStake(100, owner),
    ).to.be.revertedWithCustomError(stakingManager, 'NoStakeToUnlock');
  });

  it('can not unlock more stake than exists', async () => {
    await stakingManager.addStake(100, owner);
    await expect(stakingManager.unlockStake(101, owner))
      .to.be.revertedWithCustomError(
        stakingManager,
        'CannotUnlockMoreThanStaked',
      )
      .withArgs(100n, 101n);
  });

  it('should update unlocking state when unlocking more stake', async () => {
    await stakingManager.addStake(100, owner);
    await stakingManager.unlockStake(40, owner);

    const key = await stakingManager.getKey(owner, owner);
    const unlockingOne = await stakingManager.unlockings(key);

    await stakingManager.unlockStake(40, owner);
    const unlockingTwo = await stakingManager.unlockings(key);

    expect(unlockingTwo.unlockAt).to.be.greaterThan(unlockingOne.unlockAt);
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
    expect(unlockingTwo.unlockAt).to.be.equal(unlockingOne.unlockAt);
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

  it('cannot withdraw stake with invalid arguments', async () => {
    await expect(
      stakingManager.withdrawStake(ethers.ZeroAddress),
    ).to.be.revertedWithCustomError(
      stakingManager,
      'StakeeCannotBeZeroAddress',
    );
  });

  it("should not be able to withdraw stake that hasn't unlocked", async () => {
    await stakingManager.addStake(100, owner);
    await stakingManager.unlockStake(100, owner);
    await expect(
      stakingManager.withdrawStake(owner),
    ).to.be.revertedWithCustomError(stakingManager, 'StakeNotYetUnlocked');
  });

  it('should be able to cancel unlocking', async () => {
    await stakingManager.addStake(100, owner);
    await stakingManager.unlockStake(100, owner);
    await stakingManager.cancelUnlocking(100, owner);

    const key = await stakingManager.getKey(owner, owner);
    const unlocking = await stakingManager.unlockings(key);

    assert.equal(unlocking.amount, 0n, 'Expected unlocking to be cancelled');
  });

  it('should be able to cancel a portion of the unlocking', async () => {
    await stakingManager.addStake(100, owner);
    await stakingManager.unlockStake(100, owner);
    await stakingManager.cancelUnlocking(54, owner);

    const key = await stakingManager.getKey(owner, owner);
    const unlocking = await stakingManager.unlockings(key);

    assert.equal(
      unlocking.amount,
      46n,
      'Expected only a portion of the unlocking to be cancelled',
    );
  });

  it('cannot cancel unlocking with invalid arguments', async () => {
    await expect(
      stakingManager.cancelUnlocking(100, ethers.ZeroAddress),
    ).to.be.revertedWithCustomError(
      stakingManager,
      'StakeeCannotBeZeroAddress',
    );

    await expect(
      stakingManager.cancelUnlocking(0, owner),
    ).to.be.revertedWithCustomError(
      stakingManager,
      'CannotCancelUnlockZeroAmount',
    );
  });

  it('unlocking more than exists clears entire stake', async () => {
    await stakingManager.addStake(100, owner);
    await stakingManager.unlockStake(100, owner);
    await stakingManager.cancelUnlocking(101, owner);

    const key = await stakingManager.getKey(owner, owner);
    const unlocking = await stakingManager.unlockings(key);

    assert.equal(unlocking.amount, 0n, 'Expected unlocking to be cancelled');
  });

  it('should allow delegated stake to exceed minimum owned stake by the stakee', async () => {
    await token.transfer(await accounts[1].getAddress(), 1000);
    await token
      .connect(accounts[1])
      .approve(await stakingManager.getAddress(), 1000);
    await stakingManager.connect(accounts[1]).addStake(180, owner);
  });

  it('should not allow directory to be joined with no stake', async () => {
    await directory.addManager(owner);
    await expect(
      directory.joinNextDirectory(owner, defaultSeekerId),
    ).to.be.revertedWithCustomError(directory, 'NoStakeToJoinEpoch');
  });

  it('cannot join directory with invalid arguments', async () => {
    await directory.addManager(owner);
    await expect(
      directory.joinNextDirectory(ethers.ZeroAddress, defaultSeekerId),
    ).to.be.revertedWithCustomError(directory, 'StakeeCannotBeZeroAddress');
  });

  it('can not join directory after unlocking all stake', async () => {
    await stakingManager.addStake(1, owner);
    await stakingManager.unlockStake(1, owner);

    await directory.addManager(owner);

    await expect(
      directory.joinNextDirectory(owner, defaultSeekerId),
    ).to.be.revertedWithCustomError(directory, 'NoStakeToJoinEpoch');
  });

  it('cannot check min stake proportion with invalid arguments', async () => {
    await expect(
      stakingManager.checkMinimumStakeProportion(ethers.ZeroAddress),
    ).to.be.revertedWithCustomError(
      stakingManager,
      'StakeeCannotBeZeroAddress',
    );
  });

  it('should reduce stake when joining directory with less than minimum stake', async () => {
    await stakingManager.addStake(100, owner);

    await token.transfer(await accounts[1].getAddress(), 1000);
    await token
      .connect(accounts[1])
      .approve(await stakingManager.getAddress(), 1000);
    await stakingManager.connect(accounts[1]).addStake(180, owner);

    // after unlocking, Node will own less than 20% of stake
    await stakingManager.unlockStake(80, owner);

    await directory.addManager(owner);
    await directory.joinNextDirectory(owner, defaultSeekerId);

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

    expect(managedStake / 2n).to.equal(joinedStake);
  });

  it('should fail to join when node`s own stake is 0', async () => {
    await token.transfer(await accounts[1].getAddress(), 1000);
    await token
      .connect(accounts[1])
      .approve(await stakingManager.getAddress(), 1000);
    await stakingManager.connect(accounts[1]).addStake(180, owner);

    await directory.addManager(owner);

    await expect(
      directory.joinNextDirectory(owner, defaultSeekerId),
    ).to.be.revertedWithCustomError(directory, 'NoJoiningStakeToJoinEpoch');
  });

  it('should be able to get total stake for a stakee', async () => {
    await stakingManager.addStake(100, owner);
    for (let i = 2; i < 10; i++) {
      await token.transfer(await accounts[i].getAddress(), 1000);
      await token
        .connect(accounts[i])
        .approve(await stakingManager.getAddress(), 1000);
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
    await directory.transferOwnership(await epochsManager.getAddress());
    await epochsManager.initializeEpoch({ from: owner });

    await stakingManager.addStake(100, owner);

    const stakeEntry = await stakingManager.getStakeEntry(owner, owner);

    assert.equal(
      stakeEntry.epochId,
      1n,
      'Stake entry should track the epoch id it was updated at',
    );
  });

  it('should not be able to join directory without stake', async () => {
    await setSeekeRegistry(accounts[0], accounts[1], 1);
    await expect(epochsManager.joinNextEpoch()).to.be.revertedWithCustomError(
      rewardsManager,
      'NoStakeToCreateRewardPool',
    );
  });

  it('should not be able to join directory without setting seeker account', async () => {
    await expect(epochsManager.joinNextEpoch()).to.be.revertedWithCustomError(
      epochsManager,
      'SeekerAccountCannotBeZeroAddress',
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

    await expect(epochsManager.joinNextEpoch()).to.be.revertedWithCustomError(
      epochsManager,
      'SeekerOwnerMismatch',
    );
  });

  it('should not be able to join the next epoch more than once', async () => {
    await stakingManager.addStake(1, owner);
    await directory.addManager(owner);
    await directory.joinNextDirectory(owner, defaultSeekerId);

    await expect(
      directory.joinNextDirectory(owner, defaultSeekerId),
    ).to.be.revertedWithCustomError(directory, 'StakeeAlreadyJoinedEpoch');
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
      ethers.ZeroAddress,
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
        stake,
        1n,
        'Expected to be able to query total stake for stakee',
      );
    }

    await directory.addManager(owner);
    await directory.setCurrentDirectory(epochId);

    const totalStake = await directory.getTotalStake(1);
    assert.equal(
      totalStake,
      BigInt(expectedTotalStake),
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
        boundary,
        BigInt(i) + 1n,
        'Expected entry to hold correct boundary value',
      );
    }
  });

  it('should be able to get total managed stake', async () => {
    let expectedTotalStake = 0;
    for (let i = 0; i < accounts.length; i++) {
      await token.transfer(await accounts[i].getAddress(), 100);
      await token
        .connect(accounts[i])
        .approve(await stakingManager.getAddress(), 100);
      await stakingManager
        .connect(accounts[i])
        .addStake(i + 1, await accounts[i].getAddress());
      expectedTotalStake += i + 1;
    }

    const totalManagedStake = await stakingManager.getTotalManagedStake();

    assert.equal(
      totalManagedStake,
      BigInt(expectedTotalStake),
      'Expected to be able to query for total managed stake',
    );
  });

  it('should correctly scan accounts based on their stake proportions', async () => {
    for (let i = 0; i < 5; i++) {
      await addStakeAndJoinEpoch(accounts[i], 1, i);
    }

    await directory.addManager(owner);
    await directory.setCurrentDirectory(epochId);

    const fifthPoint = (2n ** 128n - 1n) / 5n;
    const points = [
      0n,
      fifthPoint + 1n,
      fifthPoint * 2n + 2n,
      fifthPoint * 3n + 3n,
      fifthPoint * 4n + 4n,
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
    let point = (2n ** 128n - 1n) / 8n;
    await checkScanWithEpochId(
      await accounts[0].getAddress(),
      point.toString(),
      1,
    );

    // check point of node 1, epoch 1
    point = (2n ** 128n - 1n) / 2n;
    await checkScanWithEpochId(
      await accounts[1].getAddress(),
      point.toString(),
      1,
    );

    // check point of node 2, epoch 1
    point = 2n ** 128n - 1n;
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
    point = (2n ** 128n - 1n) / 4n;
    await checkScanWithEpochId(
      await accounts[1].getAddress(),
      point.toString(),
      2,
    );

    // check point of node 3, epoch 2
    point = ((2n ** 128n - 1n) / 4n) * 3n;
    await checkScanWithEpochId(
      await accounts[3].getAddress(),
      point.toString(),
      2,
    );

    // check epoch 3 - empty directory
    await checkScanWithEpochId(ethers.ZeroAddress, '10000000', 4);
  });

  it('can not call functions that onlyManager constraint', async () => {
    await expect(
      directory.joinNextDirectory(owner, defaultSeekerId),
    ).to.be.revertedWithCustomError(directory, 'OnlyManagers');
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
    await token
      .connect(accounts[1])
      .approve(await stakingManager.getAddress(), 100);
    await stakingManager
      .connect(accounts[1])
      .addStake(1, await accounts[1].getAddress());

    await token.transfer(await accounts[2].getAddress(), 100);
    await token
      .connect(accounts[2])
      .approve(await stakingManager.getAddress(), 100);
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

    assert.equal(address, ethers.ZeroAddress, 'Expected zero address');

    // check scan with epoch id
    const addressWithEpochId = await directory.scanWithEpochId(0, epochId);

    assert.equal(
      addressWithEpochId,
      address,
      'Expected address from scan with epoch id to be the same as address from scan',
    );
  });

  it('can validate contract interface', async () => {
    const TestSyloUtils = await ethers.getContractFactory('TestSyloUtils');
    const testSyloUtils = await TestSyloUtils.deploy();

    await expect(
      testSyloUtils.validateContractInterface(
        '',
        await rewardsManager.getAddress(),
        '0x3db12b5a',
      ),
    ).to.be.revertedWithCustomError(testSyloUtils, 'ContractNameCannotBeEmpty');

    await expect(
      testSyloUtils.validateContractInterface(
        'RewardsManager',
        ethers.ZeroAddress,
        '0x3db12b5a',
      ),
    ).to.be.revertedWithCustomError(
      testSyloUtils,
      'TargetContractCannotBeZeroAddress',
    );

    await expect(
      testSyloUtils.validateContractInterface(
        'RewardsManager',
        await rewardsManager.getAddress(),
        '0x00000000',
      ),
    ).to.be.revertedWithCustomError(
      testSyloUtils,
      'InterfaceIdCannotBeZeroBytes',
    );

    await expect(
      testSyloUtils.validateContractInterface(
        'RewardsManager',
        await rewardsManager.getAddress(),
        '0x11111111',
      ),
    ).to.be.revertedWithCustomError(testSyloUtils, 'TargetNotSupportInterface');
  });

  it('reverts if seeker power has not been registered', async () => {
    await expect(
      stakingManager.calculateCapacityFromSeekerPower(111),
    ).to.revertedWithCustomError(stakingManager, 'SeekerPowerNotRegistered');
  });

  it('correctly calculates seeker staking capacity from power', async () => {
    // default multiplier is 1M
    const seekerPowers = [
      { seekerId: 10, power: 100, expectedSyloCapacity: 100000000 },
      { seekerId: 11, power: 222, expectedSyloCapacity: 222000000 },
      { seekerId: 12, power: 432, expectedSyloCapacity: 432000000 },
      { seekerId: 13, power: 3, expectedSyloCapacity: 3000000 },
      { seekerId: 14, power: 4, expectedSyloCapacity: 4000000 },
      { seekerId: 15, power: 8, expectedSyloCapacity: 8000000 },
    ];

    for (const sp of seekerPowers) {
      await seekerPowerOracle.registerSeekerPowerRestricted(
        sp.seekerId,
        sp.power,
      );

      const capacity = await stakingManager.calculateCapacityFromSeekerPower(
        sp.seekerId,
      );

      const expectedSylo = ethers.parseEther(
        sp.expectedSyloCapacity.toString(),
      );

      expect(capacity).to.equal(expectedSylo);
    }
  });

  it('returns maximum SYLO amount if seeker power is very large', async () => {
    // 1 more than the maximum sylo
    await seekerPowerOracle.registerSeekerPowerRestricted(
      111,
      MAX_SYLO_STAKE + 1n,
    );

    const capacityOne = await stakingManager.calculateCapacityFromSeekerPower(
      111,
    );

    expect(capacityOne).to.equal(MAX_SYLO_STAKE);

    // seeker_power * multiplier > maximum_sylo
    await seekerPowerOracle.registerSeekerPowerRestricted(
      222,
      MAX_SYLO_STAKE / 2n,
    );

    const capacityTwo = await stakingManager.calculateCapacityFromSeekerPower(
      222,
    );

    expect(capacityTwo).to.equal(MAX_SYLO_STAKE);
  });

  it('can set seeker power multiplier', async () => {
    await seekerPowerOracle.registerSeekerPowerRestricted(111, 1);

    const originalCapacity =
      await stakingManager.calculateCapacityFromSeekerPower(111);

    expect(originalCapacity).to.equal(ethers.parseEther('1000000'));

    await expect(
      stakingManager.setSeekerPowerMultiplier(ethers.parseEther('500000')),
    )
      .to.emit(stakingManager, 'SeekerPowerMultiplierUpdated')
      .withArgs(ethers.parseEther('500000'));

    const newCapacity = await stakingManager.calculateCapacityFromSeekerPower(
      111,
    );

    expect(newCapacity).to.equal(ethers.parseEther('500000'));

    await expect(
      stakingManager.connect(accounts[1]).setSeekerPowerMultiplier(1),
    ).to.be.revertedWith('Ownable: caller is not the owner');
  });

  it('reverts when joining directory without seeker power registered', async () => {
    await stakingManager.addStake(100, owner);

    await directory.addManager(owner);
    await expect(
      directory.joinNextDirectory(owner, 111), // unregistered seeker
    ).to.be.revertedWithCustomError(stakingManager, 'SeekerPowerNotRegistered');
  });

  it('joins directory with stake where maximum is dependent on seeker power', async () => {
    const stakeToAdd = ethers.parseEther('10000000');

    await token.approve(stakingManager.getAddress(), stakeToAdd);
    await stakingManager.addStake(stakeToAdd, owner);

    // the added stake is 10,000,000 SYLO, but the seeker power capacity
    // is 4,000,000
    await seekerPowerOracle.registerSeekerPowerRestricted(111, 4);

    await directory.addManager(owner);
    await directory.joinNextDirectory(owner, 111);

    const joinedStake = await directory.getTotalStakeForStakee(1, owner);

    expect(joinedStake).to.equal(ethers.parseEther('4000000'));
  });

  it('joins directory with stake where maximum neither exceeds seeker power capacity or proportion capacity', async () => {
    const stakeToAdd = ethers.parseEther('100000');

    await token.approve(stakingManager.getAddress(), stakeToAdd);
    await stakingManager.addStake(stakeToAdd, owner);

    await seekerPowerOracle.registerSeekerPowerRestricted(111, 1);

    // delegated stake added causes the minimum proportion to be exceeded
    const delegatedStakeToAdd = ethers.parseEther('1000000');

    await token.transfer(accounts[2], delegatedStakeToAdd);
    await token
      .connect(accounts[2])
      .approve(stakingManager.getAddress(), delegatedStakeToAdd);
    await stakingManager
      .connect(accounts[2])
      .addStake(delegatedStakeToAdd, owner);

    await directory.addManager(owner);
    await directory.joinNextDirectory(owner, 111);

    const joinedStake = await directory.getTotalStakeForStakee(1, owner);

    // The seeker staking capacity is 1M, and the proportion capacity is 500k.
    // In this case the total stake exceeds both, and the joined stake should be
    // the lesser of the two capacities.

    expect(joinedStake).to.equal(ethers.parseEther('500000'));
  });

  async function setSeekeRegistry(
    account: Signer,
    seekerAccount: Signer,
    tokenId: number,
  ) {
    await utils.setSeekerRegistry(
      registries,
      seekers,
      seekerPowerOracle,
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
    await token
      .connect(account)
      .approve(await stakingManager.getAddress(), amount);
    await stakingManager
      .connect(account)
      .addStake(amount, await account.getAddress());
    await setSeekeRegistry(account, accounts[9], seekerId);
    await epochsManager.connect(account).joinNextEpoch();
    await seekerPowerOracle.registerSeekerPowerRestricted(seekerId, MaxUint256);
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

    const mnemonic = ethers.Mnemonic.fromPhrase(
      'search topple trouble similar sorry just around connect hello range predict ahead',
    );
    const keys = [];
    for (let i = 0; i < iterations; i++) {
      keys.push(
        ethers.HDNodeWallet.fromMnemonic(mnemonic, `m/44'/60'/0'/${i}`),
      );
    }

    let i = 0;

    outputCompletion();

    console.log('collecting scan results for', iterations, 'iterations...');

    while (i < iterations) {
      const hash = crypto.createHash('sha256');
      hash.update(keys[i].publicKey);
      hash.update(Buffer.from([0])); // append epoch
      const point = BigInt(
        '0x' + hash.digest().subarray(0, 16).toString('hex'),
      );
      const address = await directory.scan(point);
      updatePoint(address);
      i++;
    }

    return points;
  }
});
