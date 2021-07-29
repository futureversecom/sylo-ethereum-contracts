const BN = require("bn.js");
const PriceManager = artifacts.require("PriceManager");
const PriceVoting = artifacts.require("PriceVoting");
const StakingManager = artifacts.require("StakingManager");
const Token = artifacts.require("SyloToken");

const utils = require('../utils');

contract('Pricing', accounts => {
  let priceManager;
  let priceVoting;
  let token;
  let stakingManager;

  before(async () => {
    token = await Token.new({ from: accounts[1] });
  });

  beforeEach(async () => {
    stakingManager = await StakingManager.new({ from: accounts[1] });
    await stakingManager.initialize(token.address, 0, { from: accounts[1] });

    priceVoting = await PriceVoting.new({ from: accounts[1] });
    await priceVoting.initialize(
      stakingManager.address,
      { from: accounts[1] }
    )

    priceManager = await PriceManager.new({ from: accounts[1] });
    await priceManager.initialize(
      stakingManager.address,
      priceVoting.address,
      { from: accounts[1] }
    );

    await token.approve(stakingManager.address, 10000, { from: accounts[1] });
  });

  it('allows a staked account to vote', async () => {
    await priceVoting.vote(1, { from: accounts[1] });
  });

  it('allows updating a vote', async () => {
    await priceVoting.vote(1, { from: accounts[1] });
    await priceVoting.vote(2, { from: accounts[1] });
  });

  it('allows withdrawing a vote', async () => {
    await priceVoting.vote(1, { from: accounts[1] });
    await priceVoting.withdraw({ from: accounts[1] });

    const price = await priceVoting.votes(accounts[1]);

    assert.equal(price.toNumber(), 0, "Expected vote to equal 0");
  });

  it('prevents vote with value set to 0', async () => {
    await stakingManager.addStake(1, accounts[1], { from: accounts[1] });
    await priceVoting.vote(0, { from: accounts[1] })
      .then(() => {
        assert.fail('Voting should fail if price is 0');
      })
      .catch(e => {
        assert.include(e.message, 'Voting price must be greater than 0');
      })
  });

  it('can determince price after votes and staking has occured', async () => {
    for (let i = 0; i < accounts.length; i++) {
      await stakingManager.addStake(10, accounts[i], { from: accounts[1] });
    }

    const prices = [20, 10, 5, 30, 45, 12, 17, 9, 24, 10];
    
    for (let i = 0; i < accounts.length; i++) {
      await priceVoting.vote(prices[i], { from: accounts[i] });
    }

    await utils.calculatePrices(priceManager, priceVoting, accounts[1]);

    const currentPrice = await priceManager.currentServicePrice();

    // Each stake is 10, and the total stake is 100. The lower boundary is 25,
    // so the boundary will be crossed when getting to the third sorted vote
    assert.equal(currentPrice.toNumber(), 10, "Calculated price does not match expected price");
  });

  it('nodes without a stake are considered in price calculation', async () => {
    for (let i = 0; i < 3; i++) {
      await stakingManager.addStake(10, accounts[i], { from: accounts[1] });
    }

    const prices = [2, 3, 4, 1];
    
    for (let i = 0; i < 4; i++) {
      await priceVoting.vote(prices[i], { from: accounts[i] });
    }

    await utils.calculatePrices(priceManager, priceVoting, accounts[1]);

    const currentPrice = await priceManager.currentServicePrice();

    // If the last node with vote 1 was counted in price, the price would be
    // 1, but if not, then it should be 2
    assert.equal(currentPrice.toNumber(), 2, "Calculated price does not match expected price");
  });

  async function voteAndStake(stakee, vote, stake, account) {
    const tokenContract = new web3.eth.Contract(Token.abi, token.address);
    await utils.sendContractTransaction(
      token.address,
      tokenContract.methods.approve(stakingManager.address, 10000),
      account
    );

    const stakingContract = new web3.eth.Contract(StakingManager.abi, stakingManager.address);
    await utils.sendContractTransaction(
      stakingManager.address,
      stakingContract.methods.addStake(stake, stakee),
      account
    );

    const votingContract = new web3.eth.Contract(PriceVoting.abi, priceVoting.address);
    await utils.sendContractTransaction(
      priceVoting.address,
      votingContract.methods.vote(vote),
      account
    );
  }
});
