package eth

import (
	"context"
	"errors"
	"math/big"

	"github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
)

type SimBackend interface {
	Backend
	Commit()
}

type simBackend struct {
	*backends.SimulatedBackend
}

func newSimBackend(sim *backends.SimulatedBackend) SimBackend {

	return &simBackend{
		SimulatedBackend: sim,
	}
}

func (b *simBackend) PendingBalanceAt(ctx context.Context, account ethcommon.Address) (*big.Int, error) {
	return nil, errors.New("Not implemented")
}

func (b *simBackend) PendingStorageAt(ctx context.Context, account ethcommon.Address, key ethcommon.Hash) ([]byte, error) {
	// Just get the current state
	return b.StorageAt(ctx, account, key, nil)
}

func (b *simBackend) PendingTransactionCount(ctx context.Context) (uint, error) {
	return 0, errors.New("Not implemented")
}

func NewSimClient(auth *bind.TransactOpts) (Client, SimBackend, error) {
	var gasLimit uint64 = 50000000

	genisis := make(core.GenesisAlloc)

	genisis[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000000)}

	sim := backends.NewSimulatedBackend(genisis, gasLimit)

	backend := newSimBackend(sim)

	tokenAddress, _, _, _ := contracts.DeploySyloToken(auth, backend)
	backend.Commit()
	ticketingAddress, _, _, _ := contracts.DeploySyloTicketing(auth, backend, tokenAddress, big.NewInt(1))
	backend.Commit()

	client, err := NewClientWithBackend(tokenAddress, ticketingAddress, backend, auth)
	if err != nil {
		return nil, nil, err
	}

	return client, backend, nil
}
