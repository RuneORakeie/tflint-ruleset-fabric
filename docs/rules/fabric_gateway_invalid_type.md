# fabric_gateway_invalid_type

Validates that the `type` attribute of `fabric_gateway` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_gateway" "example" {
    type = "VirtualNetwork" # Valid
}
```

## Validation Rules

- Must be one of: `VirtualNetwork`


## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_gateway`.

## How To Fix

Update the `type` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/platform/definitions/gateways.json)
