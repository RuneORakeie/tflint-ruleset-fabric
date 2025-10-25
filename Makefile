.PHONY: build
build:
	go build -o tflint-ruleset-fabric

.PHONY: install
install: build
	mkdir -p ~/.tflint.d/plugins
	mv ./tflint-ruleset-fabric ~/.tflint.d/plugins

.PHONY: test
test:
	go test -v ./...

.PHONY: test-manual
test-manual:
	go test -v ./rules -run '^Test[^G]'

.PHONY: test-generated
test-generated:
	go test -v ./rules -run TestGenerated

.PHONY: test-generated-discovery
test-generated-discovery:
	go test -v ./rules -run TestGeneratedRulesDiscovery

.PHONY: test-generated-all
test-generated-all:
	go test -v ./rules -run TestGeneratedRules

.PHONY: test-generated-filesystem
test-generated-filesystem:
	go test -v ./rules -run TestGeneratedRulesDiscoveryFromFilesystem

.PHONY: test-coverage
test-coverage:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -html=coverage.txt -o coverage.html

.PHONY: fmt
fmt:
	go fmt ./...
	terraform fmt -recursive examples/

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: clean
clean:
	rm -f tflint-ruleset-fabric

.PHONY: doc
doc:
	go run ./rules/generator/main.go

.PHONY: generate-rules
generate-rules:
	go run ./rules/generator \
		-specs ../fabric-rest-api-specs \
		-provider ../terraform-provider-fabric \
		-verbose

.PHONY: release
release: test build
	@echo "Build successful. Ready for release."

.PHONY: all
all: fmt lint test build

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build                      - Build the ruleset binary"
	@echo "  install                    - Install the ruleset plugin"
	@echo "  test                       - Run all tests"
	@echo "  test-manual                - Run manual rule tests only"
	@echo "  test-generated             - Run generated rule tests (all TestGenerated*)"
	@echo "  test-generated-all         - Run all TestGeneratedRules* tests"
	@echo "  test-generated-discovery   - Run rule discovery tests"
	@echo "  test-generated-filesystem  - Run filesystem discovery tests"
	@echo "  test-coverage              - Run tests with coverage report"
	@echo "  fmt                        - Format code and Terraform files"
	@echo "  lint                       - Run linter"
	@echo "  clean                      - Clean build artifacts"
	@echo "  generate-rules             - Generate rules from Fabric API specs"
	@echo "  release                    - Build release (test + build)"
	@echo "  all                        - Run fmt, lint, test, build"
	@echo "  help                       - Show this help message"
