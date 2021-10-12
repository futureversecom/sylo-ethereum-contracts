// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package parameters

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TicketingParametersMetaData contains all meta data concerning the TicketingParameters contract.
var TicketingParametersMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"baseLiveWinProb\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decayRate\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expiredWinProb\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"faceValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ticketDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_baseLiveWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_expiredWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint16\",\"name\":\"_decayRate\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_ticketDuration\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_faceValue\",\"type\":\"uint256\"}],\"name\":\"setFaceValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_baseLiveWinProb\",\"type\":\"uint128\"}],\"name\":\"setBaseLiveWinProb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_expiredWinProb\",\"type\":\"uint128\"}],\"name\":\"setExpiredWinProb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ticketDuration\",\"type\":\"uint256\"}],\"name\":\"setTicketDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610a5f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c8063a8f19c141161008c578063bcbee54311610066578063bcbee5431461019b578063dedcebda146101ec578063ef8032ef14610208578063f2fde38b1461021b57600080fd5b8063a8f19c1414610154578063a90a602714610167578063a9c1f2f11461017a57600080fd5b806387bcc0c5116100bd57806387bcc0c51461011d5780638da5cb5b146101265780639e9ceeca1461014157600080fd5b806344fd9caa146100e457806358cbd4d314610100578063715018a614610115575b600080fd5b6100ed60655481565b6040519081526020015b60405180910390f35b61011361010e3660046109cd565b61022e565b005b610113610342565b6100ed60675481565b6033546040516001600160a01b0390911681526020016100f7565b61011361014f3660046109b5565b6103a8565b61011361016236600461099b565b610457565b61011361017536600461099b565b6104e6565b6068546101889061ffff1681565b60405161ffff90911681526020016100f7565b6066546101cb9070010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff1681565b6040516fffffffffffffffffffffffffffffffff90911681526020016100f7565b6066546101cb906fffffffffffffffffffffffffffffffff1681565b6101136102163660046109b5565b610583565b61011361022936600461096d565b6105e2565b600054610100900460ff1680610247575060005460ff16155b6102af5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b600054610100900460ff161580156102d1576000805461ffff19166101011790555b6102d96106c4565b60658690556fffffffffffffffffffffffffffffffff84811670010000000000000000000000000000000002908616176066556068805461ffff851661ffff19909116179055610328826103a8565b801561033a576000805461ff00191690555b505050505050565b6033546001600160a01b0316331461039c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102a6565b6103a66000610786565b565b6033546001600160a01b031633146104025760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102a6565b600081116104525760405162461bcd60e51b815260206004820152601b60248201527f5469636b6574206475726174696f6e2063616e6e6f742062652030000000000060448201526064016102a6565b606755565b6033546001600160a01b031633146104b15760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102a6565b606680546fffffffffffffffffffffffffffffffff928316700100000000000000000000000000000000029216919091179055565b6033546001600160a01b031633146105405760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102a6565b606680547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff92909216919091179055565b6033546001600160a01b031633146105dd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102a6565b606555565b6033546001600160a01b0316331461063c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102a6565b6001600160a01b0381166106b85760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102a6565b6106c181610786565b50565b600054610100900460ff16806106dd575060005460ff16155b6107405760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102a6565b600054610100900460ff16158015610762576000805461ffff19166101011790555b61076a6107f0565b6107726108a1565b80156106c1576000805461ff001916905550565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff1680610809575060005460ff16155b61086c5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102a6565b600054610100900460ff16158015610772576000805461ffff191661010117905580156106c1576000805461ff001916905550565b600054610100900460ff16806108ba575060005460ff16155b61091d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102a6565b600054610100900460ff1615801561093f576000805461ffff19166101011790555b61077233610786565b80356fffffffffffffffffffffffffffffffff8116811461096857600080fd5b919050565b60006020828403121561097e578081fd5b81356001600160a01b0381168114610994578182fd5b9392505050565b6000602082840312156109ac578081fd5b61099482610948565b6000602082840312156109c6578081fd5b5035919050565b600080600080600060a086880312156109e4578081fd5b853594506109f460208701610948565b9350610a0260408701610948565b9250606086013561ffff81168114610a18578182fd5b94979396509194608001359291505056fea26469706673582212209f845d6c706605b264b8ce9bb84c3ca95bc18ec7ad6f3d8fb5d143b9c4b5c90564736f6c63430008040033",
}

