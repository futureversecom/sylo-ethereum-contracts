package eth

//go:generate abigen --abi ../abi/SyloTicketing.abi --pkg contracts --type SyloTicketing --out contracts/ticketing.go --bin ../bin/SyloTicketing.bin
//go:generate abigen --abi ../abi/TicketingParameters.abi --pkg contracts --type TicketingParameters --out contracts/ticketing_parameters.go --bin ../bin/TicketingParameters.bin
//go:generate abigen --abi ../abi/EpochsManager.abi --pkg contracts --type EpochsManager --out contracts/epochs_manager.go --bin ../bin/EpochsManager.bin
//go:generate abigen --abi ../abi/SyloToken.abi --pkg contracts --type SyloToken --out contracts/token.go --bin ../bin/SyloToken.bin
//go:generate abigen --abi ../abi/Directory.abi --pkg contracts --type Directory --out contracts/directory.go --bin ../bin/Directory.bin
//go:generate abigen --abi ../abi/Listings.abi --pkg contracts --type Listings --out contracts/listings.go --bin ../bin/Listings.bin
//go:generate abigen --abi ../abi/PriceManager.abi --pkg contracts --type PriceManager --out contracts/price_manager.go --bin ../bin/PriceManager.bin
//go:generate abigen --abi ../abi/PriceVoting.abi --pkg contracts --type PriceVoting --out contracts/price_voting.go --bin ../bin/PriceVoting.bin
//go:generate abigen --abi ../abi/StakingManager.abi --pkg contracts --type StakingManager --out contracts/staking_manager.go --bin ../bin/StakingManager.bin
//go:generate abigen --abi ../abi/RewardsManager.abi --pkg contracts --type RewardsManager --out contracts/rewards_manager.go --bin ../bin/RewardsManager.bin

import (
	"context"
	"fmt"
	"math/big"

	"github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Unlocking struct {
	Amount   *big.Int
	UnlockAt *big.Int
}

type Client interface {
	Address() ethcommon.Address

	// Ticketing methods

	GetTicketHash(ticket contracts.SyloTicketingTicket) ([32]byte, error)
	Deposits(account ethcommon.Address) (struct {
		Escrow   *big.Int
		Penalty  *big.Int
		UnlockAt *big.Int
	}, error)
	DepositEscrow(amount *big.Int, account ethcommon.Address) (*types.Transaction, error)
	DepositPenalty(amount *big.Int, account ethcommon.Address) (*types.Transaction, error)
	UnlockDeposits() (*types.Transaction, error)
	LockDeposits() (*types.Transaction, error)
	Redeem(ticket contracts.SyloTicketingTicket, senderRand *big.Int, redeemerRand *big.Int, sig []byte) (*types.Transaction, error)
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
	GetStakeeTotalStake(stakee ethcommon.Address) (*big.Int, error)
	GetStakerAmount(staker ethcommon.Address, stakee ethcommon.Address, blockNumber *big.Int) (*big.Int, error)

	// Directory methods
	SetCurrentDirectory(epochId [32]byte) (*types.Transaction, error)
	JoinDirectory(epochId [32]byte) (*types.Transaction, error)
	Scan(rand *big.Int) (ethcommon.Address, error)
	TransferDirectoryOwnership(newOwner ethcommon.Address) (*types.Transaction, error)

	// Listings methods

	SetListing(multiaddr string, minimumStakeAmount *big.Int) (*types.Transaction, error)
	GetListing(account ethcommon.Address) (contracts.ListingsListing, error)

	// EpochsManager methods

	InitializeEpoch() (*types.Transaction, error)
	GetCurrentActiveEpoch() (contracts.EpochsManagerEpoch, error)
	GetEpochId(*big.Int) ([32]byte, error)
	GetNextEpochId() ([32]byte, error)

	// RewardsManager methods
	GetRewardPoolBalance(epochId [32]byte, stakee ethcommon.Address) (*big.Int, error)
	GetRewardPoolStake(epochId [32]byte, stakee ethcommon.Address) (*big.Int, error)
	GetDelegatorOwedAmount(epochId [32]byte, stakee ethcommon.Address, staker ethcommon.Address) (*big.Int, error)
	IncrementRewardPool(epochId [32]byte, stakee ethcommon.Address, amount *big.Int) (*types.Transaction, error)
	InitializeRewardPool(epochId [32]byte) (*types.Transaction, error)
	ClaimRewards(epochId [][32]byte, stakee ethcommon.Address) (*types.Transaction, error)

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
	addresses Addresses

	opts *bind.TransactOpts

	// Embedded contracts
	*contracts.TicketingParametersSession
	*contracts.SyloTicketingSession
	*contracts.SyloTokenSession
	*contracts.DirectorySession
	*contracts.ListingsSession
	*contracts.PriceManagerSession
	*contracts.PriceVotingSession
	*contracts.StakingManagerSession
	*contracts.EpochsManagerSession
	*contracts.RewardsManagerSession

	backend Backend
}

