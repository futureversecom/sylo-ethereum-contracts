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

// DirectoryStake is an auto generated low-level Go binding around an user-defined struct.
type DirectoryStake struct {
	Staker common.Address
	Amount *big.Int
}

// DirectoryABI is the input ABI used to generate the binding from.
const DirectoryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"currentDirectory\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractStakingManager\",\"name\":\"stakingManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"constructDirectory\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"direcrtoryId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"point\",\"type\":\"uint128\"}],\"name\":\"scan\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"directoryId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getStakes\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structDirectory.Stake[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DirectoryBin is the compiled bytecode used for deploying new contracts.
var DirectoryBin = "0x608060405234801561001057600080fd5b50611193806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638da5cb5b1161005b5780638da5cb5b14610101578063c4d66de814610112578063f2fde38b14610125578063fa42e7ea1461013857600080fd5b80631fd05a301461008d5780634e680661146100b6578063715018a6146100e1578063859ea347146100eb575b600080fd5b6100a061009b366004610f3e565b610141565b6040516100ad919061100a565b60405180910390f35b6100c96100c4366004610fc2565b6101d7565b6040516001600160a01b0390911681526020016100ad565b6100e96103c7565b005b6100f3610432565b6040519081526020016100ad565b6033546001600160a01b03166100c9565b6100e9610120366004610e4b565b610a0c565b6100e9610133366004610e4b565b610af0565b6100f360665481565b60008281526067602090815260408083206001600160a01b03851684526001018252808320805482518185028101850190935280835260609492939192909184015b828210156101cb576000848152602090819020604080518082019091526002850290910180546001600160a01b03168252600190810154828401529083529092019101610183565b50505050905092915050565b6066546000908152606760205260408120546101f557506000919050565b60665460009081526067602052604081206002015460809061022a906fffffffffffffffffffffffffffffffff8616906110cb565b6066546000908152606760205260408120549190921c92508190610250906001906110ea565b90505b8082116103bc57600060026102688385611093565b61027291906110ab565b9050600081156102cd5760665460009081526067602052604090206102986001846110ea565b815481106102b657634e487b7160e01b600052603260045260246000fd5b9060005260206000209060020201600101546102d0565b60005b6066546000908152606760205260408120805492935090918490811061030657634e487b7160e01b600052603260045260246000fd5b906000526020600020906002020160010154905081861015801561032957508086105b1561038557606654600090815260676020526040902080548490811061035f57634e487b7160e01b600052603260045260246000fd5b60009182526020909120600290910201546001600160a01b031698975050505050505050565b8186101561039f576103986001846110ea565b93506103b4565b8086106103b4576103b1836001611093565b94505b505050610253565b506000949350505050565b6033546001600160a01b031633146104265760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6104306000610bd2565b565b6033546000906001600160a01b0316331461048f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161041d565b60408051436020820152600091016040516020818303038152906040528051906020012090506000805b606560009054906101000a90046001600160a01b03166001600160a01b03166338b9437d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561050757600080fd5b505afa15801561051b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061053f9190610ff2565b811015610968576065546040517fd2d6c0da000000000000000000000000000000000000000000000000000000008152600481018390526000916001600160a01b03169063d2d6c0da9060240160206040518083038186803b1580156105a457600080fd5b505afa1580156105b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105dc9190610e6e565b6065546040517fdf349ed50000000000000000000000000000000000000000000000000000000081526001600160a01b0380841660048301529293506000929091169063df349ed59060240160206040518083038186803b15801561064057600080fd5b505afa158015610654573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106789190610ff2565b9050600181101561068a575050610956565b60008581526067602090815260409182902082518084019093526001600160a01b0385168352919081016106be8488611093565b9052815460018082018455600093845260208085208451600290940201805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039485161781559301519201919091556065546040517f79193610000000000000000000000000000000000000000000000000000000008152858316600482015291169063791936109060240160006040518083038186803b15801561076057600080fd5b505afa158015610774573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261079c9190810190610e8a565b905060005b81518110156109455760655482516000916001600160a01b0316906382dda22d908590859081106107e257634e487b7160e01b600052603260045260246000fd5b6020026020010151876040518363ffffffff1660e01b815260040161081d9291906001600160a01b0392831681529116602082015260400190565b604080518083038186803b15801561083457600080fd5b505afa158015610848573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061086c9190610f6d565b9050606760008981526020019081526020016000206001016000866001600160a01b03166001600160a01b0316815260200190815260200160002060405180604001604052808585815181106108d257634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b0390811683529451918101919091528254600180820185556000948552938290208351600290920201805473ffffffffffffffffffffffffffffffffffffffff19169190951617845501519101558061093d81611101565b9150506107a1565b506109508286611093565b94505050505b8061096081611101565b9150506104b9565b50606560009054906101000a90046001600160a01b03166001600160a01b0316637bc742256040518163ffffffff1660e01b815260040160206040518083038186803b1580156109b757600080fd5b505afa1580156109cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ef9190610ff2565b600083815260676020526040902060020155506066819055905090565b600054610100900460ff1680610a25575060005460ff16155b610a885760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161041d565b600054610100900460ff16158015610aaa576000805461ffff19166101011790555b610ab2610c31565b6065805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0384161790558015610aec576000805461ff00191690555b5050565b6033546001600160a01b03163314610b4a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161041d565b6001600160a01b038116610bc65760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161041d565b610bcf81610bd2565b50565b603380546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff1680610c4a575060005460ff16155b610cad5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161041d565b600054610100900460ff16158015610ccf576000805461ffff19166101011790555b610cd7610cf3565b610cdf610da4565b8015610bcf576000805461ff001916905550565b600054610100900460ff1680610d0c575060005460ff16155b610d6f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161041d565b600054610100900460ff16158015610cdf576000805461ffff19166101011790558015610bcf576000805461ff001916905550565b600054610100900460ff1680610dbd575060005460ff16155b610e205760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161041d565b600054610100900460ff16158015610e42576000805461ffff19166101011790555b610cdf33610bd2565b600060208284031215610e5c578081fd5b8135610e6781611148565b9392505050565b600060208284031215610e7f578081fd5b8151610e6781611148565b60006020808385031215610e9c578182fd5b825167ffffffffffffffff80821115610eb3578384fd5b818501915085601f830112610ec6578384fd5b815181811115610ed857610ed8611132565b8060051b9150610ee9848301611062565b8181528481019084860184860187018a1015610f03578788fd5b8795505b83861015610f315780519450610f1c85611148565b84835260019590950194918601918601610f07565b5098975050505050505050565b60008060408385031215610f50578081fd5b823591506020830135610f6281611148565b809150509250929050565b600060408284031215610f7e578081fd5b6040516040810181811067ffffffffffffffff82111715610fa157610fa1611132565b604052825181526020830151610fb681611148565b60208201529392505050565b600060208284031215610fd3578081fd5b81356fffffffffffffffffffffffffffffffff81168114610e67578182fd5b600060208284031215611003578081fd5b5051919050565b602080825282518282018190526000919060409081850190868401855b8281101561105557815180516001600160a01b03168552860151868501529284019290850190600101611027565b5091979650505050505050565b604051601f8201601f1916810167ffffffffffffffff8111828210171561108b5761108b611132565b604052919050565b600082198211156110a6576110a661111c565b500190565b6000826110c657634e487b7160e01b81526012600452602481fd5b500490565b60008160001904831182151516156110e5576110e561111c565b500290565b6000828210156110fc576110fc61111c565b500390565b60006000198214156111155761111561111c565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114610bcf57600080fdfea264697066735822122068ab6a80326758a73108a8d09106f97186f9fa044c0a85f13b3b564306bcb74664736f6c63430008040033"

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

// GetStakes is a free data retrieval call binding the contract method 0x1fd05a30.
//
// Solidity: function getStakes(bytes32 directoryId, address stakee) view returns((address,uint256)[])
func (_Directory *DirectoryCaller) GetStakes(opts *bind.CallOpts, directoryId [32]byte, stakee common.Address) ([]DirectoryStake, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "getStakes", directoryId, stakee)

	if err != nil {
		return *new([]DirectoryStake), err
	}

	out0 := *abi.ConvertType(out[0], new([]DirectoryStake)).(*[]DirectoryStake)

	return out0, err

}

