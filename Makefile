GO := go
GOPATH := $(shell go env GOPATH)
GOPATH_BIN := $(GOPATH)/bin
GOLANGCI_LINT := $(GOPATH_BIN)/golangci-lint
GO_PACKAGES := $(shell go list ./... | grep -v vendor)

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
	@curl -sSfL \
	 	https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
	 	sh -s -- -b $(GOPATH_BIN) v1.46.2

lint:
	@$(GO) vet $(GO_PACKAGES)
	@$(GOLANGCI_LINT) run
	@echo "Lint successful"

format:
	@$(GO) fmt $(GO_PACKAGES)
	@$(GOLANGCI_LINT) run --fix
	@echo "Format successful"

clean:
	@rm -f nymeria
	@rm -rf vendor/
	@echo "Clean successful"

install-air:
	@echo "Make sure your GOPATH and GOPATH_BIN is set"
	@curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(GOPATH_BIN)
	@echo "Air installed successfully"	

	