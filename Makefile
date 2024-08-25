# Simple Makefile for a Go project
include $(PWD)/.env
PATH := $(USER_HOME)/go/bin:$(PATH)


# Build the application
all: build

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go


# Run the application
run:
	@go run cmd/api/main.go


# Test the application
test:
	@echo "Testing..."
	@go test ./... -v


# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main


# Live Reload
watch:
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


.PHONY: all build run test clean watch
