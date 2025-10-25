# fabric_eventhouse_invalid_description

Validates that the `description` attribute of `fabric_eventhouse` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_eventhouse" "example" {
    description = "value"
}
```

## Validation Rules

- Maximum length: 1024 characters


## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_eventhouse`.

## How To Fix

Update the `description` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/eventhouse/definitions.json)
