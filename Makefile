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

.PHONY: release
release: test build
	@echo "Build successful. Ready for release."

.PHONY: all
all: fmt lint test build