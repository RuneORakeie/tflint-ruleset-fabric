// Mapping for fabric_shortcut resource
// Auto-generated from platform/definitions/platform.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_shortcut" {
  import_path = "platform/definitions/platform.json"

  // required
  attribute "name" {
    api_ref = "CreateShortcutRequest.name"
  }

  // required
  attribute "path" {
    api_ref = "CreateShortcutRequest.path"
  }

  // required
  attribute "target" {
    api_ref = "CreateShortcutRequest.target"
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
