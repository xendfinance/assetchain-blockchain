package driverauth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/Fantom-foundation/go-opera/gossip/contract/driverauth100"
)

// GetContractBin is NodeDriverAuth contract genesis implementation bin code
// Has to be compiled with flag bin-runtime
// Built from opera-sfc 424031c81a77196f4e9d60c7d876032dd47208ce, solc 0.5.17+commit.d19bba13.Emscripten.clang, optimize-runs 10000
func GetContractBin() []byte {
	return hexutil.MustDecode(driverauth100.ContractBinRuntime)
}

// ContractAddress is the NodeDriverAuth contract address
var ContractAddress = common.HexToAddress("0xd100ae0000000000000000000000000000000000")
