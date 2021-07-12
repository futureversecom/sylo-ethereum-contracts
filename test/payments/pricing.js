const BN = require("bn.js");
const PriceManager = artifacts.require("PriceManager");
const PriceVoting = artifacts.require("PriceVoting");
const StakingManager = artifacts.require("StakingManager");
const Token = artifacts.require("SyloToken");

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

  it('prevents non staked node from voting', async () => {
    await priceVoting.vote(1, { from: accounts[1] })
      .then(() => {
        assert.fail('Voting should fail without stake');
      })
      .catch(e => {
        assert.include(e.message, 'Must have stake to vote');
      });
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
      await stakingManager.addStake(10, accounts[i],{ from: accounts[1] });
    }

    const prices = [20, 10, 5, 30, 45, 12, 17, 9, 24, 10];
    
    for (let i = 0; i < accounts.length; i++) {
      await priceVoting.vote(prices[i], { from: accounts[i] });
    }

    let sortedVotes = [];
    for (let i = 0; i < accounts.length; i++) {
      sortedVotes.push({ voter: accounts[i], price: prices[i] });
    }
    sortedVotes.sort((a, b) => a.price - b.price);

    await priceManager.calculatePrices(sortedVotes, { from: accounts[1] });

    const currentPrice = await priceManager.currentServicePrice();

    // Each stake is 10, and the total stake is 100. The lower boundary is 25,
    // so the boundary will be crossed when getting to the third sorted vote
    const expectedPrice = sortedVotes[2].price;

    assert.equal(currentPrice.toNumber(), expectedPrice, "Calculated price does not match expected price");
  });
});