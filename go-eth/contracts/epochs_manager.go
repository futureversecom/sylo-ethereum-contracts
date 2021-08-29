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
	DirectoryId     [32]byte
	FaceValue       *big.Int
	BaseLiveWinProb *big.Int
	ExpiredWinProb  *big.Int
	TicketDuration  *big.Int
	DecayRate       uint16
}

// EpochsManagerABI is the input ABI used to generate the binding from.
const EpochsManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"}],\"name\":\"NewEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"currentActiveEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"directoryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"baseLiveWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"expiredWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"ticketDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"decayRate\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"previousActiveEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"directoryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"baseLiveWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"expiredWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"ticketDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"decayRate\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractDirectory\",\"name\":\"directory\",\"type\":\"address\"},{\"internalType\":\"contractTicketingParameters\",\"name\":\"ticketingParameters\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_epochDuration\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeEpoch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentActiveEpoch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"directoryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"baseLiveWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"expiredWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"ticketDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"decayRate\",\"type\":\"uint16\"}],\"internalType\":\"structEpochsManager.Epoch\",\"name\":\"epoch\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"directoryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"baseLiveWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"expiredWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"ticketDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"decayRate\",\"type\":\"uint16\"}],\"internalType\":\"structEpochsManager.Epoch\",\"name\":\"epoch\",\"type\":\"tuple\"}],\"name\":\"getEpochId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"}],\"name\":\"getEpoch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"directoryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"baseLiveWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"expiredWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"ticketDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"decayRate\",\"type\":\"uint16\"}],\"internalType\":\"structEpochsManager.Epoch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EpochsManagerBin is the compiled bytecode used for deploying new contracts.
var EpochsManagerBin = "0x608060405234801561001057600080fd5b506112a6806100206000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80638da5cb5b11610081578063e1519a751161005b578063e1519a7514610338578063e34bf01314610340578063f2fde38b146103e757600080fd5b80638da5cb5b1461025e578063923e672414610279578063ae4a4d57146102fe57600080fd5b8063715018a6116100b2578063715018a6146100ff57806373f99a5d146101075780637e6d64a51461011a57600080fd5b80631794bb3c146100ce5780634ff0876a146100e3575b600080fd5b6100e16100dc3660046110ed565b6103fa565b005b6100ec60685481565b6040519081526020015b60405180910390f35b6100e161050a565b6100ec61011536600461112d565b610570565b6101ea6101283660046110bd565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081019190915250600090815260776020908152604091829020825161010081018452815481526001820154928101929092526002810154928201929092526003820154606082015260048201546001600160801b038082166080840152600160801b9091041660a0820152600582015460c082015260069091015461ffff1660e082015290565b6040516100f691906000610100820190508251825260208301516020830152604083015160408301526060830151606083015260808301516001600160801b0380821660808501528060a08601511660a0850152505060c083015160c083015261ffff60e08401511660e083015292915050565b6033546040516001600160a01b0390911681526020016100f6565b6070546071546072546073546074546075546076546102b396959493926001600160801b0380821693600160801b909204169161ffff1688565b6040805198895260208901979097529587019490945260608601929092526001600160801b0390811660808601521660a084015260c083015261ffff1660e0820152610100016100f6565b606954606a54606b54606c54606d54606e54606f546102b396959493926001600160801b0380821693600160801b909204169161ffff1688565b6100ec610640565b6101ea6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101919091525060408051610100810182526069548152606a546020820152606b5491810191909152606c546060820152606d546001600160801b038082166080840152600160801b9091041660a0820152606e5460c0820152606f5461ffff1660e082015290565b6100e16103f536600461109a565b610d19565b600054610100900460ff1680610413575060005460ff16155b61047b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b600054610100900460ff1615801561049d576000805461ffff19166101011790555b6104a5610dfb565b606680546001600160a01b038087167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255606780549286169290911691909117905560688290558015610504576000805461ff00191690555b50505050565b6033546001600160a01b031633146105645760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610472565b61056e6000610ebd565b565b600081600001518260200151836060015184608001518560a001518660c001518760e00151604051602001610623979695949392919096875260208701959095526040860193909352608091821b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000908116606087015290821b16607085015283015260f01b7fffff0000000000000000000000000000000000000000000000000000000000001660a082015260a20190565b604051602081830303815290604052805190602001209050919050565b606a5460695460009182916106559190611212565b9050438111156106a75760405162461bcd60e51b815260206004820152601f60248201527f43757272656e742065706f636820686173206e6f742079657420656e646564006044820152606401610472565b606654604080517f859ea34700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b03169163859ea34791600480830192602092919082900301818787803b15801561070657600080fd5b505af115801561071a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061073e91906110d5565b905060006040518061010001604052804381526020016068548152602001838152602001606760009054906101000a90046001600160a01b03166001600160a01b03166344fd9caa6040518163ffffffff1660e01b815260040160206040518083038186803b1580156107b057600080fd5b505afa1580156107c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107e891906110d5565b8152606754604080517fdedcebda00000000000000000000000000000000000000000000000000000000815290516020938401936001600160a01b039093169263dedcebda9260048082019391829003018186803b15801561084957600080fd5b505afa15801561085d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061088191906111da565b6001600160801b03168152606754604080517fbcbee54300000000000000000000000000000000000000000000000000000000815290516020938401936001600160a01b039093169263bcbee5439260048082019391829003018186803b1580156108eb57600080fd5b505afa1580156108ff573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061092391906111da565b6001600160801b03168152606754604080517f87bcc0c500000000000000000000000000000000000000000000000000000000815290516020938401936001600160a01b03909316926387bcc0c59260048082019391829003018186803b15801561098d57600080fd5b505afa1580156109a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109c591906110d5565b8152606754604080517fa9c1f2f100000000000000000000000000000000000000000000000000000000815290516020938401936001600160a01b039093169263a9c1f2f19260048082019391829003018186803b158015610a2657600080fd5b505afa158015610a3a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a5e91906111f6565b61ffff16905290506000610a7182610570565b905081607760008381526020019081526020016000206000820151816000015560208201518160010155604082015181600201556060820151816003015560808201518160040160006101000a8154816001600160801b0302191690836001600160801b0316021790555060a08201518160040160106101000a8154816001600160801b0302191690836001600160801b0316021790555060c0820151816005015560e08201518160060160006101000a81548161ffff021916908361ffff16021790555090505060696070600082015481600001556001820154816001015560028201548160020155600382015481600301556004820160009054906101000a90046001600160801b03168160040160006101000a8154816001600160801b0302191690836001600160801b031602179055506004820160109054906101000a90046001600160801b03168160040160106101000a8154816001600160801b0302191690836001600160801b03160217905550600582015481600501556006820160009054906101000a900461ffff168160060160006101000a81548161ffff021916908361ffff1602179055509050508160696000820151816000015560208201518160010155604082015181600201556060820151816003015560808201518160040160006101000a8154816001600160801b0302191690836001600160801b0316021790555060a08201518160040160106101000a8154816001600160801b0302191690836001600160801b0316021790555060c0820151816005015560e08201518160060160006101000a81548161ffff021916908361ffff1602179055509050507fddc860800a99149017c480ec51523bf4143b7215e78956ae5c31e5c568f5383a81604051610d0991815260200190565b60405180910390a1949350505050565b6033546001600160a01b03163314610d735760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610472565b6001600160a01b038116610def5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610472565b610df881610ebd565b50565b600054610100900460ff1680610e14575060005460ff16155b610e775760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610472565b600054610100900460ff16158015610e99576000805461ffff19166101011790555b610ea1610f27565b610ea9610fd8565b8015610df8576000805461ff001916905550565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff1680610f40575060005460ff16155b610fa35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610472565b600054610100900460ff16158015610ea9576000805461ffff19166101011790558015610df8576000805461ff001916905550565b600054610100900460ff1680610ff1575060005460ff16155b6110545760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610472565b600054610100900460ff16158015611076576000805461ffff19166101011790555b610ea933610ebd565b803561108a8161124b565b919050565b803561108a81611260565b6000602082840312156110ab578081fd5b81356110b681611236565b9392505050565b6000602082840312156110ce578081fd5b5035919050565b6000602082840312156110e6578081fd5b5051919050565b600080600060608486031215611101578182fd5b833561110c81611236565b9250602084013561111c81611236565b929592945050506040919091013590565b6000610100808385031215611140578182fd5b6040519081019067ffffffffffffffff8211818310171561116f57634e487b7160e01b83526041600452602483fd5b81604052833581526020840135602082015260408401356040820152606084013560608201526111a16080850161107f565b60808201526111b260a0850161107f565b60a082015260c084013560c08201526111cd60e0850161108f565b60e0820152949350505050565b6000602082840312156111eb578081fd5b81516110b68161124b565b600060208284031215611207578081fd5b81516110b681611260565b6000821982111561123157634e487b7160e01b81526011600452602481fd5b500190565b6001600160a01b0381168114610df857600080fd5b6001600160801b0381168114610df857600080fd5b61ffff81168114610df857600080fdfea264697066735822122015f3e556948b70b87293abde7395db60af4533edda1850103efbc392daa0aa4464736f6c63430008040033"

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
// Solidity: function currentActiveEpoch() view returns(uint256 startBlock, uint256 duration, bytes32 directoryId, uint256 faceValue, uint128 baseLiveWinProb, uint128 expiredWinProb, uint256 ticketDuration, uint16 decayRate)
func (_EpochsManager *EpochsManagerCaller) CurrentActiveEpoch(opts *bind.CallOpts) (struct {
	StartBlock      *big.Int
	Duration        *big.Int
	DirectoryId     [32]byte
	FaceValue       *big.Int
	BaseLiveWinProb *big.Int
	ExpiredWinProb  *big.Int
	TicketDuration  *big.Int
	DecayRate       uint16
}, error) {
	var out []interface{}
	err := _EpochsManager.contract.Call(opts, &out, "currentActiveEpoch")

	outstruct := new(struct {
		StartBlock      *big.Int
		Duration        *big.Int
		DirectoryId     [32]byte
		FaceValue       *big.Int
		BaseLiveWinProb *big.Int
		ExpiredWinProb  *big.Int
		TicketDuration  *big.Int
		DecayRate       uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartBlock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DirectoryId = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.FaceValue = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.BaseLiveWinProb = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ExpiredWinProb = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TicketDuration = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.DecayRate = *abi.ConvertType(out[7], new(uint16)).(*uint16)

	return *outstruct, err

}

// CurrentActiveEpoch is a free data retrieval call binding the contract method 0xae4a4d57.
//
// Solidity: function currentActiveEpoch() view returns(uint256 startBlock, uint256 duration, bytes32 directoryId, uint256 faceValue, uint128 baseLiveWinProb, uint128 expiredWinProb, uint256 ticketDuration, uint16 decayRate)
func (_EpochsManager *EpochsManagerSession) CurrentActiveEpoch() (struct {
	StartBlock      *big.Int
	Duration        *big.Int
	DirectoryId     [32]byte
	FaceValue       *big.Int
	BaseLiveWinProb *big.Int
	ExpiredWinProb  *big.Int
	TicketDuration  *big.Int
	DecayRate       uint16
}, error) {
	return _EpochsManager.Contract.CurrentActiveEpoch(&_EpochsManager.CallOpts)
}

// CurrentActiveEpoch is a free data retrieval call binding the contract method 0xae4a4d57.
//
// Solidity: function currentActiveEpoch() view returns(uint256 startBlock, uint256 duration, bytes32 directoryId, uint256 faceValue, uint128 baseLiveWinProb, uint128 expiredWinProb, uint256 ticketDuration, uint16 decayRate)
func (_EpochsManager *EpochsManagerCallerSession) CurrentActiveEpoch() (struct {
	StartBlock      *big.Int
	Duration        *big.Int
	DirectoryId     [32]byte
	FaceValue       *big.Int
	BaseLiveWinProb *big.Int
	ExpiredWinProb  *big.Int
	TicketDuration  *big.Int
	DecayRate       uint16
}, error) {
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
// Solidity: function getCurrentActiveEpoch() view returns((uint256,uint256,bytes32,uint256,uint128,uint128,uint256,uint16) epoch)
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
// Solidity: function getCurrentActiveEpoch() view returns((uint256,uint256,bytes32,uint256,uint128,uint128,uint256,uint16) epoch)
func (_EpochsManager *EpochsManagerSession) GetCurrentActiveEpoch() (EpochsManagerEpoch, error) {
	return _EpochsManager.Contract.GetCurrentActiveEpoch(&_EpochsManager.CallOpts)
}

// GetCurrentActiveEpoch is a free data retrieval call binding the contract method 0xe34bf013.
//
// Solidity: function getCurrentActiveEpoch() view returns((uint256,uint256,bytes32,uint256,uint128,uint128,uint256,uint16) epoch)
func (_EpochsManager *EpochsManagerCallerSession) GetCurrentActiveEpoch() (EpochsManagerEpoch, error) {
	return _EpochsManager.Contract.GetCurrentActiveEpoch(&_EpochsManager.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x7e6d64a5.
//
// Solidity: function getEpoch(bytes32 epochId) view returns((uint256,uint256,bytes32,uint256,uint128,uint128,uint256,uint16))
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
// Solidity: function getEpoch(bytes32 epochId) view returns((uint256,uint256,bytes32,uint256,uint128,uint128,uint256,uint16))
func (_EpochsManager *EpochsManagerSession) GetEpoch(epochId [32]byte) (EpochsManagerEpoch, error) {
	return _EpochsManager.Contract.GetEpoch(&_EpochsManager.CallOpts, epochId)
}

// GetEpoch is a free data retrieval call binding the contract method 0x7e6d64a5.
//
// Solidity: function getEpoch(bytes32 epochId) view returns((uint256,uint256,bytes32,uint256,uint128,uint128,uint256,uint16))
func (_EpochsManager *EpochsManagerCallerSession) GetEpoch(epochId [32]byte) (EpochsManagerEpoch, error) {
	return _EpochsManager.Contract.GetEpoch(&_EpochsManager.CallOpts, epochId)
}

// GetEpochId is a free data retrieval call binding the contract method 0x73f99a5d.
//
// Solidity: function getEpochId((uint256,uint256,bytes32,uint256,uint128,uint128,uint256,uint16) epoch) pure returns(bytes32)
func (_EpochsManager *EpochsManagerCaller) GetEpochId(opts *bind.CallOpts, epoch EpochsManagerEpoch) ([32]byte, error) {
	var out []interface{}
	err := _EpochsManager.contract.Call(opts, &out, "getEpochId", epoch)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEpochId is a free data retrieval call binding the contract method 0x73f99a5d.
//
// Solidity: function getEpochId((uint256,uint256,bytes32,uint256,uint128,uint128,uint256,uint16) epoch) pure returns(bytes32)
func (_EpochsManager *EpochsManagerSession) GetEpochId(epoch EpochsManagerEpoch) ([32]byte, error) {
	return _EpochsManager.Contract.GetEpochId(&_EpochsManager.CallOpts, epoch)
}

// GetEpochId is a free data retrieval call binding the contract method 0x73f99a5d.
//
// Solidity: function getEpochId((uint256,uint256,bytes32,uint256,uint128,uint128,uint256,uint16) epoch) pure returns(bytes32)
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

// PreviousActiveEpoch is a free data retrieval call binding the contract method 0x923e6724.
//
// Solidity: function previousActiveEpoch() view returns(uint256 startBlock, uint256 duration, bytes32 directoryId, uint256 faceValue, uint128 baseLiveWinProb, uint128 expiredWinProb, uint256 ticketDuration, uint16 decayRate)
func (_EpochsManager *EpochsManagerCaller) PreviousActiveEpoch(opts *bind.CallOpts) (struct {
	StartBlock      *big.Int
	Duration        *big.Int
	DirectoryId     [32]byte
	FaceValue       *big.Int
	BaseLiveWinProb *big.Int
	ExpiredWinProb  *big.Int
	TicketDuration  *big.Int
	DecayRate       uint16
}, error) {
	var out []interface{}
	err := _EpochsManager.contract.Call(opts, &out, "previousActiveEpoch")

	outstruct := new(struct {
		StartBlock      *big.Int
		Duration        *big.Int
		DirectoryId     [32]byte
		FaceValue       *big.Int
		BaseLiveWinProb *big.Int
		ExpiredWinProb  *big.Int
		TicketDuration  *big.Int
		DecayRate       uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartBlock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DirectoryId = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.FaceValue = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.BaseLiveWinProb = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ExpiredWinProb = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TicketDuration = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.DecayRate = *abi.ConvertType(out[7], new(uint16)).(*uint16)

	return *outstruct, err

}

// PreviousActiveEpoch is a free data retrieval call binding the contract method 0x923e6724.
//
// Solidity: function previousActiveEpoch() view returns(uint256 startBlock, uint256 duration, bytes32 directoryId, uint256 faceValue, uint128 baseLiveWinProb, uint128 expiredWinProb, uint256 ticketDuration, uint16 decayRate)
func (_EpochsManager *EpochsManagerSession) PreviousActiveEpoch() (struct {
	StartBlock      *big.Int
	Duration        *big.Int
	DirectoryId     [32]byte
	FaceValue       *big.Int
	BaseLiveWinProb *big.Int
	ExpiredWinProb  *big.Int
	TicketDuration  *big.Int
	DecayRate       uint16
}, error) {
	return _EpochsManager.Contract.PreviousActiveEpoch(&_EpochsManager.CallOpts)
}

// PreviousActiveEpoch is a free data retrieval call binding the contract method 0x923e6724.
//
// Solidity: function previousActiveEpoch() view returns(uint256 startBlock, uint256 duration, bytes32 directoryId, uint256 faceValue, uint128 baseLiveWinProb, uint128 expiredWinProb, uint256 ticketDuration, uint16 decayRate)
func (_EpochsManager *EpochsManagerCallerSession) PreviousActiveEpoch() (struct {
	StartBlock      *big.Int
	Duration        *big.Int
	DirectoryId     [32]byte
	FaceValue       *big.Int
	BaseLiveWinProb *big.Int
	ExpiredWinProb  *big.Int
	TicketDuration  *big.Int
	DecayRate       uint16
}, error) {
	return _EpochsManager.Contract.PreviousActiveEpoch(&_EpochsManager.CallOpts)
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
