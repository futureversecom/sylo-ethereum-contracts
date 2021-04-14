package eth_test

import (
	"context"
	"encoding/hex"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/dn3010/sylo-ethereum-contracts/go-eth"
	"github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

var unlockDuration = big.NewInt(10)

func startSimulatedBackend(auth *bind.TransactOpts) eth.SimBackend {
	var gasLimit uint64 = 50000000

	genisis := make(core.GenesisAlloc)

	genisis[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000000)}

	sim := backends.NewSimulatedBackend(genisis, gasLimit)

	return eth.NewSimBackend(sim)
}

func deployContracts(t *testing.T, auth *bind.TransactOpts, backend eth.SimBackend) eth.Addresses {

	var addresses eth.Addresses
	var err error
	var tx *types.Transaction

	// Deploying contracts can apparently panic if the transaction fails, so
	// we need to check for that.
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panic during deployment of contracts: %v", r)
		}
	}()

	// deploy Sylo token
	addresses.Token, tx, _, err = contracts.DeploySyloToken(auth, backend)
	if err != nil {
		t.Fatalf("could not deploy sylo token: %v", err)
	}
	backend.Commit()
	_, err = backend.TransactionReceipt(auth.Context, tx.Hash())
	if err != nil {
		t.Fatalf("could not get transaction receipt: %v", err)
	}

	// deploy ticketing
	addresses.Ticketing, tx, _, err = contracts.DeploySyloTicketing(auth, backend, addresses.Token, unlockDuration)
	if err != nil {
		t.Fatalf("could not deploy ticketing: %v", err)
	}
	backend.Commit()
	_, err = backend.TransactionReceipt(auth.Context, tx.Hash())
	if err != nil {
		t.Fatalf("could not get transaction receipt: %v", err)
	}

	// deploy directory
	addresses.Directory, tx, _, err = contracts.DeployDirectory(auth, backend, addresses.Token, unlockDuration)
	if err != nil {
		t.Fatalf("could not deploy directory: %v", err)
	}
	backend.Commit()
	_, err = backend.TransactionReceipt(auth.Context, tx.Hash())
	if err != nil {
		t.Fatalf("could not get transaction receipt: %v", err)
	}

	// deploy listing
	addresses.Listings, tx, _, err = contracts.DeployListings(auth, backend)
	if err != nil {
		t.Fatalf("could not deploy listing: %v", err)
	}
	backend.Commit()
	_, err = backend.TransactionReceipt(auth.Context, tx.Hash())
	if err != nil {
		t.Fatalf("could not get transaction receipt: %v", err)
	}

	return addresses
}

func createClientWithBackend(t *testing.T, backend eth.SimBackend, auth *bind.TransactOpts) (eth.Client, error) {
	addresses := deployContracts(t, auth, backend)
	return eth.NewClientWithBackend(addresses, backend, auth)
}

func TestClient_CanBeCreated(t *testing.T) {

	key, _ := crypto.GenerateKey()
	chainID := big.NewInt(1337)
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		t.Fatalf("could not create transaction signer: %v", err)
	}

	backend := startSimulatedBackend(auth)
	addresses := deployContracts(t, auth, backend)

	if (addresses.Token == ethcommon.Address{}) {
		t.Error("Token address is empty")
	}

	if (addresses.Ticketing == ethcommon.Address{}) {
		t.Error("ticketingAddress address is empty")
	}

	_, err = eth.NewClientWithBackend(addresses, backend, auth)
	if err != nil {
		t.Fatalf("could not create client: %v", err)
	}
}

func TestClient_LatestBlock(t *testing.T) {

	key, _ := crypto.GenerateKey()
	chainID := big.NewInt(1337)
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		t.Fatalf("could not create transaction signer: %v", err)
	}

	backend := startSimulatedBackend(auth)

	client, err := createClientWithBackend(t, backend, auth)
	if err != nil {
		t.Fatalf("could not create client: %v", err)
	}

	blockNumber, err := client.LatestBlock()
	assert.Nil(t, err, "Failed to get latest block")

	assert.GreaterOrEqual(t, blockNumber.Cmp(big.NewInt(0)), 0)
}

