import { ethers as hardhatEthers } from 'hardhat';
import { TestSeekers__factory } from '../typechain';

export async function mintSeeker(
  seekerAddress: string,
  seekerOnwer: string,
  amount: number,
): Promise<void> {
  const [deployer] = await hardhatEthers.getSigners();

  const seeker = TestSeekers__factory.connect(seekerAddress, deployer);

  for (let i = 0; i < amount; i++) {
    const tx = await seeker.mint(seekerOnwer, i);
    await tx.wait();
  }
}
