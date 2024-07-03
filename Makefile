SHELL := /bin/bash

GOBASE= $(shell pwd)
GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

#  Binary Name
BINARY := iso20022-client
VERSION_TAG := $(shell git tag --points-at HEAD || echo devel)
REV := $(shell git rev-parse HEAD || echo '')
CHANGES := $(shell test -n "$$(git status --porcelain)" && echo '+CHANGES' || true)
BUILD_TIME:= $(shell date +"%Y-%m-%d %H:%M:%S")

# List of Target OS to build the binaries
PLATFORMS=darwin linux windows
ARCHITECTURES=386 amd64
# LDFLAGS
LDFLAGS := -X 'github.com/CoreumFoundation/iso20022-client/iso20022/buildinfo.VersionTag=$(VERSION_TAG)' -X 'github.com/CoreumFoundation/iso20022-client/iso20022/buildinfo.GitCommit=$(REV)' -X 'github.com/CoreumFoundation/iso20022-client/iso20022/buildinfo.BuildTime=$(BUILD_TIME)'

ENTRY_FILE=iso20022/cmd/main.go

.PHONY: \
	help \
	clean \
	test \
	vet \
	fmt \
	env \
	debug \
	build \
	build-all \
	version \
	generate \
	lint-new \
    lint \
    lint-local

all: fmt lint vet build

help:
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@echo '    help                 Show this help screen.'
	@echo '    debug                Run the project.'
	@echo '    clean                Remove binaries.'
	@echo '    deps                 Download and install build time dependencies.'
	@echo '    test                 Run unit tests.'
	@echo '    vet                  Run go vet.'
	@echo '    fmt                  Run go fmt.'
	@echo '    env                  Display Go environment.'
	@echo '    build                Build project for current platform.'
	@echo '    build-all            Build project for all supported platforms.'
	@echo '    version              Check the Go version.'
	@echo '    generate             Generate mocks.'
	@echo '    lint-new             Lint new changes.'
	@echo '    lint                 Lint all the code.'
	@echo '    lint-local           Lint all the code using local binary and fix problems.'
	@echo ''
	@echo 'Targets run by default are: imports, fmt, lint, vet, errors and build.'
	@echo ''

clean:
	go clean -i ./...
	rm -vf $(CURDIR)/bin/$(BINARY)

deps:
	cd iso20022 && go get -v ./...

test:
	mkdir -v -p $(CURDIR)/coverage/
	cd iso20022 && go test -v -coverprofile $(CURDIR)/coverage/iso20022-cover.out -covermode=atomic ./...

vet:
	cd iso20022 && go vet -v ./...

fmt:
	go fmt ./...

env:
	@go env

debug:
	$(GORUN) -ldflags "$(LDFLAGS)" $(ENTRY_FILE)

build:
	$(GOBUILD) -ldflags "$(LDFLAGS)" -o bin/$(BINARY) $(ENTRY_FILE)

build-all:
	mkdir -v -p $(CURDIR)/bin/artifacts/$(REV)
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell GOOS=$(GOOS) GOARCH=$(GOARCH) go build -v -o bin/artifacts/$(REV)/$(GOOS)_$(GOARCH)/$(BINARY) $(ENTRY_FILE))))

version:
	@go version

generate:
	which mockgen || go install go.uber.org/mock/mockgen@v0.4.0
	cd iso20022 && go generate ./...
	make golangci-lint-local

lint-new:
	docker pull golangci/golangci-lint:v1.59.1
	docker run --rm -v $(CURDIR)/iso20022:/app -e GOPRIVATE=github.com/CoreumFoundation -w /app golangci/golangci-lint:v1.59.1 golangci-lint run --new-from-rev master ./...

lint:
	docker pull golangci/golangci-lint:v1.59.1
	docker run --rm -v $(CURDIR)/iso20022:/app -e GOPRIVATE=github.com/CoreumFoundation -w /app golangci/golangci-lint:v1.59.1 golangci-lint run ./...

lint-local:
	@if test ! -e ./bin/golangci-lint; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin v1.59.1; \
	fi
	@cd iso20022 && ../bin/golangci-lint run
