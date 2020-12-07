VERSION ?= SNAPSHOT
PACKAGES := $(shell go list ./...)

default: help

.PHONY: test
test:
	go test -v ./... -coverprofile=cover.out

integration: test

.PHONY: fix-imports
fix-imports: ## Run goimports locally to fix any issues
	goimports -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: check-imports
check-imports: ## Script to check goimports has been run on the branch
	hack/verify-goimports.sh

.PHONY: all
all: test check-imports

.PHONY: help
help: ## Help
	@echo "Please use 'make <target>' where <target> is ..."
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
