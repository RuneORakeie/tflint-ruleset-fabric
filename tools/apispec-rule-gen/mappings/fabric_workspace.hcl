// Mapping for fabric_workspace resource
// Auto-generated from platform/definitions/platform.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_workspace" {
  import_path = "platform/definitions/platform.json"

  // optional, format: uuid
  attribute "capacity_id" {
    api_ref = "CreateWorkspaceRequest.capacityId"
  }

  // optional, max 4000 chars
  attribute "description" {
    api_ref = "CreateWorkspaceRequest.description"
    max_length = 4000
  }

  // required, max 256 chars
  attribute "display_name" {
    api_ref = "CreateWorkspaceRequest.displayName"
    max_length = 256
  }

  // optional, format: uuid
  attribute "domain_id" {
    api_ref = "CreateWorkspaceRequest.domainId"
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
