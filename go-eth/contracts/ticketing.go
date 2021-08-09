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
const SyloTicketingABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"baseLiveWinProb\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decayRate\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"escrow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expiredWinProb\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"faceValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ticketDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usedTickets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractListings\",\"name\":\"listings\",\"type\":\"address\"},{\"internalType\":\"contractStakingManager\",\"name\":\"stakingManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unlockDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_baseLiveWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_expiredWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint16\",\"name\":\"_decayRate\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_ticketDuration\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newUnlockDuration\",\"type\":\"uint256\"}],\"name\":\"setUnlockDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_faceValue\",\"type\":\"uint256\"}],\"name\":\"setFaceValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_baseLiveWinProb\",\"type\":\"uint128\"}],\"name\":\"setBaseLiveWinProb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_expiredWinProb\",\"type\":\"uint128\"}],\"name\":\"setExpiredWinProb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ticketDuration\",\"type\":\"uint256\"}],\"name\":\"setTicketDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositEscrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"generationBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"senderCommit\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"redeemerCommit\",\"type\":\"bytes32\"}],\"internalType\":\"structSyloTicketing.Ticket\",\"name\":\"ticket\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"senderRand\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemerRand\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"generationBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"senderCommit\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"redeemerCommit\",\"type\":\"bytes32\"}],\"internalType\":\"structSyloTicketing.Ticket\",\"name\":\"ticket\",\"type\":\"tuple\"}],\"name\":\"calculateWinningProbability\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"generationBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"senderCommit\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"redeemerCommit\",\"type\":\"bytes32\"}],\"internalType\":\"structSyloTicketing.Ticket\",\"name\":\"ticket\",\"type\":\"tuple\"}],\"name\":\"getTicketHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SyloTicketingBin is the compiled bytecode used for deploying new contracts.
var SyloTicketingBin = "0x608060405234801561001057600080fd5b5061254d806100206000396000f3fe608060405234801561001057600080fd5b50600436106101ae5760003560e01c8063a8f19c14116100ee578063d2d1724311610097578063ef8032ef11610071578063ef8032ef14610382578063f2fde38b14610395578063fb06e747146103a8578063fc7e286d146103bb57600080fd5b8063d2d1724314610349578063dd9007691461035c578063dedcebda1461036f57600080fd5b8063bcbee543116100c8578063bcbee543146102ef578063cdba73421461032e578063d088a2311461034157600080fd5b8063a8f19c14146102a8578063a90a6027146102bb578063a9c1f2f1146102ce57600080fd5b806372b0d90c1161015b5780638a1fcd60116101355780638a1fcd601461025e5780638da5cb5b146102675780639d01393f146102825780639e9ceeca1461029557600080fd5b806372b0d90c1461023a5780637d6babb41461024d57806387bcc0c51461025557600080fd5b806344fd9caa1161018c57806344fd9caa146101e357806359a515ba146101ff578063715018a61461023257600080fd5b8063169e1ca8146101b35780633ccfd60b146101c8578063410838b0146101d0575b600080fd5b6101c66101c13660046122e8565b610405565b005b6101c6610537565b6101c66101de36600461221d565b610542565b6101ec60685481565b6040519081526020015b60405180910390f35b61022261020d366004612029565b606e6020526000908152604090205460ff1681565b60405190151581526020016101f6565b6101c6610cc4565b6101c6610248366004611f3f565b610d28565b6101ec610e3a565b6101ec606a5481565b6101ec606b5481565b6033546040516001600160a01b0390911681526020016101f6565b6101c6610290366004612041565b610f14565b6101c66102a3366004612029565b611078565b6101c66102b63660046122b6565b611127565b6101c66102c93660046122b6565b6111ad565b606c546102dc9061ffff1681565b60405161ffff90911681526020016101f6565b6069546103169070010000000000000000000000000000000090046001600160801b031681565b6040516001600160801b0390911681526020016101f6565b6101c661033c3660046122e8565b611232565b6101c66112b0565b610316610357366004612202565b61131a565b6101c661036a366004612029565b6113a3565b606954610316906001600160801b031681565b6101c6610390366004612029565b611402565b6101c66103a3366004611f3f565b611461565b6101ec6103b6366004612202565b611543565b6103ea6103c9366004611f3f565b606d6020526000908152604090208054600182015460029092015490919083565b604080519384526020840192909252908201526060016101f6565b6001600160a01b0381166000908152606d602052604090206002810154156104745760405162461bcd60e51b815260206004820152601e60248201527f43616e6e6f74206465706f736974207768696c6520756e6c6f636b696e67000060448201526064015b60405180910390fd5b8281600001600082825461048891906123bb565b90915550506065546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590526001600160a01b03909116906323b872dd906064015b602060405180830381600087803b1580156104f957600080fd5b505af115801561050d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610531919061200f565b50505050565b61054033610d28565b565b600061054d85611543565b905061055c8582868686611609565b84516001600160a01b03166000908152606d602052604090206068546001820154825461058991906123bb565b10156105fd5760405162461bcd60e51b815260206004820152602760248201527f53656e64657220646f65736e2774206861766520656e6f7567682066756e647360448201527f20746f2070617900000000000000000000000000000000000000000000000000606482015260840161046b565b6000828152606e60209081526040808320805460ff191660011790556066549189015190517f084af0b20000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169063084af0b29060240160006040518083038186803b15801561067957600080fd5b505afa15801561068d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526106b591908101906120dd565b90508060600151151560011515146107355760405162461bcd60e51b815260206004820152602960248201527f5469636b65742072656465656d6572206d757374206861766520612076616c6960448201527f64206c697374696e670000000000000000000000000000000000000000000000606482015260840161046b565b60675460208801516040517fdf349ed50000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600092919091169063df349ed59060240160206040518083038186803b15801561079c57600080fd5b505afa1580156107b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107d491906122d0565b9050806108235760405162461bcd60e51b815260206004820152601f60248201527f5469636b65742072656465656d6572206d7573742068617665207374616b6500604482015260640161046b565b82546068541115610957576065546020890151845460405163a9059cbb60e01b81526001600160a01b039283166004820152602481019190915291169063a9059cbb90604401602060405180830381600087803b15801561088357600080fd5b505af1158015610897573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108bb919061200f565b50606554600184015460405163a9059cbb60e01b81526001600160a01b039092166004830181905260248301919091529063a9059cbb90604401602060405180830381600087803b15801561090f57600080fd5b505af1158015610923573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610947919061200f565b5060008084556001840155610cba565b6068548354610966919061245f565b8355606854602083015160009161097c9161193a565b60675460208b01516040517f791936100000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292935060009291169063791936109060240160006040518083038186803b1580156109e457600080fd5b505afa1580156109f8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a209190810190611f5b565b90508160005b82518163ffffffff161015610c0c5760675483516000916001600160a01b0316906382dda22d90869063ffffffff8616908110610a7357634e487b7160e01b600052603260045260246000fd5b60200260200101518f602001516040518363ffffffff1660e01b8152600401610ab29291906001600160a01b0392831681529116602082015260400190565b604080518083038186803b158015610ac957600080fd5b505afa158015610add573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b0191906121ad565b9050600086868360000151610b169190612418565b610b2091906123f8565b9050610b2c818561245f565b60655486519195506001600160a01b03169063a9059cbb90879063ffffffff8716908110610b6a57634e487b7160e01b600052603260045260246000fd5b6020026020010151836040518363ffffffff1660e01b8152600401610ba49291906001600160a01b03929092168252602082015260400190565b602060405180830381600087803b158015610bbe57600080fd5b505af1158015610bd2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bf6919061200f565b5050508080610c04906124a2565b915050610a26565b5060008184606854610c1e919061245f565b610c2891906123bb565b60655460208e015160405163a9059cbb60e01b81526001600160a01b03918216600482015260248101849052929350169063a9059cbb90604401602060405180830381600087803b158015610c7c57600080fd5b505af1158015610c90573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cb4919061200f565b50505050505b5050505050505050565b6033546001600160a01b03163314610d1e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161046b565b6105406000611967565b336000908152606d602052604090206002810154610d885760405162461bcd60e51b815260206004820152601560248201527f4465706f73697473206e6f7420756e6c6f636b65640000000000000000000000604482015260640161046b565b43816002015410610ddb5760405162461bcd60e51b815260206004820152601a60248201527f556e6c6f636b20706572696f64206e6f7420636f6d706c657465000000000000604482015260640161046b565b60018101548154600091610dee916123bb565b600080845560018401819055600284015560655460405163a9059cbb60e01b81526001600160a01b0386811660048301526024820184905292935091169063a9059cbb906044016104df565b336000908152606d602052604081208054151580610e5c575060008160010154115b610ea85760405162461bcd60e51b815260206004820152601360248201527f4e6f7468696e6720746f20776974686472617700000000000000000000000000604482015260640161046b565b600281015415610efa5760405162461bcd60e51b815260206004820152601a60248201527f556e6c6f636b20616c726561647920696e2070726f6772657373000000000000604482015260640161046b565b606b54610f0790436123bb565b6002909101819055919050565b600054610100900460ff1680610f2d575060005460ff16155b610f905760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161046b565b600054610100900460ff16158015610fb2576000805461ffff19166101011790555b610fba6119d1565b606580546001600160a01b03808d167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255606680548c841690831617905560678054928b1692909116919091179055606b87905560688690556001600160801b038481167001000000000000000000000000000000000290861617606955606c805461ffff851661ffff1990911617905561105a82611078565b801561106c576000805461ff00191690555b50505050505050505050565b6033546001600160a01b031633146110d25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161046b565b600081116111225760405162461bcd60e51b815260206004820152601b60248201527f5469636b6574206475726174696f6e2063616e6e6f7420626520300000000000604482015260640161046b565b606a55565b6033546001600160a01b031633146111815760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161046b565b606980546001600160801b03928316700100000000000000000000000000000000029216919091179055565b6033546001600160a01b031633146112075760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161046b565b606980546fffffffffffffffffffffffffffffffff19166001600160801b0392909216919091179055565b6001600160a01b0381166000908152606d6020526040902060028101541561129c5760405162461bcd60e51b815260206004820152601e60248201527f43616e6e6f74206465706f736974207768696c6520756e6c6f636b696e670000604482015260640161046b565b8281600101600082825461048891906123bb565b336000908152606d6020526040902060028101546113105760405162461bcd60e51b815260206004820152601a60248201527f4e6f7420756e6c6f636b696e672c2063616e6e6f74206c6f636b000000000000604482015260640161046b565b6000600290910155565b60008082604001514361132d919061245f565b9050606a5481106113415750600092915050565b606954606c54600091611363916001600160801b039091169061ffff1661193a565b90506000606a5483836113769190612418565b61138091906123f8565b60695490915061139a9082906001600160801b0316612437565b95945050505050565b6033546001600160a01b031633146113fd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161046b565b606b55565b6033546001600160a01b0316331461145c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161046b565b606855565b6033546001600160a01b031633146114bb5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161046b565b6001600160a01b0381166115375760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161046b565b61154081611967565b50565b80516020808301516068546069546040808701516060880151608089015192516000986115ec98909796956001600160801b0316949101606097881b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000090811682529690971b9095166014870152602886019390935260809190911b6fffffffffffffffffffffffffffffffff1916604885015260588401526078830152609882015260b80190565b604051602081830303815290604052805190602001209050919050565b84516001600160a01b03166116605760405162461bcd60e51b815260206004820152601560248201527f5469636b65742073656e646572206973206e756c6c0000000000000000000000604482015260640161046b565b60208501516001600160a01b03166116ba5760405162461bcd60e51b815260206004820152601760248201527f5469636b65742072656465656d6572206973206e756c6c000000000000000000604482015260640161046b565b6000848152606e602052604090205460ff16156117195760405162461bcd60e51b815260206004820152601760248201527f5469636b657420616c72656164792072656465656d6564000000000000000000604482015260640161046b565b606085015160408051602081018690520160405160208183030381529060405280519060200120146117b35760405162461bcd60e51b815260206004820152602f60248201527f48617368206f662073656e64657252616e6420646f65736e2774206d6174636860448201527f2073656e64657252616e64486173680000000000000000000000000000000000606482015260840161046b565b6080850151604080516020810185905201604051602081830303815290604052805190602001201461184d5760405162461bcd60e51b815260206004820152603360248201527f48617368206f662072656465656d657252616e6420646f65736e2774206d617460448201527f63682072656465656d657252616e644861736800000000000000000000000000606482015260840161046b565b61185c81866000015186611a93565b6118ce5760405162461bcd60e51b815260206004820152602560248201527f5469636b657420646f65736e2774206861766520612076616c6964207369676e60448201527f6174757265000000000000000000000000000000000000000000000000000000606482015260840161046b565b60006118d98661131a565b90506118e6828483611abb565b6119325760405162461bcd60e51b815260206004820152601660248201527f5469636b6574206973206e6f7420612077696e6e657200000000000000000000604482015260640161046b565b505050505050565b600061271061195661ffff84166001600160801b038616612418565b61196091906123f8565b9392505050565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16806119ea575060005460ff16155b611a4d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161046b565b600054610100900460ff16158015611a6f576000805461ffff19166101011790555b611a77611b04565b611a7f611bb5565b8015611540576000805461ff001916905550565b6000826001600160a01b0316611aa98386611c5c565b6001600160a01b031614949350505050565b6000806080836001600160801b0316901b9050808585604051602001611ae2929190612317565b60408051601f1981840301815291905280516020909101201095945050505050565b600054610100900460ff1680611b1d575060005460ff16155b611b805760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161046b565b600054610100900460ff16158015611a7f576000805461ffff19166101011790558015611540576000805461ff001916905550565b600054610100900460ff1680611bce575060005460ff16155b611c315760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161046b565b600054610100900460ff16158015611c53576000805461ffff19166101011790555b611a7f33611967565b60008151604114611caf5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161046b565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115611d3c5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161046b565b601b8160ff161015611d5657611d53601b826123d3565b90505b8060ff16601b14158015611d6e57508060ff16601c14155b15611dc65760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b606482015260840161046b565b6040805160008082526020820180845289905260ff841692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015611e1a573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611e7d5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161046b565b9695505050505050565b80518015158114611e9757600080fd5b919050565b600060a08284031215611ead578081fd5b60405160a0810181811067ffffffffffffffff82111715611ed057611ed06124dc565b6040529050808235611ee1816124f2565b81526020830135611ef1816124f2565b806020830152506040830135604082015260608301356060820152608083013560808201525092915050565b80356001600160801b0381168114611e9757600080fd5b8051611e9781612507565b600060208284031215611f50578081fd5b8135611960816124f2565b60006020808385031215611f6d578182fd5b825167ffffffffffffffff80821115611f84578384fd5b818501915085601f830112611f97578384fd5b815181811115611fa957611fa96124dc565b8060051b9150611fba848301612362565b8181528481019084860184860187018a1015611fd4578788fd5b8795505b838610156120025780519450611fed856124f2565b84835260019590950194918601918601611fd8565b5098975050505050505050565b600060208284031215612020578081fd5b61196082611e87565b60006020828403121561203a578081fd5b5035919050565b60008060008060008060008060006101208a8c03121561205f578485fd5b893561206a816124f2565b985060208a013561207a816124f2565b975060408a013561208a816124f2565b965060608a0135955060808a013594506120a660a08b01611f1d565b93506120b460c08b01611f1d565b925060e08a01356120c481612507565b809250506101008a013590509295985092959850929598565b600060208083850312156120ef578182fd5b825167ffffffffffffffff80821115612106578384fd5b9084019060808287031215612119578384fd5b612121612339565b82518281111561212f578586fd5b83019150601f82018713612141578485fd5b815161215461214f82612393565b612362565b8181528886838601011115612167578687fd5b61217682878301888701612476565b825250612184838501611f34565b848201526040830151604082015261219e60608401611e87565b60608201529695505050505050565b6000604082840312156121be578081fd5b6040516040810181811067ffffffffffffffff821117156121e1576121e16124dc565b6040528251815260208301516121f6816124f2565b60208201529392505050565b600060a08284031215612213578081fd5b6119608383611e9c565b6000806000806101008587031215612233578182fd5b61223d8686611e9c565b935060a0850135925060c0850135915060e085013567ffffffffffffffff811115612266578182fd5b8501601f81018713612276578182fd5b803561228461214f82612393565b818152886020838501011115612298578384fd5b81602084016020830137908101602001929092525092959194509250565b6000602082840312156122c7578081fd5b61196082611f1d565b6000602082840312156122e1578081fd5b5051919050565b600080604083850312156122fa578182fd5b82359150602083013561230c816124f2565b809150509250929050565b60008351612329818460208801612476565b9190910191825250602001919050565b6040516080810167ffffffffffffffff8111828210171561235c5761235c6124dc565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561238b5761238b6124dc565b604052919050565b600067ffffffffffffffff8211156123ad576123ad6124dc565b50601f01601f191660200190565b600082198211156123ce576123ce6124c6565b500190565b600060ff821660ff84168060ff038211156123f0576123f06124c6565b019392505050565b60008261241357634e487b7160e01b81526012600452602481fd5b500490565b6000816000190483118215151615612432576124326124c6565b500290565b60006001600160801b0383811690831681811015612457576124576124c6565b039392505050565b600082821015612471576124716124c6565b500390565b60005b83811015612491578181015183820152602001612479565b838111156105315750506000910152565b600063ffffffff808316818114156124bc576124bc6124c6565b6001019392505050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461154057600080fd5b61ffff8116811461154057600080fdfea264697066735822122063e5bf9dab47c8cec92b9ebf3f4cc62e4f7cba6425747f61a6c93a689253583664736f6c63430008040033"

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
// Solidity: function decayRate() view returns(uint16)
func (_SyloTicketing *SyloTicketingCaller) DecayRate(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "decayRate")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DecayRate is a free data retrieval call binding the contract method 0xa9c1f2f1.
//
// Solidity: function decayRate() view returns(uint16)
func (_SyloTicketing *SyloTicketingSession) DecayRate() (uint16, error) {
	return _SyloTicketing.Contract.DecayRate(&_SyloTicketing.CallOpts)
}

