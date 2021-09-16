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

type gasInitializeRewardPool struct {
	DelegatedStakers uint64 `json:"delegated_stakers"`
	GasUsed          uint64 `json:"gas_used"`
}

func TestGasIniitializeRewardPool(t *testing.T) {
	tcs := []*gasInitializeRewardPool{
		{DelegatedStakers: 1},
		{DelegatedStakers: 10},
		{DelegatedStakers: 50},
		{DelegatedStakers: 100},
	}

	for _, tc := range tcs {
		runGasInitializeRewardPool(t, tc)
	}

	b, err := json.Marshal(tcs)
	if err != nil {
		t.Fatalf("could not marshal test results: %v", err)
	}

	writeJsonOutput(t, b, "testdata/gasInitializeRewardPool.json")
}

type gasClaimReward struct {
	DelegatedStakers uint64 `json:"delegated_stakers"`
	GasUsed          uint64 `json:"gas_used"`
}

func TestGasClaimReward(t *testing.T) {
	tcs := []*gasClaimReward{
		{DelegatedStakers: 1},
		{DelegatedStakers: 10},
		{DelegatedStakers: 50},
		{DelegatedStakers: 100},
	}

	for _, tc := range tcs {
		runGasClaimReward(t, tc)
	}

	b, err := json.Marshal(tcs)
	if err != nil {
		t.Fatalf("could not marshal test results: %v", err)
	}

	writeJsonOutput(t, b, "testdata/gasClaimRewards.json")
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

func runGasInitializeRewardPool(t *testing.T, tc *gasInitializeRewardPool) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	backend, addresses, faucet, _ := StartupEthereum(t, ctx)

	node, _ := CreateRandomClient(t, ctx, backend, addresses)
	faucet(t, node.Address(), OneEth, big.NewInt(1000000))

	List(t, ctx, backend, node, "0", big.NewInt(1))

	// populate stakes
	for i := uint64(0); i < tc.DelegatedStakers; i++ {
		c, _ := CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, c.Address(), OneEth, big.NewInt(1000000))
		approveStakingManager(t, ctx, c, big.NewInt(100))
		addStakeGas(t, ctx, c, big.NewInt(100), node.Address())
		backend.Commit()
	}

	epochId := GetNextEpochId(t, ctx, backend, node)

	// initialize reward pool
	tx, err := node.InitializeRewardPool(epochId)
	if err != nil {
		t.Fatalf("failed to distribute rewards: %v", err)
	}
	tc.GasUsed = tx.Gas()
}

func runGasClaimReward(t *testing.T, tc *gasClaimReward) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	backend, addresses, faucet, owner := StartupEthereum(t, ctx)

	node, _ := CreateRandomClient(t, ctx, backend, addresses)
	faucet(t, node.Address(), OneEth, big.NewInt(1000000))

	List(t, ctx, backend, node, "0", big.NewInt(1))

	// populate stakes
	for i := uint64(0); i < tc.DelegatedStakers; i++ {
		c, _ := CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, c.Address(), OneEth, big.NewInt(1000000))
		approveStakingManager(t, ctx, c, big.NewInt(100))
		addStakeGas(t, ctx, c, big.NewInt(100), node.Address())
		backend.Commit()
	}

	epochId := GetNextEpochId(t, ctx, backend, node)

	// join directory and initialize reward pool
	JoinDirectory(epochId, t, ctx, backend, node)
	InitializeRewardPool(epochId, t, ctx, backend, node)

	_, err := owner.TransferDirectoryOwnership(addresses.EpochsManager)
	if err != nil {
		t.Fatalf("could not transfer directory ownership: %v", err)
	}
	InitializeEpoch(t, ctx, backend, owner)

	// set up sender
	sender, senderPK := CreateRandomClient(t, ctx, backend, addresses)
	faucet(t, sender.Address(), OneEth, big.NewInt(100000))
	AddEscrow(t, ctx, backend, sender, big.NewInt(1000))
	AddPenalty(t, ctx, backend, sender, big.NewInt(1000))

	backend.Commit()

	// increment reward pool
	for i := 0; i < 10; i++ {
		ticket, sig, senderRand, nodeRand := CreateWinningTicket(t, sender, senderPK, node.Address())
		Redeem(t, ctx, backend, node, ticket, senderRand, nodeRand, sig)
	}

	// end current epoch
	InitializeEpoch(t, ctx, backend, owner)
	// advance until tickets expire
	for i := 0; i < 20; i++ {
		backend.Commit()
	}

	// distribute rewards
	tx, err := node.ClaimRewards([][32]byte{epochId}, node.Address())
	if err != nil {
		t.Fatalf("failed to distribute rewards: %v", err)
	}
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
