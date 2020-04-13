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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DirectoryABI is the input ABI used to generate the binding from.
const DirectoryABI = "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unlockDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rightAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"parent\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"left\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"right\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"unlockings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newUnlockDuration\",\"type\":\"uint256\"}],\"name\":\"setUnlockDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"addStakeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unlockStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"unlockStakeFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"lockStakeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"unstakeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rand\",\"type\":\"uint256\"}],\"name\":\"scan\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DirectoryBin is the compiled bytecode used for deploying new contracts.
var DirectoryBin = "0x608060405234801561001057600080fd5b506040516111f93803806111f983398101604081905261002f916100b9565b60006100426001600160e01b036100b516565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600180546001600160a01b0319166001600160a01b0393909316929092179091556002556100f1565b3390565b600080604083850312156100cb578182fd5b82516001600160a01b03811681146100e1578283fd5b6020939093015192949293505050565b6110f9806101006000396000f3fe608060405234801561001057600080fd5b50600436106101515760003560e01c80638fee6407116100cd578063e16d14a011610081578063ecd8c06a11610066578063ecd8c06a14610295578063f2fde38b146102a8578063f50ddb67146102bb57610151565b8063e16d14a01461026f578063eb4f16b51461028257610151565b80639341a536116100b25780639341a53614610228578063bc29855314610249578063dd9007691461025c57610151565b80638fee6407146101ef57806391f149c91461021557610151565b80636bbfea21116101245780637bc74225116101095780637bc74225146101ca5780638a1fcd60146101d25780638da5cb5b146101da57610151565b80636bbfea21146101af578063715018a6146101c257610151565b8063092a5549146101565780632cbb26191461017f5780632def6620146101945780636b5537e21461019c575b600080fd5b610169610164366004610db2565b6102ce565b6040516101769190610e60565b60405180910390f35b61019261018d366004610d24565b61053b565b005b61019261063e565b6101696101aa366004610d24565b610649565b6101926101bd366004610db2565b61065b565b61019261081d565b6101696108a9565b6101696108e1565b6101e26108e7565b6040516101769190610e0f565b6102026101fd366004610d9a565b6108f6565b604051610176979695949392919061108a565b610192610223366004610db2565b61093e565b61023b610236366004610d9a565b610970565b60405161017692919061107c565b610169610257366004610d46565b610989565b61019261026a366004610d9a565b6109bc565b6101e261027d366004610d9a565b6109f6565b610192610290366004610d9a565b610a89565b6101696102a3366004610d9a565b610a96565b6101926102b6366004610d24565b610aa2565b6101926102c9366004610d9a565b610ae0565b6000806102db3384610989565b60008181526003602052604090208054919250906103145760405162461bcd60e51b815260040161030b9061100e565b60405180910390fd5b80548511156103355760405162461bcd60e51b815260040161030b90610efd565b610343828287600003610aea565b80546104ff576000816002015482600101541161036457816006015461036a565b81600501545b60048301546000908152600360205260409020909150816103965761039183856000610b26565b6104ab565b60008281526003602052604090205b600081600201548260010154116103c05781600601546103c6565b81600501545b9050806103d357506103eb565b600081815260036020526040902090935090506103a5565b60048101546103fb838786610b26565b6004808601549083015580861461047757610417858584610b4a565b610422858584610b87565b600485018190556000818152600360205260409020610442908288610b26565b610459868684600001546000038560040154610bc4565b6000818152600360205260408120610472918890610b26565b610499565b838560050154141561048e57610472858584610b87565b610499858584610b4a565b60048201546104a85760068490555b50505b5050600082815260036020819052604082208281556001810183905560028101839055908101805473ffffffffffffffffffffffffffffffffffffffff191690556004810182905560058101829055600601555b60008281526005602052604090206002546001820154439091019081111561052957600182018190555b81548701909155925050505b92915050565b60006105473383610989565b6000818152600560205260409020600181015491925090431161057c5760405162461bcd60e51b815260040161030b90611045565b805461059a5760405162461bcd60e51b815260040161030b90610f34565b60015481546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b039092169163a9059cbb916105e691339190600401610e47565b602060405180830381600087803b15801561060057600080fd5b505af1158015610614573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106389190610d7a565b50505050565b6106473361053b565b565b60046020526000908152604090205481565b6001600160a01b0381166106815760405162461bcd60e51b815260040161030b90610ec6565b8161069e5760405162461bcd60e51b815260040161030b90610fd7565b60006106aa3383610989565b600081815260036020526040902080549192509061075d5760065460008181526003602052604090205b811561072057600081600201548260010154106106f55781600601546106fb565b81600501545b9050806107085750610720565b600081815260036020526040902090925090506106d4565b61072c81600086610b26565b50600482015560038101805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0385161790555b600481015461076c5760068290555b610777828286610aea565b6001546040517f23b872dd0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116906323b872dd906107c490339030908990600401610e23565b602060405180830381600087803b1580156107de57600080fd5b505af11580156107f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108169190610d7a565b5050505050565b610825610c1f565b6000546001600160a01b039081169116146108525760405162461bcd60e51b815260040161030b90610f6b565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36000805473ffffffffffffffffffffffffffffffffffffffff19169055565b6006546000906108bb575060006108de565b506006546000908152600360205260409020600281015460018201549154909101015b90565b60025481565b6000546001600160a01b031690565b600360208190526000918252604090912080546001820154600283015493830154600484015460058501546006909501549395929492936001600160a01b0390921692909187565b600061094a3383610989565b60008181526003602052604090209091506109658285610c23565b610638828286610aea565b6005602052600090815260409020805460019091015482565b6000828260405160200161099e929190610dd5565b60405160208183030381529060405280519060200120905092915050565b6109c4610c1f565b6000546001600160a01b039081169116146109f15760405162461bcd60e51b815260040161030b90610f6b565b600255565b600654600090610a0857506000610a84565b600060fd83610a156108a9565b0281610a1d57fe5b60065491900491505b60008181526003602052604090206001810154831015610a4b57600501549050610a26565b600181015481549303928311610a7257600301546001600160a01b03169250610a84915050565b80546006909101549203919050610a26565b919050565b610a93813361065b565b50565b600061053582336102ce565b610aaa610c1f565b6000546001600160a01b03908116911614610ad75760405162461bcd60e51b815260040161030b90610f6b565b610a9381610c7f565b610a93813361093e565b81548101825560038201546001600160a01b03166000908152600460205260408120805483019055610b2190849084908490610bc4565b505050565b8183600501541415610b3e5760058301819055610b21565b60068301819055505050565b6005830154610b5857610b21565b600580840180546000908152600360205260409020600401849055549082015560018084015490820155505050565b6006830154610b9557610b21565b600680840180546000908152600360205260409020600401849055549082015560028084015490820155505050565b600483015481811415610bd75750610638565b60008181526003602052604090206005810154861415610c005760018101805485019055610c0b565b600281018054850190555b610c1782828686610bc4565b505050505050565b3390565b60008281526005602052604090208054821415610c5457600083815260056020526040812081815560010155610b21565b80548210610c745760405162461bcd60e51b815260040161030b90610fa0565b805491909103905550565b6001600160a01b038116610ca55760405162461bcd60e51b815260040161030b90610e69565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b80356001600160a01b038116811461053557600080fd5b600060208284031215610d35578081fd5b610d3f8383610d0d565b9392505050565b60008060408385031215610d58578081fd5b610d628484610d0d565b9150610d718460208501610d0d565b90509250929050565b600060208284031215610d8b578081fd5b81518015158114610d3f578182fd5b600060208284031215610dab578081fd5b5035919050565b60008060408385031215610dc4578182fd5b82359150610d718460208501610d0d565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606093841b811682529190921b16601482015260280190565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b90815260200190565b60208082526026908201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408201527f6464726573730000000000000000000000000000000000000000000000000000606082015260800190565b6020808252600f908201527f41646472657373206973206e756c6c0000000000000000000000000000000000604082015260600190565b6020808252601e908201527f43616e6e6f7420756e6c6f636b206d6f7265207468616e207374616b65640000604082015260600190565b60208082526013908201527f4e6f20616d6f756e7420746f20756e6c6f636b00000000000000000000000000604082015260600190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252601e908201527f556e6c6f636b2068617320696e73756666696369656e7420616d6f756e740000604082015260600190565b60208082526014908201527f43616e6e6f74207374616b65206e6f7468696e67000000000000000000000000604082015260600190565b60208082526012908201527f4e6f7468696e6720746f20756e7374616b650000000000000000000000000000604082015260600190565b60208082526016908201527f5374616b65206e6f742079657420756e6c6f636b656400000000000000000000604082015260600190565b918252602082015260400190565b968752602087019590955260408601939093526001600160a01b03919091166060850152608084015260a083015260c082015260e0019056fea264697066735822122035bb674fb74e185a2d1bed70eebc3ac48df910d33c9916a2f73e821606c23e8c64736f6c63430006040033"

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
func (_Directory *DirectoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_Directory *DirectoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
// Solidity: function getKey(address staker, address stakee) constant returns(bytes32)
func (_Directory *DirectoryCaller) GetKey(opts *bind.CallOpts, staker common.Address, stakee common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "getKey", staker, stakee)
	return *ret0, err
}

