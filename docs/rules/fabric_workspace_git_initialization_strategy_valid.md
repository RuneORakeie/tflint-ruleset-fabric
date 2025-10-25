# fabric_workspace_git_initialization_strategy_valid

Validates that the initialization_strategy is one of the supported values for Microsoft Fabric workspace Git integration.

## Example

```hcl
# Valid - PreferWorkspace
resource "fabric_workspace_git" "valid_workspace" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace" # Valid
  
  git_provider_details {
    git_provider_type = "GitHub"
    owner_name        = "myorg"
    repository_name   = "myrepo"
    branch_name       = "main"
    directory_name    = "/terraform"
  }
  
  git_credentials {
    source        = "ConfiguredConnection"
    connection_id = "11111111-1111-1111-1111-111111111111"
  }
}

# Valid - PreferRemote
resource "fabric_workspace_git" "valid_remote" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferRemote" # Valid
  
  git_provider_details {
    git_provider_type = "AzureDevOps"
    organization_name = "myorg"
    project_name      = "myproject"
    repository_name   = "myrepo"
    branch_name       = "main"
    directory_name    = "/terraform"
  }
  
  git_credentials {
    source = "Automatic"
  }
}

# Invalid
resource "fabric_workspace_git" "invalid" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferGit" # Invalid - not supported
  
  git_provider_details {
    git_provider_type = "GitHub"
    owner_name        = "myorg"
    repository_name   = "myrepo"
    branch_name       = "main"
    directory_name    = "/terraform"
  }
  
  git_credentials {
    source        = "ConfiguredConnection"
    connection_id = "11111111-1111-1111-1111-111111111111"
  }
}
```

## Why

The initialization_strategy determines how conflicts are resolved when connecting a workspace to Git:

- **PreferRemote**: Items in the remote Git repository take precedence. Workspace items that conflict will be overwritten.
- **PreferWorkspace**: Items in the workspace take precedence. Git repository items that conflict will be overwritten.

Using an invalid strategy will cause the Git connection to fail. This rule ensures you're using one of the supported initialization strategies.

## Validation Rules

Must be one of:
- `PreferRemote` - Give precedence to items in the Git repository
- `PreferWorkspace` - Give precedence to items in the workspace

## How to Fix

Change the `initialization_strategy` to one of the supported values:

```hcl
resource "fabric_workspace_git" "example" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"  # or "PreferRemote"
  
  git_provider_details {
    git_provider_type = "GitHub"
    owner_name        = "myorg"
    repository_name   = "myrepo"
    branch_name       = "main"
    directory_name    = "/terraform"
  }
  
  git_credentials {
    source        = "ConfiguredConnection"
    connection_id = "11111111-1111-1111-1111-111111111111"
  }
}
```

## Configuration

```hcl
rule "fabric_workspace_git_initialization_strategy_valid" {
  enabled = true
}
```

## Attributes

| Name | Enabled | Severity | 
|------|---------|----------|
| fabric_workspace_git_initialization_strategy_valid | true | error |
