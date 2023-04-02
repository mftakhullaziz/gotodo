SHELL=/bin/bash
PACKAGE_NAME=gotodo
BUILD_DIR=./build

# Helper message lists
help:
	@echo "How to use makefile 'make <target>' where <target> is one of the following:"
	@echo "- build/service  [builds the executable]"
	@echo "- run/test       [runs unit tests]"
	@echo "- run/service    [builds and runs the program]"
	@echo "- clean          [cleans the build directory]"
	@echo "- clean/test     [cleans the cache tests]"
	@echo ""
	@echo "How to use command :"
	@echo "- make build/service"
	@echo "- make run/test"
	@echo "- make run/service"
	@echo "- make clean"
	@echo "- make clean/test"

# Targets
build/service:
	@echo "Building $(PACKAGE_NAME) ..."
	go build -o $(BUILD_DIR)/$(PACKAGE_NAME) main.go

run/build:
	@echo "Running from building $(PACKAGE_NAME) ..."
	./$(BUILD_DIR)/$(PACKAGE_NAME) main.go

run/test:
	@echo "Running unit tests for $(PACKAGE_NAME) ..."
	go test -v ./tests/...

run/service:
	@echo "Building and running $(PACKAGE_NAME) ..."
	go run main.go

clean:
	@echo "Cleaning build directory..."
	rm -rf $(BUILD_DIR)

clean/test:
	echo "Cleaning test cache"
	go clean -cache -testcache -modcache

