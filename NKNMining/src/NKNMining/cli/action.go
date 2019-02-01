package commands

import (
	"NKNMining/config"
	"github.com/urfave/cli"
	"strconv"
)

func SetAction(app *cli.App, isRemote *bool){
	app.Action = func(c *cli.Context) {
		port := c.Uint(cli_FLAG_PORT)
		logfile := c.String(cli_FLAG_LOG)
		*isRemote = c.Bool(cli_FLAG_REMOTE)


		if 0 != port && 65535 <= port {
			cli.ShowCommandHelp(c, "")
			return
		}

		config.ShellConfig.ServerPort = strconv.FormatUint(uint64(port), 10)
		config.ShellConfig.LogFile = logfile

		return
	}
}
