package storage

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/ppriyankuu/goback/internals/fs"
)

// CreateArchive creates a zip archive of the specified source directory.

// Parameters:
// - destination: The directory where the archive file will be saved.
// - source: The root directory to be archived.
// - incremental: A boolean indicating if the archive should be incremental (currently used).

// Returns:
// - string: The path to the created archive file.
// - error: An error if the archive creation fails.
func CreateArchive(destination, source string, incremental bool) (string, error) {
	// Generate a timestamp for the archive file name.
	timestamp := time.Now().Format("20060102150405")
	archivePath := filepath.Join(destination, fmt.Sprintf("backup_%s.zip", timestamp))

	// Create the archive file.
	file, err := os.Create(archivePath)
	if err != nil {
		return "", fmt.Errorf("failed to create archive file: %w", err)
	}
	defer file.Close()

	// Initialise the zip writer.
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	// Traverse the source directory to get a list of files.
	files, err := fs.TraversalDirectory(source)
	if err != nil {
		return "", fmt.Errorf("failed to traverse directory: %w", err)
	}

	// Iterate over each file in the soruce directory.
	for _, file := range files {
		// Get file information
		info, err := fs.GetFileMetadata(file)
		if err != nil {
			return "", fmt.Errorf("failed to get file info: %w", err)
		}

		// Create a zip header based on the file info.
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return "", fmt.Errorf("failed to get file info header: %w", err)
		}

		// Set the header name to the file's base name.
		header.Name = filepath.Base(file)

		// Create a writer for the zip file.
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return "", fmt.Errorf("failed to create zip header: %w", err)
		}

		// Open the source file.
		src, err := os.Open(file)
		if err != nil {
			return "", fmt.Errorf("failed to open source file: %w", err)
		}

		// Copy the file contents to the zip archive.
		_, err = io.Copy(writer, src)
		src.Close()

		// Check if the copy operation was successful.
		if err != nil {
			return "", fmt.Errorf("failed to copy file to archive: %w", err)
		}
	}

	// Return the path to the created archive file.
	return archivePath, nil
}

// CreateIncrementalArchive creates a zip archive containing only the specified changed files.

// Parameters:
// - destination: The directory where the archive file will be saved.
// - source: The root directory of the files to be archived.
// - changes: A list of file paths that have changed and need to be archived.

// Returns:
// - string: The path to the created archive file.
// - error: An error if the archive creation fails.
func CreateIncrementalArchive(destination, source string, changes []string) (string, error) {
	// Generate a timestamp for the archive file name.
	timestamp := time.Now().Format("20060102150405")
	archivePath := filepath.Join(destination, fmt.Sprintf("incremental_backup_%s.zip", timestamp))

	// Create the archive file
	file, err := os.Create(archivePath)
	if err != nil {
		return "", fmt.Errorf("failed to create archive file: %w", err)
	}
	defer file.Close()

	// Initialise the zip writer
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	// Iterate over each file that has changed.
	for _, filePath := range changes {
		// Get file information
		info, err := fs.GetFileMetadata(filePath)
		if err != nil {
			return "", fmt.Errorf("failed to get file info: %w", err)
		}

		// Create a zip header based on the file info.
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return "", fmt.Errorf("failed to get file info header: %w", err)
		}
		header.Name = filepath.Base(filePath)

		// Create a writer for the zip file
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return "", fmt.Errorf("failed to create zip header: %w", err)
		}

		// Open the source file.
		src, err := os.Open(filePath)
		if err != nil {
			return "", fmt.Errorf("failed to open source file: %w", err)
		}

		// Copy the file contents to the zip archive.
		_, err = io.Copy(writer, src)
		src.Close()

		// Check if the copy operation was successful.
		if err != nil {
			return "", fmt.Errorf("failed to copy file to archive: %w", err)
		}
	}

	// Return the path to the created archive file.
	return archivePath, nil
}

// ExtractArchive extracts the contents of a zip archive to the specified destination directory.

// Parameters:
// - archivePath: The path to the zip archive file.
// - destination: The path to the directory where the archive contents will be extracted.

// Returns:
// - error: An error if the extraction fails at any point.
func ExtractArchive(archivePath, destination string) error {
	// Open the zip archive file
	file, err := os.Open(archivePath)
	if err != nil {
		return fmt.Errorf("failed to open archive file: %w", err)
	}
	defer file.Close()

	// Retrieve file information
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get archive file info: %w", err)
	}

	// Create a new zip reader for the file.
	zipReader, err := zip.NewReader(file, info.Size())
	if err != nil {
		return fmt.Errorf("failed to create zip reader: %w", err)
	}

	// Iterate over each file in the zip archive.
	for _, zf := range zipReader.File {
		// Construct the destination file path
		dstPath := filepath.Join(destination, zf.Name)

		// Create the destination directory if it doesn't exist.
		if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
			return fmt.Errorf("failed to create destination directory: %w", err)
		}

		// Create the destination file
		dstFile, err := os.Create(dstPath)
		if err != nil {
			return fmt.Errorf("failed to create destination file: %w", err)
		}

		// Open the source file inside the zip archive.
		srcFile, err := zf.Open()
		if err != nil {
			dstFile.Close()
			return fmt.Errorf("failed to open source file: %w", err)
		}

		// Copy the contents from the source file to the destination file.
		_, err = io.Copy(dstFile, srcFile)
		srcFile.Close()
		dstFile.Close()

		// Check if the copy operation was successful.
		if err != nil {
			return fmt.Errorf("failed to copy file to destination: %w", err)
		}
	}

	return nil
}

// GetRecentBackup retrieves the most recent backup metadata from the specified destination directory.

// Parameters:
// - destination: The path to the directory where backups are stored.

// Returns:
// - Metadata: The metadata of the most recent backup file.
// - error: An error if no backups are found or if the traversal fails.
func GetRecentBackup(destination string) (Metadata, error) {
	// Traverse the specified directory to get a list of files.
	files, err := fs.TraversalDirectory(destination)
	if err != nil {
		return Metadata{}, fmt.Errorf("failed to traverse directory: %w", err)
	}

	// Initialise a slice to store backup metadata.
	var backups []Metadata

	// Iterate through the list of files.
	for _, file := range files {
		// Check if the file has a .zip extension
		if filepath.Ext(file) == ".zip" {
			// Attempt to retrieve metadata for the backup file.
			metadata, err := GetMetadata(file)
			if err != nil {
				// Skip files that fail to provide metadata.
				continue
			}
			// Append valid metadata to the backups slice.
			backups = append(backups, metadata)
		}
	}

	// Sort the backups slice based on the timestamp, newest last.
	sort.Slice(backups, func(i, j int) bool {
		return backups[i].Time.Before(backups[j].Time)
	})

	// Check if there are any backups found.
	if len(backups) == 0 {
		return Metadata{}, fmt.Errorf("no backups found")
	}

	// Return the most recent backup metadata.
	return backups[len(backups)-1], nil
}
