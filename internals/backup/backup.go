package backup

import (
	"fmt"
	"time"

	"github.com/ppriyankuu/goback/internals/cli"
	"github.com/ppriyankuu/goback/internals/storage"
)

func Backup(source, destination string, incremental bool, configPath string) error {
	config, err := cli.LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	archivePath, err := storage.CreateArchive(destination, source, incremental)
	if err != nil {
		return fmt.Errorf("failed to create backup archive: %w", err)
	}

	cli.TrackProgress("Backup created at: %s", archivePath)

	metadata := storage.Metadata{
		Source:      source,
		Destination: destination,
		Path:        archivePath,
		Time:        time.Now(),
	}
	if err := storage.StoreMetadata(metadata); err != nil {
		return fmt.Errorf("failed to store metadata: %w", err)
	}

	if err := storage.VerifyBackup(archivePath); err != nil {
		return fmt.Errorf("failed to verify backup: %w", err)
	}

	if err := storage.CleanupOldBackups(destination, config.RetentionDays); err != nil {
		return fmt.Errorf("failed to clean up old backups: %w", err)
	}

	return nil
}
