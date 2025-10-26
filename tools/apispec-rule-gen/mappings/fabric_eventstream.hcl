// Mapping for fabric_eventstream resource
// Auto-generated from eventstream/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_eventstream" {
  import_path = "eventstream/definitions.json"



  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreateEventstreamRequest.description"
    max_length = 256
  }

  // required, max 256 chars
  attribute "display_name" {
    api_ref = "CreateEventstreamRequest.displayName"
    max_length = 256
    pattern = "^[a-zA-Z0-9._-]+$"
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateEventstreamRequest.folderId"
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
