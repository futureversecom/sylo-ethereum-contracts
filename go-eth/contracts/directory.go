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

// DirectoryABI is the input ABI used to generate the binding from.
const DirectoryABI = "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unlockDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"minimumStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newUnlockDuration\",\"type\":\"uint256\"}],\"name\":\"setUnlockDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMinimumStake\",\"type\":\"uint256\"}],\"name\":\"setMinimumStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"addStakeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"unstakeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rand\",\"type\":\"uint256\"}],\"name\":\"scan\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DirectoryBin is the compiled bytecode used for deploying new contracts.
var DirectoryBin = "0x608060405234801561001057600080fd5b50604051610d7b380380610d7b8339818101604052604081101561003357600080fd5b508051602090910151600061004f6001600160e01b036100c216565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600180546001600160a01b0319166001600160a01b0393909316929092179091556002556100c6565b3390565b610ca6806100d56000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80638ee82c3211610097578063e1c98d8c11610066578063e1c98d8c1461024d578063eb4f16b514610255578063ec5ffac214610272578063f2fde38b1461027a57610100565b80638ee82c32146101e5578063bb9fe6bf1461020b578063dd90076914610213578063e16d14a01461023057610100565b8063715018a6116100d3578063715018a6146101975780637bc742251461019f5780638a1fcd60146101b95780638da5cb5b146101c157610100565b806316934fc414610105578063233e9903146101445780632def6620146101635780636bbfea211461016b575b600080fd5b61012b6004803603602081101561011b57600080fd5b50356001600160a01b03166102a0565b6040805192835260208301919091528051918290030190f35b6101616004803603602081101561015a57600080fd5b50356102b9565b005b610161610328565b6101616004803603604081101561018157600080fd5b50803590602001356001600160a01b0316610333565b61016161054f565b6101a7610610565b60408051918252519081900360200190f35b6101a761069d565b6101c96106a3565b604080516001600160a01b039092168252519081900360200190f35b610161600480360360208110156101fb57600080fd5b50356001600160a01b03166106b2565b6101a7610833565b6101616004803603602081101561022957600080fd5b50356108fe565b6101c96004803603602081101561024657600080fd5b503561096d565b610161610a53565b6101616004803603602081101561026b57600080fd5b5035610ac3565b6101a7610ad0565b6101616004803603602081101561029057600080fd5b50356001600160a01b0316610ad6565b6004602052600090815260409020805460019091015482565b6102c1610b49565b6000546001600160a01b03908116911614610323576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600355565b610331336106b2565b565b6001600160a01b03811661038e576040805162461bcd60e51b815260206004820152600f60248201527f41646472657373206973206e756c6c0000000000000000000000000000000000604482015290519081900360640190fd5b816103e0576040805162461bcd60e51b815260206004820152601460248201527f43616e6e6f74207374616b65206e6f7468696e67000000000000000000000000604482015290519081900360640190fd5b60006103eb82610b4d565b90508060010154600014610446576040805162461bcd60e51b815260206004820152601c60248201527f43616e6e6f74207374616b65207768696c6520756e6c6f636b696e6700000000604482015290519081900360640190fd5b80546104a557600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0384161790555b805483018155600154604080517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810186905290516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b15801561051e57600080fd5b505af1158015610532573d6000803e3d6000fd5b505050506040513d602081101561054857600080fd5b5050505050565b610557610b49565b6000546001600160a01b039081169116146105b9576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600080805b60055481101561069757610627610c30565b600460006005848154811061063857fe5b60009182526020808320909101546001600160a01b0316835282810193909352604091820190208151808301909252805482526001015491810191909152905061068181610b67565b1561068e57805192909201915b50600101610615565b50905090565b60025481565b6000546001600160a01b031690565b60006106bd33610b4d565b6040805180820190915281548152600182015460208201529091506106e190610b67565b610732576040805162461bcd60e51b815260206004820152601660248201527f5374616b65206e6f7420776974686472617761626c6500000000000000000000604482015290519081900360640190fd5b80546000808355600183018190555b6005548110156107c357336001600160a01b03166005828154811061076257fe5b6000918252602090912001546001600160a01b031614156107bb5760006005828154811061078c57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b600101610741565b50600154604080517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b15801561051e57600080fd5b60008061083f33610b4d565b8054909150610895576040805162461bcd60e51b815260206004820152601260248201527f4e6f7468696e6720746f20756e7374616b650000000000000000000000000000604482015290519081900360640190fd5b6001810154156108ec576040805162461bcd60e51b815260206004820152601160248201527f416c726561647920756e6c6f636b696e67000000000000000000000000000000604482015290519081900360640190fd5b60025443016001909101819055905090565b610906610b49565b6000546001600160a01b03908116911614610968576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600255565b60008060fd8361097b610610565b028161098357fe5b0490506000805b600554811015610a465761099c610c30565b60046000600584815481106109ad57fe5b60009182526020808320909101546001600160a01b031683528281019390935260409182019020815180830190925280548252600101549181019190915290506109f681610b67565b610a005750610a3e565b80519290920191828411610a3c5760058281548110610a1b57fe5b6000918252602090912001546001600160a01b03169450610a4e9350505050565b505b60010161098a565b506000925050505b919050565b6000610a5e33610b4d565b90506000816001015411610ab9576040805162461bcd60e51b815260206004820152601960248201527f4e6f7420756e6c6f636b696e672063616e6e6f74206c6f636b00000000000000604482015290519081900360640190fd5b6000600190910155565b610acd8133610333565b50565b60035481565b610ade610b49565b6000546001600160a01b03908116911614610b40576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b610acd81610b83565b3390565b6001600160a01b0316600090815260046020526040902090565b805160009015801590610b7d5750438260200151105b92915050565b6001600160a01b038116610bc85760405162461bcd60e51b8152600401808060200182810382526026815260200180610c4b6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b60405180604001604052806000815260200160008152509056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373a26469706673582212200f0c0823d6b739c72f1145d9df4cf08ddce363854988b5dbe60037d50f0298bd64736f6c63430006040033"

