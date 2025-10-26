// Mapping for fabric_schedule resource
// Auto-generated from platform/definitions/platform.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_schedule" {
  import_path = "platform/definitions/platform.json"

  // required
  attribute "configuration" {
    api_ref = "CreateScheduleRequest.configuration"
  }

  // required
  attribute "enabled" {
    api_ref = "CreateScheduleRequest.enabled"
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
