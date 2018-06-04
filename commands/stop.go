package commands

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/urfave/cli"
)

// StopCommand returns stop server cli.command
func StopCommand() cli.Command {
	return cli.Command{
		Name:     "stop",
		Usage:    "stops a functions server",
		Category: "SERVER COMMANDS",
		Action:   stop,
	}
}
func stop(c *cli.Context) error {
	cmd := exec.Command("docker", "stop", "fnserver")
	err := cmd.Run()
	if err != nil {
		return errors.New("failed to stop 'fnserver'")
	}

	fmt.Println("Successfully stopped 'fnserver'")

	return err
}
