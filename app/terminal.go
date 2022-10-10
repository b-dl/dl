package app

import (
	"fmt"
	"path"

	"github.com/b-dl/dl/logger"
	"github.com/urfave/cli/v2"
)

func terminalAction(c *cli.Context) error {
	cli.VersionPrinter(c)
	logger.Init(path.Join("log"), c.String("level"))

	cookie := c.String("cookie")

	fmt.Println(cookie)

	return nil
}
