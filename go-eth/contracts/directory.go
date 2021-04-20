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

// DirectoryNodePointer is an auto generated low-level Go binding around an user-defined struct.
type DirectoryNodePointer struct {
	Value [32]byte
}

// DirectoryABI is the input ABI used to generate the binding from.
const DirectoryABI = "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unlockDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rightAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"value_\",\"type\":\"bytes32\"}],\"internalType\":\"structDirectory.NodePointer\",\"name\":\"parent\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"value_\",\"type\":\"bytes32\"}],\"internalType\":\"structDirectory.NodePointer\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"value_\",\"type\":\"bytes32\"}],\"internalType\":\"structDirectory.NodePointer\",\"name\":\"right\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"unlockings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newUnlockDuration\",\"type\":\"uint256\"}],\"name\":\"setUnlockDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"unlockStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"cancelUnlocking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"point\",\"type\":\"uint128\"}],\"name\":\"scan\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DirectoryBin is the compiled bytecode used for deploying new contracts.
var DirectoryBin = "0x608060405234801561001057600080fd5b5060405161142238038061142283398101604081905261002f916100b9565b60006100426001600160e01b036100b516565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600180546001600160a01b0319166001600160a01b0393909316929092179091556002556100f1565b3390565b600080604083850312156100cb578182fd5b82516001600160a01b03811681146100e1578283fd5b6020939093015192949293505050565b611322806101006000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80638da5cb5b11610097578063c23a5cea11610066578063c23a5cea146101d2578063d86e697d146101e5578063dd9007691461020b578063f2fde38b1461021e576100f5565b80638da5cb5b146101835780639341a5361461018b578063a859f172146101ac578063bc298553146101bf576100f5565b80636b5537e2116100d35780636b5537e21461014b578063715018a61461016b5780637bc74225146101735780638a1fcd601461017b576100f5565b806323314c6c146100fa5780632d49aa1c1461010f5780634e68066114610122575b600080fd5b61010d610108366004610ee7565b610231565b005b61010d61011d366004610ee7565b610258565b610135610130366004610eb7565b610301565b6040516101429190610f44565b60405180910390f35b61015e610159366004610e29565b6103ec565b6040516101429190610f95565b61010d6103fe565b61015e610494565b61015e6104cc565b6101356104d2565b61019e610199366004610e9f565b6104e1565b6040516101429291906112a2565b61015e6101ba366004610ee7565b6104fa565b61015e6101cd366004610e4b565b6107f7565b61010d6101e0366004610e29565b61082a565b6101f86101f3366004610e9f565b610948565b60405161014297969594939291906112b0565b61010d610219366004610e9f565b6109ab565b61010d61022c366004610e29565b6109ef565b600061023d33836107f7565b90506102498184610abc565b6102538383610b18565b505050565b6102628282610b18565b6001546040517f23b872dd0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116906323b872dd906102af90339030908790600401610f58565b602060405180830381600087803b1580156102c957600080fd5b505af11580156102dd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102539190610e7f565b600654600090610313575060006103e7565b60006080836fffffffffffffffffffffffffffffffff16610332610494565b600654910290911c91505b6000818152600360205260409020600181015483101561038257600581015461037757600301546001600160a01b031692506103e7915050565b6005015490506103e2565b600181015481549303928310156103aa57600301546001600160a01b031692506103e7915050565b60068101546103d45760405162461bcd60e51b81526004016103cb906111fd565b60405180910390fd5b600681015490549092039190505b61033d565b919050565b60046020526000908152604090205481565b610406610c2c565b6001600160a01b03166104176104d2565b6001600160a01b03161461043d5760405162461bcd60e51b81526004016103cb906110fd565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36000805473ffffffffffffffffffffffffffffffffffffffff19169055565b6006546000906104a6575060006104c9565b506006546000908152600360205260409020600281015460018201549154909101015b90565b60025481565b6000546001600160a01b031690565b6005602052600090815260409020805460019091015482565b60008061050733846107f7565b6000818152600360205260409020909150610520610e00565b82815281546105415760405162461bcd60e51b81526004016103cb90611234565b81548611156105625760405162461bcd60e51b81526004016103cb906110c6565b61057481600001518388600003610c30565b81546107ba5760028201546001830154016105c057600482015461059c5760006006556105bb565b6004820154600090815260036020526040902081516105bb9190610c67565b610768565b6105c8610e00565b82600201548360010154116105e057826006016105e5565b826005015b6040805160208082018352925480825260009081526003909352912090915061060c610e00565b50815b610617610e00565b826002015483600101541161062f5782600601610634565b826005015b60408051602081019091529054808252909150610651575061066a565b805160009081526003602052604090209250905061060f565b610672610e00565b5060408051602081019091526004830154815281518390610697908290600090610cc5565b6004808801549082015581518814610728576106b887846000015183610d26565b6106c787846000015183610d63565b8151600488018190556000908152600360205260409020835187516106ed929190610cc5565b610707888883600001546000038460040160000154610da0565b8151600090815260036020526040902086516107239190610c67565b610753565b8251600588015414156107445761072387846000015183610d63565b61075387846000015183610d26565b60048101546107625782516006555b50505050505b600083815260036020819052604082208281556001810183905560028101839055908101805473ffffffffffffffffffffffffffffffffffffffff191690556004810182905560058101829055600601555b6000838152600560205260409020600254600182015443909101908111156107e457600182018190555b8154880190915593505050505b92915050565b6000828260405160200161080c929190610f0a565b60405160208183030381529060405280519060200120905092915050565b600061083633836107f7565b6000818152600560205260409020600181015491925090431161086b5760405162461bcd60e51b81526004016103cb9061126b565b80546108895760405162461bcd60e51b81526004016103cb90610f9e565b8054600083815260056020526040808220828155600190810192909255905490517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063a9059cbb906108ef9033908590600401610f7c565b602060405180830381600087803b15801561090957600080fd5b505af115801561091d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109419190610e7f565b5050505050565b60036020818152600092835260409283902080546001820154600283015494830154865180860188526004850154815287518087018952600586015481528851968701909852600690940154855291959094936001600160a01b03909216929187565b6109b3610c2c565b6001600160a01b03166109c46104d2565b6001600160a01b0316146109ea5760405162461bcd60e51b81526004016103cb906110fd565b600255565b6109f7610c2c565b6001600160a01b0316610a086104d2565b6001600160a01b031614610a2e5760405162461bcd60e51b81526004016103cb906110fd565b6001600160a01b038116610a545760405162461bcd60e51b81526004016103cb90611032565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b60008281526005602052604090208054821415610aed57600083815260056020526040812081815560010155610253565b80548210610b0d5760405162461bcd60e51b81526004016103cb9061118f565b805491909103905550565b6001600160a01b038116610b3e5760405162461bcd60e51b81526004016103cb9061108f565b81610b5b5760405162461bcd60e51b81526004016103cb906111c6565b336000610b6882846107f7565b6000818152600360205260409020805491925090610c12576006610b8a610e00565b600081525b815415610bde5750604080516020808201835283548083526000908152600390915291909120600281015481546001830154910111610bd15780600601610bd6565b806005015b925050610b8f565b51600483015582905560038101805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0386161790555b6004810154610c215760068290555b610941828287610c30565b3390565b81548101825560038201546001600160a01b0316600090815260046020526040812080548301905561025390849084908490610da0565b6005820154811415610c7857610cc1565b6006820154811415610c905760006006830155610cc1565b6005820154811480610ca55750600682015481145b610cc15760405162461bcd60e51b81526004016103cb90610fd5565b5050565b6005830154821415610cdd5760058301819055610253565b6006830154821415610cf55760068301819055610253565b6005830154821480610d0a5750600683015482145b6102535760405162461bcd60e51b81526004016103cb90611132565b6005830154610d3457610253565b600580840180546000908152600360205260409020600401849055549082015560018084015490820155505050565b6006830154610d7157610253565b600680840180546000908152600360205260409020600401849055549082015560028084015490820155505050565b600483018054821415610db35750610dfa565b805460009081526003602052604090206005810154861415610dde5760018101805485019055610de9565b600281018054850190555b8154610df790828686610da0565b50505b50505050565b60408051602081019091526000815290565b80356001600160a01b03811681146107f157600080fd5b600060208284031215610e3a578081fd5b610e448383610e12565b9392505050565b60008060408385031215610e5d578081fd5b610e678484610e12565b9150610e768460208501610e12565b90509250929050565b600060208284031215610e90578081fd5b81518015158114610e44578182fd5b600060208284031215610eb0578081fd5b5035919050565b600060208284031215610ec8578081fd5b81356fffffffffffffffffffffffffffffffff81168114610e44578182fd5b60008060408385031215610ef9578182fd5b82359150610e768460208501610e12565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606093841b811682529190921b16601482015260280190565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b90815260200190565b60208082526015908201527f4e6f20616d6f756e7420746f2077697468647261770000000000000000000000604082015260600190565b6020808252602f908201527f4f6c64206368696c642063616e6e6f742062652072656d6f766564202d20697460408201527f20646f6573206e6f742065786973740000000000000000000000000000000000606082015260800190565b60208082526026908201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408201527f6464726573730000000000000000000000000000000000000000000000000000606082015260800190565b6020808252600f908201527f41646472657373206973206e756c6c0000000000000000000000000000000000604082015260600190565b6020808252601e908201527f43616e6e6f7420756e6c6f636b206d6f7265207468616e207374616b65640000604082015260600190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252602f908201527f4f6c64206368696c642063616e6e6f74206265206368616e676564202d20697460408201527f20646f6573206e6f742065786973740000000000000000000000000000000000606082015260800190565b6020808252601e908201527f556e6c6f636b2068617320696e73756666696369656e7420616d6f756e740000604082015260600190565b60208082526014908201527f43616e6e6f74207374616b65206e6f7468696e67000000000000000000000000604082015260600190565b60208082526019908201527f6d697373696e67206e6f6465206f6e2074686520726967687400000000000000604082015260600190565b60208082526012908201527f4e6f7468696e6720746f20756e7374616b650000000000000000000000000000604082015260600190565b60208082526016908201527f5374616b65206e6f742079657420756e6c6f636b656400000000000000000000604082015260600190565b918252602082015260400190565b968752602087019590955260408601939093526001600160a01b039190911660608501525160808401525160a08301525160c082015260e0019056fea26469706673582212205c59c7c063a7097c8fe4fc472c4501788d8b1a749d51f8da19b537d21ef051ca64736f6c63430006040033"

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

