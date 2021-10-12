// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package directory

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

// DirectoryMetaData contains all meta data concerning the Directory contract.
var DirectoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"currentDirectory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractStakingManager\",\"name\":\"stakingManager\",\"type\":\"address\"},{\"internalType\":\"contractRewardsManager\",\"name\":\"rewardsManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"setCurrentDirectory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joinNextDirectory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"point\",\"type\":\"uint128\"}],\"name\":\"scan\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getTotalStakeForStakee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"getEntries\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506110b5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80638ce5749611610081578063dc12c9b81161005b578063dc12c9b8146101cb578063f2fde38b146101d3578063fa42e7ea146101e657600080fd5b80638ce57496146101865780638da5cb5b14610199578063d3549ebd146101aa57600080fd5b8063485cc955116100b2578063485cc9551461013e5780634e68066114610153578063715018a61461017e57600080fd5b80631bdcc9ad146100ce57806345367f231461011b575b600080fd5b6101086100dc366004610f24565b60008281526068602090815260408083206001600160a01b038516845260010190915290205492915050565b6040519081526020015b60405180910390f35b610108610129366004610ef4565b60009081526068602052604090206002015490565b61015161014c366004610e8c565b6101ef565b005b610166610161366004610ec4565b6102f9565b6040516001600160a01b039091168152602001610112565b6101516104e6565b610151610194366004610ef4565b61054c565b6033546001600160a01b0316610166565b6101bd6101b8366004610ef4565b6105ab565b604051610112929190610f48565b610151610784565b6101516101e1366004610e69565b610b03565b61010860675481565b600054610100900460ff1680610208575060005460ff16155b6102705760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b600054610100900460ff16158015610292576000805461ffff19166101011790555b61029a610be5565b606580546001600160a01b038086167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255606680549285169290911691909117905580156102f4576000805461ff00191690555b505050565b60675460009081526068602052604081205461031757506000919050565b60675460009081526068602052604081206002015460809061034c906fffffffffffffffffffffffffffffffff861690611003565b6067546000908152606860205260408120549190921c9250819061037290600190611022565b90505b8082116104de576000600261038a8385610fcb565b6103949190610fe3565b9050600081156103ef5760675460009081526068602052604090206103ba600184611022565b815481106103d857634e487b7160e01b600052603260045260246000fd5b9060005260206000209060020201600101546103f2565b60005b6067546000908152606860205260408120805492935090918490811061042857634e487b7160e01b600052603260045260246000fd5b906000526020600020906002020160010154905081861015801561044b57508086105b156104a757606754600090815260686020526040902080548490811061048157634e487b7160e01b600052603260045260246000fd5b60009182526020909120600290910201546001600160a01b031698975050505050505050565b818610156104c1576104ba600184611022565b93506104d6565b8086106104d6576104d3836001610fcb565b94505b505050610375565b505050919050565b6033546001600160a01b031633146105405760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610267565b61054a6000610ca7565b565b6033546001600160a01b031633146105a65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610267565b606755565b600081815260686020526040812054606091829167ffffffffffffffff8111156105e557634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561060e578160200160208202803683370190505b506000858152606860205260408120549192509067ffffffffffffffff81111561064857634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610671578160200160208202803683370190505b50905060005b6000868152606860205260409020548110156107795760008681526068602052604081208054839081106106bb57634e487b7160e01b600052603260045260246000fd5b60009182526020918290206040805180820190915260029092020180546001600160a01b031680835260019091015492820192909252855190925085908490811061071657634e487b7160e01b600052603260045260246000fd5b60200260200101906001600160a01b031690816001600160a01b031681525050806020015183838151811061075b57634e487b7160e01b600052603260045260246000fd5b6020908102919091010152508061077181611039565b915050610677565b509094909350915050565b6065546040517f13cdd31b0000000000000000000000000000000000000000000000000000000081523360048201819052916000916001600160a01b03909116906313cdd31b9060240160206040518083038186803b1580156107e657600080fd5b505afa1580156107fa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061081e9190610f0c565b6066546040517ff1fd0d1c0000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301529293506000929091169063f1fd0d1c9060240160206040518083038186803b15801561088257600080fd5b505afa158015610896573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ba9190610f0c565b905060006108c88284610fcb565b9050600081116109405760405162461bcd60e51b815260206004820152603760248201527f43616e206e6f74206a6f696e206469726563746f727920666f72206e6578742060448201527f65706f636820776974686f757420616e79207374616b650000000000000000006064820152608401610267565b600060675460016109519190610fcb565b60008181526068602090815260408083206001600160a01b038a168452600101909152902054909150156109ed5760405162461bcd60e51b815260206004820152602a60248201527f43616e206f6e6c79206a6f696e20746865206469726563746f7279206f6e636560448201527f207065722065706f6368000000000000000000000000000000000000000000006064820152608401610267565b600081815260686020526040812060020154610a0a908490610fcb565b9050606860008381526020019081526020016000206000016040518060400160405280886001600160a01b0316815260200183815250908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010155505082606860008481526020019081526020016000206001016000886001600160a01b03166001600160a01b0316815260200190815260200160002081905550806068600084815260200190815260200160002060020181905550505050505050565b6033546001600160a01b03163314610b5d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610267565b6001600160a01b038116610bd95760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610267565b610be281610ca7565b50565b600054610100900460ff1680610bfe575060005460ff16155b610c615760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610267565b600054610100900460ff16158015610c83576000805461ffff19166101011790555b610c8b610d11565b610c93610dc2565b8015610be2576000805461ff001916905550565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff1680610d2a575060005460ff16155b610d8d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610267565b600054610100900460ff16158015610c93576000805461ffff19166101011790558015610be2576000805461ff001916905550565b600054610100900460ff1680610ddb575060005460ff16155b610e3e5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610267565b600054610100900460ff16158015610e60576000805461ffff19166101011790555b610c9333610ca7565b600060208284031215610e7a578081fd5b8135610e858161106a565b9392505050565b60008060408385031215610e9e578081fd5b8235610ea98161106a565b91506020830135610eb98161106a565b809150509250929050565b600060208284031215610ed5578081fd5b81356fffffffffffffffffffffffffffffffff81168114610e85578182fd5b600060208284031215610f05578081fd5b5035919050565b600060208284031215610f1d578081fd5b5051919050565b60008060408385031215610f36578182fd5b823591506020830135610eb98161106a565b604080825283519082018190526000906020906060840190828701845b82811015610f8a5781516001600160a01b031684529284019290840190600101610f65565b50505083810382850152845180825285830191830190845b81811015610fbe57835183529284019291840191600101610fa2565b5090979650505050505050565b60008219821115610fde57610fde611054565b500190565b600082610ffe57634e487b7160e01b81526012600452602481fd5b500490565b600081600019048311821515161561101d5761101d611054565b500290565b60008282101561103457611034611054565b500390565b600060001982141561104d5761104d611054565b5060010190565b634e487b7160e01b600052601160045260246000fd5b6001600160a01b0381168114610be257600080fdfea2646970667358221220b00ca80b1fd637ca9fd2ef27cde11a3894c32ea39f32132954d10f6e4f3c7c1364736f6c63430008040033",
}

