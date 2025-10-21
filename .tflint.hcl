# .tflint.hcl - TFLint Configuration for Fabric Ruleset

# Enable the Fabric ruleset plugin
plugin "fabric" {
  enabled = true
  version = "0.1.0"
  source  = "github.com/RuneORakeie/tflint-ruleset-fabric"
}

# ============================================
# Fabric Workspace Naming Convention
# ============================================
rule "fabric_workspace_naming" {
  enabled = true
  
  # Optional configuration for custom naming patterns
  # pattern = "^[a-z0-9]{3,50}$"
}

# ============================================
# Fabric Workspace Capacity Assignment
# ============================================
rule "fabric_workspace_capacity_required" {
  enabled = true
}

# ============================================
# Fabric Workspace Description
# ============================================
rule "fabric_workspace_description_required" {
  enabled = true
}

# ============================================
# Fabric Role Assignment Principal
# ============================================
rule "fabric_role_assignment_principal_required" {
  enabled = true
}

# ============================================
# Fabric Git Integration Provider Validation
# ============================================
rule "fabric_git_integration_provider_valid" {
  enabled = true
  
  # Supported providers: GitHub, Azure DevOps, Bitbucket Cloud, GitLab
}

# ============================================
# Fabric Capacity Region Validation
# ============================================
rule "fabric_capacity_region_valid" {
  enabled = false  # Disabled by default, enable if needed for region validation
  
  # When enabled, validates regions against available Azure regions
}

# ============================================
# Terraform Language Rules
# ============================================
rule "terraform_required_providers" {
  enabled = true
}

rule "terraform_required_version" {
  enabled = true
}

rule "terraform_naming_convention" {
  enabled = true
  format  = "snake_case"
}

rule "terraform_typed_variables" {
  enabled = true
}

rule "terraform_documented_variables" {
  enabled = true
}

rule "terraform_documented_outputs" {
  enabled = true
}

rule "terraform_standard_module_structure" {
  enabled = true
}

# ============================================
# Global TFLint Configuration
# ============================================

# Severity level for filtering results
# Can be: error, warning, notice
# minimum_failure_severity = "warning"

# Format for output
# Can be: default, json, checkstyle, junit, sarif
format = "default"

# Module Configuration - if using modules
# module = true
# deep_check = false