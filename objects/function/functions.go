package function

import (
	"fmt"
	"os"

	"github.com/fnproject/cli/common"
	"github.com/urfave/cli"
)

// build will take the found valid function and build it
func (b *buildcmd) Build(c *cli.Context) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	fpath, ff, err := common.FindAndParseFuncfile(path)
	if err != nil {
		return err
	}

	buildArgs := c.StringSlice("build-arg")
	ff, err = common.BuildFunc(c, fpath, ff, buildArgs, b.noCache)
	if err != nil {
		return err
	}

	fmt.Printf("Function %v built successfully.\n", ff.ImageName())
	return nil
}
