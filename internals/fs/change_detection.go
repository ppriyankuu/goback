package fs

import (
	"fmt"
	"slices"
)

func DetectChanges(source, previousSource string) ([]string, error) {
	currentFiles, err := TraversalDirectory(source)
	if err != nil {
		return nil, fmt.Errorf("failed to traverse current directory: %w", err)
	}

	previousFiles, err := TraversalDirectory(previousSource)
	if err != nil {
		return nil, fmt.Errorf("failed to traverse previous directory: %w", err)
	}

	var changes []string
	for _, file := range currentFiles {
		if !contains(previousFiles, file) {
			changes = append(changes, file)
		}
	}

	return changes, nil
}

func contains(slice []string, item string) bool {
	return slices.Contains(slice, item)
}
