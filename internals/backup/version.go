package backup

import (
	"fmt"

	"github.com/ppriyankuu/goback/internals/storage"
)

func Version() (string, error) {
	metadata, err := storage.GetRecentBackup(".")
	if err != nil {
		return "", fmt.Errorf("failed to get recent metadata: %w", err)
	}

	return fmt.Sprintf("Backup version: %s (created at: %s)", metadata.Path, metadata.Time), nil
}
