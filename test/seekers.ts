import { ethers, network } from 'hardhat';
import { Signer } from 'ethers';
import { MockOracle, Seekers } from '../typechain';
import { assert, expect } from 'chai';
import utils from './utils';

describe('Listing', () => {
  let accounts: Signer[];
  let owner: string;

  let seekers: Seekers;
  let mockOracle: MockOracle;

  before(async () => {
    accounts = await ethers.getSigners();
    // first account is implicitly used as deployer of contracts in hardhat
    owner = await accounts[0].getAddress();
  });

  beforeEach(async () => {
    const Token = await ethers.getContractFactory('SyloToken');
    const token = await Token.deploy();

    const contracts = await utils.initializeContracts(owner, token.address, {
      payoutPercentage: 5000,
    });
    seekers = contracts.seekers;
    mockOracle = contracts.mockOracle;
  });

  it('can set parameters after deployment', async () => {
    await seekers.setSeekers('0x0000000000000000000000000000000000000001');
    expect(await seekers.seekers()).to.equal(
      '0x0000000000000000000000000000000000000001',
    );

    await seekers.setOracle('0x0000000000000000000000000000000000000001');
    expect(await seekers.oracle()).to.equal(
      '0x0000000000000000000000000000000000000001',
    );

    await seekers.setValidDuration(666);
    expect(await seekers.validDuration()).to.equal(666);

    await seekers.setCallbackGasLimit(666);
    expect(await seekers.callbackGasLimit()).to.equal(666);

    await seekers.setCallbackBounty(666);
    expect(await seekers.callbackBounty()).to.equal(666);
  });

  it('can make a request to the oracle and retrieve seeker owner', async () => {
    // set a mock value first
    await mockOracle.setOwner(1, owner);

    await seekers.requestVerification(1, 1);

    await mockOracle.invokeCallback();

    const seekerOwner = await seekers.ownerOf(1);

    assert.equal(seekerOwner, owner);
  });

  it('ownership check returns correct owner', async () => {
    const seekerOwner = await seekers.ownerOf(1);

    assert.equal(seekerOwner, '0x0000000000000000000000000000000000000000');
  });

  it('ownership checks correctly expire', async () => {
    // set a mock value first
    await mockOracle.setOwner(1, owner);

    await seekers.requestVerification(1, 1);

    await mockOracle.invokeCallback();

    await utils.advanceBlock(101);

    const seekerOwner = await seekers.ownerOf(1);

    assert.equal(seekerOwner, '0x0000000000000000000000000000000000000000');
  });
});
