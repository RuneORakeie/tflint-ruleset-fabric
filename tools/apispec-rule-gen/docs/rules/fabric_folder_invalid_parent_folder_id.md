# fabric_folder_invalid_parent_folder_id

Validates that the `parent_folder_id` attribute of `fabric_folder` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_folder" "example" {
    parent_folder_id = "valid-uuid" # Must be valid uuid
}
```

## Validation Rules

- Must be a valid uuid format


## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_folder`.

## How To Fix

Update the `parent_folder_id` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/platform/definitions/platform.json)
