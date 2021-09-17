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

// EpochsManagerEpoch is an auto generated low-level Go binding around an user-defined struct.
type EpochsManagerEpoch struct {
	Iteration               *big.Int
	StartBlock              *big.Int
	Duration                *big.Int
	EndBlock                *big.Int
	DefaultPayoutPercentage uint16
	FaceValue               *big.Int
	BaseLiveWinProb         *big.Int
	ExpiredWinProb          *big.Int
	TicketDuration          *big.Int
	DecayRate               uint16
}

// EpochsManagerABI is the input ABI used to generate the binding from.
const EpochsManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"}],\"name\":\"NewEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"currentActiveEpoch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDirectory\",\"name\":\"directory\",\"type\":\"address\"},{\"internalType\":\"contractListings\",\"name\":\"listings\",\"type\":\"address\"},{\"internalType\":\"contractTicketingParameters\",\"name\":\"ticketingParameters\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_epochDuration\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeEpoch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentActiveEpoch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"iteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"defaultPayoutPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"baseLiveWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"expiredWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"ticketDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"decayRate\",\"type\":\"uint16\"}],\"internalType\":\"structEpochsManager.Epoch\",\"name\":\"epoch\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"iteration\",\"type\":\"uint256\"}],\"name\":\"getEpochId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextEpochId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"}],\"name\":\"getEpoch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"iteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"defaultPayoutPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"baseLiveWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"expiredWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"ticketDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"decayRate\",\"type\":\"uint16\"}],\"internalType\":\"structEpochsManager.Epoch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EpochsManagerBin is the compiled bytecode used for deploying new contracts.
var EpochsManagerBin = "0x608060405234801561001057600080fd5b506110e1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c8063ae4a4d5711610081578063e1519a751161005b578063e1519a7514610166578063e34bf0131461016e578063f2fde38b1461017657600080fd5b8063ae4a4d5714610142578063b3e123db1461014b578063cf756fdf1461015357600080fd5b8063715018a6116100b2578063715018a6146100fd5780637e6d64a5146101075780638da5cb5b1461012757600080fd5b80634ff0876a146100ce5780635303548b146100ea575b600080fd5b6100d760695481565b6040519081526020015b60405180910390f35b6100d76100f8366004610eee565b610189565b6101056101bb565b005b61011a610115366004610eee565b610226565b6040516100e19190610fb7565b6033546040516001600160a01b0390911681526020016100e1565b6100d7606b5481565b6100d761033f565b610105610161366004610f06565b610358565b6100d7610471565b61011a610a49565b610105610184366004610ecb565b610b65565b60008160405160200161019e91815260200190565b604051602081830303815290604052805190602001209050919050565b6033546001600160a01b0316331461021a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6102246000610c47565b565b61029660405180610140016040528060008152602001600081526020016000815260200160008152602001600061ffff1681526020016000815260200160006001600160801b0316815260200160006001600160801b0316815260200160008152602001600061ffff1681525090565b506000908152606c60209081526040918290208251610140810184528154815260018201549281019290925260028101549282019290925260038201546060820152600482015461ffff9081166080830152600583015460a083015260068301546001600160801b0380821660c08501527001000000000000000000000000000000009091041660e0830152600783015461010083015260089092015490911661012082015290565b6000610353606a5460016100f89190611059565b905090565b600054610100900460ff1680610371575060005460ff16155b6103d45760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610211565b600054610100900460ff161580156103f6576000805461ffff19166101011790555b6103fe610cb1565b606680546001600160a01b038088167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556067805487841690831617905560688054928616929091169190911790556069829055801561046a576000805461ff00191690555b5050505050565b606b546000908152606c6020526040812060028101546001820154839161049791611059565b9050438111156104e95760405162461bcd60e51b815260206004820152601f60248201527f43757272656e742065706f636820686173206e6f742079657420656e646564006044820152606401610211565b6000606a5460016104fa9190611059565b90506000604051806101400160405280838152602001438152602001606954815260200160008152602001606760009054906101000a90046001600160a01b03166001600160a01b031663d2a78d7f6040518163ffffffff1660e01b815260040160206040518083038186803b15801561057357600080fd5b505afa158015610587573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105ab9190610f7d565b61ffff168152606854604080517f44fd9caa00000000000000000000000000000000000000000000000000000000815290516020938401936001600160a01b03909316926344fd9caa9260048082019391829003018186803b15801561061057600080fd5b505afa158015610624573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106489190610f9f565b8152606854604080517fdedcebda00000000000000000000000000000000000000000000000000000000815290516020938401936001600160a01b039093169263dedcebda9260048082019391829003018186803b1580156106a957600080fd5b505afa1580156106bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106e19190610f56565b6001600160801b03168152606854604080517fbcbee54300000000000000000000000000000000000000000000000000000000815290516020938401936001600160a01b039093169263bcbee5439260048082019391829003018186803b15801561074b57600080fd5b505afa15801561075f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107839190610f56565b6001600160801b03168152606854604080517f87bcc0c500000000000000000000000000000000000000000000000000000000815290516020938401936001600160a01b03909316926387bcc0c59260048082019391829003018186803b1580156107ed57600080fd5b505afa158015610801573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108259190610f9f565b8152606854604080517fa9c1f2f100000000000000000000000000000000000000000000000000000000815290516020938401936001600160a01b039093169263a9c1f2f19260048082019391829003018186803b15801561088657600080fd5b505afa15801561089a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108be9190610f7d565b61ffff169052905060006108d061033f565b6066546040517fc47534d4000000000000000000000000000000000000000000000000000000008152600481018390529192506001600160a01b03169063c47534d490602401600060405180830381600087803b15801561093057600080fd5b505af1158015610944573d6000803e3d6000fd5b5050506000828152606c6020908152604091829020855181559085015160018201558185015160028201556060850151600380830191909155608086015160048301805461ffff92831661ffff199182161790915560a0880151600585015560c088015160e08901516001600160801b03908116700100000000000000000000000000000000029116176006850155610100880151600785015561012088015160089094018054949092169316929092179091554390880155606a859055606b839055517fddc860800a99149017c480ec51523bf4143b7215e78956ae5c31e5c568f5383a9150610a389083815260200190565b60405180910390a195945050505050565b610ab960405180610140016040528060008152602001600081526020016000815260200160008152602001600061ffff1681526020016000815260200160006001600160801b0316815260200160006001600160801b0316815260200160008152602001600061ffff1681525090565b50606b546000908152606c60209081526040918290208251610140810184528154815260018201549281019290925260028101549282019290925260038201546060820152600482015461ffff9081166080830152600583015460a083015260068301546001600160801b0380821660c08501527001000000000000000000000000000000009091041660e0830152600783015461010083015260089092015490911661012082015290565b6033546001600160a01b03163314610bbf5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610211565b6001600160a01b038116610c3b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610211565b610c4481610c47565b50565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff1680610cca575060005460ff16155b610d2d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610211565b600054610100900460ff16158015610d4f576000805461ffff19166101011790555b610d57610d73565b610d5f610e24565b8015610c44576000805461ff001916905550565b600054610100900460ff1680610d8c575060005460ff16155b610def5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610211565b600054610100900460ff16158015610d5f576000805461ffff19166101011790558015610c44576000805461ff001916905550565b600054610100900460ff1680610e3d575060005460ff16155b610ea05760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610211565b600054610100900460ff16158015610ec2576000805461ffff19166101011790555b610d5f33610c47565b600060208284031215610edc578081fd5b8135610ee781611096565b9392505050565b600060208284031215610eff578081fd5b5035919050565b60008060008060808587031215610f1b578283fd5b8435610f2681611096565b93506020850135610f3681611096565b92506040850135610f4681611096565b9396929550929360600135925050565b600060208284031215610f67578081fd5b81516001600160801b0381168114610ee7578182fd5b600060208284031215610f8e578081fd5b815161ffff81168114610ee7578182fd5b600060208284031215610fb0578081fd5b5051919050565b600061014082019050825182526020830151602083015260408301516040830152606083015160608301526080830151610ff7608084018261ffff169052565b5060a083015160a083015260c083015161101c60c08401826001600160801b03169052565b5060e083015161103760e08401826001600160801b03169052565b5061010083810151908301526101209283015161ffff16929091019190915290565b60008219821115611091577f4e487b710000000000000000000000000000000000000000000000000000000081526011600452602481fd5b500190565b6001600160a01b0381168114610c4457600080fdfea26469706673582212202f512b843a83f46fdb5d8cf80cdfcbcaa53cc74c899e945fb932ecb3c89a669b64736f6c63430008040033"

