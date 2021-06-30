package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"
	"strconv"

	sylo "github.com/dn3010/sylo-ethereum-contracts/go-eth"
	sylopayments "github.com/dn3010/sylo-ethereum-contracts/go-eth"
	"github.com/dn3010/sylo-ethereum-contracts/go-eth/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/ipfs/go-log"
	cli "github.com/urfave/cli/v2"
)

var logger = log.Logger("sylo-deploy-contracts")

func main() {
	app := cli.NewApp()
	app.Name = "Deploy Sylo Smart Contracts"
	app.Usage = "Deploy the Sylo contracts to an Ethereum blockchain."
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "eth-url",
			Usage: "URL for the Etherum RPC",
			Value: "http://0.0.0.0:8545",
		},
		&cli.StringFlag{
			Name:  "eth-sk",
			Usage: "The private key that will deploy the contracts.",
			Value: "150934096e7bcd0485d154edd771b4466680038a068ccca8e8b483dce8527245",
		},
		&cli.StringFlag{
			Name:  "unlock-duration",
			Usage: "The `NUM` of blocks that must pass to unlock SYLO.",
			Value: "6",
		},
		&cli.StringFlag {
			Name: "payout-percentage",
			Usage: "The percentage of ticket rewards that is paid to delegated stakers",
			Value: "50",
		},
		&cli.BoolFlag{
			Name:  "faucet",
			Usage: "Provide a SYLO/ETH faucet service",
		},
	}
	app.Action = func(c *cli.Context) error {
		var err error
		var cancel context.CancelFunc

		m := new(syloEthMgr)

		m.ctx, cancel = context.WithCancel(context.Background())
		defer cancel()

		unlockDuration, ok := new(big.Int).SetString(c.String("unlock-duration"), 10)
		if !ok {
			logger.Errorf("could not parse integer from %s", c.String("unlock-duration"))
			logger.Infof("unlock duration will be 6")
			unlockDuration = big.NewInt(6)
		}
		payoutPercentage, err := strconv.Atoi(c.String("payout-percentage"))
		if err != nil {
			logger.Errorf("could not parse integer from %s", c.String("payout-percentage"))
			logger.Infof("payout percentage will be 50")
			payoutPercentage = 50
		}
		ethSKstr := c.String("eth-sk")
		if strings.TrimSpace(ethSKstr) == "" {
			return fmt.Errorf("ethereum secret key must be provided")
		}
		m.ethSK, err = crypto.HexToECDSA(ethSKstr)
		if err != nil {
			return fmt.Errorf("could not decode private key hex string (%s): %w", ethSKstr, err)
		}
		err = m.start(c.String("eth-url"), unlockDuration, uint8(payoutPercentage))
		if err != nil {
			return fmt.Errorf("could not execute contract deployment: %w", err)
		}

		m.faucet = c.Bool("faucet")

		http.Handle("/addresses", m.andressesHandler())
		http.Handle("/add/eth", m.ethFaucetHandler())
		http.Handle("/add/sylo", m.syloFaucetHandler())

		fmt.Println("Contracts deployed.")
		fmt.Printf("Ethereum testnet is at: %s\n", c.String("eth-url"))
		fmt.Println("Sylo contract services are at: http://0.0.0.0:7116")
		return http.ListenAndServe("0.0.0.0:7116", nil)
	}
	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
	}
}

type syloEthMgr struct {
	ctx context.Context

	ethC  *ethclient.Client
	ethSK *ecdsa.PrivateKey
	syloC sylo.Client
	opts  *bind.TransactOpts

	addrs sylopayments.Addresses

	faucet bool
}

