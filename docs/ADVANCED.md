# Advanced Topics & Configuration

## Multi-Environment Setup

### Development Configuration (.tflint.dev.hcl)

```hcl
plugin "fabric" {
  enabled = true
}

# Loose rules for development
rule "fabric_workspace_naming" {
  enabled = false
}

rule "fabric_workspace_capacity_required" {
  enabled = false
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
  enabled = false
}
```

Usage:
```bash
tflint -c .tflint.dev.hcl --recursive
```

### Production Configuration (.tflint.prod.hcl)

```hcl
plugin "fabric" {
  enabled = true
}

# Strict rules for production
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
  enabled = true
}
```

Usage:
```bash
tflint -c .tflint.prod.hcl --recursive
```

## Performance Optimization

### Reduce Rule Count

Only enable necessary rules for faster execution:

```hcl
plugin "fabric" {
  enabled = true
}

# Only critical rules
rule "fabric_workspace_capacity_required" {
  enabled = true
}

rule "fabric_role_assignment_principal_required" {
  enabled = true
}

rule "fabric_git_integration_provider_valid" {
  enabled = true
}

# Disable non-critical rules
rule "fabric_workspace_naming" {
  enabled = false
}

rule "fabric_workspace_description_required" {
  enabled = false
}

rule "fabric_capacity_region_valid" {
  enabled = false
}
```

### Run Specific Directory

```bash
# Only check specific directory
tflint terraform/prod/

# Skip certain paths
tflint --recursive --ignore-module terraform/legacy/
```

## CI/CD Integration Patterns

### GitHub Actions - Multiple Jobs

```yaml
jobs:
  lint-dev:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: tflint -c .tflint.dev.hcl --recursive terraform/dev

  lint-prod:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: tflint -c .tflint.prod.hcl --recursive terraform/prod
```

### GitLab CI

```yaml
tflint:
  image: golang:1.25-alpine
  script:
    - apk add --no-cache git
    - wget https://github.com/terraform-linters/tflint/releases/download/v0.52.0/tflint_linux_amd64.zip
    - unzip tflint_linux_amd64.zip
    - ./tflint --init
    - ./tflint --recursive
```

### Pre-commit Hook

Create `.git/hooks/pre-commit`:

```bash
#!/bin/bash
set -e

echo "Running TFLint..."
tflint --recursive

if [ $? -ne 0 ]; then
  echo "TFLint failed. Fix issues before committing."
  exit 1
fi

echo "TFLint passed ✓"
```

Make executable:
```bash
chmod +x .git/hooks/pre-commit
```

## Terraform Module Best Practices

### Reusable Module with TFLint

```hcl
# modules/fabric_workspace/main.tf
variable "workspace_config" {
  type = object({
    name        = string
    description = string
    capacity_id = string
    tags        = map(string)
  })

  validation {
    condition = can(regex("^[a-z0-9\\-]{3,50}$", var.workspace_config.name))
    error_message = "Workspace name must be 3-50 lowercase alphanumeric with hyphens"
  }

  validation {
    condition     = length(var.workspace_config.description) > 0
    error_message = "Workspace description is required"
  }

  validation {
    condition     = length(var.workspace_config.capacity_id) > 0
    error_message = "Capacity ID is required"
  }
}

resource "fabric_workspace" "workspace" {
  display_name = var.workspace_config.name
  description  = var.workspace_config.description
  capacity_id  = var.workspace_config.capacity_id
  tags         = var.workspace_config.tags
}

output "workspace_id" {
  value = fabric_workspace.workspace.id
}
```

### Module Usage

```hcl
# main.tf
module "analytics_workspace" {
  source = "./modules/fabric_workspace"

  workspace_config = {
    name        = "analytics-hub"
    description = "Central analytics platform"
    capacity_id = fabric_capacity.prod.id
    tags = {
      environment = "production"
      team        = "analytics"
    }
  }
}
```

## Advanced Patterns

### Dynamic Resource Configuration

```hcl
locals {
  workspaces = {
    dev = {
      name        = "dev-workspace"
      description = "Development environment"
      capacity_id = fabric_capacity.dev.id
    }
    prod = {
      name        = "prod-workspace"
      description = "Production environment"
      capacity_id = fabric_capacity.prod.id
    }
  }
}

resource "fabric_workspace" "workspaces" {
  for_each = local.workspaces

  display_name = each.value.name
  description  = each.value.description
  capacity_id  = each.value.capacity_id

  tags = {
    environment = each.key
  }
}
```

### Role Assignment with Dynamic Principals

