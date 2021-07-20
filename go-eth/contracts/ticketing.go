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
	GenerationBlock  *big.Int
	ReceiverRandHash [32]byte
	SenderRandHash   uint32
}

// SyloTicketingABI is the input ABI used to generate the binding from.
const SyloTicketingABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"escrow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"faceValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ticketLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usedTickets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winProb\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractListings\",\"name\":\"listings\",\"type\":\"address\"},{\"internalType\":\"contractStakingManager\",\"name\":\"stakingManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unlockDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_winProb\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ticketLength\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newUnlockDuration\",\"type\":\"uint256\"}],\"name\":\"setUnlockDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_faceValue\",\"type\":\"uint256\"}],\"name\":\"setFaceValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_winProb\",\"type\":\"uint256\"}],\"name\":\"setWinProb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositEscrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"generationBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiverRandHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"senderRandHash\",\"type\":\"uint32\"}],\"internalType\":\"structSyloTicketing.Ticket\",\"name\":\"ticket\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"receiverRand\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"generationBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiverRandHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"senderRandHash\",\"type\":\"uint32\"}],\"internalType\":\"structSyloTicketing.Ticket\",\"name\":\"ticket\",\"type\":\"tuple\"}],\"name\":\"calculatePayout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"generationBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiverRandHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"senderRandHash\",\"type\":\"uint32\"}],\"internalType\":\"structSyloTicketing.Ticket\",\"name\":\"ticket\",\"type\":\"tuple\"}],\"name\":\"getTicketHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SyloTicketingBin is the compiled bytecode used for deploying new contracts.
var SyloTicketingBin = "0x608060405234801561001057600080fd5b5061219b806100206000396000f3fe608060405234801561001057600080fd5b50600436106101825760003560e01c80638da5cb5b116100d8578063d088a2311161008c578063ef8032ef11610066578063ef8032ef146102d9578063f2fde38b146102ec578063fc7e286d146102ff57600080fd5b8063d088a231146102ab578063db1536f2146102b3578063dd900769146102c657600080fd5b8063b7b25b7c116100bd578063b7b25b7c1461027c578063b93f215714610285578063cdba73421461029857600080fd5b80638da5cb5b1461024e578063a14488191461026957600080fd5b8063715018a61161013a5780637d6babb4116101145780637d6babb41461022a57806380267b57146102325780638a1fcd601461024557600080fd5b8063715018a61461020657806372b0d90c1461020e5780637b407edb1461022157600080fd5b80633ccfd60b1161016b5780633ccfd60b146101af57806344fd9caa146101b757806359a515ba146101d357600080fd5b8063169e1ca8146101875780632b4656c81461019c575b600080fd5b61019a610195366004611f6e565b610349565b005b61019a6101aa366004611d18565b61047b565b61019a6105a6565b6101c060685481565b6040519081526020015b60405180910390f35b6101f66101e1366004611d00565b606d6020526000908152604090205460ff1681565b60405190151581526020016101ca565b61019a6105b1565b61019a61021c366004611c16565b610615565b6101c060695481565b6101c0610727565b61019a610240366004611d00565b610801565b6101c0606b5481565b6033546040516001600160a01b0390911681526020016101ca565b6101c0610277366004611eaa565b610860565b6101c0606a5481565b61019a610293366004611ec5565b6108ce565b61019a6102a6366004611f6e565b611063565b61019a6110e1565b6101c06102c1366004611eaa565b61114b565b61019a6102d4366004611d00565b611217565b61019a6102e7366004611d00565b611276565b61019a6102fa366004611c16565b6112d5565b61032e61030d366004611c16565b606c6020526000908152604090208054600182015460029092015490919083565b604080519384526020840192909252908201526060016101ca565b6001600160a01b0381166000908152606c602052604090206002810154156103b85760405162461bcd60e51b815260206004820152601e60248201527f43616e6e6f74206465706f736974207768696c6520756e6c6f636b696e67000060448201526064015b60405180910390fd5b828160000160008282546103cc9190612041565b90915550506065546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590526001600160a01b03909116906323b872dd906064015b602060405180830381600087803b15801561043d57600080fd5b505af1158015610451573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104759190611ce6565b50505050565b600054610100900460ff1680610494575060005460ff16155b6104f75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103af565b600054610100900460ff16158015610519576000805461ffff19166101011790555b6105216113b7565b606580546001600160a01b03808b167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255606680548a84169083161790556067805492891692909116919091179055606b85905560688490556069839055606a829055801561059c576000805461ff00191690555b5050505050505050565b6105af33610615565b565b6033546001600160a01b0316331461060b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103af565b6105af6000611479565b336000908152606c6020526040902060028101546106755760405162461bcd60e51b815260206004820152601560248201527f4465706f73697473206e6f7420756e6c6f636b6564000000000000000000000060448201526064016103af565b438160020154106106c85760405162461bcd60e51b815260206004820152601a60248201527f556e6c6f636b20706572696f64206e6f7420636f6d706c65746500000000000060448201526064016103af565b600181015481546000916106db91612041565b600080845560018401819055600284015560655460405163a9059cbb60e01b81526001600160a01b0386811660048301526024820184905292935091169063a9059cbb90604401610423565b336000908152606c602052604081208054151580610749575060008160010154115b6107955760405162461bcd60e51b815260206004820152601360248201527f4e6f7468696e6720746f2077697468647261770000000000000000000000000060448201526064016103af565b6002810154156107e75760405162461bcd60e51b815260206004820152601a60248201527f556e6c6f636b20616c726561647920696e2070726f677265737300000000000060448201526064016103af565b606b546107f49043612041565b6002909101819055919050565b6033546001600160a01b0316331461085b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103af565b606955565b600080606a5483604001516108759190612041565b90504381116108875750600092915050565b600083604001514361089991906120bd565b606a546108a691906120bd565b90506000606a54606854836108bb919061209e565b6108c5919061207e565b95945050505050565b60006108d98461114b565b90506108e7848285856114e3565b83516001600160a01b03166000908152606c602052604081209061090a86610860565b905080826001015483600001546109219190612041565b10156109955760405162461bcd60e51b815260206004820152602760248201527f53656e64657220646f65736e2774206861766520656e6f7567682066756e647360448201527f20746f207061790000000000000000000000000000000000000000000000000060648201526084016103af565b6000838152606d60209081526040808320805460ff191660011790556066549189015190517f084af0b20000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169063084af0b29060240160006040518083038186803b158015610a1157600080fd5b505afa158015610a25573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a4d9190810190611d85565b9050806060015115156001151514610acd5760405162461bcd60e51b815260206004820152602960248201527f5469636b6574207265636569766572206d757374206861766520612076616c6960448201527f64206c697374696e67000000000000000000000000000000000000000000000060648201526084016103af565b60675460208801516040517fdf349ed50000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600092919091169063df349ed59060240160206040518083038186803b158015610b3457600080fd5b505afa158015610b48573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b6c9190611f56565b905080610bbb5760405162461bcd60e51b815260206004820152601f60248201527f5469636b6574207265636569766572206d7573742068617665207374616b650060448201526064016103af565b83546068541115610cef576065546020890151855460405163a9059cbb60e01b81526001600160a01b039283166004820152602481019190915291169063a9059cbb90604401602060405180830381600087803b158015610c1b57600080fd5b505af1158015610c2f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c539190611ce6565b50606554600185015460405163a9059cbb60e01b81526001600160a01b039092166004830181905260248301919091529063a9059cbb90604401602060405180830381600087803b158015610ca757600080fd5b505af1158015610cbb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cdf9190611ce6565b506000808555600185015561059c565b8354610cfb9084611767565b84556020820151600090606490610d1690869060ff1661209e565b610d20919061207e565b60675460208b01516040517f791936100000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292935060009291169063791936109060240160006040518083038186803b158015610d8857600080fd5b505afa158015610d9c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610dc49190810190611c32565b90508160005b82518163ffffffff161015610fb05760675483516000916001600160a01b0316906382dda22d90869063ffffffff8616908110610e1757634e487b7160e01b600052603260045260246000fd5b60200260200101518f602001516040518363ffffffff1660e01b8152600401610e569291906001600160a01b0392831681529116602082015260400190565b604080518083038186803b158015610e6d57600080fd5b505afa158015610e81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ea59190611e55565b9050600086868360000151610eba919061209e565b610ec4919061207e565b9050610ed081856120bd565b60655486519195506001600160a01b03169063a9059cbb90879063ffffffff8716908110610f0e57634e487b7160e01b600052603260045260246000fd5b6020026020010151836040518363ffffffff1660e01b8152600401610f489291906001600160a01b03929092168252602082015260400190565b602060405180830381600087803b158015610f6257600080fd5b505af1158015610f76573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f9a9190611ce6565b5050508080610fa890612100565b915050610dca565b50600081610fbe85896120bd565b610fc89190612041565b60655460208e015160405163a9059cbb60e01b81526001600160a01b03918216600482015260248101849052929350169063a9059cbb90604401602060405180830381600087803b15801561101c57600080fd5b505af1158015611030573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110549190611ce6565b50505050505050505050505050565b6001600160a01b0381166000908152606c602052604090206002810154156110cd5760405162461bcd60e51b815260206004820152601e60248201527f43616e6e6f74206465706f736974207768696c6520756e6c6f636b696e67000060448201526064016103af565b828160010160008282546103cc9190612041565b336000908152606c6020526040902060028101546111415760405162461bcd60e51b815260206004820152601a60248201527f4e6f7420756e6c6f636b696e672c2063616e6e6f74206c6f636b00000000000060448201526064016103af565b6000600290910155565b80516020808301516068546069546040808701516060880151608089015192516000986111fa9890979695949101606097881b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000090811682529690971b9095166014870152602886019390935260488501919091526068840152608883015260e01b7fffffffff000000000000000000000000000000000000000000000000000000001660a882015260ac0190565b604051602081830303815290604052805190602001209050919050565b6033546001600160a01b031633146112715760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103af565b606b55565b6033546001600160a01b031633146112d05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103af565b606855565b6033546001600160a01b0316331461132f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103af565b6001600160a01b0381166113ab5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016103af565b6113b481611479565b50565b600054610100900460ff16806113d0575060005460ff16155b6114335760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103af565b600054610100900460ff16158015611455576000805461ffff19166101011790555b61145d61177a565b61146561182b565b80156113b4576000805461ff001916905550565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b83516001600160a01b031661153a5760405162461bcd60e51b815260206004820152601560248201527f5469636b65742073656e646572206973206e756c6c000000000000000000000060448201526064016103af565b60208401516001600160a01b03166115945760405162461bcd60e51b815260206004820152601760248201527f5469636b6574207265636569766572206973206e756c6c00000000000000000060448201526064016103af565b6000838152606d602052604090205460ff16156115f35760405162461bcd60e51b815260206004820152601760248201527f5469636b657420616c72656164792072656465656d656400000000000000000060448201526064016103af565b6060840151604080516020810185905201604051602081830303815290604052805190602001201461168d5760405162461bcd60e51b815260206004820152603360248201527f48617368206f6620726563656976657252616e6420646f65736e2774206d617460448201527f636820726563656976657252616e64486173680000000000000000000000000060648201526084016103af565b61169c818560000151856118d2565b61170e5760405162461bcd60e51b815260206004820152602560248201527f5469636b657420646f65736e2774206861766520612076616c6964207369676e60448201527f617475726500000000000000000000000000000000000000000000000000000060648201526084016103af565b61171b81836069546118fa565b6104755760405162461bcd60e51b815260206004820152601660248201527f5469636b6574206973206e6f7420612077696e6e65720000000000000000000060448201526064016103af565b600061177382846120bd565b9392505050565b600054610100900460ff1680611793575060005460ff16155b6117f65760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103af565b600054610100900460ff16158015611465576000805461ffff191661010117905580156113b4576000805461ff001916905550565b600054610100900460ff1680611844575060005460ff16155b6118a75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103af565b600054610100900460ff161580156118c9576000805461ffff19166101011790555b61146533611479565b6000826001600160a01b03166118e88386611931565b6001600160a01b031614949350505050565b6000818484604051602001611910929190611f9d565b60408051601f19818403018152919052805160209091012010949350505050565b600081516041146119845760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016103af565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115611a115760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016103af565b601b8160ff161015611a2b57611a28601b82612059565b90505b8060ff16601b14158015611a4357508060ff16601c14155b15611a9b5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b60648201526084016103af565b6040805160008082526020820180845289905260ff841692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015611aef573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611b525760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016103af565b9695505050505050565b80518015158114611b6c57600080fd5b919050565b600060a08284031215611b82578081fd5b60405160a0810181811067ffffffffffffffff82111715611ba557611ba561213a565b6040529050808235611bb681612150565b81526020830135611bc681612150565b806020830152506040830135604082015260608301356060820152608083013563ffffffff81168114611bf857600080fd5b6080919091015292915050565b805160ff81168114611b6c57600080fd5b600060208284031215611c27578081fd5b813561177381612150565b60006020808385031215611c44578182fd5b825167ffffffffffffffff80821115611c5b578384fd5b818501915085601f830112611c6e578384fd5b815181811115611c8057611c8061213a565b8060051b9150611c91848301611fe8565b8181528481019084860184860187018a1015611cab578788fd5b8795505b83861015611cd95780519450611cc485612150565b84835260019590950194918601918601611caf565b5098975050505050505050565b600060208284031215611cf7578081fd5b61177382611b5c565b600060208284031215611d11578081fd5b5035919050565b600080600080600080600060e0888a031215611d32578283fd5b8735611d3d81612150565b96506020880135611d4d81612150565b95506040880135611d5d81612150565b969995985095966060810135965060808101359560a0820135955060c0909101359350915050565b60006020808385031215611d97578182fd5b825167ffffffffffffffff80821115611dae578384fd5b9084019060808287031215611dc1578384fd5b611dc9611fbf565b825182811115611dd7578586fd5b83019150601f82018713611de9578485fd5b8151611dfc611df782612019565b611fe8565b8181528886838601011115611e0f578687fd5b611e1e828783018887016120d4565b825250611e2c838501611c05565b8482015260408301516040820152611e4660608401611b5c565b60608201529695505050505050565b600060408284031215611e66578081fd5b6040516040810181811067ffffffffffffffff82111715611e8957611e8961213a565b604052825181526020830151611e9e81612150565b60208201529392505050565b600060a08284031215611ebb578081fd5b6117738383611b71565b600080600060e08486031215611ed9578081fd5b611ee38585611b71565b925060a0840135915060c084013567ffffffffffffffff811115611f05578182fd5b8401601f81018613611f15578182fd5b8035611f23611df782612019565b818152876020838501011115611f37578384fd5b8160208401602083013783602083830101528093505050509250925092565b600060208284031215611f67578081fd5b5051919050565b60008060408385031215611f80578182fd5b823591506020830135611f9281612150565b809150509250929050565b60008351611faf8184602088016120d4565b9190910191825250602001919050565b6040516080810167ffffffffffffffff81118282101715611fe257611fe261213a565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156120115761201161213a565b604052919050565b600067ffffffffffffffff8211156120335761203361213a565b50601f01601f191660200190565b6000821982111561205457612054612124565b500190565b600060ff821660ff84168060ff0382111561207657612076612124565b019392505050565b60008261209957634e487b7160e01b81526012600452602481fd5b500490565b60008160001904831182151516156120b8576120b8612124565b500290565b6000828210156120cf576120cf612124565b500390565b60005b838110156120ef5781810151838201526020016120d7565b838111156104755750506000910152565b600063ffffffff8083168181141561211a5761211a612124565b6001019392505050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146113b457600080fdfea2646970667358221220b8546cf0f23146ee1436f97e9f985dbccae4818727aa2128d483ff88cdd7244d64736f6c63430008040033"

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

