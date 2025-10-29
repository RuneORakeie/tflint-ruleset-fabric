package apispec

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricSparkEnvironmentSettingsInvalidDriverMemory struct{ tflint.DefaultRule }

func NewFabricSparkEnvironmentSettingsInvalidDriverMemory() *FabricSparkEnvironmentSettingsInvalidDriverMemory {
	return &FabricSparkEnvironmentSettingsInvalidDriverMemory{}
}

func (r *FabricSparkEnvironmentSettingsInvalidDriverMemory) Name() string {
	return "fabric_spark_environment_settings_invalid_driver_memory"
}
func (r *FabricSparkEnvironmentSettingsInvalidDriverMemory) Enabled() bool { return true }
func (r *FabricSparkEnvironmentSettingsInvalidDriverMemory) Severity() tflint.Severity {
	return tflint.ERROR
}
func (r *FabricSparkEnvironmentSettingsInvalidDriverMemory) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/environment/definitions.json"
}

func (r *FabricSparkEnvironmentSettingsInvalidDriverMemory) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "driver_memory"},
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
		attr, ok := block.Body.Attributes["driver_memory"]
		if !ok {
			continue
		}

		var v string
		if err := runner.EvaluateExpr(attr.Expr, &v, nil); err != nil {
			continue
		}
		// TODO: add pattern/enum checks if needed
	}

	return nil
}
