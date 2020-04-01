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

func startSimulatedBackend(auth *bind.TransactOpts) SimBackend {
	var gasLimit uint64 = 50000000

	genisis := make(core.GenesisAlloc)

	genisis[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000000)}

	sim := backends.NewSimulatedBackend(genisis, gasLimit)

	return newSimBackend(sim)
}

func deployContracts(auth *bind.TransactOpts, backend SimBackend) (ethcommon.Address, ethcommon.Address, error) {

	tokenAddress, tx, _, err := contracts.DeploySyloToken(auth, backend)
	if err != nil {
		return ethcommon.Address{}, ethcommon.Address{}, err
	}
	backend.Commit()

	_, err = backend.TransactionReceipt(auth.Context, tx.Hash())

	if err != nil {
		return ethcommon.Address{}, ethcommon.Address{}, errors.Wrap(err, "Failed to deploy token")
	}

	ticketingAddress, tx, _, err := contracts.DeploySyloTicketing(auth, backend, tokenAddress, big.NewInt(1))
	if err != nil {
		return ethcommon.Address{}, ethcommon.Address{}, err
	}

	_, err = backend.TransactionReceipt(auth.Context, tx.Hash())

	if err != nil {
		return ethcommon.Address{}, ethcommon.Address{}, errors.Wrap(err, "Failed to deploy ticketing")
	}

	backend.Commit()

	return tokenAddress, ticketingAddress, nil
}

func TestClient_CanBeCreated(t *testing.T) {

	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	backend := startSimulatedBackend(auth)

	tokenAddress, ticketingAddress, err := deployContracts(auth, backend)
	assert.Nil(t, err, "Failed to deploy contracts")

	if (tokenAddress == ethcommon.Address{}) {
		t.Error("Token address is empty")
	}

	if (ticketingAddress == ethcommon.Address{}) {
		t.Error("ticketingAddress address is empty")
	}

	_, err = NewClientWithBackend(tokenAddress, ticketingAddress, backend, auth)
	assert.Nil(t, err, "Failed to init client")
}

func TestClient_LatestBlock(t *testing.T) {

	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	backend := startSimulatedBackend(auth)

	tokenAddress, ticketingAddress, err := deployContracts(auth, backend)
	assert.Nil(t, err, "Failed to deploy contracts")

	client, err := NewClientWithBackend(tokenAddress, ticketingAddress, backend, auth)
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

	tokenAddress, ticketingAddress, err := deployContracts(auth, backend)
	assert.Nil(t, err, "Failed to deploy contracts")

	client, err := NewClientWithBackend(tokenAddress, ticketingAddress, backend, auth)
	assert.Nil(t, err, "Failed to init client")

	// Approve ticketing contract to transfer funds
	tx, err := client.Approve(ticketingAddress, big.NewInt(100000))
	assert.Nil(t, err, "Failed to approve ticketing")

	backend.Commit()

	duration, _ := time.ParseDuration("10s")
	_, err = client.CheckTx(tx, duration)
	if err != nil {
		assert.Nil(t, err, "Failed to confirm approve tx")
	}

	tx, err = client.DepositEscrow(big.NewInt(1))

	backend.Commit()

	_, err = client.CheckTx(tx, duration)
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

	tokenAddress, ticketingAddress, err := deployContracts(auth, backend)
	assert.Nil(t, err, "Failed to deploy contracts")

	client, err := NewClientWithBackend(tokenAddress, ticketingAddress, backend, auth)
	assert.Nil(t, err, "Failed to init client")

	// Approve ticketing contract to transfer funds
	tx, err := client.Approve(ticketingAddress, big.NewInt(100000))
	assert.Nil(t, err, "Failed to approve ticketing")

	backend.Commit()

	duration, _ := time.ParseDuration("10s")
	_, err = client.CheckTx(tx, duration)
	if err != nil {
		assert.Nil(t, err, "Failed to confirm approve tx")
	}

	tx, err = client.DepositPenalty(big.NewInt(1))

	backend.Commit()

	_, err = client.CheckTx(tx, duration)
	if err != nil {
		assert.Nil(t, err, "Failed to confirm deposit tx")
	}

	deposit, err := client.Deposits(auth.From)
	if err != nil {
		assert.Nil(t, err, "Failed to get deposits")
	}

	assert.Equal(t, deposit.Penalty.Cmp(big.NewInt(1)), 0, "Deposit doesn't match")
}

var testPrivHex = "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"

