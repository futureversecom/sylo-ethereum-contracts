// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

interface ISyloTicketing {
    struct Deposit {
        uint256 escrow; // Balance of users escrow
        uint256 penalty; // Balance of users penalty
        uint256 unlockAt; // Block number a user can withdraw their balances
    }

    struct User {
        address main; // Main address of the ticket sender or receiver
        address delegated; // Delegated address used to sign and redeem tickets
    }

    struct Ticket {
        uint256 epochId; // The epoch this ticket is associated with
        User sender; // Ticket sender's main and delegated addresses
        User receiver; // Ticket receiver's main and delegated addresses
        address redeemer; // Address of the intended recipient
        uint256 generationBlock; // Block number the ticket was generated
        bytes32 redeemerCommit; // Hash of the secret random number of the redeemer
    }

    function setUnlockDuration(uint256 _unlockDuration) external;

    function depositEscrow(uint256 amount, address account) external;

    function depositPenalty(uint256 amount, address account) external;

    function unlockDeposits() external returns (uint256);

    function lockDeposits() external;

    function withdraw() external;

    function redeem(
        Ticket calldata ticket,
        uint256 redeemerRand,
        bytes calldata senderSig,
        bytes calldata receiverSig
    ) external;
}
