// Mapping for fabric_external_data_share_shortcut resource
// Auto-generated from platform/definitions/externaldatasharing.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_external_data_share_shortcut" {
  import_path = "platform/definitions/externaldatasharing.json"

  // required, format: uuid
  attribute "path_id" {
    api_ref = "CreateExternalDataShareShortcutRequest.pathId"
  }

  // required
  attribute "shortcut_name" {
    api_ref = "CreateExternalDataShareShortcutRequest.shortcutName"
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
