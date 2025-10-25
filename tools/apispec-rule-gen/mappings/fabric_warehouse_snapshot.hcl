// Mapping for fabric_warehouse_snapshot resource
// Auto-generated from warehouseSnapshot/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_warehouse_snapshot" {
  import_path = "warehouseSnapshot/definitions.json"

  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreateWarehouseSnapshotRequest.description"
    max_length = 256
  }

  // required
  attribute "display_name" {
    api_ref = "CreateWarehouseSnapshotRequest.displayName"
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateWarehouseSnapshotRequest.folderId"
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
