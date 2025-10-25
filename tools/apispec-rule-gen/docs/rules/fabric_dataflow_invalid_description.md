# fabric_dataflow_invalid_description

Validates that the `description` attribute of `fabric_dataflow` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_dataflow" "example" {
    description = "value"
}
```

## Validation Rules



## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_dataflow`.

## How To Fix

Update the `description` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/dataflow/definitions.json)
