// Mapping for fabric_anomaly_detector resource
// Auto-generated from anomalyDetector/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_anomaly_detector" {
  import_path = "anomalyDetector/definitions.json"


  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreateAnomalyDetectorRequest.description"
    max_length = 256
  }

  // required
  attribute "display_name" {
    api_ref = "CreateAnomalyDetectorRequest.displayName"
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateAnomalyDetectorRequest.folderId"
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
