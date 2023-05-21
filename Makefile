all: lint test

lint: linters-install
	golangci-lint run --timeout 5m

test:
	go test -covermode=atomic -race ./...

.PHONY: test lint linters-install
.DEFAULT_GOAL := all
