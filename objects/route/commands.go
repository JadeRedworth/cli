package route

import "github.com/urfave/cli"

func GetCallRoutesCommand() cli.Command {
	return cli.Command{
		Name:      "routes",
		Usage:     "call a route",
		ArgsUsage: "<app> </path> [image]",
		Action:    Call,
		Flags:     callFnFlags,
	}
}

func GetCreateRouteCommand() cli.Command {
	return cli.Command{
		Name:      "routes",
		Usage:     "Create a route in an application",
		ArgsUsage: "<app> </path>",
		Action:    create,
		Flags:     RouteFlags,
	}
}

func GetListRoutesCommand() cli.Command {
	return cli.Command{
		Name:      "routes",
		Usage:     "list routes for `app`",
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

func GetDeleteRouteCommand() cli.Command {
	return cli.Command{
		Name:      "routes",
		Usage:     "Delete a route from an application `app`",
		ArgsUsage: "<app> </path>",
		Action:    delete,
	}
}

func GetInspectRoutesCommand() cli.Command {
	return cli.Command{
		Name:      "routes",
		Usage:     "retrieve one or all routes properties",
		ArgsUsage: "<app> </path> [property.[key]]",
		Action:    inspect,
	}
}

func GetGetConfigRoutesCommand() cli.Command {
	return cli.Command{
		Name:      "routes",
		Usage:     "inspect configuration key for this route",
		ArgsUsage: "<app> </path> <key>",
		Action:    configGet,
	}
}
func GetSetConfigRoutesCommand() cli.Command {
	return cli.Command{
		Name:      "routes",
		Usage:     "store a configuration key for this route",
		ArgsUsage: "<app> </path> <key> <value>",
		Action:    configSet,
	}
}
func GetListConfigRoutesCommand() cli.Command {
	return cli.Command{
		Name:      "routes",
		Usage:     "list configuration key/value pairs for this route",
		ArgsUsage: "<app> </path>",
		Action:    configList,
	}
}
func GetUnsetConfigRoutesCommand() cli.Command {
	return cli.Command{
		Name:      "routes",
		Usage:     "remove a configuration key for this route",
		ArgsUsage: "<app> </path> <key>",
		Action:    configUnset,
	}
}

func GetUpdateRouteCommand() cli.Command {
	return cli.Command{
		Name:      "routes",
		Aliases:   []string{"u"},
		Usage:     "Update a Route in an `app`",
		ArgsUsage: "<app> </path>",
		Action:    update,
		Flags:     updateRouteFlags,
	}
}

func Call() cli.Command {
	return cli.Command{
		Name:      "call",
		Usage:     "call a remote function",
		ArgsUsage: "<app> </path>",
		Flags:     callFnFlags,
		Action:    call,
	}
}
