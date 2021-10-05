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
	MultiAddr         string
	PayoutPercentage  uint16
	MinDelegatedStake *big.Int
	Initialized       bool
}

// ListingsABI is the input ABI used to generate the binding from.
const ListingsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"defaultPayoutPercentage\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_defaultPayoutPercentage\",\"type\":\"uint16\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_defaultPayoutPercentage\",\"type\":\"uint16\"}],\"name\":\"setDefaultPayoutPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"multiAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minDelegatedStake\",\"type\":\"uint256\"}],\"name\":\"setListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getListing\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"multiAddr\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"payoutPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"minDelegatedStake\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"internalType\":\"structListings.Listing\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ListingsBin is the compiled bytecode used for deploying new contracts.
var ListingsBin = "0x608060405234801561001057600080fd5b50610b60806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063715018a61161005b578063715018a6146100f15780638da5cb5b146100f9578063d2a78d7f14610114578063f2fde38b1461013557600080fd5b8063084af0b21461008d57806313750946146100b6578063624903c4146100cb5780636f319073146100de575b600080fd5b6100a061009b36600461094c565b610148565b6040516100ad9190610a4e565b60405180910390f35b6100c96100c4366004610a2c565b61024e565b005b6100c96100d9366004610a2c565b610318565b6100c96100ec36600461097a565b610406565b6100c96104e7565b6033546040516001600160a01b0390911681526020016100ad565b6066546101229061ffff1681565b60405161ffff90911681526020016100ad565b6100c961014336600461094c565b61054d565b6040805160808101825260608082526000602083018190529282018390528101919091526001600160a01b0382166000908152606560205260409081902081516080810190925280548290829061019e90610ad9565b80601f01602080910402602001604051908101604052809291908181526020018280546101ca90610ad9565b80156102175780601f106101ec57610100808354040283529160200191610217565b820191906000526020600020905b8154815290600101906020018083116101fa57829003601f168201915b5050509183525050600182015461ffff1660208201526002820154604082015260039091015460ff16151560609091015292915050565b600054610100900460ff1680610267575060005460ff16155b6102cf5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b600054610100900460ff161580156102f1576000805461ffff19166101011790555b6102f961062f565b61030282610318565b8015610314576000805461ff00191690555b5050565b6033546001600160a01b031633146103725760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102c6565b6127108161ffff1611156103ee5760405162461bcd60e51b815260206004820152603960248201527f546865207061796f75742070657263656e74616765206d75737420626520612060448201527f76616c7565206265747765656e203020616e642031303030300000000000000060648201526084016102c6565b6066805461ffff191661ffff92909216919091179055565b81516104545760405162461bcd60e51b815260206004820152601960248201527f4d756c74696164647220737472696e6720697320656d7074790000000000000060448201526064016102c6565b6040805160808101825283815260665461ffff1660208083019190915281830184905260016060830152336000908152606582529290922081518051929384936104a192849201906108b3565b50602082015160018201805461ffff191661ffff909216919091179055604082015160028201556060909101516003909101805460ff1916911515919091179055505050565b6033546001600160a01b031633146105415760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102c6565b61054b60006106f1565b565b6033546001600160a01b031633146105a75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102c6565b6001600160a01b0381166106235760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102c6565b61062c816106f1565b50565b600054610100900460ff1680610648575060005460ff16155b6106ab5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102c6565b600054610100900460ff161580156106cd576000805461ffff19166101011790555b6106d561075b565b6106dd61080c565b801561062c576000805461ff001916905550565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff1680610774575060005460ff16155b6107d75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102c6565b600054610100900460ff161580156106dd576000805461ffff1916610101179055801561062c576000805461ff001916905550565b600054610100900460ff1680610825575060005460ff16155b6108885760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102c6565b600054610100900460ff161580156108aa576000805461ffff19166101011790555b6106dd336106f1565b8280546108bf90610ad9565b90600052602060002090601f0160209004810192826108e15760008555610927565b82601f106108fa57805160ff1916838001178555610927565b82800160010185558215610927579182015b8281111561092757825182559160200191906001019061090c565b50610933929150610937565b5090565b5b808211156109335760008155600101610938565b60006020828403121561095d578081fd5b81356001600160a01b0381168114610973578182fd5b9392505050565b6000806040838503121561098c578081fd5b823567ffffffffffffffff808211156109a3578283fd5b818501915085601f8301126109b6578283fd5b8135818111156109c8576109c8610b14565b604051601f8201601f19908116603f011681019083821181831017156109f0576109f0610b14565b81604052828152886020848701011115610a08578586fd5b82602086016020830137918201602090810195909552509694909201359450505050565b600060208284031215610a3d578081fd5b813561ffff81168114610973578182fd5b60006020808352835160808285015280518060a0860152835b81811015610a835782810184015186820160c001528301610a67565b81811115610a94578460c083880101525b509185015161ffff81166040860152916040860151606086015260608601519250610ac3608086018415159052565b601f01601f19169390930160c001949350505050565b600181811c90821680610aed57607f821691505b60208210811415610b0e57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fdfea264697066735822122095bf443ced08f75ae61fc4f93fe05604e1b0e176b5ab96065e3cfde692a80a2364736f6c63430008040033"

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

