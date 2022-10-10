package main

import (
	"fmt"
	"os"

	"github.com/b-dl/dl/app"
	"github.com/fatih/color"
)

func main() {
	if err := app.New().Run(os.Args); err != nil {
		fmt.Fprintf(
			color.Output,
			"Run %s failed: %s\n",
			color.CyanString("%s", app.Name), color.RedString("%v", err),
		)
		os.Exit(1)
	}
}
