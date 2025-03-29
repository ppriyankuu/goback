package main

import (
	"log"
	"os"

	"github.com/ppriyankuu/goback/internals/backup"
	"github.com/urfave/cli/v2"
)

// main function initializes and runs the CLI application
func main() {
	// Define the CLI application
	app := &cli.App{
		Name:  "goback",                             // Application name
		Usage: "A command-line backup utility tool", // Brief description

		// Defile the CLI flags for configuration and operations.
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config", // Path to the config file
				Aliases: []string{"c"},
				Usage:   "Path to the configuration file",
				Value:   "config.yaml", // Default config file
			},
			&cli.StringFlag{
				Name:    "source", // Source directory to backup
				Aliases: []string{"s"},
				Usage:   "Source directory to back up",
			},
			&cli.StringFlag{
				Name:    "destination", // Destination directory for backups
				Aliases: []string{"d"},
				Usage:   "Destination directory for backups",
			},
			&cli.BoolFlag{
				Name:    "incremental", // Toggle for incremental backup
				Aliases: []string{"i"},
				Usage:   "Enable incremental backup",
			},
			&cli.BoolFlag{
				Name:    "restore", // Toggle for restoration
				Aliases: []string{"r"},
				Usage:   "Restore from backup",
			},
		},

		// Define the main action for the CLI
		Action: func(c *cli.Context) error {
			// Retrieve flag values.
			configPath := c.String("config")
			source := c.String("source")
			destination := c.String("destination")
			incremental := c.Bool("incremental")
			restore := c.Bool("restore")

			// Ensure essentials flags are provided.
			if source == "" || destination == "" {
				log.Fatalf("Source and destination directories are required")
			}

			// Perform restore or backup based on the flag.
			if restore {
				return backup.Restore(source, destination)
			}

			return backup.Backup(source, destination, incremental, configPath)
		},

		// Hook to run before the main action.
		Before: func(c *cli.Context) error {
			if c.Bool("help") {
				cli.ShowAppHelp(c) // Show help message if requested.
			}
			return nil
		},
	}

	// Run the application and handle errors.
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err) // Log fatal error if the app fails.
	}
}
