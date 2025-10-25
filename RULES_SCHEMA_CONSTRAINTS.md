# TFLint Rules - Complete Reference

This document provides a comprehensive overview of all TFLint rules in this ruleset, organized by category and type.

## Overview

**Total Rules**: 70+
- **Business Logic Rules**: 17 custom rules
- **API Spec Rules**: 53 auto-generated rules

---

## Business Logic Rules

These are custom-built rules that enforce Fabric-specific best practices and governance requirements.

### Workspace Rules

#### fabric_workspace_capacity_required
- **Severity**: ERROR
- **Description**: Ensures workspaces have a capacity assigned for production use
- **Enabled by default**: Yes
- **Rationale**: Workspaces without capacity assignment cannot run workloads
- **Example**:
```hcl
# Good
resource "fabric_workspace" "example" {
  display_name = "Analytics"
  capacity_id  = fabric_capacity.prod.id
}
```

#### fabric_workspace_description_required
- **Severity**: WARNING
- **Description**: Ensures workspaces have descriptions for governance
- **Enabled by default**: Yes
- **Rationale**: Descriptions improve documentation and team collaboration
- **Example**:
```hcl
# Good
resource "fabric_workspace" "example" {
  display_name = "Analytics"
  description  = "Production analytics workspace - Data Team"
}
```

### Git Integration Rules

#### fabric_workspace_git_provider_valid
- **Severity**: ERROR
- **Description**: Validates Git provider type is GitHub or AzureDevOps
- **Enabled by default**: Yes
- **Valid values**: `GitHub`, `AzureDevOps`

#### fabric_workspace_git_directory_name_format
- **Severity**: ERROR
- **Description**: Validates directory name starts with `/` and is ≤256 chars
- **Enabled by default**: Yes
- **Example**:
```hcl
# Good
git_provider_details {
  directory_name = "/workspaces/analytics"
}
```

#### fabric_workspace_git_github_attributes_required
- **Severity**: ERROR
- **Description**: Validates GitHub git configurations have required attributes
- **Enabled by default**: Yes
- **Required for GitHub**: `owner_name`, `repository_name`, `branch_name`, `directory_name`

#### fabric_workspace_git_azdo_attributes_required
- **Severity**: ERROR
- **Description**: Validates Azure DevOps git configurations have required attributes
- **Enabled by default**: Yes
- **Required for Azure DevOps**: `organization_name`, `project_name`, `repository_name`, `branch_name`, `directory_name`

#### fabric_workspace_git_initialization_strategy_valid
- **Severity**: ERROR
- **Description**: Validates initialization strategy value
- **Enabled by default**: Yes
- **Valid values**: `PreferWorkspace`, `PreferRemote`

#### fabric_workspace_git_credentials_source_valid
- **Severity**: ERROR
- **Description**: Validates credentials source value
- **Enabled by default**: Yes
- **Valid values**: `ServicePrincipal`, `UserAuthentication`

#### fabric_workspace_git_string_lengths
- **Severity**: ERROR
- **Description**: Validates git string attribute lengths
- **Enabled by default**: Yes
- **Limits**:
  - `repository_name`: ≤100 chars
  - `branch_name`: ≤250 chars
  - `organization_name`: ≤100 chars
  - `project_name`: ≤100 chars
  - `owner_name`: ≤100 chars

### Role Assignment Rules

#### fabric_workspace_role_assignment_role
- **Severity**: ERROR
- **Description**: Validates workspace role values
- **Enabled by default**: Yes
- **Valid values**: `Admin`, `Member`, `Contributor`, `Viewer`

#### fabric_role_assignment_recommended
- **Severity**: WARNING
- **Description**: Warns when resources lack role assignments
- **Enabled by default**: Yes
- **Applies to**: Workspaces, deployment pipelines, domains, gateways
- **Rationale**: Resources without role assignments may not be accessible

### Deployment Pipeline Rules

#### fabric_deployment_pipeline_stages_count
- **Severity**: ERROR
- **Description**: Validates deployment pipelines have 2-10 stages
- **Enabled by default**: Yes
- **Limits**: Minimum 2 stages, maximum 10 stages

#### fabric_deployment_pipeline_stages_display_name_length
- **Severity**: ERROR
- **Description**: Validates stage display names are ≤256 characters
- **Enabled by default**: Yes
- **Limit**: 256 characters

#### fabric_deployment_pipeline_stages_description_length
- **Severity**: ERROR
- **Description**: Validates stage descriptions are ≤1024 characters
- **Enabled by default**: Yes
- **Limit**: 1024 characters

### Domain Rules

#### fabric_domain_contributors_scope
- **Severity**: ERROR
- **Description**: Validates domain contributor scope values
- **Enabled by default**: Yes
- **Valid values**: `Workspace`, `DomainOnly`, `SpecificWorkspaces`

### Item Rules

#### fabric_item_description_recommended
- **Severity**: WARNING
- **Description**: Recommends descriptions for all Fabric items
- **Enabled by default**: Yes
- **Applies to**: All item types (workspaces, lakehouses, notebooks, etc.)
- **Rationale**: Descriptions improve governance and documentation

### Capacity Rules

#### fabric_capacity_region_valid
- **Severity**: WARNING
- **Description**: Validates capacity regions support all Fabric workloads
- **Enabled by default**: No (opt-in)
- **Rationale**: Some regions may not have all Fabric features available
- **Regions checked**: 40+ Azure regions

---

## API Spec Rules

Auto-generated rules from the Fabric Terraform Provider schema. These rules validate schema constraints like string lengths and enum values.

### Display Name Rules (32 rules)

