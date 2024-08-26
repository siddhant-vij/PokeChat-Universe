# Include environment variables
include $(PWD)/.env

# Update PATH to include Go binaries
PATH := $(USER_HOME)/go/bin:$(PATH)

# Default target is to build the application
all: build

# Build the application
build:
	@echo "Building the application..."
	@go mod tidy
	@scripts/sqlc.sh
	@go build -o main cmd/api/main.go

# Run the application (including setup and migrations)
run: build
	@echo "Running the application..."
	@scripts/up.sh
	@go run cmd/api/main.go

# Clean the build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -f main
	@go clean

# Live reload with air (install air if not available)
watch:
	@echo "Starting live reload..."
	@scripts/up.sh
	@if which air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                echo "Watching...";\
                air; \
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

# Database migration setup (for convenience)
up:
	@scripts/up.sh

down:
	@scripts/down.sh

.PHONY: all build run clean watch up down
