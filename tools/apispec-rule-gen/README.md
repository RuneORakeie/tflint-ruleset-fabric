# API Spec Rule Generator

Automated tool for generating TFLint validation rules from HCL mapping files and Microsoft Fabric REST API specifications.

## Overview

This tool reads HCL mapping files and generates TFLint validation rules by extracting constraints from Fabric API specs and the Terraform provider schema. The generated rules validate Terraform configurations against official API requirements.

```
mappings/*.hcl          Terraform Provider         fabric-rest-api-specs/
   ├── fabric_lakehouse.hcl    ├── schema.json              ├── lakehouse/
   ├── fabric_warehouse.hcl    └── resources/               ├── warehouse/
   └── ...                         ├── fabric_lakehouse      └── ...
         ↓                         └── ...                         ↓
         └──────────────────────────────────────────────────────────
                                   ↓
                              main.go + schema.go
                                   ↓
                    ┌───────────────┴───────────────┐
                    ↓                               ↓
            rules/apispec/*.go              docs/rules/*.md
```


## Prerequisites

### 1. Clone Fabric REST API Specs

```bash
cd /path/to/your/workspace
git clone https://github.com/microsoft/fabric-rest-api-specs.git
```

### 2. Generate Mapping Files

Mapping files must exist before running the rule generator. Use the [mapping generator](../apispec-mapping-gen) tool:

```bash
cd ../apispec-mapping-gen
go run . -specs ../../../fabric-rest-api-specs -output ../apispec-rule-gen/mappings
```

See [apispec-mapping-gen/README.md](../apispec-mapping-gen/README.md) for details.

### 3. Generate Terraform Provider Schema

```bash
cd schema
terraform init
terraform providers schema -json > schema.json
cd ..
```

Your directory structure should look like:
```
workspace/
├── fabric-rest-api-specs/    ← API specs repo
└── tflint-ruleset-fabric/
    ├── tools/
    │   ├── apispec-mapping-gen/  ← Mapping generator
    │   └── apispec-rule-gen/     ← This tool
    │       ├── schema/
    │       │   └── schema.json   ← Provider schema
    │       ├── mappings/         ← HCL mapping files
    │       ├── main.go
    │       └── schema.go
    ├── rules/
    │   └── apispec/              ← Generated rules
    └── docs/
        └── rules/                ← Generated docs
```

## Quick Start

### Generate Validation Rules

**⚠️ IMPORTANT: Run from the `tools/` directory for correct file paths.**

From the **tools/ directory** (recommended):

```bash
cd tools
go run ./apispec-rule-gen -specs-path ../../fabric-rest-api-specs -base-path apispec-rule-gen -rules-path ../rules -docs-path ../docs
```

Or from the `tools/apispec-rule-gen` directory:

```bash
cd tools/apispec-rule-gen
go run . -specs-path ../../../fabric-rest-api-specs -base-path . -rules-path ../../rules -docs-path ../../docs
```

**Output**:
- Go rule files in `rules/apispec/`
- Documentation in `docs/rules/`
- Provider registry in `rules/apispec/provider.go`
- Markdown docs in `../../docs/rules/`
- Summary of generated rules and detected orphaned mappings

**What it does**:
1. Loads Terraform provider schema from `schema/schema.json`
2. Reads all mapping files from `mappings/`
3. Extracts constraints from API specs
4. Filters enum values to Terraform-supported only
5. Generates validation rules and documentation
6. Detects orphaned mappings (resources not in Terraform schema)

### Command-Line Options

| Flag | Default | Description |
|------|---------|-------------|
| `-specs-path` | (required) | Path to fabric-rest-api-specs repository |
| `-base-path` | `tools/apispec-rule-gen` | Base directory where tool expects schema/, mappings/ folders |
| `-rules-path` | `rules` | Output path for generated rules (relative to **current directory**) |
| `-docs-path` | `docs` | Output path for generated docs (relative to **current directory**) |

