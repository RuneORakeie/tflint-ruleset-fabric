// Mapping for fabric_eventhouse resource
// Auto-generated from eventhouse/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_eventhouse" {
  import_path = "eventhouse/definitions.json"



  // optional, max 1024 chars
  attribute "description" {
    api_ref = "CreateEventhouseRequest.description"
    max_length = 1024
  }

  // required, max 256 chars
  attribute "display_name" {
    api_ref = "CreateEventhouseRequest.displayName"
    max_length = 256
    pattern = "^[a-zA-Z0-9._-]+$"
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateEventhouseRequest.folderId"
  }

   // Add manual customizations below with // MANUAL: comment
 
  // MANUAL: required, format: uuid
  // Workspace that owns the Eventhouse
  attribute "workspace_id" {
    api_ref = "https://api.fabric.microsoft.com/v1/workspaces/{workspaceId}/eventhouses"
    pattern = "^[0-9a-fA-F\\-]{36}$"
  }

  // optional, enum(1 value)
  // MANUAL: enum constraint — only "Default" is currently allowed
  attribute "format" {
    valid_values = ["Default"]
    api_ref = "EventhouseDefinition.properties.format"
  }
  
  // MANUAL: boolean flag (default true) — allow both true/false
  attribute "definition_update_enabled" {
    api_ref = "CreateEventhouseRequest.definitionUpdateEnabled"
  }

  // MANUAL: mutual-exclusion rule: exactly one of creationPayload or definition must be set
  // (this may be a custom rule rather than attribute constraint)

  // MANUAL: within creationPayload (alias configuration) minimumConsumptionUnits numeric domain: 0,2.25,4.25,8.5,13,18,26,34,50 or integer 51-322
  attribute "creation_payload.minimum_consumption_units" {
    api_ref = "CreateEventhouseRequest.creationPayload.minimumConsumptionUnits"
    pattern = "^(0|2\\.25|4\\.25|8\\.5|13|18|26|34|50|(5[1-9]|[6-9]\\d|[12]\\d{2}|3[01]\\d|32[0-2]))$"
  }

  // MANUAL: definition parts map — accepted key for Default format is "EventhouseProperties.json"
  // NOTE: Key restriction enforced via pattern on map keys (tooling should apply if supported)
  attribute "definition" {
    api_ref = "CreateEventhouseRequest.definition"
    // key_pattern supports restricting map keys if your generator honors it
    // key_pattern = "^(EventhouseProperties\\.json)$"
  }

  // MANUAL: configuration.minimum_consumption_units — numeric domain:
  // Allowed values: 0, 2.25, 4.25, 8.5, 13, 18, 26, 34, 50, or any integer 51–322.
  // If nested attribute patterns are supported, the following block constrains the value.
  // Otherwise, keep as documentation for a custom rule.
  
}
