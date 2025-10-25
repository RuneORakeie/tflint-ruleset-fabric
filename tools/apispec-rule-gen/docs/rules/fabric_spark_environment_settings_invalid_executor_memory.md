# fabric_spark_environment_settings_invalid_executor_memory

Validates that the `executor_memory` attribute of `fabric_spark_environment_settings` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_spark_environment_settings" "example" {
    executor_memory = "28g" # Valid
}
```

## Validation Rules

- Must be one of: `28g`, `56g`, `112g`, `224g`, `400g`


## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_spark_environment_settings`.

## How To Fix

Update the `executor_memory` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/environment/definitions.json)
