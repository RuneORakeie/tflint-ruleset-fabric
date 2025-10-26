// Mapping for fabric_ml_model resource
// Auto-generated from mlModel/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_ml_model" {
  import_path = "mlModel/definitions.json"

  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreateMLModelRequest.description"
    max_length = 256
  }

  // required
  attribute "display_name" {
    api_ref = "CreateMLModelRequest.displayName"
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateMLModelRequest.folderId"
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
