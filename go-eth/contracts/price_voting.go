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

// PriceVotingVote is an auto generated low-level Go binding around an user-defined struct.
type PriceVotingVote struct {
	Voter common.Address
	Price *big.Int
}

// PriceVotingABI is the input ABI used to generate the binding from.
const PriceVotingABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractStakingManager\",\"name\":\"stakingManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"sortedIndexes\",\"type\":\"uint256[]\"}],\"name\":\"validateSortedVotes\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structPriceVoting.Vote[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PriceVotingBin is the compiled bytecode used for deploying new contracts.
var PriceVotingBin = "0x608060405234801561001057600080fd5b506110b2806100206000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c80638da5cb5b11610076578063d8bff5a51161005b578063d8bff5a51461015f578063da58c7d91461018d578063f2fde38b146101a057600080fd5b80638da5cb5b14610127578063c4d66de81461014c57600080fd5b80633a1231e1116100a75780633a1231e1146100f75780633ccfd60b14610117578063715018a61461011f57600080fd5b80630121b93f146100c35780630dc96015146100d8575b600080fd5b6100d66100d1366004610f16565b6101b3565b005b6100e06102a4565b6040516100ee929190610f2e565b60405180910390f35b61010a610105366004610e56565b610479565b6040516100ee9190610fb1565b6100d661080e565b6100d661096b565b6033546001600160a01b03165b6040516001600160a01b0390911681526020016100ee565b6100d661015a366004610e33565b6109d1565b61017f61016d366004610e33565b60666020526000908152604090205481565b6040519081526020016100ee565b61013461019b366004610f16565b610ab5565b6100d66101ae366004610e33565b610adf565b6000811161022e5760405162461bcd60e51b815260206004820152602360248201527f566f74696e67207072696365206d75737420626520677265617465722074686160448201527f6e2030000000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b3360009081526066602052604090205461029257606780546001810182556000919091527f9787eeb91fe3101235e4a76063c7023ecb40f923f97916639c598592fa30d6ae01805473ffffffffffffffffffffffffffffffffffffffff1916331790555b33600090815260666020526040902055565b606080600060678054905067ffffffffffffffff8111156102d557634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156102fe578160200160208202803683370190505b5060675490915060009067ffffffffffffffff81111561032e57634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610357578160200160208202803683370190505b50905060005b60675481101561046f576067818154811061038857634e487b7160e01b600052603260045260246000fd5b9060005260206000200160009054906101000a90046001600160a01b03168382815181106103c657634e487b7160e01b600052603260045260246000fd5b60200260200101906001600160a01b031690816001600160a01b031681525050606660006067838154811061040b57634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101546001600160a01b03168352820192909252604001902054825183908390811061045257634e487b7160e01b600052603260045260246000fd5b60209081029190910101528061046781611020565b91505061035d565b5090939092509050565b6067548151606091146104f45760405162461bcd60e51b815260206004820152603160248201527f4e6f7420616c6c20766f7465727320776572652070726573656e7420696e207360448201527f6f7274656420766f7465722061727261790000000000000000000000000000006064820152608401610225565b60675460009067ffffffffffffffff81111561052057634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610549578160200160208202803683370190505b50905060008060678054905067ffffffffffffffff81111561057b57634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156105c057816020015b60408051808201909152600080825260208201528152602001906001900390816105995790505b50905060005b855181101561080557600060678783815181106105f357634e487b7160e01b600052603260045260246000fd5b60200260200101518154811061061957634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101546001600160a01b03168083526066909152604090912054909150848110156106915760405162461bcd60e51b815260206004820152601e60248201527f476976656e20766f7465206172726179206973206e6f7420736f7274656400006044820152606401610225565b6001600160a01b038216600090815260666020526040902054885190955086908990859081106106d157634e487b7160e01b600052603260045260246000fd5b6020026020010151815181106106f757634e487b7160e01b600052603260045260246000fd5b60200260200101511561074c5760405162461bcd60e51b815260206004820152600f60248201527f466f756e64206475706c696361746500000000000000000000000000000000006044820152606401610225565b60018689858151811061076f57634e487b7160e01b600052603260045260246000fd5b60200260200101518151811061079557634e487b7160e01b600052603260045260246000fd5b6020026020010190151590811515815250506040518060400160405280836001600160a01b03168152602001828152508484815181106107e557634e487b7160e01b600052603260045260246000fd5b6020026020010181905250505080806107fd90611020565b9150506105c6565b50949350505050565b3360009081526066602052604081208190555b60675481101561096857336001600160a01b03166067828154811061085657634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b03161415610956576067805461088190600190611009565b8154811061089f57634e487b7160e01b600052603260045260246000fd5b600091825260209091200154606780546001600160a01b0390921691839081106108d957634e487b7160e01b600052603260045260246000fd5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550606780548061092657634e487b7160e01b600052603160045260246000fd5b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff191690550190555b8061096081611020565b915050610821565b50565b6033546001600160a01b031633146109c55760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610225565b6109cf6000610bba565b565b600054610100900460ff16806109ea575060005460ff16155b610a4d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610225565b600054610100900460ff16158015610a6f576000805461ffff19166101011790555b610a77610c19565b6065805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0384161790558015610ab1576000805461ff00191690555b5050565b60678181548110610ac557600080fd5b6000918252602090912001546001600160a01b0316905081565b6033546001600160a01b03163314610b395760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610225565b6001600160a01b038116610bb55760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610225565b610968815b603380546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff1680610c32575060005460ff16155b610c955760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610225565b600054610100900460ff16158015610cb7576000805461ffff19166101011790555b610cbf610cdb565b610cc7610d8c565b8015610968576000805461ff001916905550565b600054610100900460ff1680610cf4575060005460ff16155b610d575760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610225565b600054610100900460ff16158015610cc7576000805461ffff19166101011790558015610968576000805461ff001916905550565b600054610100900460ff1680610da5575060005460ff16155b610e085760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610225565b600054610100900460ff16158015610e2a576000805461ffff19166101011790555b610cc733610bba565b600060208284031215610e44578081fd5b8135610e4f81611067565b9392505050565b60006020808385031215610e68578182fd5b823567ffffffffffffffff80821115610e7f578384fd5b818501915085601f830112610e92578384fd5b813581811115610ea457610ea4611051565b8060051b604051601f19603f83011681018181108582111715610ec957610ec9611051565b604052828152858101935084860182860187018a1015610ee7578788fd5b8795505b83861015610f09578035855260019590950194938601938601610eeb565b5098975050505050505050565b600060208284031215610f27578081fd5b5035919050565b604080825283519082018190526000906020906060840190828701845b82811015610f705781516001600160a01b031684529284019290840190600101610f4b565b50505083810382850152845180825285830191830190845b81811015610fa457835183529284019291840191600101610f88565b5090979650505050505050565b602080825282518282018190526000919060409081850190868401855b82811015610ffc57815180516001600160a01b03168552860151868501529284019290850190600101610fce565b5091979650505050505050565b60008282101561101b5761101b61103b565b500390565b60006000198214156110345761103461103b565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461096857600080fdfea26469706673582212203e9a947154d46444db6e45e99287d8638beab945de398a9156e2295197aa287264736f6c63430008040033"

