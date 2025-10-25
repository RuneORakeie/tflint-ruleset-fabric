// Mapping for fabric_workspace_managed_private_endpoint resource
// Auto-generated from platform/definitions/managedPrivateEndpoint.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_workspace_managed_private_endpoint" {
  import_path = "platform/definitions/managedPrivateEndpoint.json"

  // required
  attribute "name" {
    api_ref = "CreateManagedPrivateEndpointRequest.name"
  }

  // required
  attribute "target_private_link_resource_id" {
    api_ref = "CreateManagedPrivateEndpointRequest.targetPrivateLinkResourceId"
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
