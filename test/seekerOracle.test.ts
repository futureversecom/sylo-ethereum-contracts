import { ethers } from 'hardhat';
import { Signer } from 'ethers';
import { SyloToken } from '../typechain-types';
import utils from './utils';
import { SyloContracts } from '../common/contracts';
import { expect } from 'chai';

describe('Seeker Power Oracle', () => {
  let accounts: Signer[];
  let deployer: string;

  let token: SyloToken;
  let contracts: SyloContracts;

  before(async () => {
    accounts = await ethers.getSigners();
    deployer = await accounts[0].getAddress();

    const Token = await ethers.getContractFactory('SyloToken');
    token = await Token.deploy();
  });

  beforeEach(async () => {
    contracts = await utils.initializeContracts(deployer, token);
  });

  it('can set oracle with owner', async () => {
    const oracle = await accounts[1].getAddress();
    await contracts.seekerPowerOracle.setOracle(oracle);

    const _oracle = await contracts.seekerPowerOracle.oracle();

    expect(_oracle).to.equal(oracle);
  });

  it('can set seeker power with oracle', async () => {
    const oracle = await accounts[1].getAddress();
    await contracts.seekerPowerOracle.setOracle(oracle);

    const seekerId = 5;

    // check power of seeker id 5
    const zeroPower = await contracts.seekerPowerOracle.seekerPowers(seekerId);
    expect(zeroPower).to.equal(0);

    const seekerPower = 111;

    // update with oracle
    await contracts.seekerPowerOracle
      .connect(accounts[1])
      .setSeekerPowerRestricted(seekerId, seekerPower);

    const updatedPower = await contracts.seekerPowerOracle.seekerPowers(
      seekerId,
    );
    expect(updatedPower).to.equal(seekerPower);
  });

  it('can set seeker power with owner', async () => {
    const seekerId = 5;

    // check power of seeker id 5
    const zeroPower = await contracts.seekerPowerOracle.seekerPowers(seekerId);
    expect(zeroPower).to.equal(0);

    const seekerPower = 111;

    // update with oracle
    await contracts.seekerPowerOracle.setSeekerPowerRestricted(
      seekerId,
      seekerPower,
    );

    const updatedPower = await contracts.seekerPowerOracle.seekerPowers(
      seekerId,
    );
    expect(updatedPower).to.equal(seekerPower);
  });

  it('only allows setSeekerPowerRestricted to be called by owner or oracle', async () => {
    await expect(
      contracts.seekerPowerOracle
        .connect(accounts[2]) // unauthorized caller
        .setSeekerPowerRestricted(1, 2),
    ).to.be.revertedWithCustomError(
      contracts.seekerPowerOracle,
      'UnauthorizedSetSeekerCall',
    );
  });

  it('can set seeker power with proof', async () => {
    const oracle = await accounts[1].getAddress();
    await contracts.seekerPowerOracle.setOracle(oracle);

    const seekerId = 111;
    const seekerPower = 222;

    const proofMessage = await contracts.seekerPowerOracle.getProofMessage(
      seekerId,
      seekerPower,
    );

    const proof = await accounts[1].signMessage(
      Buffer.from(proofMessage.slice(2), 'hex'),
    );

    await contracts.seekerPowerOracle
      .connect(accounts[2])
      .registerSeekerPower(seekerId, seekerPower, proof);

    const updatedPower = await contracts.seekerPowerOracle.seekerPowers(
      seekerId,
    );
    expect(updatedPower).to.equal(seekerPower);
  });

  it('reverts when setting seeker power with invalid proof', async () => {
    const oracle = await accounts[1].getAddress();
    await contracts.seekerPowerOracle.setOracle(oracle);

    const seekerId = 111;
    const seekerPower = 222;

    const proofMessage = await contracts.seekerPowerOracle.getProofMessage(
      seekerId,
      seekerPower,
    );

    // sign with non-oracle account
    const proof = await accounts[2].signMessage(
      Buffer.from(proofMessage.slice(2), 'hex'),
    );

    await expect(
      contracts.seekerPowerOracle
        .connect(accounts[3])
        .registerSeekerPower(seekerId, seekerPower, proof),
    ).to.be.revertedWithCustomError(
      contracts.seekerPowerOracle,
      'UnauthorizedSetSeekerCall',
    );
  });

  it('can update multiple seeker powers', async () => {
    for (let i = 1; i < 11; i++) {
      const seekerId = i;
      const seekerPower = i * 1111;

      const proofMessage = await contracts.seekerPowerOracle.getProofMessage(
        seekerId,
        seekerPower,
      );

      const proof = await accounts[0].signMessage(
        Buffer.from(proofMessage.slice(2), 'hex'),
      );

      await contracts.seekerPowerOracle
        .connect(accounts[i])
        .registerSeekerPower(seekerId, seekerPower, proof);

      const updatedPower = await contracts.seekerPowerOracle.seekerPowers(
        seekerId,
      );
      expect(updatedPower).to.equal(seekerPower);
    }
  });

  it('can update the same seeker power multiple times', async () => {
    for (let i = 0; i < 5; i++) {
      const seekerId = 1;
      const seekerPower = i * 1111;

      const proofMessage = await contracts.seekerPowerOracle.getProofMessage(
        seekerId,
        seekerPower,
      );

      const proof = await accounts[0].signMessage(
        Buffer.from(proofMessage.slice(2), 'hex'),
      );

      await contracts.seekerPowerOracle
        .connect(accounts[1])
        .registerSeekerPower(seekerId, seekerPower, proof);

      const updatedPower = await contracts.seekerPowerOracle.seekerPowers(
        seekerId,
      );
      expect(updatedPower).to.equal(seekerPower);
    }
  });
});
