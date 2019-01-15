package main

import (
	"fmt"
	"os"

	"github.com/k-kurikuri/go2ree/action"
	"github.com/urfave/cli"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	var (
		dirOnly bool
	)

	app := cli.NewApp()
	app.Name = "go2ree"
	app.Usage = "Golang For tree Command"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "d",
			Usage:       "TreeAction directories only.",
			Destination: &dirOnly,
		},
	}

	app.Action = action.TreeAction

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		fmt.Print("0 directories, 0 files")
		return 1
	}

	return 0
}
