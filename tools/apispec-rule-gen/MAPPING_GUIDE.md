# Mapping File Guide

Complete reference for creating and maintaining HCL mapping files that connect Terraform resources to Fabric API specifications.

## Overview

Mapping files define the relationship between:
- Terraform resource attributes (from `schema.json`)
- API request properties (from `fabric-rest-api-specs`)
- Validation constraints (maxLength, enum, pattern, etc.)

They serve as the "bridge" between the Terraform provider and API specs, enabling automated rule generation.

## Basic Structure

```hcl
// Comment describing the resource
mapping "fabric_<resource_name>" {
  import_path = "<api_spec_directory>/definitions.json"

  attribute "<terraform_attribute_name>" {
    api_ref = "Create<Resource>Request.<apiPropertyName>"
    // Optional constraints
    max_length = 256
    min_length = 1
    pattern = "^[a-zA-Z0-9_]+$"
    valid_values = ["value1", "value2"]
  }
}
```

## File Naming Convention

**Format**: `fabric_<resource_name>.hcl`

**Examples**:
- `fabric_lakehouse.hcl`
- `fabric_workspace.hcl`
- `fabric_kql_database.hcl`

**Important**: The filename (without `.hcl`) must exactly match the Terraform resource name in `schema.json`.

## mapping Block

### import_path

Path to the API spec file relative to `fabric-rest-api-specs/` root.

**Format**: `<directory>/definitions.json`

**Examples**:
```hcl
import_path = "lakehouse/definitions.json"
import_path = "platform/definitions/platform.json"
import_path = "sparkjobdefinition/definitions.json"
```

**Common Mistakes**:
```hcl
// ✗ Wrong - includes #/definitions/ fragment
import_path = "platform/definitions/platform.json#/definitions/CreateFolderRequest"

// ✗ Wrong - nested subdirectory that doesn't exist
import_path = "paginatedReport/definitions/paginatedReport.json"

// ✓ Correct
import_path = "paginatedReport/definitions.json"
```

## attribute Block

Each `attribute` block maps one Terraform attribute to an API property.

### Attribute Name

The Terraform attribute name from `schema.json` (snake_case).

**Examples**:
```hcl
attribute "display_name" { ... }
attribute "connectivity_type" { ... }
attribute "driver_cores" { ... }
```

### api_ref (Required)

Reference to the API property in the format: `<RequestType>.<propertyName>`

**Format**: `Create<Resource>Request.<propertyName>`

**Examples**:
```hcl
api_ref = "CreateLakehouseRequest.displayName"
api_ref = "CreateConnectionRequest.connectivityType"
api_ref = "UpdateEnvironmentSparkComputeRequest.driverCores"  // Note: Update, not Create
```

**Property Name Mapping**:
- API uses camelCase: `displayName`, `connectivityType`
- Terraform uses snake_case: `display_name`, `connectivity_type`
- The mapping connects them

**Common Mistakes**:
```hcl
// ✗ Wrong - missing request type
api_ref = "displayName"

// ✗ Wrong - old format (pre-refactoring)
api_ref = "LakehouseName"

// ✓ Correct
api_ref = "CreateLakehouseRequest.displayName"
```

### Constraint Properties

#### max_length

Maximum string length (integer).

```hcl
attribute "display_name" {
  api_ref = "CreateLakehouseRequest.displayName"
  max_length = 256
}
```

#### min_length

Minimum string length (integer).

```hcl
attribute "name" {
  api_ref = "CreateCustomPoolRequest.name"
  min_length = 1
  max_length = 200
}
```

#### pattern

Go-compatible regular expression for validation.

```hcl
attribute "display_name" {
  api_ref = "CreateSparkJobDefinitionRequest.displayName"
  max_length = 256
  pattern = "^[a-zA-Z0-9_ ]+$"  // Only letters, numbers, underscores, spaces
}
```

**Important**: Go's `regexp` package doesn't support:
- Negative lookaheads `(?!...)`
- Positive lookaheads `(?=...)`
- Lookbehinds

For complex patterns, document the intent in comments and consider implementing in the provider.

#### valid_values

Enum values (array of strings or integers).

```hcl
// String enums
attribute "connectivity_type" {
  api_ref = "CreateConnectionRequest.connectivityType"
  valid_values = ["ShareableCloud", "PersonalCloud", "VirtualNetworkGateway"]
}

// Integer enums
attribute "driver_cores" {
  api_ref = "UpdateEnvironmentSparkComputeRequest.driverCores"
  valid_values = [4, 8, 16, 32, 64]
}
```

**Enum Filtering**: 
The rule generator automatically filters `valid_values` to only include values supported by Terraform (from `schema.json`).

```
API spec:     valid_values = [7 values]
                     ↓
Terraform:    "Value must be one of : `Val1`, `Val2`"
                     ↓
Generated:    enum = ["Val1", "Val2"]  // Filtered to 2
```

