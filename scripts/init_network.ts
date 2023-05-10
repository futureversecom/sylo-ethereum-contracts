import { ethers } from 'ethers';
import { randomBytes } from 'crypto';
import contractAddress from '../deploy/ganache_deployment_phase_two.json';
import nodesConfig from './nodes.json';
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
  const provider = new ethers.providers.JsonRpcProvider('http://0.0.0.0:8545');

  const contracts = utils.conectContracts(contractAddress, provider);

  const deployer = connectSigner(
    new ethers.Wallet(nodesConfig.deployerPK),
    provider,
  );

  await setNetworkParams(contracts, deployer);
  console.log('Network params are set');

  // process relay nodes
  for (let i = 0; i < nodesConfig.relayNodes.length; i++) {
    const node = await createNode(provider, nodesConfig.relayNodes[i]);

    await contracts.token
      .connect(deployer)
      .transfer(node.signer.getAddress(), ethers.utils.parseEther('110000'));

    await addStake(contracts, node.signer);
    await registerNodes(contracts, node);
    await setSeekerRegistry(contracts, node.signer, deployer, i);
    await contracts.epochsManager.connect(node.signer).joinNextEpoch();

    console.log('Relay node', i, 'is ready');
  }

  // process incentivising nodes
  for (let i = 0; i < nodesConfig.incentivisingNodes.length; i++) {
    const node = await createNode(provider, nodesConfig.incentivisingNodes[i]);

    await contracts.token
      .connect(deployer)
      .transfer(
        node.signer.getAddress(),
        ethers.utils.parseEther('1000000000'),
      );

    await registerNodes(contracts, node);
    await depositTicketing(contracts, node.signer);
    console.log('Incentivising node', i, 'is ready');
  }

  // initialize next epoch
  await contracts.epochsManager.connect(deployer).initializeEpoch();
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
  node: ethers.Signer,
): Promise<void> {
  await contracts.token
    .connect(node)
    .approve(
      contractAddress.stakingManager,
      ethers.utils.parseEther('1000000'),
    );

  await contracts.stakingManager
    .connect(node)
    .addStake(ethers.utils.parseEther('100000'), node.getAddress());
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

async function depositTicketing(
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

async function setNetworkParams(
  contracts: utils.Contracts,
  deployer: ethers.Signer,
) {
  await contracts.ticketingParameters
    .connect(deployer)
    .setBaseLiveWinProb(WINNING_PROBABILITY);

  await contracts.ticketingParameters
    .connect(deployer)
    .setFaceValue(ethers.utils.parseEther('100'));

  await contracts.ticketingParameters
    .connect(deployer)
    .setTicketDuration(1_000_000);

  await contracts.ticketing.connect(deployer).setUnlockDuration(5);
  await contracts.stakingManager.connect(deployer).setUnlockDuration(5);
  await contracts.epochsManager.connect(deployer).setEpochDuration(10);
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
