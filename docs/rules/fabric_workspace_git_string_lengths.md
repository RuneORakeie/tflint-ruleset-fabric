# fabric_workspace_git_string_lengths

Validates that string attributes in git_provider_details do not exceed their maximum allowed lengths.

## Example

```hcl
# Valid - all strings within limits
resource "fabric_workspace_git" "valid" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "GitHub"
    owner_name        = "myorganization"           # 14 chars (max 100)
    repository_name   = "fabric-terraform-config"  # 24 chars (max 128)
    branch_name       = "main"                      # 4 chars (max 250)
    directory_name    = "/infrastructure/fabric"    # 23 chars (max 256)
  }
  
  git_credentials {
    source        = "ConfiguredConnection"
    connection_id = "11111111-1111-1111-1111-111111111111"
  }
}

# Invalid - branch_name too long
resource "fabric_workspace_git" "invalid_branch" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "GitHub"
    owner_name        = "myorg"
    repository_name   = "myrepo"
    branch_name       = "feature/very-long-branch-name-that-exceeds-the-maximum-allowed-length-of-250-characters-and-continues-for-a-very-long-time-with-lots-of-descriptive-text-that-makes-it-exceed-the-limit-set-by-microsoft-fabric-for-branch-names-in-git-integration..."  # Over 250 chars - will emit error
    directory_name    = "/terraform"
  }
  
  git_credentials {
    source        = "ConfiguredConnection"
    connection_id = "11111111-1111-1111-1111-111111111111"
  }
}

# Invalid - repository_name too long
resource "fabric_workspace_git" "invalid_repo" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "AzureDevOps"
    organization_name = "myorg"
    project_name      = "myproject"
    repository_name   = "this-is-an-extremely-long-repository-name-that-exceeds-the-128-character-limit-imposed-by-microsoft-fabric-for-azure-devops-repositories"  # Over 128 chars - will emit error
    branch_name       = "main"
    directory_name    = "/terraform"
  }
  
  git_credentials {
    source = "Automatic"
  }
}
```

## Why

Microsoft Fabric enforces maximum string lengths for Git provider details to ensure:

- **API Compatibility**: Backend systems have database column size limitations
- **UI Rendering**: Names must display properly in the Fabric portal
- **Performance**: Prevents excessive data transfer and storage
- **Git Provider Limits**: Aligns with constraints from GitHub and Azure DevOps

Exceeding these limits will cause the Git connection creation or update to fail.

## Validation Rules

The following maximum lengths are enforced:

| Attribute | Maximum Length | Applies To |
|-----------|---------------|------------|
| `branch_name` | 250 characters | All providers |
| `repository_name` | 128 characters | All providers |
| `organization_name` | 100 characters | Azure DevOps only |
| `owner_name` | 100 characters | GitHub only |
| `project_name` | 100 characters | Azure DevOps only |

**Note:** `directory_name` is validated separately by the `fabric_workspace_git_directory_name_format` rule (max 256 characters).

## How to Fix

Shorten the attribute values to be within the allowed limits:

**For GitHub:**
```hcl
resource "fabric_workspace_git" "github_example" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "GitHub"
    owner_name        = "contoso"              # Under 100 chars
    repository_name   = "fabric-config"        # Under 128 chars
    branch_name       = "main"                 # Under 250 chars
    directory_name    = "/terraform/prod"      # Under 256 chars
  }
  
  git_credentials {
    source        = "ConfiguredConnection"
    connection_id = fabric_connection.github.id
  }
}
```

**For Azure DevOps:**
```hcl
resource "fabric_workspace_git" "azdo_example" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "AzureDevOps"
    organization_name = "contoso"              # Under 100 chars
    project_name      = "DataPlatform"         # Under 100 chars
    repository_name   = "fabric-terraform"     # Under 128 chars
    branch_name       = "feature/workspace-v2" # Under 250 chars
    directory_name    = "/infra/prod"          # Under 256 chars
  }
  
  git_credentials {
    source = "Automatic"
  }
}
```

## Best Practices

- **Keep names concise**: Use abbreviations where appropriate
- **Avoid redundancy**: Don't repeat the project/organization name in the repository name
- **Use conventions**: Follow your team's naming standards for consistency
- **Branch naming**: Use conventional prefixes like `feature/`, `bugfix/`, `release/` but keep the full name under 250 chars

## Configuration

```hcl
rule "fabric_workspace_git_string_lengths" {
  enabled = true
}
```

## Attributes

| Name | Enabled | Severity | 
|------|---------|----------|
| fabric_workspace_git_string_lengths | true | error |
