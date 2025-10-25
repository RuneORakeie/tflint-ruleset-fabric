# fabric_deployment_pipeline_invalid_description

Validates that the `description` attribute of `fabric_deployment_pipeline` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_deployment_pipeline" "example" {
    description = "value"
}
```

## Validation Rules



## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_deployment_pipeline`.

## How To Fix

Update the `description` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/platform/definitions/deploymentPipelines.json)
