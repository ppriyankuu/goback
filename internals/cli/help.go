package cli

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func ShowHelp(c *cli.Context) {
	fmt.Println("Usage: goback [options]")
	fmt.Println("Options:")
	fmt.Println("  -c, --config <file>     Path to the configuration file (default: config.yaml)")
	fmt.Println("  -s, --source <dir>	   Source directory to back up")
	fmt.Println("  -d, --destination <dir> Destination directory for backups")
	fmt.Println("  -i, --incremental       Enable incremental backup")
	fmt.Println("  -r, --restore           Restore from backup")
	fmt.Println("  -h, --help              Show help documentation")
	os.Exit(0)
}
