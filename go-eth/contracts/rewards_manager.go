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

// RewardsManagerRewardPool is an auto generated low-level Go binding around an user-defined struct.
type RewardsManagerRewardPool struct {
	StakersRewardTotal            *big.Int
	InitializedAt                 *big.Int
	TotalActiveStake              *big.Int
	InitialCumulativeRewardFactor *big.Int
	CumulativeRewardFactor        *big.Int
}

// RewardsManagerABI is the input ABI used to generate the binding from.
const RewardsManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"lastClaims\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"latestActiveRewardPools\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unclaimedNodeRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unclaimedStakeRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractStakingManager\",\"name\":\"stakingManager\",\"type\":\"address\"},{\"internalType\":\"contractEpochsManager\",\"name\":\"epochsManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getStakerKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getLastClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getRewardPool\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stakersRewardTotal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initializedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalActiveStake\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"initialCumulativeRewardFactor\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"cumulativeRewardFactor\",\"type\":\"int128\"}],\"internalType\":\"structRewardsManager.RewardPool\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getRewardPoolStakersTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getRewardPoolActiveStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getUnclaimedNodeReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getUnclaimedStakeReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeNextRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"incrementRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"calculateStakerClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"toFixedPointSYLO\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"amount\",\"type\":\"int128\"}],\"name\":\"fromFixedPointSYLO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"claimStakingRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"claimStakingRewardsAsManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNodeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"removeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RewardsManagerBin is the compiled bytecode used for deploying new contracts.
var RewardsManagerBin = "0x608060405234801561001057600080fd5b50611fa1806100206000396000f3fe608060405234801561001057600080fd5b50600436106101ae5760003560e01c806398594f9f116100ee578063cec6e51f11610097578063f0b1b1b911610071578063f0b1b1b9146103f4578063f1fd0d1c1461041d578063f2fde38b1461043d578063fe0f33cb1461045057600080fd5b8063cec6e51f14610366578063d6de5bfd146103c1578063d96d5767146103e157600080fd5b8063ac18de43116100c8578063ac18de431461032d578063c0c53b8b14610340578063ce8b486f1461035357600080fd5b806398594f9f146102ff578063a58b603b14610312578063ab56b4f01461031a57600080fd5b80634c2681531161015b578063715018a611610135578063715018a6146102b657806378c041a6146102be5780638da5cb5b146102d157806396bb1fef146102ec57600080fd5b80634c2681531461027057806351e9d686146102835780635582fccb146102a357600080fd5b80633150fa8a1161018c5780633150fa8a146102245780633584b5851461023757806340768fd31461025d57600080fd5b806307b4fc01146101b35780630cac4ea8146101e65780632d06177a1461020f575b600080fd5b6101d36101c1366004611bff565b606a6020526000908152604090205481565b6040519081526020015b60405180910390f35b6101d36101f4366004611bff565b6001600160a01b031660009081526068602052604090205490565b61022261021d366004611bff565b610458565b005b6101d3610232366004611e3d565b6104d3565b61024a610245366004611c9e565b6104fe565b604051600f9190910b81526020016101dd565b61022261026b366004611c53565b610567565b6101d361027e366004611e3d565b610882565b6101d3610291366004611c9e565b606b6020526000908152604090205481565b6101d36102b1366004611d00565b6108d0565b61022261095e565b6101d36102cc366004611c1b565b6109c4565b6033546040516001600160a01b0390911681526020016101dd565b6102226102fa366004611bff565b6109f5565b61022261030d366004611c1b565b610b46565b610222610cd1565b6101d3610328366004611c1b565b610e2d565b61022261033b366004611bff565b610e54565b61022261034e366004611cb6565b610ec8565b6101d3610361366004611c1b565b610fda565b610379610374366004611e3d565b6112ae565b6040516101dd9190600060a0820190508251825260208301516020830152604083015160408301526060830151600f0b60608301526080830151600f0b608083015292915050565b6101d36103cf366004611bff565b60686020526000908152604090205481565b6101d36103ef366004611e3d565b611364565b6101d3610402366004611bff565b6001600160a01b031660009081526069602052604090205490565b6101d361042b366004611bff565b60696020526000908152604090205481565b61022261044b366004611bff565b61138e565b610222611470565b6033546001600160a01b031633146104b75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6001600160a01b03166000908152606d60205260409020439055565b6000606c60006104e38585610882565b81526020019081526020016000206000015490505b92915050565b60008061051b610516670de0b6b3a764000085611eb1565b611762565b90506000610534610516670de0b6b3a764000086611f16565b90506000610549670de0b6b3a7640000611762565b905061055e836105598484611780565b611800565b95945050505050565b336000908152606d60205260409020546105e95760405162461bcd60e51b815260206004820152603560248201527f4f6e6c79206d616e6167657273206f66207468697320636f6e7472616374206360448201527f616e2063616c6c20746869732066756e6374696f6e000000000000000000000060648201526084016104ae565b606754604080517fe34bf01300000000000000000000000000000000000000000000000000000000815290516000926001600160a01b03169163e34bf01391600480830192610140929190829003018186803b15801561064857600080fd5b505afa15801561065c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106809190611d21565b90506000606c6000610696846000015187610882565b8152602001908152602001600020905060008160020154116107205760405162461bcd60e51b815260206004820152603a60248201527f52657761726420706f6f6c20686173206e6f74206265656e20696e697469616c60448201527f697a656420666f72207468652063757272656e742065706f636800000000000060648201526084016104ae565b6000610730848460800151611845565b905061073c8185611ee4565b6001600160a01b03861660009081526068602052604081208054909190610764908490611e99565b90915550506001600160a01b03851660009081526069602052604081208054839290610791908490611e99565b90915550508154819083906000906107aa908490611e99565b90915550506003820154600f90810b900b610823576107e16107cf83600001546104fe565b6107dc84600201546104fe565b611780565b8260030160106101000a8154816fffffffffffffffffffffffffffffffff0219169083600f0b6fffffffffffffffffffffffffffffffff16021790555061087b565b60038201548254600284015461083d92600f0b9190611874565b8260030160106101000a8154816fffffffffffffffffffffffffffffffff0219169083600f0b6fffffffffffffffffffffffffffffffff1602179055505b5050505050565b600082826040516020016108b292919091825260601b6bffffffffffffffffffffffff1916602082015260340190565b60405160208183030381529060405280519060200120905092915050565b6000806108dc8361189b565b67ffffffffffffffff16905060006108fc670de0b6b3a764000083611ec5565b905060006109128561090d85611762565b6118b7565b90506000610927670de0b6b3a7640000611762565b9050600061093d61093884846118fc565b61189b565b67ffffffffffffffff1690506109538185611e99565b979650505050505050565b6033546001600160a01b031633146109b85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104ae565b6109c26000611944565b565b6040516bffffffffffffffffffffffff19606084811b8216602084015283901b1660348201526000906048016108b2565b6000610a018233610fda565b905060008111610a535760405162461bcd60e51b815260206004820152601060248201527f4e6f7468696e6720746f20636c61696d0000000000000000000000000000000060448201526064016104ae565b6001600160a01b03821660009081526069602052604081208054839290610a7b908490611ee4565b90915550506001600160a01b0382166000908152606a602052604081205490606b90610aa785336109c4565b81526020810191909152604090810160002091909155606554905163a9059cbb60e01b8152336004820152602481018390526001600160a01b039091169063a9059cbb906044015b602060405180830381600087803b158015610b0957600080fd5b505af1158015610b1d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b419190611c7e565b505050565b336000908152606d6020526040902054610bc85760405162461bcd60e51b815260206004820152603560248201527f4f6e6c79206d616e6167657273206f66207468697320636f6e7472616374206360448201527f616e2063616c6c20746869732066756e6374696f6e000000000000000000000060648201526084016104ae565b6000610bd48383610fda565b6001600160a01b0384166000908152606a6020526040812054919250606b90610bfd86866109c4565b815260208101919091526040016000205580610c1857505050565b6001600160a01b03831660009081526069602052604081208054839290610c40908490611ee4565b909155505060655460405163a9059cbb60e01b81526001600160a01b038481166004830152602482018490529091169063a9059cbb90604401602060405180830381600087803b158015610c9357600080fd5b505af1158015610ca7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ccb9190611c7e565b50505050565b336000818152606860205260408082205460665491516313cdd31b60e01b81526004810194909452926001600160a01b03909116906313cdd31b9060240160206040518083038186803b158015610d2757600080fd5b505afa158015610d3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d5f9190611e25565b905080610d935733600090815260696020526040902054610d809083611e99565b3360009081526069602052604081205591505b60008211610de35760405162461bcd60e51b815260206004820152601060248201527f4e6f7468696e6720746f20636c61696d0000000000000000000000000000000060448201526064016104ae565b3360008181526068602052604080822091909155606554905163a9059cbb60e01b81526004810192909252602482018490526001600160a01b03169063a9059cbb90604401610aef565b6000606b6000610e3d85856109c4565b815260200190815260200160002054905092915050565b6033546001600160a01b03163314610eae5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104ae565b6001600160a01b03166000908152606d6020526040812055565b600054610100900460ff1680610ee1575060005460ff16155b610f445760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016104ae565b600054610100900460ff16158015610f66576000805461ffff19166101011790555b610f6e6119ae565b606580546001600160a01b038087167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556067805485841690831617905560668054928616929091169190911790558015610ccb576000805461ff001916905550505050565b6066546040517ff731cb490000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301528381166024830152600092839291169063f731cb499060440160606040518083038186803b15801561104457600080fd5b505afa158015611058573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061107c9190611dbe565b805190915061108f5760009150506104f8565b600080606b60006110a088886109c4565b81526020019081526020016000205460016110bb9190611e99565b90505b606760009054906101000a90046001600160a01b03166001600160a01b031663b3e123db6040518163ffffffff1660e01b815260040160206040518083038186803b15801561110c57600080fd5b505afa158015611120573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111449190611e25565b811015611194576000606c600061115b848a610882565b815260200190815260200160002090506000816001015411156111815781925050611194565b508061118c81611efb565b9150506110be565b50806111a5576000925050506104f8565b6000606c60006111b58489610882565b8152602001908152602001600020905060006111d484600001516104fe565b6003830154909150600f90810b9081900b61123557611213826105596111fd86600001546104fe565b61120e866107dc89600201546104fe565b6118fc565b60038401549092507001000000000000000000000000000000009004600f0b90505b6001600160a01b0388166000908152606a6020526040812054606c90829061125d908c610882565b8152602001908152602001600020905060006112916102b18561120e8560030160109054906101000a9004600f0b87611780565b87519091506112a09082611ee4565b9a9950505050505050505050565b6112e66040518060a001604052806000815260200160008152602001600081526020016000600f0b81526020016000600f0b81525090565b606c60006112f48585610882565b81526020808201929092526040908101600020815160a0810183528154815260018201549381019390935260028101549183019190915260030154600f81810b810b810b6060840152700100000000000000000000000000000000909104810b810b900b60808201529392505050565b6000606c60006113748585610882565b815260200190815260200160002060020154905092915050565b6033546001600160a01b031633146113e85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104ae565b6001600160a01b0381166114645760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104ae565b61146d81611944565b50565b606754604080517fb3e123db00000000000000000000000000000000000000000000000000000000815290516000926001600160a01b03169163b3e123db916004808301926020929190829003018186803b1580156114ce57600080fd5b505afa1580156114e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115069190611e25565b90506000606c60006115188433610882565b8152602001908152602001600020905080600101546000146115a25760405162461bcd60e51b815260206004820152603160248201527f546865206e6578742072657761726420706f6f6c2068617320616c726561647960448201527f206265656e20696e697469616c697a656400000000000000000000000000000060648201526084016104ae565b6066546040516313cdd31b60e01b81523360048201526000916001600160a01b0316906313cdd31b9060240160206040518083038186803b1580156115e657600080fd5b505afa1580156115fa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061161e9190611e25565b9050600081116116965760405162461bcd60e51b815260206004820152602c60248201527f4d7573742068617665207374616b6520746f20696e74697469616c697a65206160448201527f2072657761726420706f6f6c000000000000000000000000000000000000000060648201526084016104ae565b436001830155336000908152606960205260409020546116b69082611e99565b6002830155336000818152606a6020526040812054606c926116d89190610882565b81526020808201929092526040908101600090812060039081015495018054700100000000000000000000000000000000909604600f90810b900b6fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff0000000000000000000000000000000090961695909517909455338452606a9091529091209190915550565b6000677fffffffffffffff82111561177957600080fd5b5060401b90565b600081600f0b6000141561179357600080fd5b600082600f0b604085600f0b901b816117bc57634e487b7160e01b600052601260045260246000fd5b0590506f7fffffffffffffffffffffffffffffff1981128015906117f057506f7fffffffffffffffffffffffffffffff8113155b6117f957600080fd5b9392505050565b6000600f83810b9083900b016f7fffffffffffffffffffffffffffffff1981128015906117f057506f7fffffffffffffffffffffffffffffff8113156117f957600080fd5b600061271061186a61ffff84166fffffffffffffffffffffffffffffffff8616611ec5565b6117f99190611eb1565b6000611893846105598661120e61188a886104fe565b6107dc886104fe565b949350505050565b60008082600f0b12156118ad57600080fd5b50600f0b60401d90565b6000600f82810b9084900b036f7fffffffffffffffffffffffffffffff1981128015906117f057506f7fffffffffffffffffffffffffffffff8113156117f957600080fd5b6000600f83810b9083900b0260401d6f7fffffffffffffffffffffffffffffff1981128015906117f057506f7fffffffffffffffffffffffffffffff8113156117f957600080fd5b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16806119c7575060005460ff16155b611a2a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016104ae565b600054610100900460ff16158015611a4c576000805461ffff19166101011790555b611a54611a70565b611a5c611b21565b801561146d576000805461ff001916905550565b600054610100900460ff1680611a89575060005460ff16155b611aec5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016104ae565b600054610100900460ff16158015611a5c576000805461ffff1916610101179055801561146d576000805461ff001916905550565b600054610100900460ff1680611b3a575060005460ff16155b611b9d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016104ae565b600054610100900460ff16158015611bbf576000805461ffff19166101011790555b611a5c33611944565b80516fffffffffffffffffffffffffffffffff81168114611be857600080fd5b919050565b805161ffff81168114611be857600080fd5b600060208284031215611c10578081fd5b81356117f981611f56565b60008060408385031215611c2d578081fd5b8235611c3881611f56565b91506020830135611c4881611f56565b809150509250929050565b60008060408385031215611c65578182fd5b8235611c7081611f56565b946020939093013593505050565b600060208284031215611c8f578081fd5b815180151581146117f9578182fd5b600060208284031215611caf578081fd5b5035919050565b600080600060608486031215611cca578081fd5b8335611cd581611f56565b92506020840135611ce581611f56565b91506040840135611cf581611f56565b809150509250925092565b600060208284031215611d11578081fd5b813580600f0b81146117f9578182fd5b60006101408284031215611d33578081fd5b611d3b611e61565b82518152602083015160208201526040830151604082015260608301516060820152611d6960808401611bed565b608082015260a083015160a0820152611d8460c08401611bc8565b60c0820152611d9560e08401611bc8565b60e08201526101008381015190820152610120611db3818501611bed565b908201529392505050565b600060608284031215611dcf578081fd5b6040516060810181811067ffffffffffffffff82111715611dfe57634e487b7160e01b83526041600452602483fd5b80604052508251815260208301516020820152604083015160408201528091505092915050565b600060208284031215611e36578081fd5b5051919050565b60008060408385031215611e4f578081fd5b823591506020830135611c4881611f56565b604051610140810167ffffffffffffffff81118282101715611e9357634e487b7160e01b600052604160045260246000fd5b60405290565b60008219821115611eac57611eac611f2a565b500190565b600082611ec057611ec0611f40565b500490565b6000816000190483118215151615611edf57611edf611f2a565b500290565b600082821015611ef657611ef6611f2a565b500390565b6000600019821415611f0f57611f0f611f2a565b5060010190565b600082611f2557611f25611f40565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b6001600160a01b038116811461146d57600080fdfea2646970667358221220fcd0b39473877c13a5ba5e8caff6980500fd41cb1a9cee409179b2abeb6ee40864736f6c63430008040033"

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

