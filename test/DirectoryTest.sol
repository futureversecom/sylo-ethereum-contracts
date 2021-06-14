// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.6.0;
pragma experimental ABIEncoderV2;

import "truffle/Assert.sol";
import "../contracts/Directory.sol";

contract TestDirectory {

  SyloToken token;
  Directory directory;

  uint256 defaultUnlockDuration = 0;
  uint128 max = (1 << 128) - 1;

  event ScanResult(uint256,uint256,uint256,address);

  function emitStake(bytes32 key) private {

    (uint256 amount, uint256 leftAmount, uint256 rightAmount, address stakee,,,) = directory.stakes(key);
    emit ScanResult(amount, leftAmount, rightAmount, stakee);
  }

  function beforeEach() public {
    token = new SyloToken();
    directory = new Directory(token, defaultUnlockDuration);

    // Allow ticketing to transfer tokens from user
    token.approve(address(directory), 100 ether);
  }

  function testScanEmptyDirectory() public {

    Assert.equal(directory.scan(0), address(0), "Expected null address");
    Assert.equal(directory.scan(max), address(0), "Expected null address");
  }

  function testDepositStake() public {
    uint256 amount = 1 ether;
    uint256 initialBalance = token.balanceOf(address(this));

    directory.addStake(amount, address(this));

    uint256 stake = directory.stakees(address(this));

    Assert.equal(stake, amount, "Expected stake to equal amount staked");

    Assert.equal(token.balanceOf(address(this)), initialBalance - amount, "Expected token balance to be reduced");
  }

  function testScanSingleStaker() public {
    directory.addStake(1 ether, address(this));

    address selected = directory.scan(0);

    Assert.equal(selected, address(this), "Expected correct scan with a single staker");
  }

  function testScanMultipleStaker() public {
    address otherStaker = 0x2074D810CDaAaf8b2D04A6E584B3fac7a4d85E15;
    directory.addStake(1 ether, address(this));
    directory.addStake(1 ether, otherStaker);

    Assert.equal(directory.getTotalStake(), 2 ether, "Expected a total stake of 2 ether");

    address selected = directory.scan(0);
    Assert.equal(selected, address(this), "Expected correct scan with multiple stakers 1");

    address selected2 = directory.scan(max);
    Assert.equal(selected2, otherStaker, "Expected correct scan with multiple stakers 2");

    address selected3 = directory.scan(max/2);
    Assert.equal(selected3, address(this), "Expected correct scan with multiple stakers 3");
  }

  function testUnlockStake() public {
    address otherStaker = 0x2074D810CDaAaf8b2D04A6E584B3fac7a4d85E15;
    directory.addStake(1 ether, address(this));
    directory.addStake(1 ether, otherStaker);

    Assert.equal(directory.getTotalStake(), 2 ether, "Unexpected total stake");

    directory.unlockStake(1 ether, address(this));
    // Not unstaked but stake should no longer be valid because of unlock period being complete

    emitStake(directory.getKey(address(this), address(this)));
    emitStake(directory.getKey(address(this), 0x2074D810CDaAaf8b2D04A6E584B3fac7a4d85E15));

    Assert.equal(directory.getTotalStake(), 1 ether, "Unexpected total stake post unlock");

    address selected = directory.scan(0);
    Assert.equal(selected, otherStaker, "Expected correct scan with a single staker");

    address selected2 = directory.scan(max);
    Assert.equal(selected2, otherStaker, "Expected correct scan with a single staker 2");
  }

  function testUnlockStakeLeaf() public {
    address otherStaker = 0x2074D810CDaAaf8b2D04A6E584B3fac7a4d85E15;
    directory.addStake(1 ether, address(this));
    directory.addStake(1 ether, otherStaker);

    Assert.equal(directory.getTotalStake(), 2 ether, "Unexpected total stake");

    directory.unlockStake(1 ether, otherStaker);
    // Not unstaked but stake should no longer be valid because of unlock period being complete

    Assert.equal(directory.getTotalStake(), 1 ether, "Unexpected total stake post unlock");

    address selected = directory.scan(0);
    Assert.equal(selected, address(this), "Expected correct scan with a single staker");

    address selected2 = directory.scan(max);
    Assert.equal(selected2, address(this), "Expected correct scan with a single staker 2");
  }

  function testUnlockStakeRoot() public {
    address addrA = 0x2074D810CDaAaf8b2D04A6E584B3fac7a4d85E15;
    address addrB = 0x43c68Be1eDeaE941Ac2baf424455aaaB0c9D440D;

    directory.addStake(10 ether, address(this));
    directory.addStake(5 ether, addrA);
    directory.addStake(5 ether, addrB);

    Assert.equal(directory.getTotalStake(), 20 ether, "Unexpected total stake");

    // Not unstaked but stake should no longer be valid because of unlock period being complete
    directory.unlockStake(10 ether, address(this));

    Assert.equal(directory.getTotalStake(), 10 ether, "Unexpected total stake post unlock");

    address selected = directory.scan(0);
    Assert.equal(selected, addrB, "Expected correct scan with a single staker");

    address selected2 = directory.scan(max);
    Assert.equal(selected2, addrA, "Expected correct scan with a single staker 2");
  }

  /* Disabled because we cannot advance the block in solidity tests */
  // function testUnstaking() public {
  //   address otherStaker = 0x2074D810CDaAaf8b2D04A6E584B3fac7a4d85E15;
  //   directory.addStake(1 ether);
  //   directory.addStakeFor(1 ether, otherStaker);

  //   Assert.equal(directory.getTotalStake(), 2 ether, "Unexpected total stake");

  //   directory.unlockStake();
  //   directory.unstake();

  //   Assert.equal(directory.getTotalStake(), 1 ether, "Unexpected total stake post unstake");

  //   address selected = directory.scan(0);
  //   Assert.equal(selected, otherStaker, "Expected correct scan with a single staker");

  //   address selected2 = directory.scan(max);
  //   Assert.equal(selected2, otherStaker, "Expected correct scan with a single staker 2");
  // }
}
