package commands

import (
	"github.com/urfave/cli"
)

func BuildCommand() cli.Command {
	return cli.Command{
		Name:        "build",
		Aliases:     []string{"b"},
		Usage:       "build command",
		Category:    "DEVELOPMENT COMMANDS",
		Subcommands: GetCommands(BuildCmds),
	}
}
