import { ethers, network } from 'hardhat';
import { SyloContracts } from '../common/contracts';
import { expect, assert } from 'chai';
import { deployContracts } from './utils';
import { ProtocolTimeManager } from '../typechain-types';
import { getInterfaceId } from './utils';

describe('Protocol time manager', () => {
  let contracts: SyloContracts;
  let protocolTimeManager: ProtocolTimeManager;

  beforeEach(async () => {
    contracts = await deployContracts();
    protocolTimeManager = contracts.protocolTimeManager;
  });

  it('cannot initialize protocol time manager with empty cycle duration', async () => {
    const factory = await ethers.getContractFactory('ProtocolTimeManager');
    const protocolTimeManager = await factory.deploy();

    await expect(
      protocolTimeManager.initialize(0, 0),
    ).to.be.revertedWithCustomError(
      protocolTimeManager,
      'CannotInitializeWithZeroCycleDuration',
    );
  });

  it('cannot initialize protocol time manager with empty period duration', async () => {
    const factory = await ethers.getContractFactory('ProtocolTimeManager');
    const protocolTimeManager = await factory.deploy();

    await expect(
      protocolTimeManager.initialize(1000, 0),
    ).to.be.revertedWithCustomError(
      protocolTimeManager,
      'CannotInitializeWithZeroPeriodDuration',
    );
  });

  it('cannot initialize protocol time manager more than once', async () => {
    await expect(protocolTimeManager.initialize(1000, 1000)).to.be.revertedWith(
      'Initializable: contract is already initialized',
    );
  });

  it('cannot set protocol start with zero value', async () => {
    await expect(
      protocolTimeManager.setProtocolStart(0),
    ).to.be.revertedWithCustomError(
      protocolTimeManager,
      'CannotSetProtocolStartWithZeroStart',
    );
  });

  it('cannot set protocol start with future start', async () => {
    const x = await ethers.provider.getBlock('latest');
    if (!x) {
      expect.fail('timestamp undefeind');
    }

    await expect(
      protocolTimeManager.setProtocolStart(x.timestamp + 1000),
    ).to.be.revertedWithCustomError(
      protocolTimeManager,
      'CannotSetStartInFuture',
    );
  });

  it('can set protocol start', async () => {
    const x = await ethers.provider.getBlock('latest');
    if (!x) {
      expect.fail('timestamp undefeind');
    }

    const start = await protocolTimeManager.getStart();
    assert.equal(Number(start), x.timestamp);

    await network.provider.send('evm_increaseTime', [3000]);
    await network.provider.send('evm_mine');

    await protocolTimeManager.setProtocolStart(x.timestamp + 1000);

    const newStart = await protocolTimeManager.getStart();
    assert.equal(Number(newStart), x.timestamp + 1000);
  });

  it('cannot set zero cycle duration', async () => {
    await expect(
      protocolTimeManager.setCycleDuration(0),
    ).to.be.revertedWithCustomError(
      protocolTimeManager,
      'CannotSetZeroCycleDuration',
    );
  });

  it('can set cycle duration', async () => {
    const cycleDuration = await protocolTimeManager.getCycleDuration();
    assert.equal(Number(cycleDuration), 1000);

    await protocolTimeManager.setCycleDuration(2000);

    const cycleDurationTwo = await protocolTimeManager.getCycleDuration();
    assert.equal(Number(cycleDurationTwo), 2000);
  });

  it('cannot set zero period duration', async () => {
    await expect(
      protocolTimeManager.setPeriodDuration(0),
    ).to.be.revertedWithCustomError(
      protocolTimeManager,
      'CannotSetZeroPeriodDuration',
    );
  });

  it('can set period duration', async () => {
    const periodDuration = await protocolTimeManager.getPeriodDuration();
    assert.equal(Number(periodDuration), 1000);

    await protocolTimeManager.setPeriodDuration(2000);

    const periodDurationTwo = await protocolTimeManager.getPeriodDuration();
    assert.equal(Number(periodDurationTwo), 2000);
  });

  it('can get both current cycle and period', async () => {
    await network.provider.send('evm_increaseTime', [3000]);
    await network.provider.send('evm_mine');

    const times = await protocolTimeManager.timeNow();
    assert.equal(Number(times[0]), 4);
    assert.equal(Number(times[1]), 4);

    await protocolTimeManager.setCycleDuration(3000);
    await protocolTimeManager.setPeriodDuration(2000);

    await network.provider.send('evm_increaseTime', [6000]);
    await network.provider.send('evm_mine');

    const timesTwo = await protocolTimeManager.timeNow();
    assert.equal(Number(timesTwo[0]), 6);
    assert.equal(Number(timesTwo[1]), 7);
  });

  it('returns one for first cycle', async () => {
    await network.provider.send('evm_increaseTime', [500]);
    await network.provider.send('evm_mine');
    const currentCycle = await protocolTimeManager.getCurrentCycle();
    assert.equal(Number(currentCycle), 1);
  });

  it('can get current cycle without updated duration', async () => {
    await network.provider.send('evm_increaseTime', [3500]);
    await network.provider.send('evm_mine');

    const currentCycle = await protocolTimeManager.getCurrentCycle();
    assert.equal(Number(currentCycle), 4);
  });

  it('can get current cycle with one updated duration', async () => {
    await network.provider.send('evm_increaseTime', [3000]);
    await network.provider.send('evm_mine');

    const currentCycle = await protocolTimeManager.getCurrentCycle();
    assert.equal(Number(currentCycle), 4);

    await protocolTimeManager.setCycleDuration(200);

    await network.provider.send('evm_increaseTime', [7000]);
    await network.provider.send('evm_mine');

    const currentCycleTwo = await protocolTimeManager.getCurrentCycle();
    assert.equal(Number(currentCycleTwo), 39);
  });

  it('can get current cycle with multiple updated durations', async () => {
    await network.provider.send('evm_increaseTime', [3000]);
    await network.provider.send('evm_mine');

    const currentCycle = await protocolTimeManager.getCurrentCycle();
    assert.equal(Number(currentCycle), 4);

    await protocolTimeManager.setCycleDuration(200);

    await network.provider.send('evm_increaseTime', [7000]);
    await network.provider.send('evm_mine');

    const currentCycleTwo = await protocolTimeManager.getCurrentCycle();
    assert.equal(Number(currentCycleTwo), 39);

    await protocolTimeManager.setCycleDuration(125);

    await network.provider.send('evm_increaseTime', [5800]);
    await network.provider.send('evm_mine');

    const currentCycleThree = await protocolTimeManager.getCurrentCycle();
    assert.equal(Number(currentCycleThree), 85);

    await protocolTimeManager.setCycleDuration(2);

    await network.provider.send('evm_increaseTime', [200]);
    await network.provider.send('evm_mine');

    const currentCycleFour = await protocolTimeManager.getCurrentCycle();
    assert.equal(Number(currentCycleFour), 185);
  });

  it('can get current period without updated duration', async () => {
    await network.provider.send('evm_increaseTime', [3000]);
    await network.provider.send('evm_mine');

    const currentPeriod = await protocolTimeManager.getCurrentPeriod();
    assert.equal(Number(currentPeriod), 4);
  });

  it('can get current period with one updated duration', async () => {
    await network.provider.send('evm_increaseTime', [3000]);
    await network.provider.send('evm_mine');

    const currentPeriod = await protocolTimeManager.getCurrentPeriod();
    assert.equal(Number(currentPeriod), 4);

    await protocolTimeManager.setPeriodDuration(200);

    await network.provider.send('evm_increaseTime', [7000]);
    await network.provider.send('evm_mine');

    const currentPeriodTwo = await protocolTimeManager.getCurrentPeriod();
    assert.equal(Number(currentPeriodTwo), 39);
  });

  it('can get current period with multiple updated durations', async () => {
    await network.provider.send('evm_increaseTime', [3000]);
    await network.provider.send('evm_mine');

    const currentPeriod = await protocolTimeManager.getCurrentPeriod();
    assert.equal(Number(currentPeriod), 4);

    await protocolTimeManager.setPeriodDuration(200);

    await network.provider.send('evm_increaseTime', [7000]);
    await network.provider.send('evm_mine');

    const currentPeriodTwo = await protocolTimeManager.getCurrentPeriod();
    assert.equal(Number(currentPeriodTwo), 39);

    await protocolTimeManager.setPeriodDuration(125);

    await network.provider.send('evm_increaseTime', [5800]);
    await network.provider.send('evm_mine');

    const currentPeriodThree = await protocolTimeManager.getCurrentPeriod();
    assert.equal(Number(currentPeriodThree), 85);

    await protocolTimeManager.setPeriodDuration(2);

    await network.provider.send('evm_increaseTime', [200]);
    await network.provider.send('evm_mine');

    const currentPeriodFour = await protocolTimeManager.getCurrentPeriod();
    assert.equal(Number(currentPeriodFour), 185);
  });

  it('supports only protocol time manager interface', async () => {
    const abi = [
      'function setProtocolStart(uint256 _start) external',
      'function setCycleDuration(uint256 duration) external',
      'function setPeriodDuration(uint256 duration) external',
      'function getCycleDuration() external returns (uint256)',
      'function getPeriodDuration() external returns (uint256)',
      'function timeNow() external returns (uint256, uint256)',
      'function getCurrentCycle() external returns (uint256)',
      'function getCurrentPeriod() external returns (uint256)',
      'function getStart() external view returns (uint256)',
    ];

    const interfaceId = getInterfaceId(abi);

    const supports = await protocolTimeManager.supportsInterface(interfaceId);

    assert.equal(
      supports,
      true,
      'Expected protocol time manager to support correct interface',
    );

    const invalidAbi = ['function foo(uint256 duration) external'];

    const invalidAbiInterfaceId = getInterfaceId(invalidAbi);

    const invalid = await protocolTimeManager.supportsInterface(
      invalidAbiInterfaceId,
    );

    assert.equal(
      invalid,
      false,
      'Expected protocol time manager to not support incorrect interface',
    );
  });
});
