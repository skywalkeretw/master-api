package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

// RunShellCommand runs a shell command and returns the output and error
func RunShellCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	// Check for errors during command execution
	if err != nil {
		return "", fmt.Errorf("failed to run command: %v\n%s", err, stderr.String())
	}

	return stdout.String(), nil
}

// Contains check to see if a string is contained in a slice of string
func Contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// GenerateTempFolder generates a temporary folder with a random name using UUID
// and creates missing folders if specified in the path
func GenerateTempFolder() (string, error) {
	// Generate a random UUID for the folder name
	randomFolderName := fmt.Sprintf("temp_folder_%s", uuid.New())

	// Get the absolute path to the system's temporary directory
	tempDir := os.TempDir()

	// Create the full path for the temporary folder
	tempFolderPath := filepath.Join(tempDir, "output", randomFolderName)

	// Create missing folders if specified in the path
	if err := os.MkdirAll(filepath.Dir(tempFolderPath), os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create missing folders: %v", err)
	}

	// Create the temporary folder
	err := os.Mkdir(tempFolderPath, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create temporary folder: %v", err)
	}

	return tempFolderPath, nil
}

// DeleteFolder deletes a folder and its contents
// Returns nil if the folder does not exist as there is nothing to delete
func DeleteFolder(folderPath string) error {
	// Check if the folder exists
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		// Folder does not exist, return nil
		return nil
	}

	// Attempt to delete the folder and its contents
	err = os.RemoveAll(folderPath)
	if err != nil {
		return fmt.Errorf("failed to delete folder: %v", err)
	}

	return nil
}

// CreateJSONFile creates a JSON file with the content of a struct
// and creates the specified directories if missing
func CreateJSONFile(filename string, data interface{}) error {
	// Get the directory path from the filename
	dir := filepath.Dir(filename)

	// Marshal the struct to JSON format
	jsonContent, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	// Create the specified directories if missing
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directories: %v", err)
	}

	// Write the JSON content to the file
	err = os.WriteFile(filename, jsonContent, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON file: %v", err)
	}

	fmt.Printf("JSON file '%s' created successfully.\n", filename)
	return nil
}

// ConvertToLowerAndReplaceSpaces removes leading and trailing whitespaces, converts a string to lowercase,
// and replaces spaces with hyphens
func TransformTitle2Filename(input string) string {
	// Remove leading and trailing whitespaces
	trimmedString := strings.TrimSpace(input)

	// Convert the string to lowercase
	lowercaseString := strings.ToLower(trimmedString)

	// Replace spaces with hyphens
	result := strings.ReplaceAll(lowercaseString, " ", "-")

	if len(result) == 0 || result == "" {
		return ""
	}
	// Append ".json" to the end of the string
	result = result + ".json"

	return result
}
