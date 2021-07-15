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

// PriceVotingVote is an auto generated low-level Go binding around an user-defined struct.
type PriceVotingVote struct {
	Voter common.Address
	Price *big.Int
}

// PriceVotingABI is the input ABI used to generate the binding from.
const PriceVotingABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractStakingManager\",\"name\":\"stakingManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sortVotes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structPriceVoting.Vote[]\",\"name\":\"sortedVotes\",\"type\":\"tuple[]\"}],\"name\":\"validateSortedVotes\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PriceVotingBin is the compiled bytecode used for deploying new contracts.
var PriceVotingBin = "0x608060405234801561001057600080fd5b506117dc806100206000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80638da5cb5b11610081578063da58c7d91161005b578063da58c7d91461018b578063efb50efb1461019e578063f2fde38b146101a657600080fd5b80638da5cb5b14610125578063c4d66de81461014a578063d8bff5a51461015d57600080fd5b80633ccfd60b116100b25780633ccfd60b14610102578063715018a61461010a5780637c5d14591461011257600080fd5b80630121b93f146100ce5780630dc96015146100e3575b600080fd5b6100e16100dc366004611560565b6101b9565b005b6100eb6102aa565b6040516100f9929190611578565b60405180910390f35b6100e161047f565b6100e16105dc565b6100e161012036600461148d565b61068d565b6033546001600160a01b03165b6040516001600160a01b0390911681526020016100f9565b6100e161015836600461146a565b610a93565b61017d61016b36600461146a565b60666020526000908152604090205481565b6040519081526020016100f9565b610132610199366004611560565b610b77565b6100eb610ba1565b6100e16101b436600461146a565b610e77565b600081116102345760405162461bcd60e51b815260206004820152602360248201527f566f74696e67207072696365206d75737420626520677265617465722074686160448201527f6e2030000000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b3360009081526066602052604090205461029857606780546001810182556000919091527f9787eeb91fe3101235e4a76063c7023ecb40f923f97916639c598592fa30d6ae01805473ffffffffffffffffffffffffffffffffffffffff1916331790555b33600090815260666020526040902055565b606080600060678054905067ffffffffffffffff8111156102db57634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610304578160200160208202803683370190505b5060675490915060009067ffffffffffffffff81111561033457634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561035d578160200160208202803683370190505b50905060005b606754811015610475576067818154811061038e57634e487b7160e01b600052603260045260246000fd5b9060005260206000200160009054906101000a90046001600160a01b03168382815181106103cc57634e487b7160e01b600052603260045260246000fd5b60200260200101906001600160a01b031690816001600160a01b031681525050606660006067838154811061041157634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101546001600160a01b03168352820192909252604001902054825183908390811061045857634e487b7160e01b600052603260045260246000fd5b60209081029190910101528061046d81611755565b915050610363565b5090939092509050565b3360009081526066602052604081208190555b6067548110156105d957336001600160a01b0316606782815481106104c757634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b031614156105c757606780546104f290600190611705565b8154811061051057634e487b7160e01b600052603260045260246000fd5b600091825260209091200154606780546001600160a01b03909216918390811061054a57634e487b7160e01b600052603260045260246000fd5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550606780548061059757634e487b7160e01b600052603160045260246000fd5b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff191690550190555b806105d181611755565b915050610492565b50565b6033546001600160a01b031633146106365760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161022b565b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36033805473ffffffffffffffffffffffffffffffffffffffff19169055565b600080825167ffffffffffffffff8111156106b857634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156106e1578160200160208202803683370190505b50905060005b8351811015610a15578284828151811061071157634e487b7160e01b600052603260045260246000fd5b602002602001015160200151101561076b5760405162461bcd60e51b815260206004820152601e60248201527f476976656e20766f7465206172726179206973206e6f7420736f727465640000604482015260640161022b565b83818151811061078b57634e487b7160e01b600052603260045260246000fd5b60200260200101516020015192506000606660008684815181106107bf57634e487b7160e01b600052603260045260246000fd5b6020026020010151600001516001600160a01b03166001600160a01b03168152602001908152602001600020541161085f5760405162461bcd60e51b815260206004820152602960248201527f466f756e6420696e76616c696420766f74657220696e20736f7274656420766f60448201527f7465722061727261790000000000000000000000000000000000000000000000606482015260840161022b565b60005b825181101561099e5784828151811061088b57634e487b7160e01b600052603260045260246000fd5b6020026020010151600001516001600160a01b03168382815181106108c057634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b031614156109455760405162461bcd60e51b815260206004820152602560248201527f466f756e64206475706c696361746520696e20736f7274656420766f7465722060448201527f6172726179000000000000000000000000000000000000000000000000000000606482015260840161022b565b60006001600160a01b031683828151811061097057634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b0316141561098c5761099e565b8061099681611755565b915050610862565b508381815181106109bf57634e487b7160e01b600052603260045260246000fd5b6020026020010151600001518282815181106109eb57634e487b7160e01b600052603260045260246000fd5b6001600160a01b039092166020928302919091019091015280610a0d81611755565b9150506106e7565b50606754835114610a8e5760405162461bcd60e51b815260206004820152603160248201527f4e6f7420616c6c20766f7465727320776572652070726573656e7420696e207360448201527f6f7274656420766f746572206172726179000000000000000000000000000000606482015260840161022b565b505050565b600054610100900460ff1680610aac575060005460ff16155b610b0f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161022b565b600054610100900460ff16158015610b31576000805461ffff19166101011790555b610b39610fb6565b6065805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0384161790558015610b73576000805461ff00191690555b5050565b60678181548110610b8757600080fd5b6000918252602090912001546001600160a01b0316905081565b606080600060678054905067ffffffffffffffff811115610bd257634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610c1757816020015b6040805180820190915260008082526020820152815260200190600190039081610bf05790505b50905060005b606754811015610cc557600060678281548110610c4a57634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101546040805180820182526001600160a01b039092168083528085526066845293205491810191909152845191925090849084908110610ca657634e487b7160e01b600052603260045260246000fd5b6020026020010181905250508080610cbd90611755565b915050610c1d565b50610cdf81600060018451610cda9190611705565b611078565b6000815167ffffffffffffffff811115610d0957634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610d32578160200160208202803683370190505b5090506000825167ffffffffffffffff811115610d5f57634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610d88578160200160208202803683370190505b50905060005b8351811015610e6c57838181518110610db757634e487b7160e01b600052603260045260246000fd5b602002602001015160000151838281518110610de357634e487b7160e01b600052603260045260246000fd5b60200260200101906001600160a01b031690816001600160a01b031681525050838181518110610e2357634e487b7160e01b600052603260045260246000fd5b602002602001015160200151828281518110610e4f57634e487b7160e01b600052603260045260246000fd5b602090810291909101015280610e6481611755565b915050610d8e565b509094909350915050565b6033546001600160a01b03163314610ed15760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161022b565b6001600160a01b038116610f4d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161022b565b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a36033805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b600054610100900460ff1680610fcf575060005460ff16155b6110325760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161022b565b600054610100900460ff16158015611054576000805461ffff19166101011790555b61105c6110ba565b61106461116b565b80156105d9576000805461ff001916905550565b80821215610a8e57600061108d84848461126d565b905061109f8484610cda6001856116ad565b6110b4846110ae836001611655565b84611078565b50505050565b600054610100900460ff16806110d3575060005460ff16155b6111365760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161022b565b600054610100900460ff16158015611064576000805461ffff191661010117905580156105d9576000805461ff001916905550565b600054610100900460ff1680611184575060005460ff16155b6111e75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161022b565b600054610100900460ff16158015611209576000805461ffff19166101011790555b6033805473ffffffffffffffffffffffffffffffffffffffff19163390811790915560405181906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35080156105d9576000805461ff001916905550565b60008084838151811061129057634e487b7160e01b600052603260045260246000fd5b6020026020010151905060006001856112a991906116ad565b9050845b6112b86001866116ad565b81136113c15782602001518782815181106112e357634e487b7160e01b600052603260045260246000fd5b60200260200101516020015110156113af57816112ff8161171c565b925050600087828151811061132457634e487b7160e01b600052603260045260246000fd5b6020026020010151905087838151811061134e57634e487b7160e01b600052603260045260246000fd5b602002602001015188838151811061137657634e487b7160e01b600052603260045260246000fd5b6020026020010181905250808884815181106113a257634e487b7160e01b600052603260045260246000fd5b6020026020010181905250505b806113b98161171c565b9150506112ad565b50856113ce826001611655565b815181106113ec57634e487b7160e01b600052603260045260246000fd5b602002602001015186858151811061141457634e487b7160e01b600052603260045260246000fd5b6020908102919091010152818661142c836001611655565b8151811061144a57634e487b7160e01b600052603260045260246000fd5b6020908102919091010152611460816001611655565b9695505050505050565b60006020828403121561147b578081fd5b813561148681611791565b9392505050565b6000602080838503121561149f578182fd5b823567ffffffffffffffff808211156114b6578384fd5b818501915085601f8301126114c9578384fd5b8135818111156114db576114db61177b565b6114e9848260051b01611624565b8181528481019250838501600683901b85018601891015611508578687fd5b8694505b8285101561155457604080828b031215611524578788fd5b61152c6115fb565b823561153781611791565b81528288013588820152855260019590950194938601930161150c565b50979650505050505050565b600060208284031215611571578081fd5b5035919050565b604080825283519082018190526000906020906060840190828701845b828110156115ba5781516001600160a01b031684529284019290840190600101611595565b50505083810382850152845180825285830191830190845b818110156115ee578351835292840192918401916001016115d2565b5090979650505050505050565b6040805190810167ffffffffffffffff8111828210171561161e5761161e61177b565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561164d5761164d61177b565b604052919050565b6000808212827f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0384138115161561168f5761168f611765565b82600160ff1b0384128116156116a7576116a7611765565b50500190565b600080831283600160ff1b018312811516156116cb576116cb611765565b837f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0183138116156116ff576116ff611765565b50500390565b60008282101561171757611717611765565b500390565b60007f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561174e5761174e611765565b5060010190565b600060001982141561174e5761174e5b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146105d957600080fdfea26469706673582212202fb752083a7cd0a50ccb0ad9cf8d0860900e99b086528a01cdd6938d7a59c42864736f6c63430008040033"

