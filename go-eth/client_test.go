package eth

import (
	"context"
	"encoding/hex"
	"math/big"
	"testing"
	"time"

	"github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var unlockDuration = big.NewInt(10)
var alwaysFailing = "failed to estimate gas needed: gas required exceeds allowance or always failing transaction"

func startSimulatedBackend(auth *bind.TransactOpts) SimBackend {
	var gasLimit uint64 = 50000000

	genisis := make(core.GenesisAlloc)

	genisis[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000000)}

	sim := backends.NewSimulatedBackend(genisis, gasLimit)

	return newSimBackend(sim)
}

func deployContracts(auth *bind.TransactOpts, backend SimBackend) (Addresses, error) {

	var addresses Addresses = Addresses{}

	tokenAddress, tx, _, err := contracts.DeploySyloToken(auth, backend)
	addresses.Token = tokenAddress
	if err != nil {
		return Addresses{}, err
	}
	backend.Commit()

	_, err = backend.TransactionReceipt(auth.Context, tx.Hash())

	if err != nil {
		return Addresses{}, errors.Wrap(err, "Failed to deploy token")
	}

	addresses.Ticketing, tx, _, err = contracts.DeploySyloTicketing(auth, backend, addresses.Token, unlockDuration)
	if err != nil {
		return Addresses{}, err
	}
	backend.Commit()

	_, err = backend.TransactionReceipt(auth.Context, tx.Hash())

	if err != nil {
		return Addresses{}, errors.Wrap(err, "Failed to deploy ticketing")
	}

	addresses.Directory, tx, _, err = contracts.DeployDirectory(auth, backend, addresses.Token, unlockDuration)
	if err != nil {
		return Addresses{}, err
	}
	backend.Commit()

	_, err = backend.TransactionReceipt(auth.Context, tx.Hash())

	if err != nil {
		return Addresses{}, errors.Wrap(err, "Failed to deploy directory")
	}

	addresses.Listings, tx, _, err = contracts.DeployListings(auth, backend)
	if err != nil {
		return Addresses{}, err
	}
	backend.Commit()

	_, err = backend.TransactionReceipt(auth.Context, tx.Hash())

	if err != nil {
		return Addresses{}, errors.Wrap(err, "Failed to deploy listings")
	}

	return addresses, nil
}

func createClientWithBackend(backend SimBackend, auth *bind.TransactOpts) (Client, error) {
	addresses, err := deployContracts(auth, backend)

	if err != nil {
		return nil, err
	}

	return NewClientWithBackend(addresses, backend, auth)
}

func TestClient_CanBeCreated(t *testing.T) {

	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	backend := startSimulatedBackend(auth)

	addresses, err := deployContracts(auth, backend)
	assert.Nil(t, err, "Failed to deploy contracts")

	if (addresses.Token == ethcommon.Address{}) {
		t.Error("Token address is empty")
	}

	if (addresses.Ticketing == ethcommon.Address{}) {
		t.Error("ticketingAddress address is empty")
	}

	_, err = NewClientWithBackend(addresses, backend, auth)
	assert.Nil(t, err, "Failed to init client")
}

func TestClient_LatestBlock(t *testing.T) {

	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	backend := startSimulatedBackend(auth)

	client, err := createClientWithBackend(backend, auth)
	assert.Nil(t, err, "Failed to init client")

	blockNumber, err := client.LatestBlock()
	assert.Nil(t, err, "Failed to get latest block")

	assert.GreaterOrEqual(t, blockNumber.Cmp(big.NewInt(0)), 0)
}

func TestClient_DepositEscrow(t *testing.T) {

	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	auth.Context = context.Background()

	backend := startSimulatedBackend(auth)

	client, err := createClientWithBackend(backend, auth)
	assert.Nil(t, err, "Failed to init client")

	// Approve ticketing contract to transfer funds
	tx, err := client.ApproveTicketing(big.NewInt(100000))
	assert.Nil(t, err, "Failed to approve ticketing")

	backend.Commit()

	duration := 10 * time.Second
	_, err = client.CheckTxTimeout(tx, duration)
	if err != nil {
		assert.Nil(t, err, "Failed to confirm approve tx")
	}

	tx, err = client.DepositEscrow(big.NewInt(1), auth.From)

	backend.Commit()

	_, err = client.CheckTxTimeout(tx, duration)
	if err != nil {
		assert.Nil(t, err, "Failed to confirm deposit tx")
	}

	deposit, err := client.Deposits(auth.From)
	if err != nil {
		assert.Nil(t, err, "Failed to get deposits")
	}

	assert.Equal(t, deposit.Escrow.Cmp(big.NewInt(1)), 0, "Deposit doesn't match")
}

