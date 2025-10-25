// Mapping for fabric_activator resource
// Auto-generated from reflex/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_activator" {
  import_path = "reflex/definitions.json"

  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreateReflexRequest.description"
    max_length = 256
  }

  // required
  attribute "display_name" {
    api_ref = "CreateReflexRequest.displayName"
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateReflexRequest.folderId"
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
