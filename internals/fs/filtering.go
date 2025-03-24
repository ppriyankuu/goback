package fs

import (
	"fmt"
	"path/filepath"
)

// FilterFiles filters out files that match any of the specified exclude patterns.

// Parameters:
// - files: A slice of the file Paths to be filtered
// - excludePatterns: A slice of patterns to exclude files that match any of them.

// Returns:
// - []string: a slice of filtered file paths that do not match any exclude patterns
// - error: An error if patterns matching fails.
func FilterFiles(files []string, excludePatterns []string) ([]string, error) {
	// Initialise a slice to hold the filtered file paths.
	var filtered []string

	// Iterate through the list of files
	for _, file := range files {
		// Assume the file should be included initially
		include := true
		// Check the file against each exclude pattern
		for _, pattern := range excludePatterns {
			// Attempt to match the current file with exclude pattern
			match, err := filepath.Match(pattern, file)
			if err != nil {
				return nil, fmt.Errorf("failed to match pattern: %w", err)
			}
			// If a match if found, mark the file for exclusion
			if match {
				include = false
				break
			}
		}

		// Add the file to the filtered list if it's not excluded
		if include {
			filtered = append(filtered, file)
		}
	}
	// Return the list of filtered files.
	return filtered, nil
}
