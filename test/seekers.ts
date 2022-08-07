import { ethers } from 'hardhat';
import { Signer } from 'ethers';
import { MockOracle, Seekers } from '../typechain';
import { assert, expect } from 'chai';
import utils from './utils';

describe('Seekers', () => {
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
    expect(await seekers._seekers()).to.equal(
      '0x0000000000000000000000000000000000000001',
    );

    await seekers.setOracle('0x0000000000000000000000000000000000000001');
    expect(await seekers._oracle()).to.equal(
      '0x0000000000000000000000000000000000000001',
    );

    await seekers.setToken('0x0000000000000000000000000000000000000001');
    expect(await seekers._oracle()).to.equal(
      '0x0000000000000000000000000000000000000001',
    );

    await seekers.setValidDuration(666);
    expect(await seekers.validDuration()).to.equal(666);

    await seekers.setCallbackGasLimit(666);
    expect(await seekers.callbackGasLimit()).to.equal(666);

    await seekers.setCallbackBounty(666);
    expect(await seekers.callbackBounty()).to.equal(666);
  });

  it('can cannot invoke call back from mock oracle - oracle address is incorrect', async () => {
    // set a mock value first
    await mockOracle.setOwner(1, owner);

    await seekers.requestVerification(1);

    await seekers.setOracle('0x0000000000000000000000000000000000000002');

    await expect(mockOracle.invokeCallback()).to.be.revertedWith(
      'Callback to sender failed',
    );
  });

  it('can make a request to the oracle and retrieve seeker owner', async () => {
    // set a mock value first
    await mockOracle.setOwner(1, owner);

    await seekers.requestVerification(1);

    await mockOracle.invokeCallback();

    const seekerOwner = await seekers.ownerOf(1);

    assert.equal(seekerOwner, owner);
  });

  it('ownership check returns correct owner', async () => {
    const seekerOwner = await seekers.ownerOf(1);

    assert.equal(seekerOwner, ethers.constants.AddressZero);
  });

  it('ownership checks correctly expire', async () => {
    // set a mock value first
    await mockOracle.setOwner(1, owner);

    await seekers.requestVerification(1);

    await mockOracle.invokeCallback();

    await utils.advanceBlock(101);

    const seekerOwner = await seekers.ownerOf(1);

    assert.equal(seekerOwner, ethers.constants.AddressZero);
  });

  it('can call request verification with fee swap', async () => {
    // set a mock value first
    await mockOracle.setOwner(1, owner);

    await seekers.requestVerificationWithFeeSwap(1, 1);

    await mockOracle.invokeCallback();

    const seekerOwner = await seekers.ownerOf(1);

    assert.equal(seekerOwner, owner);
  });

  it('can call withdraw functions', async () => {
    await seekers.withdrawAll(owner);
    await seekers.withdrawAllFee(owner);
  });

  it('can cannot request verification to mock oracle - mock oracle fail', async () => {
    const MockOracleFail = await ethers.getContractFactory('MockOracleFail');
    const mockOracleFail = await MockOracleFail.deploy();

    await seekers.setOracle(mockOracleFail.address);

    await expect(seekers.requestVerification(1)).to.be.revertedWith(
      'Oracle request failed',
    );
  });

  it('can cannot request verification with fee swap to mock oracle - mock oracle fail', async () => {
    const MockOracleFail = await ethers.getContractFactory('MockOracleFail');
    const mockOracleFail = await MockOracleFail.deploy();

    await seekers.setOracle(mockOracleFail.address);

    await expect(
      seekers.requestVerificationWithFeeSwap(1, 1),
    ).to.be.revertedWith('Oracle request failed');
  });
});
