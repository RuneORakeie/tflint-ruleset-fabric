# fabric_lakehouse_invalid_description

Validates that the `description` attribute of `fabric_lakehouse` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_lakehouse" "example" {
    description = "value"
}
```

## Validation Rules

- Maximum length: 256 characters


## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_lakehouse`.

## How To Fix

Update the `description` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/lakehouse/definitions.json)
