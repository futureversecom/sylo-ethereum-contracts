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

// SyloTokenABI is the input ABI used to generate the binding from.
const SyloTokenABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SyloTokenBin is the compiled bytecode used for deploying new contracts.
var SyloTokenBin = "0x60806040523480156200001157600080fd5b506040518060400160405280600481526020016353796c6f60e01b8152506040518060400160405280600481526020016353594c4f60e01b81525081600390805190602001906200006492919062000230565b5080516200007a90600490602084019062000230565b50506005805460ff1916601217905550620000ab336b204fce5e3e250261100000006001600160e01b03620000b116565b620002d5565b6001600160a01b0382166200010d576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b62000124600083836001600160e01b03620001c916565b6200014081600254620001ce60201b620008931790919060201c565b6002556001600160a01b038216600090815260208181526040909120546200017391839062000893620001ce821b17901c565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b505050565b60008282018381101562000229576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200027357805160ff1916838001178555620002a3565b82800160010185558215620002a3579182015b82811115620002a357825182559160200191906001019062000286565b50620002b1929150620002b5565b5090565b620002d291905b80821115620002b15760008155600101620002bc565b90565b610a3080620002e56000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80633950935111610081578063a457c2d71161005b578063a457c2d714610253578063a9059cbb1461027f578063dd62ed3e146102ab576100c9565b806339509351146101f957806370a082311461022557806395d89b411461024b576100c9565b806318160ddd116100b257806318160ddd1461018b57806323b872dd146101a5578063313ce567146101db576100c9565b806306fdde03146100ce578063095ea7b31461014b575b600080fd5b6100d66102d9565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101105781810151838201526020016100f8565b50505050905090810190601f16801561013d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101776004803603604081101561016157600080fd5b506001600160a01b03813516906020013561036f565b604080519115158252519081900360200190f35b61019361038c565b60408051918252519081900360200190f35b610177600480360360608110156101bb57600080fd5b506001600160a01b03813581169160208101359091169060400135610392565b6101e361041f565b6040805160ff9092168252519081900360200190f35b6101776004803603604081101561020f57600080fd5b506001600160a01b038135169060200135610428565b6101936004803603602081101561023b57600080fd5b50356001600160a01b031661047c565b6100d6610497565b6101776004803603604081101561026957600080fd5b506001600160a01b0381351690602001356104f8565b6101776004803603604081101561029557600080fd5b506001600160a01b038135169060200135610566565b610193600480360360408110156102c157600080fd5b506001600160a01b038135811691602001351661057a565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103655780601f1061033a57610100808354040283529160200191610365565b820191906000526020600020905b81548152906001019060200180831161034857829003601f168201915b5050505050905090565b600061038361037c6105a5565b84846105a9565b50600192915050565b60025490565b600061039f848484610695565b610415846103ab6105a5565b61041085604051806060016040528060288152602001610965602891396001600160a01b038a166000908152600160205260408120906103e96105a5565b6001600160a01b03168152602081019190915260400160002054919063ffffffff6107fc16565b6105a9565b5060019392505050565b60055460ff1690565b60006103836104356105a5565b8461041085600160006104466105a5565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff61089316565b6001600160a01b031660009081526020819052604090205490565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103655780601f1061033a57610100808354040283529160200191610365565b60006103836105056105a5565b84610410856040518060600160405280602581526020016109d6602591396001600061052f6105a5565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff6107fc16565b60006103836105736105a5565b8484610695565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b3390565b6001600160a01b0383166105ee5760405162461bcd60e51b81526004018080602001828103825260248152602001806109b26024913960400191505060405180910390fd5b6001600160a01b0382166106335760405162461bcd60e51b815260040180806020018281038252602281526020018061091d6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166106da5760405162461bcd60e51b815260040180806020018281038252602581526020018061098d6025913960400191505060405180910390fd5b6001600160a01b03821661071f5760405162461bcd60e51b81526004018080602001828103825260238152602001806108fa6023913960400191505060405180910390fd5b61072a8383836108f4565b61076d8160405180606001604052806026815260200161093f602691396001600160a01b038616600090815260208190526040902054919063ffffffff6107fc16565b6001600160a01b0380851660009081526020819052604080822093909355908416815220546107a2908263ffffffff61089316565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000818484111561088b5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610850578181015183820152602001610838565b50505050905090810190601f16801561087d5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6000828201838110156108ed576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b50505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa26469706673582212206c5b25587e478364dcd93fa1f5e68111bd25dcb9fb7be375c63e72a79c82cc0d64736f6c63430006040033"

