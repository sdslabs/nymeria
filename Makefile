PROJECT_NAME := nymeria
PACKAGE_BASE := github.com/sdslabs/nymeria
BINARY_NAME := $(PROJECT_NAME)
CMD_DIR := ./cmd/$(PROJECT_NAME)

GO := go
GOPATH := $(shell go env GOPATH)
GOPATH_BIN := $(GOPATH)/bin

BUILD_DIR := build
VERSION := $(shell git describe --tags --always --dirty)
COMMIT := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS := -ldflags "-X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.buildTime=$(BUILD_TIME) -s -w"

GOLANGCI_LINT := $(GOPATH_BIN)/golangci-lint
GOIMPORTS := $(GOPATH_BIN)/goimports
AIR := $(GOPATH_BIN)/air

SRC := $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./build/*")
GO_PACKAGES := $(shell go list ./... | grep -v vendor)

DB_HOST := $(shell grep -A6 "^db:" config.yaml | grep "host:" | head -1 | cut -d'"' -f2)
DB_PORT := $(shell grep -A6 "^db:" config.yaml | grep "port:" | head -1 | awk '{print $$2}')
DB_USER := $(shell grep -A6 "^db:" config.yaml | grep "user:" | head -1 | cut -d'"' -f2)
DB_PASS := $(shell grep -A6 "^db:" config.yaml | grep "password:" | head -1 | cut -d'"' -f2)
DB_NAME := $(shell grep -A6 "^db:" config.yaml | grep "db_name:" | head -1 | cut -d'"' -f2)

UP_MIGRATION_FILE := db/migrations/000001_init_schema.up.sql
DOWN_MIGRATION_FILE := db/migrations/000001_init_schema.down.sql

.PHONY: help all vendor build run dev test lint format clean install-tools verify verify-format

.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

all: clean vendor build

vendor:
	@$(GO) mod tidy
	@$(GO) mod vendor

build:
	@mkdir -p $(BUILD_DIR)
	@$(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)

build-debug:
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -gcflags="all=-N -l" -o $(BUILD_DIR)/$(BINARY_NAME)-debug $(CMD_DIR)

run: build
	@$(BUILD_DIR)/$(BINARY_NAME)

dev: install-air
	@$(AIR) -c .air.toml

test:
	@$(GO) test -v -race -coverprofile=coverage.out $(GO_PACKAGES)
	@$(GO) tool cover -html=coverage.out -o coverage.html

test-short:
	@$(GO) test -short $(GO_PACKAGES)

install-tools: install-golangci-lint install-goimports install-air

install-golangci-lint:
	@if [ ! -f $(GOLANGCI_LINT) ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
		sh -s -- -b $(GOPATH_BIN) v1.62.2; \
	fi

install-goimports:
	@if [ ! -f $(GOIMPORTS) ]; then \
		$(GO) install golang.org/x/tools/cmd/goimports@latest; \
	fi

install-air:
	@if [ ! -f $(AIR) ]; then \
		curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(GOPATH_BIN); \
	fi

lint: install-golangci-lint
	@$(GO) vet $(GO_PACKAGES)
	@$(GOLANGCI_LINT) run -c golangci.yaml

format: install-goimports
	@$(GOIMPORTS) -l -w -local $(PACKAGE_BASE) $(SRC)
	@$(GO) fmt $(GO_PACKAGES)

verify: verify-format lint test-short

verify-format: install-goimports
	@if [ -n "$$($(GOIMPORTS) -l -local $(PACKAGE_BASE) $(SRC))" ]; then \
		echo "ERROR: Code is not formatted properly!"; \
		$(GOIMPORTS) -l -local $(PACKAGE_BASE) $(SRC); \
		exit 1; \
	fi

clean:
	@rm -rf $(BUILD_DIR)/
	@rm -rf vendor/
	@rm -f coverage.out coverage.html

info:
	@echo "Project: $(PROJECT_NAME)"
	@echo "Version: $(VERSION)"
	@echo "Commit: $(COMMIT)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Go Version: $(shell $(GO) version)"

apply-migration:
	@PGPASSWORD=$(DB_PASS) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d $(DB_NAME) -f $(UP_MIGRATION_FILE)

rollback-migration:
	@PGPASSWORD=$(DB_PASS) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d $(DB_NAME) -f $(DOWN_MIGRATION_FILE)
