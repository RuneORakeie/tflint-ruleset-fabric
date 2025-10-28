# TFLint Fabric Ruleset - Project Structure & Contributing Guide

## Complete Project Structure

```
tflint-ruleset-fabric/
├── .github/
│   ├── workflows/
│   │   ├── tflint.yml              # CI/CD validation
│   │   ├── release.yml             # Release automation
│   │   └── test.yml                # Unit tests
│   └── ISSUE_TEMPLATE/
│       ├── bug_report.md
│       └── feature_request.md
├── rules/
│   ├── fabric_workspace_naming.go           # Workspace naming rule
│   ├── fabric_workspace_capacity.go         # Capacity requirement rule
│   ├── fabric_workspace_description.go      # Description requirement rule
│   ├── fabric_role_assignments.go           # Role assignment validation
│   ├── fabric_git_integration.go            # Git provider validation
│   ├── fabric_capacity_region.go            # Capacity region validation
│   ├── rules_test.go                        # Rule tests
│   └── generator/
│       └── main.go                          # Documentation generator
├── examples/
│   ├── valid/
│   │   ├── main.tf                          # Valid configurations
│   │   ├── variables.tf
│   │   ├── outputs.tf
│   │   └── .tflint.hcl
│   ├── invalid/
│   │   ├── main.tf                          # Invalid configurations
│   │   ├── variables.tf
│   │   └── .tflint.hcl
│   ├── complete/
│   │   ├── main.tf                          # Full example
│   │   ├── variables.tf
│   │   ├── outputs.tf
│   │   ├── terraform.tfvars.example
│   │   └── .tflint.hcl
│   └── README.md
├── docs/
│   ├── rules/
│   │   ├── README.md                        # Rules documentation
│   │   ├── fabric_workspace_naming.md
│   │   ├── fabric_workspace_capacity.md
│   │   ├── fabric_workspace_description.md
│   │   ├── fabric_role_assignments.md
│   │   ├── fabric_git_integration.md
│   │   └── fabric_capacity_region.md
│   ├── ARCHITECTURE.md
│   ├── DEVELOPMENT.md
│   └── TROUBLESHOOTING.md
├── main.go                                  # Plugin entry point
├── go.mod                                   # Go module definition
├── go.sum                                   # Go dependencies
├── Makefile                                 # Build automation
├── .tflint.hcl                              # Default config
├── .goreleaser.yml                          # Release configuration
├── .golangci.yml                            # Linting configuration
├── LICENSE                                  # MPL 2.0 license
├── README.md                                # Main documentation
├── CONTRIBUTING.md                          # Contribution guide
├── CHANGELOG.md                             # Version history
└── SECURITY.md                              # Security policy
```

## Quick Start - Setup for Development

### 1. Clone and Setup

```bash
# Clone the repository
git clone https://github.com/RuneORakeie/tflint-ruleset-fabric
cd tflint-ruleset-fabric

# Install dependencies
go mod download

# Verify setup
go mod tidy
```

### 2. Build Locally

```bash
# Build the plugin
make build

# Install for testing
make install

# Verify installation
tflint -v
# Should output: + ruleset.fabric (0.1.0)
```

### 3. Run Tests

```bash
# Run all tests
make test

# Run with coverage
go test -v -cover ./...

# Run specific test
go test -v -run TestFabricWorkspaceNaming ./rules
```

### 4. Format and Lint Code

```bash
# Format code
make fmt

# Run linter
make lint
```

## Creating a New Rule

### Step 1: Understand the Template

All rules follow this structure:

```go
type MyNewRule struct {
    tflint.DefaultRule
}

func NewMyNewRule() *MyNewRule {
    return &MyNewRule{}
}

func (r *MyNewRule) Name() string {
    return "fabric_my_rule_name"
}

func (r *MyNewRule) Enabled() bool {
    return true  // or false if disabled by default
}

func (r *MyNewRule) Severity() string {
    return tflint.ERROR  // ERROR, WARNING, or NOTICE
}

func (r *MyNewRule) Link() string {
    return "https://learn.microsoft.com/..."
}

func (r *MyNewRule) Check(runner tflint.Runner) error {
    // Rule implementation
    return nil
}
```

### Step 2: Implement Your Rule

Example: Enforce workspace tags

```go
// rules/fabric_workspace_tags_required.go
package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hcl"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

type FabricWorkspaceTags struct {
	tflint.DefaultRule
}

func NewFabricWorkspaceTags() *FabricWorkspaceTags {
	return &FabricWorkspaceTags{}
}

func (r *FabricWorkspaceTags) Name() string {
	return "fabric_workspace_tags_required"
}

func (r *FabricWorkspaceTags) Enabled() bool {
	return true
}

func (r *FabricWorkspaceTags) Severity() string {
	return tflint.WARNING
}

func (r *FabricWorkspaceTags) Link() string {
	return "https://learn.microsoft.com/en-us/fabric/admin/fabric-governance"
}

func (r *FabricWorkspaceTags) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourcesByType("fabric_workspace")
	if err != nil {
		return err
	}

	for _, resource := range resources {
		var tags map[string]interface{}
		err := resource.GetAttribute("tags", &tags)
		if err != nil || len(tags) == 0 {
			if err := runner.EmitIssue(
				r,
				"Workspace should have tags for organization and cost tracking",
				resource.GetNameRange(),
			); err != nil {
						return err
			}
		}
	}

	return nil
}
```

