package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/ppriyankuu/goback/internals/fs"
)

// CleanupOldBackups removes old backup files exceeding the specified retention period.

// Parameters:
// - destination: The directory containing the backup files.
// - retentionDays: The number of days to retain backups before deletion.

// Returns:
// - error: An error if directory traversal or file removal fails.
func CleanupOldBackups(destination string, retentionDays int) error {
	// Traverse the destination directory to get all files.
	files, err := fs.TraversalDirectory(destination)
	if err != nil {
		return fmt.Errorf("failed to traverse directory: %w", err)
	}

	// Collect metdata of all backup files.
	var backups []Metadata

	for _, file := range files {
		// Check if the file has a .zip extension
		if filepath.Ext(file) == ".zip" {
			metadata, err := GetMetadata(file)
			if err != nil {
				continue // Skip files with invalid metadata
			}
			backups = append(backups, metadata)
		}
	}

	// Sort backups by time, oldest to newest.
	sort.Slice(backups, func(i, j int) bool {
		return backups[i].Time.Before(backups[j].Time)
	})

	// Get the current time for comparison.
	now := time.Now()

	// Loop through backups and remove old ones.
	for _, backup := range backups {
		if now.Sub(backup.Time).Hours() > float64(retentionDays)*24 {
			if err := os.Remove(backup.Path); err != nil {
				return fmt.Errorf("failed to remove old backup: %w", err)
			}
		}
	}

	return nil
}

// GetMetadata retrieves metadata from a corresponding JSON file.

// Parameters:
// - archivePath: The path to the archive file.

// Returns:
// - Metadata: The metadata information extracted from the JSON file.
// - error: An error if reading or unmarshalling fails.
func GetMetadata(archivePath string) (Metadata, error) {
	// Construct the metadata file path.
	metadataPath := filepath.Join(filepath.Dir(archivePath), "metadata.json")

	// Read the metadata file
	data, err := os.ReadFile(metadataPath)
	if err != nil {
		return Metadata{}, fmt.Errorf("failed to read metadata file: %w", err)
	}

	var metadata Metadata

	// Unmarshal the JSON data into the Metadata struct.
	if err := json.Unmarshal(data, &metadata); err != nil {
		return Metadata{}, fmt.Errorf("failed to unmarshal metadata: %w", err)
	}

	// Return the extracted metadata.
	return metadata, nil
}