// DirectoryABI is the input ABI used to generate the binding from.
// Deprecated: Use DirectoryMetaData.ABI instead.
var DirectoryABI = DirectoryMetaData.ABI

// DirectoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DirectoryMetaData.Bin instead.
var DirectoryBin = DirectoryMetaData.Bin

// DeployDirectory deploys a new Ethereum contract, binding an instance of Directory to it.
func DeployDirectory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Directory, error) {
	parsed, err := DirectoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DirectoryBin), backend)
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
// Solidity: function currentDirectory() view returns(uint256)
func (_Directory *DirectoryCaller) CurrentDirectory(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "currentDirectory")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentDirectory is a free data retrieval call binding the contract method 0xfa42e7ea.
//
// Solidity: function currentDirectory() view returns(uint256)
func (_Directory *DirectorySession) CurrentDirectory() (*big.Int, error) {
	return _Directory.Contract.CurrentDirectory(&_Directory.CallOpts)
}

// CurrentDirectory is a free data retrieval call binding the contract method 0xfa42e7ea.
//
// Solidity: function currentDirectory() view returns(uint256)
func (_Directory *DirectoryCallerSession) CurrentDirectory() (*big.Int, error) {
	return _Directory.Contract.CurrentDirectory(&_Directory.CallOpts)
}

// GetEntries is a free data retrieval call binding the contract method 0xd3549ebd.
//
// Solidity: function getEntries(uint256 epochId) view returns(address[], uint256[])
func (_Directory *DirectoryCaller) GetEntries(opts *bind.CallOpts, epochId *big.Int) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "getEntries", epochId)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetEntries is a free data retrieval call binding the contract method 0xd3549ebd.
//
// Solidity: function getEntries(uint256 epochId) view returns(address[], uint256[])
func (_Directory *DirectorySession) GetEntries(epochId *big.Int) ([]common.Address, []*big.Int, error) {
	return _Directory.Contract.GetEntries(&_Directory.CallOpts, epochId)
}

