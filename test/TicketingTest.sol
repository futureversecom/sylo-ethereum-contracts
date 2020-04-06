pragma solidity ^0.6.0;
pragma experimental ABIEncoderV2;

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

    (uint256 escrow, ,) = ticketing.deposits(address(this));

    Assert.equal(escrow, amount, "Expected correct escrow amount");
    Assert.equal(token.balanceOf(address(this)), initialBalance - amount, "Expected token balance to be reduced");
  }

  function testDepositPenalty() public {
    uint256 amount = 1 ether;
    uint256 initialBalance = token.balanceOf(address(this));
    ticketing.depositPenalty(amount);

    (,uint256 penalty,) = ticketing.deposits(address(this));

    Assert.equal(penalty, amount, "Expected correct penalty amount");
    Assert.equal(token.balanceOf(address(this)), initialBalance - amount, "Expected token balance to be reduced");
  }

  function testMultipleEscrowDeposits() public {
    uint256 amount = 1 ether;
    ticketing.depositEscrow(amount);
    ticketing.depositEscrow(amount);

    (uint256 escrow, ,) = ticketing.deposits(address(this));

    Assert.equal(escrow, amount * 2, "Expected correct escrow amount");
  }

  function testMultiplePenaltyDeposits() public {
    uint256 amount = 1 ether;
    ticketing.depositPenalty(amount);
    ticketing.depositPenalty(amount);

    (,uint256 penalty,) = ticketing.deposits(address(this));

    Assert.equal(penalty, amount * 2, "Expected correct penalty amount");
  }

  function testDepositEscrowAndPenalty() public {
    uint256 amount = 1 ether;
    uint256 initialBalance = token.balanceOf(address(this));
    ticketing.depositEscrow(amount);
    ticketing.depositPenalty(amount);

    (uint256 escrow ,uint256 penalty,) = ticketing.deposits(address(this));

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
    try ticketing.unlockDeposits() {
      Assert.fail("unlock should fail");
    } catch Error(string memory reason) {
      Assert.equal(reason, "Nothing to withdraw", "Expected specific error");
    }
  }

  function testUnlockWithDeposit() public {
    testDepositEscrow();

    ticketing.unlockDeposits();

    (,, uint256 unlockAt) = ticketing.deposits(address(this));

    Assert.equal(block.number + defaultUnlockDuration, unlockAt, "Expected a different unlock time");
  }

  function testUnlockWhileUnlocked() public {
    testDepositEscrow();

    ticketing.unlockDeposits();

    try ticketing.unlockDeposits() {
      Assert.fail("unlock should fail");
    } catch Error(string memory reason) {
      Assert.equal(reason, "Unlock already in progress", "Expected specific error");
    }
  }

  function testLockWhileLocked() public {

    try ticketing.lockDeposits() {
      Assert.fail("lock should fail");
    } catch Error(string memory reason) {
      Assert.equal(reason, "Not unlocking, cannot lock", "Expected specific error");
    }
  }

  function testLockWhileInUnlockingPeriod() public {
    testDepositEscrow();

    ticketing.unlockDeposits();

    ticketing.lockDeposits();

    (,, uint256 unlockAt) = ticketing.deposits(address(this));
    Assert.equal(0, unlockAt, "Expected an unlock time of 0");
  }

  function testLockWhileCompletedUnlocking() public {
    ticketing.setUnlockDuration(0);

    testDepositEscrow();

    ticketing.unlockDeposits();

    ticketing.lockDeposits();

    (,, uint256 unlockAt) = ticketing.deposits(address(this));
    Assert.equal(0, unlockAt, "Expected an unlock time of 0");
  }

  /* Disabled because we cannot advance the block in solidity tests */
  // function testSuccessfulWithdraw() public {
  //   ticketing.setUnlockDuration(0);

  //   uint256 initialBalance = token.balanceOf(address(this));

  //   testDepositEscrow();
  //   testDepositPenalty();

  //   ticketing.unlockDeposits();

  //   ticketing.withdraw();

  //   (uint256 escrow ,uint256 penalty,) = ticketing.deposits(address(this));

  //   Assert.equal(0, escrow, "Expected escrow to be 0");
  //   Assert.equal(0, penalty, "Expected penalty to be 0");

  //   Assert.equal(initialBalance, token.balanceOf(address(this)), "Expected balance to be restored after withdrawing");
  // }

  /* Hard coded values are generated in github.com/dn3010/go-probabilistic-micropayments */
  function createTicket() internal pure returns (SyloTicketing.Ticket memory, uint256, bytes memory) {
    uint256 receiverRand = 1;

    SyloTicketing.Ticket memory t = SyloTicketing.Ticket({
      sender: 0x2074D810CDaAaf8b2D04A6E584B3fac7a4d85E15,
      receiver: 0x84f8579a947c631362c47d534f26d8E46d400157,
      senderNonce: 1,
      faceValue: 1,
      winProb: 2^256-1, // 100% chance,
      expirationBlock: 0, // Never expires
      receiverRandHash: keccak256(abi.encodePacked(receiverRand))
    });

    bytes memory sig = hex"523a703ac7588034d851be80eb2d1ca1a124154dac62930cc29a9a8eb0085c066889d26f36871af786871b1fbea42530babc6856784ca2adfef48db5ba56a03401";

    return (t, receiverRand, sig);
  }

  function testRedeemingExpiredTicket() public {

    (SyloTicketing.Ticket memory ticket, uint256 receiverRand, bytes memory sig) = createTicket();

    ticket.expirationBlock = 1;

    try ticketing.redeem(ticket, receiverRand, sig) {
      Assert.fail("Ticket should be invalid");
    } catch Error(string memory reason) {
      Assert.equal(reason, "Ticket has expired", "Expected specific error");
    }
  }

  function testRedeemingWithInvalidSig() public {

    (SyloTicketing.Ticket memory ticket, uint256 receiverRand,) = createTicket();

    bytes memory sig;

    try ticketing.redeem(ticket, receiverRand, sig) {
      Assert.fail("Ticket should be invalid");
    } catch Error(string memory reason) {
      Assert.equal(reason, "ECDSA: invalid signature length", "Expected specific error");
    }
  }

  function testRedeemingInvalidReceiverRand() public {

    (SyloTicketing.Ticket memory ticket, , bytes memory sig) = createTicket();

    try ticketing.redeem(ticket, 2, sig) {
      Assert.fail("Ticket should be invalid");
    } catch Error(string memory reason) {
      Assert.equal(reason, "Hash of receiverRand doesn't match receiverRandHash", "Expected specific error");
    }
  }
}
