package log

import "github.com/urfave/cli"

func Get() cli.Command {
	return cli.Command{
		Name:      "logs",
		Usage:     "get logs for a call. Must provide call_id or last (l) to retrieve the most recent calls logs.",
		Aliases:   []string{"lg"},
		ArgsUsage: "<app> <call-id>",
		Action:    get,
	}
}
