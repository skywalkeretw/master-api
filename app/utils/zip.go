package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ZipFolder(folderPath string, zipFilePath string) error {
	// Create a new zip file
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %v", err)
	}
	defer zipFile.Close()

	// Create a new zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Walk through the folder and add files to the zip archive
	err = filepath.Walk(folderPath, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if fileInfo.IsDir() {
			return nil
		}

		// Create a new file header
		fileHeader, err := zip.FileInfoHeader(fileInfo)
		if err != nil {
			return err
		}

		// Set the file name to be relative to the folder
		fileHeader.Name, err = filepath.Rel(folderPath, filePath)
		if err != nil {
			return err
		}

		// Create a new zip entry
		writer, err := zipWriter.CreateHeader(fileHeader)
		if err != nil {
			return err
		}

		// Open and copy the file content to the zip entry
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)
		return err
	})

	if err != nil {
		return fmt.Errorf("failed to add files to zip archive: %v", err)
	}

	fmt.Printf("Folder '%s' successfully zipped to '%s'\n", folderPath, zipFilePath)
	return nil
}
