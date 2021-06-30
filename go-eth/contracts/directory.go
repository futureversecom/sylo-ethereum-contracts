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

// DirectoryStake is an auto generated low-level Go binding around an user-defined struct.
type DirectoryStake struct {
	Amount      *big.Int
	LeftAmount  *big.Int
	RightAmount *big.Int
	Stakee      common.Address
	Parent      DirectoryStakePointer
	Left        DirectoryStakePointer
	Right       DirectoryStakePointer
}

// DirectoryStakePointer is an auto generated low-level Go binding around an user-defined struct.
type DirectoryStakePointer struct {
	Value [32]byte
}

// DirectoryABI is the input ABI used to generate the binding from.
const DirectoryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rightAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"value_\",\"type\":\"bytes32\"}],\"internalType\":\"structDirectory.StakePointer\",\"name\":\"parent\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"value_\",\"type\":\"bytes32\"}],\"internalType\":\"structDirectory.StakePointer\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"value_\",\"type\":\"bytes32\"}],\"internalType\":\"structDirectory.StakePointer\",\"name\":\"right\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"unlockings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unlockDuration\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newUnlockDuration\",\"type\":\"uint256\"}],\"name\":\"setUnlockDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"unlockStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"cancelUnlocking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"point\",\"type\":\"uint128\"}],\"name\":\"scan\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getStake\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rightAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"value_\",\"type\":\"bytes32\"}],\"internalType\":\"structDirectory.StakePointer\",\"name\":\"parent\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"value_\",\"type\":\"bytes32\"}],\"internalType\":\"structDirectory.StakePointer\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"value_\",\"type\":\"bytes32\"}],\"internalType\":\"structDirectory.StakePointer\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structDirectory.Stake\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DirectoryBin is the compiled bytecode used for deploying new contracts.
var DirectoryBin = "0x608060405234801561001057600080fd5b50611d20806100206000396000f3fe608060405234801561001057600080fd5b50600436106101515760003560e01c80638da5cb5b116100cd578063c23a5cea11610081578063cd6dc68711610066578063cd6dc6871461041c578063dd9007691461042f578063f2fde38b1461044257600080fd5b8063c23a5cea146103f6578063c8b6cbf71461040957600080fd5b80639341a536116100b25780639341a5361461034f578063a859f1721461038b578063bc2985531461039e57600080fd5b80638da5cb5b1461028a5780638fee64071461029b57600080fd5b8063715018a6116101245780637bc74225116101095780637bc742251461020457806382dda22d1461020c5780638a1fcd601461028157600080fd5b8063715018a6146101dc57806379193610146101e457600080fd5b806323314c6c146101565780632d49aa1c1461016b5780634e6806611461017e5780636b5537e2146101ae575b600080fd5b610169610164366004611bdc565b610455565b005b610169610179366004611bdc565b6104c1565b61019161018c366004611bac565b61056e565b6040516001600160a01b0390911681526020015b60405180910390f35b6101ce6101bc366004611aee565b60696020526000908152604090205481565b6040519081526020016101a5565b610169610696565b6101f76101f2366004611aee565b61073a565b6040516101a59190611c00565b6101ce6107b0565b61021f61021a366004611b11565b6107f7565b6040516101a59190600060e0820190508251825260208301516020830152604083015160408301526001600160a01b036060840151166060830152608083015151608083015260a08301515160a083015260c08301515160c083015292915050565b6101ce60665481565b6033546001600160a01b0316610191565b61030f6102a9366004611b94565b606760209081526000918252604091829020805460018201546002830154600384015486518087018852600486015481528751808801895260058701548152885197880190985260069095015486529295919490936001600160a01b0390931692919087565b604080519788526020880196909652948601939093526001600160a01b0390911660608501525160808401525160a08301525160c082015260e0016101a5565b61037661035d366004611b94565b606a602052600090815260409020805460019091015482565b604080519283526020830191909152016101a5565b6101ce610399366004611bdc565b6108d9565b6101ce6103ac366004611b11565b6040516bffffffffffffffffffffffff19606084811b8216602084015283901b16603482015260009060480160405160208183030381529060405280519060200120905092915050565b610169610404366004611aee565b610e29565b610191610417366004611b49565b610feb565b61016961042a366004611b49565b611023565b61016961043d366004611b94565b6110ff565b610169610450366004611aee565b61115e565b60006104a633836040516bffffffffffffffffffffffff19606084811b8216602084015283901b16603482015260009060480160405160208183030381529060405280519060200120905092915050565b90506104b28184611290565b6104bc838361132e565b505050565b6104cb828261132e565b6065546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018490526001600160a01b03909116906323b872dd90606401602060405180830381600087803b15801561053657600080fd5b505af115801561054a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104bc9190611b74565b606b5460009061058057506000919050565b60006080836fffffffffffffffffffffffffffffffff1661059f6107b0565b6105a99190611c65565b606b54911c91505b600081815260676020526040902060018101548310156105f45760058101546105e957600301546001600160a01b0316949350505050565b6005015490506105b1565b60018101546106039084611c84565b815490935083101561062457600301546001600160a01b0316949350505050565b600681015461067a5760405162461bcd60e51b815260206004820152601960248201527f6d697373696e67206e6f6465206f6e207468652072696768740000000000000060448201526064015b60405180910390fd5b6006810154815490925061068e9084611c84565b9250506105b1565b6033546001600160a01b031633146106f05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610671565b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380546001600160a01b0319169055565b6001600160a01b0381166000908152606860209081526040918290208054835181840281018401909452808452606093928301828280156107a457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610786575b50505050509050919050565b606b546000906107c05750600090565b606b5460009081526067602052604090206002810154600182015482546107e79190611c4d565b6107f19190611c4d565b91505090565b6107ff611a7c565b6067600061085284866040516bffffffffffffffffffffffff19606084811b8216602084015283901b16603482015260009060480160405160208183030381529060405280519060200120905092915050565b81526020808201929092526040908101600020815160e0810183528154815260018201548185015260028201548184015260038201546001600160a01b0316606082015282518085018452600483015481526080820152825180850184526005830154815260a0820152825193840190925260060154825260c08101919091529392505050565b60008061092b33846040516bffffffffffffffffffffffff19606084811b8216602084015283901b16603482015260009060480160405160208183030381529060405280519060200120905092915050565b60008181526067602090815260409182902082519182019092528281528154929350909161099b5760405162461bcd60e51b815260206004820152601260248201527f4e6f7468696e6720746f20756e7374616b6500000000000000000000000000006044820152606401610671565b81548611156109ec5760405162461bcd60e51b815260206004820152601e60248201527f43616e6e6f7420756e6c6f636b206d6f7265207468616e207374616b656400006044820152606401610671565b8051610a0f90836109ff89600019611c84565b610a0a906001611c4d565b611597565b8154610dd05781600201548260010154610a299190611c4d565b610a6557600482015415610a5b57600482015460009081526067602052604090208151610a5691906115ce565b610c1f565b6000606b55610c1f565b60008260020154836001015411610a7f5782600601610a84565b826005015b60408051602080820183529254808252600090815260679093529120909150815b60008260020154836001015411610abf5782600601610ac4565b826005015b60408051602081019091529054808252909150610ae15750610afa565b8051600090815260676020526040902092509050610aa5565b604080516020810190915260048301548152606b548551849114610b3c57600487015460009081526067602052604090208351610b3a9082908b90611680565b505b6004808801549082015581518814610bdf57610b5d87846000015183611737565b610b6c87846000015183611775565b815160048801819055600090815260676020526040902083518751610b92929190611680565b610bbe88888360000154600019610ba99190611c84565b610bb4906001611c4d565b60048501546117b3565b815160009081526067602052604090208651610bda91906115ce565b610c0a565b825160058801541415610bfb57610bda87846000015183611737565b610c0a87846000015183611775565b6004810154610c19578251606b555b50505050505b600083815260676020908152604080832083815560018101849055600281018490556003810180546001600160a01b031916905560048101849055600581018490556006018390556001600160a01b038816835260689091528120905b815463ffffffff82161015610dcd57336001600160a01b0316828263ffffffff1681548110610cbb57634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b03161415610dbb5781548290610ce690600190611c84565b81548110610d0457634e487b7160e01b600052603260045260246000fd5b9060005260206000200160009054906101000a90046001600160a01b0316828263ffffffff1681548110610d4857634e487b7160e01b600052603260045260246000fd5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081805480610d9457634e487b7160e01b600052603160045260246000fd5b600082815260209020810160001990810180546001600160a01b0319169055019055610dcd565b80610dc581611c9b565b915050610c7c565b50505b6000838152606a60205260408120606654909190610dee9043611c4d565b90508082600101541015610e0457600182018190555b87826000016000828254610e189190611c4d565b909155509098975050505050505050565b6000610e7a33836040516bffffffffffffffffffffffff19606084811b8216602084015283901b16603482015260009060480160405160208183030381529060405280519060200120905092915050565b6000818152606a602052604090206001810154919250904311610edf5760405162461bcd60e51b815260206004820152601660248201527f5374616b65206e6f742079657420756e6c6f636b6564000000000000000000006044820152606401610671565b8054610f2d5760405162461bcd60e51b815260206004820152601560248201527f4e6f20616d6f756e7420746f20776974686472617700000000000000000000006044820152606401610671565b80546000838152606a60205260408082208281556001019190915560655490517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018390526001600160a01b039091169063a9059cbb90604401602060405180830381600087803b158015610fac57600080fd5b505af1158015610fc0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fe49190611b74565b5050505050565b6068602052816000526040600020818154811061100757600080fd5b6000918252602090912001546001600160a01b03169150829050565b600054610100900460ff168061103c575060005460ff16155b61109f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610671565b600054610100900460ff161580156110c1576000805461ffff19166101011790555b6110c9611813565b606580546001600160a01b0319166001600160a01b038516179055606682905580156104bc576000805461ff0019169055505050565b6033546001600160a01b031633146111595760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610671565b606655565b6033546001600160a01b031633146111b85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610671565b6001600160a01b0381166112345760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610671565b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380546001600160a01b0319166001600160a01b0392909216919091179055565b6000828152606a6020526040902080548214156112c05750506000908152606a6020526040812081815560010155565b805482106113105760405162461bcd60e51b815260206004820152601e60248201527f556e6c6f636b2068617320696e73756666696369656e7420616d6f756e7400006044820152606401610671565b818160000160008282546113249190611c84565b9091555050505050565b6001600160a01b0381166113845760405162461bcd60e51b815260206004820152600f60248201527f41646472657373206973206e756c6c00000000000000000000000000000000006044820152606401610671565b816113d15760405162461bcd60e51b815260206004820152601460248201527f43616e6e6f74207374616b65206e6f7468696e670000000000000000000000006044820152606401610671565b33600061142382846040516bffffffffffffffffffffffff19606084811b8216602084015283901b16603482015260009060480160405160208183030381529060405280519060200120905092915050565b6000818152606760205260409020805491925090611581576001600160a01b038416600090815260686020526040902054600a10156114ca5760405162461bcd60e51b815260206004820152602e60248201527f54686973206e6f6465206861732072656163686564206974732064656c65676160448201527f746564207374616b6572206361700000000000000000000000000000000000006064820152608401610671565b604080516020810190915260008152606b905b8154156115285750604080516020808201835283548083526000908152606790915291909120600281015460018201541061151b5780600601611520565b806005015b9250506114dd565b5160048301558290556003810180546001600160a01b03199081166001600160a01b0387811691821790935560009081526068602090815260408220805460018101825590835291200180549091169185169190911790555b600481015461159057606b8290555b610fe48282875b81548101825560038201546001600160a01b031660009081526069602052604081208054830190556104bc908490849084906117b3565b60058201548114156115de575050565b60068201548114156115f557600060068301555050565b600582015481148061160a5750600682015481145b61167c5760405162461bcd60e51b815260206004820152602f60248201527f4f6c64206368696c642063616e6e6f742062652072656d6f766564202d20697460448201527f20646f6573206e6f7420657869737400000000000000000000000000000000006064820152608401610671565b5050565b60058301548214156116985760058301819055505050565b60068301548214156116b05760068301819055505050565b60058301548214806116c55750600683015482145b6104bc5760405162461bcd60e51b815260206004820152602f60248201527f4f6c64206368696c642063616e6e6f74206265206368616e676564202d20697460448201527f20646f6573206e6f7420657869737400000000000000000000000000000000006064820152608401610671565b600683015461174557505050565b60068084018054600090815260676020526040902060040193909355915491810191909155600291820154910155565b600583015461178357505050565b60058084018054600090815260676020526040902060040193909355915491810191909155600191820154910155565b6004830180548214156117c6575061180d565b8054600090815260676020526040902060058101548614156117f157600181018054850190556117fc565b600281018054850190555b815461180a908286866117b3565b50505b50505050565b600054610100900460ff168061182c575060005460ff16155b61188f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610671565b600054610100900460ff161580156118b1576000805461ffff19166101011790555b6118b96118d6565b6118c1611987565b80156118d3576000805461ff00191690555b50565b600054610100900460ff16806118ef575060005460ff16155b6119525760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610671565b600054610100900460ff161580156118c1576000805461ffff191661010117905580156118d3576000805461ff001916905550565b600054610100900460ff16806119a0575060005460ff16155b611a035760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610671565b600054610100900460ff16158015611a25576000805461ffff19166101011790555b603380546001600160a01b0319163390811790915560405181906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35080156118d3576000805461ff001916905550565b6040518060e0016040528060008152602001600081526020016000815260200160006001600160a01b03168152602001611ac56040518060200160405280600080191681525090565b815260408051602080820183526000808352818501929092528251908101835290815291015290565b600060208284031215611aff578081fd5b8135611b0a81611cd5565b9392505050565b60008060408385031215611b23578081fd5b8235611b2e81611cd5565b91506020830135611b3e81611cd5565b809150509250929050565b60008060408385031215611b5b578182fd5b8235611b6681611cd5565b946020939093013593505050565b600060208284031215611b85578081fd5b81518015158114611b0a578182fd5b600060208284031215611ba5578081fd5b5035919050565b600060208284031215611bbd578081fd5b81356fffffffffffffffffffffffffffffffff81168114611b0a578182fd5b60008060408385031215611bee578182fd5b823591506020830135611b3e81611cd5565b6020808252825182820181905260009190848201906040850190845b81811015611c415783516001600160a01b031683529284019291840191600101611c1c565b50909695505050505050565b60008219821115611c6057611c60611cbf565b500190565b6000816000190483118215151615611c7f57611c7f611cbf565b500290565b600082821015611c9657611c96611cbf565b500390565b600063ffffffff80831681811415611cb557611cb5611cbf565b6001019392505050565b634e487b7160e01b600052601160045260246000fd5b6001600160a01b03811681146118d357600080fdfea2646970667358221220437696facd8334a6cd47d82d8cb2c0619fb400038f8702ad10d5b670724461cd64736f6c63430008040033"

