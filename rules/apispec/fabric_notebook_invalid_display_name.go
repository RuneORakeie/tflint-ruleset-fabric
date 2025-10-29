package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricNotebookInvalidDisplayName struct{ tflint.DefaultRule }

func NewFabricNotebookInvalidDisplayName() *FabricNotebookInvalidDisplayName {
	return &FabricNotebookInvalidDisplayName{}
}

func (r *FabricNotebookInvalidDisplayName) Name() string {
	return "fabric_notebook_invalid_display_name"
}
func (r *FabricNotebookInvalidDisplayName) Enabled() bool             { return true }
func (r *FabricNotebookInvalidDisplayName) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricNotebookInvalidDisplayName) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/notebook/definitions.json"
}

func (r *FabricNotebookInvalidDisplayName) Check(runner tflint.Runner) error {
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
		if block.Labels[0] != "fabric_notebook" {
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

		if true && len(v) > 256 {
			if err := runner.EmitIssue(r, fmt.Sprintf("%s exceeds max length %d", "display_name", 256), attr.Expr.Range()); err != nil {
				return err
			}
		}
		if false && len(v) < 0 {
			if err := runner.EmitIssue(r, fmt.Sprintf("%s shorter than min length %d", "display_name", 0), attr.Expr.Range()); err != nil {
				return err
			}
		}
		// TODO: add pattern/enum checks if needed
	}

	return nil
}
