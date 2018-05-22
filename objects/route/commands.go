package route

import "github.com/urfave/cli"

func CallRoutes() cli.Command {
	return cli.Command{
		Name:      "routes",
		Usage:     "call a route",
		Aliases:   []string{"rt"},
		ArgsUsage: "<app> </path> [image]",
		Action:    Call,
		Flags:     callFnFlags,
	}
}

func Create() cli.Command {
	return cli.Command{
		Name:      "routes",
		Usage:     "Create a route in an application",
		Aliases:   []string{"rt"},
		ArgsUsage: "<app> </path>",
		Action:    create,
		Flags:     RouteFlags,
	}
}

func List() cli.Command {
	return cli.Command{
		Name:      "routes",
		Usage:     "list routes for `app`",
		Aliases:   []string{"rt"},
		ArgsUsage: "<app>",
		Action:    list,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "cursor",
				Usage: "pagination cursor",
			},
			cli.Int64Flag{
				Name:  "n",
				Usage: "number of routes to return",
				Value: int64(100),
			},
		},
	}
}

func Delete() cli.Command {
	return cli.Command{
		Name:      "routes",
		Usage:     "Delete a route from an application `app`",
		Aliases:   []string{"rt"},
		ArgsUsage: "<app> </path>",
		Action:    delete,
	}
}

func Inspect() cli.Command {
	return cli.Command{
		Name:      "routes",
		Usage:     "retrieve one or all routes properties",
		Aliases:   []string{"rt"},
		ArgsUsage: "<app> </path> [property.[key]]",
		Action:    inspect,
	}
}

func Update() cli.Command {
	return cli.Command{
		Name:      "routes",
		Usage:     "Update a Route in an `app`",
		Aliases:   []string{"rt"},
		ArgsUsage: "<app> </path>",
		Action:    update,
		Flags:     updateRouteFlags,
	}
}

func CallRoute() cli.Command {
	return cli.Command{
		Name:      "call",
		Usage:     "call a remote function",
		Aliases:   []string{"rt"},
		ArgsUsage: "<app> </path>",
		Flags:     callFnFlags,
		Action:    Call,
	}
}

func GetConfig() cli.Command {
	return cli.Command{
		Name:      "routes",
		Usage:     "inspect configuration key for this route",
		Aliases:   []string{"rt"},
		ArgsUsage: "<app> </path> <key>",
		Action:    getConfig,
	}
}
func SetConfig() cli.Command {
	return cli.Command{
		Name:      "routes",
		Usage:     "store a configuration key for this route",
		Aliases:   []string{"rt"},
		ArgsUsage: "<app> </path> <key> <value>",
		Action:    setConfig,
	}
}
func ListConfig() cli.Command {
	return cli.Command{
		Name:      "routes",
		Usage:     "list configuration key/value pairs for this route",
		Aliases:   []string{"rt"},
		ArgsUsage: "<app> </path>",
		Action:    listConfig,
	}
}
func UnsetConfig() cli.Command {
	return cli.Command{
		Name:      "routes",
		Usage:     "remove a configuration key for this route",
		Aliases:   []string{"rt"},
		ArgsUsage: "<app> </path> <key>",
		Action:    unsetConfig,
	}
}
