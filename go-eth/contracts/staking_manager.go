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

// StakingManagerStakeEntry is an auto generated low-level Go binding around an user-defined struct.
type StakingManagerStakeEntry struct {
	Amount    *big.Int
	UpdatedAt *big.Int
	EpochId   *big.Int
}

// StakingManagerABI is the input ABI used to generate the binding from.
const StakingManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalManagedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"unlockings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractRewardsManager\",\"name\":\"rewardsManager\",\"type\":\"address\"},{\"internalType\":\"contractEpochsManager\",\"name\":\"epochsManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unlockDuration\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newUnlockDuration\",\"type\":\"uint256\"}],\"name\":\"setUnlockDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"unlockStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"cancelUnlocking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalManagedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getStakeEntry\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"internalType\":\"structStakingManager.StakeEntry\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getCurrentStakerAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getStakeeTotalManagedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StakingManagerBin is the compiled bytecode used for deploying new contracts.
var StakingManagerBin = "0x608060405234801561001057600080fd5b506112ed806100206000396000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c80639341a536116100b2578063c78702bb11610081578063dd90076911610066578063dd900769146102ca578063f2fde38b146102dd578063f731cb49146102f057600080fd5b8063c78702bb1461027e578063cf756fdf146102b757600080fd5b80639341a536146101c4578063a859f17214610200578063bc29855314610213578063c23a5cea1461026b57600080fd5b8063715018a6116100ee578063715018a61461018f5780637ffdacd5146101975780638a1fcd60146101a05780638da5cb5b146101a957600080fd5b806313cdd31b1461012057806319fba39b1461015f57806323314c6c146101675780632d49aa1c1461017c575b600080fd5b61014c61012e366004611125565b6001600160a01b031660009081526068602052604090206001015490565b6040519081526020015b60405180910390f35b60695461014c565b61017a610175366004611220565b610325565b005b61017a61018a366004611220565b61041b565b61017a6104cd565b61014c60695481565b61014c606a5481565b6033546040516001600160a01b039091168152602001610156565b6101eb6101d23660046111a0565b606b602052600090815260409020805460019091015482565b60408051928352602083019190915201610156565b61014c61020e366004611220565b610533565b61014c610221366004611148565b6040516bffffffffffffffffffffffff19606084811b8216602084015283901b16603482015260009060480160405160208183030381529060405280519060200120905092915050565b61017a610279366004611125565b6107f4565b61014c61028c366004611148565b6001600160a01b03918216600090815260686020908152604080832093909416825291909152205490565b61017a6102c53660046111b8565b6109a7565b61017a6102d83660046111a0565b610abf565b61017a6102eb366004611125565b610b1e565b6103036102fe366004611148565b610c00565b6040805182518152602080840151908201529181015190820152606001610156565b60408051606083811b6bffffffffffffffffffffffff199081166020808501919091523390921b16603483015282516028818403018152604890920183528151918101919091206000818152606b909252919020805484141561039c576000828152606b602052604081208181556001015561040b565b805484106103f15760405162461bcd60e51b815260206004820152601e60248201527f556e6c6f636b2068617320696e73756666696369656e7420616d6f756e74000060448201526064015b60405180910390fd5b83816000016000828254610405919061125c565b90915550505b6104158484610c73565b50505050565b6104258282610c73565b6065546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018490526001600160a01b03909116906323b872dd90606401602060405180830381600087803b15801561049057600080fd5b505af11580156104a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104c89190611180565b505050565b6033546001600160a01b031633146105275760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103e8565b6105316000610ea1565b565b6001600160a01b038116600090815260686020908152604080832033845291829052822054600081116105a85760405162461bcd60e51b815260206004820152601260248201527f4e6f7468696e6720746f20756e7374616b65000000000000000000000000000060448201526064016103e8565b848110156105f85760405162461bcd60e51b815260206004820152601e60248201527f43616e6e6f7420756e6c6f636b206d6f7265207468616e207374616b6564000060448201526064016103e8565b6066546040516398594f9f60e01b81526001600160a01b038681166004830152336024830152909116906398594f9f90604401600060405180830381600087803b15801561064557600080fd5b505af1158015610659573d6000803e3d6000fd5b505050506000606760009054906101000a90046001600160a01b03166001600160a01b03166388c6abf86040518163ffffffff1660e01b815260040160206040518083038186803b1580156106ad57600080fd5b505afa1580156106c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106e59190611208565b9050604051806060016040528087846106fe919061125c565b815243602080830191909152604091820184905233600090815286825282812084518155918401516001808401919091559390920151600290910155908401805488929061074d90849061125c565b909155505060408051606087811b6bffffffffffffffffffffffff199081166020808501919091523390921b16603483015282516028818403018152604890920190925280519101206000906000818152606b60205260408120606a54929350916107b89043611244565b905080826001015410156107ce57600182018190555b888260000160008282546107e29190611244565b90915550909998505050505050505050565b60408051606083811b6bffffffffffffffffffffffff199081166020808501919091523390921b16603483015282516028818403018152604890920183528151918101919091206000818152606b9092529190206001810154431161089b5760405162461bcd60e51b815260206004820152601660248201527f5374616b65206e6f742079657420756e6c6f636b65640000000000000000000060448201526064016103e8565b80546108e95760405162461bcd60e51b815260206004820152601560248201527f4e6f20616d6f756e7420746f207769746864726177000000000000000000000060448201526064016103e8565b80546000838152606b60205260408082208281556001019190915560655490517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018390526001600160a01b039091169063a9059cbb90604401602060405180830381600087803b15801561096857600080fd5b505af115801561097c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109a09190611180565b5050505050565b600054610100900460ff16806109c0575060005460ff16155b610a235760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103e8565b600054610100900460ff16158015610a45576000805461ffff19166101011790555b610a4d610f0b565b606580546001600160a01b038088167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255606680548784169083161790556067805492861692909116919091179055606a82905580156109a0576000805461ff00191690555050505050565b6033546001600160a01b03163314610b195760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103e8565b606a55565b6033546001600160a01b03163314610b785760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103e8565b6001600160a01b038116610bf45760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016103e8565b610bfd81610ea1565b50565b610c2460405180606001604052806000815260200160008152602001600081525090565b506001600160a01b039182166000908152606860209081526040808320939094168252918252829020825160608101845281548152600182015492810192909252600201549181019190915290565b6001600160a01b038116610cc95760405162461bcd60e51b815260206004820152600f60248201527f41646472657373206973206e756c6c000000000000000000000000000000000060448201526064016103e8565b81610d165760405162461bcd60e51b815260206004820152601460248201527f43616e6e6f74207374616b65206e6f7468696e6700000000000000000000000060448201526064016103e8565b6001600160a01b0381166000908152606860209081526040808320338452918290528220549091906066546040516398594f9f60e01b81526001600160a01b0386811660048301523360248301529293509116906398594f9f90604401600060405180830381600087803b158015610d8d57600080fd5b505af1158015610da1573d6000803e3d6000fd5b505050506000606760009054906101000a90046001600160a01b03166001600160a01b03166388c6abf86040518163ffffffff1660e01b815260040160206040518083038186803b158015610df557600080fd5b505afa158015610e09573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e2d9190611208565b905060405180606001604052808684610e469190611244565b8152436020808301919091526040918201849052336000908152868252828120845181559184015160018084019190915593909201516002909101559084018054879290610e95908490611244565b90915550505050505050565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff1680610f24575060005460ff16155b610f875760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103e8565b600054610100900460ff16158015610fa9576000805461ffff19166101011790555b610fb1610fcd565b610fb961107e565b8015610bfd576000805461ff001916905550565b600054610100900460ff1680610fe6575060005460ff16155b6110495760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103e8565b600054610100900460ff16158015610fb9576000805461ffff19166101011790558015610bfd576000805461ff001916905550565b600054610100900460ff1680611097575060005460ff16155b6110fa5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103e8565b600054610100900460ff1615801561111c576000805461ffff19166101011790555b610fb933610ea1565b600060208284031215611136578081fd5b8135611141816112a2565b9392505050565b6000806040838503121561115a578081fd5b8235611165816112a2565b91506020830135611175816112a2565b809150509250929050565b600060208284031215611191578081fd5b81518015158114611141578182fd5b6000602082840312156111b1578081fd5b5035919050565b600080600080608085870312156111cd578182fd5b84356111d8816112a2565b935060208501356111e8816112a2565b925060408501356111f8816112a2565b9396929550929360600135925050565b600060208284031215611219578081fd5b5051919050565b60008060408385031215611232578182fd5b823591506020830135611175816112a2565b6000821982111561125757611257611273565b500190565b60008282101561126e5761126e611273565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6001600160a01b0381168114610bfd57600080fdfea26469706673582212202ceceedf29f8e592862fda578ed36978363b4b8e75ed8026aff44223c2eb8a7564736f6c63430008040033"

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

