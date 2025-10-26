# fabric_connection_invalid_connectivity_type

Validates that the `connectivity_type` attribute of `fabric_connection` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_connection" "example" {
    connectivity_type = "ShareableCloud" # Valid
}
```

## Validation Rules

- Must be one of: `ShareableCloud`, `VirtualNetworkGateway`


## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_connection`.

## How To Fix

Update the `connectivity_type` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/platform/definitions/connections.json)
