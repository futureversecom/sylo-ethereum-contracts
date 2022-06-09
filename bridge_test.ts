import { Api } from "@cennznet/api";
import * as hardhat from "hardhat";
import * as ethers from "ethers";
import { ERC20Peg__factory } from "./typechain/factories/ERC20Peg__factory";
import {
  ERC20__factory,
  StakingManager__factory,
  SyloToken__factory,
} from "./typechain";
import { EventProofId } from "@cennznet/types";
import { EthEventProof } from "@cennznet/api/derives/ethBridge/types";
import { CENNZnetBridge__factory } from "./typechain/factories/CENNZnetBridge__factory";
const web3 = require("web3");
const utils = require("@polkadot/util");
const utilCrypto = require("@polkadot/util-crypto");

const SYLO_TOKEN_ASSET_ID = 17076;

const RATA_SYLO_TOKEN_ASSET_ID = 17012;

// async function main() {
//   const signer = await hardhat.ethers.getSigners().then((xs) => xs[0]);

//   const syloToken = SyloToken__factory.connect(
//     "0xa5A9BC81f10Eb2ec7f881e6730D5BD5a8a893eA6",
//     signer
//   );
//   const erc20Peg = ERC20Peg__factory.connect(
//     "0x4C411B3Bf36D6DE908C6f4256a72B85E3f2B00bF",
//     signer
//   );

//   console.log("approving sylo token");
//   await syloToken.approve(erc20Peg.address, ethers.utils.parseEther("1000"));

//   const cennzAddress = cvmToAddress(signer.address);

//   console.log("depositing to peg");
//   const tx = await erc20Peg.deposit(
//     syloToken.address,
//     ethers.utils.parseEther("1000"),
//     utilCrypto.decodeAddress("5DJTrWDe5vbs1aB9GWTX93SAz99SkC21KK8L2zfYU6LFpJYJ")
//     // utilCrypto.decodeAddress(cennzAddress),
//   );

//   console.log("peg deposit tx hash:", tx.hash, " for address: ", cennzAddress);

//   process.exit();
// }

async function main() {
  console.log("testing fee proxy call...");

  const signer = await hardhat.ethers.getSigners().then((xs) => xs[0]);

  // Defines the functions, used for the abi to encode the values of that function
  const erc20Abi = ["function transfer(address who, uint256 amount)"];
  let iface = new ethers.utils.Interface(erc20Abi);
  const transferInput = iface.encodeFunctionData("transfer", [
    "0xAB8208e1adEBe8Ee6155Ef9A155E2cdA3B880cc7",
    1000,
  ]);

  const feeProxyAbi = [
    "function callWithFeePreferences(address asset, uint32 slippage, address target, bytes input)",
  ];

  const feeProxyAddress = "0x00000000000000000000000000000000000004bb";
  const feeProxy = new ethers.Contract(feeProxyAddress, feeProxyAbi, signer);

  // SYLO testnet token address derived from the generic asset Id (`17012`)
  const syloTokenAddress = "0xcCCCcCcC00004274000000000000000000000000";

  // The slippage value for exchanging between payment asset and CPAY (out of 1000)
  const slippage = 50; // 5%

  console.log("calling with fee preference...");
  // Call the fee proxy contract with the above values and input from the previous example.
  await feeProxy.callWithFeePreferences(
    syloTokenAddress,
    slippage,
    syloTokenAddress,
    transferInput
  );
}

main();

async function checkCennzSyloToken() {
  const signer = await hardhat.ethers.getSigners().then((xs) => xs[0]);

  const syloToken = SyloToken__factory.connect(getTokenAddress(), signer);

  console.log(
    await syloToken
      .balanceOf(signer.address)
      .then((b) => ethers.utils.formatEther(b.toString()))
  );
}

function getTokenAddress() {
  const prefix = "cccccccc0000" + SYLO_TOKEN_ASSET_ID.toString(16);
  const address = web3.utils.toChecksumAddress(
    prefix + "0".repeat(40 - prefix.length)
  );
  return address;
}

function _getTokenAddress(assetId: number) {
  const prefix = "cccccccc0000" + assetId.toString(16);
  const address = web3.utils.toChecksumAddress(
    prefix + "0".repeat(40 - prefix.length)
  );
  return address;
}

