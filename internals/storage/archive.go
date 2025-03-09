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

func CreateArchive(destination, source string, incremental bool) (string, error) {
	timestamp := time.Now().Format("20060102150405")
	archivePath := filepath.Join(destination, fmt.Sprintf("backup_%s.zip", timestamp))

	file, err := os.Create(archivePath)
	if err != nil {
		return "", fmt.Errorf("failed to create archive file: %w", err)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	files, err := fs.TraversalDirectory(source)
	if err != nil {
		return "", fmt.Errorf("failed to traverse directory: %w", err)
	}

	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			return "", fmt.Errorf("failed to get file info: %w", err)
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return "", fmt.Errorf("failed to get file info header: %w", err)
		}

		header.Name = filepath.Base(file)
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return "", fmt.Errorf("failed to create zip header: %w", err)
		}

		src, err := os.Open(file)
		if err != nil {
			return "", fmt.Errorf("failed to open source file: %w", err)
		}

		_, err = io.Copy(writer, src)
		src.Close()

		if err != nil {
			return "", fmt.Errorf("failed to copy file to archive: %w", err)
		}
	}

	return archivePath, nil
}

func CreateIncrementalArchive(destination, source string, changes []string) (string, error) {
	timestamp := time.Now().Format("20060102150405")
	archivePath := filepath.Join(destination, fmt.Sprintf("incremental_backup_%s.zip", timestamp))

	file, err := os.Create(archivePath)
	if err != nil {
		return "", fmt.Errorf("failed to create archive file: %w", err)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	for _, filePath := range changes {
		info, err := os.Stat(filePath)
		if err != nil {
			return "", fmt.Errorf("failed to get file info: %w", err)
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return "", fmt.Errorf("failed to get file info header: %w", err)
		}
		header.Name = filepath.Base(filePath)

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return "", fmt.Errorf("failed to create zip header: %w", err)
		}

		src, err := os.Open(filePath)
		if err != nil {
			return "", fmt.Errorf("failed to open source file: %w", err)
		}

		_, err = io.Copy(writer, src)
		src.Close()

		if err != nil {
			return "", fmt.Errorf("failed to copy file to archive: %w", err)
		}
	}

	return archivePath, nil
}

func ExtractArchive(archivePath, destination string) error {
	file, err := os.Open(archivePath)
	if err != nil {
		return fmt.Errorf("failed to open archive file: %w", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get archive file info: %w", err)
	}

	zipReader, err := zip.NewReader(file, info.Size())
	if err != nil {
		return fmt.Errorf("failed to create zip reader: %w", err)
	}

	for _, zf := range zipReader.File {
		dstPath := filepath.Join(destination, zf.Name)

		if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
			return fmt.Errorf("failed to create destination directory: %w", err)
		}

		dstFile, err := os.Create(dstPath)
		if err != nil {
			return fmt.Errorf("failed to create destination file: %w", err)
		}

		srcFile, err := zf.Open()
		if err != nil {
			dstFile.Close()
			return fmt.Errorf("failed to open source file: %w", err)
		}

		_, err = io.Copy(dstFile, srcFile)
		srcFile.Close()
		dstFile.Close()

		if err != nil {
			return fmt.Errorf("failed to copy file to destination: %w", err)
		}
	}

	return nil
}

func GetRecentBackup(destination string) (Metadata, error) {
	files, err := fs.TraversalDirectory(destination)
	if err != nil {
		return Metadata{}, fmt.Errorf("failed to traverse directory: %w", err)
	}

	var backups []Metadata
	for _, file := range files {
		if filepath.Ext(file) == ".zip" {
			metadata, err := GetMetadata(file)
			if err != nil {
				continue
			}
			backups = append(backups, metadata)
		}
	}

	sort.Slice(backups, func(i, j int) bool {
		return backups[i].Time.Before(backups[j].Time)
	})

	if len(backups) == 0 {
		return Metadata{}, fmt.Errorf("no backups found")
	}

	return backups[len(backups)-1], nil
}
