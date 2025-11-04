# TFLint Ruleset Configuration Guide

This guide explains how to configure the TFLint ruleset for Microsoft Fabric in your `.tflint.hcl` file.

## Table of Contents

- [Basic Configuration](#basic-configuration)
- [Plugin Configuration Block](#plugin-configuration-block)
- [Rule Presets](#rule-presets)
- [Naming Conventions](#naming-conventions)
- [Individual Rule Configuration](#individual-rule-configuration)
- [Examples](#examples)

---

## Basic Configuration

Add the plugin to your `.tflint.hcl` file:

```hcl
plugin "fabric" {
  enabled = true
  version = "0.1.0"
  source  = "github.com/RuneORakeie/tflint-ruleset-fabric"
}
```

Then run `tflint --init` to install the plugin.

---

## Plugin Configuration Block

The `rule "fabric"` block allows you to configure plugin-wide settings:

```hcl
rule "fabric" {
  enabled = true
  
  # Select which rule preset to use
  preset = "recommended"  # Options: minimal, recommended, all (default)
  
  # Configure naming conventions for all resources with display_name
  naming_convention {
    # General settings (apply to all resources)
    pattern = "^(dev|test|prod)-[a-z0-9-]+$"
    format  = "kebab-case"  # Options: snake_case, kebab-case, PascalCase, camelCase
    prefix  = "pcs-"
    suffix  = "-rti"
    
    # Resource-specific overrides (optional)
    pattern_overrides = {
      fabric_workspace   = "^(dev|test|prod)-ws-[a-z0-9-]+$"
      fabric_lakehouse   = "^lh-[a-z0-9-]+$"
      fabric_eventhouse  = "^eh-[a-z0-9-]+$"
    }
    
    prefix_overrides = {
      fabric_eventstream = "es-"
      fabric_eventhouse  = "eh-"
      fabric_lakehouse   = "lh-"
    }
    
    suffix_overrides = {
      fabric_workspace = "-weu"
    }
  }
}
```

---

## Rule Presets

The ruleset provides three presets that control which rules are enabled. You can select a preset using the `preset` option in the `rule "fabric"` block.

### Available Presets

#### Minimal
Core validation rules only - catches critical issues.

**Includes:**
- `fabric_workspace_capacity` - Ensures workspace has a capacity assigned
- `fabric_capacity_region_valid` - Validates capacity region

**Configuration:**
```hcl
plugin "fabric" {
  enabled = true
  version = "0.1.0"
  source  = "github.com/RuneORakeie/tflint-ruleset-fabric"
}

rule "fabric" {
  enabled = true
  preset  = "minimal"
}
```

#### Recommended
Minimal + common best practices - recommended for most projects (good for CI/CD).

**Adds to minimal:**
- `fabric_workspace_role_assignment_role` - Validates role assignments
- `fabric_deployment_pipeline_stages_count` - Enforces pipeline stage limits
- `fabric_domain_contributors_scope` - Validates domain contributor scope

**Configuration:**
```hcl
plugin "fabric" {
  enabled = true
  version = "0.1.0"
  source  = "github.com/RuneORakeie/tflint-ruleset-fabric"
}

rule "fabric" {
  enabled = true
  preset  = "recommended"
}
```

#### All
All business logic rules + all generated API validation rules (**default**).

**Adds to recommended:**
- All Git integration validation rules (7 rules)
- Description and display name recommendations
- Role assignment best practices
- Pipeline naming and description rules
- All 58 generated API spec validation rules

**Configuration:**
```hcl
plugin "fabric" {
  enabled = true
  version = "0.1.0"
  source  = "github.com/RuneORakeie/tflint-ruleset-fabric"
}

# All rules are enabled by default (preset = "all")
# No additional configuration needed, or explicitly set:
rule "fabric" {
  enabled = true
  preset  = "all"
}
```

---

## Naming Conventions

You can enforce naming conventions for all Fabric resources that have a `display_name` attribute using the `naming_convention` block. Naming conventions apply globally by default, with optional resource-specific overrides.

### Using Regex Patterns

Apply a global pattern to all resources:

```hcl
rule "fabric" {
  enabled = true
  
  naming_convention {
    pattern = "^(dev|test|prod)-[a-z0-9-]+$"
  }
}
```

Or use resource-specific patterns:

```hcl
rule "fabric" {
  enabled = true
  
  naming_convention {
    # Global pattern (applied to all resources by default)
    pattern = "^[a-z0-9-]+$"
    
    # Resource-specific overrides
    pattern_overrides = {
      fabric_workspace   = "^(dev|test|prod)-ws-[a-z0-9-]+$"
      fabric_lakehouse   = "^lh-[a-z0-9-]+$"
      fabric_eventhouse  = "^eh-[a-z0-9-]+$"
      fabric_eventstream = "^es-[a-z0-9-]+$"
    }
  }
}
```

### Using Predefined Formats

```hcl
rule "fabric" {
  enabled = true
  
  naming_convention {
    format = "kebab-case"  # Enforces kebab-case for all resources
  }
}
```

**Available formats:**
- `snake_case` - lowercase with underscores (e.g., `my_workspace`)
- `kebab-case` - lowercase with hyphens (e.g., `my-workspace`)
- `PascalCase` - uppercase first letter of each word (e.g., `MyWorkspace`)
- `camelCase` - lowercase first letter, uppercase others (e.g., `myWorkspace`)

### Using Prefix/Suffix Requirements

Apply global prefix/suffix requirements:

```hcl
rule "fabric" {
  enabled = true
  
  naming_convention {
    prefix = "pcs-"
    suffix = "-rti"
  }
}
```

Or use resource-specific overrides:

```hcl
rule "fabric" {
  enabled = true
  
  naming_convention {
    # Global prefix (applied to all resources by default)
    prefix = "pcs-"
    
    # Resource-specific overrides
    prefix_overrides = {
      fabric_workspace   = "ws-"
      fabric_lakehouse   = "lh-"
      fabric_eventstream = "es-"
      fabric_eventhouse  = "eh-"
    }
    
    suffix_overrides = {
      fabric_workspace = "-weu"  # West Europe
    }
  }
}
```

This will require:
- Workspaces: `ws-*-weu` (e.g., `ws-analytics-weu`)
- Lakehouses: `lh-*` (e.g., `lh-sales-data`)
- Eventstreams: `es-*` (e.g., `es-iot-telemetry`)
- Other resources: `pcs-*` (global prefix applied)

### Combining Rules

```hcl
rule "fabric" {
  enabled = true
  preset  = "recommended"
  
  naming_convention {
    # Apply kebab-case format to all resources
    format = "kebab-case"
    
    # Add environment prefix to all resources
    pattern = "^(dev|test|prod)-[a-z0-9-]+$"
    
    # But use specific prefixes for certain resource types
    prefix_overrides = {
      fabric_workspace = "ws-"
      fabric_lakehouse = "lh-"
    }
  }
}
```

---

## Individual Rule Configuration

You can enable or disable individual rules regardless of the preset:

### Enable a Specific Rule

```hcl
rule "fabric_item_description_recommended" {
  enabled = true
}
```

### Disable a Specific Rule

```hcl
rule "fabric_workspace_git_provider_type" {
  enabled = false
}
```

### Common Rule Configurations

#### Disable All Generated Rules (Keep Business Logic Only)

```hcl
# Disable generated rules individually
rule "fabric_activator_invalid_description" {
  enabled = false
}

rule "fabric_apache_airflow_job_invalid_description" {
  enabled = false
}

# ... (repeat for all generated rules)
```

#### Enable Only Git Integration Rules

```hcl
config {
  disabled_by_default = true
}

rule "fabric_workspace_git_provider_type" {
  enabled = true
}

rule "fabric_workspace_git_initialization_strategy_valid" {
  enabled = true
}

rule "fabric_workspace_git_directory_name_format" {
  enabled = true
}

rule "fabric_workspace_git_credentials_source_valid" {
  enabled = true
}

rule "fabric_workspace_git_azdo_attributes_required" {
  enabled = true
}

rule "fabric_workspace_git_github_attributes_required" {
  enabled = true
}

rule "fabric_workspace_git_string_lengths" {
  enabled = true
}
```

---

## Examples

### Example 1: Minimal Setup for CI/CD

Catch only critical issues in your CI pipeline:

```hcl
plugin "fabric" {
  enabled = true
  version = "0.1.0"
  source  = "github.com/RuneORakeie/tflint-ruleset-fabric"
}

rule "fabric" {
  enabled = true
  preset  = "minimal"
}
```

### Example 2: Recommended for Development

Balance between thoroughness and noise:

```hcl
plugin "fabric" {
  enabled = true
  version = "0.1.0"
  source  = "github.com/RuneORakeie/tflint-ruleset-fabric"
}

rule "fabric" {
  enabled = true
  preset  = "recommended"
}
```

### Example 3: All Rules (Default)

Enable all validation for maximum quality:

```hcl
plugin "fabric" {
  enabled = true
  version = "0.1.0"
  source  = "github.com/RuneORakeie/tflint-ruleset-fabric"
}

rule "fabric" {
  enabled = true
  preset  = "all"  # This is the default
}
```

### Example 4: Custom Naming with Recommended Preset

Enforce naming conventions with recommended rules:

```hcl
plugin "fabric" {
  enabled = true
  version = "0.1.0"
  source  = "github.com/RuneORakeie/tflint-ruleset-fabric"
}

rule "fabric" {
  enabled = true
  preset  = "recommended"
  
  naming_convention {
    format  = "kebab-case"
    pattern = "^(dev|test|prod)-[a-z0-9-]+$"
    
    prefix_overrides = {
      fabric_workspace = "ws-"
      fabric_lakehouse = "lh-"
    }
  }
}
```

### Example 5: Environment-Specific Configuration

Different naming patterns for different environments:

```hcl
plugin "fabric" {
  enabled = true
  version = "0.1.0"
  source  = "github.com/RuneORakeie/tflint-ruleset-fabric"
}

rule "fabric" {
  enabled = true
  preset  = "all"
  
  naming_convention {
    # Require dev/test/prod prefix on all resources
    pattern = "^(dev|test|prod)-[a-z0-9-]+$"
    
    # Add resource type prefixes
    prefix_overrides = {
      fabric_workspace   = "ws-"
      fabric_lakehouse   = "lh-"
      fabric_eventhouse  = "eh-"
      fabric_eventstream = "es-"
    }
    
    # Add region suffix to workspaces
    suffix_overrides = {
      fabric_workspace = "-weu"  # West Europe
    }
  }
}
```

### Example 6: Combine with Individual Rule Overrides

Use a preset but override specific rules:

```hcl
plugin "fabric" {
  enabled = true
  version = "0.1.0"
  source  = "github.com/RuneORakeie/tflint-ruleset-fabric"
}

rule "fabric" {
  enabled = true
  preset  = "recommended"
}

# Disable a specific rule from the preset
rule "fabric_deployment_pipeline_stages_count" {
  enabled = false
}

# Enable an additional rule not in the preset
rule "fabric_item_description_recommended" {
  enabled = true
}
```

---

## Rule Severity Levels

All rules in this ruleset use the following severity levels:

- **ERROR**: Critical issues that will likely cause deployment failures
- **WARNING**: Best practice violations that should be addressed
- **NOTICE**: Informational suggestions for improvement

TFLint will exit with a non-zero status code if any ERROR-level issues are found.

---

## Further Reading

- [Rule Documentation](rules/) - Detailed documentation for each rule
- [TFLint Configuration](https://github.com/terraform-linters/tflint/blob/master/docs/user-guide/config.md) - Official TFLint configuration guide
- [Plugin Configuration](https://github.com/terraform-linters/tflint/blob/master/docs/user-guide/plugins.md) - TFLint plugin configuration
