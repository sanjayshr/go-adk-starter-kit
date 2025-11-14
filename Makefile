.PHONY: build run clean test lint help

# Variables
BINARY_NAME=go-adk-starter-kit
WEB_BINARY_NAME=go-adk-web
CMD_PATH=./cmd/go-adk-starter-kit
WEB_CMD_PATH=./cmd/go-adk-web
BUILD_DIR=./bin

# Default target
help:
	@echo "Available targets:"
	@echo "  CLI Application:"
	@echo "    build           - Build the CLI application"
	@echo "    run             - Run the CLI application"
	@echo "    run-debug       - Run with debug logging"
	@echo "    run-no-logger   - Run without agent logger"
	@echo ""
	@echo "  Web Application:"
	@echo "    build-web       - Build the web launcher application"
	@echo "    run-web         - Run web API and UI together"
	@echo "    run-web-api     - Run only web API server"
	@echo "    run-webui       - Run only web UI"
	@echo ""
	@echo "  General:"
	@echo "    build-all       - Build both CLI and web applications"
	@echo "    clean           - Remove build artifacts"
	@echo "    test            - Run tests"
	@echo "    lint            - Run linter"
	@echo "    deps            - Install dependencies"
	@echo "    fmt             - Format code"
	@echo "    help            - Show this help message"

# Build the CLI application
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_PATH)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Build the web launcher application
build-web:
	@echo "Building $(WEB_BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(WEB_BINARY_NAME) $(WEB_CMD_PATH)
	@echo "Build complete: $(BUILD_DIR)/$(WEB_BINARY_NAME)"

# Build both applications
build-all: build build-web

# Run the application
run:
	@go run $(CMD_PATH)/main.go -log-level info -agent-logger=false

# Run with debug logging
run-debug:
	@go run $(CMD_PATH)/main.go -log-level debug

# Run with agent logger
run-with-agent-logger:
	@go run $(CMD_PATH)/main.go -log-level info -agent-logger=true

# Run without agent logger
run-no-logger:
	@go run $(CMD_PATH)/main.go -agent-logger=false

# Run web API server
run-web-api:
	@echo "Starting web API server..."
	@go run $(WEB_CMD_PATH)/main.go web api

# Run web UI
run-webui:
	@echo "Starting web UI..."
	@go run $(WEB_CMD_PATH)/main.go web webui

# Run both web API and UI
run-web:
	@echo "Starting web API and UI..."
	@go run $(WEB_CMD_PATH)/main.go web api webui

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@go clean
	@echo "Clean complete"

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run linter (requires golangci-lint)
lint:
	@echo "Running linter..."
	@golangci-lint run ./...

# Install dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies installed"

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Format complete"

