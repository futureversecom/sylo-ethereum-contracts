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
const PriceManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"currentServicePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentUpperPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractStakingManager\",\"name\":\"stakingManager\",\"type\":\"address\"},{\"internalType\":\"contractPriceVoting\",\"name\":\"voting\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structPriceVoting.Vote[]\",\"name\":\"sortedVotes\",\"type\":\"tuple[]\"}],\"name\":\"calculatePrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"servicePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PriceManagerBin is the compiled bytecode used for deploying new contracts.
var PriceManagerBin = "0x60806040526000606755600060685534801561001a57600080fd5b5061121f8061002a6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638da5cb5b1161005b5780638da5cb5b146100bb578063985371a3146100d6578063cdfa9e93146100df578063f2fde38b1461010757600080fd5b80630b5b820714610082578063485cc9551461009e578063715018a6146100b3575b600080fd5b61008b60675481565b6040519081526020015b60405180910390f35b6100b16100ac366004610fe2565b61011a565b005b6100b1610219565b6033546040516001600160a01b039091168152602001610095565b61008b60685481565b6100f26100ed366004610f25565b6102ca565b60408051928352602083019190915201610095565b6100b1610115366004610e3c565b6105a0565b600054610100900460ff1680610133575060005460ff16155b61019b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b600054610100900460ff161580156101bd576000805461ffff19166101011790555b6101c56106df565b606580546001600160a01b0380861673ffffffffffffffffffffffffffffffffffffffff199283161790925560668054928516929091169190911790558015610214576000805461ff00191690555b505050565b6033546001600160a01b031633146102735760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610192565b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36033805473ffffffffffffffffffffffffffffffffffffffff19169055565b600080336001600160a01b03166102e96033546001600160a01b031690565b6001600160a01b03161461033f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610192565b8251156103d1576066546040517f7c5d14590000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690637c5d14599061038f908690600401611032565b60006040518083038186803b1580156103a757600080fd5b505afa1580156103bb573d6000803e3d6000fd5b505050506103c8836107a2565b91509150915091565b600080606660009054906101000a90046001600160a01b03166001600160a01b031663efb50efb6040518163ffffffff1660e01b815260040160006040518083038186803b15801561042257600080fd5b505afa158015610436573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261045e9190810190610e5f565b915091506000825167ffffffffffffffff81111561048c57634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156104d157816020015b60408051808201909152600080825260208201528152602001906001900390816104aa5790505b50905060005b835181101561058a57604051806040016040528085838151811061050b57634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b0316815260200184838151811061054157634e487b7160e01b600052603260045260246000fd5b602002602001015181525082828151811061056c57634e487b7160e01b600052603260045260246000fd5b602002602001018190525080806105829061118d565b9150506104d7565b50610594816107a2565b94509450505050915091565b6033546001600160a01b031633146105fa5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610192565b6001600160a01b0381166106765760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610192565b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a36033805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b600054610100900460ff16806106f8575060005460ff16155b61075b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610192565b600054610100900460ff1615801561077d576000805461ffff19166101011790555b610785610c1a565b61078d610ccb565b801561079f576000805461ff00191690555b50565b600080336001600160a01b03166107c16033546001600160a01b031690565b6001600160a01b0316146108175760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610192565b606554604080517f8b0e9f3f00000000000000000000000000000000000000000000000000000000815290516000926001600160a01b031691638b0e9f3f916004808301926020929190829003018186803b15801561087557600080fd5b505afa158015610889573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ad919061101a565b9050600060646108be836019611140565b6108c89190611120565b9050600060646108d984605a611140565b6108e39190611120565b90506000805b8751811015610a6b5760655488516000916001600160a01b03169063df349ed5908b908590811061092a57634e487b7160e01b600052603260045260246000fd5b6020908102919091010151516040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b03909116600482015260240160206040518083038186803b15801561098d57600080fd5b505afa1580156109a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109c5919061101a565b9050806109d25750610a59565b8882815181106109f257634e487b7160e01b600052603260045260246000fd5b60200260200101516020015160001415610a0c5750610a59565b610a168184611108565b9250848310610a5757888281518110610a3f57634e487b7160e01b600052603260045260246000fd5b60200260200101516020015160678190555050610a6b565b505b80610a638161118d565b9150506108e9565b5050855183905b8015610c06576065546000906001600160a01b031663df349ed58a610a9860018661115f565b81518110610ab657634e487b7160e01b600052603260045260246000fd5b6020908102919091010151516040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b03909116600482015260240160206040518083038186803b158015610b1957600080fd5b505afa158015610b2d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b51919061101a565b905080610b5e5750610bf4565b88610b6a60018461115f565b81518110610b8857634e487b7160e01b600052603260045260246000fd5b60200260200101516020015160001415610ba25750610bf4565b610bac818461115f565b9250838311610bf25788610bc160018461115f565b81518110610bdf57634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516068819055505b505b80610bfe81611176565b915050610a72565b506067546068549550955050505050915091565b600054610100900460ff1680610c33575060005460ff16155b610c965760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610192565b600054610100900460ff1615801561078d576000805461ffff1916610101179055801561079f576000805461ff001916905550565b600054610100900460ff1680610ce4575060005460ff16155b610d475760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610192565b600054610100900460ff16158015610d69576000805461ffff19166101011790555b6033805473ffffffffffffffffffffffffffffffffffffffff19163390811790915560405181906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350801561079f576000805461ff001916905550565b600082601f830112610ddd578081fd5b81516020610df2610ded836110e4565b6110b3565b80838252828201915082860187848660051b8901011115610e11578586fd5b855b85811015610e2f57815184529284019290840190600101610e13565b5090979650505050505050565b600060208284031215610e4d578081fd5b8135610e58816111d4565b9392505050565b60008060408385031215610e71578081fd5b825167ffffffffffffffff80821115610e88578283fd5b818501915085601f830112610e9b578283fd5b81516020610eab610ded836110e4565b8083825282820191508286018a848660051b8901011115610eca578788fd5b8796505b84871015610ef5578051610ee1816111d4565b835260019690960195918301918301610ece565b5091880151919650909350505080821115610f0e578283fd5b50610f1b85828601610dcd565b9150509250929050565b60006020808385031215610f37578182fd5b823567ffffffffffffffff811115610f4d578283fd5b8301601f81018513610f5d578283fd5b8035610f6b610ded826110e4565b80828252848201915084840188868560061b8701011115610f8a578687fd5b8694505b83851015610fd657604080828b031215610fa6578788fd5b610fae61108a565b8235610fb9816111d4565b815282880135888201528452600195909501949286019201610f8e565b50979650505050505050565b60008060408385031215610ff4578182fd5b8235610fff816111d4565b9150602083013561100f816111d4565b809150509250929050565b60006020828403121561102b578081fd5b5051919050565b602080825282518282018190526000919060409081850190868401855b8281101561107d57815180516001600160a01b0316855286015186850152928401929085019060010161104f565b5091979650505050505050565b6040805190810167ffffffffffffffff811182821017156110ad576110ad6111be565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156110dc576110dc6111be565b604052919050565b600067ffffffffffffffff8211156110fe576110fe6111be565b5060051b60200190565b6000821982111561111b5761111b6111a8565b500190565b60008261113b57634e487b7160e01b81526012600452602481fd5b500490565b600081600019048311821515161561115a5761115a6111a8565b500290565b600082821015611171576111716111a8565b500390565b600081611185576111856111a8565b506000190190565b60006000198214156111a1576111a16111a8565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461079f57600080fdfea2646970667358221220ad6fd8f89d22375d4c4bd2ffde07460d168b412423835b55cb08f0251769118b64736f6c63430008040033"

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

