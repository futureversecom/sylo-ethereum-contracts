package eth

//go:generate abigen --abi ../abi/SyloTicketing.abi --pkg contracts --type SyloTicketing --out contracts/ticketing.go --bin ../bin/SyloTicketing.bin
//go:generate abigen --abi ../abi/SyloToken.abi --pkg contracts --type SyloToken --out contracts/token.go --bin ../bin/SyloToken.bin
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

type Unlocking struct {
	Amount   *big.Int
	UnlockAt *big.Int
}

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
	Stakes(key [32]byte) (struct {
		Amount      *big.Int
		LeftAmount  *big.Int
		RightAmount *big.Int
		Stakee      ethcommon.Address
		Parent      [32]byte
		Left        [32]byte
		Right       [32]byte
	}, error)
	GetKey(staker ethcommon.Address, stakee ethcommon.Address) ([32]byte, error)
	AddStake(amount *big.Int) (*types.Transaction, error)
	AddStakeFor(amount *big.Int, stakee ethcommon.Address) (*types.Transaction, error)
	UnlockStake(amount *big.Int) (*types.Transaction, error)
	UnlockStakeFor(amount *big.Int, stakee ethcommon.Address) (*types.Transaction, error)
	LockStake(amount *big.Int) (*types.Transaction, error)
	LockStakeFor(amount *big.Int, stakee ethcommon.Address) (*types.Transaction, error)
	Unstake() (*types.Transaction, error)
	UnstakeFor(account ethcommon.Address) (*types.Transaction, error)
	GetAmountStakedFor(stakee ethcommon.Address) (*big.Int, error)
	GetUnlockingStake(staker ethcommon.Address, stakee ethcommon.Address) (Unlocking, error)
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

type Addresses struct {
	Token     ethcommon.Address
	Ticketing ethcommon.Address
	Directory ethcommon.Address
	Listings  ethcommon.Address
}

type client struct {
	addresses Addresses

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
	addresses Addresses,
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

	return &client{
		addresses:            addresses,
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
	return c.Approve(c.addresses.Ticketing, amount)
}

func (c *client) ApproveDirectory(amount *big.Int) (*types.Transaction, error) {
	return c.Approve(c.addresses.Directory, amount)
}

func (c *client) GetAmountStakedFor(stakee ethcommon.Address) (*big.Int, error) {
	return c.Stakees(stakee)
}

func (c *client) GetUnlockingStake(staker ethcommon.Address, stakee ethcommon.Address) (Unlocking, error) {
	key, err := c.GetKey(staker, stakee)
	if err != nil {
		return Unlocking{}, err
	}
	return c.Unlockings(key)
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
