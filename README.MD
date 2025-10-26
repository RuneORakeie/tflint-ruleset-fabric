# TFLint Ruleset for Fabric Terraform Provider

A comprehensive TFLint ruleset plugin for validating Terraform configurations using the Microsoft Fabric provider. This ruleset provides **70+ validation rules** covering best practices, schema constraints, and business logic for Fabric resources.

[![License: MPL 2.0](https://img.shields.io/badge/License-MPL%202.0-brightgreen.svg)](https://opensource.org/licenses/MPL-2.0)

## Features

### 🎯 70+ Validation Rules

**Business Logic Rules (17)** - Custom governance and best practice rules:
- ✅ Workspace capacity and description requirements
- ✅ Role assignment validation across workspaces, domains, gateways, and pipelines
- ✅ Git integration validation (Azure DevOps & GitHub)
- ✅ Deployment pipeline stage validation (count, naming, descriptions)
- ✅ Domain contributor scope validation
- ✅ Capacity region availability checks

**API Spec Rules (53)** - Auto-generated from Fabric API specifications:
- ✅ Display name length validation (123-256 chars depending on resource)
- ✅ Description length validation (256-1024 chars)
- ✅ Enum value validation (roles, connectivity types, privacy levels)
- ✅ Spark environment settings validation
- ✅ Gateway configuration validation

## Requirements

- TFLint v0.42+
- Go v1.21+
- Terraform v1.0+

## Installation

### From Source

```bash
git clone https://github.com/RuneORakeie/tflint-ruleset-fabric
cd tflint-ruleset-fabric
make build
make install
```

### From Release (Future)

```bash
# Add to .tflint.hcl
plugin "fabric" {
  enabled = true
  version = "0.1.0"
  source  = "github.com/RuneORakeie/tflint-ruleset-fabric"
}

# Then run
tflint --init
```

## Quick Start

1. **Create `.tflint.hcl` in your Terraform project:**

```hcl
plugin "fabric" {
  enabled = true
}

# Enable business logic rules
rule "fabric_workspace_capacity_required" {
  enabled = true
}

rule "fabric_workspace_description_required" {
  enabled = true
}

rule "fabric_role_assignment_recommended" {
  enabled = true
}

rule "fabric_item_description_recommended" {
  enabled = true
}

# Git integration rules
rule "fabric_workspace_git_provider_valid" {
  enabled = true
}

rule "fabric_workspace_git_directory_name_format" {
  enabled = true
}
```

2. **Run TFLint:**

```bash
tflint
```

## Usage Examples

### Workspace Validation

**❌ Bad:**
```hcl
resource "fabric_workspace" "example" {
  display_name = "my-workspace"
  # Missing: capacity_id and description
}
```

**✅ Good:**
```hcl
resource "fabric_workspace" "example" {
  display_name = "analytics-workspace"
  description  = "Production analytics workspace owned by Data Team"
  capacity_id  = fabric_capacity.prod.id
}

resource "fabric_workspace_role_assignment" "admin" {
  workspace_id   = fabric_workspace.example.id
  principal_id   = "user-guid"
  principal_type = "User"
  role           = "Admin"
}
```

### Git Integration Validation

**❌ Bad:**
```hcl
resource "fabric_workspace_git" "example" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  
  git_provider_details {
    git_provider_type = "GitHub"
    repository_name   = "my-repo"
    branch_name       = "main"
    directory_name    = "workspace"  # Missing leading /
    # Missing: owner_name (required for GitHub)
  }
}
```

**✅ Good:**
```hcl
resource "fabric_workspace_git" "example" {
  workspace_id            = fabric_workspace.example.id
  initialization_strategy = "PreferWorkspace"
  credentials_source      = "ServicePrincipal"
  
  git_provider_details {
    git_provider_type = "GitHub"
    owner_name        = "my-org"
    repository_name   = "fabric-config"
    branch_name       = "main"
    directory_name    = "/workspaces/analytics"
  }
}
```

### Deployment Pipeline Validation

**❌ Bad:**
```hcl
resource "fabric_deployment_pipeline" "example" {
  display_name = "Pipeline"
  
  stages {
    display_name = "dev"
    is_public    = true
  }
  # Missing: needs at least 2 stages
}
```

**✅ Good:**
```hcl
resource "fabric_deployment_pipeline" "example" {
  display_name = "Analytics Pipeline"
  description  = "Deployment pipeline for analytics workspaces"
  
  stages {
    display_name = "Development"
    description  = "Development environment for testing"
    is_public    = true
  }
  
  stages {
    display_name = "Production"
    description  = "Production environment"
    is_public    = false
  }
}
```

## Complete Rule List

### Business Logic Rules

| Rule | Severity | Description | Enabled by Default |
|------|----------|-------------|-------------------|
| `fabric_capacity_region_valid` | WARNING | Validates capacity regions support all workloads | No |
| `fabric_deployment_pipeline_stages_count` | ERROR | Ensures 2-10 stages in deployment pipelines | Yes |
| `fabric_deployment_pipeline_stages_description_length` | ERROR | Validates stage description ≤ 1024 chars | Yes |
| `fabric_deployment_pipeline_stages_display_name_length` | ERROR | Validates stage display name ≤ 256 chars | Yes |
| `fabric_domain_contributors_scope` | ERROR | Validates contributor scope values | Yes |
| `fabric_item_description_recommended` | WARNING | Recommends descriptions for governance | Yes |
| `fabric_role_assignment_recommended` | WARNING | Warns when resources lack role assignments | Yes |
| `fabric_workspace_capacity_required` | ERROR | Requires capacity assignment | Yes |
| `fabric_workspace_description_required` | WARNING | Requires workspace descriptions | Yes |
| `fabric_workspace_git_azdo_attributes_required` | ERROR | Validates Azure DevOps git attributes | Yes |
| `fabric_workspace_git_credentials_source_valid` | ERROR | Validates credentials source | Yes |
| `fabric_workspace_git_directory_name_format` | ERROR | Validates directory name format (starts with /) | Yes |
| `fabric_workspace_git_github_attributes_required` | ERROR | Validates GitHub git attributes | Yes |
| `fabric_workspace_git_initialization_strategy_valid` | ERROR | Validates initialization strategy | Yes |
| `fabric_workspace_git_provider_valid` | ERROR | Validates git provider type | Yes |
| `fabric_workspace_git_string_lengths` | ERROR | Validates git string lengths | Yes |
| `fabric_workspace_role_assignment_role` | ERROR | Validates workspace role values | Yes |

### API Spec Rules (Auto-generated)

<details>
<summary><b>Display Name Rules (26 rules)</b></summary>

- `fabric_activator_invalid_display_name`
- `fabric_apache_airflow_job_invalid_display_name`
- `fabric_connection_invalid_display_name`
- `fabric_copy_job_invalid_display_name`
- `fabric_data_pipeline_invalid_display_name`
- `fabric_dataflow_invalid_display_name`
- `fabric_deployment_pipeline_invalid_display_name`
- `fabric_digital_twin_builder_invalid_display_name`
- `fabric_domain_invalid_display_name`
- `fabric_environment_invalid_display_name`
- `fabric_eventhouse_invalid_display_name`
- `fabric_eventstream_invalid_display_name`
- `fabric_folder_invalid_display_name`
- `fabric_gateway_invalid_display_name`
- `fabric_graphql_api_invalid_display_name`
- `fabric_kql_dashboard_invalid_display_name`
- `fabric_kql_database_invalid_display_name`
- `fabric_kql_queryset_invalid_display_name`
- `fabric_lakehouse_invalid_display_name`
- `fabric_mirrored_database_invalid_display_name`
- `fabric_ml_experiment_invalid_display_name`
- `fabric_ml_model_invalid_display_name`
- `fabric_mounted_data_factory_invalid_display_name`
- `fabric_notebook_invalid_display_name`
- `fabric_report_invalid_display_name`
- `fabric_semantic_model_invalid_display_name`
- `fabric_spark_job_definition_invalid_display_name`
- `fabric_sql_database_invalid_display_name`
- `fabric_variable_library_invalid_display_name`
- `fabric_warehouse_invalid_display_name`
- `fabric_warehouse_snapshot_invalid_display_name`
- `fabric_workspace_invalid_display_name`
</details>

<details>
<summary><b>Description Rules (21 rules)</b></summary>

- `fabric_activator_invalid_description`
- `fabric_apache_airflow_job_invalid_description`
- `fabric_copy_job_invalid_description`
- `fabric_data_pipeline_invalid_description`
- `fabric_dataflow_invalid_description`
- `fabric_deployment_pipeline_invalid_description`
- `fabric_digital_twin_builder_invalid_description`
- `fabric_domain_invalid_description`
- `fabric_environment_invalid_description`
- `fabric_eventhouse_invalid_description`
- `fabric_eventstream_invalid_description`
- `fabric_graphql_api_invalid_description`
- `fabric_kql_dashboard_invalid_description`
- `fabric_kql_database_invalid_description`
- `fabric_kql_queryset_invalid_description`
- `fabric_lakehouse_invalid_description`
- `fabric_mirrored_database_invalid_description`
- `fabric_ml_experiment_invalid_description`
- `fabric_ml_model_invalid_description`
- `fabric_mounted_data_factory_invalid_description`
- `fabric_notebook_invalid_description`
- `fabric_report_invalid_description`
- `fabric_semantic_model_invalid_description`
- `fabric_spark_job_definition_invalid_description`
- `fabric_sql_database_invalid_description`
- `fabric_variable_library_invalid_description`
- `fabric_warehouse_invalid_description`
- `fabric_warehouse_snapshot_invalid_description`
- `fabric_workspace_invalid_description`
</details>

<details>
<summary><b>Other API Spec Rules (6 rules)</b></summary>

- `fabric_connection_invalid_connectivity_type`
- `fabric_connection_invalid_privacy_level`
- `fabric_gateway_invalid_inactivity_minutes_before_sleep`
- `fabric_gateway_invalid_type`
- `fabric_spark_custom_pool_invalid_node_family`
- `fabric_spark_custom_pool_invalid_node_size`
- `fabric_spark_environment_settings_invalid_driver_cores`
- `fabric_spark_environment_settings_invalid_driver_memory`
- `fabric_spark_environment_settings_invalid_executor_cores`
- `fabric_spark_environment_settings_invalid_executor_memory`
- `fabric_spark_environment_settings_invalid_runtime_version`
</details>

## Configuration

### Rule Presets

**Minimal (Errors only):**
```hcl
plugin "fabric" {
  enabled = true
}

# Only enable ERROR severity rules
rule "fabric_workspace_capacity_required" {
  enabled = true
}

rule "fabric_deployment_pipeline_stages_count" {
  enabled = true
}
```

**Recommended (Errors + Warnings):**
```hcl
plugin "fabric" {
  enabled = true
}

# All business logic rules
rule "fabric_*" {
  enabled = true
}
```

**Strict (All rules):**
```hcl
plugin "fabric" {
  enabled = true
}

# Enable all rules including API spec rules
rule "fabric_*" {
  enabled = true
}

rule "fabric_*_invalid_*" {
  enabled = true
}
```

## Development

### Project Structure

```
├── main.go                             # Plugin entry point
├── rules/
│   ├── fabric_capacity_region.go       # Business logic rules
│   ├── fabric_workspace_*.go           # Workspace rules
│   ├── fabric_deployment_*.go          # Deployment rules
│   ├── apispec/                        # Auto-generated API rules
│   │   ├── provider.go
│   │   └── fabric_*_invalid_*.go
│   ├── business_logic_rules_test.go    # Tests
│   └── generated_rules_test.go
├── docs/
│   └── rules/                          # Rule documentation
│       └── *.md
├── tools/
│   └── apispec-rule-gen/               # Code generator
├── Makefile
└── go.mod
```

### Building

```bash
# Build the plugin
make build

# Run tests
make test

# Run linter
make lint

# Format code
make fmt
```

### Adding New Rules

1. **Business Logic Rules**: Create in `rules/` directory
2. **API Spec Rules**: Add mapping in `tools/apispec-rule-gen/mappings/`, then run generator
3. **Tests**: Add tests to appropriate test file
4. **Documentation**: Add rule documentation in `docs/rules/`

## Contributing

Contributions welcome! Please:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-rule`)
3. Add tests for new rules
4. Run `make test` and `make lint`
5. Commit changes (`git commit -m 'Add amazing rule'`)
6. Push to branch (`git push origin feature/amazing-rule`)
7. Open a Pull Request

See [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Links

- 📚 [Rule Documentation](docs/rules/)
- 🔧 [TFLint Documentation](https://github.com/terraform-linters/tflint)
- 🔌 [TFLint Plugin SDK](https://github.com/terraform-linters/tflint-plugin-sdk)
- ☁️ [Microsoft Fabric Terraform Provider](https://github.com/microsoft/terraform-provider-fabric)
- 📖 [Fabric Documentation](https://learn.microsoft.com/en-us/fabric/)

## License

Mozilla Public License 2.0 - see [LICENSE](LICENSE) file for details.

## Support

- 🐛 [GitHub Issues](https://github.com/RuneORakeie/tflint-ruleset-fabric/issues)
- 💬 [GitHub Discussions](https://github.com/RuneORakeie/tflint-ruleset-fabric/discussions)
- 📧 [Microsoft Fabric Support](https://support.fabric.microsoft.com/)

---

**Made with ❤️ for the Fabric community**
