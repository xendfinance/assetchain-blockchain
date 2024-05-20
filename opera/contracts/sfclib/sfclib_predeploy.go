package sfclib

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/Fantom-foundation/go-opera/gossip/contract/sfclib100"
)

// GetContractBin is SFCLib contract genesis implementation bin code
// Has to be compiled with flag bin-runtime
// Built from xend-contracts 6858d2716f255d815d250f90d7c8ac1ed86b2353, solc 0.5.17+commit.d19bba13.Emscripten.clang, optimize-runs 200
func GetContractBin() []byte {
	return hexutil.MustDecode(sfclib100.ContractBinRuntime)
}

// ContractAddress is the SFCLib contract address
var ContractAddress = common.HexToAddress("0xfc01face00000000000000000000000000000000")
