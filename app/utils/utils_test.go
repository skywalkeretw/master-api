package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunShellCommand(t *testing.T) {
	var tests = []struct {
		testName    string
		command     string
		args        []string
		expectedOut string
		expectedErr error
	}{
		{"EchoTest", "echo", []string{"Hello, World!"}, "Hello, World!\n", nil},
		{"LsNonexistentDir", "ls", []string{"nonexistent_directory"}, "", fmt.Errorf("failed to run command: exit status 1\nls: nonexistent_directory: No such file or directory\n")},
		{"LsCurrentDir", "ls", []string{"."}, "path\nutils.go\nutils_test.go\nzip.go\nzip_test.go\n", nil},
		{"InvalidCommand", "invalidcommand", nil, "", fmt.Errorf("")},
		{"CurlTest", "curl", []string{"-s", "https://www.example.com"}, "<!doctype html>\n<html>\n<head>\n    <title>Example Domain</title>\n\n    <meta charset=\"utf-8\" />\n    <meta http-equiv=\"Content-type\" content=\"text/html; charset=utf-8\" />\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\" />\n    <style type=\"text/css\">\n    body {\n        background-color: #f0f0f2;\n        margin: 0;\n        padding: 0;\n        font-family: -apple-system, system-ui, BlinkMacSystemFont, \"Segoe UI\", \"Open Sans\", \"Helvetica Neue\", Helvetica, Arial, sans-serif;\n        \n    }\n    div {\n        width: 600px;\n        margin: 5em auto;\n        padding: 2em;\n        background-color: #fdfdff;\n        border-radius: 0.5em;\n        box-shadow: 2px 3px 7px 2px rgba(0,0,0,0.02);\n    }\n    a:link, a:visited {\n        color: #38488f;\n        text-decoration: none;\n    }\n    @media (max-width: 700px) {\n        div {\n            margin: 0 auto;\n            width: auto;\n        }\n    }\n    </style>    \n</head>\n\n<body>\n<div>\n    <h1>Example Domain</h1>\n    <p>This domain is for use in illustrative examples in documents. You may use this\n    domain in literature without prior coordination or asking for permission.</p>\n    <p><a href=\"https://www.iana.org/domains/example\">More information...</a></p>\n</div>\n</body>\n</html>\n", nil},
		{"CatNonexistentFile", "cat", []string{"nonexistent_file.txt"}, "", fmt.Errorf("failed to run command: exit status 1\ncat: nonexistent_file.txt: No such file or directory\n")},
		{"EchoEmpty", "echo", nil, "\n", nil},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			output, err := RunShellCommand(tt.command, tt.args...)

			// Check if the command output matches the expected output
			assert.Equal(t, tt.expectedOut, output)

			// Check if the error message contains the expected error string
			assert.Contains(t, fmt.Sprintf("%v", err), fmt.Sprintf("%v", tt.expectedErr))
		})
	}
}

func TestContains(t *testing.T) {
	var tests = []struct {
		slice    []string
		str      string
		expected bool
	}{
		{[]string{"apple", "banana", "orange", "grape"}, "orange", true},
		{[]string{"apple", "banana", "orange", "grape"}, "kiwi", false},
		{[]string{"one", "two", "three"}, "two", true},
		{[]string{"red", "green", "blue"}, "purple", false},
		{[]string{"a", "b", "c"}, "b", true},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := Contains(tt.slice, tt.str)

			// Check if the result matches the expected value
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGenerateTempFolder(t *testing.T) {
	var tests = []struct {
		testName string
	}{
		{"Test1"},
		{"Test2"},
		{"Test3"},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			// Call the function to generate a temporary folder
			tempFolderPath, err := GenerateTempFolder()

			// Check if there was an error during generateTempFolder
			assert.Nil(t, err)

			// Check if the temporary folder was created
			_, err = os.Stat(tempFolderPath)
			assert.False(t, os.IsNotExist(err), "Temporary folder should be created")

			// Cleanup: Remove the created temporary folder
			if err := os.RemoveAll(tempFolderPath); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestDeleteFolder(t *testing.T) {
	var tests = []struct {
		testName    string
		folderPath  string
		expectError bool
	}{
		{"Test1", "path/to/test/folder1", false},     // Existing folder
		{"Test2", "path/to/test/nonexistent", false}, // Non-existent folder (expecting no error)
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			// Create a test folder if it's not the non-existent case
			if tt.testName != "Test2" {
				err := os.MkdirAll(tt.folderPath, os.ModePerm)
				assert.Nil(t, err, "Error creating test folder")
			}

			// Call the function to delete the folder
			err := DeleteFolder(tt.folderPath)

			// Check if the function behavior matches expectations
			if tt.expectError {
				assert.NotNil(t, err, "Expected an error but got none")
			} else {
				assert.Nil(t, err, "Unexpected error")
			}

			// If the folder should not exist, verify that it was deleted
			if !tt.expectError {
				_, err := os.Stat(tt.folderPath)
				assert.True(t, os.IsNotExist(err), "Folder should not exist after deletion")
			} else {
				// If an error was expected, verify that the folder still exists
				assert.False(t, os.IsNotExist(err), "Folder should still exist")
			}
		})
	}
}

func TestCreateJSONFile(t *testing.T) {
	type ExampleData struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	var tests = []struct {
		testName       string
		filename       string
		data           interface{}
		expectedErrMsg string
	}{
		{"Test1", "testfile1.json", ExampleData{Name: "Example1", Value: 42}, ""},
		{"Test2", "testfile2.json", ExampleData{Name: "Example2", Value: 99}, ""},
		{"Test3", "invalidpath/invalidfile.json", ExampleData{Name: "Example3", Value: 123}, "failed to write JSON file:"},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			err := CreateJSONFile(tt.filename, tt.data)

			// Check if an error occurred and if the error message contains the expected string
			if tt.expectedErrMsg != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErrMsg)
			} else {
				// If no error is expected, assert that the error is nil
				assert.NoError(t, err)

				// Check if the file was created successfully
				_, statErr := os.Stat(tt.filename)
				assert.NoError(t, statErr)

				// Remove the created file after the test
				defer os.Remove(tt.filename)
			}
		})
	}
}

func TestTransformTitle2Filename(t *testing.T) {
	var tests = []struct {
		testName       string
		input          string
		expectedOutput string
	}{
		{"Test1", " Hello World ", "hello-world.json"},
		{"Test2", " This is a Test ", "this-is-a-test.json"},
		{"Test3", "  Multiple  Spaces  ", "multiple--spaces.json"},
		{"Test4", "", ""},
		{"Test5", "   OnlySpaces  ", "onlyspaces.json"},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			output := TransformTitle2Filename(tt.input)

			// Check if the output matches the expected output
			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}
