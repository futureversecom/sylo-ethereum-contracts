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

// PriceManagerABI is the input ABI used to generate the binding from.
const PriceManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"currentServicePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentUpperPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractStakingManager\",\"name\":\"stakingManager\",\"type\":\"address\"},{\"internalType\":\"contractPriceVoting\",\"name\":\"voting\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"sortedIndexes\",\"type\":\"uint256[]\"}],\"name\":\"calculatePrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"servicePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PriceManagerBin is the compiled bytecode used for deploying new contracts.
var PriceManagerBin = "0x60806040526000606755600060685534801561001a57600080fd5b50610f508061002a6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638da5cb5b1161005b5780638da5cb5b146100bb578063985371a3146100d6578063d146f086146100df578063f2fde38b1461010757600080fd5b80630b5b820714610082578063485cc9551461009e578063715018a6146100b3575b600080fd5b61008b60675481565b6040519081526020015b60405180910390f35b6100b16100ac366004610d27565b61011a565b005b6100b1610224565b6033546040516001600160a01b039091168152602001610095565b61008b60685481565b6100f26100ed366004610ca0565b61028a565b60408051928352602083019190915201610095565b6100b1610115366004610bc2565b6103b2565b600054610100900460ff1680610133575060005460ff16155b61019b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b600054610100900460ff161580156101bd576000805461ffff19166101011790555b6101c5610494565b606580546001600160a01b038086167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556066805492851692909116919091179055801561021f576000805461ff00191690555b505050565b6033546001600160a01b0316331461027e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610192565b6102886000610556565b565b600080336001600160a01b03166102a96033546001600160a01b031690565b6001600160a01b0316146102ff5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610192565b6066546040517f3a1231e10000000000000000000000000000000000000000000000000000000081526000916001600160a01b031690633a1231e190610349908790600401610d77565b60006040518083038186803b15801561036157600080fd5b505afa158015610375573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261039d9190810190610bde565b90506103a8816105c0565b9250925050915091565b6033546001600160a01b0316331461040c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610192565b6001600160a01b0381166104885760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610192565b61049181610556565b50565b600054610100900460ff16806104ad575060005460ff16155b6105105760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610192565b600054610100900460ff16158015610532576000805461ffff19166101011790555b61053a610a34565b610542610ae5565b8015610491576000805461ff001916905550565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600080336001600160a01b03166105df6033546001600160a01b031690565b6001600160a01b0316146106355760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610192565b606554604080517f7ffdacd500000000000000000000000000000000000000000000000000000000815290516000926001600160a01b031691637ffdacd5916004808301926020929190829003018186803b15801561069357600080fd5b505afa1580156106a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106cb9190610d5f565b905060006106db826109c4610b8c565b905060006106eb83612328610b8c565b90506000805b87518110156108735760655488516000916001600160a01b0316906313cdd31b908b908590811061073257634e487b7160e01b600052603260045260246000fd5b6020908102919091010151516040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b03909116600482015260240160206040518083038186803b15801561079557600080fd5b505afa1580156107a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107cd9190610d5f565b9050806107da5750610861565b8882815181106107fa57634e487b7160e01b600052603260045260246000fd5b602002602001015160200151600014156108145750610861565b61081e8184610e39565b925084831061085f5788828151811061084757634e487b7160e01b600052603260045260246000fd5b60200260200101516020015160678190555050610873565b505b8061086b81610ebe565b9150506106f1565b505085516fffffffffffffffffffffffffffffffff8416905b8015610a20576065546000906001600160a01b03166313cdd31b8a6108b2600186610e90565b815181106108d057634e487b7160e01b600052603260045260246000fd5b6020908102919091010151516040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b03909116600482015260240160206040518083038186803b15801561093357600080fd5b505afa158015610947573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061096b9190610d5f565b9050806109785750610a0e565b88610984600184610e90565b815181106109a257634e487b7160e01b600052603260045260246000fd5b602002602001015160200151600014156109bc5750610a0e565b6109c68184610e90565b9250838311610a0c57886109db600184610e90565b815181106109f957634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516068819055505b505b80610a1881610ea7565b91505061088c565b506067546068549550955050505050915091565b600054610100900460ff1680610a4d575060005460ff16155b610ab05760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610192565b600054610100900460ff16158015610542576000805461ffff19166101011790558015610491576000805461ff001916905550565b600054610100900460ff1680610afe575060005460ff16155b610b615760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610192565b600054610100900460ff16158015610b83576000805461ffff19166101011790555b61054233610556565b6000612710610bb161ffff84166fffffffffffffffffffffffffffffffff8616610e71565b610bbb9190610e51565b9392505050565b600060208284031215610bd3578081fd5b8135610bbb81610f05565b60006020808385031215610bf0578182fd5b825167ffffffffffffffff811115610c06578283fd5b8301601f81018513610c16578283fd5b8051610c29610c2482610e15565b610de4565b80828252848201915084840188868560061b8701011115610c48578687fd5b8694505b83851015610c9457604080828b031215610c64578788fd5b610c6c610dbb565b8251610c7781610f05565b815282880151888201528452600195909501949286019201610c4c565b50979650505050505050565b60006020808385031215610cb2578182fd5b823567ffffffffffffffff811115610cc8578283fd5b8301601f81018513610cd8578283fd5b8035610ce6610c2482610e15565b80828252848201915084840188868560051b8701011115610d05578687fd5b8694505b83851015610c94578035835260019490940193918501918501610d09565b60008060408385031215610d39578081fd5b8235610d4481610f05565b91506020830135610d5481610f05565b809150509250929050565b600060208284031215610d70578081fd5b5051919050565b6020808252825182820181905260009190848201906040850190845b81811015610daf57835183529284019291840191600101610d93565b50909695505050505050565b6040805190810167ffffffffffffffff81118282101715610dde57610dde610eef565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610e0d57610e0d610eef565b604052919050565b600067ffffffffffffffff821115610e2f57610e2f610eef565b5060051b60200190565b60008219821115610e4c57610e4c610ed9565b500190565b600082610e6c57634e487b7160e01b81526012600452602481fd5b500490565b6000816000190483118215151615610e8b57610e8b610ed9565b500290565b600082821015610ea257610ea2610ed9565b500390565b600081610eb657610eb6610ed9565b506000190190565b6000600019821415610ed257610ed2610ed9565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461049157600080fdfea2646970667358221220bd1de584c86f01af2c5a34a290da9b1c0b3f282e41c5c3fe8715898c506322a964736f6c63430008040033"

