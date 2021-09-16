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
const RewardsManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractStakingManager\",\"name\":\"stakingManager\",\"type\":\"address\"},{\"internalType\":\"contractEpochsManager\",\"name\":\"epochsManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getRewardPoolStakeeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getRewardPoolStakersBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getRewardPoolBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getRewardPoolStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getDelegatorOwedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"incrementRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"}],\"name\":\"initializeRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"removeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RewardsManagerBin is the compiled bytecode used for deploying new contracts.
var RewardsManagerBin = "0x608060405234801561001057600080fd5b50611871806100206000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80637f13d2b711610097578063baf81f6511610066578063baf81f65146101dd578063c0c53b8b146101f0578063f2fde38b14610203578063f9d0ff061461021657600080fd5b80637f13d2b7146101895780638da5cb5b1461019c57806394a9c49c146101b7578063ac18de43146101ca57600080fd5b80632d06177a116100d35780632d06177a14610148578063325d64ba1461015b5780635f8190b11461016e578063715018a61461018157600080fd5b806302934c8d146100fa57806309e8c9e0146101205780632cb16fe114610135575b600080fd5b61010d61010836600461160e565b610229565b6040519081526020015b60405180910390f35b61013361012e3660046115c7565b610407565b005b61010d6101433660046115df565b610692565b61013361015636600461158b565b6106bc565b61010d6101693660046115df565b610732565b61010d61017c3660046115df565b610752565b61013361077a565b61013361019736600461164f565b6107e0565b6033546040516001600160a01b039091168152602001610117565b61010d6101c53660046115df565b610a27565b6101336101d836600461158b565b610a51565b61010d6101eb3660046115df565b610ac5565b6101336101fe366004611686565b610b26565b61013361021136600461158b565b610c39565b6101336102243660046115df565b610d1b565b606754604051637e6d64a560e01b81526004810185905260009182916001600160a01b0390911690637e6d64a5906024016101406040518083038186803b15801561027357600080fd5b505afa158015610287573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102ab91906116b5565b905060008160200151116103065760405162461bcd60e51b815260206004820152601460248201527f45706f636820646f6573206e6f7420657869737400000000000000000000000060448201526064015b60405180910390fd5b6000606860006103168888610ac5565b8152602001908152602001600020905080600101546000141561033e57600092505050610400565b6066546002820154604051635c8e77e960e11b81526001600160a01b03878116600483015288811660248301526044820192909252600092919091169063b91cefd29060640160206040518083038186803b15801561039c57600080fd5b505afa1580156103b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103d49190611752565b9050806103e75760009350505050610400565b6103fa8183600401548460010154611282565b93505050505b9392505050565b60006104138233610ac5565b60008181526068602052604090206004810154919250901561049d5760405162461bcd60e51b815260206004820152602860248201527f52657761726420706f6f6c2068617320616c7265616479206265656e20696e6960448201527f7469616c697a656400000000000000000000000000000000000000000000000060648201526084016102fd565b606754604051637e6d64a560e01b8152600481018590526000916001600160a01b031690637e6d64a5906024016101406040518083038186803b1580156104e357600080fd5b505afa1580156104f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061051b91906116b5565b905080606001516000146105715760405162461bcd60e51b815260206004820152601760248201527f45706f63682068617320616c726561647920656e64656400000000000000000060448201526064016102fd565b6066546040517f64084d4e0000000000000000000000000000000000000000000000000000000081523360048201526000916001600160a01b0316906364084d4e9060240160206040518083038186803b1580156105ce57600080fd5b505afa1580156105e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106069190611752565b90506000811161067e5760405162461bcd60e51b815260206004820152602c60248201527f4d7573742068617665207374616b6520746f20696e74697469616c697a65206160448201527f2072657761726420706f6f6c000000000000000000000000000000000000000060648201526084016102fd565b436002840155600490920191909155505050565b6000606860006106a28585610ac5565b815260200190815260200160002060040154905092915050565b6033546001600160a01b031633146107165760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102fd565b6001600160a01b03166000908152606960205260409020439055565b600061073e8383610a27565b6107488484610752565b61040091906117a2565b6000606860006107628585610ac5565b81526020810191909152604001600020549392505050565b6033546001600160a01b031633146107d45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102fd565b6107de60006112a1565b565b336000908152606960205260409020546108625760405162461bcd60e51b815260206004820152603860248201527f4f6e6c7920636f6e74726f6c6c657273206f66207468697320636f6e7472616360448201527f742063616e2063616c6c20746869732066756e6374696f6e000000000000000060648201526084016102fd565b606754604051637e6d64a560e01b8152600481018590526000916001600160a01b031690637e6d64a5906024016101406040518083038186803b1580156108a857600080fd5b505afa1580156108bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e091906116b5565b905060008160200151116109365760405162461bcd60e51b815260206004820152601460248201527f45706f636820646f6573206e6f7420657869737400000000000000000000000060448201526064016102fd565b6000606860006109468787610ac5565b8152602001908152602001600020905060008160040154116109d05760405162461bcd60e51b815260206004820152603360248201527f52657761726420706f6f6c20686173206e6f74206265656e20636f6e7374727560448201527f6374656420666f7220746869732065706f63680000000000000000000000000060648201526084016102fd565b60006109e084846080015161130b565b90506109ec81856117f9565b8260000160008282546109ff91906117a2565b9250508190555080826001016000828254610a1a91906117a2565b9091555050505050505050565b600060686000610a378585610ac5565b815260200190815260200160002060010154905092915050565b6033546001600160a01b03163314610aab5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102fd565b6001600160a01b0316600090815260696020526040812055565b60008282604051602001610b0892919091825260601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602082015260340190565b60405160208183030381529060405280519060200120905092915050565b600054610100900460ff1680610b3f575060005460ff16155b610ba25760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102fd565b600054610100900460ff16158015610bc4576000805461ffff19166101011790555b610bcc61133a565b606580546001600160a01b038087167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556067805485841690831617905560668054928616929091169190911790558015610c33576000805461ff00191690555b50505050565b6033546001600160a01b03163314610c935760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102fd565b6001600160a01b038116610d0f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102fd565b610d18816112a1565b50565b606754604051637e6d64a560e01b8152600481018490526000916001600160a01b031690637e6d64a5906024016101406040518083038186803b158015610d6157600080fd5b505afa158015610d75573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d9991906116b5565b90506000816020015111610def5760405162461bcd60e51b815260206004820152601460248201527f45706f636820646f6573206e6f7420657869737400000000000000000000000060448201526064016102fd565b6000816060015111610e435760405162461bcd60e51b815260206004820152601760248201527f45706f636820686173206e6f742079657420656e64656400000000000000000060448201526064016102fd565b8061010001518160600151610e5891906117a2565b4311610f185760405162461bcd60e51b815260206004820152606b60248201527f43616e206f6e6c7920636c61696d2072657761726473206f6e636520616c6c2060448201527f706f737369626c65207469636b6574732068617665206265656e20726564656560648201527f6d6564202865706f63682e656e64426c6f636b202b2065706f63682e7469636b60848201527f65744475726174696f6e2900000000000000000000000000000000000000000060a482015260c4016102fd565b6000610f248484610ac5565b6000818152606860205260408120600181015481549394509092610f4891906117a2565b11610fbb5760405162461bcd60e51b815260206004820152602760248201527f43616e206e6f7420636c61696d207265776172642069662062616c616e63652060448201527f6973207a65726f0000000000000000000000000000000000000000000000000060648201526084016102fd565b33600090815260038201602052604090205460ff16156110425760405162461bcd60e51b8152602060048201526024808201527f43616e206e6f7420636c61696d2062616c616e6365206d6f7265207468616e2060448201527f6f6e63650000000000000000000000000000000000000000000000000000000060648201526084016102fd565b6066546002820154604051635c8e77e960e11b81523360048201526001600160a01b0387811660248301526044820192909252600092919091169063b91cefd29060640160206040518083038186803b15801561109e57600080fd5b505afa1580156110b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110d69190611752565b905060008111806110ef5750336001600160a01b038616145b6111875760405162461bcd60e51b815260206004820152604c60248201527f4d757374206861766520686164207374616b6520666f7220746869732065706f60448201527f6368206f7220626520746865207374616b656520696e206f7264657220746f2060648201527f636c61696d207265776172640000000000000000000000000000000000000000608482015260a4016102fd565b600061119c8284600401548560010154611282565b9050336001600160a01b03871614156111be5782546111bb90826117a2565b90505b33600081815260038501602052604090819020805460ff1916600117905560655490517fa9059cbb0000000000000000000000000000000000000000000000000000000081526004810192909252602482018390526001600160a01b03169063a9059cbb90604401602060405180830381600087803b15801561124057600080fd5b505af1158015611254573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061127891906115a7565b5050505050505050565b60008261128f83866117da565b61129991906117ba565b949350505050565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600061271061133061ffff84166fffffffffffffffffffffffffffffffff86166117da565b61040091906117ba565b600054610100900460ff1680611353575060005460ff16155b6113b65760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102fd565b600054610100900460ff161580156113d8576000805461ffff19166101011790555b6113e06113fc565b6113e86114ad565b8015610d18576000805461ff001916905550565b600054610100900460ff1680611415575060005460ff16155b6114785760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102fd565b600054610100900460ff161580156113e8576000805461ffff19166101011790558015610d18576000805461ff001916905550565b600054610100900460ff16806114c6575060005460ff16155b6115295760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102fd565b600054610100900460ff1615801561154b576000805461ffff19166101011790555b6113e8336112a1565b80516fffffffffffffffffffffffffffffffff8116811461157457600080fd5b919050565b805161ffff8116811461157457600080fd5b60006020828403121561159c578081fd5b813561040081611826565b6000602082840312156115b8578081fd5b81518015158114610400578182fd5b6000602082840312156115d8578081fd5b5035919050565b600080604083850312156115f1578081fd5b82359150602083013561160381611826565b809150509250929050565b600080600060608486031215611622578081fd5b83359250602084013561163481611826565b9150604084013561164481611826565b809150509250925092565b600080600060608486031215611663578283fd5b83359250602084013561167581611826565b929592945050506040919091013590565b60008060006060848603121561169a578283fd5b83356116a581611826565b9250602084013561163481611826565b600061014082840312156116c7578081fd5b6116cf61176a565b825181526020830151602082015260408301516040820152606083015160608201526116fd60808401611579565b608082015260a083015160a082015261171860c08401611554565b60c082015261172960e08401611554565b60e08201526101008381015190820152610120611747818501611579565b908201529392505050565b600060208284031215611763578081fd5b5051919050565b604051610140810167ffffffffffffffff8111828210171561179c57634e487b7160e01b600052604160045260246000fd5b60405290565b600082198211156117b5576117b5611810565b500190565b6000826117d557634e487b7160e01b81526012600452602481fd5b500490565b60008160001904831182151516156117f4576117f4611810565b500290565b60008282101561180b5761180b611810565b500390565b634e487b7160e01b600052601160045260246000fd5b6001600160a01b0381168114610d1857600080fdfea26469706673582212204aad199c28356d9dbf23961718dcb36682a1830f0058ddf7ce3e6c74875d15b964736f6c63430008040033"

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

