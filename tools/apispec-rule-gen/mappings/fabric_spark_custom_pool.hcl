// Mapping for fabric_spark_custom_pool resource
// Auto-generated from spark/definitions.json
// DO NOT EDIT auto-generated sections directly.
// Add custom constraints with // MANUAL: comment to preserve during updates.

mapping "fabric_spark_custom_pool" {
  import_path = "spark/definitions.json"

  // required
  attribute "auto_scale" {
    api_ref = "CreateCustomPoolRequest.autoScale"
  }

  // required
  attribute "dynamic_executor_allocation" {
    api_ref = "CreateCustomPoolRequest.dynamicExecutorAllocation"
  }

  // required
  attribute "name" {
    api_ref = "CreateCustomPoolRequest.name"
  }

  // required, enum(1 values)
  attribute "node_family" {
    api_ref = "CreateCustomPoolRequest.nodeFamily"
    valid_values = ["MemoryOptimized"]
  }

  // required, enum(5 values)
  attribute "node_size" {
    api_ref = "CreateCustomPoolRequest.nodeSize"
    valid_values = ["Small", "Medium", "Large", "XLarge", "XXLarge"]
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
