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
		&cli.StringFlag{
			Name:  "face-value",
			Usage: "A uint256 value representing the face value of a winning ticket",
		},
		&cli.StringFlag{
			Name:  "win-prob",
			Usage: "The numerator used in the win probability calculation, where p = `VALUE` / 2^256",
		},
		&cli.StringFlag{
			Name:  "decay-rate",
			Usage: "The `PERCENTAGE` of a ticket's probability that may decay over it's lifetime expressed as a fraction of 10000",
		},
		&cli.StringFlag{
			Name:  "expired-win-prob",
			Usage: "The numerator used in the win probability calculation after ticket expiration, where p = `VALUE` / 2^256",
		},
		&cli.IntFlag{
			Name:  "ticket-length",
			Usage: "The `NUM` of blocks that tickets stay valid for",
		},
		&cli.IntFlag{
			Name:  "payout-percentage",
			Usage: "The `PERCENTAGE` of ticket rewards that is paid to delegated stakers expressed as a fraction of 10000",
			Value: 5000,
		},
		&cli.IntFlag{
			Name:  "epochs-duration",
			Usage: "The duration an epoch will last for in `BLOCKS`",
			Value: 80000,
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

		faceValue, ok := new(big.Int).SetString(c.String("face-value"), 10)
		if !ok {
			return fmt.Errorf("could not parse integer from %s", c.String("face-value"))
		}

		winProb, ok := new(big.Int).SetString(c.String("win-prob"), 10)
		if !ok {
			return fmt.Errorf("could not parse integer from %s", c.String("win-prob"))
		}

		decayRate := c.Int("decay-rate")
		// bound to a value between 0 and 10000
		if decayRate < 0 {
			decayRate = 0
		} else if decayRate > 10000 {
			decayRate = 10000
		}

		expiredWinProb, ok := new(big.Int).SetString(c.String("min-prob-constant"), 10)
		if !ok {
			return fmt.Errorf("could not parse integer from %s", c.String("win-prob"))
		}

		ticketLength := new(big.Int).SetInt64(int64(c.Int("ticket-length")))

		payoutPercentage := c.Int("payout-percentage")
		// bound to a value between 0 and a 10000
		if payoutPercentage < 0 {
			payoutPercentage = 0
		} else if payoutPercentage > 10000 {
			payoutPercentage = 10000
		}

		epochsDuration := new(big.Int).SetInt64(int64(c.Int("epochs-duration")))

		ethSKstr := c.String("eth-sk")
		if strings.TrimSpace(ethSKstr) == "" {
			return fmt.Errorf("ethereum secret key must be provided")
		}
		m.ethSK, err = crypto.HexToECDSA(ethSKstr)
		if err != nil {
			return fmt.Errorf("could not decode private key hex string (%s): %w", ethSKstr, err)
		}
		err = m.start(c.String("eth-url"), unlockDuration, faceValue, winProb, expiredWinProb, uint16(decayRate), ticketLength, uint16(payoutPercentage), epochsDuration)
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

func (m *syloEthMgr) start(url string, unlockDuration *big.Int, faceValue *big.Int, winProb *big.Int, expiredWinProb *big.Int, decayRate uint16, ticketLength *big.Int, payoutPercentage uint16, epochsDuration *big.Int) error {
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

	m.addrs, err = deployContracts(m.opts.Context, m.opts, m.ethC, unlockDuration, faceValue, winProb, expiredWinProb, decayRate, ticketLength, payoutPercentage, epochsDuration)
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

func deployContracts(ctx context.Context, opts *bind.TransactOpts, client *ethclient.Client, unlockDuration *big.Int, faceValue *big.Int, winProb *big.Int, expiredWinProb *big.Int, decayRate uint16, ticketLength *big.Int, payoutPercentage uint16, epochsDuration *big.Int) (addresses sylo.Addresses, err error) {
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

	// deploy staking manager
	var stakingManagerTx *types.Transaction
	var stakingManager *contracts.StakingManager
	addresses.Directory, stakingManagerTx, stakingManager, err = contracts.DeployStakingManager(opts, client)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy stakingManager: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// deploy directory
	var directoryTx *types.Transaction
	var directory *contracts.Directory
	addresses.Directory, directoryTx, directory, err = contracts.DeployDirectory(opts, client)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy directory: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// deploy epochs manager
	var epochsManagerTx *types.Transaction
	var epochsManager *contracts.EpochsManager
	addresses.EpochsManager, epochsManagerTx, epochsManager, err = contracts.DeployEpochsManager(opts, client)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy epochsManager: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// deploy rewards manager
	var rewardsManagerTx *types.Transaction
	var rewardsManager *contracts.RewardsManager
	addresses.RewardsManager, rewardsManagerTx, rewardsManager, err = contracts.DeployRewardsManager(opts, client)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy rewardsManager: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// initialize staking manager
	_, err = stakingManager.Initialize(opts, addresses.Token, addresses.RewardsManager, addresses.EpochsManager, unlockDuration)
	if err != nil {
		return addresses, fmt.Errorf("could not staking manager: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// initialize directory
	_, err = directory.Initialize(opts, addresses.StakingManager, addresses.RewardsManager)
	if err != nil {
		return addresses, fmt.Errorf("could not initialise directory: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// initialize rewards manager
	_, err = rewardsManager.Initialize(opts, addresses.Token, addresses.StakingManager, addresses.EpochsManager)
	if err != nil {
		return addresses, fmt.Errorf("could not initialise rewardsManager: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// deploy price voting
	var priceVotingTx *types.Transaction
	var priceVoting *contracts.PriceVoting
	addresses.PriceVoting, priceVotingTx, priceVoting, err = contracts.DeployPriceVoting(opts, client)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy priceVoting: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	_, err = priceVoting.Initialize(opts, addresses.StakingManager)
	if err != nil {
		return addresses, fmt.Errorf("could not initialise price voting: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// deploy price manager
	var priceManagerTx *types.Transaction
	var priceManager *contracts.PriceManager
	addresses.Directory, priceManagerTx, priceManager, err = contracts.DeployPriceManager(opts, client)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy priceManager: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	_, err = priceManager.Initialize(opts, addresses.StakingManager, addresses.PriceVoting)
	if err != nil {
		return addresses, fmt.Errorf("could not initialise price manager: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// deploy listing
	var listingTx *types.Transaction
	var listings *contracts.Listings
	addresses.Listings, listingTx, listings, err = contracts.DeployListings(opts, client)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy listing: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	_, err = listings.Initialize(opts, payoutPercentage)
	if err != nil {
		return addresses, fmt.Errorf("could not initialise listing: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// deploy ticketing parameters
	var ticketingParametersTx *types.Transaction
	var ticketingParameters *contracts.TicketingParameters
	addresses.TicketingParameters, ticketingParametersTx, ticketingParameters, err = contracts.DeployTicketingParameters(opts, client)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy ticketingParameters: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	_, err = ticketingParameters.Initialize(opts, faceValue, winProb, expiredWinProb, decayRate, ticketLength)
	if err != nil {
		return addresses, fmt.Errorf("could not initialise ticketingParameters: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	_, err = epochsManager.Initialize(opts, addresses.Directory, addresses.Listings, addresses.TicketingParameters, epochsDuration)
	if err != nil {
		return addresses, fmt.Errorf("could not initialise epochsManager: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	// deploy ticketing
	var ticketingTx *types.Transaction
	var ticketing *contracts.SyloTicketing
	addresses.Ticketing, ticketingTx, ticketing, err = contracts.DeploySyloTicketing(opts, client)
	if err != nil {
		return addresses, fmt.Errorf("could not deploy ticketing: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	_, err = ticketing.Initialize(opts, addresses.Token, addresses.Listings, addresses.StakingManager, addresses.Directory, addresses.EpochsManager, addresses.RewardsManager, unlockDuration)
	if err != nil {
		return addresses, fmt.Errorf("could not initialise ticketing: %w", err)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	_, err = rewardsManager.AddManager(opts, addresses.Ticketing)
	if err != nil {
		return addresses, fmt.Errorf("could not add ticketing contract as manager: %v", err)
	}

	// wait for receipts
	_, err = waitForReceipt(ctx, tokenTx, client)
	if err != nil {
		return addresses, fmt.Errorf("could not get token deployment receipt: %w", err)
	}
	_, err = waitForReceipt(ctx, ticketingParametersTx, client)
	if err != nil {
		return addresses, fmt.Errorf("could not get transaction receipt: %w", err)
	}
	_, err = waitForReceipt(ctx, epochsManagerTx, client)
	if err != nil {
		return addresses, fmt.Errorf("could not get transaction receipt: %w", err)
	}
	_, err = waitForReceipt(ctx, ticketingTx, client)
	if err != nil {
		return addresses, fmt.Errorf("could not get transaction receipt: %w", err)
	}
	_, err = waitForReceipt(ctx, stakingManagerTx, client)
	if err != nil {
		return addresses, fmt.Errorf("could not get transaction receipt: %w", err)
	}
	_, err = waitForReceipt(ctx, priceManagerTx, client)
	if err != nil {
		return addresses, fmt.Errorf("could not get transaction receipt: %w", err)
	}
	_, err = waitForReceipt(ctx, priceVotingTx, client)
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
	_, err = waitForReceipt(ctx, rewardsManagerTx, client)
	if err != nil {
		return addresses, fmt.Errorf("could not get transaction recept: %w", err)
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
