package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricSparkEnvironmentSettingsInvalidRuntimeVersion checks whether fabric_spark_environment_settings.runtime_version is valid
type FabricSparkEnvironmentSettingsInvalidRuntimeVersion struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewFabricRule returns a new rule instance
func NewFabricSparkEnvironmentSettingsInvalidRuntimeVersion() *FabricSparkEnvironmentSettingsInvalidRuntimeVersion {
	return &FabricSparkEnvironmentSettingsInvalidRuntimeVersion{
		resourceType:  "fabric_spark_environment_settings",
		attributeName: "runtime_version",
		enum:          []string{"1.1", "1.2", "1.3"},
	}
}

// Name returns the rule name
func (r *FabricSparkEnvironmentSettingsInvalidRuntimeVersion) Name() string {
	return "fabric_spark_environment_settings_invalid_runtime_version"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricSparkEnvironmentSettingsInvalidRuntimeVersion) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricSparkEnvironmentSettingsInvalidRuntimeVersion) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricSparkEnvironmentSettingsInvalidRuntimeVersion) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricSparkEnvironmentSettingsInvalidRuntimeVersion) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)
		if err != nil {
			return err
		}

		if err := r.validateEnum(runner, val, attribute); err != nil {
			return err
		}

	}

	return nil
}

func (r *FabricSparkEnvironmentSettingsInvalidRuntimeVersion) validateEnum(runner tflint.Runner, val string, attribute *hclext.Attribute) error {
	for _, valid := range r.enum {
		if val == valid {
			return nil
		}
	}
	return runner.EmitIssue(
		r,
		fmt.Sprintf("runtime_version must be one of: %v", r.enum),
		attribute.Expr.Range(),
	)
}
