// Mapping for fabric_mirrored_azure_databricks_catalog resource
// Auto-generated from mirroredAzureDatabricksCatalog/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_mirrored_azure_databricks_catalog" {
  import_path = "mirroredAzureDatabricksCatalog/definitions.json"

  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreateMirroredAzureDatabricksCatalogRequest.description"
    max_length = 256
  }

  // required
  attribute "display_name" {
    api_ref = "CreateMirroredAzureDatabricksCatalogRequest.displayName"
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
