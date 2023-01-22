NAME     := ports
ELF      := ${NAME}
REG_NAME := ${NAME}
REG_TAG  := $$(git rev-parse --short HEAD)
IMG      := ${REG_NAME}:${REG_TAG}
LATEST   := ${REG_NAME}:latest
SHELL := bash
#
TEST_TIMEOUT	:= 20s

# GO settings
export GO111MODULE := on
GOARCH ?= $(shell go env GOHOSTARCH 2>/dev/null)
GOOS ?= $(shell go env GOOS 2>/dev/null)
GOFLAGS ?= $(GOFLAGS:)
GOPATH := $(shell go env GOPATH)

usage:
	@echo "USAGE:"
	@echo "   make command [options]"
	@echo
	@echo "COMMANDS:"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed 's/^/   /' | sed -e 's/\\$$/AA/' | sed -e 's/#//g' | column -t -s ":" | sort -k1

clean:
	rm -f ./bin/*
	go clean
	go clean -modcache
	go clean --testcache

tools: ### Download tools required to build service
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F' ' '{print $$4}' | xargs -n 1 go install

lint: ### Run linters
ifeq (, $(shell which golangci-lint))
	$(error "No golangci-lint in PATH, consider installing it first – https://github.com/golangci/golangci-lint")
endif
	@echo "Running linters"
	golangci-lint run --timeout 10m

format: ### Gofmt & goimport code
	@echo "Format code"
	@go fmt ./cmd/... ./pkg/... ./internal/...
	@"$(shell go env GOPATH)/bin/goimports" -w ./cmd/ ./pkg/ ./internal/

gen: ### Generate mocks
	@GOSUMDB=off go generate ./pkg/... ./internal/...

compile: ### Compile service
	@echo -n "Compiling service..."
	@go build -o ./bin/${ELF} ./cmd

build: ### Build service
	@make tools
	@make gen
	@make format
	@make compile
	@echo "complete (./bin/${ELF})"

test: ### Run tests
	@echo "Running tests"
	CGO_ENABLED=1 go test -timeout ${TEST_TIMEOUT} ./pkg/... ./internal/...

precommit: ### Ensure code is ready to be committed
	@make format
	@make lint
	@make test

run: ### Run service
	@make build
	./bin/${ELF} $(ARGS)
