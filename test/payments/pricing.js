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
    await stakingManager.addStake(1, accounts[1], { from: accounts[1] });
    await priceVoting.vote(1, { from: accounts[1] });
  });

  it('allows updating a vote', async () => {
    await stakingManager.addStake(1, accounts[1], { from: accounts[1] });
    await priceVoting.vote(1, { from: accounts[1] });
    await priceVoting.vote(2, { from: accounts[1] });
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

    await priceManager.calculatePrices([], { from: accounts[1] });

    const currentPrice = await priceManager.currentServicePrice();

    // Each stake is 10, and the total stake is 100. The lower boundary is 25,
    // so the boundary will be crossed when getting to the third sorted vote
    assert.equal(currentPrice.toNumber(), 10, "Calculated price does not match expected price");
  });

  it('tests gas costs of calculating a 100 votes', async () => {
    const prices = [20, 10, 5, 30, 45, 12, 17, 9, 24, 10];
    
    for (let i = 0; i < 100; i++) {
      const account = await utils.fundRandomAccount(accounts[1], token.address);
      await voteAndStake(account.address, prices[i % prices.length], 1, account); 
    }

    // calculate the gas cost for sorting on chain
    const sortingGasCost = await priceManager.calculatePrices.estimateGas([], { from: accounts[1] });

    // calculate the gas cost for providing a sorted votes array
    const votes = await priceVoting.getVotes();
    const sortedVotes = [];
    for (let i = 0; i < votes['0'].length; i++) {
      sortedVotes.push({ voter: votes['0'][i], price: votes['1'][i] });
    }
    sortedVotes.sort((a, b) => a.price - b.price);
    const validatingGasCost = await priceManager.calculatePrices.estimateGas(sortedVotes, { from: accounts[1] });

    console.log(`Gas cost for calculating prices using on chain sort=${sortingGasCost}`);
    console.log(`Gas cost for calculating price using validating sorted array=${validatingGasCost}`)
  }).timeout(0);

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
