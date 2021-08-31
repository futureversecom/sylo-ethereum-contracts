const Token = artifacts.require("SyloToken");

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

exports.advanceBlock = async function() {
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
