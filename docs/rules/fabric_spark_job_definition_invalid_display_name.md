# fabric_spark_job_definition_invalid_display_name

Validates that the `display_name` attribute of `fabric_spark_job_definition` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_spark_job_definition" "example" {
    display_name = "valid-value" # Must match pattern: ^[a-zA-Z0-9_ ]+$
}
```

## Validation Rules

- Must match pattern: `^[a-zA-Z0-9_ ]+$`
- Maximum length: 256 characters


## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_spark_job_definition`.

## How To Fix

Update the `display_name` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/sparkjobdefinition/definitions.json)
