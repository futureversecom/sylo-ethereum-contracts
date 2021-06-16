const BN = require("bn.js");
const Token = artifacts.require("SyloToken");
const Ticketing = artifacts.require("SyloTicketing");
const eth = require('eth-lib');
const { soliditySha3, encodePacked } = require("web3-utils");

contract('Ticketing', accounts => {
  let token;
  let ticketing;
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
    ticketing = await Ticketing.new(token.address, 0, { from: accounts[0] });

    await token.approve(ticketing.address, 10000, { from: accounts[1] });
  });

  it('can redeem winning ticket', async () => {
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
    await ticketing.depositEscrow(5, accounts[0], { from: accounts[1] });
    await ticketing.depositPenalty(50, accounts[0], { from: accounts[1] });

    const { ticket, receiverRand, signature } = 
      await createWinningTicket(0, 1);

    const initialReceiverBalance = await token.balanceOf(accounts[1]);
    await ticketing.redeem(ticket, receiverRand, signature, { from: accounts[1] });

    const deposit = await ticketing.deposits.call(accounts[0]);
    assert.equal(deposit.escrow.toString(), '0', 'Expected entire escrow to be used');
    assert.equal(deposit.penalty.toString(), '45', 'Expected remaining portion of ticket face value to be substracted from penalty');

    const postRedeemBalance = await token.balanceOf(accounts[1]);
    assert.equal(
      postRedeemBalance.toString(),
      initialReceiverBalance.add(new BN(5)).toString(),
      "Expected balance of receiver to only have remaining available escrow added to it"
    );
  });

  async function createWinningTicket(sender, receiver) {
    const receiverRand = 1;
    const receiverRandHash = soliditySha3(receiverRand);

    const ticket = {
      sender: accounts[sender],
      receiver: accounts[receiver],
      faceValue: 10,
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