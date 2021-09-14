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

// RewardsManagerABI is the input ABI used to generate the binding from.
const RewardsManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractStakingManager\",\"name\":\"stakingManager\",\"type\":\"address\"},{\"internalType\":\"contractEpochsManager\",\"name\":\"epochsManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getRewardPoolBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getRewardPoolStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getDelegatorOwedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"incrementRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"}],\"name\":\"distributeReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"}],\"name\":\"initializeRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"removeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RewardsManagerBin is the compiled bytecode used for deploying new contracts.
var RewardsManagerBin = "0x608060405234801561001057600080fd5b50611bd5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c80637f13d2b71161008c578063baf81f6511610066578063baf81f65146101a1578063c0c53b8b146101b4578063d7ba4e33146101c7578063f2fde38b146101da57600080fd5b80637f13d2b7146101605780638da5cb5b14610173578063ac18de431461018e57600080fd5b80632d06177a116100bd5780632d06177a14610132578063325d64ba14610145578063715018a61461015857600080fd5b806302934c8d146100e457806309e8c9e01461010a5780632cb16fe11461011f575b600080fd5b6100f76100f23660046118c9565b6101ed565b6040519081526020015b60405180910390f35b61011d610118366004611882565b61047e565b005b6100f761012d36600461189a565b610928565b61011d610140366004611792565b610952565b6100f761015336600461189a565b6109c8565b61011d6109f0565b61011d61016e36600461190a565b610a56565b6033546040516001600160a01b039091168152602001610101565b61011d61019c366004611792565b610c66565b6100f76101af36600461189a565b610cda565b61011d6101c2366004611941565b610d3b565b61011d6101d5366004611882565b610e43565b61011d6101e8366004611792565b61135f565b606754604051637e6d64a560e01b81526004810185905260009182916001600160a01b0390911690637e6d64a5906024016101406040518083038186803b15801561023757600080fd5b505afa15801561024b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061026f9190611970565b905060008160200151116102ca5760405162461bcd60e51b815260206004820152601460248201527f45706f636820646f6573206e6f7420657869737400000000000000000000000060448201526064015b60405180910390fd5b6000606860006102da8888610cda565b81526020019081526020016000206040518060600160405290816000820154815260200160018201805480602002602001604051908101604052809291908181526020016000905b8282101561036a576000848152602090819020604080518082019091526002850290910180546001600160a01b03168252600190810154828401529083529092019101610322565b505050508152602001600282015481525050905080600001516000141561039657600092505050610477565b6000805b82602001515181101561043957856001600160a01b0316836020015182815181106103d557634e487b7160e01b600052603260045260246000fd5b6020026020010151600001516001600160a01b03161415610427578260200151818151811061041457634e487b7160e01b600052603260045260246000fd5b6020026020010151602001519150610439565b8061043181611b43565b91505061039a565b508061044b5760009350505050610477565b600061045f83600001518560800151611441565b905061047082846040015183611470565b9450505050505b9392505050565b600061048a8233610cda565b6000818152606860205260409020600281015491925090156105145760405162461bcd60e51b815260206004820152602860248201527f52657761726420706f6f6c2068617320616c7265616479206265656e20696e6960448201527f7469616c697a656400000000000000000000000000000000000000000000000060648201526084016102c1565b606754604051637e6d64a560e01b8152600481018590526000916001600160a01b031690637e6d64a5906024016101406040518083038186803b15801561055a57600080fd5b505afa15801561056e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105929190611970565b905080606001516000146105e85760405162461bcd60e51b815260206004820152601760248201527f45706f63682068617320616c726561647920656e64656400000000000000000060448201526064016102c1565b6066546040517f791936100000000000000000000000000000000000000000000000000000000081523360048201526000916001600160a01b03169063791936109060240160006040518083038186803b15801561064557600080fd5b505afa158015610659573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261068191908101906117ae565b905060008151116106fa5760405162461bcd60e51b815260206004820152602c60248201527f4d7573742068617665207374616b6520746f20696e74697469616c697a65206160448201527f2072657761726420706f6f6c000000000000000000000000000000000000000060648201526084016102c1565b60005b81518110156108845760665482516000916001600160a01b0316906382dda22d9085908590811061073e57634e487b7160e01b600052603260045260246000fd5b60209081029190910101516040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b039091166004820152336024820152604401604080518083038186803b1580156107a557600080fd5b505afa1580156107b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107dd9190611a0d565b60000151905084600101604051806040016040528085858151811061081257634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b039081168352918101949094528254600180820185556000948552938590208351600290920201805473ffffffffffffffffffffffffffffffffffffffff1916919092161781559201519101558061087c81611b43565b9150506106fd565b506066546040517fdf349ed50000000000000000000000000000000000000000000000000000000081523360048201526001600160a01b039091169063df349ed59060240160206040518083038186803b1580156108e157600080fd5b505afa1580156108f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109199190611a62565b83600201819055505050505050565b6000606860006109388585610cda565b815260200190815260200160002060020154905092915050565b6033546001600160a01b031633146109ac5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102c1565b6001600160a01b03166000908152606960205260409020439055565b6000606860006109d88585610cda565b81526020810191909152604001600020549392505050565b6033546001600160a01b03163314610a4a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102c1565b610a54600061148f565b565b33600090815260696020526040902054610ad85760405162461bcd60e51b815260206004820152603860248201527f4f6e6c7920636f6e74726f6c6c657273206f66207468697320636f6e7472616360448201527f742063616e2063616c6c20746869732066756e6374696f6e000000000000000060648201526084016102c1565b606754604051637e6d64a560e01b8152600481018590526000916001600160a01b031690637e6d64a5906024016101406040518083038186803b158015610b1e57600080fd5b505afa158015610b32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b569190611970565b90506000816020015111610bac5760405162461bcd60e51b815260206004820152601460248201527f45706f636820646f6573206e6f7420657869737400000000000000000000000060448201526064016102c1565b600060686000610bbc8787610cda565b815260200190815260200160002090506000816002015411610c465760405162461bcd60e51b815260206004820152603360248201527f52657761726420706f6f6c20686173206e6f74206265656e20636f6e7374727560448201527f6374656420666f7220746869732065706f63680000000000000000000000000060648201526084016102c1565b82816000016000828254610c5a9190611ad5565b90915550505050505050565b6033546001600160a01b03163314610cc05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102c1565b6001600160a01b0316600090815260696020526040812055565b60008282604051602001610d1d92919091825260601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602082015260340190565b60405160208183030381529060405280519060200120905092915050565b600054610100900460ff1680610d54575060005460ff16155b610db75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102c1565b600054610100900460ff16158015610dd9576000805461ffff19166101011790555b610de16114ee565b606580546001600160a01b0380871673ffffffffffffffffffffffffffffffffffffffff19928316179092556067805485841690831617905560668054928616929091169190911790558015610e3d576000805461ff00191690555b50505050565b606754604051637e6d64a560e01b8152600481018390526000916001600160a01b031690637e6d64a5906024016101406040518083038186803b158015610e8957600080fd5b505afa158015610e9d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ec19190611970565b90506000816020015111610f175760405162461bcd60e51b815260206004820152601460248201527f45706f636820646f6573206e6f7420657869737400000000000000000000000060448201526064016102c1565b6000816060015111610f6b5760405162461bcd60e51b815260206004820152601760248201527f45706f636820686173206e6f742079657420656e64656400000000000000000060448201526064016102c1565b8061010001518160600151610f809190611ad5565b43116110405760405162461bcd60e51b815260206004820152607060248201527f43616e206f6e6c7920646973747269627574652072657761726473206f6e636560448201527f20616c6c20706f737369626c65207469636b6574732068617665206265656e2060648201527f72656465656d6564202865706f63682e656e64426c6f636b202b2065706f636860848201527f2e7469636b65744475726174696f6e290000000000000000000000000000000060a482015260c4016102c1565b600061104c8333610cda565b90506000606860008381526020019081526020016000206040518060600160405290816000820154815260200160018201805480602002602001604051908101604052809291908181526020016000905b828210156110e5576000848152602090819020604080518082019091526002850290910180546001600160a01b0316825260019081015482840152908352909201910161109d565b505050508152602001600282015481525050905060008160000151116111735760405162461bcd60e51b815260206004820152602c60248201527f43616e206e6f742064697374726962757465207265776172642069662062616c60448201527f616e6365206973207a65726f000000000000000000000000000000000000000060648201526084016102c1565b600061118782600001518560800151611441565b90506000805b836020015151811015611296576000846020015182815181106111c057634e487b7160e01b600052603260045260246000fd5b6020026020010151905060006111df8260200151876040015187611470565b9050806111ed575050611284565b606554825160405163a9059cbb60e01b81526001600160a01b0391821660048201526024810184905291169063a9059cbb90604401602060405180830381600087803b15801561123c57600080fd5b505af1158015611250573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112749190611862565b5061127f8185611ad5565b935050505b8061128e81611b43565b91505061118d565b5082516000906112a7908390611b2c565b60655460405163a9059cbb60e01b8152336004820152602481018390529192506001600160a01b03169063a9059cbb90604401602060405180830381600087803b1580156112f457600080fd5b505af1158015611308573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061132c9190611862565b5060008581526068602052604081208181559061134c6001830182611708565b6002820160009055505050505050505050565b6033546001600160a01b031633146113b95760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102c1565b6001600160a01b0381166114355760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102c1565b61143e8161148f565b50565b600061271061146661ffff84166fffffffffffffffffffffffffffffffff8616611b0d565b6104779190611aed565b60008261147d8386611b0d565b6114879190611aed565b949350505050565b603380546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff1680611507575060005460ff16155b61156a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102c1565b600054610100900460ff1615801561158c576000805461ffff19166101011790555b6115946115b0565b61159c611661565b801561143e576000805461ff001916905550565b600054610100900460ff16806115c9575060005460ff16155b61162c5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102c1565b600054610100900460ff1615801561159c576000805461ffff1916610101179055801561143e576000805461ff001916905550565b600054610100900460ff168061167a575060005460ff16155b6116dd5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102c1565b600054610100900460ff161580156116ff576000805461ffff19166101011790555b61159c3361148f565b508054600082556002029060005260206000209081019061143e91905b8082111561175757805473ffffffffffffffffffffffffffffffffffffffff1916815560006001820155600201611725565b5090565b80516fffffffffffffffffffffffffffffffff8116811461177b57600080fd5b919050565b805161ffff8116811461177b57600080fd5b6000602082840312156117a3578081fd5b813561047781611b8a565b600060208083850312156117c0578182fd5b825167ffffffffffffffff808211156117d7578384fd5b818501915085601f8301126117ea578384fd5b8151818111156117fc576117fc611b74565b8060051b915061180d848301611aa4565b8181528481019084860184860187018a1015611827578788fd5b8795505b83861015611855578051945061184085611b8a565b8483526001959095019491860191860161182b565b5098975050505050505050565b600060208284031215611873578081fd5b81518015158114610477578182fd5b600060208284031215611893578081fd5b5035919050565b600080604083850312156118ac578081fd5b8235915060208301356118be81611b8a565b809150509250929050565b6000806000606084860312156118dd578081fd5b8335925060208401356118ef81611b8a565b915060408401356118ff81611b8a565b809150509250925092565b60008060006060848603121561191e578283fd5b83359250602084013561193081611b8a565b929592945050506040919091013590565b600080600060608486031215611955578283fd5b833561196081611b8a565b925060208401356118ef81611b8a565b60006101408284031215611982578081fd5b61198a611a7a565b825181526020830151602082015260408301516040820152606083015160608201526119b860808401611780565b608082015260a083015160a08201526119d360c0840161175b565b60c08201526119e460e0840161175b565b60e08201526101008381015190820152610120611a02818501611780565b908201529392505050565b600060408284031215611a1e578081fd5b6040516040810181811067ffffffffffffffff82111715611a4157611a41611b74565b604052825181526020830151611a5681611b8a565b60208201529392505050565b600060208284031215611a73578081fd5b5051919050565b604051610140810167ffffffffffffffff81118282101715611a9e57611a9e611b74565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715611acd57611acd611b74565b604052919050565b60008219821115611ae857611ae8611b5e565b500190565b600082611b0857634e487b7160e01b81526012600452602481fd5b500490565b6000816000190483118215151615611b2757611b27611b5e565b500290565b600082821015611b3e57611b3e611b5e565b500390565b6000600019821415611b5757611b57611b5e565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461143e57600080fdfea26469706673582212204803614d9562a64d78b01aade4997d57e5bab213c6ac7ad9c6ecd1331e1a926264736f6c63430008040033"

