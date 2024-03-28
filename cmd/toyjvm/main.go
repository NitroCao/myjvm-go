package main

import (
	"fmt"
	"os"
	"toyjvm/pkg/cmd"
)

var version string

func main() {
	app := cmd.NewApp(version)
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
