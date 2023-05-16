SHELL=/bin/bash
PACKAGE_NAME=gotodo
BUILD_DIR=./build

# Helper message lists
help:
	@echo "how to use makefile 'make <target>' where <target> is one of the following:"
	@echo " build/service  [builds the executable]"
	@echo " build/clean    [cleans the build directory]"
	@echo " run/unittest   [runs unit tests]"
	@echo " run/benchmark  [runs unit tests with bench]"
	@echo " run/service    [builds and runs the program]"
	@echo " run/download   [download go package from already project]"
	@echo " clean/package  [remove unused go package from already project]"
	@echo " clean/cache    [cleans the cache]"
	@echo ""
	@echo "use the command:"
	@echo " make build/service"
	@echo " make build/clean"
	@echo " make run/unittest"
	@echo " make run/benchmark"
	@echo " make run/service"
	@echo " make run/download"
	@echo " make clean/package"
	@echo " make clean/cache"

# Targets
build/service:
	@echo "build package $(PACKAGE_NAME)"
	go build -o $(BUILD_DIR)/$(PACKAGE_NAME) main.go

run/build:
	@echo "running from build $(PACKAGE_NAME)"
	./$(BUILD_DIR)/$(PACKAGE_NAME) main.go

run/unittest:
	@echo "running unit tests for $(PACKAGE_NAME)"
	go test -v ./config/... ./internal/...

run/benchmark:
	@echo "running unit tests with benchmark for $(PACKAGE_NAME)"
	go test -bench=. ./config/... ./internal/...

run/service:
	@echo "building and running $(PACKAGE_NAME)"
	go run main.go

build/clean:
	@echo "cleaning build directory"
	rm -rf $(BUILD_DIR)

clean/cache:
	@echo "cleaning cache"
	go clean -cache -testcache -modcache

run/download:
	@echo "download all package from go mod"
	go mod download

clean/package:
	@echo "remove unused package"
	go mod tidy
