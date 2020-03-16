pragma solidity ^0.6.0;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/Token.sol";
import "../contracts/Ticketing.sol";

contract TestSyloTicketing {
  SyloToken token;
  SyloTicketing ticketing;

  uint256 defaultUnlockDuration = 1;

  function beforeEach() public {
    token = new SyloToken();
    ticketing = new SyloTicketing(token, defaultUnlockDuration);

    // Allow ticketing to transfer tokens from user
    token.approve(address(ticketing), 100 ether);
  }

  function testSettingUnlockDuration() public {
    ticketing.setUnlockDuration(0);

    // TODO find away to call this from another address
  }

  function testDepositEscrow() public {
    uint256 amount = 1 ether;
    uint256 initialBalance = token.balanceOf(address(this));
    ticketing.depositEscrow(amount);

    (uint256 escrow, ,) = ticketing.getDepositDetails(address(this));

    Assert.equal(escrow, amount, "Expected correct escrow amount");
    Assert.equal(token.balanceOf(address(this)), initialBalance - amount, "Expected token balance to be reduced");
  }

  function testDepositPenalty() public {
    uint256 amount = 1 ether;
    uint256 initialBalance = token.balanceOf(address(this));
    ticketing.depositPenalty(amount);

    (,uint256 penalty,) = ticketing.getDepositDetails(address(this));

    Assert.equal(penalty, amount, "Expected correct penalty amount");
    Assert.equal(token.balanceOf(address(this)), initialBalance - amount, "Expected token balance to be reduced");
  }

  function testMultipleEscrowDeposits() public {
    uint256 amount = 1 ether;
    ticketing.depositEscrow(amount);
    ticketing.depositEscrow(amount);

    (uint256 escrow, ,) = ticketing.getDepositDetails(address(this));

    Assert.equal(escrow, amount * 2, "Expected correct escrow amount");
  }

  function testMultiplePenaltyDeposits() public {
    uint256 amount = 1 ether;
    ticketing.depositPenalty(amount);
    ticketing.depositPenalty(amount);

    (,uint256 penalty,) = ticketing.getDepositDetails(address(this));

    Assert.equal(penalty, amount * 2, "Expected correct penalty amount");
  }

  function testDepositEscrowAndPenalty() public {
    uint256 amount = 1 ether;
    uint256 initialBalance = token.balanceOf(address(this));
    ticketing.depositEscrow(amount);
    ticketing.depositPenalty(amount);

    (uint256 escrow ,uint256 penalty,) = ticketing.getDepositDetails(address(this));

    Assert.equal(escrow, amount, "Expected correct escrow amount");
    Assert.equal(penalty, amount, "Expected correct penalty amount");
    Assert.equal(token.balanceOf(address(this)), initialBalance - amount * 2, "Expected token balance to be reduced");
  }

  function testWithdrawWithoutUnlock() public {
    testDepositEscrow();

    try ticketing.withdraw() {
      Assert.fail("withdraw should fail");
    } catch Error(string memory reason) {
      Assert.equal(reason, "Deposits not unlocked", "Expected specific error");
    }
  }

  function testUnlockWithoutDeposit() public {
    try ticketing.unlock() {
      Assert.fail("unlock should fail");
    } catch Error(string memory reason) {
      Assert.equal(reason, "Nothing to withdraw", "Expected specific error");
    }
  }

  function testUnlockWithDeposit() public {
    testDepositEscrow();

    ticketing.unlock();

    (,, uint256 unlockAt) = ticketing.getDepositDetails(address(this));

    Assert.equal(block.number + defaultUnlockDuration, unlockAt, "Expected a different unlock time");
  }

  function testUnlockWhileUnlocked() public {
    testDepositEscrow();

    ticketing.unlock();

    try ticketing.unlock() {
      Assert.fail("unlock should fail");
    } catch Error(string memory reason) {
      Assert.equal(reason, "Unlock already in progress", "Expected specific error");
    }
  }

  function testLockWhileLocked() public {

    try ticketing.lock() {
      Assert.fail("lock should fail");
    } catch Error(string memory reason) {
      Assert.equal(reason, "Not unlocking, cannot lock", "Expected specific error");
    }
  }

  function testLockWhileInUnlockingPeriod() public {
    testDepositEscrow();

    ticketing.unlock();

    ticketing.lock();

    (,, uint256 unlockAt) = ticketing.getDepositDetails(address(this));
    Assert.equal(0, unlockAt, "Expected an unlock time of 0");
  }

  function testLockWhileCompletedUnlocking() public {
    ticketing.setUnlockDuration(0);

    testDepositEscrow();

    ticketing.unlock();

    ticketing.lock();

    (,, uint256 unlockAt) = ticketing.getDepositDetails(address(this));
    Assert.equal(0, unlockAt, "Expected an unlock time of 0");
  }

  function testSuccessfulWithdraw() public {
    ticketing.setUnlockDuration(0);

    uint256 initialBalance = token.balanceOf(address(this));

    testDepositEscrow();
    testDepositPenalty();

    ticketing.unlock();

    ticketing.withdraw();

    (uint256 escrow ,uint256 penalty,) = ticketing.getDepositDetails(address(this));

    Assert.equal(0, escrow, "Expected escrow to be 0");
    Assert.equal(0, penalty, "Expected penalty to be 0");

    Assert.equal(initialBalance, token.balanceOf(address(this)), "Expected balance to be restored after withdrawing");
  }

}