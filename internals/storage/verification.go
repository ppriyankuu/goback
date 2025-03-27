package storage

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// VerifyBackup verifies the integrity and contents of a zip archive.

// Parameters:
// - archivePath: The file path to the zip archive that needs to be verified.

// Returns:
// - error: An error if verification or extraction fails.
func VerifyBackup(archivePath string) error {
	// Open the zip archive file
	file, err := os.Open(archivePath)
	if err != nil {
		return fmt.Errorf("failed to open archive file: %w", err)
	}

	// Ensure the file is closed when the function exits.
	defer file.Close()

	// Get file information for the zip archive.
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get archive file info: %w", err)
	}

	// Create a new zip reader to read archive's contents.
	zipReader, err := zip.NewReader(file, info.Size())
	if err != nil {
		return fmt.Errorf("failed to create zip reader: %w", err)
	}

	// Iterate over each file inside the zip archive
	for _, file := range zipReader.File {
		// Open the current file inside the archive
		srcFile, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open source file: %w", err)
		}

		// Ensure the file is closed when done.
		defer srcFile.Close()

		// Define the destination path where the file will be extracted.
		destPath := filepath.Join("/tmp", file.Name)

		// Create the destination directory if it does not exist.
		if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
			return fmt.Errorf("failed to create destination directory: %w", err)
		}

		// Create the destination file where the content will be copied.
		destFile, err := os.Create(destPath)
		if err != nil {
			return fmt.Errorf("failed to create destination file: %w", err)
		}

		// Ensure the file is closed when done.
		defer destFile.Close()

		// Copy the contents of the current file to the destination file.
		if _, err := io.Copy(destFile, srcFile); err != nil {
			return fmt.Errorf("failed to copy file to destination: %w", err)
		}
	}

	return nil
}
