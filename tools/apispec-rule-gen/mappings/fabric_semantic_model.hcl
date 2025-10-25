// Mapping for fabric_semantic_model resource
// Auto-generated from semanticModel/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_semantic_model" {
  import_path = "semanticModel/definitions.json"

  // required
  attribute "definition" {
    api_ref = "CreateSemanticModelRequest.definition"
  }

  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreateSemanticModelRequest.description"
    max_length = 256
  }

  // required
  attribute "display_name" {
    api_ref = "CreateSemanticModelRequest.displayName"
  }

  // optional, format: uuid
  attribute "folder_id" {
    api_ref = "CreateSemanticModelRequest.folderId"
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
