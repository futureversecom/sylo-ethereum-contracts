const fs = require('fs/promises');
const path = require('path');

const PKG_DIR =
  process.env.ABI_PACKAGE ??
  path.resolve(__dirname, '../package/@sylo/ethereum-contracts-abi');
const ABI_DIR = path.resolve(__dirname, '../abi');

const syloContracts = [
  'Directory',
  'EpochsManager',
  'Listings',
  'RewardsManager',
  'StakingManager',
  'SyloTicketing',
  'SyloToken',
  'TicketingParameters',
];

async function parseAbis() {
  const abis = await Promise.all(
    syloContracts.map(async contract => {
      const src = `${ABI_DIR}/${contract}.abi`;
      const abi = await fs.readFile(src).then(b => b.toString());

      await fs.copyFile(src, `${PKG_DIR}/${contract}.abi`);

      return { contract, abi };
    }),
  );

  return abis;
}

async function writePackageIndex(abis) {
  let f = ``;

  abis.map(a => {
    f += `const ${a.contract} = ${a.abi};\n`;
  });

  f += '\n';
  f += 'module.exports = {\n';

  abis.map((a, i) => {
    f += `  ${a.contract}: ${a.contract}${i < abis.length - 1 ? ',' : ''}\n`;
  });

  f += '}';

  await fs.writeFile(`${PKG_DIR}/index.js`, f);
}

async function main() {
  const abis = await parseAbis();
  await writePackageIndex(abis);
}

main();
