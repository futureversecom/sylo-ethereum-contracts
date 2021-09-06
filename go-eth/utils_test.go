package eth

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"testing"

	"github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	OneEth           = big.NewInt(1000000000000000000)
	FaucetEthBalance = new(big.Int).Mul(OneEth, big.NewInt(10000))
	Uint128max       = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 128), big.NewInt(1)) // 2^128-1

	chainID        = big.NewInt(1337)
	unlockDuration = big.NewInt(10)
	escrowAmount   = big.NewInt(100000)
	penaltyAmount  = big.NewInt(1000)
)

func StartupEthereum(t *testing.T, ctx context.Context) (SimBackend, Addresses, FaucetF, Client) {
	ownerPK, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("could not create ecdsa key: %v", err)
	}
	ownerTransactor, err := bind.NewKeyedTransactorWithChainID(ownerPK, chainID)
	if err != nil {
		t.Fatalf("could not create transaction signer: %v", err)
	}
	ownerTransactor.Context = ctx

	backend := CreateBackend(t, ctx, ownerTransactor.From)
	addresses := DeployContracts(t, ctx, ownerTransactor, backend)

	ownerClient, err := NewClientWithBackend(addresses, backend, ownerTransactor)
	if err != nil {
		t.Fatalf("could not create client: %v", err)
	}

	// create a faucet
	faucet := MakeFaucet(t, ctx, backend, ownerClient, ownerPK)
	return backend, addresses, faucet, ownerClient
}

func CreateBackend(t *testing.T, ctx context.Context, owner common.Address) SimBackend {
	gasLimit := uint64(100000000000000)
	genesis := make(core.GenesisAlloc)
	genesis[owner] = core.GenesisAccount{Balance: FaucetEthBalance}
	return NewSimBackend(backends.NewSimulatedBackend(genesis, gasLimit))
}

func CreateRandomClient(t *testing.T, ctx context.Context, backend SimBackend, addresses Addresses) (*client, *ecdsa.PrivateKey) {
	pk, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		t.Fatalf("could not create ecdsa key: %v", err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		t.Fatalf("could not create transaction signer: %v", err)
	}
	opts.Context = ctx

	i, err := NewClientWithBackend(addresses, backend, opts)
	if err != nil {
		t.Fatalf("could not create client: %v", err)
	}

	c, ok := i.(*client)
	if !ok {
		t.Fatal("could not cast type to *client")
	}

	return c, pk
}

type FaucetF func(t *testing.T, recipient ethcommon.Address, ethAmt *big.Int, syloAmt *big.Int)

func MakeFaucet(t *testing.T, ctx context.Context, backend SimBackend, client Client, pk *ecdsa.PrivateKey) FaucetF {
	return func(t *testing.T, recipient ethcommon.Address, ethAmt *big.Int, syloAmt *big.Int) {
		if ethAmt.Cmp(big.NewInt(0)) == 1 {
			err := backend.FaucetEth(ctx, client.Address(), recipient, pk, ethAmt)
			if err != nil {
				t.Fatalf("could not faucet eth: %v", err)
			}
		}
		if syloAmt.Cmp(big.NewInt(0)) == 1 {
			tx, err := client.Transfer(recipient, syloAmt)
			if err != nil {
				t.Fatalf("could not faucet sylo: %v", err)
			}
			backend.Commit()
			_, err = client.CheckTx(ctx, tx)
			if err != nil {
				t.Fatalf("could not check sylo faucet transaction: %v", err)
			}
		}
	}
}

func TopUpDeposits(t *testing.T, ctx context.Context, backend SimBackend, client Client) {
	deposit, err := client.Deposits(client.Address())
	if err != nil {
		t.Fatalf("could not get deposits for top up: %v", err)
	}
	// make sure there is enough escrow
	if deposit.Escrow.Cmp(escrowAmount) == -1 {
		addAmount := new(big.Int).Add(escrowAmount, new(big.Int).Neg(deposit.Escrow))
		AddEscrow(t, ctx, backend, client, addAmount)
	}
	// make sure there is enough penalty
	if deposit.Penalty.Cmp(penaltyAmount) == -1 {
		addAmount := new(big.Int).Add(penaltyAmount, new(big.Int).Neg(deposit.Escrow))
		AddPenalty(t, ctx, backend, client, addAmount)
	}
}

func Stake(t *testing.T, ctx context.Context, backend SimBackend, client Client, stakeAmount *big.Int) {
	_, err := client.ApproveStakingManager(stakeAmount)
	if err != nil {
		t.Fatalf("could not approve spending: %v", err)
	}
	backend.Commit()

	tx, err := client.AddStake(stakeAmount, client.Address())
	if err != nil {
		t.Fatalf("could not add stake: %v", err)
	}
	backend.Commit()

	_, err = client.CheckTx(ctx, tx)
	if err != nil {
		t.Fatalf("could not check transaction: %v", err)
	}
}

