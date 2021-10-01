package eth

import (
	"context"
	"fmt"
	"math/big"

	"github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

type ContractParameters struct {
	// Listing parameters
	DefaultPayoutPercentage uint16

	// Deposit parameters
	UnlockDuration *big.Int

	// Ticketing parameters
	FaceValue       *big.Int
	BaseLiveWinProb *big.Int
	ExpiredWinProb  *big.Int
	DecayRate       uint16
	TicketDuration  *big.Int
}

func DeployContracts(ctx context.Context, opts *bind.TransactOpts, backend Backend, contractParams *ContractParameters) (addresses Addresses, err error) {
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

	// deploy epochs manager
	var deployEpochsManagerTx *types.Transaction
	var epochsManager *contracts.EpochsManager
	addresses.EpochsManager, deployEpochsManagerTx, epochsManager, err = contracts.DeployEpochsManager(opts, backend)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy epochs manager: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// deploy rewards manager
	var deployRewardsManagerTx *types.Transaction
	var rewardsManager *contracts.RewardsManager
	addresses.RewardsManager, deployRewardsManagerTx, rewardsManager, err = contracts.DeployRewardsManager(opts, backend)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy rewards manager: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// deploy ticketing parameters
	var deployTicketingParametersTx *types.Transaction
	var ticketingParameters *contracts.TicketingParameters
	addresses.TicketingParameters, deployTicketingParametersTx, ticketingParameters, err = contracts.DeployTicketingParameters(opts, backend)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy ticketing parameters: %w", err)
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
	WaitForReceipt(ctx, deployEpochsManagerTx, backend)
	WaitForReceipt(ctx, deployRewardsManagerTx, backend)
	WaitForReceipt(ctx, deployTicketingTx, backend)
	WaitForReceipt(ctx, deployTicketingParametersTx, backend)

	// initialise staking manager
	var initStakingManagerTx *types.Transaction
	initStakingManagerTx, err = stakingManager.Initialize(opts, addresses.Token, addresses.RewardsManager, addresses.EpochsManager, contractParams.UnlockDuration)
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
	initListingsTx, err = listings.Initialize(opts, contractParams.DefaultPayoutPercentage)
	if err != nil {
		return addresses, fmt.Errorf("could not initialize listings: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// initialise epochs manager
	var initEpochsManagerTx *types.Transaction
	initEpochsManagerTx, err = epochsManager.Initialize(opts, addresses.Directory, addresses.Listings, addresses.TicketingParameters, big.NewInt(1))
	if err != nil {
		return addresses, fmt.Errorf("could not initialize epochs manager: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// initialise rewards manager
	var initRewardsManagerTx *types.Transaction
	initRewardsManagerTx, err = rewardsManager.Initialize(opts, addresses.Token, addresses.StakingManager, addresses.EpochsManager)
	if err != nil {
		return addresses, fmt.Errorf("could not initialize epochs manager: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// initialise ticketing parameters
	var initTicketingParamtersTx *types.Transaction
	initTicketingParamtersTx, err = ticketingParameters.Initialize(
		opts,
		contractParams.FaceValue,
		contractParams.BaseLiveWinProb,
		contractParams.ExpiredWinProb,
		contractParams.DecayRate,
		contractParams.TicketDuration,
	)
	if err != nil {
		return addresses, fmt.Errorf("could not initialize ticketing parameters")
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
		contractParams.UnlockDuration,
	)
	if err != nil {
		return addresses, fmt.Errorf("could not initialize ticketing: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// wait for initializations
	WaitForReceipt(ctx, initPriceManagerTx, backend)
	WaitForReceipt(ctx, initStakingManagerTx, backend)
	WaitForReceipt(ctx, initPriceVotingTx, backend)
	WaitForReceipt(ctx, initDirectoryTx, backend)
	WaitForReceipt(ctx, initListingsTx, backend)
	WaitForReceipt(ctx, initEpochsManagerTx, backend)
	WaitForReceipt(ctx, initRewardsManagerTx, backend)
	WaitForReceipt(ctx, initTicketingParamtersTx, backend)
	WaitForReceipt(ctx, initTicketingTx, backend)

	// add manager to rewards contract
	addTicketingManagerTx, err := rewardsManager.AddManager(opts, addresses.Ticketing)
	if err != nil {
		return addresses, fmt.Errorf("could not add ticketing as manager")
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	addStakingManagerTx, err := rewardsManager.AddManager(opts, addresses.StakingManager)
	if err != nil {
		return addresses, fmt.Errorf("could not add staking as manager")
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	WaitForReceipt(ctx, addTicketingManagerTx, backend)
	WaitForReceipt(ctx, addStakingManagerTx, backend)

	opts.Nonce = nil
	return addresses, nil
}
