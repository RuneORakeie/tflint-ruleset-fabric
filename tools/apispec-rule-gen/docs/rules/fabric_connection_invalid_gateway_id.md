# fabric_connection_invalid_gateway_id

Validates that the `gateway_id` attribute of `fabric_connection` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_connection" "example" {
    gateway_id = "valid-uuid" # Must be valid uuid
}
```

## Validation Rules

- Must be a valid uuid format


## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_connection`.

## How To Fix

Update the `gateway_id` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/platform/definitions/connections.json)
