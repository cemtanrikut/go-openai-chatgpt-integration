# Application name
APP_NAME = go-openai-chatgpt-integration

# Go module
MODULE = github.com/cemtanrikut/$(APP_NAME)

# Default Go commands
GO = go

# Environment file
ENV_FILE = .env

# Default target
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build        - Build the application"
	@echo "  make run          - Run the application"
	@echo "  make test         - Run tests"
	@echo "  make lint         - Run code linting"
	@echo "  make tidy         - Tidy up dependencies"
	@echo "  make clean        - Clean build files"

# Build the application
.PHONY: build
build:
	$(GO) build -o $(APP_NAME) main.go

# Run the application
.PHONY: run
run:
	@if [ -f $(ENV_FILE) ]; then \
		source $(ENV_FILE); \
	fi; \
	$(GO) run main.go

# Run tests
.PHONY: test
test:
	$(GO) test ./... -v

# Lint the code (e.g., using golangci-lint)
.PHONY: lint
lint:
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "golangci-lint is not installed. Please install it: https://golangci-lint.run"; \
		exit 1; \
	fi
	golangci-lint run

# Tidy up dependencies
.PHONY: tidy
tidy:
	$(GO) mod tidy

# Clean build files
.PHONY: clean
clean:
	@rm -f $(APP_NAME)