// GetEntries is a free data retrieval call binding the contract method 0xd3549ebd.
//
// Solidity: function getEntries(uint256 epochId) view returns(address[], uint256[])
func (_Directory *DirectoryCallerSession) GetEntries(epochId *big.Int) ([]common.Address, []*big.Int, error) {
	return _Directory.Contract.GetEntries(&_Directory.CallOpts, epochId)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x45367f23.
//
// Solidity: function getTotalStake(uint256 epochId) view returns(uint256)
func (_Directory *DirectoryCaller) GetTotalStake(opts *bind.CallOpts, epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "getTotalStake", epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStake is a free data retrieval call binding the contract method 0x45367f23.
//
// Solidity: function getTotalStake(uint256 epochId) view returns(uint256)
func (_Directory *DirectorySession) GetTotalStake(epochId *big.Int) (*big.Int, error) {
	return _Directory.Contract.GetTotalStake(&_Directory.CallOpts, epochId)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x45367f23.
//
// Solidity: function getTotalStake(uint256 epochId) view returns(uint256)
func (_Directory *DirectoryCallerSession) GetTotalStake(epochId *big.Int) (*big.Int, error) {
	return _Directory.Contract.GetTotalStake(&_Directory.CallOpts, epochId)
}

// GetTotalStakeForStakee is a free data retrieval call binding the contract method 0x1bdcc9ad.
//
// Solidity: function getTotalStakeForStakee(uint256 epochId, address stakee) view returns(uint256)
func (_Directory *DirectoryCaller) GetTotalStakeForStakee(opts *bind.CallOpts, epochId *big.Int, stakee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "getTotalStakeForStakee", epochId, stakee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStakeForStakee is a free data retrieval call binding the contract method 0x1bdcc9ad.
//
// Solidity: function getTotalStakeForStakee(uint256 epochId, address stakee) view returns(uint256)
func (_Directory *DirectorySession) GetTotalStakeForStakee(epochId *big.Int, stakee common.Address) (*big.Int, error) {
	return _Directory.Contract.GetTotalStakeForStakee(&_Directory.CallOpts, epochId, stakee)
}

// GetTotalStakeForStakee is a free data retrieval call binding the contract method 0x1bdcc9ad.
//
// Solidity: function getTotalStakeForStakee(uint256 epochId, address stakee) view returns(uint256)
func (_Directory *DirectoryCallerSession) GetTotalStakeForStakee(epochId *big.Int, stakee common.Address) (*big.Int, error) {
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
// Solidity: function scan(uint128 point) view returns(address stakee)
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
// Solidity: function scan(uint128 point) view returns(address stakee)
func (_Directory *DirectorySession) Scan(point *big.Int) (common.Address, error) {
	return _Directory.Contract.Scan(&_Directory.CallOpts, point)
}

// Scan is a free data retrieval call binding the contract method 0x4e680661.
//
// Solidity: function scan(uint128 point) view returns(address stakee)
func (_Directory *DirectoryCallerSession) Scan(point *big.Int) (common.Address, error) {
	return _Directory.Contract.Scan(&_Directory.CallOpts, point)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address stakingManager, address rewardsManager) returns()
func (_Directory *DirectoryTransactor) Initialize(opts *bind.TransactOpts, stakingManager common.Address, rewardsManager common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "initialize", stakingManager, rewardsManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address stakingManager, address rewardsManager) returns()
func (_Directory *DirectorySession) Initialize(stakingManager common.Address, rewardsManager common.Address) (*types.Transaction, error) {
	return _Directory.Contract.Initialize(&_Directory.TransactOpts, stakingManager, rewardsManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address stakingManager, address rewardsManager) returns()
func (_Directory *DirectoryTransactorSession) Initialize(stakingManager common.Address, rewardsManager common.Address) (*types.Transaction, error) {
	return _Directory.Contract.Initialize(&_Directory.TransactOpts, stakingManager, rewardsManager)
}

// JoinNextDirectory is a paid mutator transaction binding the contract method 0xdc12c9b8.
//
// Solidity: function joinNextDirectory() returns()
func (_Directory *DirectoryTransactor) JoinNextDirectory(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "joinNextDirectory")
}

// JoinNextDirectory is a paid mutator transaction binding the contract method 0xdc12c9b8.
//
// Solidity: function joinNextDirectory() returns()
func (_Directory *DirectorySession) JoinNextDirectory() (*types.Transaction, error) {
	return _Directory.Contract.JoinNextDirectory(&_Directory.TransactOpts)
}

// JoinNextDirectory is a paid mutator transaction binding the contract method 0xdc12c9b8.
//
// Solidity: function joinNextDirectory() returns()
func (_Directory *DirectoryTransactorSession) JoinNextDirectory() (*types.Transaction, error) {
	return _Directory.Contract.JoinNextDirectory(&_Directory.TransactOpts)
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

// SetCurrentDirectory is a paid mutator transaction binding the contract method 0x8ce57496.
//
// Solidity: function setCurrentDirectory(uint256 epochId) returns()
func (_Directory *DirectoryTransactor) SetCurrentDirectory(opts *bind.TransactOpts, epochId *big.Int) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "setCurrentDirectory", epochId)
}

// SetCurrentDirectory is a paid mutator transaction binding the contract method 0x8ce57496.
//
// Solidity: function setCurrentDirectory(uint256 epochId) returns()
func (_Directory *DirectorySession) SetCurrentDirectory(epochId *big.Int) (*types.Transaction, error) {
	return _Directory.Contract.SetCurrentDirectory(&_Directory.TransactOpts, epochId)
}

// SetCurrentDirectory is a paid mutator transaction binding the contract method 0x8ce57496.
//
// Solidity: function setCurrentDirectory(uint256 epochId) returns()
func (_Directory *DirectoryTransactorSession) SetCurrentDirectory(epochId *big.Int) (*types.Transaction, error) {
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
