package fs

import (
	"fmt"
)

// PreservePermissions copies the file permissions, timestamps and ownership
// from the source file to the destination file.

// Parameters:
// - source: The path of the source file whose metadata need to be copied.
// - destination: The path of the destination file to which metadata is applied.

// Returns:
// - error: An error if any step in preserving metadata fails.
func PreservePermissions(source, destination string) error {
	// Retrive metadata of the source file.
	srcInfo, err := GetFileMetadata(source)
	if err != nil {
		return fmt.Errorf("failed to get source file info: %w", err)
	}

	// Set the metadata on the destination file.
	if err := SetFileMetadata(destination, srcInfo); err != nil {
		return fmt.Errorf("failed to preserve permissions: %w", err)
	}

	return nil
}
