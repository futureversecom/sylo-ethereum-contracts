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

func NewSimBackend(sim *backends.SimulatedBackend) SimBackend {
	return &simBackend{SimulatedBackend: sim}
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

func NewSimClients(opts []bind.TransactOpts) ([]Client, SimBackend, error) {
	var gasLimit uint64 = 50000000
	var addresses Addresses = Addresses{}

	if len(opts) < 1 {
		return nil, nil, errors.New("Please provide at least one option")
	}

	genisis := make(core.GenesisAlloc)

	for i := 0; i < len(opts); i++ {
		genisis[opts[i].From] = core.GenesisAccount{Balance: big.NewInt(1000000000000)}
	}

	sim := backends.NewSimulatedBackend(genisis, gasLimit)

	backend := NewSimBackend(sim)

	addresses.Token, _, _, _ = contracts.DeploySyloToken(&opts[0], backend)
	backend.Commit()
	addresses.Ticketing, _, _, _ = contracts.DeploySyloTicketing(&opts[0], backend, addresses.Token, big.NewInt(1))
	backend.Commit()
	addresses.Directory, _, _, _ = contracts.DeployDirectory(&opts[0], backend, addresses.Token, big.NewInt(1))
	backend.Commit()
	addresses.Listings, _, _, _ = contracts.DeployListings(&opts[0], backend)

	var clients []Client

	for i := 0; i < len(opts); i++ {
		client, err := NewClientWithBackend(
			addresses,
			backend,
			&opts[0],
		)
		clients = append(clients, client)
		if err != nil {
			return nil, nil, err
		}
	}

	return clients, backend, nil
}
