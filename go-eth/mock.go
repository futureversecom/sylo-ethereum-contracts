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

var errNotImplemented = errors.New("not implemented")

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

func (b *simBackend) ChainID(ctx context.Context) (*big.Int, error) {
	return big.NewInt(1337), nil
}

func (b *simBackend) PendingBalanceAt(ctx context.Context, account ethcommon.Address) (*big.Int, error) {
	return nil, errNotImplemented
}

func (b *simBackend) PendingStorageAt(ctx context.Context, account ethcommon.Address, key ethcommon.Hash) ([]byte, error) {
	// Just get the current state
	return b.StorageAt(ctx, account, key, nil)
}

func (b *simBackend) PendingTransactionCount(ctx context.Context) (uint, error) {
	return 0, errNotImplemented
}

func (b *simBackend) FaucetEth(ctx context.Context, from ethcommon.Address, to ethcommon.Address, signerKey *ecdsa.PrivateKey, amount *big.Int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	nonce, err := b.SimulatedBackend.PendingNonceAt(ctx, from)
	if err != nil {
		return fmt.Errorf("could not get pending nonce: %v", err)
	}

	gasLimit := uint64(21000) // in units
	gasPrice := big.NewInt(52987961)

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
	var winProb = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))
	var expiredWinProb = big.NewInt(10000)
	var decayRate = uint16(8000)
	var ticketDuration = big.NewInt(100)
	var epochsDuration = big.NewInt(80000)

	if len(opts) < 1 {
		return nil, nil, errors.New("must provide at least one option")
	}

	genisis := make(core.GenesisAlloc)

	for i := 0; i < len(opts); i++ {
		genisis[opts[i].From] = core.GenesisAccount{Balance: big.NewInt(1000000000000)}
	}

	sim := backends.NewSimulatedBackend(genisis, gasLimit)

	backend := NewSimBackend(sim)

	addresses.Token, _, _, _ = contracts.DeploySyloToken(&opts[0], backend)
	backend.Commit()

	var stakingManager *contracts.StakingManager
	addresses.StakingManager, _, stakingManager, _ = contracts.DeployStakingManager(&opts[0], backend)

	var directory *contracts.Directory
	addresses.Directory, _, directory, _ = contracts.DeployDirectory(&opts[0], backend)

	var epochsManager *contracts.EpochsManager
	addresses.EpochsManager, _, epochsManager, _ = contracts.DeployEpochsManager(&opts[0], backend)

	var rewardsManager *contracts.RewardsManager
	addresses.RewardsManager, _, rewardsManager, _ = contracts.DeployRewardsManager(&opts[0], backend)

	_, err := stakingManager.Initialize(&opts[0], addresses.Token, addresses.StakingManager, addresses.EpochsManager, big.NewInt(1))
	if err != nil {
		return nil, nil, fmt.Errorf("could not initialise listing: %w", err)
	}
	backend.Commit()

	_, err = directory.Initialize(&opts[0], addresses.StakingManager, addresses.RewardsManager)
	if err != nil {
		return nil, nil, fmt.Errorf("could not initialise directory: %w", err)
	}
	backend.Commit()

	_, err = epochsManager.Initialize(&opts[0], addresses.Directory, addresses.Listings, addresses.TicketingParameters, epochsDuration)
	if err != nil {
		return nil, nil, fmt.Errorf("could not initialise epochsManager: %w", err)
	}
	backend.Commit()

	_, err = rewardsManager.Initialize(&opts[0], addresses.Token, addresses.StakingManager, addresses.EpochsManager)
	if err != nil {
		return nil, nil, fmt.Errorf("could not initialise rewardsManager: %w", err)
	}
	backend.Commit()

	var priceVoting *contracts.PriceVoting
	addresses.PriceVoting, _, priceVoting, _ = contracts.DeployPriceVoting(&opts[0], backend)
	_, err = priceVoting.Initialize(&opts[0], addresses.StakingManager)
	if err != nil {
		return nil, nil, fmt.Errorf("could not initialise listing: %w", err)
	}
	backend.Commit()

	var priceManager *contracts.PriceManager
	addresses.PriceManager, _, priceManager, _ = contracts.DeployPriceManager(&opts[0], backend)
	_, err = priceManager.Initialize(&opts[0], addresses.StakingManager, addresses.PriceVoting)
	if err != nil {
		return nil, nil, fmt.Errorf("could not initialise price manager: %w", err)
	}
	backend.Commit()

	var listings *contracts.Listings
	addresses.Listings, _, listings, _ = contracts.DeployListings(&opts[0], backend)
	_, err = listings.Initialize(&opts[0], 50)
	if err != nil {
		return nil, nil, fmt.Errorf("could not initialise listing: %w", err)
	}
	backend.Commit()

	var ticketingParameters *contracts.TicketingParameters
	addresses.TicketingParameters, _, ticketingParameters, _ = contracts.DeployTicketingParameters(&opts[0], backend)
	_, err = ticketingParameters.Initialize(&opts[0], big.NewInt(1), winProb, expiredWinProb, decayRate, ticketDuration)
	if err != nil {
		return nil, nil, fmt.Errorf("could not initialise ticketingParameters: %w", err)
	}
	backend.Commit()

	var ticketing *contracts.SyloTicketing
	addresses.Ticketing, _, ticketing, _ = contracts.DeploySyloTicketing(&opts[0], backend)
	_, err = ticketing.Initialize(&opts[0], addresses.Token, addresses.Listings, addresses.StakingManager, addresses.Directory, addresses.EpochsManager, addresses.RewardsManager, big.NewInt(1))
	if err != nil {
		return nil, nil, fmt.Errorf("could not initialise ticketing: %w", err)
	}

	var clients []Client

	for i := 0; i < len(opts); i++ {
		client, err := NewSyloPaymentsClient(
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
