package driverauth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/Fantom-foundation/go-opera/gossip/contract/driverauth100"
)

// GetContractBin is NodeDriverAuth contract genesis implementation bin code
// Has to be compiled with flag bin-runtime
// Built from xend-contracts 6858d2716f255d815d250f90d7c8ac1ed86b2353, solc 0.5.17+commit.d19bba13.Emscripten.clang, optimize-runs 10000
func GetContractBin() []byte {
	return hexutil.MustDecode(driverauth100.ContractBinRuntime)
}

// ContractAddress is the NodeDriverAuth contract address
var ContractAddress = common.HexToAddress("0xd100ae0000000000000000000000000000000000")
