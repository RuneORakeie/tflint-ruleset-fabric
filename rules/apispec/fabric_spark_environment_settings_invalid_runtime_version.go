package apispec

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricSparkEnvironmentSettingsInvalidRuntimeVersion struct{ tflint.DefaultRule }

func NewFabricSparkEnvironmentSettingsInvalidRuntimeVersion() *FabricSparkEnvironmentSettingsInvalidRuntimeVersion {
	return &FabricSparkEnvironmentSettingsInvalidRuntimeVersion{}
}

func (r *FabricSparkEnvironmentSettingsInvalidRuntimeVersion) Name() string {
	return "fabric_spark_environment_settings_invalid_runtime_version"
}
func (r *FabricSparkEnvironmentSettingsInvalidRuntimeVersion) Enabled() bool { return true }
func (r *FabricSparkEnvironmentSettingsInvalidRuntimeVersion) Severity() tflint.Severity {
	return tflint.ERROR
}
func (r *FabricSparkEnvironmentSettingsInvalidRuntimeVersion) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/environment/definitions.json"
}

func (r *FabricSparkEnvironmentSettingsInvalidRuntimeVersion) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "runtime_version"},
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
		attr, ok := block.Body.Attributes["runtime_version"]
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
