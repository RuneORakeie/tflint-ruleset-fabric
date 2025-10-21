# Security Policy

## Reporting Security Vulnerabilities

If you discover a security vulnerability in the TFLint Fabric ruleset, please **do not** open a public GitHub issue. Instead, please report it responsibly by sending an email to: **security@example.com** with the following information:

1. **Description**: Detailed description of the vulnerability
2. **Impact**: Potential impact and severity
3. **Steps to Reproduce**: Clear steps to reproduce the issue
4. **Affected Versions**: Which versions are affected
5. **Suggested Fix**: Any suggested remediation (optional)

Please allow us 90 days to respond and issue a fix before public disclosure.

## Security Considerations

### Code Security

- ✅ No hardcoded credentials or secrets
- ✅ Input validation on all attributes
- ✅ No external API calls
- ✅ Limited dependencies (only TFLint SDK)
- ✅ Regular dependency updates
- ✅ GoSec linting enabled

### Release Security

- ✅ GPG signed releases
- ✅ SHA256 checksums provided
- ✅ SBOM generation (planned)
- ✅ Reproducible builds
- ✅ Multi-platform testing

### Plugin Security

- ✅ Runs locally only
- ✅ No data transmission
- ✅ No telemetry
- ✅ No tracking
- ✅ Open source and auditable

## Supported Versions

| Version | Status | Support Until |
|---------|--------|---------------|
| 0.1.x | Active | Current + 6 months |
| < 0.1 | Unsupported | Community only |

## Security Best Practices

### For Users

1. **Keep Updated**: Always use the latest version
   ```bash
   tflint --init  # Updates plugins
   ```

2. **Review Configurations**: Validate `.tflint.hcl` files
   ```bash
   hcl-inspect .tflint.hcl
   ```

3. **Use in CI/CD**: Enable automated checks
   ```yaml
   - run: tflint --recursive
   ```

4. **Monitor Dependencies**: Watch for Go module updates
   ```bash
   go mod tidy
   ```

5. **Scan Code**: Use Go security scanners
   ```bash
   gosec ./...
   ```

### For Developers

1. **Follow Best Practices**: Use secure coding patterns
2. **Input Validation**: Always validate Terraform attributes
3. **Error Handling**: Handle all error cases
4. **Dependencies**: Minimize external dependencies
5. **Code Review**: Request reviews before merging

## Vulnerability Management

### Response Timeline

1. **1-2 Days**: Initial acknowledgment
2. **7 Days**: Initial assessment
3. **30 Days**: Fix and testing
4. **45-90 Days**: Release and disclosure

### Disclosure Process

1. Vulnerability reported
2. Assessment and reproduction
3. Fix development and testing
4. Patch release
5. Responsible disclosure

## Compliance

This project aims to comply with:

- ✅ OWASP Top 10
- ✅ CWE/SANS Top 25
- ✅ Go Security Best Practices
- ✅ TFLint Plugin Security Guidelines

## Testing for Security

### Run Security Linter

```bash
# Install gosec
go install github.com/securego/gosec/v2/cmd/gosec@latest

# Run scan
gosec ./...
```

### Run All Checks

```bash
make lint
go test -race ./...
```

## Dependencies Security

### Current Dependencies

```
github.com/terraform-linters/tflint-plugin-sdk v0.20.1
```

### Updating Dependencies

```bash
go get -u
go mod tidy
make test
```

### Checking Vulnerabilities

```bash
go list -json -m all | nancy sleuth
```

## Public Disclosures

Any publicly disclosed vulnerabilities will be documented here:

- None reported yet

## Acknowledgments

We appreciate security researchers who responsibly disclose vulnerabilities to us. Thank you for helping keep TFLint Fabric ruleset secure.

## Additional Resources

- [Go Security Best Practices](https://golang.org/doc/effective_go#goroutines)
- [OWASP Security Guidelines](https://owasp.org/)
- [CWE Top 25](https://cwe.mitre.org/top25/)

## Questions?

For security questions or clarifications, please email: **security@example.com**

---

**Last Updated**: October 2025
**Version**: 0.1.0
