const BN = require("bn.js");
const Token = artifacts.require("SyloToken");
const TicketingParameters = artifacts.require('TicketingParameters');
const Ticketing = artifacts.require("SyloTicketing");
const RewardsManager = artifacts.require("RewardsManager");
const EpochsManager = artifacts.require("EpochsManager");
const Directory = artifacts.require("Directory");
const Listings = artifacts.require("Listings");
const StakingManager = artifacts.require("StakingManager");

exports.initializeContracts = async function(deployer, tokenAddress, opts = {}) {
  const payoutPercentage =
    opts.payoutPercentage ?
      opts.payoutPercentage :
      5000;

  const faceValue = opts.faceValue ? opts.faceValue : 15;
  const baseLiveWinProb =
    opts.baseLiveWinProb ?
      opts.baseLiveWinProb :
      (new BN(2)).pow(new BN(128)).sub(new BN(1)).toString();
  const expiredWinProb = opts.expiredWinProb ? opts.expiredWinProb : 1000;
  const decayRate = opts.decayRate ? opts.decayRate : 8000;
  const ticketDuration = opts.ticketDuration ? opts.ticketDuration : 20;

  const epochDuration = opts.epochDuration ? opts.epochDuration : 30;

  const listings = await Listings.new({ from: deployer });
  await listings.initialize(payoutPercentage), { from: deployer };

  const ticketingParameters = await TicketingParameters.new({ from: deployer });
  await ticketingParameters.initialize(
    faceValue,
    baseLiveWinProb,
    expiredWinProb,
    decayRate,
    ticketDuration,
    { from: deployer }
  );

  const epochsManager = await EpochsManager.new({ from: deployer });
  const stakingManager = await StakingManager.new({ from: deployer });
  const rewardsManager = await RewardsManager.new({ from: deployer });
  const directory = await Directory.new({ from: deployer });

  await stakingManager.initialize(
    tokenAddress,
    rewardsManager.address,
    epochsManager.address,
    0, { from: deployer }
  );
  await rewardsManager.initialize(
    tokenAddress,
    stakingManager.address,
    epochsManager.address,
    { from: deployer }
  );
  await directory.initialize(
      stakingManager.address,
      rewardsManager.address,
    { from: deployer }
  );
  await epochsManager.initialize(
    directory.address,
    listings.address,
    ticketingParameters.address,
    epochDuration,
    { from: deployer }
  );

  const ticketing = await Ticketing.new({ from: deployer })
  await ticketing.initialize(
    tokenAddress,
    listings.address,
    stakingManager.address,
    directory.address,
    epochsManager.address,
    rewardsManager.address,
    0,
    { from: deployer }
  );

  await rewardsManager.addManager(ticketing.address, { from: deployer });
  await rewardsManager.addManager(stakingManager.address, { from: deployer });

  return {
    listings,
    ticketing,
    ticketingParameters,
    directory,
    rewardsManager,
    epochsManager,
    stakingManager
  }
}

exports.fundRandomAccount = async function(funder, tokenAddress) {
  const acc = web3.eth.accounts.create();

  const contract = new web3.eth.Contract(Token.abi, tokenAddress);

  // send sylos
  await contract.methods.transfer(acc.address, 1000000).send({ from: funder });
  const b = await contract.methods.balanceOf(acc.address).call();

  // send eth
  await web3.eth.sendTransaction({
    from: funder,
    to: acc.address,
    value: '100000000000000000'
  });

  return acc;
}

exports.sendContractTransaction = async function(contractAddress, method, account) {
  var rawTx = {
    to: contractAddress,
    data: method.encodeABI(),
    value: '0x0',
    from: account.address,
    nonce: await web3.eth.getTransactionCount(account.address),
    gas: 1000000
  }

  const signedTx = await web3.eth.accounts.signTransaction(rawTx, account.privateKey);
  await web3.eth.sendSignedTransaction(signedTx.rawTransaction);
}

exports.calculatePrices = async function(priceManager, priceVoting, owner) {
  const r = await priceVoting.getVotes();
  const sortedVotes = [];
  for (let i = 0; i < r['1'].length; i++) {
    sortedVotes.push({ p: r['1'][i], i });
  }
  sortedVotes.sort((a, b) => {
    return a.p - b.p;
  });
  const sortedIndexes = sortedVotes.map(x => x.i);
  await priceManager.calculatePrices(sortedIndexes, { from: owner });
}

exports.advanceBlock = async function(i) {
  i = i ? i : 1;
  for (let j = 0; j < i; j++) {
    await new Promise((resolve, reject) => {
      web3.currentProvider.send({
        jsonrpc: '2.0',
        method: 'evm_mine',
        id: new Date().getTime()
      }, (err, result) => {
        if (err) { return reject(err) }
        const newBlockHash = web3.eth.getBlock('latest').hash

        return resolve(newBlockHash)
      })
    });
  }
}
