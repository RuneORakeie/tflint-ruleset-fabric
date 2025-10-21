# Architecture Overview

## Project Structure

```
tflint-ruleset-fabric/
├── .github/workflows/          # CI/CD pipelines
│   ├── tflint.yml             # TFLint validation
│   ├── test.yml               # Unit tests and linting
│   └── release.yml            # Automated releases
├── rules/                      # Rule implementations
│   ├── fabric_workspace_naming.go
│   ├── fabric_workspace_capacity.go
│   ├── fabric_workspace_description.go
│   ├── fabric_role_assignment_principals.go
│   ├── fabric_git_integration_validation.go
│   ├── fabric_capacity_region.go
│   └── rules_test.go          # Test suite
├── examples/                   # Terraform examples
│   ├── valid_examples.tf
│   ├── invalid_examples.tf
│   └── .tflint.hcl
├── docs/                      # Documentation
│   ├── ARCHITECTURE.md        # This file
│   ├── TROUBLESHOOTING.md
│   └── rules/
├── main.go                    # Plugin entry point
├── go.mod                     # Dependencies
├── Makefile                   # Build automation
├── .tflint.hcl               # Default config
├── .golangci.yml             # Linter config
├── .goreleaser.yml           # Release config
├── README.md                 # Main documentation
├── QUICK_START.md            # Quick start guide
├── CONTRIBUTING.md           # Development guide
└── LICENSE                   # MPL 2.0
```

## How It Works

### 1. Plugin Loading

When TFLint starts, it:
1. Loads the Fabric plugin binary
2. Calls `main()` which serves the plugin
3. Plugin registers all 6 rules
4. TFLint becomes ready to check Terraform files

### 2. Rule Execution Flow

```
User runs: tflint --recursive
    ↓
TFLint Core loads Terraform files
    ↓
For each rule:
    - Rule.Check() is called
    - Gets resources by type (fabric_workspace, etc)
    - Validates each resource
    - Emits issues if violations found
    ↓
TFLint aggregates all issues
    ↓
Results formatted and displayed
    ↓
Exit with appropriate code
```

### 3. Rule Pattern

Each rule implements the `tflint.Rule` interface:

```go
type Rule interface {
    Name() string                          // Rule identifier
    Enabled() bool                         // Default enabled status
    Severity() string                      // ERROR, WARNING, NOTICE
    Link() string                          // Documentation URL
    Check(runner tflint.Runner) error     // Validation logic
}
```

### 4. Rule Lifecycle

```
Rule Instantiation (NewFabricXXX())
    ↓
Name/Enabled/Severity/Link setup
    ↓
Check() method called
    ↓
runner.GetResourcesByType()
    ↓
Iterate resources
    ↓
Validate attributes
    ↓
runner.EmitIssue() if invalid
    ↓
Return nil
```

## Data Flow

### Input
- Terraform `.tf` files parsed by TFLint core
- HCL structure extracted
- Resources organized by type

### Processing
- Each rule receives runner
- Runner queries resources
- Rule validates attributes
- Issues created if validation fails

### Output
- Issues aggregated
- Formatted (console, JSON, SARIF)
- Displayed to user
- Exit code set

## The 6 Rules

### 1. fabric_workspace_naming
- **Type**: Validation
- **Scope**: `fabric_workspace` resources
- **Check**: Validates display_name attribute
- **Validation**: Regex pattern (3-50 chars, lowercase, alphanumeric + hyphens)
- **Severity**: WARNING

### 2. fabric_workspace_capacity_required
- **Type**: Requirement
- **Scope**: `fabric_workspace` resources
- **Check**: Validates capacity_id exists
- **Validation**: Attribute must be present and non-empty
- **Severity**: ERROR

### 3. fabric_workspace_description_required
- **Type**: Requirement
- **Scope**: `fabric_workspace` resources
- **Check**: Validates description exists
- **Validation**: Attribute must be present and non-empty
- **Severity**: WARNING

### 4. fabric_role_assignment_principal_required
- **Type**: Requirement
- **Scope**: `fabric_workspace_role_assignment` resources
- **Check**: Validates principal_id exists
- **Validation**: Attribute must be present
- **Severity**: ERROR

### 5. fabric_git_integration_provider_valid
- **Type**: Validation
- **Scope**: `fabric_workspace_git_connection` resources
- **Check**: Validates git_provider_type
- **Validation**: Must be one of approved values
- **Severity**: ERROR

