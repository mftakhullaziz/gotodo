SHELL=/bin/bash
HOME := $(shell pwd)

run/test:
	go test -v ./tests/...