package opera

import (
	"time"

	"github.com/Fantom-foundation/go-opera/inter"
)

// VitraMainNetRules returns mainnet rules
func VitraMainNetRules() Rules {
	return Rules{
		Name:      "Vitra Mainnet",
		NetworkID: VitraMainNetworkID,
		Dag:       DefaultDagRules(),
		Epochs:    VitraNetEpochsRules(),
		Economy:   DefaultEconomyRules(),
		Blocks: BlocksRules{
			MaxBlockGas:             20500000,
			MaxEmptyBlockSkipPeriod: inter.Timestamp(10 * time.Second),
		},
		Upgrades: Upgrades{
			Berlin: true,
			London: true,
			Llr:    true,
		},
	}
}

// VitraTestNetRules returns testnet rules
// equal to mainnet
func VitraTestNetRules() Rules {
	return Rules{
		Name:      "Vitra Testnet",
		NetworkID: VitraTestNetworkID,
		Dag:       DefaultDagRules(),
		Epochs:    VitraNetEpochsRules(),
		Economy:   DefaultEconomyRules(),
		Blocks: BlocksRules{
			MaxBlockGas:             20500000,
			MaxEmptyBlockSkipPeriod: inter.Timestamp(10 * time.Second),
		},
		Upgrades: Upgrades{
			Berlin: true,
			London: true,
			Llr:    true,
		},
	}
}

// VitraDevNetRules returns devnet rules
func VitraDevNetRules() Rules {
	return Rules{
		Name:      "Vitra Devnet",
		NetworkID: VitraDevNetworkID,
		Dag:       DefaultDagRules(),
		Epochs:    VitraDevEpochsRules(),
		Economy:   DefaultEconomyRules(),
		Blocks: BlocksRules{
			MaxBlockGas:             20500000,
			MaxEmptyBlockSkipPeriod: inter.Timestamp(3 * time.Second),
		},
		Upgrades: Upgrades{
			Berlin: true,
			London: true,
			Llr:    true,
		},
	}
}

func VitraNetEpochsRules() EpochsRules {
	cfg := DefaultEpochsRules()
	cfg.MaxEpochDuration = inter.Timestamp(4 * time.Hour)
	return cfg
}

func VitraDevEpochsRules() EpochsRules {
	cfg := DefaultEpochsRules()
	cfg.MaxEpochDuration = inter.Timestamp(20 * time.Minute)
	return cfg
}