// GetCurrentStakerAmount is a free data retrieval call binding the contract method 0xc78702bb.
//
// Solidity: function getCurrentStakerAmount(address stakee, address staker) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetCurrentStakerAmount(opts *bind.CallOpts, stakee common.Address, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getCurrentStakerAmount", stakee, staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentStakerAmount is a free data retrieval call binding the contract method 0xc78702bb.
//
// Solidity: function getCurrentStakerAmount(address stakee, address staker) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetCurrentStakerAmount(stakee common.Address, staker common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetCurrentStakerAmount(&_StakingManager.CallOpts, stakee, staker)
}

// GetCurrentStakerAmount is a free data retrieval call binding the contract method 0xc78702bb.
//
// Solidity: function getCurrentStakerAmount(address stakee, address staker) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetCurrentStakerAmount(stakee common.Address, staker common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetCurrentStakerAmount(&_StakingManager.CallOpts, stakee, staker)
}

// GetKey is a free data retrieval call binding the contract method 0xbc298553.
//
// Solidity: function getKey(address stakee, address staker) pure returns(bytes32)
func (_StakingManager *StakingManagerCaller) GetKey(opts *bind.CallOpts, stakee common.Address, staker common.Address) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getKey", stakee, staker)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKey is a free data retrieval call binding the contract method 0xbc298553.
//
// Solidity: function getKey(address stakee, address staker) pure returns(bytes32)
func (_StakingManager *StakingManagerSession) GetKey(stakee common.Address, staker common.Address) ([32]byte, error) {
	return _StakingManager.Contract.GetKey(&_StakingManager.CallOpts, stakee, staker)
}

