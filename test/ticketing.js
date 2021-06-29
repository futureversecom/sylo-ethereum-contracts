const BN = require("bn.js");
const Token = artifacts.require("SyloToken");
const Ticketing = artifacts.require("SyloTicketing");
const Listings = artifacts.require("Listings");
const Directory = artifacts.require("Directory");
const eth = require('eth-lib');
const { soliditySha3 } = require("web3-utils");

contract('Ticketing', accounts => {
  let token;
  let ticketing;
  let listings;
  let directory;
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

    directory = await Directory.new({ from: accounts[1] });
    await directory.initialize(token.address, 0, { from: accounts[1] });

    ticketing = await Ticketing.new({ from: accounts[1] });
    await ticketing.initialize(
      token.address, 
      listings.address, 
      directory.address, 
      0, 
      { from: accounts[1] }
    );
    await token.approve(ticketing.address, 10000, { from: accounts[1] });
    await token.approve(directory.address, 10000, { from: accounts[1] });
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

  it('can not redeem expired ticket', async () => {
    const { ticket, receiverRand, signature } = 
      await createWinningTicket(0, 1);

    ticket.expirationBlock = 1;

    await ticketing.redeem(ticket, receiverRand, signature, { from: accounts[1] })
      .then(() => {
        assert.fail('Should fail to redeem expired ticket');
      })
      .catch(e => {
        assert.include(e.message, 'Ticket has expired', 'Expected redeeming to fail due to expired ticket');
      });
  });

  it('can not redeem ticket with invalid signature', async () => {
    const { ticket, receiverRand } = 
      await createWinningTicket(0, 1);

    const signature = '0x00';

    await ticketing.redeem(ticket, receiverRand, signature, { from: accounts[1] })
      .then(() => {
        assert.fail('Should fail to redeem ticket with invalid signature');
      })
      .catch(e => {
        assert.include(e.message, 'ECDSA: invalid signature length', 'Expected redeeming to fail due to invalid signature');
      });
  });

  it('can not redeem ticket with invalid receiver rand', async () => {
    const { ticket, signature } = 
      await createWinningTicket(0, 1);

    const receiverRand = 999;

    await ticketing.redeem(ticket, receiverRand, signature, { from: accounts[1] })
      .then(() => {
        assert.fail('Should fail to redeem ticket with invalid receiver rand');
      })
      .catch(e => {
        assert.include(e.message, 'Hash of receiverRand doesn\'t match receiverRandHash', 'Expected redeeming to fail due to invalid signature');
      });
  });

  it('can redeem winning ticket', async () => {
    // simulate having account[1] as a node with a listing and a
    // directory entry
    await directory.addStake(1, accounts[1], { from: accounts[1] });
    await listings.setListing("0.0.0.0/0", 1, { from: accounts[1] });

    await ticketing.depositEscrow(50, accounts[0], { from: accounts[1] });
    await ticketing.depositPenalty(50, accounts[0], { from: accounts[1] });

    const { ticket, receiverRand, signature } = 
      await createWinningTicket(0, 1);

    const initialReceiverBalance = await token.balanceOf(accounts[1]);
    await ticketing.redeem(ticket, receiverRand, signature, { from: accounts[1] });

    const deposit = await ticketing.deposits.call(accounts[0]);
    assert.equal(deposit.escrow.toString(), '40', 'Expected ticket face value to be substracted from escrow');
    assert.equal(deposit.penalty.toString(), '50', 'Expected penalty to not be changed');

    const postRedeemBalance = await token.balanceOf(accounts[1]);
    assert.equal(
      postRedeemBalance.toString(),
      initialReceiverBalance.add(new BN(ticket.faceValue)).toString(),
      "Expected balance of receiver to have added the ticket face value"
    );
  });

  it('burns penalty on insufficient escrow', async () => {
    // simulate having account[1] as a node with a listing and a
    // directory entry
    await directory.addStake(1, accounts[1], { from: accounts[1] });
    await listings.setListing("0.0.0.0/0", 1, { from: accounts[1] });

    await ticketing.depositEscrow(5, accounts[0], { from: accounts[1] });
    await ticketing.depositPenalty(50, accounts[0], { from: accounts[1] });

    const { ticket, receiverRand, signature } = 
      await createWinningTicket(0, 1);

    const initialReceiverBalance = await token.balanceOf(accounts[1]);
    await ticketing.redeem(ticket, receiverRand, signature, { from: accounts[1] });

    const deposit = await ticketing.deposits.call(accounts[0]);
    assert.equal(deposit.escrow.toString(), '0', 'Expected entire escrow to be used');
    assert.equal(deposit.penalty.toString(), '0', 'Expected entire penalty to b burned');

    const postRedeemBalance = await token.balanceOf(accounts[1]);
    assert.equal(
      postRedeemBalance.toString(),
      initialReceiverBalance.add(new BN(5)).toString(),
      "Expected balance of receiver to only have remaining available escrow added to it"
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
    await token.approve(directory.address, 1000, { from: accounts[2] });

    // have account[2] as a delegated staker
    await directory.addStake(1, accounts[1], { from: accounts[2] });
    await listings.setListing("0.0.0.0/0", 1, { from: accounts[1] });

    await ticketing.depositEscrow(50, accounts[0], { from: accounts[1] });
    await ticketing.depositPenalty(50, accounts[0], { from: accounts[1] });

    const { ticket, receiverRand, signature } = 
      await createWinningTicket(0, 1);

    const initialDelegatorBalance = await token.balanceOf(accounts[2]);
    const initialReceiverBalance = await token.balanceOf(accounts[1]);

    await ticketing.redeem(ticket, receiverRand, signature, { from: accounts[1] });

    // The payout percentage is 50%, so the node (accounts[1]) and the
    // only delegator (account[2]) should split the reward equally
    const expectedPayout = ticket.faceValue / 2;

    const postDelegatorBalance = await token.balanceOf(accounts[2]);
    const postReceiverBalance = await token.balanceOf(accounts[1]);

    assert.equal(
      postDelegatorBalance.toString(),
      initialDelegatorBalance.add(new BN(expectedPayout)).toString(),
      "Expected balance of delegator to have added 50% of ticket face value"
    );

    assert.equal(
      postReceiverBalance.toString(),
      initialReceiverBalance.add(new BN(expectedPayout)).toString(),
      "Expected balance of receiver to have added 50% of ticket face value"
    );
  });

  it('should pay stakee remainders left from rounding down', async () => {
    // have account 2 and 3 as delegated stakers
    for (let i = 2; i < 4; i++) {
      await token.transfer(accounts[i], 1000, { from: accounts[1]} );
      await token.approve(directory.address, 1000, { from: accounts[i] });
      await directory.addStake(1, accounts[1], { from: accounts[i] });
    }

    await listings.setListing("0.0.0.0/0", 1, { from: accounts[1] });

    await ticketing.depositEscrow(50, accounts[0], { from: accounts[1] });
    await ticketing.depositPenalty(50, accounts[0], { from: accounts[1] });

    const { ticket, receiverRand, signature } = 
      await createWinningTicket(0, 1, 15); // 15 ticket as face value

    const initialDelegatorBalance = await token.balanceOf(accounts[2]);
    const initialReceiverBalance = await token.balanceOf(accounts[1]);

    await ticketing.redeem(ticket, receiverRand, signature, { from: accounts[1] });

    // The payout percentage is 50%, so due to rounding, 7 will go to
    // the delegators, and 8 will go directly to the node. For the value
    // split amongst the delegators (accounts 2 and 3), it will be split
    // proportionally but rounded down, so both will receive 3, and the remainder (1)
    // will be sent to the node
    const expectedDelegatorsPayout = parseInt(ticket.faceValue / 2);
    const expectedStakeePayout = ticket.faceValue - expectedDelegatorsPayout + 1;

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
      "Expected balance of receiver to have 9 added to their balance"
    );
  });

  it('should payout full ticket face value to node if all delegates pull out', async () => {
    await token.transfer(accounts[2], 1000, { from: accounts[1]} );
    await token.approve(directory.address, 1000, { from: accounts[2] });

    // have both the node and account[2] stake
    await directory.addStake(1, accounts[1], { from: accounts[2] });
    await directory.addStake(1, accounts[1], { from: accounts[1] });
    await listings.setListing("0.0.0.0/0", 1, { from: accounts[1] });

    await ticketing.depositEscrow(50, accounts[0], { from: accounts[1] });
    await ticketing.depositPenalty(50, accounts[0], { from: accounts[1] });

    const { ticket, receiverRand, signature } = 
      await createWinningTicket(0, 1);

    // unlock account[2] stake
    await directory.unlockStake(1, accounts[1], { from: accounts[2] });

    const initialDelegatorBalance = await token.balanceOf(accounts[2]);
    const initialReceiverBalance = await token.balanceOf(accounts[1]);

    await ticketing.redeem(ticket, receiverRand, signature, { from: accounts[1] });

    const postDelegatorBalance = await token.balanceOf(accounts[2]);
    const postReceiverBalance = await token.balanceOf(accounts[1]);

    assert.equal(
      postDelegatorBalance.toString(),
      initialDelegatorBalance.toString(),
      "Expected balance of delegator to have added 50% of ticket face value"
    );

    assert.equal(
      postReceiverBalance.toString(),
      initialReceiverBalance.add(new BN(ticket.faceValue)).toString(),
      "Expected balance of receiver to have added 50% of ticket face value"
    );
  });

  it('gas cast should not be excessive at maximum delegators', async () => {
    // reach 10 delegator limit
    for (let i = 0; i < accounts.length; i++) {
      if (i == 1) {
        await directory.addStake(1, accounts[1], { from: accounts[1] });
      } else {
        await token.transfer(accounts[i], 1000, { from: accounts[1]} );
        await token.approve(directory.address, 10000, { from: accounts[i] });
        await directory.addStake(1, accounts[1], { from: accounts[i] });
      }
    }
    
    await listings.setListing("0.0.0.0/0", 1, { from: accounts[1] });

    await ticketing.depositEscrow(50, accounts[0], { from: accounts[1] });
    await ticketing.depositPenalty(50, accounts[0], { from: accounts[1] });

    const { ticket, receiverRand, signature } = 
      await createWinningTicket(0, 1);

    await ticketing.redeem.estimateGas(
      ticket, 
      receiverRand, 
      signature, 
      { from: accounts[1],
        gas: 350000,
      }
    );
  });

  async function createWinningTicket(sender, receiver, faceValue = 10) {
    const receiverRand = 1;
    const receiverRandHash = soliditySha3(receiverRand);

    const ticket = {
      sender: accounts[sender],
      receiver: accounts[receiver],
      faceValue,
      winProb: (new BN(2)).pow(new BN(256)).sub(new BN(1)).toString(), //max win prob
      expirationBlock: 0,
      receiverRandHash,
      senderNonce: 1
    }
    
    const ticketHash = soliditySha3(
      ticket.sender,
      ticket.receiver,
      ticket.faceValue,
      { t: 'uint256', v: ticket.winProb.toString() },
      ticket.expirationBlock,
      ticket.receiverRandHash,
      { t: 'uint32', v: ticket.senderNonce }
    );

    const signature = eth.Account.sign(ticketHash, privateKeys[sender]);

    return { ticket, receiverRand, signature, ticketHash }
  }

});