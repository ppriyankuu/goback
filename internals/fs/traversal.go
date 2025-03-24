package fs

import (
	"fmt"
	"os"
	"path/filepath"
)

// TraversalDirectory walks through the given directory and collects
// all file paths (excluding directories)

// Parameters:
// - dir: The directory path to traverse.

// Returns:
// - []string: A slice containing the paths of all files found.
// - error: An error if the traversal fails.
func TraversalDirectory(dir string) ([]string, error) {
	// Initialise an empty list to store the files.
	var files []string

	// Walk the directory and process each file or folder.
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// If it's not a directory, add the file path to the list
		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to traverse directory: %w", err)
	}

	// Return the list of the file paths
	return files, nil
}
