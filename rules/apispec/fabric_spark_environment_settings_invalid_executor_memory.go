package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricSparkEnvironmentSettingsInvalidExecutorMemory checks whether fabric_spark_environment_settings.executor_memory is valid
type FabricSparkEnvironmentSettingsInvalidExecutorMemory struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewFabricRule returns a new rule instance
func NewFabricSparkEnvironmentSettingsInvalidExecutorMemory() *FabricSparkEnvironmentSettingsInvalidExecutorMemory {
	return &FabricSparkEnvironmentSettingsInvalidExecutorMemory{
		resourceType:  "fabric_spark_environment_settings",
		attributeName: "executor_memory",
		enum:          []string{"28g", "56g", "112g", "224g", "400g"},
	}
}

// Name returns the rule name
func (r *FabricSparkEnvironmentSettingsInvalidExecutorMemory) Name() string {
	return "fabric_spark_environment_settings_invalid_executor_memory"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricSparkEnvironmentSettingsInvalidExecutorMemory) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricSparkEnvironmentSettingsInvalidExecutorMemory) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricSparkEnvironmentSettingsInvalidExecutorMemory) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricSparkEnvironmentSettingsInvalidExecutorMemory) Check(runner tflint.Runner) error {
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

func (r *FabricSparkEnvironmentSettingsInvalidExecutorMemory) validateEnum(runner tflint.Runner, val string, attribute *hclext.Attribute) error {
	for _, valid := range r.enum {
		if val == valid {
			return nil
		}
	}
	return runner.EmitIssue(
		r,
		fmt.Sprintf("executor_memory must be one of: %v", r.enum),
		attribute.Expr.Range(),
	)
}
