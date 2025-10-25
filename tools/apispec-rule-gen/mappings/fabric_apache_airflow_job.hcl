// Mapping for fabric_apache_airflow_job resource
// Auto-generated from apacheAirflowJob/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_apache_airflow_job" {
  import_path = "apacheAirflowJob/definitions.json"

  // optional
  attribute "definition_manual_manual" {
    api_ref = "manual.definition_manual"
  }

  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreateApacheAirflowJobRequest.description"
    max_length = 256
  }

  // required
  attribute "display_name" {
    api_ref = "CreateApacheAirflowJobRequest.displayName"
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateApacheAirflowJobRequest.folderId"
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
