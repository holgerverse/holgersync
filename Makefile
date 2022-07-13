.DEFAULT: help

NUMBER_OF_TESTS = 10

.PHONY: help
help:
	@echo 'Makefile for "holgersync" development.'
	@echo ''
	@echo 'Usage:'
	@echo '  run-holgersync-development    - Run the devolpment command for holgersync development purpose.'
	@echo '  build-holgersync              - Build the hoglersync application.'
	@echo '  go-get-dependencies           - Get the dependencies for the holgersync application.'
	@echo '  clean-tests                   - Clean all files created by tests.'
	@echo '  create-test-env               - Create a test environment for the holgersync application.'
	@echo '                                  - NUMBER_OF_TESTS is the number of test folders to be created.'
	@echo '                                  - GIT_USERNAME is the Git username.'
	@echo '                                  - GIT_PERSONAL_ACCESS_TOKEN is the git password/PAT.'

.PHONY: run-holgersync-development
run-sync-development:
	go run main.go sync --debug --log-to-file="holgersync.log" \
		--config-file="./tests/holgersyncfile.yml"

.PHONY: build-holgersync
build-holgersync:
	go build -o holgersync ./...

.PHONY: go-get-dependencies
go-get-dependencies:
	go mod tidy

.PHONY: create-test-env
create-test-env:
	sh scripts/create_test_env.sh $(NUMBER_OF_TESTS) $(GIT_USERNAME) $(GIT_PERSONAL_ACCESS_TOKEN)

.PHONY: clean-tests
clean-tests:
	find -f tests/* | xargs rm -rf