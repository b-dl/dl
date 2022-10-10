package app

import (
	"fmt"
	"path"

	"github.com/b-dl/dl/logger"
	"github.com/b-dl/dl/request"
	"github.com/urfave/cli/v2"
)

var terminalFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "cookie",
		Aliases: []string{"c"},
		Usage:   "Cookie",
	},
	&cli.BoolFlag{
		Name:    "playlist",
		Aliases: []string{"p"},
		Usage:   "Download playlist",
	},
	&cli.StringFlag{
		Name:    "user-agent",
		Aliases: []string{"u"},
		Usage:   "Use specified User-Agent",
	},
	&cli.StringFlag{
		Name:    "refer",
		Aliases: []string{"r"},
		Usage:   "Use specified Referrer",
	},
	&cli.StringFlag{
		Name:    "format",
		Aliases: []string{"f"},
		Usage:   "Select specific format to download",
	},
	&cli.StringFlag{
		Name:    "file-name",
		Aliases: []string{"fn", "FN"},
		Usage:   "Specify the output file name",
	},
	&cli.BoolFlag{
		Name:    "multi-thread",
		Aliases: []string{"m"},
		Usage:   "Multiple threads to download single video",
	},
	&cli.UintFlag{
		Name:    "thread",
		Aliases: []string{"n"},
		Value:   10,
		Usage:   "The number of download thread (For multi-video only)",
	},
}

func terminalAction(c *cli.Context) error {
	cli.VersionPrinter(c)
	logger.Init(path.Join("log"), c.String("level"))

	cookie := c.String("cookie")

	request.SetOptions(request.RequestOptions{
		RetryTimes: int(c.Uint("retry-times")),
		Timeout:    int64(c.Uint("timeout")),
		UserAgent:  c.String("user-agent"),
		Refer:      c.String("refer"),
		Cookie:     cookie,
	})

	fmt.Println(cookie)

	return nil
}
