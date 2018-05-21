package commands

import (
	"github.com/fnproject/cli/common"
	"github.com/fnproject/cli/objects/app"
	"github.com/fnproject/cli/objects/context"
	"github.com/fnproject/cli/objects/log"
	"github.com/fnproject/cli/objects/route"
	"github.com/fnproject/cli/run"
	"github.com/urfave/cli"
)

type cmd map[string]cli.Command

var Commands = cmd{
	"list":    ListCommand(),
	"create":  CreateCommand(),
	"delete":  DeleteCommand(),
	"unset":   UnsetCommand(),
	"update":  UpdateCommand(),
	"use":     UseCommand(),
	"inspect": InspectCommand(),
	"call":    CallCommand(),
	"set":     SetCommand(),
	"get":     GetCommand(),
	"build":   BuildCommand(),
	"bump":    common.BumpCommand(),
	"deploy":  DeployCommand(),
	"push":    PushCommand(),
	"run":     run.RunCommand(),
	"test":    TestCommand(),
	"version": VersionCommand(),
	"start":   StartCommand(),
}

var CreateCmds = cmd{
	"apps":    app.GetCreateAppCommand(),
	"routes":  route.GetCreateRouteCommand(),
	"context": context.GetCreateContextCommand(),
}

var DeleteCmds = cmd{
	"apps":    app.GetDeleteAppCommand(),
	"routes":  route.GetDeleteRouteCommand(),
	"context": context.GetDeleteContextCommand(),
}

var ListCmds = cmd{
	"apps":   app.GetListAppsCommand(),
	"routes": route.GetListRoutesCommand(),
}

var CallCmds = cmd{
	"routes": route.GetCallRoutesCommand(),
}

var UseCmds = cmd{
	"context": context.GetUseContextCommand(),
}

var UnsetCmds = cmd{
	"apps":    app.GetUnsetConfigAppsCommand(),
	"routes":  route.GetUnsetConfigRoutesCommand(),
	"context": context.GetUnsetContextCommand(),
}

var UpdateCmds = cmd{
	"apps":    app.GetUpdateAppCommand(),
	"routes":  route.GetUpdateRouteCommand(),
	"context": context.GetUseContextCommand(),
}

var InspectCmds = cmd{
	"apps":   app.GetInspectAppsCommand(),
	"routes": route.GetInspectRoutesCommand(),
}

var GetCmds = cmd{
	"apps":   app.GetGetConfigAppsCommand(),
	"routes": route.GetGetConfigRoutesCommand(),
	"logs":   log.GetGetLogsCommand(),
}

var SetCmds = cmd{
	"apps":   app.GetSetConfigAppsCommnd(),
	"routes": route.GetSetConfigRoutesCommand(),
}

func GetCommands(commands map[string]cli.Command) []cli.Command {
	cmds := []cli.Command{}
	for _, cmd := range commands {
		cmds = append(cmds, cmd)
	}
	return cmds
}
