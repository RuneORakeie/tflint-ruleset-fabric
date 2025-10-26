# fabric_git_integration_provider_valid

Validates that the Git provider type in the `git_provider_details` block is one of the supported providers for Microsoft Fabric workspace Git integration.

## Example

```hcl
# Valid - Azure DevOps
resource "fabric_workspace_git" "azdo" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "AzureDevOps" # Valid provider
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

# Valid - GitHub
resource "fabric_workspace_git" "github" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "GitHub" # Valid provider
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

# Invalid - GitLab not supported
resource "fabric_workspace_git" "invalid" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "GitLab" # Invalid - not supported
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

Microsoft Fabric currently supports only specific Git providers for workspace integration. Using an unsupported provider will cause the Git connection to fail. This rule ensures you're using one of the supported Git platforms.

## Validation Rules

Must be one of:
- `AzureDevOps` - Azure DevOps (Azure Repos)
- `GitHub` - GitHub

## How to Fix

Set the `git_provider_type` within the `git_provider_details` block to one of the supported values:

**For Azure DevOps:**
```hcl
resource "fabric_workspace_git" "example" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "AzureDevOps"
    organization_name = "myorg"      # Required for AzureDevOps
    project_name      = "myproject"  # Required for AzureDevOps
    repository_name   = "myrepo"
    branch_name       = "main"
    directory_name    = "/terraform"
  }
  
  git_credentials {
    source = "Automatic"  # or "ConfiguredConnection"
  }
}
```

**For GitHub:**
```hcl
resource "fabric_workspace_git" "example" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "GitHub"
    owner_name        = "myorg"  # Required for GitHub
    repository_name   = "myrepo"
    branch_name       = "main"
    directory_name    = "/terraform"
  }
  
  git_credentials {
    source        = "ConfiguredConnection"  # GitHub only supports ConfiguredConnection
    connection_id = "11111111-1111-1111-1111-111111111111"
  }
}
```

**Note:** 
- Azure DevOps requires `organization_name` and `project_name` attributes
- GitHub requires `owner_name` attribute
- GitHub only supports `ConfiguredConnection` for git credentials
- Azure DevOps supports both `Automatic` and `ConfiguredConnection` for git credentials

## Configuration

```hcl
rule "fabric_git_integration_provider_valid" {
  enabled = true
}
```

## Attributes

| Name | Enabled | Severity | 
|------|---------|----------|
| fabric_git_integration_provider_valid | true | error |
