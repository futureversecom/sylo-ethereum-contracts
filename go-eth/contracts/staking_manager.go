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
	Amount *big.Int
	Block  *big.Int
}

// StakingManagerABI is the input ABI used to generate the binding from.
const StakingManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"unlockings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unlockDuration\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newUnlockDuration\",\"type\":\"uint256\"}],\"name\":\"setUnlockDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"unlockStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"cancelUnlocking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getStakeEntries\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"internalType\":\"structStakingManager.StakeEntry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getStakeeTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getStakerAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getCurrentStakerAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StakingManagerBin is the compiled bytecode used for deploying new contracts.
var StakingManagerBin = "0x608060405234801561001057600080fd5b50611426806100206000396000f3fe608060405234801561001057600080fd5b50600436106101365760003560e01c8063a859f172116100b2578063c23a5cea11610081578063cd6dc68711610066578063cd6dc687146102df578063dd900769146102f2578063f2fde38b1461030557600080fd5b8063c23a5cea146102b9578063c78702bb146102cc57600080fd5b8063a859f1721461021b578063b91cefd21461022e578063bc29855314610241578063bd3e20a71461029957600080fd5b80637bc74225116101095780638b0e9f3f116100ee5780638b0e9f3f146101bb5780638da5cb5b146101c45780639341a536146101df57600080fd5b80637bc74225146101aa5780638a1fcd60146101b257600080fd5b806323314c6c1461013b5780632d49aa1c1461015057806364084d4e14610163578063715018a6146101a2575b600080fd5b61014e610149366004611303565b610318565b005b61014e61015e366004611303565b61041d565b61018f61017136600461120c565b6001600160a01b031660009081526066602052604090206001015490565b6040519081526020015b60405180910390f35b61014e6104cf565b60675461018f565b61018f60685481565b61018f60675481565b6033546040516001600160a01b039091168152602001610199565b6102066101ed3660046112c0565b6069602052600090815260409020805460019091015482565b60408051928352602083019190915201610199565b61018f610229366004611303565b610535565b61018f61023c366004611260565b610706565b61018f61024f366004611228565b6040516bffffffffffffffffffffffff19606084811b8216602084015283901b16603482015260009060480160405160208183030381529060405280519060200120905092915050565b6102ac6102a7366004611228565b6109b6565b6040516101999190611327565b61014e6102c736600461120c565b610a4c565b61018f6102da366004611228565b610c0e565b61014e6102ed3660046112d8565b610c1b565b61014e6103003660046112c0565b610d0f565b61014e61031336600461120c565b610d6e565b600061036933836040516bffffffffffffffffffffffff19606084811b8216602084015283901b16603482015260009060480160405160208183030381529060405280519060200120905092915050565b600081815260696020526040902080549192509084141561039e5760008281526069602052604081208181556001015561040d565b805484106103f35760405162461bcd60e51b815260206004820152601e60248201527f556e6c6f636b2068617320696e73756666696369656e7420616d6f756e74000060448201526064015b60405180910390fd5b8381600001600082825461040791906113ae565b90915550505b6104178484610e50565b50505050565b6104278282610e50565b6065546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018490526001600160a01b03909116906323b872dd90606401602060405180830381600087803b15801561049257600080fd5b505af11580156104a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ca91906112a0565b505050565b6033546001600160a01b031633146105295760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103ea565b6105336000610f88565b565b6001600160a01b0381166000908152606660205260408120816105583385610c0e565b9050600081116105aa5760405162461bcd60e51b815260206004820152601260248201527f4e6f7468696e6720746f20756e7374616b65000000000000000000000000000060448201526064016103ea565b848110156105fa5760405162461bcd60e51b815260206004820152601e60248201527f43616e6e6f7420756e6c6f636b206d6f7265207468616e207374616b6564000060448201526064016103ea565b33600090815260208390526040908190208151808301909252908061061f88856113ae565b815243602091820152825460018181018555600094855282852084516002909302019182559290910151908201558301805487929061065f9084906113ae565b90915550506040805133606090811b6bffffffffffffffffffffffff199081166020808501919091529188901b16603483015282516028818403018152604890920190925280519101206000906000818152606960205260408120606854929350916106cb9043611376565b905080826001015410156106e157600182018190555b878260000160008282546106f59190611376565b909155509098975050505050505050565b6001600160a01b0380831660009081526066602090815260408083209387168352929052908120548061073d5760009150506109af565b6001600160a01b0380851660009081526066602090815260408083209389168352929052908120548190610773906001906113ae565b6001600160a01b038088166000908152606660209081526040808320938c16835292905290812080549293509091839081106107bf57634e487b7160e01b600052603260045260246000fd5b6000918252602091829020604080518082019091526002909202018054825260010154918101829052915086106107fc575193506109af92505050565b8183116109a657600060026108118486611376565b61081b919061138e565b90506000811561088a576001600160a01b03808a166000908152606660209081526040808320938e168352929052206108556001846113ae565b8154811061087357634e487b7160e01b600052603260045260246000fd5b90600052602060002090600202016001015461088d565b60005b6001600160a01b03808b166000908152606660209081526040808320938f16835292905290812080549293509091849081106108d957634e487b7160e01b600052603260045260246000fd5b90600052602060002090600202016001015490508189101580156108fc57508089105b1561096f576001600160a01b03808b166000908152606660209081526040808320938f168352929052206109316001856113ae565b8154811061094f57634e487b7160e01b600052603260045260246000fd5b9060005260206000209060020201600001549750505050505050506109af565b81891015610989576109826001846113ae565b945061099e565b80891061099e5761099b836001611376565b95505b5050506107fc565b60009450505050505b9392505050565b6001600160a01b0380821660009081526066602090815260408083209386168352928152828220805484518184028101840190955280855260609493919290919084015b82821015610a40578382906000526020600020906002020160405180604001604052908160008201548152602001600182015481525050815260200190600101906109fa565b50505050905092915050565b6000610a9d33836040516bffffffffffffffffffffffff19606084811b8216602084015283901b16603482015260009060480160405160208183030381529060405280519060200120905092915050565b60008181526069602052604090206001810154919250904311610b025760405162461bcd60e51b815260206004820152601660248201527f5374616b65206e6f742079657420756e6c6f636b65640000000000000000000060448201526064016103ea565b8054610b505760405162461bcd60e51b815260206004820152601560248201527f4e6f20616d6f756e7420746f207769746864726177000000000000000000000060448201526064016103ea565b80546000838152606960205260408082208281556001019190915560655490517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018390526001600160a01b039091169063a9059cbb90604401602060405180830381600087803b158015610bcf57600080fd5b505af1158015610be3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c0791906112a0565b5050505050565b60006109af838343610706565b600054610100900460ff1680610c34575060005460ff16155b610c975760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103ea565b600054610100900460ff16158015610cb9576000805461ffff19166101011790555b610cc1610ff2565b606580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038516179055606882905580156104ca576000805461ff0019169055505050565b6033546001600160a01b03163314610d695760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103ea565b606855565b6033546001600160a01b03163314610dc85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103ea565b6001600160a01b038116610e445760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016103ea565b610e4d81610f88565b50565b6001600160a01b038116610ea65760405162461bcd60e51b815260206004820152600f60248201527f41646472657373206973206e756c6c000000000000000000000000000000000060448201526064016103ea565b81610ef35760405162461bcd60e51b815260206004820152601460248201527f43616e6e6f74207374616b65206e6f7468696e6700000000000000000000000060448201526064016103ea565b6001600160a01b038116600090815260666020526040812090610f163384610c0e565b3360009081526020849052604090819020815180830190925291925080610f3d8785611376565b8152436020918201528254600181810185556000948552828520845160029093020191825592909101519082015583018054869290610f7d908490611376565b909155505050505050565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff168061100b575060005460ff16155b61106e5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103ea565b600054610100900460ff16158015611090576000805461ffff19166101011790555b6110986110b4565b6110a0611165565b8015610e4d576000805461ff001916905550565b600054610100900460ff16806110cd575060005460ff16155b6111305760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103ea565b600054610100900460ff161580156110a0576000805461ffff19166101011790558015610e4d576000805461ff001916905550565b600054610100900460ff168061117e575060005460ff16155b6111e15760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103ea565b600054610100900460ff16158015611203576000805461ffff19166101011790555b6110a033610f88565b60006020828403121561121d578081fd5b81356109af816113db565b6000806040838503121561123a578081fd5b8235611245816113db565b91506020830135611255816113db565b809150509250929050565b600080600060608486031215611274578081fd5b833561127f816113db565b9250602084013561128f816113db565b929592945050506040919091013590565b6000602082840312156112b1578081fd5b815180151581146109af578182fd5b6000602082840312156112d1578081fd5b5035919050565b600080604083850312156112ea578182fd5b82356112f5816113db565b946020939093013593505050565b60008060408385031215611315578182fd5b823591506020830135611255816113db565b602080825282518282018190526000919060409081850190868401855b8281101561136957815180518552860151868501529284019290850190600101611344565b5091979650505050505050565b60008219821115611389576113896113c5565b500190565b6000826113a957634e487b7160e01b81526012600452602481fd5b500490565b6000828210156113c0576113c06113c5565b500390565b634e487b7160e01b600052601160045260246000fd5b6001600160a01b0381168114610e4d57600080fdfea2646970667358221220552e9fb83078448d117f2ce636d5c1844bed165006c291aa7c4ebc0add46b29564736f6c63430008040033"

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
// Solidity: function getCurrentStakerAmount(address staker, address stakee) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetCurrentStakerAmount(opts *bind.CallOpts, staker common.Address, stakee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getCurrentStakerAmount", staker, stakee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentStakerAmount is a free data retrieval call binding the contract method 0xc78702bb.
//
// Solidity: function getCurrentStakerAmount(address staker, address stakee) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetCurrentStakerAmount(staker common.Address, stakee common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetCurrentStakerAmount(&_StakingManager.CallOpts, staker, stakee)
}

