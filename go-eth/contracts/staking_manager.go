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

// StakingManagerStakeOperation is an auto generated low-level Go binding around an user-defined struct.
type StakingManagerStakeOperation struct {
	Type   uint8
	Amount *big.Int
	Block  *big.Int
}

// StakingManagerABI is the input ABI used to generate the binding from.
const StakingManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"unlockings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unlockDuration\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newUnlockDuration\",\"type\":\"uint256\"}],\"name\":\"setUnlockDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"unlockStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"cancelUnlocking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getStakeOperations\",\"outputs\":[{\"components\":[{\"internalType\":\"enumStakingManager.StakeOperationType\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"internalType\":\"structStakingManager.StakeOperation[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getStakeeTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getStakerAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StakingManagerBin is the compiled bytecode used for deploying new contracts.
var StakingManagerBin = "0x608060405234801561001057600080fd5b50611302806100206000396000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c80638da5cb5b116100b2578063bc29855311610081578063cd6dc68711610066578063cd6dc687146102b1578063dd900769146102c4578063f2fde38b146102d757600080fd5b8063bc29855314610246578063c23a5cea1461029e57600080fd5b80638da5cb5b146101c95780639341a536146101e4578063a859f17214610220578063b91cefd21461023357600080fd5b8063715018a6116100ee578063715018a6146101a75780637bc74225146101af5780638a1fcd60146101b75780638b0e9f3f146101c057600080fd5b806323314c6c146101205780632d49aa1c1461013557806364084d4e146101485780636959c1c514610187575b600080fd5b61013361012e3660046111be565b6102ea565b005b6101336101433660046111be565b6103ef565b6101746101563660046110c0565b6001600160a01b031660009081526066602052604090206001015490565b6040519081526020015b60405180910390f35b61019a6101953660046110e3565b6104a1565b60405161017e91906111e2565b610133610580565b606754610174565b61017460685481565b61017460675481565b6033546040516001600160a01b03909116815260200161017e565b61020b6101f236600461117b565b6069602052600090815260409020805460019091015482565b6040805192835260208301919091520161017e565b61017461022e3660046111be565b6105e6565b61017461024136600461111b565b6107f1565b6101746102543660046110e3565b6040516bffffffffffffffffffffffff19606084811b8216602084015283901b16603482015260009060480160405160208183030381529060405280519060200120905092915050565b6101336102ac3660046110c0565b6108e2565b6101336102bf366004611193565b610aa4565b6101336102d236600461117b565b610b98565b6101336102e53660046110c0565b610bf7565b600061033b33836040516bffffffffffffffffffffffff19606084811b8216602084015283901b16603482015260009060480160405160208183030381529060405280519060200120905092915050565b6000818152606960205260409020805491925090841415610370576000828152606960205260408120818155600101556103df565b805484106103c55760405162461bcd60e51b815260206004820152601e60248201527f556e6c6f636b2068617320696e73756666696369656e7420616d6f756e74000060448201526064015b60405180910390fd5b838160000160008282546103d9919061126f565b90915550505b6103e98484610cd9565b50505050565b6103f98282610cd9565b6065546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018490526001600160a01b03909116906323b872dd90606401602060405180830381600087803b15801561046457600080fd5b505af1158015610478573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061049c919061115b565b505050565b6001600160a01b0380821660009081526066602090815260408083209386168352928152828220805484518184028101840190955280855260609493919290919084015b82821015610574576000848152602090206040805160608101909152600384029091018054829060ff16600181111561052e57634e487b7160e01b600052602160045260246000fd5b600181111561054d57634e487b7160e01b600052602160045260246000fd5b815260200160018201548152602001600282015481525050815260200190600101906104e5565b50505050905092915050565b6033546001600160a01b031633146105da5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103bc565b6105e46000610e3c565b565b6001600160a01b03811660009081526066602052604081208161060a3385436107f1565b90506000811161065c5760405162461bcd60e51b815260206004820152601260248201527f4e6f7468696e6720746f20756e7374616b65000000000000000000000000000060448201526064016103bc565b848110156106ac5760405162461bcd60e51b815260206004820152601e60248201527f43616e6e6f7420756e6c6f636b206d6f7265207468616e207374616b6564000060448201526064016103bc565b33600090815260208381526040808320815160608101835260018082528185018b90524393820193909352815480840183559185529290932082516003909402018054929390929091839160ff191690838181111561071b57634e487b7160e01b600052602160045260246000fd5b0217905550602082015181600101556040820151816002015550508482600101600082825461074a919061126f565b90915550506040805133606090811b6bffffffffffffffffffffffff199081166020808501919091529188901b16603483015282516028818403018152604890920190925280519101206000906000818152606960205260408120606854929350916107b69043611257565b905080826001015410156107cc57600182018190555b878260000160008282546107e09190611257565b909155509098975050505050505050565b6000806107fe85856104a1565b90506000805b82518110156108d857600083828151811061082f57634e487b7160e01b600052603260045260246000fd5b60200260200101519050858160400151111561084b57506108d8565b60008151600181111561086e57634e487b7160e01b600052602160045260246000fd5b141561088a5760208101516108839084611257565b92506108c5565b6001815160018111156108ad57634e487b7160e01b600052602160045260246000fd5b14156108c55760208101516108c2908461126f565b92505b50806108d081611286565b915050610804565b5095945050505050565b600061093333836040516bffffffffffffffffffffffff19606084811b8216602084015283901b16603482015260009060480160405160208183030381529060405280519060200120905092915050565b600081815260696020526040902060018101549192509043116109985760405162461bcd60e51b815260206004820152601660248201527f5374616b65206e6f742079657420756e6c6f636b65640000000000000000000060448201526064016103bc565b80546109e65760405162461bcd60e51b815260206004820152601560248201527f4e6f20616d6f756e7420746f207769746864726177000000000000000000000060448201526064016103bc565b80546000838152606960205260408082208281556001019190915560655490517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018390526001600160a01b039091169063a9059cbb90604401602060405180830381600087803b158015610a6557600080fd5b505af1158015610a79573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a9d919061115b565b5050505050565b600054610100900460ff1680610abd575060005460ff16155b610b205760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103bc565b600054610100900460ff16158015610b42576000805461ffff19166101011790555b610b4a610ea6565b606580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0385161790556068829055801561049c576000805461ff0019169055505050565b6033546001600160a01b03163314610bf25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103bc565b606855565b6033546001600160a01b03163314610c515760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103bc565b6001600160a01b038116610ccd5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016103bc565b610cd681610e3c565b50565b6001600160a01b038116610d2f5760405162461bcd60e51b815260206004820152600f60248201527f41646472657373206973206e756c6c000000000000000000000000000000000060448201526064016103bc565b81610d7c5760405162461bcd60e51b815260206004820152601460248201527f43616e6e6f74207374616b65206e6f7468696e6700000000000000000000000060448201526064016103bc565b6001600160a01b03811660009081526066602090815260408083203384528083528184208251606081018452858152808501889052439381019390935280546001808201835591865293909420825160039094020180549194929390929091839160ff19909116908381811115610e0357634e487b7160e01b600052602160045260246000fd5b02179055506020820151816001015560408201518160020155505082816001016000828254610e329190611257565b9091555050505050565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff1680610ebf575060005460ff16155b610f225760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103bc565b600054610100900460ff16158015610f44576000805461ffff19166101011790555b610f4c610f68565b610f54611019565b8015610cd6576000805461ff001916905550565b600054610100900460ff1680610f81575060005460ff16155b610fe45760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103bc565b600054610100900460ff16158015610f54576000805461ffff19166101011790558015610cd6576000805461ff001916905550565b600054610100900460ff1680611032575060005460ff16155b6110955760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103bc565b600054610100900460ff161580156110b7576000805461ffff19166101011790555b610f5433610e3c565b6000602082840312156110d1578081fd5b81356110dc816112b7565b9392505050565b600080604083850312156110f5578081fd5b8235611100816112b7565b91506020830135611110816112b7565b809150509250929050565b60008060006060848603121561112f578081fd5b833561113a816112b7565b9250602084013561114a816112b7565b929592945050506040919091013590565b60006020828403121561116c578081fd5b815180151581146110dc578182fd5b60006020828403121561118c578081fd5b5035919050565b600080604083850312156111a5578182fd5b82356111b0816112b7565b946020939093013593505050565b600080604083850312156111d0578182fd5b823591506020830135611110816112b7565b602080825282518282018190526000919060409081850190868401855b8281101561124a57815180516002811061122757634e487b7160e01b89526021600452602489fd5b8552808701518786015285015185850152606090930192908501906001016111ff565b5091979650505050505050565b6000821982111561126a5761126a6112a1565b500190565b600082821015611281576112816112a1565b500390565b600060001982141561129a5761129a6112a1565b5060010190565b634e487b7160e01b600052601160045260246000fd5b6001600160a01b0381168114610cd657600080fdfea2646970667358221220b86c4dbbce3853336ef7fe95b1d80a8b2da5a9108ddefdd3a619efa21b67e68e64736f6c63430008040033"

