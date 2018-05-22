package server

import (
	"github.com/urfave/cli"
)

type buildServerCmd struct {
	verbose bool
	noCache bool
}

func (b *buildServerCmd) flags() []cli.Flag {
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
		cli.StringFlag{
			Name:  "tag,t",
			Usage: "image name and optional tag",
		},
	}
}

func Update() cli.Command {
	return cli.Command{
		Name:    "server",
		Usage:   "pulls latest functions server",
		Aliases: []string{"sv"},
		Action:  update,
	}
}

func Start() cli.Command {
	return cli.Command{
		Name:    "server",
		Usage:   "starts a local server",
		Aliases: []string{"sv"},
		Action:  start,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "log-level",
				Usage: "--log-level debug to enable debugging",
			},
			cli.BoolFlag{
				Name:  "detach, d",
				Usage: "Run container in background.",
			},
			cli.StringFlag{
				Name:  "env-file",
				Usage: "Path to Fn server configuration file.",
			},
			cli.IntFlag{
				Name:  "port, p",
				Value: 8080,
				Usage: "Specify port number to bind to on the host.",
			},
		},
	}
}

func Build() cli.Command {
	cmd := buildServerCmd{}
	flags := append([]cli.Flag{}, cmd.flags()...)
	return cli.Command{
		Name:   "server",
		Usage:  "build custom fn server",
		Flags:  flags,
		Action: cmd.Build,
	}
}