// DeployPriceVoting deploys a new Ethereum contract, binding an instance of PriceVoting to it.
func DeployPriceVoting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PriceVoting, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceVotingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PriceVotingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PriceVoting{PriceVotingCaller: PriceVotingCaller{contract: contract}, PriceVotingTransactor: PriceVotingTransactor{contract: contract}, PriceVotingFilterer: PriceVotingFilterer{contract: contract}}, nil
}

// PriceVoting is an auto generated Go binding around an Ethereum contract.
type PriceVoting struct {
	PriceVotingCaller     // Read-only binding to the contract
	PriceVotingTransactor // Write-only binding to the contract
	PriceVotingFilterer   // Log filterer for contract events
}

// PriceVotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceVotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceVotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceVotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceVotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceVotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceVotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceVotingSession struct {
	Contract     *PriceVoting      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceVotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceVotingCallerSession struct {
	Contract *PriceVotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PriceVotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceVotingTransactorSession struct {
	Contract     *PriceVotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PriceVotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceVotingRaw struct {
	Contract *PriceVoting // Generic contract binding to access the raw methods on
}

// PriceVotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceVotingCallerRaw struct {
	Contract *PriceVotingCaller // Generic read-only contract binding to access the raw methods on
}

// PriceVotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceVotingTransactorRaw struct {
	Contract *PriceVotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceVoting creates a new instance of PriceVoting, bound to a specific deployed contract.
func NewPriceVoting(address common.Address, backend bind.ContractBackend) (*PriceVoting, error) {
	contract, err := bindPriceVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceVoting{PriceVotingCaller: PriceVotingCaller{contract: contract}, PriceVotingTransactor: PriceVotingTransactor{contract: contract}, PriceVotingFilterer: PriceVotingFilterer{contract: contract}}, nil
}

// NewPriceVotingCaller creates a new read-only instance of PriceVoting, bound to a specific deployed contract.
func NewPriceVotingCaller(address common.Address, caller bind.ContractCaller) (*PriceVotingCaller, error) {
	contract, err := bindPriceVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceVotingCaller{contract: contract}, nil
}

// NewPriceVotingTransactor creates a new write-only instance of PriceVoting, bound to a specific deployed contract.
func NewPriceVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceVotingTransactor, error) {
	contract, err := bindPriceVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceVotingTransactor{contract: contract}, nil
}

