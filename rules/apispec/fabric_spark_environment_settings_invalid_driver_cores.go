package apispec

import (
	"fmt"
	"time"
	
	"github.com/google/uuid"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricSparkEnvironmentSettingsInvalidDriverCores checks whether fabric_spark_environment_settings.driver_cores is valid
type FabricSparkEnvironmentSettingsInvalidDriverCores struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string

	format        string






}

// NewFabricRule returns a new rule instance
func NewFabricSparkEnvironmentSettingsInvalidDriverCores() *FabricSparkEnvironmentSettingsInvalidDriverCores {
	return &FabricSparkEnvironmentSettingsInvalidDriverCores{
		resourceType:  "fabric_spark_environment_settings",
		attributeName: "driver_cores",
		enum:          []string{ "4", "8", "16", "32", "64",  },

		format:        "int32",






	}
}

// Name returns the rule name
func (r *FabricSparkEnvironmentSettingsInvalidDriverCores) Name() string {
	return "fabric_spark_environment_settings_invalid_driver_cores"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricSparkEnvironmentSettingsInvalidDriverCores) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricSparkEnvironmentSettingsInvalidDriverCores) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricSparkEnvironmentSettingsInvalidDriverCores) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricSparkEnvironmentSettingsInvalidDriverCores) Check(runner tflint.Runner) error {
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

func (r *FabricSparkEnvironmentSettingsInvalidDriverCores) validateEnum(runner tflint.Runner, val string, attribute *hclext.Attribute) error {
	for _, valid := range r.enum {
		if val == valid {
			return nil
		}
	}
	return runner.EmitIssue(
		r,
		fmt.Sprintf("driver_cores must be one of: %v", r.enum),
		attribute.Expr.Range(),
	)
}



func (r *FabricSparkEnvironmentSettingsInvalidDriverCores) validateFormat(runner tflint.Runner, val string, attribute *hclext.Attribute) error {
	switch r.format {
	case "uuid":
		if _, err := uuid.Parse(val); err != nil {
			return runner.EmitIssue(
				r,
				"driver_cores must be a valid UUID",
				attribute.Expr.Range(),
			)
		}
	case "date-time":
		if _, err := time.Parse(time.RFC3339, val); err != nil {
			return runner.EmitIssue(
				r,
				"driver_cores must be a valid RFC3339 date-time",
				attribute.Expr.Range(),
			)
		}
	}
	return nil
}


