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

// DirectoryABI is the input ABI used to generate the binding from.
const DirectoryABI = "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unlockDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rightAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"parent\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"left\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"right\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"unlockings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newUnlockDuration\",\"type\":\"uint256\"}],\"name\":\"setUnlockDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"unlockStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"lockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"point\",\"type\":\"uint256\"}],\"name\":\"scan\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DirectoryBin is the compiled bytecode used for deploying new contracts.
var DirectoryBin = "0x608060405234801561001057600080fd5b5060405161110438038061110483398101604081905261002f916100b9565b60006100426001600160e01b036100b516565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600180546001600160a01b0319166001600160a01b0393909316929092179091556002556100f1565b3390565b600080604083850312156100cb578182fd5b82516001600160a01b03811681146100e1578283fd5b6020939093015192949293505050565b611004806101006000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80638fee640711610097578063dd90076911610066578063dd900769146101e5578063e16d14a0146101f8578063f2888dbb1461020b578063f2fde38b1461021e576100f5565b80638fee6407146101785780639341a5361461019e578063a859f172146101bf578063bc298553146101d2576100f5565b8063715018a6116100d3578063715018a61461014b5780637bc74225146101535780638a1fcd601461015b5780638da5cb5b14610163576100f5565b80632d49aa1c146100fa57806333f4909f1461010f5780636b5537e214610122575b600080fd5b61010d610108366004610cbd565b610231565b005b61010d61011d366004610cbd565b6103fc565b610135610130366004610c2f565b610434565b6040516101429190610d6b565b60405180910390f35b61010d610446565b6101356104d2565b61013561050a565b61016b610510565b6040516101429190610d1a565b61018b610186366004610ca5565b61051f565b6040516101429796959493929190610f95565b6101b16101ac366004610ca5565b610567565b604051610142929190610f87565b6101356101cd366004610cbd565b610580565b6101356101e0366004610c51565b6107f3565b61010d6101f3366004610ca5565b610826565b61016b610206366004610ca5565b610860565b61010d610219366004610c2f565b6108f3565b61010d61022c366004610c2f565b6109b8565b6001600160a01b0381166102605760405162461bcd60e51b815260040161025790610dd1565b60405180910390fd5b8161027d5760405162461bcd60e51b815260040161025790610ee2565b600061028933836107f3565b600081815260036020526040902080549192509061033c5760065460008181526003602052604090205b81156102ff57600081600201548260010154106102d45781600601546102da565b81600501545b9050806102e757506102ff565b600081815260036020526040902090925090506102b3565b61030b816000866109f9565b50600482015560038101805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0385161790555b600481015461034b5760068290555b610356828286610a1e565b6001546040517f23b872dd0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116906323b872dd906103a390339030908990600401610d2e565b602060405180830381600087803b1580156103bd57600080fd5b505af11580156103d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103f59190610c85565b5050505050565b600061040833836107f3565b60008181526003602052604090209091506104238285610a55565b61042e828286610a1e565b50505050565b60046020526000908152604090205481565b61044e610ab1565b6000546001600160a01b0390811691161461047b5760405162461bcd60e51b815260040161025790610e76565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36000805473ffffffffffffffffffffffffffffffffffffffff19169055565b6006546000906104e457506000610507565b506006546000908152600360205260409020600281015460018201549154909101015b90565b60025481565b6000546001600160a01b031690565b600360208190526000918252604090912080546001820154600283015493830154600484015460058501546006909501549395929492936001600160a01b0390921692909187565b6005602052600090815260409020805460019091015482565b60008061058d33846107f3565b60008181526003602052604090208054919250906105bd5760405162461bcd60e51b815260040161025790610f19565b80548511156105de5760405162461bcd60e51b815260040161025790610e08565b6105ec828287600003610a1e565b80546107b7576000816002015482600101541161060d578160060154610613565b81600501545b600483015460009081526003602052604090209091508161064e5761063a818560006109f9565b60048301546106495760006006555b610763565b60008281526003602052604090205b6000816002015482600101541161067857816006015461067e565b81600501545b90508061068b57506106a3565b6000818152600360205260409020909350905061065d565b60048101546106b38387866109f9565b6004808601549083015580861461072f576106cf858584610ab5565b6106da858584610af2565b6004850181905560008181526003602052604090206106fa9085886109f9565b610711868684600001546000038560040154610b2f565b600081815260036020526040812061072a9188906109f9565b610751565b83856005015414156107465761072a858584610af2565b610751858584610ab5565b60048201546107605760068490555b50505b5050600082815260036020819052604082208281556001810183905560028101839055908101805473ffffffffffffffffffffffffffffffffffffffff191690556004810182905560058101829055600601555b6000828152600560205260409020600254600182015443909101908111156107e157600182018190555b81548701909155925050505b92915050565b60008282604051602001610808929190610ce0565b60405160208183030381529060405280519060200120905092915050565b61082e610ab1565b6000546001600160a01b0390811691161461085b5760405162461bcd60e51b815260040161025790610e76565b600255565b600654600090610872575060006108ee565b600060fd8361087f6104d2565b028161088757fe5b60065491900491505b600081815260036020526040902060018101548310156108b557600501549050610890565b6001810154815493039283116108dc57600301546001600160a01b031692506108ee915050565b80546006909101549203919050610890565b919050565b60006108ff33836107f3565b600081815260056020526040902060018101549192509043116109345760405162461bcd60e51b815260040161025790610f50565b80546109525760405162461bcd60e51b815260040161025790610e3f565b8054600083815260056020526040808220828155600190810192909255905490517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063a9059cbb906103a39033908590600401610d52565b6109c0610ab1565b6000546001600160a01b039081169116146109ed5760405162461bcd60e51b815260040161025790610e76565b6109f681610b8a565b50565b8183600501541415610a115760058301819055610a19565b600683018190555b505050565b81548101825560038201546001600160a01b03166000908152600460205260408120805483019055610a1990849084908490610b2f565b60008281526005602052604090208054821415610a8657600083815260056020526040812081815560010155610a19565b80548210610aa65760405162461bcd60e51b815260040161025790610eab565b805491909103905550565b3390565b6005830154610ac357610a19565b600580840180546000908152600360205260409020600401849055549082015560018084015490820155505050565b6006830154610b0057610a19565b600680840180546000908152600360205260409020600401849055549082015560028084015490820155505050565b600483015481811415610b42575061042e565b60008181526003602052604090206005810154861415610b6b5760018101805485019055610b76565b600281018054850190555b610b8282828686610b2f565b505050505050565b6001600160a01b038116610bb05760405162461bcd60e51b815260040161025790610d74565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b80356001600160a01b03811681146107ed57600080fd5b600060208284031215610c40578081fd5b610c4a8383610c18565b9392505050565b60008060408385031215610c63578081fd5b610c6d8484610c18565b9150610c7c8460208501610c18565b90509250929050565b600060208284031215610c96578081fd5b81518015158114610c4a578182fd5b600060208284031215610cb6578081fd5b5035919050565b60008060408385031215610ccf578182fd5b82359150610c7c8460208501610c18565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606093841b811682529190921b16601482015260280190565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b90815260200190565b60208082526026908201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408201527f6464726573730000000000000000000000000000000000000000000000000000606082015260800190565b6020808252600f908201527f41646472657373206973206e756c6c0000000000000000000000000000000000604082015260600190565b6020808252601e908201527f43616e6e6f7420756e6c6f636b206d6f7265207468616e207374616b65640000604082015260600190565b60208082526013908201527f4e6f20616d6f756e7420746f20756e6c6f636b00000000000000000000000000604082015260600190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252601e908201527f556e6c6f636b2068617320696e73756666696369656e7420616d6f756e740000604082015260600190565b60208082526014908201527f43616e6e6f74207374616b65206e6f7468696e67000000000000000000000000604082015260600190565b60208082526012908201527f4e6f7468696e6720746f20756e7374616b650000000000000000000000000000604082015260600190565b60208082526016908201527f5374616b65206e6f742079657420756e6c6f636b656400000000000000000000604082015260600190565b918252602082015260400190565b968752602087019590955260408601939093526001600160a01b03919091166060850152608084015260a083015260c082015260e0019056fea2646970667358221220589a4fc41355e51e02e8e1f583a110f6353cea621c7e89bf16864f348f0a3f5464736f6c63430006040033"

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

