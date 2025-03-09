package main

import (
	"log"
	"os"

	"github.com/ppriyankuu/goback/internals/backup"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "goback",
		Usage: "A command-line backup utility tool",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Path to the configuration file",
				Value:   "config.yaml",
			},
			&cli.StringFlag{
				Name:    "source",
				Aliases: []string{"s"},
				Usage:   "Source directory to back up",
			},
			&cli.StringFlag{
				Name:    "destination",
				Aliases: []string{"d"},
				Usage:   "Destination directory for backups",
			},
			&cli.BoolFlag{
				Name:    "incremental",
				Aliases: []string{"i"},
				Usage:   "Enable incremental backup",
			},
			&cli.BoolFlag{
				Name:    "restore",
				Aliases: []string{"r"},
				Usage:   "Restore from backup",
			},
		},
		Action: func(c *cli.Context) error {
			configPath := c.String("config")
			source := c.String("source")
			destination := c.String("destination")
			incremental := c.Bool("incremental")
			restore := c.Bool("restore")

			if source == "" || destination == "" {
				log.Fatalf("Source and destination directories are required")
			}

			if restore {
				return backup.Restore(source, destination)
			}

			return backup.Backup(source, destination, incremental, configPath)
		},
		Before: func(c *cli.Context) error {
			if c.Bool("help") {
				cli.ShowAppHelp(c)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