// DeployDirectory deploys a new Ethereum contract, binding an instance of Directory to it.
func DeployDirectory(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, _unlockDuration *big.Int) (common.Address, *types.Transaction, *Directory, error) {
	parsed, err := abi.JSON(strings.NewReader(DirectoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DirectoryBin), backend, token, _unlockDuration)
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
func (_Directory *DirectoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_Directory *DirectoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() constant returns(uint256)
func (_Directory *DirectoryCaller) GetTotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "getTotalStake")
	return *ret0, err
}

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() constant returns(uint256)
func (_Directory *DirectorySession) GetTotalStake() (*big.Int, error) {
	return _Directory.Contract.GetTotalStake(&_Directory.CallOpts)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() constant returns(uint256)
func (_Directory *DirectoryCallerSession) GetTotalStake() (*big.Int, error) {
	return _Directory.Contract.GetTotalStake(&_Directory.CallOpts)
}

// MinimumStake is a free data retrieval call binding the contract method 0xec5ffac2.
//
// Solidity: function minimumStake() constant returns(uint256)
func (_Directory *DirectoryCaller) MinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "minimumStake")
	return *ret0, err
}

// MinimumStake is a free data retrieval call binding the contract method 0xec5ffac2.
//
// Solidity: function minimumStake() constant returns(uint256)
func (_Directory *DirectorySession) MinimumStake() (*big.Int, error) {
	return _Directory.Contract.MinimumStake(&_Directory.CallOpts)
}

// MinimumStake is a free data retrieval call binding the contract method 0xec5ffac2.
//
// Solidity: function minimumStake() constant returns(uint256)
func (_Directory *DirectoryCallerSession) MinimumStake() (*big.Int, error) {
	return _Directory.Contract.MinimumStake(&_Directory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Directory *DirectoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Directory *DirectorySession) Owner() (common.Address, error) {
	return _Directory.Contract.Owner(&_Directory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Directory *DirectoryCallerSession) Owner() (common.Address, error) {
	return _Directory.Contract.Owner(&_Directory.CallOpts)
}

// Scan is a free data retrieval call binding the contract method 0xe16d14a0.
//
// Solidity: function scan(uint256 rand) constant returns(address)
func (_Directory *DirectoryCaller) Scan(opts *bind.CallOpts, rand *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "scan", rand)
	return *ret0, err
}

// Scan is a free data retrieval call binding the contract method 0xe16d14a0.
//
// Solidity: function scan(uint256 rand) constant returns(address)
func (_Directory *DirectorySession) Scan(rand *big.Int) (common.Address, error) {
	return _Directory.Contract.Scan(&_Directory.CallOpts, rand)
}

// Scan is a free data retrieval call binding the contract method 0xe16d14a0.
//
// Solidity: function scan(uint256 rand) constant returns(address)
func (_Directory *DirectoryCallerSession) Scan(rand *big.Int) (common.Address, error) {
	return _Directory.Contract.Scan(&_Directory.CallOpts, rand)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) constant returns(uint256 amount, uint256 unlockAt)
func (_Directory *DirectoryCaller) Stakes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount   *big.Int
	UnlockAt *big.Int
}, error) {
	ret := new(struct {
		Amount   *big.Int
		UnlockAt *big.Int
	})
	out := ret
	err := _Directory.contract.Call(opts, out, "stakes", arg0)
	return *ret, err
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) constant returns(uint256 amount, uint256 unlockAt)
func (_Directory *DirectorySession) Stakes(arg0 common.Address) (struct {
	Amount   *big.Int
	UnlockAt *big.Int
}, error) {
	return _Directory.Contract.Stakes(&_Directory.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) constant returns(uint256 amount, uint256 unlockAt)
func (_Directory *DirectoryCallerSession) Stakes(arg0 common.Address) (struct {
	Amount   *big.Int
	UnlockAt *big.Int
}, error) {
	return _Directory.Contract.Stakes(&_Directory.CallOpts, arg0)
}

// UnlockDuration is a free data retrieval call binding the contract method 0x8a1fcd60.
//
// Solidity: function unlockDuration() constant returns(uint256)
func (_Directory *DirectoryCaller) UnlockDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "unlockDuration")
	return *ret0, err
}