func Vote(t *testing.T, ctx context.Context, backend SimBackend, client Client, vote *big.Int) {
	tx, err := client.Vote(vote)
	if err != nil {
		t.Fatalf("could not add stake: %v", err)
	}
	backend.Commit()

	_, err = client.CheckTx(ctx, tx)
	if err != nil {
		t.Fatalf("could not check transaction: %v", err)
	}
}

type SortedVote struct {
	Voter ethcommon.Address
	Price *big.Int
	Index *big.Int
}

func CalculatePrices(t *testing.T, ctx context.Context, backend SimBackend, client Client) *types.Transaction {
	voters, prices, err := client.GetVotes()
	if err != nil {
		t.Fatalf("could not retrieve votes: %v", err)
	}

	sortedVotes := []SortedVote{}

	for i := 0; i < len(voters); i++ {
		sortedVotes = append(
			sortedVotes,
			SortedVote{voters[i], prices[i], big.NewInt(int64(i))},
		)
	}
	sort.Slice(sortedVotes, func(i, j int) bool {
		return sortedVotes[i].Price.Cmp(sortedVotes[j].Price) == -1
	})

	sortedIndexes := []*big.Int{}
	for i := 0; i < len(voters); i++ {
		sortedIndexes = append(sortedIndexes, sortedVotes[i].Index)
	}

	tx, err := client.CalculatePrices(sortedIndexes)
	if err != nil {
		t.Fatalf("could not add stake: %v", err)
	}
	backend.Commit()

	_, err = client.CheckTx(ctx, tx)
	if err != nil {
		t.Fatalf("could not check transaction: %v", err)
	}

	return tx
}

func ConstructDirectory(t *testing.T, ctx context.Context, backend SimBackend, client Client) {
	tx, err := client.ConstructDirectory()
	if err != nil {
		t.Fatalf("could not add stake: %v", err)
	}
	backend.Commit()

	_, err = client.CheckTx(ctx, tx)
	if err != nil {
		t.Fatalf("could not check transaction: %v", err)
	}
}

func InitializeEpoch(t *testing.T, ctx context.Context, backend SimBackend, client Client) {
	tx, err := client.InitializeEpoch()
	if err != nil {
		t.Fatalf("could not initializ epoch: %v", err)
	}
	backend.Commit()

	_, err = client.CheckTx(ctx, tx)
	if err != nil {
		t.Fatalf("could not check transaction: %v", err)
	}
}

func DelegateStake(t *testing.T, ctx context.Context, backend SimBackend, client Client, stakee ethcommon.Address, stakeAmount *big.Int) {
	_, err := client.ApproveStakingManager(stakeAmount)
	if err != nil {
		t.Fatalf("could not approve spending: %v", err)
	}
	backend.Commit()

	tx, err := client.AddStake(stakeAmount, stakee)
	if err != nil {
		t.Fatalf("could not add stake: %v", err)
	}
	backend.Commit()

	_, err = client.CheckTx(ctx, tx)
	if err != nil {
		t.Fatalf("could not check transaction: %v", err)
	}
}

func List(t *testing.T, ctx context.Context, backend SimBackend, client Client, multiAddr string, minimumStakeAmount *big.Int) {
	tx, err := client.SetListing(multiAddr, minimumStakeAmount)
	if err != nil {
		t.Fatalf("could not add stake: %v", err)
	}
	backend.Commit()

	_, err = client.CheckTx(ctx, tx)
	if err != nil {
		t.Fatalf("could not check transaction: %v", err)
	}
}

func WaitForUnlockAt(t *testing.T, ctx context.Context, backend SimBackend, client Client) bool {
	unlocking, err := client.GetUnlockingStake(client.Address(), client.Address())
	if err != nil {
		t.Fatalf("could not check unlocking status: %v", err)
	}
	for {
		select {
		case <-ctx.Done():
			break
		default:
		}
		n, err := client.LatestBlock()
		if err != nil {
			t.Fatalf("could not get latest block: %v", err)
		}
		if BigIntsEqual(unlocking.UnlockAt, big.NewInt(0)) {
			// nothing to wait for
			return false
		}
		if n.Cmp(unlocking.UnlockAt) != -1 {
			// unlock block reached
			return true
		}
		backend.Commit()
	}
}

