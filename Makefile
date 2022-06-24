GO := go

.PHONY: help vendor build run dev clean

help:
	@echo "Accountv2 make help"
	@echo ""
	@echo "vendor: Downloads the dependencies in the vendor folder"
	@echo "build: Builds the binary of the server"
	@echo "run: Runs the binary of the server"
	@echo "dev: Combines build and run commands"
	@echo "clean: Removes the vendor directory and binary"

vendor:
	@${GO} mod vendor

build:
	@${GO} build -o nymeria ./cmd/server/main.go

run:
	@./nymeria

dev: build run

clean:
	@rm nymeria
	@rm -rf vendor/