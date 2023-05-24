GO := go
GOPATH := $(shell go env GOPATH)
GOPATH_BIN := $(GOPATH)/bin
GOLANGCI_LINT := $(GOPATH_BIN)/golangci-lint
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
GOIMPORTS := $(GOPATH_BIN)/goimports
GO_PACKAGES = $(shell go list ./... | grep -v vendor)
PACKAGE_BASE := github.com/sdslabs/nymeria

.PHONY: help vendor build run dev lint format clean

help:
	@echo "Nymeria make help"
	@echo ""
	@echo "vendor: Downloads the dependencies in the vendor folder"
	@echo "build: Builds the binary of the server"
	@echo "run: Runs the binary of the server"
	@echo "dev: Combines build and run commands"
	@echo "lint: Lints the code using vet and golangci-lint"
	@echo "format: Formats the code using fmt and golangci-lint"
	@echo "clean: Removes the vendor directory and binary"

vendor:
	@${GO} mod tidy
	@${GO} mod vendor
	@echo "Vendor downloaded successfully"

build:
	@${GO} build -o nymeria ./cmd/nymeria/main.go
	@echo "Binary built successfully"

run:
	@./nymeria

dev:
	@$(GOPATH_BIN)/air -c .air.toml

install-golangci-lint:
	@echo "=====> Installing golangci-lint..."
	@curl -sSfL \
	 	https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
	 	sh -s -- -b $(GOPATH_BIN) v1.52.2

lint: install-golangci-lint
	@$(GO) vet $(GO_PACKAGES)
	@$(GOLANGCI_LINT) run -c golangci.yaml
	@echo "Lint successful"

install-goimports:
	@echo "=====> Installing formatter..."
	@$(GO) install golang.org/x/tools/cmd/goimports@latest

format: install-goimports
	@echo "=====> Formatting code..."
	@$(GOIMPORTS) -l -w -local ${PACKAGE_BASE} $(SRC)
	@echo "Format successful"

## verify: Run format and lint checks
verify: verify-format lint

## verify-format: Verify the format
verify-format: install-goimports
	@echo "=====> Verifying format..."
	$(if $(shell $(GOIMPORTS) -l -local ${PACKAGE_BASE} ${SRC}), @echo ERROR: Format verification failed! && $(GOIMPORTS) -l -local ${PACKAGE_BASE} ${SRC} && exit 1)

clean:
	@rm -f nymeria
	@rm -rf vendor/
	@echo "Clean successful"

install-air:
	@echo "Make sure your GOPATH and GOPATH_BIN is set"
	@curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(GOPATH_BIN)
	@echo "Air installed successfully"	

	