// GetKey is a free data retrieval call binding the contract method 0xbc298553.
//
// Solidity: function getKey(address staker, address stakee) constant returns(bytes32)
func (_Directory *DirectorySession) GetKey(staker common.Address, stakee common.Address) ([32]byte, error) {
	return _Directory.Contract.GetKey(&_Directory.CallOpts, staker, stakee)
}

// GetKey is a free data retrieval call binding the contract method 0xbc298553.
//
// Solidity: function getKey(address staker, address stakee) constant returns(bytes32)
func (_Directory *DirectoryCallerSession) GetKey(staker common.Address, stakee common.Address) ([32]byte, error) {
	return _Directory.Contract.GetKey(&_Directory.CallOpts, staker, stakee)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() constant returns(uint256)
func (_Directory *DirectoryCaller) GetTotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "getTotalStake")
	return *ret0, err
}

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() constant returns(uint256)
func (_Directory *DirectorySession) GetTotalStake() (*big.Int, error) {
	return _Directory.Contract.GetTotalStake(&_Directory.CallOpts)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() constant returns(uint256)
func (_Directory *DirectoryCallerSession) GetTotalStake() (*big.Int, error) {
	return _Directory.Contract.GetTotalStake(&_Directory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Directory *DirectoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Directory *DirectorySession) Owner() (common.Address, error) {
	return _Directory.Contract.Owner(&_Directory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Directory *DirectoryCallerSession) Owner() (common.Address, error) {
	return _Directory.Contract.Owner(&_Directory.CallOpts)
}

// Scan is a free data retrieval call binding the contract method 0xe16d14a0.
//
// Solidity: function scan(uint256 rand) constant returns(address)
func (_Directory *DirectoryCaller) Scan(opts *bind.CallOpts, rand *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "scan", rand)
	return *ret0, err
}

// Scan is a free data retrieval call binding the contract method 0xe16d14a0.
//
// Solidity: function scan(uint256 rand) constant returns(address)
func (_Directory *DirectorySession) Scan(rand *big.Int) (common.Address, error) {
	return _Directory.Contract.Scan(&_Directory.CallOpts, rand)
}

// Scan is a free data retrieval call binding the contract method 0xe16d14a0.
//
// Solidity: function scan(uint256 rand) constant returns(address)
func (_Directory *DirectoryCallerSession) Scan(rand *big.Int) (common.Address, error) {
	return _Directory.Contract.Scan(&_Directory.CallOpts, rand)
}

// Stakees is a free data retrieval call binding the contract method 0x6b5537e2.
//
// Solidity: function stakees(address ) constant returns(uint256)
func (_Directory *DirectoryCaller) Stakees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "stakees", arg0)
	return *ret0, err
}

