// Mapping for fabric_spark_environment_settings resource
// Note: This uses UpdateEnvironmentSparkComputeRequest (not Create) - unusual pattern
mapping "fabric_spark_environment_settings" {
  import_path = "environment/definitions.json"

  // optional, enum(5 values)
  attribute "driver_cores" {
    api_ref = "UpdateEnvironmentSparkComputeRequest.driverCores"
    valid_values = [4, 8, 16, 32, 64]
  }

  // optional, enum(5 values)
  attribute "driver_memory" {
    api_ref = "UpdateEnvironmentSparkComputeRequest.driverMemory"
    valid_values = ["28g", "56g", "112g", "224g", "400g"]
  }

  // optional, enum(5 values)
  attribute "executor_cores" {
    api_ref = "UpdateEnvironmentSparkComputeRequest.executorCores"
    valid_values = [4, 8, 16, 32, 64]
  }

  // optional, enum(5 values)
  attribute "executor_memory" {
    api_ref = "UpdateEnvironmentSparkComputeRequest.executorMemory"
    valid_values = ["28g", "56g", "112g", "224g", "400g"]
  }

  // optional, enum(3 values)
  attribute "runtime_version" {
    api_ref = "UpdateEnvironmentSparkComputeRequest.runtimeVersion"
    valid_values = ["1.1", "1.2", "1.3"]
  }

  // Note: dynamic_executor_allocation, pool, and spark_properties are nested/complex types
  // that cannot be validated with current tool

}
