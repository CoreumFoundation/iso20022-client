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
ARCHITECTURES=amd64 arm64
# LDFLAGS
LDFLAGS := -X 'github.com/CoreumFoundation/iso20022-client/iso20022/buildinfo.VersionTag=$(VERSION_TAG)' -X 'github.com/CoreumFoundation/iso20022-client/iso20022/buildinfo.GitCommit=$(REV)' -X 'github.com/CoreumFoundation/iso20022-client/iso20022/buildinfo.BuildTime=$(BUILD_TIME)'

ENTRY_FILE=iso20022/cmd/main.go

.PHONY: \
	help \
	clean \
	test \
	integration-test \
	vet \
	fmt \
	env \
	debug \
	build \
	build-all \
	version \
	generate \
    lint

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
	@echo '    integration-test     Run integration tests.'
	@echo '    vet                  Run go vet.'
	@echo '    fmt                  Run go fmt.'
	@echo '    env                  Display Go environment.'
	@echo '    build                Build project for current platform.'
	@echo '    build-all            Build project for all supported platforms.'
	@echo '    version              Check the Go version.'
	@echo '    generate             Generate ISO20022 messages and mocks.'
	@echo '    lint                 Lint the code.'
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

integration-test:
	mkdir -v -p $(CURDIR)/coverage/
	cd integration-tests && go test -v -coverprofile $(CURDIR)/coverage/integration-cover.out -covermode=atomic ./...

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

generate: generate-messages generate-mocks lint

generate-mocks:
	which mockgen || go install go.uber.org/mock/mockgen@v0.4.0
	cd iso20022 && go generate ./...

generate-messages:
	@which moovio_xsd2go || cd iso20022-messages && go install github.com/moov-io/xsd2go/cli/moovio_xsd2go
	@cd iso20022-messages && moovio_xsd2go convert \
		xsd/xmldsig-core-schema.xsd \
		github.com/CoreumFoundation/iso20022-client/iso20022-messages \
		gen \
		--template-name=internal/templates/xmldsig.tgo \
		--xmlns-override="http://www.w3.org/2000/09/xmldsig#=xmldsig"
	@cd iso20022-messages && moovio_xsd2go convert \
		xsd/messages.xsd \
		github.com/CoreumFoundation/iso20022-client/iso20022-messages \
		gen \
		--template-name=internal/templates/model.tgo \
		--template-name=internal/templates/write.tgo \
		--template-name=internal/templates/validate.tgo \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:head.001.001.01=head_001_001_01" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:head.001.001.02=head_001_001_02" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:head.001.001.04=head_001_001_04" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:head.002.001.01=head_002_001_01" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.002.001.07=pacs_002_001_07" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08=pacs_002_001_08" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10=pacs_002_001_10" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.002.001.11=pacs_002_001_11" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.002.001.12=pacs_002_001_12" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.002.001.14=pacs_002_001_14" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.003.001.11=pacs_003_001_11" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08=pacs_003_001_08" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10=pacs_004_001_10" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.004.001.13=pacs_004_001_13" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.007.001.10=pacs_007_001_10" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.007.001.13=pacs_007_001_13" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.06=pacs_008_001_06" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08=pacs_008_001_08" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09=pacs_008_001_09" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.12=pacs_008_001_12" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08=pacs_009_001_08" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.009.001.09=pacs_009_001_09" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.009.001.11=pacs_009_001_11" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.010.001.04=pacs_010_001_04" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.010.001.06=pacs_010_001_06" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03=pacs_028_001_03" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.028.001.04=pacs_028_001_04" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.028.001.06=pacs_028_001_06" \
		--xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.029.001.02=pacs_029_001_02"
	@find ./iso20022-messages/gen -name '*.go' -exec gofmt -w {} \; -exec goimports -w {} \;

lint:
	@if test ! -e ./bin/golangci-lint; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin v1.59.1; \
	fi
	@cd iso20022 && ../bin/golangci-lint run --timeout 15m
