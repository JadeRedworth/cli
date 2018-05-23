package commands

import (
	"os"

	"github.com/fnproject/cli/client"
	"github.com/fnproject/cli/objects/route"
	"github.com/fnproject/cli/run"
	"github.com/urfave/cli"
)

func CallCommand() cli.Command {
	return cli.Command{
		Name:      "call",
		Usage:     "call a remote function",
		Aliases:   []string{"cl"},
		ArgsUsage: "<app> </path>",
		Flags:     route.CallFnFlags,
		Category:  "DEVELOPMENT COMMANDS",
		Action:    Call,
	}
}

func Call(c *cli.Context) error {
	appName := c.Args().Get(0)
	route := route.CleanRoutePath(c.Args().Get(1))
	content := run.Stdin()

	return client.CallFN(appName, route, content, os.Stdout, c.String("method"), c.StringSlice("e"), c.String("content-type"), c.Bool("display-call-id"))
}
