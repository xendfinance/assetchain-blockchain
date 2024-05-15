package launcher

import (
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/ethereum/go-ethereum/params"

	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/genesis"
	"github.com/Fantom-foundation/go-opera/opera/genesisstore"
)

var (
	Bootnodes = map[uint64][]string{
		opera.AssetChainMainnetRules().NetworkID: {
			"enode://d4aefe078fb585b8b5bf0611e9c408d71a321eab00d74f4edfba1b5db09b6ae24721ddf639cb5be11b9b2a8e6d8e1e4987ac013f500c2902c3916acde8030a7c@nodes-3.rpcmainnet.xendrwachain.com:3000",
			"enode://4f2d722dedbf36999b0aa7698fd7a345a0a31ecb408c13c574a963f17fb6bcb74dbeeec57d70a350fd088cba6d22a51212c9bccb84c5e8631710483fcfac5b93@nodes-2.rpcmainnet.xendrwachain.com:3000",
			"enode://fc22780489d26d92e2e1ed0be5bbb68d947ee7742f800c663c8bd91eb64e687a6ed1f962fb9db0e6af38075d16f0716245b074efbab6d0a7418796cc1904b844@nodes-5.rpcmainnet.xendrwachain.com:3000",
			"enode://3ef14e9c1547204856d3ea4380edb4c6c3a258c9497a7c3b02850a254969abf47d25c48c4ec444099fe560355a0fd36e911e3b95aa742fe8267e099ef2b06858@nodes-4.rpcmainnet.xendrwachain.com:3000",
			"enode://ca6b5af34617013bf1272e052bbf570fda2827affce1cc50d3ba4b55df32163ace111f0859af55c46b6ecf8b7ab3b94f47ad2e423c5439edf09e1bdabf56bf05@nodes-1.rpcmainnet.xendrwachain.com:3000",
		},

		opera.AssetChainTestnetRules().NetworkID: {
			"enode://d4aefe078fb585b8b5bf0611e9c408d71a321eab00d74f4edfba1b5db09b6ae24721ddf639cb5be11b9b2a8e6d8e1e4987ac013f500c2902c3916acde8030a7c@nodes-3.rpcmainnet.xendrwachain.com:3000",
			"enode://4f2d722dedbf36999b0aa7698fd7a345a0a31ecb408c13c574a963f17fb6bcb74dbeeec57d70a350fd088cba6d22a51212c9bccb84c5e8631710483fcfac5b93@nodes-2.rpcmainnet.xendrwachain.com:3000",
			"enode://fc22780489d26d92e2e1ed0be5bbb68d947ee7742f800c663c8bd91eb64e687a6ed1f962fb9db0e6af38075d16f0716245b074efbab6d0a7418796cc1904b844@nodes-5.rpcmainnet.xendrwachain.com:3000",
			"enode://3ef14e9c1547204856d3ea4380edb4c6c3a258c9497a7c3b02850a254969abf47d25c48c4ec444099fe560355a0fd36e911e3b95aa742fe8267e099ef2b06858@nodes-4.rpcmainnet.xendrwachain.com:3000",
			"enode://ca6b5af34617013bf1272e052bbf570fda2827affce1cc50d3ba4b55df32163ace111f0859af55c46b6ecf8b7ab3b94f47ad2e423c5439edf09e1bdabf56bf05@nodes-1.rpcmainnet.xendrwachain.com:3000",
		},
	}

	mainnetHeader = genesis.Header{
		NetworkName: opera.AssetChainName,
		NetworkID:   opera.AssetChainProdNetworkID,
		GenesisID:   hash.HexToHash("0x6f9e74ee2b3a052d173bb09ef06f5a213569367b454af1a05cdf33a12a8c9e8c"),
	}

	testnetHeader = genesis.Header{
		NetworkName: opera.AssetChainName,
		NetworkID:   opera.AssetChainTestNetworkID,
		GenesisID:   hash.HexToHash("0x2b0e07eb61740e8409859c4dc2f76a7ce11df34bab4f686bc24aa2b2edf829eb"),
	}

	AllowedOperaGenesis = []GenesisTemplate{
		{
			Name:   opera.AssetChainName,
			Header: mainnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0): hash.HexToHash("0xac853075b052427b5247aad8f9c164c48e29acc19b0c27298bc19b1564bc10d1"),
				genesisstore.BlocksSection(0): hash.HexToHash("0x31537b677879927f57b1c7617331c64ef72fe86d7c8f5f6cdb3c747b7e50261b"),
				genesisstore.EvmSection(0):    hash.HexToHash("0x0ac3143552e971f7db1a3e4c9ddcaf6ce240bb56bc878e3b6b48a317dc837bec"),
			},
		},

		{
			Name:   opera.AssetChainName,
			Header: testnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0): hash.HexToHash("0x071260c77ff4775a72611ed79ee0055a1c5861583e86200d6750675b4740e95f"),
				genesisstore.BlocksSection(0): hash.HexToHash("0x6feebb87208c21b7c01f8c0b00b8724045bf1f4ed3c6149843aeea1e6c3a3c3c"),
				genesisstore.EvmSection(0):    hash.HexToHash("0x9a9f886edefcf0a6916c986ce1d4c9ce4318c3d9ecd92ad8075d4bfafab18148"),
			},
		},
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
