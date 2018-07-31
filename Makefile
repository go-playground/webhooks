GOCMD=go

linters-install:
	@gometalinter --version >/dev/null 2>&1 || { \
		echo "installing linting tools..."; \
		$(GOCMD) get github.com/alecthomas/gometalinter; \
		gometalinter --install; \
	}

lint: linters-install
	@gofmt -l . >gofmt.test 2>&1 && if [ -s gofmt.test ]; then echo "Fix formatting using 'gofmt -s -w .' for:"; cat gofmt.test; exit 1; fi && rm gofmt.test
	gometalinter --vendor --disable-all --enable=vet --enable=vetshadow --enable=golint --enable=megacheck --enable=ineffassign --enable=misspell --enable=errcheck --enable=goconst ./...

test:
	$(GOCMD) test -cover -race ./...

.PHONY: test lint linters-install