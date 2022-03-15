import { ethers } from "hardhat";
import { Signer } from "ethers";
import { Channels } from "../typechain";
import { assert, expect } from "chai";

describe('Channel', () => {
  let accounts: Signer[];

  let channels: Channels;

  before(async () => {
    accounts = await ethers.getSigners();
  });

  beforeEach(async () => {
    const Channels = await ethers.getContractFactory("Channels");
    channels = (await Channels.deploy()) as Channels;
    await channels.initialize({ from: await accounts[0].getAddress() });
  });

  it('can set channel', async () => {
    await expect(channels.setChannel(1))
      .to.emit(channels, 'NewChannelUpdated')
      .withArgs(1);

    const channel = await channels.getChannel();
    assert.equal(channel, 1, "Expected channel to be 1");
  });

  it('cannot set channel greater than 48', async () => {
    await expect(channels.setChannel(49))
      .to.be.revertedWith("Channels cannot be more than 48");
  });
});
