# fabric_mounted_data_factory_invalid_display_name

Validates that the `display_name` attribute of `fabric_mounted_data_factory` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_mounted_data_factory" "example" {
    display_name = "value"
}
```

## Validation Rules



## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_mounted_data_factory`.

## How To Fix

Update the `display_name` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/mountedDataFactory/definitions.json)
