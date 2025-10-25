# fabric_workspace_capacity_required

Ensures that workspaces have a capacity assigned for production use.

## Example

```hcl
resource "fabric_workspace" "valid" {
  display_name = "My Workspace"
  capacity_id  = fabric_capacity.example.id # Good - capacity assigned
}

resource "fabric_workspace" "invalid" {
  display_name = "My Workspace"
  # Missing capacity_id - will emit error
}
```

## Why

Microsoft Fabric workspaces should have a capacity assigned for production use. Without a capacity assignment:
- The workspace will use shared capacity (trial mode)
- Performance may be limited
- Some features may be unavailable
- It's not suitable for production workloads

Assigning a capacity ensures predictable performance and access to all Fabric features.

## How to Fix

Add the `capacity_id` attribute to your workspace configuration, referencing a valid Fabric capacity resource.

```hcl
resource "fabric_capacity" "production" {
  display_name = "Production Capacity"
  region       = "eastus"
  sku          = "F64"
}

resource "fabric_workspace" "my_workspace" {
  display_name = "My Workspace"
  capacity_id  = fabric_capacity.production.id
}
```

## Configuration

```hcl
rule "fabric_workspace_capacity_required" {
  enabled = true
}
```

## Attributes

| Name | Enabled | Severity | 
|------|---------|----------|
| fabric_workspace_capacity_required | true | error |