// DeployPriceManager deploys a new Ethereum contract, binding an instance of PriceManager to it.
func DeployPriceManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PriceManager, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PriceManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PriceManager{PriceManagerCaller: PriceManagerCaller{contract: contract}, PriceManagerTransactor: PriceManagerTransactor{contract: contract}, PriceManagerFilterer: PriceManagerFilterer{contract: contract}}, nil
}

// PriceManager is an auto generated Go binding around an Ethereum contract.
type PriceManager struct {
	PriceManagerCaller     // Read-only binding to the contract
	PriceManagerTransactor // Write-only binding to the contract
	PriceManagerFilterer   // Log filterer for contract events
}

// PriceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceManagerSession struct {
	Contract     *PriceManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceManagerCallerSession struct {
	Contract *PriceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PriceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceManagerTransactorSession struct {
	Contract     *PriceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PriceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceManagerRaw struct {
	Contract *PriceManager // Generic contract binding to access the raw methods on
}

// PriceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceManagerCallerRaw struct {
	Contract *PriceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// PriceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceManagerTransactorRaw struct {
	Contract *PriceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceManager creates a new instance of PriceManager, bound to a specific deployed contract.
func NewPriceManager(address common.Address, backend bind.ContractBackend) (*PriceManager, error) {
	contract, err := bindPriceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceManager{PriceManagerCaller: PriceManagerCaller{contract: contract}, PriceManagerTransactor: PriceManagerTransactor{contract: contract}, PriceManagerFilterer: PriceManagerFilterer{contract: contract}}, nil
}

// NewPriceManagerCaller creates a new read-only instance of PriceManager, bound to a specific deployed contract.
func NewPriceManagerCaller(address common.Address, caller bind.ContractCaller) (*PriceManagerCaller, error) {
	contract, err := bindPriceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceManagerCaller{contract: contract}, nil
}

// NewPriceManagerTransactor creates a new write-only instance of PriceManager, bound to a specific deployed contract.
func NewPriceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceManagerTransactor, error) {
	contract, err := bindPriceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceManagerTransactor{contract: contract}, nil
}

// NewPriceManagerFilterer creates a new log filterer instance of PriceManager, bound to a specific deployed contract.
func NewPriceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceManagerFilterer, error) {
	contract, err := bindPriceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceManagerFilterer{contract: contract}, nil
}

// bindPriceManager binds a generic wrapper to an already deployed contract.
func bindPriceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceManager *PriceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceManager.Contract.PriceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceManager *PriceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceManager.Contract.PriceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceManager *PriceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceManager.Contract.PriceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceManager *PriceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceManager *PriceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceManager *PriceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceManager.Contract.contract.Transact(opts, method, params...)
}

