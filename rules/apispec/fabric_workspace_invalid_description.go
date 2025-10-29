package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricWorkspaceInvalidDescription struct{ tflint.DefaultRule }

func NewFabricWorkspaceInvalidDescription() *FabricWorkspaceInvalidDescription {
	return &FabricWorkspaceInvalidDescription{}
}

func (r *FabricWorkspaceInvalidDescription) Name() string {
	return "fabric_workspace_invalid_description"
}
func (r *FabricWorkspaceInvalidDescription) Enabled() bool             { return true }
func (r *FabricWorkspaceInvalidDescription) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricWorkspaceInvalidDescription) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/platform/definitions/platform.json"
}

func (r *FabricWorkspaceInvalidDescription) Check(runner tflint.Runner) error {
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
		if block.Labels[0] != "fabric_workspace" {
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
		if len(v) > 4000 {
			if err := runner.EmitIssue(r,
				fmt.Sprintf("%s exceeds max length %d", "description", 4000),
				attr.Expr.Range()); err != nil {
				return err
			}
		}
		// TODO: add pattern/enum checks if needed
	}

	return nil
}
