// Mapping for fabric_notebook resource
// Auto-generated from notebook/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_notebook" {
  import_path = "notebook/definitions.json"

  // optional
  attribute "definition" {
    api_ref = "CreateNotebookRequest.definition"
  }

  // optional, max 1021 chars
  attribute "description" {
    api_ref = "CreateNotebookRequest.description"
    max_length = 1021
  }

  // required, max 256 chars
  attribute "display_name" {
    api_ref = "CreateNotebookRequest.displayName"
    max_length = 256
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateNotebookRequest.folderId"
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
