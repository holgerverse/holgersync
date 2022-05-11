.DEFAULT: help

.PHONY: help
help:
	@echo 'Makefile for "holgersync" development.'
	@echo ''
	@echo 'Usage:'
	@echo '  build-holgersync       - Build the hoglersync application.'
	@echo '  run-holgerdocs         - Run the holgerdocs subcommand on the test path.'
	@echo '  act                    - Run all GitHub actions workflows locally. (Requires Act installed)'

.PHONY: run-holgerdocs
run-holgerdocs:
	go run cmd/*.go holgerdocs terraform --module-path tests

.PHONY: build-holgersync
build-holgersync:
	go build -o holgersync ./...

.PHONY: act
act:
	act --workflows ".github/workflows/"