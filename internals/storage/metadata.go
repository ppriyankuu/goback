package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Metadata holds the details of a backup operation
type Metadata struct {
	Source      string    `json:"source"`
	Destination string    `json:"destination"`
	Path        string    `json:"path"`
	Time        time.Time `json:"time"`
}

// StoreMetadata saves the metadata to a JSON file.

// Parameters:
// - metadata: The metadata information to store.

// Returns:
// - error: An error if marshaling or writing fails.
func StoreMetadata(metatdata Metadata) error {
	// Construct the metadata file path
	metadataPath := filepath.Join(metatdata.Destination, "metadata.json")

	// Marshal metadata to JSON
	data, err := json.Marshal(metatdata)
	if err != nil {
		return fmt.Errorf("failed to marshal metadata: %w", err)
	}

	// Write JSON data to the metadata file.
	if err := os.WriteFile(metadataPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write metadata file: %w", err)
	}

	return nil
}

// GetRecentMetadata retrieves the most recent metadata from a JSON file.

// Parameters:
// - destination: The directory containing the metadata file.

// Returns:
// - Metadata: The metadata information from the file.
// - error: An error if reading or unmarshalling fails.
func GetRecentMetadata(destination string) (Metadata, error) {
	// Construct the metadata file path
	metadataPath := filepath.Join(destination, "metadata.json")

	// Read the metadata file
	data, err := os.ReadFile(metadataPath)
	if err != nil {
		return Metadata{}, fmt.Errorf("failed to read metadata file: %w", err)
	}

	var metadata Metadata

	// Unmarshal JSON data into the Metadata struct.
	if err := json.Unmarshal(data, &metadata); err != nil {
		return Metadata{}, fmt.Errorf("failed to unmarshal metadata: %w", err)
	}

	// Return the metadata on successful retrieval.
	return metadata, nil
}
