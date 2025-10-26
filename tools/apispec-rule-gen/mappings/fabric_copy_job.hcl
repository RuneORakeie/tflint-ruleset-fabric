// Mapping for fabric_copy_job resource
// Auto-generated from copyJob/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_copy_job" {
  import_path = "copyJob/definitions.json"


  // optional, max 1021 chars
  attribute "description" {
    api_ref = "CreateCopyJobRequest.description"
    max_length = 1021
  }

  // required, max 256 chars
  attribute "display_name" {
    api_ref = "CreateCopyJobRequest.displayName"
    max_length = 256
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateCopyJobRequest.folderId"
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
