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

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		t.Fatalf("could not indent test results: %v", err)
	}
	f, err := os.Create(path.Clean("testdata/gasAddStakeOut.json"))
	if err != nil {
		t.Fatalf("could not create test result file: %v", err)
	}
	defer f.Close()
	_, err = out.WriteTo(f)
	if err != nil {
		t.Fatalf("could not write test results to file: %v", err)
	}
}

func runGasAddStake(t *testing.T, tc *gasAddStakeTest) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	backend, addresses, faucet, _ := StartupEthereum(t, ctx)

	// populate stake tree
	for i := uint64(0); i < tc.PriorStakers; i++ {
		c, _ := CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, c.Address(), OneEth, big.NewInt(1000000))
		approveDirectoryGas(t, ctx, c, tc.PriorStakerAmount)
		addStakeGas(t, ctx, c, tc.PriorStakerAmount, c.Address())
	}

	c, _ := CreateRandomClient(t, ctx, backend, addresses)
	faucet(t, c.Address(), OneEth, big.NewInt(1000000))
	tc.GasUsed =
		approveDirectoryGas(t, ctx, c, tc.StakeAmount) +
			addStakeGas(t, ctx, c, tc.StakeAmount, c.Address())
}

func approveDirectoryGas(t *testing.T, ctx context.Context, c *client, amount *big.Int) uint64 {
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
