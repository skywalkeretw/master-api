# Define your Go application name and source code directory.
APP_NAME := api

# Define the paths for Swag and Kind configuration files.
SWAG_DIR := ./docs
KIND_CONFIG := ./deployment/kind-config.yml

# Define the Kind cluster name.
KIND_CLUSTER_NAME := master

# Define Go related variables.
GO := go
GO_BUILD := $(GO) build
GO_TEST := $(GO) test
GO_RUN := $(GO) run

# Define the main Go file for your application.
MAIN_FILE := main.go

# Define the Swag command to generate API documentation.
SWAG := swag

# Define the Kind command.
KIND := kind

# Define the Docker-related variables.
DOCKER := docker
DOCKER_IMAGE_NAME := api
DOCKER_IMAGE_TAG := v1
DOCKER_BUILD_ARGS := -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

# Define the path to your Kubernetes manifests directory.
API_DEPLOYMENT := deployment/api.yml

# Default target to build and run your Go application.
.PHONY: run
run:
	$(GO_RUN) $(MAIN_FILE)

# Target to build your Go application.
.PHONY: build
build:
	$(GO_BUILD) -o $(APP_NAME) $(MAIN_FILE)

# Target to generate Swagger documentation using Swag.
.PHONY: swagger
swagger:
	$(SWAG) init --parseDependency

# Target to create a Kind cluster with the specified name and configuration.
.PHONY: create-cluster
create-cluster: create-kind-cluster deploy-rabbitmq docker-build-load deploy-api

create-kind-cluster:
	kind create cluster --name master --config deployment/kind-config.yml
# ./deployment/deploy-cluster.sh

# Target to delete the Kind cluster with the specified name.
.PHONY: delete-cluster
delete-cluster:
	$(KIND) delete cluster --name $(KIND_CLUSTER_NAME)

# Target to build the Docker image for your Go application.
.PHONY: docker-build
docker-build:
	$(DOCKER) build $(DOCKER_BUILD_ARGS) .

.PHONY: docker-build-load
docker-build-load: docker-build
	kind load docker-image --name $(KIND_CLUSTER_NAME) $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

# Target to apply Kubernetes manifests using kubectl.
.PHONY: deploy-api
deploy-api:
	kubectl apply -f $(API_DEPLOYMENT)

.PHONY: deploy-rabbitmq
deploy-rabbitmq:
	kubectl apply -f "https://github.com/rabbitmq/cluster-operator/releases/latest/download/cluster-operator.yml"
	kubectl apply -f https://raw.githubusercontent.com/rabbitmq/cluster-operator/main/docs/examples/hello-world/rabbitmq.yaml


# Target to apply Kubernetes manifests using kubectl.
.PHONY: deploy-registry
deploy-registry:
	pwd
	./deployment/deploy-registry.sh ${KIND_CLUSTER_NAME}

# Target to run tests for your Go application.
.PHONY: test
test:
	$(GO_TEST) ./...

# Clean up generated files and artifacts.
.PHONY: clean
clean:
	rm -f $(APP_NAME) $(SWAG_DIR)/docs.go
