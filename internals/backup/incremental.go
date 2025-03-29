package backup

import (
	"fmt"
	"time"

	"github.com/ppriyankuu/goback/internals/fs"
	"github.com/ppriyankuu/goback/internals/storage"
)

// IncrementalBackup performs an incremental backup by detecting changes since the last backup.

// Parameters:
// - source: The source directory or file to back up.
// - destination: The destination directory where the backup will be stored.

// Returns:
// - error: An error if any step in the process fails.
func IncrementalBackup(source, destination string) error {
	// Retrieve the most recent metadata for the specified destination.
	metadata, err := storage.GetRecentMetadata(destination)
	if err != nil {
		return fmt.Errorf("failed to get recent metadata: %w", err)
	}

	// Detect changes in the source directory since the last backup.
	changes, err := fs.DetectChanges(source, metadata.Source)
	if err != nil {
		return fmt.Errorf("failed to detect changes: %w", err)
	}

	// Create an incremental archive containing only the detected changes.
	archivePath, err := storage.CreateIncrementalArchive(destination, source, changes)
	if err != nil {
		return fmt.Errorf("failed to create incremental archive: %w", err)
	}

	// Create new metadata for the incremental backup.
	newMetadata := storage.Metadata{
		Source:      source,
		Destination: destination,
		Path:        archivePath,
		Time:        time.Now(),
	}

	// Store the new metadata for future reference
	if err := storage.StoreMetadata(newMetadata); err != nil {
		return fmt.Errorf("failed to store metadata: %w", err)
	}

	return nil
}