func TestClient_DepositEscrow(t *testing.T) {

	key, _ := crypto.GenerateKey()
	chainID := big.NewInt(1337)
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		t.Fatalf("could not create transaction signer: %v", err)
	}

	auth.Context = context.Background()

	backend := startSimulatedBackend(auth)

	client, err := createClientWithBackend(t, backend, auth)
	if err != nil {
		t.Fatalf("could not create client: %v", err)
	}

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
	if err != nil {
		t.Fatalf("could not deposit escrow: %v", err)
	}

	backend.Commit()

	_, err = client.CheckTxTimeout(tx, duration)
	if err != nil {
		t.Fatalf("could not confirm deposit transaction: %v", err)
	}

	deposit, err := client.Deposits(auth.From)
	if err != nil {
		t.Fatalf("could not get deposits: %v", err)
	}
	if !bigIntsEqual(deposit.Escrow, big.NewInt(1)) {
		t.Fatalf("escrow deposit does not match: got %v: expected %v", deposit.Escrow, big.NewInt(1))
	}
}

func TestClient_DepositPenalty(t *testing.T) {

	key, _ := crypto.GenerateKey()
	chainID := big.NewInt(1337)
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		t.Fatalf("could not create transaction signer: %v", err)
	}

	auth.Context = context.Background()

	backend := startSimulatedBackend(auth)

	client, err := createClientWithBackend(t, backend, auth)
	if err != nil {
		t.Fatalf("could not create client: %v", err)
	}

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
	if err != nil {
		t.Fatalf("could not deposit penalty: %v", err)
	}

	backend.Commit()

	_, err = client.CheckTxTimeout(tx, duration)
	if err != nil {
		assert.Nil(t, err, "Failed to confirm deposit tx")
	}

	deposit, err := client.Deposits(auth.From)
	if err != nil {
		assert.Nil(t, err, "Failed to get deposits")
	}
	if !bigIntsEqual(deposit.Penalty, big.NewInt(1)) {
		t.Fatalf("penalty deposit does not match: got %v: expected %v", deposit.Penalty, big.NewInt(1))
	}
}

func TestClient_WithdrawTicketing(t *testing.T) {
	key, _ := crypto.GenerateKey()
	chainID := big.NewInt(1337)
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		t.Fatalf("could not create transaction signer: %v", err)
	}

	auth.Context = context.Background()

	backend := startSimulatedBackend(auth)

	client, err := createClientWithBackend(t, backend, auth)
	if err != nil {
		t.Fatalf("could not create client: %v", err)
	}

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
	if err != nil {
		t.Fatalf("could not check transaction: %v", err)
	}

	_, err = client.Withdraw()
	if err == nil {
		t.Fatalf("expected error because unlock period isn't complete")
	}
	if !strings.HasSuffix(err.Error(), "Unlock period not complete") {
		t.Fatalf("could not withdraw: %v", err)
	}

	// Advance enough blocks for the unlock period to end
	for i := uint64(0); i < unlockDuration.Uint64(); i++ {
		backend.Commit()
	}

	_, err = client.Withdraw()
	assert.Nil(t, err, "Failed to withdraw deposit")

}

var testPrivHex = "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"

func TestClient_RedeemTicket(t *testing.T) {

	key, _ := crypto.HexToECDSA(testPrivHex)
	chainID := big.NewInt(1337)
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		t.Fatalf("could not create transaction signer: %v", err)
	}

	auth.Context = context.Background()

	backend := startSimulatedBackend(auth)

	client, err := createClientWithBackend(t, backend, auth)
	if err != nil {
		t.Fatalf("could not create client: %v", err)
	}

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
	chainID := big.NewInt(1337)
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		t.Fatalf("could not create transaction signer: %v", err)
	}

	auth.Context = context.Background()

	backend := startSimulatedBackend(auth)

	client, err := createClientWithBackend(t, backend, auth)
	if err != nil {
		t.Fatalf("could not create client: %v", err)
	}

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

	_, err = client.Redeem(ticket, receiverRand, sig)
	if err == nil {
		t.Fatalf("expected error because ticket has already been used")
	}
	if !strings.HasSuffix(err.Error(), "Ticket already redeemed") {
		t.Fatalf("could not redeem: %v", err)
	}
}

