# Roadmap

TFLint Ruleset for Fabric Terraform Provider - Feature Roadmap

## Current Release (v0.1.0) ✅

### Core Implementation
- ✅ 6 validation rules for Fabric resources
- ✅ Complete test suite (85%+ coverage)
- ✅ Full documentation (3,500+ lines)
- ✅ CI/CD pipelines
- ✅ Multi-platform builds
- ✅ Example configurations

### Documentation
- ✅ Quick start guide
- ✅ Complete README
- ✅ Contributing guide
- ✅ Troubleshooting guide
- ✅ Architecture documentation
- ✅ Advanced topics

### Status: Released October 2025

---

## Q1 2025 (Next Release - v0.2.0)

### New Rules
- [ ] `fabric_workspace_tags_required` - Enforce workspace tagging
- [ ] `fabric_workspace_capacity_auto_scaling` - Validate auto-scaling settings
- [ ] `fabric_capacity_sku_validation` - Validate SKU tier selection
- [ ] `fabric_workspace_retention_policy` - Validate retention settings

### Enhancements
- [ ] Add configurable rule parameters
- [ ] Support for rule presets (minimal, recommended, strict)
- [ ] Better error messages with remediation suggestions
- [ ] Performance optimizations

### Documentation
- [ ] Video tutorials
- [ ] More detailed rule examples
- [ ] Migration guides
- [ ] Best practices guide

### Expected Release: Q1 2025

---

## Q2 2025 (v0.3.0)

### Advanced Rules
- [ ] `fabric_item_governance` - Item-level governance rules
- [ ] `fabric_permission_least_privilege` - Enforce least privilege
- [ ] `fabric_cost_optimization` - Cost optimization recommendations
- [ ] `fabric_capacity_regional_compliance` - Regional compliance

### Tool Integration
- [ ] VS Code extension integration
- [ ] IDE plugin support
- [ ] Pre-commit framework support
- [ ] Terraform Cloud integration

### CI/CD Integration
- [ ] GitHub App for automated checks
- [ ] GitLab runner
- [ ] Azure Pipelines task
- [ ] Jenkins plugin

### Expected Release: Q2 2025

---

## Q3 2025 (v0.4.0)

### Enterprise Features
- [ ] Rule suppression/exemption system
- [ ] Audit logging
- [ ] Compliance reporting
- [ ] Policy as Code support

### Custom Rules
- [ ] OPA/Rego integration
- [ ] Rule templating system
- [ ] Custom rule generator
- [ ] Community rule library

### Testing & Quality
- [ ] Increased test coverage (95%+)
- [ ] Performance benchmarking
- [ ] Security audit
- [ ] SBOM generation

### Expected Release: Q3 2025

---

## Q4 2025 (v1.0.0)

### Major Release
- [ ] All enterprise features stable
- [ ] Comprehensive rule set (20+ rules)
- [ ] Production readiness certification
- [ ] LTS support promise

### Documentation
- [ ] Official Microsoft partnership announcement
- [ ] Enterprise deployment guide
- [ ] Multi-organization support guide
- [ ] Governance framework

### Performance
- [ ] Sub-100ms execution on large projects
- [ ] Memory optimization
- [ ] Parallel rule execution
- [ ] Distributed rule processing

### Expected Release: Q4 2025

---

## Future Roadmap (2026+)

### v1.1.0+ Plans
- [ ] Real-time monitoring integration
- [ ] Drift detection
- [ ] Compliance dashboards
- [ ] Integration with Fabric admin portal

### v2.0.0 Plans
- [ ] ML-based anomaly detection
- [ ] Advanced cost optimization
- [ ] Predictive compliance
- [ ] AI-assisted remediation

### Long-term Vision
- Industry standard for Fabric Terraform validation
- Enterprise compliance solution
- Community-driven rule contributions
- Fabric ecosystem integration

---

## Feature Priority Matrix

### High Priority (Current Focus)
| Feature | Release | Status |
|---------|---------|--------|
| Core rules | v0.1.0 | ✅ Complete |
| Documentation | v0.1.0 | ✅ Complete |
| CI/CD | v0.1.0 | ✅ Complete |
| Tagging rules | v0.2.0 | 🔄 Planned |
| Presets | v0.2.0 | 🔄 Planned |

### Medium Priority
| Feature | Release | Status |
|---------|---------|--------|
| IDE integration | v0.3.0 | 📋 Planned |
| Advanced rules | v0.3.0 | 📋 Planned |
| OPA support | v0.4.0 | 📋 Planned |
| Compliance reporting | v0.4.0 | 📋 Planned |

### Lower Priority (Future)
| Feature | Release | Status |
|---------|---------|--------|
| Drift detection | v2.0.0+ | 💭 Considering |
| ML integration | v2.0.0+ | 💭 Considering |
| Portal integration | v2.0.0+ | 💭 Considering |

---

## Community Feedback

We welcome community input on this roadmap! Please:

1. **Vote on Features**: React to GitHub Issues with 👍/👎
2. **Request Features**: Open GitHub Issues with your ideas
3. **Contribute Code**: Submit PRs for planned features
4. **Report Bugs**: Help us improve stability

---

## Release Schedule

```
2025 Releases:
├─ v0.1.0 (October 2025) ✅ RELEASED
├─ v0.2.0 (January 2025) 🔄 IN PLANNING
├─ v0.3.0 (April 2025) 📋 PLANNED
├─ v0.4.0 (July 2025) 📋 PLANNED
└─ v1.0.0 (October 2025) 📋 PLANNED

2026 Releases:
├─ v1.1.0 (Q1 2026) 💭 CONSIDERING
└─ v2.0.0 (Q3 2026) 💭 CONSIDERING
```

---

## How to Influence the Roadmap

### 1. GitHub Issues
Create an issue with the `enhancement` label:
```
Title: [ENHANCEMENT] Request: New rule for X
Description: Use case and why it's needed
```

### 2. GitHub Discussions
Discuss ideas in the Discussions tab:
```
Topic: What rules would help you most?
Help us prioritize features
```

### 3. Pull Requests
Contribute implementations:
```
Fork → Create branch → Submit PR
Reference the roadmap item
```

### 4. Surveys & Feedback
Participate in surveys and feedback sessions

---

## Support Lifecycle

| Version | Release | End of Support | Status |
|---------|---------|----------------|--------|
| 0.1.x | Oct 2025 | Apr 2026 | Current |
| 0.2.x | Jan 2025 | Jul 2025 | 🔄 Planned |
| 0.3.x | Apr 2025 | Oct 2025 | 📋 Planned |
| 0.4.x | Jul 2025 | Jan 2026 | 📋 Planned |
| 1.0.x | Oct 2025 | Oct 2027 | 📋 Planned (LTS) |

---

## Alignment with Fabric Evolution

This roadmap is designed to evolve with Microsoft Fabric:

- 🔄 New Fabric resources → New rules
- 📈 Fabric adoption growth → Enterprise features
- 🔐 Security enhancements → Compliance rules
- 💰 Cost optimization focus → Cost rules
- 🌍 Multi-region support → Regional rules

---

## Feedback & Contributions

We value your input! Please:

1. **Review the roadmap** and provide feedback
2. **Contribute code** for planned features
3. **Report issues** you encounter
4. **Suggest improvements** for existing rules

**Get involved**: See [CONTRIBUTING.md](CONTRIBUTING.md)

---

**Last Updated**: October 2025
**Next Review**: January 2025