// CurrentServicePrice is a free data retrieval call binding the contract method 0x0b5b8207.
//
// Solidity: function currentServicePrice() view returns(uint256)
func (_PriceManager *PriceManagerCaller) CurrentServicePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceManager.contract.Call(opts, &out, "currentServicePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentServicePrice is a free data retrieval call binding the contract method 0x0b5b8207.
//
// Solidity: function currentServicePrice() view returns(uint256)
func (_PriceManager *PriceManagerSession) CurrentServicePrice() (*big.Int, error) {
	return _PriceManager.Contract.CurrentServicePrice(&_PriceManager.CallOpts)
}

// CurrentServicePrice is a free data retrieval call binding the contract method 0x0b5b8207.
//
// Solidity: function currentServicePrice() view returns(uint256)
func (_PriceManager *PriceManagerCallerSession) CurrentServicePrice() (*big.Int, error) {
	return _PriceManager.Contract.CurrentServicePrice(&_PriceManager.CallOpts)
}

// CurrentUpperPrice is a free data retrieval call binding the contract method 0x985371a3.
//
// Solidity: function currentUpperPrice() view returns(uint256)
func (_PriceManager *PriceManagerCaller) CurrentUpperPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceManager.contract.Call(opts, &out, "currentUpperPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentUpperPrice is a free data retrieval call binding the contract method 0x985371a3.
//
// Solidity: function currentUpperPrice() view returns(uint256)
func (_PriceManager *PriceManagerSession) CurrentUpperPrice() (*big.Int, error) {
	return _PriceManager.Contract.CurrentUpperPrice(&_PriceManager.CallOpts)
}

// CurrentUpperPrice is a free data retrieval call binding the contract method 0x985371a3.
//
// Solidity: function currentUpperPrice() view returns(uint256)
func (_PriceManager *PriceManagerCallerSession) CurrentUpperPrice() (*big.Int, error) {
	return _PriceManager.Contract.CurrentUpperPrice(&_PriceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceManager *PriceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceManager *PriceManagerSession) Owner() (common.Address, error) {
	return _PriceManager.Contract.Owner(&_PriceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceManager *PriceManagerCallerSession) Owner() (common.Address, error) {
	return _PriceManager.Contract.Owner(&_PriceManager.CallOpts)
}

// CalculatePrices is a paid mutator transaction binding the contract method 0xd146f086.
//
// Solidity: function calculatePrices(uint256[] sortedIndexes) returns(uint256 servicePrice, uint256 upperPrice)
func (_PriceManager *PriceManagerTransactor) CalculatePrices(opts *bind.TransactOpts, sortedIndexes []*big.Int) (*types.Transaction, error) {
	return _PriceManager.contract.Transact(opts, "calculatePrices", sortedIndexes)
}

// CalculatePrices is a paid mutator transaction binding the contract method 0xd146f086.
//
// Solidity: function calculatePrices(uint256[] sortedIndexes) returns(uint256 servicePrice, uint256 upperPrice)
func (_PriceManager *PriceManagerSession) CalculatePrices(sortedIndexes []*big.Int) (*types.Transaction, error) {
	return _PriceManager.Contract.CalculatePrices(&_PriceManager.TransactOpts, sortedIndexes)
}

// CalculatePrices is a paid mutator transaction binding the contract method 0xd146f086.
//
// Solidity: function calculatePrices(uint256[] sortedIndexes) returns(uint256 servicePrice, uint256 upperPrice)
func (_PriceManager *PriceManagerTransactorSession) CalculatePrices(sortedIndexes []*big.Int) (*types.Transaction, error) {
	return _PriceManager.Contract.CalculatePrices(&_PriceManager.TransactOpts, sortedIndexes)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address stakingManager, address voting) returns()
func (_PriceManager *PriceManagerTransactor) Initialize(opts *bind.TransactOpts, stakingManager common.Address, voting common.Address) (*types.Transaction, error) {
	return _PriceManager.contract.Transact(opts, "initialize", stakingManager, voting)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address stakingManager, address voting) returns()
func (_PriceManager *PriceManagerSession) Initialize(stakingManager common.Address, voting common.Address) (*types.Transaction, error) {
	return _PriceManager.Contract.Initialize(&_PriceManager.TransactOpts, stakingManager, voting)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address stakingManager, address voting) returns()
func (_PriceManager *PriceManagerTransactorSession) Initialize(stakingManager common.Address, voting common.Address) (*types.Transaction, error) {
	return _PriceManager.Contract.Initialize(&_PriceManager.TransactOpts, stakingManager, voting)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceManager *PriceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceManager *PriceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceManager.Contract.RenounceOwnership(&_PriceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceManager *PriceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceManager.Contract.RenounceOwnership(&_PriceManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceManager *PriceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PriceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceManager *PriceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceManager.Contract.TransferOwnership(&_PriceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceManager *PriceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceManager.Contract.TransferOwnership(&_PriceManager.TransactOpts, newOwner)
}

// PriceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PriceManager contract.
type PriceManagerOwnershipTransferredIterator struct {
	Event *PriceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PriceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceManagerOwnershipTransferred)
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
		it.Event = new(PriceManagerOwnershipTransferred)
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
func (it *PriceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the PriceManager contract.
type PriceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceManager *PriceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PriceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PriceManagerOwnershipTransferredIterator{contract: _PriceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceManager *PriceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PriceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceManagerOwnershipTransferred)
				if err := _PriceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PriceManager *PriceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*PriceManagerOwnershipTransferred, error) {
	event := new(PriceManagerOwnershipTransferred)
	if err := _PriceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