// NewPriceVotingFilterer creates a new log filterer instance of PriceVoting, bound to a specific deployed contract.
func NewPriceVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceVotingFilterer, error) {
	contract, err := bindPriceVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceVotingFilterer{contract: contract}, nil
}

// bindPriceVoting binds a generic wrapper to an already deployed contract.
func bindPriceVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceVotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceVoting *PriceVotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceVoting.Contract.PriceVotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceVoting *PriceVotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceVoting.Contract.PriceVotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceVoting *PriceVotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceVoting.Contract.PriceVotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceVoting *PriceVotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceVoting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceVoting *PriceVotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceVoting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceVoting *PriceVotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceVoting.Contract.contract.Transact(opts, method, params...)
}

// GetVotes is a free data retrieval call binding the contract method 0x0dc96015.
//
// Solidity: function getVotes() view returns(address[], uint256[])
func (_PriceVoting *PriceVotingCaller) GetVotes(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _PriceVoting.contract.Call(opts, &out, "getVotes")

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetVotes is a free data retrieval call binding the contract method 0x0dc96015.
//
// Solidity: function getVotes() view returns(address[], uint256[])
func (_PriceVoting *PriceVotingSession) GetVotes() ([]common.Address, []*big.Int, error) {
	return _PriceVoting.Contract.GetVotes(&_PriceVoting.CallOpts)
}

// GetVotes is a free data retrieval call binding the contract method 0x0dc96015.
//
// Solidity: function getVotes() view returns(address[], uint256[])
func (_PriceVoting *PriceVotingCallerSession) GetVotes() ([]common.Address, []*big.Int, error) {
	return _PriceVoting.Contract.GetVotes(&_PriceVoting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceVoting *PriceVotingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceVoting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceVoting *PriceVotingSession) Owner() (common.Address, error) {
	return _PriceVoting.Contract.Owner(&_PriceVoting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceVoting *PriceVotingCallerSession) Owner() (common.Address, error) {
	return _PriceVoting.Contract.Owner(&_PriceVoting.CallOpts)
}

// SortVotes is a free data retrieval call binding the contract method 0xefb50efb.
//
// Solidity: function sortVotes() view returns(address[], uint256[])
func (_PriceVoting *PriceVotingCaller) SortVotes(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _PriceVoting.contract.Call(opts, &out, "sortVotes")

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// SortVotes is a free data retrieval call binding the contract method 0xefb50efb.
//
// Solidity: function sortVotes() view returns(address[], uint256[])
func (_PriceVoting *PriceVotingSession) SortVotes() ([]common.Address, []*big.Int, error) {
	return _PriceVoting.Contract.SortVotes(&_PriceVoting.CallOpts)
}

// SortVotes is a free data retrieval call binding the contract method 0xefb50efb.
//
// Solidity: function sortVotes() view returns(address[], uint256[])
func (_PriceVoting *PriceVotingCallerSession) SortVotes() ([]common.Address, []*big.Int, error) {
	return _PriceVoting.Contract.SortVotes(&_PriceVoting.CallOpts)
}

// ValidateSortedVotes is a free data retrieval call binding the contract method 0x7c5d1459.
//
// Solidity: function validateSortedVotes((address,uint256)[] sortedVotes) view returns()
func (_PriceVoting *PriceVotingCaller) ValidateSortedVotes(opts *bind.CallOpts, sortedVotes []PriceVotingVote) error {
	var out []interface{}
	err := _PriceVoting.contract.Call(opts, &out, "validateSortedVotes", sortedVotes)

	if err != nil {
		return err
	}

	return err

}

// ValidateSortedVotes is a free data retrieval call binding the contract method 0x7c5d1459.
//
// Solidity: function validateSortedVotes((address,uint256)[] sortedVotes) view returns()
func (_PriceVoting *PriceVotingSession) ValidateSortedVotes(sortedVotes []PriceVotingVote) error {
	return _PriceVoting.Contract.ValidateSortedVotes(&_PriceVoting.CallOpts, sortedVotes)
}

// ValidateSortedVotes is a free data retrieval call binding the contract method 0x7c5d1459.
//
// Solidity: function validateSortedVotes((address,uint256)[] sortedVotes) view returns()
func (_PriceVoting *PriceVotingCallerSession) ValidateSortedVotes(sortedVotes []PriceVotingVote) error {
	return _PriceVoting.Contract.ValidateSortedVotes(&_PriceVoting.CallOpts, sortedVotes)
}

// Voters is a free data retrieval call binding the contract method 0xda58c7d9.
//
// Solidity: function voters(uint256 ) view returns(address)
func (_PriceVoting *PriceVotingCaller) Voters(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PriceVoting.contract.Call(opts, &out, "voters", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Voters is a free data retrieval call binding the contract method 0xda58c7d9.
//
// Solidity: function voters(uint256 ) view returns(address)
func (_PriceVoting *PriceVotingSession) Voters(arg0 *big.Int) (common.Address, error) {
	return _PriceVoting.Contract.Voters(&_PriceVoting.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xda58c7d9.
//
// Solidity: function voters(uint256 ) view returns(address)
func (_PriceVoting *PriceVotingCallerSession) Voters(arg0 *big.Int) (common.Address, error) {
	return _PriceVoting.Contract.Voters(&_PriceVoting.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(uint256)
func (_PriceVoting *PriceVotingCaller) Votes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceVoting.contract.Call(opts, &out, "votes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(uint256)
func (_PriceVoting *PriceVotingSession) Votes(arg0 common.Address) (*big.Int, error) {
	return _PriceVoting.Contract.Votes(&_PriceVoting.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(uint256)
func (_PriceVoting *PriceVotingCallerSession) Votes(arg0 common.Address) (*big.Int, error) {
	return _PriceVoting.Contract.Votes(&_PriceVoting.CallOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address stakingManager) returns()
func (_PriceVoting *PriceVotingTransactor) Initialize(opts *bind.TransactOpts, stakingManager common.Address) (*types.Transaction, error) {
	return _PriceVoting.contract.Transact(opts, "initialize", stakingManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address stakingManager) returns()
func (_PriceVoting *PriceVotingSession) Initialize(stakingManager common.Address) (*types.Transaction, error) {
	return _PriceVoting.Contract.Initialize(&_PriceVoting.TransactOpts, stakingManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address stakingManager) returns()
func (_PriceVoting *PriceVotingTransactorSession) Initialize(stakingManager common.Address) (*types.Transaction, error) {
	return _PriceVoting.Contract.Initialize(&_PriceVoting.TransactOpts, stakingManager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceVoting *PriceVotingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceVoting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceVoting *PriceVotingSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceVoting.Contract.RenounceOwnership(&_PriceVoting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceVoting *PriceVotingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceVoting.Contract.RenounceOwnership(&_PriceVoting.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceVoting *PriceVotingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PriceVoting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceVoting *PriceVotingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceVoting.Contract.TransferOwnership(&_PriceVoting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceVoting *PriceVotingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceVoting.Contract.TransferOwnership(&_PriceVoting.TransactOpts, newOwner)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 price) returns()
func (_PriceVoting *PriceVotingTransactor) Vote(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _PriceVoting.contract.Transact(opts, "vote", price)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 price) returns()
func (_PriceVoting *PriceVotingSession) Vote(price *big.Int) (*types.Transaction, error) {
	return _PriceVoting.Contract.Vote(&_PriceVoting.TransactOpts, price)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 price) returns()
func (_PriceVoting *PriceVotingTransactorSession) Vote(price *big.Int) (*types.Transaction, error) {
	return _PriceVoting.Contract.Vote(&_PriceVoting.TransactOpts, price)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PriceVoting *PriceVotingTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceVoting.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PriceVoting *PriceVotingSession) Withdraw() (*types.Transaction, error) {
	return _PriceVoting.Contract.Withdraw(&_PriceVoting.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PriceVoting *PriceVotingTransactorSession) Withdraw() (*types.Transaction, error) {
	return _PriceVoting.Contract.Withdraw(&_PriceVoting.TransactOpts)
}

// PriceVotingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PriceVoting contract.
type PriceVotingOwnershipTransferredIterator struct {
	Event *PriceVotingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PriceVotingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceVotingOwnershipTransferred)
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
		it.Event = new(PriceVotingOwnershipTransferred)
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
func (it *PriceVotingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceVotingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceVotingOwnershipTransferred represents a OwnershipTransferred event raised by the PriceVoting contract.
type PriceVotingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceVoting *PriceVotingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PriceVotingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceVoting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PriceVotingOwnershipTransferredIterator{contract: _PriceVoting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceVoting *PriceVotingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PriceVotingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceVoting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceVotingOwnershipTransferred)
				if err := _PriceVoting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PriceVoting *PriceVotingFilterer) ParseOwnershipTransferred(log types.Log) (*PriceVotingOwnershipTransferred, error) {
	event := new(PriceVotingOwnershipTransferred)
	if err := _PriceVoting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
