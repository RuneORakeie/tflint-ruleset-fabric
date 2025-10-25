# Roadmap

TFLint Ruleset for Fabric Terraform Provider - Development Roadmap

## Current Status (v0.1.0-dev)

### ✅ Completed Features

**70+ Validation Rules**
- ✅ 17 business logic rules for governance
- ✅ 53 auto-generated API spec rules
- ✅ Comprehensive test coverage (business logic rules)
- ✅ Full documentation for all rules

**Infrastructure**
- ✅ Rule generation framework
- ✅ Testing framework
- ✅ Documentation generator
- ✅ Project structure and tooling

**Documentation**
- ✅ Individual rule documentation (70+ pages)
- ✅ README with examples
- ✅ Contributing guidelines
- ✅ Code of conduct

### 🔄 In Progress
- Testing framework for API spec rules
- CI/CD pipeline setup
- Release automation

---

## Q1 2026 (v0.1.0 - First Official Release)

### Goals
- First stable release with 70+ rules
- Complete test coverage
- CI/CD automation
- Initial community release

### Planned Features
- ✅ All 70+ rules fully tested
- ✅ GitHub Actions CI/CD
- ✅ Automated releases via GoReleaser
- ✅ GitHub release artifacts
- ✅ Installation via TFLint plugin system

### Testing & Quality
- Complete test coverage for all rules
- Integration test suite
- Example configurations
- Documentation review

**Target Release**: January 2026

---

## Q2 2026 (v0.2.0)

### New Rules
- [ ] `fabric_workspace_tags_required` - Enforce workspace tagging
- [ ] `fabric_naming_convention` - Configurable naming patterns
- [ ] `fabric_lakehouse_shortcut_validation` - Shortcut configuration rules
- [ ] `fabric_semantic_model_validation` - Semantic model rules

### Enhancements
- [ ] Configurable rule parameters
- [ ] Rule presets (minimal, recommended, strict)
- [ ] Better error messages with fix suggestions
- [ ] Performance optimizations

### Documentation
- [ ] Video tutorials
- [ ] Migration guides
- [ ] Best practices guide
- [ ] Rule cookbook

**Target Release**: April 2026

---

## Q3 2026 (v0.3.0)

### Advanced Rules
- [ ] `fabric_capacity_optimization` - Cost optimization recommendations
- [ ] `fabric_permission_least_privilege` - Permission validation
- [ ] `fabric_data_governance` - Data classification rules
- [ ] `fabric_compliance_validation` - Compliance checks

### Tool Integration
- [ ] VS Code extension
- [ ] Pre-commit hooks
- [ ] Terraform Cloud integration
- [ ] Azure DevOps pipeline task

### Features
- [ ] Custom rule templates
- [ ] Rule suppression system
- [ ] Configuration validation
- [ ] Bulk rule enable/disable

**Target Release**: July 2026

---

## Q4 2026 (v0.4.0)

### Enterprise Features
- [ ] Audit logging
- [ ] Compliance reporting
- [ ] Policy as Code support
- [ ] Multi-workspace validation

### Advanced Validation
- [ ] Cross-resource dependency checks
- [ ] Capacity utilization warnings
- [ ] Cost estimation
- [ ] Security scanning

### Community
- [ ] Community rule library
- [ ] Rule contribution framework
- [ ] Public rule registry
- [ ] Community voting on features

**Target Release**: October 2026

---

## 2027+ (Long-term Vision)

### v1.0.0 - Enterprise Ready
- [ ] 100+ validation rules
- [ ] LTS support
- [ ] Enterprise SLA
- [ ] Professional support options
- [ ] Certification program

### Advanced Features
- [ ] Real-time monitoring integration
- [ ] Drift detection
- [ ] Automatic remediation suggestions
- [ ] Integration with Fabric admin portal

### Innovation
- [ ] ML-based anomaly detection
- [ ] Predictive compliance
- [ ] AI-assisted rule creation
- [ ] Natural language rule queries

---

## Feature Requests & Priorities