// DefaultPayoutPercentage is a free data retrieval call binding the contract method 0xd2a78d7f.
//
// Solidity: function defaultPayoutPercentage() view returns(uint16)
func (_Listings *ListingsCaller) DefaultPayoutPercentage(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Listings.contract.Call(opts, &out, "defaultPayoutPercentage")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DefaultPayoutPercentage is a free data retrieval call binding the contract method 0xd2a78d7f.
//
// Solidity: function defaultPayoutPercentage() view returns(uint16)
func (_Listings *ListingsSession) DefaultPayoutPercentage() (uint16, error) {
	return _Listings.Contract.DefaultPayoutPercentage(&_Listings.CallOpts)
}

// DefaultPayoutPercentage is a free data retrieval call binding the contract method 0xd2a78d7f.
//
// Solidity: function defaultPayoutPercentage() view returns(uint16)
func (_Listings *ListingsCallerSession) DefaultPayoutPercentage() (uint16, error) {
	return _Listings.Contract.DefaultPayoutPercentage(&_Listings.CallOpts)
}

// GetListing is a free data retrieval call binding the contract method 0x084af0b2.
//
// Solidity: function getListing(address account) view returns((string,uint16,uint256,bool))
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
// Solidity: function getListing(address account) view returns((string,uint16,uint256,bool))
func (_Listings *ListingsSession) GetListing(account common.Address) (ListingsListing, error) {
	return _Listings.Contract.GetListing(&_Listings.CallOpts, account)
}

// GetListing is a free data retrieval call binding the contract method 0x084af0b2.
//
// Solidity: function getListing(address account) view returns((string,uint16,uint256,bool))
func (_Listings *ListingsCallerSession) GetListing(account common.Address) (ListingsListing, error) {
	return _Listings.Contract.GetListing(&_Listings.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Listings *ListingsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Listings.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Listings *ListingsSession) Owner() (common.Address, error) {
	return _Listings.Contract.Owner(&_Listings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Listings *ListingsCallerSession) Owner() (common.Address, error) {
	return _Listings.Contract.Owner(&_Listings.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x13750946.
//
// Solidity: function initialize(uint16 _defaultPayoutPercentage) returns()
func (_Listings *ListingsTransactor) Initialize(opts *bind.TransactOpts, _defaultPayoutPercentage uint16) (*types.Transaction, error) {
	return _Listings.contract.Transact(opts, "initialize", _defaultPayoutPercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0x13750946.
//
// Solidity: function initialize(uint16 _defaultPayoutPercentage) returns()
func (_Listings *ListingsSession) Initialize(_defaultPayoutPercentage uint16) (*types.Transaction, error) {
	return _Listings.Contract.Initialize(&_Listings.TransactOpts, _defaultPayoutPercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0x13750946.
//
// Solidity: function initialize(uint16 _defaultPayoutPercentage) returns()
func (_Listings *ListingsTransactorSession) Initialize(_defaultPayoutPercentage uint16) (*types.Transaction, error) {
	return _Listings.Contract.Initialize(&_Listings.TransactOpts, _defaultPayoutPercentage)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Listings *ListingsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Listings.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Listings *ListingsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Listings.Contract.RenounceOwnership(&_Listings.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Listings *ListingsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Listings.Contract.RenounceOwnership(&_Listings.TransactOpts)
}

// SetDefaultPayoutPercentage is a paid mutator transaction binding the contract method 0x624903c4.
//
// Solidity: function setDefaultPayoutPercentage(uint16 _defaultPayoutPercentage) returns()
func (_Listings *ListingsTransactor) SetDefaultPayoutPercentage(opts *bind.TransactOpts, _defaultPayoutPercentage uint16) (*types.Transaction, error) {
	return _Listings.contract.Transact(opts, "setDefaultPayoutPercentage", _defaultPayoutPercentage)
}

// SetDefaultPayoutPercentage is a paid mutator transaction binding the contract method 0x624903c4.
//
// Solidity: function setDefaultPayoutPercentage(uint16 _defaultPayoutPercentage) returns()
func (_Listings *ListingsSession) SetDefaultPayoutPercentage(_defaultPayoutPercentage uint16) (*types.Transaction, error) {
	return _Listings.Contract.SetDefaultPayoutPercentage(&_Listings.TransactOpts, _defaultPayoutPercentage)
}

// SetDefaultPayoutPercentage is a paid mutator transaction binding the contract method 0x624903c4.
//
// Solidity: function setDefaultPayoutPercentage(uint16 _defaultPayoutPercentage) returns()
func (_Listings *ListingsTransactorSession) SetDefaultPayoutPercentage(_defaultPayoutPercentage uint16) (*types.Transaction, error) {
	return _Listings.Contract.SetDefaultPayoutPercentage(&_Listings.TransactOpts, _defaultPayoutPercentage)
}

// SetListing is a paid mutator transaction binding the contract method 0x6f319073.
//
// Solidity: function setListing(string multiAddr, uint256 minDelegatedStake) returns()
func (_Listings *ListingsTransactor) SetListing(opts *bind.TransactOpts, multiAddr string, minDelegatedStake *big.Int) (*types.Transaction, error) {
	return _Listings.contract.Transact(opts, "setListing", multiAddr, minDelegatedStake)
}

// SetListing is a paid mutator transaction binding the contract method 0x6f319073.
//
// Solidity: function setListing(string multiAddr, uint256 minDelegatedStake) returns()
func (_Listings *ListingsSession) SetListing(multiAddr string, minDelegatedStake *big.Int) (*types.Transaction, error) {
	return _Listings.Contract.SetListing(&_Listings.TransactOpts, multiAddr, minDelegatedStake)
}

// SetListing is a paid mutator transaction binding the contract method 0x6f319073.
//
// Solidity: function setListing(string multiAddr, uint256 minDelegatedStake) returns()
func (_Listings *ListingsTransactorSession) SetListing(multiAddr string, minDelegatedStake *big.Int) (*types.Transaction, error) {
	return _Listings.Contract.SetListing(&_Listings.TransactOpts, multiAddr, minDelegatedStake)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Listings *ListingsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Listings.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Listings *ListingsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Listings.Contract.TransferOwnership(&_Listings.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Listings *ListingsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Listings.Contract.TransferOwnership(&_Listings.TransactOpts, newOwner)
}

// ListingsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Listings contract.
type ListingsOwnershipTransferredIterator struct {
	Event *ListingsOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ListingsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ListingsOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ListingsOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ListingsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ListingsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ListingsOwnershipTransferred represents a OwnershipTransferred event raised by the Listings contract.
type ListingsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Listings *ListingsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ListingsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Listings.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ListingsOwnershipTransferredIterator{contract: _Listings.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Listings *ListingsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ListingsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Listings.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ListingsOwnershipTransferred)
				if err := _Listings.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Listings *ListingsFilterer) ParseOwnershipTransferred(log types.Log) (*ListingsOwnershipTransferred, error) {
	event := new(ListingsOwnershipTransferred)
	if err := _Listings.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
