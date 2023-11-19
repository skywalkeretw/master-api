package openapi

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/skywalkeretw/master-api/app/utils"
)

var outputPath string
var servers = []string{"go-server", "python-flask", "nodejs-server"}
var clients = []string{"go", "python", "Node.js"}

// https://github.com/swagger-api/swagger-codegen/wiki/Server-stub-generator-HOWTO#go-server
func GenerateServerStub(swaggerSpec, language string) {
	utils.Contains(servers, language)
	cmd := exec.Command("swagger-codegen", "generate", "-i", swaggerSpec, "-l", language, "-o", outputPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GenerateClient(swaggerSpec, language string) {
	utils.Contains(clients, language)
	cmd := exec.Command("swagger-codegen", "generate", "-i", swaggerSpec, "-l", language, "-o", outputPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
