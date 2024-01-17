import { ethers } from 'ethers';
import contractAddress from '../deployments/ganache_deployment_phase_two.json';
import nodesConfig from './nodes.json';
import * as Contracts from '../common/contracts';
import * as utils from './utils';

async function main() {
  const provider = new ethers.JsonRpcProvider('http://0.0.0.0:8545');

  const contracts = Contracts.connectContracts(contractAddress, provider);

  const deployer = connectSigner(
    new ethers.Wallet(nodesConfig.deployerPK),
    provider,
  );

  await utils.setNetworkParams(contracts, deployer);
  console.log('Network params are set');

  // process relay nodes
  for (let i = 0; i < nodesConfig.relayNodes.length; i++) {
    const node = await createNode(provider, nodesConfig.relayNodes[i]);

    await contracts.syloToken
      .connect(deployer)
      .transfer(node.signer.getAddress(), ethers.parseEther('110000'));

    await utils.updateFuturepassRegistrar(contracts, node.signer);
    await utils.addStake(contracts, node.signer);
    await utils.registerNodes(contracts, node);
    await utils.setSeekerRegistry(contracts, node.signer, deployer, i);
    await contracts.epochsManager
      .connect(node.signer)
      .joinNextEpoch({ gasLimit: 1_000_000 });

    console.log('Relay node', i, 'is ready');
  }

  // process incentivising nodes
  for (let i = 0; i < nodesConfig.incentivisingNodes.length; i++) {
    const node = await createNode(provider, nodesConfig.incentivisingNodes[i]);

    await contracts.syloToken
      .connect(deployer)
      .transfer(node.signer.getAddress(), ethers.parseEther('1000000000'));

    await utils.updateFuturepassRegistrar(contracts, node.signer);
    await utils.registerNodes(contracts, node);
    await utils.depositTicketing(contracts, node.signer);

    if (
      nodesConfig.incentivisingNodes[i].authorizedAccount.address.length > 0
    ) {
      await utils.authorizeAccount(
        contracts,
        node.signer,
        nodesConfig.incentivisingNodes[i].authorizedAccount.address,
      );
    }

    console.log('Incentivising node', i, 'is ready');
  }

  // initialize next epoch
  await contracts.epochsManager
    .connect(deployer)
    .initializeEpoch({ gasLimit: 1_000_000 });
}

async function createNode(
  provider: ethers.JsonRpcProvider,
  nodeConfig: utils.NodeConfig,
): Promise<utils.Node> {
  const newNode = connectSigner(
    new ethers.Wallet(nodeConfig.privateKey),
    provider,
  );

  return {
    signer: newNode,
    publicEndPoint: nodeConfig.publicEndpoint,
  };
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
