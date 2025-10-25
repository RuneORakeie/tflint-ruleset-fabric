// Mapping for fabric_kql_queryset resource
// Auto-generated from kqlQueryset/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_kql_queryset" {
  import_path = "kqlQueryset/definitions.json"

  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreateKQLQuerysetRequest.description"
    max_length = 256
  }

  // required
  attribute "display_name" {
    api_ref = "CreateKQLQuerysetRequest.displayName"
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateKQLQuerysetRequest.folderId"
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
