MODULE   := github.com/l2D/claude-code-command-fix
VERSION  := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT   := $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
BUILD_TIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
LDFLAGS  := -s -w \
	-X '$(MODULE)/internal/version.Version=$(VERSION)' \
	-X '$(MODULE)/internal/version.CommitSHA=$(COMMIT)' \
	-X '$(MODULE)/internal/version.BuildTime=$(BUILD_TIME)'

.DEFAULT_GOAL := help

.PHONY: build test lint fmt tidy coverage clean install help

build: ## Build both binaries
	go build -ldflags "$(LDFLAGS)" -o bin/claude-fix ./cmd/claude-fix/
	go build -ldflags "$(LDFLAGS)" -o bin/claude-command-fix ./cmd/claude-command-fix/

test: ## Run tests with race detector
	go test -v -race ./...

lint: ## Run golangci-lint
	golangci-lint run ./...

fmt: ## Format code
	gofmt -w .
	goimports -w -local $(MODULE) .

tidy: ## Tidy go modules
	go mod tidy

coverage: ## Generate test coverage report
	mkdir -p coverage
	go test -race -coverprofile=coverage/coverage.out ./...
	go tool cover -html=coverage/coverage.out -o coverage/coverage.html
	@echo "Coverage report: coverage/coverage.html"

clean: ## Remove build artifacts
	rm -rf bin/ dist/ coverage/

install: build ## Install binaries to GOPATH/bin
	cp bin/claude-fix $(GOPATH)/bin/claude-fix
	cp bin/claude-command-fix $(GOPATH)/bin/claude-command-fix

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
