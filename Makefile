# Makefile for 5G Core Network Simulator

# Variables
SHELL := /bin/bash
SERVICES := amf smf upf pcf udm ausf nrf nssf
BUILD_DIR := bin
DOCKER_REGISTRY ?=
DOCKER_TAG ?= latest
GO_BUILD_FLAGS := -ldflags="-s -w"
KUBECTL_CONTEXT ?= minikube

.PHONY: all build clean test docker-build docker-push deploy-local deploy-prod run-local help

# Default target
all: build

# Show help information
help:
	@echo "5G Core Network Simulator - Available targets:"
	@echo "  all          : Build all services (default)"
	@echo "  build        : Build all service binaries"
	@echo "  clean        : Remove build artifacts"
	@echo "  test         : Run all tests"
	@echo "  docker-build : Build Docker images for all services"
	@echo "  docker-push  : Push Docker images to registry"
	@echo "  deploy-local : Deploy to local Kubernetes cluster"
	@echo "  deploy-prod  : Deploy to production Kubernetes cluster"
	@echo "  run-local    : Run all services locally using Docker Compose"
	@echo "  run-<service>: Run a specific service locally (e.g., run-nrf)"
	@echo ""
	@echo "Variables:"
	@echo "  DOCKER_REGISTRY : Docker registry prefix (default: none)"
	@echo "  DOCKER_TAG      : Docker image tag (default: latest)"
	@echo "  KUBECTL_CONTEXT : Kubernetes context (default: minikube)"

# Create build directory
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# Build all services
build: $(BUILD_DIR)
	@echo "Building services..."
	@for service in $(SERVICES); do \
		echo "Building $$service..." ; \
		go build $(GO_BUILD_FLAGS) -o $(BUILD_DIR)/$$service cmd/$$service/main.go ; \
	done
	@echo "Build complete."