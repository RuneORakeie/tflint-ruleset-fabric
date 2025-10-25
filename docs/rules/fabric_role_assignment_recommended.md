# fabric_role_assignment_recommended

Warns when resources are created without role assignments. Without role assignments, resources may only be accessible to admins or not accessible at all.

## Example

```hcl
# Good - workspace with role assignment
resource "fabric_workspace" "example" {
  display_name = "My Workspace"
  capacity_id  = fabric_capacity.example.id
}

resource "fabric_workspace_role_assignment" "team" {
  workspace_id = fabric_workspace.example.id
  principal_id = "00000000-0000-0000-0000-000000000000"
  role         = "Contributor"
}

# Warning - workspace without role assignment
resource "fabric_workspace" "no_roles" {
  display_name = "No Access Workspace"
  capacity_id  = fabric_capacity.example.id
  # No role assignments defined - will emit warning
}
```

## Why

Role assignments are critical for resource access control in Microsoft Fabric. Without explicit role assignments:

- **Limited Access**: Only tenant admins may be able to access the resource
- **Collaboration Issues**: Team members won't be able to work with the resource
- **Security Concerns**: Access may not be properly governed
- **Operational Problems**: May cause deployment or runtime access issues

This rule helps ensure that access is properly configured from the start.

## Applies To

This rule checks the following resource types:
- `fabric_workspace` (looks for `fabric_workspace_role_assignment`)
- `fabric_deployment_pipeline` (looks for `fabric_deployment_pipeline_role_assignment`)
- `fabric_domain` (looks for `fabric_domain_role_assignment`)
- `fabric_gateway` (looks for `fabric_gateway_role_assignment`)

## How to Fix

Create corresponding role assignment resources for your Fabric resources:

```hcl
resource "fabric_workspace" "my_workspace" {
  display_name = "My Workspace"
  capacity_id  = fabric_capacity.production.id
}

# Add role assignments
resource "fabric_workspace_role_assignment" "developers" {
  workspace_id = fabric_workspace.my_workspace.id
  principal_id = "dev-team-group-id"
  role         = "Contributor"
}

resource "fabric_workspace_role_assignment" "viewers" {
  workspace_id = fabric_workspace.my_workspace.id
  principal_id = "viewers-group-id"
  role         = "Viewer"
}
```

## Configuration

```hcl
rule "fabric_role_assignment_recommended" {
  enabled = true
}
```

## Attributes

| Name | Enabled | Severity | 
|------|---------|----------|
| fabric_role_assignment_recommended | true | warning |
