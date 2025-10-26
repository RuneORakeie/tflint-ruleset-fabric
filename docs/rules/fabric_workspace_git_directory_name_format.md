# fabric_workspace_git_directory_name_format

Validates that the directory_name in git_provider_details starts with a forward slash (/) and does not exceed 256 characters.

## Example

```hcl
# Valid
resource "fabric_workspace_git" "valid" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "GitHub"
    owner_name        = "myorg"
    repository_name   = "myrepo"
    branch_name       = "main"
    directory_name    = "/terraform/fabric" # Valid - starts with /
  }
  
  git_credentials {
    source        = "ConfiguredConnection"
    connection_id = "11111111-1111-1111-1111-111111111111"
  }
}

# Invalid - missing leading slash
resource "fabric_workspace_git" "invalid_format" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "GitHub"
    owner_name        = "myorg"
    repository_name   = "myrepo"
    branch_name       = "main"
    directory_name    = "terraform/fabric" # Invalid - must start with /
  }
  
  git_credentials {
    source        = "ConfiguredConnection"
    connection_id = "11111111-1111-1111-1111-111111111111"
  }
}

# Invalid - too long
resource "fabric_workspace_git" "invalid_length" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "GitHub"
    owner_name        = "myorg"
    repository_name   = "myrepo"
    branch_name       = "main"
    directory_name    = "/very/long/path/that/exceeds/the/maximum/allowed/length/..." # Over 256 chars
  }
  
  git_credentials {
    source        = "ConfiguredConnection"
    connection_id = "11111111-1111-1111-1111-111111111111"
  }
}
```

## Why

Microsoft Fabric requires the directory_name to follow specific formatting rules:

1. **Must start with `/`**: The path must be an absolute path starting from the repository root
2. **Maximum 256 characters**: This is a platform limitation to ensure compatibility with file system and API constraints

Using an incorrectly formatted directory_name will cause the Git connection to fail.

## Validation Rules

- Must start with forward slash `/`
- Maximum length: 256 characters

## How to Fix

Ensure your directory_name starts with `/` and is not longer than 256 characters:

```hcl
resource "fabric_workspace_git" "example" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "GitHub"
    owner_name        = "myorg"
    repository_name   = "myrepo"
    branch_name       = "main"
    directory_name    = "/fabric/workspaces/prod"  # Starts with /, under 256 chars
  }
  
  git_credentials {
    source        = "ConfiguredConnection"
    connection_id = "11111111-1111-1111-1111-111111111111"
  }
}
```

**Valid examples:**
- `/` (root directory)
- `/terraform`
- `/infrastructure/fabric`
- `/workspaces/production/analytics`

**Invalid examples:**
- `terraform` (missing leading `/`)
- `./terraform` (relative path)
- `~/terraform` (home directory notation)

## Configuration

```hcl
rule "fabric_workspace_git_directory_name_format" {
  enabled = true
}
```

## Attributes

| Name | Enabled | Severity | 
|------|---------|----------|
| fabric_workspace_git_directory_name_format | true | error |
