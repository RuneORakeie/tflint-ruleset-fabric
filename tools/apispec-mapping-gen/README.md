# API Spec Mapping Generator

Automated tool for generating HCL mapping files from Microsoft Fabric REST API specifications.

## Overview

This tool scans Fabric API specs and generates mapping files that connect Terraform resource attributes to API properties with their validation constraints. These mapping files are then consumed by the [rule generator](../apispec-rule-gen) to create TFLint validation rules.

## Workflow

```
fabric-rest-api-specs/
   ├── lakehouse/definitions.json
   ├── warehouse/definitions.json
   └── ...
         ↓
    apispec-mapping-gen
         ↓
    ../apispec-rule-gen/mappings/
         ├── fabric_lakehouse.hcl
         ├── fabric_warehouse.hcl
         └── ...
```

## Prerequisites

### Clone Fabric REST API Specs

```bash
cd /path/to/your/workspace
git clone https://github.com/microsoft/fabric-rest-api-specs.git
```

Your directory structure should look like:
```
workspace/
├── fabric-rest-api-specs/    ← API specs repo
└── tflint-ruleset-fabric/
    └── tools/
        ├── apispec-mapping-gen/  ← This tool
        └── apispec-rule-gen/
            └── mappings/         ← Output directory
```

## Usage

### Generate Mapping Files

From the `tools/apispec-mapping-gen` directory:

```bash
go run . -specs ../../../fabric-rest-api-specs -output ../apispec-rule-gen/mappings
```

### Command-Line Options

| Flag | Default | Description |
|------|---------|-------------|
| `-specs` | `../fabric-rest-api-specs` | Path to fabric-rest-api-specs repository |
| `-output` | `mappings` | Output directory for generated mapping files |
| `-skip-existing` | `true` | Skip files that already exist (prevents overwriting manual edits) |

### Examples

**Generate all mappings (skip existing)**:
```bash
go run . -specs ../../../fabric-rest-api-specs -output ../apispec-rule-gen/mappings
```

**Regenerate everything (overwrite all)**:
```bash
go run . -specs ../../../fabric-rest-api-specs -output ../apispec-rule-gen/mappings -skip-existing=false
```

**Use custom paths**:
```bash
go run . -specs /path/to/api/specs -output /path/to/output
```

## Output

The tool generates `.hcl` mapping files following this structure:

```hcl
// fabric_lakehouse.hcl
mapping "fabric_lakehouse" {
  import_path = "lakehouse/definitions.json"

  attribute "display_name" {
    api_ref = "CreateLakehouseRequest.displayName"
    max_length = 256
    min_length = 1
  }

  attribute "description" {
    api_ref = "CreateLakehouseRequest.description"
    max_length = 256
  }
}
```

### What It Discovers

The tool automatically extracts:

- **maxLength** - Maximum string length
- **minLength** - Minimum string length
- **pattern** - Regular expression pattern
- **enum** - Allowed values
- **maximum** - Maximum numeric value
- **minimum** - Minimum numeric value
- **format** - String format (uuid, uri, email, etc.)
- **readOnly** - Read-only properties

## Resource Name Mapping

The tool converts API spec names to Terraform resource names:

### Directory Prefixes

Some API directories add prefixes to resource names:

```go
"spark" → "spark_"  // spark/customPool → fabric_spark_custom_pool
```

### Resource Name Overrides

Some resources need special naming:

| API Spec Name | Terraform Name |
|---------------|----------------|
| `graphqlapi` | `graphql_api` |
| `kqldatabase` | `kql_database` |
| `mlmodel` | `ml_model` |
| `mlexperiment` | `ml_experiment` |

These mappings are defined in `main.go` and can be extended as needed.

## Manual Customization

### When to Customize

After initial generation, you may need to manually edit mapping files to:

1. **Fix incorrect API references** - Tool may pick wrong Create*Request type
2. **Add missing constraints** - API specs don't always have all validations
3. **Override auto-detected values** - Refine constraints based on testing
4. **Handle special cases** - Block attributes, nested objects, etc.

### Protection from Overwriting

By default (`-skip-existing=true`), the tool will **NOT overwrite** existing mapping files. This protects manual customizations.

**Output when skipping**:
```
Skipped fabric_lakehouse.hcl (already exists)
Skipped fabric_warehouse.hcl (already exists)
```

**To regenerate specific files**:
```bash
rm ../apispec-rule-gen/mappings/fabric_lakehouse.hcl
go run . -specs ../../../fabric-rest-api-specs -output ../apispec-rule-gen/mappings
```

## Common Patterns

### Standard Resource

Most resources follow this pattern:

```hcl
mapping "fabric_resource_name" {
  import_path = "resourcetype/definitions.json"
  
  attribute "display_name" {
    api_ref = "CreateResourceRequest.displayName"
    max_length = 256
    min_length = 1
  }
  
  attribute "description" {
    api_ref = "CreateResourceRequest.description"
    max_length = 256
  }
}
```

### Enum Validation

```hcl
attribute "connectivity_type" {
  api_ref = "CreateConnectionRequest.connectivityType"
  valid_values = ["ShareableCloud", "VirtualNetworkGateway"]
}
```

### Pattern Validation

```hcl
attribute "name" {
  api_ref = "CreateResourceRequest.name"
  pattern = "^[a-zA-Z0-9_]+$"
  max_length = 64
}
```

## Troubleshooting

### Missing API Spec Files

**Error**: `open ../fabric-rest-api-specs: no such file or directory`

**Solution**: Clone the API specs repository or adjust the `-specs` path.

### Wrong Create*Request Type

**Symptom**: Generated mapping has incorrect `api_ref` value.

**Solution**: Manually edit the mapping file and change the request type (e.g., `UpdateRequest` → `CreateRequest`).

### Missing Constraints

**Symptom**: API spec doesn't include all validation constraints.

**Solution**: Manually add constraints to the mapping file based on documentation or testing.

## Next Steps

After generating mapping files:

1. **Review generated mappings** - Check for incorrect API references
2. **Customize as needed** - Add missing constraints or fix errors
3. **Generate rules** - Use [apispec-rule-gen](../apispec-rule-gen) to create validation rules

## Reference

For detailed mapping file syntax and examples, see [MAPPING_GUIDE.md](../apispec-rule-gen/MAPPING_GUIDE.md).
