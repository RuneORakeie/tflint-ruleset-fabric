# TFLint Ruleset for Fabric Terraform Provider

This is a TFLint ruleset plugin for validating Terraform configurations using the Microsoft Fabric provider. It provides best practices and error detection for Fabric resources including workspaces, role assignments, Git integrations, and capacities.

## Features

This ruleset includes the following rules:

| Rule Name | Severity | Description |
|-----------|----------|-------------|
| `fabric_workspace_naming` | WARNING | Validates workspace names follow naming conventions (3-50 chars, lowercase, hyphens only) |
| `fabric_workspace_capacity_required` | ERROR | Ensures workspaces have capacity assigned for production |
| `fabric_workspace_description_required` | WARNING | Ensures workspaces have descriptions for governance |
| `fabric_role_assignment_principal_required` | ERROR | Validates role assignments specify principals |
| `fabric_git_integration_provider_valid` | ERROR | Validates Git provider type is one of supported providers |
| `fabric_capacity_region_valid` | WARNING | Validates capacity regions are available (disabled by default) |

## Requirements

- TFLint v0.42+
- Go v1.25+
- Terraform v1.0+

## Installation

### From Source

```bash
git clone https://github.com/RuneORakeie/tflint-ruleset-fabric
cd tflint-ruleset-fabric
make install
```

### From Release

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

## Configuration

Add to your `.tflint.hcl`:

```hcl
plugin "fabric" {
  enabled = true
  version = "0.1.0"
}

rule "fabric_workspace_naming" {
  enabled = true
}

rule "fabric_workspace_capacity_required" {
  enabled = true
}

rule "fabric_workspace_description_required" {
  enabled = true
}

rule "fabric_role_assignment_principal_required" {
  enabled = true
}

rule "fabric_git_integration_provider_valid" {
  enabled = true
}

rule "fabric_capacity_region_valid" {
  enabled = false  # Enable to validate regions
}
```

## Usage

After installation, run TFLint normally:

```bash
tflint
```

To check only Fabric rules:

```bash
tflint --only fabric_workspace_naming
```

To filter by severity:

```bash
tflint --minimum-failure-severity error
```

## Examples

### Valid Workspace Configuration

```hcl
resource "fabric_workspace" "example" {
  display_name = "my-analytics-workspace"
  description  = "Production analytics workspace for business unit"
  capacity_id  = data.fabric_capacity.prod.id
}
```

### Valid Role Assignment

```hcl
resource "fabric_workspace_role_assignment" "admin" {
  workspace_id   = fabric_workspace.example.id
  principal_id   = "00000000-0000-0000-0000-000000000001"
  role           = "Admin"
  principal_type = "User"
}
```

### Valid Git Integration

```hcl
resource "fabric_workspace_git_connection" "example" {
  workspace_id       = fabric_workspace.example.id
  git_provider_type  = "GitHub"
  repository_name    = "my-fabric-repo"
  branch_name        = "main"
  organization_name  = "my-org"
}
```

## Development

### Building the Plugin

```bash
make build
```

### Running Tests

```bash
make test
```

### Code Formatting

```bash
make fmt
```

### Linting

```bash
make lint
```

## Testing with Examples

The `examples/` directory contains sample Terraform configurations for testing:

```bash
cd examples
tflint
```

## Contributing

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## Project Structure

```
├── main.go                          # Plugin entry point
├── rules/
│   ├── fabric_workspace_naming.go
│   ├── fabric_workspace_capacity.go
│   ├── fabric_workspace_description.go
│   ├── fabric_role_assignments.go
│   ├── fabric_git_integration.go
│   ├── fabric_capacity_region.go
│   └── generator/
│       └── main.go                 # Rule documentation generator
├── examples/
│   ├── valid/
│   └── invalid/
├── Makefile
├── go.mod
├── go.sum
├── README.md
└── LICENSE
```

## Links

- [TFLint Documentation](https://github.com/terraform-linters/tflint)
- [TFLint Plugin SDK](https://github.com/terraform-linters/tflint-plugin-sdk)
- [Microsoft Fabric Terraform Provider](https://github.com/microsoft/terraform-provider-fabric)
- [Fabric Documentation](https://learn.microsoft.com/en-us/fabric/)

## License

This project is licensed under the Mozilla Public License 2.0 - see LICENSE file for details.

## Support

For issues and questions:
- [GitHub Issues](https://github.com/RuneORakeie/tflint-ruleset-fabric/issues)
- [Microsoft Fabric Support](https://support.fabric.microsoft.com/)