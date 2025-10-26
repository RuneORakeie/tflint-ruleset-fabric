# fabric_eventhouse_invalid_display_name

Validates that the `display_name` attribute of `fabric_eventhouse` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_eventhouse" "example" {
    display_name = "valid-value" # Must match pattern: ^[a-zA-Z0-9._-]+$
}
```

## Validation Rules

- Must match pattern: `^[a-zA-Z0-9._-]+$`
- Maximum length: 256 characters


## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_eventhouse`.

## How To Fix

Update the `display_name` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/eventhouse/definitions.json)
