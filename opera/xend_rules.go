package opera

import (
	"time"

	"github.com/Fantom-foundation/go-opera/inter"
)

// AssetChainMainnetRules returns mainnet rules
func AssetChainMainnetRules() Rules {
	return Rules{
		Name:      AssetChainName,
		NetworkID: AssetChainProdNetworkID,
		Dag:       DefaultDagRules(),
		Epochs:    XendNetEpochsRules(),
		Economy:   DefaultEconomyRules(),
		Blocks: BlocksRules{
			MaxBlockGas:             20500000,
			MaxEmptyBlockSkipPeriod: inter.Timestamp(60 * time.Second),
		},
		Upgrades: Upgrades{
			Berlin: true,
			London: true,
			Llr:    true,
		},
	}
}

// AssetChainTestnetRules returns testnet rules
// equal to mainnet
func AssetChainTestnetRules() Rules {
	return Rules{
		Name:      AssetChainName,
		NetworkID: AssetChainTestNetworkID,
		Dag:       DefaultDagRules(),
		Epochs:    XendNetEpochsRules(),
		Economy:   DefaultEconomyRules(),
		Blocks: BlocksRules{
			MaxBlockGas:             20500000,
			MaxEmptyBlockSkipPeriod: inter.Timestamp(60 * time.Second),
		},
		Upgrades: Upgrades{
			Berlin: true,
			London: true,
			Llr:    true,
		},
	}
}

// XendDevNetRules returns devnet rules
func XendDevNetRules() Rules {
	return Rules{
		Name:      AssetChainName,
		NetworkID: AssetChainDevNetworkID,
		Dag:       DefaultDagRules(),
		Epochs:    XendDevEpochsRules(),
		Economy:   DefaultEconomyRules(),
		Blocks: BlocksRules{
			MaxBlockGas:             20500000,
			MaxEmptyBlockSkipPeriod: inter.Timestamp(60 * time.Second),
		},
		Upgrades: Upgrades{
			Berlin: true,
			London: true,
			Llr:    true,
		},
	}
}

func XendNetEpochsRules() EpochsRules {
	cfg := DefaultEpochsRules()
	cfg.MaxEpochDuration = inter.Timestamp(4 * time.Hour)
	return cfg
}

func XendDevEpochsRules() EpochsRules {
	cfg := DefaultEpochsRules()
	cfg.MaxEpochDuration = inter.Timestamp(4 * time.Hour)
	return cfg
}
