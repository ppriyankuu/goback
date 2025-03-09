package backup

import (
	"fmt"

	"github.com/ppriyankuu/goback/internals/cli"
	"github.com/ppriyankuu/goback/internals/storage"
)

func Restore(source, destination string) error {
	_, err := cli.LoadConfig("config.yaml")
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	recentBackup, err := storage.GetRecentBackup(destination)
	if err != nil {
		return fmt.Errorf("failed to get recent backup: %w", err)
	}

	if err := storage.ExtractArchive(recentBackup.Path, destination); err != nil {
		return fmt.Errorf("failed to extract backup: %w", err)
	}

	cli.TrackProgress("Restored from: %s", recentBackup)

	return nil
}