// DeployRewardsManager deploys a new Ethereum contract, binding an instance of RewardsManager to it.
func DeployRewardsManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RewardsManager, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardsManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RewardsManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RewardsManager{RewardsManagerCaller: RewardsManagerCaller{contract: contract}, RewardsManagerTransactor: RewardsManagerTransactor{contract: contract}, RewardsManagerFilterer: RewardsManagerFilterer{contract: contract}}, nil
}

// RewardsManager is an auto generated Go binding around an Ethereum contract.
type RewardsManager struct {
	RewardsManagerCaller     // Read-only binding to the contract
	RewardsManagerTransactor // Write-only binding to the contract
	RewardsManagerFilterer   // Log filterer for contract events
}

// RewardsManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardsManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardsManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardsManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardsManagerSession struct {
	Contract     *RewardsManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardsManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardsManagerCallerSession struct {
	Contract *RewardsManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// RewardsManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardsManagerTransactorSession struct {
	Contract     *RewardsManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RewardsManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardsManagerRaw struct {
	Contract *RewardsManager // Generic contract binding to access the raw methods on
}

// RewardsManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardsManagerCallerRaw struct {
	Contract *RewardsManagerCaller // Generic read-only contract binding to access the raw methods on
}

// RewardsManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardsManagerTransactorRaw struct {
	Contract *RewardsManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardsManager creates a new instance of RewardsManager, bound to a specific deployed contract.
func NewRewardsManager(address common.Address, backend bind.ContractBackend) (*RewardsManager, error) {
	contract, err := bindRewardsManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardsManager{RewardsManagerCaller: RewardsManagerCaller{contract: contract}, RewardsManagerTransactor: RewardsManagerTransactor{contract: contract}, RewardsManagerFilterer: RewardsManagerFilterer{contract: contract}}, nil
}

// NewRewardsManagerCaller creates a new read-only instance of RewardsManager, bound to a specific deployed contract.
func NewRewardsManagerCaller(address common.Address, caller bind.ContractCaller) (*RewardsManagerCaller, error) {
	contract, err := bindRewardsManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardsManagerCaller{contract: contract}, nil
}

// NewRewardsManagerTransactor creates a new write-only instance of RewardsManager, bound to a specific deployed contract.
func NewRewardsManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardsManagerTransactor, error) {
	contract, err := bindRewardsManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardsManagerTransactor{contract: contract}, nil
}

// NewRewardsManagerFilterer creates a new log filterer instance of RewardsManager, bound to a specific deployed contract.
func NewRewardsManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardsManagerFilterer, error) {
	contract, err := bindRewardsManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardsManagerFilterer{contract: contract}, nil
}

