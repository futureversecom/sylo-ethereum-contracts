import { ethers } from 'hardhat';
import { BigNumberish, Signer } from 'ethers';
import { toWei } from 'web3-utils';
import { Registries, SyloToken, TestSeekers } from '../typechain-types';
import { randomBytes } from 'crypto';
import { SyloContracts } from '../common/contracts';

type Options = {
  faceValue?: BigNumberish;
  payoutPercentage?: number;
  baseLiveWinProb?: BigNumberish;
  expiredWinProb?: BigNumberish;
  decayRate?: number;
  ticketDuration?: number;
  epochDuration?: number;
  minimumStakeProportion?: number;
  unlockDuration?: number;
};

const initializeContracts = async function (
  deployer: string,
  syloToken: SyloToken,
  opts: Options = {},
): Promise<SyloContracts> {
  const payoutPercentage = opts.payoutPercentage ? opts.payoutPercentage : 5000;

  const faceValue = opts.faceValue ?? toWei('15');
  const baseLiveWinProb = opts.baseLiveWinProb ?? 2n ** 128n - 1n;
  const expiredWinProb = opts.expiredWinProb ?? 1000;
  const decayRate = opts.decayRate ?? 8000;
  const ticketDuration = opts.ticketDuration ?? 20;

  const epochDuration = opts.epochDuration ?? 30;

  const unlockDuration = opts.unlockDuration ?? 10;

  const minimumStakeProportion = opts.minimumStakeProportion ?? 2000;

  const tokenAddress = await syloToken.getAddress();

  console.log('1');
  const SeekersFactory = await ethers.getContractFactory('TestSeekers');
  const seekers = await SeekersFactory.deploy();
  console.log('1');

  const RegistriesFactory = await ethers.getContractFactory('Registries');
  const registries = await RegistriesFactory.deploy();
  await registries.initialize(await seekers.getAddress(), payoutPercentage, {
    from: deployer,
  });
  console.log('1');

  const TicketingParametersFactory = await ethers.getContractFactory(
    'TicketingParameters',
  );
  const ticketingParameters = await TicketingParametersFactory.deploy();
  await ticketingParameters.initialize(
    faceValue,
    baseLiveWinProb,
    expiredWinProb,
    decayRate,
    ticketDuration,
    { from: deployer },
  );
  console.log('1');

  const EpochsManagerFactory = await ethers.getContractFactory('EpochsManager');
  const epochsManager = await EpochsManagerFactory.deploy();
  console.log('1');

  const StakingManagerFactory = await ethers.getContractFactory(
    'StakingManager',
  );
  const stakingManager = await StakingManagerFactory.deploy();
  console.log('1');

  const RewardsManagerFactory = await ethers.getContractFactory(
    'RewardsManager',
  );
  const rewardsManager = await RewardsManagerFactory.deploy();
  console.log('1');

  const DirectoryFactory = await ethers.getContractFactory('Directory');
  const directory = await DirectoryFactory.deploy();
  console.log('1');

  const AuthorizedAccountFactory = await ethers.getContractFactory(
    'AuthorizedAccounts',
  );
  const authorizedAccounts = await AuthorizedAccountFactory.deploy();
  console.log('1');

  await stakingManager.initialize(
    tokenAddress,
    await rewardsManager.getAddress(),
    await epochsManager.getAddress(),
    unlockDuration,
    minimumStakeProportion,
    { from: deployer },
  );
  console.log('2');

  await rewardsManager.initialize(
    tokenAddress,
    await stakingManager.getAddress(),
    await epochsManager.getAddress(),
    { from: deployer },
  );
  console.log('1');

  await directory.initialize(
    await stakingManager.getAddress(),
    await rewardsManager.getAddress(),
    {
      from: deployer,
    },
  );
  console.log('hero');

  await epochsManager.initialize(
    await seekers.getAddress(),
    await directory.getAddress(),
    await registries.getAddress(),
    await ticketingParameters.getAddress(),
    epochDuration,
    { from: deployer },
  );
  console.log('1');

  await authorizedAccounts.initialize({ from: deployer });
  console.log('1');

  const TicketingFactory = await ethers.getContractFactory('SyloTicketing');
  const syloTicketing = await TicketingFactory.deploy();
  await syloTicketing.initialize(
    tokenAddress,
    await registries.getAddress(),
    await stakingManager.getAddress(),
    await directory.getAddress(),
    await epochsManager.getAddress(),
    await rewardsManager.getAddress(),
    await authorizedAccounts.getAddress(),
    unlockDuration,
    { from: deployer },
  );

  await rewardsManager.addManager(await syloTicketing.getAddress(), {
    from: deployer,
  });
  await rewardsManager.addManager(await stakingManager.getAddress(), {
    from: deployer,
  });
  await rewardsManager.addManager(await epochsManager.getAddress(), {
    from: deployer,
  });

  await directory.addManager(await epochsManager.getAddress());
  console.log('made it');
  return {
    syloToken,
    authorizedAccounts,
    registries,
    ticketingParameters,
    epochsManager,
    stakingManager,
    rewardsManager,
    directory,
    syloTicketing,
    seekers,
  };
};

const advanceBlock = async function (i: number): Promise<void> {
  i = i || 1;
  for (let j = 0; j < i; j++) {
    await ethers.provider.send('evm_mine', []);
  }
};

async function setSeekerRegistry(
  registries: Registries,
  seekers: TestSeekers,
  account: Signer,
  seekerAccount: Signer,
  tokenId: number,
): Promise<void> {
  if (!(await seekers.exists(tokenId))) {
    await seekers.mint(await seekerAccount.getAddress(), tokenId);
  }

  const nonce = randomBytes(32);

  const accountAddress = await account.getAddress();
  const proofMessage = await registries.getProofMessage(
    tokenId,
    accountAddress,
    nonce,
  );

  const signature = await seekerAccount.signMessage(
    Buffer.from(proofMessage.slice(2), 'hex'),
  );

  await registries.connect(account).register('0.0.0.0/0');

  await registries
    .connect(account)
    .setSeekerAccount(
      await seekerAccount.getAddress(),
      tokenId,
      nonce,
      signature,
    );
}

export default {
  initializeContracts,
  advanceBlock,
  setSeekerRegistry,
};
