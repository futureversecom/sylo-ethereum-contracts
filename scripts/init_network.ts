import { ethers } from 'ethers';
import { randomBytes } from 'crypto';
import contractAddress from '../deploy/ganache_deployment_phase_two.json';
import Nodes from './nodes.json';
import * as utils from './utils';

const WINNING_PROBABILITY = ethers.BigNumber.from(2).pow(128).sub(1);

type Node = {
  signer: ethers.Signer;
  publicEndPoint: string;
};

type NodeConfig = {
  privateKey: string;
  publicEndpoint: string;
};

async function main() {
  let nodeList: Node[] = [];

  const provider = new ethers.providers.JsonRpcProvider('http://0.0.0.0:8545');

  const contracts = await utils.conectContracts(contractAddress, provider);

  const deployerAccount = connectSigner(
    new ethers.Wallet(Nodes.deployerPK),
    provider,
  );

  const incentivisedNode = connectSigner(
    new ethers.Wallet(Nodes.incentiveNodePK),
    provider,
  );

  for (const nodeConfig of Nodes.relayNodes) {
    nodeList.push(await createNode(provider, nodeConfig));
  }

  for (let i = 0; i < Nodes.relayNodes.length; i++) {
    await addStake(contracts, nodeList[i].signer, deployerAccount);
  }

  for (let i = 0; i < Nodes.relayNodes.length; i++) {
    await registerNodes(contracts, nodeList[i]);
  }

  for (let i = 0; i < Nodes.relayNodes.length; i++) {
    await setSeekerRegistry(contracts, nodeList[i].signer, deployerAccount, i);
  }

  await contracts.token
    .connect(deployerAccount)
    .transfer(
      incentivisedNode.getAddress(),
      ethers.utils.parseEther('1000000000'),
    );

  await setNetworkIncentives(contracts, incentivisedNode);

  for (let i = 1; i < Nodes.relayNodes.length; i++) {
    await contracts.epochsManager.connect(nodeList[i].signer).joinNextEpoch();
  }

  await initNetwork(contracts, deployerAccount);
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
  contracts: utils.Contracts,
  nodes: ethers.Signer,
  deployer: ethers.Signer,
): Promise<void> {
  await contracts.token
    .connect(deployer)
    .transfer(nodes.getAddress(), ethers.utils.parseEther('110000'));

  await contracts.token
    .connect(nodes)
    .approve(
      contractAddress.stakingManager,
      ethers.utils.parseEther('1000000'),
    );

  await contracts.stakingManager
    .connect(nodes)
    .addStake(ethers.utils.parseEther('100000'), nodes.getAddress());
}

async function registerNodes(
  contracts: utils.Contracts,
  nodes: Node,
): Promise<void> {
  if (nodes.publicEndPoint != '') {
    await contracts.registries
      .connect(nodes.signer)
      .register(nodes.publicEndPoint);
  }
}

async function setSeekerRegistry(
  contracts: utils.Contracts,
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
  contracts: utils.Contracts,
  incentivisedNode: ethers.Signer,
) {
  await contracts.token
    .connect(incentivisedNode)
    .approve(contractAddress.ticketing, ethers.utils.parseEther('1000000000'));

  await contracts.ticketing
    .connect(incentivisedNode)
    .depositEscrow(
      ethers.utils.parseEther('1000000'),
      incentivisedNode.getAddress(),
    );

  await contracts.ticketing
    .connect(incentivisedNode)
    .depositPenalty(
      ethers.utils.parseEther('100000'),
      incentivisedNode.getAddress(),
    );
}

async function initNetwork(
  contracts: utils.Contracts,
  deployer: ethers.Signer,
) {
  await contracts.ticketingParameters
    .connect(deployer)
    .setBaseLiveWinProb(WINNING_PROBABILITY);
  await contracts.ticketingParameters
    .connect(deployer)
    .setFaceValue(ethers.utils.parseEther('10000'));
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
    const tx = await sendTx(t);
    await tx.wait(1);
    return tx;
  };

  return s;
}

main();
