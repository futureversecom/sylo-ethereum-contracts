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
	StartBlock      *big.Int
	Duration        *big.Int
	EndBlock        *big.Int
	DirectoryId     [32]byte
	FaceValue       *big.Int
	BaseLiveWinProb *big.Int
	ExpiredWinProb  *big.Int
	TicketDuration  *big.Int
	DecayRate       uint16
}

// EpochsManagerABI is the input ABI used to generate the binding from.
const EpochsManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"}],\"name\":\"NewEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"currentActiveEpoch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDirectory\",\"name\":\"directory\",\"type\":\"address\"},{\"internalType\":\"contractTicketingParameters\",\"name\":\"ticketingParameters\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_epochDuration\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeEpoch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentActiveEpoch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"directoryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"baseLiveWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"expiredWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"ticketDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"decayRate\",\"type\":\"uint16\"}],\"internalType\":\"structEpochsManager.Epoch\",\"name\":\"epoch\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"directoryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"baseLiveWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"expiredWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"ticketDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"decayRate\",\"type\":\"uint16\"}],\"internalType\":\"structEpochsManager.Epoch\",\"name\":\"epoch\",\"type\":\"tuple\"}],\"name\":\"getEpochId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"}],\"name\":\"getEpoch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"directoryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"baseLiveWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"expiredWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"ticketDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"decayRate\",\"type\":\"uint16\"}],\"internalType\":\"structEpochsManager.Epoch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EpochsManagerBin is the compiled bytecode used for deploying new contracts.
var EpochsManagerBin = "0x608060405234801561001057600080fd5b5061110b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c8063ae4a4d5711610076578063e34bf0131161005b578063e34bf0131461022e578063edcbc4e61461031f578063f2fde38b1461033257600080fd5b8063ae4a4d571461021d578063e1519a751461022657600080fd5b8063715018a6116100a7578063715018a6146100f45780637e6d64a5146100fc5780638da5cb5b1461020257600080fd5b80631794bb3c146100c35780634ff0876a146100d8575b600080fd5b6100d66100d1366004610eaa565b610345565b005b6100e160685481565b6040519081526020015b60405180910390f35b6100d6610455565b6101f561010a366004610e7a565b6040805161012081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810191909152506000908152606a602090815260409182902082516101208101845281548152600182015492810192909252600281015492820192909252600382015460608201526004820154608082015260058201546fffffffffffffffffffffffffffffffff80821660a08401527001000000000000000000000000000000009091041660c0820152600682015460e082015260079091015461ffff1661010082015290565b6040516100eb9190610fad565b6033546040516001600160a01b0390911681526020016100eb565b6100e160695481565b6100e16104bb565b6101f56040805161012081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810191909152506069546000908152606a602090815260409182902082516101208101845281548152600182015492810192909252600281015492820192909252600382015460608201526004820154608082015260058201546fffffffffffffffffffffffffffffffff80821660a08401527001000000000000000000000000000000009091041660c0820152600682015460e082015260079091015461ffff1661010082015290565b6100e161032d366004610eea565b610a05565b6100d6610340366004610e57565b610ad6565b600054610100900460ff168061035e575060005460ff16155b6103c65760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b600054610100900460ff161580156103e8576000805461ffff19166101011790555b6103f0610bb8565b606680546001600160a01b038087167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560678054928616929091169190911790556068829055801561044f576000805461ff00191690555b50505050565b6033546001600160a01b031633146104af5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103bd565b6104b96000610c7a565b565b6069546000908152606a602052604081206001810154815483916104de9161106e565b9050438111156105305760405162461bcd60e51b815260206004820152601f60248201527f43757272656e742065706f636820686173206e6f742079657420656e6465640060448201526064016103bd565b606654604080517f859ea34700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b03169163859ea34791600480830192602092919082900301818787803b15801561058f57600080fd5b505af11580156105a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c79190610e92565b90506000604051806101200160405280438152602001606854815260200160008152602001838152602001606760009054906101000a90046001600160a01b03166001600160a01b03166344fd9caa6040518163ffffffff1660e01b815260040160206040518083038186803b15801561064057600080fd5b505afa158015610654573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106789190610e92565b8152606754604080517fdedcebda00000000000000000000000000000000000000000000000000000000815290516020938401936001600160a01b039093169263dedcebda9260048082019391829003018186803b1580156106d957600080fd5b505afa1580156106ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107119190610f75565b6fffffffffffffffffffffffffffffffff168152606754604080517fbcbee54300000000000000000000000000000000000000000000000000000000815290516020938401936001600160a01b039093169263bcbee5439260048082019391829003018186803b15801561078457600080fd5b505afa158015610798573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107bc9190610f75565b6fffffffffffffffffffffffffffffffff168152606754604080517f87bcc0c500000000000000000000000000000000000000000000000000000000815290516020938401936001600160a01b03909316926387bcc0c59260048082019391829003018186803b15801561082f57600080fd5b505afa158015610843573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108679190610e92565b8152606754604080517fa9c1f2f100000000000000000000000000000000000000000000000000000000815290516020938401936001600160a01b039093169263a9c1f2f19260048082019391829003018186803b1580156108c857600080fd5b505afa1580156108dc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109009190610f91565b61ffff1690529050600061091382610a05565b6000818152606a60209081526040918290208551815590850151600182015581850151600280830191909155606086015160038301556080860151600483015560a086015160c08701516fffffffffffffffffffffffffffffffff90811670010000000000000000000000000000000002911617600583015560e086015160068301556101008601516007909201805461ffff90931661ffff199093169290921790915543908801556069829055519091507fddc860800a99149017c480ec51523bf4143b7215e78956ae5c31e5c568f5383a906109f49083815260200190565b60405180910390a195945050505050565b60008160000151826020015183608001518460a001518560c001518660e00151876101000151604051602001610ab9979695949392919096875260208701959095526040860193909352608091821b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000908116606087015290821b16607085015283015260f01b7fffff0000000000000000000000000000000000000000000000000000000000001660a082015260a20190565b604051602081830303815290604052805190602001209050919050565b6033546001600160a01b03163314610b305760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103bd565b6001600160a01b038116610bac5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016103bd565b610bb581610c7a565b50565b600054610100900460ff1680610bd1575060005460ff16155b610c345760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103bd565b600054610100900460ff16158015610c56576000805461ffff19166101011790555b610c5e610ce4565b610c66610d95565b8015610bb5576000805461ff001916905550565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff1680610cfd575060005460ff16155b610d605760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103bd565b600054610100900460ff16158015610c66576000805461ffff19166101011790558015610bb5576000805461ff001916905550565b600054610100900460ff1680610dae575060005460ff16155b610e115760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103bd565b600054610100900460ff16158015610e33576000805461ffff19166101011790555b610c6633610c7a565b8035610e47816110a7565b919050565b8035610e47816110c5565b600060208284031215610e68578081fd5b8135610e7381611092565b9392505050565b600060208284031215610e8b578081fd5b5035919050565b600060208284031215610ea3578081fd5b5051919050565b600080600060608486031215610ebe578182fd5b8335610ec981611092565b92506020840135610ed981611092565b929592945050506040919091013590565b60006101208284031215610efc578081fd5b610f04611036565b8235815260208301356020820152604083013560408201526060830135606082015260808301356080820152610f3c60a08401610e3c565b60a0820152610f4d60c08401610e3c565b60c082015260e083013560e0820152610100610f6a818501610e4c565b908201529392505050565b600060208284031215610f86578081fd5b8151610e73816110a7565b600060208284031215610fa2578081fd5b8151610e73816110c5565b600061012082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015260a08301516fffffffffffffffffffffffffffffffff80821660a08501528060c08601511660c0850152505060e083015160e08301526101008084015161102e8285018261ffff169052565b505092915050565b604051610120810167ffffffffffffffff8111828210171561106857634e487b7160e01b600052604160045260246000fd5b60405290565b6000821982111561108d57634e487b7160e01b81526011600452602481fd5b500190565b6001600160a01b0381168114610bb557600080fd5b6fffffffffffffffffffffffffffffffff81168114610bb557600080fd5b61ffff81168114610bb557600080fdfea2646970667358221220847763ee74b53ecf7aa8ed99cd04a09927135571a5b90c2a20ff5e61b7bd347d64736f6c63430008040033"

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
// Solidity: function getCurrentActiveEpoch() view returns((uint256,uint256,uint256,bytes32,uint256,uint128,uint128,uint256,uint16) epoch)
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
// Solidity: function getCurrentActiveEpoch() view returns((uint256,uint256,uint256,bytes32,uint256,uint128,uint128,uint256,uint16) epoch)
func (_EpochsManager *EpochsManagerSession) GetCurrentActiveEpoch() (EpochsManagerEpoch, error) {
	return _EpochsManager.Contract.GetCurrentActiveEpoch(&_EpochsManager.CallOpts)
}