// DeployPriceVoting deploys a new Ethereum contract, binding an instance of PriceVoting to it.
func DeployPriceVoting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PriceVoting, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceVotingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PriceVotingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PriceVoting{PriceVotingCaller: PriceVotingCaller{contract: contract}, PriceVotingTransactor: PriceVotingTransactor{contract: contract}, PriceVotingFilterer: PriceVotingFilterer{contract: contract}}, nil
}

// PriceVoting is an auto generated Go binding around an Ethereum contract.
type PriceVoting struct {
	PriceVotingCaller     // Read-only binding to the contract
	PriceVotingTransactor // Write-only binding to the contract
	PriceVotingFilterer   // Log filterer for contract events
}

// PriceVotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceVotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceVotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceVotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceVotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceVotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceVotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceVotingSession struct {
	Contract     *PriceVoting      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceVotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceVotingCallerSession struct {
	Contract *PriceVotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PriceVotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceVotingTransactorSession struct {
	Contract     *PriceVotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PriceVotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceVotingRaw struct {
	Contract *PriceVoting // Generic contract binding to access the raw methods on
}

// PriceVotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceVotingCallerRaw struct {
	Contract *PriceVotingCaller // Generic read-only contract binding to access the raw methods on
}

// PriceVotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceVotingTransactorRaw struct {
	Contract *PriceVotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceVoting creates a new instance of PriceVoting, bound to a specific deployed contract.
func NewPriceVoting(address common.Address, backend bind.ContractBackend) (*PriceVoting, error) {
	contract, err := bindPriceVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceVoting{PriceVotingCaller: PriceVotingCaller{contract: contract}, PriceVotingTransactor: PriceVotingTransactor{contract: contract}, PriceVotingFilterer: PriceVotingFilterer{contract: contract}}, nil
}

// NewPriceVotingCaller creates a new read-only instance of PriceVoting, bound to a specific deployed contract.
func NewPriceVotingCaller(address common.Address, caller bind.ContractCaller) (*PriceVotingCaller, error) {
	contract, err := bindPriceVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceVotingCaller{contract: contract}, nil
}

// NewPriceVotingTransactor creates a new write-only instance of PriceVoting, bound to a specific deployed contract.
func NewPriceVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceVotingTransactor, error) {
	contract, err := bindPriceVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceVotingTransactor{contract: contract}, nil
}

