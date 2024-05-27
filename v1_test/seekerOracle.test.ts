import { ethers } from 'hardhat';
import { Signer } from 'ethers';
import { SyloToken } from '../typechain-types';
import utils from './utils';
import { SyloContracts } from '../common/contracts';
import { expect } from 'chai';
import { randomBytes } from 'crypto';

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

  it('seeker power oracle cannot be initialized with invalid arguments', async () => {
    const SeekerPowerOracle = await ethers.getContractFactory(
      'SeekerPowerOracle',
    );
    const seekerPowerOracle = await SeekerPowerOracle.deploy();

    await expect(
      seekerPowerOracle.initialize(ethers.ZeroAddress),
    ).to.be.revertedWithCustomError(
      seekerPowerOracle,
      'OracleCannotBeZeroAddress',
    );
  });

  it('seeker power oracle cannot be initialized twice', async () => {
    await expect(
      contracts.seekerPowerOracle.initialize(deployer),
    ).to.be.revertedWith('Initializable: contract is already initialized');
  });

  it('can set oracle with owner', async () => {
    const oracle = await accounts[1].getAddress();
    await contracts.seekerPowerOracle.setOracle(oracle);

    const _oracle = await contracts.seekerPowerOracle.oracle();

    expect(_oracle).to.equal(oracle);
  });

  it('reverts with setting oracle with non-owner', async () => {
    await expect(
      contracts.seekerPowerOracle.connect(accounts[3]).setOracle(deployer),
    ).to.be.revertedWith('Ownable: caller is not the owner');
  });

  it('can set seeker power with oracle', async () => {
    const oracle = await accounts[1].getAddress();
    await contracts.seekerPowerOracle.setOracle(oracle);

    const seekerId = 5;

    // check power of seeker id 5
    const zeroPower = await contracts.seekerPowerOracle.getSeekerPower(
      seekerId,
    );
    expect(zeroPower).to.equal(0);

    const seekerPower = 111;

    // update with oracle
    await contracts.seekerPowerOracle
      .connect(accounts[1])
      .registerSeekerPowerRestricted(seekerId, seekerPower);

    const updatedPower = await contracts.seekerPowerOracle.getSeekerPower(
      seekerId,
    );
    expect(updatedPower).to.equal(seekerPower);
  });

  it('can set seeker power with owner', async () => {
    const seekerId = 5;

    // check power of seeker id 5
    const zeroPower = await contracts.seekerPowerOracle.getSeekerPower(
      seekerId,
    );
    expect(zeroPower).to.equal(0);

    const seekerPower = 111;

    // update with oracle
    await contracts.seekerPowerOracle.registerSeekerPowerRestricted(
      seekerId,
      seekerPower,
    );

    const updatedPower = await contracts.seekerPowerOracle.getSeekerPower(
      seekerId,
    );
    expect(updatedPower).to.equal(seekerPower);
  });

  it('only allows registerSeekerPowerRestricted to be called by the oracle', async () => {
    await expect(
      contracts.seekerPowerOracle
        .connect(accounts[2]) // unauthorized caller
        .registerSeekerPowerRestricted(1, 2),
    ).to.be.revertedWithCustomError(
      contracts.seekerPowerOracle,
      'UnauthorizedRegisterSeekerPowerCall',
    );
  });

  it('can set seeker power with proof', async () => {
    const oracle = await accounts[1].getAddress();
    await contracts.seekerPowerOracle.setOracle(oracle);

    const seekerId = 111;
    const seekerPower = 222;
    const nonce = randomBytes(32);

    const proofMessage = await contracts.seekerPowerOracle.getProofMessage(
      seekerId,
      seekerPower,
      nonce,
    );

    const proof = await accounts[1].signMessage(
      Buffer.from(proofMessage.slice(2), 'hex'),
    );

    await contracts.seekerPowerOracle
      .connect(accounts[2])
      .registerSeekerPower(seekerId, seekerPower, nonce, proof);

    const updatedPower = await contracts.seekerPowerOracle.getSeekerPower(
      seekerId,
    );
    expect(updatedPower).to.equal(seekerPower);
  });

  it('reverts when setting seeker power with invalid proof', async () => {
    const oracle = await accounts[1].getAddress();
    await contracts.seekerPowerOracle.setOracle(oracle);

    const seekerId = 111;
    const seekerPower = 222;
    const nonce = randomBytes(32);

    const proofMessage = await contracts.seekerPowerOracle.getProofMessage(
      seekerId,
      seekerPower,
      nonce,
    );

    // sign with non-oracle account
    const invalidOracleProof = await accounts[2].signMessage(
      Buffer.from(proofMessage.slice(2), 'hex'),
    );

    await expect(
      contracts.seekerPowerOracle.registerSeekerPower(
        seekerId,
        seekerPower,
        nonce,
        invalidOracleProof,
      ),
    ).to.be.revertedWithCustomError(
      contracts.seekerPowerOracle,
      'UnauthorizedRegisterSeekerPowerCall',
    );

    const validProof = await accounts[1].signMessage(
      Buffer.from(proofMessage.slice(2), 'hex'),
    );

    await contracts.seekerPowerOracle.registerSeekerPower(
      seekerId,
      seekerPower,
      nonce,
      validProof,
    );

    await expect(
      contracts.seekerPowerOracle.registerSeekerPower(
        seekerId,
        seekerPower,
        nonce,
        validProof,
      ),
    ).to.be.revertedWithCustomError(
      contracts.seekerPowerOracle,
      'NonceCannotBeReused',
    );
  });

  it('can not set seeker power to 0', async () => {
    const seekerId = 111;

    await expect(
      contracts.seekerPowerOracle.registerSeekerPowerRestricted(seekerId, 0),
    ).to.be.revertedWithCustomError(
      contracts.seekerPowerOracle,
      'PowerCannotBeZero',
    );

    const seekerPower = 0;
    const nonce = randomBytes(32);

    const proofMessage = await contracts.seekerPowerOracle.getProofMessage(
      seekerId,
      seekerPower,
      nonce,
    );

    const proof = await accounts[1].signMessage(
      Buffer.from(proofMessage.slice(2), 'hex'),
    );

    await expect(
      contracts.seekerPowerOracle
        .connect(accounts[2])
        .registerSeekerPower(seekerId, seekerPower, nonce, proof),
    ).to.be.revertedWithCustomError(
      contracts.seekerPowerOracle,
      'PowerCannotBeZero',
    );
  });

  it('can update multiple seeker powers', async () => {
    for (let i = 1; i < 11; i++) {
      const seekerId = i;
      const seekerPower = i * 1111;
      const nonce = randomBytes(32);

      const proofMessage = await contracts.seekerPowerOracle.getProofMessage(
        seekerId,
        seekerPower,
        nonce,
      );

      const proof = await accounts[0].signMessage(
        Buffer.from(proofMessage.slice(2), 'hex'),
      );

      await contracts.seekerPowerOracle
        .connect(accounts[i])
        .registerSeekerPower(seekerId, seekerPower, nonce, proof);

      const updatedPower = await contracts.seekerPowerOracle.getSeekerPower(
        seekerId,
      );
      expect(updatedPower).to.equal(seekerPower);
    }
  });

  it('can update the same seeker power multiple times', async () => {
    for (let i = 1; i < 6; i++) {
      const seekerId = 1;
      const seekerPower = i * 1111;
      const nonce = randomBytes(32);

      const proofMessage = await contracts.seekerPowerOracle.getProofMessage(
        seekerId,
        seekerPower,
        nonce,
      );

      const proof = await accounts[0].signMessage(
        Buffer.from(proofMessage.slice(2), 'hex'),
      );

      await contracts.seekerPowerOracle
        .connect(accounts[1])
        .registerSeekerPower(seekerId, seekerPower, nonce, proof);

      const updatedPower = await contracts.seekerPowerOracle.getSeekerPower(
        seekerId,
      );
      expect(updatedPower).to.equal(seekerPower);
    }
  });
});