// DeployEpochsManager deploys a new Ethereum contract, binding an instance of EpochsManager to it.
func DeployEpochsManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EpochsManager, error) {
	parsed, err := abi.JSON(strings.NewReader(EpochsManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EpochsManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EpochsManager{EpochsManagerCaller: EpochsManagerCaller{contract: contract}, EpochsManagerTransactor: EpochsManagerTransactor{contract: contract}, EpochsManagerFilterer: EpochsManagerFilterer{contract: contract}}, nil
}

// EpochsManager is an auto generated Go binding around an Ethereum contract.
type EpochsManager struct {
	EpochsManagerCaller     // Read-only binding to the contract
	EpochsManagerTransactor // Write-only binding to the contract
	EpochsManagerFilterer   // Log filterer for contract events
}

// EpochsManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EpochsManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EpochsManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EpochsManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EpochsManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EpochsManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EpochsManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EpochsManagerSession struct {
	Contract     *EpochsManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EpochsManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EpochsManagerCallerSession struct {
	Contract *EpochsManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EpochsManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EpochsManagerTransactorSession struct {
	Contract     *EpochsManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EpochsManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EpochsManagerRaw struct {
	Contract *EpochsManager // Generic contract binding to access the raw methods on
}

// EpochsManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EpochsManagerCallerRaw struct {
	Contract *EpochsManagerCaller // Generic read-only contract binding to access the raw methods on
}

// EpochsManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EpochsManagerTransactorRaw struct {
	Contract *EpochsManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEpochsManager creates a new instance of EpochsManager, bound to a specific deployed contract.
func NewEpochsManager(address common.Address, backend bind.ContractBackend) (*EpochsManager, error) {
	contract, err := bindEpochsManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EpochsManager{EpochsManagerCaller: EpochsManagerCaller{contract: contract}, EpochsManagerTransactor: EpochsManagerTransactor{contract: contract}, EpochsManagerFilterer: EpochsManagerFilterer{contract: contract}}, nil
}

// NewEpochsManagerCaller creates a new read-only instance of EpochsManager, bound to a specific deployed contract.
func NewEpochsManagerCaller(address common.Address, caller bind.ContractCaller) (*EpochsManagerCaller, error) {
	contract, err := bindEpochsManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EpochsManagerCaller{contract: contract}, nil
}

// NewEpochsManagerTransactor creates a new write-only instance of EpochsManager, bound to a specific deployed contract.
func NewEpochsManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*EpochsManagerTransactor, error) {
	contract, err := bindEpochsManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EpochsManagerTransactor{contract: contract}, nil
}

// NewEpochsManagerFilterer creates a new log filterer instance of EpochsManager, bound to a specific deployed contract.
func NewEpochsManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*EpochsManagerFilterer, error) {
	contract, err := bindEpochsManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EpochsManagerFilterer{contract: contract}, nil
}

// bindEpochsManager binds a generic wrapper to an already deployed contract.
func bindEpochsManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EpochsManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EpochsManager *EpochsManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EpochsManager.Contract.EpochsManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EpochsManager *EpochsManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EpochsManager.Contract.EpochsManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EpochsManager *EpochsManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EpochsManager.Contract.EpochsManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EpochsManager *EpochsManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EpochsManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EpochsManager *EpochsManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EpochsManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EpochsManager *EpochsManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EpochsManager.Contract.contract.Transact(opts, method, params...)
}

