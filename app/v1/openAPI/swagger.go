package openapi

import (
	"fmt"
	"path/filepath"

	"golang.org/x/text/cases"
	lang "golang.org/x/text/language"

	"github.com/skywalkeretw/master-api/app/utils"
)

var servers = []string{"go-server", "python-flask", "nodejs-server"}
var clients = []string{"go", "python", "Node.js"}

// https://github.com/swagger-api/swagger-codegen/wiki/Server-stub-generator-HOWTO#go-server
func GenerateServerStub(swaggerSpecPath, language string) (string, error) {
	utils.Contains(servers, language)
	serverCodeTmpDirPath, err := utils.GenerateTempFolder()
	if err != nil {
		return "", err
	}

	swaggerCodgenCMD, err := utils.RunShellCommand("swagger-codegen", "generate", "-i", swaggerSpecPath, "-l", language, "-o", serverCodeTmpDirPath)
	if err != nil {
		return "", err
	}
	fmt.Println(swaggerCodgenCMD)
	fmt.Println(serverCodeTmpDirPath)
	zipPath := fmt.Sprintf("/generate/%s.zip", language)
	err = utils.ZipFolder(serverCodeTmpDirPath, zipPath)
	if err != nil {
		return "", err
	}
	// err = utils.DeleteFolder(serverCodeTmpDirPath)
	// if err != nil {
	// 	fmt.Println("failed to delete folder")
	// 	// return "", err
	// }
	// err = os.Remove(swaggerSpecPath)
	// if err != nil {
	// 	fmt.Println("failed to delete swagger json")
	// 	// return "", err
	// }
	return zipPath, nil
}

func GenerateClient(swaggerSpecPath, name, language string) (string, error) {
	utils.Contains(clients, language)
	clientCodeTmpDirPath, err := utils.GenerateTempFolder()
	if err != nil {
		return "", err
	}

	_, err = utils.RunShellCommand("swagger-codegen", "generate", "-i", swaggerSpecPath, "-l", language, "-o", clientCodeTmpDirPath)
	if err != nil {
		return "", err
	}

	// create file to apstract the swagger package
	switch language {
	case "go":
		err = utils.CopyFileToFolder("/templates/go/functioncall.go", clientCodeTmpDirPath)
		if err != nil {
			return "", err
		}
		err = utils.ReplacePlaceholder(filepath.Join(clientCodeTmpDirPath, "functioncall.go"), "{{FUNCTION_NAME}}", cases.Title(lang.English, cases.Compact).String(name))
		if err != nil {
			return "", err
		}
		err = utils.ReplacePlaceholder(filepath.Join(clientCodeTmpDirPath, "functioncall.go"), "{{FUNCTION_RETURN}}", "string")
		if err != nil {
			return "", err
		}
	}

	zipPath := fmt.Sprintf("/generate/%s-%s.zip", name, language)
	err = utils.ZipFolder(clientCodeTmpDirPath, zipPath)
	if err != nil {
		return "", err
	}
	// err = utils.DeleteFolder(clientCodeTmpDirPath)
	// if err != nil {
	// 	fmt.Println("failed to delete folder")
	// 	// return "", err
	// }
	// err = os.Remove(swaggerSpecPath)
	// if err != nil {
	// 	fmt.Println("failed to delete swagger json")
	// 	// return "", err
	// }
	return zipPath, nil
}

// getStringFromInterface checks if the interface contains a string of a valid OpenAPI type and returns it
func isValidOpenAPIType(strValue string) (string, error) {

	if utils.IsJSONObject(strValue) {
		strValue = "object"
	}
	// Check if the string is a valid OpenAPI type
	openAPITypes := map[string]bool{
		"string":  true,
		"number":  true,
		"integer": true,
		"boolean": true,
		"array":   true,
		"object":  true,
	}

	_, valid := openAPITypes[strValue]
	if !valid {
		return "", fmt.Errorf("'%s' is not a valid OpenAPI type", strValue)
	}

	return strValue, nil
}
