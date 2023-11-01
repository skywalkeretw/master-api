// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/app": {
            "get": {
                "description": "List all Kubernetes deployments in the cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deployments"
                ],
                "summary": "List Kubernetes Deployments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/kubernets.DeploymentInfo"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/pods": {
            "get": {
                "description": "List all Kubernetes pods in the cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List Kubernetes Pods",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/kubernets.PodInfo"
                            }
                        }
                    }
                }
            }
        },
        "/healthcheck": {
            "get": {
                "description": "Perform a healthcheck",
                "produces": [
                    "application/json"
                ],
                "summary": "Healthcheck",
                "operationId": "healthcheck",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/liveness": {
            "get": {
                "description": "Check if the service is live",
                "produces": [
                    "application/json"
                ],
                "summary": "Liveness",
                "operationId": "liveness",
                "responses": {
                    "200": {
                        "description": "Live",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/readiness": {
            "get": {
                "description": "Check if the service is ready",
                "produces": [
                    "application/json"
                ],
                "summary": "Readiness",
                "operationId": "readiness",
                "responses": {
                    "200": {
                        "description": "Ready",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "kubernets.DeploymentInfo": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                }
            }
        },
        "kubernets.PodInfo": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Name still to be choosen",
	Description:      "REST API for interacting with the platform",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
