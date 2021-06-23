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
const ListingsABI = "[{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"multiAddr\",\"type\":\"string\"}],\"internalType\":\"structListings.Listing\",\"name\":\"listing\",\"type\":\"tuple\"}],\"name\":\"setListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getListing\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"multiAddr\",\"type\":\"string\"}],\"internalType\":\"structListings.Listing\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ListingsBin is the compiled bytecode used for deploying new contracts.
var ListingsBin = "0x608060405234801561001057600080fd5b50610539806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063084af0b2146100465780634ec1d70b1461006f5780638129fc1c14610084575b600080fd5b610059610054366004610307565b61008c565b60405161006691906103fc565b60405180910390f35b61008261007d366004610342565b610162565b005b61008261018e565b60408051602080820183526060825273ffffffffffffffffffffffffffffffffffffffff84166000908152600182528390208351918201909352825491929091829082906100d9906104b2565b80601f0160208091040260200160405190810160405280929190818152602001828054610105906104b2565b80156101525780601f1061012757610100808354040283529160200191610152565b820191906000526020600020905b81548152906001019060200180831161013557829003601f168201915b5050505050815250509050919050565b33600090815260016020908152604090912082518051849361018892849291019061026e565b50505050565b600054610100900460ff16806101a7575060005460ff16155b610237576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840160405180910390fd5b600054610100900460ff16158015610259576000805461ffff19166101011790555b801561026b576000805461ff00191690555b50565b82805461027a906104b2565b90600052602060002090601f01602090048101928261029c57600085556102e2565b82601f106102b557805160ff19168380011785556102e2565b828001600101855582156102e2579182015b828111156102e25782518255916020019190600101906102c7565b506102ee9291506102f2565b5090565b5b808211156102ee57600081556001016102f3565b600060208284031215610318578081fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461033b578182fd5b9392505050565b60006020808385031215610354578182fd5b823567ffffffffffffffff8082111561036b578384fd5b818501915082828703121561037e578384fd5b610386610458565b823582811115610394578586fd5b80840193505086601f8401126103a8578485fd5b8235828111156103ba576103ba6104ed565b6103cc601f8201601f19168601610481565b925080835287858286010111156103e1578586fd5b80858501868501378201909301939093525090815292915050565b60006020808352835181828501528051806040860152835b8181101561043057828101840151868201606001528301610414565b818111156104415784606083880101525b50601f01601f191693909301606001949350505050565b6040516020810167ffffffffffffffff8111828210171561047b5761047b6104ed565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156104aa576104aa6104ed565b604052919050565b600181811c908216806104c657607f821691505b602082108114156104e757634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fdfea26469706673582212203b945837de723ea9cd74af970d35dbfad7becb8bbf4f04fb7a1e5dd8db7d039764736f6c63430008040033"

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
func (_Listings *ListingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_Listings *ListingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
// Solidity: function getListing(address account) view returns((string))
func (_Listings *ListingsCaller) GetListing(opts *bind.CallOpts, account common.Address) (ListingsListing, error) {
	var out []interface{}
	err := _Listings.contract.Call(opts, &out, "getListing", account)

	if err != nil {
		return *new(ListingsListing), err
	}

	out0 := *abi.ConvertType(out[0], new(ListingsListing)).(*ListingsListing)

	return out0, err

}

// GetListing is a free data retrieval call binding the contract method 0x084af0b2.
//
// Solidity: function getListing(address account) view returns((string))
func (_Listings *ListingsSession) GetListing(account common.Address) (ListingsListing, error) {
	return _Listings.Contract.GetListing(&_Listings.CallOpts, account)
}

// GetListing is a free data retrieval call binding the contract method 0x084af0b2.
//
// Solidity: function getListing(address account) view returns((string))
func (_Listings *ListingsCallerSession) GetListing(account common.Address) (ListingsListing, error) {
	return _Listings.Contract.GetListing(&_Listings.CallOpts, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Listings *ListingsTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Listings.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Listings *ListingsSession) Initialize() (*types.Transaction, error) {
	return _Listings.Contract.Initialize(&_Listings.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Listings *ListingsTransactorSession) Initialize() (*types.Transaction, error) {
	return _Listings.Contract.Initialize(&_Listings.TransactOpts)
}

// SetListing is a paid mutator transaction binding the contract method 0x4ec1d70b.
//
// Solidity: function setListing((string) listing) returns()
func (_Listings *ListingsTransactor) SetListing(opts *bind.TransactOpts, listing ListingsListing) (*types.Transaction, error) {
	return _Listings.contract.Transact(opts, "setListing", listing)
}

// SetListing is a paid mutator transaction binding the contract method 0x4ec1d70b.
//
// Solidity: function setListing((string) listing) returns()
func (_Listings *ListingsSession) SetListing(listing ListingsListing) (*types.Transaction, error) {
	return _Listings.Contract.SetListing(&_Listings.TransactOpts, listing)
}

// SetListing is a paid mutator transaction binding the contract method 0x4ec1d70b.
//
// Solidity: function setListing((string) listing) returns()
func (_Listings *ListingsTransactorSession) SetListing(listing ListingsListing) (*types.Transaction, error) {
	return _Listings.Contract.SetListing(&_Listings.TransactOpts, listing)
}