func TestClient_RedeemTicket(t *testing.T) {

	key, _ := crypto.HexToECDSA(testPrivHex);
	auth := bind.NewKeyedTransactor(key)

	auth.Context = context.Background()

	backend := startSimulatedBackend(auth)

	tokenAddress, ticketingAddress, err := deployContracts(auth, backend)
	assert.Nil(t, err, "Failed to deploy contracts")

	client, err := NewClientWithBackend(tokenAddress, ticketingAddress, backend, auth)
	assert.Nil(t, err, "Failed to init client")

	// Approve ticketing contract to transfer funds
	_, err = client.Approve(ticketingAddress, big.NewInt(100000))
	assert.Nil(t, err, "Failed to approve ticketing")

	backend.Commit()

	_, err = client.DepositEscrow(big.NewInt(100000))
	assert.Nil(t, err, "Failed to deposit escrow")
	backend.Commit()

	receiverRandTmp, _ := hex.DecodeString("b10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6")
	receiverRandHash := [32]byte{}
	copy(receiverRandHash[:], receiverRandTmp)

	ticket := contracts.SyloTicketingTicket{
		Sender: ethcommon.HexToAddress("0x970E8128AB834E8EAC17Ab8E3812F010678CF791"),
		Receiver: ethcommon.HexToAddress("0x34D743d137a8cc298349F993b22B03Fea15c30c2"),
		ReceiverRandHash: receiverRandHash,
		FaceValue: big.NewInt(1),
		WinProb: new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1)), // 2^256-1
		ExpirationBlock: big.NewInt(0),
		SenderNonce: 1,
	}

	sig, _ := hex.DecodeString("fe733162c570e2cb5cd9e0975110ea846e0cdba80c354344f6221d65ff9084ad29f37e486285023bb8c320ffe2c1e532635df485c4f3537993252f81fe943a2a00")
	receiverRand := big.NewInt(1)

	tx, err := client.Redeem(ticket, receiverRand, sig)
	assert.Nil(t, err, "Failed to redeem ticket")


	backend.Commit()

	duration, _ := time.ParseDuration("10s")
	_, err = client.CheckTx(tx, duration)
	assert.Nil(t, err, "Failed to confirm redeem")
}

func TestClient_ReplayTicket(t *testing.T) {

	key, _ := crypto.HexToECDSA(testPrivHex);
	auth := bind.NewKeyedTransactor(key)

	auth.Context = context.Background()

	backend := startSimulatedBackend(auth)

	tokenAddress, ticketingAddress, err := deployContracts(auth, backend)
	assert.Nil(t, err, "Failed to deploy contracts")

	client, err := NewClientWithBackend(tokenAddress, ticketingAddress, backend, auth)
	assert.Nil(t, err, "Failed to init client")

	// Approve ticketing contract to transfer funds
	_, err = client.Approve(ticketingAddress, big.NewInt(100000))
	assert.Nil(t, err, "Failed to approve ticketing")

	backend.Commit()

	_, err = client.DepositEscrow(big.NewInt(100000))
	assert.Nil(t, err, "Failed to deposit escrow")
	backend.Commit()

	receiverRandTmp, _ := hex.DecodeString("b10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6")
	receiverRandHash := [32]byte{}
	copy(receiverRandHash[:], receiverRandTmp)

	ticket := contracts.SyloTicketingTicket{
		Sender: ethcommon.HexToAddress("0x970E8128AB834E8EAC17Ab8E3812F010678CF791"),
		Receiver: ethcommon.HexToAddress("0x34D743d137a8cc298349F993b22B03Fea15c30c2"),
		ReceiverRandHash: receiverRandHash,
		FaceValue: big.NewInt(1),
		WinProb: new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1)), // 2^256-1
		ExpirationBlock: big.NewInt(0),
		SenderNonce: 1,
	}

	sig, _ := hex.DecodeString("fe733162c570e2cb5cd9e0975110ea846e0cdba80c354344f6221d65ff9084ad29f37e486285023bb8c320ffe2c1e532635df485c4f3537993252f81fe943a2a00")
	receiverRand := big.NewInt(1)

	tx, err := client.Redeem(ticket, receiverRand, sig)
	assert.Nil(t, err, "Failed to redeem ticket")


	backend.Commit()

	duration, _ := time.ParseDuration("10s")
	_, err = client.CheckTx(tx, duration)
	assert.Nil(t, err, "Failed to confirm redeem")

	tx, err = client.Redeem(ticket, receiverRand, sig)
	if assert.Error(t, err) {
		// Transaction should always fail because ticket has already been used
		assert.Equal(t, "failed to estimate gas needed: gas required exceeds allowance or always failing transaction", err.Error())
	}
}
