import * as etherHRE from 'hardhat';
import { ethers } from 'ethers';
import * as factories from '../typechain-types';
import EpochsManagerContractABI from '../artifacts/contracts/epochs/EpochsManager.sol/EpochsManager.json';
// import TicketingParametersABI from '../artifacts/contracts/payments/ticketing/TicketingParameters.sol/TicketingParameters.json';
import contractAddress from '../deploy/localhost_deployment_phase_two.json';

const WINNING_PROBABILITY = ethers.BigNumber.from(2).pow(128).sub(1);

async function getActiveEpoch() {
  // const provider = new ethers.providers.JsonRpcProvider(
  //   'https://porcini.rootnet.app/', // https://porcini.au.rootnet.app // https://porcini.rootnet.app/
  // );
  /*
   hardhat / ganache testnet
  */
  const provider = new ethers.providers.JsonRpcProvider(
    'http://127.0.0.1:8545/',
  );

  const deployer = connectSigner(
    new ethers.Wallet(
      '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80',
    ),
    provider,
  );

  // const epochsManagerContract = new ethers.Contract(
  //   contractAddress.epochsManager,
  //   EpochsManagerContractABI.abi,
  //   wallet,
  // );

  const epochsManagerContract = factories.EpochsManager__factory.connect(
    contractAddress.epochsManager, // 0xd9f69a6dE82630558E468f219865bFe7247ba35E // 0x69e0A00Bd41F733BF6AE088692Fe916f94e9a9CF
    provider,
  );

  // for (let i = 0; i < 10; i++) {
  //   // const activeEpoch = await epochsManagerContract
  //   //   .connect(wallet)
  //   //   .getCurrentActiveEpoch();
  //   // console.log(activeEpoch[0].toNumber());

  //   // console.log(await provider.getBlockNumber());
  //   // await provider.send('evm_mine', []);
  //   console.log(await provider.getBlockNumber());
  //   provider.on('block', (blockNumber: number) => {
  //     console.log('New block mined ', blockNumber);
  //   });

  const activeEpoch1 = await epochsManagerContract
    .connect(deployer)
    .getCurrentActiveEpoch();
  console.log('active epoch ', activeEpoch1[0].toNumber());

  console.log('here 5');

  const blockNumber = await provider.getBlockNumber();

  await epochsManagerContract
    .connect(deployer)
    .initializeEpoch({ gasLimit: 1_000_000 });
  await provider.send('evm_mine', []);
  console.log('here 6');

  const activeEpoch = await epochsManagerContract
    .connect(deployer)
    .getCurrentActiveEpoch();
  console.log('active epoch ', activeEpoch[0].toNumber());

  await new Promise(resolve => setTimeout(resolve, 5000));

  console.log('here 7');

  const blockNumber2 = await provider.getBlockNumber();
  console.log(`New block : ${blockNumber2}`);

  await epochsManagerContract
    .connect(deployer)
    .initializeEpoch({ gasLimit: 1_000_000 });
  await provider.send('evm_mine', []);

  console.log('here 8');

  const activeEpoch2 = await epochsManagerContract
    .connect(deployer)
    .getCurrentActiveEpoch();
  console.log('active epoch ', activeEpoch2[0].toNumber());

  await new Promise(resolve => setTimeout(resolve, 5000));

  // Keep listening for new blocks indefinitely
  // while (true) {
  //   const blockNumber = await provider.getBlockNumber();
  //   console.log(`New block mined: ${blockNumber}`);

  //   // provider.on('block', async (blockNumber: number) => {
  //   await epochsManagerContract
  //     .connect(wallet)
  //     .initializeEpoch({ gasLimit: 1_000_000 });

  //   console.log('inited epoch');

  //   const activeEpoch = await epochsManagerContract
  //     .connect(wallet)
  //     .getCurrentActiveEpoch();

  //   console.log('active epoch ', activeEpoch[0].toNumber());
  //   // });

  //   // Wait for a short interval before checking again
  //   // You can adjust the interval as needed
  //   await new Promise(resolve => setTimeout(resolve, 5000)); // Wait for 5 seconds before checking again
  // }
}

function connectSigner(
  wallet: ethers.Wallet,
  provider: ethers.providers.Provider,
): ethers.Wallet {
  const s = wallet.connect(provider);

  const sendTx = s.sendTransaction.bind(s);

  s.sendTransaction = async t => {
    const tx = await sendTx(t);
    await tx.wait(1);
    return tx;
  };

  return s;
}

getActiveEpoch();

// async function listenForNewBlocks() {
//   const provider = new ethers.providers.JsonRpcProvider(
//     'https://porcini.rootnet.app/',
//   );

//   // Subscribe to the "block" event
//   provider.on('block', (blockNumber: number) => {
//     console.log(`New block mined: ${blockNumber}`);
//   });

//   // Keep the process running to continue listening
//   // This can be stopped using a signal or other mechanism if needed
//   await new Promise(() => {});
// }

// listenForNewBlocks().catch(error => {
//   console.error('Error while listening for new blocks:', error);
// });