// GetStakes is a free data retrieval call binding the contract method 0x1fd05a30.
//
// Solidity: function getStakes(bytes32 directoryId, address stakee) view returns((address,uint256)[])
func (_Directory *DirectorySession) GetStakes(directoryId [32]byte, stakee common.Address) ([]DirectoryStake, error) {
	return _Directory.Contract.GetStakes(&_Directory.CallOpts, directoryId, stakee)
}

// GetStakes is a free data retrieval call binding the contract method 0x1fd05a30.
//
// Solidity: function getStakes(bytes32 directoryId, address stakee) view returns((address,uint256)[])
func (_Directory *DirectoryCallerSession) GetStakes(directoryId [32]byte, stakee common.Address) ([]DirectoryStake, error) {
	return _Directory.Contract.GetStakes(&_Directory.CallOpts, directoryId, stakee)
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

// ConstructDirectory is a paid mutator transaction binding the contract method 0x859ea347.
//
// Solidity: function constructDirectory() returns(bytes32 direcrtoryId)
func (_Directory *DirectoryTransactor) ConstructDirectory(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "constructDirectory")
}

// ConstructDirectory is a paid mutator transaction binding the contract method 0x859ea347.
//
// Solidity: function constructDirectory() returns(bytes32 direcrtoryId)
func (_Directory *DirectorySession) ConstructDirectory() (*types.Transaction, error) {
	return _Directory.Contract.ConstructDirectory(&_Directory.TransactOpts)
}

// ConstructDirectory is a paid mutator transaction binding the contract method 0x859ea347.
//
// Solidity: function constructDirectory() returns(bytes32 direcrtoryId)
func (_Directory *DirectoryTransactorSession) ConstructDirectory() (*types.Transaction, error) {
	return _Directory.Contract.ConstructDirectory(&_Directory.TransactOpts)
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
