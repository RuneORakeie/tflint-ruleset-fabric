# Detect OS
ifeq ($(OS),Windows_NT)
	BINARY_NAME := tflint-ruleset-fabric.exe
	EXE_EXT := .exe
	PLUGINS_DIR := $(USERPROFILE)\.tflint.d\plugins
	MKDIR := powershell -NoProfile -Command "New-Item -ItemType Directory -Path '$(PLUGINS_DIR)' -Force | Out-Null"
	CP := powershell -NoProfile -Command "Copy-Item -Force '$(BINARY_NAME)' '$(PLUGINS_DIR)\tflint-ruleset-fabric$(EXE_EXT)'"
	RM := powershell -NoProfile -Command "Remove-Item -Force -ErrorAction SilentlyContinue"
else
	BINARY_NAME := tflint-ruleset-fabric
	EXE_EXT :=
	PLUGINS_DIR := $(HOME)/.tflint.d/plugins
	MKDIR := mkdir -p $(PLUGINS_DIR)
	CP := cp -f $(BINARY_NAME) $(PLUGINS_DIR)/tflint-ruleset-fabric$(EXE_EXT)
	RM := rm -f
endif

.PHONY: build
build:
	go build -v -o $(BINARY_NAME)

.PHONY: install
install: build
	@$(MKDIR)
	@$(CP)
	@echo "Installed to $(PLUGINS_DIR)"

.PHONY: uninstall
uninstall:
	@$(RM) "$(PLUGINS_DIR)/tflint-ruleset-fabric$(EXE_EXT)" || true

# Ensure e2e installs the plugin first
.PHONY: e2e
e2e: install
	cd integration && go test -v

.PHONY: test
test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic $(shell go list ./... | grep -v /integration)

.PHONY: test-coverage
test-coverage:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic $(shell go list ./... | grep -v /integration)
	go tool cover -html=coverage.txt -o coverage.html

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: clean
clean:
	rm -f tflint-ruleset-fabric$(EXE_EXT)
	rm -f coverage.txt coverage.html
	rm -rf dist/

.PHONY: goreleaser-snapshot
goreleaser-snapshot:
	goreleaser release --snapshot --clean --skip=sign

.PHONY: goreleaser-check
goreleaser-check:
	goreleaser check

.PHONY: all
all: fmt lint test build

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build                - Build the ruleset binary"
	@echo "  install              - Install plugin to ~/.tflint.d/plugins"
	@echo "  test                 - Run all tests"
	@echo "  test-coverage        - Run tests with coverage report"
	@echo "  fmt                  - Format code"
	@echo "  lint                 - Run linter"
	@echo "  clean                - Clean build artifacts"
	@echo "  goreleaser-snapshot  - Test goreleaser locally"
	@echo "  goreleaser-check     - Validate .goreleaser.yml"
	@echo "  all                  - Run fmt, lint, test, build"
	@echo "  help                 - Show this help message"
