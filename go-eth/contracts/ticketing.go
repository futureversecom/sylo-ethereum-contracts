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
	Sender          common.Address
	Redeemer        common.Address
	GenerationBlock *big.Int
	SenderCommit    [32]byte
	RedeemerCommit  [32]byte
}

// SyloTicketingABI is the input ABI used to generate the binding from.
const SyloTicketingABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"baseLiveWinProb\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decayRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"escrow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expiredWinProb\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"faceValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ticketDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usedTickets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractListings\",\"name\":\"listings\",\"type\":\"address\"},{\"internalType\":\"contractStakingManager\",\"name\":\"stakingManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unlockDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_baseLiveWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_expiredWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"_decayRate\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_ticketDuration\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newUnlockDuration\",\"type\":\"uint256\"}],\"name\":\"setUnlockDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_faceValue\",\"type\":\"uint256\"}],\"name\":\"setFaceValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_baseLiveWinProb\",\"type\":\"uint128\"}],\"name\":\"setBaseLiveWinProb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_expiredWinProb\",\"type\":\"uint128\"}],\"name\":\"setExpiredWinProb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositEscrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"generationBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"senderCommit\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"redeemerCommit\",\"type\":\"bytes32\"}],\"internalType\":\"structSyloTicketing.Ticket\",\"name\":\"ticket\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"senderRand\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemerRand\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"generationBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"senderCommit\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"redeemerCommit\",\"type\":\"bytes32\"}],\"internalType\":\"structSyloTicketing.Ticket\",\"name\":\"ticket\",\"type\":\"tuple\"}],\"name\":\"calculateWinningProbability\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"generationBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"senderCommit\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"redeemerCommit\",\"type\":\"bytes32\"}],\"internalType\":\"structSyloTicketing.Ticket\",\"name\":\"ticket\",\"type\":\"tuple\"}],\"name\":\"getTicketHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SyloTicketingBin is the compiled bytecode used for deploying new contracts.
var SyloTicketingBin = "0x608060405234801561001057600080fd5b50612474806100206000396000f3fe608060405234801561001057600080fd5b50600436106101a35760003560e01c8063a8f19c14116100ee578063d2d1724311610097578063ef8032ef11610071578063ef8032ef14610362578063f2fde38b14610375578063fb06e74714610388578063fc7e286d1461039b57600080fd5b8063d2d1724314610329578063dd9007691461033c578063dedcebda1461034f57600080fd5b8063bcbee543116100c8578063bcbee543146102cf578063cdba73421461030e578063d088a2311461032157600080fd5b8063a8f19c141461028a578063a90a60271461029d578063a9c1f2f1146102b057600080fd5b8063715018a61161015057806387bcc0c51161012a57806387bcc0c51461025d5780638a1fcd60146102665780638da5cb5b1461026f57600080fd5b8063715018a61461023a57806372b0d90c146102425780637d6babb41461025557600080fd5b8063410838b011610181578063410838b0146101d857806344fd9caa146101eb57806359a515ba1461020757600080fd5b8063169e1ca8146101a857806317e9370a146101bd5780633ccfd60b146101d0575b600080fd5b6101bb6101b6366004612210565b6103e5565b005b6101bb6101cb366004611f69565b610517565b6101bb610675565b6101bb6101e6366004612145565b610680565b6101f460685481565b6040519081526020015b60405180910390f35b61022a610215366004611f51565b606e6020526000908152604090205460ff1681565b60405190151581526020016101fe565b6101bb610e05565b6101bb610250366004611e67565b610e69565b6101f4610f7b565b6101f4606a5481565b6101f4606b5481565b6033546040516001600160a01b0390911681526020016101fe565b6101bb6102983660046121de565b611055565b6101bb6102ab3660046121de565b6110db565b606c546102bd9060ff1681565b60405160ff90911681526020016101fe565b6069546102f69070010000000000000000000000000000000090046001600160801b031681565b6040516001600160801b0390911681526020016101fe565b6101bb61031c366004612210565b611160565b6101bb6111de565b6102f661033736600461212a565b611248565b6101bb61034a366004611f51565b6112d0565b6069546102f6906001600160801b031681565b6101bb610370366004611f51565b61132f565b6101bb610383366004611e67565b61138e565b6101f461039636600461212a565b611470565b6103ca6103a9366004611e67565b606d6020526000908152604090208054600182015460029092015490919083565b604080519384526020840192909252908201526060016101fe565b6001600160a01b0381166000908152606d602052604090206002810154156104545760405162461bcd60e51b815260206004820152601e60248201527f43616e6e6f74206465706f736974207768696c6520756e6c6f636b696e67000060448201526064015b60405180910390fd5b8281600001600082825461046891906122e3565b90915550506065546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590526001600160a01b03909116906323b872dd906064015b602060405180830381600087803b1580156104d957600080fd5b505af11580156104ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105119190611f37565b50505050565b600054610100900460ff1680610530575060005460ff16155b6105935760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161044b565b600054610100900460ff161580156105b5576000805461ffff19166101011790555b6105bd611536565b606580546001600160a01b03808d167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255606680548c841690831617905560678054928b1692909116919091179055606b87905560688690556001600160801b038481167001000000000000000000000000000000000290861617606955606c805460ff851660ff19909116179055606a8290558015610669576000805461ff00191690555b50505050505050505050565b61067e33610e69565b565b600061068b85611470565b905061069a85828686866115f8565b84516001600160a01b03166000908152606d60205260409020606854600182015482546106c791906122e3565b101561073b5760405162461bcd60e51b815260206004820152602760248201527f53656e64657220646f65736e2774206861766520656e6f7567682066756e647360448201527f20746f2070617900000000000000000000000000000000000000000000000000606482015260840161044b565b6000828152606e60209081526040808320805460ff191660011790556066549189015190517f084af0b20000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169063084af0b29060240160006040518083038186803b1580156107b757600080fd5b505afa1580156107cb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526107f39190810190612005565b90508060600151151560011515146108735760405162461bcd60e51b815260206004820152602960248201527f5469636b65742072656465656d6572206d757374206861766520612076616c6960448201527f64206c697374696e670000000000000000000000000000000000000000000000606482015260840161044b565b60675460208801516040517fdf349ed50000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600092919091169063df349ed59060240160206040518083038186803b1580156108da57600080fd5b505afa1580156108ee573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061091291906121f8565b9050806109615760405162461bcd60e51b815260206004820152601f60248201527f5469636b65742072656465656d6572206d7573742068617665207374616b6500604482015260640161044b565b82546068541115610a95576065546020890151845460405163a9059cbb60e01b81526001600160a01b039283166004820152602481019190915291169063a9059cbb90604401602060405180830381600087803b1580156109c157600080fd5b505af11580156109d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109f99190611f37565b50606554600184015460405163a9059cbb60e01b81526001600160a01b039092166004830181905260248301919091529063a9059cbb90604401602060405180830381600087803b158015610a4d57600080fd5b505af1158015610a61573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a859190611f37565b5060008084556001840155610dfb565b6068548354610aa49190612387565b83556068546020830151600091610abd9160ff16611929565b60675460208b01516040517f791936100000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292935060009291169063791936109060240160006040518083038186803b158015610b2557600080fd5b505afa158015610b39573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610b619190810190611e83565b90508160005b82518163ffffffff161015610d4d5760675483516000916001600160a01b0316906382dda22d90869063ffffffff8616908110610bb457634e487b7160e01b600052603260045260246000fd5b60200260200101518f602001516040518363ffffffff1660e01b8152600401610bf39291906001600160a01b0392831681529116602082015260400190565b604080518083038186803b158015610c0a57600080fd5b505afa158015610c1e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c4291906120d5565b9050600086868360000151610c579190612340565b610c619190612320565b9050610c6d8185612387565b60655486519195506001600160a01b03169063a9059cbb90879063ffffffff8716908110610cab57634e487b7160e01b600052603260045260246000fd5b6020026020010151836040518363ffffffff1660e01b8152600401610ce59291906001600160a01b03929092168252602082015260400190565b602060405180830381600087803b158015610cff57600080fd5b505af1158015610d13573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d379190611f37565b5050508080610d45906123ca565b915050610b67565b5060008184606854610d5f9190612387565b610d6991906122e3565b60655460208e015160405163a9059cbb60e01b81526001600160a01b03918216600482015260248101849052929350169063a9059cbb90604401602060405180830381600087803b158015610dbd57600080fd5b505af1158015610dd1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610df59190611f37565b50505050505b5050505050505050565b6033546001600160a01b03163314610e5f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161044b565b61067e6000611951565b336000908152606d602052604090206002810154610ec95760405162461bcd60e51b815260206004820152601560248201527f4465706f73697473206e6f7420756e6c6f636b65640000000000000000000000604482015260640161044b565b43816002015410610f1c5760405162461bcd60e51b815260206004820152601a60248201527f556e6c6f636b20706572696f64206e6f7420636f6d706c657465000000000000604482015260640161044b565b60018101548154600091610f2f916122e3565b600080845560018401819055600284015560655460405163a9059cbb60e01b81526001600160a01b0386811660048301526024820184905292935091169063a9059cbb906044016104bf565b336000908152606d602052604081208054151580610f9d575060008160010154115b610fe95760405162461bcd60e51b815260206004820152601360248201527f4e6f7468696e6720746f20776974686472617700000000000000000000000000604482015260640161044b565b60028101541561103b5760405162461bcd60e51b815260206004820152601a60248201527f556e6c6f636b20616c726561647920696e2070726f6772657373000000000000604482015260640161044b565b606b5461104890436122e3565b6002909101819055919050565b6033546001600160a01b031633146110af5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161044b565b606980546001600160801b03928316700100000000000000000000000000000000029216919091179055565b6033546001600160a01b031633146111355760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161044b565b606980546fffffffffffffffffffffffffffffffff19166001600160801b0392909216919091179055565b6001600160a01b0381166000908152606d602052604090206002810154156111ca5760405162461bcd60e51b815260206004820152601e60248201527f43616e6e6f74206465706f736974207768696c6520756e6c6f636b696e670000604482015260640161044b565b8281600101600082825461046891906122e3565b336000908152606d60205260409020600281015461123e5760405162461bcd60e51b815260206004820152601a60248201527f4e6f7420756e6c6f636b696e672c2063616e6e6f74206c6f636b000000000000604482015260640161044b565b6000600290910155565b60008082604001514361125b9190612387565b9050606a54811061126f5750600092915050565b606954606c54600091611290916001600160801b039091169060ff16611929565b90506000606a5483836112a39190612340565b6112ad9190612320565b6069549091506112c79082906001600160801b031661235f565b95945050505050565b6033546001600160a01b0316331461132a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161044b565b606b55565b6033546001600160a01b031633146113895760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161044b565b606855565b6033546001600160a01b031633146113e85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161044b565b6001600160a01b0381166114645760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161044b565b61146d81611951565b50565b805160208083015160685460695460408087015160608801516080890151925160009861151998909796956001600160801b0316949101606097881b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000090811682529690971b9095166014870152602886019390935260809190911b6fffffffffffffffffffffffffffffffff1916604885015260588401526078830152609882015260b80190565b604051602081830303815290604052805190602001209050919050565b600054610100900460ff168061154f575060005460ff16155b6115b25760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161044b565b600054610100900460ff161580156115d4576000805461ffff19166101011790555b6115dc6119bb565b6115e4611a6c565b801561146d576000805461ff001916905550565b84516001600160a01b031661164f5760405162461bcd60e51b815260206004820152601560248201527f5469636b65742073656e646572206973206e756c6c0000000000000000000000604482015260640161044b565b60208501516001600160a01b03166116a95760405162461bcd60e51b815260206004820152601760248201527f5469636b65742072656465656d6572206973206e756c6c000000000000000000604482015260640161044b565b6000848152606e602052604090205460ff16156117085760405162461bcd60e51b815260206004820152601760248201527f5469636b657420616c72656164792072656465656d6564000000000000000000604482015260640161044b565b606085015160408051602081018690520160405160208183030381529060405280519060200120146117a25760405162461bcd60e51b815260206004820152602f60248201527f48617368206f662073656e64657252616e6420646f65736e2774206d6174636860448201527f2073656e64657252616e64486173680000000000000000000000000000000000606482015260840161044b565b6080850151604080516020810185905201604051602081830303815290604052805190602001201461183c5760405162461bcd60e51b815260206004820152603360248201527f48617368206f662072656465656d657252616e6420646f65736e2774206d617460448201527f63682072656465656d657252616e644861736800000000000000000000000000606482015260840161044b565b61184b81866000015186611b13565b6118bd5760405162461bcd60e51b815260206004820152602560248201527f5469636b657420646f65736e2774206861766520612076616c6964207369676e60448201527f6174757265000000000000000000000000000000000000000000000000000000606482015260840161044b565b60006118c886611248565b90506118d5828483611b3b565b6119215760405162461bcd60e51b815260206004820152601660248201527f5469636b6574206973206e6f7420612077696e6e657200000000000000000000604482015260640161044b565b505050505050565b60006064611940836001600160801b038616612340565b61194a9190612320565b9392505050565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16806119d4575060005460ff16155b611a375760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161044b565b600054610100900460ff161580156115e4576000805461ffff1916610101179055801561146d576000805461ff001916905550565b600054610100900460ff1680611a85575060005460ff16155b611ae85760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161044b565b600054610100900460ff16158015611b0a576000805461ffff19166101011790555b6115e433611951565b6000826001600160a01b0316611b298386611b84565b6001600160a01b031614949350505050565b6000806080836001600160801b0316901b9050808585604051602001611b6292919061223f565b60408051601f1981840301815291905280516020909101201095945050505050565b60008151604114611bd75760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161044b565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115611c645760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161044b565b601b8160ff161015611c7e57611c7b601b826122fb565b90505b8060ff16601b14158015611c9657508060ff16601c14155b15611cee5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b606482015260840161044b565b6040805160008082526020820180845289905260ff841692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015611d42573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611da55760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161044b565b9695505050505050565b80518015158114611dbf57600080fd5b919050565b600060a08284031215611dd5578081fd5b60405160a0810181811067ffffffffffffffff82111715611df857611df8612404565b6040529050808235611e098161241a565b81526020830135611e198161241a565b806020830152506040830135604082015260608301356060820152608083013560808201525092915050565b80356001600160801b0381168114611dbf57600080fd5b8051611dbf8161242f565b600060208284031215611e78578081fd5b813561194a8161241a565b60006020808385031215611e95578182fd5b825167ffffffffffffffff80821115611eac578384fd5b818501915085601f830112611ebf578384fd5b815181811115611ed157611ed1612404565b8060051b9150611ee284830161228a565b8181528481019084860184860187018a1015611efc578788fd5b8795505b83861015611f2a5780519450611f158561241a565b84835260019590950194918601918601611f00565b5098975050505050505050565b600060208284031215611f48578081fd5b61194a82611daf565b600060208284031215611f62578081fd5b5035919050565b60008060008060008060008060006101208a8c031215611f87578485fd5b8935611f928161241a565b985060208a0135611fa28161241a565b975060408a0135611fb28161241a565b965060608a0135955060808a01359450611fce60a08b01611e45565b9350611fdc60c08b01611e45565b925060e08a0135611fec8161242f565b809250506101008a013590509295985092959850929598565b60006020808385031215612017578182fd5b825167ffffffffffffffff8082111561202e578384fd5b9084019060808287031215612041578384fd5b612049612261565b825182811115612057578586fd5b83019150601f82018713612069578485fd5b815161207c612077826122bb565b61228a565b818152888683860101111561208f578687fd5b61209e8287830188870161239e565b8252506120ac838501611e5c565b84820152604083015160408201526120c660608401611daf565b60608201529695505050505050565b6000604082840312156120e6578081fd5b6040516040810181811067ffffffffffffffff8211171561210957612109612404565b60405282518152602083015161211e8161241a565b60208201529392505050565b600060a0828403121561213b578081fd5b61194a8383611dc4565b600080600080610100858703121561215b578182fd5b6121658686611dc4565b935060a0850135925060c0850135915060e085013567ffffffffffffffff81111561218e578182fd5b8501601f8101871361219e578182fd5b80356121ac612077826122bb565b8181528860208385010111156121c0578384fd5b81602084016020830137908101602001929092525092959194509250565b6000602082840312156121ef578081fd5b61194a82611e45565b600060208284031215612209578081fd5b5051919050565b60008060408385031215612222578182fd5b8235915060208301356122348161241a565b809150509250929050565b6000835161225181846020880161239e565b9190910191825250602001919050565b6040516080810167ffffffffffffffff8111828210171561228457612284612404565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156122b3576122b3612404565b604052919050565b600067ffffffffffffffff8211156122d5576122d5612404565b50601f01601f191660200190565b600082198211156122f6576122f66123ee565b500190565b600060ff821660ff84168060ff03821115612318576123186123ee565b019392505050565b60008261233b57634e487b7160e01b81526012600452602481fd5b500490565b600081600019048311821515161561235a5761235a6123ee565b500290565b60006001600160801b038381169083168181101561237f5761237f6123ee565b039392505050565b600082821015612399576123996123ee565b500390565b60005b838110156123b95781810151838201526020016123a1565b838111156105115750506000910152565b600063ffffffff808316818114156123e4576123e46123ee565b6001019392505050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461146d57600080fd5b60ff8116811461146d57600080fdfea26469706673582212203df0d0ef8b918919583f4194a6b1717d55cb6ea92617ef6826735c97b46241d664736f6c63430008040033"

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

