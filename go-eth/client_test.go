package eth_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"math/big"
	"strings"
	"testing"
	"time"

	sylopayments "github.com/dn3010/sylo-ethereum-contracts/go-eth"
	"github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	escrowAmount := big.NewInt(100000)
	penaltyAmount := big.NewInt(1000)
	unlockDuration := big.NewInt(10)

	backend, addresses, faucet, _ := sylopayments.StartupEthereum(t, ctx)

	t.Run("client can be created", func(t *testing.T) {
		sylopayments.CreateRandomClient(t, ctx, backend, addresses)
	})

	t.Run("can use faucet", func(t *testing.T) {
		aliceClient, _ := sylopayments.CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, aliceClient.Address(), sylopayments.OneEth, big.NewInt(1000000))
	})

	t.Run("can get latest block", func(t *testing.T) {
		aliceClient, _ := sylopayments.CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, aliceClient.Address(), sylopayments.OneEth, big.NewInt(0))
		blockNumberA, err := aliceClient.LatestBlock()
		if err != nil {
			t.Fatalf("could not get latest block: %v", err)
		}
		backend.Commit()
		blockNumberB, err := aliceClient.LatestBlock()
		if err != nil {
			t.Fatalf("could not get latest block: %v", err)
		}
		if !sylopayments.BigIntsEqual(blockNumberA.Add(blockNumberA, big.NewInt(1)), blockNumberB) {
			t.Fatalf("block number did not advance")
		}
	})

	t.Run("can deposit escrow", func(t *testing.T) {
		aliceClient, _ := sylopayments.CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, aliceClient.Address(), sylopayments.OneEth, big.NewInt(1000000))

		sylopayments.AddEscrow(t, ctx, backend, aliceClient, escrowAmount)

		deposit, err := aliceClient.Deposits(aliceClient.Address())
		if err != nil {
			t.Fatalf("could not get deposits: %v", err)
		}
		if !sylopayments.BigIntsEqual(deposit.Escrow, escrowAmount) {
			t.Fatalf("escrow deposit does not match: got %v: expected %v", deposit.Escrow, escrowAmount)
		}
	})

	t.Run("can deposit penalty", func(t *testing.T) {
		aliceClient, _ := sylopayments.CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, aliceClient.Address(), sylopayments.OneEth, big.NewInt(1000000))

		sylopayments.AddPenalty(t, ctx, backend, aliceClient, penaltyAmount)

		deposit, err := aliceClient.Deposits(aliceClient.Address())
		if err != nil {
			t.Fatalf("could not get deposits: %v", err)
		}
		if !sylopayments.BigIntsEqual(deposit.Penalty, penaltyAmount) {
			t.Fatalf("penalty deposit does not match: got %v: expected %v", deposit.Penalty, penaltyAmount)
		}
	})

	t.Run("can withdraw ticketing", func(t *testing.T) {
		aliceClient, _ := sylopayments.CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, aliceClient.Address(), sylopayments.OneEth, big.NewInt(1000000))
		sylopayments.TopUpDeposits(t, ctx, backend, aliceClient)

		tx, err := aliceClient.UnlockDeposits()
		if err != nil {
			t.Fatalf("could not unlock ticketing escrow: %v", err)
		}
		backend.Commit()

		_, err = aliceClient.CheckTx(ctx, tx)
		if err != nil {
			t.Fatalf("could not check transaction: %v", err)
		}

		_, err = aliceClient.Withdraw()
		if err == nil {
			t.Fatalf("expected error because unlock period isn't complete")
		}
		if !strings.HasSuffix(err.Error(), "Unlock period not complete") {
			t.Fatalf("could not withdraw: %v", err)
		}

		// advance enough blocks for the unlock period to end
		for i := uint64(0); i < unlockDuration.Uint64(); i++ {
			backend.Commit()
		}

		tx, err = aliceClient.Withdraw()
		if err != nil {
			t.Fatalf("could not withdraw: %v", err)
		}
		backend.Commit()
		_, err = aliceClient.CheckTx(ctx, tx)
		if err != nil {
			t.Fatalf("could not confirm transaction: %v", err)
		}

		deposit, err := aliceClient.Deposits(aliceClient.Address())
		if err != nil {
			t.Fatalf("could not get deposits: %v", err)
		}
		if !sylopayments.BigIntsEqual(deposit.Escrow, big.NewInt(0)) {
			t.Fatalf("escrow should be withdrawn")
		}
		if !sylopayments.BigIntsEqual(deposit.Penalty, big.NewInt(0)) {
			t.Fatalf("penalty should be withdrawn")
		}
	})

	t.Run("can redeem ticket", func(t *testing.T) {
		aliceClient, alicePK := sylopayments.CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, aliceClient.Address(), sylopayments.OneEth, big.NewInt(1000000))
		sylopayments.TopUpDeposits(t, ctx, backend, aliceClient)

		bobClient, _ := sylopayments.CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, bobClient.Address(), sylopayments.OneEth, big.NewInt(0))

		sylopayments.List(t, ctx, backend, bobClient, "0.0.0.0/0", big.NewInt(1))

		sylopayments.DelegateStake(t, ctx, backend, aliceClient, bobClient.Address(), big.NewInt(600))

		bobRand := big.NewInt(1)
		var bobRandHash [32]byte
		copy(bobRandHash[:], crypto.Keccak256(bobRand.FillBytes(bobRandHash[:])))

		ticket := contracts.SyloTicketingTicket{
			Sender:           aliceClient.Address(),
			Receiver:         bobClient.Address(),
			ReceiverRandHash: bobRandHash,
			FaceValue:        big.NewInt(1),
			WinProb:          sylopayments.Uint256max, // always win
			ExpirationBlock:  big.NewInt(0),
			SenderNonce:      1,
		}

		ticketHash, err := aliceClient.GetTicketHash(ticket)
		if err != nil {
			t.Fatalf("could not get ticket hash: %v", err)
		}

		sig, err := crypto.Sign(ticketHash[:], alicePK)
		if err != nil {
			t.Fatalf("could not sign hash: %v", err)
		}

		aliceDepositsBefore, err := aliceClient.Deposits(aliceClient.Address())
		if err != nil {
			t.Fatalf("could not get deposits for alice: %v", err)
		}

		bobBalanceBefore, err := bobClient.BalanceOf(bobClient.Address())
		if err != nil {
			t.Fatalf("could not get balance for bob: %v", err)
		}

		tx, err := bobClient.Redeem(ticket, bobRand, sig)
		if err != nil {
			t.Fatalf("could not redeem ticket: %v", err)
		}
		backend.Commit()

		_, err = bobClient.CheckTx(ctx, tx)
		if err != nil {
			t.Fatalf("could not check transaction: %v", err)
		}

		aliceDepositsAfter, err := aliceClient.Deposits(aliceClient.Address())
		if err != nil {
			t.Fatalf("could not get deposits for alice: %v", err)
		}

		bobBalanceAfter, err := bobClient.BalanceOf(bobClient.Address())
		if err != nil {
			t.Fatalf("could not get balance for bob: %v", err)
		}

		if !sylopayments.BigIntsEqual(aliceDepositsAfter.Escrow, new(big.Int).Add(aliceDepositsBefore.Escrow, new(big.Int).Neg(ticket.FaceValue))) {
			t.Fatalf("alice's escrow is %v: expected %v", aliceDepositsAfter.Escrow, new(big.Int).Add(aliceDepositsBefore.Escrow, new(big.Int).Neg(ticket.FaceValue)))
		}
		if !sylopayments.BigIntsEqual(bobBalanceAfter, new(big.Int).Add(bobBalanceBefore, ticket.FaceValue)) {
			t.Fatalf("bob's balance is %v: expected %v", bobBalanceAfter, new(big.Int).Add(bobBalanceBefore, ticket.FaceValue))
		}
	})

	t.Run("cannot replay ticket", func(t *testing.T) {
		aliceClient, alicePK := sylopayments.CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, aliceClient.Address(), sylopayments.OneEth, big.NewInt(1000000))
		sylopayments.TopUpDeposits(t, ctx, backend, aliceClient)

		bobClient, _ := sylopayments.CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, bobClient.Address(), sylopayments.OneEth, big.NewInt(0))

		sylopayments.List(t, ctx, backend, bobClient, "0.0.0.0/0", big.NewInt(1))

		sylopayments.DelegateStake(t, ctx, backend, aliceClient, bobClient.Address(), big.NewInt(600))

		bobRand := big.NewInt(1)
		var bobRandHash [32]byte
		copy(bobRandHash[:], crypto.Keccak256(bobRand.FillBytes(bobRandHash[:])))

		ticket := contracts.SyloTicketingTicket{
			Sender:           aliceClient.Address(),
			Receiver:         bobClient.Address(),
			ReceiverRandHash: bobRandHash,
			FaceValue:        big.NewInt(1),
			WinProb:          sylopayments.Uint256max, // always win
			ExpirationBlock:  big.NewInt(0),
			SenderNonce:      1,
		}

		ticketHash, err := aliceClient.GetTicketHash(ticket)
		if err != nil {
			t.Fatalf("could not get ticket hash: %v", err)
		}

		sig, err := crypto.Sign(ticketHash[:], alicePK)
		if err != nil {
			t.Fatalf("could not sign hash: %v", err)
		}

		// good redemption
		tx, err := bobClient.Redeem(ticket, bobRand, sig)
		if err != nil {
			t.Fatalf("could not redeem ticket: %v", err)
		}
		backend.Commit()

		_, err = bobClient.CheckTx(ctx, tx)
		if err != nil {
			t.Fatalf("could not confirm transaction: %v", err)
		}

		// replay redemption
		_, err = bobClient.Redeem(ticket, bobRand, sig)
		if err == nil {
			t.Fatalf("expected error because ticket has already been used")
		}
		if !strings.HasSuffix(err.Error(), "Ticket already redeemed") {
			t.Fatalf("could not redeem: %v", err)
		}
	})

	t.Run("can stake and unstake", func(t *testing.T) {
		stakeAmount := big.NewInt(1000)

		aliceClient, _ := sylopayments.CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, aliceClient.Address(), sylopayments.OneEth, big.NewInt(1000000))

		sylopayments.Stake(t, ctx, backend, aliceClient, stakeAmount)
		defer sylopayments.UnstakeAll(t, ctx, backend, aliceClient)

		tx, err := aliceClient.UnlockStake(stakeAmount, aliceClient.Address())
		if err != nil {
			t.Fatalf("could not unlock stake: %v", err)
		}
		backend.Commit()

		_, err = aliceClient.WithdrawStake(aliceClient.Address())
		if err == nil {
			t.Fatalf("expected error because stake not yet unlocked")
		}
		if !strings.HasSuffix(err.Error(), "Stake not yet unlocked") {
			t.Fatalf("could not unstake: %v", err)
		}

		_, err = aliceClient.CheckTx(ctx, tx)
		if err != nil {
			t.Fatalf("could not check transaction: %v", err)
		}

		// all the stake should be unlocking
		unlocking, err := aliceClient.GetUnlockingStake(aliceClient.Address(), aliceClient.Address())
		if err != nil {
			t.Fatalf("could not check unlocking status: %v", err)
		}

		if !sylopayments.BigIntsEqual(unlocking.Amount, stakeAmount) {
			t.Fatalf("unlocking amount should be zero")
		}

		// advance enough blocks for the unlock period to end
		for i := uint64(0); i < unlockDuration.Uint64(); i++ {
			backend.Commit()
		}

		balanceBefore, err := aliceClient.BalanceOf(aliceClient.Address())
		if err != nil {
			t.Fatalf("could not check balance: %v", err)
		}

		// return the unstaked amount
		tx, err = aliceClient.WithdrawStake(aliceClient.Address())
		if err != nil {
			t.Fatalf("could not unstake: %v", err)
		}
		backend.Commit()

		_, err = aliceClient.CheckTx(ctx, tx)
		if err != nil {
			t.Fatalf("could not check transaction: %v", err)
		}

		// no stake should be unlocking anymore
		unlocking, err = aliceClient.GetUnlockingStake(aliceClient.Address(), aliceClient.Address())
		if err != nil {
			t.Fatalf("could not check unlocking status: %v", err)
		}

		if !sylopayments.BigIntsEqual(unlocking.Amount, big.NewInt(0)) {
			t.Fatalf("unlocking amount should be zero")
		}
		if !sylopayments.BigIntsEqual(unlocking.UnlockAt, big.NewInt(0)) {
			t.Fatalf("unlocking at should be zero")
		}

		balanceAfter, err := aliceClient.BalanceOf(aliceClient.Address())
		if err != nil {
			t.Fatalf("could not check balance: %v", err)
		}

		// check the token balance has increased
		if !sylopayments.BigIntsEqual(balanceAfter, new(big.Int).Add(balanceBefore, stakeAmount)) {
			t.Fatalf("expected stake to be returned")
		}

		// try to return the unstaked amount again
		_, err = aliceClient.WithdrawStake(aliceClient.Address())
		if err == nil {
			t.Fatalf("expected error because should not be able to unstake again")
		}
		if !strings.HasSuffix(err.Error(), "No amount to withdraw") {
			t.Fatalf("could not unstake: %v", err)
		}
	})

	t.Run("can cancel unstaking", func(t *testing.T) {
		stakeAmount := big.NewInt(1000)

		aliceClient, _ := sylopayments.CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, aliceClient.Address(), sylopayments.OneEth, big.NewInt(1000000))

		sylopayments.Stake(t, ctx, backend, aliceClient, stakeAmount)
		defer sylopayments.UnstakeAll(t, ctx, backend, aliceClient)

		tx, err := aliceClient.UnlockStake(stakeAmount, aliceClient.Address())
		if err != nil {
			t.Fatalf("could not unlock stake: %v", err)
		}
		backend.Commit()

		_, err = aliceClient.CheckTx(ctx, tx)
		if err != nil {
			t.Fatalf("could not check transaction: %v", err)
		}

		// locking the unlocking amount should cancel the unlocking
		_, err = aliceClient.CancelUnlocking(stakeAmount, aliceClient.Address())
		if err != nil {
			t.Fatalf("could not lock stake: %v", err)
		}
		backend.Commit()

		// no stake should be unlocking
		unlocking, err := aliceClient.GetUnlockingStake(aliceClient.Address(), aliceClient.Address())
		if err != nil {
			t.Fatalf("could not check unlocking status: %v", err)
		}
		if !sylopayments.BigIntsEqual(unlocking.Amount, big.NewInt(0)) {
			t.Fatalf("unlocking amount should be zero")
		}
		if !sylopayments.BigIntsEqual(unlocking.UnlockAt, big.NewInt(0)) {
			t.Fatalf("unlockAt should be zero")
		}

		// unlock the stake again
		_, err = aliceClient.UnlockStake(stakeAmount, aliceClient.Address())
		if err != nil {
			t.Fatalf("could not unlock stake: %v", err)
		}
		backend.Commit()

		if !sylopayments.WaitForUnlockAt(t, ctx, backend, aliceClient) {
			t.Fatalf("nothing to wait for")
		}

		// locking the unlocked amount should restake
		_, err = aliceClient.CancelUnlocking(stakeAmount, aliceClient.Address())
		if err != nil {
			t.Fatalf("could not lock stake: %v", err)
		}
		backend.Commit()

		// no stake should be unlocking
		unlocking, err = aliceClient.GetUnlockingStake(aliceClient.Address(), aliceClient.Address())
		if err != nil {
			t.Fatalf("could not check unlocking status: %v", err)
		}
		if !sylopayments.BigIntsEqual(unlocking.Amount, big.NewInt(0)) {
			t.Fatalf("unlocking amount should be zero")
		}
		if !sylopayments.BigIntsEqual(unlocking.UnlockAt, big.NewInt(0)) {
			t.Fatalf("unlocking at should be zero")
		}
		stakedAmount, err := aliceClient.GetAmountStaked(aliceClient.Address())
		if err != nil {
			t.Fatalf("could not check amount staked: %v", err)
		}
		if !sylopayments.BigIntsEqual(stakedAmount, stakeAmount) {
			t.Fatalf("should have the original amount staked")
		}
	})
}