## Special Cases

### Merged Resources

Some Terraform resources combine properties from multiple API request types.

**Example**: `fabric_connection`

```hcl
// fabric_connection.hcl
// Properties merged from: CreateConnectionRequest, 
//                        CreateCloudConnectionRequest,
//                        CreateVirtualNetworkGatewayConnectionRequest

mapping "fabric_connection" {
  import_path = "platform/definitions/connections.json"

  // Base properties from CreateConnectionRequest
  attribute "connectivity_type" {
    api_ref = "CreateConnectionRequest.connectivityType"
    valid_values = ["ShareableCloud", "VirtualNetworkGateway", ...]
  }
  
  // Properties shared across multiple types map to single attribute
  attribute "gateway_id" {
    api_ref = "CreateOnPremisesConnectionRequest.gatewayId"
    // Note: Also in CreateVirtualNetworkGatewayConnectionRequest.gatewayId
    // but maps to same Terraform attribute
  }
  
  attribute "credential_details" {
    api_ref = "CreateCloudConnectionRequest.credentialDetails"
    // Note: Also in other connection types, maps to same attribute
  }
}
```

**Configured in** `analyze_specs.go`:
```go
var resourceMergeRules = map[string][]string{
    "connection": {
        "CreateCloudConnectionRequest",
        "CreateVirtualNetworkGatewayConnectionRequest",
        "CreateOnPremisesConnectionRequest",
    },
}
```

### Update-Based Resources

Most resources use `Create*Request`, but some use `Update*Request`.

**Example**: `fabric_spark_environment_settings`

```hcl
// fabric_spark_environment_settings.hcl
// Note: This uses UpdateEnvironmentSparkComputeRequest (not Create) - unusual pattern

mapping "fabric_spark_environment_settings" {
  import_path = "environment/definitions.json"

  attribute "driver_cores" {
    api_ref = "UpdateEnvironmentSparkComputeRequest.driverCores"
    valid_values = [4, 8, 16, 32, 64]
  }
  
  attribute "runtime_version" {
    api_ref = "UpdateEnvironmentSparkComputeRequest.runtimeVersion"
    valid_values = ["1.1", "1.2", "1.3"]
  }
}
```

**Why**: This resource configures Spark compute settings for an existing environment, so the API operation is an update, not a create.

### Resources with No Create Request

Some mapping files reference resources without `Create*Request` in the API spec.

**Example**: Resources not yet in the API spec or custom mappings.

```hcl
// fabric_spark_job_definition.hcl
mapping "fabric_spark_job_definition" {
  import_path = "sparkjobdefinition/definitions.json"

  // MANUAL: verified in portal
  attribute "display_name" {
    api_ref = "CreateSparkJobDefinitionRequest.displayName"
    max_length = 256
    pattern = "^[a-zA-Z0-9_ ]+$"
  }

  // MANUAL: API spec says max 256, but portal silently truncates at 1021
  attribute "description" {
    api_ref = "CreateSparkJobDefinitionRequest.description"
    max_length = 1021
    warn_on_exceed = true
  }
}
```

## Manual Customizations

Mark manual customizations with `// MANUAL:` comments to preserve them when re-running `analyze_specs.go`.

### Pattern Example

```hcl
attribute "display_name" {
  api_ref = "CreateSparkJobDefinitionRequest.displayName"
  max_length = 256
  // MANUAL: verified in portal - only letters, numbers, underscores, spaces
  pattern = "^[a-zA-Z0-9_ ]+$"
}
```

### Constraint Override

```hcl
attribute "description" {
  api_ref = "CreateSparkJobDefinitionRequest.description"
  // MANUAL: API spec says 256, but portal actually truncates at 1021
  max_length = 1021
  warn_on_exceed = true
}
```

### Complex Patterns

```hcl
attribute "display_name" {
  api_ref = "CreateFolderRequest.displayName"
  max_length = 255
  // MANUAL: Original pattern used negative lookaheads not supported by Go regexp
  // Pattern should enforce:
  //   - No leading/trailing spaces
  //   - No special chars: ~"#.&*:<>?/\{\|}
  //   - Not named: $recycle.bin, recycled, recycler
  // Consider implementing custom validation in Terraform provider
}
```

## Comments and Documentation

### File Header

```hcl
// Mapping for fabric_<resource> resource
// Auto-generated from <directory>/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.
```

### Inline Comments

```hcl
// optional, max 256 chars
attribute "description" {
  api_ref = "CreateLakehouseRequest.description"
  max_length = 256
}

// required, enum(5 values)
attribute "driver_cores" {
  api_ref = "UpdateEnvironmentSparkComputeRequest.driverCores"
  valid_values = [4, 8, 16, 32, 64]
}
```

### Merge Documentation

```hcl
// Properties merged from: CreateConnectionRequest, CreateCloudConnectionRequest
mapping "fabric_connection" {
  ...
}
```