// CalculatePrices is a paid mutator transaction binding the contract method 0xcdfa9e93.
//
// Solidity: function calculatePrices((address,uint256)[] sortedVotes) returns(uint256 servicePrice, uint256 upperPrice)
func (_PriceManager *PriceManagerTransactor) CalculatePrices(opts *bind.TransactOpts, sortedVotes []PriceVotingVote) (*types.Transaction, error) {
	return _PriceManager.contract.Transact(opts, "calculatePrices", sortedVotes)
}

// CalculatePrices is a paid mutator transaction binding the contract method 0xcdfa9e93.
//
// Solidity: function calculatePrices((address,uint256)[] sortedVotes) returns(uint256 servicePrice, uint256 upperPrice)
func (_PriceManager *PriceManagerSession) CalculatePrices(sortedVotes []PriceVotingVote) (*types.Transaction, error) {
	return _PriceManager.Contract.CalculatePrices(&_PriceManager.TransactOpts, sortedVotes)
}

// CalculatePrices is a paid mutator transaction binding the contract method 0xcdfa9e93.
//
// Solidity: function calculatePrices((address,uint256)[] sortedVotes) returns(uint256 servicePrice, uint256 upperPrice)
func (_PriceManager *PriceManagerTransactorSession) CalculatePrices(sortedVotes []PriceVotingVote) (*types.Transaction, error) {
	return _PriceManager.Contract.CalculatePrices(&_PriceManager.TransactOpts, sortedVotes)
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
