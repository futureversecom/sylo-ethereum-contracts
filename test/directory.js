const Directory = artifacts.require("Directory");
const Token = artifacts.require("SyloToken");

contract('Directory', accounts => {
  let token;
  let directory;

  before(async () => {
    token = await Token.new({ from: accounts[1] });
  });

  beforeEach(async () => {
    directory = await Directory.new(token.address, 0, { from: accounts[0] });

    await token.approve(directory.address, 10000, { from: accounts[1] });
  });

  it('should be able to unstake the root', async () => {
    await directory.addStake(1, accounts[1], { from: accounts[1] });
    await directory.addStake(1, accounts[2],{ from: accounts[1] });

    const totalStake = await directory.getTotalStake();
    assert.equal(totalStake, 2, "Expected a total stake of 2");

    await directory.unlockStake(1, accounts[1], { from: accounts[1] });

    const totalStake2 = await directory.getTotalStake();
    assert.equal(totalStake2.toNumber(), 1, "Expected a total stake of 1");
  });

  it('should be able to unstake a leaf', async () => {
    await directory.addStake(1, accounts[1], { from: accounts[1] });
    await directory.addStake(1, accounts[2],{ from: accounts[1] });

    const totalStake = await directory.getTotalStake();
    assert.equal(totalStake, 2, "Expected a total stake of 2");

    await directory.unlockStake(1, accounts[2], { from: accounts[1] });

    const totalStake2 = await directory.getTotalStake();
    assert.equal(totalStake2.toNumber(), 1, "Expected a total stake of 1");
  });

  it('should be able to unstake a node with parent and child', async () => {
    for (let i = 0; i < accounts.length; i++) {
      await directory.addStake(1, accounts[i],{ from: accounts[1] });
    }

    const totalStake = await directory.getTotalStake();
    assert.equal(totalStake, accounts.length, "Expected a total stake of 2");

    await directory.unlockStake(1, accounts[2], { from: accounts[1] });

    const totalStake2 = await directory.getTotalStake();
    assert.equal(totalStake2.toNumber(), accounts.length - 1, "Expected a total stake of 1");
  });
});