// Stakees is a free data retrieval call binding the contract method 0x6b5537e2.
//
// Solidity: function stakees(address ) constant returns(uint256)
func (_Directory *DirectorySession) Stakees(arg0 common.Address) (*big.Int, error) {
	return _Directory.Contract.Stakees(&_Directory.CallOpts, arg0)
}

// Stakees is a free data retrieval call binding the contract method 0x6b5537e2.
//
// Solidity: function stakees(address ) constant returns(uint256)
func (_Directory *DirectoryCallerSession) Stakees(arg0 common.Address) (*big.Int, error) {
	return _Directory.Contract.Stakees(&_Directory.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x8fee6407.
//
// Solidity: function stakes(bytes32 ) constant returns(uint256 amount, uint256 leftAmount, uint256 rightAmount, address stakee, bytes32 parent, bytes32 left, bytes32 right)
func (_Directory *DirectoryCaller) Stakes(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Amount      *big.Int
	LeftAmount  *big.Int
	RightAmount *big.Int
	Stakee      common.Address
	Parent      [32]byte
	Left        [32]byte
	Right       [32]byte
}, error) {
	ret := new(struct {
		Amount      *big.Int
		LeftAmount  *big.Int
		RightAmount *big.Int
		Stakee      common.Address
		Parent      [32]byte
		Left        [32]byte
		Right       [32]byte
	})
	out := ret
	err := _Directory.contract.Call(opts, out, "stakes", arg0)
	return *ret, err
}

// Stakes is a free data retrieval call binding the contract method 0x8fee6407.
//
// Solidity: function stakes(bytes32 ) constant returns(uint256 amount, uint256 leftAmount, uint256 rightAmount, address stakee, bytes32 parent, bytes32 left, bytes32 right)
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
// Solidity: function stakes(bytes32 ) constant returns(uint256 amount, uint256 leftAmount, uint256 rightAmount, address stakee, bytes32 parent, bytes32 left, bytes32 right)
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
// Solidity: function unlockDuration() constant returns(uint256)
func (_Directory *DirectoryCaller) UnlockDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Directory.contract.Call(opts, out, "unlockDuration")
	return *ret0, err
}