// GetKey is a free data retrieval call binding the contract method 0xbc298553.
//
// Solidity: function getKey(address stakee, address staker) pure returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) GetKey(stakee common.Address, staker common.Address) ([32]byte, error) {
	return _StakingManager.Contract.GetKey(&_StakingManager.CallOpts, stakee, staker)
}

// GetStakeEntry is a free data retrieval call binding the contract method 0xf731cb49.
//
// Solidity: function getStakeEntry(address stakee, address staker) view returns((uint256,uint256,uint256))
func (_StakingManager *StakingManagerCaller) GetStakeEntry(opts *bind.CallOpts, stakee common.Address, staker common.Address) (StakingManagerStakeEntry, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getStakeEntry", stakee, staker)

	if err != nil {
		return *new(StakingManagerStakeEntry), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingManagerStakeEntry)).(*StakingManagerStakeEntry)

	return out0, err

}

// GetStakeEntry is a free data retrieval call binding the contract method 0xf731cb49.
//
// Solidity: function getStakeEntry(address stakee, address staker) view returns((uint256,uint256,uint256))
func (_StakingManager *StakingManagerSession) GetStakeEntry(stakee common.Address, staker common.Address) (StakingManagerStakeEntry, error) {
	return _StakingManager.Contract.GetStakeEntry(&_StakingManager.CallOpts, stakee, staker)
}

