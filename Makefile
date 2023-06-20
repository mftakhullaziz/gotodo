SHELL=/bin/bash
PACKAGE_NAME=gotodo
BUILD_DIR=./build

# Helper message lists
help:
	@echo " Helper for 'make <target>' where <target> is one of the following:"
	@echo ""
	@echo " build/service  [builds the executable]"
	@echo " build/clean    [cleans the build directory]"
	@echo " run/unittest   [run unit testing default]"
	@echo " run/benchmark  [run unit testing with benchmark]"
	@echo " run/coverage   [run unit testing with coverage]"	
	@echo " run/service    [builds and runs the program]"
	@echo " run/download   [download go package from already project]"
	@echo " clean/package  [remove unused go package from already project]"
	@echo " clean/cache    [cleans the cache]"
	@echo ""
	@echo " Use the example command:"
	@echo ""
	@echo " make build/service"
	@echo " make build/clean"
	@echo " make run/unittest"
	@echo " make run/benchmark"
	@echo " make run/coverage"
	@echo " make run/service"
	@echo " make run/download"
	@echo " make clean/package"
	@echo " make clean/cache"

# Targets
build/service:
	@echo "Build Go Package $(PACKAGE_NAME)"
	go build -o $(BUILD_DIR)/$(PACKAGE_NAME) main.go

run/build:
	@echo "Running from Go Package $(PACKAGE_NAME)"
	./$(BUILD_DIR)/$(PACKAGE_NAME) main.go

run/testing:
	@echo "Running Go Unit Test $(PACKAGE_NAME)"
	go test -v ./config/ ./internal/adapters/handlers/tasks/ ./internal/adapters/handlers/accounts/ ./internal/adapters/handlers/accounts/login/ ./internal/adapters/handlers/accounts/register/ ./internal/persistence/record/

run/grc-testing:
	@echo "Running Go Unit Test $(PACKAGE_NAME)"
	grc go test -v ./config/ ./internal/adapters/handlers/tasks/ ./internal/adapters/handlers/accounts/ ./internal/adapters/handlers/accounts/login/ ./internal/adapters/handlers/accounts/register/ ./internal/persistence/record/

run/benchmark:
	@echo "Running Go Benchmark $(PACKAGE_NAME)"
	grc go test -bench=. ./config/ ./internal/adapters/handlers/tasks/ ./internal/adapters/handlers/accounts/ ./internal/adapters/handlers/accounts/login/ ./internal/adapters/handlers/accounts/register/ ./internal/persistence/record/

run/coverage:
	@echo "Running Go Coverage $(PACKAGE_NAME)"
	grc go test -cover ./config/ ./internal/adapters/handlers/tasks/ ./internal/adapters/handlers/accounts/ ./internal/adapters/handlers/accounts/login/ ./internal/adapters/handlers/accounts/register/ ./internal/persistence/record/

run/service:
	@echo "Running Service Go $(PACKAGE_NAME)"
	go run main.go

build/clean:
	@echo "Clear Build Dir"
	rm -rf $(BUILD_DIR)

clean/cache:
	@echo "Clean Cache and Test Cache"
	go clean -cache -testcache -modcache

run/download:
	@echo "Download Package from Mod"
	go mod download

clean/package:
	@echo "Remove Unused Package"
	go mod tidy
