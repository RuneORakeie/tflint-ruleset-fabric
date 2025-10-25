# fabric_spark_environment_settings_invalid_runtime_version

Validates that the `runtime_version` attribute of `fabric_spark_environment_settings` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_spark_environment_settings" "example" {
    runtime_version = "1.1" # Valid
}
```

## Validation Rules

- Must be one of: `1.1`, `1.2`, `1.3`


## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_spark_environment_settings`.

## How To Fix

Update the `runtime_version` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/environment/definitions.json)