// DeployStakingManager deploys a new Ethereum contract, binding an instance of StakingManager to it.
func DeployStakingManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StakingManager, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StakingManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingManager{StakingManagerCaller: StakingManagerCaller{contract: contract}, StakingManagerTransactor: StakingManagerTransactor{contract: contract}, StakingManagerFilterer: StakingManagerFilterer{contract: contract}}, nil
}

// StakingManager is an auto generated Go binding around an Ethereum contract.
type StakingManager struct {
	StakingManagerCaller     // Read-only binding to the contract
	StakingManagerTransactor // Write-only binding to the contract
	StakingManagerFilterer   // Log filterer for contract events
}

// StakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingManagerSession struct {
	Contract     *StakingManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingManagerCallerSession struct {
	Contract *StakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingManagerTransactorSession struct {
	Contract     *StakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingManagerRaw struct {
	Contract *StakingManager // Generic contract binding to access the raw methods on
}

// StakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingManagerCallerRaw struct {
	Contract *StakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingManagerTransactorRaw struct {
	Contract *StakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingManager creates a new instance of StakingManager, bound to a specific deployed contract.
func NewStakingManager(address common.Address, backend bind.ContractBackend) (*StakingManager, error) {
	contract, err := bindStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingManager{StakingManagerCaller: StakingManagerCaller{contract: contract}, StakingManagerTransactor: StakingManagerTransactor{contract: contract}, StakingManagerFilterer: StakingManagerFilterer{contract: contract}}, nil
}

// NewStakingManagerCaller creates a new read-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*StakingManagerCaller, error) {
	contract, err := bindStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerCaller{contract: contract}, nil
}

// NewStakingManagerTransactor creates a new write-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingManagerTransactor, error) {
	contract, err := bindStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerTransactor{contract: contract}, nil
}

// NewStakingManagerFilterer creates a new log filterer instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingManagerFilterer, error) {
	contract, err := bindStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingManagerFilterer{contract: contract}, nil
}

