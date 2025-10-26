// Mapping for fabric_environment resource
// Auto-generated from environment/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_environment" {
  import_path = "environment/definitions.json"

  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreateEnvironmentRequest.description"
    max_length = 256
  }

  // required
  attribute "display_name" {
    api_ref = "CreateEnvironmentRequest.displayName"
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateEnvironmentRequest.folderId"
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
