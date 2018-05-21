package log

import "github.com/urfave/cli"

func GetGetLogsCommand() cli.Command {
	return cli.Command{
		Name:      "logs",
		Usage:     "get logs for a call. Must provide call_id or last (l) to retrieve the most recent calls logs.",
		ArgsUsage: "<app> <call-id>",
		Action:    get,
	}
}
