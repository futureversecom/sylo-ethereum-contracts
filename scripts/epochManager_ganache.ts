import * as etherHRE from 'hardhat';
import { ethers } from 'ethers';
import * as factories from '../typechain-types';
import EpochsManagerContractABI from '../artifacts/contracts/epochs/EpochsManager.sol/EpochsManager.json';
// import TicketingParametersABI from '../artifacts/contracts/payments/ticketing/TicketingParameters.sol/TicketingParameters.json';
import contractAddress from '../deployments/porcini_deployment_phase_two.json';

// const WINNING_PROBABILITY = ethers.BigNumber.from(2).pow(128).sub(1);

async function getActiveEpoch() {
  const provider = new ethers.JsonRpcProvider(
    'https://porcini.rootnet.app/', // https://porcini.au.rootnet.app // https://porcini.rootnet.app/
  );
  // /*
  //  hardhat / ganache testnet
  // */
  // const provider = new ethers.providers.JsonRpcProvider(
  //   'http://127.0.0.1:8545/',
  // );

  const deployer = connectSigner(
    new ethers.Wallet(
      '8e8cd52d61e3f01e75988ca45d2fb9adfeb06b68ab08a3b67525a9179ecae6f6',
    ),
    provider,
  );

  // const epochsManagerContract = new ethers.Contract(
  //   contractAddress.epochsManager,
  //   EpochsManagerContractABI.abi,
  //   wallet,
  // );

  const epochsManagerContract = factories.EpochsManager__factory.connect(
    '0x38839485d4a8aE7644878Bf3C7666e0b40AC6e94', // 0xd9f69a6dE82630558E468f219865bFe7247ba35E // 0x69e0A00Bd41F733BF6AE088692Fe916f94e9a9CF
    provider,
  );

  const ticketContract = factories.SyloTicketing__factory.connect(
    '0xf97f621C812C160497003D2fa215335E3E1DA794', // 0xd9f69a6dE82630558E468f219865bFe7247ba35E // 0x69e0A00Bd41F733BF6AE088692Fe916f94e9a9CF
    provider,
  );

  const token = factories.SyloToken__factory.connect(
    contractAddress.syloToken, // 0xd9f69a6dE82630558E468f219865bFe7247ba35E // 0x69e0A00Bd41F733BF6AE088692Fe916f94e9a9CF
    provider,
  );

  const param = factories.TicketingParameters__factory.connect(
    contractAddress.ticketingParameters, // 0xd9f69a6dE82630558E468f219865bFe7247ba35E // 0x69e0A00Bd41F733BF6AE088692Fe916f94e9a9CF
    provider,
  );

  const [deployerr] = await etherHRE.ethers.getSigners();

  const x = await param.getTicketingParameters();
  console.log('params before ', x);

  await param.connect(deployerr).setFaceValue(ethers.parseEther('1'));

  // const y = await param.getTicketingParameters();
  // console.log('params after ', y);

  // ten 10000000000000000000n
  //one? 1000000000000000000n
  const depo = await ticketContract
    .connect(deployer)
    .deposits(deployer.address);
  console.log(depo);
  console.log('addy ', deployer.address);

  // await token
  //   .connect(deployer)
  //   .approve(contractAddress.syloTicketing, ethers.parseEther('1000000000'), {
  //     gasLimit: 1_000_000,
  //   });

  //   console.log('approved ');

  // await ticketContract
  //   .connect(deployer)
  //   .depositPenalty(ethers.parseEther('10000'), deployer.address);

  // await ticketContract
  //   .connect(deployer)
  //   .depositEscrow(ethers.parseEther('10000'), deployer.address);

  const depotwo = await ticketContract
    .connect(deployer)
    .deposits(deployer.address);
  console.log('depo two ', depotwo);

  const activeEpoch1 = await epochsManagerContract
    .connect(deployer)
    .getCurrentActiveEpoch();
  console.log('active epoch ', activeEpoch1[0]);

  console.log('here 5');

  // const blockNumber = await provider.getBlockNumber();

  await epochsManagerContract.connect(deployerr).initializeEpoch();
  // await new Promise(resolve => setTimeout(resolve, 5000));

  console.log('here 6');

  const activeEpoch = await epochsManagerContract
    .connect(deployer)
    .getCurrentActiveEpoch();
  console.log('active epoch ', activeEpoch[0]);

  // console.log('here 7');

  // const blockNumber2 = await provider.getBlockNumber();
  // console.log(`New block : ${blockNumber2}`);

  // await epochsManagerContract
  //   .connect(deployer)
  //   .initializeEpoch({ gasLimit: 1_000_000 });
  // await provider.send('evm_mine', []);

  // console.log('here 8');

  // const activeEpoch2 = await epochsManagerContract
  //   .connect(deployer)
  //   .getCurrentActiveEpoch();
  // console.log('active epoch ', activeEpoch2[0].toNumber());

  // await new Promise(resolve => setTimeout(resolve, 5000));

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
  provider: ethers.Provider,
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
