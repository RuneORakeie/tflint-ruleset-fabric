// Mapping for fabric_dashboard resource
// Auto-generated from dashboard/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_dashboard" {
  import_path = "dashboard/definitions.json"

  // optional
  attribute "definition_manual_manual" {
    api_ref = "manual.definition_manual"
  }

  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreateDashboardRequest.description"
    max_length = 256
  }

  // required
  attribute "display_name" {
    api_ref = "CreateDashboardRequest.displayName"
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateDashboardRequest.folderId"
  }

  // Add manual customizations below with // MANUAL: comment
  // Example:
  // // MANUAL: custom constraint
  // attribute "display_name" {
  //   api_ref = "CreateXxxRequest.displayName"
  //   max_length = 256
  //   pattern = "^[a-zA-Z0-9_]+$"
  //   warn_on_exceed = true
  // }
}
