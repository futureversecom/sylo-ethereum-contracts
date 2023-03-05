import { ethers } from 'ethers';
import * as factories from '../typechain-types';
import * as fs from 'fs/promises';
import { randomBytes } from 'crypto';
import contractAddress from './ganache.json';
import NodePK from './nodes.json';

var WINNING_PROBABILITY = ethers.BigNumber.from(2).pow(128).sub(1);

type Nodes = {
  deployerAccount: ethers.Wallet;
  incentivisedNode: ethers.Wallet;
  nodeAccountOne: ethers.Wallet;
  nodeAccountTwo: ethers.Wallet;
  nodeAccountThree: ethers.Wallet;
  nodeAccountFive: ethers.Wallet;
};

type Contracts = {
  stakingManager: factories.contracts.staking.StakingManager;
  token: factories.contracts.SyloToken;
  seekers: factories.contracts.mocks.TestSeekers;
  registries: factories.contracts.Registries;
  ticketing: factories.contracts.payments.SyloTicketing;
  ticketingParameters: factories.contracts.payments.ticketing.TicketingParameters;
  epochsManager: factories.contracts.epochs.EpochsManager;
};

async function main() {
  const provider = new ethers.providers.JsonRpcProvider('http://0.0.0.0:8545');

  const contracts = await conectContracts(provider);

  const nodes = await connectNodeAccounts(provider);

  await addStake(contracts, nodes);

  await registerNodes(contracts, nodes);

  await setSeekerRegistry(
    contracts,
    nodes.deployerAccount,
    nodes.deployerAccount,
    0,
    0,
  );
  await setSeekerRegistry(
    contracts,
    nodes.nodeAccountOne,
    nodes.deployerAccount,
    1,
    1,
  );
  await setSeekerRegistry(
    contracts,
    nodes.nodeAccountTwo,
    nodes.deployerAccount,
    2,
    2,
  );
  await setSeekerRegistry(
    contracts,
    nodes.nodeAccountThree,
    nodes.deployerAccount,
    3,
    3,
  );
  await setSeekerRegistry(
    contracts,
    nodes.nodeAccountFive,
    nodes.deployerAccount,
    4,
    4,
  );

  await setNetworkIncentives(contracts, nodes);

  await contracts.epochsManager.connect(nodes.deployerAccount).joinNextEpoch();
  await contracts.epochsManager.connect(nodes.nodeAccountOne).joinNextEpoch();
  await contracts.epochsManager.connect(nodes.nodeAccountTwo).joinNextEpoch();
  await contracts.epochsManager.connect(nodes.nodeAccountThree).joinNextEpoch();
  await contracts.epochsManager.connect(nodes.nodeAccountFive).joinNextEpoch();

  await initEpoch(contracts, nodes);
}

async function initEpoch(contracts: Contracts, nodes: Nodes) {
  await contracts.ticketingParameters
    .connect(nodes.deployerAccount)
    .setBaseLiveWinProb(WINNING_PROBABILITY);
  await contracts.ticketingParameters
    .connect(nodes.deployerAccount)
    .setFaceValue(10000);
  await contracts.ticketingParameters
    .connect(nodes.deployerAccount)
    .setTicketDuration(20);
  await contracts.epochsManager
    .connect(nodes.deployerAccount)
    .setEpochDuration(20);
  await contracts.epochsManager
    .connect(nodes.deployerAccount)
    .initializeEpoch();
}

async function setNetworkIncentives(contracts: Contracts, nodes: Nodes) {
  await contracts.token
    .connect(nodes.deployerAccount)
    .transfer(nodes.incentivisedNode.address, 1100000);
  await contracts.token
    .connect(nodes.incentivisedNode)
    .approve(contractAddress.ticketing, 1100000);
  await contracts.ticketing
    .connect(nodes.incentivisedNode)
    .depositEscrow(1000000, nodes.incentivisedNode.address);
  await contracts.ticketing
    .connect(nodes.incentivisedNode)
    .depositPenalty(100000, nodes.incentivisedNode.address);
}

async function setSeekerRegistry(
  contracts: Contracts,
  nodeAccount: ethers.Wallet,
  seekerAccount: ethers.Wallet,
  endpoint: number,
  tokenId: number,
): Promise<void> {
  if (!(await contracts.seekers.exists(tokenId))) {
    await contracts.seekers
      .connect(seekerAccount)
      .mint(await seekerAccount.getAddress(), tokenId);
  }

  const nonce = randomBytes(32);

  const accountAddress = await nodeAccount.getAddress();
  const proofMessage = await contracts.registries.getProofMessage(
    tokenId,
    accountAddress,
    nonce,
  );

  const signature = await seekerAccount.signMessage(
    Buffer.from(proofMessage.slice(2), 'hex'),
  );

  await contracts.registries
    .connect(nodeAccount)
    .setSeekerAccount(
      await seekerAccount.getAddress(),
      tokenId,
      nonce,
      signature,
    );
}

