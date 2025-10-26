# fabric_ml_experiment_invalid_display_name

Validates that the `display_name` attribute of `fabric_ml_experiment` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_ml_experiment" "example" {
    display_name = "value"
}
```

## Validation Rules



## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_ml_experiment`.

## How To Fix

Update the `display_name` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/mlExperiment/definitions.json)
