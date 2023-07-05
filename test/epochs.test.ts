import { ethers } from 'hardhat';
import { Signer } from 'ethers';
import { EpochsManager, SyloToken } from '../typechain-types';
import utils from './utils';
import { assert, expect } from 'chai';
import { SyloContracts } from '../common/contracts';

describe('Epochs', () => {
  let accounts: Signer[];
  let owner: string;

  let token: SyloToken;
  let epochsManager: EpochsManager;
  let contracts: SyloContracts;

  before(async () => {
    accounts = await ethers.getSigners();
    // first account is implicitly used as deployer of contracts in hardhat
    owner = await accounts[0].getAddress();

    const Token = await ethers.getContractFactory('SyloToken');
    token = await Token.deploy();
  });

  beforeEach(async () => {
    contracts = await utils.initializeContracts(owner, token);
    epochsManager = contracts.epochsManager;
    await contracts.directory.transferOwnership(
      await epochsManager.getAddress(),
    );
  });

  it('epoch manager cannot be intialized twice', async () => {
    await expect(
      epochsManager.initialize(
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        0,
      ),
    ).to.be.revertedWith('Initializable: contract is already initialized');
  });

  it('epoch manager cannot be intialized with invalid arguments', async () => {
    const EpochsManager = await ethers.getContractFactory('EpochsManager');
    epochsManager = await EpochsManager.deploy();

    await expect(
      epochsManager.initialize(
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        ethers.ZeroAddress,
        0,
      ),
    ).to.be.revertedWithCustomError(
      epochsManager,
      'RootSeekerCannotBeZeroAddress',
    );
  });

  it('can set epoch duration', async () => {
    await expect(epochsManager.setEpochDuration(777))
      .to.emit(epochsManager, 'EpochDurationUpdated')
      .withArgs(777);

    const epochDuration = await epochsManager.epochDuration();
    assert.equal(epochDuration, 777n, 'Expected epoch duration to be updated');
  });

  it('can not set epoch duration to zero', async () => {
    await expect(
      epochsManager.setEpochDuration(0),
    ).to.be.revertedWithCustomError(epochsManager, 'EpochDurationCannotBeZero');
  });

  it('not owner cannot set epoch duration', async () => {
    await expect(
      epochsManager.connect(accounts[1]).setEpochDuration(777),
    ).to.be.revertedWith('Ownable: caller is not the owner');
  });

  it('can initialize next epoch', async () => {
    await expect(epochsManager.initializeEpoch())
      .to.emit(epochsManager, 'NewEpoch')
      .withArgs(1);

    let currentIteration = await epochsManager.currentIteration();
    assert.equal(
      currentIteration,
      1n,
      'Expected fist epoch id to be correctly set',
    );

    await utils.advanceBlock(31);

    await epochsManager.initializeEpoch();
    currentIteration = await epochsManager.currentIteration();
    assert.equal(
      currentIteration,
      2n,
      'Expected second epoch id to be correctly set',
    );
  });

  it('can not initialize next epoch before current one had ended', async () => {
    await epochsManager.initializeEpoch();
    await expect(epochsManager.initializeEpoch())
      .to.be.revertedWithCustomError(epochsManager, 'EpochHasNotEnded')
      .withArgs(1);
  });

  it('correctly updates the epoch parameters every epoch', async () => {
    await epochsManager.initializeEpoch();
    await utils.advanceBlock(31);

    // change a couple of the parameters
    await contracts.ticketingParameters.setFaceValue(2222);
    await contracts.ticketingParameters.setDecayRate(1111);

    await epochsManager.initializeEpoch();

    const epochInfo = await epochsManager.getCurrentActiveEpoch();

    assert.equal(epochInfo[0], 2n, 'Expected epoch id to be 2');
    assert.equal(
      epochInfo[1].faceValue,
      2222n,
      'Expected face value to change',
    );

    assert.equal(
      epochInfo[1].decayRate,
      1111n,
      'Expected decay rate to change',
    );
  });
});
