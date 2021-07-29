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
	SenderNonce      uint32
}

// SyloTicketingABI is the input ABI used to generate the binding from.
const SyloTicketingABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"baseWinProb\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decayRate\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"escrow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"faceValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minProbConstant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ticketLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usedTickets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractListings\",\"name\":\"listings\",\"type\":\"address\"},{\"internalType\":\"contractStakingManager\",\"name\":\"stakingManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unlockDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseWinProb\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProbConstant\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ticketLength\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newUnlockDuration\",\"type\":\"uint256\"}],\"name\":\"setUnlockDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_faceValue\",\"type\":\"uint256\"}],\"name\":\"setFaceValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseWinProb\",\"type\":\"uint256\"}],\"name\":\"setWinProb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositEscrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"generationBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiverRandHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"senderNonce\",\"type\":\"uint32\"}],\"internalType\":\"structSyloTicketing.Ticket\",\"name\":\"ticket\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"receiverRand\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"generationBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiverRandHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"senderNonce\",\"type\":\"uint32\"}],\"internalType\":\"structSyloTicketing.Ticket\",\"name\":\"ticket\",\"type\":\"tuple\"}],\"name\":\"calculateWinningProbability\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"generationBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiverRandHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"senderNonce\",\"type\":\"uint32\"}],\"internalType\":\"structSyloTicketing.Ticket\",\"name\":\"ticket\",\"type\":\"tuple\"}],\"name\":\"calculatePayout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"generationBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiverRandHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"senderNonce\",\"type\":\"uint32\"}],\"internalType\":\"structSyloTicketing.Ticket\",\"name\":\"ticket\",\"type\":\"tuple\"}],\"name\":\"getTicketHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SyloTicketingBin is the compiled bytecode used for deploying new contracts.
var SyloTicketingBin = "0x608060405234801561001057600080fd5b506122fa806100206000396000f3fe608060405234801561001057600080fd5b50600436106101a35760003560e01c8063a1448819116100ee578063db1536f211610097578063df8fa43011610071578063df8fa43014610328578063ef8032ef1461033b578063f2fde38b1461034e578063fc7e286d1461036157600080fd5b8063db1536f2146102f9578063dbcaefc61461030c578063dd9007691461031557600080fd5b8063b93f2157116100c8578063b93f2157146102cb578063cdba7342146102de578063d088a231146102f157600080fd5b8063a14488191461028a578063a9c1f2f11461029d578063b7b25b7c146102c257600080fd5b8063715018a61161015057806380267b571161012a57806380267b57146102535780638a1fcd60146102665780638da5cb5b1461026f57600080fd5b8063715018a61461023057806372b0d90c146102385780637d6babb41461024b57600080fd5b806344fd9caa1161018157806344fd9caa146101e15780634a6ffa3a146101ea57806359a515ba146101fd57600080fd5b8063169e1ca8146101a85780633ccfd60b146101bd57806340d6f2ae146101c5575b600080fd5b6101bb6101b63660046120a1565b6103ab565b005b6101bb6104dd565b6101ce60695481565b6040519081526020015b60405180910390f35b6101ce60685481565b6101ce6101f8366004611fdd565b6104e8565b61022061020b366004611e2a565b606f6020526000908152604090205460ff1681565b60405190151581526020016101d8565b6101bb6105a1565b6101bb610246366004611d40565b610605565b6101ce610717565b6101bb610261366004611e2a565b6107f1565b6101ce606d5481565b6033546040516001600160a01b0390911681526020016101d8565b6101ce610298366004611fdd565b610850565b606a546102ad9063ffffffff1681565b60405163ffffffff90911681526020016101d8565b6101ce606c5481565b6101bb6102d9366004611ff8565b6108be565b6101bb6102ec3660046120a1565b611054565b6101bb6110d2565b6101ce610307366004611fdd565b61113c565b6101ce606b5481565b6101bb610323366004611e2a565b611208565b6101bb610336366004611e42565b611267565b6101bb610349366004611e2a565b611398565b6101bb61035c366004611d40565b6113f7565b61039061036f366004611d40565b606e6020526000908152604090208054600182015460029092015490919083565b604080519384526020840192909252908201526060016101d8565b6001600160a01b0381166000908152606e6020526040902060028101541561041a5760405162461bcd60e51b815260206004820152601e60248201527f43616e6e6f74206465706f736974207768696c6520756e6c6f636b696e67000060448201526064015b60405180910390fd5b8281600001600082825461042e9190612174565b90915550506065546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590526001600160a01b03909116906323b872dd906064015b602060405180830381600087803b15801561049f57600080fd5b505af11580156104b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d79190611e10565b50505050565b6104e633610605565b565b6000808260400151436104fb919061221c565b90506000606c5482606461050f91906121d1565b61051991906121b1565b606954606a549192506000916105369063ffffffff1660646121f0565b63ffffffff1661054691906121b1565b9050600061055483836121d1565b9050606b54600019610566919061221c565b606954111561057c575060001995945050505050565b606b548160695461058d919061221c565b6105979190612174565b9695505050505050565b6033546001600160a01b031633146105fb5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610411565b6104e660006114d9565b336000908152606e6020526040902060028101546106655760405162461bcd60e51b815260206004820152601560248201527f4465706f73697473206e6f7420756e6c6f636b656400000000000000000000006044820152606401610411565b438160020154106106b85760405162461bcd60e51b815260206004820152601a60248201527f556e6c6f636b20706572696f64206e6f7420636f6d706c6574650000000000006044820152606401610411565b600181015481546000916106cb91612174565b600080845560018401819055600284015560655460405163a9059cbb60e01b81526001600160a01b0386811660048301526024820184905292935091169063a9059cbb90604401610485565b336000908152606e602052604081208054151580610739575060008160010154115b6107855760405162461bcd60e51b815260206004820152601360248201527f4e6f7468696e6720746f207769746864726177000000000000000000000000006044820152606401610411565b6002810154156107d75760405162461bcd60e51b815260206004820152601a60248201527f556e6c6f636b20616c726561647920696e2070726f67726573730000000000006044820152606401610411565b606d546107e49043612174565b6002909101819055919050565b6033546001600160a01b0316331461084b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610411565b606955565b600080606c5483604001516108659190612174565b90504381116108775750600092915050565b6000836040015143610889919061221c565b606c54610896919061221c565b90506000606c54606854836108ab91906121d1565b6108b591906121b1565b95945050505050565b60006108c98461113c565b90506108d784828585611543565b83516001600160a01b03166000908152606e60205260408120906108fa86610850565b905080826001015483600001546109119190612174565b10156109855760405162461bcd60e51b815260206004820152602760248201527f53656e64657220646f65736e2774206861766520656e6f7567682066756e647360448201527f20746f20706179000000000000000000000000000000000000000000000000006064820152608401610411565b6000838152606f60209081526040808320805460ff191660011790556066549189015190517f084af0b20000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169063084af0b29060240160006040518083038186803b158015610a0157600080fd5b505afa158015610a15573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a3d9190810190611eb8565b9050806060015115156001151514610abd5760405162461bcd60e51b815260206004820152602960248201527f5469636b6574207265636569766572206d757374206861766520612076616c6960448201527f64206c697374696e6700000000000000000000000000000000000000000000006064820152608401610411565b60675460208801516040517fdf349ed50000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600092919091169063df349ed59060240160206040518083038186803b158015610b2457600080fd5b505afa158015610b38573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b5c9190612089565b905080610bab5760405162461bcd60e51b815260206004820152601f60248201527f5469636b6574207265636569766572206d7573742068617665207374616b65006044820152606401610411565b83546068541115610cdf576065546020890151855460405163a9059cbb60e01b81526001600160a01b039283166004820152602481019190915291169063a9059cbb90604401602060405180830381600087803b158015610c0b57600080fd5b505af1158015610c1f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c439190611e10565b50606554600185015460405163a9059cbb60e01b81526001600160a01b039092166004830181905260248301919091529063a9059cbb90604401602060405180830381600087803b158015610c9757600080fd5b505af1158015610cab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ccf9190611e10565b506000808555600185015561104a565b8354610ceb90846117d9565b84556020820151600090606490610d0690869060ff166121d1565b610d1091906121b1565b60675460208b01516040517f791936100000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015292935060009291169063791936109060240160006040518083038186803b158015610d7857600080fd5b505afa158015610d8c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610db49190810190611d5c565b90508160005b82518163ffffffff161015610fa05760675483516000916001600160a01b0316906382dda22d90869063ffffffff8616908110610e0757634e487b7160e01b600052603260045260246000fd5b60200260200101518f602001516040518363ffffffff1660e01b8152600401610e469291906001600160a01b0392831681529116602082015260400190565b604080518083038186803b158015610e5d57600080fd5b505afa158015610e71573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e959190611f88565b9050600086868360000151610eaa91906121d1565b610eb491906121b1565b9050610ec0818561221c565b60655486519195506001600160a01b03169063a9059cbb90879063ffffffff8716908110610efe57634e487b7160e01b600052603260045260246000fd5b6020026020010151836040518363ffffffff1660e01b8152600401610f389291906001600160a01b03929092168252602082015260400190565b602060405180830381600087803b158015610f5257600080fd5b505af1158015610f66573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f8a9190611e10565b5050508080610f989061225f565b915050610dba565b50600081610fae858961221c565b610fb89190612174565b60655460208e015160405163a9059cbb60e01b81526001600160a01b03918216600482015260248101849052929350169063a9059cbb90604401602060405180830381600087803b15801561100c57600080fd5b505af1158015611020573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110449190611e10565b50505050505b5050505050505050565b6001600160a01b0381166000908152606e602052604090206002810154156110be5760405162461bcd60e51b815260206004820152601e60248201527f43616e6e6f74206465706f736974207768696c6520756e6c6f636b696e6700006044820152606401610411565b8281600101600082825461042e9190612174565b336000908152606e6020526040902060028101546111325760405162461bcd60e51b815260206004820152601a60248201527f4e6f7420756e6c6f636b696e672c2063616e6e6f74206c6f636b0000000000006044820152606401610411565b6000600290910155565b80516020808301516068546069546040808701516060880151608089015192516000986111eb9890979695949101606097881b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000090811682529690971b9095166014870152602886019390935260488501919091526068840152608883015260e01b7fffffffff000000000000000000000000000000000000000000000000000000001660a882015260ac0190565b604051602081830303815290604052805190602001209050919050565b6033546001600160a01b031633146112625760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610411565b606d55565b600054610100900460ff1680611280575060005460ff16155b6112e35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610411565b600054610100900460ff16158015611305576000805461ffff19166101011790555b61130d6117ec565b606580546001600160a01b03808c167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255606680548b841690831617905560678054928a1692909116919091179055606d86905560688590556069849055606b839055606c829055801561138d576000805461ff00191690555b505050505050505050565b6033546001600160a01b031633146113f25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610411565b606855565b6033546001600160a01b031633146114515760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610411565b6001600160a01b0381166114cd5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610411565b6114d6816114d9565b50565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b83516001600160a01b031661159a5760405162461bcd60e51b815260206004820152601560248201527f5469636b65742073656e646572206973206e756c6c00000000000000000000006044820152606401610411565b60208401516001600160a01b03166115f45760405162461bcd60e51b815260206004820152601760248201527f5469636b6574207265636569766572206973206e756c6c0000000000000000006044820152606401610411565b6000838152606f602052604090205460ff16156116535760405162461bcd60e51b815260206004820152601760248201527f5469636b657420616c72656164792072656465656d65640000000000000000006044820152606401610411565b606084015160408051602081018590520160405160208183030381529060405280519060200120146116ed5760405162461bcd60e51b815260206004820152603360248201527f48617368206f6620726563656976657252616e6420646f65736e2774206d617460448201527f636820726563656976657252616e6448617368000000000000000000000000006064820152608401610411565b6116fc818560000151856118ae565b61176e5760405162461bcd60e51b815260206004820152602560248201527f5469636b657420646f65736e2774206861766520612076616c6964207369676e60448201527f61747572650000000000000000000000000000000000000000000000000000006064820152608401610411565b6000611779856104e8565b90506117868284836118d6565b6117d25760405162461bcd60e51b815260206004820152601660248201527f5469636b6574206973206e6f7420612077696e6e6572000000000000000000006044820152606401610411565b5050505050565b60006117e5828461221c565b9392505050565b600054610100900460ff1680611805575060005460ff16155b6118685760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610411565b600054610100900460ff1615801561188a576000805461ffff19166101011790555b61189261190d565b61189a6119be565b80156114d6576000805461ff001916905550565b6000826001600160a01b03166118c48386611a65565b6001600160a01b031614949350505050565b60008184846040516020016118ec9291906120d0565b60408051601f19818403018152919052805160209091012010949350505050565b600054610100900460ff1680611926575060005460ff16155b6119895760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610411565b600054610100900460ff1615801561189a576000805461ffff191661010117905580156114d6576000805461ff001916905550565b600054610100900460ff16806119d7575060005460ff16155b611a3a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610411565b600054610100900460ff16158015611a5c576000805461ffff19166101011790555b61189a336114d9565b60008151604114611ab85760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610411565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115611b455760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610411565b601b8160ff161015611b5f57611b5c601b8261218c565b90505b8060ff16601b14158015611b7757508060ff16601c14155b15611bcf5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610411565b6040805160008082526020820180845289905260ff841692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015611c23573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166105975760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610411565b80518015158114611c9657600080fd5b919050565b600060a08284031215611cac578081fd5b60405160a0810181811067ffffffffffffffff82111715611ccf57611ccf612299565b6040529050808235611ce0816122af565b81526020830135611cf0816122af565b806020830152506040830135604082015260608301356060820152608083013563ffffffff81168114611d2257600080fd5b6080919091015292915050565b805160ff81168114611c9657600080fd5b600060208284031215611d51578081fd5b81356117e5816122af565b60006020808385031215611d6e578182fd5b825167ffffffffffffffff80821115611d85578384fd5b818501915085601f830112611d98578384fd5b815181811115611daa57611daa612299565b8060051b9150611dbb84830161211b565b8181528481019084860184860187018a1015611dd5578788fd5b8795505b83861015611e035780519450611dee856122af565b84835260019590950194918601918601611dd9565b5098975050505050505050565b600060208284031215611e21578081fd5b6117e582611c86565b600060208284031215611e3b578081fd5b5035919050565b600080600080600080600080610100898b031215611e5e578384fd5b8835611e69816122af565b97506020890135611e79816122af565b96506040890135611e89816122af565b979a96995096976060810135975060808101359660a0820135965060c0820135955060e0909101359350915050565b60006020808385031215611eca578182fd5b825167ffffffffffffffff80821115611ee1578384fd5b9084019060808287031215611ef4578384fd5b611efc6120f2565b825182811115611f0a578586fd5b83019150601f82018713611f1c578485fd5b8151611f2f611f2a8261214c565b61211b565b8181528886838601011115611f42578687fd5b611f5182878301888701612233565b825250611f5f838501611d2f565b8482015260408301516040820152611f7960608401611c86565b60608201529695505050505050565b600060408284031215611f99578081fd5b6040516040810181811067ffffffffffffffff82111715611fbc57611fbc612299565b604052825181526020830151611fd1816122af565b60208201529392505050565b600060a08284031215611fee578081fd5b6117e58383611c9b565b600080600060e0848603121561200c578081fd5b6120168585611c9b565b925060a0840135915060c084013567ffffffffffffffff811115612038578182fd5b8401601f81018613612048578182fd5b8035612056611f2a8261214c565b81815287602083850101111561206a578384fd5b8160208401602083013783602083830101528093505050509250925092565b60006020828403121561209a578081fd5b5051919050565b600080604083850312156120b3578182fd5b8235915060208301356120c5816122af565b809150509250929050565b600083516120e2818460208801612233565b9190910191825250602001919050565b6040516080810167ffffffffffffffff8111828210171561211557612115612299565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561214457612144612299565b604052919050565b600067ffffffffffffffff82111561216657612166612299565b50601f01601f191660200190565b6000821982111561218757612187612283565b500190565b600060ff821660ff84168060ff038211156121a9576121a9612283565b019392505050565b6000826121cc57634e487b7160e01b81526012600452602481fd5b500490565b60008160001904831182151516156121eb576121eb612283565b500290565b600063ffffffff8083168185168183048111821515161561221357612213612283565b02949350505050565b60008282101561222e5761222e612283565b500390565b60005b8381101561224e578181015183820152602001612236565b838111156104d75750506000910152565b600063ffffffff8083168181141561227957612279612283565b6001019392505050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146114d657600080fdfea26469706673582212206b78ac61d93c17f4320721e0415c1386ae421596b6ab403a3e9b916f127f76fe64736f6c63430008040033"

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

