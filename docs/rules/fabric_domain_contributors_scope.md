# fabric_domain_contributors_scope

Validates that domain contributors_scope values are one of the supported scope options.

## Example

```hcl
resource "fabric_domain" "valid" {
  display_name        = "Sales Domain"
  contributors_scope  = "SpecificUsersAndGroups" # Valid scope
}

resource "fabric_domain" "invalid" {
  display_name        = "Marketing Domain"
  contributors_scope  = "Everyone" # Invalid - not a valid scope
}
```

## Why

Microsoft Fabric domains use the `contributors_scope` setting to control who can contribute items to the domain. Only specific scope values are supported by the Fabric API. Using an invalid scope will cause the domain creation or update to fail.

## Validation Rules

Must be one of:
- `AdminsOnly` - Only Fabric administrators can contribute items to this domain
- `AllTenant` - All users in the tenant can contribute items to this domain
- `SpecificUsersAndGroups` - Only specific users and groups can contribute items (requires additional role assignments)

## How to Fix

Change the `contributors_scope` to one of the supported values:

```hcl
# Example: Restrict contributions to admins only
resource "fabric_domain" "executive" {
  display_name        = "Executive Domain"
  description         = "Domain for executive-level reports and dashboards"
  contributors_scope  = "AdminsOnly"
}

# Example: Allow all tenant users to contribute
resource "fabric_domain" "collaborative" {
  display_name        = "Collaboration Domain"
  description         = "Open domain for cross-team collaboration"
  contributors_scope  = "AllTenant"
}

# Example: Restrict to specific users/groups
resource "fabric_domain" "sales" {
  display_name        = "Sales Domain"
  description         = "Domain for sales team analytics"
  contributors_scope  = "SpecificUsersAndGroups"
}

# When using SpecificUsersAndGroups, define role assignments
resource "fabric_domain_role_assignment" "sales_contributors" {
  domain_id    = fabric_domain.sales.id
  principal_id = "sales-team-group-id"
  role         = "Contributor"
}
```

## Scope Comparison

| Scope | Who Can Contribute | Use Case |
|-------|-------------------|----------|
| AdminsOnly | Fabric administrators only | Highly controlled domains with certified content |
| AllTenant | All tenant users | Collaborative environments, experimentation |
| SpecificUsersAndGroups | Designated users/groups | Team-specific domains with controlled access |

## Configuration

```hcl
rule "fabric_domain_contributors_scope" {
  enabled = true
}
```

## Attributes

| Name | Enabled | Severity | 
|------|---------|----------|
| fabric_domain_contributors_scope | true | error |
