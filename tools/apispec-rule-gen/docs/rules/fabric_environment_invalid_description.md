# fabric_environment_invalid_description

Validates that the `description` attribute of `fabric_environment` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_environment" "example" {
    description = "value"
}
```

## Validation Rules



## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_environment`.

## How To Fix

Update the `description` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/environment/definitions.json)
