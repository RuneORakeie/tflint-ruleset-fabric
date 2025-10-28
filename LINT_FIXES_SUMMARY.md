# Linting Fixes Summary - Round 2

All 6 remaining golangci-lint issues have been fixed.

## Fixed Issues (Round 2)

### errcheck (3 issues)
1. **rules/fabric_deployment_pipeline_stages_description_length.go:63** - Added error check for `runner.EmitIssue`
2. **rules/fabric_deployment_pipeline_stages_display_name_length.go:63** - Added error check for `runner.EmitIssue`
3. **rules/fabric_domain_contributors_scope.go:59** - Added error check for `runner.EmitIssue`

### goimports (3 issues)
1. **main.go:4** - Fixed import grouping to separate third-party from local packages
2. **rules/fabric_capacity_region.go:6** - Fixed import grouping to separate third-party from local packages
3. **rules/fabric_deployment_pipeline_stages_count.go:6** - Fixed import grouping to separate third-party from local packages

## Import Grouping Standard

Based on `.golangci.yml` configuration with `local-prefixes: github.com/RuneORakeie/tflint-ruleset-fabric`, imports must be grouped as:
1. Standard library imports (e.g., `fmt`, `strings`)
2. Third-party imports (e.g., `github.com/terraform-linters/...`)
3. Blank line separator
4. Local project imports (e.g., `github.com/RuneORakeie/tflint-ruleset-fabric/...`)

## Total Fixes Across Both Rounds

- **errcheck**: 11 issues fixed
- **goimports**: 6 issues fixed  
- **staticcheck**: 1 issue fixed
- **unused**: 2 issues fixed

**Total**: 20 linting issues resolved

All files should now pass golangci-lint checks.