// bindStakingManager binds a generic wrapper to an already deployed contract.
func bindStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.StakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transact(opts, method, params...)
}

// GetKey is a free data retrieval call binding the contract method 0xbc298553.
//
// Solidity: function getKey(address staker, address stakee) pure returns(bytes32)
func (_StakingManager *StakingManagerCaller) GetKey(opts *bind.CallOpts, staker common.Address, stakee common.Address) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getKey", staker, stakee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKey is a free data retrieval call binding the contract method 0xbc298553.
//
// Solidity: function getKey(address staker, address stakee) pure returns(bytes32)
func (_StakingManager *StakingManagerSession) GetKey(staker common.Address, stakee common.Address) ([32]byte, error) {
	return _StakingManager.Contract.GetKey(&_StakingManager.CallOpts, staker, stakee)
}

// GetKey is a free data retrieval call binding the contract method 0xbc298553.
//
// Solidity: function getKey(address staker, address stakee) pure returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) GetKey(staker common.Address, stakee common.Address) ([32]byte, error) {
	return _StakingManager.Contract.GetKey(&_StakingManager.CallOpts, staker, stakee)
}

// GetStakeOperations is a free data retrieval call binding the contract method 0x6959c1c5.
//
// Solidity: function getStakeOperations(address staker, address stakee) view returns((uint8,uint256,uint256)[])
func (_StakingManager *StakingManagerCaller) GetStakeOperations(opts *bind.CallOpts, staker common.Address, stakee common.Address) ([]StakingManagerStakeOperation, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getStakeOperations", staker, stakee)

	if err != nil {
		return *new([]StakingManagerStakeOperation), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingManagerStakeOperation)).(*[]StakingManagerStakeOperation)

	return out0, err

}

