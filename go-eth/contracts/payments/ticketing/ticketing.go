// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ticketing

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EpochsManagerEpoch is an auto generated low-level Go binding around an user-defined struct.
type EpochsManagerEpoch struct {
	Iteration               *big.Int
	StartBlock              *big.Int
	Duration                *big.Int
	EndBlock                *big.Int
	DefaultPayoutPercentage uint16
	FaceValue               *big.Int
	BaseLiveWinProb         *big.Int
	ExpiredWinProb          *big.Int
	TicketDuration          *big.Int
	DecayRate               uint16
}

// SyloTicketingTicket is an auto generated low-level Go binding around an user-defined struct.
type SyloTicketingTicket struct {
	EpochId         *big.Int
	Sender          common.Address
	Redeemer        common.Address
	GenerationBlock *big.Int
	SenderCommit    [32]byte
	RedeemerCommit  [32]byte
}

// SyloTicketingMetaData contains all meta data concerning the SyloTicketing contract.
var SyloTicketingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"escrow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usedTickets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractListings\",\"name\":\"listings\",\"type\":\"address\"},{\"internalType\":\"contractStakingManager\",\"name\":\"stakingManager\",\"type\":\"address\"},{\"internalType\":\"contractDirectory\",\"name\":\"directory\",\"type\":\"address\"},{\"internalType\":\"contractEpochsManager\",\"name\":\"epochsManager\",\"type\":\"address\"},{\"internalType\":\"contractRewardsManager\",\"name\":\"rewardsManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unlockDuration\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newUnlockDuration\",\"type\":\"uint256\"}],\"name\":\"setUnlockDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositEscrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"generationBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"senderCommit\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"redeemerCommit\",\"type\":\"bytes32\"}],\"internalType\":\"structSyloTicketing.Ticket\",\"name\":\"ticket\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"senderRand\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemerRand\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"generationBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"senderCommit\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"redeemerCommit\",\"type\":\"bytes32\"}],\"internalType\":\"structSyloTicketing.Ticket\",\"name\":\"ticket\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"iteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"defaultPayoutPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"faceValue\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"baseLiveWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"expiredWinProb\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"ticketDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"decayRate\",\"type\":\"uint16\"}],\"internalType\":\"structEpochsManager.Epoch\",\"name\":\"epoch\",\"type\":\"tuple\"}],\"name\":\"calculateWinningProbability\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"generationBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"senderCommit\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"redeemerCommit\",\"type\":\"bytes32\"}],\"internalType\":\"structSyloTicketing.Ticket\",\"name\":\"ticket\",\"type\":\"tuple\"}],\"name\":\"getTicketHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506122cd806100206000396000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c80638a1fcd60116100b2578063d088a23111610081578063dd90076911610066578063dd90076914610252578063f2fde38b14610265578063fc7e286d1461027857600080fd5b8063d088a23114610237578063d8024d321461023f57600080fd5b80638a1fcd60146101cc5780638da5cb5b146101d5578063a57c4bf9146101f0578063cdba73421461022457600080fd5b8063715018a6116100ee578063715018a61461018857806372b0d90c146101905780637761134a146101a35780637d6babb4146101c457600080fd5b80631460e39014610120578063169e1ca8146101355780633ccfd60b1461014857806359a515ba14610150575b600080fd5b61013361012e366004611caf565b6102c2565b005b610133610143366004612041565b610408565b610133610535565b61017361015e366004611c97565b606d6020526000908152604090205460ff1681565b60405190151581526020015b60405180910390f35b610133610540565b61013361019e366004611c61565b6105a4565b6101b66101b1366004611ea8565b6106b6565b60405190815260200161017f565b6101b6610750565b6101b6606b5481565b6033546040516001600160a01b03909116815260200161017f565b6102036101fe366004611ec3565b61082a565b6040516fffffffffffffffffffffffffffffffff909116815260200161017f565b610133610232366004612041565b6108a8565b610133610926565b61013361024d366004611f8f565b610990565b610133610260366004611c97565b610f2e565b610133610273366004611c61565b610f8d565b6102a7610286366004611c61565b606c6020526000908152604090208054600182015460029092015490919083565b6040805193845260208401929092529082015260600161017f565b600054610100900460ff16806102db575060005460ff16155b6103435760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b600054610100900460ff16158015610365576000805461ffff19166101011790555b61036d61106f565b606580547fffffffffffffffffffffffff00000000000000000000000000000000000000009081166001600160a01b038b8116919091179092556066805482168a8416179055606780548216898416179055606880548216888416179055606a8054821687841617905560698054909116918516919091179055606b82905580156103fe576000805461ff00191690555b5050505050505050565b6001600160a01b0381166000908152606c602052604090206002810154156104725760405162461bcd60e51b815260206004820152601e60248201527f43616e6e6f74206465706f736974207768696c6520756e6c6f636b696e670000604482015260640161033a565b828160000160008282546104869190612138565b90915550506065546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590526001600160a01b03909116906323b872dd906064015b602060405180830381600087803b1580156104f757600080fd5b505af115801561050b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061052f9190611c7d565b50505050565b61053e336105a4565b565b6033546001600160a01b0316331461059a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161033a565b61053e6000611131565b336000908152606c6020526040902060028101546106045760405162461bcd60e51b815260206004820152601560248201527f4465706f73697473206e6f7420756e6c6f636b65640000000000000000000000604482015260640161033a565b438160020154106106575760405162461bcd60e51b815260206004820152601a60248201527f556e6c6f636b20706572696f64206e6f7420636f6d706c657465000000000000604482015260640161033a565b6001810154815460009161066a91612138565b600080845560018401819055600284015560655460405163a9059cbb60e01b81526001600160a01b0386811660048301526024820184905292935091169063a9059cbb906044016104dd565b80516020808301516040808501516060860151608087015160a08801519351600097610733979096959101958652606094851b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000090811660208801529390941b909216603485015260488401526068830152608882015260a80190565b604051602081830303815290604052805190602001209050919050565b336000908152606c602052604081208054151580610772575060008160010154115b6107be5760405162461bcd60e51b815260206004820152601360248201527f4e6f7468696e6720746f20776974686472617700000000000000000000000000604482015260640161033a565b6002810154156108105760405162461bcd60e51b815260206004820152601a60248201527f556e6c6f636b20616c726561647920696e2070726f6772657373000000000000604482015260640161033a565b606b5461081d9043612138565b6002909101819055919050565b60008083606001514361083d91906121e5565b905082610100015181106108555760009150506108a2565b600061086a8460c0015185610120015161119b565b9050600084610100015183836108809190612195565b61088a9190612175565b9050808560c0015161089c91906121b4565b93505050505b92915050565b6001600160a01b0381166000908152606c602052604090206002810154156109125760405162461bcd60e51b815260206004820152601e60248201527f43616e6e6f74206465706f736974207768696c6520756e6c6f636b696e670000604482015260640161033a565b828160010160008282546104869190612138565b336000908152606c6020526040902060028101546109865760405162461bcd60e51b815260206004820152601a60248201527f4e6f7420756e6c6f636b696e672c2063616e6e6f74206c6f636b000000000000604482015260640161033a565b6000600290910155565b606a5484516040517fbc0bc6ba00000000000000000000000000000000000000000000000000000000815260048101919091526000916001600160a01b03169063bc0bc6ba906024016101406040518083038186803b1580156109f257600080fd5b505afa158015610a06573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a2a9190611d3b565b90506000816020015111610aa65760405162461bcd60e51b815260206004820152602860248201527f5469636b65742773206173736f6369617465642065706f636820646f6573206e60448201527f6f74206578697374000000000000000000000000000000000000000000000000606482015260840161033a565b8060200151856060015110158015610ad857506000816060015111610acc576001610ad8565b80606001518560600151105b610b4a5760405162461bcd60e51b815260206004820152603a60248201527f54686973207469636b657420776173206e6f742067656e65726174656420647560448201527f72696e672069742773206173736f6369617465642065706f6368000000000000606482015260840161033a565b6000610b55866106b6565b9050610b658682878787876111d1565b60665460408781015190517f084af0b20000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600092919091169063084af0b29060240160006040518083038186803b158015610bcc57600080fd5b505afa158015610be0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610c089190810190611dd8565b9050806060015115156001151514610c885760405162461bcd60e51b815260206004820152602960248201527f5469636b65742072656465656d6572206d757374206861766520612076616c6960448201527f64206c697374696e670000000000000000000000000000000000000000000000606482015260840161033a565b6000828152606d6020526040808220805460ff1916600117905560685489518a83015192517f1bdcc9ad0000000000000000000000000000000000000000000000000000000081526001600160a01b0390921692631bdcc9ad92610d029291906004019182526001600160a01b0316602082015260400190565b60206040518083038186803b158015610d1a57600080fd5b505afa158015610d2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d529190612029565b905060008111610dca5760405162461bcd60e51b815260206004820152603d60248201527f5469636b65742072656465656d6572206d7573742068617665206a6f696e656460448201527f20746865206469726563746f727920666f7220746869732065706f6368000000606482015260840161033a565b60695488516040808b015190517fd96d57670000000000000000000000000000000000000000000000000000000081526000936001600160a01b03169263d96d576792610e2b926004019182526001600160a01b0316602082015260400190565b60206040518083038186803b158015610e4357600080fd5b505afa158015610e57573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e7b9190612029565b905060008111610f195760405162461bcd60e51b815260206004820152604660248201527f5469636b65742072656465656d6572206d757374206861766520696e6974696160448201527f6c697a65642074686569722072657761726420706f6f6c20666f72207468697360648201527f2065706f63680000000000000000000000000000000000000000000000000000608482015260a40161033a565b610f23858a611507565b505050505050505050565b6033546001600160a01b03163314610f885760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161033a565b606b55565b6033546001600160a01b03163314610fe75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161033a565b6001600160a01b0381166110635760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161033a565b61106c81611131565b50565b600054610100900460ff1680611088575060005460ff16155b6110eb5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161033a565b600054610100900460ff1615801561110d576000805461ffff19166101011790555b611115611604565b61111d6116b5565b801561106c576000805461ff001916905550565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60006127106111c061ffff84166fffffffffffffffffffffffffffffffff8616612195565b6111ca9190612175565b9392505050565b60208601516001600160a01b031661122b5760405162461bcd60e51b815260206004820152601560248201527f5469636b65742073656e646572206973206e756c6c0000000000000000000000604482015260640161033a565b60408601516001600160a01b03166112855760405162461bcd60e51b815260206004820152601760248201527f5469636b65742072656465656d6572206973206e756c6c000000000000000000604482015260640161033a565b6000858152606d602052604090205460ff16156112e45760405162461bcd60e51b815260206004820152601760248201527f5469636b657420616c72656164792072656465656d6564000000000000000000604482015260640161033a565b6080860151604080516020810187905201604051602081830303815290604052805190602001201461137e5760405162461bcd60e51b815260206004820152602f60248201527f48617368206f662073656e64657252616e6420646f65736e2774206d6174636860448201527f2073656e64657252616e64486173680000000000000000000000000000000000606482015260840161033a565b60a086015160408051602081018690520160405160208183030381529060405280519060200120146114185760405162461bcd60e51b815260206004820152603360248201527f48617368206f662072656465656d657252616e6420646f65736e2774206d617460448201527f63682072656465656d657252616e644861736800000000000000000000000000606482015260840161033a565b6114278287602001518761175c565b6114995760405162461bcd60e51b815260206004820152602560248201527f5469636b657420646f65736e2774206861766520612076616c6964207369676e60448201527f6174757265000000000000000000000000000000000000000000000000000000606482015260840161033a565b60006114a5878361082a565b90506114b2838583611784565b6114fe5760405162461bcd60e51b815260206004820152601660248201527f5469636b6574206973206e6f7420612077696e6e657200000000000000000000604482015260640161033a565b50505050505050565b600061152b82602001516001600160a01b03166000908152606c6020526040902090565b905080600001548360a0015111156115ec5761155082604001518283600001546117ea565b606554600182015460405163a9059cbb60e01b815261dead600482015260248101919091526001600160a01b039091169063a9059cbb90604401602060405180830381600087803b1580156115a457600080fd5b505af11580156115b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115dc9190611c7d565b5060008082556001820155505050565b6115ff8260400151828560a001516117ea565b505050565b600054610100900460ff168061161d575060005460ff16155b6116805760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161033a565b600054610100900460ff1615801561111d576000805461ffff1916610101179055801561106c576000805461ff001916905550565b600054610100900460ff16806116ce575060005460ff16155b6117315760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161033a565b600054610100900460ff16158015611753576000805461ffff19166101011790555b61111d33611131565b6000826001600160a01b03166117728386611974565b6001600160a01b031614949350505050565b600080826fffffffffffffffffffffffffffffffff166080846fffffffffffffffffffffffffffffffff16901b1790508085856040516020016117c8929190612070565b60408051601f1981840301815291905280516020909101201095945050505050565b81548111156118615760405162461bcd60e51b815260206004820152603260248201527f5370656e64657220646f6573206e6f74206861766520656e6f75676820746f2060448201527f7472616e7366657220746f207265776172640000000000000000000000000000606482015260840161033a565b815461186e9082906121e5565b825560655460695460405163a9059cbb60e01b81526001600160a01b0391821660048201526024810184905291169063a9059cbb90604401602060405180830381600087803b1580156118c057600080fd5b505af11580156118d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118f89190611c7d565b506069546040517f40768fd30000000000000000000000000000000000000000000000000000000081526001600160a01b03858116600483015260248201849052909116906340768fd390604401600060405180830381600087803b15801561196057600080fd5b505af11580156114fe573d6000803e3d6000fd5b600081516041146119c75760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161033a565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115611a545760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161033a565b601b8160ff161015611a6e57611a6b601b82612150565b90505b8060ff16601b14158015611a8657508060ff16601c14155b15611ade5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b606482015260840161033a565b6040805160008082526020820180845289905260ff841692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015611b32573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661089c5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161033a565b80518015158114611ba557600080fd5b919050565b600060c08284031215611bbb578081fd5b60405160c0810181811067ffffffffffffffff82111715611bde57611bde61223e565b604052823581529050806020830135611bf681612254565b60208201526040830135611c0981612254565b80604083015250606083013560608201526080830135608082015260a083013560a08201525092915050565b8035611ba581612269565b8051611ba581612269565b8035611ba581612287565b8051611ba581612287565b600060208284031215611c72578081fd5b81356111ca81612254565b600060208284031215611c8e578081fd5b6111ca82611b95565b600060208284031215611ca8578081fd5b5035919050565b600080600080600080600060e0888a031215611cc9578283fd5b8735611cd481612254565b96506020880135611ce481612254565b95506040880135611cf481612254565b94506060880135611d0481612254565b93506080880135611d1481612254565b925060a0880135611d2481612254565b8092505060c0880135905092959891949750929550565b60006101408284031215611d4d578081fd5b611d55612092565b82518152602083015160208201526040830151604082015260608301516060820152611d8360808401611c56565b608082015260a083015160a0820152611d9e60c08401611c40565b60c0820152611daf60e08401611c40565b60e08201526101008381015190820152610120611dcd818501611c56565b908201529392505050565b60006020808385031215611dea578182fd5b825167ffffffffffffffff80821115611e01578384fd5b9084019060808287031215611e14578384fd5b611e1c6120bc565b825182811115611e2a578586fd5b83019150601f82018713611e3c578485fd5b8151611e4f611e4a82612110565b6120df565b8181528886838601011115611e62578687fd5b611e71828783018887016121fc565b825250611e7f838501611c56565b8482015260408301516040820152611e9960608401611b95565b60608201529695505050505050565b600060c08284031215611eb9578081fd5b6111ca8383611baa565b600080828403610200811215611ed7578283fd5b611ee18585611baa565b92506101408060bf1983011215611ef6578283fd5b611efe612092565b915060c0850135825260e08501356020830152610100808601356040840152610120808701356060850152611f34838801611c4b565b608085015261016087013560a0850152611f516101808801611c35565b60c0850152611f636101a08801611c35565b60e08501526101c087013582850152611f7f6101e08801611c4b565b9084015250929590945092505050565b6000806000806101208587031215611fa5578182fd5b611faf8686611baa565b935060c0850135925060e0850135915061010085013567ffffffffffffffff811115611fd9578182fd5b8501601f81018713611fe9578182fd5b8035611ff7611e4a82612110565b81815288602083850101111561200b578384fd5b81602084016020830137908101602001929092525092959194509250565b60006020828403121561203a578081fd5b5051919050565b60008060408385031215612053578182fd5b82359150602083013561206581612254565b809150509250929050565b600083516120828184602088016121fc565b9190910191825250602001919050565b604051610140810167ffffffffffffffff811182821017156120b6576120b661223e565b60405290565b6040516080810167ffffffffffffffff811182821017156120b6576120b661223e565b604051601f8201601f1916810167ffffffffffffffff811182821017156121085761210861223e565b604052919050565b600067ffffffffffffffff82111561212a5761212a61223e565b50601f01601f191660200190565b6000821982111561214b5761214b612228565b500190565b600060ff821660ff84168060ff0382111561216d5761216d612228565b019392505050565b60008261219057634e487b7160e01b81526012600452602481fd5b500490565b60008160001904831182151516156121af576121af612228565b500290565b60006fffffffffffffffffffffffffffffffff838116908316818110156121dd576121dd612228565b039392505050565b6000828210156121f7576121f7612228565b500390565b60005b838110156122175781810151838201526020016121ff565b8381111561052f5750506000910152565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461106c57600080fd5b6fffffffffffffffffffffffffffffffff8116811461106c57600080fd5b61ffff8116811461106c57600080fdfea2646970667358221220ca5043333a6a99500b8d48d1921227c0eeb0a9000554c45377b3de9bfa2a58c164736f6c63430008040033",
}