**Path Behavior:**
- When run from **repository root**: `-base-path tools/apispec-rule-gen` (default) reads from tool directory, writes to `rules/` and `docs/`
- When run from **tool directory**: Must specify `-base-path .` AND `-rules-path ../../rules -docs-path ../../docs` to write to correct locations

## Tool Details

### Rule Generation Process

**How it works**:
1. Loads Terraform provider schema from `schema/schema.json`
2. Reads all `.hcl` mapping files from `mappings/` directory
3. For each mapping:
   - Validates that the resource exists in Terraform schema
   - Extracts constraints from API spec files
   - Filters enum values to Terraform-supported only
   - Merges constraints from three sources (see below)
4. Generates Go validation rules in `../../rules/apispec/`
5. Generates Markdown documentation in `../../docs/rules/`
6. Reports orphaned mappings (resources not in schema)

### Key Features

#### 1. Enum Filtering
Filters API enum values to only those supported by Terraform:

```
API Spec: connectivity_type = ["ShareableCloud", "PersonalCloud", 
                               "OnPremisesGateway", "VirtualNetworkGateway", ...]
                                        ↓
Terraform schema.json: "Value must be one of : `ShareableCloud`, `VirtualNetworkGateway`"
                                        ↓
Generated Rule: enum = ["ShareableCloud", "VirtualNetworkGateway"]
```

Messages during generation:
```
ℹ️  Filtered fabric_connection.connectivity_type enum from 7 to 2 values (Terraform-supported only)
```

#### 2. Orphaned Mapping Detection
Warns about mapping files without corresponding Terraform resources:

```
⚠️  WARNING: Found mapping files without corresponding Terraform resources:
  - fabric_paginated_report (no resource found in schema.json)
  
   These mappings will be SKIPPED during rule generation.
```

#### 3. Constraint Merging
Combines constraints from three sources (priority order):
1. Manual constraints in mapping file (highest priority)
2. API spec definitions
3. Terraform schema.json (inferred from descriptions)

#### 4. API Reference Validation
Validates that all `api_ref` values point to valid properties:

```
✓ Valid:   api_ref = "CreateLakehouseRequest.displayName"
✗ Invalid: api_ref = "displayName"  ← Missing request type
✗ Invalid: api_ref = "LakehouseName" ← Old format
```

Warnings are shown for invalid references:
```
Warning: Invalid API reference 'displayName' for fabric_lakehouse.display_name 
         (cannot infer request object)
```

## Mapping File Format

See `MAPPING_GUIDE.md` for detailed documentation on the HCL mapping format.

Basic structure:

```hcl
mapping "fabric_lakehouse" {
  import_path = "lakehouse/definitions.json"

  attribute "display_name" {
    api_ref = "CreateLakehouseRequest.displayName"
    max_length = 256
  }
  
  attribute "description" {
    api_ref = "CreateLakehouseRequest.description"
    max_length = 256
  }
}
```

## Common Workflows

### Adding a New Resource

1. **Auto-generate mapping** (if API spec exists):
   ```bash
   go run analyze_specs.go -specs ../../fabric-rest-api-specs
   ```

2. **Review generated mapping** in `mappings/fabric_<resource>.hcl`

3. **Customize if needed** (add `// MANUAL:` comment to preserve on re-generation)

4. **Generate rules**:
   ```bash
   go run . schema.go -specs-path ../../fabric-rest-api-specs
   ```

### Updating Existing Mappings

1. **Check API spec** for new constraints

2. **Update mapping file**:
   - Add `// MANUAL:` comment for custom constraints
   - Update `api_ref` if API changed

3. **Regenerate rules**:
   ```bash
   go run . schema.go -specs-path ../../fabric-rest-api-specs
   ```

### Fixing Warnings

Common warnings and solutions:

| Warning | Cause | Fix |
|---------|-------|-----|
| `Invalid API reference` | Missing request type prefix | Add `CreateXxxRequest.` prefix |
| `Property 'xxx' not found` | Wrong request type | Check API spec for correct type |
| `Could not read spec file` | Wrong path or `#/definitions/` fragment | Fix `import_path` |
| `resource exists in API spec but not in Terraform` | Attribute not in schema.json | Check if attribute is computed/read-only |

## Troubleshooting

### Enum filtering not working

**Symptom**: Enum values not being filtered

**Check**:
1. Schema.json has enum description pattern: `Value must be one of : `Val1`, `Val2``
2. Resource name in mapping matches schema.json exactly
3. Attribute name matches schema.json

### Rules not generating

**Check**:
1. Mapping file exists in `mappings/`
2. Resource exists in `schema/schema.json`
3. No parse errors in mapping file (check HCL syntax)
4. `api_ref` values are valid

### Orphaned mapping warnings

**Causes**:
1. Resource not yet in Terraform provider
2. Filename doesn't match resource name
3. Resource name override needed in `analyze_specs.go`

## File Reference

### Generated Files
- `mappings/*.hcl` - Resource-to-API mappings (auto-generated + manual)
- `../../rules/apispec/*.go` - Validation rules
- `../../docs/rules/*.md` - Rule documentation
- `../../rules/apispec/provider.go` - Rule registration

### Source Files
- `main.go` - Rule generator
- `schema.go` - Terraform type definitions
- `analyze_specs.go` - Mapping generator
- `*.tmpl` - Code generation templates

### Data Files
- `schema/schema.json` - Terraform provider schema
- `../../fabric-rest-api-specs/` - API specifications

## Advanced Topics

### Custom Patterns

Go's `regexp` package doesn't support negative lookaheads. Complex validation patterns should be:

1. **Documented in mapping** with comment explaining intent
2. **Implemented in provider** if critical
3. **Simplified** to Go-compatible regex if possible

Example:
```hcl
attribute "display_name" {
  api_ref = "CreateFolderRequest.displayName"
  max_length = 255
  // Note: Original pattern used negative lookaheads not supported by Go regexp
  // Pattern enforces: no leading/trailing spaces, no special chars, 
  // not named $recycle.bin/recycled/recycler
  // Consider implementing custom validation in Terraform provider
}
```

### Update-Based Resources

Some resources use `Update*Request` instead of `Create*Request`:

```hcl
// fabric_spark_environment_settings.hcl
// Note: This uses UpdateEnvironmentSparkComputeRequest (not Create) - unusual pattern
mapping "fabric_spark_environment_settings" {
  import_path = "environment/definitions.json"
  
  attribute "driver_cores" {
    api_ref = "UpdateEnvironmentSparkComputeRequest.driverCores"
    valid_values = [4, 8, 16, 32, 64]
  }
}
```

### Merged Resources

Resources that combine properties from multiple request types:

```hcl
// fabric_connection.hcl
// Properties merged from: CreateConnectionRequest, 
//                        CreateCloudConnectionRequest,
//                        CreateVirtualNetworkGatewayConnectionRequest
mapping "fabric_connection" {
  attribute "connectivity_type" {
    api_ref = "CreateConnectionRequest.connectivityType"
    valid_values = [...]
  }
  
  // Properties that appear in multiple types map to same attribute
  attribute "gateway_id" {
    api_ref = "CreateOnPremisesConnectionRequest.gatewayId"
  }
}
```

## Contributing

When adding new mappings or fixing issues:

1. **Test generation**: Ensure rules generate without errors
2. **Check warnings**: Fix all validation warnings
3. **Document special cases**: Add comments for unusual patterns
4. **Mark manual changes**: Use `// MANUAL:` prefix for custom constraints

## See Also

- `MAPPING_GUIDE.md` - Detailed mapping file reference
- `../../RULES_SCHEMA_CONSTRAINTS.md` - Generated rules overview
- Terraform Provider: https://registry.terraform.io/providers/Azure/fabric
