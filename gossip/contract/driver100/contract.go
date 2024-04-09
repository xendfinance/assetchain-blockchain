// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package driver100

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
	_ = abi.ConvertType
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"AdvanceEpochs\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"diff\",\"type\":\"bytes\"}],\"name\":\"UpdateNetworkRules\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"UpdateNetworkVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"UpdateValidatorPubkey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"UpdateValidatorWeight\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"backend\",\"type\":\"address\"}],\"name\":\"UpdatedBackend\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_backend\",\"type\":\"address\"}],\"name\":\"setBackend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_backend\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_evmWriterAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"copyCode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"with\",\"type\":\"address\"}],\"name\":\"swapCode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"diff\",\"type\":\"uint256\"}],\"name\":\"incNonce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"diff\",\"type\":\"bytes\"}],\"name\":\"updateNetworkRules\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateNetworkVersion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"advanceEpochs\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"updateValidatorWeight\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"updateValidatorPubkey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"}],\"name\":\"setGenesisValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupFromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earlyUnlockPenalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"setGenesisDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"deactivateValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nextValidatorIDs\",\"type\":\"uint256[]\"}],\"name\":\"sealEpochValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"offlineTimes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"offlineBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"uptimes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"originatedTxsFee\",\"type\":\"uint256[]\"}],\"name\":\"sealEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"offlineTimes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"offlineBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"uptimes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"originatedTxsFee\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"usedGas\",\"type\":\"uint256\"}],\"name\":\"sealEpochV1\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612543806100206000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c806379bead38116100a2578063d6a0c7af11610071578063d6a0c7af1461073e578063da7fc24f146107a2578063e08d7e66146107e6578063e30443bc1461085f578063ebdf104c146108ad57610116565b806379bead38146104bd5780637f52e13e1461050b578063a4066fbe1461068d578063b9cc6b1c146106c557610116565b8063242a6e3f116100e9578063242a6e3f1461027a578063267ab446146102fd57806339e503ab1461032b578063485cc955146103835780634feb92f3146103e757610116565b806307690b2a1461011b5780630aeeca001461017f57806318f628d4146101ad5780631e702f8314610242575b600080fd5b61017d6004803603604081101561013157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a25565b005b6101ab6004803603602081101561019557600080fd5b8101908080359060200190929190505050610bd9565b005b61024060048036036101208110156101c457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190505050610cd6565b005b6102786004803603604081101561025857600080fd5b810190808035906020019092919080359060200190929190505050610e7c565b005b6102fb6004803603604081101561029057600080fd5b8101908080359060200190929190803590602001906401000000008111156102b757600080fd5b8201836020820111156102c957600080fd5b803590602001918460018302840111640100000000831117156102eb57600080fd5b9091929391929390505050610fb7565b005b6103296004803603602081101561031357600080fd5b81019080803590602001909291905050506110e3565b005b6103816004803603606081101561034157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291905050506111e0565b005b6103e56004803603604081101561039957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611371565b005b6104bb60048036036101008110156103fe57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561044557600080fd5b82018360208201111561045757600080fd5b8035906020019184600183028401116401000000008311171561047957600080fd5b90919293919293908035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190505050611537565b005b610509600480360360408110156104d357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611701565b005b61068b600480360360a081101561052157600080fd5b810190808035906020019064010000000081111561053e57600080fd5b82018360208201111561055057600080fd5b8035906020019184602083028401116401000000008311171561057257600080fd5b90919293919293908035906020019064010000000081111561059357600080fd5b8201836020820111156105a557600080fd5b803590602001918460208302840111640100000000831117156105c757600080fd5b9091929391929390803590602001906401000000008111156105e857600080fd5b8201836020820111156105fa57600080fd5b8035906020019184602083028401116401000000008311171561061c57600080fd5b90919293919293908035906020019064010000000081111561063d57600080fd5b82018360208201111561064f57600080fd5b8035906020019184602083028401116401000000008311171561067157600080fd5b909192939192939080359060200190929190505050611889565b005b6106c3600480360360408110156106a357600080fd5b810190808035906020019092919080359060200190929190505050611a9f565b005b61073c600480360360208110156106db57600080fd5b81019080803590602001906401000000008111156106f857600080fd5b82018360208201111561070a57600080fd5b8035906020019184600183028401116401000000008311171561072c57600080fd5b9091929391929390505050611b9e565b005b6107a06004803603604081101561075457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611cc8565b005b6107e4600480360360208110156107b857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611e7c565b005b61085d600480360360208110156107fc57600080fd5b810190808035906020019064010000000081111561081957600080fd5b82018360208201111561082b57600080fd5b8035906020019184602083028401116401000000008311171561084d57600080fd5b9091929391929390505050611fc6565b005b6108ab6004803603604081101561087557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050612128565b005b610a23600480360360808110156108c357600080fd5b81019080803590602001906401000000008111156108e057600080fd5b8201836020820111156108f257600080fd5b8035906020019184602083028401116401000000008311171561091457600080fd5b90919293919293908035906020019064010000000081111561093557600080fd5b82018360208201111561094757600080fd5b8035906020019184602083028401116401000000008311171561096957600080fd5b90919293919293908035906020019064010000000081111561098a57600080fd5b82018360208201111561099c57600080fd5b803590602001918460208302840111640100000000831117156109be57600080fd5b9091929391929390803590602001906401000000008111156109df57600080fd5b8201836020820111156109f157600080fd5b80359060200191846020830284011164010000000083111715610a1357600080fd5b90919293919293905050506122b0565b005b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ae8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166307690b2a83836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b158015610bbd57600080fd5b505af1158015610bd1573d6000803e3d6000fd5b505050505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c9c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b7f0151256d62457b809bbc891b1f81c6dd0b9987552c70ce915b519750cd434dd1816040518082815260200191505060405180910390a150565b600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d78576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f6e6f742063616c6c61626c65000000000000000000000000000000000000000081525060200191505060405180910390fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318f628d48a8a8a8a8a8a8a8a8a6040518a63ffffffff1660e01b8152600401808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018881526020018781526020018681526020018581526020018481526020018381526020018281526020019950505050505050505050600060405180830381600087803b158015610e5957600080fd5b505af1158015610e6d573d6000803e3d6000fd5b50505050505050505050505050565b600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f1e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f6e6f742063616c6c61626c65000000000000000000000000000000000000000081525060200191505060405180910390fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631e702f8383836040518363ffffffff1660e01b81526004018083815260200182815260200192505050600060405180830381600087803b158015610f9b57600080fd5b505af1158015610faf573d6000803e3d6000fd5b505050505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461107a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b827f0f0ef1ab97439def0a9d2c6d9dc166207f1b13b99e62b442b2993d6153c63a6e838360405180806020018281038252848482818152602001925080828437600081840152601f19601f820116905080830192505050935050505060405180910390a2505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146111a6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b7f2ccdfd47cf0c1f1069d949f1789bb79b2f12821f021634fc835af1de66ea2feb816040518082815260200191505060405180910390a150565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146112a3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166339e503ab8484846040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019350505050600060405180830381600087803b15801561135457600080fd5b505af1158015611368573d6000803e3d6000fd5b50505050505050565b600060019054906101000a900460ff1680611390575061138f6124c9565b5b806113a757506000809054906101000a900460ff16155b6113fc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e8152602001806124e1602e913960400191505060405180910390fd5b60008060019054906101000a900460ff16159050801561144c576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b82603460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff167f64ee8f7bfc37fc205d7194ee3d64947ab7b57e663cd0d1abd3ef24503583069360405160405180910390a281603560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080156115325760008060016101000a81548160ff0219169083151502179055505b505050565b600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146115d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f6e6f742063616c6c61626c65000000000000000000000000000000000000000081525060200191505060405180910390fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634feb92f38a8a8a8a8a8a8a8a8a6040518a63ffffffff1660e01b8152600401808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001898152602001806020018781526020018681526020018581526020018481526020018381526020018281038252898982818152602001925080828437600081840152601f19601f8201169050808301925050509a5050505050505050505050600060405180830381600087803b1580156116de57600080fd5b505af11580156116f2573d6000803e3d6000fd5b50505050505050505050505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146117c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166379bead3883836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561186d57600080fd5b505af1158015611881573d6000803e3d6000fd5b505050505050565b600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461192b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f6e6f742063616c6c61626c65000000000000000000000000000000000000000081525060200191505060405180910390fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663592fe0c08a8a8a8a8a8a8a8a8a6040518a63ffffffff1660e01b8152600401808060200180602001806020018060200186815260200185810385528e8e82818152602001925060200280828437600081840152601f19601f82011690508083019250505085810384528c8c82818152602001925060200280828437600081840152601f19601f82011690508083019250505085810383528a8a82818152602001925060200280828437600081840152601f19601f8201169050808301925050508581038252888882818152602001925060200280828437600081840152601f19601f8201169050808301925050509d5050505050505050505050505050600060405180830381600087803b158015611a7c57600080fd5b505af1158015611a90573d6000803e3d6000fd5b50505050505050505050505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611b62576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b817fb975807576e3b1461be7de07ebf7d20e4790ed802d7a0c4fdd0a1a13df72a935826040518082815260200191505060405180910390a25050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611c61576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b7f47d10eed096a44e3d0abc586c7e3a5d6cb5358cc90e7d437cd0627f7e765fb99828260405180806020018281038252848482818152602001925080828437600081840152601f19601f820116905080830192505050935050505060405180910390a15050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611d8b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d6a0c7af83836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b158015611e6057600080fd5b505af1158015611e74573d6000803e3d6000fd5b505050505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611f3f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff167f64ee8f7bfc37fc205d7194ee3d64947ab7b57e663cd0d1abd3ef24503583069360405160405180910390a280603460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612068576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f6e6f742063616c6c61626c65000000000000000000000000000000000000000081525060200191505060405180910390fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e08d7e6683836040518363ffffffff1660e01b815260040180806020018281038252848482818152602001925060200280828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b15801561210c57600080fd5b505af1158015612120573d6000803e3d6000fd5b505050505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146121eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e30443bc83836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561229457600080fd5b505af11580156122a8573d6000803e3d6000fd5b505050505050565b600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612352576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f6e6f742063616c6c61626c65000000000000000000000000000000000000000081525060200191505060405180910390fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663592fe0c0898989898989898963322adc3a6040518a63ffffffff1660e01b8152600401808060200180602001806020018060200186815260200185810385528e8e82818152602001925060200280828437600081840152601f19601f82011690508083019250505085810384528c8c82818152602001925060200280828437600081840152601f19601f82011690508083019250505085810383528a8a82818152602001925060200280828437600081840152601f19601f8201169050808301925050508581038252888882818152602001925060200280828437600081840152601f19601f8201169050808301925050509d5050505050505050505050505050600060405180830381600087803b1580156124a757600080fd5b505af11580156124bb573d6000803e3d6000fd5b505050505050505050505050565b6000803090506000813b905060008114925050509056fe436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564a265627a7a723158200e3b4f3f8ff8f9f35bf6a4819537ff47f280a0f7820eaeed0a40415efda42d5964736f6c63430005100032",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// AdvanceEpochs is a paid mutator transaction binding the contract method 0x0aeeca00.
//
// Solidity: function advanceEpochs(uint256 num) returns()
func (_Contract *ContractTransactor) AdvanceEpochs(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "advanceEpochs", num)
}

