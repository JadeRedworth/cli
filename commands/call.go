package commands

import (
	"net/url"
	"os"
	"path"

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

	u := url.URL{
		Scheme: "http",
		Host:   client.Host(),
	}
	u.Path = path.Join(u.Path, "r", appName, route)
	content := run.Stdin()

	return client.CallFN(u.String(), content, os.Stdout, c.String("method"), c.StringSlice("e"), c.String("content-type"), c.Bool("display-call-id"))
}