// CalculateStakerClaim is a free data retrieval call binding the contract method 0xce8b486f.
//
// Solidity: function calculateStakerClaim(address stakee, address staker) view returns(uint256)
func (_RewardsManager *RewardsManagerCaller) CalculateStakerClaim(opts *bind.CallOpts, stakee common.Address, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "calculateStakerClaim", stakee, staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateStakerClaim is a free data retrieval call binding the contract method 0xce8b486f.
//
// Solidity: function calculateStakerClaim(address stakee, address staker) view returns(uint256)
func (_RewardsManager *RewardsManagerSession) CalculateStakerClaim(stakee common.Address, staker common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.CalculateStakerClaim(&_RewardsManager.CallOpts, stakee, staker)
}

// CalculateStakerClaim is a free data retrieval call binding the contract method 0xce8b486f.
//
// Solidity: function calculateStakerClaim(address stakee, address staker) view returns(uint256)
func (_RewardsManager *RewardsManagerCallerSession) CalculateStakerClaim(stakee common.Address, staker common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.CalculateStakerClaim(&_RewardsManager.CallOpts, stakee, staker)
}

// FromFixedPointSYLO is a free data retrieval call binding the contract method 0x5582fccb.
//
// Solidity: function fromFixedPointSYLO(int128 amount) pure returns(uint256)
func (_RewardsManager *RewardsManagerCaller) FromFixedPointSYLO(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "fromFixedPointSYLO", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FromFixedPointSYLO is a free data retrieval call binding the contract method 0x5582fccb.
//
// Solidity: function fromFixedPointSYLO(int128 amount) pure returns(uint256)
func (_RewardsManager *RewardsManagerSession) FromFixedPointSYLO(amount *big.Int) (*big.Int, error) {
	return _RewardsManager.Contract.FromFixedPointSYLO(&_RewardsManager.CallOpts, amount)
}

// FromFixedPointSYLO is a free data retrieval call binding the contract method 0x5582fccb.
//
// Solidity: function fromFixedPointSYLO(int128 amount) pure returns(uint256)
func (_RewardsManager *RewardsManagerCallerSession) FromFixedPointSYLO(amount *big.Int) (*big.Int, error) {
	return _RewardsManager.Contract.FromFixedPointSYLO(&_RewardsManager.CallOpts, amount)
}

// GetKey is a free data retrieval call binding the contract method 0x4c268153.
//
// Solidity: function getKey(uint256 epochId, address stakee) pure returns(bytes32)
func (_RewardsManager *RewardsManagerCaller) GetKey(opts *bind.CallOpts, epochId *big.Int, stakee common.Address) ([32]byte, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "getKey", epochId, stakee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKey is a free data retrieval call binding the contract method 0x4c268153.
//
// Solidity: function getKey(uint256 epochId, address stakee) pure returns(bytes32)
func (_RewardsManager *RewardsManagerSession) GetKey(epochId *big.Int, stakee common.Address) ([32]byte, error) {
	return _RewardsManager.Contract.GetKey(&_RewardsManager.CallOpts, epochId, stakee)
}

// GetKey is a free data retrieval call binding the contract method 0x4c268153.
//
// Solidity: function getKey(uint256 epochId, address stakee) pure returns(bytes32)
func (_RewardsManager *RewardsManagerCallerSession) GetKey(epochId *big.Int, stakee common.Address) ([32]byte, error) {
	return _RewardsManager.Contract.GetKey(&_RewardsManager.CallOpts, epochId, stakee)
}

// GetLastClaim is a free data retrieval call binding the contract method 0xab56b4f0.
//
// Solidity: function getLastClaim(address stakee, address staker) view returns(uint256)
func (_RewardsManager *RewardsManagerCaller) GetLastClaim(opts *bind.CallOpts, stakee common.Address, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "getLastClaim", stakee, staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastClaim is a free data retrieval call binding the contract method 0xab56b4f0.
//
// Solidity: function getLastClaim(address stakee, address staker) view returns(uint256)
func (_RewardsManager *RewardsManagerSession) GetLastClaim(stakee common.Address, staker common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.GetLastClaim(&_RewardsManager.CallOpts, stakee, staker)
}

// GetLastClaim is a free data retrieval call binding the contract method 0xab56b4f0.
//
// Solidity: function getLastClaim(address stakee, address staker) view returns(uint256)
func (_RewardsManager *RewardsManagerCallerSession) GetLastClaim(stakee common.Address, staker common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.GetLastClaim(&_RewardsManager.CallOpts, stakee, staker)
}

// GetRewardPool is a free data retrieval call binding the contract method 0xcec6e51f.
//
// Solidity: function getRewardPool(uint256 epochId, address stakee) view returns((uint256,uint256,uint256,int128,int128))
func (_RewardsManager *RewardsManagerCaller) GetRewardPool(opts *bind.CallOpts, epochId *big.Int, stakee common.Address) (RewardsManagerRewardPool, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "getRewardPool", epochId, stakee)

	if err != nil {
		return *new(RewardsManagerRewardPool), err
	}

	out0 := *abi.ConvertType(out[0], new(RewardsManagerRewardPool)).(*RewardsManagerRewardPool)

	return out0, err

}

// GetRewardPool is a free data retrieval call binding the contract method 0xcec6e51f.
//
// Solidity: function getRewardPool(uint256 epochId, address stakee) view returns((uint256,uint256,uint256,int128,int128))
func (_RewardsManager *RewardsManagerSession) GetRewardPool(epochId *big.Int, stakee common.Address) (RewardsManagerRewardPool, error) {
	return _RewardsManager.Contract.GetRewardPool(&_RewardsManager.CallOpts, epochId, stakee)
}

// GetRewardPool is a free data retrieval call binding the contract method 0xcec6e51f.
//
// Solidity: function getRewardPool(uint256 epochId, address stakee) view returns((uint256,uint256,uint256,int128,int128))
func (_RewardsManager *RewardsManagerCallerSession) GetRewardPool(epochId *big.Int, stakee common.Address) (RewardsManagerRewardPool, error) {
	return _RewardsManager.Contract.GetRewardPool(&_RewardsManager.CallOpts, epochId, stakee)
}

// GetRewardPoolActiveStake is a free data retrieval call binding the contract method 0xd96d5767.
//
// Solidity: function getRewardPoolActiveStake(uint256 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerCaller) GetRewardPoolActiveStake(opts *bind.CallOpts, epochId *big.Int, stakee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "getRewardPoolActiveStake", epochId, stakee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardPoolActiveStake is a free data retrieval call binding the contract method 0xd96d5767.
//
// Solidity: function getRewardPoolActiveStake(uint256 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerSession) GetRewardPoolActiveStake(epochId *big.Int, stakee common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.GetRewardPoolActiveStake(&_RewardsManager.CallOpts, epochId, stakee)
}

// GetRewardPoolActiveStake is a free data retrieval call binding the contract method 0xd96d5767.
//
// Solidity: function getRewardPoolActiveStake(uint256 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerCallerSession) GetRewardPoolActiveStake(epochId *big.Int, stakee common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.GetRewardPoolActiveStake(&_RewardsManager.CallOpts, epochId, stakee)
}

// GetRewardPoolStakersTotal is a free data retrieval call binding the contract method 0x3150fa8a.
//
// Solidity: function getRewardPoolStakersTotal(uint256 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerCaller) GetRewardPoolStakersTotal(opts *bind.CallOpts, epochId *big.Int, stakee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "getRewardPoolStakersTotal", epochId, stakee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardPoolStakersTotal is a free data retrieval call binding the contract method 0x3150fa8a.
//
// Solidity: function getRewardPoolStakersTotal(uint256 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerSession) GetRewardPoolStakersTotal(epochId *big.Int, stakee common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.GetRewardPoolStakersTotal(&_RewardsManager.CallOpts, epochId, stakee)
}

