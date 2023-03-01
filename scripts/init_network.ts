import { ethers } from 'ethers';
import * as factories from '../typechain-types';
import * as fs from 'fs/promises';
import { randomBytes } from 'crypto';
import contractAddress from './ganache.json';
import NodeData from './nodes.json';

var WINNING_PROBABILITY = ethers.BigNumber.from(2).pow(128).sub(1);

type Node = {
  signer: ethers.Signer;
  publicEndPoint: string;
};

type NodeConfig = {
  privateKey: string;
  publicEndpoint: string;
};

const nodeConfigs: NodeConfig[] = [
  {
    privateKey: NodeData.incentNode,
    publicEndpoint: NodeData.incentNodePublicEndPoint,
  },
  {
    privateKey: NodeData.deployer,
    publicEndpoint: '',
  },
  {
    privateKey: NodeData.nodeOnePK,
    publicEndpoint: NodeData.nodeOnePublicEndPoint,
  },
  {
    privateKey: NodeData.nodeTwoPK,
    publicEndpoint: NodeData.nodeTwoPublicEndPoint,
  },
  {
    privateKey: NodeData.nodeThreePK,
    publicEndpoint: NodeData.nodeThreePublicEndPoint,
  },
  {
    privateKey: NodeData.nodeFivePK,
    publicEndpoint: NodeData.nodeFivePublicEndPoint,
  },
];

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
  let nodeList: Node[] = [];

  const provider = new ethers.providers.JsonRpcProvider('http://0.0.0.0:8545');

  const contracts = await conectContracts(provider);

  for (let i = 0; i < nodeConfigs.length; i++) {
    nodeList[i] = await createNode(provider, nodeConfigs[i]);
  }

  const deployerAccount = nodeList[1];
  const incentivisedNode = nodeList[0];

  for (let i = 0; i < nodeConfigs.length; i++) {
    await addStake(contracts, nodeList[i].signer, deployerAccount.signer);
  }

  for (let i = 0; i < nodeConfigs.length; i++) {
    await registerNodes(contracts, nodeList[i]);
  }

  for (let i = 0; i < nodeConfigs.length; i++) {
    await setSeekerRegistry(
      contracts,
      nodeList[i].signer,
      deployerAccount.signer,
      i,
    );
  }

  await setNetworkIncentives(
    contracts,
    deployerAccount.signer,
    incentivisedNode.signer,
  );

  for (let i = 1; i < nodeConfigs.length; i++) {
    await contracts.epochsManager.connect(nodeList[i].signer).joinNextEpoch();
  }

  await initEpoch(contracts, deployerAccount.signer);
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

async function createNode(
  provider: ethers.providers.JsonRpcProvider,
  nodeConfig: NodeConfig,
): Promise<Node> {
  const newNode = connectSigner(
    new ethers.Wallet(nodeConfig.privateKey),
    provider,
  );

  return {
    signer: newNode,
    publicEndPoint: nodeConfig.publicEndpoint,
  };
}

async function addStake(
  contracts: Contracts,
  nodes: ethers.Signer,
  deployer: ethers.Signer,
): Promise<void> {
  if (nodes.getAddress() == deployer.getAddress()) {
    await contracts.token
      .connect(deployer)
      .approve(contractAddress.stakingManager, 90000000000000);
    await contracts.stakingManager
      .connect(nodes)
      .addStake(100000, nodes.getAddress());
  } else {
    await contracts.token
      .connect(deployer)
      .transfer(nodes.getAddress(), 1100000);
    await contracts.token
      .connect(nodes)
      .approve(contractAddress.stakingManager, 90000000000000);
    await contracts.stakingManager
      .connect(nodes)
      .addStake(100000, nodes.getAddress());
  }
}

async function registerNodes(contracts: Contracts, nodes: Node): Promise<void> {
  if (nodes.publicEndPoint != '') {
    await contracts.registries
      .connect(nodes.signer)
      .register(nodes.publicEndPoint);
  }
}

async function setSeekerRegistry(
  contracts: Contracts,
  nodeAccount: ethers.Signer,
  seekerAccount: ethers.Signer,
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

async function setNetworkIncentives(
  contracts: Contracts,
  deployer: ethers.Signer,
  incentivisedNode: ethers.Signer,
) {
  await contracts.token
    .connect(deployer)
    .transfer(incentivisedNode.getAddress(), 1100000);
  await contracts.token
    .connect(incentivisedNode)
    .approve(contractAddress.ticketing, 1100000);
  await contracts.ticketing
    .connect(incentivisedNode)
    .depositEscrow(1000000, incentivisedNode.getAddress());
  await contracts.ticketing
    .connect(incentivisedNode)
    .depositPenalty(100000, incentivisedNode.getAddress());
}

async function initEpoch(contracts: Contracts, deployer: ethers.Signer) {
  await contracts.ticketingParameters
    .connect(deployer)
    .setBaseLiveWinProb(WINNING_PROBABILITY);
  await contracts.ticketingParameters.connect(deployer).setFaceValue(10000);
  await contracts.ticketingParameters.connect(deployer).setTicketDuration(20);
  await contracts.epochsManager.connect(deployer).setEpochDuration(20);
  await contracts.epochsManager.connect(deployer).initializeEpoch();
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
