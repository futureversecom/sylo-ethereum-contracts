package eth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

type SimBackend interface {
	Backend
	Commit()
	FaucetEth(ctx context.Context, from ethcommon.Address, to ethcommon.Address, signerKey *ecdsa.PrivateKey, amount *big.Int) error
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

func (b *simBackend) FaucetEth(ctx context.Context, from ethcommon.Address, to ethcommon.Address, signerKey *ecdsa.PrivateKey, amount *big.Int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	nonce, err := b.SimulatedBackend.PendingNonceAt(ctx, from)
	if err != nil {
		return fmt.Errorf("could not get pending nonce: %v", err)
	}

	gasLimit := uint64(21000) // in units
	gasPrice, err := b.SimulatedBackend.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("could not get suggested gas price: %v", err)
	}

	var data []byte
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, signerKey)
	if err != nil {
		return fmt.Errorf("could not sign transaction: %v", err)
	}

	err = b.SimulatedBackend.SendTransaction(ctx, signedTx)
	if err != nil {
		return fmt.Errorf("could not send transaction: %v", err)
	}
	b.Commit()

	return nil
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