All display name rules validate that display names don't exceed their maximum length:

- Most resources: 123 characters
- Deployment pipelines: 246 characters
- Domains: 40 characters
- Pipeline stages: 256 characters

**Resources with display name validation:**
- fabric_activator
- fabric_apache_airflow_job
- fabric_connection
- fabric_copy_job
- fabric_data_pipeline
- fabric_dataflow
- fabric_deployment_pipeline
- fabric_digital_twin_builder
- fabric_domain
- fabric_environment
- fabric_eventhouse
- fabric_eventstream
- fabric_folder
- fabric_gateway
- fabric_graphql_api
- fabric_kql_dashboard
- fabric_kql_database
- fabric_kql_queryset
- fabric_lakehouse
- fabric_mirrored_database
- fabric_ml_experiment
- fabric_ml_model
- fabric_mounted_data_factory
- fabric_notebook
- fabric_report
- fabric_semantic_model
- fabric_spark_job_definition
- fabric_sql_database
- fabric_variable_library
- fabric_warehouse
- fabric_warehouse_snapshot
- fabric_workspace

### Description Rules (29 rules)

All description rules validate that descriptions don't exceed 256 characters (except pipeline stage descriptions which can be up to 1024 characters).

**Resources with description validation:**
- fabric_activator
- fabric_apache_airflow_job
- fabric_connection
- fabric_copy_job
- fabric_data_pipeline
- fabric_dataflow
- fabric_deployment_pipeline
- fabric_digital_twin_builder
- fabric_domain
- fabric_environment
- fabric_eventhouse
- fabric_eventstream
- fabric_graphql_api
- fabric_kql_dashboard
- fabric_kql_database
- fabric_kql_queryset
- fabric_lakehouse
- fabric_mirrored_database
- fabric_ml_experiment
- fabric_ml_model
- fabric_mounted_data_factory
- fabric_notebook
- fabric_report
- fabric_semantic_model
- fabric_spark_job_definition
- fabric_sql_database
- fabric_variable_library
- fabric_warehouse
- fabric_warehouse_snapshot
- fabric_workspace

### Enum Validation Rules

#### fabric_connection_invalid_connectivity_type
- **Valid values**: `ShareableCloud`, `VirtualNetworkGateway`

#### fabric_connection_invalid_privacy_level
- **Valid values**: `None`, `Organizational`, `Private`, `Public`

#### fabric_gateway_invalid_type
- **Valid values**: Various gateway types

### Spark Environment Rules

#### fabric_spark_environment_settings_invalid_driver_cores
- **Valid values**: 1, 2, 4, 8

#### fabric_spark_environment_settings_invalid_driver_memory
- **Valid values**: Memory size strings (e.g., "4g", "8g", "16g")

#### fabric_spark_environment_settings_invalid_executor_cores
- **Valid values**: 1, 2, 4, 8

#### fabric_spark_environment_settings_invalid_executor_memory
- **Valid values**: Memory size strings (e.g., "4g", "8g", "16g")

#### fabric_spark_environment_settings_invalid_runtime_version
- **Description**: Validates Spark runtime version

### Spark Pool Rules

#### fabric_spark_custom_pool_invalid_node_family
- **Description**: Validates node family for custom Spark pools

#### fabric_spark_custom_pool_invalid_node_size
- **Description**: Validates node size for custom Spark pools

### Gateway Rules

#### fabric_gateway_invalid_inactivity_minutes_before_sleep
- **Description**: Validates inactivity timeout value

---

## Rule Categories by Severity

### ERROR Rules (51)
Critical validation failures that should prevent deployment:
- All schema constraint violations
- Missing required attributes
- Invalid enum values
- Incorrect string formats

### WARNING Rules (19)
Best practice violations that should be addressed:
- Missing descriptions
- Missing role assignments
- Suboptimal configurations
- Regional availability concerns

---

## Rule Naming Conventions

Rules follow these naming patterns:

**Business Logic Rules:**
- `fabric_<resource>_<attribute>_<constraint>`
- Examples: `fabric_workspace_capacity_required`, `fabric_workspace_git_directory_name_format`

**API Spec Rules:**
- `fabric_<resource>_invalid_<attribute>`
- Examples: `fabric_connection_invalid_display_name`, `fabric_workspace_invalid_description`

---

## Enabling/Disabling Rules

### Enable All Rules
```hcl
plugin "fabric" {
  enabled = true
}
```

### Enable Specific Categories
```hcl
# Only business logic rules
rule "fabric_workspace_*" {
  enabled = true
}

rule "fabric_deployment_pipeline_*" {
  enabled = true
}

# Only API spec rules
rule "fabric_*_invalid_*" {
  enabled = true
}
```

### Disable Specific Rules
```hcl
rule "fabric_capacity_region_valid" {
  enabled = false
}

rule "fabric_item_description_recommended" {
  enabled = false
}
```

---

## Testing Rules

All rules have comprehensive test coverage:

**Business Logic Rules**: `rules/business_logic_rules_test.go`
**API Spec Rules**: `rules/generated_rules_test.go`

Run tests:
```bash
make test
```

---

## Documentation

Each rule has dedicated documentation in `docs/rules/`:
- Description and rationale
- Severity level
- Configuration examples
- Valid/invalid examples
- Related links

View rule docs: [docs/rules/](../docs/rules/)

---

## Contributing New Rules

See [CONTRIBUTING.md](../CONTRIBUTING.md) for guidelines on:
- Adding business logic rules
- Generating API spec rules
- Writing tests
- Documenting rules

---

**Last Updated**: January 2026
