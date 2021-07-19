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

type gasAddStakeTest struct {
	PriorStakers      uint64   `json:"prior_stakers"`
	PriorStakerAmount *big.Int `json:"prior_staker_amount"`
	StakeAmount       *big.Int `json:"stake_amount"`
	GasUsed           uint64   `json:"gas_used"`
}

func TestGasAddStake(t *testing.T) {
	tcs := []*gasAddStakeTest{
		{PriorStakers: 0, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000)},
		{PriorStakers: 10, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000)},
		{PriorStakers: 100, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000)},
		{PriorStakers: 1000, PriorStakerAmount: big.NewInt(1000), StakeAmount: big.NewInt(1000)},
	}

	for _, tc := range tcs {
		runGasAddStake(t, tc)
	}

	b, err := json.Marshal(tcs)
	if err != nil {
		t.Fatalf("could not marshal test results: %v", err)
	}

	writeJsonOutput(t, b, "testdata/gasAddStakeOut.json")
}

type gasCalculatePricesTest struct {
	Nodes   uint64 `json:"voters"`
	GasUsed uint64 `json:"gas_used"`
}

func TestGasCalculatePrices(t *testing.T) {
	tcs := []*gasCalculatePricesTest{
		{Nodes: 1},
		{Nodes: 100},
		{Nodes: 200},
		{Nodes: 300},
	}

	for _, tc := range tcs {
		runGasCalculatePrices(t, tc)
	}

	b, err := json.Marshal(tcs)
	if err != nil {
		t.Fatalf("could not marshal test results: %v", err)
	}

	writeJsonOutput(t, b, "testdata/gasCalculatePricesOut.json")
}

func runGasAddStake(t *testing.T, tc *gasAddStakeTest) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	backend, addresses, faucet, _ := StartupEthereum(t, ctx)

	// populate stakes
	for i := uint64(0); i < tc.PriorStakers; i++ {
		c, _ := CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, c.Address(), OneEth, big.NewInt(1000000))
		approveStakingManager(t, ctx, c, tc.PriorStakerAmount)
		addStakeGas(t, ctx, c, tc.PriorStakerAmount, c.Address())
	}

	c, _ := CreateRandomClient(t, ctx, backend, addresses)
	faucet(t, c.Address(), OneEth, big.NewInt(1000000))
	tc.GasUsed =
		approveStakingManager(t, ctx, c, tc.StakeAmount) +
			addStakeGas(t, ctx, c, tc.StakeAmount, c.Address())
}

func approveStakingManager(t *testing.T, ctx context.Context, c *client, amount *big.Int) uint64 {
	tx, err := c.ApproveStakingManager(amount)
	if err != nil {
		t.Fatalf("could not approve spending: %v", err)
	}
	return tx.Gas()
}

func addStakeGas(t *testing.T, ctx context.Context, c *client, amount *big.Int, stakee ethcommon.Address) uint64 {
	tx, err := c.AddStake(amount, stakee)
	if err != nil {
		t.Fatalf("could not add stake: %v", err)
	}
	return tx.Gas()
}

func voteGas(t *testing.T, ctx context.Context, c *client, price *big.Int) uint64 {
	tx, err := c.Vote(price)
	if err != nil {
		t.Fatalf("could not vote: %v", err)
	}
	return tx.Gas()
}

func runGasCalculatePrices(t *testing.T, tc *gasCalculatePricesTest) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	backend, addresses, faucet, owner := StartupEthereum(t, ctx)

	// populate stakes and votes
	for i := uint64(0); i < tc.Nodes; i++ {
		c, _ := CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, c.Address(), OneEth, big.NewInt(1000000))
		approveStakingManager(t, ctx, c, big.NewInt(100))
		addStakeGas(t, ctx, c, big.NewInt(100), c.Address())
		voteGas(t, ctx, c, big.NewInt(1))
		backend.Commit()
	}

	tx := CalculatePrices(t, ctx, backend, owner)

	tc.GasUsed = tx.Gas()
}

func writeJsonOutput(t *testing.T, b []byte, filename string) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	if err != nil {
		t.Fatalf("could not indent test results: %v", err)
	}
	f, err := os.Create(path.Clean(filename))
	if err != nil {
		t.Fatalf("could not create test result file: %v", err)
	}
	defer f.Close()
	_, err = out.WriteTo(f)
	if err != nil {
		t.Fatalf("could not write test results to file: %v", err)
	}
}
