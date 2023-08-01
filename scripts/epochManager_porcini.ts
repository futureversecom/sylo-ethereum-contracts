import * as factories from '../typechain-types';
import contractAddress from '../deployments/porcini_deployment_phase_two.json';
// import EpochsManagerContractABI from '../artifacts/contracts/epochs/EpochsManager.sol/EpochsManager.json'
import { ethers } from 'hardhat';

async function getActiveEpoch() {
  /*
   porcini testnet
  */
  const provider = new ethers.JsonRpcProvider('http://0.0.0.0:8545');

  const [deployer] = await ethers.getSigners();

  // const epochsManagerContract = new ethers.Contract(
  //   contractAddress.epochsManager,
  //   EpochsManagerContractABI.abi,
  // );

  const epochsManagerContract = factories.EpochsManager__factory.connect(
    contractAddress.epochsManager, // 0xd9f69a6dE82630558E468f219865bFe7247ba35E // 0x69e0A00Bd41F733BF6AE088692Fe916f94e9a9CF
    provider,
  );

  // const activeEpoch1 = await epochsManagerContract
  //   .connect(deployer)
  //   .getCurrentActiveEpoch();
  // console.log('getting current active epoch');
  // console.log('active epoch ', activeEpoch1[0].toString());
  // console.log('current block', await provider.getBlockNumber());
  // console.log('start block', activeEpoch1[1].startBlock);
  // console.log('end block', activeEpoch1[1].endBlock);
  // console.log('duration', activeEpoch1[1].duration);

  // const contracty = await ethers.getContractAt(
  //   'EpochsManager',
  //   '0x69e0A00Bd41F733BF6AE088692Fe916f94e9a9CF',
  // );
  // console.log('owner ', contracty.owner);
  // console.log('deployer ', deployer.address);

  // await epochsManagerContract
  //   .connect(deployer)
  //   .initializeEpoch({ gasLimit: 1_000_000 });
  // console.log('inited epoch');

  // const activeEpoch2 = await epochsManagerContract
  //   .connect(deployer)
  //   .getCurrentActiveEpoch();
  // console.log('getting current active epoch');
  // console.log('active epoch ', activeEpoch2[0].toString());
  //--------
  while (true) {
    await epochsManagerContract
      .connect(deployer)
      .initializeEpoch({ gasLimit: 1_000_000 });
    // .then(tx => tx.wait());
    console.log('inited epoch');
    const activeEpoch1 = await epochsManagerContract
      .connect(deployer)
      .getCurrentActiveEpoch();
    console.log('getting current active epoch');
    console.log('active epoch ', activeEpoch1[0].toString());
    console.log('current block', await provider.getBlockNumber());
    // await new Promise(resolve => setTimeout(resolve, 1500));
  }
}

getActiveEpoch();
