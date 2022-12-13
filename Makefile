.PHONY: help
help: ## Show help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
%:
	@:

.PHONY: tests
tests: ## Run all tests except dev tests.
	go test -v ./...

.PHONY: dev_tests
dev_tests: ## Run all tests includes dev one.
	DEV_MODE=true go test -v ./...