// Mapping for fabric_restore_point resource
// Auto-generated from warehouse/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_restore_point" {
  import_path = "warehouse/definitions.json"

  // optional, max 512 chars
  attribute "description" {
    api_ref = "CreateRestorePointRequest.description"
    max_length = 512
  }

  // optional, max 128 chars
  attribute "display_name" {
    api_ref = "CreateRestorePointRequest.displayName"
    max_length = 128
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
