// Mapping for fabric_data_pipeline resource
// Auto-generated from dataPipeline/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_data_pipeline" {
  import_path = "dataPipeline/definitions.json"


  // optional, max 1024 chars
  attribute "description" {
    api_ref = "CreateDataPipelineRequest.description"
    max_length = 1024
  }

  // required, max 256 chars
  attribute "display_name" {
    api_ref = "CreateDataPipelineRequest.displayName"
    max_length = 256
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateDataPipelineRequest.folderId"
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
