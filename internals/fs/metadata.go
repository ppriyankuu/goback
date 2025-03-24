package fs

import (
	"fmt"
	"os"
	"syscall"
)

// GetFileMetaData retrieves metadata for the specified file path.

// Parameters:
// - path: The file path for which metadata is to be retrieved.

// Returns:
// - os.FileInfo: The file information, including size, mode, and modifaction time.
// - error: An error if the file doesn't exist or metadata retrieval fails.
func GetFileMetadata(path string) (os.FileInfo, error) {
	// Use os.Stat to get file metadata
	return os.Stat(path)
}

// SetFileMetaData sets the file metadata such as permissions, timestamps, and ownerships.

// Parameters:
// - path: The file path for which metadata is to be set.
// - info: the os.FileInfo containing the metadata to reply

// Returns:
// - error: An error if any step in setting metadata fails.
func SetFileMetadata(path string, info os.FileInfo) error {
	// Set the file mode (permissions)
	if err := os.Chmod(path, info.Mode()); err != nil {
		return fmt.Errorf("failed to set file mod: %w", err)
	}

	// Set the file modification and access time
	if err := os.Chtimes(path, info.ModTime(), info.ModTime()); err != nil {
		return fmt.Errorf("failed to set file times: %w", err)
	}

	// Retrieve system-specific file metadata for ownership information.
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return fmt.Errorf("failed to retrieve file ownership information")
	}

	// Set the file onwership (user and group IDs)
	if err := os.Lchown(path, int(stat.Uid), int(stat.Gid)); err != nil {
		return fmt.Errorf("failed to set file ownership: %w", err)
	}

	return nil
}
