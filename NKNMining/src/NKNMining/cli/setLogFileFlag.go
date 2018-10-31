package commands

import "github.com/urfave/cli"

type SetLogFlag struct{}

func (*SetLogFlag) NewFlag() cli.Flag {
	return &(cli.StringFlag{
		Name:   cli_FLAG_LOG,
		Usage:  "set log file name, default is NKNMining.log",
		Hidden: false,
		Value: 	"NKNMining.log",
	})
}


