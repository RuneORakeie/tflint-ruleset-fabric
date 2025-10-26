# Tools

Development and code generation tools for tflint-ruleset-fabric.

## Tools Overview

### apispec-rule-gen

Generates TFLint validation rules from Fabric REST API specifications.

**Purpose**: Automatically create rules that validate Terraform resources against the official Fabric API constraints (enums, patterns, min/max values, etc.)

**Quick Start**:
```bash
cd apispec-rule-gen
go run . -specs-path /path/to/fabric-rest-api-specs
```

See [apispec-rule-gen/README.md](./apispec-rule-gen/README.md) for details.

## Architecture

Inspired by [tflint-ruleset-azurerm](https://github.com/terraform-linters/tflint-ruleset-azurerm/tree/master/tools):

1. **Mapping Files** (HCL) - Define resource-to-spec relationships
2. **Spec Loader** - Parse Swagger/OpenAPI specs
3. **Constraint Extractor** - Extract validation rules
4. **Code Generator** - Generate Go rules and docs

## Workflow

1. Create/update mapping files in `apispec-rule-gen/mappings/`
2. Run generator to create rules
3. Generated files:
   - `rules/apispec/*.go` - Rule implementations
   - `docs/rules/*.md` - Rule documentation
   - `rules/apispec/provider.go` - Rule registry

## Dependencies

- Fabric REST API Specs (external)
- Terraform Provider Fabric (for schema validation)
- TFLint Plugin SDK

## Development

```bash
# Install dependencies
cd apispec-rule-gen
go mod download

# Run generator
go run . -specs-path ../../../fabric-rest-api-specs

# Verify generated code
cd ../..
go test ./rules/...
```
