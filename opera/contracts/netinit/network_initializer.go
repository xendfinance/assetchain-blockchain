package netinit

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/Fantom-foundation/go-opera/gossip/contract/netinit100"
)

// GetContractBin is NetworkInitializer contract genesis implementation bin code
// Has to be compiled with flag bin-runtime
// Built from opera-sfc 424031c81a77196f4e9d60c7d876032dd47208ce, solc 0.5.17+commit.d19bba13.Emscripten.clang, optimize-runs 10000
func GetContractBin() []byte {
	return hexutil.MustDecode(netinit100.ContractBinRuntime)
}

// ContractAddress is the NetworkInitializer contract address
var ContractAddress = common.HexToAddress("0xd1005eed00000000000000000000000000000000")
