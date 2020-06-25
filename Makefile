# all: run a complete build
.PHONY: all
all: \
	go-lint \
	go-review \
	go-test \
	mod-tidy \
	go-build \
	git-verify-nodiff

export GO111MODULE := on

include tools/git-verify-nodiff/rules.mk
include tools/golangci-lint/rules.mk
include tools/goreview/rules.mk

# go-build: ensure the library can be cross-compiled to supported OSes
.PHONY: go-build
go-build:
	GOOS=darwin go build ./...
	GOOS=windows go build ./...
	GOOS=linux go build ./...

.PHONY: go-lint
go-lint: $(golangci_lint)
	$(golangci_lint) run

.PHONY: go-review
go-review: $(goreview)
	$(goreview) -c 1 ./...

# go-test: run Go test suite
.PHONY: go-test
go-test:
	go test -count 1 -cover -race ./...

# mod-tidy: ensure Go module files are in sync
.PHONY: mod-tidy
mod-tidy:
	go mod tidy
