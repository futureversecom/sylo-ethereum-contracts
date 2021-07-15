package eth

import (
	"bytes"
	"context"
	"encoding/json"
	"math/big"
	"os"
	"path"
	"testing"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

type gasTests struct {
	Description string           `json:"description"`
	Tests       []json.Marshaler `json:"test"`
}

type gasAddStakeTest struct {
	PriorStakers      uint64   `json:"prior_stakers"`
	PriorStakerAmount *big.Int `json:"prior_staker_amount"`
	StakeAmount       *big.Int `json:"stake_amounts"`
	GasUsed           uint64   `json:"gas_used"`
}

func (t *gasAddStakeTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(*t)
}

func TestGasAddStake(t *testing.T) {
	suite := []*gasTests{
		{
			Description: "vary stake amount",
			Tests: []json.Marshaler{
				&gasAddStakeTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(10)},
				&gasAddStakeTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(100)},
				&gasAddStakeTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000)},
			},
		},
		{
			Description: "vary stake tree size",
			Tests: []json.Marshaler{
				&gasAddStakeTest{PriorStakers: 0, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000)},
				&gasAddStakeTest{PriorStakers: 1, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000)},
				&gasAddStakeTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000)},
				&gasAddStakeTest{PriorStakers: 100, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000)},
			},
		},
		{
			Description: "vary prior stake amount",
			Tests: []json.Marshaler{
				&gasAddStakeTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(10), StakeAmount: big.NewInt(1000)},
				&gasAddStakeTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(100), StakeAmount: big.NewInt(1000)},
				&gasAddStakeTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000)},
			},
		},
	}
	runner(t, suite, runGasAddStake, "gasAddStake.json")
}

func runGasAddStake(t *testing.T, i interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tc, ok := i.(*gasAddStakeTest)
	if !ok {
		t.Fatal("received an invalid test case")
	}

	backend, addresses, faucet := StartupEthereum(t, ctx)

	// populate stake tree
	for i := uint64(0); i < tc.PriorStakers; i++ {
		c, _ := CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, c.Address(), OneEth, big.NewInt(1000000))
		approveDirectoryGas(t, ctx, c, tc.PriorStakerAmount)
		addStakeGas(t, ctx, c, tc.PriorStakerAmount, c.Address())
	}

	c, _ := CreateRandomClient(t, ctx, backend, addresses)
	faucet(t, c.Address(), OneEth, big.NewInt(1000000))
	approveDirectoryGas(t, ctx, c, tc.StakeAmount)
	tc.GasUsed = addStakeGas(t, ctx, c, tc.StakeAmount, c.Address())
}

type gasUnlockStakeTest struct {
	PriorStakers      uint64   `json:"prior_stakers"`
	PriorStakerAmount *big.Int `json:"prior_staker_amount"`
	StakeAmount       *big.Int `json:"stake_amount"`
	UnlockAmount      *big.Int `json:"unlock_amount"`
	GasUsed           uint64   `json:"gas_used"`
}

func (t *gasUnlockStakeTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(*t)
}

func TestGasUnlockStake(t *testing.T) {
	suite := []*gasTests{
		{
			Description: "vary unlock amount",
			Tests: []json.Marshaler{
				&gasUnlockStakeTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(100)},
				&gasUnlockStakeTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(500)},
				&gasUnlockStakeTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(999)},
				&gasUnlockStakeTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(1000)},
			},
		},
		{
			Description: "vary prior stakers",
			Tests: []json.Marshaler{
				&gasUnlockStakeTest{PriorStakers: 0, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(100)},
				&gasUnlockStakeTest{PriorStakers: 1, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(100)},
				&gasUnlockStakeTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(100)},
				&gasUnlockStakeTest{PriorStakers: 100, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(100)},
			},
		},
	}
	runner(t, suite, runGasUnlockStake, "gasUnlockStake.json")
}

