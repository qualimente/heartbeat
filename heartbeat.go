package main

import (
	"fmt"
	"github.com/urfave/cli"
	"time"
)

func RunPeriodically(c *cli.Context) error {
	var period time.Duration = 1 * time.Second

	fmt.Printf("Running %s periodically\n", c.App.Name)
	for {
		go func() {
			PrintHeartbeat()
		}()

		time.Sleep(period)
	}

	return nil
}

func PrintHeartbeat() {
	now := time.Now().UTC()
	fmt.Printf("%v Every heartbeat bears your name\n", now.Format(time.RFC3339))
}
