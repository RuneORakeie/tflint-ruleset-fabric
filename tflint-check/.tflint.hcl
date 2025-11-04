# .tflint.hcl - TFLint Configuration for Fabric Ruleset

# Enable the Fabric ruleset plugin
plugin "fabric" {
  enabled = true
}

# ============================================
# Fabric Workspace Capacity Assignment
# ============================================
rule "fabric_workspace_capacity_required" {
  enabled = true
}

# ============================================
# Terraform Language Rules
# ============================================

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

# Module Configuration
config {
  # module = true              # Enable module inspection
  # force = false              # Continue on errors
  # disabled_by_default = false  # Enable all rules by default
  
  # Note: deep_check is deprecated and removed in newer TFLint versions
}
