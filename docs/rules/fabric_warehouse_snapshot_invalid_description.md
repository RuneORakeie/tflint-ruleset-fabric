# fabric_warehouse_snapshot_invalid_description

Validates that the `description` attribute of `fabric_warehouse_snapshot` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_warehouse_snapshot" "example" {
    description = "value"
}
```

## Validation Rules

- Maximum length: 256 characters


## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_warehouse_snapshot`.

## How To Fix

Update the `description` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/warehouseSnapshot/definitions.json)
