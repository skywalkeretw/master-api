# [Master] API for A Serverless Platform for Deploying, Running, and Integrating Microservices

Welcome to our Go-based API, designed to simplify your interaction with OpenAPI RESTful and AsyncAPI messaging services. This powerful system is deployed to kubernetes and extends the functionality of the container plafrom by enabling you to effortlessly generate adapter code, modify configurations, delete services, and access comprehensive documentation.

## Features

- **RESTful Service Integration:** With our API, you can seamlessly interact with RESTful services, making it easy to perform common tasks like making HTTP requests, handling responses, and more.

- **Messaging Service Integration:** Our API supports messaging services, allowing you to effortlessly connect and communicate with various message queues, making your application more robust and responsive.

- **Generate Adapter Code:** Simplify the integration process by generating adapter code tailored to your specific needs.

- **Configuration Management:** Modify your service configurations on the fly. Our API provides endpoints to adjust settings, ensuring your applications are always aligned with your requirements.

- **Service Deletion:** Need to clean up or reconfigure your services? You can easily delete services using our API, streamlining the management of your resources.

- **Extensive Documentation:** Access comprehensive Swagger documentation, making it simple to understand and navigate the API endpoints and functionality.

- **Platform Agnostic:** Our API leverages Kubernetes under the hood, providing a platform-agnostic experience. You can use this API on any infrastructure that supports Kubernetes, ensuring flexibility and portability.

## Prerequisites

Before you proceed, make sure you have the following prerequisites installed:

- [Go](https://golang.org/dl/) The API is developed in go. 
    ```bash
    brew install go
    ```
- [Docker](https://www.docker.com/get-started) The API is packaged as a OCI container using Docker. Docker is also used as the test environment.
    ```bash
    brew install --cask docker
    ```
- [Kind (Kubernetes in Docker)](https://kind.sigs.k8s.io/) Kind is used to test the api in a local kubernetes cluster.
    ```bash
    brew install kind
    ```
- [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) Kubectl is required to interact with the Kind cluster or any other Kubernetes cluster.
    ```bash
    brew install kubectl
    ```
- [Swag](https://github.com/swaggo/swag) Swag is used to generate the OpenAPI Spec used for documenting the API
    ```bash
    brew tap swaggo/swag && brew install swag
    ```

## Usage

This API is equipped with a Makefile that simplifies various development and deployment tasks. You can use it for building, testing, running, and deploying the API. Additionally, it includes targets for creating and managing a Kubernetes cluster with Kind for local production like testing and generating Swagger documentation.

## Makefile Targets

**Build and Run the Go Application:**

```bash
make run
```

**Build the Go Application:**

```bash
make build
```

**Generate Swagger Documentation:**

```bash
make swagger
```

**Create a Kind Cluster:**

```bash
make create-cluster
```

**Delete a Kind Cluster:**

```bash
make delete-cluster
```

**Build the Docker Image:**

```bash
make docker-build
```

**Deploy the Application to Kubernetes:**

```bash
make make deploy-api
```

**Run Tests:**

```bash
make test
```

**Clean Up:**

```bash
make clean
```

## Configuration

- **APP_NAME:** The name of your Go application.
- **SWAG_DIR:** The directory containing Swagger documentation.
- **KIND_CONFIG:** The Kind cluster configuration file.
- **KIND_CLUSTER_NAME:** The name of your Kind cluster.
- **DOCKER_IMAGE_NAME:** The name of the Docker image for your - application.
- **DOCKER_IMAGE_TAG:** The Docker image tag.
- **KUBE_MANIFESTS:** The path to your Kubernetes manifest files.
- **MAIN_FILE:** The main Go file for your application.

## Getting Started

1. Make sure you have all the prerequisites installed.
1. Use the provided Makefile targets to build, run, and deploy your Go application as needed.
1. Customize the Makefile variables to match your project's configuration.
1. Refer to the Makefile documentation for more details on each target and their usage.

# Brain storming:

use `npm install -g @stoplight/prism-cli` to mock the api based on open api `prism mock ./openapi.yml -p 8080`

1. Generate the Server based on 

https://swagger.io/docs/open-source-tools/swagger-codegen/
https://github.com/swagger-api/swagger-codegen#homebrew
https://www.asyncapi.com/docs/tutorials/getting-started˝
file:///Users/lukeroy/Downloads/OpenAPI+Specification+Zero+to+Master.pdf