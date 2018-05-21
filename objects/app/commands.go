package app

import (
	"github.com/fnproject/cli/client"
	"github.com/urfave/cli"
)

type app client.FnClient

func GetCreateAppCommand() cli.Command {
	return cli.Command{
		Name:      "apps",
		Usage:     "Create a new application",
		ArgsUsage: "<app>",
		Action:    createApp,
		Flags: []cli.Flag{
			cli.StringSliceFlag{
				Name:  "config",
				Usage: "application configuration",
			},
		},
	}
}

func GetListAppsCommand() cli.Command {
	return cli.Command{
		Name:   "apps",
		Usage:  "List all applications ",
		Action: listApps,
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

func GetDeleteAppCommand() cli.Command {
	return cli.Command{
		Name:   "apps",
		Usage:  "Delete an application",
		Action: deleteApps,
	}
}

func GetInspectAppsCommand() cli.Command {
	return cli.Command{
		Name:      "apps",
		Usage:     "retrieve one or all apps properties",
		ArgsUsage: "<app> [property.[key]]",
		Action:    inspectApps,
	}
}

func GetUpdateAppCommand() cli.Command {
	return cli.Command{
		Name:      "apps",
		Usage:     "update an application",
		ArgsUsage: "<app>",
		Action:    updateApps,
		Flags: []cli.Flag{
			cli.StringSliceFlag{
				Name:  "config,c",
				Usage: "route configuration",
			},
		},
	}
}

func GetSetConfigAppsCommnd() cli.Command {
	return cli.Command{
		Name:      "apps",
		Usage:     "store a configuration key for this application",
		ArgsUsage: "<app> <key> <value>",
		Action:    configSetApps,
	}
}

func GetListConfigAppsCommand() cli.Command {
	return cli.Command{
		Name:      "apps",
		Usage:     "list configuration key/value pairs for this application",
		ArgsUsage: "<app>",
		Action:    configListApps,
	}
}

func GetGetConfigAppsCommand() cli.Command {
	return cli.Command{
		Name:      "apps",
		Usage:     "inspect configuration key for this application",
		ArgsUsage: "<app> <key>",
		Action:    configGetApps,
	}
}

func GetUnsetConfigAppsCommand() cli.Command {
	return cli.Command{
		Name:      "apps",
		Usage:     "remove a configuration key for this application",
		ArgsUsage: "<app> <key>",
		Action:    configUnsetApps,
	}
}