func TestClient_Unstake(t *testing.T) {

	key, _ := crypto.HexToECDSA(testPrivHex)
	chainID := big.NewInt(1337)
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		t.Fatalf("could not create transaction signer: %v", err)
	}

	auth.Context = context.Background()

	backend := startSimulatedBackend(auth)

	client, err := createClientWithBackend(t, backend, auth)
	if err != nil {
		t.Fatalf("could not create client: %v", err)
	}

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
	if err == nil {
		t.Fatalf("expected error because stake not yet unlocked")
	}
	if !strings.HasSuffix(err.Error(), "Stake not yet unlocked") {
		t.Fatalf("could not unstake: %v", err)
	}

	_, err = client.CheckTxTimeout(tx, duration)
	assert.Nil(t, err, "Failed to confirm unlcok stake")

	// Advance enough blocks for the unlock period to end
	for i := uint64(0); i < unlockDuration.Uint64(); i++ {
		backend.Commit()
	}

	balanceBefore, err := client.BalanceOf(auth.From)
	if err != nil {
		t.Fatalf("could not check balance: %v", err)
	}

	tx, err = client.Unstake(auth.From)
	assert.Nil(t, err, "Failed to unstake")

	backend.Commit()

	_, err = client.CheckTxTimeout(tx, duration)
	assert.Nil(t, err, "Failed to confirm unstake")

	// Check that unlocking state is reset
	unlocking, err := client.GetUnlockingStake(auth.From, auth.From)
	assert.Nil(t, err, "Should be able to get unlocking")

	assert.Zero(t, unlocking.Amount.Uint64(), "Unlocking should be cleared")
	assert.Zero(t, unlocking.UnlockAt.Uint64(), "Unlocking should be cleared")

	balanceAfter, err := client.BalanceOf(auth.From)
	if err != nil {
		t.Fatalf("could not check balance: %v", err)
	}

	// Check the token balance has increased
	assert.Equal(t, balanceAfter.Cmp(new(big.Int).Add(balanceBefore, stakeAmount)), 0, "Expected stake to be returned")

	// Should not be able to unstake again
	_, err = client.Unstake(auth.From)
	if err == nil {
		t.Fatalf("expected error because should not be able to unstake again")
	}
	if !strings.HasSuffix(err.Error(), "No amount to unlock") {
		t.Fatalf("could not unstake: %v", err)
	}

}

func TestClient_CancelUnstaking(t *testing.T) {

	key, _ := crypto.HexToECDSA(testPrivHex)
	chainID := big.NewInt(1337)
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		t.Fatalf("could not create transaction signer: %v", err)
	}

	auth.Context = context.Background()

	backend := startSimulatedBackend(auth)

	client, err := createClientWithBackend(t, backend, auth)
	if err != nil {
		t.Fatalf("could not create client: %v", err)
	}

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

	_, err = client.LockStake(big.NewInt(1000), auth.From)
	assert.Nil(t, err, "Should be able to lock stake")

	backend.Commit()

	_, err = client.UnlockStake(big.NewInt(1000), auth.From)
	assert.Nil(t, err, "Failed to unlock stake")

	backend.Commit()

	// Advance enough blocks for the unlock period to end
	for i := uint64(0); i < unlockDuration.Uint64(); i++ {
		backend.Commit()
	}

	_, err = client.LockStake(big.NewInt(1000), auth.From)
	assert.Nil(t, err, "Should be able to lock stake")

	backend.Commit()
}

func bigIntsEqual(x *big.Int, y *big.Int) bool {
	return x.Cmp(y) == 0
}
