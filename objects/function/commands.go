package function

import (
	"github.com/urfave/cli"
)

func Build() cli.Command {
	cmd := buildcmd{}
	flags := append([]cli.Flag{}, cmd.flags()...)
	return cli.Command{
		Name:    "function",
		Usage:   "build function version",
		Aliases: []string{"func"},
		Flags:   flags,
		Action:  cmd.Build,
	}
}

type buildcmd struct {
	verbose bool
	noCache bool
}

func (b *buildcmd) flags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:        "v",
			Usage:       "verbose mode",
			Destination: &b.verbose,
		},
		cli.BoolFlag{
			Name:        "no-cache",
			Usage:       "Don't use docker cache",
			Destination: &b.noCache,
		},
		cli.StringSliceFlag{
			Name:  "build-arg",
			Usage: "set build-time variables",
		},
	}
}
