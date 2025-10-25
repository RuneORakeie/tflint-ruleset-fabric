# fabric_gateway_invalid_capacity_id

Validates that the `capacity_id` attribute of `fabric_gateway` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_gateway" "example" {
    capacity_id = "valid-uuid" # Must be valid uuid
}
```

## Validation Rules

- Must be a valid uuid format


## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_gateway`.

## How To Fix

Update the `capacity_id` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/platform/definitions/gateways.json)
