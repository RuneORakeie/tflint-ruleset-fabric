# fabric_domain_invalid_parent_domain_id

Validates that the `parent_domain_id` attribute of `fabric_domain` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_domain" "example" {
    parent_domain_id = "valid-uuid" # Must be valid uuid
}
```

## Validation Rules

- Must be a valid uuid format


## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_domain`.

## How To Fix

Update the `parent_domain_id` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/admin/definitions/domains.json)
