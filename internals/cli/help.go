package cli

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// ShowHelp displays the usage information and available options for the application.

// Parameters:
// - c: The CLI content that contains command-line arguments and flags.
func ShowHelp(c *cli.Context) {
	// Print the usage information.
	fmt.Println("Usage: goback [options]")

	// List the available options.
	fmt.Println("Options:")
	fmt.Println("  -c, --config <file>     Path to the configuration file (default: config.yaml)")
	fmt.Println("  -s, --source <dir>	   Source directory to back up")
	fmt.Println("  -d, --destination <dir> Destination directory for backups")
	fmt.Println("  -i, --incremental       Enable incremental backup")
	fmt.Println("  -r, --restore           Restore from backup")
	fmt.Println("  -h, --help              Show help documentation")

	// Exit the program after showing help.
	os.Exit(0)
}
