.DEFAULT: help

.PHONY: help
help:
	@echo 'Makefile for "holgersync" development.'
	@echo ''
	@echo 'Usage:'
	@echo '  run-holgersync-test    - Run the devolpment command for holgersync development purpose.'
	@echo '  build-holgersync       - Build the hoglersync application.'

.PHONY: run-development
run-sync-development:
	go run cmd/main.go sync --debug --log-to-file="holgersync.log"

.PHONY: build-holgersync
build-holgersync:
	go build -o holgersync ./...