func TestClient_DepositPenalty(t *testing.T) {

	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	auth.Context = context.Background()

	backend := startSimulatedBackend(auth)

	client, err := createClientWithBackend(backend, auth)
	assert.Nil(t, err, "Failed to init client")

	// Approve ticketing contract to transfer funds
	tx, err := client.ApproveTicketing(big.NewInt(100000))
	assert.Nil(t, err, "Failed to approve ticketing")

	backend.Commit()

	duration := 10 * time.Second
	_, err = client.CheckTxTimeout(tx, duration)
	if err != nil {
		assert.Nil(t, err, "Failed to confirm approve tx")
	}

	tx, err = client.DepositPenalty(big.NewInt(1), auth.From)

	backend.Commit()

	_, err = client.CheckTxTimeout(tx, duration)
	if err != nil {
		assert.Nil(t, err, "Failed to confirm deposit tx")
	}

	deposit, err := client.Deposits(auth.From)
	if err != nil {
		assert.Nil(t, err, "Failed to get deposits")
	}

	assert.Equal(t, deposit.Penalty.Cmp(big.NewInt(1)), 0, "Deposit doesn't match")
}

func TestClient_WithdrawTicketing(t *testing.T) {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	auth.Context = context.Background()

	backend := startSimulatedBackend(auth)

	client, err := createClientWithBackend(backend, auth)
	assert.Nil(t, err, "Failed to init client")

	_, err = client.ApproveTicketing(big.NewInt(100000))
	assert.Nil(t, err, "Failed to approve ticketing")

	backend.Commit()

	_, err = client.DepositEscrow(big.NewInt(10), auth.From)
	assert.Nil(t, err, "Failed to deposit escrow")

	backend.Commit()

	tx, err := client.UnlockDeposits()
	assert.Nil(t, err, "Failed to unlock deposit")

	backend.Commit()

	_, err = client.CheckTxTimeout(tx, 10*time.Second)

	// Expect error because unlock period isn't complete
	tx, err = client.Withdraw()
	if assert.Error(t, err) {
		assert.Equal(t, alwaysFailing, err.Error())
	}

	// Advance enough blocks for the unlock period to end
	for i := uint64(0); i < unlockDuration.Uint64(); i++ {
		backend.Commit()
	}

	tx, err = client.Withdraw()
	assert.Nil(t, err, "Failed to withdraw deposit")

}

var testPrivHex = "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"

