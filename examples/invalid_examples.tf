# examples/invalid/main.tf
# Invalid Fabric Terraform Configuration Examples - These will trigger TFLint rules

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
# INVALID: Multiple violations in single resource
# ============================================
resource "fabric_workspace" "multiple_issues" {
  display_name = "WS123!@#"  # WARNING: Invalid naming convention
  # description missing - WARNING
  # capacity_id missing - ERROR
  
}

# ============================================
# INVALID: Capacity with unsupported region (WARNING)
# Triggers: fabric_capacity_region_valid (when enabled)
# ============================================
resource "fabric_capacity" "invalid_region" {
  name       = "capacity-invalid-region"
  location             = "mars-1"  # Not a real Azure region
  sku          {
    name = "F16"
    tier = "Fabric"
  }
  administration_members   = ["user@company.com"]
}

# ============================================
# INVALID: Multiple missing fields
# ============================================
resource "fabric_workspace" "minimal_broken" {
  display_name = "X"  # Too short - WARNING
  # No description - WARNING
  # No capacity_id - ERROR
}

resource "fabric_workspace_role_assignment" "incomplete" {
  workspace_id = "workspace-123"
  role         = "Editor"
  # No principal_id - ERROR
  # No principal_type - Could be ERROR depending on provider requirements
}

# ============================================
# INVALID: Git provider variations that fail
# ============================================
resource "fabric_workspace_git_connection" "bad_provider_1" {
  workspace_id       = "workspace-id"
  git_provider_type  = "git_hub"  # Wrong format
  repository_name    = "repo"
  branch_name        = "main"
  organization_name  = "org"
}

resource "fabric_workspace_git_connection" "bad_provider_2" {
  workspace_id       = "workspace-id"
  git_provider_type  = "Bitbucket"  # Should be "Bitbucket Cloud"
  repository_name    = "repo"
  branch_name        = "main"
  organization_name  = "org"
} Workspace without capacity (ERROR)
# Triggers: fabric_workspace_capacity_required
# ============================================
resource "fabric_workspace" "no_capacity_workspace" {
  display_name = "workspace-without-capacity"
  description  = "This workspace has no capacity assigned"
  # capacity_id is missing - ERROR
}

# ============================================
# INVALID: Workspace naming violation (WARNING)
# Triggers: fabric_workspace_naming
# ============================================
resource "fabric_workspace" "BAD_NAMING" {
  display_name = "BAD-NAMING-WITH-CAPS"  # Upper case not allowed
  description  = "Workspace with invalid naming"
  capacity_id  = "capacity-id-here"
}

resource "fabric_workspace" "invalid_name_special_chars" {
  display_name = "workspace@with#special$chars"  # Special characters not allowed
  description  = "Invalid special characters"
  capacity_id  = "capacity-id-here"
}

resource "fabric_workspace" "short_name" {
  display_name = "ab"  # Too short, minimum is 3 characters
  description  = "Name too short"
  capacity_id  = "capacity-id-here"
}

resource "fabric_workspace" "very_long_name" {
  display_name = "this-is-a-very-long-workspace-name-that-exceeds-the-fifty-character-limit"  # Too long
  description  = "Name exceeds 50 characters"
  capacity_id  = "capacity-id-here"
}

# ============================================
# INVALID: Workspace without description (WARNING)
# Triggers: fabric_workspace_description_required
# ============================================
resource "fabric_workspace" "no_description" {
  display_name = "workspace-no-description"
  # description is missing - WARNING
  capacity_id  = "capacity-id-here"
}

# ============================================
# INVALID: Role assignment without principal (ERROR)
# Triggers: fabric_role_assignment_principal_required
# ============================================
resource "fabric_workspace_role_assignment" "missing_principal" {
  workspace_id = "workspace-id"
  # principal is missing - ERROR
  role         = "Admin"
}

# ============================================
# INVALID: Git integration with invalid provider (ERROR)
# Triggers: fabric_git_integration_provider_valid
# ============================================
resource "fabric_workspace_git_connection" "invalid_provider" {
  workspace_id       = "workspace-id"
  git_provider_type  = "Subversion"  # Not a supported provider
  repository_name    = "my-repo"
  branch_name        = "main"
  organization_name  = "my-org"
}

resource "fabric_workspace_git_connection" "typo_provider" {
  workspace_id       = "workspace-id"
  git_provider_type  = "GitHUB"  # Should be "GitHub" - case sensitive
  repository_name    = "my-repo"
  branch_name        = "main"
  organization_name  = "my-org"
}

