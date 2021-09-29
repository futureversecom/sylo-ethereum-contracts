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
const RewardsManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"lastClaims\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"latestActiveRewardPools\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unclaimedNodeRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unclaimedStakeRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractStakingManager\",\"name\":\"stakingManager\",\"type\":\"address\"},{\"internalType\":\"contractEpochsManager\",\"name\":\"epochsManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getStakerKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getLastClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getRewardPool\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stakersRewardTotal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initializedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalActiveStake\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"initialCumulativeRewardFactor\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"cumulativeRewardFactor\",\"type\":\"int128\"}],\"internalType\":\"structRewardsManager.RewardPool\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getRewardPoolStakersTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getRewardPoolActiveStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getUnclaimedNodeReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"getUnclaimedStakeReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeNextRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"incrementRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"calculateStakerClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"}],\"name\":\"claimStakingRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"claimStakingRewardsAsManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNodeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"removeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RewardsManagerBin is the compiled bytecode used for deploying new contracts.
var RewardsManagerBin = "0x608060405234801561001057600080fd5b50611dee806100206000396000f3fe608060405234801561001057600080fd5b50600436106101985760003560e01c8063a58b603b116100e3578063d6de5bfd1161008c578063f1fd0d1c11610066578063f1fd0d1c146103ce578063f2fde38b146103ee578063fe0f33cb1461040157600080fd5b8063d6de5bfd14610372578063d96d576714610392578063f0b1b1b9146103a557600080fd5b8063c0c53b8b116100bd578063c0c53b8b146102f1578063ce8b486f14610304578063cec6e51f1461031757600080fd5b8063a58b603b146102c3578063ab56b4f0146102cb578063ac18de43146102de57600080fd5b806351e9d686116101455780638da5cb5b1161011f5780638da5cb5b1461028257806396bb1fef1461029d57806398594f9f146102b057600080fd5b806351e9d68614610247578063715018a61461026757806378c041a61461026f57600080fd5b80633150fa8a116101765780633150fa8a1461020e57806340768fd3146102215780634c2681531461023457600080fd5b806307b4fc011461019d5780630cac4ea8146101d05780632d06177a146101f9575b600080fd5b6101bd6101ab366004611a8b565b606a6020526000908152604090205481565b6040519081526020015b60405180910390f35b6101bd6101de366004611a8b565b6001600160a01b031660009081526068602052604090205490565b61020c610207366004611a8b565b610409565b005b6101bd61021c366004611ca8565b610484565b61020c61022f366004611adf565b6104af565b6101bd610242366004611ca8565b6107ca565b6101bd610255366004611b2a565b606b6020526000908152604090205481565b61020c610818565b6101bd61027d366004611aa7565b61087e565b6033546040516001600160a01b0390911681526020016101c7565b61020c6102ab366004611a8b565b6108af565b61020c6102be366004611aa7565b610a00565b61020c610b8b565b6101bd6102d9366004611aa7565b610ce7565b61020c6102ec366004611a8b565b610d0e565b61020c6102ff366004611b42565b610d82565b6101bd610312366004611aa7565b610e94565b61032a610325366004611ca8565b61117f565b6040516101c79190600060a0820190508251825260208301516020830152604083015160408301526060830151600f0b60608301526080830151600f0b608083015292915050565b6101bd610380366004611a8b565b60686020526000908152604090205481565b6101bd6103a0366004611ca8565b611235565b6101bd6103b3366004611a8b565b6001600160a01b031660009081526069602052604090205490565b6101bd6103dc366004611a8b565b60696020526000908152604090205481565b61020c6103fc366004611a8b565b61125f565b61020c611341565b6033546001600160a01b031633146104685760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6001600160a01b03166000908152606d60205260409020439055565b6000606c600061049485856107ca565b81526020019081526020016000206000015490505b92915050565b336000908152606d60205260409020546105315760405162461bcd60e51b815260206004820152603560248201527f4f6e6c79206d616e6167657273206f66207468697320636f6e7472616374206360448201527f616e2063616c6c20746869732066756e6374696f6e0000000000000000000000606482015260840161045f565b606754604080517fe34bf01300000000000000000000000000000000000000000000000000000000815290516000926001600160a01b03169163e34bf01391600480830192610140929190829003018186803b15801561059057600080fd5b505afa1580156105a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c89190611b8c565b90506000606c60006105de8460000151876107ca565b8152602001908152602001600020905060008160020154116106685760405162461bcd60e51b815260206004820152603a60248201527f52657761726420706f6f6c20686173206e6f74206265656e20696e697469616c60448201527f697a656420666f72207468652063757272656e742065706f6368000000000000606482015260840161045f565b6000610678848460800151611633565b90506106848185611d5b565b6001600160a01b038616600090815260686020526040812080549091906106ac908490611d04565b90915550506001600160a01b038516600090815260696020526040812080548392906106d9908490611d04565b90915550508154819083906000906106f2908490611d04565b90915550506003820154600f90810b900b61076b576107296107178360000154611669565b6107248460020154611669565b611687565b8260030160106101000a8154816fffffffffffffffffffffffffffffffff0219169083600f0b6fffffffffffffffffffffffffffffffff1602179055506107c3565b60038201548254600284015461078592600f0b9190611700565b8260030160106101000a8154816fffffffffffffffffffffffffffffffff0219169083600f0b6fffffffffffffffffffffffffffffffff1602179055505b5050505050565b600082826040516020016107fa92919091825260601b6bffffffffffffffffffffffff1916602082015260340190565b60405160208183030381529060405280519060200120905092915050565b6033546001600160a01b031633146108725760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161045f565b61087c6000611727565b565b6040516bffffffffffffffffffffffff19606084811b8216602084015283901b1660348201526000906048016107fa565b60006108bb8233610e94565b90506000811161090d5760405162461bcd60e51b815260206004820152601060248201527f4e6f7468696e6720746f20636c61696d00000000000000000000000000000000604482015260640161045f565b6001600160a01b03821660009081526069602052604081208054839290610935908490611d5b565b90915550506001600160a01b0382166000908152606a602052604081205490606b90610961853361087e565b81526020810191909152604090810160002091909155606554905163a9059cbb60e01b8152336004820152602481018390526001600160a01b039091169063a9059cbb906044015b602060405180830381600087803b1580156109c357600080fd5b505af11580156109d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109fb9190611b0a565b505050565b336000908152606d6020526040902054610a825760405162461bcd60e51b815260206004820152603560248201527f4f6e6c79206d616e6167657273206f66207468697320636f6e7472616374206360448201527f616e2063616c6c20746869732066756e6374696f6e0000000000000000000000606482015260840161045f565b6000610a8e8383610e94565b6001600160a01b0384166000908152606a6020526040812054919250606b90610ab7868661087e565b815260208101919091526040016000205580610ad257505050565b6001600160a01b03831660009081526069602052604081208054839290610afa908490611d5b565b909155505060655460405163a9059cbb60e01b81526001600160a01b038481166004830152602482018490529091169063a9059cbb90604401602060405180830381600087803b158015610b4d57600080fd5b505af1158015610b61573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b859190611b0a565b50505050565b336000818152606860205260408082205460665491516313cdd31b60e01b81526004810194909452926001600160a01b03909116906313cdd31b9060240160206040518083038186803b158015610be157600080fd5b505afa158015610bf5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c199190611c90565b905080610c4d5733600090815260696020526040902054610c3a9083611d04565b3360009081526069602052604081205591505b60008211610c9d5760405162461bcd60e51b815260206004820152601060248201527f4e6f7468696e6720746f20636c61696d00000000000000000000000000000000604482015260640161045f565b3360008181526068602052604080822091909155606554905163a9059cbb60e01b81526004810192909252602482018490526001600160a01b03169063a9059cbb906044016109a9565b6000606b6000610cf7858561087e565b815260200190815260200160002054905092915050565b6033546001600160a01b03163314610d685760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161045f565b6001600160a01b03166000908152606d6020526040812055565b600054610100900460ff1680610d9b575060005460ff16155b610dfe5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161045f565b600054610100900460ff16158015610e20576000805461ffff19166101011790555b610e28611791565b606580546001600160a01b038087167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556067805485841690831617905560668054928616929091169190911790558015610b85576000805461ff001916905550505050565b6066546040517ff731cb490000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301528381166024830152600092839291169063f731cb499060440160606040518083038186803b158015610efe57600080fd5b505afa158015610f12573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f369190611c29565b8051909150610f495760009150506104a9565b600080606b6000610f5a888861087e565b8152602001908152602001600020546001610f759190611d04565b90505b606760009054906101000a90046001600160a01b03166001600160a01b031663b3e123db6040518163ffffffff1660e01b815260040160206040518083038186803b158015610fc657600080fd5b505afa158015610fda573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ffe9190611c90565b81101561104e576000606c6000611015848a6107ca565b8152602001908152602001600020905060008160010154111561103b578192505061104e565b508061104681611d72565b915050610f78565b508061105f576000925050506104a9565b6000606c600061106f84896107ca565b81526020019081526020016000209050600061108e8460000151611669565b6003830154909150600f90810b9081900b6110f4576110d2826110cd6110b78660000154611669565b6110c8866107248960020154611669565b611853565b61189b565b60038401549092507001000000000000000000000000000000009004600f0b90505b6001600160a01b0388166000908152606a6020526040812054606c90829061111c908c6107ca565b815260200190815260200160002090506000611155611150856110c88560030160109054906101000a9004600f0b87611687565b6118e0565b67ffffffffffffffff1690508660000151816111719190611d5b565b9a9950505050505050505050565b6111b76040518060a001604052806000815260200160008152602001600081526020016000600f0b81526020016000600f0b81525090565b606c60006111c585856107ca565b81526020808201929092526040908101600020815160a0810183528154815260018201549381019390935260028101549183019190915260030154600f81810b810b810b6060840152700100000000000000000000000000000000909104810b810b900b60808201529392505050565b6000606c600061124585856107ca565b815260200190815260200160002060020154905092915050565b6033546001600160a01b031633146112b95760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161045f565b6001600160a01b0381166113355760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161045f565b61133e81611727565b50565b606754604080517fb3e123db00000000000000000000000000000000000000000000000000000000815290516000926001600160a01b03169163b3e123db916004808301926020929190829003018186803b15801561139f57600080fd5b505afa1580156113b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113d79190611c90565b90506000606c60006113e984336107ca565b8152602001908152602001600020905080600101546000146114735760405162461bcd60e51b815260206004820152603160248201527f546865206e6578742072657761726420706f6f6c2068617320616c726561647960448201527f206265656e20696e697469616c697a6564000000000000000000000000000000606482015260840161045f565b6066546040516313cdd31b60e01b81523360048201526000916001600160a01b0316906313cdd31b9060240160206040518083038186803b1580156114b757600080fd5b505afa1580156114cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114ef9190611c90565b9050600081116115675760405162461bcd60e51b815260206004820152602c60248201527f4d7573742068617665207374616b6520746f20696e74697469616c697a65206160448201527f2072657761726420706f6f6c0000000000000000000000000000000000000000606482015260840161045f565b436001830155336000908152606960205260409020546115879082611d04565b6002830155336000818152606a6020526040812054606c926115a991906107ca565b81526020808201929092526040908101600090812060039081015495018054700100000000000000000000000000000000909604600f90810b900b6fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff0000000000000000000000000000000090961695909517909455338452606a9091529091209190915550565b600061271061165861ffff84166fffffffffffffffffffffffffffffffff8616611d3c565b6116629190611d1c565b9392505050565b6000677fffffffffffffff82111561168057600080fd5b5060401b90565b600081600f0b6000141561169a57600080fd5b600082600f0b604085600f0b901b816116c357634e487b7160e01b600052601260045260246000fd5b0590506f7fffffffffffffffffffffffffffffff1981128015906116f757506f7fffffffffffffffffffffffffffffff8113155b61166257600080fd5b600061171f846110cd866110c861171688611669565b61072488611669565b949350505050565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16806117aa575060005460ff16155b61180d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161045f565b600054610100900460ff1615801561182f576000805461ffff19166101011790555b6118376118fc565b61183f6119ad565b801561133e576000805461ff001916905550565b6000600f83810b9083900b0260401d6f7fffffffffffffffffffffffffffffff1981128015906116f757506f7fffffffffffffffffffffffffffffff81131561166257600080fd5b6000600f83810b9083900b016f7fffffffffffffffffffffffffffffff1981128015906116f757506f7fffffffffffffffffffffffffffffff81131561166257600080fd5b60008082600f0b12156118f257600080fd5b50600f0b60401d90565b600054610100900460ff1680611915575060005460ff16155b6119785760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161045f565b600054610100900460ff1615801561183f576000805461ffff1916610101179055801561133e576000805461ff001916905550565b600054610100900460ff16806119c6575060005460ff16155b611a295760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161045f565b600054610100900460ff16158015611a4b576000805461ffff19166101011790555b61183f33611727565b80516fffffffffffffffffffffffffffffffff81168114611a7457600080fd5b919050565b805161ffff81168114611a7457600080fd5b600060208284031215611a9c578081fd5b813561166281611da3565b60008060408385031215611ab9578081fd5b8235611ac481611da3565b91506020830135611ad481611da3565b809150509250929050565b60008060408385031215611af1578182fd5b8235611afc81611da3565b946020939093013593505050565b600060208284031215611b1b578081fd5b81518015158114611662578182fd5b600060208284031215611b3b578081fd5b5035919050565b600080600060608486031215611b56578081fd5b8335611b6181611da3565b92506020840135611b7181611da3565b91506040840135611b8181611da3565b809150509250925092565b60006101408284031215611b9e578081fd5b611ba6611ccc565b82518152602083015160208201526040830151604082015260608301516060820152611bd460808401611a79565b608082015260a083015160a0820152611bef60c08401611a54565b60c0820152611c0060e08401611a54565b60e08201526101008381015190820152610120611c1e818501611a79565b908201529392505050565b600060608284031215611c3a578081fd5b6040516060810181811067ffffffffffffffff82111715611c6957634e487b7160e01b83526041600452602483fd5b80604052508251815260208301516020820152604083015160408201528091505092915050565b600060208284031215611ca1578081fd5b5051919050565b60008060408385031215611cba578182fd5b823591506020830135611ad481611da3565b604051610140810167ffffffffffffffff81118282101715611cfe57634e487b7160e01b600052604160045260246000fd5b60405290565b60008219821115611d1757611d17611d8d565b500190565b600082611d3757634e487b7160e01b81526012600452602481fd5b500490565b6000816000190483118215151615611d5657611d56611d8d565b500290565b600082821015611d6d57611d6d611d8d565b500390565b6000600019821415611d8657611d86611d8d565b5060010190565b634e487b7160e01b600052601160045260246000fd5b6001600160a01b038116811461133e57600080fdfea26469706673582212209f089fd6db5b28f10a30e264589fd3480ab0e4b25299d1d10fca653572f31bea64736f6c63430008040033"

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