// GetKey is a free data retrieval call binding the contract method 0xbc298553.
//
// Solidity: function getKey(address staker, address stakee) pure returns(bytes32)
func (_Directory *DirectoryCaller) GetKey(opts *bind.CallOpts, staker common.Address, stakee common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "getKey", staker, stakee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKey is a free data retrieval call binding the contract method 0xbc298553.
//
// Solidity: function getKey(address staker, address stakee) pure returns(bytes32)
func (_Directory *DirectorySession) GetKey(staker common.Address, stakee common.Address) ([32]byte, error) {
	return _Directory.Contract.GetKey(&_Directory.CallOpts, staker, stakee)
}

// GetKey is a free data retrieval call binding the contract method 0xbc298553.
//
// Solidity: function getKey(address staker, address stakee) pure returns(bytes32)
func (_Directory *DirectoryCallerSession) GetKey(staker common.Address, stakee common.Address) ([32]byte, error) {
	return _Directory.Contract.GetKey(&_Directory.CallOpts, staker, stakee)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() view returns(uint256)
func (_Directory *DirectoryCaller) GetTotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "getTotalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() view returns(uint256)
func (_Directory *DirectorySession) GetTotalStake() (*big.Int, error) {
	return _Directory.Contract.GetTotalStake(&_Directory.CallOpts)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() view returns(uint256)
func (_Directory *DirectoryCallerSession) GetTotalStake() (*big.Int, error) {
	return _Directory.Contract.GetTotalStake(&_Directory.CallOpts)
}

// Nodes is a free data retrieval call binding the contract method 0xd86e697d.
//
// Solidity: function nodes(bytes32 ) view returns(uint256 amount, uint256 leftAmount, uint256 rightAmount, address stakee, (bytes32) parent, (bytes32) left, (bytes32) right)
func (_Directory *DirectoryCaller) Nodes(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Amount      *big.Int
	LeftAmount  *big.Int
	RightAmount *big.Int
	Stakee      common.Address
	Parent      DirectoryNodePointer
	Left        DirectoryNodePointer
	Right       DirectoryNodePointer
}, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "nodes", arg0)

	outstruct := new(struct {
		Amount      *big.Int
		LeftAmount  *big.Int
		RightAmount *big.Int
		Stakee      common.Address
		Parent      DirectoryNodePointer
		Left        DirectoryNodePointer
		Right       DirectoryNodePointer
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LeftAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RightAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Stakee = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Parent = *abi.ConvertType(out[4], new(DirectoryNodePointer)).(*DirectoryNodePointer)
	outstruct.Left = *abi.ConvertType(out[5], new(DirectoryNodePointer)).(*DirectoryNodePointer)
	outstruct.Right = *abi.ConvertType(out[6], new(DirectoryNodePointer)).(*DirectoryNodePointer)

	return *outstruct, err

}

// Nodes is a free data retrieval call binding the contract method 0xd86e697d.
//
// Solidity: function nodes(bytes32 ) view returns(uint256 amount, uint256 leftAmount, uint256 rightAmount, address stakee, (bytes32) parent, (bytes32) left, (bytes32) right)
func (_Directory *DirectorySession) Nodes(arg0 [32]byte) (struct {
	Amount      *big.Int
	LeftAmount  *big.Int
	RightAmount *big.Int
	Stakee      common.Address
	Parent      DirectoryNodePointer
	Left        DirectoryNodePointer
	Right       DirectoryNodePointer
}, error) {
	return _Directory.Contract.Nodes(&_Directory.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0xd86e697d.
//
// Solidity: function nodes(bytes32 ) view returns(uint256 amount, uint256 leftAmount, uint256 rightAmount, address stakee, (bytes32) parent, (bytes32) left, (bytes32) right)
func (_Directory *DirectoryCallerSession) Nodes(arg0 [32]byte) (struct {
	Amount      *big.Int
	LeftAmount  *big.Int
	RightAmount *big.Int
	Stakee      common.Address
	Parent      DirectoryNodePointer
	Left        DirectoryNodePointer
	Right       DirectoryNodePointer
}, error) {
	return _Directory.Contract.Nodes(&_Directory.CallOpts, arg0)
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

// Stakees is a free data retrieval call binding the contract method 0x6b5537e2.
//
// Solidity: function stakees(address ) view returns(uint256)
func (_Directory *DirectoryCaller) Stakees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "stakees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Stakees is a free data retrieval call binding the contract method 0x6b5537e2.
//
// Solidity: function stakees(address ) view returns(uint256)
func (_Directory *DirectorySession) Stakees(arg0 common.Address) (*big.Int, error) {
	return _Directory.Contract.Stakees(&_Directory.CallOpts, arg0)
}

// Stakees is a free data retrieval call binding the contract method 0x6b5537e2.
//
// Solidity: function stakees(address ) view returns(uint256)
func (_Directory *DirectoryCallerSession) Stakees(arg0 common.Address) (*big.Int, error) {
	return _Directory.Contract.Stakees(&_Directory.CallOpts, arg0)
}

// UnlockDuration is a free data retrieval call binding the contract method 0x8a1fcd60.
//
// Solidity: function unlockDuration() view returns(uint256)
func (_Directory *DirectoryCaller) UnlockDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "unlockDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockDuration is a free data retrieval call binding the contract method 0x8a1fcd60.
//
// Solidity: function unlockDuration() view returns(uint256)
func (_Directory *DirectorySession) UnlockDuration() (*big.Int, error) {
	return _Directory.Contract.UnlockDuration(&_Directory.CallOpts)
}

// UnlockDuration is a free data retrieval call binding the contract method 0x8a1fcd60.
//
// Solidity: function unlockDuration() view returns(uint256)
func (_Directory *DirectoryCallerSession) UnlockDuration() (*big.Int, error) {
	return _Directory.Contract.UnlockDuration(&_Directory.CallOpts)
}

// Unlockings is a free data retrieval call binding the contract method 0x9341a536.
//
// Solidity: function unlockings(bytes32 ) view returns(uint256 amount, uint256 unlockAt)
func (_Directory *DirectoryCaller) Unlockings(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Amount   *big.Int
	UnlockAt *big.Int
}, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "unlockings", arg0)

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
func (_Directory *DirectorySession) Unlockings(arg0 [32]byte) (struct {
	Amount   *big.Int
	UnlockAt *big.Int
}, error) {
	return _Directory.Contract.Unlockings(&_Directory.CallOpts, arg0)
}

// Unlockings is a free data retrieval call binding the contract method 0x9341a536.
//
// Solidity: function unlockings(bytes32 ) view returns(uint256 amount, uint256 unlockAt)
func (_Directory *DirectoryCallerSession) Unlockings(arg0 [32]byte) (struct {
	Amount   *big.Int
	UnlockAt *big.Int
}, error) {
	return _Directory.Contract.Unlockings(&_Directory.CallOpts, arg0)
}

// AddStake is a paid mutator transaction binding the contract method 0x2d49aa1c.
//
// Solidity: function addStake(uint256 amount, address stakee) returns()
func (_Directory *DirectoryTransactor) AddStake(opts *bind.TransactOpts, amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "addStake", amount, stakee)
}

// AddStake is a paid mutator transaction binding the contract method 0x2d49aa1c.
//
// Solidity: function addStake(uint256 amount, address stakee) returns()
func (_Directory *DirectorySession) AddStake(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.AddStake(&_Directory.TransactOpts, amount, stakee)
}

// AddStake is a paid mutator transaction binding the contract method 0x2d49aa1c.
//
// Solidity: function addStake(uint256 amount, address stakee) returns()
func (_Directory *DirectoryTransactorSession) AddStake(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.AddStake(&_Directory.TransactOpts, amount, stakee)
}

// CancelUnlocking is a paid mutator transaction binding the contract method 0x23314c6c.
//
// Solidity: function cancelUnlocking(uint256 amount, address stakee) returns()
func (_Directory *DirectoryTransactor) CancelUnlocking(opts *bind.TransactOpts, amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "cancelUnlocking", amount, stakee)
}

// CancelUnlocking is a paid mutator transaction binding the contract method 0x23314c6c.
//
// Solidity: function cancelUnlocking(uint256 amount, address stakee) returns()
func (_Directory *DirectorySession) CancelUnlocking(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.CancelUnlocking(&_Directory.TransactOpts, amount, stakee)
}

// CancelUnlocking is a paid mutator transaction binding the contract method 0x23314c6c.
//
// Solidity: function cancelUnlocking(uint256 amount, address stakee) returns()
func (_Directory *DirectoryTransactorSession) CancelUnlocking(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.CancelUnlocking(&_Directory.TransactOpts, amount, stakee)
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

// UnlockStake is a paid mutator transaction binding the contract method 0xa859f172.
//
// Solidity: function unlockStake(uint256 amount, address stakee) returns(uint256)
func (_Directory *DirectoryTransactor) UnlockStake(opts *bind.TransactOpts, amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "unlockStake", amount, stakee)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xa859f172.
//
// Solidity: function unlockStake(uint256 amount, address stakee) returns(uint256)
func (_Directory *DirectorySession) UnlockStake(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.UnlockStake(&_Directory.TransactOpts, amount, stakee)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xa859f172.
//
// Solidity: function unlockStake(uint256 amount, address stakee) returns(uint256)
func (_Directory *DirectoryTransactorSession) UnlockStake(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.UnlockStake(&_Directory.TransactOpts, amount, stakee)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address stakee) returns()
func (_Directory *DirectoryTransactor) WithdrawStake(opts *bind.TransactOpts, stakee common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "withdrawStake", stakee)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address stakee) returns()
func (_Directory *DirectorySession) WithdrawStake(stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.WithdrawStake(&_Directory.TransactOpts, stakee)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address stakee) returns()
func (_Directory *DirectoryTransactorSession) WithdrawStake(stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.WithdrawStake(&_Directory.TransactOpts, stakee)
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
