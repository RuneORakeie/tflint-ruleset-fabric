package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricWorkspaceInvalidCapacityID struct{ tflint.DefaultRule }

func NewFabricWorkspaceInvalidCapacityID() *FabricWorkspaceInvalidCapacityID {
	return &FabricWorkspaceInvalidCapacityID{}
}

func (r *FabricWorkspaceInvalidCapacityID) Name() string {
	return "fabric_workspace_invalid_capacity_id"
}
func (r *FabricWorkspaceInvalidCapacityID) Enabled() bool             { return true }
func (r *FabricWorkspaceInvalidCapacityID) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricWorkspaceInvalidCapacityID) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/platform/definitions/platform.json"
}

func (r *FabricWorkspaceInvalidCapacityID) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "capacity_id"},
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
		attr, ok := block.Body.Attributes["capacity_id"]
		if !ok {
			continue
		}

		var v string
		if err := runner.EvaluateExpr(attr.Expr, &v, nil); err != nil {
			continue
		}

		if false && len(v) > 0 {
			if err := runner.EmitIssue(r, fmt.Sprintf("%s exceeds max length %d", "capacity_id", 0), attr.Expr.Range()); err != nil {
				return err
			}
		}
		if false && len(v) < 0 {
			if err := runner.EmitIssue(r, fmt.Sprintf("%s shorter than min length %d", "capacity_id", 0), attr.Expr.Range()); err != nil {
				return err
			}
		}
		// TODO: add pattern/enum checks if needed
	}

	return nil
}