// AdvanceEpochs is a paid mutator transaction binding the contract method 0x0aeeca00.
//
// Solidity: function advanceEpochs(uint256 num) returns()
func (_Contract *ContractSession) AdvanceEpochs(num *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AdvanceEpochs(&_Contract.TransactOpts, num)
}

// AdvanceEpochs is a paid mutator transaction binding the contract method 0x0aeeca00.
//
// Solidity: function advanceEpochs(uint256 num) returns()
func (_Contract *ContractTransactorSession) AdvanceEpochs(num *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AdvanceEpochs(&_Contract.TransactOpts, num)
}

// CopyCode is a paid mutator transaction binding the contract method 0xd6a0c7af.
//
// Solidity: function copyCode(address acc, address from) returns()
func (_Contract *ContractTransactor) CopyCode(opts *bind.TransactOpts, acc common.Address, from common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "copyCode", acc, from)
}

// CopyCode is a paid mutator transaction binding the contract method 0xd6a0c7af.
//
// Solidity: function copyCode(address acc, address from) returns()
func (_Contract *ContractSession) CopyCode(acc common.Address, from common.Address) (*types.Transaction, error) {
	return _Contract.Contract.CopyCode(&_Contract.TransactOpts, acc, from)
}

// CopyCode is a paid mutator transaction binding the contract method 0xd6a0c7af.
//
// Solidity: function copyCode(address acc, address from) returns()
func (_Contract *ContractTransactorSession) CopyCode(acc common.Address, from common.Address) (*types.Transaction, error) {
	return _Contract.Contract.CopyCode(&_Contract.TransactOpts, acc, from)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
//
// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
func (_Contract *ContractTransactor) DeactivateValidator(opts *bind.TransactOpts, validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deactivateValidator", validatorID, status)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
//
// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
func (_Contract *ContractSession) DeactivateValidator(validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DeactivateValidator(&_Contract.TransactOpts, validatorID, status)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
//
// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
func (_Contract *ContractTransactorSession) DeactivateValidator(validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DeactivateValidator(&_Contract.TransactOpts, validatorID, status)
}

// IncNonce is a paid mutator transaction binding the contract method 0x79bead38.
//
// Solidity: function incNonce(address acc, uint256 diff) returns()
func (_Contract *ContractTransactor) IncNonce(opts *bind.TransactOpts, acc common.Address, diff *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "incNonce", acc, diff)
}

// IncNonce is a paid mutator transaction binding the contract method 0x79bead38.
//
// Solidity: function incNonce(address acc, uint256 diff) returns()
func (_Contract *ContractSession) IncNonce(acc common.Address, diff *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.IncNonce(&_Contract.TransactOpts, acc, diff)
}

// IncNonce is a paid mutator transaction binding the contract method 0x79bead38.
//
// Solidity: function incNonce(address acc, uint256 diff) returns()
func (_Contract *ContractTransactorSession) IncNonce(acc common.Address, diff *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.IncNonce(&_Contract.TransactOpts, acc, diff)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _backend, address _evmWriterAddress) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, _backend common.Address, _evmWriterAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", _backend, _evmWriterAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _backend, address _evmWriterAddress) returns()
func (_Contract *ContractSession) Initialize(_backend common.Address, _evmWriterAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _backend, _evmWriterAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _backend, address _evmWriterAddress) returns()
func (_Contract *ContractTransactorSession) Initialize(_backend common.Address, _evmWriterAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _backend, _evmWriterAddress)
}

// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
//
// Solidity: function sealEpoch(uint256[] offlineTimes, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
func (_Contract *ContractTransactor) SealEpoch(opts *bind.TransactOpts, offlineTimes []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sealEpoch", offlineTimes, offlineBlocks, uptimes, originatedTxsFee)
}

// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
//
// Solidity: function sealEpoch(uint256[] offlineTimes, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
func (_Contract *ContractSession) SealEpoch(offlineTimes []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SealEpoch(&_Contract.TransactOpts, offlineTimes, offlineBlocks, uptimes, originatedTxsFee)
}

// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
//
// Solidity: function sealEpoch(uint256[] offlineTimes, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
func (_Contract *ContractTransactorSession) SealEpoch(offlineTimes []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SealEpoch(&_Contract.TransactOpts, offlineTimes, offlineBlocks, uptimes, originatedTxsFee)
}

// SealEpochV1 is a paid mutator transaction binding the contract method 0x7f52e13e.
//
// Solidity: function sealEpochV1(uint256[] offlineTimes, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee, uint256 usedGas) returns()
func (_Contract *ContractTransactor) SealEpochV1(opts *bind.TransactOpts, offlineTimes []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int, usedGas *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sealEpochV1", offlineTimes, offlineBlocks, uptimes, originatedTxsFee, usedGas)
}

// SealEpochV1 is a paid mutator transaction binding the contract method 0x7f52e13e.
//
// Solidity: function sealEpochV1(uint256[] offlineTimes, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee, uint256 usedGas) returns()
func (_Contract *ContractSession) SealEpochV1(offlineTimes []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int, usedGas *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SealEpochV1(&_Contract.TransactOpts, offlineTimes, offlineBlocks, uptimes, originatedTxsFee, usedGas)
}

// SealEpochV1 is a paid mutator transaction binding the contract method 0x7f52e13e.
//
// Solidity: function sealEpochV1(uint256[] offlineTimes, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee, uint256 usedGas) returns()
func (_Contract *ContractTransactorSession) SealEpochV1(offlineTimes []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int, usedGas *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SealEpochV1(&_Contract.TransactOpts, offlineTimes, offlineBlocks, uptimes, originatedTxsFee, usedGas)
}

// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
//
// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
func (_Contract *ContractTransactor) SealEpochValidators(opts *bind.TransactOpts, nextValidatorIDs []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sealEpochValidators", nextValidatorIDs)
}

// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
//
// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
func (_Contract *ContractSession) SealEpochValidators(nextValidatorIDs []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SealEpochValidators(&_Contract.TransactOpts, nextValidatorIDs)
}

// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
//
// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
func (_Contract *ContractTransactorSession) SealEpochValidators(nextValidatorIDs []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SealEpochValidators(&_Contract.TransactOpts, nextValidatorIDs)
}

// SetBackend is a paid mutator transaction binding the contract method 0xda7fc24f.
//
// Solidity: function setBackend(address _backend) returns()
func (_Contract *ContractTransactor) SetBackend(opts *bind.TransactOpts, _backend common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setBackend", _backend)
}

// SetBackend is a paid mutator transaction binding the contract method 0xda7fc24f.
//
// Solidity: function setBackend(address _backend) returns()
func (_Contract *ContractSession) SetBackend(_backend common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetBackend(&_Contract.TransactOpts, _backend)
}

// SetBackend is a paid mutator transaction binding the contract method 0xda7fc24f.
//
// Solidity: function setBackend(address _backend) returns()
func (_Contract *ContractTransactorSession) SetBackend(_backend common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetBackend(&_Contract.TransactOpts, _backend)
}

// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
//
// Solidity: function setBalance(address acc, uint256 value) returns()
func (_Contract *ContractTransactor) SetBalance(opts *bind.TransactOpts, acc common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setBalance", acc, value)
}

// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
//
// Solidity: function setBalance(address acc, uint256 value) returns()
func (_Contract *ContractSession) SetBalance(acc common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetBalance(&_Contract.TransactOpts, acc, value)
}

// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
//
// Solidity: function setBalance(address acc, uint256 value) returns()
func (_Contract *ContractTransactorSession) SetBalance(acc common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetBalance(&_Contract.TransactOpts, acc, value)
}

// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
//
// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
func (_Contract *ContractTransactor) SetGenesisDelegation(opts *bind.TransactOpts, delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setGenesisDelegation", delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
}

// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
//
// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
func (_Contract *ContractSession) SetGenesisDelegation(delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisDelegation(&_Contract.TransactOpts, delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
}

// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
//
// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
func (_Contract *ContractTransactorSession) SetGenesisDelegation(delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisDelegation(&_Contract.TransactOpts, delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
}

// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
//
// Solidity: function setGenesisValidator(address _auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
func (_Contract *ContractTransactor) SetGenesisValidator(opts *bind.TransactOpts, _auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setGenesisValidator", _auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
}

// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
//
// Solidity: function setGenesisValidator(address _auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
func (_Contract *ContractSession) SetGenesisValidator(_auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisValidator(&_Contract.TransactOpts, _auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
}

// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
//
// Solidity: function setGenesisValidator(address _auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
func (_Contract *ContractTransactorSession) SetGenesisValidator(_auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisValidator(&_Contract.TransactOpts, _auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
}

// SetStorage is a paid mutator transaction binding the contract method 0x39e503ab.
//
// Solidity: function setStorage(address acc, bytes32 key, bytes32 value) returns()
func (_Contract *ContractTransactor) SetStorage(opts *bind.TransactOpts, acc common.Address, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setStorage", acc, key, value)
}

// SetStorage is a paid mutator transaction binding the contract method 0x39e503ab.
//
// Solidity: function setStorage(address acc, bytes32 key, bytes32 value) returns()
func (_Contract *ContractSession) SetStorage(acc common.Address, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SetStorage(&_Contract.TransactOpts, acc, key, value)
}

// SetStorage is a paid mutator transaction binding the contract method 0x39e503ab.
//
// Solidity: function setStorage(address acc, bytes32 key, bytes32 value) returns()
func (_Contract *ContractTransactorSession) SetStorage(acc common.Address, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SetStorage(&_Contract.TransactOpts, acc, key, value)
}

// SwapCode is a paid mutator transaction binding the contract method 0x07690b2a.
//
// Solidity: function swapCode(address acc, address with) returns()
func (_Contract *ContractTransactor) SwapCode(opts *bind.TransactOpts, acc common.Address, with common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "swapCode", acc, with)
}

// SwapCode is a paid mutator transaction binding the contract method 0x07690b2a.
//
// Solidity: function swapCode(address acc, address with) returns()
func (_Contract *ContractSession) SwapCode(acc common.Address, with common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SwapCode(&_Contract.TransactOpts, acc, with)
}

// SwapCode is a paid mutator transaction binding the contract method 0x07690b2a.
//
// Solidity: function swapCode(address acc, address with) returns()
func (_Contract *ContractTransactorSession) SwapCode(acc common.Address, with common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SwapCode(&_Contract.TransactOpts, acc, with)
}

// UpdateNetworkRules is a paid mutator transaction binding the contract method 0xb9cc6b1c.
//
// Solidity: function updateNetworkRules(bytes diff) returns()
func (_Contract *ContractTransactor) UpdateNetworkRules(opts *bind.TransactOpts, diff []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateNetworkRules", diff)
}

// UpdateNetworkRules is a paid mutator transaction binding the contract method 0xb9cc6b1c.
//
// Solidity: function updateNetworkRules(bytes diff) returns()
func (_Contract *ContractSession) UpdateNetworkRules(diff []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateNetworkRules(&_Contract.TransactOpts, diff)
}

// UpdateNetworkRules is a paid mutator transaction binding the contract method 0xb9cc6b1c.
//
// Solidity: function updateNetworkRules(bytes diff) returns()
func (_Contract *ContractTransactorSession) UpdateNetworkRules(diff []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateNetworkRules(&_Contract.TransactOpts, diff)
}

// UpdateNetworkVersion is a paid mutator transaction binding the contract method 0x267ab446.
//
// Solidity: function updateNetworkVersion(uint256 version) returns()
func (_Contract *ContractTransactor) UpdateNetworkVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateNetworkVersion", version)
}

// UpdateNetworkVersion is a paid mutator transaction binding the contract method 0x267ab446.
//
// Solidity: function updateNetworkVersion(uint256 version) returns()
func (_Contract *ContractSession) UpdateNetworkVersion(version *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateNetworkVersion(&_Contract.TransactOpts, version)
}

// UpdateNetworkVersion is a paid mutator transaction binding the contract method 0x267ab446.
//
// Solidity: function updateNetworkVersion(uint256 version) returns()
func (_Contract *ContractTransactorSession) UpdateNetworkVersion(version *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateNetworkVersion(&_Contract.TransactOpts, version)
}

// UpdateValidatorPubkey is a paid mutator transaction binding the contract method 0x242a6e3f.
//
// Solidity: function updateValidatorPubkey(uint256 validatorID, bytes pubkey) returns()
func (_Contract *ContractTransactor) UpdateValidatorPubkey(opts *bind.TransactOpts, validatorID *big.Int, pubkey []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateValidatorPubkey", validatorID, pubkey)
}

// UpdateValidatorPubkey is a paid mutator transaction binding the contract method 0x242a6e3f.
//
// Solidity: function updateValidatorPubkey(uint256 validatorID, bytes pubkey) returns()
func (_Contract *ContractSession) UpdateValidatorPubkey(validatorID *big.Int, pubkey []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateValidatorPubkey(&_Contract.TransactOpts, validatorID, pubkey)
}

// UpdateValidatorPubkey is a paid mutator transaction binding the contract method 0x242a6e3f.
//
// Solidity: function updateValidatorPubkey(uint256 validatorID, bytes pubkey) returns()
func (_Contract *ContractTransactorSession) UpdateValidatorPubkey(validatorID *big.Int, pubkey []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateValidatorPubkey(&_Contract.TransactOpts, validatorID, pubkey)
}

// UpdateValidatorWeight is a paid mutator transaction binding the contract method 0xa4066fbe.
//
// Solidity: function updateValidatorWeight(uint256 validatorID, uint256 value) returns()
func (_Contract *ContractTransactor) UpdateValidatorWeight(opts *bind.TransactOpts, validatorID *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateValidatorWeight", validatorID, value)
}

// UpdateValidatorWeight is a paid mutator transaction binding the contract method 0xa4066fbe.
//
// Solidity: function updateValidatorWeight(uint256 validatorID, uint256 value) returns()
func (_Contract *ContractSession) UpdateValidatorWeight(validatorID *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateValidatorWeight(&_Contract.TransactOpts, validatorID, value)
}

// UpdateValidatorWeight is a paid mutator transaction binding the contract method 0xa4066fbe.
//
// Solidity: function updateValidatorWeight(uint256 validatorID, uint256 value) returns()
func (_Contract *ContractTransactorSession) UpdateValidatorWeight(validatorID *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateValidatorWeight(&_Contract.TransactOpts, validatorID, value)
}

// ContractAdvanceEpochsIterator is returned from FilterAdvanceEpochs and is used to iterate over the raw logs and unpacked data for AdvanceEpochs events raised by the Contract contract.
type ContractAdvanceEpochsIterator struct {
	Event *ContractAdvanceEpochs // Event containing the contract specifics and raw log

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
func (it *ContractAdvanceEpochsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAdvanceEpochs)
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
		it.Event = new(ContractAdvanceEpochs)
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
func (it *ContractAdvanceEpochsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAdvanceEpochsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAdvanceEpochs represents a AdvanceEpochs event raised by the Contract contract.
type ContractAdvanceEpochs struct {
	Num *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAdvanceEpochs is a free log retrieval operation binding the contract event 0x0151256d62457b809bbc891b1f81c6dd0b9987552c70ce915b519750cd434dd1.
//
// Solidity: event AdvanceEpochs(uint256 num)
func (_Contract *ContractFilterer) FilterAdvanceEpochs(opts *bind.FilterOpts) (*ContractAdvanceEpochsIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AdvanceEpochs")
	if err != nil {
		return nil, err
	}
	return &ContractAdvanceEpochsIterator{contract: _Contract.contract, event: "AdvanceEpochs", logs: logs, sub: sub}, nil
}

// WatchAdvanceEpochs is a free log subscription operation binding the contract event 0x0151256d62457b809bbc891b1f81c6dd0b9987552c70ce915b519750cd434dd1.
//
// Solidity: event AdvanceEpochs(uint256 num)
func (_Contract *ContractFilterer) WatchAdvanceEpochs(opts *bind.WatchOpts, sink chan<- *ContractAdvanceEpochs) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AdvanceEpochs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAdvanceEpochs)
				if err := _Contract.contract.UnpackLog(event, "AdvanceEpochs", log); err != nil {
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

// ParseAdvanceEpochs is a log parse operation binding the contract event 0x0151256d62457b809bbc891b1f81c6dd0b9987552c70ce915b519750cd434dd1.
//
// Solidity: event AdvanceEpochs(uint256 num)
func (_Contract *ContractFilterer) ParseAdvanceEpochs(log types.Log) (*ContractAdvanceEpochs, error) {
	event := new(ContractAdvanceEpochs)
	if err := _Contract.contract.UnpackLog(event, "AdvanceEpochs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdateNetworkRulesIterator is returned from FilterUpdateNetworkRules and is used to iterate over the raw logs and unpacked data for UpdateNetworkRules events raised by the Contract contract.
type ContractUpdateNetworkRulesIterator struct {
	Event *ContractUpdateNetworkRules // Event containing the contract specifics and raw log

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
func (it *ContractUpdateNetworkRulesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdateNetworkRules)
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
		it.Event = new(ContractUpdateNetworkRules)
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
func (it *ContractUpdateNetworkRulesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdateNetworkRulesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdateNetworkRules represents a UpdateNetworkRules event raised by the Contract contract.
type ContractUpdateNetworkRules struct {
	Diff []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUpdateNetworkRules is a free log retrieval operation binding the contract event 0x47d10eed096a44e3d0abc586c7e3a5d6cb5358cc90e7d437cd0627f7e765fb99.
//
// Solidity: event UpdateNetworkRules(bytes diff)
func (_Contract *ContractFilterer) FilterUpdateNetworkRules(opts *bind.FilterOpts) (*ContractUpdateNetworkRulesIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateNetworkRules")
	if err != nil {
		return nil, err
	}
	return &ContractUpdateNetworkRulesIterator{contract: _Contract.contract, event: "UpdateNetworkRules", logs: logs, sub: sub}, nil
}

// WatchUpdateNetworkRules is a free log subscription operation binding the contract event 0x47d10eed096a44e3d0abc586c7e3a5d6cb5358cc90e7d437cd0627f7e765fb99.
//
// Solidity: event UpdateNetworkRules(bytes diff)
func (_Contract *ContractFilterer) WatchUpdateNetworkRules(opts *bind.WatchOpts, sink chan<- *ContractUpdateNetworkRules) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateNetworkRules")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdateNetworkRules)
				if err := _Contract.contract.UnpackLog(event, "UpdateNetworkRules", log); err != nil {
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

// ParseUpdateNetworkRules is a log parse operation binding the contract event 0x47d10eed096a44e3d0abc586c7e3a5d6cb5358cc90e7d437cd0627f7e765fb99.
//
// Solidity: event UpdateNetworkRules(bytes diff)
func (_Contract *ContractFilterer) ParseUpdateNetworkRules(log types.Log) (*ContractUpdateNetworkRules, error) {
	event := new(ContractUpdateNetworkRules)
	if err := _Contract.contract.UnpackLog(event, "UpdateNetworkRules", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdateNetworkVersionIterator is returned from FilterUpdateNetworkVersion and is used to iterate over the raw logs and unpacked data for UpdateNetworkVersion events raised by the Contract contract.
type ContractUpdateNetworkVersionIterator struct {
	Event *ContractUpdateNetworkVersion // Event containing the contract specifics and raw log

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
func (it *ContractUpdateNetworkVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdateNetworkVersion)
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
		it.Event = new(ContractUpdateNetworkVersion)
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
func (it *ContractUpdateNetworkVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdateNetworkVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdateNetworkVersion represents a UpdateNetworkVersion event raised by the Contract contract.
type ContractUpdateNetworkVersion struct {
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateNetworkVersion is a free log retrieval operation binding the contract event 0x2ccdfd47cf0c1f1069d949f1789bb79b2f12821f021634fc835af1de66ea2feb.
//
// Solidity: event UpdateNetworkVersion(uint256 version)
func (_Contract *ContractFilterer) FilterUpdateNetworkVersion(opts *bind.FilterOpts) (*ContractUpdateNetworkVersionIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateNetworkVersion")
	if err != nil {
		return nil, err
	}
	return &ContractUpdateNetworkVersionIterator{contract: _Contract.contract, event: "UpdateNetworkVersion", logs: logs, sub: sub}, nil
}

// WatchUpdateNetworkVersion is a free log subscription operation binding the contract event 0x2ccdfd47cf0c1f1069d949f1789bb79b2f12821f021634fc835af1de66ea2feb.
//
// Solidity: event UpdateNetworkVersion(uint256 version)
func (_Contract *ContractFilterer) WatchUpdateNetworkVersion(opts *bind.WatchOpts, sink chan<- *ContractUpdateNetworkVersion) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateNetworkVersion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdateNetworkVersion)
				if err := _Contract.contract.UnpackLog(event, "UpdateNetworkVersion", log); err != nil {
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

// ParseUpdateNetworkVersion is a log parse operation binding the contract event 0x2ccdfd47cf0c1f1069d949f1789bb79b2f12821f021634fc835af1de66ea2feb.
//
// Solidity: event UpdateNetworkVersion(uint256 version)
func (_Contract *ContractFilterer) ParseUpdateNetworkVersion(log types.Log) (*ContractUpdateNetworkVersion, error) {
	event := new(ContractUpdateNetworkVersion)
	if err := _Contract.contract.UnpackLog(event, "UpdateNetworkVersion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdateValidatorPubkeyIterator is returned from FilterUpdateValidatorPubkey and is used to iterate over the raw logs and unpacked data for UpdateValidatorPubkey events raised by the Contract contract.
type ContractUpdateValidatorPubkeyIterator struct {
	Event *ContractUpdateValidatorPubkey // Event containing the contract specifics and raw log

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
func (it *ContractUpdateValidatorPubkeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdateValidatorPubkey)
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
		it.Event = new(ContractUpdateValidatorPubkey)
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
func (it *ContractUpdateValidatorPubkeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdateValidatorPubkeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdateValidatorPubkey represents a UpdateValidatorPubkey event raised by the Contract contract.
type ContractUpdateValidatorPubkey struct {
	ValidatorID *big.Int
	Pubkey      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateValidatorPubkey is a free log retrieval operation binding the contract event 0x0f0ef1ab97439def0a9d2c6d9dc166207f1b13b99e62b442b2993d6153c63a6e.
//
// Solidity: event UpdateValidatorPubkey(uint256 indexed validatorID, bytes pubkey)
func (_Contract *ContractFilterer) FilterUpdateValidatorPubkey(opts *bind.FilterOpts, validatorID []*big.Int) (*ContractUpdateValidatorPubkeyIterator, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateValidatorPubkey", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractUpdateValidatorPubkeyIterator{contract: _Contract.contract, event: "UpdateValidatorPubkey", logs: logs, sub: sub}, nil
}

// WatchUpdateValidatorPubkey is a free log subscription operation binding the contract event 0x0f0ef1ab97439def0a9d2c6d9dc166207f1b13b99e62b442b2993d6153c63a6e.
//
// Solidity: event UpdateValidatorPubkey(uint256 indexed validatorID, bytes pubkey)
func (_Contract *ContractFilterer) WatchUpdateValidatorPubkey(opts *bind.WatchOpts, sink chan<- *ContractUpdateValidatorPubkey, validatorID []*big.Int) (event.Subscription, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateValidatorPubkey", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdateValidatorPubkey)
				if err := _Contract.contract.UnpackLog(event, "UpdateValidatorPubkey", log); err != nil {
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

// ParseUpdateValidatorPubkey is a log parse operation binding the contract event 0x0f0ef1ab97439def0a9d2c6d9dc166207f1b13b99e62b442b2993d6153c63a6e.
//
// Solidity: event UpdateValidatorPubkey(uint256 indexed validatorID, bytes pubkey)
func (_Contract *ContractFilterer) ParseUpdateValidatorPubkey(log types.Log) (*ContractUpdateValidatorPubkey, error) {
	event := new(ContractUpdateValidatorPubkey)
	if err := _Contract.contract.UnpackLog(event, "UpdateValidatorPubkey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdateValidatorWeightIterator is returned from FilterUpdateValidatorWeight and is used to iterate over the raw logs and unpacked data for UpdateValidatorWeight events raised by the Contract contract.
type ContractUpdateValidatorWeightIterator struct {
	Event *ContractUpdateValidatorWeight // Event containing the contract specifics and raw log

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
func (it *ContractUpdateValidatorWeightIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdateValidatorWeight)
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
		it.Event = new(ContractUpdateValidatorWeight)
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
func (it *ContractUpdateValidatorWeightIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdateValidatorWeightIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdateValidatorWeight represents a UpdateValidatorWeight event raised by the Contract contract.
type ContractUpdateValidatorWeight struct {
	ValidatorID *big.Int
	Weight      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateValidatorWeight is a free log retrieval operation binding the contract event 0xb975807576e3b1461be7de07ebf7d20e4790ed802d7a0c4fdd0a1a13df72a935.
//
// Solidity: event UpdateValidatorWeight(uint256 indexed validatorID, uint256 weight)
func (_Contract *ContractFilterer) FilterUpdateValidatorWeight(opts *bind.FilterOpts, validatorID []*big.Int) (*ContractUpdateValidatorWeightIterator, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateValidatorWeight", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractUpdateValidatorWeightIterator{contract: _Contract.contract, event: "UpdateValidatorWeight", logs: logs, sub: sub}, nil
}

// WatchUpdateValidatorWeight is a free log subscription operation binding the contract event 0xb975807576e3b1461be7de07ebf7d20e4790ed802d7a0c4fdd0a1a13df72a935.
//
// Solidity: event UpdateValidatorWeight(uint256 indexed validatorID, uint256 weight)
func (_Contract *ContractFilterer) WatchUpdateValidatorWeight(opts *bind.WatchOpts, sink chan<- *ContractUpdateValidatorWeight, validatorID []*big.Int) (event.Subscription, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateValidatorWeight", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdateValidatorWeight)
				if err := _Contract.contract.UnpackLog(event, "UpdateValidatorWeight", log); err != nil {
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

// ParseUpdateValidatorWeight is a log parse operation binding the contract event 0xb975807576e3b1461be7de07ebf7d20e4790ed802d7a0c4fdd0a1a13df72a935.
//
// Solidity: event UpdateValidatorWeight(uint256 indexed validatorID, uint256 weight)
func (_Contract *ContractFilterer) ParseUpdateValidatorWeight(log types.Log) (*ContractUpdateValidatorWeight, error) {
	event := new(ContractUpdateValidatorWeight)
	if err := _Contract.contract.UnpackLog(event, "UpdateValidatorWeight", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdatedBackendIterator is returned from FilterUpdatedBackend and is used to iterate over the raw logs and unpacked data for UpdatedBackend events raised by the Contract contract.
type ContractUpdatedBackendIterator struct {
	Event *ContractUpdatedBackend // Event containing the contract specifics and raw log

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
func (it *ContractUpdatedBackendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdatedBackend)
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
		it.Event = new(ContractUpdatedBackend)
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
func (it *ContractUpdatedBackendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdatedBackendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdatedBackend represents a UpdatedBackend event raised by the Contract contract.
type ContractUpdatedBackend struct {
	Backend common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdatedBackend is a free log retrieval operation binding the contract event 0x64ee8f7bfc37fc205d7194ee3d64947ab7b57e663cd0d1abd3ef245035830693.
//
// Solidity: event UpdatedBackend(address indexed backend)
func (_Contract *ContractFilterer) FilterUpdatedBackend(opts *bind.FilterOpts, backend []common.Address) (*ContractUpdatedBackendIterator, error) {

	var backendRule []interface{}
	for _, backendItem := range backend {
		backendRule = append(backendRule, backendItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdatedBackend", backendRule)
	if err != nil {
		return nil, err
	}
	return &ContractUpdatedBackendIterator{contract: _Contract.contract, event: "UpdatedBackend", logs: logs, sub: sub}, nil
}

// WatchUpdatedBackend is a free log subscription operation binding the contract event 0x64ee8f7bfc37fc205d7194ee3d64947ab7b57e663cd0d1abd3ef245035830693.
//
// Solidity: event UpdatedBackend(address indexed backend)
func (_Contract *ContractFilterer) WatchUpdatedBackend(opts *bind.WatchOpts, sink chan<- *ContractUpdatedBackend, backend []common.Address) (event.Subscription, error) {

	var backendRule []interface{}
	for _, backendItem := range backend {
		backendRule = append(backendRule, backendItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdatedBackend", backendRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdatedBackend)
				if err := _Contract.contract.UnpackLog(event, "UpdatedBackend", log); err != nil {
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

// ParseUpdatedBackend is a log parse operation binding the contract event 0x64ee8f7bfc37fc205d7194ee3d64947ab7b57e663cd0d1abd3ef245035830693.
//
// Solidity: event UpdatedBackend(address indexed backend)
func (_Contract *ContractFilterer) ParseUpdatedBackend(log types.Log) (*ContractUpdatedBackend, error) {
	event := new(ContractUpdatedBackend)
	if err := _Contract.contract.UnpackLog(event, "UpdatedBackend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var ContractBinRuntime = "0x608060405234801561001057600080fd5b50600436106101165760003560e01c806379bead38116100a2578063d6a0c7af11610071578063d6a0c7af1461073e578063da7fc24f146107a2578063e08d7e66146107e6578063e30443bc1461085f578063ebdf104c146108ad57610116565b806379bead38146104bd5780637f52e13e1461050b578063a4066fbe1461068d578063b9cc6b1c146106c557610116565b8063242a6e3f116100e9578063242a6e3f1461027a578063267ab446146102fd57806339e503ab1461032b578063485cc955146103835780634feb92f3146103e757610116565b806307690b2a1461011b5780630aeeca001461017f57806318f628d4146101ad5780631e702f8314610242575b600080fd5b61017d6004803603604081101561013157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a25565b005b6101ab6004803603602081101561019557600080fd5b8101908080359060200190929190505050610bd9565b005b61024060048036036101208110156101c457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190505050610cd6565b005b6102786004803603604081101561025857600080fd5b810190808035906020019092919080359060200190929190505050610e7c565b005b6102fb6004803603604081101561029057600080fd5b8101908080359060200190929190803590602001906401000000008111156102b757600080fd5b8201836020820111156102c957600080fd5b803590602001918460018302840111640100000000831117156102eb57600080fd5b9091929391929390505050610fb7565b005b6103296004803603602081101561031357600080fd5b81019080803590602001909291905050506110e3565b005b6103816004803603606081101561034157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291905050506111e0565b005b6103e56004803603604081101561039957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611371565b005b6104bb60048036036101008110156103fe57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561044557600080fd5b82018360208201111561045757600080fd5b8035906020019184600183028401116401000000008311171561047957600080fd5b90919293919293908035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190505050611537565b005b610509600480360360408110156104d357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611701565b005b61068b600480360360a081101561052157600080fd5b810190808035906020019064010000000081111561053e57600080fd5b82018360208201111561055057600080fd5b8035906020019184602083028401116401000000008311171561057257600080fd5b90919293919293908035906020019064010000000081111561059357600080fd5b8201836020820111156105a557600080fd5b803590602001918460208302840111640100000000831117156105c757600080fd5b9091929391929390803590602001906401000000008111156105e857600080fd5b8201836020820111156105fa57600080fd5b8035906020019184602083028401116401000000008311171561061c57600080fd5b90919293919293908035906020019064010000000081111561063d57600080fd5b82018360208201111561064f57600080fd5b8035906020019184602083028401116401000000008311171561067157600080fd5b909192939192939080359060200190929190505050611889565b005b6106c3600480360360408110156106a357600080fd5b810190808035906020019092919080359060200190929190505050611a9f565b005b61073c600480360360208110156106db57600080fd5b81019080803590602001906401000000008111156106f857600080fd5b82018360208201111561070a57600080fd5b8035906020019184600183028401116401000000008311171561072c57600080fd5b9091929391929390505050611b9e565b005b6107a06004803603604081101561075457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611cc8565b005b6107e4600480360360208110156107b857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611e7c565b005b61085d600480360360208110156107fc57600080fd5b810190808035906020019064010000000081111561081957600080fd5b82018360208201111561082b57600080fd5b8035906020019184602083028401116401000000008311171561084d57600080fd5b9091929391929390505050611fc6565b005b6108ab6004803603604081101561087557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050612128565b005b610a23600480360360808110156108c357600080fd5b81019080803590602001906401000000008111156108e057600080fd5b8201836020820111156108f257600080fd5b8035906020019184602083028401116401000000008311171561091457600080fd5b90919293919293908035906020019064010000000081111561093557600080fd5b82018360208201111561094757600080fd5b8035906020019184602083028401116401000000008311171561096957600080fd5b90919293919293908035906020019064010000000081111561098a57600080fd5b82018360208201111561099c57600080fd5b803590602001918460208302840111640100000000831117156109be57600080fd5b9091929391929390803590602001906401000000008111156109df57600080fd5b8201836020820111156109f157600080fd5b80359060200191846020830284011164010000000083111715610a1357600080fd5b90919293919293905050506122b0565b005b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ae8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166307690b2a83836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b158015610bbd57600080fd5b505af1158015610bd1573d6000803e3d6000fd5b505050505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c9c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b7f0151256d62457b809bbc891b1f81c6dd0b9987552c70ce915b519750cd434dd1816040518082815260200191505060405180910390a150565b600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d78576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f6e6f742063616c6c61626c65000000000000000000000000000000000000000081525060200191505060405180910390fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318f628d48a8a8a8a8a8a8a8a8a6040518a63ffffffff1660e01b8152600401808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018881526020018781526020018681526020018581526020018481526020018381526020018281526020019950505050505050505050600060405180830381600087803b158015610e5957600080fd5b505af1158015610e6d573d6000803e3d6000fd5b50505050505050505050505050565b600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f1e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f6e6f742063616c6c61626c65000000000000000000000000000000000000000081525060200191505060405180910390fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631e702f8383836040518363ffffffff1660e01b81526004018083815260200182815260200192505050600060405180830381600087803b158015610f9b57600080fd5b505af1158015610faf573d6000803e3d6000fd5b505050505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461107a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b827f0f0ef1ab97439def0a9d2c6d9dc166207f1b13b99e62b442b2993d6153c63a6e838360405180806020018281038252848482818152602001925080828437600081840152601f19601f820116905080830192505050935050505060405180910390a2505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146111a6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b7f2ccdfd47cf0c1f1069d949f1789bb79b2f12821f021634fc835af1de66ea2feb816040518082815260200191505060405180910390a150565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146112a3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166339e503ab8484846040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019350505050600060405180830381600087803b15801561135457600080fd5b505af1158015611368573d6000803e3d6000fd5b50505050505050565b600060019054906101000a900460ff1680611390575061138f6124c9565b5b806113a757506000809054906101000a900460ff16155b6113fc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e8152602001806124e1602e913960400191505060405180910390fd5b60008060019054906101000a900460ff16159050801561144c576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b82603460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff167f64ee8f7bfc37fc205d7194ee3d64947ab7b57e663cd0d1abd3ef24503583069360405160405180910390a281603560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080156115325760008060016101000a81548160ff0219169083151502179055505b505050565b600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146115d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f6e6f742063616c6c61626c65000000000000000000000000000000000000000081525060200191505060405180910390fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634feb92f38a8a8a8a8a8a8a8a8a6040518a63ffffffff1660e01b8152600401808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001898152602001806020018781526020018681526020018581526020018481526020018381526020018281038252898982818152602001925080828437600081840152601f19601f8201169050808301925050509a5050505050505050505050600060405180830381600087803b1580156116de57600080fd5b505af11580156116f2573d6000803e3d6000fd5b50505050505050505050505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146117c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166379bead3883836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561186d57600080fd5b505af1158015611881573d6000803e3d6000fd5b505050505050565b600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461192b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f6e6f742063616c6c61626c65000000000000000000000000000000000000000081525060200191505060405180910390fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663592fe0c08a8a8a8a8a8a8a8a8a6040518a63ffffffff1660e01b8152600401808060200180602001806020018060200186815260200185810385528e8e82818152602001925060200280828437600081840152601f19601f82011690508083019250505085810384528c8c82818152602001925060200280828437600081840152601f19601f82011690508083019250505085810383528a8a82818152602001925060200280828437600081840152601f19601f8201169050808301925050508581038252888882818152602001925060200280828437600081840152601f19601f8201169050808301925050509d5050505050505050505050505050600060405180830381600087803b158015611a7c57600080fd5b505af1158015611a90573d6000803e3d6000fd5b50505050505050505050505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611b62576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b817fb975807576e3b1461be7de07ebf7d20e4790ed802d7a0c4fdd0a1a13df72a935826040518082815260200191505060405180910390a25050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611c61576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b7f47d10eed096a44e3d0abc586c7e3a5d6cb5358cc90e7d437cd0627f7e765fb99828260405180806020018281038252848482818152602001925080828437600081840152601f19601f820116905080830192505050935050505060405180910390a15050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611d8b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d6a0c7af83836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b158015611e6057600080fd5b505af1158015611e74573d6000803e3d6000fd5b505050505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611f3f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff167f64ee8f7bfc37fc205d7194ee3d64947ab7b57e663cd0d1abd3ef24503583069360405160405180910390a280603460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612068576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f6e6f742063616c6c61626c65000000000000000000000000000000000000000081525060200191505060405180910390fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e08d7e6683836040518363ffffffff1660e01b815260040180806020018281038252848482818152602001925060200280828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b15801561210c57600080fd5b505af1158015612120573d6000803e3d6000fd5b505050505050565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146121eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616c6c6572206973206e6f7420746865206261636b656e640000000000000081525060200191505060405180910390fd5b603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e30443bc83836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561229457600080fd5b505af11580156122a8573d6000803e3d6000fd5b505050505050565b600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612352576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f6e6f742063616c6c61626c65000000000000000000000000000000000000000081525060200191505060405180910390fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663592fe0c0898989898989898963322adc3a6040518a63ffffffff1660e01b8152600401808060200180602001806020018060200186815260200185810385528e8e82818152602001925060200280828437600081840152601f19601f82011690508083019250505085810384528c8c82818152602001925060200280828437600081840152601f19601f82011690508083019250505085810383528a8a82818152602001925060200280828437600081840152601f19601f8201169050808301925050508581038252888882818152602001925060200280828437600081840152601f19601f8201169050808301925050509d5050505050505050505050505050600060405180830381600087803b1580156124a757600080fd5b505af11580156124bb573d6000803e3d6000fd5b505050505050505050505050565b6000803090506000813b905060008114925050509056fe436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564a265627a7a723158200e3b4f3f8ff8f9f35bf6a4819537ff47f280a0f7820eaeed0a40415efda42d5964736f6c63430005100032"
