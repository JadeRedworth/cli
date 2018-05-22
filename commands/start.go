package commands

import (
	"github.com/urfave/cli"
)

func StartCommand() cli.Command {
	return cli.Command{
		Name:        "start",
		Usage:       "start",
		Aliases:     []string{"st"},
		Category:    "DEVELOPMENT COMMANDS",
		Subcommands: GetCommands(StartCmds),
	}
}
