// Mapping for fabric_dataflow resource
// Auto-generated from dataflow/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_dataflow" {
  import_path = "dataflow/definitions.json"

  // optional, max 3988 chars
  attribute "description" {
    api_ref = "CreateDataflowRequest.description"
    max_length = 3988
  }

  // required, max 256 chars
  attribute "display_name" {
    api_ref = "CreateDataflowRequest.displayName"
    max_length = 256
    pattern = "^[a-zA-Z0-9\\s()\\[\\]{}+\\-=_#]+$"
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateDataflowRequest.folderId"
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
