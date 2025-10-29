package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricDataflowInvalidDisplayName struct{ tflint.DefaultRule }

func NewFabricDataflowInvalidDisplayName() *FabricDataflowInvalidDisplayName {
	return &FabricDataflowInvalidDisplayName{}
}

func (r *FabricDataflowInvalidDisplayName) Name() string {
	return "fabric_dataflow_invalid_display_name"
}
func (r *FabricDataflowInvalidDisplayName) Enabled() bool             { return true }
func (r *FabricDataflowInvalidDisplayName) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricDataflowInvalidDisplayName) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/dataflow/definitions.json"
}

func (r *FabricDataflowInvalidDisplayName) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "display_name"},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, block := range content.Blocks {
		if block.Labels[0] != "fabric_dataflow" {
			continue
		}
		attr, ok := block.Body.Attributes["display_name"]
		if !ok {
			continue
		}

		var v string
		if err := runner.EvaluateExpr(attr.Expr, &v, nil); err != nil {
			continue
		}
		if len(v) > 256 {
			if err := runner.EmitIssue(r,
				fmt.Sprintf("%s exceeds max length %d", "display_name", 256),
				attr.Expr.Range()); err != nil {
				return err
			}
		}
		// TODO: add pattern/enum checks if needed
	}

	return nil
}
