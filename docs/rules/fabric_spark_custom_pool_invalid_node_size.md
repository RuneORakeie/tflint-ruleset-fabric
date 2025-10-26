# fabric_spark_custom_pool_invalid_node_size

Validates that the `node_size` attribute of `fabric_spark_custom_pool` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_spark_custom_pool" "example" {
    node_size = "Small" # Valid
}
```

## Validation Rules

- Must be one of: `Small`, `Medium`, `Large`, `XLarge`, `XXLarge`


## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_spark_custom_pool`.

## How To Fix

Update the `node_size` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/spark/definitions.json)
