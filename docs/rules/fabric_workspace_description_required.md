# fabric_workspace_description_required

Ensures that workspaces have descriptions for governance and documentation purposes.

## Example

```hcl
resource "fabric_workspace" "valid" {
  display_name = "My Workspace"
  description  = "Production workspace for sales analytics" # Good - has description
  capacity_id  = fabric_capacity.example.id
}

resource "fabric_workspace" "invalid" {
  display_name = "My Workspace"
  capacity_id  = fabric_capacity.example.id
  # Missing description - will emit warning
}
```

## Why

Descriptions are essential for:
- **Governance**: Understanding the purpose and ownership of resources
- **Documentation**: Making it easier for teams to navigate the Fabric environment
- **Compliance**: Meeting organizational documentation requirements
- **Collaboration**: Helping new team members understand resource purposes

In large Fabric deployments with multiple workspaces, descriptions become critical for maintaining clarity and organization.

## How to Fix

Add a meaningful `description` attribute to your workspace configuration that explains the purpose, owner, and any relevant business context.

```hcl
resource "fabric_workspace" "my_workspace" {
  display_name = "Sales Analytics"
  description  = "Production workspace for sales team analytics and reporting. Owner: sales-team@company.com"
  capacity_id  = fabric_capacity.production.id
}
```

## Configuration

```hcl
rule "fabric_workspace_description_required" {
  enabled = true
}
```

## Attributes

| Name | Enabled | Severity | 
|------|---------|----------|
| fabric_workspace_description_required | true | warning |