func (m *syloEthMgr) start(url string, unlockDuration *big.Int, payoutPercentage uint8) error {
	var err error

	ctx, cancel := context.WithTimeout(m.ctx, 3*time.Minute)
	defer cancel()

	m.ethC, err = ethclient.Dial(url)
	if err != nil {
		return fmt.Errorf("failed to dial ethereum client: %w", err)
	}
	chainID, err := func() (*big.Int, error) {
		for {
			// give ethClient time to come online
			select {
			case <-ctx.Done():
				return nil, fmt.Errorf("could not find eth client: %w", ctx.Err())
			case <-time.After(250 * time.Millisecond):
			}
			chainID, err := m.ethC.ChainID(ctx)
			if err != nil {
				logger.Debugf("waiting for eth client: could not get chain id: %v", err)
				continue
			}
			return chainID, nil
		}
	}()
	if err != nil {
		return fmt.Errorf("eth client did not come online: %w", err)
	}

	m.opts, err = bind.NewKeyedTransactorWithChainID(m.ethSK, chainID)
	if err != nil {
		return fmt.Errorf("could not create trasactor: %w", err)
	}
	m.opts.Context = m.ctx

	m.addrs, err = deployContracts(m.opts.Context, m.opts, m.ethC, unlockDuration, payoutPercentage)
	if err != nil {
		return fmt.Errorf("could not deploy contracts: %w", err)
	}

	m.syloC, err = sylopayments.NewClient(m.ctx, m.addrs, m.ethC, m.opts)
	if err != nil {
		return fmt.Errorf("could not set up sylo payments: %w", err)
	}

	return nil
}

func (m *syloEthMgr) andressesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(m.addrs)
		if err != nil {
			http.Error(w, "could not encode addresses", http.StatusInternalServerError)
		}
	}
}

func (m *syloEthMgr) getEth(req *faucetRequest) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic while fauceting Eth: %w", r)
		}
	}()

	gasLimit := uint64(21000) // in units
	gasPrice, err := m.ethC.SuggestGasPrice(m.ctx)
	if err != nil {
		return fmt.Errorf("could not get suggested gas price: %v", err)
	}

	nonce, err := m.ethC.PendingNonceAt(m.ctx, m.opts.From)
	if err != nil {
		return fmt.Errorf("could not get pending nonce: %v", err)
	}

	var data []byte
	tx, err := func() (*types.Transaction, error) {
		tx := types.NewTransaction(nonce, req.Recipient, req.Amount, gasLimit, gasPrice, data)
		return types.SignTx(tx, types.HomesteadSigner{}, m.ethSK)
	}()
	if err != nil {
		return fmt.Errorf("could not sign transaction: %v", err)
	}

	err = m.ethC.SendTransaction(m.ctx, tx)
	if err != nil {
		return fmt.Errorf("could not send transaction: %v", err)
	}

	// increment transactor nonce
	m.opts.Nonce.Add(m.opts.Nonce, big.NewInt(1))

	waitCtx, waitCancel := context.WithTimeout(m.ctx, time.Minute)
	defer waitCancel()

	_, err = waitForReceipt(waitCtx, tx, m.ethC)
	if err != nil {
		return fmt.Errorf("could not get receipt: %w", err)
	}

	return nil
}

func (m *syloEthMgr) getSylo(req *faucetRequest) error {
	nonce, err := m.ethC.PendingNonceAt(context.Background(), m.syloC.Address())
	if err != nil {
		return fmt.Errorf("could not get pending nonce: %w", err)
	}
	m.opts.Nonce.SetUint64(nonce)

	tx, err := m.syloC.Transfer(req.Recipient, req.Amount)
	if err != nil {
		return fmt.Errorf("could not transfer sylo: %w", err)
	}

	waitCtx, waitCancel := context.WithTimeout(m.ctx, 30*time.Second)
	defer waitCancel()

	_, err = m.syloC.CheckTx(waitCtx, tx)
	if err != nil {
		return fmt.Errorf("could not confirm transaction: %w", err)
	}

	return nil
}

type faucetRequest struct {
	Recipient common.Address `json:"recipient"`
	Amount    *big.Int       `json:"amount"`
}