### 6. fabric_capacity_region_valid
- **Type**: Validation
- **Scope**: `fabric_capacity` resources
- **Check**: Validates region
- **Validation**: Must be valid Azure region
- **Severity**: WARNING
- **Status**: Disabled by default

## Build & Release

### Local Development

```bash
make build           # Build plugin
make install         # Install to ~/.tflint.d/plugins
make test            # Run tests
make lint            # Run linter
make fmt             # Format code
```

### CI/CD Pipeline

**On Push/PR:**
- test.yml runs tests on multiple OS/Go versions
- Code coverage uploaded to Codecov
- GoLint checks code quality

**On Tag (v*.*.*):**
- release.yml builds multi-platform binaries
- Signs with GPG
- Creates GitHub release
- Uploads checksums

## Testing Strategy

### Unit Tests
- Located in `rules/rules_test.go`
- Uses TFLint test helper
- Tests valid and invalid configurations
- 30+ test cases
- Target: 85%+ coverage

### Test Execution

```go
runner := helper.TestRunner(t, cases)
rule := NewFabricXXX()
err := rule.Check(runner)
```

### Test Coverage

```bash
go test -cover ./...
go test -coverprofile=coverage.txt ./...
go tool cover -html=coverage.txt
```

## Dependencies

### Direct Dependencies

- `github.com/terraform-linters/tflint-plugin-sdk v0.20.1`
  - Provides Rule interface
  - Provides TestRunner helper
  - Provides HCL parsing utilities

### Go Version

- Minimum: Go 1.25
- Tested: 1.24, 1.25
- Reason: Latest TFLint SDK requires 1.25+

## Configuration

### .tflint.hcl

```hcl
plugin "fabric" {
  enabled = true
  version = "0.1.0"
  source  = "github.com/RuneORakeie/tflint-ruleset-fabric"
}

rule "fabric_workspace_naming" {
  enabled = true
}

rule "fabric_capacity_region_valid" {
  enabled = false
}
```

### Environment-Specific Configs

```bash
tflint -c .tflint.dev.hcl   # Loose rules
tflint -c .tflint.prod.hcl  # Strict rules
```

## Performance

### Resource Usage

- Memory: ~20MB
- Startup: ~50ms
- Per-rule execution: ~100ms per 10 resources
- Typical project: <1 second

### Optimization Strategies

- Rules execute once per resource type
- Early exits on invalid attributes
- No external API calls
- Regex pattern compiled once

## Security

### Code Security

- GoSec linting enabled
- No hardcoded credentials
- No external dependencies beyond TFLint SDK
- Input validation on all attributes

### Release Security

- GPG signed releases
- SHA256 checksums
- Verified on download
- SBOM generation (planned)

## Extensibility

### Adding New Rules

1. Create new file in `rules/`
2. Implement `tflint.Rule` interface
3. Register in `main.go`
4. Add tests
5. Document

### Example Structure

```go
type MyNewRule struct {
    tflint.DefaultRule
}

func NewMyNewRule() *MyNewRule {
    return &MyNewRule{}
}

func (r *MyNewRule) Name() string {
    return "fabric_my_rule"
}

func (r *MyNewRule) Enabled() bool {
    return true
}

func (r *MyNewRule) Severity() string {
    return tflint.WARNING
}

func (r *MyNewRule) Link() string {
    return "https://docs.example.com"
}

func (r *MyNewRule) Check(runner tflint.Runner) error {
    // Implementation
    return nil
}
```

## Maintenance

### Release Cycle

- Minor releases: Monthly
- Patch releases: As needed
- Major releases: Quarterly
- Security: Immediately

### Version Numbering

- MAJOR.MINOR.PATCH (semantic versioning)
- 0.1.0 = Initial release
- Tags: v0.1.0, v0.2.0, etc.

## Troubleshooting Architecture Issues

### Plugin Not Loading

Check:
- Plugin binary in `~/.tflint.d/plugins`
- Correct module path
- Compatible TFLint version

### Rule Not Firing

Check:
- Rule enabled in config
- Resource type matches
- Attribute names correct
- Test cases pass

### Performance Issues

Check:
- Number of resources
- Resource type filtering
- Rule complexity
- External dependencies

## Related Documentation

- [README.md](../README.md) - Feature overview
- [QUICK_START.md](../QUICK_START.md) - Getting started
- [CONTRIBUTING.md](../CONTRIBUTING.md) - Development
- [TROUBLESHOOTING.md](TROUBLESHOOTING.md) - Problem solving
