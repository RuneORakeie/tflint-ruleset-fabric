# fabric_workspace_managed_private_endpoint_invalid_target_private_link_resource_id

Validates that the `target_private_link_resource_id` attribute of `fabric_workspace_managed_private_endpoint` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_workspace_managed_private_endpoint" "example" {
    target_private_link_resource_id = "value"
}
```

## Validation Rules



## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_workspace_managed_private_endpoint`.

## How To Fix

Update the `target_private_link_resource_id` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/platform/definitions/managedPrivateEndpoint.json)
