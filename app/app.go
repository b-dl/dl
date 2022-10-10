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
				Name:  "cmd",
				Usage: "Terminal downloader",
				Flags: append([]cli.Flag{
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
						Name:    "stream-format",
						Aliases: []string{"f"},
						Usage:   "Select specific stream to download",
					},
					&cli.StringFlag{
						Name:    "output-name",
						Aliases: []string{"on", "ON"},
						Usage:   "Specify the output file name",
					},
					&cli.UintFlag{
						Name:  "start",
						Value: 1,
						Usage: "Define the starting item of a playlist or a file input",
					},
					&cli.UintFlag{
						Name:  "end",
						Value: 0,
						Usage: "Define the ending item of a playlist or a file input",
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
						Usage:   "The number of download thread (only works for multiple-parts video)",
					},
				}, baseFlags...),
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
