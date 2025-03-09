package backup

import (
	"fmt"
	"time"

	"github.com/ppriyankuu/goback/internals/fs"
	"github.com/ppriyankuu/goback/internals/storage"
)

func IncrementalBackup(source, destination string) error {
	metadata, err := storage.GetRecentMetadata(destination)
	if err != nil {
		return fmt.Errorf("failed to get recent metadata: %w", err)
	}

	changes, err := fs.DetectChanges(source, metadata.Source)
	if err != nil {
		return fmt.Errorf("failed to detect changes: %w", err)
	}

	archivePath, err := storage.CreateIncrementalArchive(destination, source, changes)
	if err != nil {
		return fmt.Errorf("failed to create incremental archive: %w", err)
	}

	newMetadata := storage.Metadata{
		Source:      source,
		Destination: destination,
		Path:        archivePath,
		Time:        time.Now(),
	}

	if err := storage.StoreMetadata(newMetadata); err != nil {
		return fmt.Errorf("failed to store metadata: %w", err)
	}

	return nil
}
