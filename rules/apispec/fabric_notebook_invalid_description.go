package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricNotebookInvalidDescription struct{ tflint.DefaultRule }

func NewFabricNotebookInvalidDescription() *FabricNotebookInvalidDescription {
	return &FabricNotebookInvalidDescription{}
}

func (r *FabricNotebookInvalidDescription) Name() string {
	return "fabric_notebook_invalid_description"
}
func (r *FabricNotebookInvalidDescription) Enabled() bool             { return true }
func (r *FabricNotebookInvalidDescription) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricNotebookInvalidDescription) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/notebook/definitions.json"
}

func (r *FabricNotebookInvalidDescription) Check(runner tflint.Runner) error {
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
		if block.Labels[0] != "fabric_notebook" {
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
		if len(v) > 1021 {
			if err := runner.EmitIssue(r,
				fmt.Sprintf("%s exceeds max length %d", "description", 1021),
				attr.Expr.Range()); err != nil {
				return err
			}
		}
		// TODO: add pattern/enum checks if needed
	}

	return nil
}