```hcl
locals {
  role_assignments = {
    admin = {
      role           = "Admin"
      principal_type = "User"
      principal_ids  = ["user1@company.com", "user2@company.com"]
    }
    editor = {
      role           = "Editor"
      principal_type = "Group"
      principal_ids  = ["editors-group-id"]
    }
  }
}

resource "fabric_workspace_role_assignment" "assignments" {
  for_each = merge([
    for role, config in local.role_assignments : {
      for principal_id in config.principal_ids :
      "${role}-${principal_id}" => {
        role           = config.role
        principal_type = config.principal_type
        principal_id   = principal_id
      }
    }
  ]...)

  workspace_id   = fabric_workspace.prod.id
  principal_id   = each.value.principal_id
  role           = each.value.role
  principal_type = each.value.principal_type
}
```

## Debugging and Troubleshooting

### Enable Debug Logging

```bash
TF_LOG=debug tflint --recursive 2>&1 | tee tflint-debug.log
```

### Filter Log Output

```bash
TF_LOG=debug tflint --only fabric_workspace_naming 2>&1 | grep -i workspace
```

### Test Single Rule

```bash
tflint --only fabric_workspace_naming terraform/main.tf
```

### Verbose Output

```bash
tflint --recursive -f json | jq '.[] | select(.rule.name == "fabric_workspace_naming")'
```

## Custom Configuration Examples

### Minimal Configuration (Errors Only)

```hcl
plugin "fabric" {
  enabled = true
}

rule "fabric_workspace_capacity_required" {
  enabled = true
}

rule "fabric_role_assignment_principal_required" {
  enabled = true
}

rule "fabric_git_integration_provider_valid" {
  enabled = true
}
```

### Recommended Configuration (Default)

```hcl
plugin "fabric" {
  enabled = true
}

# All rules with defaults
```

### Maximum Strictness

```hcl
plugin "fabric" {
  enabled = true
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
  enabled = true
}
```

## Performance Benchmarking

### Measure Execution Time

```bash
time tflint --recursive
```

### Profile by Rule

```bash
for rule in fabric_workspace_naming fabric_workspace_capacity_required \
            fabric_workspace_description_required fabric_role_assignment_principal_required \
            fabric_git_integration_provider_valid fabric_capacity_region_valid; do
  echo "Testing $rule..."
  time tflint --only $rule --recursive
done
```

### Memory Usage

```bash
# Linux
/usr/bin/time -v tflint --recursive

# macOS
time tflint --recursive
```

## Integration with Other Tools

### Terraform Cloud/Enterprise

Set environment variable before run:
```bash
export TFE_TOKEN=your-token
tflint --recursive
```

### Pre-commit Framework

`.pre-commit-config.yaml`:
```yaml
repos:
  - repo: https://github.com/terraform-linters/pre-commit-hooks
    rev: v1.0.0
    hooks:
      - id: tflint
        args: ['--init', '--recursive']
```

### IDE Integration

**VS Code Settings (.vscode/settings.json)**:
```json
{
  "[hcl]": {
    "editor.defaultFormatter": "hashicorp.terraform",
    "editor.formatOnSave": true
  },
  "terraform.lintPath": "tflint",
  "terraform.lintConfig": ".tflint.hcl"
}
```

## Troubleshooting Advanced Issues

### Plugin Version Mismatch

```bash
# Check versions
tflint -v

# Reinstall
make clean install

# Verify
tflint -v | grep fabric
```

### Configuration Not Applying

```bash
# Check config syntax
hcl-inspect .tflint.hcl

# Use explicit path
tflint -c .tflint.prod.hcl --recursive

# Validate HCL
terraform fmt .tflint.hcl
```

### Rules Not Firing on Specific Files

```bash
# Check file is being parsed
tflint --format json | jq '.[] | select(.filename == "your-file.tf")'

# Run with debug
TF_LOG=debug tflint --only fabric_workspace_naming --recursive 2>&1 | grep "your-file"
```

## Best Practices

1. **Use environment-specific configs** for dev/prod differences
2. **Enable rules incrementally** rather than all at once
3. **Document custom configurations** in your team
4. **Monitor performance** with larger projects
5. **Review logs regularly** for patterns
6. **Update regularly** to get new rules and bug fixes
7. **Test before committing** with pre-commit hooks
8. **Integrate early** in CI/CD pipeline

## See Also

- [README.md](README.md) - Main documentation
- [QUICK_START.md](QUICK_START.md) - Getting started
- [TROUBLESHOOTING.md](docs/TROUBLESHOOTING.md) - Common issues
- [CONTRIBUTING.md](CONTRIBUTING.md) - Development