// UnlockDuration is a free data retrieval call binding the contract method 0x8a1fcd60.
//
// Solidity: function unlockDuration() constant returns(uint256)
func (_Directory *DirectorySession) UnlockDuration() (*big.Int, error) {
	return _Directory.Contract.UnlockDuration(&_Directory.CallOpts)
}

// UnlockDuration is a free data retrieval call binding the contract method 0x8a1fcd60.
//
// Solidity: function unlockDuration() constant returns(uint256)
func (_Directory *DirectoryCallerSession) UnlockDuration() (*big.Int, error) {
	return _Directory.Contract.UnlockDuration(&_Directory.CallOpts)
}

// Unlockings is a free data retrieval call binding the contract method 0x9341a536.
//
// Solidity: function unlockings(bytes32 ) constant returns(uint256 amount, uint256 unlockAt)
func (_Directory *DirectoryCaller) Unlockings(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Amount   *big.Int
	UnlockAt *big.Int
}, error) {
	ret := new(struct {
		Amount   *big.Int
		UnlockAt *big.Int
	})
	out := ret
	err := _Directory.contract.Call(opts, out, "unlockings", arg0)
	return *ret, err
}

// Unlockings is a free data retrieval call binding the contract method 0x9341a536.
//
// Solidity: function unlockings(bytes32 ) constant returns(uint256 amount, uint256 unlockAt)
func (_Directory *DirectorySession) Unlockings(arg0 [32]byte) (struct {
	Amount   *big.Int
	UnlockAt *big.Int
}, error) {
	return _Directory.Contract.Unlockings(&_Directory.CallOpts, arg0)
}

// Unlockings is a free data retrieval call binding the contract method 0x9341a536.
//
// Solidity: function unlockings(bytes32 ) constant returns(uint256 amount, uint256 unlockAt)
func (_Directory *DirectoryCallerSession) Unlockings(arg0 [32]byte) (struct {
	Amount   *big.Int
	UnlockAt *big.Int
}, error) {
	return _Directory.Contract.Unlockings(&_Directory.CallOpts, arg0)
}

// AddStake is a paid mutator transaction binding the contract method 0xeb4f16b5.
//
// Solidity: function addStake(uint256 amount) returns()
func (_Directory *DirectoryTransactor) AddStake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "addStake", amount)
}

// AddStake is a paid mutator transaction binding the contract method 0xeb4f16b5.
//
// Solidity: function addStake(uint256 amount) returns()
func (_Directory *DirectorySession) AddStake(amount *big.Int) (*types.Transaction, error) {
	return _Directory.Contract.AddStake(&_Directory.TransactOpts, amount)
}

// AddStake is a paid mutator transaction binding the contract method 0xeb4f16b5.
//
// Solidity: function addStake(uint256 amount) returns()
func (_Directory *DirectoryTransactorSession) AddStake(amount *big.Int) (*types.Transaction, error) {
	return _Directory.Contract.AddStake(&_Directory.TransactOpts, amount)
}

// AddStakeFor is a paid mutator transaction binding the contract method 0x6bbfea21.
//
// Solidity: function addStakeFor(uint256 amount, address stakee) returns()
func (_Directory *DirectoryTransactor) AddStakeFor(opts *bind.TransactOpts, amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "addStakeFor", amount, stakee)
}

// AddStakeFor is a paid mutator transaction binding the contract method 0x6bbfea21.
//
// Solidity: function addStakeFor(uint256 amount, address stakee) returns()
func (_Directory *DirectorySession) AddStakeFor(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.AddStakeFor(&_Directory.TransactOpts, amount, stakee)
}

// AddStakeFor is a paid mutator transaction binding the contract method 0x6bbfea21.
//
// Solidity: function addStakeFor(uint256 amount, address stakee) returns()
func (_Directory *DirectoryTransactorSession) AddStakeFor(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.AddStakeFor(&_Directory.TransactOpts, amount, stakee)
}

// LockStake is a paid mutator transaction binding the contract method 0xf50ddb67.
//
// Solidity: function lockStake(uint256 amount) returns()
func (_Directory *DirectoryTransactor) LockStake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "lockStake", amount)
}

