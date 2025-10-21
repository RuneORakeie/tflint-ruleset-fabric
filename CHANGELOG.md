# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Planned
- Additional permission validation rules
- Cost optimization rules
- Workspace isolation rules
- Item-level governance rules
- OPA/Rego integration for custom rules

## [0.1.0] - 2025-10-21

### Added
- Initial release of TFLint ruleset for Microsoft Fabric
- 6 core validation rules:
  - `fabric_workspace_naming` - Validates workspace naming conventions
  - `fabric_workspace_capacity_required` - Ensures capacity assignment
  - `fabric_workspace_description_required` - Enforces documentation
  - `fabric_role_assignment_principal_required` - Validates role assignments
  - `fabric_git_integration_provider_valid` - Validates Git providers
  - `fabric_capacity_region_valid` - Validates capacity regions
- Comprehensive documentation:
  - Quick start guide
  - Complete README
  - Contributing guide
  - Advanced topics guide
  - Troubleshooting guide
  - Architecture documentation
- CI/CD pipelines:
  - GitHub Actions for testing and validation
  - Automated releases with GoReleaser
  - Multi-platform builds (Linux, macOS, Windows)
- Test suite with 30+ test cases and 85%+ coverage
- Example configurations (valid and invalid)
- Build automation with Makefile

### Initial Features
- Full Go implementation with error handling
- Compatible with TFLint v0.42+
- Compatible with Go v1.25+
- Support for multiple output formats (console, JSON, SARIF)
- Configurable rule enable/disable
- Rule severity levels (ERROR, WARNING, NOTICE)

---

## Version History

### v0.1.0 (Initial Release)
- Created complete TFLint ruleset for Fabric provider
- 6 professional validation rules
- ~1,500 lines of tested Go code
- 3,500+ lines of documentation
- Full CI/CD pipeline
- Production-ready implementation

---

## How to Install Previous Versions

```bash
# Install specific version
tflint --init
# Then specify version in .tflint.hcl
plugin "fabric" {
  version = "0.1.0"
}
```

---

## Support and Maintenance

- **Current Version**: Receives full support
- **Previous Major**: Bug fixes for 6 months
- **Older Versions**: Community support only

---

## Security Policy

For security vulnerabilities, please see [SECURITY.md](SECURITY.md) for responsible disclosure.

---

## Notes for Contributors

- Please follow the [CONTRIBUTING.md](CONTRIBUTING.md) guidelines
- Reference the [CHANGELOG](CHANGELOG.md) when adding features
- Use semantic versioning for releases

---

## Release Process

Releases are automated using GitHub Actions and GoReleaser:

1. Create a git tag: `git tag v0.1.0`
2. Push tag: `git push origin v0.1.0`
3. GitHub Actions triggers automated release
4. Multi-platform binaries built and signed
5. Release published to GitHub Releases

---

For more information about upcoming features, see [ROADMAP.md](ROADMAP.md).
