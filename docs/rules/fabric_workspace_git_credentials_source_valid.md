# fabric_workspace_git_credentials_source_valid

Validates that the git_credentials.source value is appropriate for the selected Git provider type.

## Example

```hcl
# Valid - GitHub with ConfiguredConnection
resource "fabric_workspace_git" "github_valid" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "GitHub"
    owner_name        = "myorg"
    repository_name   = "myrepo"
    branch_name       = "main"
    directory_name    = "/terraform"
  }
  
  git_credentials {
    source        = "ConfiguredConnection"  # Valid for GitHub
    connection_id = "11111111-1111-1111-1111-111111111111"
  }
}

# Valid - AzureDevOps with Automatic
resource "fabric_workspace_git" "azdo_automatic" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "AzureDevOps"
    organization_name = "myorg"
    project_name      = "myproject"
    repository_name   = "myrepo"
    branch_name       = "main"
    directory_name    = "/terraform"
  }
  
  git_credentials {
    source = "Automatic"  # Valid for AzureDevOps (uses user identity)
  }
}

# Valid - AzureDevOps with ConfiguredConnection
resource "fabric_workspace_git" "azdo_configured" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "AzureDevOps"
    organization_name = "myorg"
    project_name      = "myproject"
    repository_name   = "myrepo"
    branch_name       = "main"
    directory_name    = "/terraform"
  }
  
  git_credentials {
    source        = "ConfiguredConnection"  # Valid for AzureDevOps
    connection_id = "11111111-1111-1111-1111-111111111111"
  }
}

# Invalid - GitHub with Automatic (not supported)
resource "fabric_workspace_git" "github_invalid" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "GitHub"
    owner_name        = "myorg"
    repository_name   = "myrepo"
    branch_name       = "main"
    directory_name    = "/terraform"
  }
  
  git_credentials {
    source = "Automatic"  # Invalid - GitHub doesn't support Automatic
  }
}
```

## Why

Different Git providers support different credential sources:

**GitHub:**
- Only supports `ConfiguredConnection` (requires a configured connection resource)
- Does not support `Automatic` (user identity passthrough)

**Azure DevOps:**
- Supports `ConfiguredConnection` (requires a configured connection resource)
- Supports `Automatic` (uses the signed-in user's identity)

Using an unsupported credential source will cause the Git connection to fail.

## Validation Rules

**For GitHub:**
- `source` must be `ConfiguredConnection`

**For Azure DevOps:**
- `source` must be `ConfiguredConnection` OR `Automatic`

## How to Fix

**For GitHub**, use ConfiguredConnection:
```hcl
resource "fabric_workspace_git" "github_example" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "GitHub"
    owner_name        = "myorg"
    repository_name   = "myrepo"
    branch_name       = "main"
    directory_name    = "/terraform"
  }
  
  git_credentials {
    source        = "ConfiguredConnection"
    connection_id = fabric_connection.github.id
  }
}
```

**For Azure DevOps**, use either ConfiguredConnection or Automatic:
```hcl
# Option 1: Automatic (user identity)
resource "fabric_workspace_git" "azdo_automatic" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "AzureDevOps"
    organization_name = "myorg"
    project_name      = "myproject"
    repository_name   = "myrepo"
    branch_name       = "main"
    directory_name    = "/terraform"
  }
  
  git_credentials {
    source = "Automatic"  # Uses signed-in user's identity
  }
}

# Option 2: ConfiguredConnection (service principal or managed identity)
resource "fabric_workspace_git" "azdo_configured" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "AzureDevOps"
    organization_name = "myorg"
    project_name      = "myproject"
    repository_name   = "myrepo"
    branch_name       = "main"
    directory_name    = "/terraform"
  }
  
  git_credentials {
    source        = "ConfiguredConnection"
    connection_id = fabric_connection.azdo.id
  }
}
```

## Important Notes

- **Service Principal Authentication**: Only supported when using `ConfiguredConnection`
- **User Identity**: `Automatic` only works with user authentication, not service principals
- **GitHub Limitation**: GitHub requires a configured connection and does not support automatic user identity passthrough

## Configuration

```hcl
rule "fabric_workspace_git_credentials_source_valid" {
  enabled = true
}
```

## Attributes

| Name | Enabled | Severity | 
|------|---------|----------|
| fabric_workspace_git_credentials_source_valid | true | error |