### Step 3: Register in main.go

```go
func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "fabric",
			Version: Version,
			Rules: []tflint.Rule{
				// ... existing rules ...
				rules.NewFabricWorkspaceTags(),  // Add your new rule
			},
		},
	})
}
```

### Step 4: Write Tests

```go
// In rules/rules_test.go
func TestFabricWorkspaceTags(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "valid - tags present",
			Content: `
resource "fabric_workspace" "example" {
  display_name = "test-workspace"
  description  = "Test"
  capacity_id  = "capacity-123"
  
  tags = {
    environment = "prod"
  }
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "invalid - no tags",
			Content: `
resource "fabric_workspace" "example" {
  display_name = "test-workspace"
  description  = "Test"
  capacity_id  = "capacity-123"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewFabricWorkspaceTags(),
					Message: "Workspace should have tags for organization and cost tracking",
				},
			},
		},
	}

	runner := helper.TestRunner(t, cases)
	rule := NewFabricWorkspaceTags()

	if err := rule.Check(runner); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
}
```

### Step 5: Add Documentation

Create `docs/rules/fabric_workspace_tags_required.md`:

```markdown
# fabric_workspace_tags_required

Ensures that fabric_workspace resources have tags defined for organization and tracking purposes.

## Configuration

This rule has no configuration parameters.

## Examples

### Passing

```hcl
resource "fabric_workspace" "example" {
  display_name = "my-workspace"
  description  = "Workspace description"
  capacity_id  = fabric_capacity.prod.id
  
  tags = {
    environment = "production"
    team        = "analytics"
  }
}
```

### Failing

```hcl
resource "fabric_workspace" "example" {
  display_name = "my-workspace"
  description  = "Workspace description"
  capacity_id  = fabric_capacity.prod.id
  # tags are missing
}
```

## Why This Matters

Tags are essential for:
- Cost tracking and billing
- Resource organization
- Compliance and governance
- Automation and filtering

## Related Rules

- `fabric_workspace_naming`
- `fabric_workspace_description_required`
```

## Contributing Workflow

### 1. Fork and Branch

```bash
# Fork the repository on GitHub
git clone https://github.com/RuneORakeie/tflint-ruleset-fabric
cd tflint-ruleset-fabric

# Create feature branch
git checkout -b feat/my-new-rule
```

### 2. Make Changes

```bash
# Create new rule file
touch rules/fabric_my_new_rule.go

# Implement the rule
# Add tests
# Update main.go
```

### 3. Test Locally

```bash
# Run all tests
make test

# Build and install
make install

# Test with examples
cd examples/invalid
tflint
cd ../valid
tflint
```

### 4. Format and Lint

```bash
# Format code
make fmt

# Lint
make lint
```

### 5. Commit and Push

```bash
# Commit with clear message
git add .
git commit -m "feat: add fabric_workspace_tags_required rule"

# Push to your fork
git push origin feat/my-new-rule
```

### 6. Create Pull Request

- Go to GitHub
- Create PR from your branch to `main`
- Fill in PR template
- Wait for CI/CD checks to pass
- Request review

## Pull Request Checklist

- [ ] Code follows project style guide
- [ ] Tests added/updated
- [ ] Documentation updated
- [ ] Changelog entry added
- [ ] No breaking changes (or clearly documented)
- [ ] CI/CD checks pass
- [ ] Commit messages are clear

## Coding Standards

### Go Code Style

```go
// Use clear variable names
resourceID := resource.GetAttribute("id")

// Comment exported functions
// MyFunction does something important
func MyFunction() {}

// Keep functions focused and small
// Aim for <50 lines per function

// Use early returns
if err != nil {
    return err
}

// Group related constants
const (
    GitHub        = "GitHub"
    AzureDevOps   = "Azure DevOps"
    BitbucketCloud = "Bitbucket Cloud"
)
```

### Documentation Style

- Use clear, concise language
- Include examples (both valid and invalid)
- Explain the "why" not just the "what"
- Reference official Microsoft Fabric docs
- Keep links up to date

### Test Requirements

- Minimum 80% code coverage
- Test both valid and invalid cases
- Use descriptive test names
- Test error conditions

## Release Process

Releases follow semantic versioning: MAJOR.MINOR.PATCH

### Patch Release (0.1.1)

Bug fixes and minor improvements

```bash
git tag v0.1.1
git push origin v0.1.1
```

### Minor Release (0.2.0)

New features, backward compatible

### Major Release (1.0.0)

Breaking changes

## Troubleshooting Development

### Plugin not loading

```bash
# Check plugin directory
ls -la ~/.tflint.d/plugins/

# Rebuild and reinstall
make clean install

# Verify
tflint -v
```

### Tests failing

```bash
# Run with verbose output
go test -v ./...

# Check for race conditions
go test -race ./...

# Run specific test
go test -v -run TestName ./rules
```

### Build errors

```bash
# Update dependencies
go get -u

# Tidy modules
go mod tidy

# Clean build
make clean build
```

## Getting Help

- **Documentation**: See `docs/` directory
- **Issues**: Open on GitHub with details
- **Discussions**: Use GitHub Discussions
- **Security**: See SECURITY.md

## Code of Conduct

Be respectful and inclusive. See CODE_OF_CONDUCT.md for details.

## License

By contributing, you agree that your contributions will be licensed under the Mozilla Public License 2.0.