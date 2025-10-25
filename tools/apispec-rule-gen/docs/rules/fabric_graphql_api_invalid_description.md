# fabric_graphql_api_invalid_description

Validates that the `description` attribute of `fabric_graphql_api` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_graphql_api" "example" {
    description = "value"
}
```

## Validation Rules



## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_graphql_api`.

## How To Fix

Update the `description` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/graphQLApi/definitions.json)