// BaseLiveWinProb is a free data retrieval call binding the contract method 0xdedcebda.
//
// Solidity: function baseLiveWinProb() view returns(uint128)
func (_SyloTicketing *SyloTicketingCaller) BaseLiveWinProb(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "baseLiveWinProb")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseLiveWinProb is a free data retrieval call binding the contract method 0xdedcebda.
//
// Solidity: function baseLiveWinProb() view returns(uint128)
func (_SyloTicketing *SyloTicketingSession) BaseLiveWinProb() (*big.Int, error) {
	return _SyloTicketing.Contract.BaseLiveWinProb(&_SyloTicketing.CallOpts)
}

// BaseLiveWinProb is a free data retrieval call binding the contract method 0xdedcebda.
//
// Solidity: function baseLiveWinProb() view returns(uint128)
func (_SyloTicketing *SyloTicketingCallerSession) BaseLiveWinProb() (*big.Int, error) {
	return _SyloTicketing.Contract.BaseLiveWinProb(&_SyloTicketing.CallOpts)
}

// CalculateWinningProbability is a free data retrieval call binding the contract method 0xd2d17243.
//
// Solidity: function calculateWinningProbability((address,address,uint256,bytes32,bytes32) ticket) view returns(uint128)
func (_SyloTicketing *SyloTicketingCaller) CalculateWinningProbability(opts *bind.CallOpts, ticket SyloTicketingTicket) (*big.Int, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "calculateWinningProbability", ticket)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateWinningProbability is a free data retrieval call binding the contract method 0xd2d17243.
//
// Solidity: function calculateWinningProbability((address,address,uint256,bytes32,bytes32) ticket) view returns(uint128)
func (_SyloTicketing *SyloTicketingSession) CalculateWinningProbability(ticket SyloTicketingTicket) (*big.Int, error) {
	return _SyloTicketing.Contract.CalculateWinningProbability(&_SyloTicketing.CallOpts, ticket)
}

