package eth_test

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/dn3010/sylo-ethereum-contracts/go-eth"
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
	gasLimit         = uint64(50000000)
	oneEth           = big.NewInt(1000000000000000000)
	faucetEthBalance = new(big.Int).Mul(oneEth, big.NewInt(1000))
	chainID          = big.NewInt(1337)
	unlockDuration   = big.NewInt(10)
	escrowAmount     = big.NewInt(100000)
	penaltyAmount    = big.NewInt(1000)
	uint256max       = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1)) // 2^256-1
)

func TestClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	backend, addresses, faucet := startupEthereum(t, ctx)
	_ = faucet // in case it isn't used

	t.Run("client can be created", func(t *testing.T) {
		createRandomClient(t, ctx, backend, addresses)
	})

	t.Run("can use faucet", func(t *testing.T) {
		aliceClient, _ := createRandomClient(t, ctx, backend, addresses)
		err := faucet(aliceClient.Address(), oneEth, big.NewInt(1000000))
		if err != nil {
			t.Fatalf("could not faucet: %v", err)
		}
	})

	t.Run("cannot faucet more eth than is available", func(t *testing.T) {
		aliceClient, _ := createRandomClient(t, ctx, backend, addresses)
		tooMuchEth := new(big.Int).Add(oneEth, faucetEthBalance)
		err := faucet(aliceClient.Address(), tooMuchEth, big.NewInt(0))
		if err == nil {
			t.Fatalf("should not be able to faucet this much: %v", err)
		}
	})

	t.Run("can get latest block", func(t *testing.T) {
		aliceClient, _ := createRandomClient(t, ctx, backend, addresses)
		err := faucet(aliceClient.Address(), oneEth, big.NewInt(0))
		if err != nil {
			t.Fatalf("could not faucet: %v", err)
		}
		blockNumberA, err := aliceClient.LatestBlock()
		if err != nil {
			t.Fatalf("could not get latest block: %v", err)
		}
		backend.Commit()
		blockNumberB, err := aliceClient.LatestBlock()
		if err != nil {
			t.Fatalf("could not get latest block: %v", err)
		}
		if !bigIntsEqual(blockNumberA.Add(blockNumberA, big.NewInt(1)), blockNumberB) {
			t.Fatalf("block number did not advance")
		}
	})

	t.Run("can deposit escrow", func(t *testing.T) {
		aliceClient, _ := createRandomClient(t, ctx, backend, addresses)
		err := faucet(aliceClient.Address(), oneEth, big.NewInt(1000000))
		if err != nil {
			t.Fatalf("could not faucet: %v", err)
		}

		addEscrow(t, ctx, backend, aliceClient, escrowAmount)

		deposit, err := aliceClient.Deposits(aliceClient.Address())
		if err != nil {
			t.Fatalf("could not get deposits: %v", err)
		}
		if !bigIntsEqual(deposit.Escrow, escrowAmount) {
			t.Fatalf("escrow deposit does not match: got %v: expected %v", deposit.Escrow, escrowAmount)
		}
	})

	t.Run("can deposit penalty", func(t *testing.T) {
		aliceClient, _ := createRandomClient(t, ctx, backend, addresses)
		err := faucet(aliceClient.Address(), oneEth, big.NewInt(1000000))
		if err != nil {
			t.Fatalf("could not faucet: %v", err)
		}

		addPenalty(t, ctx, backend, aliceClient, penaltyAmount)

		deposit, err := aliceClient.Deposits(aliceClient.Address())
		if err != nil {
			t.Fatalf("could not get deposits: %v", err)
		}
		if !bigIntsEqual(deposit.Penalty, penaltyAmount) {
			t.Fatalf("penalty deposit does not match: got %v: expected %v", deposit.Penalty, penaltyAmount)
		}
	})

	t.Run("can withdraw ticketing", func(t *testing.T) {
		aliceClient, _ := createRandomClient(t, ctx, backend, addresses)
		err := faucet(aliceClient.Address(), oneEth, big.NewInt(1000000))
		if err != nil {
			t.Fatalf("could not faucet: %v", err)
		}
		topUpDeposits(t, ctx, backend, aliceClient)

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
		if !bigIntsEqual(deposit.Escrow, big.NewInt(0)) {
			t.Fatalf("escrow should be withdrawn")
		}
		if !bigIntsEqual(deposit.Penalty, big.NewInt(0)) {
			t.Fatalf("penalty should be withdrawn")
		}
	})

	t.Run("can redeem ticket", func(t *testing.T) {
		aliceClient, alicePK := createRandomClient(t, ctx, backend, addresses)
		err := faucet(aliceClient.Address(), oneEth, big.NewInt(1000000))
		if err != nil {
			t.Fatalf("could not faucet: %v", err)
		}
		topUpDeposits(t, ctx, backend, aliceClient)

		bobClient, _ := createRandomClient(t, ctx, backend, addresses)
		err = faucet(bobClient.Address(), oneEth, big.NewInt(0))
		if err != nil {
			t.Fatalf("could not faucet: %v", err)
		}

		bobRand := big.NewInt(1)
		var bobRandHash [32]byte
		copy(bobRandHash[:], crypto.Keccak256(bobRand.FillBytes(bobRandHash[:])))

		ticket := contracts.SyloTicketingTicket{
			Sender:           aliceClient.Address(),
			Receiver:         bobClient.Address(),
			ReceiverRandHash: bobRandHash,
			FaceValue:        big.NewInt(1),
			WinProb:          uint256max, // always win
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

		if !bigIntsEqual(aliceDepositsAfter.Escrow, new(big.Int).Add(aliceDepositsBefore.Escrow, new(big.Int).Neg(ticket.FaceValue))) {
			t.Fatalf("alice's escrow is %v: expected %v", aliceDepositsAfter.Escrow, new(big.Int).Add(aliceDepositsBefore.Escrow, new(big.Int).Neg(ticket.FaceValue)))
		}
		if !bigIntsEqual(bobBalanceAfter, new(big.Int).Add(bobBalanceBefore, ticket.FaceValue)) {
			t.Fatalf("bob's balance is %v: expected %v", bobBalanceAfter, new(big.Int).Add(bobBalanceBefore, ticket.FaceValue))
		}
	})

	t.Run("cannot replay ticket", func(t *testing.T) {
		aliceClient, alicePK := createRandomClient(t, ctx, backend, addresses)
		err := faucet(aliceClient.Address(), oneEth, big.NewInt(1000000))
		if err != nil {
			t.Fatalf("could not faucet: %v", err)
		}
		topUpDeposits(t, ctx, backend, aliceClient)

		bobClient, _ := createRandomClient(t, ctx, backend, addresses)
		err = faucet(bobClient.Address(), oneEth, big.NewInt(0))
		if err != nil {
			t.Fatalf("could not faucet: %v", err)
		}

		bobRand := big.NewInt(1)
		var bobRandHash [32]byte
		copy(bobRandHash[:], crypto.Keccak256(bobRand.FillBytes(bobRandHash[:])))

		ticket := contracts.SyloTicketingTicket{
			Sender:           aliceClient.Address(),
			Receiver:         bobClient.Address(),
			ReceiverRandHash: bobRandHash,
			FaceValue:        big.NewInt(1),
			WinProb:          uint256max, // always win
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

		aliceClient, _ := createRandomClient(t, ctx, backend, addresses)
		err := faucet(aliceClient.Address(), oneEth, big.NewInt(1000000))
		if err != nil {
			t.Fatalf("could not faucet: %v", err)
		}

		stake(t, ctx, backend, aliceClient, stakeAmount)
		defer unstakeAll(t, ctx, backend, aliceClient)

		tx, err := aliceClient.UnlockStake(stakeAmount, aliceClient.Address())
		if err != nil {
			t.Fatalf("could not unlock stake: %v", err)
		}
		backend.Commit()

		_, err = aliceClient.Unstake(aliceClient.Address())
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

		if !bigIntsEqual(unlocking.Amount, stakeAmount) {
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
		tx, err = aliceClient.Unstake(aliceClient.Address())
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

		if !bigIntsEqual(unlocking.Amount, big.NewInt(0)) {
			t.Fatalf("unlocking amount should be zero")
		}
		if !bigIntsEqual(unlocking.UnlockAt, big.NewInt(0)) {
			t.Fatalf("unlocking at should be zero")
		}

		balanceAfter, err := aliceClient.BalanceOf(aliceClient.Address())
		if err != nil {
			t.Fatalf("could not check balance: %v", err)
		}

		// check the token balance has increased
		if !bigIntsEqual(balanceAfter, new(big.Int).Add(balanceBefore, stakeAmount)) {
			t.Fatalf("expected stake to be returned")
		}

		// try to return the unstaked amount again
		_, err = aliceClient.Unstake(aliceClient.Address())
		if err == nil {
			t.Fatalf("expected error because should not be able to unstake again")
		}
		if !strings.HasSuffix(err.Error(), "No amount to unlock") {
			t.Fatalf("could not unstake: %v", err)
		}
	})

	t.Run("can cancel unstaking", func(t *testing.T) {
		stakeAmount := big.NewInt(1000)

		aliceClient, _ := createRandomClient(t, ctx, backend, addresses)
		err := faucet(aliceClient.Address(), oneEth, big.NewInt(1000000))
		if err != nil {
			t.Fatalf("could not faucet: %v", err)
		}

		stake(t, ctx, backend, aliceClient, stakeAmount)
		defer unstakeAll(t, ctx, backend, aliceClient)

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
		_, err = aliceClient.LockStake(stakeAmount, aliceClient.Address())
		if err != nil {
			t.Fatalf("could not lock stake: %v", err)
		}
		backend.Commit()

		// no stake should be unlocking
		unlocking, err := aliceClient.GetUnlockingStake(aliceClient.Address(), aliceClient.Address())
		if err != nil {
			t.Fatalf("could not check unlocking status: %v", err)
		}
		if !bigIntsEqual(unlocking.Amount, big.NewInt(0)) {
			t.Fatalf("unlocking amount should be zero")
		}
		if !bigIntsEqual(unlocking.UnlockAt, big.NewInt(0)) {
			t.Fatalf("unlocking at should be zero")
		}

		// unlock the stake again
		_, err = aliceClient.UnlockStake(stakeAmount, aliceClient.Address())
		if err != nil {
			t.Fatalf("could not unlock stake: %v", err)
		}
		backend.Commit()

		if !waitForUnlockAt(t, ctx, backend, aliceClient) {
			t.Fatalf("nothing to wait for")
		}

		// locking the unlocked amount should restake
		_, err = aliceClient.LockStake(stakeAmount, aliceClient.Address())
		if err != nil {
			t.Fatalf("could not lock stake: %v", err)
		}
		backend.Commit()

		// no stake should be unlocking
		unlocking, err = aliceClient.GetUnlockingStake(aliceClient.Address(), aliceClient.Address())
		if err != nil {
			t.Fatalf("could not check unlocking status: %v", err)
		}
		if !bigIntsEqual(unlocking.Amount, big.NewInt(0)) {
			t.Fatalf("unlocking amount should be zero")
		}
		if !bigIntsEqual(unlocking.UnlockAt, big.NewInt(0)) {
			t.Fatalf("unlocking at should be zero")
		}
	})
}

func startupEthereum(t *testing.T, ctx context.Context) (eth.SimBackend, eth.Addresses, faucetF) {
	ownerPK, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("could not create ecdsa key: %v", err)
	}
	ownerTransactor, err := bind.NewKeyedTransactorWithChainID(ownerPK, chainID)
	if err != nil {
		t.Fatalf("could not create transaction signer: %v", err)
	}
	ownerTransactor.Context = ctx

	backend := createBackend(t, ctx, ownerTransactor.From)
	addresses := deployContracts(t, ctx, ownerTransactor, backend)

	ownerClient, err := eth.NewClientWithBackend(addresses, backend, ownerTransactor)
	if err != nil {
		t.Fatalf("could not create client: %v", err)
	}

	// create a faucet
	faucet := makeFaucet(t, ctx, backend, ownerClient, ownerPK)
	return backend, addresses, faucet
}

// createBackend will make a genesis block and use it to generate a new
// simulated ethereum backend.
func createBackend(t *testing.T, ctx context.Context, owner common.Address) eth.SimBackend {
	genesis := make(core.GenesisAlloc)
	genesis[owner] = core.GenesisAccount{Balance: faucetEthBalance}
	return eth.NewSimBackend(backends.NewSimulatedBackend(genesis, gasLimit))
}

// createRandomClient will generate a new ecdsa key and use it to create a Sylo
// ethereum client.
func createRandomClient(t *testing.T, ctx context.Context, backend eth.SimBackend, addresses eth.Addresses) (eth.Client, *ecdsa.PrivateKey) {
	pk, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("could not create ecdsa key: %v", err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		t.Fatalf("could not create transaction signer: %v", err)
	}
	opts.Context = ctx

	client, err := eth.NewClientWithBackend(addresses, backend, opts)
	if err != nil {
		t.Fatalf("could not create client: %v", err)
	}

	return client, pk
}

type faucetF func(recipient ethcommon.Address, ethAmt *big.Int, syloAmt *big.Int) error

func makeFaucet(t *testing.T, ctx context.Context, backend eth.SimBackend, client eth.Client, pk *ecdsa.PrivateKey) faucetF {
	return func(recipient ethcommon.Address, ethAmt *big.Int, syloAmt *big.Int) error {
		if ethAmt.Cmp(big.NewInt(0)) == 1 {
			err := backend.FaucetEth(ctx, client.Address(), recipient, pk, ethAmt)
			if err != nil {
				return fmt.Errorf("could not faucet eth: %v", err)
			}
		}
		if syloAmt.Cmp(big.NewInt(0)) == 1 {
			tx, err := client.Transfer(recipient, syloAmt)
			if err != nil {
				return fmt.Errorf("could not faucet sylo: %v", err)
			}
			backend.Commit()
			_, err = client.CheckTx(ctx, tx)
			if err != nil {
				return fmt.Errorf("could not check sylo faucet transaction: %v", err)
			}
		}
		return nil
	}
}

// topUpDeposits will ensure that both the escrow and penalty accounts have
// enough Sylo.
func topUpDeposits(t *testing.T, ctx context.Context, backend eth.SimBackend, client eth.Client) {
	deposit, err := client.Deposits(client.Address())
	if err != nil {
		t.Fatalf("could not get deposits for top up: %v", err)
	}
	// make sure there is enough escrow
	if deposit.Escrow.Cmp(escrowAmount) == -1 {
		addAmount := new(big.Int).Add(escrowAmount, new(big.Int).Neg(deposit.Escrow))
		addEscrow(t, ctx, backend, client, addAmount)
	}
	// make sure there is enough penalty
	if deposit.Penalty.Cmp(penaltyAmount) == -1 {
		addAmount := new(big.Int).Add(penaltyAmount, new(big.Int).Neg(deposit.Escrow))
		addPenalty(t, ctx, backend, client, addAmount)
	}
}

// stake will add the stakeAmount to the stake tree.
func stake(t *testing.T, ctx context.Context, backend eth.SimBackend, client eth.Client, stakeAmount *big.Int) {
	_, err := client.ApproveDirectory(stakeAmount)
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

// waitForUnlockedAt will wait for any unlocking stake to be ready to unstake.
//
// Will return true once the unlockedAt block is reached, or false if the
// unlockedAt block is 0.
func waitForUnlockAt(t *testing.T, ctx context.Context, backend eth.SimBackend, client eth.Client) bool {
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
		if bigIntsEqual(unlocking.UnlockAt, big.NewInt(0)) {
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

// unstakeAll will unlock all stake and wait for it to be unlocked and
// withdrawn.
func unstakeAll(t *testing.T, ctx context.Context, backend eth.SimBackend, client eth.Client) {
	stakeAmount, err := client.GetAmountStaked(client.Address())
	if err != nil {
		t.Fatalf("could not get staked amount: %v", err)
	}

	_, err = client.UnlockStake(stakeAmount, client.Address())
	if err == nil {
		// wait for unlocking
		if waitForUnlockAt(t, ctx, backend, client) {
			// return the unstaked amount
			tx, err := client.Unstake(client.Address())
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
		// nothing to do
	} else {
		// error unlocking
		t.Fatalf("could not unlock stake: %v", err)
	}

	// check that all the stake is gone
	stakeAmount, err = client.GetAmountStaked(client.Address())
	if err != nil {
		t.Fatalf("could not get staked amount: %v", err)
	}
	if !bigIntsEqual(stakeAmount, big.NewInt(0)) {
		t.Fatalf("all stake should be removed")
	}

	unlocking, err := client.GetUnlockingStake(client.Address(), client.Address())
	if err != nil {
		t.Fatalf("could not check unlocking status: %v", err)
	}
	if !bigIntsEqual(unlocking.Amount, big.NewInt(0)) {
		t.Fatalf("unlocking amount should be zero")
	}
	if !bigIntsEqual(unlocking.UnlockAt, big.NewInt(0)) {
		t.Fatalf("unlocking at should be zero")
	}
}

func deployContracts(t *testing.T, ctx context.Context, transactor *bind.TransactOpts, backend eth.SimBackend) eth.Addresses {
	var addresses eth.Addresses
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

	// deploy ticketing
	addresses.Ticketing, tx, _, err = contracts.DeploySyloTicketing(transactor, backend, addresses.Token, unlockDuration)
	if err != nil {
		t.Fatalf("could not deploy ticketing: %v", err)
	}
	backend.Commit()
	_, err = backend.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		t.Fatalf("could not get transaction receipt: %v", err)
	}

	// deploy directory
	addresses.Directory, tx, _, err = contracts.DeployDirectory(transactor, backend, addresses.Token, unlockDuration)
	if err != nil {
		t.Fatalf("could not deploy directory: %v", err)
	}
	backend.Commit()
	_, err = backend.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		t.Fatalf("could not get transaction receipt: %v", err)
	}

	// deploy listing
	addresses.Listings, tx, _, err = contracts.DeployListings(transactor, backend)
	if err != nil {
		t.Fatalf("could not deploy listing: %v", err)
	}
	backend.Commit()
	_, err = backend.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		t.Fatalf("could not get transaction receipt: %v", err)
	}

	return addresses
}

func addEscrow(t *testing.T, ctx context.Context, backend eth.SimBackend, client eth.Client, escrowAmount *big.Int) {
	err := addDeposit(ctx, backend, client, escrowAmount, client.DepositEscrow)
	if err != nil {
		t.Fatalf("could not add escrow amount: %v", err)
	}
}

func addPenalty(t *testing.T, ctx context.Context, backend eth.SimBackend, client eth.Client, penaltyAmount *big.Int) {
	err := addDeposit(ctx, backend, client, penaltyAmount, client.DepositPenalty)
	if err != nil {
		t.Fatalf("could not add penalty amount: %v", err)
	}
}

type depositF func(amount *big.Int, account ethcommon.Address) (*types.Transaction, error)

func addDeposit(ctx context.Context, backend eth.SimBackend, client eth.Client, amount *big.Int, f depositF) error {
	tx, err := client.ApproveTicketing(amount)
	if err != nil {
		return fmt.Errorf("could not approve penalty amount: %v", err)
	}
	backend.Commit()

	_, err = client.CheckTx(ctx, tx)
	if err != nil {
		return fmt.Errorf("could not check transaction: %v", err)
	}

	tx, err = f(amount, client.Address())
	if err != nil {
		return fmt.Errorf("could not deposit penalty: %v", err)
	}
	backend.Commit()

	_, err = client.CheckTx(ctx, tx)
	if err != nil {
		return fmt.Errorf("could not confirm penalty deposit transaction: %v", err)
	}

	return nil
}

func bigIntsEqual(x *big.Int, y *big.Int) bool {
	return x.Cmp(y) == 0
}
