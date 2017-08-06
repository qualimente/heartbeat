package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "heartbeat"
	app.Usage = "Produce a periodic heartbeat message"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.BoolTFlag{
			Name:  "verbose",
			Usage: "Print verbose output",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "Print the version and exit",
			Action: func(c *cli.Context) error {
				fmt.Printf("heartbeat version %s\n", app.Version)
				return nil
			},
		},
		{
			Name:   "run",
			Usage:  "Run the program",
			Action: RunPeriodically,
		},
	}

	app.Run(os.Args)
}
