// Mapping for fabric_spark_job_definition resource
// Auto-generated from sparkjobdefinition/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_spark_job_definition" {
  import_path = "sparkjobdefinition/definitions.json"

  // optional
  attribute "definition_manual_manual" {
    api_ref = "manual.definition_manual"
  }

  // optional, max 1021 chars
  attribute "description" {
    api_ref = "CreateSparkJobDefinitionRequest.description"
    max_length = 1021
  }

  // required, max 256 chars
  attribute "display_name" {
    api_ref = "CreateSparkJobDefinitionRequest.displayName"
    max_length = 256
    pattern = "^[a-zA-Z0-9_ ]+$"
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateSparkJobDefinitionRequest.folderId"
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
