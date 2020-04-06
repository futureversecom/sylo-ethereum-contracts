// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ListingsListing is an auto generated low-level Go binding around an user-defined struct.
type ListingsListing struct {
	SyloId    [32]byte
	MultiAddr [32]byte
}

// ListingsABI is the input ABI used to generate the binding from.
const ListingsABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"syloId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"multiAddr\",\"type\":\"bytes32\"}],\"internalType\":\"structListings.Listing\",\"name\":\"listing\",\"type\":\"tuple\"}],\"name\":\"setListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getListing\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"syloId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"multiAddr\",\"type\":\"bytes32\"}],\"internalType\":\"structListings.Listing\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ListingsBin is the compiled bytecode used for deploying new contracts.
var ListingsBin = "0x608060405234801561001057600080fd5b506101c7806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063084af0b21461003b5780636243102014610064575b600080fd5b61004e6100493660046100f5565b610079565b60405161005b919061017a565b60405180910390f35b610077610072366004610130565b6100bf565b005b6100816100de565b5073ffffffffffffffffffffffffffffffffffffffff1660009081526020818152604091829020825180840190935280548352600101549082015290565b3360009081526020818152604090912082518155910151600190910155565b604080518082019091526000808252602082015290565b600060208284031215610106578081fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610129578182fd5b9392505050565b600060408284031215610141578081fd5b6040516040810181811067ffffffffffffffff82111715610160578283fd5b604052823581526020928301359281019290925250919050565b81518152602091820151918101919091526040019056fea2646970667358221220017471bcca4f71a12c3cd8eba11b64bf0d05419d084409903ac3fe6044ac6c4064736f6c63430006040033"

// DeployListings deploys a new Ethereum contract, binding an instance of Listings to it.
func DeployListings(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Listings, error) {
	parsed, err := abi.JSON(strings.NewReader(ListingsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ListingsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Listings{ListingsCaller: ListingsCaller{contract: contract}, ListingsTransactor: ListingsTransactor{contract: contract}, ListingsFilterer: ListingsFilterer{contract: contract}}, nil
}

// Listings is an auto generated Go binding around an Ethereum contract.
type Listings struct {
	ListingsCaller     // Read-only binding to the contract
	ListingsTransactor // Write-only binding to the contract
	ListingsFilterer   // Log filterer for contract events
}

// ListingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ListingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ListingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ListingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ListingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ListingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ListingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ListingsSession struct {
	Contract     *Listings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ListingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ListingsCallerSession struct {
	Contract *ListingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ListingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ListingsTransactorSession struct {
	Contract     *ListingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ListingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ListingsRaw struct {
	Contract *Listings // Generic contract binding to access the raw methods on
}

// ListingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ListingsCallerRaw struct {
	Contract *ListingsCaller // Generic read-only contract binding to access the raw methods on
}

// ListingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ListingsTransactorRaw struct {
	Contract *ListingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewListings creates a new instance of Listings, bound to a specific deployed contract.
func NewListings(address common.Address, backend bind.ContractBackend) (*Listings, error) {
	contract, err := bindListings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Listings{ListingsCaller: ListingsCaller{contract: contract}, ListingsTransactor: ListingsTransactor{contract: contract}, ListingsFilterer: ListingsFilterer{contract: contract}}, nil
}

// NewListingsCaller creates a new read-only instance of Listings, bound to a specific deployed contract.
func NewListingsCaller(address common.Address, caller bind.ContractCaller) (*ListingsCaller, error) {
	contract, err := bindListings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ListingsCaller{contract: contract}, nil
}

// NewListingsTransactor creates a new write-only instance of Listings, bound to a specific deployed contract.
func NewListingsTransactor(address common.Address, transactor bind.ContractTransactor) (*ListingsTransactor, error) {
	contract, err := bindListings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ListingsTransactor{contract: contract}, nil
}

// NewListingsFilterer creates a new log filterer instance of Listings, bound to a specific deployed contract.
func NewListingsFilterer(address common.Address, filterer bind.ContractFilterer) (*ListingsFilterer, error) {
	contract, err := bindListings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ListingsFilterer{contract: contract}, nil
}

// bindListings binds a generic wrapper to an already deployed contract.
func bindListings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ListingsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Listings *ListingsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Listings.Contract.ListingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Listings *ListingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Listings.Contract.ListingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Listings *ListingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Listings.Contract.ListingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Listings *ListingsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Listings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Listings *ListingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Listings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Listings *ListingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Listings.Contract.contract.Transact(opts, method, params...)
}

// GetListing is a free data retrieval call binding the contract method 0x084af0b2.
//
// Solidity: function getListing(address account) constant returns(ListingsListing)
func (_Listings *ListingsCaller) GetListing(opts *bind.CallOpts, account common.Address) (ListingsListing, error) {
	var (
		ret0 = new(ListingsListing)
	)
	out := ret0
	err := _Listings.contract.Call(opts, out, "getListing", account)
	return *ret0, err
}

// GetListing is a free data retrieval call binding the contract method 0x084af0b2.
//
// Solidity: function getListing(address account) constant returns(ListingsListing)
func (_Listings *ListingsSession) GetListing(account common.Address) (ListingsListing, error) {
	return _Listings.Contract.GetListing(&_Listings.CallOpts, account)
}

// GetListing is a free data retrieval call binding the contract method 0x084af0b2.
//
// Solidity: function getListing(address account) constant returns(ListingsListing)
func (_Listings *ListingsCallerSession) GetListing(account common.Address) (ListingsListing, error) {
	return _Listings.Contract.GetListing(&_Listings.CallOpts, account)
}

// SetListing is a paid mutator transaction binding the contract method 0x62431020.
//
// Solidity: function setListing(ListingsListing listing) returns()
func (_Listings *ListingsTransactor) SetListing(opts *bind.TransactOpts, listing ListingsListing) (*types.Transaction, error) {
	return _Listings.contract.Transact(opts, "setListing", listing)
}

// SetListing is a paid mutator transaction binding the contract method 0x62431020.
//
// Solidity: function setListing(ListingsListing listing) returns()
func (_Listings *ListingsSession) SetListing(listing ListingsListing) (*types.Transaction, error) {
	return _Listings.Contract.SetListing(&_Listings.TransactOpts, listing)
}

// SetListing is a paid mutator transaction binding the contract method 0x62431020.
//
// Solidity: function setListing(ListingsListing listing) returns()
func (_Listings *ListingsTransactorSession) SetListing(listing ListingsListing) (*types.Transaction, error) {
	return _Listings.Contract.SetListing(&_Listings.TransactOpts, listing)
}
