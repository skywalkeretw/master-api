package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"unicode/utf8"

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
	randomFolderName := fmt.Sprintf("gen_code_%s", uuid.New())

	// Create the full path for the temporary folder
	tempFolderPath := filepath.Join("generate", "output", randomFolderName)

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
func TransformTitle2FilenamePath(input ...string) string {
	// Remove leading and trailing whitespaces
	trimmedString := strings.TrimSpace(input[len(input)-1])

	// Convert the string to lowercase
	lowercaseString := strings.ToLower(trimmedString)

	// Replace spaces with hyphens
	result := strings.ReplaceAll(lowercaseString, " ", "-")

	if len(result) == 0 || result == "" {
		return ""
	}
	// Append ".json" to the end of the string
	result = result + ".json"
	slice := input[:len(input)-1]
	slice = append(slice, result)
	path := filepath.Join(slice...)
	return path
}

// Int32Ptr functionto get a pointer of type int32
func Int32Ptr(i int) *int32 {
	i32 := int32(i)
	return &i32
}

func IsJSONObject(input string) bool {
	var jsonObject map[string]interface{}
	err := json.Unmarshal([]byte(input), &jsonObject)
	return err == nil
}

func JsonToMap(jsonString string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// truncateString truncates the input string to a maximum of 120 characters followed by ellipsis
func TruncateString(input string) string {
	const maxChars = 120

	if utf8.RuneCountInString(input) <= maxChars {
		return input
	}

	// Truncate to 120 characters and add ellipsis
	runes := []rune(input[:maxChars])
	return string(runes) + "..."
}

func GetEnvSting(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func GetEnvInt(key string, defaultValue int) int {
	envValue, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	value, err := strconv.Atoi(envValue)

	if err != nil {
		return defaultValue
	}

	return value
}

func StringToBool(str string) bool {
	switch strings.ToLower(str) {
	case "true", "t", "yes", "y", "1":
		return true
	case "false", "f", "no", "n", "0":
		return false
	default:
		return false
	}
}

// Custom validation function for allowed languages
func ValidateAllowedLanguages(language string) bool {
	allowedLanguages := []string{"golang", "python", "javascript"}
	for _, allowed := range allowedLanguages {
		if language == allowed {
			return true
		}
	}
	return false
}
