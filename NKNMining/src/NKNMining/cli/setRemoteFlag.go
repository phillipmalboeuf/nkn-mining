package commands

import "github.com/urfave/cli"

type SetRemoteFlag struct{}

func (*SetRemoteFlag) NewFlag() cli.Flag {
	return &(cli.BoolFlag{
		Name:   cli_FLAG_REMOTE,
		Usage:  "run NKNMining in remote mode",
		Hidden: false,
	})
}
