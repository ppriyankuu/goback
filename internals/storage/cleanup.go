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

func CleanupOldBackups(destination string, retentionDays int) error {
	files, err := fs.TraversalDirectory(destination)
	if err != nil {
		return fmt.Errorf("failed to traverse directory: %w", err)
	}

	var backups []Metadata
	for _, file := range files {
		if filepath.Ext(file) == ".zip" {
			metadata, err := GetMetadata(file)
			if err != nil {
				continue
			}
			backups = append(backups, metadata)
		}
	}

	sort.Slice(backups, func(i, j int) bool {
		return backups[i].Time.Before(backups[j].Time)
	})

	now := time.Now()
	for _, backup := range backups {
		if now.Sub(backup.Time).Hours() > float64(retentionDays)*24 {
			if err := os.Remove(backup.Path); err != nil {
				return fmt.Errorf("failed to remove old backup: %w", err)
			}
		}
	}

	return nil
}

func GetMetadata(archivePath string) (Metadata, error) {
	metadataPath := filepath.Join(filepath.Dir(archivePath), "metadata.json")
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
