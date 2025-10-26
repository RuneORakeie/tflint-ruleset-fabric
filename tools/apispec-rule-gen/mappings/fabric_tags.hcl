// Mapping for fabric_tags resource
// Auto-generated from admin/definitions/tags.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_tags" {
  import_path = "admin/definitions/tags.json"

  // required
  attribute "create_tags_request" {
    api_ref = "CreateTagsRequest.createTagsRequest"
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
