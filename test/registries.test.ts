import { ethers } from 'hardhat';
import { Signer } from 'ethers';
import { Registries, TestSeekers } from '../typechain-types';
import { assert, expect } from 'chai';
import utils from './utils';
import { randomBytes } from 'crypto';

describe('Registries', () => {
  let accounts: Signer[];
  let owner: string;

  let registries: Registries;
  let seekers: TestSeekers;

  before(async () => {
    accounts = await ethers.getSigners();
    // first account is implicitly used as deployer of contracts in hardhat
    owner = await accounts[0].getAddress();
  });

  beforeEach(async () => {
    const Token = await ethers.getContractFactory('SyloToken');
    const token = await Token.deploy();

    const contracts = await utils.initializeContracts(owner, token.address, {
      payoutPercentage: 5000,
    });
    registries = contracts.registries;
    seekers = contracts.seekers;
  });

  it('registries cannot initialize twice', async () => {
    await expect(
      registries.initialize(seekers.address, 5000),
    ).to.be.revertedWith('Initializable: contract is already initialized');
  });

  it('registries cannot initialize with invalid arguments', async () => {
    const Registries = await ethers.getContractFactory('Registries');
    registries = await Registries.deploy();

    await expect(
      registries.initialize(ethers.constants.AddressZero, 5000),
    ).to.be.revertedWithCustomError(
      registries,
      'RootSeekersCannotBeZeroAddress',
    );

    await expect(
      registries.initialize(seekers.address, 10001),
    ).to.be.revertedWithCustomError(registries, 'PercentageCannotExceed10000');
  });

  it('can allow owner to set default payout percentage', async () => {
    await expect(registries.setDefaultPayoutPercentage(2000))
      .to.emit(registries, 'DefaultPayoutPercentageUpdated')
      .withArgs(2000);

    const p = await registries.defaultPayoutPercentage();
    assert.equal(
      p,
      2000,
      'Expected default payout percentage to be correctly updated',
    );
  });

  it('not owner cannot set default payout percentage', async () => {
    await expect(
      registries.connect(accounts[1]).setDefaultPayoutPercentage(2000),
    ).to.be.revertedWith('Ownable: caller is not the owner');
  });

  it('can set registry', async () => {
    await registries.register('http://api');

    const registry = await registries.getRegistry(owner);

    assert.equal(
      registry.publicEndpoint,
      'http://api',
      'Expected registries to have correct address',
    );
  });

  it('can set registry twice', async () => {
    await registries.register('http://api');
    await registries.register('http://api2');

    const registry = await registries.getRegistry(owner);

    assert.equal(
      registry.publicEndpoint,
      'http://api2',
      'Expected registries to have correct address',
    );
  });

  it('can retrieve all registered nodes', async () => {
    await registries.register('http://api');
    await registries.connect(accounts[1]).register('http://api');

    const nodes = await registries.getNodes();

    assert.deepEqual(nodes, [
      await accounts[0].getAddress(),
      await accounts[1].getAddress(),
    ]);
  });

  it('can query total number of registered nodes', async () => {
    await registries.register('http://api');
    await registries.connect(accounts[1]).register('http://api');

    const n = await registries.getTotalNodes();

    assert.equal(n.toNumber(), 2);
  });

  it('can retrieve a list of registries', async () => {
    const addresses = await Promise.all(accounts.map(a => a.getAddress()));

    for (let i = 0; i < 20; i++) {
      await registries.connect(accounts[i]).register(`http://api/${i}`);
    }

    const result = await registries.getRegistries(0, 20);

    assert.deepEqual(result[0], addresses, 'Expected 20 registries returned');

    for (let i = 0; i < 20; i++) {
      assert.equal(
        result[0][i],
        addresses[i],
        'Expected correct registry to be returned',
      );
      assert.equal(
        result[1][i].publicEndpoint,
        `http://api/${i}`,
        'Expected correct registry to be returned',
      );
    }
  });

  it('can retrieve a list of registries with start and end indexes', async () => {
    const addresses = await Promise.all(accounts.map(a => a.getAddress()));

    for (let i = 0; i < 20; i++) {
      await registries.connect(accounts[i]).register(`http://api/${i}`);
    }

    const result = await registries.getRegistries(5, 10);

    assert.deepEqual(
      result[0],
      addresses.slice(5, 10),
      'Expected only accounts 5 to 9 to be returned from query',
    );

    for (let i = 5; i < 10; i++) {
      assert.equal(
        result[0][i - 5],
        addresses[i],
        'Expected correct registry to be returned',
      );
      assert.equal(
        result[1][i - 5].publicEndpoint,
        `http://api/${i}`,
        'Expected correct registry to be returned',
      );
    }

    await expect(registries.getRegistries(8, 5)).to.be.revertedWithCustomError(
      registries,
      'EndMustBeGreaterThanStart',
    );

    await expect(registries.getRegistries(8, 21))
      .to.be.revertedWithCustomError(registries, 'EndCannotExceedNumberOfNodes')
      .withArgs(20);
  });

  it('requires default payout percentage to not exceed 100%', async () => {
    await expect(
      registries.setDefaultPayoutPercentage(10001),
    ).to.be.revertedWithCustomError(registries, 'PercentageCannotExceed10000');
  });

  it('can set seeker account with valid proof', async () => {
    const seekerAccount = accounts[1];
    const seekerAddress = await seekerAccount.getAddress();

    await utils.setSeekerRegistry(
      registries,
      seekers,
      accounts[0],
      accounts[1],
      1,
    );

    const registry = await registries.getRegistry(owner);

    expect(registry.seekerAccount).to.equal(seekerAddress);
    expect(registry.seekerId).to.equal(1);
  });

  it('fails to set seeker account when seekerAccount is zero address', async () => {
    await expect(
      registries.setSeekerAccount(
        ethers.constants.AddressZero,
        1,
        randomBytes(32),
        '0x',
      ),
    ).to.be.revertedWithCustomError(
      registries,
      'SeekerAccountCannotBeZeroAddress',
    );
  });

  it('fails to set seeker account when seekerId greater than max seeker id', async () => {
    const seekerAccount = accounts[1];
    const seekerAddress = await seekerAccount.getAddress();

    await expect(
      registries.setSeekerAccount(seekerAddress, 50000, randomBytes(32), '0x'),
    ).to.be.revertedWithCustomError(registries, 'SeekerIdOutOfRange');
  });

  it('fails to set seeker account when reusing signature', async () => {
    const seekerAccount = accounts[1];
    const seekerAddress = await seekerAccount.getAddress();

    const tokenId = 1;
    await seekers.mint(seekerAddress, tokenId);

    const nonce = randomBytes(32);
    const proofMessage = await registries.getProofMessage(
      tokenId,
      await accounts[0].getAddress(),
      nonce,
    );

    const signature = await seekerAccount.signMessage(
      Buffer.from(proofMessage.slice(2), 'hex'),
    );

    await registries.connect(accounts[0]).register('0.0.0.0/0');

    // first attempt should be valid
    await registries.setSeekerAccount(seekerAddress, 1, nonce, signature);

    // second attempt should fail due to nonce reuse
    await expect(
      registries.setSeekerAccount(seekerAddress, 1, nonce, signature),
    ).to.be.revertedWithCustomError(registries, 'NonceCannotBeReused');
  });

  it('fails to set seeker account with invalid proof', async () => {
    const seekerAccount = accounts[1];
    const seekerAddress = await seekerAccount.getAddress();

    const tokenId = 1;
    await seekers.mint(seekerAddress, tokenId);

    const nonce = randomBytes(32);
    const proofMessage = await registries.getProofMessage(
      tokenId,
      await accounts[0].getAddress(),
      nonce,
    );

    // sign proof with wrong account
    const proof = await accounts[2].signMessage(proofMessage);

    await expect(
      registries.setSeekerAccount(seekerAddress, tokenId, nonce, proof),
    ).to.be.revertedWithCustomError(
      registries,
      'ProofNotSignedBySeekerAccount',
    );
  });

  it("fails to set seeker account if seeker isn't owned by account", async () => {
    const seekerAccount = accounts[1];
    const seekerAddress = await seekerAccount.getAddress();

    const tokenId = 1;
    await seekers.mint(await accounts[2].getAddress(), tokenId);

    const nonce = randomBytes(32);
    const proofMessage = await registries.getProofMessage(
      tokenId,
      await accounts[0].getAddress(),
      nonce,
    );

    const signature = await seekerAccount.signMessage(
      Buffer.from(proofMessage.slice(2), 'hex'),
    );

    await expect(
      registries.setSeekerAccount(seekerAddress, tokenId, nonce, signature),
    ).to.be.revertedWithCustomError(registries, 'SeekerAccountMustOwnSeekerId');
  });

  it('can revoke seeker account', async () => {
    const seekerAccount = accounts[1];

    await utils.setSeekerRegistry(
      registries,
      seekers,
      accounts[0],
      accounts[1],
      1,
    );

    await registries.connect(seekerAccount).revokeSeekerAccount(owner);

    const registry = await registries.getRegistry(owner);

    expect(registry.seekerAccount).to.equal(ethers.constants.AddressZero);
  });

  it('can only revoke seeker account as seeker account', async () => {
    await utils.setSeekerRegistry(
      registries,
      seekers,
      accounts[0],
      accounts[1],
      1,
    );

    await expect(
      registries.revokeSeekerAccount(owner),
    ).to.be.revertedWithCustomError(registries, 'SeekerAccountMustBeMsgSender');
  });

  it('requires registry to not have empty public endpoint string', async () => {
    await expect(registries.register('')).to.be.revertedWithCustomError(
      registries,
      'PublicEndpointCannotBeEmpty',
    );
  });

  it('registered node seeker id will be reset if a new node registers with the same seeker id', async () => {
    const seekerAccount = accounts[2];

    const accountOne = accounts[3];
    const accountAddressOne = await accountOne.getAddress();

    const accountTwo = accounts[4];
    const accountAddressTwo = await accountTwo.getAddress();

    const tokenID = 100;

    await utils.setSeekerRegistry(
      registries,
      seekers,
      accountOne,
      seekerAccount,
      tokenID,
    );

    await utils.setSeekerRegistry(
      registries,
      seekers,
      accountTwo,
      seekerAccount,
      tokenID,
    );

    const regoSeekerAccountOne = await registries.getRegistry(
      accountAddressOne,
    );
    const regoSeekerAccountTwo = await registries.getRegistry(
      accountAddressTwo,
    );

    // tests that the registry for both seeker accounts don't have the same seekerID
    expect(regoSeekerAccountOne.seekerId).to.equal(0);
    expect(regoSeekerAccountTwo.seekerId).is.equal(tokenID);
  });

  it('has the correct prefix message', async () => {
    const lineOne =
      "🤖 Hi frend! 🤖\n\n📜 Signing this message proves that you're the owner of this Seeker NFT and allows your Seeker to be used to operate your Seeker's Node. It's a simple but important step to ensure smooth operation.\n\nThis request will not trigger a blockchain transaction or cost any gas fees.\n\n🔥 Your node's address: ";
    const lineTwo = '\n\n🆔 Your seeker id: ';
    const lineThree =
      '\n\n📦 A unique random value which secures this message: ';

    const account = accounts[0];
    const accountAddress = await account.getAddress();
    const tokenId = 100;
    const nonce = randomBytes(32);

    const proofMessageHexString = await registries.getProofMessage(
      tokenId,
      await account.getAddress(),
      nonce,
    );

    const proofMessage = `${lineOne}${accountAddress.toLowerCase()}${lineTwo}${tokenId}${lineThree}${
      '0x' + nonce.toString('hex')
    }`;

    const proofMessageString = Buffer.from(
      proofMessageHexString.slice(2),
      'hex',
    ).toString('utf8');

    expect(proofMessageString).to.equal(proofMessage);
  });
});
