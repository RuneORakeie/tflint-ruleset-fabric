package apispec

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricSparkEnvironmentSettingsInvalidExecutorCores checks whether fabric_spark_environment_settings.executor_cores is valid
type FabricSparkEnvironmentSettingsInvalidExecutorCores struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string

	format string
}

// NewFabricRule returns a new rule instance
func NewFabricSparkEnvironmentSettingsInvalidExecutorCores() *FabricSparkEnvironmentSettingsInvalidExecutorCores {
	return &FabricSparkEnvironmentSettingsInvalidExecutorCores{
		resourceType:  "fabric_spark_environment_settings",
		attributeName: "executor_cores",
		enum:          []string{"4", "8", "16", "32", "64"},

		format: "int32",
	}
}

// Name returns the rule name
func (r *FabricSparkEnvironmentSettingsInvalidExecutorCores) Name() string {
	return "fabric_spark_environment_settings_invalid_executor_cores"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricSparkEnvironmentSettingsInvalidExecutorCores) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricSparkEnvironmentSettingsInvalidExecutorCores) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricSparkEnvironmentSettingsInvalidExecutorCores) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricSparkEnvironmentSettingsInvalidExecutorCores) Check(runner tflint.Runner) error {
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

		if err := r.validateFormat(runner, val, attribute); err != nil {
			return err
		}

	}

	return nil
}

func (r *FabricSparkEnvironmentSettingsInvalidExecutorCores) validateEnum(runner tflint.Runner, val string, attribute *hclext.Attribute) error {
	for _, valid := range r.enum {
		if val == valid {
			return nil
		}
	}
	return runner.EmitIssue(
		r,
		fmt.Sprintf("executor_cores must be one of: %v", r.enum),
		attribute.Expr.Range(),
	)
}

func (r *FabricSparkEnvironmentSettingsInvalidExecutorCores) validateFormat(runner tflint.Runner, val string, attribute *hclext.Attribute) error {
	switch r.format {
	case "uuid":
		if _, err := uuid.Parse(val); err != nil {
			return runner.EmitIssue(
				r,
				"executor_cores must be a valid UUID",
				attribute.Expr.Range(),
			)
		}
	case "date-time":
		if _, err := time.Parse(time.RFC3339, val); err != nil {
			return runner.EmitIssue(
				r,
				"executor_cores must be a valid RFC3339 date-time",
				attribute.Expr.Range(),
			)
		}
	}
	return nil
}
