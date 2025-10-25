# fabric_spark_custom_pool_invalid_node_family

Validates that the `node_family` attribute of `fabric_spark_custom_pool` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_spark_custom_pool" "example" {
    node_family = "MemoryOptimized" # Valid
}
```

## Validation Rules

- Must be one of: `MemoryOptimized`


## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_spark_custom_pool`.

## How To Fix

Update the `node_family` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/spark/definitions.json)