// bindRewardsManager binds a generic wrapper to an already deployed contract.
func bindRewardsManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardsManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardsManager *RewardsManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardsManager.Contract.RewardsManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardsManager *RewardsManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardsManager.Contract.RewardsManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardsManager *RewardsManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardsManager.Contract.RewardsManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardsManager *RewardsManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardsManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardsManager *RewardsManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardsManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardsManager *RewardsManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardsManager.Contract.contract.Transact(opts, method, params...)
}

// GetDelegatorOwedAmount is a free data retrieval call binding the contract method 0x02934c8d.
//
// Solidity: function getDelegatorOwedAmount(bytes32 epochId, address stakee, address staker) view returns(uint256)
func (_RewardsManager *RewardsManagerCaller) GetDelegatorOwedAmount(opts *bind.CallOpts, epochId [32]byte, stakee common.Address, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "getDelegatorOwedAmount", epochId, stakee, staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDelegatorOwedAmount is a free data retrieval call binding the contract method 0x02934c8d.
//
// Solidity: function getDelegatorOwedAmount(bytes32 epochId, address stakee, address staker) view returns(uint256)
func (_RewardsManager *RewardsManagerSession) GetDelegatorOwedAmount(epochId [32]byte, stakee common.Address, staker common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.GetDelegatorOwedAmount(&_RewardsManager.CallOpts, epochId, stakee, staker)
}

// GetDelegatorOwedAmount is a free data retrieval call binding the contract method 0x02934c8d.
//
// Solidity: function getDelegatorOwedAmount(bytes32 epochId, address stakee, address staker) view returns(uint256)
func (_RewardsManager *RewardsManagerCallerSession) GetDelegatorOwedAmount(epochId [32]byte, stakee common.Address, staker common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.GetDelegatorOwedAmount(&_RewardsManager.CallOpts, epochId, stakee, staker)
}

// GetKey is a free data retrieval call binding the contract method 0xbaf81f65.
//
// Solidity: function getKey(bytes32 epochId, address stakee) pure returns(bytes32)
func (_RewardsManager *RewardsManagerCaller) GetKey(opts *bind.CallOpts, epochId [32]byte, stakee common.Address) ([32]byte, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "getKey", epochId, stakee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKey is a free data retrieval call binding the contract method 0xbaf81f65.
//
// Solidity: function getKey(bytes32 epochId, address stakee) pure returns(bytes32)
func (_RewardsManager *RewardsManagerSession) GetKey(epochId [32]byte, stakee common.Address) ([32]byte, error) {
	return _RewardsManager.Contract.GetKey(&_RewardsManager.CallOpts, epochId, stakee)
}

// GetKey is a free data retrieval call binding the contract method 0xbaf81f65.
//
// Solidity: function getKey(bytes32 epochId, address stakee) pure returns(bytes32)
func (_RewardsManager *RewardsManagerCallerSession) GetKey(epochId [32]byte, stakee common.Address) ([32]byte, error) {
	return _RewardsManager.Contract.GetKey(&_RewardsManager.CallOpts, epochId, stakee)
}

// GetRewardPoolBalance is a free data retrieval call binding the contract method 0x325d64ba.
//
// Solidity: function getRewardPoolBalance(bytes32 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerCaller) GetRewardPoolBalance(opts *bind.CallOpts, epochId [32]byte, stakee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "getRewardPoolBalance", epochId, stakee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardPoolBalance is a free data retrieval call binding the contract method 0x325d64ba.
//
// Solidity: function getRewardPoolBalance(bytes32 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerSession) GetRewardPoolBalance(epochId [32]byte, stakee common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.GetRewardPoolBalance(&_RewardsManager.CallOpts, epochId, stakee)
}

// GetRewardPoolBalance is a free data retrieval call binding the contract method 0x325d64ba.
//
// Solidity: function getRewardPoolBalance(bytes32 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerCallerSession) GetRewardPoolBalance(epochId [32]byte, stakee common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.GetRewardPoolBalance(&_RewardsManager.CallOpts, epochId, stakee)
}