func TestScan(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	backend, addresses, faucet, ownerClient := sylopayments.StartupEthereum(t, ctx)

	zeroHalves := big.NewInt(0)
	oneHalf, _ := new(big.Int).SetString("170141183460469231731687303715884105727", 10)        // 1 * ((2 << 128) // 2 - 1) // 2 + 0
	oneHalfPlusOne, _ := new(big.Int).SetString("170141183460469231731687303715884105728", 10) // 1 * ((2 << 128) // 2 - 1) // 2 + 1
	twoHalves, _ := new(big.Int).SetString("340282366920938463463374607431768211455", 10)      // 2 * ((2 << 128) // 2 - 1) // 2 + 0

	zeroThirds := big.NewInt(0)
	oneThird, _ := new(big.Int).SetString("113427455640312821154458202477256070485", 10)         // 1 * ((3 << 128) // 3 - 1) // 3 + 0
	oneThirdPlusOne, _ := new(big.Int).SetString("113427455640312821154458202477256070486", 10)  // 1 * ((3 << 128) // 3 - 1) // 3 + 1
	twoThirds, _ := new(big.Int).SetString("226854911280625642308916404954512140970", 10)        // 2 * ((3 << 128) // 3 - 1) // 3 + 0
	twoThirdsPlusOne, _ := new(big.Int).SetString("226854911280625642308916404954512140971", 10) // 2 * ((3 << 128) // 3 - 1) // 3 + 1
	threeThirds, _ := new(big.Int).SetString("340282366920938463463374607431768211455", 10)      // 3 * ((3 << 128) // 3 - 1) // 3 + 0

	t.Run("can scan empty stake tree", func(t *testing.T) {
		aliceClient, _ := sylopayments.CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, aliceClient.Address(), sylopayments.OneEth, big.NewInt(1000000))
		backend.Commit()

		scanTests := []*big.Int{
			big.NewInt(0),
			new(big.Int).Add(new(big.Int).Lsh(big.NewInt(1), 128), big.NewInt(-1)),
		}
		zeroAddr := ethcommon.BytesToAddress([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

		for _, i := range scanTests {
			a, err := aliceClient.Scan(i)
			if err != nil {
				t.Fatalf("could not scan %v: %v", i, err)
			}
			if !bytes.Equal(a.Bytes(), zeroAddr.Bytes()) {
				t.Fatalf("should scan our own address")
			}
		}
	})

	t.Run("can stake and be scanned", func(t *testing.T) {
		aliceClient, _ := sylopayments.CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, aliceClient.Address(), sylopayments.OneEth, big.NewInt(1000000))
		backend.Commit()

		stakeAmount := big.NewInt(100)
		sylopayments.Stake(t, ctx, backend, aliceClient, stakeAmount)
		defer sylopayments.UnstakeAll(t, ctx, backend, aliceClient)

		voteAmount := big.NewInt(1)
		sylopayments.Vote(t, ctx, backend, aliceClient, voteAmount)

		sylopayments.CalculatePrices(t, ctx, backend, ownerClient)
		sylopayments.ConstructDirectory(t, ctx, backend, ownerClient)

		aliceNode, _ := sylopayments.GetNode(t, aliceClient)
		if !sylopayments.BigIntsEqual(aliceNode.Amount, stakeAmount) {
			t.Fatalf("stake amount is not correct")
		}

		scanTests := []*big.Int{
			big.NewInt(0),
			new(big.Int).Lsh(big.NewInt(1), 64),
			new(big.Int).Add(new(big.Int).Lsh(big.NewInt(1), 128), big.NewInt(-1)),
		}

		for _, i := range scanTests {
			a, err := aliceClient.Scan(i)
			if err != nil {
				t.Fatalf("could not scan %v: %v", i, err)
			}
			if !bytes.Equal(a.Bytes(), aliceClient.Address().Bytes()) {
				t.Fatalf("should scan our own address")
			}
		}
	})

	t.Run("can stake 2 nodes and be scanned", func(t *testing.T) {
		aliceClient, _ := sylopayments.CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, aliceClient.Address(), sylopayments.OneEth, big.NewInt(1000000))
		backend.Commit()

		aliceStakeAmount := big.NewInt(100)
		sylopayments.Stake(t, ctx, backend, aliceClient, aliceStakeAmount)
		defer sylopayments.UnstakeAll(t, ctx, backend, aliceClient)

		aliceNode, _ := sylopayments.GetNode(t, aliceClient)
		if !sylopayments.BigIntsEqual(aliceNode.Amount, aliceStakeAmount) {
			t.Fatalf("stake amount is not correct")
		}

		bobClient, _ := sylopayments.CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, bobClient.Address(), sylopayments.OneEth, big.NewInt(1000000))
		backend.Commit()

		bobStakeAmount := big.NewInt(100)
		sylopayments.Stake(t, ctx, backend, bobClient, bobStakeAmount)
		defer sylopayments.UnstakeAll(t, ctx, backend, bobClient)

		aliceNode, _ = sylopayments.GetNode(t, aliceClient)
		if !sylopayments.BigIntsEqual(aliceNode.Amount, aliceStakeAmount) {
			t.Fatalf("alice stake amount is not correct")
		}

		voteAmount := big.NewInt(1)
		sylopayments.Vote(t, ctx, backend, aliceClient, voteAmount)
		sylopayments.Vote(t, ctx, backend, bobClient, voteAmount)

		sylopayments.CalculatePrices(t, ctx, backend, ownerClient)

		sylopayments.ConstructDirectory(t, ctx, backend, ownerClient)

		bobNode, _ := sylopayments.GetNode(t, bobClient)
		if !sylopayments.BigIntsEqual(bobNode.Amount, bobStakeAmount) {
			t.Fatalf("bob stake amount is not correct")
		}

		scanTests := [](struct {
			desc string
			val  *big.Int
			addr ethcommon.Address
		}){
			// bob
			{desc: "zeroHalves should scan to alice", val: zeroHalves, addr: aliceClient.Address()},
			{desc: "oneHalf should scan to alice", val: oneHalf, addr: aliceClient.Address()},
			// alice
			{desc: "oneHalfPlusOne should scan to bob", val: oneHalfPlusOne, addr: bobClient.Address()},
			{desc: "twoHalves should scan to bob", val: twoHalves, addr: bobClient.Address()},
		}

		for _, scanTest := range scanTests {
			t.Run(scanTest.desc, func(t *testing.T) {
				a, err := aliceClient.Scan(scanTest.val)
				if err != nil {
					t.Fatalf("could not scan %v: %v", scanTest.val, err)
				}
				if !bytes.Equal(a.Bytes(), scanTest.addr.Bytes()) {
					t.Fatalf("scanned the wrong address for %v", scanTest.val)
				}
			})
		}
	})

	t.Run("can stake 3 nodes and be scanned", func(t *testing.T) {
		var err error

		// stake Alice
		aliceClient, _ := sylopayments.CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, aliceClient.Address(), sylopayments.OneEth, big.NewInt(1000000))
		backend.Commit()

		aliceStakeAmount := big.NewInt(100)
		sylopayments.Stake(t, ctx, backend, aliceClient, aliceStakeAmount)
		defer sylopayments.UnstakeAll(t, ctx, backend, aliceClient)

		aliceNode, _ := sylopayments.GetNode(t, aliceClient)
		if !sylopayments.BigIntsEqual(aliceNode.Amount, aliceStakeAmount) {
			t.Fatalf("stake amount is not correct")
		}

		// stake Bob
		bobClient, _ := sylopayments.CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, bobClient.Address(), sylopayments.OneEth, big.NewInt(1000000))
		backend.Commit()

		bobStakeAmount := big.NewInt(100)
		sylopayments.Stake(t, ctx, backend, bobClient, bobStakeAmount)
		defer sylopayments.UnstakeAll(t, ctx, backend, bobClient)

		aliceNode, _ = sylopayments.GetNode(t, aliceClient)
		if !sylopayments.BigIntsEqual(aliceNode.Amount, aliceStakeAmount) {
			t.Fatalf("alice stake amount is not correct")
		}
		bobNode, _ := sylopayments.GetNode(t, bobClient)
		if !sylopayments.BigIntsEqual(bobNode.Amount, bobStakeAmount) {
			t.Fatalf("bob stake amount is not correct")
		}

		// stake Charlie
		charlieClient, _ := sylopayments.CreateRandomClient(t, ctx, backend, addresses)
		faucet(t, charlieClient.Address(), sylopayments.OneEth, big.NewInt(1000000))
		backend.Commit()

		charlieStakeAmount := big.NewInt(100)
		sylopayments.Stake(t, ctx, backend, charlieClient, charlieStakeAmount)
		defer sylopayments.UnstakeAll(t, ctx, backend, charlieClient)

		aliceNode, _ = sylopayments.GetNode(t, aliceClient)
		if !sylopayments.BigIntsEqual(aliceNode.Amount, aliceStakeAmount) {
			t.Fatalf("alice stake amount is not correct")
		}
		bobNode, _ = sylopayments.GetNode(t, bobClient)
		if !sylopayments.BigIntsEqual(bobNode.Amount, bobStakeAmount) {
			t.Fatalf("bob stake amount is not correct")
		}

		charlieNode, _ := sylopayments.GetNode(t, charlieClient)
		if !sylopayments.BigIntsEqual(charlieNode.Amount, charlieStakeAmount) {
			t.Fatalf("charlie stake amount is not correct")
		}

		voteAmount := big.NewInt(1)
		sylopayments.Vote(t, ctx, backend, aliceClient, voteAmount)
		sylopayments.Vote(t, ctx, backend, bobClient, voteAmount)
		sylopayments.Vote(t, ctx, backend, charlieClient, voteAmount)

		sylopayments.CalculatePrices(t, ctx, backend, ownerClient)

		sylopayments.ConstructDirectory(t, ctx, backend, ownerClient)

		scanTests := [](struct {
			desc string
			val  *big.Int
			addr ethcommon.Address
		}){
			// charlie
			{desc: "zeroThirds should scan to alice", val: zeroThirds, addr: aliceClient.Address()},
			{desc: "oneThird should scan to alice", val: oneThird, addr: aliceClient.Address()},
			// alice
			{desc: "oneThirdPlusOne should scan to bob", val: oneThirdPlusOne, addr: bobClient.Address()},
			{desc: "twoThirds should scan to bob", val: twoThirds, addr: bobClient.Address()},
			// bob
			{desc: "twoThirdsPlusOne should scan to charlie", val: twoThirdsPlusOne, addr: charlieClient.Address()},
			{desc: "threeThirds should scan to charlie", val: threeThirds, addr: charlieClient.Address()},
		}

		for _, scanTest := range scanTests {
			t.Run(scanTest.desc, func(t *testing.T) {
				a, err := aliceClient.Scan(scanTest.val)
				if err != nil {
					t.Fatalf("could not scan %v: %v", scanTest.val, err)
				}
				if !bytes.Equal(a.Bytes(), scanTest.addr.Bytes()) {
					t.Fatalf("scanned the wrong address for %v", scanTest.val)
				}
			})
		}

		// unlock Alice's stake
		tx, err := aliceClient.UnlockStake(aliceStakeAmount, aliceClient.Address())
		if err != nil {
			t.Fatalf("could not unlock stake: %v", err)
		}
		backend.Commit()
		_, err = aliceClient.CheckTx(ctx, tx)
		if err != nil {
			t.Fatalf("could not check transaction: %v", err)
		}

		bobNode, _ = sylopayments.GetNode(t, bobClient)
		if !sylopayments.BigIntsEqual(bobNode.Amount, bobStakeAmount) {
			t.Fatalf("bob's stake amount is not correct")
		}

		// reconstruct directory
		sylopayments.ConstructDirectory(t, ctx, backend, ownerClient)

		scanTests = [](struct {
			desc string
			val  *big.Int
			addr ethcommon.Address
		}){
			// charlie
			{desc: "zeroHalves should scan to bob now", val: zeroHalves, addr: bobClient.Address()},
			{desc: "oneHalfMinusOne should scan to bob now", val: oneHalf, addr: bobClient.Address()},
			// bob
			{desc: "oneHalf should scan to charli now", val: oneHalfPlusOne, addr: charlieClient.Address()},
			{desc: "twoHalves should scan to charlie still", val: twoHalves, addr: charlieClient.Address()},
		}

		for _, scanTest := range scanTests {
			t.Run(scanTest.desc, func(t *testing.T) {
				a, err := aliceClient.Scan(scanTest.val)
				if err != nil {
					t.Fatalf("could not scan %v: %v", scanTest.val, err)
				}
				if !bytes.Equal(a.Bytes(), scanTest.addr.Bytes()) {
					t.Fatalf("scanned the wrong address for %v", scanTest.val)
				}
			})
		}
	})
}

func prettyPrintNodeInfo(t *testing.T, ctx context.Context, client sylopayments.Client, desc string) {
	key, err := client.GetKey(client.Address(), client.Address())
	if err != nil {
		t.Fatalf("could not get key: %v", err)
	}
	node, err := client.Stakes(key)
	if err != nil {
		t.Fatalf("could not get node info: %v", err)
	}
	keyStr := base64.RawStdEncoding.EncodeToString(key[:])
	t.Logf("%s (%v): Stake amount=%v", desc, keyStr, node.Amount)
}

var _ = prettyPrintNodeInfo
