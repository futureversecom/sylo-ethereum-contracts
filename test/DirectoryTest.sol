pragma solidity ^0.6.0;
pragma experimental ABIEncoderV2;

import "truffle/Assert.sol";
import "../contracts/Directory.sol";

contract TestDirectory {

  SyloToken token;
  Directory directory;

  uint256 defaultUnlockDuration = 0;
  uint256 max = 2^256-1;

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

    directory.addStake(amount);

    Assert.equal(token.balanceOf(address(this)), initialBalance - amount, "Expected token balance to be reduced");
  }

  function testScanSingleStaker() public {
    directory.addStake(1 ether);

    address selected = directory.scan(0);

    Assert.equal(selected, address(this), "Expected correct scan with a single staker");
  }

  function testScanMultipleStaker() public {
    address otherStaker = 0x2074D810CDaAaf8b2D04A6E584B3fac7a4d85E15;
    directory.addStake(1 ether);
    directory.addStakeFor(1 ether, otherStaker);

    address selected = directory.scan(0);
    Assert.equal(selected, address(this), "Expected correct scan with multiple stakers 1");

    address selected2 = directory.scan(max);
    Assert.equal(selected2, otherStaker, "Expected correct scan with multiple stakers 2");

    address selected3 = directory.scan(max/2);
    Assert.equal(selected3, address(this), "Expected correct scan with multiple stakers 2");
  }

  function testUnlockedStake() public {
    address otherStaker = 0x2074D810CDaAaf8b2D04A6E584B3fac7a4d85E15;
    directory.addStake(1 ether);
    directory.addStakeFor(1 ether, otherStaker);

    Assert.equal(directory.getTotalStake(), 2 ether, "Unexpected total stake");

    directory.unlockStake();
    // Not unstaked but stake should no longer be valid because of unlock period being complete

    Assert.equal(directory.getTotalStake(), 1 ether, "Unexpected total stake post unlock");

    address selected = directory.scan(0);
    Assert.equal(selected, otherStaker, "Expected correct scan with a single staker");

    address selected2 = directory.scan(max);
    Assert.equal(selected2, otherStaker, "Expected correct scan with a single staker 2");
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