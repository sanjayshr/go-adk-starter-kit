.PHONY: build run clean test lint help

# Variables
BINARY_NAME=go-adk-starter-kit
CMD_PATH=./cmd/go-adk-starter-kit
BUILD_DIR=./bin

# Default target
help:
	@echo "Available targets:"
	@echo "  build       - Build the application"
	@echo "  run         - Run the application"
	@echo "  clean       - Remove build artifacts"
	@echo "  test        - Run tests"
	@echo "  lint        - Run linter"
	@echo "  help        - Show this help message"

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_PATH)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Run the application
run:
	@go run $(CMD_PATH)/main.go -log-level info -agent-logger=false

# Run with debug logging
run-debug:
	@go run $(CMD_PATH)/main.go -log-level debug

# Run with agent logger
run-with-agent-logger:
	@go run $(CMD_PATH)/main.go -log-level info -agent-logger=true

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