func runGasUnlockStake(t *testing.T, i interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tc, ok := i.(*gasUnlockStakeTest)
	if !ok {
		t.Fatal("received an invalid test case")
	}

	backend, addresses, faucet := StartupEthereum(t, ctx)

	// populate stake tree
	for i := uint64(0); i < tc.PriorStakers; i++ {
		c, _ := CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, c.Address(), OneEth, big.NewInt(1000000))
		approveDirectoryGas(t, ctx, c, tc.PriorStakerAmount)
		addStakeGas(t, ctx, c, tc.PriorStakerAmount, c.Address())
	}

	c, _ := CreateRandomClient(t, ctx, backend, addresses)
	faucet(t, c.Address(), OneEth, big.NewInt(1000000))
	approveDirectoryGas(t, ctx, c, tc.StakeAmount)
	addStakeGas(t, ctx, c, tc.StakeAmount, c.Address())
	backend.Commit()
	tc.GasUsed = unlockStakeGas(t, ctx, c, tc.UnlockAmount, c.Address())
}

type gasCancelUnlockingTest struct {
	PriorStakers      uint64   `json:"prior_stakers"`
	PriorStakerAmount *big.Int `json:"prior_staker_amount"`
	StakeAmount       *big.Int `json:"stake_amount"`
	UnlockAmount      *big.Int `json:"unlock_amount"`
	CancelAmount      *big.Int `json:"cancel_amount"`
	GasUsed           uint64   `json:"gas_used"`
}

func (t *gasCancelUnlockingTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(*t)
}

func TestGasCancelUnlocking(t *testing.T) {
	suite := []*gasTests{
		{
			Description: "vary unlock amount",
			Tests: []json.Marshaler{
				&gasCancelUnlockingTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(100), CancelAmount: big.NewInt(100)},
				&gasCancelUnlockingTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(500), CancelAmount: big.NewInt(100)},
				&gasCancelUnlockingTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(999), CancelAmount: big.NewInt(100)},
				&gasCancelUnlockingTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(1000), CancelAmount: big.NewInt(100)},
			},
		},
		{
			Description: "vary prior stakers",
			Tests: []json.Marshaler{
				&gasCancelUnlockingTest{PriorStakers: 0, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(100), CancelAmount: big.NewInt(100)},
				&gasCancelUnlockingTest{PriorStakers: 1, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(100), CancelAmount: big.NewInt(100)},
				&gasCancelUnlockingTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(100), CancelAmount: big.NewInt(100)},
				&gasCancelUnlockingTest{PriorStakers: 100, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(100), CancelAmount: big.NewInt(100)},
			},
		},
		{
			Description: "vary cancel stakers",
			Tests: []json.Marshaler{
				&gasCancelUnlockingTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(100), CancelAmount: big.NewInt(1)},
				&gasCancelUnlockingTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(100), CancelAmount: big.NewInt(10)},
				&gasCancelUnlockingTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(100), CancelAmount: big.NewInt(100)},
			},
		},
	}
	runner(t, suite, runGasCancelUnlocking, "gasCancelUnlocking.json")
}

func runGasCancelUnlocking(t *testing.T, i interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tc, ok := i.(*gasCancelUnlockingTest)
	if !ok {
		t.Fatal("received an invalid test case")
	}

	backend, addresses, faucet := StartupEthereum(t, ctx)

	// populate stake tree
	for i := uint64(0); i < tc.PriorStakers; i++ {
		c, _ := CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, c.Address(), OneEth, big.NewInt(1000000))
		approveDirectoryGas(t, ctx, c, tc.PriorStakerAmount)
		addStakeGas(t, ctx, c, tc.PriorStakerAmount, c.Address())
	}

	c, _ := CreateRandomClient(t, ctx, backend, addresses)
	faucet(t, c.Address(), OneEth, big.NewInt(1000000))
	approveDirectoryGas(t, ctx, c, tc.StakeAmount)
	addStakeGas(t, ctx, c, tc.StakeAmount, c.Address())
	backend.Commit()
	unlockStakeGas(t, ctx, c, tc.UnlockAmount, c.Address())
	backend.Commit()
	tc.GasUsed = cancelUnlockingGas(t, ctx, c, tc.CancelAmount, c.Address())
}

type gasWithdrawStakeTest struct {
	PriorStakers      uint64   `json:"prior_stakers"`
	PriorStakerAmount *big.Int `json:"prior_staker_amount"`
	StakeAmount       *big.Int `json:"stake_amount"`
	UnlockAmount      *big.Int `json:"unlock_amount"`
	GasUsed           uint64   `json:"gas_used"`
}

func (t *gasWithdrawStakeTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(*t)
}

