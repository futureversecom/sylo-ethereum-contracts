import { ethers } from 'hardhat';
import { Signer } from 'ethers';
import { EpochsManager, SyloToken } from '../typechain';
import utils, { Contracts } from './utils';
import { assert, expect } from 'chai';

describe('Epochs', () => {
  let accounts: Signer[];
  let owner: string;

  let token: SyloToken;
  let epochsManager: EpochsManager;
  let contracts: Contracts;

  before(async () => {
    accounts = await ethers.getSigners();
    // first account is implicitly used as deployer of contracts in hardhat
    owner = await accounts[0].getAddress();

    const Token = await ethers.getContractFactory('SyloToken');
    token = await Token.deploy();
  });

  beforeEach(async () => {
    contracts = await utils.initializeContracts(owner, token.address);
    epochsManager = contracts.epochsManager;
    await contracts.directory.transferOwnership(epochsManager.address);
  });

  it('can set epoch duration', async () => {
    await expect(epochsManager.setEpochDuration(778))
      .to.emit(epochsManager, 'EpochDurationUpdated')
      .withArgs(778);

    await epochsManager.setEpochDuration(777);
    const epochDuration = await epochsManager.epochDuration();
    assert.equal(
      epochDuration.toNumber(),
      777,
      'Expected epoch duration to be updated',
    );
  });

  it('can initialize next epoch', async () => {
    await expect(epochsManager.initializeEpoch())
      .to.emit(epochsManager, 'NewEpoch')
      .withArgs(1);

    let currentIteration = await epochsManager.currentIteration();
    assert.equal(
      currentIteration.toNumber(),
      1,
      'Expected fist epoch id to be correctly set',
    );

    await utils.advanceBlock(31);

    await epochsManager.initializeEpoch();
    currentIteration = await epochsManager.currentIteration();
    assert.equal(
      currentIteration.toNumber(),
      2,
      'Expected second epoch id to be correctly set',
    );
  });

  it('can not initialize next epoch before current one had ended', async () => {
    await epochsManager.initializeEpoch();
    await expect(epochsManager.initializeEpoch()).to.be.revertedWith(
      'Current epoch has not yet ended',
    );
  });

  it('correctly updates the epoch parameters every epoch', async () => {
    await epochsManager.initializeEpoch();
    await utils.advanceBlock(31);

    // change a couple of the parameters
    await contracts.ticketingParameters.setFaceValue(2222);
    await contracts.ticketingParameters.setDecayRate(1111);

    await epochsManager.initializeEpoch();

    const epoch = await epochsManager.getCurrentActiveEpoch();

    assert.equal(
      epoch.faceValue.toNumber(),
      2222,
      'Expected face value to change',
    );

    assert.equal(epoch.decayRate, 1111, 'Expected decay rate to change');
  });
});