// CalculatePayout is a free data retrieval call binding the contract method 0xa1448819.
//
// Solidity: function calculatePayout((address,address,uint256,bytes32,uint32) ticket) view returns(uint256)
func (_SyloTicketing *SyloTicketingCaller) CalculatePayout(opts *bind.CallOpts, ticket SyloTicketingTicket) (*big.Int, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "calculatePayout", ticket)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePayout is a free data retrieval call binding the contract method 0xa1448819.
//
// Solidity: function calculatePayout((address,address,uint256,bytes32,uint32) ticket) view returns(uint256)
func (_SyloTicketing *SyloTicketingSession) CalculatePayout(ticket SyloTicketingTicket) (*big.Int, error) {
	return _SyloTicketing.Contract.CalculatePayout(&_SyloTicketing.CallOpts, ticket)
}

// CalculatePayout is a free data retrieval call binding the contract method 0xa1448819.
//
// Solidity: function calculatePayout((address,address,uint256,bytes32,uint32) ticket) view returns(uint256)
func (_SyloTicketing *SyloTicketingCallerSession) CalculatePayout(ticket SyloTicketingTicket) (*big.Int, error) {
	return _SyloTicketing.Contract.CalculatePayout(&_SyloTicketing.CallOpts, ticket)
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

// FaceValue is a free data retrieval call binding the contract method 0x44fd9caa.
//
// Solidity: function faceValue() view returns(uint256)
func (_SyloTicketing *SyloTicketingCaller) FaceValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "faceValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FaceValue is a free data retrieval call binding the contract method 0x44fd9caa.
//
// Solidity: function faceValue() view returns(uint256)
func (_SyloTicketing *SyloTicketingSession) FaceValue() (*big.Int, error) {
	return _SyloTicketing.Contract.FaceValue(&_SyloTicketing.CallOpts)
}

// FaceValue is a free data retrieval call binding the contract method 0x44fd9caa.
//
// Solidity: function faceValue() view returns(uint256)
func (_SyloTicketing *SyloTicketingCallerSession) FaceValue() (*big.Int, error) {
	return _SyloTicketing.Contract.FaceValue(&_SyloTicketing.CallOpts)
}

// GetTicketHash is a free data retrieval call binding the contract method 0xdb1536f2.
//
// Solidity: function getTicketHash((address,address,uint256,bytes32,uint32) ticket) view returns(bytes32)
func (_SyloTicketing *SyloTicketingCaller) GetTicketHash(opts *bind.CallOpts, ticket SyloTicketingTicket) ([32]byte, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "getTicketHash", ticket)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTicketHash is a free data retrieval call binding the contract method 0xdb1536f2.
//
// Solidity: function getTicketHash((address,address,uint256,bytes32,uint32) ticket) view returns(bytes32)
func (_SyloTicketing *SyloTicketingSession) GetTicketHash(ticket SyloTicketingTicket) ([32]byte, error) {
	return _SyloTicketing.Contract.GetTicketHash(&_SyloTicketing.CallOpts, ticket)
}

// GetTicketHash is a free data retrieval call binding the contract method 0xdb1536f2.
//
// Solidity: function getTicketHash((address,address,uint256,bytes32,uint32) ticket) view returns(bytes32)
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

// TicketLength is a free data retrieval call binding the contract method 0xb7b25b7c.
//
// Solidity: function ticketLength() view returns(uint256)
func (_SyloTicketing *SyloTicketingCaller) TicketLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "ticketLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TicketLength is a free data retrieval call binding the contract method 0xb7b25b7c.
//
// Solidity: function ticketLength() view returns(uint256)
func (_SyloTicketing *SyloTicketingSession) TicketLength() (*big.Int, error) {
	return _SyloTicketing.Contract.TicketLength(&_SyloTicketing.CallOpts)
}

// TicketLength is a free data retrieval call binding the contract method 0xb7b25b7c.
//
// Solidity: function ticketLength() view returns(uint256)
func (_SyloTicketing *SyloTicketingCallerSession) TicketLength() (*big.Int, error) {
	return _SyloTicketing.Contract.TicketLength(&_SyloTicketing.CallOpts)
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

// WinProb is a free data retrieval call binding the contract method 0x7b407edb.
//
// Solidity: function winProb() view returns(uint256)
func (_SyloTicketing *SyloTicketingCaller) WinProb(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "winProb")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WinProb is a free data retrieval call binding the contract method 0x7b407edb.
//
// Solidity: function winProb() view returns(uint256)
func (_SyloTicketing *SyloTicketingSession) WinProb() (*big.Int, error) {
	return _SyloTicketing.Contract.WinProb(&_SyloTicketing.CallOpts)
}

// WinProb is a free data retrieval call binding the contract method 0x7b407edb.
//
// Solidity: function winProb() view returns(uint256)
func (_SyloTicketing *SyloTicketingCallerSession) WinProb() (*big.Int, error) {
	return _SyloTicketing.Contract.WinProb(&_SyloTicketing.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0x2b4656c8.
//
// Solidity: function initialize(address token, address listings, address stakingManager, uint256 _unlockDuration, uint256 _faceValue, uint256 _winProb, uint256 _ticketLength) returns()
func (_SyloTicketing *SyloTicketingTransactor) Initialize(opts *bind.TransactOpts, token common.Address, listings common.Address, stakingManager common.Address, _unlockDuration *big.Int, _faceValue *big.Int, _winProb *big.Int, _ticketLength *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "initialize", token, listings, stakingManager, _unlockDuration, _faceValue, _winProb, _ticketLength)
}

// Initialize is a paid mutator transaction binding the contract method 0x2b4656c8.
//
// Solidity: function initialize(address token, address listings, address stakingManager, uint256 _unlockDuration, uint256 _faceValue, uint256 _winProb, uint256 _ticketLength) returns()
func (_SyloTicketing *SyloTicketingSession) Initialize(token common.Address, listings common.Address, stakingManager common.Address, _unlockDuration *big.Int, _faceValue *big.Int, _winProb *big.Int, _ticketLength *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.Initialize(&_SyloTicketing.TransactOpts, token, listings, stakingManager, _unlockDuration, _faceValue, _winProb, _ticketLength)
}

// Initialize is a paid mutator transaction binding the contract method 0x2b4656c8.
//
// Solidity: function initialize(address token, address listings, address stakingManager, uint256 _unlockDuration, uint256 _faceValue, uint256 _winProb, uint256 _ticketLength) returns()
func (_SyloTicketing *SyloTicketingTransactorSession) Initialize(token common.Address, listings common.Address, stakingManager common.Address, _unlockDuration *big.Int, _faceValue *big.Int, _winProb *big.Int, _ticketLength *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.Initialize(&_SyloTicketing.TransactOpts, token, listings, stakingManager, _unlockDuration, _faceValue, _winProb, _ticketLength)
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

// Redeem is a paid mutator transaction binding the contract method 0xb93f2157.
//
// Solidity: function redeem((address,address,uint256,bytes32,uint32) ticket, uint256 receiverRand, bytes sig) returns()
func (_SyloTicketing *SyloTicketingTransactor) Redeem(opts *bind.TransactOpts, ticket SyloTicketingTicket, receiverRand *big.Int, sig []byte) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "redeem", ticket, receiverRand, sig)
}

// Redeem is a paid mutator transaction binding the contract method 0xb93f2157.
//
// Solidity: function redeem((address,address,uint256,bytes32,uint32) ticket, uint256 receiverRand, bytes sig) returns()
func (_SyloTicketing *SyloTicketingSession) Redeem(ticket SyloTicketingTicket, receiverRand *big.Int, sig []byte) (*types.Transaction, error) {
	return _SyloTicketing.Contract.Redeem(&_SyloTicketing.TransactOpts, ticket, receiverRand, sig)
}

// Redeem is a paid mutator transaction binding the contract method 0xb93f2157.
//
// Solidity: function redeem((address,address,uint256,bytes32,uint32) ticket, uint256 receiverRand, bytes sig) returns()
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

// SetFaceValue is a paid mutator transaction binding the contract method 0xef8032ef.
//
// Solidity: function setFaceValue(uint256 _faceValue) returns()
func (_SyloTicketing *SyloTicketingTransactor) SetFaceValue(opts *bind.TransactOpts, _faceValue *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "setFaceValue", _faceValue)
}

// SetFaceValue is a paid mutator transaction binding the contract method 0xef8032ef.
//
// Solidity: function setFaceValue(uint256 _faceValue) returns()
func (_SyloTicketing *SyloTicketingSession) SetFaceValue(_faceValue *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.SetFaceValue(&_SyloTicketing.TransactOpts, _faceValue)
}

// SetFaceValue is a paid mutator transaction binding the contract method 0xef8032ef.
//
// Solidity: function setFaceValue(uint256 _faceValue) returns()
func (_SyloTicketing *SyloTicketingTransactorSession) SetFaceValue(_faceValue *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.SetFaceValue(&_SyloTicketing.TransactOpts, _faceValue)
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

// SetWinProb is a paid mutator transaction binding the contract method 0x80267b57.
//
// Solidity: function setWinProb(uint256 _winProb) returns()
func (_SyloTicketing *SyloTicketingTransactor) SetWinProb(opts *bind.TransactOpts, _winProb *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "setWinProb", _winProb)
}

// SetWinProb is a paid mutator transaction binding the contract method 0x80267b57.
//
// Solidity: function setWinProb(uint256 _winProb) returns()
func (_SyloTicketing *SyloTicketingSession) SetWinProb(_winProb *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.SetWinProb(&_SyloTicketing.TransactOpts, _winProb)
}

// SetWinProb is a paid mutator transaction binding the contract method 0x80267b57.
//
// Solidity: function setWinProb(uint256 _winProb) returns()
func (_SyloTicketing *SyloTicketingTransactorSession) SetWinProb(_winProb *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.SetWinProb(&_SyloTicketing.TransactOpts, _winProb)
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
