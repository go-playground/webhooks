GOPATH=$(shell go env GOPATH)

all: lint test

linters-install:
	@echo "+ $@"
	@$(GOPATH)/bin/golangci-lint --version >/dev/null 2>&1 || { \
  		echo "Install golangci-lint..."; \
		curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin; \
	}

lint: linters-install
	@echo "+ $@"
	$(GOPATH)/bin/golangci-lint run ./...

test:
	@echo "+ $@"
	GO111MODULE=on go test -covermode=atomic -race ./...

.PHONY: test lint linters-install
.DEFAULT_GOAL := all
