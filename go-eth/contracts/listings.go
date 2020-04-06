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
	MultiAddr string
}

// ListingsABI is the input ABI used to generate the binding from.
const ListingsABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"multiAddr\",\"type\":\"string\"}],\"internalType\":\"structListings.Listing\",\"name\":\"listing\",\"type\":\"tuple\"}],\"name\":\"setListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getListing\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"multiAddr\",\"type\":\"string\"}],\"internalType\":\"structListings.Listing\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ListingsBin is the compiled bytecode used for deploying new contracts.
var ListingsBin = "0x608060405234801561001057600080fd5b506103c0806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063084af0b21461003b5780634ec1d70b14610064575b600080fd5b61004e610049366004610216565b610079565b60405161005b9190610307565b60405180910390f35b610077610072366004610251565b61013e565b005b610081610168565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260208181526040918290208251815460026001821615610100026000190190911604601f8101849004840282018501855292810183815290939192849284919084018282801561012e5780601f106101035761010080835404028352916020019161012e565b820191906000526020600020905b81548152906001019060200180831161011157829003601f168201915b5050505050815250509050919050565b3360009081526020818152604090912082518051849361016292849291019061017b565b50505050565b6040518060200160405280606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106101bc57805160ff19168380011785556101e9565b828001600101855582156101e9579182015b828111156101e95782518255916020019190600101906101ce565b506101f59291506101f9565b5090565b61021391905b808211156101f557600081556001016101ff565b90565b600060208284031215610227578081fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461024a578182fd5b9392505050565b60006020808385031215610263578182fd5b823567ffffffffffffffff8082111561027a578384fd5b81850183818803121561028b578485fd5b61029484610363565b92508035828111156102a4578586fd5b80820188601f8201126102b5578687fd5b80359250838311156102c5578687fd5b6102d7601f8401601f19168701610363565b935082845288868483010111156102ec578687fd5b82868201878601375050810190920192909252815292915050565b60006020808352835181828501528051806040860152835b8181101561033b5782810184015186820160600152830161031f565b8181111561034c5784606083880101525b50601f01601f191693909301606001949350505050565b60405181810167ffffffffffffffff8111828210171561038257600080fd5b60405291905056fea264697066735822122056f781367fa8d44b160ecd483757b93cb5e1fe943518b83f4b9566d7a737f26064736f6c63430006040033"

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

// SetListing is a paid mutator transaction binding the contract method 0x4ec1d70b.
//
// Solidity: function setListing(ListingsListing listing) returns()
func (_Listings *ListingsTransactor) SetListing(opts *bind.TransactOpts, listing ListingsListing) (*types.Transaction, error) {
	return _Listings.contract.Transact(opts, "setListing", listing)
}

// SetListing is a paid mutator transaction binding the contract method 0x4ec1d70b.
//
// Solidity: function setListing(ListingsListing listing) returns()
func (_Listings *ListingsSession) SetListing(listing ListingsListing) (*types.Transaction, error) {
	return _Listings.Contract.SetListing(&_Listings.TransactOpts, listing)
}

// SetListing is a paid mutator transaction binding the contract method 0x4ec1d70b.
//
// Solidity: function setListing(ListingsListing listing) returns()
func (_Listings *ListingsTransactorSession) SetListing(listing ListingsListing) (*types.Transaction, error) {
	return _Listings.Contract.SetListing(&_Listings.TransactOpts, listing)
}
