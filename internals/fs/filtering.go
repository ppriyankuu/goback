package fs

import (
	"fmt"
	"path/filepath"
)

func FilterFiles(files []string, excludePatterns []string) ([]string, error) {
	var filtered []string
	for _, file := range files {
		include := true
		for _, pattern := range excludePatterns {
			match, err := filepath.Match(pattern, file)
			if err != nil {
				return nil, fmt.Errorf("failed to match pattern: %w", err)
			}
			if match {
				include = false
				break
			}
		}
		if include {
			filtered = append(filtered, file)
		}
	}
	return filtered, nil
}