// BaseWinProb is a free data retrieval call binding the contract method 0x40d6f2ae.
//
// Solidity: function baseWinProb() view returns(uint256)
func (_SyloTicketing *SyloTicketingCaller) BaseWinProb(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "baseWinProb")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseWinProb is a free data retrieval call binding the contract method 0x40d6f2ae.
//
// Solidity: function baseWinProb() view returns(uint256)
func (_SyloTicketing *SyloTicketingSession) BaseWinProb() (*big.Int, error) {
	return _SyloTicketing.Contract.BaseWinProb(&_SyloTicketing.CallOpts)
}

// BaseWinProb is a free data retrieval call binding the contract method 0x40d6f2ae.
//
// Solidity: function baseWinProb() view returns(uint256)
func (_SyloTicketing *SyloTicketingCallerSession) BaseWinProb() (*big.Int, error) {
	return _SyloTicketing.Contract.BaseWinProb(&_SyloTicketing.CallOpts)
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

// CalculateWinningProbability is a free data retrieval call binding the contract method 0x4a6ffa3a.
//
// Solidity: function calculateWinningProbability((address,address,uint256,bytes32,uint32) ticket) view returns(uint256)
func (_SyloTicketing *SyloTicketingCaller) CalculateWinningProbability(opts *bind.CallOpts, ticket SyloTicketingTicket) (*big.Int, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "calculateWinningProbability", ticket)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateWinningProbability is a free data retrieval call binding the contract method 0x4a6ffa3a.
//
// Solidity: function calculateWinningProbability((address,address,uint256,bytes32,uint32) ticket) view returns(uint256)
func (_SyloTicketing *SyloTicketingSession) CalculateWinningProbability(ticket SyloTicketingTicket) (*big.Int, error) {
	return _SyloTicketing.Contract.CalculateWinningProbability(&_SyloTicketing.CallOpts, ticket)
}

// CalculateWinningProbability is a free data retrieval call binding the contract method 0x4a6ffa3a.
//
// Solidity: function calculateWinningProbability((address,address,uint256,bytes32,uint32) ticket) view returns(uint256)
func (_SyloTicketing *SyloTicketingCallerSession) CalculateWinningProbability(ticket SyloTicketingTicket) (*big.Int, error) {
	return _SyloTicketing.Contract.CalculateWinningProbability(&_SyloTicketing.CallOpts, ticket)
}

// DecayRate is a free data retrieval call binding the contract method 0xa9c1f2f1.
//
// Solidity: function decayRate() view returns(uint32)
func (_SyloTicketing *SyloTicketingCaller) DecayRate(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "decayRate")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DecayRate is a free data retrieval call binding the contract method 0xa9c1f2f1.
//
// Solidity: function decayRate() view returns(uint32)
func (_SyloTicketing *SyloTicketingSession) DecayRate() (uint32, error) {
	return _SyloTicketing.Contract.DecayRate(&_SyloTicketing.CallOpts)
}

// DecayRate is a free data retrieval call binding the contract method 0xa9c1f2f1.
//
// Solidity: function decayRate() view returns(uint32)
func (_SyloTicketing *SyloTicketingCallerSession) DecayRate() (uint32, error) {
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

// MinProbConstant is a free data retrieval call binding the contract method 0xdbcaefc6.
//
// Solidity: function minProbConstant() view returns(uint256)
func (_SyloTicketing *SyloTicketingCaller) MinProbConstant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "minProbConstant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinProbConstant is a free data retrieval call binding the contract method 0xdbcaefc6.
//
// Solidity: function minProbConstant() view returns(uint256)
func (_SyloTicketing *SyloTicketingSession) MinProbConstant() (*big.Int, error) {
	return _SyloTicketing.Contract.MinProbConstant(&_SyloTicketing.CallOpts)
}

// MinProbConstant is a free data retrieval call binding the contract method 0xdbcaefc6.
//
// Solidity: function minProbConstant() view returns(uint256)
func (_SyloTicketing *SyloTicketingCallerSession) MinProbConstant() (*big.Int, error) {
	return _SyloTicketing.Contract.MinProbConstant(&_SyloTicketing.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0xdf8fa430.
//
// Solidity: function initialize(address token, address listings, address stakingManager, uint256 _unlockDuration, uint256 _faceValue, uint256 _baseWinProb, uint256 _minProbConstant, uint256 _ticketLength) returns()
func (_SyloTicketing *SyloTicketingTransactor) Initialize(opts *bind.TransactOpts, token common.Address, listings common.Address, stakingManager common.Address, _unlockDuration *big.Int, _faceValue *big.Int, _baseWinProb *big.Int, _minProbConstant *big.Int, _ticketLength *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "initialize", token, listings, stakingManager, _unlockDuration, _faceValue, _baseWinProb, _minProbConstant, _ticketLength)
}

// Initialize is a paid mutator transaction binding the contract method 0xdf8fa430.
//
// Solidity: function initialize(address token, address listings, address stakingManager, uint256 _unlockDuration, uint256 _faceValue, uint256 _baseWinProb, uint256 _minProbConstant, uint256 _ticketLength) returns()
func (_SyloTicketing *SyloTicketingSession) Initialize(token common.Address, listings common.Address, stakingManager common.Address, _unlockDuration *big.Int, _faceValue *big.Int, _baseWinProb *big.Int, _minProbConstant *big.Int, _ticketLength *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.Initialize(&_SyloTicketing.TransactOpts, token, listings, stakingManager, _unlockDuration, _faceValue, _baseWinProb, _minProbConstant, _ticketLength)
}

// Initialize is a paid mutator transaction binding the contract method 0xdf8fa430.
//
// Solidity: function initialize(address token, address listings, address stakingManager, uint256 _unlockDuration, uint256 _faceValue, uint256 _baseWinProb, uint256 _minProbConstant, uint256 _ticketLength) returns()
func (_SyloTicketing *SyloTicketingTransactorSession) Initialize(token common.Address, listings common.Address, stakingManager common.Address, _unlockDuration *big.Int, _faceValue *big.Int, _baseWinProb *big.Int, _minProbConstant *big.Int, _ticketLength *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.Initialize(&_SyloTicketing.TransactOpts, token, listings, stakingManager, _unlockDuration, _faceValue, _baseWinProb, _minProbConstant, _ticketLength)
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
// Solidity: function setWinProb(uint256 _baseWinProb) returns()
func (_SyloTicketing *SyloTicketingTransactor) SetWinProb(opts *bind.TransactOpts, _baseWinProb *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "setWinProb", _baseWinProb)
}

// SetWinProb is a paid mutator transaction binding the contract method 0x80267b57.
//
// Solidity: function setWinProb(uint256 _baseWinProb) returns()
func (_SyloTicketing *SyloTicketingSession) SetWinProb(_baseWinProb *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.SetWinProb(&_SyloTicketing.TransactOpts, _baseWinProb)
}

// SetWinProb is a paid mutator transaction binding the contract method 0x80267b57.
//
// Solidity: function setWinProb(uint256 _baseWinProb) returns()
func (_SyloTicketing *SyloTicketingTransactorSession) SetWinProb(_baseWinProb *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.SetWinProb(&_SyloTicketing.TransactOpts, _baseWinProb)
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
