package commands

import (
	"github.com/urfave/cli"
)

func CallCommand() cli.Command {
	return cli.Command{
		Name:        "call",
		Aliases:     []string{"cl"},
		Usage:       "call command",
		Category:    "DEVELOPMENT COMMANDS",
		Hidden:      false,
		ArgsUsage:   "<command>",
		Subcommands: GetCommands(CallCmds),
	}
}
