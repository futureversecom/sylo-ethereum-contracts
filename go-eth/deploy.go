package eth

import (
	"context"
	"fmt"
	"math/big"

	"github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

func DeployContracts(ctx context.Context, opts *bind.TransactOpts, backend Backend, unlockDuration *big.Int, baseLiveWinProb *big.Int) (addresses Addresses, err error) {
	// Deploying contracts can apparently panic if the transaction fails, so
	// we need to check for that.
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic during deployment of contracts: %w", r)
		}
	}()

	nonce, err := backend.NonceAt(ctx, opts.From, nil)
	if err != nil {
		return addresses, fmt.Errorf("could not get nonce: %w", err)
	}
	opts.Nonce = new(big.Int).SetUint64(nonce)

	// deploy Sylo token
	var deployTokenTx *types.Transaction
	addresses.Token, deployTokenTx, _, err = contracts.DeploySyloToken(opts, backend)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy sylo token: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// deploy staking manager
	var deployStakingManagerTx *types.Transaction
	var stakingManager *contracts.StakingManager
	addresses.StakingManager, deployStakingManagerTx, stakingManager, err = contracts.DeployStakingManager(opts, backend)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy staking manager: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// deploy price voting
	var deployPriceVotingTx *types.Transaction
	var priceVoting *contracts.PriceVoting
	addresses.PriceVoting, deployPriceVotingTx, priceVoting, err = contracts.DeployPriceVoting(opts, backend)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy price voting: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// deploy price manager
	var deployPriceManagerTx *types.Transaction
	var priceManager *contracts.PriceManager
	addresses.PriceManager, deployPriceManagerTx, priceManager, err = contracts.DeployPriceManager(opts, backend)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy price manager: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// deploy directory
	var deployDirectoryTx *types.Transaction
	var directory *contracts.Directory
	addresses.Directory, deployDirectoryTx, directory, err = contracts.DeployDirectory(opts, backend)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy directory: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// deploy listings
	var deployListingsTx *types.Transaction
	var listings *contracts.Listings
	addresses.Listings, deployListingsTx, listings, err = contracts.DeployListings(opts, backend)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy listings: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// deploy ticketing
	var deployTicketingTx *types.Transaction
	var ticketing *contracts.SyloTicketing
	addresses.Ticketing, deployTicketingTx, ticketing, err = contracts.DeploySyloTicketing(opts, backend)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy ticketing: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// wait for deployments
	WaitForReceipt(ctx, deployTokenTx, backend)
	WaitForReceipt(ctx, deployStakingManagerTx, backend)
	WaitForReceipt(ctx, deployPriceVotingTx, backend)
	WaitForReceipt(ctx, deployPriceManagerTx, backend)
	WaitForReceipt(ctx, deployDirectoryTx, backend)
	WaitForReceipt(ctx, deployListingsTx, backend)
	WaitForReceipt(ctx, deployTicketingTx, backend)

	// initialise staking manager
	var initStakingManagerTx *types.Transaction
	initStakingManagerTx, err = stakingManager.Initialize(opts, addresses.Token, addresses.RewardsManager, addresses.EpochsManager, unlockDuration)
	if err != nil {
		return addresses, fmt.Errorf("could not initialize staking: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// initialise price voting
	var initPriceVotingTx *types.Transaction
	initPriceVotingTx, err = priceVoting.Initialize(opts, addresses.StakingManager)
	if err != nil {
		return addresses, fmt.Errorf("could not initialize price: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// initialise price manager
	var initPriceManagerTx *types.Transaction
	initPriceManagerTx, err = priceManager.Initialize(opts, addresses.StakingManager, addresses.PriceVoting)
	if err != nil {
		return addresses, fmt.Errorf("could not initialize price manager contract: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// initialize directory
	var initDirectoryTx *types.Transaction
	initDirectoryTx, err = directory.Initialize(opts, addresses.StakingManager, addresses.RewardsManager)
	if err != nil {
		return addresses, fmt.Errorf("could not initialize directory: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// initialise listings
	var initListingsTx *types.Transaction
	initListingsTx, err = listings.Initialize(opts, 50)
	if err != nil {
		return addresses, fmt.Errorf("could not initialize listings: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// initialise ticketing
	var initTicketingTx *types.Transaction
	initTicketingTx, err = ticketing.Initialize(
		opts,
		addresses.Token,
		addresses.Listings,
		addresses.StakingManager,
		addresses.Directory,
		addresses.EpochsManager,
		addresses.RewardsManager,
		unlockDuration,
	)
	if err != nil {
		return addresses, fmt.Errorf("could not initialize ticketing: %w", err)
	}

	// wait for initializations
	WaitForReceipt(ctx, initPriceManagerTx, backend)
	WaitForReceipt(ctx, initStakingManagerTx, backend)
	WaitForReceipt(ctx, initPriceVotingTx, backend)
	WaitForReceipt(ctx, initDirectoryTx, backend)
	WaitForReceipt(ctx, initListingsTx, backend)
	WaitForReceipt(ctx, initTicketingTx, backend)

	opts.Nonce = nil
	return addresses, nil
}
