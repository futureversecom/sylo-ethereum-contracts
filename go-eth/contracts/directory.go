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

// DirectoryABI is the input ABI used to generate the binding from.
const DirectoryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"currentDirectory\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractStakingManager\",\"name\":\"stakingManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"}],\"name\":\"setCurrentDirectory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"}],\"name\":\"joinDirectory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"point\",\"type\":\"uint128\"}],\"name\":\"scan\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getTotalStakeForStakee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"}],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"}],\"name\":\"getEntries\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DirectoryBin is the compiled bytecode used for deploying new contracts.
var DirectoryBin = "0x608060405234801561001057600080fd5b50610f3d806100206000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c8063a019fef811610081578063f27a7d4b1161005b578063f27a7d4b146101bd578063f2fde38b146101de578063fa42e7ea146101f157600080fd5b8063a019fef814610184578063c47534d414610197578063c4d66de8146101aa57600080fd5b80634e680661116100b25780634e6806611461013e578063715018a6146101695780638da5cb5b1461017357600080fd5b80630e332e82146100ce5780633234fef51461011b575b600080fd5b6101086100dc366004610d59565b60008281526067602090815260408083206001600160a01b038516845260010190915290205492915050565b6040519081526020015b60405180910390f35b610108610129366004610d41565b60009081526067602052604090206002015490565b61015161014c366004610d88565b6101fa565b6040516001600160a01b039091168152602001610112565b6101716103ea565b005b6033546001600160a01b0316610151565b610171610192366004610d41565b610455565b6101716101a5366004610d41565b6106a7565b6101716101b8366004610d1e565b610706565b6101d06101cb366004610d41565b6107ea565b604051610112929190610dd0565b6101716101ec366004610d1e565b6109c3565b61010860665481565b60665460009081526067602052604081205461021857506000919050565b60665460009081526067602052604081206002015460809061024d906fffffffffffffffffffffffffffffffff861690610e8b565b6066546000908152606760205260408120549190921c9250819061027390600190610eaa565b90505b8082116103df576000600261028b8385610e53565b6102959190610e6b565b9050600081156102f05760665460009081526067602052604090206102bb600184610eaa565b815481106102d957634e487b7160e01b600052603260045260246000fd5b9060005260206000209060020201600101546102f3565b60005b6066546000908152606760205260408120805492935090918490811061032957634e487b7160e01b600052603260045260246000fd5b906000526020600020906002020160010154905081861015801561034c57508086105b156103a857606654600090815260676020526040902080548490811061038257634e487b7160e01b600052603260045260246000fd5b60009182526020909120600290910201546001600160a01b031698975050505050505050565b818610156103c2576103bb600184610eaa565b93506103d7565b8086106103d7576103d4836001610e53565b94505b505050610276565b506000949350505050565b6033546001600160a01b031633146104495760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6104536000610aa5565b565b6065546040517fdf349ed50000000000000000000000000000000000000000000000000000000081523360048201819052916000916001600160a01b039091169063df349ed59060240160206040518083038186803b1580156104b757600080fd5b505afa1580156104cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ef9190610db8565b9050600081116105675760405162461bcd60e51b815260206004820152603760248201527f43616e206e6f74206a6f696e206469726563746f727920666f72206e6578742060448201527f65706f636820776974686f757420616e79207374616b650000000000000000006064820152608401610440565b60008381526067602090815260408083206001600160a01b0386168452600101909152902054156106005760405162461bcd60e51b815260206004820152602a60248201527f43616e206f6e6c79206a6f696e20746865206469726563746f7279206f6e636560448201527f207065722065706f6368000000000000000000000000000000000000000000006064820152608401610440565b60008381526067602052604081206002015461061d908390610e53565b6000858152606760208181526040808420815180830183526001600160a01b03998a16808252818501888152835460018082018655858a52878a2094516002928302909501805473ffffffffffffffffffffffffffffffffffffffff191695909e16949094178d5590519b83019b909b558652810183529084209690965596909152909452500155565b6033546001600160a01b031633146107015760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610440565b606655565b600054610100900460ff168061071f575060005460ff16155b6107825760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610440565b600054610100900460ff161580156107a4576000805461ffff19166101011790555b6107ac610b04565b6065805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03841617905580156107e6576000805461ff00191690555b5050565b600081815260676020526040812054606091829167ffffffffffffffff81111561082457634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561084d578160200160208202803683370190505b506000858152606760205260408120549192509067ffffffffffffffff81111561088757634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156108b0578160200160208202803683370190505b50905060005b6000868152606760205260409020548110156109b85760008681526067602052604081208054839081106108fa57634e487b7160e01b600052603260045260246000fd5b60009182526020918290206040805180820190915260029092020180546001600160a01b031680835260019091015492820192909252855190925085908490811061095557634e487b7160e01b600052603260045260246000fd5b60200260200101906001600160a01b031690816001600160a01b031681525050806020015183838151811061099a57634e487b7160e01b600052603260045260246000fd5b602090810291909101015250806109b081610ec1565b9150506108b6565b509094909350915050565b6033546001600160a01b03163314610a1d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610440565b6001600160a01b038116610a995760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610440565b610aa281610aa5565b50565b603380546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff1680610b1d575060005460ff16155b610b805760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610440565b600054610100900460ff16158015610ba2576000805461ffff19166101011790555b610baa610bc6565b610bb2610c77565b8015610aa2576000805461ff001916905550565b600054610100900460ff1680610bdf575060005460ff16155b610c425760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610440565b600054610100900460ff16158015610bb2576000805461ffff19166101011790558015610aa2576000805461ff001916905550565b600054610100900460ff1680610c90575060005460ff16155b610cf35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610440565b600054610100900460ff16158015610d15576000805461ffff19166101011790555b610bb233610aa5565b600060208284031215610d2f578081fd5b8135610d3a81610ef2565b9392505050565b600060208284031215610d52578081fd5b5035919050565b60008060408385031215610d6b578081fd5b823591506020830135610d7d81610ef2565b809150509250929050565b600060208284031215610d99578081fd5b81356fffffffffffffffffffffffffffffffff81168114610d3a578182fd5b600060208284031215610dc9578081fd5b5051919050565b604080825283519082018190526000906020906060840190828701845b82811015610e125781516001600160a01b031684529284019290840190600101610ded565b50505083810382850152845180825285830191830190845b81811015610e4657835183529284019291840191600101610e2a565b5090979650505050505050565b60008219821115610e6657610e66610edc565b500190565b600082610e8657634e487b7160e01b81526012600452602481fd5b500490565b6000816000190483118215151615610ea557610ea5610edc565b500290565b600082821015610ebc57610ebc610edc565b500390565b6000600019821415610ed557610ed5610edc565b5060010190565b634e487b7160e01b600052601160045260246000fd5b6001600160a01b0381168114610aa257600080fdfea26469706673582212201e902b3654b908a1fe3057afca3cca11cddab19a9814db8802a4cd690e33415764736f6c63430008040033"