// Scan is a free data retrieval call binding the contract method 0xe16d14a0.
//
// Solidity: function scan(uint256 point) view returns(address)
func (_Directory *DirectoryCaller) Scan(opts *bind.CallOpts, point *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "scan", point)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Scan is a free data retrieval call binding the contract method 0xe16d14a0.
//
// Solidity: function scan(uint256 point) view returns(address)
func (_Directory *DirectorySession) Scan(point *big.Int) (common.Address, error) {
	return _Directory.Contract.Scan(&_Directory.CallOpts, point)
}

// Scan is a free data retrieval call binding the contract method 0xe16d14a0.
//
// Solidity: function scan(uint256 point) view returns(address)
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

// Stakes is a free data retrieval call binding the contract method 0x8fee6407.
//
// Solidity: function stakes(bytes32 ) view returns(uint256 amount, uint256 leftAmount, uint256 rightAmount, address stakee, bytes32 parent, bytes32 left, bytes32 right)
func (_Directory *DirectoryCaller) Stakes(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Amount      *big.Int
	LeftAmount  *big.Int
	RightAmount *big.Int
	Stakee      common.Address
	Parent      [32]byte
	Left        [32]byte
	Right       [32]byte
}, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "stakes", arg0)

	outstruct := new(struct {
		Amount      *big.Int
		LeftAmount  *big.Int
		RightAmount *big.Int
		Stakee      common.Address
		Parent      [32]byte
		Left        [32]byte
		Right       [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = out[0].(*big.Int)
	outstruct.LeftAmount = out[1].(*big.Int)
	outstruct.RightAmount = out[2].(*big.Int)
	outstruct.Stakee = out[3].(common.Address)
	outstruct.Parent = out[4].([32]byte)
	outstruct.Left = out[5].([32]byte)
	outstruct.Right = out[6].([32]byte)

	return *outstruct, err

}

// Stakes is a free data retrieval call binding the contract method 0x8fee6407.
//
// Solidity: function stakes(bytes32 ) view returns(uint256 amount, uint256 leftAmount, uint256 rightAmount, address stakee, bytes32 parent, bytes32 left, bytes32 right)
func (_Directory *DirectorySession) Stakes(arg0 [32]byte) (struct {
	Amount      *big.Int
	LeftAmount  *big.Int
	RightAmount *big.Int
	Stakee      common.Address
	Parent      [32]byte
	Left        [32]byte
	Right       [32]byte
}, error) {
	return _Directory.Contract.Stakes(&_Directory.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x8fee6407.
//
// Solidity: function stakes(bytes32 ) view returns(uint256 amount, uint256 leftAmount, uint256 rightAmount, address stakee, bytes32 parent, bytes32 left, bytes32 right)
func (_Directory *DirectoryCallerSession) Stakes(arg0 [32]byte) (struct {
	Amount      *big.Int
	LeftAmount  *big.Int
	RightAmount *big.Int
	Stakee      common.Address
	Parent      [32]byte
	Left        [32]byte
	Right       [32]byte
}, error) {
	return _Directory.Contract.Stakes(&_Directory.CallOpts, arg0)
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

	outstruct.Amount = out[0].(*big.Int)
	outstruct.UnlockAt = out[1].(*big.Int)

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

// LockStake is a paid mutator transaction binding the contract method 0x33f4909f.
//
// Solidity: function lockStake(uint256 amount, address stakee) returns()
func (_Directory *DirectoryTransactor) LockStake(opts *bind.TransactOpts, amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "lockStake", amount, stakee)
}

// LockStake is a paid mutator transaction binding the contract method 0x33f4909f.
//
// Solidity: function lockStake(uint256 amount, address stakee) returns()
func (_Directory *DirectorySession) LockStake(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.LockStake(&_Directory.TransactOpts, amount, stakee)
}

// LockStake is a paid mutator transaction binding the contract method 0x33f4909f.
//
// Solidity: function lockStake(uint256 amount, address stakee) returns()
func (_Directory *DirectoryTransactorSession) LockStake(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.LockStake(&_Directory.TransactOpts, amount, stakee)
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

// Unstake is a paid mutator transaction binding the contract method 0xf2888dbb.
//
// Solidity: function unstake(address stakee) returns()
func (_Directory *DirectoryTransactor) Unstake(opts *bind.TransactOpts, stakee common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "unstake", stakee)
}

// Unstake is a paid mutator transaction binding the contract method 0xf2888dbb.
//
// Solidity: function unstake(address stakee) returns()
func (_Directory *DirectorySession) Unstake(stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.Unstake(&_Directory.TransactOpts, stakee)
}

// Unstake is a paid mutator transaction binding the contract method 0xf2888dbb.
//
// Solidity: function unstake(address stakee) returns()
func (_Directory *DirectoryTransactorSession) Unstake(stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.Unstake(&_Directory.TransactOpts, stakee)
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
