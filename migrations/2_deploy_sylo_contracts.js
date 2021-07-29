const { deployProxy } = require('@openzeppelin/truffle-upgrades');

const Token = artifacts.require('SyloToken');
const Listings = artifacts.require('Listings');
const Directory = artifacts.require('Directory');
const Ticketing = artifacts.require('SyloTicketing');

module.exports = async function (deployer) {
  if (deployer.network == 'develop') {
    await deployer.deploy(Token);
    const token = await Token.deployed();
    await deployProxy(Listings, [50], { deployer });
    await deployProxy(Directory, [token.address, 0], { deployer });
    await deployProxy(Ticketing, [token.address, 0], { deployer });
  }
};