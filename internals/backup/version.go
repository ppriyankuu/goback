package backup

import (
	"fmt"

	"github.com/ppriyankuu/goback/internals/storage"
)

// Version retrieves and returns the version information of the most recent backup.

// Returns:
// - string: The formatted version information.
// - error: An error if fetching recent metadata fails.
func Version() (string, error) {
	// Retrieve the most recent backup metadata.
	metadata, err := storage.GetRecentBackup(".")
	if err != nil {
		return "", fmt.Errorf("failed to get recent metadata: %w", err)
	}

	// Format and return the version.
	return fmt.Sprintf("Backup version: %s (created at: %s)", metadata.Path, metadata.Time), nil
}
