// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "../IAuthorizedAccounts.sol";

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

    enum SignatureType {
        Main,
        Authorized,
        AttachedAuthorized
    }

    struct UserSignature {
        SignatureType sigType;
        bytes signature;
        // This field will only be present if the sig type is `AttachedAuthorized`
        IAuthorizedAccounts.AttachedAuthorizedAccount attachedAuthorizedAccount;
    }

    struct Ticket {
        uint256 epochId; // The epoch this ticket is associated with
        User sender; // Ticket sender's main and delegated addresses
        User receiver; // Ticket receiver's main and delegated addresses
        address redeemer; // Address of the intended recipient
        uint256 generationBlock; // Block number the ticket was generated
        bytes32 redeemerCommit; // Hash of the secret random number of the redeemer
    }

    // A type of ticket that does not explicit state the receiver address.
    struct MultiReceiverTicket {
        uint256 epochId; // The epoch this ticket is associated with
        User sender; // Ticket sender's main and delegated addresses
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

    function redeemV2(
        Ticket calldata ticket,
        uint256 redeemerRand,
        UserSignature calldata senderSig,
        UserSignature calldata receiverSig
    ) external;

    function redeemMultiReceiver(
        MultiReceiverTicket calldata ticket,
        uint256 redeemerRand,
        User calldata receiver,
        bytes calldata senderSig,
        bytes calldata receiverSig
    ) external;

    function redeemMultiReceiverV2(
        MultiReceiverTicket calldata ticket,
        uint256 redeemerRand,
        User calldata receiver,
        UserSignature calldata senderSig,
        UserSignature calldata receiverSig
    ) external;
}
