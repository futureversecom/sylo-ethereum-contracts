import { ethers } from 'hardhat';
import { SeekerStatsOracle } from '../../typechain-types';
import { SyloContracts } from '../../common/contracts';
import { Signer } from 'ethers';
import { expect, assert } from 'chai';
import { deployContracts } from '../utils';

class Seeker {
  constructor(
    public seekerId: number,
    public rank: number,
    public attr_reactor: number,
    public attr_cores: number,
    public attr_durability: number,
    public attr_sensors: number,
    public attr_storage: number,
    public attr_chip: number,
  ) {}
}

describe.only('Seeker Stats', () => {
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
      'OracleCannotBeZeroAddress',
    );
  });

  it('cannot initialize seeker stats oracle more than once', async () => {
    await expect(
      seekerStatsOracle.initialize(await accounts[19].getAddress()),
    ).to.be.revertedWith('Initializable: contract is already initialized');
  });

  it('can set oracle account as owner', async () => {
    assert.equal(
      await seekerStatsOracle.SeekerStatsOracleAccount(),
      await accounts[19].getAddress(),
    );
    await seekerStatsOracle.setOracle(await accounts[18].getAddress());
    assert.equal(
      await seekerStatsOracle.SeekerStatsOracleAccount(),
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
      'OracleCannotBeZeroAddress',
    );
  });

  it('can register seeker', async () => {
    const seeker = new Seeker(10, 2, 10, 20, 30, 40, 50, 60);

    const proofMessage = await seekerStatsOracle.createStatsMessage(seeker);
    const signature = await accounts[19].signMessage(
      Buffer.from(proofMessage.slice(2), 'hex'),
    );
    await seekerStatsOracle.registerSeeker(seeker, signature);
  });

  it('can register seeker restricted', async () => {
    const seeker = new Seeker(10, 2, 10, 20, 30, 40, 50, 60);

    await seekerStatsOracle
      .connect(accounts[19])
      .registerSeekerRestricted(seeker);
  });

  it('cannot register seeker restricted from non-oracle account', async () => {
    const seeker = new Seeker(10, 2, 10, 20, 30, 40, 50, 60);

    await expect(
      seekerStatsOracle.registerSeekerRestricted(seeker),
    ).to.be.revertedWithCustomError(
      seekerStatsOracle,
      'UnauthorizedRegisterSeekerStatsCall',
    );
  });

  it('cannot register seeker from non-oracle account', async () => {
    const seeker = new Seeker(20, 2, 10, 20, 30, 40, 50, 60);

    const proofMessage = await seekerStatsOracle.createStatsMessage(seeker);
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

    const proofMessage = await seekerStatsOracle.createStatsMessage(seeker);
    const proofMessageTwo = await seekerStatsOracle.createStatsMessage(
      seekerTwo,
    );

    const signature = await accounts[19].signMessage(
      Buffer.from(proofMessage.slice(2), 'hex'),
    );
    const signatureTwo = await accounts[19].signMessage(
      Buffer.from(proofMessageTwo.slice(2), 'hex'),
    );

    await seekerStatsOracle.registerSeeker(seeker, signature);

    const fetchedSeeker = await seekerStatsOracle.seekers(10);
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

    const fetchedSeekerTwo = await seekerStatsOracle.seekers(10);
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
    const seekerList = await createAndRegisterSeeker(15);

    const attributeConverageExpected = calculateAttributesCoverage(seekerList);
    const attributeCoverage =
      await seekerStatsOracle.calculateAttributeCoverage(seekerList);
    const formatedCoverage = ethers.formatEther(attributeCoverage);

    assert.equal(
      Number(formatedCoverage).toFixed(0),
      attributeConverageExpected.toFixed(0),
    );
  });

  function calculateAttributesCoverage(seekers: Seeker[]): number {
    const angleRadians = Math.sin((2 * Math.PI) / 6 + 2 * Math.PI);

    let totalCoverage = 0;

    for (const seeker of seekers) {
      const {
        attr_reactor,
        attr_cores,
        attr_durability,
        attr_sensors,
        attr_storage,
        attr_chip,
      } = seeker;

      totalCoverage += (attr_reactor * angleRadians * attr_cores) / 2;
      totalCoverage += (attr_cores * angleRadians * attr_durability) / 2;
      totalCoverage += (attr_durability * angleRadians * attr_sensors) / 2;
      totalCoverage += (attr_sensors * angleRadians * attr_storage) / 2;
      totalCoverage += (attr_storage * angleRadians * attr_chip) / 2;
      totalCoverage += (attr_chip * angleRadians * attr_reactor) / 2;
    }

    return totalCoverage;
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
      const proofMessage = await seekerStatsOracle.createStatsMessage(
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
      seeker1.attr_reactor === seeker2.attr_reactor &&
      seeker1.attr_cores === seeker2.attr_cores &&
      seeker1.attr_durability === seeker2.attr_durability &&
      seeker1.attr_sensors === seeker2.attr_sensors &&
      seeker1.attr_storage === seeker2.attr_storage &&
      seeker1.attr_chip === seeker2.attr_chip
    );
  }
});
