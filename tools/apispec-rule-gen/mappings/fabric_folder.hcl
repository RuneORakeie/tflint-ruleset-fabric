// Mapping for fabric_folder resource
// Auto-generated from platform/definitions/platform.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_folder" {
  import_path = "platform/definitions/platform.json"

  // required, max 255 chars
  attribute "display_name" {
    api_ref = "CreateFolderRequest.displayName"
    max_length = 255
  }

  // optional, format: uuid
  attribute "parent_folder_id" {
    api_ref = "CreateFolderRequest.parentFolderId"
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
