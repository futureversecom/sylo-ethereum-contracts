// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

interface ISyloTicketing {
    struct Deposit {
        uint256 escrow; // Balance of users escrow
        uint256 penalty; // Balance of users penalty
        uint256 unlockAt; // Block number a user can withdraw their balances
    }

    struct Ticket {
        uint256 epochId; // The epoch this ticket is associated with
        address sender; // Address of the ticket sender
        address delegatedSender; // Address of the ticket's signer if not the original sender (optional)
        address redeemer; // Address of the intended recipient
        uint256 generationBlock; // Block number the ticket was generated
        bytes32 senderCommit; // Hash of the secret random number of the sender
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
        uint256 senderRand,
        uint256 redeemerRand,
        bytes calldata sig
    ) external;
}
