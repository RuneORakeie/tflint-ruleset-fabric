# fabric_variable_library_invalid_description

Validates that the `description` attribute of `fabric_variable_library` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_variable_library" "example" {
    description = "value"
}
```

## Validation Rules

- Maximum length: 256 characters


## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_variable_library`.

## How To Fix

Update the `description` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/variableLibrary/definitions.json)
