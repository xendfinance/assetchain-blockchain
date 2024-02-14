package launcher

import (
	"github.com/ethereum/go-ethereum/cmd/utils"
	"gopkg.in/urfave/cli.v1"
)

var InitNetCommand = cli.Command{
	Name:     "network",
	Usage:    "network command [command options] [arguments...]",
	Category: "network COMMANDS",

	Subcommands: []cli.Command{
		{
			Name:      "new",
			Usage:     "Create a new network",
			Action:    utils.MigrateFlags(newVitraNetwork),
			ArgsUsage: "network new <val_num> [flags]",
			Flags: []cli.Flag{
				DataDirFlag,
				utils.KeyStoreDirFlag,
				utils.PasswordFileFlag,
				ValidatorsFileFlag,
				NetworkTypeFlag,
			},
			Description: `
opera network new
`,
		},
	},
}

func newVitraNetwork(ctx *cli.Context) error {
	return nil
}
