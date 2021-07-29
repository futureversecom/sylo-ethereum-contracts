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

// SyloTicketingTicket is an auto generated low-level Go binding around an user-defined struct.
type SyloTicketingTicket struct {
	Sender           common.Address
	Receiver         common.Address
	FaceValue        *big.Int
	WinProb          *big.Int
	ExpirationBlock  *big.Int
	ReceiverRandHash [32]byte
	SenderNonce      uint32
}

// SyloTicketingABI is the input ABI used to generate the binding from.
const SyloTicketingABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"escrow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usedTickets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractListings\",\"name\":\"listings\",\"type\":\"address\"},{\"internalType\":\"contractStakingManager\",\"name\":\"stakingManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unlockDuration\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newUnlockDuration\",\"type\":\"uint256\"}],\"name\":\"setUnlockDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositEscrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"winProb\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiverRandHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"senderNonce\",\"type\":\"uint32\"}],\"internalType\":\"structSyloTicketing.Ticket\",\"name\":\"ticket\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"receiverRand\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"winProb\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiverRandHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"senderNonce\",\"type\":\"uint32\"}],\"internalType\":\"structSyloTicketing.Ticket\",\"name\":\"ticket\",\"type\":\"tuple\"}],\"name\":\"getTicketHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SyloTicketingBin is the compiled bytecode used for deploying new contracts.
var SyloTicketingBin = "0x608060405234801561001057600080fd5b50611ff3806100206000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80638a1fcd6011610097578063d088a23111610066578063d088a231146101fb578063dd90076914610203578063f2fde38b14610216578063fc7e286d1461022957600080fd5b80638a1fcd60146101b15780638da5cb5b146101ba578063cdba7342146101d5578063cf756fdf146101e857600080fd5b8063715018a6116100d3578063715018a61461016d57806371e6af6e1461017557806372b0d90c146101965780637d6babb4146101a957600080fd5b8063169e1ca8146101055780633ccfd60b1461011a5780634e055fe21461012257806359a515ba14610135575b600080fd5b610118610113366004611dc6565b610273565b005b6101186103a5565b610118610130366004611d1b565b6103b0565b610158610143366004611b73565b606a6020526000908152604090205460ff1681565b60405190151581526020015b60405180910390f35b610118610b4b565b610188610183366004611d00565b610baf565b604051908152602001610164565b6101186101a4366004611a89565b610c7d565b610188610d8f565b61018860685481565b6033546040516001600160a01b039091168152602001610164565b6101186101e3366004611dc6565b610e69565b6101186101f6366004611b8b565b610ee7565b610118611000565b610118610211366004611b73565b61106a565b610118610224366004611a89565b6110c9565b610258610237366004611a89565b60696020526000908152604090208054600182015460029092015490919083565b60408051938452602084019290925290820152606001610164565b6001600160a01b03811660009081526069602052604090206002810154156102e25760405162461bcd60e51b815260206004820152601e60248201527f43616e6e6f74206465706f736974207768696c6520756e6c6f636b696e67000060448201526064015b60405180910390fd5b828160000160008282546102f69190611e99565b90915550506065546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590526001600160a01b03909116906323b872dd906064015b602060405180830381600087803b15801561036757600080fd5b505af115801561037b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061039f9190611b59565b50505050565b6103ae33610c7d565b565b60006103bb84610baf565b90506103c9848285856111ab565b83516001600160a01b03166000908152606960205260409081902090850151600182015482546103f99190611e99565b101561046d5760405162461bcd60e51b815260206004820152602760248201527f53656e64657220646f65736e2774206861766520656e6f7567682066756e647360448201527f20746f207061790000000000000000000000000000000000000000000000000060648201526084016102d9565b6000828152606a60209081526040808320805460ff191660011790556066549188015190517f084af0b20000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169063084af0b29060240160006040518083038186803b1580156104e957600080fd5b505afa1580156104fd573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526105259190810190611bdb565b90508060600151151560011515146105a55760405162461bcd60e51b815260206004820152602960248201527f5469636b6574207265636569766572206d757374206861766520612076616c6960448201527f64206c697374696e67000000000000000000000000000000000000000000000060648201526084016102d9565b60675460208701516040517fdf349ed50000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600092919091169063df349ed59060240160206040518083038186803b15801561060c57600080fd5b505afa158015610620573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106449190611dae565b9050806106935760405162461bcd60e51b815260206004820152601f60248201527f5469636b6574207265636569766572206d7573742068617665207374616b650060448201526064016102d9565b8254604088015111156107c9576065546020880151845460405163a9059cbb60e01b81526001600160a01b039283166004820152602481019190915291169063a9059cbb90604401602060405180830381600087803b1580156106f557600080fd5b505af1158015610709573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061072d9190611b59565b50606554600184015460405163a9059cbb60e01b81526001600160a01b039092166004830181905260248301919091529063a9059cbb90604401602060405180830381600087803b15801561078157600080fd5b505af1158015610795573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107b99190611b59565b5060008084556001840155610b42565b604087015183546107d991611492565b8355604087015160208301516000916064916107f8919060ff16611ef6565b6108029190611ed6565b60675460208a01516040517f791936100000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292935060009291169063791936109060240160006040518083038186803b15801561086a57600080fd5b505afa15801561087e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526108a69190810190611aa5565b90508160005b82518163ffffffff161015610a925760675483516000916001600160a01b0316906382dda22d90869063ffffffff86169081106108f957634e487b7160e01b600052603260045260246000fd5b60200260200101518e602001516040518363ffffffff1660e01b81526004016109389291906001600160a01b0392831681529116602082015260400190565b604080518083038186803b15801561094f57600080fd5b505afa158015610963573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109879190611cab565b905060008686836000015161099c9190611ef6565b6109a69190611ed6565b90506109b28185611f15565b60655486519195506001600160a01b03169063a9059cbb90879063ffffffff87169081106109f057634e487b7160e01b600052603260045260246000fd5b6020026020010151836040518363ffffffff1660e01b8152600401610a2a9291906001600160a01b03929092168252602082015260400190565b602060405180830381600087803b158015610a4457600080fd5b505af1158015610a58573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a7c9190611b59565b5050508080610a8a90611f58565b9150506108ac565b50600081848c60400151610aa69190611f15565b610ab09190611e99565b60655460208d015160405163a9059cbb60e01b81526001600160a01b03918216600482015260248101849052929350169063a9059cbb90604401602060405180830381600087803b158015610b0457600080fd5b505af1158015610b18573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b3c9190611b59565b50505050505b50505050505050565b6033546001600160a01b03163314610ba55760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102d9565b6103ae60006114a5565b80516020808301516040808501516060860151608087015160a088015160c08901519451600098610c60989097969101606097881b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000090811682529690971b9095166014870152602886019390935260488501919091526068840152608883015260e01b7fffffffff000000000000000000000000000000000000000000000000000000001660a882015260ac0190565b604051602081830303815290604052805190602001209050919050565b3360009081526069602052604090206002810154610cdd5760405162461bcd60e51b815260206004820152601560248201527f4465706f73697473206e6f7420756e6c6f636b6564000000000000000000000060448201526064016102d9565b43816002015410610d305760405162461bcd60e51b815260206004820152601a60248201527f556e6c6f636b20706572696f64206e6f7420636f6d706c65746500000000000060448201526064016102d9565b60018101548154600091610d4391611e99565b600080845560018401819055600284015560655460405163a9059cbb60e01b81526001600160a01b0386811660048301526024820184905292935091169063a9059cbb9060440161034d565b3360009081526069602052604081208054151580610db1575060008160010154115b610dfd5760405162461bcd60e51b815260206004820152601360248201527f4e6f7468696e6720746f2077697468647261770000000000000000000000000060448201526064016102d9565b600281015415610e4f5760405162461bcd60e51b815260206004820152601a60248201527f556e6c6f636b20616c726561647920696e2070726f677265737300000000000060448201526064016102d9565b606854610e5c9043611e99565b6002909101819055919050565b6001600160a01b0381166000908152606960205260409020600281015415610ed35760405162461bcd60e51b815260206004820152601e60248201527f43616e6e6f74206465706f736974207768696c6520756e6c6f636b696e67000060448201526064016102d9565b828160010160008282546102f69190611e99565b600054610100900460ff1680610f00575060005460ff16155b610f635760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102d9565b600054610100900460ff16158015610f85576000805461ffff19166101011790555b610f8d61150f565b606580546001600160a01b038088167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560668054878416908316179055606780549286169290911691909117905560688290558015610ff9576000805461ff00191690555b5050505050565b33600090815260696020526040902060028101546110605760405162461bcd60e51b815260206004820152601a60248201527f4e6f7420756e6c6f636b696e672c2063616e6e6f74206c6f636b00000000000060448201526064016102d9565b6000600290910155565b6033546001600160a01b031633146110c45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102d9565b606855565b6033546001600160a01b031633146111235760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102d9565b6001600160a01b03811661119f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102d9565b6111a8816114a5565b50565b83516001600160a01b03166112025760405162461bcd60e51b815260206004820152601560248201527f5469636b65742073656e646572206973206e756c6c000000000000000000000060448201526064016102d9565b60208401516001600160a01b031661125c5760405162461bcd60e51b815260206004820152601760248201527f5469636b6574207265636569766572206973206e756c6c00000000000000000060448201526064016102d9565b60808401511580611271575043846080015110155b6112bd5760405162461bcd60e51b815260206004820152601260248201527f5469636b6574206861732065787069726564000000000000000000000000000060448201526064016102d9565b6000838152606a602052604090205460ff161561131c5760405162461bcd60e51b815260206004820152601760248201527f5469636b657420616c72656164792072656465656d656400000000000000000060448201526064016102d9565b60a084015160408051602081018590520160405160208183030381529060405280519060200120146113b65760405162461bcd60e51b815260206004820152603360248201527f48617368206f6620726563656976657252616e6420646f65736e2774206d617460448201527f636820726563656976657252616e64486173680000000000000000000000000060648201526084016102d9565b6113c5818560000151856115d1565b6114375760405162461bcd60e51b815260206004820152602560248201527f5469636b657420646f65736e2774206861766520612076616c6964207369676e60448201527f617475726500000000000000000000000000000000000000000000000000000060648201526084016102d9565b611446818386606001516115f9565b61039f5760405162461bcd60e51b815260206004820152601660248201527f5469636b6574206973206e6f7420612077696e6e65720000000000000000000060448201526064016102d9565b600061149e8284611f15565b9392505050565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff1680611528575060005460ff16155b61158b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102d9565b600054610100900460ff161580156115ad576000805461ffff19166101011790555b6115b5611630565b6115bd6116e1565b80156111a8576000805461ff001916905550565b6000826001600160a01b03166115e78386611788565b6001600160a01b031614949350505050565b600081848460405160200161160f929190611df5565b60408051601f19818403018152919052805160209091012010949350505050565b600054610100900460ff1680611649575060005460ff16155b6116ac5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102d9565b600054610100900460ff161580156115bd576000805461ffff191661010117905580156111a8576000805461ff001916905550565b600054610100900460ff16806116fa575060005460ff16155b61175d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102d9565b600054610100900460ff1615801561177f576000805461ffff19166101011790555b6115bd336114a5565b600081516041146117db5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016102d9565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156118685760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016102d9565b601b8160ff1610156118825761187f601b82611eb1565b90505b8060ff16601b1415801561189a57508060ff16601c14155b156118f25760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b60648201526084016102d9565b6040805160008082526020820180845289905260ff841692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015611946573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166119a95760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016102d9565b9695505050505050565b805180151581146119c357600080fd5b919050565b600060e082840312156119d9578081fd5b60405160e0810181811067ffffffffffffffff821117156119fc576119fc611f92565b6040529050808235611a0d81611fa8565b81526020830135611a1d81611fa8565b8060208301525060408301356040820152606083013560608201526080830135608082015260a083013560a0820152611a5860c08401611a64565b60c08201525092915050565b803563ffffffff811681146119c357600080fd5b805160ff811681146119c357600080fd5b600060208284031215611a9a578081fd5b813561149e81611fa8565b60006020808385031215611ab7578182fd5b825167ffffffffffffffff80821115611ace578384fd5b818501915085601f830112611ae1578384fd5b815181811115611af357611af3611f92565b8060051b9150611b04848301611e40565b8181528481019084860184860187018a1015611b1e578788fd5b8795505b83861015611b4c5780519450611b3785611fa8565b84835260019590950194918601918601611b22565b5098975050505050505050565b600060208284031215611b6a578081fd5b61149e826119b3565b600060208284031215611b84578081fd5b5035919050565b60008060008060808587031215611ba0578283fd5b8435611bab81611fa8565b93506020850135611bbb81611fa8565b92506040850135611bcb81611fa8565b9396929550929360600135925050565b60006020808385031215611bed578182fd5b825167ffffffffffffffff80821115611c04578384fd5b9084019060808287031215611c17578384fd5b611c1f611e17565b825182811115611c2d578586fd5b83019150601f82018713611c3f578485fd5b8151611c52611c4d82611e71565b611e40565b8181528886838601011115611c65578687fd5b611c7482878301888701611f2c565b825250611c82838501611a78565b8482015260408301516040820152611c9c606084016119b3565b60608201529695505050505050565b600060408284031215611cbc578081fd5b6040516040810181811067ffffffffffffffff82111715611cdf57611cdf611f92565b604052825181526020830151611cf481611fa8565b60208201529392505050565b600060e08284031215611d11578081fd5b61149e83836119c8565b60008060006101208486031215611d30578081fd5b611d3a85856119c8565b925060e0840135915061010084013567ffffffffffffffff811115611d5d578182fd5b8401601f81018613611d6d578182fd5b8035611d7b611c4d82611e71565b818152876020838501011115611d8f578384fd5b8160208401602083013783602083830101528093505050509250925092565b600060208284031215611dbf578081fd5b5051919050565b60008060408385031215611dd8578182fd5b823591506020830135611dea81611fa8565b809150509250929050565b60008351611e07818460208801611f2c565b9190910191825250602001919050565b6040516080810167ffffffffffffffff81118282101715611e3a57611e3a611f92565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715611e6957611e69611f92565b604052919050565b600067ffffffffffffffff821115611e8b57611e8b611f92565b50601f01601f191660200190565b60008219821115611eac57611eac611f7c565b500190565b600060ff821660ff84168060ff03821115611ece57611ece611f7c565b019392505050565b600082611ef157634e487b7160e01b81526012600452602481fd5b500490565b6000816000190483118215151615611f1057611f10611f7c565b500290565b600082821015611f2757611f27611f7c565b500390565b60005b83811015611f47578181015183820152602001611f2f565b8381111561039f5750506000910152565b600063ffffffff80831681811415611f7257611f72611f7c565b6001019392505050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146111a857600080fdfea26469706673582212201d76edf28e3bcfddd1332349c5d4d8fe22b887e12fdf8294ec0b54d00c134ca164736f6c63430008040033"

