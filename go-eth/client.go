package eth

//go:generate abigen --abi ../abi/EpochsManager.abi --pkg manager --type EpochsManager --out contracts/epochs/manager/manager.go --bin ../bin/EpochsManager.bin
//go:generate abigen --abi ../abi/PriceManager.abi --pkg manager --type PriceManager --out contracts/payments/pricing/manager/manager.go --bin ../bin/PriceManager.bin
//go:generate abigen --abi ../abi/PriceVoting.abi --pkg voting --type PriceVoting --out contracts/payments/pricing/voting/voting.go --bin ../bin/PriceVoting.bin
//go:generate abigen --abi ../abi/TicketingParameters.abi --pkg parameters --type TicketingParameters --out contracts/payments/ticketing/parameters/parameters.go --bin ../bin/TicketingParameters.bin
//go:generate abigen --abi ../abi/RewardsManager.abi --pkg rewardsmanager --type RewardsManager --out contracts/payments/ticketing/rewardsmanager/rewardsmanager.go --bin ../bin/RewardsManager.bin
//go:generate abigen --abi ../abi/SyloTicketing.abi --pkg ticketing --type SyloTicketing --out contracts/payments/ticketing/ticketing.go --bin ../bin/SyloTicketing.bin
//go:generate abigen --abi ../abi/Directory.abi --pkg directory --type Directory --out contracts/staking/directory/directory.go --bin ../bin/Directory.bin
//go:generate abigen --abi ../abi/StakingManager.abi --pkg manager --type StakingManager --out contracts/staking/manager/manager.go --bin ../bin/StakingManager.bin
//go:generate abigen --abi ../abi/Listings.abi --pkg listings --type Listings --out contracts/listings/listings.go --bin ../bin/Listings.bin
//go:generate abigen --abi ../abi/SyloToken.abi --pkg token --type SyloToken --out contracts/token/token.go --bin ../bin/SyloToken.bin

