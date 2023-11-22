package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZipFolder(t *testing.T) {

	var tests = []struct {
		testFiles []struct {
			name    string
			content string
		}
		zipFileName string
	}{
		{
			testFiles: []struct {
				name    string
				content string
			}{
				{"file1.txt", "Test content 1"},
				{"file2.txt", "Test content 2"},
			},
			zipFileName: "output.zip",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			// Create a temporary directory for testing
			tempDir := t.TempDir()

			// Create some test files in the temporary directory
			for _, testfile := range tt.testFiles {
				filePath := filepath.Join(tempDir, testfile.name)
				if err := os.WriteFile(filePath, []byte(testfile.content), 0644); err != nil {
					t.Fatal(err)
				}
			}

			zipFile := filepath.Join(tempDir, tt.zipFileName)

			err := ZipFolder(tempDir, zipFile)

			// Check if there was an error during zipFolder
			assert.Nil(t, err)

			// Check if the zip file was created
			_, err = os.Stat(zipFile)
			assert.False(t, os.IsNotExist(err), "Zip file should be created")

			// Cleanup: Remove the created zip file
			if err := os.Remove(zipFile); err != nil {
				t.Fatal(err)
			}
		})
	}
}
