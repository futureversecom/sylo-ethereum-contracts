import { ethers } from 'hardhat';
import { SeekerStatsOracle } from '../../typechain-types';
import { SyloContracts } from '../../common/contracts';
import { Signer } from 'ethers';
import { expect, assert } from 'chai';
import { deployContracts, getInterfaceId } from '../utils';

class Seeker {
  constructor(
    public seekerId: number,
    public rank: number,
    public attrReactor: number,
    public attrCores: number,
    public attrDurability: number,
    public attrSensors: number,
    public attrStorage: number,
    public attrChip: number,
  ) {}
}

describe('Seeker Stats', () => {
  let accounts: Signer[];
  let contracts: SyloContracts;
  let seekerStatsOracle: SeekerStatsOracle;

  beforeEach(async () => {
    accounts = await ethers.getSigners();
    contracts = await deployContracts();
    seekerStatsOracle = contracts.seekerStatsOracle;
    await seekerStatsOracle.setOracle(await accounts[19].getAddress());
  });

  it('cannot initialize seeker stats oracle with invalid arguemnts', async () => {
    const factory = await ethers.getContractFactory('SeekerStatsOracle');
    const seekerStatsOracle = await factory.deploy();

    await expect(
      seekerStatsOracle.initialize(ethers.ZeroAddress),
    ).to.be.revertedWithCustomError(
      seekerStatsOracle,
      'OracleAddressCannotBeNil',
    );
  });

  it('cannot initialize seeker stats oracle more than once', async () => {
    await expect(
      seekerStatsOracle.initialize(await accounts[19].getAddress()),
    ).to.be.revertedWith('Initializable: contract is already initialized');
  });

  it('can set oracle account as owner', async () => {
    assert.equal(
      await seekerStatsOracle.oracle(),
      await accounts[19].getAddress(),
    );

    const oracle = await accounts[18].getAddress();
    await expect(seekerStatsOracle.setOracle(oracle))
      .to.emit(seekerStatsOracle, 'OracleUpdated')
      .withArgs(oracle);
    assert.equal(
      await seekerStatsOracle.oracle(),
      await accounts[18].getAddress(),
    );
  });

  it('cannot set oracle account as non-owner', async () => {
    await expect(
      seekerStatsOracle
        .connect(accounts[19])
        .setOracle(await accounts[18].getAddress()),
    ).to.be.revertedWith('Ownable: caller is not the owner');
  });

  it('cannot set oracle account as zero address', async () => {
    await expect(
      seekerStatsOracle.setOracle(ethers.ZeroAddress),
    ).to.be.revertedWithCustomError(
      seekerStatsOracle,
      'OracleAddressCannotBeNil',
    );
  });

  it('can register seeker', async () => {
    const seeker = new Seeker(10, 2, 10, 20, 30, 40, 50, 60);

    const proofMessage = await seekerStatsOracle.createProofMessage(seeker);
    const signature = await accounts[19].signMessage(
      Buffer.from(proofMessage.slice(2), 'hex'),
    );

    await expect(seekerStatsOracle.registerSeeker(seeker, signature))
      .to.emit(seekerStatsOracle, 'SeekerStatsUpdated')
      .withArgs(10n, 10n, 20n, 30n, 40n, 50n, 60n);
  });

  it('can register seeker restricted', async () => {
    const seeker = new Seeker(10, 2, 10, 20, 30, 40, 50, 60);

    await expect(
      seekerStatsOracle.connect(accounts[19]).registerSeekerRestricted(seeker),
    )
      .to.emit(seekerStatsOracle, 'SeekerStatsUpdated')
      .withArgs(10n, 10n, 20n, 30n, 40n, 50n, 60n);
  });

  it('cannot register seeker restricted from non-oracle account', async () => {
    const seeker = new Seeker(10, 2, 10, 20, 30, 40, 50, 60);

    await expect(
      seekerStatsOracle.registerSeekerRestricted(seeker),
    ).to.be.revertedWithCustomError(
      seekerStatsOracle,
      'UnauthorizedRegisterSeekerStats',
    );
  });

  it('cannot register seeker from non-oracle account', async () => {
    const seeker = new Seeker(20, 2, 10, 20, 30, 40, 50, 60);

    const proofMessage = await seekerStatsOracle.createProofMessage(seeker);
    const signature = await accounts[18].signMessage(
      Buffer.from(proofMessage.slice(2), 'hex'),
    );
    await expect(
      seekerStatsOracle.registerSeeker(seeker, signature),
    ).to.be.revertedWithCustomError(
      seekerStatsOracle,
      'InvalidSignatureForSeekerProof',
    );
  });

  it('can update registered seeker', async () => {
    const seeker = new Seeker(10, 2, 1, 1, 1, 1, 1, 1);
    const seekerTwo = new Seeker(10, 2, 10, 10, 10, 10, 10, 10);

    const proofMessage = await seekerStatsOracle.createProofMessage(seeker);
    const proofMessageTwo = await seekerStatsOracle.createProofMessage(
      seekerTwo,
    );

    const signature = await accounts[19].signMessage(
      Buffer.from(proofMessage.slice(2), 'hex'),
    );
    const signatureTwo = await accounts[19].signMessage(
      Buffer.from(proofMessageTwo.slice(2), 'hex'),
    );

    await seekerStatsOracle.registerSeeker(seeker, signature);

    const fetchedSeeker = await seekerStatsOracle.seekerStats(10);
    const newSeekerOne = new Seeker(
      Number(fetchedSeeker[0]),
      Number(fetchedSeeker[1]),
      Number(fetchedSeeker[2]),
      Number(fetchedSeeker[3]),
      Number(fetchedSeeker[4]),
      Number(fetchedSeeker[5]),
      Number(fetchedSeeker[6]),
      Number(fetchedSeeker[7]),
    );

    assert.equal(Number(fetchedSeeker.seekerId), 10);

    await seekerStatsOracle.registerSeeker(seekerTwo, signatureTwo);

    const fetchedSeekerTwo = await seekerStatsOracle.seekerStats(10);
    const newSeekerTwo = new Seeker(
      Number(fetchedSeekerTwo[0]),
      Number(fetchedSeekerTwo[1]),
      Number(fetchedSeekerTwo[2]),
      Number(fetchedSeekerTwo[3]),
      Number(fetchedSeekerTwo[4]),
      Number(fetchedSeekerTwo[5]),
      Number(fetchedSeekerTwo[6]),
      Number(fetchedSeekerTwo[7]),
    );

    assert.equal(compareSeekers(newSeekerOne, newSeekerTwo), false);
  });

  it('cannot register seeker from with invalid proof', async () => {
    const seeker = new Seeker(20, 2, 10, 20, 30, 40, 50, 60);

    const proofMessage = 'invalid message';
    const signature = await accounts[18].signMessage(
      Buffer.from(proofMessage.slice(2), 'hex'),
    );
    await expect(
      seekerStatsOracle.registerSeeker(seeker, signature),
    ).to.be.revertedWithCustomError(
      seekerStatsOracle,
      'InvalidSignatureForSeekerProof',
    );
  });

  it('cannot calculate converage with unregistered seeker', async () => {
    const seeker = new Seeker(30, 2, 10, 20, 30, 40, 50, 60);

    await expect(seekerStatsOracle.calculateAttributeCoverage([seeker]))
      .to.be.revertedWithCustomError(seekerStatsOracle, 'SeekerNotRegistered')
      .withArgs(seeker.seekerId);
  });

  it('can calculate converage with registered seeker', async () => {
    const seekerList = await createAndRegisterSeeker(1);

    const attributeConverageExpected = calculateAttributesCoverage(seekerList);
    const attributeCoverage =
      await seekerStatsOracle.calculateAttributeCoverage(seekerList);
    const formatedCoverage = ethers.formatEther(attributeCoverage);

    assert.equal(
      Number(formatedCoverage).toFixed(1),
      attributeConverageExpected.toFixed(1),
    );
  });

  it('can calculate converage with multiple registered seeker', async () => {
    const seekerList = await createAndRegisterSeeker(5);

    const attributeConverageExpected = calculateAttributesCoverage(seekerList);
    const attributeCoverage =
      await seekerStatsOracle.calculateAttributeCoverage(seekerList);
    const formatedCoverage = ethers.formatEther(attributeCoverage);

    assert.equal(
      Number(formatedCoverage).toFixed(0),
      attributeConverageExpected.toFixed(0),
    );
  });

  it('supports only seeker stats oracle interface', async () => {
    const abi = [
      'function setOracle(address _seekerStatsOracleAccount) external',
      'function createProofMessage((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) calldata seeker) external pure returns (bytes memory)',
      'function registerSeekerRestricted((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) calldata seeker) external',
      'function registerSeeker((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) calldata seeker, bytes calldata signature) external',
      'function calculateAttributeCoverage((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] calldata seekersList) external view returns (int256)',
    ];

    const interfaceId = getInterfaceId(abi);

    const supports = await seekerStatsOracle.supportsInterface(interfaceId);

    assert.equal(
      supports,
      true,
      'Expected seeker stats oracle to support correct interface',
    );

    const invalidAbi = ['function foo(uint256 duration) external'];

    const invalidAbiInterfaceId = getInterfaceId(invalidAbi);

    const invalid = await seekerStatsOracle.supportsInterface(
      invalidAbiInterfaceId,
    );

    assert.equal(
      invalid,
      false,
      'Expected seeker stats oracle to not support incorrect interface',
    );
  });

  function calculateAttributesCoverage(seekers: Seeker[]): number {
    const angleRadians = Math.sin((2 * Math.PI) / 6 + 2 * Math.PI);

    let coverage = 0;

    let totalReactor = 0;
    let totalCores = 0;
    let totalDurability = 0;
    let totalSensors = 0;
    let totalStorage = 0;
    let totalChip = 0;

    for (const seeker of seekers) {
      const {
        attrReactor,
        attrCores,
        attrDurability,
        attrSensors,
        attrStorage,
        attrChip,
      } = seeker;

      totalReactor += attrReactor;
      totalCores += attrCores;
      totalDurability += attrDurability;
      totalSensors += attrSensors;
      totalStorage += attrStorage;
      totalChip += attrChip;
    }

    coverage += (totalReactor * angleRadians * totalCores) / 2;
    coverage += (totalCores * angleRadians * totalDurability) / 2;
    coverage += (totalDurability * angleRadians * totalSensors) / 2;
    coverage += (totalSensors * angleRadians * totalStorage) / 2;
    coverage += (totalStorage * angleRadians * totalChip) / 2;
    coverage += (totalChip * angleRadians * totalReactor) / 2;

    return coverage;
  }

  async function createAndRegisterSeeker(amount: number): Promise<Seeker[]> {
    const seekerList: Seeker[] = [];
    for (let i = 0; i < amount; i++) {
      const newSeeker = new Seeker(
        i,
        i,
        i + 10,
        i + 20,
        i + 30,
        i + 40,
        i + 50,
        i + 60,
      );
      const proofMessage = await seekerStatsOracle.createProofMessage(
        newSeeker,
      );
      const signature = await accounts[19].signMessage(
        Buffer.from(proofMessage.slice(2), 'hex'),
      );
      seekerList.push(newSeeker);
      await seekerStatsOracle.registerSeeker(newSeeker, signature);
    }

    return seekerList;
  }

  function compareSeekers(seeker1: Seeker, seeker2: Seeker): boolean {
    return (
      seeker1.seekerId === seeker2.seekerId &&
      seeker1.rank === seeker2.rank &&
      seeker1.attrReactor === seeker2.attrReactor &&
      seeker1.attrCores === seeker2.attrCores &&
      seeker1.attrDurability === seeker2.attrDurability &&
      seeker1.attrSensors === seeker2.attrSensors &&
      seeker1.attrStorage === seeker2.attrStorage &&
      seeker1.attrChip === seeker2.attrChip
    );
  }
});
