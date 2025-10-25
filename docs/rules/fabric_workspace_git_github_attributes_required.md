# fabric_workspace_git_github_attributes_required

Validates that owner_name is provided when git_provider_type is "GitHub".

## Example

```hcl
# Valid - GitHub with owner_name
resource "fabric_workspace_git" "valid" {
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
    source        = "ConfiguredConnection"
    connection_id = "11111111-1111-1111-1111-111111111111"
  }
}

# Invalid - missing owner_name
resource "fabric_workspace_git" "invalid" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "GitHub"
    # Missing owner_name - will emit error
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

GitHub organizes repositories using a two-level hierarchy:
1. **Owner** - The user or organization that owns the repository
2. **Repository** - The actual Git repository

To uniquely identify a GitHub repository, Fabric requires the owner name in addition to the repository name. The owner can be:
- A **user account** (e.g., `octocat`)
- An **organization** (e.g., `microsoft`, `azure`)

Without the owner_name, Fabric cannot locate the correct repository since multiple users/organizations could have repositories with the same name.

## Validation Rules

When `git_provider_type = "GitHub"`:
- `owner_name` is **REQUIRED**

**Note:** This attribute should NOT be set when using Azure DevOps (Azure DevOps uses `organization_name` and `project_name` instead).

## How to Fix

Add `owner_name` to your git_provider_details block:

```hcl
resource "fabric_workspace_git" "example" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "GitHub"
    owner_name        = "mycompany"        # Your GitHub user or organization
    repository_name   = "fabric-config"
    branch_name       = "main"
    directory_name    = "/workspaces"
  }
  
  git_credentials {
    source        = "ConfiguredConnection"
    connection_id = fabric_connection.github.id
  }
}
```

## Finding Your GitHub Owner Name

Your GitHub repository URL reveals the required values:

```
https://github.com/{owner_name}/{repository_name}
```

**Examples:**

**Organization Repository:**
- URL: `https://github.com/microsoft/fabric-samples`
- `owner_name`: `microsoft`
- `repository_name`: `fabric-samples`

**Personal Repository:**
- URL: `https://github.com/octocat/hello-world`
- `owner_name`: `octocat`
- `repository_name`: `hello-world`

## Configuration

```hcl
rule "fabric_workspace_git_github_attributes_required" {
  enabled = true
}
```

## Attributes

| Name | Enabled | Severity | 
|------|---------|----------|
| fabric_workspace_git_github_attributes_required | true | error |
