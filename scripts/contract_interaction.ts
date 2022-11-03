import { ethers as hardhatEthers } from 'hardhat';
import { Contract } from 'ethers';
import web3 from 'web3';
import { SyloToken } from '../typechain';

export async function main() {
  const [deployer] = await hardhatEthers.getSigners();
  console.log('deployer', deployer.address);

  const xrpTokenAddress = web3.utils.toChecksumAddress(
    '0xCCCCCCCC00000002000000000000000000000000',
  );
  const erc20Abi = [
    'event Transfer(address indexed from, address indexed to, uint256 value)',
    'event Approval(address indexed owner, address indexed spender, uint256 value)',
    'function approve(address spender, uint256 amount) public returns (bool)',
    'function allowance(address owner, address spender) public view returns (uint256)',
    'function balanceOf(address who) public view returns (uint256)',
    'function name() public view returns (string memory)',
    'function symbol() public view returns (string memory)',
    'function decimals() public view returns (uint8)',
    'function transfer(address who, uint256 amount)',
  ];
  const xrpToken = new Contract(
    xrpTokenAddress,
    erc20Abi,
    deployer,
  ) as SyloToken;
  const balance = await xrpToken.balanceOf(deployer.address);
  console.log('balance', deployer.address, balance.toString());
}

main();
