import { ethers } from 'hardhat';
import { MockOracle__factory, Seekers__factory } from '../typechain';

async function seekersUpdateMockOracle() {
  const [deployer] = await ethers.getSigners();

  const mockOracle = '0xc9F30AB40D0A8E0EdfC5E1eD22F822A0bcCaADe7';

  const seeker = Seekers__factory.connect(
    '0xec0Fc52F22eAdE86b6E2f93dDa9f5fE7f11Cf4f6',
    deployer,
  );

  const tx = await seeker.connect(deployer).setOracle(mockOracle);
  await tx.wait(2);
}

async function seekerOwnerCheck(seekerId: number) {
  const [deployer] = await ethers.getSigners();

  const seeker = Seekers__factory.connect(
    '0xec0Fc52F22eAdE86b6E2f93dDa9f5fE7f11Cf4f6',
    deployer,
  );

  console.log('Seeker Owner', seekerId, await seeker.ownerOf(seekerId));
}

async function requestSeekerOwner(seekerId: number) {
  const [deployer] = await ethers.getSigners();

  const seeker = Seekers__factory.connect(
    '0xec0Fc52F22eAdE86b6E2f93dDa9f5fE7f11Cf4f6',
    deployer,
  );

  const tx = await seeker.requestVerification(seekerId);
  await tx.wait(2);

  console.log('Successfully Requested');
}

async function invokeCallback() {
  const [deployer] = await ethers.getSigners();

  const mockOracle = MockOracle__factory.connect(
    '0xc9F30AB40D0A8E0EdfC5E1eD22F822A0bcCaADe7',
    deployer,
  );

  const tx = await mockOracle.invokeCallback();
  await tx.wait(2);

  console.log('Successfully Invoked');
}

async function setSeekerOwner(seekerId: number, seekerAddress: string) {
  const [deployer] = await ethers.getSigners();

  const mockOracle = MockOracle__factory.connect(
    '0xc9F30AB40D0A8E0EdfC5E1eD22F822A0bcCaADe7',
    deployer,
  );

  const tx = await mockOracle.setOwner(seekerId, seekerAddress);
  await tx.wait(2);

  console.log('Successfully Set Seeker Owner');
}

async function main() {
  // step 1: mock 5 seekers ownership - only need to call once
  // for (let i = 0; i < 5; i++) {
  //   await setSeekerOwner(i, '0xA5A5A6e97528a6BA1EE04f27582d37E9b612f6C3');
  // }

  // step 2: request verification for seeker 0 - should be called from the node
  // await requestSeekerOwner(0);

  // step 3: invoke callback from mock oracle
  await invokeCallback();

  // step 4: check seeker owner from seekers contract
  await seekerOwnerCheck(0);
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error);
    process.exit(1);
  });