// SyloTicketingABI is the input ABI used to generate the binding from.
// Deprecated: Use SyloTicketingMetaData.ABI instead.
var SyloTicketingABI = SyloTicketingMetaData.ABI

// SyloTicketingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SyloTicketingMetaData.Bin instead.
var SyloTicketingBin = SyloTicketingMetaData.Bin

// DeploySyloTicketing deploys a new Ethereum contract, binding an instance of SyloTicketing to it.
func DeploySyloTicketing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SyloTicketing, error) {
	parsed, err := SyloTicketingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SyloTicketingBin), backend)
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

// CalculateWinningProbability is a free data retrieval call binding the contract method 0xa57c4bf9.
//
// Solidity: function calculateWinningProbability((uint256,address,address,uint256,bytes32,bytes32) ticket, (uint256,uint256,uint256,uint256,uint16,uint256,uint128,uint128,uint256,uint16) epoch) view returns(uint128)
func (_SyloTicketing *SyloTicketingCaller) CalculateWinningProbability(opts *bind.CallOpts, ticket SyloTicketingTicket, epoch EpochsManagerEpoch) (*big.Int, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "calculateWinningProbability", ticket, epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateWinningProbability is a free data retrieval call binding the contract method 0xa57c4bf9.
//
// Solidity: function calculateWinningProbability((uint256,address,address,uint256,bytes32,bytes32) ticket, (uint256,uint256,uint256,uint256,uint16,uint256,uint128,uint128,uint256,uint16) epoch) view returns(uint128)
func (_SyloTicketing *SyloTicketingSession) CalculateWinningProbability(ticket SyloTicketingTicket, epoch EpochsManagerEpoch) (*big.Int, error) {
	return _SyloTicketing.Contract.CalculateWinningProbability(&_SyloTicketing.CallOpts, ticket, epoch)
}

// CalculateWinningProbability is a free data retrieval call binding the contract method 0xa57c4bf9.
//
// Solidity: function calculateWinningProbability((uint256,address,address,uint256,bytes32,bytes32) ticket, (uint256,uint256,uint256,uint256,uint16,uint256,uint128,uint128,uint256,uint16) epoch) view returns(uint128)
func (_SyloTicketing *SyloTicketingCallerSession) CalculateWinningProbability(ticket SyloTicketingTicket, epoch EpochsManagerEpoch) (*big.Int, error) {
	return _SyloTicketing.Contract.CalculateWinningProbability(&_SyloTicketing.CallOpts, ticket, epoch)
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

// GetTicketHash is a free data retrieval call binding the contract method 0x7761134a.
//
// Solidity: function getTicketHash((uint256,address,address,uint256,bytes32,bytes32) ticket) pure returns(bytes32)
func (_SyloTicketing *SyloTicketingCaller) GetTicketHash(opts *bind.CallOpts, ticket SyloTicketingTicket) ([32]byte, error) {
	var out []interface{}
	err := _SyloTicketing.contract.Call(opts, &out, "getTicketHash", ticket)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTicketHash is a free data retrieval call binding the contract method 0x7761134a.
//
// Solidity: function getTicketHash((uint256,address,address,uint256,bytes32,bytes32) ticket) pure returns(bytes32)
func (_SyloTicketing *SyloTicketingSession) GetTicketHash(ticket SyloTicketingTicket) ([32]byte, error) {
	return _SyloTicketing.Contract.GetTicketHash(&_SyloTicketing.CallOpts, ticket)
}

// GetTicketHash is a free data retrieval call binding the contract method 0x7761134a.
//
// Solidity: function getTicketHash((uint256,address,address,uint256,bytes32,bytes32) ticket) pure returns(bytes32)
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

// Initialize is a paid mutator transaction binding the contract method 0x1460e390.
//
// Solidity: function initialize(address token, address listings, address stakingManager, address directory, address epochsManager, address rewardsManager, uint256 _unlockDuration) returns()
func (_SyloTicketing *SyloTicketingTransactor) Initialize(opts *bind.TransactOpts, token common.Address, listings common.Address, stakingManager common.Address, directory common.Address, epochsManager common.Address, rewardsManager common.Address, _unlockDuration *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "initialize", token, listings, stakingManager, directory, epochsManager, rewardsManager, _unlockDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0x1460e390.
//
// Solidity: function initialize(address token, address listings, address stakingManager, address directory, address epochsManager, address rewardsManager, uint256 _unlockDuration) returns()
func (_SyloTicketing *SyloTicketingSession) Initialize(token common.Address, listings common.Address, stakingManager common.Address, directory common.Address, epochsManager common.Address, rewardsManager common.Address, _unlockDuration *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.Initialize(&_SyloTicketing.TransactOpts, token, listings, stakingManager, directory, epochsManager, rewardsManager, _unlockDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0x1460e390.
//
// Solidity: function initialize(address token, address listings, address stakingManager, address directory, address epochsManager, address rewardsManager, uint256 _unlockDuration) returns()
func (_SyloTicketing *SyloTicketingTransactorSession) Initialize(token common.Address, listings common.Address, stakingManager common.Address, directory common.Address, epochsManager common.Address, rewardsManager common.Address, _unlockDuration *big.Int) (*types.Transaction, error) {
	return _SyloTicketing.Contract.Initialize(&_SyloTicketing.TransactOpts, token, listings, stakingManager, directory, epochsManager, rewardsManager, _unlockDuration)
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

// Redeem is a paid mutator transaction binding the contract method 0xd8024d32.
//
// Solidity: function redeem((uint256,address,address,uint256,bytes32,bytes32) ticket, uint256 senderRand, uint256 redeemerRand, bytes sig) returns()
func (_SyloTicketing *SyloTicketingTransactor) Redeem(opts *bind.TransactOpts, ticket SyloTicketingTicket, senderRand *big.Int, redeemerRand *big.Int, sig []byte) (*types.Transaction, error) {
	return _SyloTicketing.contract.Transact(opts, "redeem", ticket, senderRand, redeemerRand, sig)
}

// Redeem is a paid mutator transaction binding the contract method 0xd8024d32.
//
// Solidity: function redeem((uint256,address,address,uint256,bytes32,bytes32) ticket, uint256 senderRand, uint256 redeemerRand, bytes sig) returns()
func (_SyloTicketing *SyloTicketingSession) Redeem(ticket SyloTicketingTicket, senderRand *big.Int, redeemerRand *big.Int, sig []byte) (*types.Transaction, error) {
	return _SyloTicketing.Contract.Redeem(&_SyloTicketing.TransactOpts, ticket, senderRand, redeemerRand, sig)
}

// Redeem is a paid mutator transaction binding the contract method 0xd8024d32.
//
// Solidity: function redeem((uint256,address,address,uint256,bytes32,bytes32) ticket, uint256 senderRand, uint256 redeemerRand, bytes sig) returns()
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
