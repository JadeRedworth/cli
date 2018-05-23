package app

import (
	"github.com/fnproject/cli/client"
	"github.com/urfave/cli"
)

type app client.FnClient

func Create() cli.Command {
	return cli.Command{
		Name:      "apps",
		Usage:     "Create a new application",
		Aliases:   []string{"a"},
		ArgsUsage: "<app>",
		Action:    create,
		Flags: []cli.Flag{
			cli.StringSliceFlag{
				Name:  "config",
				Usage: "application configuration",
			},
			cli.StringSliceFlag{
				Name:  "annotation",
				Usage: "application annotations",
			},
		},
	}
}

func List() cli.Command {
	return cli.Command{
		Name:    "apps",
		Usage:   "List all applications",
		Aliases: []string{"a"},
		Action:  list,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "cursor",
				Usage: "pagination cursor",
			},
			cli.Int64Flag{
				Name:  "n",
				Usage: "number of apps to return",
				Value: int64(100),
			},
		},
	}
}

func Delete() cli.Command {
	return cli.Command{
		Name:    "apps",
		Usage:   "Delete an application",
		Aliases: []string{"a"},
		Action:  delete,
	}
}

func Inspect() cli.Command {
	return cli.Command{
		Name:      "apps",
		Usage:     "retrieve one or all apps properties",
		Aliases:   []string{"a"},
		ArgsUsage: "<app> [property.[key]]",
		Action:    inspect,
	}
}

func Update() cli.Command {
	return cli.Command{
		Name:      "apps",
		Usage:     "update an application",
		Aliases:   []string{"a"},
		ArgsUsage: "<app>",
		Action:    update,
		Flags: []cli.Flag{
			cli.StringSliceFlag{
				Name:  "config,c",
				Usage: "route configuration",
			},
			cli.StringSliceFlag{
				Name:  "annotation",
				Usage: "application annotations",
			},
		},
	}
}

func SetConfig() cli.Command {
	return cli.Command{
		Name:      "apps",
		Usage:     "store a configuration key for this application",
		Aliases:   []string{"a"},
		ArgsUsage: "<app> <key> <value>",
		Action:    setConfig,
	}
}

func ListConfig() cli.Command {
	return cli.Command{
		Name:      "apps",
		Usage:     "list configuration key/value pairs for this application",
		Aliases:   []string{"a"},
		ArgsUsage: "<app>",
		Action:    listConfig,
	}
}

func GetConfig() cli.Command {
	return cli.Command{
		Name:      "apps",
		Usage:     "inspect configuration key for this application",
		Aliases:   []string{"a"},
		ArgsUsage: "<app> <key>",
		Action:    getConfig,
	}
}

func UnsetConfig() cli.Command {
	return cli.Command{
		Name:      "apps",
		Usage:     "remove a configuration key for this application",
		Aliases:   []string{"a"},
		ArgsUsage: "<app> <key>",
		Action:    unsetConfig,
	}
}