// DeployDirectory deploys a new Ethereum contract, binding an instance of Directory to it.
func DeployDirectory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Directory, error) {
	parsed, err := abi.JSON(strings.NewReader(DirectoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DirectoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Directory{DirectoryCaller: DirectoryCaller{contract: contract}, DirectoryTransactor: DirectoryTransactor{contract: contract}, DirectoryFilterer: DirectoryFilterer{contract: contract}}, nil
}

// Directory is an auto generated Go binding around an Ethereum contract.
type Directory struct {
	DirectoryCaller     // Read-only binding to the contract
	DirectoryTransactor // Write-only binding to the contract
	DirectoryFilterer   // Log filterer for contract events
}

// DirectoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DirectoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DirectoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DirectoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DirectoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DirectoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DirectorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DirectorySession struct {
	Contract     *Directory        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DirectoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DirectoryCallerSession struct {
	Contract *DirectoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DirectoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DirectoryTransactorSession struct {
	Contract     *DirectoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DirectoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DirectoryRaw struct {
	Contract *Directory // Generic contract binding to access the raw methods on
}

// DirectoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DirectoryCallerRaw struct {
	Contract *DirectoryCaller // Generic read-only contract binding to access the raw methods on
}

// DirectoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DirectoryTransactorRaw struct {
	Contract *DirectoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDirectory creates a new instance of Directory, bound to a specific deployed contract.
func NewDirectory(address common.Address, backend bind.ContractBackend) (*Directory, error) {
	contract, err := bindDirectory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Directory{DirectoryCaller: DirectoryCaller{contract: contract}, DirectoryTransactor: DirectoryTransactor{contract: contract}, DirectoryFilterer: DirectoryFilterer{contract: contract}}, nil
}

// NewDirectoryCaller creates a new read-only instance of Directory, bound to a specific deployed contract.
func NewDirectoryCaller(address common.Address, caller bind.ContractCaller) (*DirectoryCaller, error) {
	contract, err := bindDirectory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DirectoryCaller{contract: contract}, nil
}

// NewDirectoryTransactor creates a new write-only instance of Directory, bound to a specific deployed contract.
func NewDirectoryTransactor(address common.Address, transactor bind.ContractTransactor) (*DirectoryTransactor, error) {
	contract, err := bindDirectory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DirectoryTransactor{contract: contract}, nil
}

// NewDirectoryFilterer creates a new log filterer instance of Directory, bound to a specific deployed contract.
func NewDirectoryFilterer(address common.Address, filterer bind.ContractFilterer) (*DirectoryFilterer, error) {
	contract, err := bindDirectory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DirectoryFilterer{contract: contract}, nil
}

// bindDirectory binds a generic wrapper to an already deployed contract.
func bindDirectory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DirectoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Directory *DirectoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Directory.Contract.DirectoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Directory *DirectoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Directory.Contract.DirectoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Directory *DirectoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Directory.Contract.DirectoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Directory *DirectoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Directory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Directory *DirectoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Directory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Directory *DirectoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Directory.Contract.contract.Transact(opts, method, params...)
}

// CurrentDirectory is a free data retrieval call binding the contract method 0xfa42e7ea.
//
// Solidity: function currentDirectory() view returns(bytes32)
func (_Directory *DirectoryCaller) CurrentDirectory(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "currentDirectory")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentDirectory is a free data retrieval call binding the contract method 0xfa42e7ea.
//
// Solidity: function currentDirectory() view returns(bytes32)
func (_Directory *DirectorySession) CurrentDirectory() ([32]byte, error) {
	return _Directory.Contract.CurrentDirectory(&_Directory.CallOpts)
}

// CurrentDirectory is a free data retrieval call binding the contract method 0xfa42e7ea.
//
// Solidity: function currentDirectory() view returns(bytes32)
func (_Directory *DirectoryCallerSession) CurrentDirectory() ([32]byte, error) {
	return _Directory.Contract.CurrentDirectory(&_Directory.CallOpts)
}

// GetEntries is a free data retrieval call binding the contract method 0xf27a7d4b.
//
// Solidity: function getEntries(bytes32 epochId) view returns(address[], uint256[])
func (_Directory *DirectoryCaller) GetEntries(opts *bind.CallOpts, epochId [32]byte) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "getEntries", epochId)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetEntries is a free data retrieval call binding the contract method 0xf27a7d4b.
//
// Solidity: function getEntries(bytes32 epochId) view returns(address[], uint256[])
func (_Directory *DirectorySession) GetEntries(epochId [32]byte) ([]common.Address, []*big.Int, error) {
	return _Directory.Contract.GetEntries(&_Directory.CallOpts, epochId)
}

// GetEntries is a free data retrieval call binding the contract method 0xf27a7d4b.
//
// Solidity: function getEntries(bytes32 epochId) view returns(address[], uint256[])
func (_Directory *DirectoryCallerSession) GetEntries(epochId [32]byte) ([]common.Address, []*big.Int, error) {
	return _Directory.Contract.GetEntries(&_Directory.CallOpts, epochId)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x3234fef5.
//
// Solidity: function getTotalStake(bytes32 epochId) view returns(uint256)
func (_Directory *DirectoryCaller) GetTotalStake(opts *bind.CallOpts, epochId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "getTotalStake", epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStake is a free data retrieval call binding the contract method 0x3234fef5.
//
// Solidity: function getTotalStake(bytes32 epochId) view returns(uint256)
func (_Directory *DirectorySession) GetTotalStake(epochId [32]byte) (*big.Int, error) {
	return _Directory.Contract.GetTotalStake(&_Directory.CallOpts, epochId)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x3234fef5.
//
// Solidity: function getTotalStake(bytes32 epochId) view returns(uint256)
func (_Directory *DirectoryCallerSession) GetTotalStake(epochId [32]byte) (*big.Int, error) {
	return _Directory.Contract.GetTotalStake(&_Directory.CallOpts, epochId)
}

// GetTotalStakeForStakee is a free data retrieval call binding the contract method 0x0e332e82.
//
// Solidity: function getTotalStakeForStakee(bytes32 epochId, address stakee) view returns(uint256)
func (_Directory *DirectoryCaller) GetTotalStakeForStakee(opts *bind.CallOpts, epochId [32]byte, stakee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "getTotalStakeForStakee", epochId, stakee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStakeForStakee is a free data retrieval call binding the contract method 0x0e332e82.
//
// Solidity: function getTotalStakeForStakee(bytes32 epochId, address stakee) view returns(uint256)
func (_Directory *DirectorySession) GetTotalStakeForStakee(epochId [32]byte, stakee common.Address) (*big.Int, error) {
	return _Directory.Contract.GetTotalStakeForStakee(&_Directory.CallOpts, epochId, stakee)
}

// GetTotalStakeForStakee is a free data retrieval call binding the contract method 0x0e332e82.
//
// Solidity: function getTotalStakeForStakee(bytes32 epochId, address stakee) view returns(uint256)
func (_Directory *DirectoryCallerSession) GetTotalStakeForStakee(epochId [32]byte, stakee common.Address) (*big.Int, error) {
	return _Directory.Contract.GetTotalStakeForStakee(&_Directory.CallOpts, epochId, stakee)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Directory *DirectoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Directory *DirectorySession) Owner() (common.Address, error) {
	return _Directory.Contract.Owner(&_Directory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Directory *DirectoryCallerSession) Owner() (common.Address, error) {
	return _Directory.Contract.Owner(&_Directory.CallOpts)
}

// Scan is a free data retrieval call binding the contract method 0x4e680661.
//
// Solidity: function scan(uint128 point) view returns(address)
func (_Directory *DirectoryCaller) Scan(opts *bind.CallOpts, point *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "scan", point)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Scan is a free data retrieval call binding the contract method 0x4e680661.
//
// Solidity: function scan(uint128 point) view returns(address)
func (_Directory *DirectorySession) Scan(point *big.Int) (common.Address, error) {
	return _Directory.Contract.Scan(&_Directory.CallOpts, point)
}

// Scan is a free data retrieval call binding the contract method 0x4e680661.
//
// Solidity: function scan(uint128 point) view returns(address)
func (_Directory *DirectoryCallerSession) Scan(point *big.Int) (common.Address, error) {
	return _Directory.Contract.Scan(&_Directory.CallOpts, point)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address stakingManager) returns()
func (_Directory *DirectoryTransactor) Initialize(opts *bind.TransactOpts, stakingManager common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "initialize", stakingManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address stakingManager) returns()
func (_Directory *DirectorySession) Initialize(stakingManager common.Address) (*types.Transaction, error) {
	return _Directory.Contract.Initialize(&_Directory.TransactOpts, stakingManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address stakingManager) returns()
func (_Directory *DirectoryTransactorSession) Initialize(stakingManager common.Address) (*types.Transaction, error) {
	return _Directory.Contract.Initialize(&_Directory.TransactOpts, stakingManager)
}

// JoinDirectory is a paid mutator transaction binding the contract method 0xa019fef8.
//
// Solidity: function joinDirectory(bytes32 epochId) returns()
func (_Directory *DirectoryTransactor) JoinDirectory(opts *bind.TransactOpts, epochId [32]byte) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "joinDirectory", epochId)
}

// JoinDirectory is a paid mutator transaction binding the contract method 0xa019fef8.
//
// Solidity: function joinDirectory(bytes32 epochId) returns()
func (_Directory *DirectorySession) JoinDirectory(epochId [32]byte) (*types.Transaction, error) {
	return _Directory.Contract.JoinDirectory(&_Directory.TransactOpts, epochId)
}

// JoinDirectory is a paid mutator transaction binding the contract method 0xa019fef8.
//
// Solidity: function joinDirectory(bytes32 epochId) returns()
func (_Directory *DirectoryTransactorSession) JoinDirectory(epochId [32]byte) (*types.Transaction, error) {
	return _Directory.Contract.JoinDirectory(&_Directory.TransactOpts, epochId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Directory *DirectoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Directory *DirectorySession) RenounceOwnership() (*types.Transaction, error) {
	return _Directory.Contract.RenounceOwnership(&_Directory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Directory *DirectoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Directory.Contract.RenounceOwnership(&_Directory.TransactOpts)
}

// SetCurrentDirectory is a paid mutator transaction binding the contract method 0xc47534d4.
//
// Solidity: function setCurrentDirectory(bytes32 epochId) returns()
func (_Directory *DirectoryTransactor) SetCurrentDirectory(opts *bind.TransactOpts, epochId [32]byte) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "setCurrentDirectory", epochId)
}

// SetCurrentDirectory is a paid mutator transaction binding the contract method 0xc47534d4.
//
// Solidity: function setCurrentDirectory(bytes32 epochId) returns()
func (_Directory *DirectorySession) SetCurrentDirectory(epochId [32]byte) (*types.Transaction, error) {
	return _Directory.Contract.SetCurrentDirectory(&_Directory.TransactOpts, epochId)
}

// SetCurrentDirectory is a paid mutator transaction binding the contract method 0xc47534d4.
//
// Solidity: function setCurrentDirectory(bytes32 epochId) returns()
func (_Directory *DirectoryTransactorSession) SetCurrentDirectory(epochId [32]byte) (*types.Transaction, error) {
	return _Directory.Contract.SetCurrentDirectory(&_Directory.TransactOpts, epochId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Directory *DirectoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Directory *DirectorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Directory.Contract.TransferOwnership(&_Directory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Directory *DirectoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Directory.Contract.TransferOwnership(&_Directory.TransactOpts, newOwner)
}

// DirectoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Directory contract.
type DirectoryOwnershipTransferredIterator struct {
	Event *DirectoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DirectoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectoryOwnershipTransferred)
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
		it.Event = new(DirectoryOwnershipTransferred)
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
func (it *DirectoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectoryOwnershipTransferred represents a OwnershipTransferred event raised by the Directory contract.
type DirectoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Directory *DirectoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DirectoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Directory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DirectoryOwnershipTransferredIterator{contract: _Directory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Directory *DirectoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DirectoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Directory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectoryOwnershipTransferred)
				if err := _Directory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Directory *DirectoryFilterer) ParseOwnershipTransferred(log types.Log) (*DirectoryOwnershipTransferred, error) {
	event := new(DirectoryOwnershipTransferred)
	if err := _Directory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