// GetRewardPoolStake is a free data retrieval call binding the contract method 0x2cb16fe1.
//
// Solidity: function getRewardPoolStake(bytes32 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerCaller) GetRewardPoolStake(opts *bind.CallOpts, epochId [32]byte, stakee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "getRewardPoolStake", epochId, stakee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardPoolStake is a free data retrieval call binding the contract method 0x2cb16fe1.
//
// Solidity: function getRewardPoolStake(bytes32 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerSession) GetRewardPoolStake(epochId [32]byte, stakee common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.GetRewardPoolStake(&_RewardsManager.CallOpts, epochId, stakee)
}

// GetRewardPoolStake is a free data retrieval call binding the contract method 0x2cb16fe1.
//
// Solidity: function getRewardPoolStake(bytes32 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerCallerSession) GetRewardPoolStake(epochId [32]byte, stakee common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.GetRewardPoolStake(&_RewardsManager.CallOpts, epochId, stakee)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RewardsManager *RewardsManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RewardsManager *RewardsManagerSession) Owner() (common.Address, error) {
	return _RewardsManager.Contract.Owner(&_RewardsManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RewardsManager *RewardsManagerCallerSession) Owner() (common.Address, error) {
	return _RewardsManager.Contract.Owner(&_RewardsManager.CallOpts)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_RewardsManager *RewardsManagerTransactor) AddManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _RewardsManager.contract.Transact(opts, "addManager", manager)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_RewardsManager *RewardsManagerSession) AddManager(manager common.Address) (*types.Transaction, error) {
	return _RewardsManager.Contract.AddManager(&_RewardsManager.TransactOpts, manager)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_RewardsManager *RewardsManagerTransactorSession) AddManager(manager common.Address) (*types.Transaction, error) {
	return _RewardsManager.Contract.AddManager(&_RewardsManager.TransactOpts, manager)
}