// TicketingParametersABI is the input ABI used to generate the binding from.
// Deprecated: Use TicketingParametersMetaData.ABI instead.
var TicketingParametersABI = TicketingParametersMetaData.ABI

// TicketingParametersBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TicketingParametersMetaData.Bin instead.
var TicketingParametersBin = TicketingParametersMetaData.Bin

// DeployTicketingParameters deploys a new Ethereum contract, binding an instance of TicketingParameters to it.
func DeployTicketingParameters(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TicketingParameters, error) {
	parsed, err := TicketingParametersMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TicketingParametersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TicketingParameters{TicketingParametersCaller: TicketingParametersCaller{contract: contract}, TicketingParametersTransactor: TicketingParametersTransactor{contract: contract}, TicketingParametersFilterer: TicketingParametersFilterer{contract: contract}}, nil
}

// TicketingParameters is an auto generated Go binding around an Ethereum contract.
type TicketingParameters struct {
	TicketingParametersCaller     // Read-only binding to the contract
	TicketingParametersTransactor // Write-only binding to the contract
	TicketingParametersFilterer   // Log filterer for contract events
}

// TicketingParametersCaller is an auto generated read-only Go binding around an Ethereum contract.
type TicketingParametersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TicketingParametersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TicketingParametersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TicketingParametersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TicketingParametersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TicketingParametersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TicketingParametersSession struct {
	Contract     *TicketingParameters // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TicketingParametersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TicketingParametersCallerSession struct {
	Contract *TicketingParametersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TicketingParametersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TicketingParametersTransactorSession struct {
	Contract     *TicketingParametersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TicketingParametersRaw is an auto generated low-level Go binding around an Ethereum contract.
type TicketingParametersRaw struct {
	Contract *TicketingParameters // Generic contract binding to access the raw methods on
}

// TicketingParametersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TicketingParametersCallerRaw struct {
	Contract *TicketingParametersCaller // Generic read-only contract binding to access the raw methods on
}

// TicketingParametersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TicketingParametersTransactorRaw struct {
	Contract *TicketingParametersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTicketingParameters creates a new instance of TicketingParameters, bound to a specific deployed contract.
func NewTicketingParameters(address common.Address, backend bind.ContractBackend) (*TicketingParameters, error) {
	contract, err := bindTicketingParameters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TicketingParameters{TicketingParametersCaller: TicketingParametersCaller{contract: contract}, TicketingParametersTransactor: TicketingParametersTransactor{contract: contract}, TicketingParametersFilterer: TicketingParametersFilterer{contract: contract}}, nil
}

// NewTicketingParametersCaller creates a new read-only instance of TicketingParameters, bound to a specific deployed contract.
func NewTicketingParametersCaller(address common.Address, caller bind.ContractCaller) (*TicketingParametersCaller, error) {
	contract, err := bindTicketingParameters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TicketingParametersCaller{contract: contract}, nil
}

// NewTicketingParametersTransactor creates a new write-only instance of TicketingParameters, bound to a specific deployed contract.
func NewTicketingParametersTransactor(address common.Address, transactor bind.ContractTransactor) (*TicketingParametersTransactor, error) {
	contract, err := bindTicketingParameters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TicketingParametersTransactor{contract: contract}, nil
}

// NewTicketingParametersFilterer creates a new log filterer instance of TicketingParameters, bound to a specific deployed contract.
func NewTicketingParametersFilterer(address common.Address, filterer bind.ContractFilterer) (*TicketingParametersFilterer, error) {
	contract, err := bindTicketingParameters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TicketingParametersFilterer{contract: contract}, nil
}

// bindTicketingParameters binds a generic wrapper to an already deployed contract.
func bindTicketingParameters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TicketingParametersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TicketingParameters *TicketingParametersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TicketingParameters.Contract.TicketingParametersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TicketingParameters *TicketingParametersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TicketingParameters.Contract.TicketingParametersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TicketingParameters *TicketingParametersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TicketingParameters.Contract.TicketingParametersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TicketingParameters *TicketingParametersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TicketingParameters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TicketingParameters *TicketingParametersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TicketingParameters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TicketingParameters *TicketingParametersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TicketingParameters.Contract.contract.Transact(opts, method, params...)
}

