import { ethers } from 'ethers';
import { randomBytes } from 'crypto';
import contractAddress from '../deployments/ganache_deployment_phase_two.json';
import nodesConfig from './nodes.json';
import * as utils from './utils';

const WINNING_PROBABILITY = BigInt(2) ** BigInt(128) - BigInt(1);

type Node = {
  signer: ethers.Signer;
  publicEndPoint: string;
};

type NodeConfig = {
  privateKey: string;
  publicEndpoint: string;
};

async function main() {
  const provider = new ethers.JsonRpcProvider('http://0.0.0.0:8545');

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
      .transfer(node.signer.getAddress(), ethers.parseEther('110000'));

    await addStake(contracts, node.signer);
    await registerNodes(contracts, node);
    await setSeekerRegistry(contracts, node.signer, deployer, i);
    await contracts.epochsManager
      .connect(node.signer)
      .joinNextEpoch({ gasLimit: 1_000_000 });

    console.log('Relay node', i, 'is ready');
  }

  // process incentivising nodes
  for (let i = 0; i < nodesConfig.incentivisingNodes.length; i++) {
    const node = await createNode(provider, nodesConfig.incentivisingNodes[i]);

    await contracts.token
      .connect(deployer)
      .transfer(node.signer.getAddress(), ethers.parseEther('1000000000'));

    await registerNodes(contracts, node);
    await depositTicketing(contracts, node.signer);
    console.log('Incentivising node', i, 'is ready');
  }

  // initialize next epoch
  await contracts.epochsManager
    .connect(deployer)
    .initializeEpoch({ gasLimit: 1_000_000 });
}

async function createNode(
  provider: ethers.JsonRpcProvider,
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
    .approve(contractAddress.stakingManager, ethers.parseEther('1000000'), {
      gasLimit: 1_000_000,
    });

  await contracts.stakingManager
    .connect(node)
    .addStake(ethers.parseEther('100000'), node.getAddress(), {
      gasLimit: 1_000_000,
    });
}

async function registerNodes(
  contracts: utils.Contracts,
  nodes: Node,
): Promise<void> {
  if (nodes.publicEndPoint != '') {
    await contracts.registries
      .connect(nodes.signer)
      .register(nodes.publicEndPoint, { gasLimit: 1_000_000 });
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
      .mint(await seekerAccount.getAddress(), tokenId, { gasLimit: 1_000_000 });
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
      { gasLimit: 1_000_000 },
    );
}

async function depositTicketing(
  contracts: utils.Contracts,
  incentivisedNode: ethers.Signer,
) {
  await contracts.token
    .connect(incentivisedNode)
    .approve(contractAddress.ticketing, ethers.parseEther('1000000000'), {
      gasLimit: 1_000_000,
    });

  await contracts.ticketing
    .connect(incentivisedNode)
    .depositEscrow(
      ethers.parseEther('1000000'),
      incentivisedNode.getAddress(),
      { gasLimit: 1_000_000 },
    );

  await contracts.ticketing
    .connect(incentivisedNode)
    .depositPenalty(
      ethers.parseEther('100000'),
      incentivisedNode.getAddress(),
      { gasLimit: 1_000_000 },
    );
}

async function setNetworkParams(
  contracts: utils.Contracts,
  deployer: ethers.Signer,
) {
  await contracts.ticketingParameters
    .connect(deployer)
    .setBaseLiveWinProb(WINNING_PROBABILITY, { gasLimit: 1_000_000 });

  await contracts.ticketingParameters
    .connect(deployer)
    .setFaceValue(ethers.parseEther('100'), { gasLimit: 1_000_000 });

  await contracts.ticketingParameters
    .connect(deployer)
    .setTicketDuration(1_000_000, { gasLimit: 1_000_000 });

  await contracts.ticketing
    .connect(deployer)
    .setUnlockDuration(5, { gasLimit: 1_000_000 });

  await contracts.stakingManager
    .connect(deployer)
    .setUnlockDuration(5, { gasLimit: 1_000_000 });

  await contracts.epochsManager
    .connect(deployer)
    .setEpochDuration(10, { gasLimit: 1_000_000 });
}

export function connectSigner(
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

main();