func NewClient(
	ctx context.Context,
	addresses Addresses,
	eth *ethclient.Client,
	opts *bind.TransactOpts,
) (Client, error) {

	chainID, err := eth.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get chain id: %w", err)
	}

	signer := types.NewEIP155Signer(chainID)

	backend, err := NewBackend(eth, signer)
	if err != nil {
		return nil, fmt.Errorf("could not get backend: %w", err)
	}

	return NewClientWithBackend(
		addresses,
		backend,
		opts,
	)
}

func NewClientWithBackend(
	addresses Addresses,
	backend Backend,
	opts *bind.TransactOpts,
) (Client, error) {

	syloToken, err := contracts.NewSyloToken(addresses.Token, backend)
	if err != nil {
		return nil, err
	}

	TokenSession := &contracts.SyloTokenSession{
		Contract:     syloToken,
		TransactOpts: *opts,
	}

	ticketingParamters, err := contracts.NewTicketingParameters(addresses.TicketingParameters, backend)
	if err != nil {
		return nil, err
	}

	TicketingParametersSession := &contracts.TicketingParametersSession{
		Contract:     ticketingParamters,
		TransactOpts: *opts,
	}

	syloTicketing, err := contracts.NewSyloTicketing(addresses.Ticketing, backend)
	if err != nil {
		return nil, err
	}

	TicketingSession := &contracts.SyloTicketingSession{
		Contract:     syloTicketing,
		TransactOpts: *opts,
	}

	directory, err := contracts.NewDirectory(addresses.Directory, backend)
	if err != nil {
		return nil, err
	}

	DirectorySession := &contracts.DirectorySession{
		Contract:     directory,
		TransactOpts: *opts,
	}

	listings, err := contracts.NewListings(addresses.Listings, backend)
	if err != nil {
		return nil, err
	}

	ListingsSession := &contracts.ListingsSession{
		Contract:     listings,
		TransactOpts: *opts,
	}

	epochsManager, err := contracts.NewEpochsManager(addresses.EpochsManager, backend)
	if err != nil {
		return nil, err
	}

	EpochsManagerSession := &contracts.EpochsManagerSession{
		Contract:     epochsManager,
		TransactOpts: *opts,
	}

	priceManager, err := contracts.NewPriceManager(addresses.PriceManager, backend)
	if err != nil {
		return nil, err
	}

	PriceManagerSession := &contracts.PriceManagerSession{
		Contract:     priceManager,
		TransactOpts: *opts,
	}

	priceVoting, err := contracts.NewPriceVoting(addresses.PriceVoting, backend)
	if err != nil {
		return nil, err
	}

	PriceVotingSession := &contracts.PriceVotingSession{
		Contract:     priceVoting,
		TransactOpts: *opts,
	}

	stakingManager, err := contracts.NewStakingManager(addresses.StakingManager, backend)
	if err != nil {
		return nil, err
	}

	StakingManagerSession := &contracts.StakingManagerSession{
		Contract:     stakingManager,
		TransactOpts: *opts,
	}

	rewardsManager, err := contracts.NewRewardsManager(addresses.RewardsManager, backend)
	if err != nil {
		return nil, err
	}

	RewardsManagerSession := &contracts.RewardsManagerSession{
		Contract:     rewardsManager,
		TransactOpts: *opts,
	}

	return &client{
		addresses:                  addresses,
		backend:                    backend,
		TicketingParametersSession: TicketingParametersSession,
		SyloTicketingSession:       TicketingSession,
		SyloTokenSession:           TokenSession,
		PriceManagerSession:        PriceManagerSession,
		EpochsManagerSession:       EpochsManagerSession,
		PriceVotingSession:         PriceVotingSession,
		StakingManagerSession:      StakingManagerSession,
		DirectorySession:           DirectorySession,
		ListingsSession:            ListingsSession,
		RewardsManagerSession:      RewardsManagerSession,
		opts:                       opts,
	}, nil
}

func (c *client) Address() ethcommon.Address {
	return c.opts.From
}

func (c *client) ApproveTicketing(amount *big.Int) (*types.Transaction, error) {
	return c.Approve(c.addresses.Ticketing, amount)
}

func (c *client) ApproveStakingManager(amount *big.Int) (*types.Transaction, error) {
	return c.Approve(c.addresses.StakingManager, amount)
}

func (c *client) GetAmountStaked(stakee ethcommon.Address) (*big.Int, error) {
	return c.GetStakeeTotalStake(stakee)
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
		return receipt.BlockNumber, fmt.Errorf("Tx %v failed with status %v, %v", tx.Hash().Hex(), receipt.Status, receipt.PostState)
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
