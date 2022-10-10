package app

import (
	"fmt"
	"os"
	"sort"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var Version = "0.0.1"

const (
	Name = "dl"
)

var baseFlags = []cli.Flag{
	&cli.UintFlag{
		Name:    "timeout",
		Aliases: []string{"t", "T"},
		Value:   20,
		Usage:   "Request timed out (s)",
		EnvVars: []string{"DL_TIMEOUT"},
	},
	&cli.UintFlag{
		Name:    "retry-times",
		Aliases: []string{"rt", "RT"},
		Value:   5,
		Usage:   "How many times to retry when the download failed",
		EnvVars: []string{"DL_RETRY_TIMES"},
	},
	&cli.StringFlag{
		Name:    "level",
		Aliases: []string{"l", "L"},
		Value:   "info",
		Usage:   "Log level",
		EnvVars: []string{"DL_LEVEL"},
	},
	&cli.StringFlag{
		Name:    "output-path",
		Aliases: []string{"op", "OP"},
		Value:   ".",
		Usage:   "Specify the output path",
		EnvVars: []string{"DL_PATH"},
	},
	&cli.UintFlag{
		Name:    "chunk-size",
		Aliases: []string{"cs", "CS"},
		Value:   1,
		Usage:   "HTTP chunk size for downloading (in MB)",
		EnvVars: []string{"DL_CHUNK_SIZE"},
	},
}

func init() {
	cli.VersionPrinter = func(c *cli.Context) {
		blue := color.New(color.FgBlue)
		cyan := color.New(color.FgCyan)
		fmt.Fprintf(
			color.Output,
			"%s: version %s, A fast browser download parser\n",
			cyan.Sprintf(Name),
			blue.Sprintf(c.App.Version),
		)
	}
}

func New() *cli.App {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(
			color.Output,
			"%v\n",
			color.New(color.FgRed).Sprintln(err.Error()),
		)
		os.Exit(1)
	}

	logrus.SetReportCaller(true)

	app := &cli.App{
		Name:    Name,
		Usage:   "A fast browser download parser.",
		Version: Version,
		Commands: []*cli.Command{
			{
				Name:   "cmd",
				Usage:  "Terminal downloader",
				Flags:  append(terminalFlags, baseFlags...),
				Action: terminalAction,
			},
			{
				Name:  "serve",
				Usage: "Deploy as a service",
				Flags: append([]cli.Flag{
					&cli.UintFlag{
						Name:    "port",
						Aliases: []string{"p", "P"},
						Value:   10101,
						Usage:   "Service port",
						EnvVars: []string{"DL_PORT"},
					},
				}, baseFlags...),
				Action: serveAction,
			},
		},
		UseShortOptionHandling: true,
		EnableBashCompletion:   true,
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	return app
}