// UnlockDuration is a free data retrieval call binding the contract method 0x8a1fcd60.
//
// Solidity: function unlockDuration() constant returns(uint256)
func (_Directory *DirectorySession) UnlockDuration() (*big.Int, error) {
	return _Directory.Contract.UnlockDuration(&_Directory.CallOpts)
}

// UnlockDuration is a free data retrieval call binding the contract method 0x8a1fcd60.
//
// Solidity: function unlockDuration() constant returns(uint256)
func (_Directory *DirectoryCallerSession) UnlockDuration() (*big.Int, error) {
	return _Directory.Contract.UnlockDuration(&_Directory.CallOpts)
}

// AddStake is a paid mutator transaction binding the contract method 0xeb4f16b5.
//
// Solidity: function addStake(uint256 amount) returns()
func (_Directory *DirectoryTransactor) AddStake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "addStake", amount)
}

// AddStake is a paid mutator transaction binding the contract method 0xeb4f16b5.
//
// Solidity: function addStake(uint256 amount) returns()
func (_Directory *DirectorySession) AddStake(amount *big.Int) (*types.Transaction, error) {
	return _Directory.Contract.AddStake(&_Directory.TransactOpts, amount)
}

// AddStake is a paid mutator transaction binding the contract method 0xeb4f16b5.
//
// Solidity: function addStake(uint256 amount) returns()
func (_Directory *DirectoryTransactorSession) AddStake(amount *big.Int) (*types.Transaction, error) {
	return _Directory.Contract.AddStake(&_Directory.TransactOpts, amount)
}

// AddStakeFor is a paid mutator transaction binding the contract method 0x6bbfea21.
//
// Solidity: function addStakeFor(uint256 amount, address staker) returns()
func (_Directory *DirectoryTransactor) AddStakeFor(opts *bind.TransactOpts, amount *big.Int, staker common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "addStakeFor", amount, staker)
}

// AddStakeFor is a paid mutator transaction binding the contract method 0x6bbfea21.
//
// Solidity: function addStakeFor(uint256 amount, address staker) returns()
func (_Directory *DirectorySession) AddStakeFor(amount *big.Int, staker common.Address) (*types.Transaction, error) {
	return _Directory.Contract.AddStakeFor(&_Directory.TransactOpts, amount, staker)
}

// AddStakeFor is a paid mutator transaction binding the contract method 0x6bbfea21.
//
// Solidity: function addStakeFor(uint256 amount, address staker) returns()
func (_Directory *DirectoryTransactorSession) AddStakeFor(amount *big.Int, staker common.Address) (*types.Transaction, error) {
	return _Directory.Contract.AddStakeFor(&_Directory.TransactOpts, amount, staker)
}

// LockStake is a paid mutator transaction binding the contract method 0xe1c98d8c.
//
// Solidity: function lockStake() returns()
func (_Directory *DirectoryTransactor) LockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "lockStake")
}

// LockStake is a paid mutator transaction binding the contract method 0xe1c98d8c.
//
// Solidity: function lockStake() returns()
func (_Directory *DirectorySession) LockStake() (*types.Transaction, error) {
	return _Directory.Contract.LockStake(&_Directory.TransactOpts)
}

