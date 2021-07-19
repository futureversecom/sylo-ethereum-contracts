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

// SyloTokenABI is the input ABI used to generate the binding from.
const SyloTokenABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SyloTokenBin is the compiled bytecode used for deploying new contracts.
<<<<<<< HEAD
var SyloTokenBin = "0x60806040523480156200001157600080fd5b506040518060400160405280600481526020016353796c6f60e01b8152506040518060400160405280600481526020016353594c4f60e01b81525081600390805190602001906200006492919062000189565b5080516200007a90600490602084019062000189565b5050506200009b336b204fce5e3e25026110000000620000a160201b60201c565b62000291565b6001600160a01b038216620000fc5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b80600260008282546200011091906200022f565b90915550506001600160a01b038216600090815260208190526040812080548392906200013f9084906200022f565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b828054620001979062000254565b90600052602060002090601f016020900481019282620001bb576000855562000206565b82601f10620001d657805160ff191683800117855562000206565b8280016001018555821562000206579182015b8281111562000206578251825591602001919060010190620001e9565b506200021492915062000218565b5090565b5b8082111562000214576000815560010162000219565b600082198211156200024f57634e487b7160e01b81526011600452602481fd5b500190565b600181811c908216806200026957607f821691505b602082108114156200028b57634e487b7160e01b600052602260045260246000fd5b50919050565b61098680620002a16000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80633950935111610081578063a457c2d71161005b578063a457c2d714610187578063a9059cbb1461019a578063dd62ed3e146101ad57600080fd5b8063395093511461014357806370a082311461015657806395d89b411461017f57600080fd5b806318160ddd116100b257806318160ddd1461010f57806323b872dd14610121578063313ce5671461013457600080fd5b806306fdde03146100ce578063095ea7b3146100ec575b600080fd5b6100d66101e6565b6040516100e3919061089e565b60405180910390f35b6100ff6100fa366004610875565b610278565b60405190151581526020016100e3565b6002545b6040519081526020016100e3565b6100ff61012f36600461083a565b61028e565b604051601281526020016100e3565b6100ff610151366004610875565b610352565b6101136101643660046107e7565b6001600160a01b031660009081526020819052604090205490565b6100d661038e565b6100ff610195366004610875565b61039d565b6100ff6101a8366004610875565b61044e565b6101136101bb366004610808565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f590610915565b80601f016020809104026020016040519081016040528092919081815260200182805461022190610915565b801561026e5780601f106102435761010080835404028352916020019161026e565b820191906000526020600020905b81548152906001019060200180831161025157829003601f168201915b5050505050905090565b600061028533848461045b565b50600192915050565b600061029b8484846105b3565b6001600160a01b03841660009081526001602090815260408083203384529091529020548281101561033a5760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206160448201527f6c6c6f77616e636500000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b610347853385840361045b565b506001949350505050565b3360008181526001602090815260408083206001600160a01b038716845290915281205490916102859185906103899086906108f1565b61045b565b6060600480546101f590610915565b3360009081526001602090815260408083206001600160a01b0386168452909152812054828110156104375760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152608401610331565b610444338585840361045b565b5060019392505050565b60006102853384846105b3565b6001600160a01b0383166104d65760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610331565b6001600160a01b0382166105525760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610331565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6001600160a01b03831661062f5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610331565b6001600160a01b0382166106ab5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610331565b6001600160a01b0383166000908152602081905260409020548181101561073a5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610331565b6001600160a01b038085166000908152602081905260408082208585039055918516815290812080548492906107719084906108f1565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516107bd91815260200190565b60405180910390a350505050565b80356001600160a01b03811681146107e257600080fd5b919050565b6000602082840312156107f8578081fd5b610801826107cb565b9392505050565b6000806040838503121561081a578081fd5b610823836107cb565b9150610831602084016107cb565b90509250929050565b60008060006060848603121561084e578081fd5b610857846107cb565b9250610865602085016107cb565b9150604084013590509250925092565b60008060408385031215610887578182fd5b610890836107cb565b946020939093013593505050565b6000602080835283518082850152825b818110156108ca578581018301518582016040015282016108ae565b818111156108db5783604083870101525b50601f01601f1916929092016040019392505050565b6000821982111561091057634e487b7160e01b81526011600452602481fd5b500190565b600181811c9082168061092957607f821691505b6020821081141561094a57634e487b7160e01b600052602260045260246000fd5b5091905056fea26469706673582212202c406a30abdfc7a325854a97641350157865ee36746b3c719f497c395c79d8e564736f6c63430008040033"
=======
var SyloTokenBin = "0x60806040523480156200001157600080fd5b506040518060400160405280600481526020016353796c6f60e01b8152506040518060400160405280600481526020016353594c4f60e01b81525081600390805190602001906200006492919062000189565b5080516200007a90600490602084019062000189565b5050506200009b336b204fce5e3e25026110000000620000a160201b60201c565b62000291565b6001600160a01b038216620000fc5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b80600260008282546200011091906200022f565b90915550506001600160a01b038216600090815260208190526040812080548392906200013f9084906200022f565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b828054620001979062000254565b90600052602060002090601f016020900481019282620001bb576000855562000206565b82601f10620001d657805160ff191683800117855562000206565b8280016001018555821562000206579182015b8281111562000206578251825591602001919060010190620001e9565b506200021492915062000218565b5090565b5b8082111562000214576000815560010162000219565b600082198211156200024f57634e487b7160e01b81526011600452602481fd5b500190565b600181811c908216806200026957607f821691505b602082108114156200028b57634e487b7160e01b600052602260045260246000fd5b50919050565b6109b480620002a16000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80633950935111610081578063a457c2d71161005b578063a457c2d714610187578063a9059cbb1461019a578063dd62ed3e146101ad57600080fd5b8063395093511461014357806370a082311461015657806395d89b411461017f57600080fd5b806318160ddd116100b257806318160ddd1461010f57806323b872dd14610121578063313ce5671461013457600080fd5b806306fdde03146100ce578063095ea7b3146100ec575b600080fd5b6100d66101e6565b6040516100e391906108ab565b60405180910390f35b6100ff6100fa366004610882565b610278565b60405190151581526020016100e3565b6002545b6040519081526020016100e3565b6100ff61012f366004610847565b61028e565b604051601281526020016100e3565b6100ff610151366004610882565b610359565b6101136101643660046107f4565b6001600160a01b031660009081526020819052604090205490565b6100d6610390565b6100ff610195366004610882565b61039f565b6100ff6101a8366004610882565b610452565b6101136101bb366004610815565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f59061092d565b80601f01602080910402602001604051908101604052809291908181526020018280546102219061092d565b801561026e5780601f106102435761010080835404028352916020019161026e565b820191906000526020600020905b81548152906001019060200180831161025157829003601f168201915b5050505050905090565b600061028533848461045f565b50600192915050565b600061029b8484846105b7565b6001600160a01b03841660009081526001602090815260408083203384529091529020548281101561033a5760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206160448201527f6c6c6f77616e636500000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b61034e85336103498685610916565b61045f565b506001949350505050565b3360008181526001602090815260408083206001600160a01b038716845290915281205490916102859185906103499086906108fe565b6060600480546101f59061092d565b3360009081526001602090815260408083206001600160a01b0386168452909152812054828110156104395760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152608401610331565b61044833856103498685610916565b5060019392505050565b60006102853384846105b7565b6001600160a01b0383166104da5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610331565b6001600160a01b0382166105565760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610331565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6001600160a01b0383166106335760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610331565b6001600160a01b0382166106af5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610331565b6001600160a01b0383166000908152602081905260409020548181101561073e5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610331565b6107488282610916565b6001600160a01b03808616600090815260208190526040808220939093559085168152908120805484929061077e9084906108fe565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516107ca91815260200190565b60405180910390a350505050565b80356001600160a01b03811681146107ef57600080fd5b919050565b600060208284031215610805578081fd5b61080e826107d8565b9392505050565b60008060408385031215610827578081fd5b610830836107d8565b915061083e602084016107d8565b90509250929050565b60008060006060848603121561085b578081fd5b610864846107d8565b9250610872602085016107d8565b9150604084013590509250925092565b60008060408385031215610894578182fd5b61089d836107d8565b946020939093013593505050565b6000602080835283518082850152825b818110156108d7578581018301518582016040015282016108bb565b818111156108e85783604083870101525b50601f01601f1916929092016040019392505050565b6000821982111561091157610911610968565b500190565b60008282101561092857610928610968565b500390565b600181811c9082168061094157607f821691505b6020821081141561096257634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fdfea26469706673582212202054608209602a26ee1c22703417dc09732083061fa54710a26a562b0886e17964736f6c63430008040033"
>>>>>>> 3746208 (extract faceValue and winprob to onchain values)