// GetRewardPoolStakersTotal is a free data retrieval call binding the contract method 0x3150fa8a.
//
// Solidity: function getRewardPoolStakersTotal(uint256 epochId, address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerCallerSession) GetRewardPoolStakersTotal(epochId *big.Int, stakee common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.GetRewardPoolStakersTotal(&_RewardsManager.CallOpts, epochId, stakee)
}

// GetStakerKey is a free data retrieval call binding the contract method 0x78c041a6.
//
// Solidity: function getStakerKey(address stakee, address staker) pure returns(bytes32)
func (_RewardsManager *RewardsManagerCaller) GetStakerKey(opts *bind.CallOpts, stakee common.Address, staker common.Address) ([32]byte, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "getStakerKey", stakee, staker)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStakerKey is a free data retrieval call binding the contract method 0x78c041a6.
//
// Solidity: function getStakerKey(address stakee, address staker) pure returns(bytes32)
func (_RewardsManager *RewardsManagerSession) GetStakerKey(stakee common.Address, staker common.Address) ([32]byte, error) {
	return _RewardsManager.Contract.GetStakerKey(&_RewardsManager.CallOpts, stakee, staker)
}

// GetStakerKey is a free data retrieval call binding the contract method 0x78c041a6.
//
// Solidity: function getStakerKey(address stakee, address staker) pure returns(bytes32)
func (_RewardsManager *RewardsManagerCallerSession) GetStakerKey(stakee common.Address, staker common.Address) ([32]byte, error) {
	return _RewardsManager.Contract.GetStakerKey(&_RewardsManager.CallOpts, stakee, staker)
}

