package openapi

import (
	"fmt"
	"os"

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

	_, err = utils.RunShellCommand("swagger-codegen", "generate", "-i", swaggerSpecPath, "-l", language, "-o", serverCodeTmpDirPath)
	if err != nil {
		return "", err
	}

	zipPath := fmt.Sprintf("/generate/%s.zip", language)
	err = utils.ZipFolder(serverCodeTmpDirPath, zipPath)
	if err != nil {
		return "", err
	}
	err = utils.DeleteFolder(serverCodeTmpDirPath)
	if err != nil {
		fmt.Println("failed to delete folder")
		// return "", err
	}
	err = os.Remove(swaggerSpecPath)
	if err != nil {
		fmt.Println("failed to delete swagger json")
		// return "", err
	}
	return zipPath, nil
}

func GenerateClient(swaggerSpecPath, language string) (string, error) {
	utils.Contains(clients, language)
	clientCodeTmpDirPath, err := utils.GenerateTempFolder()
	if err != nil {
		return "", err
	}

	_, err = utils.RunShellCommand("swagger-codegen", "generate", "-i", swaggerSpecPath, "-l", language, "-o", clientCodeTmpDirPath)
	if err != nil {
		return "", err
	}

	zipPath := fmt.Sprintf("/generate/%s.zip", language)
	err = utils.ZipFolder(clientCodeTmpDirPath, zipPath)
	if err != nil {
		return "", err
	}
	err = utils.DeleteFolder(clientCodeTmpDirPath)
	if err != nil {
		fmt.Println("failed to delete folder")
		// return "", err
	}
	err = os.Remove(swaggerSpecPath)
	if err != nil {
		fmt.Println("failed to delete swagger json")
		// return "", err
	}
	return zipPath, nil
}
