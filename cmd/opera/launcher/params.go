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
			"enode://94369667584ee7fe73ba644eda2569d21a8533a1b9d055cc836207100cd8c6e9755eedde7234b4c966bbd21535ef20bd9bcd5e3f531af8d0481a5f48f5ec6ffc@rpcmainnet.xendrwachain.com:3000",
			"enode://d53073829b8f21fdda806f10af2a3ee4970a984720321d24f19dd7d72c5fe95fc15c412c34d642b0744f89016c1b0db19838e14253da0699222a2ec9e1a2c58c@nodes-1.rpcmainnet.xendrwachain.com:3000",
			"enode://1e74983e622982a20d2811f0c88d1b56ed1a55cac7da41e980347afca401b21135e9dc23996075ffe100f321c258cbc6737820b0cf65853c1a41240d10a0b06a@nodes-2.rpcmainnet.xendrwachain.com:3000",
			"enode://055e58a5df6777b4f54e4a83570e6eedf6334fe77d3aa7feb4d72650f377c65c46b4fb05be485e52703d32c7d586ba9dd5c4f5795404e2ab960a2ab926d9dafc@nodes-3.rpcmainnet.xendrwachain.com:3000",
			"enode://8572091416d2f9edc13e161b70f6c31ce2d39744e6afde710868f3bce71c2c80514098f4ea6df17683995d6fd9373e382a19cbc7471cc5a7b96982475eb1cb60@nodes-4.rpcmainnet.xendrwachain.com:3000",
			"enode://81a0595ca25a2aab64aed73be0a20a705f45c90955591fcdf6ad4a904599040c4ecc5d153543b99fcad2cbb03f566f1a771b134cbeff58cef32786ce99993028@nodes-5.rpcmainnet.xendrwachain.com:3000",
		},

		opera.AssetChainTestnetRules().NetworkID: {
			"enode://aaec8a1aa57ac5518a34a95366bcd047ba7f8d350610c2c27fa355791737622ed5e027542f795e3819ceb05bad93c0dc8ccb4cbfd7d5adc34d1c4b3d1a8e65aa@rpctestnet.xendrwachain.com:3000",
			"enode://54ca286342bcd0518372f87161ae1153e8f4fc0c8710709451f9f637f1496291010728ed6a30d8494e40a75ca855e3c2d27ea6a711fb4b9015ea9e2beefa7053@nlb.rpctestnet.xendrwachain.com:3000",
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
