import { ethers } from 'hardhat';
import { BigNumberish, MaxUint256, Signer } from 'ethers';
import { toWei } from 'web3-utils';
import {
  Registries,
  SeekerPowerOracle,
  SyloToken,
  TestSeekers,
} from '../typechain-types';
import { randomBytes } from 'crypto';
import { SyloContracts } from '../common/contracts';

type Options = {
  faceValue?: BigNumberish;
  payoutPercentage?: number;
  baseLiveWinProb?: BigNumberish;
  expiredWinProb?: BigNumberish;
  decayRate?: number;
  ticketDuration?: number;
  initialEpoch?: number;
  epochDuration?: number;
  minimumStakeProportion?: number;
  seekerPowerMultiplier?: BigNumberish;
  unlockDuration?: number;
  seekerPowerOracleAccount?: string;
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

  const seekerPowerMultiplier =
    opts.seekerPowerMultiplier ?? ethers.parseEther('1000000');

  const seekerPowerOracleAccount = opts.seekerPowerOracleAccount ?? deployer;

  const tokenAddress = await syloToken.getAddress();

  const SeekersFactory = await ethers.getContractFactory('TestSeekers');
  const seekers = await SeekersFactory.deploy();

  const FuturepassRegistrarFactory = await ethers.getContractFactory(
    'TestFuturepassRegistrar',
  );
  const futurepassRegistrar = await FuturepassRegistrarFactory.deploy();

  const RegistriesFactory = await ethers.getContractFactory('Registries');
  const registries = await RegistriesFactory.deploy();
  await registries.initialize(await seekers.getAddress(), payoutPercentage, {
    from: deployer,
  });

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

  const EpochsManagerFactory = await ethers.getContractFactory('EpochsManager');
  const epochsManager = await EpochsManagerFactory.deploy();

  const StakingManagerFactory = await ethers.getContractFactory(
    'StakingManager',
  );
  const stakingManager = await StakingManagerFactory.deploy();

  const RewardsManagerFactory = await ethers.getContractFactory(
    'RewardsManager',
  );
  const rewardsManager = await RewardsManagerFactory.deploy();

  const DirectoryFactory = await ethers.getContractFactory('Directory');
  const directory = await DirectoryFactory.deploy();

  const AuthorizedAccountFactory = await ethers.getContractFactory(
    'AuthorizedAccounts',
  );
  const authorizedAccounts = await AuthorizedAccountFactory.deploy();

  const SeekerPowerOracleFactory = await ethers.getContractFactory(
    'SeekerPowerOracle',
  );
  const seekerPowerOracle = await SeekerPowerOracleFactory.deploy();

  await stakingManager.initialize(
    tokenAddress,
    await rewardsManager.getAddress(),
    await epochsManager.getAddress(),
    await seekerPowerOracle.getAddress(),
    unlockDuration,
    minimumStakeProportion,
    seekerPowerMultiplier,
    { from: deployer },
  );
  await rewardsManager.initialize(
    tokenAddress,
    await stakingManager.getAddress(),
    await epochsManager.getAddress(),
    { from: deployer },
  );
  await directory.initialize(
    await stakingManager.getAddress(),
    await rewardsManager.getAddress(),
    {
      from: deployer,
    },
  );
  await epochsManager.initialize(
    await seekers.getAddress(),
    await directory.getAddress(),
    await registries.getAddress(),
    await ticketingParameters.getAddress(),
    opts.initialEpoch ?? (await ethers.provider.getBlockNumber()) + 1,
    epochDuration,
    { from: deployer },
  );
  await authorizedAccounts.initialize({ from: deployer });
  await seekerPowerOracle.initialize(seekerPowerOracleAccount, {
    from: deployer,
  });

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
    await futurepassRegistrar.getAddress(),
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
    seekerPowerOracle,
    futurepassRegistrar,
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
  seekerPowerOracle: SeekerPowerOracle,
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

  // set the seeker power default to the max value
  await seekerPowerOracle.registerSeekerPowerRestricted(tokenId, MaxUint256);
}

export default {
  initializeContracts,
  advanceBlock,
  setSeekerRegistry,
};