// UnstakeAll will unlock all stake and wait for it to be unlocked and
// withdrawn.
func UnstakeAll(t *testing.T, ctx context.Context, backend SimBackend, client Client) {
	stakeAmount, err := client.GetAmountStaked(client.Address())
	if err != nil {
		t.Fatalf("could not get staked amount: %v", err)
	}

	tx, err := client.UnlockStake(stakeAmount, client.Address())
	if err == nil {
		backend.Commit()
		_, err = client.CheckTx(ctx, tx)
		if err != nil {
			t.Fatalf("could not check transaction: %v", err)
		}
		// wait for unlocking
		if WaitForUnlockAt(t, ctx, backend, client) {
			// withdraw the unstaked amount
			tx, err := client.WithdrawStake(client.Address())
			if err != nil {
				t.Fatalf("could not unstake: %v", err)
			}
			backend.Commit()
			_, err = client.CheckTx(ctx, tx)
			if err != nil {
				t.Fatalf("could not check transaction: %v", err)
			}
		}
	} else if strings.HasSuffix(err.Error(), "Nothing to unstake") {
		// withdraw any unstaked amount
		tx, err := client.WithdrawStake(client.Address())
		if err == nil {
			backend.Commit()
			_, err = client.CheckTx(ctx, tx)
			if err != nil {
				t.Fatalf("could not check transaction: %v", err)
			}
		} else if strings.HasSuffix(err.Error(), "No amount to withdraw") {
			// nothing to do
		} else {
			t.Fatalf("could not unstake: %v", err)
		}
	} else {
		// error unlocking
		t.Fatalf("could not unlock stake: %v", err)
	}

	// check that all the stake is gone
	stakeAmount, err = client.GetAmountStaked(client.Address())
	if err != nil {
		t.Fatalf("could not get staked amount: %v", err)
	}
	if !BigIntsEqual(stakeAmount, big.NewInt(0)) {
		t.Fatalf("all stake should be removed: got %v", stakeAmount)
	}

	unlocking, err := client.GetUnlockingStake(client.Address(), client.Address())
	if err != nil {
		t.Fatalf("could not check unlocking status: %v", err)
	}
	if !BigIntsEqual(unlocking.Amount, big.NewInt(0)) {
		t.Fatalf("unlocking amount should be zero: got %v", unlocking.Amount)
	}
	if !BigIntsEqual(unlocking.UnlockAt, big.NewInt(0)) {
		t.Fatalf("unlocking at should be zero")
	}
}

