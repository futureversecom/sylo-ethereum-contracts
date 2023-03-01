import { ethers } from 'ethers';
import { TestSeekers__factory } from '../typechain-types';

const args = require('minimist')(process.argv.slice(2), {
  string: ['account_pk', 'seeker_contract', 'evm'],
  number: ['amount'],
}) as Args;

type Args = {
  evm: string;
  account_pk: string;
  seeker_contract: string;
  amount: number;
};

async function main() {
  // if (args.evm == null) {
  //   throw new Error('Must provide `--evm` arg');
  // }

  // if (args.account_pk == null) {
  //   throw new Error('Must provide `--account_pk` arg');
  // }

  // if (args.seeker_contract == null) {
  //   throw new Error('Must provide `--seeker_contract` arg');
  // }

  // if (args.amount == null) {
  //   throw new Error('Must provide `--amount` arg');
  // }

  const provider = new ethers.providers.JsonRpcProvider('http://0.0.0.0:8545');

  const account = new ethers.Wallet(
    '0x150934096e7bcd0485d154edd771b4466680038a068ccca8e8b483dce8527245',
  ).connect(provider);

  const seeker = TestSeekers__factory.connect(
    '0x49C537a88016186Ef41713239799Fc975F9e9aFA',
    account,
  );

  for (let i = 0; i < 10; i++) {
    const tx = await seeker.mint(account.address, i);
    await tx.wait();
  }

  console.log('Minted all requested seekers');
}

main().catch(e => console.log(e));
