package opera

import (
	"time"

	"github.com/Fantom-foundation/go-opera/inter"
)

// XendMainNetRules returns mainnet rules
func XendMainNetRules() Rules {
	return Rules{
		Name:      "Xend Mainnet",
		NetworkID: XendMainNetworkID,
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

// XendTestNetRules returns testnet rules
// equal to mainnet
func XendTestNetRules() Rules {
	return Rules{
		Name:      "Xend Testnet",
		NetworkID: XendTestNetworkID,
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
		Name:      "Xend Devnet",
		NetworkID: XendDevNetworkID,
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