### High Priority
| Feature | Target Release | Status |
|---------|---------------|--------|
| Complete test coverage | v0.1.0 | 🔄 In Progress |
| CI/CD pipeline | v0.1.0 | 🔄 In Progress |
| Rule presets | v0.2.0 | 📋 Planned |
| Configurable parameters | v0.2.0 | 📋 Planned |

### Medium Priority
| Feature | Target Release | Status |
|---------|---------------|--------|
| VS Code extension | v0.3.0 | 📋 Planned |
| Custom rules | v0.3.0 | 📋 Planned |
| Compliance reports | v0.4.0 | 📋 Planned |
| Policy as Code | v0.4.0 | 📋 Planned |

### Community Requested
| Feature | Votes | Status |
|---------|-------|--------|
| TBD | TBD | 💭 Gathering feedback |

---

## Contributing to the Roadmap

We welcome community input! You can influence the roadmap by:

### 1. **GitHub Issues**
Create feature requests with the `enhancement` label:
```
Title: [FEATURE] Request for <feature>
Description: 
- Use case
- Expected behavior
- Benefits
```

### 2. **GitHub Discussions**
Participate in roadmap discussions:
- Vote on features using 👍 reactions
- Share your use cases
- Propose new rules

### 3. **Pull Requests**
Contribute implementations:
- Fork the repository
- Implement your feature
- Submit PR with tests and docs

### 4. **Community Meetings**
Join quarterly roadmap review sessions (TBD)

---

## Release Schedule

```
2026 Releases:
├─ v0.1.0 (January 2026)   - First stable release
├─ v0.2.0 (April 2026)      - Enhanced rules & config
├─ v0.3.0 (July 2026)       - Advanced features
└─ v0.4.0 (October 2026)    - Enterprise features

2027 Releases:
├─ v0.5.0 (Q1 2027)         - TBD based on feedback
├─ v0.6.0 (Q2 2027)         - TBD
└─ v1.0.0 (Q4 2027)         - LTS Release
```

---

## Support Lifecycle

| Version | Release Date | End of Support | Status |
|---------|-------------|----------------|--------|
| 0.1.x | Jan 2026 | Jul 2026 | 📋 Planned |
| 0.2.x | Apr 2026 | Oct 2026 | 📋 Planned |
| 0.3.x | Jul 2026 | Jan 2027 | 📋 Planned |
| 0.4.x | Oct 2026 | Apr 2027 | 📋 Planned |
| 1.0.x | Q4 2027 | Q4 2029 | 📋 Planned (LTS) |

---

## Alignment with Fabric Evolution

This roadmap evolves with Microsoft Fabric:

- 🆕 **New Fabric resources** → New validation rules
- 📈 **Fabric adoption growth** → Enterprise features
- 🔐 **Security enhancements** → Security rules
- 💰 **Cost optimization** → Cost validation rules
- 🌍 **Multi-region expansion** → Regional compliance

We monitor:
- [Fabric Terraform Provider updates](https://github.com/microsoft/terraform-provider-fabric)
- [Fabric release notes](https://learn.microsoft.com/en-us/fabric/release-plan/)
- Community feedback

---

## Metrics & Goals

### v0.1.0 Goals
- ✅ 70+ rules implemented
- 🔄 90%+ test coverage
- 🔄 100% documented rules
- 📋 10+ community stars
- 📋 5+ contributors

### v0.2.0 Goals
- 📋 100+ rules
- 📋 95%+ test coverage
- 📋 50+ community stars
- 📋 Rule presets
- 📋 Configurable parameters

### v1.0.0 Goals
- 📋 150+ rules
- 📋 99%+ test coverage
- 📋 500+ community stars
- 📋 50+ contributors
- 📋 Enterprise adoption

---

## Get Involved

We need your help to make this the best Fabric validation tool!

**Ways to contribute:**
- 🐛 Report bugs
- 💡 Suggest features
- 📝 Improve documentation
- 🧪 Add test cases
- 🔧 Implement rules
- 📣 Spread the word

**See**: [CONTRIBUTING.md](CONTRIBUTING.md)

---

**Last Updated**: January 2026  
**Next Review**: April 2026

**Status Legend:**
- ✅ Complete
- 🔄 In Progress
- 📋 Planned
- 💭 Under Consideration
