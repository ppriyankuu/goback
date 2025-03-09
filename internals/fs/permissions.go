package fs

import (
	"fmt"
	"os"
	"syscall"
)

func PreservePermissions(source, destination string) error {
	srcInfo, err := os.Stat(source)
	if err != nil {
		return fmt.Errorf("failed to get source file info: %w", err)
	}

	if err := os.Chmod(destination, srcInfo.Mode()); err != nil {
		return fmt.Errorf("failed to set destination file mode: %w", err)
	}
	if err := os.Chtimes(destination, srcInfo.ModTime(), srcInfo.ModTime()); err != nil {
		return fmt.Errorf("failed to set destination file times: %w", err)
	}

	stat, ok := srcInfo.Sys().(*syscall.Stat_t)
	if !ok {
		return fmt.Errorf("failed to retrieve file ownership information")
	}

	if err := os.Lchown(destination, int(stat.Uid), int(stat.Gid)); err != nil {
		return fmt.Errorf("failed to set destination file ownership: %w", err)
	}

	return nil
}
