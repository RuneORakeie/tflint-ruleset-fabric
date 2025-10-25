// Mapping for fabric_item resource
// Auto-generated from platform/definitions/platform.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_item" {
  import_path = "platform/definitions/platform.json"

  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreateItemRequest.description"
    max_length = 256
  }

  // required
  attribute "display_name" {
    api_ref = "CreateItemRequest.displayName"
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateItemRequest.folderId"
  }

  // required
  attribute "type" {
    api_ref = "CreateItemRequest.type"
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