// DistributeReward is a paid mutator transaction binding the contract method 0xd7ba4e33.
//
// Solidity: function distributeReward(bytes32 epochId) returns()
func (_RewardsManager *RewardsManagerTransactor) DistributeReward(opts *bind.TransactOpts, epochId [32]byte) (*types.Transaction, error) {
	return _RewardsManager.contract.Transact(opts, "distributeReward", epochId)
}

// DistributeReward is a paid mutator transaction binding the contract method 0xd7ba4e33.
//
// Solidity: function distributeReward(bytes32 epochId) returns()
func (_RewardsManager *RewardsManagerSession) DistributeReward(epochId [32]byte) (*types.Transaction, error) {
	return _RewardsManager.Contract.DistributeReward(&_RewardsManager.TransactOpts, epochId)
}

// DistributeReward is a paid mutator transaction binding the contract method 0xd7ba4e33.
//
// Solidity: function distributeReward(bytes32 epochId) returns()
func (_RewardsManager *RewardsManagerTransactorSession) DistributeReward(epochId [32]byte) (*types.Transaction, error) {
	return _RewardsManager.Contract.DistributeReward(&_RewardsManager.TransactOpts, epochId)
}

// IncrementRewardPool is a paid mutator transaction binding the contract method 0x7f13d2b7.
//
// Solidity: function incrementRewardPool(bytes32 epochId, address stakee, uint256 amount) returns()
func (_RewardsManager *RewardsManagerTransactor) IncrementRewardPool(opts *bind.TransactOpts, epochId [32]byte, stakee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardsManager.contract.Transact(opts, "incrementRewardPool", epochId, stakee, amount)
}

// IncrementRewardPool is a paid mutator transaction binding the contract method 0x7f13d2b7.
//
// Solidity: function incrementRewardPool(bytes32 epochId, address stakee, uint256 amount) returns()
func (_RewardsManager *RewardsManagerSession) IncrementRewardPool(epochId [32]byte, stakee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardsManager.Contract.IncrementRewardPool(&_RewardsManager.TransactOpts, epochId, stakee, amount)
}