// GetStakeOperations is a free data retrieval call binding the contract method 0x6959c1c5.
//
// Solidity: function getStakeOperations(address staker, address stakee) view returns((uint8,uint256,uint256)[])
func (_StakingManager *StakingManagerSession) GetStakeOperations(staker common.Address, stakee common.Address) ([]StakingManagerStakeOperation, error) {
	return _StakingManager.Contract.GetStakeOperations(&_StakingManager.CallOpts, staker, stakee)
}

// GetStakeOperations is a free data retrieval call binding the contract method 0x6959c1c5.
//
// Solidity: function getStakeOperations(address staker, address stakee) view returns((uint8,uint256,uint256)[])
func (_StakingManager *StakingManagerCallerSession) GetStakeOperations(staker common.Address, stakee common.Address) ([]StakingManagerStakeOperation, error) {
	return _StakingManager.Contract.GetStakeOperations(&_StakingManager.CallOpts, staker, stakee)
}

// GetStakeeTotalStake is a free data retrieval call binding the contract method 0x64084d4e.
//
// Solidity: function getStakeeTotalStake(address stakee) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetStakeeTotalStake(opts *bind.CallOpts, stakee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getStakeeTotalStake", stakee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeeTotalStake is a free data retrieval call binding the contract method 0x64084d4e.
//
// Solidity: function getStakeeTotalStake(address stakee) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetStakeeTotalStake(stakee common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetStakeeTotalStake(&_StakingManager.CallOpts, stakee)
}

// GetStakeeTotalStake is a free data retrieval call binding the contract method 0x64084d4e.
//
// Solidity: function getStakeeTotalStake(address stakee) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetStakeeTotalStake(stakee common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetStakeeTotalStake(&_StakingManager.CallOpts, stakee)
}

// GetStakerAmount is a free data retrieval call binding the contract method 0xb91cefd2.
//
// Solidity: function getStakerAmount(address staker, address stakee, uint256 blockNumber) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetStakerAmount(opts *bind.CallOpts, staker common.Address, stakee common.Address, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getStakerAmount", staker, stakee, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakerAmount is a free data retrieval call binding the contract method 0xb91cefd2.
//
// Solidity: function getStakerAmount(address staker, address stakee, uint256 blockNumber) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetStakerAmount(staker common.Address, stakee common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _StakingManager.Contract.GetStakerAmount(&_StakingManager.CallOpts, staker, stakee, blockNumber)
}