func TestGasWithdrawStake(t *testing.T) {
	suite := []*gasTests{
		{
			Description: "vary unlock amount",
			Tests: []json.Marshaler{
				&gasWithdrawStakeTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(100)},
				&gasWithdrawStakeTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(500)},
				&gasWithdrawStakeTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(999)},
				&gasWithdrawStakeTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(1000)},
			},
		},
		{
			Description: "vary prior stakers",
			Tests: []json.Marshaler{
				&gasWithdrawStakeTest{PriorStakers: 0, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(100)},
				&gasWithdrawStakeTest{PriorStakers: 1, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(100)},
				&gasWithdrawStakeTest{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(100)},
				&gasWithdrawStakeTest{PriorStakers: 100, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000), UnlockAmount: big.NewInt(100)},
			},
		},
	}
	runner(t, suite, runGasWithdrawStake, "gasWithdrawStake.json")
}

func runGasWithdrawStake(t *testing.T, i interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tc, ok := i.(*gasWithdrawStakeTest)
	if !ok {
		t.Fatal("received an invalid test case")
	}

	backend, addresses, faucet := StartupEthereum(t, ctx)

	// populate stake tree
	for i := uint64(0); i < tc.PriorStakers; i++ {
		c, _ := CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, c.Address(), OneEth, big.NewInt(1000000))
		approveDirectoryGas(t, ctx, c, tc.PriorStakerAmount)
		addStakeGas(t, ctx, c, tc.PriorStakerAmount, c.Address())
	}

	c, _ := CreateRandomClient(t, ctx, backend, addresses)
	faucet(t, c.Address(), OneEth, big.NewInt(1000000))
	approveDirectoryGas(t, ctx, c, tc.StakeAmount)
	addStakeGas(t, ctx, c, tc.StakeAmount, c.Address())
	backend.Commit()
	unlockStakeGas(t, ctx, c, tc.UnlockAmount, c.Address())
	for i := 0; i < 11; i++ {
		backend.Commit()
	}
	tc.GasUsed = withdrawStakeGas(t, ctx, c, c.Address())
}

func addStakeGas(t *testing.T, ctx context.Context, c *client, amount *big.Int, stakee ethcommon.Address) uint64 {
	tx, err := c.AddStake(amount, stakee)
	if err != nil {
		t.Fatalf("could not add stake: %v", err)
	}
	return tx.Gas()
}

func unlockStakeGas(t *testing.T, ctx context.Context, c *client, amount *big.Int, stakee ethcommon.Address) uint64 {
	tx, err := c.UnlockStake(amount, stakee)
	if err != nil {
		t.Fatalf("could not unlock stake: %v", err)
	}
	return tx.Gas()
}

func cancelUnlockingGas(t *testing.T, ctx context.Context, c *client, amount *big.Int, stakee ethcommon.Address) uint64 {
	tx, err := c.CancelUnlocking(amount, stakee)
	if err != nil {
		t.Fatalf("could not cancel unlocking: %v", err)
	}
	return tx.Gas()
}

func withdrawStakeGas(t *testing.T, ctx context.Context, c *client, stakee ethcommon.Address) uint64 {
	tx, err := c.WithdrawStake(stakee)
	if err != nil {
		t.Fatalf("could not withdraw stake: %v", err)
	}
	return tx.Gas()
}

func approveDirectoryGas(t *testing.T, ctx context.Context, c *client, amount *big.Int) uint64 {
	tx, err := c.ApproveDirectory(amount)
	if err != nil {
		t.Fatalf("could not approve spending: %v", err)
	}
	return tx.Gas()
}

func runner(t *testing.T, suite []*gasTests, f func(t *testing.T, tc interface{}), outFile string) {
	for _, s := range suite {
		for _, tc := range s.Tests {
			f(t, tc)
		}
	}

	b, err := json.Marshal(suite)
	if err != nil {
		t.Fatalf("could not marshal test results: %v", err)
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		t.Fatalf("could not indent test results: %v", err)
	}
	p, err := os.Create(path.Join("testdata", path.Clean(outFile)))
	if err != nil {
		t.Fatalf("could not create test result file: %v", err)
	}
	defer p.Close()
	_, err = out.WriteTo(p)
	if err != nil {
		t.Fatalf("could not write test results to file: %v", err)
	}
}
