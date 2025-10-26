# fabric_gateway_invalid_inactivity_minutes_before_sleep

Validates that the `inactivity_minutes_before_sleep` attribute of `fabric_gateway` resources is valid according to the Fabric API specification.

## Example

```hcl
resource "fabric_gateway" "example" {
    inactivity_minutes_before_sleep = "30" # Valid
}
```

## Validation Rules

- Must be one of: `30`, `60`, `90`, `120`, `150`, `240`, `360`, `480`, `720`, `1440`


## Why

This rule ensures compliance with the Fabric REST API specification for `fabric_gateway`.

## How To Fix

Update the `inactivity_minutes_before_sleep` attribute to conform to the validation rules above.

## Reference

- [Fabric API Spec](https://github.com/microsoft/fabric-rest-api-specs/tree/main/platform/definitions/gateways.json)