// GetUnclaimedNodeReward is a free data retrieval call binding the contract method 0x0cac4ea8.
//
// Solidity: function getUnclaimedNodeReward(address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerCaller) GetUnclaimedNodeReward(opts *bind.CallOpts, stakee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "getUnclaimedNodeReward", stakee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnclaimedNodeReward is a free data retrieval call binding the contract method 0x0cac4ea8.
//
// Solidity: function getUnclaimedNodeReward(address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerSession) GetUnclaimedNodeReward(stakee common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.GetUnclaimedNodeReward(&_RewardsManager.CallOpts, stakee)
}

// GetUnclaimedNodeReward is a free data retrieval call binding the contract method 0x0cac4ea8.
//
// Solidity: function getUnclaimedNodeReward(address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerCallerSession) GetUnclaimedNodeReward(stakee common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.GetUnclaimedNodeReward(&_RewardsManager.CallOpts, stakee)
}

// GetUnclaimedStakeReward is a free data retrieval call binding the contract method 0xf0b1b1b9.
//
// Solidity: function getUnclaimedStakeReward(address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerCaller) GetUnclaimedStakeReward(opts *bind.CallOpts, stakee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "getUnclaimedStakeReward", stakee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnclaimedStakeReward is a free data retrieval call binding the contract method 0xf0b1b1b9.
//
// Solidity: function getUnclaimedStakeReward(address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerSession) GetUnclaimedStakeReward(stakee common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.GetUnclaimedStakeReward(&_RewardsManager.CallOpts, stakee)
}

