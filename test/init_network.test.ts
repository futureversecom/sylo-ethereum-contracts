import { assert, expect } from 'chai';
import { ethers } from 'hardhat';
import { Signer } from 'ethers';
import ContractAddresses from '../scripts/ganache.json';
import {
  EpochsManager,
  Registries,
  StakingManager,
  SyloToken,
  TestSeekers,
  TicketingParameters,
  SyloTicketing,
} from '../typechain-types';
import utils from './utils';

describe('init network', () => {
  let accounts: Signer[];
  let owner: string;

  let registries: Registries;
  let seekers: TestSeekers;
  let epochsManager: EpochsManager;
  let ticketing: SyloTicketing;
  let ticketingParameter: TicketingParameters;
  let token: SyloToken;
  let stakingManager: StakingManager;

  before(async () => {
    accounts = await ethers.getSigners();
    // first account is implicitly used as deployer of contracts in hardhat
    owner = await accounts[0].getAddress();

    const Token = await ethers.getContractFactory('SyloToken');
    token = await Token.deploy();
  });

  beforeEach(async () => {
    const contracts = await utils.initializeContracts(owner, token.address);
    console.log(owner);
    epochsManager = contracts.epochsManager;
    stakingManager = contracts.stakingManager;
    registries = contracts.registries;
    seekers = contracts.seekers;
  });

  it('can add stake to nodes', async () => {
    // Approve and add stake Node one
    for (let i = 1; i <= 4; i++) {
      await token.transfer(accounts[i].getAddress(), 1100000);
      await token
        .connect(accounts[i])
        .approve(stakingManager.address, 90000000000000);
      await stakingManager
        .connect(accounts[i])
        .addStake(100000, accounts[i].getAddress());
    }
    for (let i = 1; i <= 4; i++) {
      expect(
        await stakingManager.getStakeeTotalManagedStake(
          accounts[i].getAddress(),
        ),
      ).to.be.equal(100000);
    }
  });

  it.only('can register nodes', async () => {
    await registries.connect(accounts[1]).register('http://0.0.0.0/29170');
    await registries.connect(accounts[2]).register('http://0.0.0.0/29171');
    await registries.connect(accounts[3]).register('http://0.0.0.0/29172');
    await registries.connect(accounts[4]).register('http://0.0.0.0/29173');

    const registryOne = await registries.getRegistry(accounts[1].getAddress());
    const registryTwo = await registries.getRegistry(accounts[2].getAddress());
    const registryThree = await registries.getRegistry(
      accounts[3].getAddress(),
    );
    const registryFour = await registries.getRegistry(accounts[4].getAddress());

    expect(registryOne.publicEndpoint).to.be.equal('http://0.0.0.0/29170');
    expect(registryTwo.publicEndpoint).to.be.equal('http://0.0.0.0/29171');
    expect(registryThree.publicEndpoint).to.be.equal('http://0.0.0.0/29172');
    expect(registryFour.publicEndpoint).to.be.equal('http://0.0.0.0/29173');
  });
});
