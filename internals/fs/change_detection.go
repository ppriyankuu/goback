package fs

import (
	"fmt"
	"slices"
)

// DetectChanges compares the contents of two directories and returns a list of files
// that are present in the current directory but missing from the previous directory.

// Parameters:
// - source: The path to the current directory
// - previousSource: The path to the previous directory for comparison

// Returns:
// - []string: A slice of strings containing the names of files that have changed or are new.
// - error: An error if the directory traversal fails for either path.
func DetectChanges(source, previousSource string) ([]string, error) {
	// Traverse the current directory and get a list of all files
	currentFiles, err := TraversalDirectory(source)
	if err != nil {
		return nil, fmt.Errorf("failed to traverse current directory: %w", err)
	}

	// Traverse the previous directory and get a list of all files
	previousFiles, err := TraversalDirectory(previousSource)
	if err != nil {
		return nil, fmt.Errorf("failed to traverse previous directory: %w", err)
	}

	// Initialise a slice to store the names of files that have changed or are new.
	var changes []string
	for _, file := range currentFiles {
		// If the files is not present inside previousFiles, it is considered to be altered
		if !contains(previousFiles, file) {
			// Add the new or changed file to the changes slice.
			changes = append(changes, file)
		}
	}

	// Return the list of changed files
	return changes, nil
}

// contains checks if a specific item exists within a slice of strings

// Parameters:
// - slice: the slice of strings to search within.
// - item: the string item to look for.

// Returns:
// - bool: True if the item is found, otherwise false.
func contains(slice []string, item string) bool {
	// Use slices package to check for item existence.
	return slices.Contains(slice, item)
}
