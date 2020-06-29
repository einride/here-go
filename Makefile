# all: run a complete build
.PHONY: all
all: \
	openapi-generate \
	go-lint \
	go-review \
	go-test \
	go-build \
	mod-tidy \
	git-verify-nodiff

export GO111MODULE := on

include tools/git-verify-nodiff/rules.mk
include tools/golangci-lint/rules.mk
include tools/goreview/rules.mk

openapi_base_path := $(PWD)/pkg/openapi

.PHONY: openapi-generate
openapi-generate:
	# Routing v8
	docker run --rm \
		-v $(openapi_base_path)/routing/v8:/gen \
		-u $(shell id -u) \
		-e "JAVA_OPTS=-Dmodels -DmodelDocs=false -Dapis -DapiDocs=false -DsupportingFiles=client.go,configuration.go,response.go" \
		openapitools/openapi-generator-cli:v4.3.1 generate \
		-p "packageName=routingv8" \
		-p "enumClassPrefix=true" \
		-i https://developer.here.com/documentation/routing-api/8.3.1/swagger/v8.yaml \
		-g go \
		-o /gen \
		--http-user-agent "einride/here-go"
	# TODO: Find a way to not generate these folders
	rm -rf $(openapi_base_path)/routing/v8/api
	rm -rf $(openapi_base_path)/routing/v8/.openapi-generator
	rm -rf $(openapi_base_path)/routing/v8/.openapi-generator-ignore

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
	$(goreview) -c 1 ./pkg/routing/...

# go-test: run Go test suite
.PHONY: go-test
go-test:
	go test -count 1 -cover -race ./...

# mod-tidy: ensure Go module files are in sync
.PHONY: mod-tidy
mod-tidy:
	go mod tidy
