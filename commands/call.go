package commands

import (
	"github.com/fnproject/cli/objects/route"
	"github.com/urfave/cli"
)

func CallCommand() cli.Command {
	return cli.Command{
		Name:     "call",
		Aliases:  []string{"cl"},
		Usage:    "call command",
		Category: "DEVELOPMENT COMMANDS",
		Action:   route.Call,
	}
}