// GetStakeEntry is a free data retrieval call binding the contract method 0xf731cb49.
//
// Solidity: function getStakeEntry(address stakee, address staker) view returns((uint256,uint256,uint256))
func (_StakingManager *StakingManagerCallerSession) GetStakeEntry(stakee common.Address, staker common.Address) (StakingManagerStakeEntry, error) {
	return _StakingManager.Contract.GetStakeEntry(&_StakingManager.CallOpts, stakee, staker)
}

// GetStakeeTotalManagedStake is a free data retrieval call binding the contract method 0x13cdd31b.
//
// Solidity: function getStakeeTotalManagedStake(address stakee) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetStakeeTotalManagedStake(opts *bind.CallOpts, stakee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getStakeeTotalManagedStake", stakee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeeTotalManagedStake is a free data retrieval call binding the contract method 0x13cdd31b.
//
// Solidity: function getStakeeTotalManagedStake(address stakee) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetStakeeTotalManagedStake(stakee common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetStakeeTotalManagedStake(&_StakingManager.CallOpts, stakee)
}

// GetStakeeTotalManagedStake is a free data retrieval call binding the contract method 0x13cdd31b.
//
// Solidity: function getStakeeTotalManagedStake(address stakee) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetStakeeTotalManagedStake(stakee common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetStakeeTotalManagedStake(&_StakingManager.CallOpts, stakee)
}

// GetTotalManagedStake is a free data retrieval call binding the contract method 0x19fba39b.
//
// Solidity: function getTotalManagedStake() view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetTotalManagedStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getTotalManagedStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalManagedStake is a free data retrieval call binding the contract method 0x19fba39b.
//
// Solidity: function getTotalManagedStake() view returns(uint256)
func (_StakingManager *StakingManagerSession) GetTotalManagedStake() (*big.Int, error) {
	return _StakingManager.Contract.GetTotalManagedStake(&_StakingManager.CallOpts)
}

// GetTotalManagedStake is a free data retrieval call binding the contract method 0x19fba39b.
//
// Solidity: function getTotalManagedStake() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetTotalManagedStake() (*big.Int, error) {
	return _StakingManager.Contract.GetTotalManagedStake(&_StakingManager.CallOpts)
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

// TotalManagedStake is a free data retrieval call binding the contract method 0x7ffdacd5.
//
// Solidity: function totalManagedStake() view returns(uint256)
func (_StakingManager *StakingManagerCaller) TotalManagedStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "totalManagedStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalManagedStake is a free data retrieval call binding the contract method 0x7ffdacd5.
//
// Solidity: function totalManagedStake() view returns(uint256)
func (_StakingManager *StakingManagerSession) TotalManagedStake() (*big.Int, error) {
	return _StakingManager.Contract.TotalManagedStake(&_StakingManager.CallOpts)
}

// TotalManagedStake is a free data retrieval call binding the contract method 0x7ffdacd5.
//
// Solidity: function totalManagedStake() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) TotalManagedStake() (*big.Int, error) {
	return _StakingManager.Contract.TotalManagedStake(&_StakingManager.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address token, address rewardsManager, address epochsManager, uint256 _unlockDuration) returns()
func (_StakingManager *StakingManagerTransactor) Initialize(opts *bind.TransactOpts, token common.Address, rewardsManager common.Address, epochsManager common.Address, _unlockDuration *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "initialize", token, rewardsManager, epochsManager, _unlockDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address token, address rewardsManager, address epochsManager, uint256 _unlockDuration) returns()
func (_StakingManager *StakingManagerSession) Initialize(token common.Address, rewardsManager common.Address, epochsManager common.Address, _unlockDuration *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.Initialize(&_StakingManager.TransactOpts, token, rewardsManager, epochsManager, _unlockDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address token, address rewardsManager, address epochsManager, uint256 _unlockDuration) returns()
func (_StakingManager *StakingManagerTransactorSession) Initialize(token common.Address, rewardsManager common.Address, epochsManager common.Address, _unlockDuration *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.Initialize(&_StakingManager.TransactOpts, token, rewardsManager, epochsManager, _unlockDuration)
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
