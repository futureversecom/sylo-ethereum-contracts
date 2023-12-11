import { ethers } from 'ethers';
import * as Contracts from '../common/contracts';
import contractAddress from '../deployments/ganache_deployment_phase_two.json';
import { randomBytes } from 'crypto';
import { Permission } from '../common/enum';

export const MAX_WINNING_PROBABILITY = 2n ** 128n - 1n;

export type Node = {
  signer: ethers.Signer;
  publicEndPoint: string;
};

export type NodeConfig = {
  privateKey: string;
  publicEndpoint: string;
};

export type IncentivisingNodeConfig =
  | NodeConfig
  | {
      authorizedAccount: {
        address: string;
        description: string;
      };
    };

export async function updateFuturepassRegistrar(
  contracts: Contracts.SyloContracts,
  node: ethers.Signer,
): Promise<void> {
  await contracts.futurepassRegistrar.connect(node).create(node.getAddress(), {
    gasLimit: 1_000_000,
  });
}

export async function addStake(
  contracts: Contracts.SyloContracts,
  node: ethers.Signer,
): Promise<void> {
  await contracts.syloToken
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

export async function registerNodes(
  contracts: Contracts.SyloContracts,
  nodes: Node,
): Promise<void> {
  if (nodes.publicEndPoint != '') {
    await contracts.registries
      .connect(nodes.signer)
      .register(nodes.publicEndPoint, { gasLimit: 1_000_000 });
  }
}

export async function setSeekerRegistry(
  contracts: Contracts.SyloContracts,
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

export async function depositTicketing(
  contracts: Contracts.SyloContracts,
  incentivisedNode: ethers.Signer,
) {
  await contracts.syloToken
    .connect(incentivisedNode)
    .approve(contractAddress.syloTicketing, ethers.parseEther('1000000000'), {
      gasLimit: 1_000_000,
    });

  await contracts.syloTicketing
    .connect(incentivisedNode)
    .depositEscrow(
      ethers.parseEther('1000000'),
      incentivisedNode.getAddress(),
      { gasLimit: 1_000_000 },
    );

  await contracts.syloTicketing
    .connect(incentivisedNode)
    .depositPenalty(
      ethers.parseEther('100000'),
      incentivisedNode.getAddress(),
      { gasLimit: 1_000_000 },
    );
}

export async function authorizeAccount(
  contracts: Contracts.SyloContracts,
  main: ethers.Signer,
  authorized: string,
) {
  await contracts.authorizedAccounts
    .connect(main)
    .authorizeAccount(authorized, [Permission.PersonalSign], {
      gasLimit: 1_000_000,
    });
}

export async function setNetworkParams(
  contracts: Contracts.SyloContracts,
  deployer: ethers.Signer,
) {
  await contracts.ticketingParameters
    .connect(deployer)
    .setBaseLiveWinProb(MAX_WINNING_PROBABILITY, { gasLimit: 1_000_000 });

  await contracts.ticketingParameters
    .connect(deployer)
    .setFaceValue(ethers.parseEther('100'), { gasLimit: 1_000_000 });

  await contracts.ticketingParameters
    .connect(deployer)
    .setTicketDuration(1_000_000, { gasLimit: 1_000_000 });

  await contracts.syloTicketing
    .connect(deployer)
    .setUnlockDuration(5, { gasLimit: 1_000_000 });

  await contracts.stakingManager
    .connect(deployer)
    .setUnlockDuration(5, { gasLimit: 1_000_000 });

  await contracts.epochsManager
    .connect(deployer)
    .setEpochDuration(10, { gasLimit: 1_000_000 });
}
