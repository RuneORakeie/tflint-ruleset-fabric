// Mapping for fabric_lakehouse resource
// Auto-generated from lakehouse/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_lakehouse" {
  import_path = "lakehouse/definitions.json"

  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreateLakehouseRequest.description"
    max_length = 256
  }

  // required, max 123 chars
  attribute "display_name" {
    api_ref = "CreateLakehouseRequest.displayName"
    max_length = 123
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateLakehouseRequest.folderId"
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