// DeploySyloTicketing deploys a new Ethereum contract, binding an instance of SyloTicketing to it.
func DeploySyloTicketing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SyloTicketing, error) {
	parsed, err := abi.JSON(strings.NewReader(SyloTicketingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SyloTicketingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SyloTicketing{SyloTicketingCaller: SyloTicketingCaller{contract: contract}, SyloTicketingTransactor: SyloTicketingTransactor{contract: contract}, SyloTicketingFilterer: SyloTicketingFilterer{contract: contract}}, nil
}

// SyloTicketing is an auto generated Go binding around an Ethereum contract.
type SyloTicketing struct {
	SyloTicketingCaller     // Read-only binding to the contract
	SyloTicketingTransactor // Write-only binding to the contract
	SyloTicketingFilterer   // Log filterer for contract events
}

// SyloTicketingCaller is an auto generated read-only Go binding around an Ethereum contract.
type SyloTicketingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyloTicketingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SyloTicketingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyloTicketingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SyloTicketingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyloTicketingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SyloTicketingSession struct {
	Contract     *SyloTicketing    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SyloTicketingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SyloTicketingCallerSession struct {
	Contract *SyloTicketingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SyloTicketingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SyloTicketingTransactorSession struct {
	Contract     *SyloTicketingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SyloTicketingRaw is an auto generated low-level Go binding around an Ethereum contract.
type SyloTicketingRaw struct {
	Contract *SyloTicketing // Generic contract binding to access the raw methods on
}

// SyloTicketingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SyloTicketingCallerRaw struct {
	Contract *SyloTicketingCaller // Generic read-only contract binding to access the raw methods on
}

// SyloTicketingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SyloTicketingTransactorRaw struct {
	Contract *SyloTicketingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSyloTicketing creates a new instance of SyloTicketing, bound to a specific deployed contract.
func NewSyloTicketing(address common.Address, backend bind.ContractBackend) (*SyloTicketing, error) {
	contract, err := bindSyloTicketing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SyloTicketing{SyloTicketingCaller: SyloTicketingCaller{contract: contract}, SyloTicketingTransactor: SyloTicketingTransactor{contract: contract}, SyloTicketingFilterer: SyloTicketingFilterer{contract: contract}}, nil
}

// NewSyloTicketingCaller creates a new read-only instance of SyloTicketing, bound to a specific deployed contract.
func NewSyloTicketingCaller(address common.Address, caller bind.ContractCaller) (*SyloTicketingCaller, error) {
	contract, err := bindSyloTicketing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SyloTicketingCaller{contract: contract}, nil
}

// NewSyloTicketingTransactor creates a new write-only instance of SyloTicketing, bound to a specific deployed contract.
func NewSyloTicketingTransactor(address common.Address, transactor bind.ContractTransactor) (*SyloTicketingTransactor, error) {
	contract, err := bindSyloTicketing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SyloTicketingTransactor{contract: contract}, nil
}

// NewSyloTicketingFilterer creates a new log filterer instance of SyloTicketing, bound to a specific deployed contract.
func NewSyloTicketingFilterer(address common.Address, filterer bind.ContractFilterer) (*SyloTicketingFilterer, error) {
	contract, err := bindSyloTicketing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SyloTicketingFilterer{contract: contract}, nil
}

// bindSyloTicketing binds a generic wrapper to an already deployed contract.
func bindSyloTicketing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SyloTicketingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SyloTicketing *SyloTicketingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SyloTicketing.Contract.SyloTicketingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SyloTicketing *SyloTicketingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyloTicketing.Contract.SyloTicketingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SyloTicketing *SyloTicketingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SyloTicketing.Contract.SyloTicketingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SyloTicketing *SyloTicketingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SyloTicketing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SyloTicketing *SyloTicketingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyloTicketing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SyloTicketing *SyloTicketingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SyloTicketing.Contract.contract.Transact(opts, method, params...)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256 escrow, uint256 penalty, uint256 unlockAt)
func (_SyloTicketing *SyloTicketingCaller) Deposits(opts *bind.CallOpts, arg0 common.Address) (struct {
	Escrow   *big.Int
	Penalty  *big.Int
	UnlockAt *big.Int
}, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "deposits", arg0)

	outstruct := new(struct {
		Escrow   *big.Int
		Penalty  *big.Int
		UnlockAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Escrow = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Penalty = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UnlockAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256 escrow, uint256 penalty, uint256 unlockAt)
func (_SyloTicketing *SyloTicketingSession) Deposits(arg0 common.Address) (struct {
	Escrow   *big.Int
	Penalty  *big.Int
	UnlockAt *big.Int
}, error) {
	return _SyloTicketing.Contract.Deposits(&_SyloTicketing.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256 escrow, uint256 penalty, uint256 unlockAt)
func (_SyloTicketing *SyloTicketingCallerSession) Deposits(arg0 common.Address) (struct {
	Escrow   *big.Int
	Penalty  *big.Int
	UnlockAt *big.Int
}, error) {
	return _SyloTicketing.Contract.Deposits(&_SyloTicketing.CallOpts, arg0)
}

// GetTicketHash is a free data retrieval call binding the contract method 0x71e6af6e.
//
// Solidity: function getTicketHash((address,address,uint256,uint256,uint256,bytes32,uint32) ticket) pure returns(bytes32)
func (_SyloTicketing *SyloTicketingCaller) GetTicketHash(opts *bind.CallOpts, ticket SyloTicketingTicket) ([32]byte, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "getTicketHash", ticket)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTicketHash is a free data retrieval call binding the contract method 0x71e6af6e.
//
// Solidity: function getTicketHash((address,address,uint256,uint256,uint256,bytes32,uint32) ticket) pure returns(bytes32)
func (_SyloTicketing *SyloTicketingSession) GetTicketHash(ticket SyloTicketingTicket) ([32]byte, error) {
	return _SyloTicketing.Contract.GetTicketHash(&_SyloTicketing.CallOpts, ticket)
}

// GetTicketHash is a free data retrieval call binding the contract method 0x71e6af6e.
//
// Solidity: function getTicketHash((address,address,uint256,uint256,uint256,bytes32,uint32) ticket) pure returns(bytes32)
func (_SyloTicketing *SyloTicketingCallerSession) GetTicketHash(ticket SyloTicketingTicket) ([32]byte, error) {
	return _SyloTicketing.Contract.GetTicketHash(&_SyloTicketing.CallOpts, ticket)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SyloTicketing *SyloTicketingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SyloTicketing *SyloTicketingSession) Owner() (common.Address, error) {
	return _SyloTicketing.Contract.Owner(&_SyloTicketing.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SyloTicketing *SyloTicketingCallerSession) Owner() (common.Address, error) {
	return _SyloTicketing.Contract.Owner(&_SyloTicketing.CallOpts)
}

// UnlockDuration is a free data retrieval call binding the contract method 0x8a1fcd60.
//
// Solidity: function unlockDuration() view returns(uint256)
func (_SyloTicketing *SyloTicketingCaller) UnlockDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "unlockDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockDuration is a free data retrieval call binding the contract method 0x8a1fcd60.
//
// Solidity: function unlockDuration() view returns(uint256)
func (_SyloTicketing *SyloTicketingSession) UnlockDuration() (*big.Int, error) {
	return _SyloTicketing.Contract.UnlockDuration(&_SyloTicketing.CallOpts)
}

// UnlockDuration is a free data retrieval call binding the contract method 0x8a1fcd60.
//
// Solidity: function unlockDuration() view returns(uint256)
func (_SyloTicketing *SyloTicketingCallerSession) UnlockDuration() (*big.Int, error) {
	return _SyloTicketing.Contract.UnlockDuration(&_SyloTicketing.CallOpts)
}

// UsedTickets is a free data retrieval call binding the contract method 0x59a515ba.
//
// Solidity: function usedTickets(bytes32 ) view returns(bool)
func (_SyloTicketing *SyloTicketingCaller) UsedTickets(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "usedTickets", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedTickets is a free data retrieval call binding the contract method 0x59a515ba.
//
// Solidity: function usedTickets(bytes32 ) view returns(bool)
func (_SyloTicketing *SyloTicketingSession) UsedTickets(arg0 [32]byte) (bool, error) {
	return _SyloTicketing.Contract.UsedTickets(&_SyloTicketing.CallOpts, arg0)
}

// UsedTickets is a free data retrieval call binding the contract method 0x59a515ba.
//
// Solidity: function usedTickets(bytes32 ) view returns(bool)
func (_SyloTicketing *SyloTicketingCallerSession) UsedTickets(arg0 [32]byte) (bool, error) {
	return _SyloTicketing.Contract.UsedTickets(&_SyloTicketing.CallOpts, arg0)
}

// DepositEscrow is a paid mutator transaction binding the contract method 0x169e1ca8.
//
// Solidity: function depositEscrow(uint256 amount, address account) returns()
func (_SyloTicketing *SyloTicketingTransactor) DepositEscrow(opts *bind.TransactOpts, amount *big.Int, account common.Address) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "depositEscrow", amount, account)
}

// DepositEscrow is a paid mutator transaction binding the contract method 0x169e1ca8.
//
// Solidity: function depositEscrow(uint256 amount, address account) returns()
func (_SyloTicketing *SyloTicketingSession) DepositEscrow(amount *big.Int, account common.Address) (*types.Transaction, error) {
	return _SyloTicketing.Contract.DepositEscrow(&_SyloTicketing.TransactOpts, amount, account)
}

// DepositEscrow is a paid mutator transaction binding the contract method 0x169e1ca8.
//
// Solidity: function depositEscrow(uint256 amount, address account) returns()
func (_SyloTicketing *SyloTicketingTransactorSession) DepositEscrow(amount *big.Int, account common.Address) (*types.Transaction, error) {
	return _SyloTicketing.Contract.DepositEscrow(&_SyloTicketing.TransactOpts, amount, account)
}

// DepositPenalty is a paid mutator transaction binding the contract method 0xcdba7342.
//
// Solidity: function depositPenalty(uint256 amount, address account) returns()
func (_SyloTicketing *SyloTicketingTransactor) DepositPenalty(opts *bind.TransactOpts, amount *big.Int, account common.Address) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "depositPenalty", amount, account)
}

// DepositPenalty is a paid mutator transaction binding the contract method 0xcdba7342.
//
// Solidity: function depositPenalty(uint256 amount, address account) returns()
func (_SyloTicketing *SyloTicketingSession) DepositPenalty(amount *big.Int, account common.Address) (*types.Transaction, error) {
	return _SyloTicketing.Contract.DepositPenalty(&_SyloTicketing.TransactOpts, amount, account)
}

// DepositPenalty is a paid mutator transaction binding the contract method 0xcdba7342.
//
// Solidity: function depositPenalty(uint256 amount, address account) returns()
func (_SyloTicketing *SyloTicketingTransactorSession) DepositPenalty(amount *big.Int, account common.Address) (*types.Transaction, error) {
	return _SyloTicketing.Contract.DepositPenalty(&_SyloTicketing.TransactOpts, amount, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address token, address listings, address stakingManager, uint256 _unlockDuration) returns()
func (_SyloTicketing *SyloTicketingTransactor) Initialize(opts *bind.TransactOpts, token common.Address, listings common.Address, stakingManager common.Address, _unlockDuration *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "initialize", token, listings, stakingManager, _unlockDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address token, address listings, address stakingManager, uint256 _unlockDuration) returns()
func (_SyloTicketing *SyloTicketingSession) Initialize(token common.Address, listings common.Address, stakingManager common.Address, _unlockDuration *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.Initialize(&_SyloTicketing.TransactOpts, token, listings, stakingManager, _unlockDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address token, address listings, address stakingManager, uint256 _unlockDuration) returns()
func (_SyloTicketing *SyloTicketingTransactorSession) Initialize(token common.Address, listings common.Address, stakingManager common.Address, _unlockDuration *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.Initialize(&_SyloTicketing.TransactOpts, token, listings, stakingManager, _unlockDuration)
}

// LockDeposits is a paid mutator transaction binding the contract method 0xd088a231.
//
// Solidity: function lockDeposits() returns()
func (_SyloTicketing *SyloTicketingTransactor) LockDeposits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "lockDeposits")
}

// LockDeposits is a paid mutator transaction binding the contract method 0xd088a231.
//
// Solidity: function lockDeposits() returns()
func (_SyloTicketing *SyloTicketingSession) LockDeposits() (*types.Transaction, error) {
	return _SyloTicketing.Contract.LockDeposits(&_SyloTicketing.TransactOpts)
}

// LockDeposits is a paid mutator transaction binding the contract method 0xd088a231.
//
// Solidity: function lockDeposits() returns()
func (_SyloTicketing *SyloTicketingTransactorSession) LockDeposits() (*types.Transaction, error) {
	return _SyloTicketing.Contract.LockDeposits(&_SyloTicketing.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x4e055fe2.
//
// Solidity: function redeem((address,address,uint256,uint256,uint256,bytes32,uint32) ticket, uint256 receiverRand, bytes sig) returns()
func (_SyloTicketing *SyloTicketingTransactor) Redeem(opts *bind.TransactOpts, ticket SyloTicketingTicket, receiverRand *big.Int, sig []byte) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "redeem", ticket, receiverRand, sig)
}

// Redeem is a paid mutator transaction binding the contract method 0x4e055fe2.
//
// Solidity: function redeem((address,address,uint256,uint256,uint256,bytes32,uint32) ticket, uint256 receiverRand, bytes sig) returns()
func (_SyloTicketing *SyloTicketingSession) Redeem(ticket SyloTicketingTicket, receiverRand *big.Int, sig []byte) (*types.Transaction, error) {
	return _SyloTicketing.Contract.Redeem(&_SyloTicketing.TransactOpts, ticket, receiverRand, sig)
}

// Redeem is a paid mutator transaction binding the contract method 0x4e055fe2.
//
// Solidity: function redeem((address,address,uint256,uint256,uint256,bytes32,uint32) ticket, uint256 receiverRand, bytes sig) returns()
func (_SyloTicketing *SyloTicketingTransactorSession) Redeem(ticket SyloTicketingTicket, receiverRand *big.Int, sig []byte) (*types.Transaction, error) {
	return _SyloTicketing.Contract.Redeem(&_SyloTicketing.TransactOpts, ticket, receiverRand, sig)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SyloTicketing *SyloTicketingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SyloTicketing *SyloTicketingSession) RenounceOwnership() (*types.Transaction, error) {
	return _SyloTicketing.Contract.RenounceOwnership(&_SyloTicketing.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SyloTicketing *SyloTicketingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SyloTicketing.Contract.RenounceOwnership(&_SyloTicketing.TransactOpts)
}

// SetUnlockDuration is a paid mutator transaction binding the contract method 0xdd900769.
//
// Solidity: function setUnlockDuration(uint256 newUnlockDuration) returns()
func (_SyloTicketing *SyloTicketingTransactor) SetUnlockDuration(opts *bind.TransactOpts, newUnlockDuration *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "setUnlockDuration", newUnlockDuration)
}

// SetUnlockDuration is a paid mutator transaction binding the contract method 0xdd900769.
//
// Solidity: function setUnlockDuration(uint256 newUnlockDuration) returns()
func (_SyloTicketing *SyloTicketingSession) SetUnlockDuration(newUnlockDuration *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.SetUnlockDuration(&_SyloTicketing.TransactOpts, newUnlockDuration)
}

// SetUnlockDuration is a paid mutator transaction binding the contract method 0xdd900769.
//
// Solidity: function setUnlockDuration(uint256 newUnlockDuration) returns()
func (_SyloTicketing *SyloTicketingTransactorSession) SetUnlockDuration(newUnlockDuration *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.SetUnlockDuration(&_SyloTicketing.TransactOpts, newUnlockDuration)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SyloTicketing *SyloTicketingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SyloTicketing *SyloTicketingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SyloTicketing.Contract.TransferOwnership(&_SyloTicketing.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SyloTicketing *SyloTicketingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SyloTicketing.Contract.TransferOwnership(&_SyloTicketing.TransactOpts, newOwner)
}

// UnlockDeposits is a paid mutator transaction binding the contract method 0x7d6babb4.
//
// Solidity: function unlockDeposits() returns(uint256)
func (_SyloTicketing *SyloTicketingTransactor) UnlockDeposits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "unlockDeposits")
}

// UnlockDeposits is a paid mutator transaction binding the contract method 0x7d6babb4.
//
// Solidity: function unlockDeposits() returns(uint256)
func (_SyloTicketing *SyloTicketingSession) UnlockDeposits() (*types.Transaction, error) {
	return _SyloTicketing.Contract.UnlockDeposits(&_SyloTicketing.TransactOpts)
}

// UnlockDeposits is a paid mutator transaction binding the contract method 0x7d6babb4.
//
// Solidity: function unlockDeposits() returns(uint256)
func (_SyloTicketing *SyloTicketingTransactorSession) UnlockDeposits() (*types.Transaction, error) {
	return _SyloTicketing.Contract.UnlockDeposits(&_SyloTicketing.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_SyloTicketing *SyloTicketingTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_SyloTicketing *SyloTicketingSession) Withdraw() (*types.Transaction, error) {
	return _SyloTicketing.Contract.Withdraw(&_SyloTicketing.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_SyloTicketing *SyloTicketingTransactorSession) Withdraw() (*types.Transaction, error) {
	return _SyloTicketing.Contract.Withdraw(&_SyloTicketing.TransactOpts)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x72b0d90c.
//
// Solidity: function withdrawTo(address account) returns()
func (_SyloTicketing *SyloTicketingTransactor) WithdrawTo(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "withdrawTo", account)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x72b0d90c.
//
// Solidity: function withdrawTo(address account) returns()
func (_SyloTicketing *SyloTicketingSession) WithdrawTo(account common.Address) (*types.Transaction, error) {
	return _SyloTicketing.Contract.WithdrawTo(&_SyloTicketing.TransactOpts, account)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x72b0d90c.
//
// Solidity: function withdrawTo(address account) returns()
func (_SyloTicketing *SyloTicketingTransactorSession) WithdrawTo(account common.Address) (*types.Transaction, error) {
	return _SyloTicketing.Contract.WithdrawTo(&_SyloTicketing.TransactOpts, account)
}

// SyloTicketingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SyloTicketing contract.
type SyloTicketingOwnershipTransferredIterator struct {
	Event *SyloTicketingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SyloTicketingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyloTicketingOwnershipTransferred)
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
		it.Event = new(SyloTicketingOwnershipTransferred)
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
func (it *SyloTicketingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyloTicketingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyloTicketingOwnershipTransferred represents a OwnershipTransferred event raised by the SyloTicketing contract.
type SyloTicketingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SyloTicketing *SyloTicketingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SyloTicketingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SyloTicketing.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SyloTicketingOwnershipTransferredIterator{contract: _SyloTicketing.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SyloTicketing *SyloTicketingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SyloTicketingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SyloTicketing.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyloTicketingOwnershipTransferred)
				if err := _SyloTicketing.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SyloTicketing *SyloTicketingFilterer) ParseOwnershipTransferred(log types.Log) (*SyloTicketingOwnershipTransferred, error) {
	event := new(SyloTicketingOwnershipTransferred)
	if err := _SyloTicketing.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
