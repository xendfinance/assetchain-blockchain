package sfc

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/Fantom-foundation/go-opera/gossip/contract/sfc100"
)

// GetContractBin is SFC contract genesis implementation bin code
// Has to be compiled with flag bin-runtime
// Built from xend-contracts 6858d2716f255d815d250f90d7c8ac1ed86b2353, solc 0.5.17+commit.d19bba13.Emscripten.clang, optimize-runs 200
func GetContractBin() []byte {
	return hexutil.MustDecode(sfc100.ContractBinRuntime)
}

// ContractAddress is the SFC contract address
var ContractAddress = common.HexToAddress("0xfc00face00000000000000000000000000000000")
