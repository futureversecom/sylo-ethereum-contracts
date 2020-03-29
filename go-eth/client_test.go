package eth

import (
	"context"
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