// IncrementRewardPool is a paid mutator transaction binding the contract method 0x7f13d2b7.
//
// Solidity: function incrementRewardPool(bytes32 epochId, address stakee, uint256 amount) returns()
func (_RewardsManager *RewardsManagerTransactorSession) IncrementRewardPool(epochId [32]byte, stakee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardsManager.Contract.IncrementRewardPool(&_RewardsManager.TransactOpts, epochId, stakee, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address token, address stakingManager, address epochsManager) returns()
func (_RewardsManager *RewardsManagerTransactor) Initialize(opts *bind.TransactOpts, token common.Address, stakingManager common.Address, epochsManager common.Address) (*types.Transaction, error) {
	return _RewardsManager.contract.Transact(opts, "initialize", token, stakingManager, epochsManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address token, address stakingManager, address epochsManager) returns()
func (_RewardsManager *RewardsManagerSession) Initialize(token common.Address, stakingManager common.Address, epochsManager common.Address) (*types.Transaction, error) {
	return _RewardsManager.Contract.Initialize(&_RewardsManager.TransactOpts, token, stakingManager, epochsManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address token, address stakingManager, address epochsManager) returns()
func (_RewardsManager *RewardsManagerTransactorSession) Initialize(token common.Address, stakingManager common.Address, epochsManager common.Address) (*types.Transaction, error) {
	return _RewardsManager.Contract.Initialize(&_RewardsManager.TransactOpts, token, stakingManager, epochsManager)
}

// InitializeRewardPool is a paid mutator transaction binding the contract method 0x09e8c9e0.
//
// Solidity: function initializeRewardPool(bytes32 epochId) returns()
func (_RewardsManager *RewardsManagerTransactor) InitializeRewardPool(opts *bind.TransactOpts, epochId [32]byte) (*types.Transaction, error) {
	return _RewardsManager.contract.Transact(opts, "initializeRewardPool", epochId)
}

// InitializeRewardPool is a paid mutator transaction binding the contract method 0x09e8c9e0.
//
// Solidity: function initializeRewardPool(bytes32 epochId) returns()
func (_RewardsManager *RewardsManagerSession) InitializeRewardPool(epochId [32]byte) (*types.Transaction, error) {
	return _RewardsManager.Contract.InitializeRewardPool(&_RewardsManager.TransactOpts, epochId)
}

// InitializeRewardPool is a paid mutator transaction binding the contract method 0x09e8c9e0.
//
// Solidity: function initializeRewardPool(bytes32 epochId) returns()
func (_RewardsManager *RewardsManagerTransactorSession) InitializeRewardPool(epochId [32]byte) (*types.Transaction, error) {
	return _RewardsManager.Contract.InitializeRewardPool(&_RewardsManager.TransactOpts, epochId)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_RewardsManager *RewardsManagerTransactor) RemoveManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _RewardsManager.contract.Transact(opts, "removeManager", manager)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_RewardsManager *RewardsManagerSession) RemoveManager(manager common.Address) (*types.Transaction, error) {
	return _RewardsManager.Contract.RemoveManager(&_RewardsManager.TransactOpts, manager)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_RewardsManager *RewardsManagerTransactorSession) RemoveManager(manager common.Address) (*types.Transaction, error) {
	return _RewardsManager.Contract.RemoveManager(&_RewardsManager.TransactOpts, manager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RewardsManager *RewardsManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardsManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RewardsManager *RewardsManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _RewardsManager.Contract.RenounceOwnership(&_RewardsManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RewardsManager *RewardsManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RewardsManager.Contract.RenounceOwnership(&_RewardsManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RewardsManager *RewardsManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RewardsManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RewardsManager *RewardsManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RewardsManager.Contract.TransferOwnership(&_RewardsManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RewardsManager *RewardsManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RewardsManager.Contract.TransferOwnership(&_RewardsManager.TransactOpts, newOwner)
}

// RewardsManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RewardsManager contract.
type RewardsManagerOwnershipTransferredIterator struct {
	Event *RewardsManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RewardsManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsManagerOwnershipTransferred)
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
		it.Event = new(RewardsManagerOwnershipTransferred)
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
func (it *RewardsManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsManagerOwnershipTransferred represents a OwnershipTransferred event raised by the RewardsManager contract.
type RewardsManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RewardsManager *RewardsManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RewardsManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RewardsManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RewardsManagerOwnershipTransferredIterator{contract: _RewardsManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RewardsManager *RewardsManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RewardsManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RewardsManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsManagerOwnershipTransferred)
				if err := _RewardsManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RewardsManager *RewardsManagerFilterer) ParseOwnershipTransferred(log types.Log) (*RewardsManagerOwnershipTransferred, error) {
	event := new(RewardsManagerOwnershipTransferred)
	if err := _RewardsManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
