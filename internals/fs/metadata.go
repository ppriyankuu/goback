package fs

import (
	"fmt"
	"os"
	"syscall"
)

func GetFileMetadata(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

func SetFileMetadata(path string, info os.FileInfo) error {
	if err := os.Chmod(path, info.Mode()); err != nil {
		return fmt.Errorf("failed to set file mod: %w", err)
	}
	if err := os.Chtimes(path, info.ModTime(), info.ModTime()); err != nil {
		return fmt.Errorf("failed to set file times: %w", err)
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return fmt.Errorf("failed to retrieve file ownership information")
	}
	if err := os.Lchown(path, int(stat.Uid), int(stat.Gid)); err != nil {
		return fmt.Errorf("failed to set file ownership: %w", err)
	}

	return nil
}