import (
	"context"
	"fmt"
	"math/big"

	epochsmanager "github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts/epochs/manager"
	"github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts/listings"
	pricemanager "github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts/payments/pricing/manager"
	"github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts/payments/pricing/voting"
	"github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts/payments/ticketing"
	"github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts/payments/ticketing/parameters"
	"github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts/payments/ticketing/rewardsmanager"
	"github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts/staking/directory"
	stakingmanager "github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts/staking/manager"
	"github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts/token"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Unlocking struct {
	Amount   *big.Int
	UnlockAt *big.Int
}

type Client interface {
	Address() ethcommon.Address

	// Ticketing methods

	GetTicketHash(ticket ticketing.SyloTicketingTicket) ([32]byte, error)
	Deposits(account ethcommon.Address) (struct {
		Escrow   *big.Int
		Penalty  *big.Int
		UnlockAt *big.Int
	}, error)
	DepositEscrow(amount *big.Int, account ethcommon.Address) (*types.Transaction, error)
	DepositPenalty(amount *big.Int, account ethcommon.Address) (*types.Transaction, error)
	UnlockDeposits() (*types.Transaction, error)
	LockDeposits() (*types.Transaction, error)
	Redeem(ticket ticketing.SyloTicketingTicket, senderRand *big.Int, redeemerRand *big.Int, sig []byte) (*types.Transaction, error)
	Withdraw() (*types.Transaction, error)
	WithdrawTo(account ethcommon.Address) (*types.Transaction, error)

	// Token methods

	Allowance(owner ethcommon.Address, spender ethcommon.Address) (*big.Int, error)
	Approve(spender ethcommon.Address, amount *big.Int) (*types.Transaction, error)
	DecreaseAllowance(spender ethcommon.Address, subtractedValue *big.Int) (*types.Transaction, error)
	IncreaseAllowance(spender ethcommon.Address, addedValue *big.Int) (*types.Transaction, error)
	BalanceOf(account ethcommon.Address) (*big.Int, error)
	Transfer(recipient ethcommon.Address, amount *big.Int) (*types.Transaction, error)
	TransferFrom(sender ethcommon.Address, recipient ethcommon.Address, amount *big.Int) (*types.Transaction, error)

	// PriceManager methods

	CalculatePrices(sortedVotes []*big.Int) (*types.Transaction, error)

	// PriceVoting methods

	Vote(price *big.Int) (*types.Transaction, error)
	GetVotes() ([]ethcommon.Address, []*big.Int, error)
	WithdrawVote() (*types.Transaction, error)

	// StakingManager methods

	AddStake(amount *big.Int, stakee ethcommon.Address) (*types.Transaction, error)
	UnlockStake(amount *big.Int, stakee ethcommon.Address) (*types.Transaction, error)
	CancelUnlocking(amount *big.Int, stakee ethcommon.Address) (*types.Transaction, error)
	WithdrawStake(account ethcommon.Address) (*types.Transaction, error)
	GetAmountStaked(stakee ethcommon.Address) (*big.Int, error)
	GetUnlockingStake(staker ethcommon.Address, stakee ethcommon.Address) (Unlocking, error)
	GetCurrentStakerAmount(staker ethcommon.Address, stakee ethcommon.Address) (*big.Int, error)

	// Directory methods

	SetCurrentDirectory(epochId *big.Int) (*types.Transaction, error)
	JoinNextDirectory() (*types.Transaction, error)
	Scan(rand *big.Int) (ethcommon.Address, error)
	TransferDirectoryOwnership(newOwner ethcommon.Address) (*types.Transaction, error)

	// Listings methods

	SetListing(multiaddr string, minimumStakeAmount *big.Int) (*types.Transaction, error)
	GetListing(account ethcommon.Address) (listings.ListingsListing, error)

	// EpochsManager methods

	InitializeEpoch() (*types.Transaction, error)
	GetCurrentActiveEpoch() (epochsmanager.EpochsManagerEpoch, error)
	GetNextEpochId() (*big.Int, error)

	// RewardsManager methods

	GetUnclaimedNodeReward(stakee ethcommon.Address) (*big.Int, error)
	GetUnclaimedStakeReward(stakee ethcommon.Address) (*big.Int, error)
	GetRewardPoolActiveStake(epochId *big.Int, stakee ethcommon.Address) (*big.Int, error)
	InitializeNextRewardPool() (*types.Transaction, error)
	ClaimStakingRewards(stakee ethcommon.Address) (*types.Transaction, error)

	// Alias for Approve but uses the ticketingAddress or directoryAddress as the spender

	ApproveTicketing(amount *big.Int) (*types.Transaction, error)
	ApproveStakingManager(amount *big.Int) (*types.Transaction, error)

	// Utils

	LatestBlock() (*big.Int, error)
	CheckTx(ctx context.Context, tx *types.Transaction) (*big.Int, error)
	EthBalance(ctx context.Context) (*big.Int, error)
	SyloBalance(ctx context.Context) (*big.Int, error)
}

type Addresses struct {
	Token               ethcommon.Address
	Ticketing           ethcommon.Address
	TicketingParameters ethcommon.Address
	Directory           ethcommon.Address
	Listings            ethcommon.Address
	PriceManager        ethcommon.Address
	PriceVoting         ethcommon.Address
	StakingManager      ethcommon.Address
	EpochsManager       ethcommon.Address
	RewardsManager      ethcommon.Address
}

type client struct {
	address   ethcommon.Address
	contracts Addresses
	backend   Backend

	// Embedded contracts
	*parameters.TicketingParametersSession
	*ticketing.SyloTicketingSession
	*token.SyloTokenSession
	*directory.DirectorySession
	*listings.ListingsSession
	*pricemanager.PriceManagerSession
	*voting.PriceVotingSession
	*stakingmanager.StakingManagerSession
	*epochsmanager.EpochsManagerSession
	*rewardsmanager.RewardsManagerSession
}

func NewSyloPaymentsClient(
	contractAddrs Addresses,
	backend Backend,
	opts *bind.TransactOpts,
) (Client, error) {
	c := new(client)
	c.address = opts.From
	c.contracts = contractAddrs
	c.backend = backend

	// sylo token
	syloToken, err := token.NewSyloToken(c.contracts.Token, backend)
	if err != nil {
		return nil, err
	}
	c.SyloTokenSession = &token.SyloTokenSession{
		Contract:     syloToken,
		TransactOpts: *opts,
	}

	// ticketing parameters
	ticketingParamters, err := parameters.NewTicketingParameters(c.contracts.TicketingParameters, backend)
	if err != nil {
		return nil, err
	}
	c.TicketingParametersSession = &parameters.TicketingParametersSession{
		Contract:     ticketingParamters,
		TransactOpts: *opts,
	}

	// sylo ticketing
	syloTicketing, err := ticketing.NewSyloTicketing(c.contracts.Ticketing, backend)
	if err != nil {
		return nil, err
	}
	c.SyloTicketingSession = &ticketing.SyloTicketingSession{
		Contract:     syloTicketing,
		TransactOpts: *opts,
	}

	// staking directory
	stakingDirectory, err := directory.NewDirectory(c.contracts.Directory, backend)
	if err != nil {
		return nil, err
	}
	c.DirectorySession = &directory.DirectorySession{
		Contract:     stakingDirectory,
		TransactOpts: *opts,
	}

	// multiaddr listings
	multiaddrListings, err := listings.NewListings(c.contracts.Listings, backend)
	if err != nil {
		return nil, err
	}
	c.ListingsSession = &listings.ListingsSession{
		Contract:     multiaddrListings,
		TransactOpts: *opts,
	}

	// epochs manager
	epochsManager, err := epochsmanager.NewEpochsManager(c.contracts.EpochsManager, backend)
	if err != nil {
		return nil, err
	}
	c.EpochsManagerSession = &epochsmanager.EpochsManagerSession{
		Contract:     epochsManager,
		TransactOpts: *opts,
	}

	// price manager
	priceManager, err := pricemanager.NewPriceManager(c.contracts.PriceManager, backend)
	if err != nil {
		return nil, err
	}
	c.PriceManagerSession = &pricemanager.PriceManagerSession{
		Contract:     priceManager,
		TransactOpts: *opts,
	}

	// service price voting
	priceVoting, err := voting.NewPriceVoting(c.contracts.PriceVoting, backend)
	if err != nil {
		return nil, err
	}
	c.PriceVotingSession = &voting.PriceVotingSession{
		Contract:     priceVoting,
		TransactOpts: *opts,
	}

	// staking manager
	stakingManager, err := stakingmanager.NewStakingManager(contractAddrs.StakingManager, backend)
	if err != nil {
		return nil, err
	}
	c.StakingManagerSession = &stakingmanager.StakingManagerSession{
		Contract:     stakingManager,
		TransactOpts: *opts,
	}

	// rewards manager
	rewardsManager, err := rewardsmanager.NewRewardsManager(c.contracts.RewardsManager, backend)
	if err != nil {
		return nil, err
	}
	c.RewardsManagerSession = &rewardsmanager.RewardsManagerSession{
		Contract:     rewardsManager,
		TransactOpts: *opts,
	}

	return c, nil
}

func (c *client) Address() ethcommon.Address {
	return c.address
}

func (c *client) ApproveTicketing(amount *big.Int) (*types.Transaction, error) {
	return c.Approve(c.contracts.Ticketing, amount)
}

func (c *client) ApproveStakingManager(amount *big.Int) (*types.Transaction, error) {
	return c.Approve(c.contracts.StakingManager, amount)
}

func (c *client) GetAmountStaked(stakee ethcommon.Address) (*big.Int, error) {
	return c.GetStakeeTotalManagedStake(stakee)
}

func (c *client) Withdraw() (*types.Transaction, error) {
	return c.SyloTicketingSession.Withdraw()
}

func (c *client) WithdrawVote() (*types.Transaction, error) {
	return c.PriceVotingSession.Withdraw()
}

func (c *client) GetStakeKey(staker ethcommon.Address, stakee ethcommon.Address) ([32]byte, error) {
	return c.StakingManagerSession.GetKey(staker, stakee)
}

func (c *client) GetUnlockingStake(staker ethcommon.Address, stakee ethcommon.Address) (Unlocking, error) {
	key, err := c.StakingManagerSession.GetKey(staker, stakee)
	if err != nil {
		return Unlocking{}, err
	}
	return c.Unlockings(key)
}

func (c *client) TransferDirectoryOwnership(newOwner ethcommon.Address) (*types.Transaction, error) {
	return c.DirectorySession.TransferOwnership(newOwner)
}

func (c *client) LatestBlock() (*big.Int, error) {
	if c == nil {
		return nil, fmt.Errorf("client cannot be nil")
	}
	if c.backend == nil {
		return nil, fmt.Errorf("backend cannot be nil")
	}

	header, err := c.backend.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return header.Number, nil
}

func (c *client) CheckTx(ctx context.Context, tx *types.Transaction) (*big.Int, error) {
	if c == nil {
		return nil, fmt.Errorf("client cannot be nil")
	}

	if tx == nil {
		return nil, fmt.Errorf("transaction cannot be nil")
	}

	if c.backend == nil {
		return nil, fmt.Errorf("client backend cannot be nil")
	}

	receipt, err := bind.WaitMined(ctx, c.backend, tx)
	if err != nil {
		return nil, err
	}

	if receipt.Status == uint64(0) {
		return receipt.BlockNumber, fmt.Errorf("tx %v failed with status %v, %v", tx.Hash().Hex(), receipt.Status, receipt.PostState)
	}

	return receipt.BlockNumber, nil
}

func (c *client) EthBalance(ctx context.Context) (*big.Int, error) {
	if c == nil {
		return nil, fmt.Errorf("client cannot be nil")
	}

	if c.backend == nil {
		return nil, fmt.Errorf("client backend cannot be nil")
	}

	return c.backend.BalanceAt(ctx, c.Address(), nil)
}

func (c *client) SyloBalance(ctx context.Context) (*big.Int, error) {
	if c == nil {
		return nil, fmt.Errorf("client cannot be nil")
	}

	return c.BalanceOf(c.Address())
}

// ETH is set to 10^18
const ETH = int64(1e18)

// Eth calculates the `big.Int` for a specific ETH amount.
func Eth(amt int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(amt), new(big.Int).SetInt64(ETH))
}
