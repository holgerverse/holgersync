.DEFAULT: help

.PHONY: help
help:
	@echo 'Makefile for "holgersync" development.'
	@echo ''
	@echo 'Usage:'
	@echo '  build-holgersync       - Build the hoglersync application.'
	@echo '  run-holgerdocs         - Run the holgerdocs subcommand on the test path.'

.PHONY: run-holgerdocs
run-holgerdocs:
	go run cmd/main.go holgerdocs --modulepath tests

.PHONY: build-holgersync
build-holgersync:
	go build -o holgersync ./...