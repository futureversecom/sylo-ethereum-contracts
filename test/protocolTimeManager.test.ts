import { ethers, network } from 'hardhat';
import { SyloContracts } from '../common/contracts';
import { expect, assert } from 'chai';
import { deployContracts } from './utils';
import { ProtocolTimeManager } from '../typechain-types';
import { getInterfaceId } from './utils';

describe.only('Protocol time manager', () => {
  let contracts: SyloContracts;
  let protocolTimeManager: ProtocolTimeManager;

  beforeEach(async () => {
    contracts = await deployContracts();
    protocolTimeManager = contracts.protocolTimeManager;
  });

  it('cannot initialize protocol time manager with empty cycle duration', async () => {
    const factory = await ethers.getContractFactory('ProtocolTimeManager');
    const protocolTimeManagerTemp = await factory.deploy();

    await expect(
      protocolTimeManagerTemp.initialize(0, 0),
    ).to.be.revertedWithCustomError(
      protocolTimeManagerTemp,
      'CannotInitializeWithZeroCycleDuration',
    );
  });

  it('cannot initialize protocol time manager with empty period duration', async () => {
    const factory = await ethers.getContractFactory('ProtocolTimeManager');
    const protocolTimeManagerTemp = await factory.deploy();

    await expect(
      protocolTimeManagerTemp.initialize(1000, 0),
    ).to.be.revertedWithCustomError(
      protocolTimeManagerTemp,
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

  it('can set protocol start', async () => {
    const latestBlock = await ethers.provider.getBlock('latest');
    if (!latestBlock) {
      expect.fail('timestamp undefeind');
    }

    const start = await protocolTimeManager.getStart();
    assert.equal(Number(start), 0);

    await protocolTimeManager.setProtocolStart(1000);

    const newStart = await protocolTimeManager.getStart();
    assert.equal(Number(newStart), latestBlock.timestamp + 1001);
  });

  it('cannot set zero cycle duration', async () => {
    await expect(
      protocolTimeManager.setCycleDuration(0),
    ).to.be.revertedWithCustomError(
      protocolTimeManager,
      'CannotSetZeroCycleDuration',
    );
  });

  it('cannot set duplicate cycle duration', async () => {
    await expect(
      protocolTimeManager.setCycleDuration(1000),
    ).to.be.revertedWithCustomError(
      protocolTimeManager,
      'CannotSetDuplicateCycleDuration',
    );
  });

  it('can set cycle duration', async () => {
    await protocolTimeManager.setProtocolStart(100);

    await increaseTime(200);

    await protocolTimeManager.setCycleDuration(2000);

    await checkCycle(1);
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
    await protocolTimeManager.setProtocolStart(100);

    await increaseTime(200);

    await checkCycle(1);

    await protocolTimeManager.setPeriodDuration(2000);

    await increaseTime(1000);

    await checkPeriod(2);
  });

  it('Cannot get cycle without protocol start', async () => {
    await expect(protocolTimeManager.getCurrentCycle()).revertedWithCustomError(
      protocolTimeManager,
      'ProtocolHasNotBegun',
    );
  });

  it('Cannot get period without protocol start', async () => {
    await expect(
      protocolTimeManager.getCurrentPeriod(),
    ).revertedWithCustomError(protocolTimeManager, 'ProtocolHasNotBegun');
  });

  it('returns one for first cycle', async () => {
    await protocolTimeManager.setProtocolStart(100);

    await increaseTime(500);

    await checkCycle(1);
  });

  it('can get current cycle without updated duration', async () => {
    await protocolTimeManager.setProtocolStart(100);

    await increaseTime(3500);

    await checkCycle(4);
  });

  it('can get current cycle with one updated duration', async () => {
    await protocolTimeManager.setProtocolStart(100);

    // Timeline     -> 1000
    // Timeline P   -> 900
    // 100 (1) -> 1100
    await increaseTime(1000);
    await checkCycle(1);

    // Timeline     -> 1500
    // Timeline P   -> 1400
    // 1100 (2) -> 2100
    await increaseTime(500);
    console.log('first check');
    await checkCycle(2);

    await protocolTimeManager.setCycleDuration(200);

    console.log('second check');
    await checkCycle(2);

    // Timeline     -> 2500
    // Timeline P   -> 2400
    // 2100 (3) -> 2300
    // 2300 (4) -> 2500
    // 2500 (5) -> 2700
    await increaseTime(1000);
    await checkCycle(4);
  });

  it('can get current cycle with multiple updated durations', async () => {
    await protocolTimeManager.setProtocolStart(100);

    await increaseTime(150);

    await checkCycle(1);

    await increaseTime(2850);

    await checkCycle(3);

    await protocolTimeManager.setCycleDuration(200);

    await increaseTime(1000);

    await checkCycle(8);

    await protocolTimeManager.setCycleDuration(125);

    await increaseTime(2000);

    await checkCycle(24);

    await protocolTimeManager.setCycleDuration(2);

    await increaseTime(200);

    await checkCycle(125);
  });

  it('cycle test one', async () => {
    // Protocol start -> 800
    // Timeline       -> 0
    await protocolTimeManager.setProtocolStart(800);

    // Timeline       -> 500
    // Timeline P     -> 0
    await increaseTime(500);

    await protocolTimeManager.setCycleDuration(150);

    // Timeline       -> 500
    // Timeline P     -> 0

    await expect(protocolTimeManager.getCurrentCycle()).revertedWithCustomError(
      protocolTimeManager,
      'ProtocolHasNotBegun',
    );

    // Timeline       -> 1000
    // Timeline P     -> 200
    // 800 (1) -> 950 (2) -> 1100
    await increaseTime(500);
    await checkCycle(2);
    await checkCycle(2);

    // Timeline       -> 1500
    // Timeline P     -> 700
    // 1100 (3) -> 1250 (4) -> 1400 (5) -> 1550
    await increaseTime(500);
    await checkCycle(5);

    // CycleDuration  -> 50
    // Timeline       -> 1500
    // Timeline P     -> 700
    await protocolTimeManager.setCycleDuration(50);

    // Timeline       -> 1550
    // Timeline P     -> 750
    // 1550 (6) -> 1600
    await increaseTime(50);
    await checkCycle(6);

    // Timeline       -> 1600
    // Timeline P     -> 800
    // 1600 (7) -> 1650
    await increaseTime(50);
    await checkCycle(7);

    // Timeline       -> 1650
    // Timeline P     -> 850
    // 1650 (8) -> 1700
    await increaseTime(50);
    await checkCycle(8);

    // Timeline       -> 1700
    // Timeline P     -> 900
    // 1700 (9) -> 1750
    await increaseTime(50);
    console.log('time since start ehrehrehe');
    await checkCycle(9);

    // CycleDuration  -> 850
    // Timeline       -> 1700
    // Timeline P     -> 900
    await protocolTimeManager.setCycleDuration(850);

    // Timeline       -> 1750
    // Timeline P     -> 950
    // 1750 (10) -> 2600)
    await increaseTime(50);
    console.log('time since start ehrehrehe two');
    await checkCycle(10);

    // Timeline       -> 4200
    // Timeline P     -> 3400
    // 2600 (11) -> 3450 (12) -> 4300
    await increaseTime(2450);
    await checkCycle(12);

    await protocolTimeManager.setCycleDuration(75);

    // Timeline       -> 4300
    // Timeline P     -> 3500
    // 4300 (13) -> 4375
    await increaseTime(100);
    console.log('chekcing 13z');
    await checkCycle(13);

    // Timeline       -> 5750
    // Timeline P     -> 4950
    // 4375 (14) -> 4450 (15) -> 4525 (16) -> 4600 (17)
    // -> 4675 (18) -> 4750 (19) -> 4825 (20) -> 4900 (21)
    // -> 4975 (22)
    await increaseTime(1450);
    await checkCycle(22);
  });

  it('returns one for first period', async () => {
    await protocolTimeManager.setProtocolStart(100);

    await increaseTime(500);

    await checkPeriod(1);
  });

  it('can get current period without updated duration', async () => {
    await protocolTimeManager.setProtocolStart(100);

    await increaseTime(3500);

    await checkPeriod(4);
  });

  it('can get current period with one updated duration', async () => {
    await protocolTimeManager.setProtocolStart(100);

    // Timeline     -> 1000
    // Timeline P   -> 900
    // 100 (1) -> 1100
    await increaseTime(1000);
    await checkPeriod(1);

    // Timeline     -> 1500
    // Timeline P   -> 1400
    // 1100 (2) -> 2100
    await increaseTime(500);
    await checkPeriod(2);

    await protocolTimeManager.setPeriodDuration(200);

    // Timeline     -> 2600
    // Timeline P   -> 2500
    // 2100 (3) -> 2200
    // 2200 (4) -> 2400
    // 2400 (5) -> 2600
    await increaseTime(1000);
    await checkPeriod(5);
  });

  it('can get current period with multiple updated durations', async () => {
    await protocolTimeManager.setProtocolStart(100);

    await increaseTime(150);

    await checkPeriod(1);

    await increaseTime(2850);

    await checkPeriod(3);

    await protocolTimeManager.setPeriodDuration(200);

    await increaseTime(1000);

    await checkPeriod(8);

    await protocolTimeManager.setPeriodDuration(125);

    await increaseTime(2000);

    await checkPeriod(24);

    await protocolTimeManager.setPeriodDuration(2);

    await increaseTime(200);

    await checkPeriod(125);
  });

  it('period test one', async () => {
    // Protocol start -> 800
    // Timeline       -> 0
    await protocolTimeManager.setProtocolStart(800);

    // Timeline       -> 500
    // Timeline P     -> 0
    await increaseTime(500);

    await protocolTimeManager.setPeriodDuration(150);

    // Timeline       -> 500
    // Timeline P     -> 0

    await expect(
      protocolTimeManager.getCurrentPeriod(),
    ).revertedWithCustomError(protocolTimeManager, 'ProtocolHasNotBegun');

    // Timeline       -> 1000
    // Timeline P     -> 200
    // 800 (1) -> 950 (2) -> 1100
    await increaseTime(500);
    await checkPeriod(2);
    await checkPeriod(2);

    // Timeline       -> 1500
    // Timeline P     -> 700
    // 1100 (3) -> 1250 (4) -> 1400 (5) -> 1550
    await increaseTime(500);
    await checkPeriod(5);

    // CycleDuration  -> 50
    // Timeline       -> 1500
    // Timeline P     -> 700
    await protocolTimeManager.setPeriodDuration(50);

    // Timeline       -> 1550
    // Timeline P     -> 750
    // 1550 (6) -> 1600
    await increaseTime(50);
    await checkPeriod(6);

    // Timeline       -> 1600
    // Timeline P     -> 800
    // 1600 (7) -> 1650
    await increaseTime(50);
    await checkPeriod(7);

    // Timeline       -> 1650
    // Timeline P     -> 850
    // 1650 (8) -> 1700
    await increaseTime(50);
    await checkPeriod(8);

    // Timeline       -> 1700
    // Timeline P     -> 900
    // 1700 (9) -> 1750
    await increaseTime(50);
    await checkPeriod(9);

    // CycleDuration  -> 850
    // Timeline       -> 1700
    // Timeline P     -> 900
    await protocolTimeManager.setPeriodDuration(850);

    // Timeline       -> 1750
    // Timeline P     -> 950
    // 1750 (10) -> 2600)
    await increaseTime(50);
    await checkPeriod(10);

    // Timeline       -> 4200
    // Timeline P     -> 3400
    // 2600 (11) -> 3450 (12) -> 4300
    await increaseTime(2450);
    await checkPeriod(12);

    await protocolTimeManager.setPeriodDuration(75);

    // Timeline       -> 4300
    // Timeline P     -> 3500
    // 4300 (13) -> 4375
    await increaseTime(100);
    await checkPeriod(13);

    // Timeline       -> 5750
    // Timeline P     -> 4950
    // 4375 (14) -> 4450 (15) -> 4525 (16) -> 4600 (17)
    // -> 4675 (18) -> 4750 (19) -> 4825 (20) -> 4900 (21)
    // -> 4975 (22)
    await increaseTime(1450);
    await checkPeriod(22);
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

  async function increaseTime(time: number) {
    await network.provider.send('evm_increaseTime', [time]);
    await network.provider.send('evm_mine');
  }

  async function checkCycle(cycle: number) {
    const currentCycle = await protocolTimeManager.getCurrentCycle();
    assert.equal(Number(currentCycle), cycle);
  }

  async function checkPeriod(period: number) {
    const currentPeriod = await protocolTimeManager.getCurrentPeriod();
    assert.equal(Number(currentPeriod), period);
  }
});
