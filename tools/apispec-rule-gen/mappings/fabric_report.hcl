// Mapping for fabric_report resource
// Auto-generated from report/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_report" {
  import_path = "report/definitions.json"

  // required
  attribute "definition" {
    api_ref = "CreateReportRequest.definition"
  }

  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreateReportRequest.description"
    max_length = 256
  }

  // required
  attribute "display_name" {
    api_ref = "CreateReportRequest.displayName"
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateReportRequest.folderId"
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