// LockStake is a paid mutator transaction binding the contract method 0xf50ddb67.
//
// Solidity: function lockStake(uint256 amount) returns()
func (_Directory *DirectorySession) LockStake(amount *big.Int) (*types.Transaction, error) {
	return _Directory.Contract.LockStake(&_Directory.TransactOpts, amount)
}

// LockStake is a paid mutator transaction binding the contract method 0xf50ddb67.
//
// Solidity: function lockStake(uint256 amount) returns()
func (_Directory *DirectoryTransactorSession) LockStake(amount *big.Int) (*types.Transaction, error) {
	return _Directory.Contract.LockStake(&_Directory.TransactOpts, amount)
}

// LockStakeFor is a paid mutator transaction binding the contract method 0x91f149c9.
//
// Solidity: function lockStakeFor(uint256 amount, address stakee) returns()
func (_Directory *DirectoryTransactor) LockStakeFor(opts *bind.TransactOpts, amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "lockStakeFor", amount, stakee)
}

// LockStakeFor is a paid mutator transaction binding the contract method 0x91f149c9.
//
// Solidity: function lockStakeFor(uint256 amount, address stakee) returns()
func (_Directory *DirectorySession) LockStakeFor(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.LockStakeFor(&_Directory.TransactOpts, amount, stakee)
}

// LockStakeFor is a paid mutator transaction binding the contract method 0x91f149c9.
//
// Solidity: function lockStakeFor(uint256 amount, address stakee) returns()
func (_Directory *DirectoryTransactorSession) LockStakeFor(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.LockStakeFor(&_Directory.TransactOpts, amount, stakee)
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

// UnlockStake is a paid mutator transaction binding the contract method 0xecd8c06a.
//
// Solidity: function unlockStake(uint256 amount) returns(uint256)
func (_Directory *DirectoryTransactor) UnlockStake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "unlockStake", amount)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xecd8c06a.
//
// Solidity: function unlockStake(uint256 amount) returns(uint256)
func (_Directory *DirectorySession) UnlockStake(amount *big.Int) (*types.Transaction, error) {
	return _Directory.Contract.UnlockStake(&_Directory.TransactOpts, amount)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xecd8c06a.
//
// Solidity: function unlockStake(uint256 amount) returns(uint256)
func (_Directory *DirectoryTransactorSession) UnlockStake(amount *big.Int) (*types.Transaction, error) {
	return _Directory.Contract.UnlockStake(&_Directory.TransactOpts, amount)
}

// UnlockStakeFor is a paid mutator transaction binding the contract method 0x092a5549.
//
// Solidity: function unlockStakeFor(uint256 amount, address stakee) returns(uint256)
func (_Directory *DirectoryTransactor) UnlockStakeFor(opts *bind.TransactOpts, amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "unlockStakeFor", amount, stakee)
}

// UnlockStakeFor is a paid mutator transaction binding the contract method 0x092a5549.
//
// Solidity: function unlockStakeFor(uint256 amount, address stakee) returns(uint256)
func (_Directory *DirectorySession) UnlockStakeFor(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.UnlockStakeFor(&_Directory.TransactOpts, amount, stakee)
}

// UnlockStakeFor is a paid mutator transaction binding the contract method 0x092a5549.
//
// Solidity: function unlockStakeFor(uint256 amount, address stakee) returns(uint256)
func (_Directory *DirectoryTransactorSession) UnlockStakeFor(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.UnlockStakeFor(&_Directory.TransactOpts, amount, stakee)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Directory *DirectoryTransactor) Unstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "unstake")
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Directory *DirectorySession) Unstake() (*types.Transaction, error) {
	return _Directory.Contract.Unstake(&_Directory.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Directory *DirectoryTransactorSession) Unstake() (*types.Transaction, error) {
	return _Directory.Contract.Unstake(&_Directory.TransactOpts)
}

// UnstakeFor is a paid mutator transaction binding the contract method 0x2cbb2619.
//
// Solidity: function unstakeFor(address stakee) returns()
func (_Directory *DirectoryTransactor) UnstakeFor(opts *bind.TransactOpts, stakee common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "unstakeFor", stakee)
}

// UnstakeFor is a paid mutator transaction binding the contract method 0x2cbb2619.
//
// Solidity: function unstakeFor(address stakee) returns()
func (_Directory *DirectorySession) UnstakeFor(stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.UnstakeFor(&_Directory.TransactOpts, stakee)
}

// UnstakeFor is a paid mutator transaction binding the contract method 0x2cbb2619.
//
// Solidity: function unstakeFor(address stakee) returns()
func (_Directory *DirectoryTransactorSession) UnstakeFor(stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.UnstakeFor(&_Directory.TransactOpts, stakee)
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
	return event, nil
}
