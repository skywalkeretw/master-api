# Define your Go application name and source code directory.
APP_NAME := api

# Define the paths for Swag and Kind configuration files.
SWAG_DIR := ./docs
KIND_CONFIG := ./deployment/kind-config.yml

# Define the Kind cluster name.
KIND_CLUSTER_NAME := my-kind-cluster

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
DOCKER_IMAGE_NAME := myapp-image
DOCKER_IMAGE_TAG := latest
DOCKER_BUILD_ARGS := -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

# Define the path to your Kubernetes manifests directory.
KUBE_MANIFESTS := deployment/api.yml

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
	$(SWAG) init -g $(MAIN_FILE) -d $(SWAG_DIR)

# Target to create a Kind cluster with the specified name and configuration.
.PHONY: create-cluster
create-cluster:
	$(KIND) create cluster --name $(KIND_CLUSTER_NAME) --config $(KIND_CONFIG)

# Target to delete the Kind cluster with the specified name.
.PHONY: delete-cluster
delete-cluster:
	$(KIND) delete cluster --name $(KIND_CLUSTER_NAME)

# Target to build the Docker image for your Go application.
.PHONY: docker-build
docker-build: build
	$(DOCKER) build $(DOCKER_BUILD_ARGS) .

# Target to apply Kubernetes manifests using kubectl.
.PHONY: kubectl-apply
kubectl-apply:
	kubectl apply -f $(KUBE_MANIFESTS)

# Target to delete resources in the cluster.
.PHONY: kubectl-delete
kubectl-delete:
	kubectl delete -f $(KUBE_MANIFESTS)

# Target to run tests for your Go application.
.PHONY: test
test:
	$(GO_TEST) ./...

# Default target when you run "make" without specifying a target.
.PHONY: all
all: build run

# Clean up generated files and artifacts.
.PHONY: clean
clean:
	rm -f $(APP_NAME) $(SWAG_DIR)/docs.go
