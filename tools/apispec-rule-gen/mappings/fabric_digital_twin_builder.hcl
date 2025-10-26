// Mapping for fabric_digital_twin_builder resource
// Auto-generated from digitalTwinBuilder/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_digital_twin_builder" {
  import_path = "digitalTwinBuilder/definitions.json"

  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreateDigitalTwinBuilderRequest.description"
    max_length = 256
  }

  // required
  attribute "display_name" {
    api_ref = "CreateDigitalTwinBuilderRequest.displayName"
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
