package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricDataflowInvalidDescription struct{ tflint.DefaultRule }

func NewFabricDataflowInvalidDescription() *FabricDataflowInvalidDescription {
	return &FabricDataflowInvalidDescription{}
}

func (r *FabricDataflowInvalidDescription) Name() string {
	return "fabric_dataflow_invalid_description"
}
func (r *FabricDataflowInvalidDescription) Enabled() bool             { return true }
func (r *FabricDataflowInvalidDescription) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricDataflowInvalidDescription) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/dataflow/definitions.json"
}

func (r *FabricDataflowInvalidDescription) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "description"},
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
		attr, ok := block.Body.Attributes["description"]
		if !ok {
			continue
		}

		var v string
		if err := runner.EvaluateExpr(attr.Expr, &v, nil); err != nil {
			continue
		}
		if len(v) > 3988 {
			if err := runner.EmitIssue(r,
				fmt.Sprintf("%s exceeds max length %d", "description", 3988),
				attr.Expr.Range()); err != nil {
				return err
			}
		}
		// TODO: add pattern/enum checks if needed
	}

	return nil
}
