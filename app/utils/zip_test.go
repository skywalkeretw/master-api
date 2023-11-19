package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZipFolder(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Create some test files in the temporary directory
	testFile1 := filepath.Join(tempDir, "file1.txt")
	testFile2 := filepath.Join(tempDir, "file2.txt")

	if err := os.WriteFile(testFile1, []byte("Test content 1"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(testFile2, []byte("Test content 2"), 0644); err != nil {
		t.Fatal(err)
	}

	// Define the expected zip file path
	expectedZipFile := filepath.Join(tempDir, "output.zip")

	var tests = []struct {
		sourceFolder string
		zipFile      string
	}{
		{tempDir, expectedZipFile},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			err := ZipFolder(tt.sourceFolder, tt.zipFile)

			// Check if there was an error during zipFolder
			assert.Nil(t, err)

			// Check if the zip file was created
			_, err = os.Stat(tt.zipFile)
			assert.False(t, os.IsNotExist(err), "Zip file should be created")

			// Cleanup: Remove the created zip file
			if err := os.Remove(tt.zipFile); err != nil {
				t.Fatal(err)
			}
		})
	}
}
