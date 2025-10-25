# fabric_workspace_managed_private_endpoint_invalid_name

Validates that the `name` attribute of `fabric_workspace_managed_private_endpoint` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_workspace_managed_private_endpoint" "example" {
    name = "value"
}
```

## Validation Rules



## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_workspace_managed_private_endpoint`.

## How To Fix

Update the `name` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/platform/definitions/managedPrivateEndpoint.json)
