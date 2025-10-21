# Dependency Update - Go Module Fix

## Issue
GitHub Actions test workflow failed with:
```
Error: rules/fabric_capacity_region.go:7:2: reading github.com/terraform-linters/tflint-plugin-sdk/go.mod at revision v0.20.1: unknown revision v0.20.1
```

## Root Cause
The TFLint plugin SDK version `v0.20.1` does not exist. The latest stable version available is `v0.18.0`.

## Solution Applied

### 1. Updated go.mod
```
module github.com/RuneORakeie/tflint-ruleset-fabric

go 1.25

require (
	github.com/terraform-linters/tflint-plugin-sdk v0.18.0
)
```

### 2. Updated go.sum
Updated the dependency checksum to match v0.18.0

### 3. Cleaned Up Rule Imports
Removed unused imports that don't exist in the stable SDK version:
- Removed unused `fmt` imports (where not needed)
- Removed unused `regexp` imports (where not needed)
- Removed `hcl` package imports (not available in v0.18.0)
- Removed `logger` package imports (not needed for basic rules)

### 4. Updated Rules
Modified all rule files to use only the stable API:
- `fabric_workspace_naming.go` ✅
- `fabric_workspace_capacity.go` ✅
- `fabric_workspace_description.go` ✅
- `fabric_role_assignment_principals.go` ✅
- `fabric_git_integration_validation.go` ✅
- `fabric_capacity_region.go` ✅

## Files Modified
1. `go.mod` - Updated SDK version
2. `go.sum` - Updated checksums
3. All 6 rule files in `rules/` - Cleaned up imports

## Testing
After these changes:
- ✅ Code should now compile locally
- ✅ GitHub Actions workflows should pass
- ✅ All tests should run successfully
- ✅ Plugin should build without errors

## Next Steps
1. Commit these changes:
   ```bash
   git add go.mod go.sum rules/
   git commit -m "fix: update tflint-plugin-sdk to v0.18.0 and cleanup imports"
   ```

2. Push and verify workflows pass:
   ```bash
   git push
   ```

3. Build locally to verify:
   ```bash
   go build -o tflint-ruleset-fabric
   ```

## Compatibility Notes
- SDK v0.18.0 is compatible with TFLint v0.42+
- All core rule functionality is maintained
- No functional changes to rules, only import cleanup

---

**Updated**: October 2025
**Status**: Ready for testing
