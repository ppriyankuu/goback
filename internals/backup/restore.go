package backup

import (
	"fmt"

	"github.com/ppriyankuu/goback/internals/cli"
	"github.com/ppriyankuu/goback/internals/storage"
)

// Restore restores data from the most recent backup to the specified destination.

// Parameters:
// - source: The source directory or file to restore.
// - destination: The destination directory where the backup will be restored.

// Returns:
// - error: An error if any step in the restore process fails.
func Restore(source, destination string) error {
	// Load the configuration from the specified YAML file.
	_, err := cli.LoadConfig("config.yaml")
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	// Retrieve the most recent backup for the specified destination.
	recentBackup, err := storage.GetRecentBackup(destination)
	if err != nil {
		return fmt.Errorf("failed to get recent backup: %w", err)
	}

	// Extract the most recent backup archive to the destination directory.
	if err := storage.ExtractArchive(recentBackup.Path, destination); err != nil {
		return fmt.Errorf("failed to extract backup: %w", err)
	}

	// Track and log the progress of the restore operation.
	cli.TrackProgress("Restored from: %s", recentBackup)

	return nil
}