func TestClient_RedeemTicket(t *testing.T) {

	key, _ := crypto.HexToECDSA(testPrivHex)
	auth := bind.NewKeyedTransactor(key)

	auth.Context = context.Background()

	backend := startSimulatedBackend(auth)

	client, err := createClientWithBackend(backend, auth)
	assert.Nil(t, err, "Failed to init client")

	// Approve ticketing contract to transfer funds
	_, err = client.ApproveTicketing(big.NewInt(100000))
	assert.Nil(t, err, "Failed to approve ticketing")

	backend.Commit()

	_, err = client.DepositEscrow(big.NewInt(100000), auth.From)
	assert.Nil(t, err, "Failed to deposit escrow")
	backend.Commit()

	receiverRandTmp, _ := hex.DecodeString("b10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6")
	receiverRandHash := [32]byte{}
	copy(receiverRandHash[:], receiverRandTmp)

	ticket := contracts.SyloTicketingTicket{
		Sender:           ethcommon.HexToAddress("0x970E8128AB834E8EAC17Ab8E3812F010678CF791"),
		Receiver:         ethcommon.HexToAddress("0x34D743d137a8cc298349F993b22B03Fea15c30c2"),
		ReceiverRandHash: receiverRandHash,
		FaceValue:        big.NewInt(1),
		WinProb:          new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1)), // 2^256-1
		ExpirationBlock:  big.NewInt(0),
		SenderNonce:      1,
	}

	sig, _ := hex.DecodeString("fe733162c570e2cb5cd9e0975110ea846e0cdba80c354344f6221d65ff9084ad29f37e486285023bb8c320ffe2c1e532635df485c4f3537993252f81fe943a2a00")
	receiverRand := big.NewInt(1)

	depositBefore, err := client.Deposits(ticket.Sender)
	assert.Nil(t, err, "Failed to get deposits")

	balanceBefore, err := client.BalanceOf(ticket.Receiver)
	assert.Nil(t, err, "Failed to get balance")

	tx, err := client.Redeem(ticket, receiverRand, sig)
	assert.Nil(t, err, "Failed to redeem ticket")

	backend.Commit()

	duration := 10 * time.Second
	_, err = client.CheckTxTimeout(tx, duration)
	assert.Nil(t, err, "Failed to confirm redeem")

	depositAfter, err := client.Deposits(ticket.Sender)
	assert.Nil(t, err, "Failed to get deposits")

	balanceAfter, err := client.BalanceOf(ticket.Receiver)
	assert.Nil(t, err, "Failed to get balance")

	assert.Equal(t, new(big.Int).Add(depositAfter.Escrow, ticket.FaceValue).Cmp(depositBefore.Escrow), 0, "Deposit should decrease")
	assert.Equal(t, new(big.Int).Add(balanceBefore, ticket.FaceValue).Cmp(balanceAfter), 0, "Balance should increase")
}

func TestClient_ReplayTicket(t *testing.T) {

	key, _ := crypto.HexToECDSA(testPrivHex)
	auth := bind.NewKeyedTransactor(key)

	auth.Context = context.Background()

	backend := startSimulatedBackend(auth)

	client, err := createClientWithBackend(backend, auth)
	assert.Nil(t, err, "Failed to init client")

	// Approve ticketing contract to transfer funds
	_, err = client.ApproveTicketing(big.NewInt(100000))
	assert.Nil(t, err, "Failed to approve ticketing")

	backend.Commit()

	_, err = client.DepositEscrow(big.NewInt(100000), auth.From)
	assert.Nil(t, err, "Failed to deposit escrow")
	backend.Commit()

	receiverRandTmp, _ := hex.DecodeString("b10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6")
	receiverRandHash := [32]byte{}
	copy(receiverRandHash[:], receiverRandTmp)

	ticket := contracts.SyloTicketingTicket{
		Sender:           ethcommon.HexToAddress("0x970E8128AB834E8EAC17Ab8E3812F010678CF791"),
		Receiver:         ethcommon.HexToAddress("0x34D743d137a8cc298349F993b22B03Fea15c30c2"),
		ReceiverRandHash: receiverRandHash,
		FaceValue:        big.NewInt(1),
		WinProb:          new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1)), // 2^256-1
		ExpirationBlock:  big.NewInt(0),
		SenderNonce:      1,
	}

	sig, _ := hex.DecodeString("fe733162c570e2cb5cd9e0975110ea846e0cdba80c354344f6221d65ff9084ad29f37e486285023bb8c320ffe2c1e532635df485c4f3537993252f81fe943a2a00")
	receiverRand := big.NewInt(1)

	tx, err := client.Redeem(ticket, receiverRand, sig)
	assert.Nil(t, err, "Failed to redeem ticket")

	backend.Commit()

	duration := 10 * time.Second
	_, err = client.CheckTxTimeout(tx, duration)
	assert.Nil(t, err, "Failed to confirm redeem")

	tx, err = client.Redeem(ticket, receiverRand, sig)
	if assert.Error(t, err) {
		// Transaction should always fail because ticket has already been used
		assert.Equal(t, alwaysFailing, err.Error())
	}
}

