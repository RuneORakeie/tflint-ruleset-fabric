package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricSparkEnvironmentSettingsInvalidDriverCores struct{ tflint.DefaultRule }

func NewFabricSparkEnvironmentSettingsInvalidDriverCores() *FabricSparkEnvironmentSettingsInvalidDriverCores {
	return &FabricSparkEnvironmentSettingsInvalidDriverCores{}
}

func (r *FabricSparkEnvironmentSettingsInvalidDriverCores) Name() string {
	return "fabric_spark_environment_settings_invalid_driver_cores"
}
func (r *FabricSparkEnvironmentSettingsInvalidDriverCores) Enabled() bool { return true }
func (r *FabricSparkEnvironmentSettingsInvalidDriverCores) Severity() tflint.Severity {
	return tflint.ERROR
}
func (r *FabricSparkEnvironmentSettingsInvalidDriverCores) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/environment/definitions.json"
}

func (r *FabricSparkEnvironmentSettingsInvalidDriverCores) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "driver_cores"},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, block := range content.Blocks {
		if block.Labels[0] != "fabric_spark_environment_settings" {
			continue
		}
		attr, ok := block.Body.Attributes["driver_cores"]
		if !ok {
			continue
		}

		var v string
		if err := runner.EvaluateExpr(attr.Expr, &v, nil); err != nil {
			continue
		}

		if false && len(v) > 0 {
			if err := runner.EmitIssue(r, fmt.Sprintf("%s exceeds max length %d", "driver_cores", 0), attr.Expr.Range()); err != nil {
				return err
			}
		}
		if false && len(v) < 0 {
			if err := runner.EmitIssue(r, fmt.Sprintf("%s shorter than min length %d", "driver_cores", 0), attr.Expr.Range()); err != nil {
				return err
			}
		}
		// TODO: add pattern/enum checks if needed
	}

	return nil
}