// GetCurrentActiveEpoch is a free data retrieval call binding the contract method 0xe34bf013.
//
// Solidity: function getCurrentActiveEpoch() view returns((uint256,uint256,uint256,bytes32,uint256,uint128,uint128,uint256,uint16) epoch)
func (_EpochsManager *EpochsManagerCallerSession) GetCurrentActiveEpoch() (EpochsManagerEpoch, error) {
	return _EpochsManager.Contract.GetCurrentActiveEpoch(&_EpochsManager.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x7e6d64a5.
//
// Solidity: function getEpoch(bytes32 epochId) view returns((uint256,uint256,uint256,bytes32,uint256,uint128,uint128,uint256,uint16))
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
// Solidity: function getEpoch(bytes32 epochId) view returns((uint256,uint256,uint256,bytes32,uint256,uint128,uint128,uint256,uint16))
func (_EpochsManager *EpochsManagerSession) GetEpoch(epochId [32]byte) (EpochsManagerEpoch, error) {
	return _EpochsManager.Contract.GetEpoch(&_EpochsManager.CallOpts, epochId)
}

// GetEpoch is a free data retrieval call binding the contract method 0x7e6d64a5.
//
// Solidity: function getEpoch(bytes32 epochId) view returns((uint256,uint256,uint256,bytes32,uint256,uint128,uint128,uint256,uint16))
func (_EpochsManager *EpochsManagerCallerSession) GetEpoch(epochId [32]byte) (EpochsManagerEpoch, error) {
	return _EpochsManager.Contract.GetEpoch(&_EpochsManager.CallOpts, epochId)
}

// GetEpochId is a free data retrieval call binding the contract method 0xedcbc4e6.
//
// Solidity: function getEpochId((uint256,uint256,uint256,bytes32,uint256,uint128,uint128,uint256,uint16) epoch) pure returns(bytes32)
func (_EpochsManager *EpochsManagerCaller) GetEpochId(opts *bind.CallOpts, epoch EpochsManagerEpoch) ([32]byte, error) {
	var out []interface{}
	err := _EpochsManager.contract.Call(opts, &out, "getEpochId", epoch)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEpochId is a free data retrieval call binding the contract method 0xedcbc4e6.
//
// Solidity: function getEpochId((uint256,uint256,uint256,bytes32,uint256,uint128,uint128,uint256,uint16) epoch) pure returns(bytes32)
func (_EpochsManager *EpochsManagerSession) GetEpochId(epoch EpochsManagerEpoch) ([32]byte, error) {
	return _EpochsManager.Contract.GetEpochId(&_EpochsManager.CallOpts, epoch)
}

// GetEpochId is a free data retrieval call binding the contract method 0xedcbc4e6.
//
// Solidity: function getEpochId((uint256,uint256,uint256,bytes32,uint256,uint128,uint128,uint256,uint16) epoch) pure returns(bytes32)
func (_EpochsManager *EpochsManagerCallerSession) GetEpochId(epoch EpochsManagerEpoch) ([32]byte, error) {
	return _EpochsManager.Contract.GetEpochId(&_EpochsManager.CallOpts, epoch)
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

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address directory, address ticketingParameters, uint256 _epochDuration) returns()
func (_EpochsManager *EpochsManagerTransactor) Initialize(opts *bind.TransactOpts, directory common.Address, ticketingParameters common.Address, _epochDuration *big.Int) (*types.Transaction, error) {
	return _EpochsManager.contract.Transact(opts, "initialize", directory, ticketingParameters, _epochDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address directory, address ticketingParameters, uint256 _epochDuration) returns()
func (_EpochsManager *EpochsManagerSession) Initialize(directory common.Address, ticketingParameters common.Address, _epochDuration *big.Int) (*types.Transaction, error) {
	return _EpochsManager.Contract.Initialize(&_EpochsManager.TransactOpts, directory, ticketingParameters, _epochDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address directory, address ticketingParameters, uint256 _epochDuration) returns()
func (_EpochsManager *EpochsManagerTransactorSession) Initialize(directory common.Address, ticketingParameters common.Address, _epochDuration *big.Int) (*types.Transaction, error) {
	return _EpochsManager.Contract.Initialize(&_EpochsManager.TransactOpts, directory, ticketingParameters, _epochDuration)
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