func TestClient_Unstake(t *testing.T) {

	key, _ := crypto.HexToECDSA(testPrivHex)
	auth := bind.NewKeyedTransactor(key)

	auth.Context = context.Background()

	backend := startSimulatedBackend(auth)

	client, err := createClientWithBackend(backend, auth)
	assert.Nil(t, err, "Failed to init client")

	// Approve ticketing contract to transfer funds
	_, err = client.ApproveDirectory(big.NewInt(100000))
	assert.Nil(t, err, "Failed to approve ticketing")

	backend.Commit()

	stakeAmount := big.NewInt(1000)

	tx, err := client.AddStake(stakeAmount, auth.From)
	assert.Nil(t, err, "Failed to add stake")

	backend.Commit()

	duration := 10 * time.Second
	_, err = client.CheckTxTimeout(tx, duration)
	assert.Nil(t, err, "Failed to confirm add stake")

	tx, err = client.UnlockStake(stakeAmount, auth.From)
	assert.Nil(t, err, "Failed to unlock stake")

	backend.Commit()

	_, err = client.Unstake(auth.From)
	if assert.Error(t, err) {
		assert.Equal(t, alwaysFailing, err.Error())
	}

	_, err = client.CheckTxTimeout(tx, duration)
	assert.Nil(t, err, "Failed to confirm unlcok stake")

	// Advance enough blocks for the unlock period to end
	for i := uint64(0); i < unlockDuration.Uint64(); i++ {
		backend.Commit()
	}

	balanceBefore, err := client.BalanceOf(auth.From);

	tx, err = client.Unstake(auth.From)
	assert.Nil(t, err, "Failed to unstake")

	backend.Commit()

	_, err = client.CheckTxTimeout(tx, duration)
	assert.Nil(t, err, "Failed to confirm unstake")

	// Check that unlocking state is reset
	unlocking, err := client.GetUnlockingStake(auth.From, auth.From);
	assert.Nil(t, err, "Should be able to get unlocking");

	assert.Zero(t, unlocking.Amount.Uint64(), "Unlocking should be cleared")
	assert.Zero(t, unlocking.UnlockAt.Uint64(), "Unlocking should be cleared")

	balanceAfter, err := client.BalanceOf(auth.From);

	// Check the token balance has increased
	assert.Equal(t, balanceAfter.Cmp(new(big.Int).Add(balanceBefore, stakeAmount)), 0, "Expected stake to be returned")

	// Should not be able to unstake again
	tx, err = client.Unstake(auth.From)

	if (assert.Error(t, err)) {
		assert.Equal(t, alwaysFailing, err.Error())
	}

}

func TestClient_CancelUnstaking(t *testing.T) {

	key, _ := crypto.HexToECDSA(testPrivHex)
	auth := bind.NewKeyedTransactor(key)

	auth.Context = context.Background()

	backend := startSimulatedBackend(auth)

	client, err := createClientWithBackend(backend, auth)
	assert.Nil(t, err, "Failed to init client")

	// Approve ticketing contract to transfer funds
	_, err = client.ApproveDirectory(big.NewInt(100000))
	assert.Nil(t, err, "Failed to approve ticketing")

	backend.Commit()

	tx, err := client.AddStake(big.NewInt(1000), auth.From)
	assert.Nil(t, err, "Failed to add stake")

	backend.Commit()

	duration := 10 * time.Second
	_, err = client.CheckTxTimeout(tx, duration)
	assert.Nil(t, err, "Failed to confirm add stake")

	tx, err = client.UnlockStake(big.NewInt(1000), auth.From)
	assert.Nil(t, err, "Failed to unlock stake")

	backend.Commit()

	_, err = client.CheckTxTimeout(tx, duration)
	assert.Nil(t, err, "Failed to confirm unlcok stake")

	tx, err = client.LockStake(big.NewInt(1000), auth.From)
	assert.Nil(t, err, "Should be able to lock stake")

	backend.Commit()

	tx, err = client.UnlockStake(big.NewInt(1000), auth.From)
	assert.Nil(t, err, "Failed to unlock stake")

	backend.Commit()

	// Advance enough blocks for the unlock period to end
	for i := uint64(0); i < unlockDuration.Uint64(); i++ {
		backend.Commit()
	}

	tx, err = client.LockStake(big.NewInt(1000), auth.From)
	assert.Nil(t, err, "Should be able to lock stake")

	backend.Commit()
}
