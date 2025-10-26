# fabric_connection_invalid_privacy_level

Validates that the `privacy_level` attribute of `fabric_connection` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_connection" "example" {
    privacy_level = "None" # Valid
}
```

## Validation Rules

- Must be one of: `None`, `Private`, `Organizational`, `Public`


## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_connection`.

## How To Fix

Update the `privacy_level` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/platform/definitions/connections.json)
