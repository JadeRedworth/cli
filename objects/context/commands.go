package context

import "github.com/urfave/cli"

func GetCreateContextCommand() cli.Command {
	return cli.Command{
		Name:      "context",
		Usage:     "create a new context",
		ArgsUsage: "<context>",
		Action:    createContext,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "provider",
				Usage: "context provider",
			},
			cli.StringFlag{
				Name:  "api-url",
				Usage: "context api url",
			},
			cli.StringFlag{
				Name:  "registry",
				Usage: "context registry",
			},
		},
	}
}

func GetListContextCommand() cli.Command {
	return cli.Command{
		Name:   "contexts",
		Usage:  "list contexts",
		Action: listContext,
	}
}

func GetDeleteContextCommand() cli.Command {
	return cli.Command{
		Name:      "context",
		Usage:     "delete a context",
		ArgsUsage: "<context>",
		Action:    deleteCtx,
	}
}

func (ctxMap ContextMap) GetUpdateContextCommand() cli.Command {
	return cli.Command{
		Name:      "context",
		Usage:     "update context files",
		ArgsUsage: "<key> <value>",
		Action:    ctxMap.updateCtx,
	}
}

func GetUseContextCommand() cli.Command {
	return cli.Command{
		Name:      "context",
		Usage:     "use context for future invocations",
		ArgsUsage: "<context>",
		Action:    useCtx,
	}
}

func GetUnsetContextCommand() cli.Command {
	return cli.Command{
		Name:   "context",
		Usage:  "unset current-context",
		Action: unsetCtx,
	}
}
