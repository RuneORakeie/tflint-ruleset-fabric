// Mapping for fabric_mirrored_database resource
// Auto-generated from mirroredDatabase/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_mirrored_database" {
  import_path = "mirroredDatabase/definitions.json"

  // required
  attribute "definition" {
    api_ref = "CreateMirroredDatabaseRequest.definition"
  }

  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreateMirroredDatabaseRequest.description"
    max_length = 256
  }

  // required
  attribute "display_name" {
    api_ref = "CreateMirroredDatabaseRequest.displayName"
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateMirroredDatabaseRequest.folderId"
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
