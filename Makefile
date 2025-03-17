# Configurable variables
BINARY_DIR := bin
GO_FILES := $(shell rg -t go "package main" --files-with-matches 2>/dev/null)
MAIN_DIRS := $(sort $(dir $(GO_FILES)))
PROJECT_NAME := $(shell basename $(shell pwd))
DOCKER_IMAGE_NAME := $(PROJECT_NAME)
DOCKER_IMAGE_TAG := latest
GORELEASER_CONFIG_URL := https://raw.githubusercontent.com/goreleaser/goreleaser/main/www/docs/static/examples/go-binary.yaml
GORELEASER_VERSION := latest
GO_TEST_FLAGS := -v
SKIP_BUILD := false

# Simplified gum log commands
LOG_INFO := gum log --level info
LOG_DEBUG := gum log --level debug
LOG_ERROR := gum log --level error

.PHONY: all deps build test release image

# Default target
all: deps build

# Install dependencies
deps:
	@$(LOG_INFO) "Installing dependencies..."
	@go mod tidy || ($(LOG_ERROR) "Failed to tidy dependencies" && exit 1)
	@go mod download || ($(LOG_ERROR) "Failed to download dependencies" && exit 1)
	@$(LOG_DEBUG) "Dependencies installed successfully"

# Build all main packages
build:
	@if [ "$(SKIP_BUILD)" = "true" ]; then \
		$(LOG_INFO) "Skipping build as SKIP_BUILD=true"; \
	else \
		$(LOG_INFO) "Building binaries..."; \
		mkdir -p $(BINARY_DIR); \
		for dir in $(MAIN_DIRS); do \
			$(LOG_INFO) "Building $$dir"; \
			go build -o $(BINARY_DIR)/$$(basename $$dir) ./$$dir || ($(LOG_ERROR) "Build failed for $$dir" && exit 1); \
		done; \
		$(LOG_DEBUG) "Build completed successfully"; \
	fi

# Run tests with gum logging
test:
	@$(LOG_INFO) "Running tests..."
	@go test $(GO_TEST_FLAGS) ./... || ($(LOG_ERROR) "Tests failed" && exit 1)
	@$(LOG_DEBUG) "Tests completed successfully"

# Build Docker image if Dockerfile exists
image:
	@if [ -f Dockerfile ]; then \
		$(LOG_INFO) "Building Docker image..."; \
		docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) . || ($(LOG_ERROR) "Docker build failed" && exit 1); \
		$(LOG_DEBUG) "Docker image built successfully: $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)"; \
	else \
		$(LOG_ERROR) "Error: Dockerfile not found"; \
		exit 1; \
	fi

# Release using goreleaser
release:
	@$(LOG_INFO) "Preparing release..."
	@curl -s -o .goreleaser.yml $(GORELEASER_CONFIG_URL) || ($(LOG_ERROR) "Failed to download goreleaser config" && exit 1)
	@$(LOG_DEBUG) "Downloaded goreleaser config from $(GORELEASER_CONFIG_URL)"
	@go install github.com/goreleaser/goreleaser@$(GORELEASER_VERSION) || ($(LOG_ERROR) "Failed to install goreleaser" && exit 1)
	@$(LOG_DEBUG) "Installed goreleaser $(GORELEASER_VERSION)"
	@goreleaser release --rm-dist || ($(LOG_ERROR) "Release failed" && exit 1)
	@$(LOG_INFO) "Release completed successfully"
