// Mapping for fabric_domain resource
// Auto-generated from admin/definitions/domains.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_domain" {
  import_path = "admin/definitions/domains.json"

  // optional, format: uuid
  attribute "default_label_id_domain" {
    api_ref = "UpdateDomainRequest.defaultLabelId"
  }

  // optional
  attribute "default_label_id_domain" {
    api_ref = "UpdateDomainRequest.defaultLabelId"
  }

  // optional
  attribute "default_label_id_domain_manual" {
    api_ref = "manual.default_label_id_domain"
  }

  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreateDomainRequest.description"
    max_length = 256
  }

  // required, max 40 chars
  attribute "display_name" {
    api_ref = "CreateDomainRequest.displayName"
    max_length = 40
  }

  // optional, format: uuid
  attribute "parent_domain_id" {
    api_ref = "CreateDomainRequest.parentDomainId"
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
