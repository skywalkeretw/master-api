package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Parse command line arguments
	filePath := flag.String("file", "", "Path to the source code file")
	flag.Parse()

	// Check if the file path is provided
	if *filePath == "" {
		fmt.Println("Please provide the path to the source code file using the -file flag.")
		return
	}

	// Read the source code file
	annotations, err := detectAnnotations(*filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Println(annotations)
	// Print detected annotations
	fmt.Println("Detected Annotations:")
	for group, data := range annotations {
		fmt.Printf("key: %s", group)
		for key, value := range data {
			fmt.Printf("key: %s -> value: %s", key, value)

		}
	}
}

func detectAnnotations(filePath string) (map[string]map[string]string, error) {
	// Open the source code file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	annotations := make(map[string]map[string]string)
	var currentFunction string

	for lineNum := 1; scanner.Scan(); lineNum++ {
		line := scanner.Text()

		// Check if the line contains an annotation starting with "@"
		if strings.Contains(line, "@") {
			// Find the position of ":"
			colonIndex := strings.Index(line, ":")
			if colonIndex != -1 {
				key := strings.TrimSpace(line[strings.Index(line, "@")+1 : colonIndex])
				value := strings.TrimSpace(line[colonIndex+1:])

				// Check if the annotation is above a function
				if strings.HasPrefix(line, "func ") {
					currentFunction = strings.TrimSpace(strings.TrimPrefix(line, "func "))
					annotations[currentFunction] = make(map[string]string)
				}

				// Group the annotation under the current function
				if currentFunction != "" {
					annotations[currentFunction][key] = value
				}
			}
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return annotations, nil
}
