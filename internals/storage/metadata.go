package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Metadata struct {
	Source      string    `json:"source"`
	Destination string    `json:"destination"`
	Path        string    `json:"path"`
	Time        time.Time `json:"time"`
}

func StoreMetadata(metatdata Metadata) error {
	metadataPath := filepath.Join(metatdata.Destination, "metadata.json")
	data, err := json.Marshal(metatdata)
	if err != nil {
		return fmt.Errorf("failed to marshal metadata: %w", err)
	}

	if err := os.WriteFile(metadataPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write metadata file: %w", err)
	}

	return nil
}

func GetRecentMetadata(destination string) (Metadata, error) {
	metadataPath := filepath.Join(destination, "metadata.json")
	data, err := os.ReadFile(metadataPath)
	if err != nil {
		return Metadata{}, fmt.Errorf("failed to read metadata file: %w", err)
	}

	var metadata Metadata
	if err := json.Unmarshal(data, &metadata); err != nil {
		return Metadata{}, fmt.Errorf("failed to unmarshal metadata: %w", err)
	}

	return metadata, nil
}
