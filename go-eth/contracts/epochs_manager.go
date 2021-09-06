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
	StartBlock              *big.Int
	Duration                *big.Int
	EndBlock                *big.Int
	DirectoryId             [32]byte
	DefaultPayoutPercentage uint16
	FaceValue               *big.Int
	BaseLiveWinProb         *big.Int
	ExpiredWinProb          *big.Int
	TicketDuration          *big.Int
	DecayRate               uint16
}

// EpochsManagerABI is the input ABI used to generate the binding from.
const EpochsManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"}],\"name\":\"NewEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"currentActiveEpoch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDirectory\",\"name\":\"directory\",\"type\":\"address\"},{\"internalType\":\"contractListings\",\"name\":\"listings\",\"type\":\"address\"},{\"internalType\":\"contractTicketingParameters\",\"name\":\"ticketingParameters\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_epochDuration\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeEpoch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentActiveEpoch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"directoryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"defaultPayoutPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"baseLiveWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"expiredWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"ticketDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"decayRate\",\"type\":\"uint16\"}],\"internalType\":\"structEpochsManager.Epoch\",\"name\":\"epoch\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"directoryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"defaultPayoutPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"baseLiveWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"expiredWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"ticketDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"decayRate\",\"type\":\"uint16\"}],\"internalType\":\"structEpochsManager.Epoch\",\"name\":\"epoch\",\"type\":\"tuple\"}],\"name\":\"getEpochId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"}],\"name\":\"getEpoch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"directoryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"defaultPayoutPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"baseLiveWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"expiredWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"ticketDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"decayRate\",\"type\":\"uint16\"}],\"internalType\":\"structEpochsManager.Epoch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EpochsManagerBin is the compiled bytecode used for deploying new contracts.
var EpochsManagerBin = "0x608060405234801561001057600080fd5b50611306806100206000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c8063ae4a4d5711610076578063e1519a751161005b578063e1519a751461031c578063e34bf01314610324578063f2fde38b1461042d57600080fd5b8063ae4a4d5714610300578063cf756fdf1461030957600080fd5b8063715018a6116100a7578063715018a6146101bd5780637e6d64a5146101c75780638da5cb5b146102e557600080fd5b806325d50f03146100c35780634ff0876a146101b4575b600080fd5b6101a16100d13660046110a8565b805160208083015160608085015160808087015160a088015160c089015160e08a01516101008b0151610120909b015160408051808c019c909c528b810199909952968a01959095527fffff00000000000000000000000000000000000000000000000000000000000060f093841b81168a86015260828a01929092526fffffffffffffffffffffffffffffffff1990841b811660a28a01529390921b90921660b287015260c28601969096521b90931660e2830152825180830360c401815260e4909201909252805191012090565b6040519081526020015b60405180910390f35b6101a160695481565b6101c5610440565b005b6102d86101d5366004611028565b6040805161014081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526101008101829052610120810191909152506000908152606b60209081526040918290208251610140810184528154815260018201549281019290925260028101549282019290925260038201546060820152600482015461ffff9081166080830152600583015460a083015260068301546fffffffffffffffffffffffffffffffff80821660c08501527001000000000000000000000000000000009091041660e0830152600783015461010083015260089092015490911661012082015290565b6040516101ab919061117d565b6033546040516001600160a01b0390911681526020016101ab565b6101a1606a5481565b6101c5610317366004611058565b6104ab565b6101a16105c4565b6102d86040805161014081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081019190915250606a546000908152606b60209081526040918290208251610140810184528154815260018201549281019290925260028101549282019290925260038201546060820152600482015461ffff9081166080830152600583015460a083015260068301546fffffffffffffffffffffffffffffffff80821660c08501527001000000000000000000000000000000009091041660e0830152600783015461010083015260089092015490911661012082015290565b6101c561043b366004611005565b610c84565b6033546001600160a01b0316331461049f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6104a96000610d66565b565b600054610100900460ff16806104c4575060005460ff16155b6105275760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610496565b600054610100900460ff16158015610549576000805461ffff19166101011790555b610551610dd0565b606680546001600160a01b038088167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255606780548784169083161790556068805492861692909116919091179055606982905580156105bd576000805461ff00191690555b5050505050565b606a546000908152606b602052604081206001810154815483916105e791611269565b9050438111156106395760405162461bcd60e51b815260206004820152601f60248201527f43757272656e742065706f636820686173206e6f742079657420656e646564006044820152606401610496565b606654604080517f859ea34700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b03169163859ea34791600480830192602092919082900301818787803b15801561069857600080fd5b505af11580156106ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106d09190611040565b90506000604051806101400160405280438152602001606954815260200160008152602001838152602001606760009054906101000a90046001600160a01b03166001600160a01b031663d2a78d7f6040518163ffffffff1660e01b815260040160206040518083038186803b15801561074957600080fd5b505afa15801561075d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107819190611161565b61ffff168152606854604080517f44fd9caa00000000000000000000000000000000000000000000000000000000815290516020938401936001600160a01b03909316926344fd9caa9260048082019391829003018186803b1580156107e657600080fd5b505afa1580156107fa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061081e9190611040565b8152606854604080517fdedcebda00000000000000000000000000000000000000000000000000000000815290516020938401936001600160a01b039093169263dedcebda9260048082019391829003018186803b15801561087f57600080fd5b505afa158015610893573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108b79190611145565b6fffffffffffffffffffffffffffffffff168152606854604080517fbcbee54300000000000000000000000000000000000000000000000000000000815290516020938401936001600160a01b039093169263bcbee5439260048082019391829003018186803b15801561092a57600080fd5b505afa15801561093e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109629190611145565b6fffffffffffffffffffffffffffffffff168152606854604080517f87bcc0c500000000000000000000000000000000000000000000000000000000815290516020938401936001600160a01b03909316926387bcc0c59260048082019391829003018186803b1580156109d557600080fd5b505afa1580156109e9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a0d9190611040565b8152606854604080517fa9c1f2f100000000000000000000000000000000000000000000000000000000815290516020938401936001600160a01b039093169263a9c1f2f19260048082019391829003018186803b158015610a6e57600080fd5b505afa158015610a82573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aa69190611161565b61ffff169052805160208083015160608085015160808087015160a088015160c089015160e08a01516101008b01516101208c015160408051808d019d909d528c81019a909a52978b01969096527fffff00000000000000000000000000000000000000000000000000000000000060f094851b81168b87015260828b01939093526fffffffffffffffffffffffffffffffff1991851b821660a28b015290931b90921660b288015260c287019290925291901b1660e2840152805180840360c401815260e4909301905281519101209091506000906000818152606b6020908152604091829020855181559085015160018201558185015160028083019190915560608601516003830155608086015160048301805461ffff92831661ffff199182161790915560a0880151600585015560c088015160e08901516fffffffffffffffffffffffffffffffff908116700100000000000000000000000000000000029116176006850155610100880151600785015561012088015160089094018054949092169316929092179091554390880155606a829055519091507fddc860800a99149017c480ec51523bf4143b7215e78956ae5c31e5c568f5383a90610c739083815260200190565b60405180910390a195945050505050565b6033546001600160a01b03163314610cde5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610496565b6001600160a01b038116610d5a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610496565b610d6381610d66565b50565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff1680610de9575060005460ff16155b610e4c5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610496565b600054610100900460ff16158015610e6e576000805461ffff19166101011790555b610e76610e92565b610e7e610f43565b8015610d63576000805461ff001916905550565b600054610100900460ff1680610eab575060005460ff16155b610f0e5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610496565b600054610100900460ff16158015610e7e576000805461ffff19166101011790558015610d63576000805461ff001916905550565b600054610100900460ff1680610f5c575060005460ff16155b610fbf5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610496565b600054610100900460ff16158015610fe1576000805461ffff19166101011790555b610e7e33610d66565b8035610ff5816112a2565b919050565b8035610ff5816112c0565b600060208284031215611016578081fd5b81356110218161128d565b9392505050565b600060208284031215611039578081fd5b5035919050565b600060208284031215611051578081fd5b5051919050565b6000806000806080858703121561106d578283fd5b84356110788161128d565b935060208501356110888161128d565b925060408501356110988161128d565b9396929550929360600135925050565b600061014082840312156110ba578081fd5b6110c2611231565b823581526020830135602082015260408301356040820152606083013560608201526110f060808401610ffa565b608082015260a083013560a082015261110b60c08401610fea565b60c082015261111c60e08401610fea565b60e0820152610100838101359082015261012061113a818501610ffa565b908201529392505050565b600060208284031215611156578081fd5b8151611021816112a2565b600060208284031215611172578081fd5b8151611021816112c0565b6000610140820190508251825260208301516020830152604083015160408301526060830151606083015260808301516111bd608084018261ffff169052565b5060a083015160a083015260c08301516111eb60c08401826fffffffffffffffffffffffffffffffff169052565b5060e083015161120f60e08401826fffffffffffffffffffffffffffffffff169052565b5061010083810151908301526101209283015161ffff16929091019190915290565b604051610140810167ffffffffffffffff8111828210171561126357634e487b7160e01b600052604160045260246000fd5b60405290565b6000821982111561128857634e487b7160e01b81526011600452602481fd5b500190565b6001600160a01b0381168114610d6357600080fd5b6fffffffffffffffffffffffffffffffff81168114610d6357600080fd5b61ffff81168114610d6357600080fdfea2646970667358221220d458fb827810524c5f282e589f2dc5938e7f0be44fbe5651da800489af3f10d364736f6c63430008040033"

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
// Solidity: function getCurrentActiveEpoch() view returns((uint256,uint256,uint256,bytes32,uint16,uint256,uint128,uint128,uint256,uint16) epoch)
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
// Solidity: function getCurrentActiveEpoch() view returns((uint256,uint256,uint256,bytes32,uint16,uint256,uint128,uint128,uint256,uint16) epoch)
func (_EpochsManager *EpochsManagerSession) GetCurrentActiveEpoch() (EpochsManagerEpoch, error) {
	return _EpochsManager.Contract.GetCurrentActiveEpoch(&_EpochsManager.CallOpts)
}

