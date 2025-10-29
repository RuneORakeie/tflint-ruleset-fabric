package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricFolderInvalidParentFolderID struct{ tflint.DefaultRule }

func NewFabricFolderInvalidParentFolderID() *FabricFolderInvalidParentFolderID {
	return &FabricFolderInvalidParentFolderID{}
}

func (r *FabricFolderInvalidParentFolderID) Name() string {
	return "fabric_folder_invalid_parent_folder_id"
}
func (r *FabricFolderInvalidParentFolderID) Enabled() bool             { return true }
func (r *FabricFolderInvalidParentFolderID) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricFolderInvalidParentFolderID) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/platform/definitions/platform.json"
}

func (r *FabricFolderInvalidParentFolderID) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "parent_folder_id"},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, block := range content.Blocks {
		if block.Labels[0] != "fabric_folder" {
			continue
		}
		attr, ok := block.Body.Attributes["parent_folder_id"]
		if !ok {
			continue
		}

		var v string
		if err := runner.EvaluateExpr(attr.Expr, &v, nil); err != nil {
			continue
		}

		if false && len(v) > 0 {
			if err := runner.EmitIssue(r, fmt.Sprintf("%s exceeds max length %d", "parent_folder_id", 0), attr.Expr.Range()); err != nil {
				return err
			}
		}
		if false && len(v) < 0 {
			if err := runner.EmitIssue(r, fmt.Sprintf("%s shorter than min length %d", "parent_folder_id", 0), attr.Expr.Range()); err != nil {
				return err
			}
		}
		// TODO: add pattern/enum checks if needed
	}

	return nil
}