func DeployContracts(t *testing.T, ctx context.Context, transactor *bind.TransactOpts, backend SimBackend) Addresses {
	var addresses Addresses
	var err error
	var tx *types.Transaction

	// Deploying contracts can apparently panic if the transaction fails, so
	// we need to check for that.
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panic during deployment of contracts: %v", r)
		}
	}()

	// deploy Sylo token
	addresses.Token, tx, _, err = contracts.DeploySyloToken(transactor, backend)
	if err != nil {
		t.Fatalf("could not deploy sylo token: %v", err)
	}
	backend.Commit()
	_, err = backend.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		t.Fatalf("could not get transaction receipt: %v", err)
	}

	// deploy staking manager
	var stakingManager *contracts.StakingManager
	addresses.StakingManager, _, stakingManager, err = contracts.DeployStakingManager(transactor, backend)
	if err != nil {
		t.Fatalf("could not deploy staking manager")
	}

	_, err = stakingManager.Initialize(transactor, addresses.Token, unlockDuration)
	if err != nil {
		t.Fatalf("could not initialize directory contract: %v", err)
	}
	backend.Commit()

	// deploy price voting
	var priceVoting *contracts.PriceVoting
	addresses.PriceVoting, _, priceVoting, err = contracts.DeployPriceVoting(transactor, backend)
	if err != nil {
		t.Fatalf("could not deploy price voting")
	}

	_, err = priceVoting.Initialize(transactor, addresses.StakingManager)
	if err != nil {
		t.Fatalf("could not initialize price voting contract: %v", err)
	}
	backend.Commit()

	// deploy price maanger
	var priceManager *contracts.PriceManager
	addresses.PriceManager, _, priceManager, err = contracts.DeployPriceManager(transactor, backend)
	if err != nil {
		t.Fatalf("could not deploy price manager")
	}

	_, err = priceManager.Initialize(transactor, addresses.StakingManager, addresses.PriceVoting)
	if err != nil {
		t.Fatalf("could not initialize price manager contract: %v", err)
	}
	backend.Commit()

	// deploy directory
	var directory *contracts.Directory
	addresses.Directory, tx, directory, err = contracts.DeployDirectory(transactor, backend)
	if err != nil {
		t.Fatalf("could not deploy directory: %v", err)
	}

	_, err = directory.Initialize(transactor, addresses.StakingManager)
	if err != nil {
		t.Fatalf("could not initialize directory contract: %v", err)
	}

	backend.Commit()
	_, err = backend.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		t.Fatalf("could not get transaction receipt: %v", err)
	}

	// deploy listing
	var listings *contracts.Listings
	addresses.Listings, tx, listings, err = contracts.DeployListings(transactor, backend)
	if err != nil {
		t.Fatalf("could not deploy listing: %v", err)
	}

	_, err = listings.Initialize(transactor, 50)
	if err != nil {
		t.Fatalf("could not get listings receipt: %v", err)
	}

	backend.Commit()
	_, err = backend.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		t.Fatalf("could not get transaction receipt: %v", err)
	}

	// deploy ticketing parameters
	var ticketingParameters *contracts.TicketingParameters
	addresses.TicketingParameters, tx, ticketingParameters, err = contracts.DeployTicketingParameters(transactor, backend)
	if err != nil {
		t.Fatalf("could not deploy ticketingParameters: %v", err)
	}

	_, err = ticketingParameters.Initialize(transactor, big.NewInt(1), Uint128max, big.NewInt(10000), uint16(8000), big.NewInt(100))
	if err != nil {
		t.Fatalf("could not initialize ticket contract: %v", err)
	}

	backend.Commit()
	_, err = backend.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		t.Fatalf("could not get transaction receipt: %v", err)
	}

	// deploy epochs manager
	var epochsManager *contracts.EpochsManager
	addresses.EpochsManager, tx, epochsManager, err = contracts.DeployEpochsManager(transactor, backend)
	if err != nil {
		t.Fatalf("could not deploy epochsManager: %v", err)
	}

	_, err = epochsManager.Initialize(transactor, addresses.Directory, addresses.Listings, addresses.TicketingParameters, big.NewInt(1))
	if err != nil {
		t.Fatalf("could not initialize ticket contract: %v", err)
	}

	backend.Commit()
	_, err = backend.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		t.Fatalf("could not get transaction receipt: %v", err)
	}

	// deploy ticketing
	var ticketing *contracts.SyloTicketing
	addresses.Ticketing, tx, ticketing, err = contracts.DeploySyloTicketing(transactor, backend)
	if err != nil {
		t.Fatalf("could not deploy ticketing: %v", err)
	}

	_, err = ticketing.Initialize(transactor, addresses.Token, addresses.Listings, addresses.Directory, addresses.EpochsManager, unlockDuration)
	if err != nil {
		t.Fatalf("could not initialize ticket contract: %v", err)
	}

	backend.Commit()
	_, err = backend.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		t.Fatalf("could not get transaction receipt: %v", err)
	}

	return addresses
}

func AddEscrow(t *testing.T, ctx context.Context, backend SimBackend, client Client, escrowAmount *big.Int) {
	err := addDeposit(ctx, backend, client, escrowAmount, client.DepositEscrow)
	if err != nil {
		t.Fatalf("could not add escrow amount: %v", err)
	}
}

func AddPenalty(t *testing.T, ctx context.Context, backend SimBackend, client Client, penaltyAmount *big.Int) {
	err := addDeposit(ctx, backend, client, penaltyAmount, client.DepositPenalty)
	if err != nil {
		t.Fatalf("could not add penalty amount: %v", err)
	}
}

type depositF func(amount *big.Int, account ethcommon.Address) (*types.Transaction, error)

func addDeposit(ctx context.Context, backend SimBackend, client Client, amount *big.Int, f depositF) error {
	tx, err := client.ApproveTicketing(amount)
	if err != nil {
		return fmt.Errorf("could not approve ticketing amount: %w", err)
	}
	backend.Commit()

	_, err = client.CheckTx(ctx, tx)
	if err != nil {
		return fmt.Errorf("could not check transaction: %w", err)
	}

	tx, err = f(amount, client.Address())
	if err != nil {
		return fmt.Errorf("could not deposit: %w", err)
	}
	backend.Commit()

	_, err = client.CheckTx(ctx, tx)
	if err != nil {
		return fmt.Errorf("could not confirm deposit transaction: %w", err)
	}

	return nil
}

func BigIntsEqual(x *big.Int, y *big.Int) bool {
	return x.Cmp(y) == 0
}

func GetNode(t *testing.T, client Client) (struct {
	Amount *big.Int
	Stakee ethcommon.Address
}, []byte) {
	key, err := client.GetKey(client.Address(), client.Address())
	if err != nil {
		t.Fatalf("could not get key: %v", err)
	}
	node, err := client.Stakes(key)
	if err != nil {
		t.Fatalf("could not get node info: %v", err)
	}
	return node, key[:]
}
