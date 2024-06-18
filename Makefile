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
	default \
	clean \
	clean-artifacts \
	clean-releases \
	tools \
	test \
	vet \
	errors \
	lint \
	imports \
	fmt \
	env \
	debug \
	build \
	build-all \
	doc \
	release \
	package-release \
	version \
	golangci-lint-new \
    golangci-lint \
    golangci-lint-local

all: imports fmt lint vet errors build

help:
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@echo '    help                 Show this help screen.'
	@echo '    debug                Run the project.'
	@echo '    clean                Remove binaries, artifacts and releases.'
	@echo '    clean-artifacts      Remove build artifacts only.'
	@echo '    clean-releases       Remove releases only.'
	@echo '    tools                Install tools needed by the project.'
	@echo '    deps                 Download and install build time dependencies.'
	@echo '    test                 Run unit tests.'
	@echo '    vet                  Run go vet.'
	@echo '    errors               Run errcheck.'
	@echo '    lint                 Run golint.'
	@echo '    imports              Run goimports.'
	@echo '    fmt                  Run go fmt.'
	@echo '    env                  Display Go environment.'
	@echo '    build                Build project for current platform.'
	@echo '    build-all            Build project for all supported platforms.'
	@echo '    docker-build         Build project using docker.'
	@echo '    doc                  Start Go documentation server on port 8080.'
	@echo '    version              Check the Go version.'
	@echo '    golangci-lint-new    Lint new changes.'
	@echo '    golangci-lint        Lint all the code.'
	@echo '    golangci-lint-local  Lint all the code using local binary and fix problems.'
	@echo ''
	@echo 'Targets run by default are: imports, fmt, lint, vet, errors and build.'
	@echo ''

print-%:
	@echo $* = $($*)

clean: clean-artifacts clean-releases
	go clean -i ./...
	rm -vf $(CURDIR)/bin/$(BINARY)

clean-artifacts:
	rm -Rf bin/artifacts

clean-releases:
	rm -Rf bin/releases

tools:
	go get golang.org/x/tools/cmd/goimports
	go get github.com/kisielk/errcheck
	go get github.com/golang/lint/golint
	go get github.com/axw/gocov/gocov
	go get github.com/matm/gocov-html
	go get github.com/tools/godep
	go get github.com/mitchellh/gox

deps:
	go get -v ./...

test:
	cd iso20022 && go test -v ./...

vet:
	go vet -v ./...

errors:
	errcheck -ignoretests -blank ./...

lint:
	golint ./..

imports:
	goimports -l -w .

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

release: package-release

package-release:
	@test -x $(CURDIR)/bin/artifacts/$(REV) || exit 1
	mkdir -v -p $(CURDIR)/bin/releases/$(REV)
	for release in $$(find $(CURDIR)/bin/artifacts/$(REV) -mindepth 1 -maxdepth 1 -type d 2>/dev/null); do \
  		platform=$$(basename $$release); \
  		pushd $$release &>/dev/null; \
  		zip $(CURDIR)/bin/releases/$(REV)/$(BINARY)_$${platform}.zip $(BINARY); \
  		popd &>/dev/null; \
  	done

doc:
	godoc -index

version:
	@go version

mockgen-go:
	which mockgen || go install go.uber.org/mock/mockgen@v0.4.0
	cd iso20022 && go generate ./...
	make golangci-lint-local

golangci-lint-new:
	docker pull golangci/golangci-lint:v1.59.1
	docker run --rm -v $$(pwd):/app -v ~/.netrc:/root/.netrc -e GOPRIVATE=github.com/CoreumFoundation -w /app golangci/golangci-lint:v1.59.1 golangci-lint run --new-from-rev master ./...

golangci-lint:
	docker pull golangci/golangci-lint:1.59.1
	docker run --rm -v $$(pwd):/app -v ~/.netrc:/root/.netrc -e GOPRIVATE=github.com/CoreumFoundation -w /app golangci/golangci-lint:1.59.1 golangci-lint run ./...

golangci-lint-local:
	@if test ! -e ./bin/golangci-lint; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin v1.59.1; \
	fi
	@cd iso20022 && ../bin/golangci-lint run
