package brprocessor

import (
	"time"

	"github.com/Fantom-foundation/lachesis-base/inter/dag"
	"github.com/Fantom-foundation/lachesis-base/utils/cachescale"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type Config struct {
	BufferLimit dag.Metric

	SemaphoreTimeout time.Duration

	MaxTasks int
}

func DefaultConfig(scale cachescale.Func) Config {
	return Config{
		BufferLimit: dag.Metric{
			Num:  10000,
			Size: scale.U64(15 * opt.MiB),
		},
		SemaphoreTimeout: 10 * time.Second,
		MaxTasks:         512,
	}
}
