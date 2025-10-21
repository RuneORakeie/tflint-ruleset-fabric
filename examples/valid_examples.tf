# examples/valid/main.tf
# Valid Fabric Terraform Configuration Examples

terraform {
  required_providers {
    fabric = {
      source  = "microsoft/fabric"
      version = "~> 1"
    }
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.0"
    }
  }
}

# ============================================
# VALID: Workspace with proper configuration
# ============================================
resource "fabric_workspace" "prod_workspace" {
  display_name = "prod-analytics-platform"
  description  = "Production analytics workspace for financial reporting and KPI tracking"
  capacity_id  = fabric_capacity.prod.id
}

# ============================================
# VALID: Workspace for development
# ============================================
resource "fabric_workspace" "dev_workspace" {
  display_name = "dev-analytics-workspace"
  description  = "Development environment for analytics development and testing"
  capacity_id  = fabric_capacity.dev.id
}

# ============================================
# VALID: Capacity configuration
# ============================================
resource "azurerm_fabric_capacity" "prod" {
  name                   = "prodcapacity"
  location               = "eastus"
  resource_group_name    = "rg-fabric-prod"
  administration_members = ["user@company.com"]

  sku {
    name = "F64"
    tier = "Fabric"
  }

  tags = {
    environment = "production"
  }
}

resource "azurerm_fabric_capacity" "dev" {
  name                   = "devcapacity"
  location               = "eastus"
  administration_members = ["dev-team@company.com"]
  resource_group_name    = "rg-fabric-dev"

  sku {
    name = "Trial"
    tier = "Fabric"
  }

  tags = {
    environment = "development"
  }
}

# ============================================
# VALID: Role assignments with principal_id
# ============================================
resource "fabric_workspace_role_assignment" "admin_user" {
  workspace_id = fabric_workspace.prod_workspace.id
  principal = {
    id   = "00000000-0000-0000-0000-000000000001"
    type = "User"
  }
  role = "Admin"
}

resource "fabric_workspace_role_assignment" "editor_group" {
  workspace_id = fabric_workspace.prod_workspace.id
  principal = {
    id   = "00000000-0000-0000-0000-000000000002"
    type = "Group"
  }
  role = "Contributor"
}

resource "fabric_workspace_role_assignment" "viewer_group" {
  workspace_id = fabric_workspace.prod_workspace.id
  principal = {
    id   = "00000000-0000-0000-0000-000000000003"
    type = "Group"
  }
  role = "Viewer"
}

# ============================================
# VALID: Git integration with supported provider
# ============================================
resource "fabric_workspace_git_connection" "github_prod" {
  workspace_id      = fabric_workspace.prod_workspace.id
  git_provider_type = "GitHub"
  repository_name   = "fabric-analytics-repo"
  branch_name       = "main"
  organization_name = "company-org"
}

resource "fabric_workspace_git_connection" "azure_devops" {
  workspace_id      = fabric_workspace.dev_workspace.id
  git_provider_type = "Azure DevOps"
  repository_name   = "fabric-dev-repo"
  branch_name       = "develop"
  organization_name = "company-devops"
}

# ============================================
# VALID: Workspace with all recommended settings
# ============================================
resource "fabric_workspace" "complete_example" {
  display_name = "analytics-hub-central"
  description  = "Central analytics hub for enterprise reporting and dashboards"
  capacity_id  = fabric_capacity.prod.id
}

# Assign roles for complete example
resource "fabric_workspace_role_assignment" "complete_admin" {
  workspace_id = fabric_workspace.complete_example.id
  principal = {
    id   = "00000000-0000-0000-0000-000000000001"
    type = "User"
  }
  role = "Admin"
}

resource "fabric_workspace_role_assignment" "complete_editors" {
  workspace_id = fabric_workspace.complete_example.id
  principal = {
    id   = "00000000-0000-0000-0000-000000000004"
    type = "Group"
  }
  role = "Contributor"
}

# Connect to Git for CI/CD
resource "fabric_workspace_git_connection" "complete_git" {
  workspace_id      = fabric_workspace.complete_example.id
  git_provider_type = "GitHub"
  repository_name   = "fabric-enterprise-analytics"
  branch_name       = "main"
  organization_name = "enterprise-org"
}
