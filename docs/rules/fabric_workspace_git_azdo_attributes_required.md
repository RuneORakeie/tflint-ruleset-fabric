# fabric_workspace_git_azdo_attributes_required

Validates that organization_name and project_name are provided when git_provider_type is "AzureDevOps".

## Example

```hcl
# Valid - AzureDevOps with all required attributes
resource "fabric_workspace_git" "valid" {
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
    source = "Automatic"
  }
}

# Invalid - missing organization_name
resource "fabric_workspace_git" "missing_org" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "AzureDevOps"
    # Missing organization_name - will emit error
    project_name      = "myproject"
    repository_name   = "myrepo"
    branch_name       = "main"
    directory_name    = "/terraform"
  }
  
  git_credentials {
    source = "Automatic"
  }
}

# Invalid - missing project_name
resource "fabric_workspace_git" "missing_project" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "AzureDevOps"
    organization_name = "myorg"
    # Missing project_name - will emit error
    repository_name   = "myrepo"
    branch_name       = "main"
    directory_name    = "/terraform"
  }
  
  git_credentials {
    source = "Automatic"
  }
}
```

## Why

Azure DevOps organizes repositories in a three-level hierarchy:
1. **Organization** - The top-level container (e.g., `dev.azure.com/mycompany`)
2. **Project** - Contains multiple repositories, work items, and pipelines
3. **Repository** - The actual Git repository

To uniquely identify an Azure DevOps repository, Fabric requires both the organization name and project name in addition to the repository name. Without these, Fabric cannot locate the correct repository.

## Validation Rules

When `git_provider_type = "AzureDevOps"`:
- `organization_name` is **REQUIRED**
- `project_name` is **REQUIRED**

**Note:** These attributes should NOT be set when using GitHub (GitHub uses `owner_name` instead).

## How to Fix

Add both `organization_name` and `project_name` to your git_provider_details block:

```hcl
resource "fabric_workspace_git" "example" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "AzureDevOps"
    organization_name = "mycompany"          # Your Azure DevOps organization
    project_name      = "infrastructure"     # The project containing your repo
    repository_name   = "fabric-config"
    branch_name       = "main"
    directory_name    = "/workspaces"
  }
  
  git_credentials {
    source = "Automatic"  # or "ConfiguredConnection"
  }
}
```

## Finding Your Azure DevOps Values

Your Azure DevOps URL structure reveals the required values:

```
https://dev.azure.com/{organization_name}/{project_name}/_git/{repository_name}
```

**Example:**
- URL: `https://dev.azure.com/contoso/DataPlatform/_git/fabric-terraform`
- `organization_name`: `contoso`
- `project_name`: `DataPlatform`
- `repository_name`: `fabric-terraform`

## Configuration

```hcl
rule "fabric_workspace_git_azdo_attributes_required" {
  enabled = true
}
```

## Attributes

| Name | Enabled | Severity | 
|------|---------|----------|
| fabric_workspace_git_azdo_attributes_required | true | error |
