const BN = require("bn.js");
const Token = artifacts.require("SyloToken");
const crypto = require("crypto");
const sodium = require('libsodium-wrappers-sumo');
const Ticketing = artifacts.require("SyloTicketing");
const Listings = artifacts.require("Listings");
const StakingManager = artifacts.require("StakingManager");
const eth = require('eth-lib');
const { soliditySha3 } = require("web3-utils");

contract('Ticketing', accounts => {
  const faceValue = 15;

  const baseLiveWinProb = (new BN(2)).pow(new BN(256)).sub(new BN(1)).toString();
  const expiredWinProb = 1000;
  const decayRate = 80;
  const ticketDuration = 100;

  let token;
  let ticketing;
  let listings;
  let stakingManager;
  // private keys generated from default truffle mnemonic
  // use these to sign tickets
  const privateKeys =
    [ '0xc87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3',
      '0xae6ae8e5ccbfb04590405997ee2d52d2b330726137b875053c36d94e974d162f',
      '0x0dbbe8e4ae425a6d2687f1a7e3ba17bc98c673636790f1b8ad91193c05875ef1',
      '0xc88b703fb08cbea894b6aeff5a544fb92e78a18e19814cd85da83b71f772aa6c',
      '0x388c684f0ba1ef5017716adb5d21a053ea8e90277d0868337519f97bede61418',
      '0x659cbb0e2411a44db63778987b1e22153c086a95eb6b18bdf89de078917abc63',
      '0x82d052c865f5763aad42add438569276c00d3d88a2d062d36b2bae914d58b8c8',
      '0xaa3680d5d48a8283413f7a108367c7299ca73f553735860a87b08f39395618b7',
      '0x0f62d96d6675f32685bbdb8ac13cda7c23436f63efbb9d07700d8669ff12b7c4',
      '0x8d5366123cb560bb606379f90a0bfd4769eecc0557f1b362dcae9012b548b1e5]'
    ]

  before(async () => {
    token = await Token.new({ from: accounts[1] });
  });

  beforeEach(async () => {
    listings = await Listings.new({ from: accounts[1] });
    await listings.initialize(50), { from: accounts[1] };

    stakingManager = await StakingManager.new({ from: accounts[1] });
    await stakingManager.initialize(token.address, 0, { from: accounts[1] });

    ticketing = await Ticketing.new({ from: accounts[1] });
    await ticketing.initialize(
      token.address, 
      listings.address, 
      stakingManager.address, 
      0,
      faceValue,
      baseLiveWinProb,
      expiredWinProb,
      decayRate,
      ticketDuration,
      { from: accounts[1] }
    );
    await token.approve(ticketing.address, 10000, { from: accounts[1] });
    await token.approve(stakingManager.address, 10000, { from: accounts[1] });
  });

  it('should be able to deposit escrow', async () => {
    await ticketing.depositEscrow(50, accounts[0], { from: accounts[1] });

    const deposit = await ticketing.deposits(accounts[0]);
    assert.equal(deposit.escrow.toString(), '50', 'Expected 50 in escrow');
  });

  it('should be able to deposit penalty', async () => {
    await ticketing.depositPenalty(50, accounts[0], { from: accounts[1] });

    const deposit = await ticketing.deposits(accounts[0]);
    assert.equal(deposit.penalty.toString(), '50', 'Expected 50 in escrow');
  });

  it('should be able to deposit escrow multiple times', async () => {
    await ticketing.depositEscrow(50, accounts[0], { from: accounts[1] });
    await ticketing.depositEscrow(50, accounts[0], { from: accounts[1] });

    const deposit = await ticketing.deposits(accounts[0]);
    assert.equal(deposit.escrow.toString(), '100', 'Expected 100 in escrow');
  });

  it('should be able to depost to penalty multiple times', async () => {
    await ticketing.depositPenalty(50, accounts[0], { from: accounts[1] });
    await ticketing.depositPenalty(50, accounts[0], { from: accounts[1] });

    const deposit = await ticketing.deposits(accounts[0]);
    assert.equal(deposit.penalty.toString(), '100', 'Expected 100 in penalty');
  });

  it('should fail to withdraw without unlocking', async () => {
    await ticketing.depositEscrow(50, accounts[0], { from: accounts[1] });

    await ticketing.withdraw({ from: accounts[0] })
      .then(() => {
        assert.fail('Withdrawing should fail');
      })
      .catch(e => {
        assert.include(e.message, 'Deposits not unlocked', 'Withdraw should fail due to not being unlocked');
      });
  });

  it('should fail to unlock without deposit', async () => {
    await ticketing.unlockDeposits({ from: accounts[0] })
      .then(() => {
        assert.fail('Withdrawing should fail');
      })
      .catch(e => {
        assert.include(e.message, 'Nothing to withdraw', 'Unlock should fail due to no deposit');
      });
  });
  
  it('should be able to unlock', async () => {
    await ticketing.depositEscrow(50, accounts[0], { from: accounts[1] });
    await ticketing.unlockDeposits({ from: accounts[0] });

    const deposit = await ticketing.deposits(accounts[0]);
    assert.isAbove(deposit.unlockAt.toNumber(), 0, 'Expected deposit to go into unlocking phase');
  });

  it('should fail to unlock if already unlocking', async () => {
    await ticketing.depositEscrow(50, accounts[0], { from: accounts[1] });
    await ticketing.unlockDeposits({ from: accounts[0] });

    await ticketing.unlockDeposits({ from: accounts[0] })
      .then(() => {
        assert.fail('Withdrawing should fail');
      })
      .catch(e => {
        assert.include(e.message, 'Unlock already in progress', 'Unlock should fail due to already unlocking');
      });
  });

  it('should fail to lock if already locked', async () => {
    await ticketing.depositEscrow(50, accounts[0], { from: accounts[1] });
    await ticketing.lockDeposits({ from: accounts[0] })
      .then(() => {
        assert.fail('Locking should fail');
      })
      .catch(e => {
        assert.include(e.message, 'Not unlocking, cannot lock', 'Expect lock to fail as it deposit is already locked');
      });
  });

  it('should be able to lock deposit while it is unlocked', async () => {
    await ticketing.depositEscrow(50, accounts[0], { from: accounts[1] });
    await ticketing.unlockDeposits({ from: accounts[0] });
    await ticketing.lockDeposits({ from: accounts[0] });

    const deposit = await ticketing.deposits(accounts[0]);
    assert.equal(deposit.unlockAt.toString(), '0', 'Expected deposit to go into unlocking phase');
  });

  it('can not redeem ticket with invalid signature', async () => {
    const { ticket, senderRand, redeemerRand } = 
      await createWinningTicket(0, 1);

    const signature = '0x00';

    await ticketing.redeem(ticket, senderRand, redeemerRand, signature, { from: accounts[1] })
      .then(() => {
        assert.fail('Should fail to redeem ticket with invalid signature');
      })
      .catch(e => {
        assert.include(e.message, 'ECDSA: invalid signature length', 'Expected redeeming to fail due to invalid signature');
      });
  });

  it('can not redeem ticket with invalid sender rand', async () => {
    const { ticket, redeemerRand, signature } = 
      await createWinningTicket(0, 1);

    const senderRand = 999;

    await ticketing.redeem(ticket, senderRand, redeemerRand, signature, { from: accounts[1] })
      .then(() => {
        assert.fail('Should fail to redeem ticket with invalid redeemer rand');
      })
      .catch(e => {
        assert.include(e.message, 'Hash of senderRand doesn\'t match senderRandHash', 'Expected redeeming to fail due to invalid signature');
      });
  });

  it('can not redeem ticket with invalid redeemer rand', async () => {
    const { ticket, senderRand, signature } = 
      await createWinningTicket(0, 1);

    const redeemerRand = 999;

    await ticketing.redeem(ticket, senderRand, redeemerRand, signature, { from: accounts[1] })
      .then(() => {
        assert.fail('Should fail to redeem ticket with invalid redeemer rand');
      })
      .catch(e => {
        assert.include(e.message, 'Hash of redeemerRand doesn\'t match redeemerRandHash', 'Expected redeeming to fail due to invalid signature');
      });
  });

  it('can redeem winning ticket', async () => {
    // simulate having account[1] as a node with a listing and a
    // stakingManager entry
    await stakingManager.addStake(1, accounts[1], { from: accounts[1] });
    await listings.setListing("0.0.0.0/0", 1, { from: accounts[1] });

    await ticketing.depositEscrow(50, accounts[0], { from: accounts[1] });
    await ticketing.depositPenalty(50, accounts[0], { from: accounts[1] });

    
    const { ticket, senderRand, redeemerRand, signature } = 
    await createWinningTicket(0, 1);
    
    const initialReceiverBalance = await token.balanceOf(accounts[1]);
    await ticketing.redeem(ticket, senderRand, redeemerRand, signature, { from: accounts[1] });

    const deposit = await ticketing.deposits.call(accounts[0]);
    assert.equal(deposit.escrow.toString(), '35', 'Expected ticket payout to be substracted from escrow');
    assert.equal(deposit.penalty.toString(), '50', 'Expected penalty to not be changed');

    const postRedeemBalance = await token.balanceOf(accounts[1]);
    assert.equal(
      postRedeemBalance.toString(),
      initialReceiverBalance.add(new BN(faceValue)).toString(),
      "Expected balance of redeemer to have added the ticket face value"
    );
  });

  it('burns penalty on insufficient escrow', async () => {
    // simulate having account[1] as a node with a listing and a
    // stakingManager entry
    await stakingManager.addStake(1, accounts[1], { from: accounts[1] });
    await listings.setListing("0.0.0.0/0", 1, { from: accounts[1] });

    await ticketing.depositEscrow(5, accounts[0], { from: accounts[1] });
    await ticketing.depositPenalty(50, accounts[0], { from: accounts[1] });

    const { ticket, senderRand, redeemerRand, signature } = 
      await createWinningTicket(0, 1);

    const initialReceiverBalance = await token.balanceOf(accounts[1]);
    await ticketing.redeem(ticket, senderRand, redeemerRand, signature, { from: accounts[1] });

    const deposit = await ticketing.deposits.call(accounts[0]);
    assert.equal(deposit.escrow.toString(), '0', 'Expected entire escrow to be used');
    assert.equal(deposit.penalty.toString(), '0', 'Expected entire penalty to b burned');

    const postRedeemBalance = await token.balanceOf(accounts[1]);
    assert.equal(
      postRedeemBalance.toString(),
      initialReceiverBalance.add(new BN(5)).toString(),
      "Expected balance of redeemer to only have remaining available escrow added to it"
    );

    const ticketingBalance = await token.balanceOf(ticketing.address);
    assert.equal(
      ticketingBalance.toString(),
      '0',
      'Expected all tokens to be transferred'
    );
  });

  it('should payout delegated stakers on redeeming ticket', async () => {
    await token.transfer(accounts[2], 1000, { from: accounts[1]} );
    await token.approve(stakingManager.address, 1000, { from: accounts[2] });

    // have account[2] as a delegated staker
    await stakingManager.addStake(1, accounts[1], { from: accounts[2] });
    await listings.setListing("0.0.0.0/0", 1, { from: accounts[1] });

    await ticketing.depositEscrow(50, accounts[0], { from: accounts[1] });
    await ticketing.depositPenalty(50, accounts[0], { from: accounts[1] });

    const { ticket, senderRand, redeemerRand, signature } = 
      await createWinningTicket(0, 1);

    const initialDelegatorBalance = await token.balanceOf(accounts[2]);
    const initialReceiverBalance = await token.balanceOf(accounts[1]);

    await ticketing.redeem(ticket, senderRand, redeemerRand, signature, { from: accounts[1] });

    // The payout percentage is 50%, so the node (accounts[1]) and the
    // only delegator (account[2]) should split the reward equally
    const expectedPayout = faceValue / 2;

    const postDelegatorBalance = await token.balanceOf(accounts[2]);
    const postReceiverBalance = await token.balanceOf(accounts[1]);

    assert.equal(
      postDelegatorBalance.toString(),
      initialDelegatorBalance.add(new BN(expectedPayout)).toString(),
      "Expected balance of delegator to have added 50% of ticket face value"
    );

    assert.equal(
      postReceiverBalance.toString(),
      initialReceiverBalance.add(new BN(expectedPayout + 1 /* add 1 due to rounding  */)).toString(),
      "Expected balance of redeemer to have added 50% of ticket face value"
    );
  });

  it('should pay stakee remainders left from rounding down', async () => {
    // have account 2 and 3 as delegated stakers
    for (let i = 2; i < 4; i++) {
      await token.transfer(accounts[i], 1000, { from: accounts[1]} );
      await token.approve(stakingManager.address, 1000, { from: accounts[i] });
      await stakingManager.addStake(1, accounts[1], { from: accounts[i] });
    }

    await listings.setListing("0.0.0.0/0", 1, { from: accounts[1] });

    await ticketing.depositEscrow(50, accounts[0], { from: accounts[1] });
    await ticketing.depositPenalty(50, accounts[0], { from: accounts[1] });

    const { ticket, senderRand, redeemerRand, signature } = 
      await createWinningTicket(0, 1);

    const initialDelegatorBalance = await token.balanceOf(accounts[2]);
    const initialReceiverBalance = await token.balanceOf(accounts[1]);

    await ticketing.redeem(ticket, senderRand, redeemerRand, signature, { from: accounts[1] });

    // The payout percentage is 50%, so due to rounding, 7 will go to
    // the delegators, and 8 will go directly to the node. For the value
    // split amongst the delegators (accounts 2 and 3), it will be split
    // proportionally but rounded down, so both will receive 3, and the remainder (1)
    // will be sent to the node
    const expectedDelegatorsPayout = parseInt(faceValue / 2);
    const expectedStakeePayout = faceValue - expectedDelegatorsPayout + 1;

    const postDelegatorBalance = await token.balanceOf(accounts[2]);
    const postReceiverBalance = await token.balanceOf(accounts[1]);

    assert.equal(
      postDelegatorBalance.toString(),
      initialDelegatorBalance.add(new BN(parseInt(expectedDelegatorsPayout / 2))).toString(),
      "Expected balance of delegator to have 3 added to their balance"
    );

    assert.equal(
      postReceiverBalance.toString(),
      initialReceiverBalance.add(new BN(expectedStakeePayout)).toString(),
      "Expected balance of redeemer to have 9 added to their balance"
    );
  });

  it('should payout full ticket face value to node if all delegates pull out', async () => {
    await token.transfer(accounts[2], 1000, { from: accounts[1]} );
    await token.approve(stakingManager.address, 1000, { from: accounts[2] });

    // have both the node and account[2] stake
    await stakingManager.addStake(1, accounts[1], { from: accounts[2] });
    await stakingManager.addStake(1, accounts[1], { from: accounts[1] });
    await listings.setListing("0.0.0.0/0", 1, { from: accounts[1] });

    await ticketing.depositEscrow(50, accounts[0], { from: accounts[1] });
    await ticketing.depositPenalty(50, accounts[0], { from: accounts[1] });

    // unlock account[2] stake
    await stakingManager.unlockStake(1, accounts[1], { from: accounts[2] });

    const initialDelegatorBalance = await token.balanceOf(accounts[2]);
    const initialReceiverBalance = await token.balanceOf(accounts[1]);

    const { ticket, senderRand, redeemerRand, signature } = 
      await createWinningTicket(0, 1);

    await ticketing.redeem(ticket, senderRand, redeemerRand, signature, { from: accounts[1] });

    const postDelegatorBalance = await token.balanceOf(accounts[2]);
    const postReceiverBalance = await token.balanceOf(accounts[1]);

    assert.equal(
      postDelegatorBalance.toString(),
      initialDelegatorBalance.toString(),
      "Expected balance of delegator to have added 50% of ticket face value"
    );

    assert.equal(
      postReceiverBalance.toString(),
      initialReceiverBalance.add(new BN(faceValue)).toString(),
      "Expected balance of redeemer to have added 50% of ticket face value"
    );
  });

  it('gas cast should not be excessive at maximum delegators', async () => {
    // reach 10 delegator limit
    for (let i = 0; i < accounts.length; i++) {
      if (i == 1) {
        await stakingManager.addStake(1, accounts[1], { from: accounts[1] });
      } else {
        await token.transfer(accounts[i], 1000, { from: accounts[1]} );
        await token.approve(stakingManager.address, 10000, { from: accounts[i] });
        await stakingManager.addStake(1, accounts[1], { from: accounts[i] });
      }
    }
    
    await listings.setListing("0.0.0.0/0", 1, { from: accounts[1] });

    await ticketing.depositEscrow(50, accounts[0], { from: accounts[1] });
    await ticketing.depositPenalty(50, accounts[0], { from: accounts[1] });

    const { ticket, senderRand, redeemerRand, signature } = 
      await createWinningTicket(0, 1);

    await ticketing.redeem.estimateGas(
      ticket, 
      senderRand,
      redeemerRand, 
      signature, 
      { from: accounts[1],
        gas: 350000,
      }
    );
  });

  it('should decay winning probability as ticket approaches expiry', async () => {
    await stakingManager.addStake(1, accounts[1], { from: accounts[1] });
    await listings.setListing("0.0.0.0/0", 1, { from: accounts[1] });

    // deploy another ticketing contract with simpler parameters
    ticketing = await Ticketing.new({ from: accounts[1] });
    await ticketing.initialize(
      token.address, 
      listings.address, 
      stakingManager.address, 
      0,
      faceValue,
      100000,
      1000,
      80,
      100,
      { from: accounts[1] }
    );
    await token.approve(ticketing.address, 10000, { from: accounts[1] });

    await ticketing.depositEscrow(50, accounts[0], { from: accounts[1] });
    await ticketing.depositPenalty(50, accounts[0], { from: accounts[1] });

    const { ticket } = 
      await createWinningTicket(0, 1);

    // advance the block halfway to ticket expiry
    for (let i = 0; i < 51; i++) {
      await advanceBlock();
    }

    // check if the probability has decayed 50% of the maximum decayed value (80%)
    const expectedProbability = new BN(100000 - (0.5 * 0.8 * 100000));

    const decayedProbability = await ticketing.calculateWinningProbability(ticket);

    assert.equal(
      decayedProbability.toString(), 
      expectedProbability.toString(), 
      "Expected probablity of ticket winning to decay"
    );
  });

  it('returns 0 winning probability if ticket has expired', async () => {
    await stakingManager.addStake(1, accounts[1], { from: accounts[1] });
    await listings.setListing("0.0.0.0/0", 1, { from: accounts[1] });

    const { ticket } = 
      await createWinningTicket(0, 1);

    // advance the block all the way to ticket expiry
    for (let i = 0; i < 101; i++) {
      await advanceBlock();
    }

    const p = await ticketing.calculateWinningProbability(ticket);

    assert.equal(
      '0',
      p.toString(),
      'Expected probabilit to be 0'
    );
  });

  it('simulates scenario between sender, node, and oracle', async () => {
    const sender = accounts[0];
    const node = accounts[1];

    // set up the node's stake and listing
    await stakingManager.addStake(1, node, { from: accounts[1] });
    await listings.setListing("0.0.0.0/0", 1, { from: accounts[1] });

    // set up the sender's escrow
    await ticketing.depositEscrow(50, sender, { from: accounts[1] });
    await ticketing.depositPenalty(50, sender, { from: accounts[1] });

    // have the node and sender generate random numbers
    const nodeRand = crypto.randomBytes(32);
    const senderRand = crypto.randomBytes(32);

    // create commits from those random numbers
    const nodeCommit = soliditySha3(nodeRand);
    const senderCommit = soliditySha3(senderRand);

    // create the ticket to be given to the node
    const ticket = {
      sender: accounts[0],
      redeemer: accounts[1],
      generationBlock: new BN((await web3.eth.getBlockNumber()) + 1).toString(),
      senderCommit,
      redeemerCommit: nodeCommit
    };

    // have sender sign the hash of the ticket
    const ticketHash = await ticketing.getTicketHash(ticket);
    const signature = eth.Account.sign(ticketHash, privateKeys[0]);

    // establish the oracle
    const oracle = sodium.crypto_sign_keypair('uint8array');

    // encrypt senderRandom to create the key
    const key = sodium.crypto_box_seal(
      senderRand,
      sodium.crypto_sign_ed25519_pk_to_curve25519(oracle.publicKey),
      'uint8array'
    );

    // have oracle decrypt the key and reveal the random number to the node
    const revealedSenderRand = sodium.crypto_box_seal_open(
      key,
      sodium.crypto_sign_ed25519_pk_to_curve25519(oracle.publicKey),
      sodium.crypto_sign_ed25519_sk_to_curve25519(oracle.privateKey),
      'uint8array'
    );

    // once secret has been revealed, the node can now redeem the ticket
    await ticketing.redeem(
      ticket,
      revealedSenderRand,
      nodeRand,
      signature,
      { from: node }
    );
  });


  async function advanceBlock() {
    return await new Promise((resolve, reject) => {
      web3.currentProvider.send({
        jsonrpc: '2.0',
        method: 'evm_mine',
        id: new Date().getTime()
      }, (err, result) => {
        if (err) { return reject(err) }
        const newBlockHash = web3.eth.getBlock('latest').hash
  
        return resolve(newBlockHash)
      })
    })
  }

  async function createWinningTicket(sender, redeemer) {
    const senderRand = 1;
    const senderCommit = soliditySha3(senderRand);

    const redeemerRand = 1;
    const redeemerCommit = soliditySha3(redeemerRand);

    const generationBlock = await web3.eth.getBlockNumber();

    const ticket = {
      sender: accounts[sender],
      redeemer: accounts[redeemer],
      generationBlock: new BN(generationBlock + 1).toString(),
      senderCommit,
      redeemerCommit
    };

    const ticketHash = await ticketing.getTicketHash(ticket);

    const signature = eth.Account.sign(ticketHash, privateKeys[sender]);

    return { ticket, senderRand, redeemerRand, signature, ticketHash }
  }

});