// DeployDirectory deploys a new Ethereum contract, binding an instance of Directory to it.
func DeployDirectory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Directory, error) {
	parsed, err := abi.JSON(strings.NewReader(DirectoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DirectoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Directory{DirectoryCaller: DirectoryCaller{contract: contract}, DirectoryTransactor: DirectoryTransactor{contract: contract}, DirectoryFilterer: DirectoryFilterer{contract: contract}}, nil
}

// Directory is an auto generated Go binding around an Ethereum contract.
type Directory struct {
	DirectoryCaller     // Read-only binding to the contract
	DirectoryTransactor // Write-only binding to the contract
	DirectoryFilterer   // Log filterer for contract events
}

// DirectoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DirectoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DirectoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DirectoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DirectoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DirectoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DirectorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DirectorySession struct {
	Contract     *Directory        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DirectoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DirectoryCallerSession struct {
	Contract *DirectoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DirectoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DirectoryTransactorSession struct {
	Contract     *DirectoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DirectoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DirectoryRaw struct {
	Contract *Directory // Generic contract binding to access the raw methods on
}

// DirectoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DirectoryCallerRaw struct {
	Contract *DirectoryCaller // Generic read-only contract binding to access the raw methods on
}

// DirectoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DirectoryTransactorRaw struct {
	Contract *DirectoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDirectory creates a new instance of Directory, bound to a specific deployed contract.
func NewDirectory(address common.Address, backend bind.ContractBackend) (*Directory, error) {
	contract, err := bindDirectory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Directory{DirectoryCaller: DirectoryCaller{contract: contract}, DirectoryTransactor: DirectoryTransactor{contract: contract}, DirectoryFilterer: DirectoryFilterer{contract: contract}}, nil
}

// NewDirectoryCaller creates a new read-only instance of Directory, bound to a specific deployed contract.
func NewDirectoryCaller(address common.Address, caller bind.ContractCaller) (*DirectoryCaller, error) {
	contract, err := bindDirectory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DirectoryCaller{contract: contract}, nil
}

// NewDirectoryTransactor creates a new write-only instance of Directory, bound to a specific deployed contract.
func NewDirectoryTransactor(address common.Address, transactor bind.ContractTransactor) (*DirectoryTransactor, error) {
	contract, err := bindDirectory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DirectoryTransactor{contract: contract}, nil
}

// NewDirectoryFilterer creates a new log filterer instance of Directory, bound to a specific deployed contract.
func NewDirectoryFilterer(address common.Address, filterer bind.ContractFilterer) (*DirectoryFilterer, error) {
	contract, err := bindDirectory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DirectoryFilterer{contract: contract}, nil
}

// bindDirectory binds a generic wrapper to an already deployed contract.
func bindDirectory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DirectoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Directory *DirectoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Directory.Contract.DirectoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Directory *DirectoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Directory.Contract.DirectoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Directory *DirectoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Directory.Contract.DirectoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Directory *DirectoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Directory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Directory *DirectoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Directory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Directory *DirectoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Directory.Contract.contract.Transact(opts, method, params...)
}

// GetKey is a free data retrieval call binding the contract method 0xbc298553.
//
// Solidity: function getKey(address staker, address stakee) pure returns(bytes32)
func (_Directory *DirectoryCaller) GetKey(opts *bind.CallOpts, staker common.Address, stakee common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "getKey", staker, stakee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKey is a free data retrieval call binding the contract method 0xbc298553.
//
// Solidity: function getKey(address staker, address stakee) pure returns(bytes32)
func (_Directory *DirectorySession) GetKey(staker common.Address, stakee common.Address) ([32]byte, error) {
	return _Directory.Contract.GetKey(&_Directory.CallOpts, staker, stakee)
}

// GetKey is a free data retrieval call binding the contract method 0xbc298553.
//
// Solidity: function getKey(address staker, address stakee) pure returns(bytes32)
func (_Directory *DirectoryCallerSession) GetKey(staker common.Address, stakee common.Address) ([32]byte, error) {
	return _Directory.Contract.GetKey(&_Directory.CallOpts, staker, stakee)
}

// GetStake is a free data retrieval call binding the contract method 0x82dda22d.
//
// Solidity: function getStake(address stakee, address staker) view returns((uint256,uint256,uint256,address,(bytes32),(bytes32),(bytes32)))
func (_Directory *DirectoryCaller) GetStake(opts *bind.CallOpts, stakee common.Address, staker common.Address) (DirectoryStake, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "getStake", stakee, staker)

	if err != nil {
		return *new(DirectoryStake), err
	}

	out0 := *abi.ConvertType(out[0], new(DirectoryStake)).(*DirectoryStake)

	return out0, err

}

// GetStake is a free data retrieval call binding the contract method 0x82dda22d.
//
// Solidity: function getStake(address stakee, address staker) view returns((uint256,uint256,uint256,address,(bytes32),(bytes32),(bytes32)))
func (_Directory *DirectorySession) GetStake(stakee common.Address, staker common.Address) (DirectoryStake, error) {
	return _Directory.Contract.GetStake(&_Directory.CallOpts, stakee, staker)
}

// GetStake is a free data retrieval call binding the contract method 0x82dda22d.
//
// Solidity: function getStake(address stakee, address staker) view returns((uint256,uint256,uint256,address,(bytes32),(bytes32),(bytes32)))
func (_Directory *DirectoryCallerSession) GetStake(stakee common.Address, staker common.Address) (DirectoryStake, error) {
	return _Directory.Contract.GetStake(&_Directory.CallOpts, stakee, staker)
}

// GetStakers is a free data retrieval call binding the contract method 0x79193610.
//
// Solidity: function getStakers(address stakee) view returns(address[])
func (_Directory *DirectoryCaller) GetStakers(opts *bind.CallOpts, stakee common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "getStakers", stakee)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStakers is a free data retrieval call binding the contract method 0x79193610.
//
// Solidity: function getStakers(address stakee) view returns(address[])
func (_Directory *DirectorySession) GetStakers(stakee common.Address) ([]common.Address, error) {
	return _Directory.Contract.GetStakers(&_Directory.CallOpts, stakee)
}

// GetStakers is a free data retrieval call binding the contract method 0x79193610.
//
// Solidity: function getStakers(address stakee) view returns(address[])
func (_Directory *DirectoryCallerSession) GetStakers(stakee common.Address) ([]common.Address, error) {
	return _Directory.Contract.GetStakers(&_Directory.CallOpts, stakee)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() view returns(uint256)
func (_Directory *DirectoryCaller) GetTotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "getTotalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() view returns(uint256)
func (_Directory *DirectorySession) GetTotalStake() (*big.Int, error) {
	return _Directory.Contract.GetTotalStake(&_Directory.CallOpts)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() view returns(uint256)
func (_Directory *DirectoryCallerSession) GetTotalStake() (*big.Int, error) {
	return _Directory.Contract.GetTotalStake(&_Directory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Directory *DirectoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Directory *DirectorySession) Owner() (common.Address, error) {
	return _Directory.Contract.Owner(&_Directory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Directory *DirectoryCallerSession) Owner() (common.Address, error) {
	return _Directory.Contract.Owner(&_Directory.CallOpts)
}

// Scan is a free data retrieval call binding the contract method 0x4e680661.
//
// Solidity: function scan(uint128 point) view returns(address)
func (_Directory *DirectoryCaller) Scan(opts *bind.CallOpts, point *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "scan", point)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Scan is a free data retrieval call binding the contract method 0x4e680661.
//
// Solidity: function scan(uint128 point) view returns(address)
func (_Directory *DirectorySession) Scan(point *big.Int) (common.Address, error) {
	return _Directory.Contract.Scan(&_Directory.CallOpts, point)
}

// Scan is a free data retrieval call binding the contract method 0x4e680661.
//
// Solidity: function scan(uint128 point) view returns(address)
func (_Directory *DirectoryCallerSession) Scan(point *big.Int) (common.Address, error) {
	return _Directory.Contract.Scan(&_Directory.CallOpts, point)
}

// Stakees is a free data retrieval call binding the contract method 0x6b5537e2.
//
// Solidity: function stakees(address ) view returns(uint256)
func (_Directory *DirectoryCaller) Stakees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "stakees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Stakees is a free data retrieval call binding the contract method 0x6b5537e2.
//
// Solidity: function stakees(address ) view returns(uint256)
func (_Directory *DirectorySession) Stakees(arg0 common.Address) (*big.Int, error) {
	return _Directory.Contract.Stakees(&_Directory.CallOpts, arg0)
}

// Stakees is a free data retrieval call binding the contract method 0x6b5537e2.
//
// Solidity: function stakees(address ) view returns(uint256)
func (_Directory *DirectoryCallerSession) Stakees(arg0 common.Address) (*big.Int, error) {
	return _Directory.Contract.Stakees(&_Directory.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xc8b6cbf7.
//
// Solidity: function stakers(address , uint256 ) view returns(address)
func (_Directory *DirectoryCaller) Stakers(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "stakers", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Stakers is a free data retrieval call binding the contract method 0xc8b6cbf7.
//
// Solidity: function stakers(address , uint256 ) view returns(address)
func (_Directory *DirectorySession) Stakers(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Directory.Contract.Stakers(&_Directory.CallOpts, arg0, arg1)
}

// Stakers is a free data retrieval call binding the contract method 0xc8b6cbf7.
//
// Solidity: function stakers(address , uint256 ) view returns(address)
func (_Directory *DirectoryCallerSession) Stakers(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Directory.Contract.Stakers(&_Directory.CallOpts, arg0, arg1)
}

// Stakes is a free data retrieval call binding the contract method 0x8fee6407.
//
// Solidity: function stakes(bytes32 ) view returns(uint256 amount, uint256 leftAmount, uint256 rightAmount, address stakee, (bytes32) parent, (bytes32) left, (bytes32) right)
func (_Directory *DirectoryCaller) Stakes(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Amount      *big.Int
	LeftAmount  *big.Int
	RightAmount *big.Int
	Stakee      common.Address
	Parent      DirectoryStakePointer
	Left        DirectoryStakePointer
	Right       DirectoryStakePointer
}, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "stakes", arg0)

	outstruct := new(struct {
		Amount      *big.Int
		LeftAmount  *big.Int
		RightAmount *big.Int
		Stakee      common.Address
		Parent      DirectoryStakePointer
		Left        DirectoryStakePointer
		Right       DirectoryStakePointer
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LeftAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RightAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Stakee = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Parent = *abi.ConvertType(out[4], new(DirectoryStakePointer)).(*DirectoryStakePointer)
	outstruct.Left = *abi.ConvertType(out[5], new(DirectoryStakePointer)).(*DirectoryStakePointer)
	outstruct.Right = *abi.ConvertType(out[6], new(DirectoryStakePointer)).(*DirectoryStakePointer)

	return *outstruct, err

}

// Stakes is a free data retrieval call binding the contract method 0x8fee6407.
//
// Solidity: function stakes(bytes32 ) view returns(uint256 amount, uint256 leftAmount, uint256 rightAmount, address stakee, (bytes32) parent, (bytes32) left, (bytes32) right)
func (_Directory *DirectorySession) Stakes(arg0 [32]byte) (struct {
	Amount      *big.Int
	LeftAmount  *big.Int
	RightAmount *big.Int
	Stakee      common.Address
	Parent      DirectoryStakePointer
	Left        DirectoryStakePointer
	Right       DirectoryStakePointer
}, error) {
	return _Directory.Contract.Stakes(&_Directory.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x8fee6407.
//
// Solidity: function stakes(bytes32 ) view returns(uint256 amount, uint256 leftAmount, uint256 rightAmount, address stakee, (bytes32) parent, (bytes32) left, (bytes32) right)
func (_Directory *DirectoryCallerSession) Stakes(arg0 [32]byte) (struct {
	Amount      *big.Int
	LeftAmount  *big.Int
	RightAmount *big.Int
	Stakee      common.Address
	Parent      DirectoryStakePointer
	Left        DirectoryStakePointer
	Right       DirectoryStakePointer
}, error) {
	return _Directory.Contract.Stakes(&_Directory.CallOpts, arg0)
}

// UnlockDuration is a free data retrieval call binding the contract method 0x8a1fcd60.
//
// Solidity: function unlockDuration() view returns(uint256)
func (_Directory *DirectoryCaller) UnlockDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "unlockDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockDuration is a free data retrieval call binding the contract method 0x8a1fcd60.
//
// Solidity: function unlockDuration() view returns(uint256)
func (_Directory *DirectorySession) UnlockDuration() (*big.Int, error) {
	return _Directory.Contract.UnlockDuration(&_Directory.CallOpts)
}

// UnlockDuration is a free data retrieval call binding the contract method 0x8a1fcd60.
//
// Solidity: function unlockDuration() view returns(uint256)
func (_Directory *DirectoryCallerSession) UnlockDuration() (*big.Int, error) {
	return _Directory.Contract.UnlockDuration(&_Directory.CallOpts)
}

// Unlockings is a free data retrieval call binding the contract method 0x9341a536.
//
// Solidity: function unlockings(bytes32 ) view returns(uint256 amount, uint256 unlockAt)
func (_Directory *DirectoryCaller) Unlockings(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Amount   *big.Int
	UnlockAt *big.Int
}, error) {
	var out []interface{}
	err := _Directory.contract.Call(opts, &out, "unlockings", arg0)

	outstruct := new(struct {
		Amount   *big.Int
		UnlockAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnlockAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Unlockings is a free data retrieval call binding the contract method 0x9341a536.
//
// Solidity: function unlockings(bytes32 ) view returns(uint256 amount, uint256 unlockAt)
func (_Directory *DirectorySession) Unlockings(arg0 [32]byte) (struct {
	Amount   *big.Int
	UnlockAt *big.Int
}, error) {
	return _Directory.Contract.Unlockings(&_Directory.CallOpts, arg0)
}

// Unlockings is a free data retrieval call binding the contract method 0x9341a536.
//
// Solidity: function unlockings(bytes32 ) view returns(uint256 amount, uint256 unlockAt)
func (_Directory *DirectoryCallerSession) Unlockings(arg0 [32]byte) (struct {
	Amount   *big.Int
	UnlockAt *big.Int
}, error) {
	return _Directory.Contract.Unlockings(&_Directory.CallOpts, arg0)
}

// AddStake is a paid mutator transaction binding the contract method 0x2d49aa1c.
//
// Solidity: function addStake(uint256 amount, address stakee) returns()
func (_Directory *DirectoryTransactor) AddStake(opts *bind.TransactOpts, amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "addStake", amount, stakee)
}

// AddStake is a paid mutator transaction binding the contract method 0x2d49aa1c.
//
// Solidity: function addStake(uint256 amount, address stakee) returns()
func (_Directory *DirectorySession) AddStake(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.AddStake(&_Directory.TransactOpts, amount, stakee)
}

// AddStake is a paid mutator transaction binding the contract method 0x2d49aa1c.
//
// Solidity: function addStake(uint256 amount, address stakee) returns()
func (_Directory *DirectoryTransactorSession) AddStake(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.AddStake(&_Directory.TransactOpts, amount, stakee)
}

// CancelUnlocking is a paid mutator transaction binding the contract method 0x23314c6c.
//
// Solidity: function cancelUnlocking(uint256 amount, address stakee) returns()
func (_Directory *DirectoryTransactor) CancelUnlocking(opts *bind.TransactOpts, amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "cancelUnlocking", amount, stakee)
}

// CancelUnlocking is a paid mutator transaction binding the contract method 0x23314c6c.
//
// Solidity: function cancelUnlocking(uint256 amount, address stakee) returns()
func (_Directory *DirectorySession) CancelUnlocking(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.CancelUnlocking(&_Directory.TransactOpts, amount, stakee)
}

// CancelUnlocking is a paid mutator transaction binding the contract method 0x23314c6c.
//
// Solidity: function cancelUnlocking(uint256 amount, address stakee) returns()
func (_Directory *DirectoryTransactorSession) CancelUnlocking(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.CancelUnlocking(&_Directory.TransactOpts, amount, stakee)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address token, uint256 _unlockDuration) returns()
func (_Directory *DirectoryTransactor) Initialize(opts *bind.TransactOpts, token common.Address, _unlockDuration *big.Int) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "initialize", token, _unlockDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address token, uint256 _unlockDuration) returns()
func (_Directory *DirectorySession) Initialize(token common.Address, _unlockDuration *big.Int) (*types.Transaction, error) {
	return _Directory.Contract.Initialize(&_Directory.TransactOpts, token, _unlockDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address token, uint256 _unlockDuration) returns()
func (_Directory *DirectoryTransactorSession) Initialize(token common.Address, _unlockDuration *big.Int) (*types.Transaction, error) {
	return _Directory.Contract.Initialize(&_Directory.TransactOpts, token, _unlockDuration)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Directory *DirectoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Directory *DirectorySession) RenounceOwnership() (*types.Transaction, error) {
	return _Directory.Contract.RenounceOwnership(&_Directory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Directory *DirectoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Directory.Contract.RenounceOwnership(&_Directory.TransactOpts)
}

// SetUnlockDuration is a paid mutator transaction binding the contract method 0xdd900769.
//
// Solidity: function setUnlockDuration(uint256 newUnlockDuration) returns()
func (_Directory *DirectoryTransactor) SetUnlockDuration(opts *bind.TransactOpts, newUnlockDuration *big.Int) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "setUnlockDuration", newUnlockDuration)
}

// SetUnlockDuration is a paid mutator transaction binding the contract method 0xdd900769.
//
// Solidity: function setUnlockDuration(uint256 newUnlockDuration) returns()
func (_Directory *DirectorySession) SetUnlockDuration(newUnlockDuration *big.Int) (*types.Transaction, error) {
	return _Directory.Contract.SetUnlockDuration(&_Directory.TransactOpts, newUnlockDuration)
}

// SetUnlockDuration is a paid mutator transaction binding the contract method 0xdd900769.
//
// Solidity: function setUnlockDuration(uint256 newUnlockDuration) returns()
func (_Directory *DirectoryTransactorSession) SetUnlockDuration(newUnlockDuration *big.Int) (*types.Transaction, error) {
	return _Directory.Contract.SetUnlockDuration(&_Directory.TransactOpts, newUnlockDuration)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Directory *DirectoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Directory *DirectorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Directory.Contract.TransferOwnership(&_Directory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Directory *DirectoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Directory.Contract.TransferOwnership(&_Directory.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xa859f172.
//
// Solidity: function unlockStake(uint256 amount, address stakee) returns(uint256)
func (_Directory *DirectoryTransactor) UnlockStake(opts *bind.TransactOpts, amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "unlockStake", amount, stakee)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xa859f172.
//
// Solidity: function unlockStake(uint256 amount, address stakee) returns(uint256)
func (_Directory *DirectorySession) UnlockStake(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.UnlockStake(&_Directory.TransactOpts, amount, stakee)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xa859f172.
//
// Solidity: function unlockStake(uint256 amount, address stakee) returns(uint256)
func (_Directory *DirectoryTransactorSession) UnlockStake(amount *big.Int, stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.UnlockStake(&_Directory.TransactOpts, amount, stakee)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address stakee) returns()
func (_Directory *DirectoryTransactor) WithdrawStake(opts *bind.TransactOpts, stakee common.Address) (*types.Transaction, error) {
	return _Directory.contract.Transact(opts, "withdrawStake", stakee)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address stakee) returns()
func (_Directory *DirectorySession) WithdrawStake(stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.WithdrawStake(&_Directory.TransactOpts, stakee)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address stakee) returns()
func (_Directory *DirectoryTransactorSession) WithdrawStake(stakee common.Address) (*types.Transaction, error) {
	return _Directory.Contract.WithdrawStake(&_Directory.TransactOpts, stakee)
}

// DirectoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Directory contract.
type DirectoryOwnershipTransferredIterator struct {
	Event *DirectoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DirectoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectoryOwnershipTransferred)
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
		it.Event = new(DirectoryOwnershipTransferred)
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
func (it *DirectoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectoryOwnershipTransferred represents a OwnershipTransferred event raised by the Directory contract.
type DirectoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Directory *DirectoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DirectoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Directory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DirectoryOwnershipTransferredIterator{contract: _Directory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Directory *DirectoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DirectoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Directory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectoryOwnershipTransferred)
				if err := _Directory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Directory *DirectoryFilterer) ParseOwnershipTransferred(log types.Log) (*DirectoryOwnershipTransferred, error) {
	event := new(DirectoryOwnershipTransferred)
	if err := _Directory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