// GetStakerAmount is a free data retrieval call binding the contract method 0xb91cefd2.
//
// Solidity: function getStakerAmount(address staker, address stakee, uint256 blockNumber) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetStakerAmount(staker common.Address, stakee common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _StakingManager.Contract.GetStakerAmount(&_StakingManager.CallOpts, staker, stakee, blockNumber)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetTotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getTotalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() view returns(uint256)
func (_StakingManager *StakingManagerSession) GetTotalStake() (*big.Int, error) {
	return _StakingManager.Contract.GetTotalStake(&_StakingManager.CallOpts)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetTotalStake() (*big.Int, error) {
	return _StakingManager.Contract.GetTotalStake(&_StakingManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerSession) Owner() (common.Address, error) {
	return _StakingManager.Contract.Owner(&_StakingManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerCallerSession) Owner() (common.Address, error) {
	return _StakingManager.Contract.Owner(&_StakingManager.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_StakingManager *StakingManagerCaller) TotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "totalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_StakingManager *StakingManagerSession) TotalStake() (*big.Int, error) {
	return _StakingManager.Contract.TotalStake(&_StakingManager.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) TotalStake() (*big.Int, error) {
	return _StakingManager.Contract.TotalStake(&_StakingManager.CallOpts)
}

// UnlockDuration is a free data retrieval call binding the contract method 0x8a1fcd60.
//
// Solidity: function unlockDuration() view returns(uint256)
func (_StakingManager *StakingManagerCaller) UnlockDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "unlockDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockDuration is a free data retrieval call binding the contract method 0x8a1fcd60.
//
// Solidity: function unlockDuration() view returns(uint256)
func (_StakingManager *StakingManagerSession) UnlockDuration() (*big.Int, error) {
	return _StakingManager.Contract.UnlockDuration(&_StakingManager.CallOpts)
}

// UnlockDuration is a free data retrieval call binding the contract method 0x8a1fcd60.
//
// Solidity: function unlockDuration() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) UnlockDuration() (*big.Int, error) {
	return _StakingManager.Contract.UnlockDuration(&_StakingManager.CallOpts)
}

// Unlockings is a free data retrieval call binding the contract method 0x9341a536.
//
// Solidity: function unlockings(bytes32 ) view returns(uint256 amount, uint256 unlockAt)
func (_StakingManager *StakingManagerCaller) Unlockings(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Amount   *big.Int
	UnlockAt *big.Int
}, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "unlockings", arg0)

	outstruct := new(struct {
		Amount   *big.Int
		UnlockAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnlockAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Unlockings is a free data retrieval call binding the contract method 0x9341a536.
//
// Solidity: function unlockings(bytes32 ) view returns(uint256 amount, uint256 unlockAt)
func (_StakingManager *StakingManagerSession) Unlockings(arg0 [32]byte) (struct {
	Amount   *big.Int
	UnlockAt *big.Int
}, error) {
	return _StakingManager.Contract.Unlockings(&_StakingManager.CallOpts, arg0)
}

// Unlockings is a free data retrieval call binding the contract method 0x9341a536.
//
// Solidity: function unlockings(bytes32 ) view returns(uint256 amount, uint256 unlockAt)
func (_StakingManager *StakingManagerCallerSession) Unlockings(arg0 [32]byte) (struct {
	Amount   *big.Int
	UnlockAt *big.Int
}, error) {
	return _StakingManager.Contract.Unlockings(&_StakingManager.CallOpts, arg0)
}

// AddStake is a paid mutator transaction binding the contract method 0x2d49aa1c.
//
// Solidity: function addStake(uint256 amount, address stakee) returns()
func (_StakingManager *StakingManagerTransactor) AddStake(opts *bind.TransactOpts, amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "addStake", amount, stakee)
}

// AddStake is a paid mutator transaction binding the contract method 0x2d49aa1c.
//
// Solidity: function addStake(uint256 amount, address stakee) returns()
func (_StakingManager *StakingManagerSession) AddStake(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.AddStake(&_StakingManager.TransactOpts, amount, stakee)
}

// AddStake is a paid mutator transaction binding the contract method 0x2d49aa1c.
//
// Solidity: function addStake(uint256 amount, address stakee) returns()
func (_StakingManager *StakingManagerTransactorSession) AddStake(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.AddStake(&_StakingManager.TransactOpts, amount, stakee)
}

// CancelUnlocking is a paid mutator transaction binding the contract method 0x23314c6c.
//
// Solidity: function cancelUnlocking(uint256 amount, address stakee) returns()
func (_StakingManager *StakingManagerTransactor) CancelUnlocking(opts *bind.TransactOpts, amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "cancelUnlocking", amount, stakee)
}

// CancelUnlocking is a paid mutator transaction binding the contract method 0x23314c6c.
//
// Solidity: function cancelUnlocking(uint256 amount, address stakee) returns()
func (_StakingManager *StakingManagerSession) CancelUnlocking(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.CancelUnlocking(&_StakingManager.TransactOpts, amount, stakee)
}

// CancelUnlocking is a paid mutator transaction binding the contract method 0x23314c6c.
//
// Solidity: function cancelUnlocking(uint256 amount, address stakee) returns()
func (_StakingManager *StakingManagerTransactorSession) CancelUnlocking(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.CancelUnlocking(&_StakingManager.TransactOpts, amount, stakee)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address token, uint256 _unlockDuration) returns()
func (_StakingManager *StakingManagerTransactor) Initialize(opts *bind.TransactOpts, token common.Address, _unlockDuration *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "initialize", token, _unlockDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address token, uint256 _unlockDuration) returns()
func (_StakingManager *StakingManagerSession) Initialize(token common.Address, _unlockDuration *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.Initialize(&_StakingManager.TransactOpts, token, _unlockDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address token, uint256 _unlockDuration) returns()
func (_StakingManager *StakingManagerTransactorSession) Initialize(token common.Address, _unlockDuration *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.Initialize(&_StakingManager.TransactOpts, token, _unlockDuration)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceOwnership(&_StakingManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceOwnership(&_StakingManager.TransactOpts)
}

// SetUnlockDuration is a paid mutator transaction binding the contract method 0xdd900769.
//
// Solidity: function setUnlockDuration(uint256 newUnlockDuration) returns()
func (_StakingManager *StakingManagerTransactor) SetUnlockDuration(opts *bind.TransactOpts, newUnlockDuration *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setUnlockDuration", newUnlockDuration)
}

// SetUnlockDuration is a paid mutator transaction binding the contract method 0xdd900769.
//
// Solidity: function setUnlockDuration(uint256 newUnlockDuration) returns()
func (_StakingManager *StakingManagerSession) SetUnlockDuration(newUnlockDuration *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetUnlockDuration(&_StakingManager.TransactOpts, newUnlockDuration)
}

// SetUnlockDuration is a paid mutator transaction binding the contract method 0xdd900769.
//
// Solidity: function setUnlockDuration(uint256 newUnlockDuration) returns()
func (_StakingManager *StakingManagerTransactorSession) SetUnlockDuration(newUnlockDuration *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetUnlockDuration(&_StakingManager.TransactOpts, newUnlockDuration)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.TransferOwnership(&_StakingManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.TransferOwnership(&_StakingManager.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xa859f172.
//
// Solidity: function unlockStake(uint256 amount, address stakee) returns(uint256)
func (_StakingManager *StakingManagerTransactor) UnlockStake(opts *bind.TransactOpts, amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "unlockStake", amount, stakee)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xa859f172.
//
// Solidity: function unlockStake(uint256 amount, address stakee) returns(uint256)
func (_StakingManager *StakingManagerSession) UnlockStake(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.UnlockStake(&_StakingManager.TransactOpts, amount, stakee)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xa859f172.
//
// Solidity: function unlockStake(uint256 amount, address stakee) returns(uint256)
func (_StakingManager *StakingManagerTransactorSession) UnlockStake(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.UnlockStake(&_StakingManager.TransactOpts, amount, stakee)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address stakee) returns()
func (_StakingManager *StakingManagerTransactor) WithdrawStake(opts *bind.TransactOpts, stakee common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "withdrawStake", stakee)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address stakee) returns()
func (_StakingManager *StakingManagerSession) WithdrawStake(stakee common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.WithdrawStake(&_StakingManager.TransactOpts, stakee)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address stakee) returns()
func (_StakingManager *StakingManagerTransactorSession) WithdrawStake(stakee common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.WithdrawStake(&_StakingManager.TransactOpts, stakee)
}

// StakingManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingManager contract.
type StakingManagerOwnershipTransferredIterator struct {
	Event *StakingManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerOwnershipTransferred)
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
		it.Event = new(StakingManagerOwnershipTransferred)
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
func (it *StakingManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerOwnershipTransferred represents a OwnershipTransferred event raised by the StakingManager contract.
type StakingManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingManager *StakingManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerOwnershipTransferredIterator{contract: _StakingManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingManager *StakingManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerOwnershipTransferred)
				if err := _StakingManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseOwnershipTransferred(log types.Log) (*StakingManagerOwnershipTransferred, error) {
	event := new(StakingManagerOwnershipTransferred)
	if err := _StakingManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
