plugin "fabric" {
  enabled = true
}

# Enable all rules
rule "fabric_workspace_capacity_required" { enabled = true }
rule "fabric_workspace_description_required" { enabled = true }
rule "fabric_workspace_naming" { enabled = true }
rule "fabric_role_assignment_principal_required" { enabled = true }
rule "fabric_git_integration_provider_valid" { enabled = true }

