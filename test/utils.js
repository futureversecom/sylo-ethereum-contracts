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