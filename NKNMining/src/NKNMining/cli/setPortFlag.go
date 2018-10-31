package commands

import "github.com/urfave/cli"

type SetPortFlag struct{}

func (*SetPortFlag) NewFlag() cli.Flag {
	return &(cli.UintFlag{
		Name:   cli_FLAG_PORT,
		Usage:  "set node shell server port, default is 8181",
		Value: 8181,
		Hidden: false,
	})
}