// NewPriceVotingFilterer creates a new log filterer instance of PriceVoting, bound to a specific deployed contract.
func NewPriceVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceVotingFilterer, error) {
	contract, err := bindPriceVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceVotingFilterer{contract: contract}, nil
}

// bindPriceVoting binds a generic wrapper to an already deployed contract.
func bindPriceVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceVotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceVoting *PriceVotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceVoting.Contract.PriceVotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceVoting *PriceVotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceVoting.Contract.PriceVotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceVoting *PriceVotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceVoting.Contract.PriceVotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceVoting *PriceVotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceVoting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceVoting *PriceVotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceVoting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceVoting *PriceVotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceVoting.Contract.contract.Transact(opts, method, params...)
}

// GetVotes is a free data retrieval call binding the contract method 0x0dc96015.
//
// Solidity: function getVotes() view returns(address[], uint256[])
func (_PriceVoting *PriceVotingCaller) GetVotes(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _PriceVoting.contract.Call(opts, &out, "getVotes")

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetVotes is a free data retrieval call binding the contract method 0x0dc96015.
//
// Solidity: function getVotes() view returns(address[], uint256[])
func (_PriceVoting *PriceVotingSession) GetVotes() ([]common.Address, []*big.Int, error) {
	return _PriceVoting.Contract.GetVotes(&_PriceVoting.CallOpts)
}

// GetVotes is a free data retrieval call binding the contract method 0x0dc96015.
//
// Solidity: function getVotes() view returns(address[], uint256[])
func (_PriceVoting *PriceVotingCallerSession) GetVotes() ([]common.Address, []*big.Int, error) {
	return _PriceVoting.Contract.GetVotes(&_PriceVoting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceVoting *PriceVotingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceVoting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceVoting *PriceVotingSession) Owner() (common.Address, error) {
	return _PriceVoting.Contract.Owner(&_PriceVoting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceVoting *PriceVotingCallerSession) Owner() (common.Address, error) {
	return _PriceVoting.Contract.Owner(&_PriceVoting.CallOpts)
}

// ValidateSortedVotes is a free data retrieval call binding the contract method 0x3a1231e1.
//
// Solidity: function validateSortedVotes(uint256[] sortedIndexes) view returns((address,uint256)[])
func (_PriceVoting *PriceVotingCaller) ValidateSortedVotes(opts *bind.CallOpts, sortedIndexes []*big.Int) ([]PriceVotingVote, error) {
	var out []interface{}
	err := _PriceVoting.contract.Call(opts, &out, "validateSortedVotes", sortedIndexes)

	if err != nil {
		return *new([]PriceVotingVote), err
	}

	out0 := *abi.ConvertType(out[0], new([]PriceVotingVote)).(*[]PriceVotingVote)

	return out0, err

}

// ValidateSortedVotes is a free data retrieval call binding the contract method 0x3a1231e1.
//
// Solidity: function validateSortedVotes(uint256[] sortedIndexes) view returns((address,uint256)[])
func (_PriceVoting *PriceVotingSession) ValidateSortedVotes(sortedIndexes []*big.Int) ([]PriceVotingVote, error) {
	return _PriceVoting.Contract.ValidateSortedVotes(&_PriceVoting.CallOpts, sortedIndexes)
}

// ValidateSortedVotes is a free data retrieval call binding the contract method 0x3a1231e1.
//
// Solidity: function validateSortedVotes(uint256[] sortedIndexes) view returns((address,uint256)[])
func (_PriceVoting *PriceVotingCallerSession) ValidateSortedVotes(sortedIndexes []*big.Int) ([]PriceVotingVote, error) {
	return _PriceVoting.Contract.ValidateSortedVotes(&_PriceVoting.CallOpts, sortedIndexes)
}

// Voters is a free data retrieval call binding the contract method 0xda58c7d9.
//
// Solidity: function voters(uint256 ) view returns(address)
func (_PriceVoting *PriceVotingCaller) Voters(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PriceVoting.contract.Call(opts, &out, "voters", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Voters is a free data retrieval call binding the contract method 0xda58c7d9.
//
// Solidity: function voters(uint256 ) view returns(address)
func (_PriceVoting *PriceVotingSession) Voters(arg0 *big.Int) (common.Address, error) {
	return _PriceVoting.Contract.Voters(&_PriceVoting.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xda58c7d9.
//
// Solidity: function voters(uint256 ) view returns(address)
func (_PriceVoting *PriceVotingCallerSession) Voters(arg0 *big.Int) (common.Address, error) {
	return _PriceVoting.Contract.Voters(&_PriceVoting.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(uint256)
func (_PriceVoting *PriceVotingCaller) Votes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceVoting.contract.Call(opts, &out, "votes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(uint256)
func (_PriceVoting *PriceVotingSession) Votes(arg0 common.Address) (*big.Int, error) {
	return _PriceVoting.Contract.Votes(&_PriceVoting.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(uint256)
func (_PriceVoting *PriceVotingCallerSession) Votes(arg0 common.Address) (*big.Int, error) {
	return _PriceVoting.Contract.Votes(&_PriceVoting.CallOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address stakingManager) returns()
func (_PriceVoting *PriceVotingTransactor) Initialize(opts *bind.TransactOpts, stakingManager common.Address) (*types.Transaction, error) {
	return _PriceVoting.contract.Transact(opts, "initialize", stakingManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address stakingManager) returns()
func (_PriceVoting *PriceVotingSession) Initialize(stakingManager common.Address) (*types.Transaction, error) {
	return _PriceVoting.Contract.Initialize(&_PriceVoting.TransactOpts, stakingManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address stakingManager) returns()
func (_PriceVoting *PriceVotingTransactorSession) Initialize(stakingManager common.Address) (*types.Transaction, error) {
	return _PriceVoting.Contract.Initialize(&_PriceVoting.TransactOpts, stakingManager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceVoting *PriceVotingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceVoting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceVoting *PriceVotingSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceVoting.Contract.RenounceOwnership(&_PriceVoting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceVoting *PriceVotingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceVoting.Contract.RenounceOwnership(&_PriceVoting.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceVoting *PriceVotingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PriceVoting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceVoting *PriceVotingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceVoting.Contract.TransferOwnership(&_PriceVoting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceVoting *PriceVotingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceVoting.Contract.TransferOwnership(&_PriceVoting.TransactOpts, newOwner)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 price) returns()
func (_PriceVoting *PriceVotingTransactor) Vote(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _PriceVoting.contract.Transact(opts, "vote", price)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 price) returns()
func (_PriceVoting *PriceVotingSession) Vote(price *big.Int) (*types.Transaction, error) {
	return _PriceVoting.Contract.Vote(&_PriceVoting.TransactOpts, price)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 price) returns()
func (_PriceVoting *PriceVotingTransactorSession) Vote(price *big.Int) (*types.Transaction, error) {
	return _PriceVoting.Contract.Vote(&_PriceVoting.TransactOpts, price)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PriceVoting *PriceVotingTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceVoting.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PriceVoting *PriceVotingSession) Withdraw() (*types.Transaction, error) {
	return _PriceVoting.Contract.Withdraw(&_PriceVoting.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PriceVoting *PriceVotingTransactorSession) Withdraw() (*types.Transaction, error) {
	return _PriceVoting.Contract.Withdraw(&_PriceVoting.TransactOpts)
}

// PriceVotingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PriceVoting contract.
type PriceVotingOwnershipTransferredIterator struct {
	Event *PriceVotingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PriceVotingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceVotingOwnershipTransferred)
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
		it.Event = new(PriceVotingOwnershipTransferred)
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
func (it *PriceVotingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceVotingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceVotingOwnershipTransferred represents a OwnershipTransferred event raised by the PriceVoting contract.
type PriceVotingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceVoting *PriceVotingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PriceVotingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceVoting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PriceVotingOwnershipTransferredIterator{contract: _PriceVoting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceVoting *PriceVotingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PriceVotingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceVoting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceVotingOwnershipTransferred)
				if err := _PriceVoting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PriceVoting *PriceVotingFilterer) ParseOwnershipTransferred(log types.Log) (*PriceVotingOwnershipTransferred, error) {
	event := new(PriceVotingOwnershipTransferred)
	if err := _PriceVoting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
