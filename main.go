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
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "format",
					Usage: "Specify the output format: text, json",
				}, cli.IntFlag{
					Name:  "period",
					Usage: "Specify the period between heartbeat messages as an integer in seconds, default: 1 second",
					Value: 1,
				},
			},
		},
	}

	app.Run(os.Args)
}
