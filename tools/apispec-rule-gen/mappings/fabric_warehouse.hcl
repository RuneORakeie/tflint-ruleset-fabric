// Mapping for fabric_warehouse resource
// Auto-generated from warehouse/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_warehouse" {
  import_path = "warehouse/definitions.json"

  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreateWarehouseRequest.description"
    max_length = 256
  }

  // required
  attribute "display_name" {
    api_ref = "CreateWarehouseRequest.displayName"
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateWarehouseRequest.folderId"
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
