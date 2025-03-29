package backup

import (
	"fmt"
	"time"

	"github.com/ppriyankuu/goback/internals/cli"
	"github.com/ppriyankuu/goback/internals/storage"
)

// Backup performs a full or incremental backup based on the provided flag.

// Parameters:
// - source: The source directory or file to backup.
// - destination: The destination directory where the backup will be stored.
// - incremental: A boolean indicating whether to perform an incremental backup.
// - configPath: The path to the config file.

// Returns:
// - error: An error if performing the backup fails.
func Backup(source, destination string, incremental bool, configPath string) error {
	// Load the configuration from the specified file.
	config, err := cli.LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	// Create a backup archive, either full or incremental based on the flag.
	archivePath, err := storage.CreateArchive(destination, source, incremental)
	if err != nil {
		return fmt.Errorf("failed to create backup archive: %w", err)
	}

	// Track and log the progress of the backup operation
	cli.TrackProgress("Backup created at: %s", archivePath)

	// Create metadata for the backup
	metadata := storage.Metadata{
		Source:      source,
		Destination: destination,
		Path:        archivePath,
		Time:        time.Now(),
	}

	// Store the metadata for future reference.
	if err := storage.StoreMetadata(metadata); err != nil {
		return fmt.Errorf("failed to store metadata: %w", err)
	}

	// Verify the integrity of the backup archive.
	if err := storage.VerifyBackup(archivePath); err != nil {
		return fmt.Errorf("failed to verify backup: %w", err)
	}

	// Clean up old backups based on the retention policy (time period)
	if err := storage.CleanupOldBackups(destination, config.RetentionDays); err != nil {
		return fmt.Errorf("failed to clean up old backups: %w", err)
	}

	return nil
}