// GetRewardPoolStakeeBalance is a free data retrieval call binding the contract method 0x5f8190b1.
//
// Solidity: function getRewardPoolStakeeBalance(bytes32 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerCaller) GetRewardPoolStakeeBalance(opts *bind.CallOpts, epochId [32]byte, stakee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "getRewardPoolStakeeBalance", epochId, stakee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardPoolStakeeBalance is a free data retrieval call binding the contract method 0x5f8190b1.
//
// Solidity: function getRewardPoolStakeeBalance(bytes32 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerSession) GetRewardPoolStakeeBalance(epochId [32]byte, stakee common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.GetRewardPoolStakeeBalance(&_RewardsManager.CallOpts, epochId, stakee)
}

// GetRewardPoolStakeeBalance is a free data retrieval call binding the contract method 0x5f8190b1.
//
// Solidity: function getRewardPoolStakeeBalance(bytes32 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerCallerSession) GetRewardPoolStakeeBalance(epochId [32]byte, stakee common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.GetRewardPoolStakeeBalance(&_RewardsManager.CallOpts, epochId, stakee)
}

// GetRewardPoolStakersBalance is a free data retrieval call binding the contract method 0x94a9c49c.
//
// Solidity: function getRewardPoolStakersBalance(bytes32 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerCaller) GetRewardPoolStakersBalance(opts *bind.CallOpts, epochId [32]byte, stakee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "getRewardPoolStakersBalance", epochId, stakee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardPoolStakersBalance is a free data retrieval call binding the contract method 0x94a9c49c.
//
// Solidity: function getRewardPoolStakersBalance(bytes32 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerSession) GetRewardPoolStakersBalance(epochId [32]byte, stakee common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.GetRewardPoolStakersBalance(&_RewardsManager.CallOpts, epochId, stakee)
}

