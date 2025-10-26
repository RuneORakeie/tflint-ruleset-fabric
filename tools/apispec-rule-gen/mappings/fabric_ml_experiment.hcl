// Mapping for fabric_ml_experiment resource
// Auto-generated from mlExperiment/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_ml_experiment" {
  import_path = "mlExperiment/definitions.json"

  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreateMLExperimentRequest.description"
    max_length = 256
  }

  // required
  attribute "display_name" {
    api_ref = "CreateMLExperimentRequest.displayName"
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateMLExperimentRequest.folderId"
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