function cvmToAddress(cvmAddress: string) {
  var message = utils.stringToU8a("cvm:");
  message = utils.u8aConcat(
    message,
    new Array(7).fill(0),
    utils.hexToU8a(cvmAddress)
  );
  let checkSum = message.reduce((a: any, b: any) => a ^ b, 0);
  message = utils.u8aConcat(message, new Array(1).fill(checkSum));

  return utilCrypto.encodeAddress(message, 42);
}

async function testStakeDeposit() {
  const signer = await hardhat.ethers.getSigners().then((xs) => xs[0]);

  const syloToken = SyloToken__factory.connect(getTokenAddress(), signer);

  const staking = StakingManager__factory.connect(
    "0x4EafFD4c012A7D4215790225776564Ebd69fa0D2",
    signer
  );

  const originalBalance = await syloToken.balanceOf(signer.address);

  console.log(`original balance: ${ethers.utils.formatEther(originalBalance)}`);

  await syloToken.approve(staking.address, ethers.utils.parseEther("1000"));

  console.log("getting stake entry");

  const se = await staking.getStakeEntry(signer.address, signer.address);

  console.log(se);

  console.log("depositing stake...");

  const tx = await staking.addStake(
    ethers.utils.parseEther("10"),
    signer.address
  );

  console.log("sent add stake tx", tx.hash);

  const postBalance = await syloToken.balanceOf(signer.address);

  console.log(ethers.utils.formatEther(postBalance));
}

async function withdraw() {
  const signer = await hardhat.ethers.getSigners().then((xs) => xs[0]);

  const syloToken = SyloToken__factory.connect(
    "0x262EA359Ee8E01f03c9022f1Ae0889665f6a8EF2",
    signer
  );
  const erc20Peg = ERC20Peg__factory.connect(
    "0xa39E871e6e24f2d1Dd6AdA830538aBBE7b30F78F",
    signer
  );
  const bridge = CENNZnetBridge__factory.connect(
    "0x6484A31Df401792c784cD93aAAb3E933B406DdB3",
    signer
  );

  const api = await Api.create({
    provider: "wss://nikau.centrality.me/public/ws",
  });

  const cennzAddress = cvmToAddress(signer.address);

  console.log("connected to nikau...");

  const nonce = await api.rpc.system.accountNextIndex(
    utilCrypto.decodeAddress(cennzAddress)
  );
  const amount = 100;
  const ethBeneficiary = signer.address;
  const call = api.tx.erc20Peg.withdraw(
    SYLO_TOKEN_ASSET_ID,
    amount,
    ethBeneficiary
  );

  const payload = api.createType("EthWalletCall", { call, nonce }).toU8a();
  const signature = await signer.signMessage(payload);

  console.log("created payload and signature", payload, signature);

  const eventProofId: EventProofId = await new Promise(async (resolve) => {
    const txHash = await api.tx.ethWallet
      .call(call, signer.address, signature)
      .send(async ({ status, events }) => {
        if (status.isInBlock) {
          for (const {
            event: { method, section, data },
          } of events) {
            if (section === "erc20Peg" && method == "Erc20Withdraw") {
              console.log("*******************************************");
              console.log("Withdraw claim on CENNZnet side successfully");
              console.log(data[0]);
              resolve(data[0] as any);
            }
          }
        }
      });

    console.log("txHash", txHash);
  });

  const eventProof: EthEventProof = await new Promise(async (resolve) => {
    const unsubHeads = await api.rpc.chain.subscribeNewHeads(async () => {
      console.log(`Waiting till event proof is fetched....`);
      const eventProof = await api.derive.ethBridge.eventProof(eventProofId);

      console.log("Event proof", eventProof);

      if (eventProof !== null) {
        console.log("Event proof found;::", eventProof);
        unsubHeads();
        resolve(eventProof);
      }
    });
  });

  console.log("sending erc20 peg withdraw");

  const b = await signer.getBalance();
  console.log(b.toString());

  const value = await bridge.verificationFee();

  console.log(value.toString());

  const withdrawTx = await erc20Peg.withdraw(
    syloToken.address,
    amount,
    signer.address,
    eventProof,
    { value }
  );

  await withdrawTx.wait(1);

  console.log(withdrawTx.hash);
}

// deployTestContract();

// testStakeDeposit();

// withdraw();

// checkCennzSyloToken();

// main();