// LockStake is a paid mutator transaction binding the contract method 0xe1c98d8c.
//
// Solidity: function lockStake() returns()
func (_Directory *DirectoryTransactorSession) LockStake() (*types.Transaction, error) {
	return _Directory.Contract.LockStake(&_Directory.TransactOpts)
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

// SetMinimumStake is a paid mutator transaction binding the contract method 0x233e9903.
//
// Solidity: function setMinimumStake(uint256 newMinimumStake) returns()
func (_Directory *DirectoryTransactor) SetMinimumStake(opts *bind.TransactOpts, newMinimumStake *big.Int) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "setMinimumStake", newMinimumStake)
}

// SetMinimumStake is a paid mutator transaction binding the contract method 0x233e9903.
//
// Solidity: function setMinimumStake(uint256 newMinimumStake) returns()
func (_Directory *DirectorySession) SetMinimumStake(newMinimumStake *big.Int) (*types.Transaction, error) {
	return _Directory.Contract.SetMinimumStake(&_Directory.TransactOpts, newMinimumStake)
}

// SetMinimumStake is a paid mutator transaction binding the contract method 0x233e9903.
//
// Solidity: function setMinimumStake(uint256 newMinimumStake) returns()
func (_Directory *DirectoryTransactorSession) SetMinimumStake(newMinimumStake *big.Int) (*types.Transaction, error) {
	return _Directory.Contract.SetMinimumStake(&_Directory.TransactOpts, newMinimumStake)
}

// SetUnlockDuration is a paid mutator transaction binding the contract method 0xdd900769.
//
// Solidity: function setUnlockDuration(uint256 newUnlockDuration) returns()
func (_Directory *DirectoryTransactor) SetUnlockDuration(opts *bind.TransactOpts, newUnlockDuration *big.Int) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "setUnlockDuration", newUnlockDuration)
}

// SetUnlockDuration is a paid mutator transaction binding the contract method 0xdd900769.
//
// Solidity: function setUnlockDuration(uint256 newUnlockDuration) returns()
func (_Directory *DirectorySession) SetUnlockDuration(newUnlockDuration *big.Int) (*types.Transaction, error) {
	return _Directory.Contract.SetUnlockDuration(&_Directory.TransactOpts, newUnlockDuration)
}

// SetUnlockDuration is a paid mutator transaction binding the contract method 0xdd900769.
//
// Solidity: function setUnlockDuration(uint256 newUnlockDuration) returns()
func (_Directory *DirectoryTransactorSession) SetUnlockDuration(newUnlockDuration *big.Int) (*types.Transaction, error) {
	return _Directory.Contract.SetUnlockDuration(&_Directory.TransactOpts, newUnlockDuration)
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

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns(uint256)
func (_Directory *DirectoryTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns(uint256)
func (_Directory *DirectorySession) UnlockStake() (*types.Transaction, error) {
	return _Directory.Contract.UnlockStake(&_Directory.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns(uint256)
func (_Directory *DirectoryTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _Directory.Contract.UnlockStake(&_Directory.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Directory *DirectoryTransactor) Unstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "unstake")
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Directory *DirectorySession) Unstake() (*types.Transaction, error) {
	return _Directory.Contract.Unstake(&_Directory.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Directory *DirectoryTransactorSession) Unstake() (*types.Transaction, error) {
	return _Directory.Contract.Unstake(&_Directory.TransactOpts)
}

// UnstakeTo is a paid mutator transaction binding the contract method 0x8ee82c32.
//
// Solidity: function unstakeTo(address account) returns()
func (_Directory *DirectoryTransactor) UnstakeTo(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "unstakeTo", account)
}

// UnstakeTo is a paid mutator transaction binding the contract method 0x8ee82c32.
//
// Solidity: function unstakeTo(address account) returns()
func (_Directory *DirectorySession) UnstakeTo(account common.Address) (*types.Transaction, error) {
	return _Directory.Contract.UnstakeTo(&_Directory.TransactOpts, account)
}

// UnstakeTo is a paid mutator transaction binding the contract method 0x8ee82c32.
//
// Solidity: function unstakeTo(address account) returns()
func (_Directory *DirectoryTransactorSession) UnstakeTo(account common.Address) (*types.Transaction, error) {
	return _Directory.Contract.UnstakeTo(&_Directory.TransactOpts, account)
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
	return event, nil
}
