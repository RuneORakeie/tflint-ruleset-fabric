package apispec

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricConnectionInvalidConnectivityType struct{ tflint.DefaultRule }

func NewFabricConnectionInvalidConnectivityType() *FabricConnectionInvalidConnectivityType {
	return &FabricConnectionInvalidConnectivityType{}
}

func (r *FabricConnectionInvalidConnectivityType) Name() string {
	return "fabric_connection_invalid_connectivity_type"
}
func (r *FabricConnectionInvalidConnectivityType) Enabled() bool             { return true }
func (r *FabricConnectionInvalidConnectivityType) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricConnectionInvalidConnectivityType) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/platform/definitions/connections.json"
}

func (r *FabricConnectionInvalidConnectivityType) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "connectivity_type"},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, block := range content.Blocks {
		if block.Labels[0] != "fabric_connection" {
			continue
		}
		attr, ok := block.Body.Attributes["connectivity_type"]
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