// GetUnclaimedStakeReward is a free data retrieval call binding the contract method 0xf0b1b1b9.
//
// Solidity: function getUnclaimedStakeReward(address stakee) view returns(uint256)
func (_RewardsManager *RewardsManagerCallerSession) GetUnclaimedStakeReward(stakee common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.GetUnclaimedStakeReward(&_RewardsManager.CallOpts, stakee)
}

// LastClaims is a free data retrieval call binding the contract method 0x51e9d686.
//
// Solidity: function lastClaims(bytes32 ) view returns(uint256)
func (_RewardsManager *RewardsManagerCaller) LastClaims(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "lastClaims", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastClaims is a free data retrieval call binding the contract method 0x51e9d686.
//
// Solidity: function lastClaims(bytes32 ) view returns(uint256)
func (_RewardsManager *RewardsManagerSession) LastClaims(arg0 [32]byte) (*big.Int, error) {
	return _RewardsManager.Contract.LastClaims(&_RewardsManager.CallOpts, arg0)
}

// LastClaims is a free data retrieval call binding the contract method 0x51e9d686.
//
// Solidity: function lastClaims(bytes32 ) view returns(uint256)
func (_RewardsManager *RewardsManagerCallerSession) LastClaims(arg0 [32]byte) (*big.Int, error) {
	return _RewardsManager.Contract.LastClaims(&_RewardsManager.CallOpts, arg0)
}

// LatestActiveRewardPools is a free data retrieval call binding the contract method 0x07b4fc01.
//
// Solidity: function latestActiveRewardPools(address ) view returns(uint256)
func (_RewardsManager *RewardsManagerCaller) LatestActiveRewardPools(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "latestActiveRewardPools", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestActiveRewardPools is a free data retrieval call binding the contract method 0x07b4fc01.
//
// Solidity: function latestActiveRewardPools(address ) view returns(uint256)
func (_RewardsManager *RewardsManagerSession) LatestActiveRewardPools(arg0 common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.LatestActiveRewardPools(&_RewardsManager.CallOpts, arg0)
}

// LatestActiveRewardPools is a free data retrieval call binding the contract method 0x07b4fc01.
//
// Solidity: function latestActiveRewardPools(address ) view returns(uint256)
func (_RewardsManager *RewardsManagerCallerSession) LatestActiveRewardPools(arg0 common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.LatestActiveRewardPools(&_RewardsManager.CallOpts, arg0)
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

// ToFixedPointSYLO is a free data retrieval call binding the contract method 0x3584b585.
//
// Solidity: function toFixedPointSYLO(uint256 amount) pure returns(int128)
func (_RewardsManager *RewardsManagerCaller) ToFixedPointSYLO(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "toFixedPointSYLO", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToFixedPointSYLO is a free data retrieval call binding the contract method 0x3584b585.
//
// Solidity: function toFixedPointSYLO(uint256 amount) pure returns(int128)
func (_RewardsManager *RewardsManagerSession) ToFixedPointSYLO(amount *big.Int) (*big.Int, error) {
	return _RewardsManager.Contract.ToFixedPointSYLO(&_RewardsManager.CallOpts, amount)
}

// ToFixedPointSYLO is a free data retrieval call binding the contract method 0x3584b585.
//
// Solidity: function toFixedPointSYLO(uint256 amount) pure returns(int128)
func (_RewardsManager *RewardsManagerCallerSession) ToFixedPointSYLO(amount *big.Int) (*big.Int, error) {
	return _RewardsManager.Contract.ToFixedPointSYLO(&_RewardsManager.CallOpts, amount)
}

// UnclaimedNodeRewards is a free data retrieval call binding the contract method 0xd6de5bfd.
//
// Solidity: function unclaimedNodeRewards(address ) view returns(uint256)
func (_RewardsManager *RewardsManagerCaller) UnclaimedNodeRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "unclaimedNodeRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnclaimedNodeRewards is a free data retrieval call binding the contract method 0xd6de5bfd.
//
// Solidity: function unclaimedNodeRewards(address ) view returns(uint256)
func (_RewardsManager *RewardsManagerSession) UnclaimedNodeRewards(arg0 common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.UnclaimedNodeRewards(&_RewardsManager.CallOpts, arg0)
}

// UnclaimedNodeRewards is a free data retrieval call binding the contract method 0xd6de5bfd.
//
// Solidity: function unclaimedNodeRewards(address ) view returns(uint256)
func (_RewardsManager *RewardsManagerCallerSession) UnclaimedNodeRewards(arg0 common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.UnclaimedNodeRewards(&_RewardsManager.CallOpts, arg0)
}

// UnclaimedStakeRewards is a free data retrieval call binding the contract method 0xf1fd0d1c.
//
// Solidity: function unclaimedStakeRewards(address ) view returns(uint256)
func (_RewardsManager *RewardsManagerCaller) UnclaimedStakeRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardsManager.contract.Call(opts, &out, "unclaimedStakeRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnclaimedStakeRewards is a free data retrieval call binding the contract method 0xf1fd0d1c.
//
// Solidity: function unclaimedStakeRewards(address ) view returns(uint256)
func (_RewardsManager *RewardsManagerSession) UnclaimedStakeRewards(arg0 common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.UnclaimedStakeRewards(&_RewardsManager.CallOpts, arg0)
}

// UnclaimedStakeRewards is a free data retrieval call binding the contract method 0xf1fd0d1c.
//
// Solidity: function unclaimedStakeRewards(address ) view returns(uint256)
func (_RewardsManager *RewardsManagerCallerSession) UnclaimedStakeRewards(arg0 common.Address) (*big.Int, error) {
	return _RewardsManager.Contract.UnclaimedStakeRewards(&_RewardsManager.CallOpts, arg0)
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

// ClaimNodeRewards is a paid mutator transaction binding the contract method 0xa58b603b.
//
// Solidity: function claimNodeRewards() returns()
func (_RewardsManager *RewardsManagerTransactor) ClaimNodeRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardsManager.contract.Transact(opts, "claimNodeRewards")
}

// ClaimNodeRewards is a paid mutator transaction binding the contract method 0xa58b603b.
//
// Solidity: function claimNodeRewards() returns()
func (_RewardsManager *RewardsManagerSession) ClaimNodeRewards() (*types.Transaction, error) {
	return _RewardsManager.Contract.ClaimNodeRewards(&_RewardsManager.TransactOpts)
}

// ClaimNodeRewards is a paid mutator transaction binding the contract method 0xa58b603b.
//
// Solidity: function claimNodeRewards() returns()
func (_RewardsManager *RewardsManagerTransactorSession) ClaimNodeRewards() (*types.Transaction, error) {
	return _RewardsManager.Contract.ClaimNodeRewards(&_RewardsManager.TransactOpts)
}

// ClaimStakingRewards is a paid mutator transaction binding the contract method 0x96bb1fef.
//
// Solidity: function claimStakingRewards(address stakee) returns()
func (_RewardsManager *RewardsManagerTransactor) ClaimStakingRewards(opts *bind.TransactOpts, stakee common.Address) (*types.Transaction, error) {
	return _RewardsManager.contract.Transact(opts, "claimStakingRewards", stakee)
}

// ClaimStakingRewards is a paid mutator transaction binding the contract method 0x96bb1fef.
//
// Solidity: function claimStakingRewards(address stakee) returns()
func (_RewardsManager *RewardsManagerSession) ClaimStakingRewards(stakee common.Address) (*types.Transaction, error) {
	return _RewardsManager.Contract.ClaimStakingRewards(&_RewardsManager.TransactOpts, stakee)
}

// ClaimStakingRewards is a paid mutator transaction binding the contract method 0x96bb1fef.
//
// Solidity: function claimStakingRewards(address stakee) returns()
func (_RewardsManager *RewardsManagerTransactorSession) ClaimStakingRewards(stakee common.Address) (*types.Transaction, error) {
	return _RewardsManager.Contract.ClaimStakingRewards(&_RewardsManager.TransactOpts, stakee)
}

// ClaimStakingRewardsAsManager is a paid mutator transaction binding the contract method 0x98594f9f.
//
// Solidity: function claimStakingRewardsAsManager(address stakee, address staker) returns()
func (_RewardsManager *RewardsManagerTransactor) ClaimStakingRewardsAsManager(opts *bind.TransactOpts, stakee common.Address, staker common.Address) (*types.Transaction, error) {
	return _RewardsManager.contract.Transact(opts, "claimStakingRewardsAsManager", stakee, staker)
}

// ClaimStakingRewardsAsManager is a paid mutator transaction binding the contract method 0x98594f9f.
//
// Solidity: function claimStakingRewardsAsManager(address stakee, address staker) returns()
func (_RewardsManager *RewardsManagerSession) ClaimStakingRewardsAsManager(stakee common.Address, staker common.Address) (*types.Transaction, error) {
	return _RewardsManager.Contract.ClaimStakingRewardsAsManager(&_RewardsManager.TransactOpts, stakee, staker)
}

// ClaimStakingRewardsAsManager is a paid mutator transaction binding the contract method 0x98594f9f.
//
// Solidity: function claimStakingRewardsAsManager(address stakee, address staker) returns()
func (_RewardsManager *RewardsManagerTransactorSession) ClaimStakingRewardsAsManager(stakee common.Address, staker common.Address) (*types.Transaction, error) {
	return _RewardsManager.Contract.ClaimStakingRewardsAsManager(&_RewardsManager.TransactOpts, stakee, staker)
}

// IncrementRewardPool is a paid mutator transaction binding the contract method 0x40768fd3.
//
// Solidity: function incrementRewardPool(address stakee, uint256 amount) returns()
func (_RewardsManager *RewardsManagerTransactor) IncrementRewardPool(opts *bind.TransactOpts, stakee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardsManager.contract.Transact(opts, "incrementRewardPool", stakee, amount)
}

// IncrementRewardPool is a paid mutator transaction binding the contract method 0x40768fd3.
//
// Solidity: function incrementRewardPool(address stakee, uint256 amount) returns()
func (_RewardsManager *RewardsManagerSession) IncrementRewardPool(stakee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardsManager.Contract.IncrementRewardPool(&_RewardsManager.TransactOpts, stakee, amount)
}

// IncrementRewardPool is a paid mutator transaction binding the contract method 0x40768fd3.
//
// Solidity: function incrementRewardPool(address stakee, uint256 amount) returns()
func (_RewardsManager *RewardsManagerTransactorSession) IncrementRewardPool(stakee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardsManager.Contract.IncrementRewardPool(&_RewardsManager.TransactOpts, stakee, amount)
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

// InitializeNextRewardPool is a paid mutator transaction binding the contract method 0xfe0f33cb.
//
// Solidity: function initializeNextRewardPool() returns()
func (_RewardsManager *RewardsManagerTransactor) InitializeNextRewardPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardsManager.contract.Transact(opts, "initializeNextRewardPool")
}

// InitializeNextRewardPool is a paid mutator transaction binding the contract method 0xfe0f33cb.
//
// Solidity: function initializeNextRewardPool() returns()
func (_RewardsManager *RewardsManagerSession) InitializeNextRewardPool() (*types.Transaction, error) {
	return _RewardsManager.Contract.InitializeNextRewardPool(&_RewardsManager.TransactOpts)
}

// InitializeNextRewardPool is a paid mutator transaction binding the contract method 0xfe0f33cb.
//
// Solidity: function initializeNextRewardPool() returns()
func (_RewardsManager *RewardsManagerTransactorSession) InitializeNextRewardPool() (*types.Transaction, error) {
	return _RewardsManager.Contract.InitializeNextRewardPool(&_RewardsManager.TransactOpts)
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
