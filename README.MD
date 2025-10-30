# TFLint Ruleset for Fabric Terraform Provider
[![Build Status](https://github.com/RuneORakeie/tflint-ruleset-fabric/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/RuneORakeie/tflint-ruleset-fabric/actions)
[![GitHub release](https://img.shields.io/github/release/RuneORakeie/tflint-ruleset-fabric.svg)](https://github.com/RuneORakeie/tflint-ruleset-fabric/releases/latest)
[![codecov](https://codecov.io/gh/RuneORakeie/tflint-ruleset-fabric/branch/main/graph/badge.svg?token=BNJI183RPY)](https://codecov.io/gh/RuneORakeie/tflint-ruleset-fabric)
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL%202.0-brightgreen.svg)](https://opensource.org/licenses/MPL-2.0)

A comprehensive TFLint ruleset plugin for validating Terraform configurations using the Microsoft Fabric provider. This ruleset provides **70+ validation rules** covering best practices, schema constraints, and business logic for Fabric resources.

## Features

### 🎯 70+ Validation Rules

**Business Logic Rules (17)** - Custom governance and best practice rules:
- ✅ Workspace capacity requirements
- ✅ Role assignment validation across workspaces, domains, gateways, and pipelines
- ✅ Git integration validation (Azure DevOps & GitHub)
- ✅ Deployment pipeline stage validation (count, naming, descriptions)
- ✅ Domain contributor scope validation
- ✅ Capacity region availability checks

**API Spec Rules (58)** - Auto-generated from Fabric API specifications:
- ✅ Display name length validation (123-256 chars depending on resource)
- ✅ Description length validation (256-1024 chars)
- ✅ Enum value validation (roles, connectivity types, privacy levels)
- ✅ Spark environment settings validation
- ✅ Gateway configuration validation

## Requirements

- TFLint v0.42+
- Go v1.24+
- Terraform v1.0+

## Installation

### From Release

```sh
# Add to .tflint.hcl
plugin "fabric" {
  enabled = true
  version = "0.1.0"
  source  = "github.com/RuneORakeie/tflint-ruleset-fabric"
  signing_key = <<-KEY
-----BEGIN PGP PUBLIC KEY BLOCK-----

mDMEaPIddRYJKwYBBAHaRw8BAQdA97T9Pw1tHwmuoCHx5Od2jRErjYQLMkGtjzHa
n/OtVO+0LlJ1bmUgT3ZsaWVuIFJha2VpZSA8cnVuZS5yYWtlaWVAdGlldG9ldnJ5
LmNvbT6ImQQTFgoAQRYhBDnXUAR+SpStNuTHwcE/gmmRiGC6BQJo8h11AhsDBQkD
wmcABQsJCAcCAiICBhUKCQgLAgQWAgMBAh4HAheAAAoJEME/gmmRiGC6clQBAN4b
xNef3oASq2YWhyDJGoKavDGsTr9CJZkZ66EBPx+zAQDhrFEpylND2qtxQG8x0O5C
cXAcWbXr1DcxbJMEngtZCLg4BGjyHXUSCisGAQQBl1UBBQEBB0B7RETmxgBs7+xU
+MC1lY8ns3V2jnkf/JDAqDymDCcEcgMBCAeIfgQYFgoAJhYhBDnXUAR+SpStNuTH
wcE/gmmRiGC6BQJo8h11AhsMBQkDwmcAAAoJEME/gmmRiGC6XP4A/2x6serOfr3l
8uRNXHoxoKfKAxtcT4JjXGb3NMuc1DfVAP4zy+0wZb8/6jpRriVphm63Y0LjVFv0
3mMMNhrYbuc0AA==
=4jLr
-----END PGP PUBLIC KEY BLOCK-----
KEY  
}

# Then run
tflint --init
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
│   │   ├── fabric_*_invalid_*.go
│   │   └── generated_rules_test.go     # Tests
│   ├── business_logic_rules_test.go    # Tests
├── docs/
│   └── rules/                          # Rule documentation
│       └── *.md
├── tools/
│   └── apispec-rule-gen/               # Code generator
│   └── apispec-mapping-gen/            # Mapping from Fabric API Specs generator
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