// DeploySyloToken deploys a new Ethereum contract, binding an instance of SyloToken to it.
func DeploySyloToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SyloToken, error) {
	parsed, err := abi.JSON(strings.NewReader(SyloTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SyloTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SyloToken{SyloTokenCaller: SyloTokenCaller{contract: contract}, SyloTokenTransactor: SyloTokenTransactor{contract: contract}, SyloTokenFilterer: SyloTokenFilterer{contract: contract}}, nil
}

// SyloToken is an auto generated Go binding around an Ethereum contract.
type SyloToken struct {
	SyloTokenCaller     // Read-only binding to the contract
	SyloTokenTransactor // Write-only binding to the contract
	SyloTokenFilterer   // Log filterer for contract events
}

// SyloTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SyloTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyloTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SyloTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyloTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SyloTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyloTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SyloTokenSession struct {
	Contract     *SyloToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SyloTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SyloTokenCallerSession struct {
	Contract *SyloTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SyloTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SyloTokenTransactorSession struct {
	Contract     *SyloTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SyloTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SyloTokenRaw struct {
	Contract *SyloToken // Generic contract binding to access the raw methods on
}

// SyloTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SyloTokenCallerRaw struct {
	Contract *SyloTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SyloTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SyloTokenTransactorRaw struct {
	Contract *SyloTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSyloToken creates a new instance of SyloToken, bound to a specific deployed contract.
func NewSyloToken(address common.Address, backend bind.ContractBackend) (*SyloToken, error) {
	contract, err := bindSyloToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SyloToken{SyloTokenCaller: SyloTokenCaller{contract: contract}, SyloTokenTransactor: SyloTokenTransactor{contract: contract}, SyloTokenFilterer: SyloTokenFilterer{contract: contract}}, nil
}

// NewSyloTokenCaller creates a new read-only instance of SyloToken, bound to a specific deployed contract.
func NewSyloTokenCaller(address common.Address, caller bind.ContractCaller) (*SyloTokenCaller, error) {
	contract, err := bindSyloToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SyloTokenCaller{contract: contract}, nil
}

// NewSyloTokenTransactor creates a new write-only instance of SyloToken, bound to a specific deployed contract.
func NewSyloTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SyloTokenTransactor, error) {
	contract, err := bindSyloToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SyloTokenTransactor{contract: contract}, nil
}

// NewSyloTokenFilterer creates a new log filterer instance of SyloToken, bound to a specific deployed contract.
func NewSyloTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*SyloTokenFilterer, error) {
	contract, err := bindSyloToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SyloTokenFilterer{contract: contract}, nil
}

// bindSyloToken binds a generic wrapper to an already deployed contract.
func bindSyloToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SyloTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SyloToken *SyloTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SyloToken.Contract.SyloTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SyloToken *SyloTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyloToken.Contract.SyloTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SyloToken *SyloTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SyloToken.Contract.SyloTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SyloToken *SyloTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SyloToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SyloToken *SyloTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyloToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SyloToken *SyloTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SyloToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SyloToken *SyloTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SyloToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SyloToken *SyloTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SyloToken.Contract.Allowance(&_SyloToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SyloToken *SyloTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SyloToken.Contract.Allowance(&_SyloToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SyloToken *SyloTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SyloToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SyloToken *SyloTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SyloToken.Contract.BalanceOf(&_SyloToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SyloToken *SyloTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SyloToken.Contract.BalanceOf(&_SyloToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SyloToken *SyloTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SyloToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SyloToken *SyloTokenSession) Decimals() (uint8, error) {
	return _SyloToken.Contract.Decimals(&_SyloToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SyloToken *SyloTokenCallerSession) Decimals() (uint8, error) {
	return _SyloToken.Contract.Decimals(&_SyloToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SyloToken *SyloTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SyloToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SyloToken *SyloTokenSession) Name() (string, error) {
	return _SyloToken.Contract.Name(&_SyloToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SyloToken *SyloTokenCallerSession) Name() (string, error) {
	return _SyloToken.Contract.Name(&_SyloToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SyloToken *SyloTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SyloToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SyloToken *SyloTokenSession) Symbol() (string, error) {
	return _SyloToken.Contract.Symbol(&_SyloToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SyloToken *SyloTokenCallerSession) Symbol() (string, error) {
	return _SyloToken.Contract.Symbol(&_SyloToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SyloToken *SyloTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyloToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SyloToken *SyloTokenSession) TotalSupply() (*big.Int, error) {
	return _SyloToken.Contract.TotalSupply(&_SyloToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SyloToken *SyloTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _SyloToken.Contract.TotalSupply(&_SyloToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SyloToken *SyloTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SyloToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SyloToken *SyloTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SyloToken.Contract.Approve(&_SyloToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SyloToken *SyloTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SyloToken.Contract.Approve(&_SyloToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SyloToken *SyloTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SyloToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SyloToken *SyloTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SyloToken.Contract.DecreaseAllowance(&_SyloToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SyloToken *SyloTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SyloToken.Contract.DecreaseAllowance(&_SyloToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SyloToken *SyloTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SyloToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SyloToken *SyloTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SyloToken.Contract.IncreaseAllowance(&_SyloToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SyloToken *SyloTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SyloToken.Contract.IncreaseAllowance(&_SyloToken.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SyloToken *SyloTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SyloToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SyloToken *SyloTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SyloToken.Contract.Transfer(&_SyloToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SyloToken *SyloTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SyloToken.Contract.Transfer(&_SyloToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SyloToken *SyloTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SyloToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SyloToken *SyloTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SyloToken.Contract.TransferFrom(&_SyloToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SyloToken *SyloTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SyloToken.Contract.TransferFrom(&_SyloToken.TransactOpts, sender, recipient, amount)
}

// SyloTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SyloToken contract.
type SyloTokenApprovalIterator struct {
	Event *SyloTokenApproval // Event containing the contract specifics and raw log

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
func (it *SyloTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyloTokenApproval)
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
		it.Event = new(SyloTokenApproval)
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
func (it *SyloTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyloTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyloTokenApproval represents a Approval event raised by the SyloToken contract.
type SyloTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SyloToken *SyloTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SyloTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SyloToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SyloTokenApprovalIterator{contract: _SyloToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SyloToken *SyloTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SyloTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SyloToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyloTokenApproval)
				if err := _SyloToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SyloToken *SyloTokenFilterer) ParseApproval(log types.Log) (*SyloTokenApproval, error) {
	event := new(SyloTokenApproval)
	if err := _SyloToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyloTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SyloToken contract.
type SyloTokenTransferIterator struct {
	Event *SyloTokenTransfer // Event containing the contract specifics and raw log

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
func (it *SyloTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyloTokenTransfer)
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
		it.Event = new(SyloTokenTransfer)
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
func (it *SyloTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyloTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyloTokenTransfer represents a Transfer event raised by the SyloToken contract.
type SyloTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SyloToken *SyloTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SyloTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SyloToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SyloTokenTransferIterator{contract: _SyloToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SyloToken *SyloTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SyloTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SyloToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyloTokenTransfer)
				if err := _SyloToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SyloToken *SyloTokenFilterer) ParseTransfer(log types.Log) (*SyloTokenTransfer, error) {
	event := new(SyloTokenTransfer)
	if err := _SyloToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
