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
const RewardsManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractStakingManager\",\"name\":\"stakingManager\",\"type\":\"address\"},{\"internalType\":\"contractEpochsManager\",\"name\":\"epochsManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getRewardPoolStakeeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getRewardPoolStakersBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getRewardPoolBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getRewardPoolStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getDelegatorOwedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"incrementRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"}],\"name\":\"initializeRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"epochIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"epochId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"calculateClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"removeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RewardsManagerBin is the compiled bytecode used for deploying new contracts.
var RewardsManagerBin = "0x608060405234801561001057600080fd5b50611a5a806100206000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80637f13d2b711610097578063af0d160f11610066578063af0d160f146101fb578063baf81f651461020e578063c0c53b8b14610221578063f2fde38b1461023457600080fd5b80637f13d2b7146101a75780638da5cb5b146101ba57806394a9c49c146101d5578063ac18de43146101e857600080fd5b8063325d64ba116100d3578063325d64ba1461016657806350c480ce146101795780635f8190b11461018c578063715018a61461019f57600080fd5b806302934c8d1461010557806309e8c9e01461012b5780632cb16fe1146101405780632d06177a14610153575b600080fd5b6101186101133660046117a3565b610247565b6040519081526020015b60405180910390f35b61013e61013936600461175c565b610425565b005b61011861014e366004611774565b6106b0565b61013e610161366004611666565b6106da565b610118610174366004611774565b610750565b61013e610187366004611682565b610770565b61011861019a366004611774565b6108d8565b61013e610900565b61013e6101b53660046117e4565b610966565b6033546040516001600160a01b039091168152602001610122565b6101186101e3366004611774565b610bad565b61013e6101f6366004611666565b610bd7565b610118610209366004611774565b610c4b565b61011861021c366004611774565b6110fd565b61013e61022f36600461181b565b61115e565b61013e610242366004611666565b611270565b606754604051637e6d64a560e01b81526004810185905260009182916001600160a01b0390911690637e6d64a5906024016101406040518083038186803b15801561029157600080fd5b505afa1580156102a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102c9919061184a565b905060008160200151116103245760405162461bcd60e51b815260206004820152601460248201527f45706f636820646f6573206e6f7420657869737400000000000000000000000060448201526064015b60405180910390fd5b60006068600061033488886110fd565b8152602001908152602001600020905080600101546000141561035c5760009250505061041e565b6066546002820154604051635c8e77e960e11b81526001600160a01b03878116600483015288811660248301526044820192909252600092919091169063b91cefd29060640160206040518083038186803b1580156103ba57600080fd5b505afa1580156103ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103f291906118e7565b905080610405576000935050505061041e565b6104188183600401548460010154611352565b93505050505b9392505050565b600061043182336110fd565b6000818152606860205260409020600481015491925090156104bb5760405162461bcd60e51b815260206004820152602860248201527f52657761726420706f6f6c2068617320616c7265616479206265656e20696e6960448201527f7469616c697a6564000000000000000000000000000000000000000000000000606482015260840161031b565b606754604051637e6d64a560e01b8152600481018590526000916001600160a01b031690637e6d64a5906024016101406040518083038186803b15801561050157600080fd5b505afa158015610515573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610539919061184a565b9050806060015160001461058f5760405162461bcd60e51b815260206004820152601760248201527f45706f63682068617320616c726561647920656e646564000000000000000000604482015260640161031b565b6066546040517f64084d4e0000000000000000000000000000000000000000000000000000000081523360048201526000916001600160a01b0316906364084d4e9060240160206040518083038186803b1580156105ec57600080fd5b505afa158015610600573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061062491906118e7565b90506000811161069c5760405162461bcd60e51b815260206004820152602c60248201527f4d7573742068617665207374616b6520746f20696e74697469616c697a65206160448201527f2072657761726420706f6f6c0000000000000000000000000000000000000000606482015260840161031b565b436002840155600490920191909155505050565b6000606860006106c085856110fd565b815260200190815260200160002060040154905092915050565b6033546001600160a01b031633146107345760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161031b565b6001600160a01b03166000908152606960205260409020439055565b600061075c8383610bad565b61076684846108d8565b61041e919061195a565b6000805b83518110156108345760006107b08583815181106107a257634e487b7160e01b600052603260045260246000fd5b602002602001015185610c4b565b905080156108145760006107eb8684815181106107dd57634e487b7160e01b600052603260045260246000fd5b6020026020010151866110fd565b60009081526068602090815260408083203384526003019091529020805460ff19166001179055505b61081e818461195a565b925050808061082c906119c8565b915050610774565b506065546040517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018390526001600160a01b039091169063a9059cbb90604401602060405180830381600087803b15801561089a57600080fd5b505af11580156108ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d2919061173c565b50505050565b6000606860006108e885856110fd565b81526020810191909152604001600020549392505050565b6033546001600160a01b0316331461095a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161031b565b6109646000611371565b565b336000908152606960205260409020546109e85760405162461bcd60e51b815260206004820152603860248201527f4f6e6c7920636f6e74726f6c6c657273206f66207468697320636f6e7472616360448201527f742063616e2063616c6c20746869732066756e6374696f6e0000000000000000606482015260840161031b565b606754604051637e6d64a560e01b8152600481018590526000916001600160a01b031690637e6d64a5906024016101406040518083038186803b158015610a2e57600080fd5b505afa158015610a42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a66919061184a565b90506000816020015111610abc5760405162461bcd60e51b815260206004820152601460248201527f45706f636820646f6573206e6f74206578697374000000000000000000000000604482015260640161031b565b600060686000610acc87876110fd565b815260200190815260200160002090506000816004015411610b565760405162461bcd60e51b815260206004820152603360248201527f52657761726420706f6f6c20686173206e6f74206265656e20636f6e7374727560448201527f6374656420666f7220746869732065706f636800000000000000000000000000606482015260840161031b565b6000610b668484608001516113db565b9050610b7281856119b1565b826000016000828254610b85919061195a565b9250508190555080826001016000828254610ba0919061195a565b9091555050505050505050565b600060686000610bbd85856110fd565b815260200190815260200160002060010154905092915050565b6033546001600160a01b03163314610c315760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161031b565b6001600160a01b0316600090815260696020526040812055565b606754604051637e6d64a560e01b81526004810184905260009182916001600160a01b0390911690637e6d64a5906024016101406040518083038186803b158015610c9557600080fd5b505afa158015610ca9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ccd919061184a565b90506000816020015111610d235760405162461bcd60e51b815260206004820152601460248201527f45706f636820646f6573206e6f74206578697374000000000000000000000000604482015260640161031b565b6000816060015111610d775760405162461bcd60e51b815260206004820152601760248201527f45706f636820686173206e6f742079657420656e646564000000000000000000604482015260640161031b565b8061010001518160600151610d8c919061195a565b4311610e4c5760405162461bcd60e51b815260206004820152606b60248201527f43616e206f6e6c7920636c61696d2072657761726473206f6e636520616c6c2060448201527f706f737369626c65207469636b6574732068617665206265656e20726564656560648201527f6d6564202865706f63682e656e64426c6f636b202b2065706f63682e7469636b60848201527f65744475726174696f6e2900000000000000000000000000000000000000000060a482015260c40161031b565b6000610e5885856110fd565b6000818152606860205260408120600181015481549394509092610e7c919061195a565b11610eef5760405162461bcd60e51b815260206004820152602760248201527f43616e206e6f7420636c61696d207265776172642069662062616c616e63652060448201527f6973207a65726f00000000000000000000000000000000000000000000000000606482015260840161031b565b33600090815260038201602052604090205460ff1615610f765760405162461bcd60e51b8152602060048201526024808201527f43616e206e6f7420636c61696d2062616c616e6365206d6f7265207468616e2060448201527f6f6e636500000000000000000000000000000000000000000000000000000000606482015260840161031b565b6066546002820154604051635c8e77e960e11b81523360048201526001600160a01b0388811660248301526044820192909252600092919091169063b91cefd29060640160206040518083038186803b158015610fd257600080fd5b505afa158015610fe6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061100a91906118e7565b905060008111806110235750336001600160a01b038716145b6110bb5760405162461bcd60e51b815260206004820152604c60248201527f4d757374206861766520686164207374616b6520666f7220746869732065706f60448201527f6368206f7220626520746865207374616b656520696e206f7264657220746f2060648201527f636c61696d207265776172640000000000000000000000000000000000000000608482015260a40161031b565b60006110d08284600401548560010154611352565b9050336001600160a01b03881614156110f25782546110ef908261195a565b90505b979650505050505050565b6000828260405160200161114092919091825260601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602082015260340190565b60405160208183030381529060405280519060200120905092915050565b600054610100900460ff1680611177575060005460ff16155b6111da5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161031b565b600054610100900460ff161580156111fc576000805461ffff19166101011790555b61120461140a565b606580546001600160a01b038087167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560678054858416908316179055606680549286169290911691909117905580156108d2576000805461ff001916905550505050565b6033546001600160a01b031633146112ca5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161031b565b6001600160a01b0381166113465760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161031b565b61134f81611371565b50565b60008261135f8386611992565b6113699190611972565b949350505050565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600061271061140061ffff84166fffffffffffffffffffffffffffffffff8616611992565b61041e9190611972565b600054610100900460ff1680611423575060005460ff16155b6114865760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161031b565b600054610100900460ff161580156114a8576000805461ffff19166101011790555b6114b06114cc565b6114b861157d565b801561134f576000805461ff001916905550565b600054610100900460ff16806114e5575060005460ff16155b6115485760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161031b565b600054610100900460ff161580156114b8576000805461ffff1916610101179055801561134f576000805461ff001916905550565b600054610100900460ff1680611596575060005460ff16155b6115f95760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161031b565b600054610100900460ff1615801561161b576000805461ffff19166101011790555b6114b833611371565b803561162f81611a0f565b919050565b80516fffffffffffffffffffffffffffffffff8116811461162f57600080fd5b805161ffff8116811461162f57600080fd5b600060208284031215611677578081fd5b813561041e81611a0f565b60008060408385031215611694578081fd5b823567ffffffffffffffff808211156116ab578283fd5b818501915085601f8301126116be578283fd5b81356020828211156116d2576116d26119f9565b8160051b92506116e3818401611929565b8281528181019085830185870184018b10156116fd578788fd5b8796505b8487101561171f578035835260019690960195918301918301611701565b50965061172f9050878201611624565b9450505050509250929050565b60006020828403121561174d578081fd5b8151801515811461041e578182fd5b60006020828403121561176d578081fd5b5035919050565b60008060408385031215611786578182fd5b82359150602083013561179881611a0f565b809150509250929050565b6000806000606084860312156117b7578081fd5b8335925060208401356117c981611a0f565b915060408401356117d981611a0f565b809150509250925092565b6000806000606084860312156117f8578283fd5b83359250602084013561180a81611a0f565b929592945050506040919091013590565b60008060006060848603121561182f578283fd5b833561183a81611a0f565b925060208401356117c981611a0f565b6000610140828403121561185c578081fd5b6118646118ff565b8251815260208301516020820152604083015160408201526060830151606082015261189260808401611654565b608082015260a083015160a08201526118ad60c08401611634565b60c08201526118be60e08401611634565b60e082015261010083810151908201526101206118dc818501611654565b908201529392505050565b6000602082840312156118f8578081fd5b5051919050565b604051610140810167ffffffffffffffff81118282101715611923576119236119f9565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715611952576119526119f9565b604052919050565b6000821982111561196d5761196d6119e3565b500190565b60008261198d57634e487b7160e01b81526012600452602481fd5b500490565b60008160001904831182151516156119ac576119ac6119e3565b500290565b6000828210156119c3576119c36119e3565b500390565b60006000198214156119dc576119dc6119e3565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461134f57600080fdfea2646970667358221220f1ac04609f5e1974f10bb302dc4fc53a52742ee5fb03aac10faeb02d79c36c4664736f6c63430008040033"

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

