package launcher

import (
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/ethereum/go-ethereum/params"

	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/genesis"
	"github.com/Fantom-foundation/go-opera/opera/genesisstore"
)

var (
	Bootnodes = map[string][]string{
		"Vitra Mainnet": {
			"enode://60792ba55d984a46f0a2ba469089de6ca16dfad812a621944f70248e02208c23ddf8f823fb05c12b5d36964f38d73055a89e32b5d19c90495a10470b60e37512@nodes-1.vitrachain-rpc.com:3000",
			"enode://9009824396c6c2fbe976481ad07f92d5cb696e59a9a35fb20ef318452ba77e58b3f5d8903062ecc49a4cc6630d6250967e3fd44dec8f1ac2903440018891c164@nodes-2.vitrachain-rpc.com:3000",
			"enode://56f1707fd0460654840d2ad726e9bafa653cc0b39289a2d838530a44b355b9927ebe7db99222d3a9aed977ad0c193e56e0a774c80f19bae6344dce758f81a6de@nodes-3.vitrachain-rpc.com:3000",
			"enode://a9ff87401f8c19715ceaa618ebe0d68e76b82ae57a19d8c792cb25637b7edf8a4e9d9cd933cde47102acc2cb5a6d79dd1d01c162e85a9d4d9dac89f9b4e1bbd1@nodes-4.vitrachain-rpc.com:3000",
			"enode://da2917206f04de4c1a18d66ebbdc7d1e6c74e892a39762280634dc5c45602c5ce1a325016ec2225fe2eb1957f81602a0b4fa1ac9ccb49d9c600399b4ff3b78a1@nodes-5.vitrachain-rpc.com:3000",
		},

		"Vitra Testnet": {
			"enode://05b63dd3b86bf547cad299180df834dca5bc995dda1eb5803ce980d02447d73d3c0aa28bebc346f11a43a8e40c714a2ffed24bffd8bdfbbaef3c1d2e4428a66d@nlb.vitrafoundation-rpc.com:3000",
		},
	}

	mainnetHeader = genesis.Header{
		GenesisID:   hash.HexToHash("0x6f9e74ee2b3a052d173bb09ef06f5a213569367b454af1a05cdf33a12a8c9e8c"),
		NetworkID:   opera.VitraMainNetworkID,
		NetworkName: "Vitra Mainnet",
	}

	testnetHeader = genesis.Header{
		GenesisID:   hash.HexToHash("0x2b0e07eb61740e8409859c4dc2f76a7ce11df34bab4f686bc24aa2b2edf829eb"),
		NetworkID:   opera.VitraTestNetworkID,
		NetworkName: "Vitra Testnet",
	}

	AllowedOperaGenesis = []GenesisTemplate{

		{
			Name:   "Vitra Mainnet",
			Header: mainnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0): hash.HexToHash("0xac853075b052427b5247aad8f9c164c48e29acc19b0c27298bc19b1564bc10d1"),
				genesisstore.BlocksSection(0): hash.HexToHash("0x31537b677879927f57b1c7617331c64ef72fe86d7c8f5f6cdb3c747b7e50261b"),
				genesisstore.EvmSection(0):    hash.HexToHash("0x0ac3143552e971f7db1a3e4c9ddcaf6ce240bb56bc878e3b6b48a317dc837bec"),
			},
		},

		{
			Name:   "Vitra Testnet",
			Header: testnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0): hash.HexToHash("0x071260c77ff4775a72611ed79ee0055a1c5861583e86200d6750675b4740e95f"),
				genesisstore.BlocksSection(0): hash.HexToHash("0x6feebb87208c21b7c01f8c0b00b8724045bf1f4ed3c6149843aeea1e6c3a3c3c"),
				genesisstore.EvmSection(0):    hash.HexToHash("0x9a9f886edefcf0a6916c986ce1d4c9ce4318c3d9ecd92ad8075d4bfafab18148"),
			},
		},

		// {
		// 	Name:   "Mainnet-5577 with pruned MPT",
		// 	Header: mainnetHeader,
		// 	Hashes: genesis.Hashes{
		// 		genesisstore.EpochsSection(0): hash.HexToHash("0x945d8084b4e6e1e78cfe9472fefca3f6ecc7041765dfed24f64e9946252f569a"),
		// 		genesisstore.BlocksSection(0): hash.HexToHash("0xe3ec041f3cca79928aa4abef588b48e96ff3cfa3908b2268af3ac5496c722fec"),
		// 		genesisstore.EvmSection(0):    hash.HexToHash("0x12dd52ac21fee5d76b47a64386e73187d5260e448e8044f38c6c73eaa627e4b5"),
		// 	},
		// },
		// {
		// 	Name:   "Testnet-2458 with pruned MPT",
		// 	Header: testnetHeader,
		// 	Hashes: genesis.Hashes{
		// 		genesisstore.EpochsSection(0): hash.HexToHash("0x4a5caf86d7f5a31dad91f2cbd44db052c602f515e5319f828adb585a7a6723d6"),
		// 		genesisstore.BlocksSection(0): hash.HexToHash("0x07eadb81c1e2a1b5c444c8c2430c6873380f447de64790b25abe9e7fa6874f65"),
		// 		genesisstore.EvmSection(0):    hash.HexToHash("0xa96e006ae17d15e1244c3e7ff4d556e5a3849e70df7a81704787f3273f37c9b1"),
		// 	},
		// },

		// {
		// 	Name:   "Testnet-6226 with full MPT",
		// 	Header: testnetHeader,
		// 	Hashes: genesis.Hashes{
		// 		genesisstore.EpochsSection(0): hash.HexToHash("0x61040a80f16755b86d67f13880f55c1238d307e2e1c6fc87951eb3bdee0bdff2"),
		// 		genesisstore.BlocksSection(0): hash.HexToHash("0x12010621d8cf4dcd4ea357e98eac61edf9517a6df752cb2d929fca69e56bd8d1"),
		// 		genesisstore.EvmSection(0):    hash.HexToHash("0x9227c80bf56e4af08dc32cb6043cc43672f2be8177d550ab34a7a9f57f4f104b"),
		// 	},
		// },
		// {
		// 	Name:   "Testnet-16200 without MPT",
		// 	Header: testnetHeader,
		// 	Hashes: genesis.Hashes{
		// 		genesisstore.EpochsSection(0): hash.HexToHash("0x61040a80f16755b86d67f13880f55c1238d307e2e1c6fc87951eb3bdee0bdff2"),
		// 		genesisstore.BlocksSection(0): hash.HexToHash("0x12010621d8cf4dcd4ea357e98eac61edf9517a6df752cb2d929fca69e56bd8d1"),
		// 		genesisstore.EpochsSection(1): hash.HexToHash("0x7d651ed0e0f3e92ffd89cb52112598db54afd8bf3050bc083ff0bfe1b98948fd"),
		// 		genesisstore.BlocksSection(1): hash.HexToHash("0xd72e9bf39c645df8d978955fab8997a7e9cd7cb5812c007e2bb4b51a8c570a90"),
		// 	},
		// },
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