// CurrentActiveEpoch is a free data retrieval call binding the contract method 0xae4a4d57.
//
// Solidity: function currentActiveEpoch() view returns(bytes32)
func (_EpochsManager *EpochsManagerCaller) CurrentActiveEpoch(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EpochsManager.contract.Call(opts, &out, "currentActiveEpoch")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentActiveEpoch is a free data retrieval call binding the contract method 0xae4a4d57.
//
// Solidity: function currentActiveEpoch() view returns(bytes32)
func (_EpochsManager *EpochsManagerSession) CurrentActiveEpoch() ([32]byte, error) {
	return _EpochsManager.Contract.CurrentActiveEpoch(&_EpochsManager.CallOpts)
}

// CurrentActiveEpoch is a free data retrieval call binding the contract method 0xae4a4d57.
//
// Solidity: function currentActiveEpoch() view returns(bytes32)
func (_EpochsManager *EpochsManagerCallerSession) CurrentActiveEpoch() ([32]byte, error) {
	return _EpochsManager.Contract.CurrentActiveEpoch(&_EpochsManager.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_EpochsManager *EpochsManagerCaller) EpochDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EpochsManager.contract.Call(opts, &out, "epochDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_EpochsManager *EpochsManagerSession) EpochDuration() (*big.Int, error) {
	return _EpochsManager.Contract.EpochDuration(&_EpochsManager.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_EpochsManager *EpochsManagerCallerSession) EpochDuration() (*big.Int, error) {
	return _EpochsManager.Contract.EpochDuration(&_EpochsManager.CallOpts)
}

// GetCurrentActiveEpoch is a free data retrieval call binding the contract method 0xe34bf013.
//
// Solidity: function getCurrentActiveEpoch() view returns((uint256,uint256,uint256,uint256,uint16,uint256,uint128,uint128,uint256,uint16) epoch)
func (_EpochsManager *EpochsManagerCaller) GetCurrentActiveEpoch(opts *bind.CallOpts) (EpochsManagerEpoch, error) {
	var out []interface{}
	err := _EpochsManager.contract.Call(opts, &out, "getCurrentActiveEpoch")

	if err != nil {
		return *new(EpochsManagerEpoch), err
	}

	out0 := *abi.ConvertType(out[0], new(EpochsManagerEpoch)).(*EpochsManagerEpoch)

	return out0, err

}

// GetCurrentActiveEpoch is a free data retrieval call binding the contract method 0xe34bf013.
//
// Solidity: function getCurrentActiveEpoch() view returns((uint256,uint256,uint256,uint256,uint16,uint256,uint128,uint128,uint256,uint16) epoch)
func (_EpochsManager *EpochsManagerSession) GetCurrentActiveEpoch() (EpochsManagerEpoch, error) {
	return _EpochsManager.Contract.GetCurrentActiveEpoch(&_EpochsManager.CallOpts)
}

// GetCurrentActiveEpoch is a free data retrieval call binding the contract method 0xe34bf013.
//
// Solidity: function getCurrentActiveEpoch() view returns((uint256,uint256,uint256,uint256,uint16,uint256,uint128,uint128,uint256,uint16) epoch)
func (_EpochsManager *EpochsManagerCallerSession) GetCurrentActiveEpoch() (EpochsManagerEpoch, error) {
	return _EpochsManager.Contract.GetCurrentActiveEpoch(&_EpochsManager.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x7e6d64a5.
//
// Solidity: function getEpoch(bytes32 epochId) view returns((uint256,uint256,uint256,uint256,uint16,uint256,uint128,uint128,uint256,uint16))
func (_EpochsManager *EpochsManagerCaller) GetEpoch(opts *bind.CallOpts, epochId [32]byte) (EpochsManagerEpoch, error) {
	var out []interface{}
	err := _EpochsManager.contract.Call(opts, &out, "getEpoch", epochId)

	if err != nil {
		return *new(EpochsManagerEpoch), err
	}

	out0 := *abi.ConvertType(out[0], new(EpochsManagerEpoch)).(*EpochsManagerEpoch)

	return out0, err

}

// GetEpoch is a free data retrieval call binding the contract method 0x7e6d64a5.
//
// Solidity: function getEpoch(bytes32 epochId) view returns((uint256,uint256,uint256,uint256,uint16,uint256,uint128,uint128,uint256,uint16))
func (_EpochsManager *EpochsManagerSession) GetEpoch(epochId [32]byte) (EpochsManagerEpoch, error) {
	return _EpochsManager.Contract.GetEpoch(&_EpochsManager.CallOpts, epochId)
}

// GetEpoch is a free data retrieval call binding the contract method 0x7e6d64a5.
//
// Solidity: function getEpoch(bytes32 epochId) view returns((uint256,uint256,uint256,uint256,uint16,uint256,uint128,uint128,uint256,uint16))
func (_EpochsManager *EpochsManagerCallerSession) GetEpoch(epochId [32]byte) (EpochsManagerEpoch, error) {
	return _EpochsManager.Contract.GetEpoch(&_EpochsManager.CallOpts, epochId)
}

// GetEpochId is a free data retrieval call binding the contract method 0x5303548b.
//
// Solidity: function getEpochId(uint256 iteration) pure returns(bytes32)
func (_EpochsManager *EpochsManagerCaller) GetEpochId(opts *bind.CallOpts, iteration *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _EpochsManager.contract.Call(opts, &out, "getEpochId", iteration)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEpochId is a free data retrieval call binding the contract method 0x5303548b.
//
// Solidity: function getEpochId(uint256 iteration) pure returns(bytes32)
func (_EpochsManager *EpochsManagerSession) GetEpochId(iteration *big.Int) ([32]byte, error) {
	return _EpochsManager.Contract.GetEpochId(&_EpochsManager.CallOpts, iteration)
}

// GetEpochId is a free data retrieval call binding the contract method 0x5303548b.
//
// Solidity: function getEpochId(uint256 iteration) pure returns(bytes32)
func (_EpochsManager *EpochsManagerCallerSession) GetEpochId(iteration *big.Int) ([32]byte, error) {
	return _EpochsManager.Contract.GetEpochId(&_EpochsManager.CallOpts, iteration)
}

// GetNextEpochId is a free data retrieval call binding the contract method 0xb3e123db.
//
// Solidity: function getNextEpochId() view returns(bytes32)
func (_EpochsManager *EpochsManagerCaller) GetNextEpochId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EpochsManager.contract.Call(opts, &out, "getNextEpochId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNextEpochId is a free data retrieval call binding the contract method 0xb3e123db.
//
// Solidity: function getNextEpochId() view returns(bytes32)
func (_EpochsManager *EpochsManagerSession) GetNextEpochId() ([32]byte, error) {
	return _EpochsManager.Contract.GetNextEpochId(&_EpochsManager.CallOpts)
}

// GetNextEpochId is a free data retrieval call binding the contract method 0xb3e123db.
//
// Solidity: function getNextEpochId() view returns(bytes32)
func (_EpochsManager *EpochsManagerCallerSession) GetNextEpochId() ([32]byte, error) {
	return _EpochsManager.Contract.GetNextEpochId(&_EpochsManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EpochsManager *EpochsManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EpochsManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EpochsManager *EpochsManagerSession) Owner() (common.Address, error) {
	return _EpochsManager.Contract.Owner(&_EpochsManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EpochsManager *EpochsManagerCallerSession) Owner() (common.Address, error) {
	return _EpochsManager.Contract.Owner(&_EpochsManager.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address directory, address listings, address ticketingParameters, uint256 _epochDuration) returns()
func (_EpochsManager *EpochsManagerTransactor) Initialize(opts *bind.TransactOpts, directory common.Address, listings common.Address, ticketingParameters common.Address, _epochDuration *big.Int) (*types.Transaction, error) {
	return _EpochsManager.contract.Transact(opts, "initialize", directory, listings, ticketingParameters, _epochDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address directory, address listings, address ticketingParameters, uint256 _epochDuration) returns()
func (_EpochsManager *EpochsManagerSession) Initialize(directory common.Address, listings common.Address, ticketingParameters common.Address, _epochDuration *big.Int) (*types.Transaction, error) {
	return _EpochsManager.Contract.Initialize(&_EpochsManager.TransactOpts, directory, listings, ticketingParameters, _epochDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address directory, address listings, address ticketingParameters, uint256 _epochDuration) returns()
func (_EpochsManager *EpochsManagerTransactorSession) Initialize(directory common.Address, listings common.Address, ticketingParameters common.Address, _epochDuration *big.Int) (*types.Transaction, error) {
	return _EpochsManager.Contract.Initialize(&_EpochsManager.TransactOpts, directory, listings, ticketingParameters, _epochDuration)
}

// InitializeEpoch is a paid mutator transaction binding the contract method 0xe1519a75.
//
// Solidity: function initializeEpoch() returns(bytes32)
func (_EpochsManager *EpochsManagerTransactor) InitializeEpoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EpochsManager.contract.Transact(opts, "initializeEpoch")
}

// InitializeEpoch is a paid mutator transaction binding the contract method 0xe1519a75.
//
// Solidity: function initializeEpoch() returns(bytes32)
func (_EpochsManager *EpochsManagerSession) InitializeEpoch() (*types.Transaction, error) {
	return _EpochsManager.Contract.InitializeEpoch(&_EpochsManager.TransactOpts)
}

// InitializeEpoch is a paid mutator transaction binding the contract method 0xe1519a75.
//
// Solidity: function initializeEpoch() returns(bytes32)
func (_EpochsManager *EpochsManagerTransactorSession) InitializeEpoch() (*types.Transaction, error) {
	return _EpochsManager.Contract.InitializeEpoch(&_EpochsManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EpochsManager *EpochsManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EpochsManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EpochsManager *EpochsManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _EpochsManager.Contract.RenounceOwnership(&_EpochsManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EpochsManager *EpochsManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EpochsManager.Contract.RenounceOwnership(&_EpochsManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EpochsManager *EpochsManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EpochsManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EpochsManager *EpochsManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EpochsManager.Contract.TransferOwnership(&_EpochsManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EpochsManager *EpochsManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EpochsManager.Contract.TransferOwnership(&_EpochsManager.TransactOpts, newOwner)
}

// EpochsManagerNewEpochIterator is returned from FilterNewEpoch and is used to iterate over the raw logs and unpacked data for NewEpoch events raised by the EpochsManager contract.
type EpochsManagerNewEpochIterator struct {
	Event *EpochsManagerNewEpoch // Event containing the contract specifics and raw log

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
func (it *EpochsManagerNewEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochsManagerNewEpoch)
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
		it.Event = new(EpochsManagerNewEpoch)
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
func (it *EpochsManagerNewEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochsManagerNewEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochsManagerNewEpoch represents a NewEpoch event raised by the EpochsManager contract.
type EpochsManagerNewEpoch struct {
	EpochId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewEpoch is a free log retrieval operation binding the contract event 0xddc860800a99149017c480ec51523bf4143b7215e78956ae5c31e5c568f5383a.
//
// Solidity: event NewEpoch(bytes32 epochId)
func (_EpochsManager *EpochsManagerFilterer) FilterNewEpoch(opts *bind.FilterOpts) (*EpochsManagerNewEpochIterator, error) {

	logs, sub, err := _EpochsManager.contract.FilterLogs(opts, "NewEpoch")
	if err != nil {
		return nil, err
	}
	return &EpochsManagerNewEpochIterator{contract: _EpochsManager.contract, event: "NewEpoch", logs: logs, sub: sub}, nil
}

// WatchNewEpoch is a free log subscription operation binding the contract event 0xddc860800a99149017c480ec51523bf4143b7215e78956ae5c31e5c568f5383a.
//
// Solidity: event NewEpoch(bytes32 epochId)
func (_EpochsManager *EpochsManagerFilterer) WatchNewEpoch(opts *bind.WatchOpts, sink chan<- *EpochsManagerNewEpoch) (event.Subscription, error) {

	logs, sub, err := _EpochsManager.contract.WatchLogs(opts, "NewEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochsManagerNewEpoch)
				if err := _EpochsManager.contract.UnpackLog(event, "NewEpoch", log); err != nil {
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

// ParseNewEpoch is a log parse operation binding the contract event 0xddc860800a99149017c480ec51523bf4143b7215e78956ae5c31e5c568f5383a.
//
// Solidity: event NewEpoch(bytes32 epochId)
func (_EpochsManager *EpochsManagerFilterer) ParseNewEpoch(log types.Log) (*EpochsManagerNewEpoch, error) {
	event := new(EpochsManagerNewEpoch)
	if err := _EpochsManager.contract.UnpackLog(event, "NewEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EpochsManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EpochsManager contract.
type EpochsManagerOwnershipTransferredIterator struct {
	Event *EpochsManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EpochsManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochsManagerOwnershipTransferred)
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
		it.Event = new(EpochsManagerOwnershipTransferred)
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
func (it *EpochsManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochsManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochsManagerOwnershipTransferred represents a OwnershipTransferred event raised by the EpochsManager contract.
type EpochsManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EpochsManager *EpochsManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EpochsManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EpochsManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EpochsManagerOwnershipTransferredIterator{contract: _EpochsManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EpochsManager *EpochsManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EpochsManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EpochsManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochsManagerOwnershipTransferred)
				if err := _EpochsManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EpochsManager *EpochsManagerFilterer) ParseOwnershipTransferred(log types.Log) (*EpochsManagerOwnershipTransferred, error) {
	event := new(EpochsManagerOwnershipTransferred)
	if err := _EpochsManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