// GetRewardPoolStakersBalance is a free data retrieval call binding the contract method 0x94a9c49c.
//
// Solidity: function getRewardPoolStakersBalance(bytes32 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerCallerSession) GetRewardPoolStakersBalance(epochId [32]byte, stakee common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.GetRewardPoolStakersBalance(&_RewardsManager.CallOpts, epochId, stakee)
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

// ClaimReward is a paid mutator transaction binding the contract method 0xf9d0ff06.
//
// Solidity: function claimReward(bytes32 epochId, address stakee) returns()
func (_RewardsManager *RewardsManagerTransactor) ClaimReward(opts *bind.TransactOpts, epochId [32]byte, stakee common.Address) (*types.Transaction, error) {
	return _RewardsManager.contract.Transact(opts, "claimReward", epochId, stakee)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xf9d0ff06.
//
// Solidity: function claimReward(bytes32 epochId, address stakee) returns()
func (_RewardsManager *RewardsManagerSession) ClaimReward(epochId [32]byte, stakee common.Address) (*types.Transaction, error) {
	return _RewardsManager.Contract.ClaimReward(&_RewardsManager.TransactOpts, epochId, stakee)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xf9d0ff06.
//
// Solidity: function claimReward(bytes32 epochId, address stakee) returns()
func (_RewardsManager *RewardsManagerTransactorSession) ClaimReward(epochId [32]byte, stakee common.Address) (*types.Transaction, error) {
	return _RewardsManager.Contract.ClaimReward(&_RewardsManager.TransactOpts, epochId, stakee)
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
