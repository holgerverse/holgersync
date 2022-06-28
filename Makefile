.DEFAULT: help

.PHONY: help
help:
	@echo 'Makefile for "holgersync" development.'
	@echo ''
	@echo 'Usage:'
	@echo '  run-holgersync-development    - Run the devolpment command for holgersync development purpose.'
	@echo '  build-holgersync              - Build the hoglersync application.'
	@echo '  go-get-dependencies           - Get the dependencies for the holgersync application.'
	@echo '  clean-tests				   - Clean all files created by tests.'

.PHONY: run-holgersync-development
run-sync-development:
	go run cmd/main.go sync --debug --log-to-file="holgersync.log" \
		--config-file="./tests/holgersyncfile.yml"

.PHONY: build-holgersync
build-holgersync:
	go build -o holgersync ./...

.PHONY: go-get-dependencies
go-get-dependencies:
	go mod tidy

.PHONY: clean-tests
clean-tests:
	find -f tests/test_folder* | xargs rm -rf