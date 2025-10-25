# fabric_workspace_role_assignment_role

Validates that workspace role assignment role values are one of the supported workspace roles.

## Example

```hcl
resource "fabric_workspace_role_assignment" "valid" {
  workspace_id = fabric_workspace.example.id
  principal_id = "00000000-0000-0000-0000-000000000000"
  role         = "Contributor" # Valid role
}

resource "fabric_workspace_role_assignment" "invalid" {
  workspace_id = fabric_workspace.example.id
  principal_id = "00000000-0000-0000-0000-000000000000"
  role         = "Owner" # Invalid - not a valid workspace role
}
```

## Why

Microsoft Fabric workspaces support specific predefined roles, each with different permission levels. Using an invalid role will cause the role assignment to fail. This rule ensures you're using one of the supported workspace roles.

## Validation Rules

Must be one of:
- `Admin` - Full control including workspace settings, content creation/editing, and managing permissions
- `Contributor` - Can create and edit content, but cannot manage workspace settings or permissions
- `Member` - Can view and use content, create personal content (like reports from datasets)
- `Viewer` - Read-only access to workspace content

## How to Fix

Change the `role` to one of the supported values:

```hcl
# Example: Assign different roles to different groups
resource "fabric_workspace_role_assignment" "admins" {
  workspace_id = fabric_workspace.sales.id
  principal_id = "admin-group-id"
  role         = "Admin"
}

resource "fabric_workspace_role_assignment" "developers" {
  workspace_id = fabric_workspace.sales.id
  principal_id = "dev-group-id"
  role         = "Contributor"
}

resource "fabric_workspace_role_assignment" "analysts" {
  workspace_id = fabric_workspace.sales.id
  principal_id = "analyst-group-id"
  role         = "Member"
}

resource "fabric_workspace_role_assignment" "business_users" {
  workspace_id = fabric_workspace.sales.id
  principal_id = "users-group-id"
  role         = "Viewer"
}
```

## Role Comparison

| Role | Create/Edit Content | Manage Settings | Share | Manage Permissions |
|------|-------------------|----------------|-------|-------------------|
| Admin | ✓ | ✓ | ✓ | ✓ |
| Contributor | ✓ | ✗ | ✓ | ✗ |
| Member | Limited | ✗ | ✗ | ✗ |
| Viewer | ✗ | ✗ | ✗ | ✗ |

## Configuration

```hcl
rule "fabric_workspace_role_assignment_role" {
  enabled = true
}
```

## Attributes

| Name | Enabled | Severity | 
|------|---------|----------|
| fabric_workspace_role_assignment_role | true | error |