// DeploySyloToken deploys a new Ethereum contract, binding an instance of SyloToken to it.
func DeploySyloToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SyloToken, error) {
	parsed, err := abi.JSON(strings.NewReader(SyloTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SyloTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SyloToken{SyloTokenCaller: SyloTokenCaller{contract: contract}, SyloTokenTransactor: SyloTokenTransactor{contract: contract}, SyloTokenFilterer: SyloTokenFilterer{contract: contract}}, nil
}

// SyloToken is an auto generated Go binding around an Ethereum contract.
type SyloToken struct {
	SyloTokenCaller     // Read-only binding to the contract
	SyloTokenTransactor // Write-only binding to the contract
	SyloTokenFilterer   // Log filterer for contract events
}

// SyloTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SyloTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyloTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SyloTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyloTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SyloTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyloTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SyloTokenSession struct {
	Contract     *SyloToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SyloTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SyloTokenCallerSession struct {
	Contract *SyloTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SyloTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SyloTokenTransactorSession struct {
	Contract     *SyloTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SyloTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SyloTokenRaw struct {
	Contract *SyloToken // Generic contract binding to access the raw methods on
}

// SyloTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SyloTokenCallerRaw struct {
	Contract *SyloTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SyloTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SyloTokenTransactorRaw struct {
	Contract *SyloTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSyloToken creates a new instance of SyloToken, bound to a specific deployed contract.
func NewSyloToken(address common.Address, backend bind.ContractBackend) (*SyloToken, error) {
	contract, err := bindSyloToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SyloToken{SyloTokenCaller: SyloTokenCaller{contract: contract}, SyloTokenTransactor: SyloTokenTransactor{contract: contract}, SyloTokenFilterer: SyloTokenFilterer{contract: contract}}, nil
}

// NewSyloTokenCaller creates a new read-only instance of SyloToken, bound to a specific deployed contract.
func NewSyloTokenCaller(address common.Address, caller bind.ContractCaller) (*SyloTokenCaller, error) {
	contract, err := bindSyloToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SyloTokenCaller{contract: contract}, nil
}

// NewSyloTokenTransactor creates a new write-only instance of SyloToken, bound to a specific deployed contract.
func NewSyloTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SyloTokenTransactor, error) {
	contract, err := bindSyloToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SyloTokenTransactor{contract: contract}, nil
}

// NewSyloTokenFilterer creates a new log filterer instance of SyloToken, bound to a specific deployed contract.
func NewSyloTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*SyloTokenFilterer, error) {
	contract, err := bindSyloToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SyloTokenFilterer{contract: contract}, nil
}

// bindSyloToken binds a generic wrapper to an already deployed contract.
func bindSyloToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SyloTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SyloToken *SyloTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SyloToken.Contract.SyloTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SyloToken *SyloTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyloToken.Contract.SyloTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SyloToken *SyloTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SyloToken.Contract.SyloTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SyloToken *SyloTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SyloToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SyloToken *SyloTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyloToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SyloToken *SyloTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SyloToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SyloToken *SyloTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SyloToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SyloToken *SyloTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SyloToken.Contract.Allowance(&_SyloToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SyloToken *SyloTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SyloToken.Contract.Allowance(&_SyloToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SyloToken *SyloTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SyloToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SyloToken *SyloTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SyloToken.Contract.BalanceOf(&_SyloToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SyloToken *SyloTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SyloToken.Contract.BalanceOf(&_SyloToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SyloToken *SyloTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SyloToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SyloToken *SyloTokenSession) Decimals() (uint8, error) {
	return _SyloToken.Contract.Decimals(&_SyloToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SyloToken *SyloTokenCallerSession) Decimals() (uint8, error) {
	return _SyloToken.Contract.Decimals(&_SyloToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SyloToken *SyloTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SyloToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SyloToken *SyloTokenSession) Name() (string, error) {
	return _SyloToken.Contract.Name(&_SyloToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SyloToken *SyloTokenCallerSession) Name() (string, error) {
	return _SyloToken.Contract.Name(&_SyloToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SyloToken *SyloTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SyloToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SyloToken *SyloTokenSession) Symbol() (string, error) {
	return _SyloToken.Contract.Symbol(&_SyloToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SyloToken *SyloTokenCallerSession) Symbol() (string, error) {
	return _SyloToken.Contract.Symbol(&_SyloToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SyloToken *SyloTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyloToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SyloToken *SyloTokenSession) TotalSupply() (*big.Int, error) {
	return _SyloToken.Contract.TotalSupply(&_SyloToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SyloToken *SyloTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _SyloToken.Contract.TotalSupply(&_SyloToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SyloToken *SyloTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SyloToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SyloToken *SyloTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SyloToken.Contract.Approve(&_SyloToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SyloToken *SyloTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SyloToken.Contract.Approve(&_SyloToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SyloToken *SyloTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SyloToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SyloToken *SyloTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SyloToken.Contract.DecreaseAllowance(&_SyloToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SyloToken *SyloTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SyloToken.Contract.DecreaseAllowance(&_SyloToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SyloToken *SyloTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SyloToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SyloToken *SyloTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SyloToken.Contract.IncreaseAllowance(&_SyloToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SyloToken *SyloTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SyloToken.Contract.IncreaseAllowance(&_SyloToken.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SyloToken *SyloTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SyloToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SyloToken *SyloTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SyloToken.Contract.Transfer(&_SyloToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SyloToken *SyloTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SyloToken.Contract.Transfer(&_SyloToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SyloToken *SyloTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SyloToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SyloToken *SyloTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SyloToken.Contract.TransferFrom(&_SyloToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SyloToken *SyloTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SyloToken.Contract.TransferFrom(&_SyloToken.TransactOpts, sender, recipient, amount)
}

// SyloTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SyloToken contract.
type SyloTokenApprovalIterator struct {
	Event *SyloTokenApproval // Event containing the contract specifics and raw log

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
func (it *SyloTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyloTokenApproval)
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
		it.Event = new(SyloTokenApproval)
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
func (it *SyloTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyloTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyloTokenApproval represents a Approval event raised by the SyloToken contract.
type SyloTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SyloToken *SyloTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SyloTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SyloToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SyloTokenApprovalIterator{contract: _SyloToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SyloToken *SyloTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SyloTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SyloToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyloTokenApproval)
				if err := _SyloToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SyloToken *SyloTokenFilterer) ParseApproval(log types.Log) (*SyloTokenApproval, error) {
	event := new(SyloTokenApproval)
	if err := _SyloToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyloTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SyloToken contract.
type SyloTokenTransferIterator struct {
	Event *SyloTokenTransfer // Event containing the contract specifics and raw log

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
func (it *SyloTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyloTokenTransfer)
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
		it.Event = new(SyloTokenTransfer)
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
func (it *SyloTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyloTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyloTokenTransfer represents a Transfer event raised by the SyloToken contract.
type SyloTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SyloToken *SyloTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SyloTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SyloToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SyloTokenTransferIterator{contract: _SyloToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SyloToken *SyloTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SyloTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SyloToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyloTokenTransfer)
				if err := _SyloToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SyloToken *SyloTokenFilterer) ParseTransfer(log types.Log) (*SyloTokenTransfer, error) {
	event := new(SyloTokenTransfer)
	if err := _SyloToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