async function registerNodes(
  contracts: Contracts,
  nodes: Nodes,
): Promise<void> {
  await contracts.registries
    .connect(nodes.nodeAccountOne)
    .register('http://localhost/28901/public/metadata');
  await contracts.registries
    .connect(nodes.nodeAccountOne)
    .register('http://localhost/28903/public/metadata');
  await contracts.registries
    .connect(nodes.nodeAccountOne)
    .register('http://localhost/28905/public/metadata');
  await contracts.registries
    .connect(nodes.nodeAccountOne)
    .register('http://localhost/28907/public/metadata');
}

async function addStake(contracts: Contracts, nodes: Nodes): Promise<void> {
  await contracts.token
    .connect(nodes.deployerAccount)
    .approve(contractAddress.stakingManager, 90000000000000);
  await contracts.stakingManager
    .connect(nodes.deployerAccount)
    .addStake(100000, nodes.deployerAccount.address);

  // Approve and add stake Node one
  await contracts.token
    .connect(nodes.deployerAccount)
    .transfer(nodes.nodeAccountOne.address, 1100000);
  await contracts.token
    .connect(nodes.nodeAccountOne)
    .approve(contractAddress.stakingManager, 90000000000000);
  await contracts.stakingManager
    .connect(nodes.nodeAccountOne)
    .addStake(100000, nodes.nodeAccountOne.address);

  // Approve and add stake Node two
  await contracts.token
    .connect(nodes.deployerAccount)
    .transfer(nodes.nodeAccountTwo.address, 1100000);
  await contracts.token
    .connect(nodes.nodeAccountTwo)
    .approve(contractAddress.stakingManager, 90000000000000);
  await contracts.stakingManager
    .connect(nodes.nodeAccountTwo)
    .addStake(100000, nodes.nodeAccountTwo.address);

  // Approve and add stake Node three
  await contracts.token
    .connect(nodes.deployerAccount)
    .transfer(nodes.nodeAccountThree.address, 1100000);
  await contracts.token
    .connect(nodes.nodeAccountThree)
    .approve(contractAddress.stakingManager, 90000000000000);
  await contracts.stakingManager
    .connect(nodes.nodeAccountThree)
    .addStake(100000, nodes.nodeAccountThree.address);

  // Approve and add stake Node four
  await contracts.token
    .connect(nodes.deployerAccount)
    .transfer(nodes.nodeAccountFive.address, 1100000);
  await contracts.token
    .connect(nodes.nodeAccountFive)
    .approve(contractAddress.stakingManager, 90000000000000);
  await contracts.stakingManager
    .connect(nodes.nodeAccountFive)
    .addStake(100000, nodes.nodeAccountFive.address);
}

async function conectContracts(provider: ethers.providers.JsonRpcProvider) {
  const stakingManager = factories.StakingManager__factory.connect(
    contractAddress.stakingManager,
    provider,
  );

  const token = factories.SyloToken__factory.connect(
    contractAddress.token,
    provider,
  );

  const seekers = factories.TestSeekers__factory.connect(
    contractAddress.seekers,
    provider,
  );

  const registries = factories.Registries__factory.connect(
    contractAddress.registries,
    provider,
  );

  const ticketing = factories.SyloTicketing__factory.connect(
    contractAddress.ticketing,
    provider,
  );

  const ticketingParameters = factories.TicketingParameters__factory.connect(
    contractAddress.ticketingParameters,
    provider,
  );

  const epochsManager = factories.EpochsManager__factory.connect(
    contractAddress.epochsManager,
    provider,
  );

  return {
    token,
    stakingManager,
    seekers,
    registries,
    contractAddress,
    ticketing,
    ticketingParameters,
    epochsManager,
  } as Contracts;
}

async function connectNodeAccounts(provider: ethers.providers.JsonRpcProvider) {
  const deployerAccount = connectSigner(
    new ethers.Wallet(NodePK.deployer),
    provider,
  );

  const nodeAccountOne = connectSigner(
    new ethers.Wallet(NodePK.node1),
    provider,
  );

  const nodeAccountTwo = connectSigner(
    new ethers.Wallet(NodePK.node2),
    provider,
  );

  const nodeAccountThree = connectSigner(
    new ethers.Wallet(NodePK.node3),
    provider,
  );

  const nodeAccountFive = connectSigner(
    new ethers.Wallet(NodePK.node5),
    provider,
  );

  const incentivisedNode = connectSigner(
    new ethers.Wallet(NodePK.incentNode),
    provider,
  );

  return {
    deployerAccount,
    incentivisedNode,
    nodeAccountOne,
    nodeAccountTwo,
    nodeAccountThree,
    nodeAccountFive,
  } as Nodes;
}

export function connectSigner(
  wallet: ethers.Wallet,
  provider: ethers.providers.Provider,
): ethers.Wallet {
  const s = wallet.connect(provider);

  const sendTx = s.sendTransaction.bind(s);

  s.sendTransaction = async t => {
    t.gasLimit = 500000;
    const tx = await sendTx(t);
    await tx.wait(1);
    return tx;
  };

  return s;
}

main();
