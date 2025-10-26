# fabric_spark_environment_settings_invalid_executor_cores

Validates that the `executor_cores` attribute of `fabric_spark_environment_settings` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_spark_environment_settings" "example" {
    executor_cores = "4" # Valid
}
```

## Validation Rules

- Must be one of: `4`, `8`, `16`, `32`, `64`
- Must be a valid int32 format


## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_spark_environment_settings`.

## How To Fix

Update the `executor_cores` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/environment/definitions.json)