// CalculateClaim is a free data retrieval call binding the contract method 0xaf0d160f.
//
// Solidity: function calculateClaim(bytes32 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerCaller) CalculateClaim(opts *bind.CallOpts, epochId [32]byte, stakee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "calculateClaim", epochId, stakee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateClaim is a free data retrieval call binding the contract method 0xaf0d160f.
//
// Solidity: function calculateClaim(bytes32 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerSession) CalculateClaim(epochId [32]byte, stakee common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.CalculateClaim(&_RewardsManager.CallOpts, epochId, stakee)
}

// CalculateClaim is a free data retrieval call binding the contract method 0xaf0d160f.
//
// Solidity: function calculateClaim(bytes32 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerCallerSession) CalculateClaim(epochId [32]byte, stakee common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.CalculateClaim(&_RewardsManager.CallOpts, epochId, stakee)
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

// ClaimRewards is a paid mutator transaction binding the contract method 0x50c480ce.
//
// Solidity: function claimRewards(bytes32[] epochIds, address stakee) returns()
func (_RewardsManager *RewardsManagerTransactor) ClaimRewards(opts *bind.TransactOpts, epochIds [][32]byte, stakee common.Address) (*types.Transaction, error) {
	return _RewardsManager.contract.Transact(opts, "claimRewards", epochIds, stakee)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x50c480ce.
//
// Solidity: function claimRewards(bytes32[] epochIds, address stakee) returns()
func (_RewardsManager *RewardsManagerSession) ClaimRewards(epochIds [][32]byte, stakee common.Address) (*types.Transaction, error) {
	return _RewardsManager.Contract.ClaimRewards(&_RewardsManager.TransactOpts, epochIds, stakee)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x50c480ce.
//
// Solidity: function claimRewards(bytes32[] epochIds, address stakee) returns()
func (_RewardsManager *RewardsManagerTransactorSession) ClaimRewards(epochIds [][32]byte, stakee common.Address) (*types.Transaction, error) {
	return _RewardsManager.Contract.ClaimRewards(&_RewardsManager.TransactOpts, epochIds, stakee)
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