// BaseLiveWinProb is a free data retrieval call binding the contract method 0xdedcebda.
//
// Solidity: function baseLiveWinProb() view returns(uint128)
func (_TicketingParameters *TicketingParametersCaller) BaseLiveWinProb(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TicketingParameters.contract.Call(opts, &out, "baseLiveWinProb")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseLiveWinProb is a free data retrieval call binding the contract method 0xdedcebda.
//
// Solidity: function baseLiveWinProb() view returns(uint128)
func (_TicketingParameters *TicketingParametersSession) BaseLiveWinProb() (*big.Int, error) {
	return _TicketingParameters.Contract.BaseLiveWinProb(&_TicketingParameters.CallOpts)
}

// BaseLiveWinProb is a free data retrieval call binding the contract method 0xdedcebda.
//
// Solidity: function baseLiveWinProb() view returns(uint128)
func (_TicketingParameters *TicketingParametersCallerSession) BaseLiveWinProb() (*big.Int, error) {
	return _TicketingParameters.Contract.BaseLiveWinProb(&_TicketingParameters.CallOpts)
}

// DecayRate is a free data retrieval call binding the contract method 0xa9c1f2f1.
//
// Solidity: function decayRate() view returns(uint16)
func (_TicketingParameters *TicketingParametersCaller) DecayRate(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _TicketingParameters.contract.Call(opts, &out, "decayRate")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DecayRate is a free data retrieval call binding the contract method 0xa9c1f2f1.
//
// Solidity: function decayRate() view returns(uint16)
func (_TicketingParameters *TicketingParametersSession) DecayRate() (uint16, error) {
	return _TicketingParameters.Contract.DecayRate(&_TicketingParameters.CallOpts)
}

// DecayRate is a free data retrieval call binding the contract method 0xa9c1f2f1.
//
// Solidity: function decayRate() view returns(uint16)
func (_TicketingParameters *TicketingParametersCallerSession) DecayRate() (uint16, error) {
	return _TicketingParameters.Contract.DecayRate(&_TicketingParameters.CallOpts)
}

// ExpiredWinProb is a free data retrieval call binding the contract method 0xbcbee543.
//
// Solidity: function expiredWinProb() view returns(uint128)
func (_TicketingParameters *TicketingParametersCaller) ExpiredWinProb(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TicketingParameters.contract.Call(opts, &out, "expiredWinProb")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpiredWinProb is a free data retrieval call binding the contract method 0xbcbee543.
//
// Solidity: function expiredWinProb() view returns(uint128)
func (_TicketingParameters *TicketingParametersSession) ExpiredWinProb() (*big.Int, error) {
	return _TicketingParameters.Contract.ExpiredWinProb(&_TicketingParameters.CallOpts)
}

// ExpiredWinProb is a free data retrieval call binding the contract method 0xbcbee543.
//
// Solidity: function expiredWinProb() view returns(uint128)
func (_TicketingParameters *TicketingParametersCallerSession) ExpiredWinProb() (*big.Int, error) {
	return _TicketingParameters.Contract.ExpiredWinProb(&_TicketingParameters.CallOpts)
}

// FaceValue is a free data retrieval call binding the contract method 0x44fd9caa.
//
// Solidity: function faceValue() view returns(uint256)
func (_TicketingParameters *TicketingParametersCaller) FaceValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TicketingParameters.contract.Call(opts, &out, "faceValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FaceValue is a free data retrieval call binding the contract method 0x44fd9caa.
//
// Solidity: function faceValue() view returns(uint256)
func (_TicketingParameters *TicketingParametersSession) FaceValue() (*big.Int, error) {
	return _TicketingParameters.Contract.FaceValue(&_TicketingParameters.CallOpts)
}

// FaceValue is a free data retrieval call binding the contract method 0x44fd9caa.
//
// Solidity: function faceValue() view returns(uint256)
func (_TicketingParameters *TicketingParametersCallerSession) FaceValue() (*big.Int, error) {
	return _TicketingParameters.Contract.FaceValue(&_TicketingParameters.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TicketingParameters *TicketingParametersCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TicketingParameters.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TicketingParameters *TicketingParametersSession) Owner() (common.Address, error) {
	return _TicketingParameters.Contract.Owner(&_TicketingParameters.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TicketingParameters *TicketingParametersCallerSession) Owner() (common.Address, error) {
	return _TicketingParameters.Contract.Owner(&_TicketingParameters.CallOpts)
}

// TicketDuration is a free data retrieval call binding the contract method 0x87bcc0c5.
//
// Solidity: function ticketDuration() view returns(uint256)
func (_TicketingParameters *TicketingParametersCaller) TicketDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TicketingParameters.contract.Call(opts, &out, "ticketDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TicketDuration is a free data retrieval call binding the contract method 0x87bcc0c5.
//
// Solidity: function ticketDuration() view returns(uint256)
func (_TicketingParameters *TicketingParametersSession) TicketDuration() (*big.Int, error) {
	return _TicketingParameters.Contract.TicketDuration(&_TicketingParameters.CallOpts)
}

// TicketDuration is a free data retrieval call binding the contract method 0x87bcc0c5.
//
// Solidity: function ticketDuration() view returns(uint256)
func (_TicketingParameters *TicketingParametersCallerSession) TicketDuration() (*big.Int, error) {
	return _TicketingParameters.Contract.TicketDuration(&_TicketingParameters.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x58cbd4d3.
//
// Solidity: function initialize(uint256 _faceValue, uint128 _baseLiveWinProb, uint128 _expiredWinProb, uint16 _decayRate, uint256 _ticketDuration) returns()
func (_TicketingParameters *TicketingParametersTransactor) Initialize(opts *bind.TransactOpts, _faceValue *big.Int, _baseLiveWinProb *big.Int, _expiredWinProb *big.Int, _decayRate uint16, _ticketDuration *big.Int) (*types.Transaction, error) {
	return _TicketingParameters.contract.Transact(opts, "initialize", _faceValue, _baseLiveWinProb, _expiredWinProb, _decayRate, _ticketDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0x58cbd4d3.
//
// Solidity: function initialize(uint256 _faceValue, uint128 _baseLiveWinProb, uint128 _expiredWinProb, uint16 _decayRate, uint256 _ticketDuration) returns()
func (_TicketingParameters *TicketingParametersSession) Initialize(_faceValue *big.Int, _baseLiveWinProb *big.Int, _expiredWinProb *big.Int, _decayRate uint16, _ticketDuration *big.Int) (*types.Transaction, error) {
	return _TicketingParameters.Contract.Initialize(&_TicketingParameters.TransactOpts, _faceValue, _baseLiveWinProb, _expiredWinProb, _decayRate, _ticketDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0x58cbd4d3.
//
// Solidity: function initialize(uint256 _faceValue, uint128 _baseLiveWinProb, uint128 _expiredWinProb, uint16 _decayRate, uint256 _ticketDuration) returns()
func (_TicketingParameters *TicketingParametersTransactorSession) Initialize(_faceValue *big.Int, _baseLiveWinProb *big.Int, _expiredWinProb *big.Int, _decayRate uint16, _ticketDuration *big.Int) (*types.Transaction, error) {
	return _TicketingParameters.Contract.Initialize(&_TicketingParameters.TransactOpts, _faceValue, _baseLiveWinProb, _expiredWinProb, _decayRate, _ticketDuration)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TicketingParameters *TicketingParametersTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TicketingParameters.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TicketingParameters *TicketingParametersSession) RenounceOwnership() (*types.Transaction, error) {
	return _TicketingParameters.Contract.RenounceOwnership(&_TicketingParameters.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TicketingParameters *TicketingParametersTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TicketingParameters.Contract.RenounceOwnership(&_TicketingParameters.TransactOpts)
}

// SetBaseLiveWinProb is a paid mutator transaction binding the contract method 0xa90a6027.
//
// Solidity: function setBaseLiveWinProb(uint128 _baseLiveWinProb) returns()
func (_TicketingParameters *TicketingParametersTransactor) SetBaseLiveWinProb(opts *bind.TransactOpts, _baseLiveWinProb *big.Int) (*types.Transaction, error) {
	return _TicketingParameters.contract.Transact(opts, "setBaseLiveWinProb", _baseLiveWinProb)
}

// SetBaseLiveWinProb is a paid mutator transaction binding the contract method 0xa90a6027.
//
// Solidity: function setBaseLiveWinProb(uint128 _baseLiveWinProb) returns()
func (_TicketingParameters *TicketingParametersSession) SetBaseLiveWinProb(_baseLiveWinProb *big.Int) (*types.Transaction, error) {
	return _TicketingParameters.Contract.SetBaseLiveWinProb(&_TicketingParameters.TransactOpts, _baseLiveWinProb)
}

// SetBaseLiveWinProb is a paid mutator transaction binding the contract method 0xa90a6027.
//
// Solidity: function setBaseLiveWinProb(uint128 _baseLiveWinProb) returns()
func (_TicketingParameters *TicketingParametersTransactorSession) SetBaseLiveWinProb(_baseLiveWinProb *big.Int) (*types.Transaction, error) {
	return _TicketingParameters.Contract.SetBaseLiveWinProb(&_TicketingParameters.TransactOpts, _baseLiveWinProb)
}

// SetExpiredWinProb is a paid mutator transaction binding the contract method 0xa8f19c14.
//
// Solidity: function setExpiredWinProb(uint128 _expiredWinProb) returns()
func (_TicketingParameters *TicketingParametersTransactor) SetExpiredWinProb(opts *bind.TransactOpts, _expiredWinProb *big.Int) (*types.Transaction, error) {
	return _TicketingParameters.contract.Transact(opts, "setExpiredWinProb", _expiredWinProb)
}

// SetExpiredWinProb is a paid mutator transaction binding the contract method 0xa8f19c14.
//
// Solidity: function setExpiredWinProb(uint128 _expiredWinProb) returns()
func (_TicketingParameters *TicketingParametersSession) SetExpiredWinProb(_expiredWinProb *big.Int) (*types.Transaction, error) {
	return _TicketingParameters.Contract.SetExpiredWinProb(&_TicketingParameters.TransactOpts, _expiredWinProb)
}

// SetExpiredWinProb is a paid mutator transaction binding the contract method 0xa8f19c14.
//
// Solidity: function setExpiredWinProb(uint128 _expiredWinProb) returns()
func (_TicketingParameters *TicketingParametersTransactorSession) SetExpiredWinProb(_expiredWinProb *big.Int) (*types.Transaction, error) {
	return _TicketingParameters.Contract.SetExpiredWinProb(&_TicketingParameters.TransactOpts, _expiredWinProb)
}

// SetFaceValue is a paid mutator transaction binding the contract method 0xef8032ef.
//
// Solidity: function setFaceValue(uint256 _faceValue) returns()
func (_TicketingParameters *TicketingParametersTransactor) SetFaceValue(opts *bind.TransactOpts, _faceValue *big.Int) (*types.Transaction, error) {
	return _TicketingParameters.contract.Transact(opts, "setFaceValue", _faceValue)
}

// SetFaceValue is a paid mutator transaction binding the contract method 0xef8032ef.
//
// Solidity: function setFaceValue(uint256 _faceValue) returns()
func (_TicketingParameters *TicketingParametersSession) SetFaceValue(_faceValue *big.Int) (*types.Transaction, error) {
	return _TicketingParameters.Contract.SetFaceValue(&_TicketingParameters.TransactOpts, _faceValue)
}

// SetFaceValue is a paid mutator transaction binding the contract method 0xef8032ef.
//
// Solidity: function setFaceValue(uint256 _faceValue) returns()
func (_TicketingParameters *TicketingParametersTransactorSession) SetFaceValue(_faceValue *big.Int) (*types.Transaction, error) {
	return _TicketingParameters.Contract.SetFaceValue(&_TicketingParameters.TransactOpts, _faceValue)
}

// SetTicketDuration is a paid mutator transaction binding the contract method 0x9e9ceeca.
//
// Solidity: function setTicketDuration(uint256 _ticketDuration) returns()
func (_TicketingParameters *TicketingParametersTransactor) SetTicketDuration(opts *bind.TransactOpts, _ticketDuration *big.Int) (*types.Transaction, error) {
	return _TicketingParameters.contract.Transact(opts, "setTicketDuration", _ticketDuration)
}

// SetTicketDuration is a paid mutator transaction binding the contract method 0x9e9ceeca.
//
// Solidity: function setTicketDuration(uint256 _ticketDuration) returns()
func (_TicketingParameters *TicketingParametersSession) SetTicketDuration(_ticketDuration *big.Int) (*types.Transaction, error) {
	return _TicketingParameters.Contract.SetTicketDuration(&_TicketingParameters.TransactOpts, _ticketDuration)
}

// SetTicketDuration is a paid mutator transaction binding the contract method 0x9e9ceeca.
//
// Solidity: function setTicketDuration(uint256 _ticketDuration) returns()
func (_TicketingParameters *TicketingParametersTransactorSession) SetTicketDuration(_ticketDuration *big.Int) (*types.Transaction, error) {
	return _TicketingParameters.Contract.SetTicketDuration(&_TicketingParameters.TransactOpts, _ticketDuration)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TicketingParameters *TicketingParametersTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TicketingParameters.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TicketingParameters *TicketingParametersSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TicketingParameters.Contract.TransferOwnership(&_TicketingParameters.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TicketingParameters *TicketingParametersTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TicketingParameters.Contract.TransferOwnership(&_TicketingParameters.TransactOpts, newOwner)
}

// TicketingParametersOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TicketingParameters contract.
type TicketingParametersOwnershipTransferredIterator struct {
	Event *TicketingParametersOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TicketingParametersOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TicketingParametersOwnershipTransferred)
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
		it.Event = new(TicketingParametersOwnershipTransferred)
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
func (it *TicketingParametersOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TicketingParametersOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TicketingParametersOwnershipTransferred represents a OwnershipTransferred event raised by the TicketingParameters contract.
type TicketingParametersOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TicketingParameters *TicketingParametersFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TicketingParametersOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TicketingParameters.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TicketingParametersOwnershipTransferredIterator{contract: _TicketingParameters.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TicketingParameters *TicketingParametersFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TicketingParametersOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TicketingParameters.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TicketingParametersOwnershipTransferred)
				if err := _TicketingParameters.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TicketingParameters *TicketingParametersFilterer) ParseOwnershipTransferred(log types.Log) (*TicketingParametersOwnershipTransferred, error) {
	event := new(TicketingParametersOwnershipTransferred)
	if err := _TicketingParameters.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
