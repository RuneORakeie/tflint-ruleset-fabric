.PHONY: build
build:
	go build -v -o tflint-ruleset-fabric

.PHONY: install
install: build
	mkdir -p ~/.tflint.d/plugins
	cp ./tflint-ruleset-fabric ~/.tflint.d/plugins/  # Changed from mv to cp

.PHONY: test
test:
	go test -v ./...

.PHONY: test-coverage
test-coverage:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -html=coverage.txt -o coverage.html

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: clean
clean:
	rm -f tflint-ruleset-fabric
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