// CalculateWinningProbability is a free data retrieval call binding the contract method 0xd2d17243.
//
// Solidity: function calculateWinningProbability((address,address,uint256,bytes32,bytes32) ticket) view returns(uint128)
func (_SyloTicketing *SyloTicketingCallerSession) CalculateWinningProbability(ticket SyloTicketingTicket) (*big.Int, error) {
	return _SyloTicketing.Contract.CalculateWinningProbability(&_SyloTicketing.CallOpts, ticket)
}

// DecayRate is a free data retrieval call binding the contract method 0xa9c1f2f1.
//
// Solidity: function decayRate() view returns(uint8)
func (_SyloTicketing *SyloTicketingCaller) DecayRate(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "decayRate")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DecayRate is a free data retrieval call binding the contract method 0xa9c1f2f1.
//
// Solidity: function decayRate() view returns(uint8)
func (_SyloTicketing *SyloTicketingSession) DecayRate() (uint8, error) {
	return _SyloTicketing.Contract.DecayRate(&_SyloTicketing.CallOpts)
}

// DecayRate is a free data retrieval call binding the contract method 0xa9c1f2f1.
//
// Solidity: function decayRate() view returns(uint8)
func (_SyloTicketing *SyloTicketingCallerSession) DecayRate() (uint8, error) {
	return _SyloTicketing.Contract.DecayRate(&_SyloTicketing.CallOpts)
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

// ExpiredWinProb is a free data retrieval call binding the contract method 0xbcbee543.
//
// Solidity: function expiredWinProb() view returns(uint128)
func (_SyloTicketing *SyloTicketingCaller) ExpiredWinProb(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "expiredWinProb")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpiredWinProb is a free data retrieval call binding the contract method 0xbcbee543.
//
// Solidity: function expiredWinProb() view returns(uint128)
func (_SyloTicketing *SyloTicketingSession) ExpiredWinProb() (*big.Int, error) {
	return _SyloTicketing.Contract.ExpiredWinProb(&_SyloTicketing.CallOpts)
}

// ExpiredWinProb is a free data retrieval call binding the contract method 0xbcbee543.
//
// Solidity: function expiredWinProb() view returns(uint128)
func (_SyloTicketing *SyloTicketingCallerSession) ExpiredWinProb() (*big.Int, error) {
	return _SyloTicketing.Contract.ExpiredWinProb(&_SyloTicketing.CallOpts)
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

// GetTicketHash is a free data retrieval call binding the contract method 0xfb06e747.
//
// Solidity: function getTicketHash((address,address,uint256,bytes32,bytes32) ticket) view returns(bytes32)
func (_SyloTicketing *SyloTicketingCaller) GetTicketHash(opts *bind.CallOpts, ticket SyloTicketingTicket) ([32]byte, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "getTicketHash", ticket)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTicketHash is a free data retrieval call binding the contract method 0xfb06e747.
//
// Solidity: function getTicketHash((address,address,uint256,bytes32,bytes32) ticket) view returns(bytes32)
func (_SyloTicketing *SyloTicketingSession) GetTicketHash(ticket SyloTicketingTicket) ([32]byte, error) {
	return _SyloTicketing.Contract.GetTicketHash(&_SyloTicketing.CallOpts, ticket)
}

// GetTicketHash is a free data retrieval call binding the contract method 0xfb06e747.
//
// Solidity: function getTicketHash((address,address,uint256,bytes32,bytes32) ticket) view returns(bytes32)
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

// TicketDuration is a free data retrieval call binding the contract method 0x87bcc0c5.
//
// Solidity: function ticketDuration() view returns(uint256)
func (_SyloTicketing *SyloTicketingCaller) TicketDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "ticketDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TicketDuration is a free data retrieval call binding the contract method 0x87bcc0c5.
//
// Solidity: function ticketDuration() view returns(uint256)
func (_SyloTicketing *SyloTicketingSession) TicketDuration() (*big.Int, error) {
	return _SyloTicketing.Contract.TicketDuration(&_SyloTicketing.CallOpts)
}

// TicketDuration is a free data retrieval call binding the contract method 0x87bcc0c5.
//
// Solidity: function ticketDuration() view returns(uint256)
func (_SyloTicketing *SyloTicketingCallerSession) TicketDuration() (*big.Int, error) {
	return _SyloTicketing.Contract.TicketDuration(&_SyloTicketing.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0x17e9370a.
//
// Solidity: function initialize(address token, address listings, address stakingManager, uint256 _unlockDuration, uint256 _faceValue, uint128 _baseLiveWinProb, uint128 _expiredWinProb, uint8 _decayRate, uint256 _ticketDuration) returns()
func (_SyloTicketing *SyloTicketingTransactor) Initialize(opts *bind.TransactOpts, token common.Address, listings common.Address, stakingManager common.Address, _unlockDuration *big.Int, _faceValue *big.Int, _baseLiveWinProb *big.Int, _expiredWinProb *big.Int, _decayRate uint8, _ticketDuration *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "initialize", token, listings, stakingManager, _unlockDuration, _faceValue, _baseLiveWinProb, _expiredWinProb, _decayRate, _ticketDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0x17e9370a.
//
// Solidity: function initialize(address token, address listings, address stakingManager, uint256 _unlockDuration, uint256 _faceValue, uint128 _baseLiveWinProb, uint128 _expiredWinProb, uint8 _decayRate, uint256 _ticketDuration) returns()
func (_SyloTicketing *SyloTicketingSession) Initialize(token common.Address, listings common.Address, stakingManager common.Address, _unlockDuration *big.Int, _faceValue *big.Int, _baseLiveWinProb *big.Int, _expiredWinProb *big.Int, _decayRate uint8, _ticketDuration *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.Initialize(&_SyloTicketing.TransactOpts, token, listings, stakingManager, _unlockDuration, _faceValue, _baseLiveWinProb, _expiredWinProb, _decayRate, _ticketDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0x17e9370a.
//
// Solidity: function initialize(address token, address listings, address stakingManager, uint256 _unlockDuration, uint256 _faceValue, uint128 _baseLiveWinProb, uint128 _expiredWinProb, uint8 _decayRate, uint256 _ticketDuration) returns()
func (_SyloTicketing *SyloTicketingTransactorSession) Initialize(token common.Address, listings common.Address, stakingManager common.Address, _unlockDuration *big.Int, _faceValue *big.Int, _baseLiveWinProb *big.Int, _expiredWinProb *big.Int, _decayRate uint8, _ticketDuration *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.Initialize(&_SyloTicketing.TransactOpts, token, listings, stakingManager, _unlockDuration, _faceValue, _baseLiveWinProb, _expiredWinProb, _decayRate, _ticketDuration)
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

// Redeem is a paid mutator transaction binding the contract method 0x410838b0.
//
// Solidity: function redeem((address,address,uint256,bytes32,bytes32) ticket, uint256 senderRand, uint256 redeemerRand, bytes sig) returns()
func (_SyloTicketing *SyloTicketingTransactor) Redeem(opts *bind.TransactOpts, ticket SyloTicketingTicket, senderRand *big.Int, redeemerRand *big.Int, sig []byte) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "redeem", ticket, senderRand, redeemerRand, sig)
}

// Redeem is a paid mutator transaction binding the contract method 0x410838b0.
//
// Solidity: function redeem((address,address,uint256,bytes32,bytes32) ticket, uint256 senderRand, uint256 redeemerRand, bytes sig) returns()
func (_SyloTicketing *SyloTicketingSession) Redeem(ticket SyloTicketingTicket, senderRand *big.Int, redeemerRand *big.Int, sig []byte) (*types.Transaction, error) {
	return _SyloTicketing.Contract.Redeem(&_SyloTicketing.TransactOpts, ticket, senderRand, redeemerRand, sig)
}

// Redeem is a paid mutator transaction binding the contract method 0x410838b0.
//
// Solidity: function redeem((address,address,uint256,bytes32,bytes32) ticket, uint256 senderRand, uint256 redeemerRand, bytes sig) returns()
func (_SyloTicketing *SyloTicketingTransactorSession) Redeem(ticket SyloTicketingTicket, senderRand *big.Int, redeemerRand *big.Int, sig []byte) (*types.Transaction, error) {
	return _SyloTicketing.Contract.Redeem(&_SyloTicketing.TransactOpts, ticket, senderRand, redeemerRand, sig)
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

// SetBaseLiveWinProb is a paid mutator transaction binding the contract method 0xa90a6027.
//
// Solidity: function setBaseLiveWinProb(uint128 _baseLiveWinProb) returns()
func (_SyloTicketing *SyloTicketingTransactor) SetBaseLiveWinProb(opts *bind.TransactOpts, _baseLiveWinProb *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "setBaseLiveWinProb", _baseLiveWinProb)
}

// SetBaseLiveWinProb is a paid mutator transaction binding the contract method 0xa90a6027.
//
// Solidity: function setBaseLiveWinProb(uint128 _baseLiveWinProb) returns()
func (_SyloTicketing *SyloTicketingSession) SetBaseLiveWinProb(_baseLiveWinProb *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.SetBaseLiveWinProb(&_SyloTicketing.TransactOpts, _baseLiveWinProb)
}

// SetBaseLiveWinProb is a paid mutator transaction binding the contract method 0xa90a6027.
//
// Solidity: function setBaseLiveWinProb(uint128 _baseLiveWinProb) returns()
func (_SyloTicketing *SyloTicketingTransactorSession) SetBaseLiveWinProb(_baseLiveWinProb *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.SetBaseLiveWinProb(&_SyloTicketing.TransactOpts, _baseLiveWinProb)
}

// SetExpiredWinProb is a paid mutator transaction binding the contract method 0xa8f19c14.
//
// Solidity: function setExpiredWinProb(uint128 _expiredWinProb) returns()
func (_SyloTicketing *SyloTicketingTransactor) SetExpiredWinProb(opts *bind.TransactOpts, _expiredWinProb *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "setExpiredWinProb", _expiredWinProb)
}

// SetExpiredWinProb is a paid mutator transaction binding the contract method 0xa8f19c14.
//
// Solidity: function setExpiredWinProb(uint128 _expiredWinProb) returns()
func (_SyloTicketing *SyloTicketingSession) SetExpiredWinProb(_expiredWinProb *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.SetExpiredWinProb(&_SyloTicketing.TransactOpts, _expiredWinProb)
}

// SetExpiredWinProb is a paid mutator transaction binding the contract method 0xa8f19c14.
//
// Solidity: function setExpiredWinProb(uint128 _expiredWinProb) returns()
func (_SyloTicketing *SyloTicketingTransactorSession) SetExpiredWinProb(_expiredWinProb *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.SetExpiredWinProb(&_SyloTicketing.TransactOpts, _expiredWinProb)
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