// GetCurrentActiveEpoch is a free data retrieval call binding the contract method 0xe34bf013.
//
// Solidity: function getCurrentActiveEpoch() view returns((uint256,uint256,uint256,bytes32,uint16,uint256,uint128,uint128,uint256,uint16) epoch)
func (_EpochsManager *EpochsManagerCallerSession) GetCurrentActiveEpoch() (EpochsManagerEpoch, error) {
	return _EpochsManager.Contract.GetCurrentActiveEpoch(&_EpochsManager.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x7e6d64a5.
//
// Solidity: function getEpoch(bytes32 epochId) view returns((uint256,uint256,uint256,bytes32,uint16,uint256,uint128,uint128,uint256,uint16))
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
// Solidity: function getEpoch(bytes32 epochId) view returns((uint256,uint256,uint256,bytes32,uint16,uint256,uint128,uint128,uint256,uint16))
func (_EpochsManager *EpochsManagerSession) GetEpoch(epochId [32]byte) (EpochsManagerEpoch, error) {
	return _EpochsManager.Contract.GetEpoch(&_EpochsManager.CallOpts, epochId)
}

// GetEpoch is a free data retrieval call binding the contract method 0x7e6d64a5.
//
// Solidity: function getEpoch(bytes32 epochId) view returns((uint256,uint256,uint256,bytes32,uint16,uint256,uint128,uint128,uint256,uint16))
func (_EpochsManager *EpochsManagerCallerSession) GetEpoch(epochId [32]byte) (EpochsManagerEpoch, error) {
	return _EpochsManager.Contract.GetEpoch(&_EpochsManager.CallOpts, epochId)
}

// GetEpochId is a free data retrieval call binding the contract method 0x25d50f03.
//
// Solidity: function getEpochId((uint256,uint256,uint256,bytes32,uint16,uint256,uint128,uint128,uint256,uint16) epoch) pure returns(bytes32)
func (_EpochsManager *EpochsManagerCaller) GetEpochId(opts *bind.CallOpts, epoch EpochsManagerEpoch) ([32]byte, error) {
	var out []interface{}
	err := _EpochsManager.contract.Call(opts, &out, "getEpochId", epoch)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEpochId is a free data retrieval call binding the contract method 0x25d50f03.
//
// Solidity: function getEpochId((uint256,uint256,uint256,bytes32,uint16,uint256,uint128,uint128,uint256,uint16) epoch) pure returns(bytes32)
func (_EpochsManager *EpochsManagerSession) GetEpochId(epoch EpochsManagerEpoch) ([32]byte, error) {
	return _EpochsManager.Contract.GetEpochId(&_EpochsManager.CallOpts, epoch)
}

// GetEpochId is a free data retrieval call binding the contract method 0x25d50f03.
//
// Solidity: function getEpochId((uint256,uint256,uint256,bytes32,uint16,uint256,uint128,uint128,uint256,uint16) epoch) pure returns(bytes32)
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
