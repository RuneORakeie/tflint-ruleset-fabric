// Mapping for fabric_paginated_report resource
// Auto-generated from paginatedReport/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.
mapping "fabric_paginated_report" {
  import_path = "paginatedReport/definitions.json"

  // optional, max 256 chars
  attribute "description" {
    api_ref = "CreatePaginatedReportRequest.description"
    max_length = 256
  }

}