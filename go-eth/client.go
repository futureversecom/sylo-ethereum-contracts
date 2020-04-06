package eth

//go:generate abigen --abi ../abi/SyloTicketing.abi --pkg contracts --type SyloTicketing --out contracts/syloTicketing.go --bin ../bin/SyloTicketing.bin
//go:generate abigen --abi ../abi/SyloToken.abi --pkg contracts --type SyloToken --out contracts/syloToken.go --bin ../bin/SyloToken.bin
//go:generate abigen --abi ../abi/Directory.abi --pkg contracts --type Directory --out contracts/directory.go --bin ../bin/Directory.bin
//go:generate abigen --abi ../abi/Listings.abi --pkg contracts --type Listings --out contracts/listings.go --bin ../bin/Listings.bin

import (
	"context"
	"math/big"
	"time"

	"github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/pkg/errors"
)

type Client interface {
	Address() ethcommon.Address

	// Ticketing methods
	Deposits(account ethcommon.Address) (struct {
		Escrow   *big.Int
		Penalty  *big.Int
		UnlockAt *big.Int
	}, error)
	//UnlockDuration() (*big.Int, error)  // Conflicting names
	DepositEscrow(amount *big.Int) (*types.Transaction, error)
	DepositEscrowFor(amount *big.Int, account ethcommon.Address) (*types.Transaction, error)
	DepositPenalty(amount *big.Int) (*types.Transaction, error)
	DepositPenaltyFor(amount *big.Int, account ethcommon.Address) (*types.Transaction, error)
	UnlockDeposits() (*types.Transaction, error)
	LockDeposits() (*types.Transaction, error)
	Redeem(ticket contracts.SyloTicketingTicket, receiverRand *big.Int, sig []byte) (*types.Transaction, error)
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

	// Directory methods
	Stakes(account ethcommon.Address) (struct {
		Amount   *big.Int
		UnlockAt *big.Int
	}, error)
	AddStake(amount *big.Int) (*types.Transaction, error)
	AddStakeFor(amount *big.Int, staker ethcommon.Address) (*types.Transaction, error)
	UnlockStake() (*types.Transaction, error)
	LockStake() (*types.Transaction, error)
	Unstake() (*types.Transaction, error)
	UnstakeTo(account ethcommon.Address) (*types.Transaction, error)
	Scan(rand *big.Int) (ethcommon.Address, error)

	// Listings methods
	SetListing(contracts.ListingsListing) (*types.Transaction, error)
	GetListing(account ethcommon.Address) (contracts.ListingsListing, error)

	// Alias for Approve but uses the ticketingAddress or directoryAddress as the spender
	ApproveTicketing(amount *big.Int) (*types.Transaction, error)
	ApproveDirectory(amount *big.Int) (*types.Transaction, error)

	//Utils
	LatestBlock() (*big.Int, error)
	CheckTx(tx *types.Transaction, timeout time.Duration) (*big.Int, error)
}

type client struct {
	ticketingAddress ethcommon.Address
	tokenAddress     ethcommon.Address
	directoryAddress ethcommon.Address
	listingsAddress  ethcommon.Address

	opts *bind.TransactOpts

	// Embedded contracts
	*contracts.SyloTicketingSession
	*contracts.SyloTokenSession
	*contracts.DirectorySession
	*contracts.ListingsSession

	backend Backend
}

func NewClient(
	ctx context.Context,
	ticketingAddress ethcommon.Address,
	tokenAddress ethcommon.Address,
	directoryAddress ethcommon.Address,
	listingsAddress ethcommon.Address,
	eth *ethclient.Client,
	opts *bind.TransactOpts,
) (Client, error) {

	chainID, err := eth.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	signer := types.NewEIP155Signer(chainID)

	backend, err := NewBackend(eth, signer)
	if err != nil {
		return nil, err
	}

	return NewClientWithBackend(
		ticketingAddress,
		tokenAddress,
		directoryAddress,
		listingsAddress,
		backend,
		opts,
	)
}

func NewClientWithBackend(
	tokenAddress ethcommon.Address,
	ticketingAddress ethcommon.Address,
	directoryAddress ethcommon.Address,
	listingsAddress ethcommon.Address,
	backend Backend,
	opts *bind.TransactOpts,
) (Client, error) {

	syloToken, err := contracts.NewSyloToken(tokenAddress, backend)
	if err != nil {
		return nil, err
	}

	TokenSession := &contracts.SyloTokenSession{
		Contract:     syloToken,
		TransactOpts: *opts,
	}

	syloTicketing, err := contracts.NewSyloTicketing(ticketingAddress, backend)
	if err != nil {
		return nil, err
	}

	TicketingSession := &contracts.SyloTicketingSession{
		Contract:     syloTicketing,
		TransactOpts: *opts,
	}

	directory, err := contracts.NewDirectory(directoryAddress, backend)
	if err != nil {
		return nil, err
	}

	DirectorySession := &contracts.DirectorySession{
		Contract:     directory,
		TransactOpts: *opts,
	}

	listings, err := contracts.NewListings(listingsAddress, backend)
	if err != nil {
		return nil, err
	}

	ListingsSession := &contracts.ListingsSession{
		Contract:     listings,
		TransactOpts: *opts,
	}

	return &client{
		ticketingAddress:     ticketingAddress,
		tokenAddress:         tokenAddress,
		directoryAddress:     directoryAddress,
		listingsAddress:      listingsAddress,
		backend:              backend,
		SyloTicketingSession: TicketingSession,
		SyloTokenSession:     TokenSession,
		DirectorySession:     DirectorySession,
		ListingsSession:      ListingsSession,
		opts:                 opts,
	}, nil
}

func (c *client) Address() ethcommon.Address {
	return c.opts.From
}

func (c *client) ApproveTicketing(amount *big.Int) (*types.Transaction, error) {
	return c.Approve(c.ticketingAddress, amount)
}

func (c *client) ApproveDirectory(amount *big.Int) (*types.Transaction, error) {
	return c.Approve(c.directoryAddress, amount)
}

func (c *client) LatestBlock() (*big.Int, error) {

	header, err := c.backend.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return header.Number, nil
}

func (c *client) CheckTx(tx *types.Transaction, timeout time.Duration) (*big.Int, error) {

	ctxT, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	receipt, err := bind.WaitMined(ctxT, c.backend, tx)
	if err != nil {
		return nil, err
	}

	if receipt.Status == uint64(0) {
		return receipt.BlockNumber, errors.Errorf("Tx %v failed with status %v, %v", tx.Hash().Hex(), receipt.Status, receipt.PostState)
	}

	return receipt.BlockNumber, nil
}