func (f *syloEthMgr) ethFaucetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !f.faucet {
			http.Error(w, "eth faucet disabled", http.StatusForbidden)
			return
		}

		req := new(faucetRequest)
		switch r.Method {
		case "POST":
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				logger.Errorf("could not decode eth request: %v", err)
				http.Error(w, fmt.Sprintf("could not decode eth request: %v", err), http.StatusBadRequest)
				return
			}
			err = f.getEth(req)
			if err != nil {
				logger.Errorf("could not faucet eth: %v", err)
				http.Error(w, "could not faucet eth", http.StatusInternalServerError)
				return
			}
		default:
			http.Error(w, "only post requests are accepted", http.StatusInternalServerError)
			return
		}
	}
}

func (f *syloEthMgr) syloFaucetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !f.faucet {
			http.Error(w, "sylo faucet disabled", http.StatusForbidden)
			return
		}

		req := new(faucetRequest)
		switch r.Method {
		case "POST":
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				logger.Errorf("could not decode sylo request: %v", err)
				http.Error(w, fmt.Sprintf("could not decode sylo request: %v", err), http.StatusBadRequest)
				return
			}
			err = f.getSylo(req)
			if err != nil {
				logger.Errorf("could not faucet sylo: %v", err)
				http.Error(w, "could not faucet sylo", http.StatusInternalServerError)
				return
			}
		default:
			http.Error(w, "only post requests are accepted", http.StatusInternalServerError)
			return
		}
	}
}

func deployContracts(ctx context.Context, opts *bind.TransactOpts, client *ethclient.Client, unlockDuration *big.Int, payoutPercentage uint8) (addresses sylo.Addresses, err error) {
	// Deploying contracts can apparently panic if the transaction fails, so
	// we need to check for that.
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic during deployment of contracts: %w", r)
		}
	}()

	// get nonce
	nonce, err := client.PendingNonceAt(ctx, opts.From)
	if err != nil {
		return addresses, fmt.Errorf("could not get next nonce: %w", err)
	}

	// deploy Sylo token
	opts.Nonce = new(big.Int).SetUint64(nonce)
	var tokenTx *types.Transaction
	addresses.Token, tokenTx, _, err = contracts.DeploySyloToken(opts, client)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy sylo token: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// deploy directory
	var directoryTx *types.Transaction
	directory := &contracts.Directory{}
	addresses.Directory, directoryTx, directory, err = contracts.DeployDirectory(opts, client)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy directory: %w", err)
	}
	directory.Initialize(opts, addresses.Token, unlockDuration)
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// deploy listing
	var listingTx *types.Transaction
	listings := &contracts.Listings{}
	addresses.Listings, listingTx, listings, err = contracts.DeployListings(opts, client)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy listing: %w", err)
	}
	listings.Initialize(opts, payoutPercentage)
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// deploy ticketing
	var ticketingTx *types.Transaction
	ticketing := &contracts.SyloTicketing{}
	addresses.Ticketing, ticketingTx, ticketing, err = contracts.DeploySyloTicketing(opts, client)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy ticketing: %w", err)
	}
	ticketing.Initialize(opts, addresses.Token, addresses.Listings, addresses.Directory, unlockDuration)
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// wait for receipts
	_, err = waitForReceipt(ctx, tokenTx, client)
	if err != nil {
		return addresses, fmt.Errorf("could not get token deployment receipt: %w", err)
	}
	_, err = waitForReceipt(ctx, ticketingTx, client)
	if err != nil {
		return addresses, fmt.Errorf("could not get transaction receipt: %w", err)
	}
	_, err = waitForReceipt(ctx, directoryTx, client)
	if err != nil {
		return addresses, fmt.Errorf("could not get transaction receipt: %w", err)
	}
	_, err = waitForReceipt(ctx, listingTx, client)
	if err != nil {
		return addresses, fmt.Errorf("could not get transaction receipt: %w", err)
	}

	return addresses, nil
}

func waitForReceipt(parent context.Context, tx *types.Transaction, client *ethclient.Client) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(parent, 30*time.Second)
	defer cancel()

	for {
		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err == nil {
			logger.Info("Got receipt for contract")
			return receipt, nil
		}
		select {
		case <-ctx.Done():
			logger.Error("Could not get receipt contract before deadline")
			return nil, context.DeadlineExceeded
		case <-time.After(3 * time.Second):
		}
	}
}