## Complete Examples

### Simple Resource

```hcl
// Mapping for fabric_lakehouse resource
// Auto-generated from lakehouse/definitions.json

mapping "fabric_lakehouse" {
  import_path = "lakehouse/definitions.json"

  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreateLakehouseRequest.description"
    max_length = 256
  }

  // required
  attribute "display_name" {
    api_ref = "CreateLakehouseRequest.displayName"
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateLakehouseRequest.folderId"
  }
}
```

### Resource with Enums

```hcl
// Mapping for fabric_gateway resource
mapping "fabric_gateway" {
  import_path = "platform/definitions/gateway.json"

  // required
  attribute "display_name" {
    api_ref = "CreateVirtualNetworkGatewayRequest.displayName"
    max_length = 200
  }

  // required, format: uuid
  attribute "capacity_id" {
    api_ref = "CreateVirtualNetworkGatewayRequest.capacityId"
  }

  // optional, enum(10 values)
  attribute "inactivity_minutes_before_sleep" {
    api_ref = "CreateVirtualNetworkGatewayRequest.inactivityMinutesBeforeSleep"
    valid_values = [30, 60, 90, 120, 150, 240, 360, 480, 720, 1440]
  }

  // required, enum(3 values) - filters to 1
  attribute "type" {
    api_ref = "CreateGatewayRequest.type"
    valid_values = ["OnPremises", "OnPremisesPersonal", "VirtualNetwork"]
    // Note: Terraform only supports "VirtualNetwork" currently
  }
}
```

### Resource with Manual Constraints

```hcl
// Mapping for fabric_spark_job_definition resource
mapping "fabric_spark_job_definition" {
  import_path = "sparkjobdefinition/definitions.json"

  // MANUAL: verified in portal
  // MANUAL: only letters, numbers, underscores (_), or spaces
  attribute "display_name" {
    api_ref = "CreateSparkJobDefinitionRequest.displayName"
    max_length = 256
    pattern = "^[a-zA-Z0-9_ ]+$"
  }

  // MANUAL: API spec says max 256, but portal silently truncates at 1021
  // MANUAL: Using 1021 based on observed behavior
  attribute "description" {
    api_ref = "CreateSparkJobDefinitionRequest.description"
    max_length = 1021
    warn_on_exceed = true
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateSparkJobDefinitionRequest.folderId"
  }

  // optional
  attribute "definition" {
    api_ref = "CreateSparkJobDefinitionRequest.definition"
  }
}
```

## Validation and Testing

### Running the Generator

```bash
# Generate rules and check for warnings
go run . schema.go -specs-path ../../fabric-rest-api-specs 2>&1 | grep -i warning
```

### Common Validation Errors

| Error | Meaning | Fix |
|-------|---------|-----|
| `Invalid API reference 'displayName'` | Missing request type | Add `CreateXxxRequest.` prefix |
| `Property 'xxx' not found in 'CreateXxxRequest'` | Property doesn't exist in spec | Check API spec for correct property name |
| `Could not read spec file ...#/definitions/...` | Invalid import_path | Remove `#/definitions/...` fragment |
| `resource exists in API spec but not in Terraform` | Attribute not in schema.json | Attribute may be computed/read-only |

### Checking Generated Rules

After generation, verify:

1. **Rule files created**: `../../rules/apispec/fabric_<resource>_invalid_<attribute>.go`
2. **No errors in output**: Check for compilation or validation errors
3. **Enum filtering**: Look for "Filtered ... enum from X to Y values" messages
4. **Correct constraints**: Open generated `.go` files to verify constraints

## Best Practices

### 1. Always Use Full API References

```hcl
// ✓ Good
api_ref = "CreateLakehouseRequest.displayName"

// ✗ Bad
api_ref = "displayName"
```

### 2. Document Manual Changes

```hcl
// ✓ Good
// MANUAL: Portal behavior differs from API spec
attribute "description" {
  max_length = 1021
}

// ✗ Bad - no explanation
attribute "description" {
  max_length = 1021
}
```

### 3. Keep Constraints in Sync with API

When API specs change:
- Re-run `analyze_specs.go` to update auto-generated constraints
- Review `.new` files for conflicts
- Merge changes carefully

### 4. Use Comments for Complex Cases

```hcl
// Note: gatewayId appears in multiple request types but maps to single attribute
attribute "gateway_id" {
  api_ref = "CreateOnPremisesConnectionRequest.gatewayId"
}
```

### 5. Validate After Changes

Always regenerate rules after modifying mappings:
```bash
go run . schema.go -specs-path ../../fabric-rest-api-specs
```

## See Also

- `README.md` - Tool usage and workflows
- `analyze_specs.go` - Auto-generation logic
- `main.go` - Rule generation logic
- `../../fabric-rest-api-specs/` - API specifications