// DecayRate is a free data retrieval call binding the contract method 0xa9c1f2f1.
//
// Solidity: function decayRate() view returns(uint16)
func (_SyloTicketing *SyloTicketingCallerSession) DecayRate() (uint16, error) {
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

// Initialize is a paid mutator transaction binding the contract method 0x9d01393f.
//
// Solidity: function initialize(address token, address listings, address stakingManager, uint256 _unlockDuration, uint256 _faceValue, uint128 _baseLiveWinProb, uint128 _expiredWinProb, uint16 _decayRate, uint256 _ticketDuration) returns()
func (_SyloTicketing *SyloTicketingTransactor) Initialize(opts *bind.TransactOpts, token common.Address, listings common.Address, stakingManager common.Address, _unlockDuration *big.Int, _faceValue *big.Int, _baseLiveWinProb *big.Int, _expiredWinProb *big.Int, _decayRate uint16, _ticketDuration *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "initialize", token, listings, stakingManager, _unlockDuration, _faceValue, _baseLiveWinProb, _expiredWinProb, _decayRate, _ticketDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0x9d01393f.
//
// Solidity: function initialize(address token, address listings, address stakingManager, uint256 _unlockDuration, uint256 _faceValue, uint128 _baseLiveWinProb, uint128 _expiredWinProb, uint16 _decayRate, uint256 _ticketDuration) returns()
func (_SyloTicketing *SyloTicketingSession) Initialize(token common.Address, listings common.Address, stakingManager common.Address, _unlockDuration *big.Int, _faceValue *big.Int, _baseLiveWinProb *big.Int, _expiredWinProb *big.Int, _decayRate uint16, _ticketDuration *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.Initialize(&_SyloTicketing.TransactOpts, token, listings, stakingManager, _unlockDuration, _faceValue, _baseLiveWinProb, _expiredWinProb, _decayRate, _ticketDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0x9d01393f.
//
// Solidity: function initialize(address token, address listings, address stakingManager, uint256 _unlockDuration, uint256 _faceValue, uint128 _baseLiveWinProb, uint128 _expiredWinProb, uint16 _decayRate, uint256 _ticketDuration) returns()
func (_SyloTicketing *SyloTicketingTransactorSession) Initialize(token common.Address, listings common.Address, stakingManager common.Address, _unlockDuration *big.Int, _faceValue *big.Int, _baseLiveWinProb *big.Int, _expiredWinProb *big.Int, _decayRate uint16, _ticketDuration *big.Int) (*types.Transaction, error) {
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

// SetTicketDuration is a paid mutator transaction binding the contract method 0x9e9ceeca.
//
// Solidity: function setTicketDuration(uint256 _ticketDuration) returns()
func (_SyloTicketing *SyloTicketingTransactor) SetTicketDuration(opts *bind.TransactOpts, _ticketDuration *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "setTicketDuration", _ticketDuration)
}

// SetTicketDuration is a paid mutator transaction binding the contract method 0x9e9ceeca.
//
// Solidity: function setTicketDuration(uint256 _ticketDuration) returns()
func (_SyloTicketing *SyloTicketingSession) SetTicketDuration(_ticketDuration *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.SetTicketDuration(&_SyloTicketing.TransactOpts, _ticketDuration)
}

// SetTicketDuration is a paid mutator transaction binding the contract method 0x9e9ceeca.
//
// Solidity: function setTicketDuration(uint256 _ticketDuration) returns()
func (_SyloTicketing *SyloTicketingTransactorSession) SetTicketDuration(_ticketDuration *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.SetTicketDuration(&_SyloTicketing.TransactOpts, _ticketDuration)
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
