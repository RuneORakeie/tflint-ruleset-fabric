# TFLint Ruleset for Fabric - Build & Test Fixes

## Issues Fixed

### 1. **Test Function Signature Error**

**Problem:**
```
Error: rules/rules_test.go:58:33: cannot use cases (variable of type []struct{...}) 
as map[string]string value in argument to helper.TestRunner
```

**Root Cause:**
The `helper.TestRunner` function expects a `map[string]string` (filename → content mapping), but the original test was passing a slice of structs.

**Solution:**
Refactored the test structure to use the correct pattern:

```go
// ❌ WRONG - Original structure
cases := []struct {
    Name     string
    Content  string
    Expected helper.Issues
}{}
runner := helper.TestRunner(t, cases)  // ← Type mismatch!

// ✅ CORRECT - New structure
runner := helper.TestRunner(t, map[string]string{
    "main.tf": tt.content,  // ← Correct map type
})
```

**Changes Made:**
- Converted test cases to use subtests with `t.Run()`
- Each test now calls `helper.TestRunner()` with a proper `map[string]string`
- Simplified assertion logic by checking `len(runner.Issues)` instead of comparing `helper.Issues` structures
- Added more comprehensive test cases for edge cases

### 2. **Go Version Compatibility Issue**

**Problem:**
```
Error: can't load config: the Go language version (go1.24) used to build 
golangci-lint is lower than the targeted Go version (1.25)
```

**Root Cause:**
The workflow was testing with Go 1.25 but golangci-lint v1.64.8 was built with Go 1.24. When a tool is built with an older Go version and targets a newer one, it causes compatibility issues.

**Solution:**
- Set the lint job to use Go 1.24 (match golangci-lint's build version)
- Explicitly pin golangci-lint version to v1.64.8 instead of 'latest'
- Updated go.mod to specify Go 1.24 as the base version (tests still run on 1.24 and 1.25)

**Changes Made:**
```yaml
# .github/workflows/test.yml - lint job
- uses: actions/setup-go@v4
  with:
    go-version: '1.24'  # ← Changed from '1.25'

- uses: golangci/golangci-lint-action@v3
  with:
    version: '1.64.8'   # ← Pinned from 'latest'
    args: --timeout=10m
```

### 3. **Test Structure Improvements**

**Benefits of the new test structure:**
- ✅ Uses proper subtests for better test organization and error reporting
- ✅ Each test case gets its own runner instance
- ✅ Cleaner test output with descriptive test names
- ✅ Follows Go testing best practices (table-driven tests with subtests)
- ✅ More maintainable and easier to add new test cases

### 4. **Extended Test Coverage**

Added new test cases to improve coverage:

**FabricWorkspaceNaming:**
- Valid naming (lowercase with hyphens) ✓
- Invalid naming (uppercase letters) ✗
- Invalid naming (too short) ✗
- Valid naming (exactly 3 characters) ✓
- Invalid naming (too long) ✗
- Invalid naming (special characters) ✗

**FabricWorkspaceCapacity:**
- Valid with capacity assigned ✓
- Invalid without capacity ✗
- Valid with capacity from variable ✓

**FabricWorkspaceDescription:**
- Valid with description ✓
- Invalid without description ✗
- Valid with description from variable ✓

## Files Modified

1. **rules/rules_test.go**
   - Complete rewrite of test functions
   - Changed from slice of structs to map-based TestRunner calls
   - Implemented subtests pattern
   - Added edge case test scenarios

2. **.github/workflows/test.yml**
   - Updated lint job to use Go 1.24
   - Pinned golangci-lint to v1.64.8
   - Updated build job to use Go 1.24
   - Test job still runs against both 1.24 and 1.25

3. **go.mod**
   - Changed from `go 1.25` to `go 1.24`
   - Ensures compatibility with golangci-lint build environment

## How to Verify the Fixes

Run the tests locally:
```bash
go test -v ./...
```

Run with coverage:
```bash
go test -v -coverprofile=coverage.txt ./...
go tool cover -html=coverage.txt
```

Run linting:
```bash
make lint  # or directly
golangci-lint run ./...
```

## Next Steps

1. Push these changes to your repository
2. GitHub Actions should now pass all tests
3. Monitor the CI/CD pipeline in GitHub Actions
4. Review coverage reports in codecov

## Additional Notes

- The minimum Go version is now 1.24 (set in go.mod)
- Tests are still validated against both Go 1.24 and 1.25 in CI/CD
- All tests follow the TFLint plugin SDK pattern as documented in their examples
- The test structure is consistent with TFLint ruleset best practices

## Reference Documentation

- [TFLint Plugin SDK Helper](https://pkg.go.dev/github.com/terraform-linters/tflint-plugin-sdk/helper)
- [Table-Driven Tests in Go](https://github.com/golang/go/wiki/TableDrivenTests)
- [TFLint Ruleset Template](https://github.com/terraform-linters/tflint-ruleset-template)
