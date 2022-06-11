.DEFAULT: help

.PHONY: help
help:
	@echo 'Makefile for "holgersync" development.'
	@echo ''
	@echo 'Usage:'
	@echo '  build-holgersync       - Build the hoglersync application.'
	@echo '  run-holgersync         - Run the holgersync subcommand on the test path.'

.PHONY: run-holgersync
run-holgerdocs:
	go run cmd/*.go

.PHONY: build-holgersync
build-holgersync:
	go build -o holgersync ./...