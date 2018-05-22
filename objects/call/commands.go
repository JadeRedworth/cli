package call

import "github.com/urfave/cli"

func Get() cli.Command {
	return cli.Command{
		Name:      "calls",
		Usage:     "get function call info per app",
		Aliases:   []string{"cl"},
		ArgsUsage: "<app> <call-id>",
		Action:    get,
	}
}

func List() cli.Command {
	return cli.Command{
		Name:      "calls",
		Usage:     "list all calls for the specific app. Route is optional",
		Aliases:   []string{"cl"},
		ArgsUsage: "<app>",
		Action:    list,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "path",
				Usage: "function's path",
			},
			cli.StringFlag{
				Name:  "cursor",
				Usage: "pagination cursor",
			},
			cli.StringFlag{
				Name:  "from-time",
				Usage: "'start' timestamp",
			},
			cli.StringFlag{
				Name:  "to-time",
				Usage: "'stop' timestamp",
			},
			cli.Int64Flag{
				Name:  "n",
				Usage: "number of calls to return",
				Value: int64(100),
			},
		},
	}
}