// GetCurrentStakerAmount is a free data retrieval call binding the contract method 0xc78702bb.
//
// Solidity: function getCurrentStakerAmount(address staker, address stakee) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetCurrentStakerAmount(staker common.Address, stakee common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetCurrentStakerAmount(&_StakingManager.CallOpts, staker, stakee)
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

// GetStakeEntries is a free data retrieval call binding the contract method 0xbd3e20a7.
//
// Solidity: function getStakeEntries(address staker, address stakee) view returns((uint256,uint256)[])
func (_StakingManager *StakingManagerCaller) GetStakeEntries(opts *bind.CallOpts, staker common.Address, stakee common.Address) ([]StakingManagerStakeEntry, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getStakeEntries", staker, stakee)

	if err != nil {
		return *new([]StakingManagerStakeEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingManagerStakeEntry)).(*[]StakingManagerStakeEntry)

	return out0, err

}

// GetStakeEntries is a free data retrieval call binding the contract method 0xbd3e20a7.
//
// Solidity: function getStakeEntries(address staker, address stakee) view returns((uint256,uint256)[])
func (_StakingManager *StakingManagerSession) GetStakeEntries(staker common.Address, stakee common.Address) ([]StakingManagerStakeEntry, error) {
	return _StakingManager.Contract.GetStakeEntries(&_StakingManager.CallOpts, staker, stakee)
}

// GetStakeEntries is a free data retrieval call binding the contract method 0xbd3e20a7.
//
// Solidity: function getStakeEntries(address staker, address stakee) view returns((uint256,uint256)[])
func (_StakingManager *StakingManagerCallerSession) GetStakeEntries(staker common.Address, stakee common.Address) ([]StakingManagerStakeEntry, error) {
	return _StakingManager.Contract.GetStakeEntries(&_StakingManager.CallOpts, staker, stakee)
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
