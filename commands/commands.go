package commands

import (
	"github.com/fnproject/cli/common"
	"github.com/fnproject/cli/objects/app"
	"github.com/fnproject/cli/objects/call"
	"github.com/fnproject/cli/objects/context"
	"github.com/fnproject/cli/objects/log"
	"github.com/fnproject/cli/objects/route"
	"github.com/fnproject/cli/objects/server"
	"github.com/fnproject/cli/run"
	"github.com/urfave/cli"
)

type cmd map[string]cli.Command

// Commands map of all top-level commands
var Commands = cmd{
	"build":        BuildCommand(),
	"build-server": BuildServerCommand(),
	"bump":         common.BumpCommand(),
	"call":         CallCommand(),
	"configure":    ConfigureCommand(),
	"create":       CreateCommand(),
	"delete":       DeleteCommand(),
	"deploy":       DeployCommand(),
	"get":          GetCommand(),
	"init":         InitCommand(),
	"inspect":      InspectCommand(),
	"lamba":        LambdaCommand(),
	"list":         ListCommand(),
	"push":         PushCommand(),
	"run":          run.RunCommand(),
	"start":        StartCommand(),
	"stop":         StopCommand(),
	"test":         TestCommand(),
	"unset":        UnsetCommand(),
	"update":       UpdateCommand(),
	"use":          UseCommand(),
}

// CreateCmds cli.command map
var CreateCmds = cmd{
	"apps":    app.Create(),
	"routes":  route.Create(),
	"context": context.Create(),
}

// ConfigCmds cli.command map
var ConfigCmds = cmd{
	"apps":   app.SetConfig(),
	"routes": route.SetConfig(),
}

// ConfigListCmds cli.command map
var ConfigListCmds = cmd{
	"apps":   app.ListConfig(),
	"routes": route.ListConfig(),
}

// ConfigGetCmds cli.command map
var ConfigGetCmds = cmd{
	"apps":   app.GetConfig(),
	"routes": route.GetConfig(),
}

// ConfigSetCmds cli.command map
var ConfigSetCmds = cmd{
	"apps":   app.SetConfig(),
	"routes": route.SetConfig(),
}

// ConfigUnsetCmds cli.command map
var ConfigUnsetCmds = cmd{
	"apps":   app.UnsetConfig(),
	"routes": route.UnsetConfig(),
}

// DeleteCmds cli.command map
var DeleteCmds = cmd{
	"apps":    app.Delete(),
	"routes":  route.Delete(),
	"context": context.Delete(),
}

// GetCmds cli.command map
var GetCmds = cmd{
	"config": ConfigCommand("get"),
	"logs":   log.Get(),
	"calls":  call.Get(),
}

// InspectCmds cli.command map
var InspectCmds = cmd{
	"apps":   app.Inspect(),
	"routes": route.Inspect(),
}

// ListCmds cli.command map
var ListCmds = cmd{
	"config":  ConfigCommand("list"),
	"apps":    app.List(),
	"routes":  route.List(),
	"calls":   call.List(),
	"context": context.List(),
}

// UnsetCmds cli.command map
var UnsetCmds = cmd{
	"config":  ConfigCommand("unset"),
	"context": context.Unset(),
}

// UpdateCmds cli.command map
var UpdateCmds = cmd{
	"apps":    app.Update(),
	"routes":  route.Update(),
	"context": context.Update(),
	"server":  server.Update(),
}

// UseCmds cli.command map
var UseCmds = cmd{
	"context": context.Use(),
}

// GetCommands returns a list of cli.commands
func GetCommands(commands map[string]cli.Command) []cli.Command {
	cmds := []cli.Command{}
	for _, cmd := range commands {
		cmds = append(cmds, cmd)
	}
	return cmds
}
