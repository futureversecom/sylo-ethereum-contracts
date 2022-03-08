import { ethers, upgrades, network } from 'hardhat';
import * as fs from 'fs/promises';

async function main() {
  const [deployer] = await ethers.getSigners();
  console.log(`Beginning upgrade of epochs manager to ${network.name} by deployer: ${deployer.address}`);

  const deployedAddresses = await fs.readFile(`${__dirname}/${network.name}_deployment_phase_two.json`)
    .then(b => JSON.parse(b.toString()));

  const EpochsManager = await ethers.getContractFactory("EpochsManager");
  const upgradedEpochsManager = await upgrades.upgradeProxy(
    deployedAddresses.epochsManager,
    EpochsManager
  );

  console.log(`Successfully upgraded epochs manager. Address: ${upgradedEpochsManager.address}`);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });