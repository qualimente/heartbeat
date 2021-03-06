package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"time"
)

func RunPeriodically(c *cli.Context) error {

	log.SetFormatter(_makeFormatter(c.String("format")))

	log.WithFields(log.Fields{
		"appName": c.App.Name,
	}).Info("Running periodically")

	period := c.Duration("period")

	for {
		go func() {
			PrintHeartbeat()
		}()

		time.Sleep(period)
	}

}

func PrintHeartbeat() {
	log.WithFields(log.Fields{
		"type": "heartbeat",
	}).Info("Every heartbeat bears your name")
}

func _makeFormatter(format string) log.Formatter {
	switch format {
	case "text":
		return &log.TextFormatter{DisableColors: true}
	case "json":
		return &log.JSONFormatter{}
	default:
		return &log.JSONFormatter{}
	}
}
